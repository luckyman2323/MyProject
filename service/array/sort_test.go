package array

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	nums := []int{9, 8, 5, 4, 7, 6, 3, 0, 1, 2}

	insertSort(nums) //使用冒泡排序
	fmt.Println(nums)
}
