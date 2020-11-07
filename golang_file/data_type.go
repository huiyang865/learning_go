package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// 定义bool类型
	gender := false
	fmt.Println(gender)

	// go语言有int8, int16, int32, int64, uint8, uint16, uint32, uint64整数类型
	var age int8 = 18
	var ageBk int = 28 // go语言中int是一种动态类型，根据机器字节数确定，在64位机器上为int64，32位机器上为int32
	fmt.Println(age, ageBk)

	// go语言有float32,float64浮点类型
	var weight float32 = 64.88 // 浮点没有float类型，必须指明是float32, float64类型
	fmt.Println(weight)
	fmt.Println(math.MaxFloat64) // float64比int64最大值远大很多，float64底层存储和int存储方式不一样
	weightedBk := 72.33
	fmt.Printf("%T\n", weightedBk)

	// byte和rune
	// byte类型 uint8的别称，主要用来做字节转换，静态语言中做中文处理容易出错，用rune处理较多
	var a byte = 6
	fmt.Println(a)
	// rune是int32的别称
	var b rune = 6
	fmt.Println(b)

	// 字符
	aChar := 'a' // 单引号表示字符
	fmt.Printf("%T\n", aChar+1)
	fmt.Printf("a+1=%c\n", aChar+1) //python 中没有字符概念，只有字符串

	// 类型转换
	// 基本类型转换
	// valueOfTypeB = (typeB) valueOfTypeA
	//　1. go 语言中不支持变量间隐式转换
	aa := int(3.0)
	fmt.Println(aa)
	var ab int = 5.0 // 5.0是常量，支持隐式转换
	// var ab int = 5.１　隐式转换不能将5.1转换成整数，可以用显式转换
	fmt.Println(ab)
	var ac float64 = 4.5
	fmt.Printf("%T\n", ac)

	// var ad int = ac 不能将ac float64类型隐式转换成int类型
	var ad int = int(ac)
	fmt.Println(ad)

	// var ae int = int("1") 不能将string类型直接显式转换成int
	// 2. 应用strconv字符串和int之间的转换，Itoa, Atoi
	fmt.Println("a" + strconv.Itoa(aa))
	// fmt.Println("a" + strconv.Itoa(b)) // 只能转int类型，不能转float64,或者int64等类型
	fmt.Println("a" + strconv.Itoa(int(b)))
	var i int
	i, _ = strconv.Atoi("3")
	fmt.Println(3 + i)

	// 3. Parse类函数用途更广，可以转换string和bool等类型
	ba, _ := strconv.ParseBool("true")
	fmt.Println(ba)
}
