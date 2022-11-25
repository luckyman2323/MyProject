package service

import (
	"math/rand"
	"time"
)

func GetRandInt(i int) int {
	rand.Seed(time.Now().UnixNano())
	//随机生成100以内的正整数
	return rand.Intn(i)
}

func GetRandIntArry(i int) []int {
	rand.Seed(time.Now().UnixNano())
	res := rand.Perm(i)
	return res
}
