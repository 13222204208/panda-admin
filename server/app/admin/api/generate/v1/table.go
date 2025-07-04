package v1

import (
	"server/app/admin/api/common/page"

	"github.com/gogf/gf/v2/frame/g"
)

type TableInfo struct {
	TableName    string `json:"tableName" dc:"表名"`
	TableComment string `json:"tableComment" dc:"表描述"`
	CreateTime   string `json:"createTime" dc:"创建时间"`
	UpdateTime   string `json:"updateTime" dc:"修改时间"`
}

type GetTablesReq struct {
	g.Meta `path:"/table" method:"get" tags:"Generate" summary:"获取数据表列表"`
	page.ReqPage
}

type GetTablesRes struct {
	List []*TableInfo `json:"list" dc:"数据表列表"`
	page.ResPage
}

// ImportTablesReq 导入表结构请求参数
type ImportTablesReq struct {
	g.Meta `path:"/table/import" method:"post" tags:"Generate" summary:"导入表数据"`
	Tables []TableInfo `json:"tables" v:"required#请输入表名" dc:"要导入的表信息列表"`
}

// ImportTablesRes 导入表结构响应参数
type ImportTablesRes struct {
	Count   int      `json:"count" dc:"成功导入的表数量"`
	Tables  []string `json:"tables,omitempty" dc:"导入的表名列表"`
	Errors  []string `json:"errors,omitempty" dc:"导入失败的错误信息"`
	Success bool     `json:"success" dc:"导入是否成功"`
}

// ColumnInfo 字段信息
type ColumnInfo struct {
	ColumnName    string `json:"columnName" dc:"字段名"`
	DataType      string `json:"dataType" dc:"数据类型"`
	ColumnType    string `json:"columnType" dc:"完整字段类型"`
	IsNullable    string `json:"isNullable" dc:"是否可为空"`
	ColumnDefault string `json:"columnDefault" dc:"默认值"`
	ColumnComment string `json:"columnComment" dc:"字段注释"`
	ColumnKey     string `json:"columnKey" dc:"键类型"`
	Extra         string `json:"extra" dc:"额外信息"`
}

// TableWithColumnsInfo 表及其字段信息
type TableWithColumnsInfo struct {
	TableName    string       `json:"tableName" dc:"表名"`
	TableComment string       `json:"tableComment" dc:"表描述"`
	Columns      []ColumnInfo `json:"columns" dc:"字段列表"`
}

// GetTablesWithColumnsReq 获取表和字段信息请求
type GetTablesWithColumnsReq struct {
	g.Meta `path:"/table/columns" method:"get" tags:"Generate" summary:"获取所有表和字段信息（过滤时间字段）"`
}

// GetTablesWithColumnsRes 获取表和字段信息响应
type GetTablesWithColumnsRes struct {
	List []TableWithColumnsInfo `json:"list" dc:"表和字段信息列表"`
}

// GetTableColumnsReq 根据表名获取字段信息请求
type GetTableColumnsReq struct {
	g.Meta    `path:"/table/{tableName}/columns" method:"get" tags:"Generate" summary:"根据表名获取字段信息"`
	TableName string `json:"tableName" v:"required#请输入表名" dc:"表名"`
}

// GetTableColumnsRes 根据表名获取字段信息响应
type GetTableColumnsRes struct {
	TableName string       `json:"tableName" dc:"表名"`
	Columns   []ColumnInfo `json:"columns" dc:"字段列表"`
}
