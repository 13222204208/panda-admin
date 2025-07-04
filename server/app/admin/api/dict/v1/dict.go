package v1

import (
	"server/app/admin/api/common/page"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DictCommon 字典公共字段
type DictCommon struct {
	Title     *string `json:"title,omitempty" v:"required#请输入字典标题" dc:"字典标题"`
	DictType  *string `json:"dictType,omitempty" v:"required#请输入字典类型" dc:"字典类型"`
	DictLabel *string `json:"dictLabel,omitempty" v:"required#请输入字典标签" dc:"字典标签"`
	DictValue *string `json:"dictValue,omitempty" v:"required#请输入字典值" dc:"字典值"`
	Sort      *int    `json:"sort,omitempty" v:"min:0#排序值不能小于0" dc:"排序值"`
	Status    *int    `json:"status,omitempty" v:"in:0,1#请选择字典状态|状态只能是0或1" dc:"状态（1启用，0禁用）"`
	Remark    *string `json:"remark,omitempty" dc:"备注说明"`
}

// GetListReq 查询字典列表请求参数
type GetListReq struct {
	g.Meta `path:"/dict" method:"get" tags:"字典管理" summary:"获取字典列表"`

	DictType  string `json:"dictType" dc:"字典类型"`
	DictLabel string `json:"dictLabel" dc:"字典标签"`
	Status    *int   `json:"status" v:"in:0,1#状态只能是0或1" dc:"状态（1启用，0禁用）"`
	page.ReqPage
}

// GetListRes 查询字典列表返回参数
type GetListRes struct {
	List []DictInfo `json:"list" dc:"字典列表"`
	page.ResPage
}

// UpdateReq 更新字典请求参数
type UpdateReq struct {
	g.Meta `path:"/dict/{id}" method:"put" tags:"字典管理" summary:"更新字典"`
	Id     uint64 `json:"id" v:"required#请输入字典ID" dc:"字典ID"`
	DictCommon
}

// UpdateRes 更新字典返回参数
type UpdateRes struct{}

// DeleteReq 删除字典请求参数
type DeleteReq struct {
	g.Meta `path:"/dict/{id}" method:"delete" tags:"字典管理" summary:"删除字典"`
	Id     uint64 `json:"id" v:"required#请输入字典ID" dc:"字典ID"`
}

// DeleteRes 删除字典返回参数
type DeleteRes struct{}

// BatchDeleteReq 批量删除字典请求参数
type BatchDeleteReq struct {
	g.Meta `path:"/dict/batch-delete" method:"post" tags:"字典管理" summary:"批量删除字典"`
	Ids    []uint64 `json:"ids" v:"required#请选择要删除的字典" dc:"字典ID列表"`
}

// BatchDeleteRes 批量删除字典返回参数
type BatchDeleteRes struct{}

// BatchCreateReq 批量创建字典请求参数
type BatchCreateReq struct {
	g.Meta    `path:"/dict/batch" method:"post" tags:"字典管理" summary:"批量创建字典"`
	Title     string          `json:"title" v:"required#请输入字典标题" dc:"字典标题"`
	DictType  string          `json:"dictType" v:"required#请输入字典类型" dc:"字典类型"`
	DictItems []BatchDictItem `json:"dictItems" v:"required#请添加字典项" dc:"字典项列表"`
}

// BatchCreateRes 批量创建字典返回参数
type BatchCreateRes struct{}

// BatchDictItem 批量字典项
type BatchDictItem struct {
	DictLabel string  `json:"dictLabel" v:"required#请输入字典标签" dc:"字典标签"`
	DictValue string  `json:"dictValue" v:"required#请输入字典值" dc:"字典值"`
	Sort      *int    `json:"sort,omitempty" v:"min:0#排序值不能小于0" dc:"排序值"`
	Remark    *string `json:"remark,omitempty" dc:"备注说明"`
}

// GetOptionsReq 获取字典选项请求参数
type GetOptionsReq struct {
	g.Meta   `path:"/dict/options/{dictType}" method:"get" tags:"字典管理" summary:"根据字典类型获取字典选项"`
	DictType string `json:"dictType" v:"required#请输入字典类型" dc:"字典类型"`
}

// GetOptionsRes 获取字典选项返回参数
type GetOptionsRes struct {
	Options []DictOption `json:"options" dc:"字典选项列表"`
}

// DictOption 字典选项
type DictOption struct {
	Label string `json:"label" dc:"字典标签"`
	Value string `json:"value" dc:"字典值"`
}

// DictInfo 字典信息
type DictInfo struct {
	Id        uint64      `json:"id" dc:"主键ID"`
	Title     string      `json:"title" dc:"字典标题"`
	DictType  string      `json:"dictType" dc:"字典类型"`
	DictLabel string      `json:"dictLabel" dc:"字典标签"`
	DictValue string      `json:"dictValue" dc:"字典值"`
	Sort      int         `json:"sort" dc:"排序值"`
	Status    int         `json:"status" dc:"状态（1启用，0禁用）"`
	Remark    string      `json:"remark" dc:"备注说明"`
	CreatedAt *gtime.Time `json:"createTime" dc:"创建时间"`
}

// GetDistinctTypesReq 获取不重复的字典类型和标题请求参数
type GetDistinctTypesReq struct {
	g.Meta `path:"/dict/types" method:"get" tags:"字典管理" summary:"获取不重复的字典类型和标题"`
}

// GetDistinctTypesRes 获取不重复的字典类型和标题返回参数
type GetDistinctTypesRes struct {
	Types []DictTypeInfo `json:"types" dc:"字典类型列表"`
}

// DictTypeInfo 字典类型信息
type DictTypeInfo struct {
	Title    string `json:"title" dc:"字典标题"`
	DictType string `json:"dictType" dc:"字典类型"`
}
