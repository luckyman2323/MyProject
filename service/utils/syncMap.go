package service

import (
	"fmt"
	"sync"
	"time"
)

var ServiceMap sync.Map

func CheckEndpoint() {
	ServiceMap.Range(func(key, value interface{}) bool {
		fmt.Println("===", key)
		time.Sleep(1*time.Second)
		return true
	})
}
