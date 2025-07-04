package generate

import (
	"context"
	"encoding/json"
	"fmt"
	"server/app/admin/api/common/page"
	v1 "server/app/admin/api/generate/v1"
	"server/app/admin/internal/consts"
	"server/app/admin/internal/dao"
	generateLibrary "server/app/admin/internal/library/generate"
	"server/app/admin/internal/model/do"
	"server/utility"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sGenerate struct{}

func New() *sGenerate {
	return &sGenerate{}
}

// GetTables 获取数据库中所有表信息
func (s *sGenerate) GetTables(ctx context.Context, req v1.GetTablesReq) (res *v1.GetTablesRes, err error) {
	// 调用库函数获取表元数据
	result, err := generateLibrary.GetTablesMeta("mysql", req.CurrentPage, req.PageSize)
	if err != nil {
		return nil, err
	}
	res = &v1.GetTablesRes{
		List: make([]*v1.TableInfo, 0, len(result.List)),
		ResPage: page.ResPage{
			Total:       result.Total,
			CurrentPage: req.CurrentPage,
		},
	}

	for _, table := range result.List {
		res.List = append(res.List, &v1.TableInfo{
			TableName:    table.Name,
			TableComment: table.Comment,
			CreateTime:   utility.ParseAndFormatTime(table.CreateTime),
			UpdateTime:   utility.ParseAndFormatTime(table.UpdateTime),
		})
	}

	return
}

// ImportTables 导入表信息到代码生成记录
func (s *sGenerate) ImportTables(ctx context.Context, req v1.ImportTablesReq) (res *v1.ImportTablesRes, err error) {
	if len(req.Tables) == 0 {
		return &v1.ImportTablesRes{
			Count:   0,
			Success: false,
			Errors:  []string{"没有提供要导入的表信息"},
		}, nil
	}

	res = &v1.ImportTablesRes{
		Tables: make([]string, 0),
		Errors: make([]string, 0),
	}

	// 使用事务处理批量导入
	err = dao.CodeGenRecord.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, table := range req.Tables {
			// 验证表名
			if strings.TrimSpace(table.TableName) == "" {
				res.Errors = append(res.Errors, "表名不能为空")
				continue
			}

			// 检查表是否已存在于代码生成记录中
			existCount, err := dao.CodeGenRecord.Ctx(ctx).Where(dao.CodeGenRecord.Columns().TableName, table.TableName).Count()
			if err != nil {
				res.Errors = append(res.Errors, fmt.Sprintf("检查表 %s 时出错: %v", table.TableName, err))
				continue
			}

			if existCount > 0 {
				res.Errors = append(res.Errors, fmt.Sprintf("表 %s 已存在于代码生成记录中", table.TableName))
				continue
			}

			// 获取表的详细信息（包括字段信息）
			tableInfo, err := s.getTableDetailInfo(table.TableName)
			if err != nil {
				res.Errors = append(res.Errors, fmt.Sprintf("获取表 %s 详细信息失败: %v", table.TableName, err))
				continue
			}

			// 生成包名和模块名
			packageName := s.generatePackageName(table.TableName)
			moduleName := s.generateModuleName(table.TableName)

			// 准备插入数据
			record := &do.CodeGenRecord{
				TableName:    table.TableName,
				TableComment: table.TableComment,
				PackageName:  packageName,
				ModuleName:   moduleName,
				Options: fmt.Sprintf(`{
					"tableName": "%s",
					"tableComment": "%s",
					"moduleName": "%s",
					"hasAdd": true,
					"hasDelete": true,
					"hasEdit": true,
					"hasQuery": true,
					"hasExport": false,
					"hasBatchDelete": false,
					"parentMenuId": null,
					"menuName": "",
					"menuIcon": "",
					"joinTables": []
				}`, table.TableName, table.TableComment, moduleName), // 默认配置
				Columns: tableInfo.ColumnsJson,
			}

			// 插入记录
			_, err = dao.CodeGenRecord.Ctx(ctx).TX(tx).Data(record).Insert()
			if err != nil {
				res.Errors = append(res.Errors, fmt.Sprintf("插入表 %s 记录失败: %v", table.TableName, err))
				continue
			}

			// 成功导入
			res.Tables = append(res.Tables, table.TableName)
			res.Count++

			g.Log().Infof(ctx, "成功导入表 %s 到代码生成记录", table.TableName)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// 设置成功状态
	res.Success = res.Count > 0

	g.Log().Infof(ctx, "导入表信息完成，成功: %d, 失败: %d", res.Count, len(res.Errors))

	return res, nil
}

// getTableDetailInfo 获取表的详细信息，包括字段信息
func (s *sGenerate) getTableDetailInfo(tableName string) (*TableDetailInfo, error) {
	// 这里可以调用现有的库函数获取表的字段信息
	// 假设有一个获取表字段的方法
	// 获取表字段信息
	columns, err := generateLibrary.GetTableColumns(tableName)
	if err != nil {
		return nil, err
	}

	// 转换为前端需要的字段格式
	var fieldList []map[string]interface{}
	for i, col := range columns {
		// 判断字段类型对应的表单组件
		htmlType := "input"
		if strings.Contains(strings.ToLower(col.DataType), "text") {
			htmlType = "textarea"
		} else if strings.Contains(strings.ToLower(col.DataType), "int") {
			htmlType = "number"
		} else if strings.Contains(strings.ToLower(col.DataType), "date") {
			htmlType = "date"
		}

		// 判断是否为主键
		isPk := col.ColumnKey == "PRI"
		// 判断是否自增
		isIncrement := strings.Contains(col.Extra, "auto_increment")
		// 判断是否必填
		isRequired := false

		fieldInfo := map[string]interface{}{
			consts.FieldKeyID:              i + 1,
			consts.FieldKeyFieldName:       col.ColumnName,        // 字段名
			consts.FieldKeyFieldComment:    col.ColumnComment,     // 字段描述
			consts.FieldKeyHtmlType:        htmlType,              // 表单组件
			consts.FieldKeyDictType:        consts.DefaultDictType,        // 绑定的字典
			consts.FieldKeyValidationRules: consts.DefaultValidationRules, // 验证规则
			consts.FieldKeyIsEdit:          !isPk && !isIncrement, // 编辑
			consts.FieldKeyIsRequired:      isRequired,            // 必填
			consts.FieldKeyIsUnique:        consts.DefaultIsUnique,        // 唯一
			consts.FieldKeyIsList:          consts.DefaultIsList,          // 列表
			consts.FieldKeyIsQuery:         consts.DefaultIsQuery,         // 查询
			consts.FieldKeyQueryType:       consts.DefaultQueryType,       // 查询方式
		}
		fieldList = append(fieldList, fieldInfo)
	}

	// 将字段信息转换为JSON
	columnsJson, err := json.Marshal(fieldList)
	if err != nil {
		return nil, fmt.Errorf("序列化字段信息失败: %v", err)
	}

	return &TableDetailInfo{
		ColumnsJson: string(columnsJson),
	}, nil
}

// generatePackageName 根据表名生成包名
func (s *sGenerate) generatePackageName(tableName string) string {
	// 移除表前缀，转换为小写
	name := strings.ToLower(tableName)
	// 移除常见的表前缀
	prefixes := []string{"t_", "tb_", "tbl_"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(name, prefix) {
			name = strings.TrimPrefix(name, prefix)
			break
		}
	}
	return name
}

// generateModuleName 根据表名生成模块名
func (s *sGenerate) generateModuleName(tableName string) string {
	// 简单的模块名生成逻辑
	name := s.generatePackageName(tableName)
	// 可以根据业务需求进行更复杂的模块名生成
	return name
}

// TableDetailInfo 表详细信息
type TableDetailInfo struct {
	ColumnsJson string
}

func (s *sGenerate) GetTablesWithColumns(ctx context.Context, req v1.GetTablesWithColumnsReq) (res *v1.GetTablesWithColumnsRes, err error) {
	// 调用库函数获取表和字段信息
	tablesWithColumns, err := generateLibrary.GetAllTablesWithFilteredColumns()
	if err != nil {
		return nil, err
	}

	// 转换为API响应格式
	result := make([]v1.TableWithColumnsInfo, 0, len(tablesWithColumns))
	for _, table := range tablesWithColumns {
		// 转换字段信息
		columns := make([]v1.ColumnInfo, 0, len(table.Columns))
		for _, col := range table.Columns {
			columns = append(columns, v1.ColumnInfo{
				ColumnName:    col.ColumnName,
				DataType:      col.DataType,
				ColumnType:    col.ColumnType,
				IsNullable:    col.IsNullable,
				ColumnDefault: col.ColumnDefault,
				ColumnComment: col.ColumnComment,
				ColumnKey:     col.ColumnKey,
				Extra:         col.Extra,
			})
		}

		result = append(result, v1.TableWithColumnsInfo{
			TableName:    table.TableName,
			TableComment: table.TableComment,
			Columns:      columns,
		})
	}

	res = &v1.GetTablesWithColumnsRes{
		List: result,
	}

	return
}

// GetTableColumns 根据表名获取字段信息
func (s *sGenerate) GetTableColumns(ctx context.Context, req v1.GetTableColumnsReq) (res *v1.GetTableColumnsRes, err error) {
	// 调用库函数获取指定表的字段信息
	columns, err := generateLibrary.GetTableColumns(req.TableName)
	if err != nil {
		return nil, fmt.Errorf("获取表 %s 的字段信息失败: %v", req.TableName, err)
	}

	// 转换为API响应格式
	resColumns := make([]v1.ColumnInfo, 0, len(columns))
	for _, col := range columns {
		resColumns = append(resColumns, v1.ColumnInfo{
			ColumnName:    col.ColumnName,
			DataType:      col.DataType,
			ColumnType:    col.ColumnType,
			IsNullable:    col.IsNullable,
			ColumnDefault: col.ColumnDefault,
			ColumnComment: col.ColumnComment,
			ColumnKey:     col.ColumnKey,
			Extra:         col.Extra,
		})
	}

	res = &v1.GetTableColumnsRes{
		TableName: req.TableName,
		Columns:   resColumns,
	}

	return res, nil
}
