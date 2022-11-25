package array

import "fmt"

func ArrayAndSlice() {
	//数组声明
	// var a1 = [3]int{1, 2, 3}
	// var a2 = [...]int{1, 2, 3}
	// var a3 = [...]int{0: 1, 1: 2, 2: 2}
	// var a4 = [3]int{}

	//切片声明
	// var s1 []int
	// var s2 = []int{1, 2, 3}
	// var s3 = make([]int, 3)

	var arry = [4]int{1, 2, 3, 4}

	slice := arry[1:3]

	slice[0] = 5

	fmt.Println("arry=", arry)
	fmt.Println("slice=", slice, "len=", len(slice), "cap=", cap(slice))

	slice = arry[3:]

	fmt.Println("arry=", arry)
	fmt.Println("slice=", slice, "len=", len(slice), "cap=", cap(slice))

	slice = arry[:]

	fmt.Println("arry=", arry)
	fmt.Println("slice=", slice, "len=", len(slice), "cap=", cap(slice))

	slice = append(slice, []int{3}...)

	fmt.Println("arry=", arry)
	fmt.Println("slice=", slice, "len=", len(slice), "cap=", cap(slice))

	slice[0] = 7

	fmt.Println("arry=", arry)
	fmt.Println("slice=", slice, "len=", len(slice), "cap=", cap(slice))

}
