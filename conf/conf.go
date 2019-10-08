package conf

import (
	"giligili/cache"
	"giligili/model"
	"giligili/tasks"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化相关配置
func Init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		panic(err)
	}

	//	连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()

	// 开启定时任务
	tasks.CronJob()
}
