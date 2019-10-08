package main

import (
	"giligili/conf"
	"giligili/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
