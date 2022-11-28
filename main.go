package main

import (
	"myproject/logs"
	"time"
)

func init() {
	logs.InitLogger()
}

func main() {
	var str = 1 << 28

	// res, err := strconv.ParseInt(str, 10, 64)
	logs.Logger.Info("res: %v, err: %s", str, nil)
	time.Sleep(time.Second)
}
