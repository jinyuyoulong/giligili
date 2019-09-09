package service

import (
	"giligili/model"
	"giligili/serializer"
)

// 视频详情
type ListVideoService struct {
	// binging 绑定 required 必填 mix rune 的 数量
	Title string `from:"title" json:"title" binding:"required,min=2,max=100"`
	Info  string `from:"info" json:"info" binding:"max=3000"`
	URL   string `from:"url" json:"url" binding:"max=3000"`
}

// 视频创建
func (service *ListVideoService) ListVideo() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
