package service

import "fmt"

func DeferTest() {
	var i int
	err := fmt.Errorf("111111")

	defer func(t1 int, t2 error){
		fmt.Println("i=", t1, "err=", t2)
	}(i, err)

	defer func(){
		fmt.Println("i=", i, "err=", err)
	}()

	err = fmt.Errorf("222222")
	i = 1

}