package service

import (
	"giligili/model"
	"giligili/serializer"
)

// 视频删除
type DeleteVideoService struct {
}

func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "视频删除失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{}
}
