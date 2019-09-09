package model

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
	"os"
)

type Video struct {
	gorm.Model
	Title  string
	Info   string
	URL    string
	Avatar string
}

// AvatarURL 实现签名,私有访问
func (v *Video) AvatarURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signURL, _ := bucket.SignURL(v.Avatar, oss.HTTPGet, 600)
	return signURL
}
