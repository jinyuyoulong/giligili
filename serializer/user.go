package serializer

import "giligili/model"

type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	NickName  string `json:"nick_name"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

type UserResponse struct {
}

func BuildUser(user *model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		NickName:  user.NickName,
		Status:    user.Status,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
