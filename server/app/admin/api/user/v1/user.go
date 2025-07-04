package v1

import (
	"server/app/admin/api/common/page"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserCommon 用户公共字段
type UserCommon struct {
	Title        *string `json:"title,omitempty" dc:"职位名称"`
	DepartmentId *uint64 `json:"departmentId,omitempty" dc:"所属部门ID"`
	Nickname     *string `json:"nickname,omitempty" dc:"昵称"`
	Username     *string `json:"username,omitempty" dc:"用户名"`
	Password     *string `json:"password,omitempty" v:"length:6,20#请输入密码|密码长度应在6-20位之间" dc:"密码（加密存储）"`
	Phone        *string `json:"phone,omitempty" v:"phone#请输入正确的手机号格式" dc:"联系电话"`
	Email        *string `json:"email,omitempty" v:"email#请输入正确的邮箱格式" dc:"邮箱地址"`
	Sex          *int    `json:"sex,omitempty" v:"in:0,1,2#性别只能是0,1,2" dc:"性别（0未知，1男，2女）"`
	Status       *int    `json:"status,omitempty" v:"in:0,1#请选择用户状态|状态只能是0或1" dc:"状态（1启用，0禁用）"`
	Remark       *string `json:"remark,omitempty" dc:"备注"`
}

// GetListReq 查询用户列表请求参数
type GetListReq struct {
	g.Meta `path:"/user" method:"get" tags:"用户管理" summary:"获取用户列表"`
	page.ReqPage
	DepartmentId *uint64 `json:"departmentId" dc:"所属部门ID"`
	Username     string  `json:"username" dc:"用户名"`
	Phone        string  `json:"phone" dc:"联系电话"`
	Status       *int    `json:"status" v:"in:0,1#状态只能是0或1" dc:"状态（1启用，0禁用）"`
}

// GetListRes 查询用户列表返回参数
type GetListRes struct {
	page.ResPage
	List []UserInfo `json:"list" dc:"用户列表"`
}

// CreateReq 创建用户请求参数
type CreateReq struct {
	g.Meta `path:"/user" method:"post" tags:"用户管理" summary:"创建用户"`
	UserCommon
}

// CreateRes 创建用户返回参数
type CreateRes struct {
	Id uint64 `json:"id" dc:"用户ID"`
}

// UpdateReq 更新用户请求参数
type UpdateReq struct {
	g.Meta `path:"/user/{id}" method:"put" tags:"用户管理" summary:"更新用户"`
	Id     uint64 `json:"id" v:"required#请输入用户ID" dc:"用户ID"`
	UserCommon
}

// 部门详情
type Dept struct {
	Id   uint64 `json:"id" dc:"主键ID"`
	Name string `json:"name" dc:"部门名称"`
}

// UpdateRes 更新用户返回参数
type UpdateRes struct{}

// DeleteReq 删除用户请求参数
type DeleteReq struct {
	g.Meta `path:"/user/{id}" method:"delete" tags:"用户管理" summary:"删除用户"`
	Id     uint64 `json:"id" v:"required#请输入用户ID" dc:"用户ID"`
}

// DeleteRes 删除用户返回参数
type DeleteRes struct{}

// GetDetailReq 获取用户详情请求参数
type GetDetailReq struct {
	g.Meta `path:"/user/detail" method:"get" tags:"用户管理" summary:"获取用户详情"`
}

// GetDetailRes 获取用户详情返回参数
type GetDetailRes struct {
	UserInfo
}

// ResetPasswordReq 重置密码请求参数
type ResetPasswordReq struct {
	g.Meta      `path:"/user/{id}/reset-password" method:"put" tags:"用户管理" summary:"重置用户密码"`
	Id          uint64 `json:"id" v:"required#请输入用户ID" dc:"用户ID"`
	OldPassword string `json:"oldPassword" dc:"旧密码"`
	Password    string `json:"password" v:"required|length:6,20#请输入密码|密码长度应在6-20位之间" dc:"新密码"`
}

// ResetPasswordRes 重置密码返回参数
type ResetPasswordRes struct{}

// UserInfo 用户信息
type UserInfo struct {
	Id           uint64      `json:"id" dc:"主键ID"`
	Title        string      `json:"title" dc:"职位名称"`
	Avatar       string      `json:"avatar" dc:"头像"`
	DepartmentId uint64      `json:"departmentId" dc:"所属部门ID"`
	Dept         Dept        `json:"dept" dc:"所属部门"`
	Nickname     string      `json:"nickname" dc:"昵称"`
	Username     string      `json:"username" dc:"用户名"`
	Phone        string      `json:"phone" dc:"联系电话"`
	Email        string      `json:"email" dc:"邮箱地址"`
	Sex          int         `json:"sex" dc:"性别（0未知，1男，2女）"`
	Status       int         `json:"status" dc:"状态（1启用，0禁用）"`
	Remark       string      `json:"remark" dc:"备注"`
	CreatedAt    *gtime.Time `json:"createTime" dc:"创建时间"`
}

// RoleInfo 角色信息（用于用户详情中的角色列表）
type RoleInfo struct {
	Id   uint64 `json:"id" dc:"角色ID"`
	Name string `json:"name" dc:"角色名称"`
	Code string `json:"code" dc:"角色编码"`
}

// BatchDeleteReq 批量删除用户请求参数
type BatchDeleteReq struct {
	g.Meta `path:"/user/batch" method:"delete" tags:"用户管理" summary:"批量删除用户"`
	Ids    []uint64 `json:"ids" v:"required#请选择要删除的用户" dc:"用户ID列表"`
}

// BatchDeleteRes 批量删除用户返回参数
type BatchDeleteRes struct{}

// GetRoleIdsReq 获取用户角色ID列表请求参数
type GetRoleIdsReq struct {
	g.Meta `path:"/user-role-ids" method:"post" tags:"用户管理" summary:"获取用户对应的角色ID列表"`
	UserId uint64 `json:"userId" v:"required#请输入用户ID" dc:"用户ID"`
}

// GetRoleIdsRes 获取用户角色ID列表返回参数
type GetRoleIdsRes struct {
	RoleIds []uint64 `json:"roleIds" dc:"角色ID列表"`
}

// AssignRolesReq 分配用户角色请求参数
type AssignRolesReq struct {
	g.Meta  `path:"/user/assign-roles" method:"post" tags:"用户管理" summary:"分配用户角色"`
	UserId  uint64   `json:"userId" v:"required#请输入用户ID" dc:"用户ID"`
	RoleIds []uint64 `json:"roleIds" dc:"角色ID列表"`
}

// AssignRolesRes 分配用户角色返回参数
type AssignRolesRes struct{}

// UploadAvatarReq 上传用户头像请求参数
type UploadAvatarReq struct {
	g.Meta `path:"/user/{id}/avatar" method:"post" tags:"用户管理" summary:"上传用户头像"`
	Id     uint64 `json:"id" v:"required#请输入用户ID" dc:"用户ID"`
	Avatar string `json:"avatar" v:"required#请上传头像文件" dc:"头像文件base64或URL"`
}

// UploadAvatarRes 上传用户头像返回参数
type UploadAvatarRes struct {
	AvatarUrl string `json:"avatarUrl" dc:"头像URL"`
}
