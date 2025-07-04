package v1

import (
	"server/app/admin/api/common/page"
	"server/app/admin/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// MemberCommon 会员信息表公共字段
type MemberCommon struct {
	Username *string `json:"username,omitempty" v:"required#请输入用户名" dc:"用户名"`
	Email *string `json:"email,omitempty"  dc:"邮箱"`
	Mobile *string `json:"mobile,omitempty" v:"required#请输入手机号" dc:"手机号"`
}


// CreateMemberReq 创建会员信息表请求
type CreateMemberReq struct {
    g.Meta `path:"/member" method:"post" tags:"会员信息表" summary:"创建会员信息表"`
    MemberCommon
}

// CreateMemberRes 创建会员信息表响应
type CreateMemberRes struct {}



// UpdateMemberReq 更新会员信息表请求
type UpdateMemberReq struct {
    g.Meta `path:"/member/{id}" method:"put" tags:"会员信息表" summary:"更新会员信息表"`
    Id uint64 `json:"id" v:"required#请输入ID" dc:"ID"`
    MemberCommon
}

// UpdateMemberRes 更新会员信息表响应
type UpdateMemberRes struct {}



// DeleteMemberReq 删除会员信息表请求
type DeleteMemberReq struct {
    g.Meta `path:"/member/{id}" method:"delete" tags:"会员信息表" summary:"删除会员信息表"`
    Id uint64 `json:"id" v:"required#请输入ID" dc:"ID"`
}

// DeleteMemberRes 删除会员信息表响应
type DeleteMemberRes struct {}



// BatchDeleteMemberReq 批量删除会员信息表请求
type BatchDeleteMemberReq struct {
    g.Meta `path:"/member/batch" method:"delete" tags:"会员信息表" summary:"批量删除会员信息表"`
    Ids []uint64 `json:"ids" v:"required#请输入ID列表" dc:"ID列表"`
}

// BatchDeleteMemberRes 批量删除会员信息表响应
type BatchDeleteMemberRes struct {}



// GetMemberListReq 获取会员信息表列表请求
type GetMemberListReq struct {
    g.Meta `path:"/member" method:"get" tags:"会员信息表" summary:"获取会员信息表列表"`
    page.ReqPage
    Username string `json:"username" dc:"用户名"`
    Email string `json:"email" dc:"邮箱"`
}

// GetMemberListRes 获取会员信息表列表响应
type GetMemberListRes struct {
    List []*entity.Member `json:"list" dc:"会员信息表列表"`
    page.ResPage
}


