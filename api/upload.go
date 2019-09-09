package api

import (
	"giligili/service"
	"github.com/gin-gonic/gin"
)

func UploadToken(c *gin.Context) {
	service := service.UploadTokenService{}
	err := c.ShouldBind(&service)
	if err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
