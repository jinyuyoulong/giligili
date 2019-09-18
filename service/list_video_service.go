package service

import (
	"fmt"
	"giligili/model"
	"giligili/serializer"
)

// 视频列表服务
type ListVideoService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// 视频创建
func (service *ListVideoService) List() serializer.Response {
	// 第一版逻辑
	//var videos []model.Video
	//err := model.DB.Find(&videos).Error
	//if err != nil {
	//	return serializer.Response{
	//		Status: 50000,
	//		Msg:    "数据库连接错误",
	//		Error:  err.Error(),
	//	}
	//}
	//
	//return serializer.Response{
	//	Data: serializer.BuildVideos(videos),
	//}
	fmt.Printf("limit:%d",service.Limit)

	var videos []model.Video
	total := 0

	if service.Limit == 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Video{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status:50000,
			Msg:"数据库连接错误",
			Error:err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&videos).Error; err != nil {
		return serializer.Response{
			Status:50000,
			Msg:"数据库连接错误",
			Error:err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildVideos(videos),uint(total))
}
