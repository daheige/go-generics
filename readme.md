# go1.18 generics
    go generics要求安装go1.18.1+版本

# goland 升级到最新版本
    推荐GoLand 2022.1 版本

# goimports 包升级
    go install -v golang.org/x/tools/cmd/goimports@latest

# generics tutorial
https://golang.google.cn/doc/tutorial/generics

# generics 设计理念
    类型参数
        一般在泛型函数中定义，比如 `func Print[T any](s []T)`

    类型约束 type constraint 
        可以对于int,int64,float64,float32,string,chan,map,slice 等不同类型进行约束

    类型自动推导
        对于不能自动推导的需要显式指定类型，比如说 New[int]() 创建一个collection

    类型集合
        新语法，他的标识符是 “~”，完整用法是 ~T。
        ~T 是指底层类型为 T 的所有类型的集合。
        比如说 ~int | ~int8 | ~int16 | ~int32 | ~int64
        这个集合类型可以嵌套到别的类型中

# generics 约束
    - comparable 约束，使得传入的参数必须能进行 == 与 != 操作
    - any 约束，可以是任意类型，在go底层就是type any = interface{}
    - interface 约束其实有三种: 可选类型约定、method方法约束、可选和方法混合约束

# go1.18 generics 实战
    常见的泛型使用方式：
    1.generics func (printslice)
    2.generics type constraint eg: string,int,map,chan,slice (sum)
    3.generics struct (collection)
    4.generics method (collection)

# 参数推导
```go
func Map[F, T any](s []F, f func(F) T) []T { ... }
```

# 联合元素
```go
// 联合元素 约束多个类型
type Number interface {
    int8 | int16 | int32 | int | int64 | float32 | float64
}
```
# 近似元素和嵌入约束
```go
// Signed is a constraint whose type set is any signed integer type.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint whose type set is any unsigned integer type.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Float is a constraint whose type set is any floating point type.
type Float interface {
	~float32 | ~float64
}

// Ordered is a constraint whose type set is any ordered type.
// That is, any type that supports the < operator.
type Ordered interface {
	Signed | Unsigned | Float | ~string
}
```

# 接口类型
```go
// Stringish 的类型集将是字符串类型和所有实现 fmt.Stringer 的类型，
// 任何这些类型（包括 fmt.Stringer 本身）将被允许作为这个约束的类型参数。
type Stringish interface {
	string | fmt.Stringer
}
```

# go generics设计说明

    底层使用Go应用程序的知名无服务器数据库平台Planetscale，在Go 1.18刚发布不久，其工程师便针对新的泛型功能进行测试，
    并在自家博客发布文章提到，Go的泛型功能可能使程序代码变慢，而该篇文章在网络被热烈讨论，因为Go开发者盼望了泛型功能许久，
    最后拿到的却与期待有所落差。

    而Google技术总监同时也是Go泛型主要设计者的Ian Lance Taylor，则在Go博客搬出2021年，他在Go Day与GopherCon的讲题
    When To Use Generics，来说明Go泛型的正确使用时机。
    Planetscale的工程师提到，之所以Go泛型可能使程序代码变慢的原因，是因为Go泛型的实例，
    并非使用完全单态化（Monomorphization），而是采用了一种称为GCShape Stenciling with Dictionaries的
    部分单态化技术，该方法虽然可大幅减少程序代码的量，但是却会在特定的情况下，使程序代码变慢。
    
    而Ian Lance Taylor提到，Go的程序开发通用准则，要开发者通过编写程序代码来撰写Go程序，而非定义类型，因此谈到泛型，
    要是开发者通过定义类型参数约束来撰写程序，那一开始就走错路了，应该从编写函数开始，而在之后就能够自然地添加类型参数，
    因为类型参数在该情况展现的好处将会非常明显。
    
    Ian Lance Taylor列出4个类型参数可能有用的情况：
    第一是使用语言定义的特殊容器类型，进行操作的函数，像是slice、map和channel，当函数具有这些类型的参数，
    并且函数程序代码没有对元素类型进行任何特定假设，则类型参数可能有用。

    另一个参数类型有用的情况，是用于通用数据结构，像是连接串行（Linked List）或是二元树（Ｂinary Tree），
    这些类似slice或map但又非语言内置的通用数据结构。
    
    第三则是类型参数首选函数而非方法，开发者可以将方法转换成函数，会比将方法添加到类型中简单的多。
    最后一个类型参数可发挥效果的情况，是不同类型需要实现通用方法，且不同类型的实例看起来都相同时。

    Ian Lance Taylor也提醒了不适合使用类型参数的时机，包括不要用类型参数代替接口类型（Interface Type），且当方法实例不同时，
    也不要使用类型参数。他提到Go 1.18实例泛型的方法，使得类型参数通常不会比接口类型快，因此不需要为了执行速度更改程序代码。
    
    Ian Lance Taylor最后给出了简单的泛型使用方针，就是当开发者发现自己多次编写相同的程序代码，而副本不同之处的唯一区别，
    仅在于使用了不同类型，便可以考虑使用类型参数。

# go1.18 设计变化点
https://golang.google.cn/doc/go1.18

大的变化：
1. go generics
2. go work
3. core library
   eg: New debug/buildinfo package and New net/netip package
