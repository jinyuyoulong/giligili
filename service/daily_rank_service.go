package service

import (
	"giligili/cache"
	"giligili/model"
	"giligili/serializer"
)

type DailyRankService struct {
	
}

func (d *DailyRankService) Get() serializer.Response {
	var videos  []model.Video

	vids, _ := cache.RedisClient.ZRevRange(cache.DailyRankKey, 0, 9).Result()
	if len(vids) > 1 {
		model.DB.Where("id in (?)",vids)
	}

	return serializer.Response{
		Data:
	}
}