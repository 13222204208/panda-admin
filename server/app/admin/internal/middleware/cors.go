package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// CORS 跨域中间件
func CORS(r *ghttp.Request) {
	// 设置CORS头
	r.Response.CORSDefault()
	
	// 或者自定义CORS设置
	// r.Response.Header().Set("Access-Control-Allow-Origin", "http://localhost:8848")
	// r.Response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// r.Response.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")
	// r.Response.Header().Set("Access-Control-Allow-Credentials", "true")
	
	// 处理预检请求
	if r.Method == "OPTIONS" {
		r.Response.WriteHeader(200)
		return
	}
	
	r.Middleware.Next()
}