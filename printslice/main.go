package main

import (
	"fmt"
)

// print slice
// 泛型函数定义
func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println("value:", v)
	}
}

func main() {
	fmt.Println("Hello, 世界")
	var i any = 1 // 特意类型
	i = 12.3
	fmt.Println(i)
	printSlice([]int{1, 2, 3, 4})
}
