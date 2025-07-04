package v1

import (
	"server/app/admin/api/common/page"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CodeGenRecordInfo 代码生成记录信息
type CodeGenRecordInfo struct {
	Id           uint64      `json:"id" dc:"主键ID"`
	TableName    string      `json:"tableName" dc:"数据表名称"`
	TableComment string      `json:"tableComment" dc:"表注释"`
	PackageName  string      `json:"packageName" dc:"生成的Go包名"`
	ModuleName   string      `json:"moduleName" dc:"模块名"`
	Options      string      `json:"options" dc:"配置选项"`
	Columns      string      `json:"columns" dc:"表字段"`
	Status       int         `json:"status" dc:"生成状态（1成功，0失败）"`
	CreatedAt    *gtime.Time `json:"createdAt" dc:"生成时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// GetCodeGenRecordListReq 获取代码生成记录列表请求
type GetCodeGenRecordListReq struct {
	g.Meta `path:"/generate/record" method:"get" tags:"代码生成" summary:"获取代码生成记录列表"`
	page.ReqPage
	TableName  string `json:"tableName" dc:"表名称（模糊搜索）"`
	ModuleName string `json:"moduleName" dc:"模块名（模糊搜索）"`
	Status     *int   `json:"status" v:"in:0,1#状态只能是0或1" dc:"生成状态（1成功，0失败）"`
}

// GetCodeGenRecordListRes 获取代码生成记录列表响应
type GetCodeGenRecordListRes struct {
	page.ResPage
	List []CodeGenRecordInfo `json:"list" dc:"代码生成记录列表"`
}

// GetCodeGenRecordDetailReq 获取代码生成记录详情请求
type GetCodeGenRecordDetailReq struct {
	g.Meta `path:"/generate/record/{id}" method:"get" tags:"代码生成" summary:"获取代码生成记录详情"`
	Id     uint64 `json:"id" v:"required|min:1#记录ID不能为空|记录ID必须大于0" dc:"记录ID"`
}

// GetCodeGenRecordDetailRes 获取代码生成记录详情响应
type GetCodeGenRecordDetailRes struct {
	CodeGenRecordInfo
}

// DeleteCodeGenRecordReq 删除代码生成记录请求
type DeleteCodeGenRecordReq struct {
	g.Meta `path:"/generate/record/{id}" method:"delete" tags:"代码生成" summary:"删除代码生成记录"`
	Id     uint64 `json:"id" v:"required#请输入记录ID" dc:"记录ID"`
}

// DeleteCodeGenRecordRes 删除代码生成记录响应
type DeleteCodeGenRecordRes struct{}

// UpdateCodeGenRecordReq 更新代码生成记录请求
type UpdateCodeGenRecordReq struct {
	g.Meta       `path:"/generate/record/{id}" method:"put" tags:"代码生成" summary:"更新代码生成记录"`
	Id           uint64 `json:"id" v:"required#请输入记录ID" dc:"记录ID"`
	TableName    string `json:"tableName" v:"required#请输入表名称" dc:"数据表名称"`
	TableComment string `json:"tableComment" dc:"表注释"`
	PackageName  string `json:"packageName" v:"required#请输入包名" dc:"生成的Go包名"`
	ModuleName   string `json:"moduleName" v:"required#请输入模块名" dc:"模块名"`
	Options      string `json:"options" dc:"配置选项"`
	Columns      string `json:"columns" dc:"表字段"`
}

// UpdateCodeGenRecordRes 更新代码生成记录响应
type UpdateCodeGenRecordRes struct{}
