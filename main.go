package main

import (
	"giligili/conf"
	"giligili/server"
	"giligili/tasks"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")

	// 开启定时任务
	tasks.CronJob()
}
