package tasks

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"reflect"
	"runtime"
	"time"
)

var Cron *cron.Cron

func Run(job func () error) {
	println("run")
	// 获得job 执行的时间差
	from := time.Now().UnixNano()
	err := job()
	// 获得job 执行的时间差
	to := time.Now().UnixNano()
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	if err != nil{
		fmt.Printf("%s error:%dms\n", jobName, (to-from)/int64(time.Millisecond))
	}else {
		fmt.Printf("%s success:%dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}
func CronJob()  {
	if Cron == nil {
		//Cron = cron.New(cron.WithSeconds())
		Cron = cron.New()
	}

	//Cron.AddFunc("*/1 * * * * *", func() {
	//	Run(RestartDailyRank)
	//})
	Cron.AddFunc("* * */1 * *", func() {
		Run(RestartDailyRank)
	})
	Cron.Start()

	fmt.Println("crontab start...")
}
