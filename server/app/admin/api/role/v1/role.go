package v1

import (
	"server/app/admin/api/common/page"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleCommon 角色公共字段
type RoleCommon struct {
	Name   *string `json:"name,omitempty" v:"required#请输入角色名称" dc:"角色名称"`
	Code   *string `json:"code,omitempty" v:"required#请输入角色编码" dc:"角色编码（唯一）"`
	Status *int    `json:"status,omitempty" v:"in:0,1#请选择角色状态|状态只能是0或1" dc:"状态（1启用，0禁用）"`
	Remark *string `json:"remark,omitempty" dc:"备注"`
}

// GetListReq 查询角色列表请求参数
type GetListReq struct {
	g.Meta `path:"/role" method:"get" tags:"角色管理" summary:"获取角色列表"`
	page.ReqPage
	Name   string `json:"name" dc:"角色名称"`
	Code   string `json:"code" dc:"角色编码"`
	Status *int   `json:"status" v:"in:0,1#状态只能是0或1" dc:"状态（1启用，0禁用）"`
}

// GetListRes 查询角色列表返回参数
type GetListRes struct {
	page.ResPage
	List []RoleInfo `json:"list" dc:"角色列表"`
}

// CreateReq 创建角色请求参数
type CreateReq struct {
	g.Meta `path:"/role" method:"post" tags:"角色管理" summary:"创建角色"`
	RoleCommon
}

// CreateRes 创建角色返回参数
type CreateRes struct {
	Id uint64 `json:"id" dc:"角色ID"`
}

// UpdateReq 更新角色请求参数
type UpdateReq struct {
	g.Meta `path:"/role/{id}" method:"put" tags:"角色管理" summary:"更新角色"`
	Id     uint64 `json:"id" v:"required#请输入角色ID" dc:"角色ID"`
	RoleCommon
}

// UpdateRes 更新角色返回参数
type UpdateRes struct{}

// DeleteReq 删除角色请求参数
type DeleteReq struct {
	g.Meta `path:"/role/{id}" method:"delete" tags:"角色管理" summary:"删除角色"`
	Id     uint64 `json:"id" v:"required#请输入角色ID" dc:"角色ID"`
}

// DeleteRes 删除角色返回参数
type DeleteRes struct{}

// GetAllReq 获取所有角色请求参数
type GetAllReq struct {
	g.Meta `path:"/all-role" method:"get" tags:"角色管理" summary:"获取所有角色列表"`
}

// GetAllRes 获取所有角色返回参数
type GetAllRes struct {
	List []RoleInfo `json:"list" dc:"角色列表"`
}

// RoleInfo 角色信息
type RoleInfo struct {
	Id        uint64      `json:"id" dc:"主键ID"`
	Name      string      `json:"name" dc:"角色名称"`
	Code      string      `json:"code" dc:"角色编码（唯一）"`
	Status    int         `json:"status" dc:"状态（1启用，0禁用）"`
	Remark    string      `json:"remark" dc:"备注"`
	CreatedAt *gtime.Time `json:"createTime" dc:"创建时间"`
}

// AssignMenusReq 分配角色菜单权限请求参数
type AssignMenusReq struct {
	g.Meta  `path:"/role/{id}/menus" method:"post" tags:"角色管理" summary:"分配角色菜单权限"`
	Id      uint64   `json:"id" v:"required#请输入角色ID" dc:"角色ID"`
	MenuIds []uint64 `json:"menuIds" dc:"菜单ID列表"`
}

// AssignMenusRes 分配角色菜单权限返回参数
type AssignMenusRes struct{}

// GetRoleMenuIdsReq 获取角色菜单ID列表请求参数
type GetRoleMenuIdsReq struct {
	g.Meta `path:"/role/{id}/menu-ids" method:"get" tags:"角色管理" summary:"获取角色菜单ID列表"`
	Id     uint64 `json:"id" v:"required#请输入角色ID" dc:"角色ID"`
}

// GetRoleMenuIdsRes 获取角色菜单ID列表返回参数
type GetRoleMenuIdsRes struct {
	MenuIds []uint64 `json:"menuIds" dc:"菜单ID列表"`
}
