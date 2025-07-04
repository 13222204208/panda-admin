package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DepartmentCommon 部门公共字段
type DepartmentCommon struct {
	ParentId  *uint64 `json:"parentId,omitempty" v:"min:0#父级部门ID不能小于0" dc:"父级部门ID"`
	Name      *string `json:"name,omitempty" v:"required#请输入部门名称" dc:"部门名称"`
	Principal *string `json:"principal,omitempty" dc:"负责人名称"`
	Phone     *string `json:"phone,omitempty" dc:"联系电话"`
	Email     *string `json:"email,omitempty" v:"email#请输入正确的邮箱格式" dc:"邮箱地址"`
	Sort      *int    `json:"sort,omitempty" v:"min:0#排序号不能小于0" dc:"排序号"`
	Status    *int    `json:"status,omitempty" v:"in:0,1#请选择部门状态|状态只能是0或1" dc:"状态（1启用，0禁用）"`
	Remark    *string `json:"remark,omitempty" dc:"备注"`
}

// GetListReq 查询部门列表请求参数
type GetListReq struct {
	g.Meta `path:"/department" method:"get" tags:"部门管理" summary:"获取部门列表"`

	ParentId *uint64 `json:"parentId" dc:"父级部门ID"`
	Name     string  `json:"name" dc:"部门名称"`
	Status   *int    `json:"status" v:"in:0,1#状态只能是0或1" dc:"状态（1启用，0禁用）"`
}

// GetListRes 查询部门列表返回参数
type GetListRes struct {
	List []DepartmentInfo `json:"list" dc:"部门列表"`
}

// CreateReq 创建部门请求参数
type CreateReq struct {
	g.Meta `path:"/department" method:"post" tags:"部门管理" summary:"创建部门"`
	DepartmentCommon
}

// CreateRes 创建部门返回参数
type CreateRes struct {
	Id uint64 `json:"id" dc:"部门ID"`
}

// UpdateReq 更新部门请求参数
type UpdateReq struct {
	g.Meta `path:"/department/{id}" method:"put" tags:"部门管理" summary:"更新部门"`
	Id     uint64 `json:"id" v:"required#请输入部门ID" dc:"部门ID"`
	DepartmentCommon
}

// UpdateRes 更新部门返回参数
type UpdateRes struct{}

// DeleteReq 删除部门请求参数
type DeleteReq struct {
	g.Meta `path:"/department/{id}" method:"delete" tags:"部门管理" summary:"删除部门"`
	Id     uint64 `json:"id" v:"required#请输入部门ID" dc:"部门ID"`
}

// DeleteRes 删除部门返回参数
type DeleteRes struct{}

// DepartmentInfo 部门信息
type DepartmentInfo struct {
	Id        uint64      `json:"id" dc:"主键ID"`
	ParentId  uint64      `json:"parentId" dc:"父级部门ID"`
	Name      string      `json:"name" dc:"部门名称"`
	Principal string      `json:"principal" dc:"负责人名称"`
	Phone     string      `json:"phone" dc:"联系电话"`
	Email     string      `json:"email" dc:"邮箱地址"`
	Sort      int         `json:"sort" dc:"排序号"`
	Status    int         `json:"status" dc:"状态（1启用，0禁用）"`
	Remark    string      `json:"remark" dc:"备注"`
	CreatedAt *gtime.Time `json:"createTime" dc:"创建时间"`
}
