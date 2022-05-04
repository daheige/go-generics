package main

import (
	"fmt"
)

// 泛型结构体
type collection[T any] struct {
	collect []T
}

// New 创建一个 collection[T any] struct ptr
func New[T any]() *collection[T] {
	c := &collection[T]{
		collect: []T{},
	}

	return c
}

// Append append value to collection
// *collection[T] 作为方法的接收者
func (c *collection[T]) Append(value ...T) {
	c.collect = append(c.collect, value...)
}

// Map 传递一个回调函数处理value 得到一个新的collection
func (c *collection[T]) Map(trans func(value T) T) *collection[T] {
	var list []T
	for _, val := range c.collect {
		list = append(list, trans(val))
	}

	return &collection[T]{
		collect: list,
	}
}

// Result 得到集合中的元素
func (c *collection[T]) Result() []T {
	return c.collect
}

// Display 打印value
func (c *collection[T]) Display() {
	for _, val := range c.collect {
		fmt.Println("current value: ", val)
	}
}

func main() {
	var c = New[int]() // New后面需要显式指定类型
	c.Append(1, 2, 3)
	c.Append(4, 5, 6)
	c.Display()

	fmt.Println("map callback")
	c.Map(func(val int) int {
		return val * 2
	}).Display()

	result := c.Result()
	fmt.Println("result: ", result)
}
