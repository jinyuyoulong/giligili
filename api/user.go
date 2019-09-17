package api

import (
	"giligili/serializer"
	"giligili/service"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {

}
func UserLogout(c *gin.Context) {

}

func UserRegister(c *gin.Context) {
	service := service.UserRegisteService{}

	if err := c.ShouldBind(&service); err == nil{
		if user, err := service.Register(); err != nil{
			c.JSON(200,err)
		}else {
			res := serializer.BuildUserResponse(&user)
			c.JSON(200,res)
		}

	}else {
		c.JSON(200,ErrorResponse(err))
	}

}

func UserMe(c *gin.Context) {

}
