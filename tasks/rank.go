package tasks

import "giligili/cache"

// 一天执行一次，重新计算日排行
func RestartDailyRank() error {
	return cache.RedisClient.Del(cache.DailyRankKey).Err()
}
