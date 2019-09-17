package model

import (
	"giligili/cache"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

type Video struct {
	gorm.Model
	Title  string
	Info   string
	URL    string
	Avatar string
	UserID uint
}

// AvatarURL 实现签名,私有访问
func (v *Video) AvatarURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signURL, _ := bucket.SignURL(v.Avatar, oss.HTTPGet, 600)
	return signURL
}

// View 点击数
func (v *Video) View() uint64 {
	resultStr, _ := cache.RedisClient.Get(cache.VedioViewKey(v.ID)).Result()
	count, _ := strconv.ParseUint(resultStr, 10, 64)
	return count
}

// AddView 视频浏览
func (v *Video) AddView() {
	// 增加视频点击数
	cache.RedisClient.Incr(cache.VedioViewKey(v.ID))
	// 增加排行榜点击数
	cache.RedisClient.ZIncrBy(cache.DailyRankKey,1,strconv.Itoa(int(v.ID)))
}