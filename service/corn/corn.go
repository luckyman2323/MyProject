package service

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func AddCorn() {
	// 定时任务
	cronService := cron.New()
	defer func() {
		cronService.Stop()
	}()

	cronService.AddFunc("0 */1 * * * ?", func() {
		fmt.Println("==========", time.Now())
	})

	fmt.Println("start==========", time.Now())
	cronService.Start()
	time.Sleep(10 * time.Minute)
}
