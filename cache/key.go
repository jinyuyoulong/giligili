package cache

import (
	"fmt"
	"strconv"
)

// 本文件用来管理Redis 中的 key

const (
	// 日排行
	DailyRankKey = "rank:daily"
)

// VedioViewKey 视频点击数的 key
// view:video view: redis 的通常做法，用来做命名空间
// view:video:1 一号视频的点击数
// view:artical
func VedioViewKey(id uint) string  {
	return fmt.Sprintf("view:video:%s",strconv.Itoa(int(id)))
}
