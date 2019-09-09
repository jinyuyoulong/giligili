package service

import (
	"giligili/model"
	"giligili/serializer"
)

// 视频投稿的服务
type CreateVideoService struct {
	// binging 绑定 required 必填 mix rune 的 数量
	Title  string `from:"title" json:"title" binding:"required,min=2,max=100"`
	Info   string `from:"info" json:"info" binding:"max=3000"`
	URL    string `from:"url" json:"url" binding:"max=3000"`
	Avatar string `from:"avatar" json:"avatar"`
}

// 视频创建
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title:  service.Title,
		Info:   service.Info,
		URL:    service.URL,
		Avatar: service.Avatar,
	}

	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "视频保存失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}

// 错误码 定义一般规则
// 4 开头 用户操作错误
// 5 开头 系统错误
