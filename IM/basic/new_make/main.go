package main

import "fmt"

func main() {
	test1()
}

//在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，
//否则我们的值就没办法存储。
//而对于值类型的声明不需要分配内存空间， 是因为它们在声明的时候已经默认分配好了内存空间。
//要分配内存， 就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存
package main

import "fmt"

func main() {
	test1()
}

//在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，
//否则我们的值就没办法存储。
//而对于值类型的声明不需要分配内存空间， 是因为它们在声明的时候已经默认分配好了内存空间。
//要分配内存， 就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存

func test1() {
	var a int
	a = 10//而对于值类型的声明不需要分配内存空间， 是因为它们在声明的时候已经默认分配好了内存空间。
	fmt.Println(a)

	var v *int
	v = new(int)
	*v = 11
	fmt.Println(*v)
}

func test2() {
	defer func() {
		fmt.Println(recover())
	}()
	var v *int
	*v = 11
	fmt.Println(*v)
}

func test3() {
	defer func() {
		fmt.Println(recover())
	}()

	var v *int

	fmt.Println(v) //<nil>
	fmt.Println(*v) //地址不存在, 无法取值
}

//new(T) 为一个 T 类型新值分配空间并将此空间初始化为 T 的零值，返回的是新值的地址
func test4()  {
	v := new(int)
	*v = 10
	fmt.Println(v)
	fmt.Println(*v)
}

func test5() {
	var m *map[int]int
	fmt.Printf("%T", m)
	fmt.Println()

	m = new(map[int]int)
	fmt.Println(m)
	fmt.Printf("%T", m)
}

func new_map_test() {
	//使用new来创建map时，返回的内容是一个指针，
	//这个指针指向了一个所有字段全为0的值map对象，需要初始化后才能使用，
	//而使用make来创建map时，返回的内容是一个引用，可以直接使用。

	//使用new创建一个map指针
	ma := new(map[string]int)
	//第一种初始化方法
	*ma = map[string]int{}

	(*ma)["a"] = 44
	fmt.Println(*ma)

	//第二种初始化方法
	*ma = make(map[string]int, 0)
	(*ma)["b"] = 55
	fmt.Println(*ma)

	//第三种初始化方法
	mb := make(map[string]int, 0)
	mb["c"] = 66
	*ma = mb
	(*ma)["d"] = 77
	fmt.Println(*ma)
}


func test1() {
	var a int
	a = 10//而对于值类型的声明不需要分配内存空间， 是因为它们在声明的时候已经默认分配好了内存空间。
	fmt.Println(a)

	var v *int
	v = new(int)
	*v = 11
	fmt.Println(*v)
}

func test2() {
	defer func() {
		fmt.Println(recover())
	}()
	var v *int
	*v = 11
	fmt.Println(*v)
}

func test3() {
	defer func() {
		fmt.Println(recover())
	}()

	var v *int

	fmt.Println(v) //<nil>
	fmt.Println(*v) //地址不存在, 无法取值
}

//new(T) 为一个 T 类型新值分配空间并将此空间初始化为 T 的零值，返回的是新值的地址
func test4()  {
	v := new(int)
	*v = 10
	fmt.Println(v)
	fmt.Println(*v)
}

func test5() {
	var m *map[int]int
	fmt.Printf("%T", m)
	fmt.Println()

	m = new(map[int]int)
	fmt.Println(m)
	fmt.Printf("%T", m)
}

func new_map_test() {
	//使用new来创建map时，返回的内容是一个指针，
	//这个指针指向了一个所有字段全为0的值map对象，需要初始化后才能使用，
	//而使用make来创建map时，返回的内容是一个引用，可以直接使用。

	//使用new创建一个map指针
	ma := new(map[string]int)
	//第一种初始化方法
	*ma = map[string]int{}

	(*ma)["a"] = 44
	fmt.Println(*ma)

	//第二种初始化方法
	*ma = make(map[string]int, 0)
	(*ma)["b"] = 55
	fmt.Println(*ma)

	//第三种初始化方法
	mb := make(map[string]int, 0)
	mb["c"] = 66
	*ma = mb
	(*ma)["d"] = 77
	fmt.Println(*ma)
}

