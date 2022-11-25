package array

import "fmt"

//bubbleSort 冒泡排序
func bubbleSort(data []int) {
	r := len(data) - 1
	for i := 0; i < r; i++ {
		for j := r; j > i; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
			fmt.Println("====", j, "======", data)
		}
	}
}

//insertSort 插入排序
func insertSort(data []int) {
	r := len(data) - 1
	for i := 1; i <= r; i++ {
		for j := i; j > 0 && data[j] < data[j-1]; j-- {
			fmt.Println("====", j, "======", data)
			data[j], data[j-1] = data[j-1], data[j]
		}

	}
}

//selectSort 选择排序
func selectSort(data []int) {
	r := len(data) - 1
	for i := 0; i < r; i++ {
		min := i
		for j := i + 1; j <= r; j++ {
			if data[j] < data[j-1] {
				min = j
			}
		}
		i, min = min, i
	}
}
