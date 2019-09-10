package api

// 此项目中 api 包是控制器
// 控制器中尽量少的写业务代码
// 主要职责是做组合，分发
// service:接收事情，处理事情；

import (
	"giligili/service"
	"github.com/gin-gonic/gin"
)

func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ShowVideo(c *gin.Context) {
	videoService := service.ShowVideoService{}
	res := videoService.Show(c.Param("id"))
	c.JSON(200, res)
}

func ListVideo(c *gin.Context) {
	service := service.ListVideoService{}

	res := service.ListVideo()
	c.JSON(200, res)
}

func UpdateVideo(c *gin.Context) {
	service := service.UpdateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func DeleteVideo(c *gin.Context) {
	service := service.DeleteVideoService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}

func RankDaily(c *gin.Context) {

}