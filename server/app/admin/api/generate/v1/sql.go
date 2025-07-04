package v1

import "github.com/gogf/gf/v2/frame/g"

// GenerateSqlReq 根据提示词生成SQL语句请求参数
type GenerateSqlReq struct {
	g.Meta `path:"/sql/generate" method:"post" tags:"Generate" summary:"根据提示词生成SQL语句"`
	Prompt string `json:"prompt" v:"required#请输入提示词" dc:"提示词描述"`
}

// GenerateSqlRes 根据提示词生成SQL语句响应参数
type GenerateSqlRes struct {
	Sql string `json:"sql" dc:"生成的SQL语句"`
}

// ExecuteSqlReq 执行SQL语句请求参数
type ExecuteSqlReq struct {
	g.Meta `path:"/sql/execute" method:"post" tags:"Generate" summary:"执行SQL语句"`
	Sql    string `json:"sql" v:"required#请输入SQL语句" dc:"要执行的SQL语句"`
}

// ExecuteSqlRes 执行SQL语句响应参数
type ExecuteSqlRes struct {
}
