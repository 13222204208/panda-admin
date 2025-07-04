package generate

import (
	"context"
	"fmt"
	"strings"
	"time"

	v1 "server/app/admin/api/generate/v1"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/guonaihong/gout"
)

// GenerateSqlRequest 外部API请求结构
type GenerateSqlRequest struct {
	Prompt string `json:"prompt"`
}

// GenerateSqlResponse 外部API响应结构
type GenerateSqlResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Data struct {
		Sql string `json:"sql"`
	} `json:"data"`
}

// GenerateSql 根据提示词生成SQL语句
func (s *sGenerate) GenerateSql(ctx context.Context, req v1.GenerateSqlReq) (res *v1.GenerateSqlRes, err error) {
	// 获取配置
	var (
		apiURL     = g.Cfg().MustGet(ctx, "sql.generate.url", "http://deepseek.jingqi888.cn/generate/sql").String()
		timeout    = g.Cfg().MustGet(ctx, "sql.generate.timeout", 60000).Int() // 默认60秒
		maxRetries = g.Cfg().MustGet(ctx, "sql.generate.retries", 3).Int()     // 默认重试3次
	)

	// 记录请求信息
	g.Log().Info(ctx, "SQL生成API请求开始", g.Map{
		"url":     apiURL,
		"prompt":  req.Prompt,
		"timeout": timeout,
	})

	// 构建请求数据
	reqData := GenerateSqlRequest{
		Prompt: req.Prompt,
	}

	// 重试机制
	for attempt := 1; attempt <= maxRetries; attempt++ {
		g.Log().Info(ctx, fmt.Sprintf("尝试第%d次请求", attempt))

		// 发送HTTP请求
		var responseBody string
		var code int
		err = gout.POST(apiURL).
			SetJSON(reqData).
			SetHeader(gout.H{
				"Content-Type": "application/json",
				"Accept":       "application/json",
				"User-Agent":   "Panda-SQL-Generator/1.0",
			}).
			SetTimeout(time.Duration(timeout) * time.Millisecond).
			BindBody(&responseBody).
			Code(&code).
			Do()

		// 如果请求成功，跳出重试循环
		if err == nil && code == 200 {
			// 记录成功响应
			g.Log().Info(ctx, "API响应成功", g.Map{
				"attempt":  attempt,
				"response": responseBody,
			})

			// 检查响应是否为空
			if responseBody == "" {
				return nil, fmt.Errorf("API返回空响应")
			}

			// 使用GoFrame的JSON解析
			var response GenerateSqlResponse
			parseErr := gjson.DecodeTo(responseBody, &response)
			if parseErr != nil {
				return nil, fmt.Errorf("JSON解析失败: %v，原始响应: %s", parseErr, responseBody)
			}

			if response.Code != 0 {
				return nil, fmt.Errorf("SQL生成API返回错误，错误码: %d，错误信息: %s", response.Code, response.Msg)
			}

			// 清理SQL内容中的markdown标记
			sqlContent := cleanSqlContent(response.Data.Sql)

			// 返回成功结果
			res = &v1.GenerateSqlRes{
				Sql: sqlContent,
			}
			return res, nil
		}

		// 记录失败信息
		g.Log().Warning(ctx, fmt.Sprintf("第%d次请求失败", attempt), g.Map{
			"error": err,
			"code":  code,
			"body":  responseBody,
		})

		// 如果不是最后一次尝试，等待后重试
		if attempt < maxRetries {
			waitTime := time.Duration(attempt*2) * time.Second // 递增等待时间
			g.Log().Info(ctx, fmt.Sprintf("等待%v后重试", waitTime))
			time.Sleep(waitTime)
		}
	}

	// 所有重试都失败
	if err != nil {
		return nil, fmt.Errorf("SQL生成API请求失败(已重试%d次): %v", maxRetries, err)
	}

	return nil, fmt.Errorf("SQL生成API请求失败，(已重试%d次)", maxRetries)
}

// cleanSqlContent 清理SQL内容中的markdown标记
func cleanSqlContent(sql string) string {
	// 移除可能的markdown代码块标记
	sql = strings.ReplaceAll(sql, "```sql", "")
	sql = strings.ReplaceAll(sql, "```", "")

	// 移除首尾空白字符
	sql = strings.TrimSpace(sql)

	return sql
}

// ExecuteSql 执行SQL语句（仅允许创建表语句）
func (s *sGenerate) ExecuteSql(ctx context.Context, req v1.ExecuteSqlReq) (res *v1.ExecuteSqlRes, err error) {
	// 记录请求信息
	g.Log().Info(ctx, "SQL执行请求开始", g.Map{
		"sql": req.Sql,
	})

	// 清理SQL语句
	sqlStatement := strings.TrimSpace(req.Sql)
	if sqlStatement == "" {
		return res, gerror.New("sql语句不能为空")
	}

	// 转换为大写进行检查
	sqlUpper := strings.ToUpper(sqlStatement)

	// 安全检查：只允许CREATE TABLE语句
	if !strings.HasPrefix(sqlUpper, "CREATE TABLE") {
		g.Log().Warning(ctx, "SQL执行被拒绝：不是CREATE TABLE语句", g.Map{
			"sql": sqlStatement,
		})
		return res, gerror.New("为了安全考虑，只允许执行CREATE TABLE语句")
	}

	// 额外安全检查：禁止包含危险关键词
	dangerousKeywords := []string{"DROP", "DELETE", "TRUNCATE", "ALTER", "INSERT"}
	for _, keyword := range dangerousKeywords {
		if strings.Contains(sqlUpper, keyword) {
			g.Log().Warning(ctx, "SQL执行被拒绝：包含危险关键词", g.Map{
				"sql":     sqlStatement,
				"keyword": keyword,
			})
			return res, gerror.New(fmt.Sprintf("CREATE TABLE语句中不能包含 %s 关键词", keyword))
		}
	}

	// 获取数据库连接
	db := g.DB()

	// 执行CREATE TABLE语句
	result, err := db.Exec(ctx, sqlStatement)
	if err != nil {
		g.Log().Error(ctx, "CREATE TABLE执行失败", g.Map{
			"sql":   sqlStatement,
			"error": err,
		})
		return res, gerror.New(fmt.Sprintf("CREATE TABLE执行失败: %v", err))
	}

	// 获取影响行数
	affectedRows, err := result.RowsAffected()

	g.Log().Info(ctx, "CREATE TABLE执行成功", g.Map{
		"sql":          sqlStatement,
		"affectedRows": affectedRows,
	})

	return res, nil
}
