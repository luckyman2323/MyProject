package main

import(
	"fmt"
	"math/rand"
	"sort"
	"time"
)
func main(){

	var intarr []int

	rand.Seed(time.Now().UnixNano())

	var len  = 10
	for i := 0; i < len; i++{
		intarr = append(intarr, rand.Intn(100))
	}

	fmt.Println(intarr)
	sort.Ints(intarr)
	fmt.Println(intarr)

}



