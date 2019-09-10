package service

import (
	"giligili/model"
	"giligili/serializer"
)

// 视频详情
type ShowVideoService struct {
	// binging 绑定 required 必填 mix rune 的 数量
	Title string `from:"title" json:"title" binding:"required,min=2,max=100"`
	Info  string `from:"info" json:"info" binding:"max=3000"`
	URL   string `from:"url" json:"url" binding:"max=3000"`
}

// 视频创建
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	//处理视频被观看的一系问题
	video.AddView()

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
