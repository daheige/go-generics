package main

import (
	"fmt"
	"strconv"
)

// sumInt64s takes a slice of string to int64 values.
func sumInt64s(s []int64) int64 {
	var sum int64
	for _, v := range s {
		sum += v
	}

	return sum
}

// sumFloats takes a slice of string to float64 values.
func sumFloats(s []float64) float64 {
	var sum float64
	for _, v := range s {
		sum += v
	}

	return sum
}

// =================1.Add non-generic functions===============
// 泛型函数定义
// sumInt64sOrFloats 在函数定义的时候，对参数进行类型约束
// Add a generic function to handle multiple types
// It supports both int64 and float64
// as types for slice values.
// 这里定义T泛型参数约定只能是int64 | float64
//
// Each type parameter has a type constraint that acts
// as a kind of meta-type for the type parameter.
// Each type constraint specifies the permissible
// type arguments that calling code can use for the respective type parameter.
func sumInt64sOrFloats[T int64 | float64](v []T) T {
	var sum T
	for _, val := range v {
		sum += val
	}

	return sum
}

// ==========2.Add a generic function to handle multiple types====
// sumWithMap 通过定义多个类型约束来定义泛型函数
// 声明一个泛型map 并对里面的value进行求和
// comparable 约束，使得传入的参数必须能进行 == 与 != 操作
// 这里定义来K是可比较的类型，对于V类型约束为int64 | float64
// 函数返回值是一个V泛型类型
func sumWithMap[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, val := range m {
		sum += val
	}

	return sum
}

// =============3.Declare a type constraint===========
// 泛型 Map
// H 声明一个泛型map
// type H[K string, V any] map[K]V
// 下面的key可以任意可比较的类型
// 这里的v是any约束，any可以是任意类型
type H[K comparable, V any] map[K]V

func foo() {
	m := H[string, int64]{
		"a": 1,
		"b": 2,
	}
	fmt.Println("m:", m)

	m2 := H[int, int]{
		1: 1,
		2: 2,
	}
	fmt.Println("m2: ", m2)

	m3 := H[string, string]{
		"a": "abc",
		"b": "23ac",
	}
	fmt.Println("m3: ", m3)
}

// 声明一个泛型通道
// Ch 通道Ch中的value是任意类型
type Ch[T any] chan T

// type constraint
// Number 类型约束多个类型
// 联合元素 约束多个类型
type Number interface {
	int8 | int16 | int32 | int | int64 | float32 | float64
}

// 近似元素
// 他的类型集是 ~int，也就是所有类型为 int 的类型（如：int、int8、int16、int32、int64）
// 都能够满足这个类型约束的条件，包括底层类型为 int 类型的（例如：类型别名）。
type AnyInt interface{ ~int }

// SignedInteger 新语法，他的标识符是 “~”，完整用法是 ~T。~T 是指底层类型为 T 的所有类型的集合
type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// NumberSlice number slice 定义指定类型约束的泛型切片
// 这里T是一个 Number 约束，NumberSlice 是一个复杂的泛型切片
type NumberSlice[T Number] []T

// 下面的函数是无法进行编译的
// cannot use generic type NumberSlice[T Number] without instantiation
// func printNumberSlice[T NumberSlice](s []T) {
//
// }

// 下面的也是无法进行编译的
// func printNumberSlice[T NumberSlice](s T) {
//
// }

// func printNumberSlice[V any,T NumberSlice[V]](s T) {
//
// }

// 指定类型的泛型切片遍历
func printNumberSlice[T NumberSlice[int]](s ...T) {
	for _, val := range s {
		fmt.Println("current value: ", val)
		for _, v := range val {
			fmt.Println("v: ", v)
		}
	}
}

func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println("value:", v)
	}
}

// 这里定义泛型函数，参数T类型约束为Number
// 在函数名字后面先要通过[]方式约定T是什么类型约束
func sumNumbers[T Number](nums ...T) T {
	var sum T
	for _, val := range nums {
		sum += val
	}

	return sum
}

// sumWithMap2 V是Number类型
// K 是 go 底层的 comparable 可比较的类型
func sumWithMap2[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, val := range m {
		sum += val
	}

	return sum
}

func printNumbers[T Number](s []T) {
	for _, v := range s {
		fmt.Printf("number: %v type:%T\n", v, v)
	}
}

// List 定义通用的slice泛型切片
// 声明变量的时候需要写var s = List[int]{1,2,3} 这种显式声明格式
type List[T any] []T

func main() {
	fmt.Println("=======Non-Generic=======")
	fmt.Println("=====sumInt64s====")
	fmt.Println(sumInt64s([]int64{1, 2, 3, 4}))
	fmt.Println("=====sumFloats====")
	fmt.Println(sumFloats([]float64{1.1, 2.1, 3.2, 4.1}))

	fmt.Println("=========with func generic=====")
	fmt.Println(sumInt64sOrFloats([]int64{1, 2, 3, 4, 5}))
	fmt.Println(sumInt64sOrFloats([]float64{1.1, 2.1, 3.2, 4.1}))

	fmt.Println("=====sumWithMap func generic========")
	fmt.Println("sumWithMap result:", sumWithMap(map[string]int64{
		"a": 1,
		"b": 2,
		"c": 3,
	}))
	fmt.Println("map generics")
	foo()

	// chan generics
	chanGenerics()

	// Number sum
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		sumNumbers(1, 2, 3),
		sumNumbers(1.2, 1.2, 1.3))

	fmt.Println("sumWithMap2 result:", sumWithMap2(map[string]int64{
		"a": 1,
		"b": 2,
		"c": 3,
	}))

	// =====printNumbers=====
	printNumbers([]int64{1, 2, 3})
	printNumbers([]float64{1.1, 2, 3})
	printNumbers([]int{1, 2, 3})

	// ==========NumberSlice=======
	fmt.Println("=====NumberSlice generics======")
	var s = NumberSlice[int]{1, 2, 3}
	printSlice(s)
	var s2 = NumberSlice[int]{4, 5, 6}
	printNumberSlice(s, s2)

	fmt.Println("=====slice generics====")
	var l = List[int]{1, 2, 3}
	printSlice(l)
	var l2 = List[float64]{1, 2, 3}
	printSlice(l2)
}

func chanGenerics() {
	ch := make(Ch[int], 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
	for val := range ch {
		fmt.Println("ch value: ", val)
	}

	ch2 := make(Ch[string], 10)
	for i := 0; i < 10; i++ {
		ch2 <- "hello: " + strconv.Itoa(i)
	}

	close(ch2)
	for val := range ch2 {
		fmt.Println("ch2 value: ", val)
	}
}
