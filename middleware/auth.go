package middleware

import (
	"giligili/model"
	"giligili/serializer"
	"github.com/gin-gonic/gin"
)

func CurrentUser() gin.HandlerFunc {
	
	return func(c *gin.Context) {
		uid, _ := c.Cookie("user_id")
		if uid != "" {
			user, err := model.GetUser(uid)
			if err == nil{
				c.Set("user",&user)
			}
		}
		c.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.Response{
			Status: 401,
			Msg:    "需要登录",
		})
		c.Abort()
	}
}
