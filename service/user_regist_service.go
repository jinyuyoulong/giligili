package service

import (
	"giligili/model"
	"giligili/serializer"
)

type UserRegisteService struct {
	UserName string `from:"user_name" json:"user_name" binding:"required,min=2,max=30"`
	NickName string `from:"password" json:"nick_name" binding:"required,min=2,max=40"`
	Password string `from:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `from:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

func (service *UserRegisteService) Register() (model.User, *serializer.Response){
	err := service.Valid()
	if err != nil{
		return model.User{}, err
	}

	user := model.User{
		UserName:service.UserName,
		NickName:service.NickName,
		Status:model.Active,
	}

	// 加密密吗

	if err := user.SetPassword(service.Password); err != nil {
		return user,&serializer.Response{
			Status:40002,
			Msg:"密码加密失败",
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return user,&serializer.Response{
			Status:40002,
			Msg:"注册失败",
		}
	}

	return  user,nil

}

func (service *UserRegisteService)Valid() *serializer.Response {
	if service.Password != service.PasswordConfirm {
		return &serializer.Response{
			Status:40001,
			Msg:"两次密码不一致",
		}
	}

	// select count(*) from user where nickname = "bilibili"
	// 昵称检查
	count := 0
	model.DB.Model(&model.User{}).Where("nick_name = ?",service.NickName).Count(&count)

	if count > 0{
		return &serializer.Response{
			Status:40001,
			Msg:"昵称被占用",
		}
	}
	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?",service.UserName).Count(&count)
	if count > 0{
		return &serializer.Response{
			Status:40001,
			Msg:"用户名已经注册",
		}
	}

	return nil
}