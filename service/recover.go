package service

import (
	"fmt"
	"log"
)

func tast2() {
	fmt.Println("sssssssss")
}
func tast(x int) {
	defer func() {
		fmt.Println("eeeeeeeeee")
		if err:=recover();err !=nil{
			fmt.Println(err)
		}
	}()
	var a [10]int
	a[x] = 1222
	log.Println(a)
}
