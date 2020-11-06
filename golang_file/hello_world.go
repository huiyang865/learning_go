package main // 首行必须是package，静态语言特性

import "fmt" //双引号


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
}


// go run hello_world.go 执行脚本
// go build hello_world.go　编译脚本
