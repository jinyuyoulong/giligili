package service

import (
	"giligili/model"
	"giligili/serializer"
)

type UserLoginService struct {
	UserName string `from:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `from:"password" json:"password" binding:"required,min=8,max=40"`
}

func (service *UserLoginService) Login() (model.User, *serializer.Response) {
	var user model.User
	if err := model.DB.Where("user_name = ?", user.UserName).First(&user).Error; err != nil {
		return user, &serializer.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}

	if !user.CheckPassword(service.Password) {
		return user, &serializer.Response{
			Status: 40001,
			Msg:    "账号或密码错误",
		}
	}
	return user, nil
}
