package service

import (
	"fmt"
	"giligili/serializer"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"path/filepath"

	"github.com/google/uuid"
	"os"
)

type UploadTokenService struct {
	Filename string `from:"filename" json:"filename"`
}

func (service *UploadTokenService) Post() serializer.Response {
	end_point := os.Getenv("OSS_END_POINT")
	fmt.Printf("%v\n", end_point)

	client, err := oss.New(end_point, os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	// 带可选参数的签名直传
	options := []oss.Option{
		oss.ContentType("image/png"),
	}

	ext := filepath.Ext(service.Filename)
	// 随机数 防止文件名重名，避免冲突
	key := "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ext

	// 签名直传
	signPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	// 查看 image
	signGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)

	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signPutURL,
			"get": signGetURL,
		},
	}
}
