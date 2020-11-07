package main // 首行必须是package，静态语言特性

import "fmt" //双引号

func test() (int, error) {
	return 0, nil
}

//变量的作用域
// 全局变量，只能用var定义
var aaa = 30

func var_scope() {
	// 局部变量
	c := 10
	fmt.Println(c, aaa)
}

// func 定义函数, 将主要逻辑放到main函数中，不需要手动调用main函数
func main() {
	// 字符串只能用双引号，单引号会报错
	// println('hello world!')
	println("hello world!") // 打印一行，换行
	// print("hello world!")　//　不换行
	fmt.Println("hello world!") // 使用更多

	// 变量的定义
	// 静态语言的变量与动态语言变量定义差异较大，python中不需要定义变量
	var i int = 4
	i = 20
	fmt.Println(i)

	var j = "hello world!" // 根据值推断变量类型
	// i = "hello world" 会报错，i类型不可改变
	fmt.Println(j)

	k := "hello world!" // 更接近python使用模式
	fmt.Println(k)

	// 多变量定义
	var a, b, c int = 10, 20, 30
	fmt.Println(a, b, c)

	// 匿名变量，go语言中，变量被定义后一定要使用，否则会报错
	_, err := test()
	if err != nil {
		fmt.Println("函数调用成功")
	}

	// go中常量不可被修改
	const PI = 3.1415926
	r := 2.0
	ci := 2 * PI * r
	fmt.Println(ci)

	const (
		xx int = 16
		yy
		ss = "abc"
		zz
	)
	fmt.Println(xx, yy, ss, zz)

	// const常量的iota，常量计数器
	// 1. iota只能再常量组中使用
	// 2. 各个常量const组之间互相不干扰
	// 3. 没有表达式的常量定义，复用上一行的表达式
	// 4. 从第一行开始，iota逐行加1
	const (
		ca = iota
		cb = 10 // 每一行iota加一
		cc
		cd = iota
		ce
		cf, ch = iota, iota // iota代表的是行数，不是常量个数
		cj     = iota
	)
	fmt.Println(ca, cb, cc, cd, ce, cf, ch, cj)

	var_scope()
}

// go run hello_world.go 执行脚本
// go build hello_world.go　编译脚本
