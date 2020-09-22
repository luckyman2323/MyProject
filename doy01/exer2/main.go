package main

import "fmt"

func main() {

	var arr1 = [9]int{33, 20, 17, 30, 90, 43, 58, 50, 90}

	BubSort(&arr1)

	fmt.Println("arr1降序排序后为:", arr1)

}

func BubSort(arr1 *[9]int) {
	var temp int = 0 //替换变量
	for i := 0; i < len(*arr1)-1; i++ {
		for j := 0; j < len(*arr1)-1-i; j++ {
			if (*arr1)[j] < (*arr1)[j+1] {
				temp = (*arr1)[j]
				(*arr1)[j] = (*arr1)[j+1]
				(*arr1)[j+1] = temp
			}
		}
	}
}
