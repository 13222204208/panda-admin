package middleware

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
)

const (
	// JwtSecretKey JWT密钥，实际项目中应该从配置文件读取
	JwtSecretKey = "your-secret-key-here-change-in-production"
	// CtxUsername 上下文中用户名的键
	CtxUsername = "username"
	// CtxUserID 上下文中用户ID的键
	CtxUserID = "id"
)

// JWTClaims JWT声明
type JWTClaims struct {
	UserID   uint64 `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Auth JWT认证中间件
func Auth(r *ghttp.Request) {
	// 跳过登录和刷新令牌接口
	if r.URL.Path == "/admin/login" || r.URL.Path == "/admin/refresh-token" || r.URL.Path == "/admin/captcha" {
		r.Middleware.Next()
		return
	}

	// 获取Authorization头
	auth := r.Header.Get("Authorization")
	if auth == "" {
		r.Response.WriteJsonExit(g.Map{
			"code":    401,
			"message": "未提供认证令牌",
		})
		return
	}

	// 解析Bearer token
	if !strings.HasPrefix(auth, "Bearer ") {
		r.Response.WriteJsonExit(g.Map{
			"code":    401,
			"message": "认证令牌格式错误",
		})
		return
	}

	tokenString := strings.TrimPrefix(auth, "Bearer ")

	// 验证token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecretKey), nil
	})

	if err != nil || !token.Valid {
		r.Response.WriteJsonExit(g.Map{
			"code":    401,
			"message": "无效的认证令牌",
		})
		return
	}

	if claims, ok := token.Claims.(*JWTClaims); ok {
		// 将用户信息存储到上下文中
		ctx := context.WithValue(r.Context(), CtxUserID, claims.UserID)
		ctx = context.WithValue(ctx, CtxUsername, claims.Username)
		r.SetCtx(ctx)
	}

	r.Middleware.Next()
}
