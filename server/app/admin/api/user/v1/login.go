package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LoginReq 登录请求参数
type LoginReq struct {
	g.Meta     `path:"/login" method:"post" tags:"用户认证" summary:"用户登录"`
	Username   string `json:"username" v:"required#请输入用户名" dc:"用户名"`
	Password   string `json:"password" v:"required#请输入密码" dc:"密码"`
	CaptchaId  string `p:"captchaId" v:"required" dc:"验证码ID"`
	VerifyCode string `p:"verifyCode" v:"required" dc:"验证码"`
}

// LoginRes 登录返回参数
type LoginRes struct {
	Avatar       string      `json:"avatar" dc:"头像"`
	Username     string      `json:"username" dc:"用户名"`
	Nickname     string      `json:"nickname" dc:"昵称"`
	Roles        []string    `json:"roles" dc:"角色列表"`
	Permissions  []string    `json:"permissions" dc:"权限列表"`
	AccessToken  string      `json:"accessToken" dc:"访问令牌"`
	RefreshToken string      `json:"refreshToken" dc:"刷新令牌"`
	Expires      *gtime.Time `json:"expires" dc:"过期时间"`
}

// RefreshTokenReq 刷新令牌请求参数
type RefreshTokenReq struct {
	g.Meta       `path:"/refresh-token" method:"post" tags:"用户认证" summary:"刷新令牌"`
	RefreshToken string `json:"refreshToken" v:"required#请提供刷新令牌" dc:"刷新令牌"`
}

// RefreshTokenRes 刷新令牌返回参数
type RefreshTokenRes struct {
	AccessToken  string      `json:"accessToken" dc:"访问令牌"`
	RefreshToken string      `json:"refreshToken" dc:"刷新令牌"`
	Expires      *gtime.Time `json:"expires" dc:"过期时间"`
}

type CaptchaReq struct {
	g.Meta `path:"/captcha" method:"get" tags:"登陆验证码" summary:"获取验证码"`
}

type CaptchaRes struct {
	CaptchaId  string `json:"captchaId"`
	CaptchaImg string `json:"captchaImg"`
}
