package middleware

//默认禁止跨域请求，此处设置白名单，允许

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// cross 跨域请求
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	// "http://localhost" 只可以读，写入失败
	// "http://localhost:8080" 跨区请求，写入成功。
	config.AllowOrigins = []string{"http://localhost:8080", "http://10.211.55.5:3001", "http://localhost:3001"}
	config.AllowCredentials = true

	return cors.New(config)
}
