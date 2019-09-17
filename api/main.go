package api

import (
	"encoding/json"
	"fmt"
	"giligili/conf"
	"giligili/model"
	"giligili/serializer"
	"gopkg.in/go-playground/validator.v8"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "Pong",
	})
}

// 获取当前用户
func CurrentUser(c *gin.Context) *model.User {

	if user, _ := c.Get("user"); user != nil {

		// 断言，转为user类型指针
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.Response{
				Status: 40001,
				Msg:    fmt.Sprintf("%s%s", field, tag),
				Error:  fmt.Sprint(err),
			}
		}
	}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 40001,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	return serializer.Response{
		Status: 40001,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
