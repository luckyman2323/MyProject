package main


import (
	"fmt"
)

func main() {

	queue := new(ArrayQueue)
	queue.Push("1")
	queue.Push("2")
	queue.Push("3")
	fmt.Println("size:", queue.Size())
	fmt.Println("pop:", queue.Pop())
	fmt.Println("pop:", queue.Pop())
	fmt.Println("size:", queue.Size())
	queue.Push("4")
	fmt.Println("pop:", queue.Pop())

}

// 数组队，先进先出
type ArrayQueue struct {
	array []string   // 底层切片
	size  int        // 栈的元素数量
}

// 队大小
func (queue *ArrayQueue) Size() int {
	return queue.size
}


// 入队
func (queue *ArrayQueue) Push(v string) {
	queue.array = append(queue.array, v)
	queue.size = queue.size + 1
}

//出队
func (queue *ArrayQueue) Pop() string {

	if queue.size == 0 {
		panic("empty")
	}

	// 队首元素
	v := queue.array[0]

	//建立新的数组
	newArray := make([]string, queue.size-1, queue.size-1)
	for i := 1; i <= queue.size-1; i++ {
		newArray[i-1] = queue.array[i]
	}
	queue.array = newArray

	// 栈中元素数量-1
	queue.size = queue.size - 1
	return v
}

// 获取队首元素
func (queue *ArrayQueue) Peek() string {
	// 队中元素已空
	if queue.size == 0 {
		panic("empty")
	}

	// 队首元素值
	v := queue.array[0]
	return v
}

