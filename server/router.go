package server

import (
	"giligili/api"
	"giligili/middleware"
	"os"

	//"giligili/middleware"
	"github.com/gin-gonic/gin"
	//"os"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件顺序不能变
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET"))) // session 为了保持登录状态
	r.Use(middleware.Cors())                               // 跨域
	r.Use(middleware.CurrentUser())                        // 登录中间件，验证登录

	v1 := r.Group("/api/v1")
	{
		// 面向接口编程：先把接口定义出来

		v1.POST("/ping", api.Ping)
		v1.GET("/ping", api.Ping)

		// 视频投稿
		// 视频详情
		// 视频列表
		// 视频更新
		// 视频删除

		v1.POST("/user/login", api.UserLogin)
		v1.POST("/user/register", api.UserRegister)

		authed := r.Group("/")
		// 需要登录保护的
		authed.Use(middleware.AuthRequired())
		{
			//	user routing
			authed.GET("user/me", api.UserMe)
			authed.DELETE("user/logout", api.UserLogout)
		}

		v1.POST("/videos", api.CreateVideo)
		v1.GET("/video/:id", api.ShowVideo)
		v1.GET("/videos", api.ListVideo)
		v1.PUT("/video/:id", api.UpdateVideo)
		v1.DELETE("/video/:id", api.DeleteVideo)

		v1.POST("/upload/token", api.UploadToken)
	}
	return r
}
