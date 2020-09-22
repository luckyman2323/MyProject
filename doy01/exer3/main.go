package main

import (
	"fmt"
)

func main() {

	stack := new(ArrayStack)
	stack.Push("1")
	stack.Push("2")
	stack.Push("3")
	fmt.Println("size:", stack.Size())
	fmt.Println("pop:", stack.Pop())
	fmt.Println("pop:", stack.Pop())
	fmt.Println("size:", stack.Size())
	stack.Push("4")
	fmt.Println("pop:", stack.Pop())

}

// 数组栈，后进先出
type ArrayStack struct {
	array []string   // 底层切片
	size  int        // 栈的元素数量
}

// 栈大小
func (stack *ArrayStack) Size() int {
	return stack.size
}


// 入栈
func (stack *ArrayStack) Push(v string) {
	stack.array = append(stack.array, v)
	stack.size = stack.size + 1
}

//出栈
func (stack *ArrayStack) Pop() string {

	if stack.size == 0 {
		panic("empty")
	}

	// 栈顶元素
	v := stack.array[stack.size-1]

	//建立新的数组
	newArray := make([]string, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArray[i] = stack.array[i]
	}
	stack.array = newArray

	// 栈中元素数量-1
	stack.size = stack.size - 1
	return v
}

// 获取栈顶元素
func (stack *ArrayStack) Peek() string {
	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 栈顶元素值
	v := stack.array[stack.size-1]
	return v
}