/**
 * @File : 15_变量逃逸分析
 * @Author: Octal_H
 * @Date : 2019/11/13
 * @Desc :变量逃逸（Escape Analysis）——自动决定变量分配方式，提高运行效率
 */

package main

import "fmt"

// 声明空结构体测试结构体逃逸情况
type Data struct {
}

func main() {

	//1) 逃逸分析
	// 声明a变量并打印
	//var a int
	//// 调用void()函数
	//void()
	//// 打印a变量的值和dummy()函数返回
	//fmt.Println(a, dummy(0))
	/*
		go run -gcflags "-m -l" 15_变量逃逸分析.go
		# command-line-arguments
		.\15_变量逃逸分析.go:20:13: main ... argument does not escape
		.\15_变量逃逸分析.go:20:13: a escapes to heap
		.\15_变量逃逸分析.go:20:22: dummy(0) escapes to heap
		0 0
	*/

	fmt.Println("------------------")
	//取地址发生逃逸
	fmt.Println(dummy2())
	/*
		>go run -gcflags "-m -l" 15_变量逃逸分析.go
		# command-line-arguments
		.\15_变量逃逸分析.go:41:6: moved to heap: c
		.\15_变量逃逸分析.go:33:13: main ... argument does not escape
		.\15_变量逃逸分析.go:33:14: "------------------" escapes to heap
		.\15_变量逃逸分析.go:35:13: main ... argument does not escape
		.\15_变量逃逸分析.go:35:20: dummy2() escapes to heap

	*/

}

func dummy2() *Data {
	// 实例化c为Data类型
	var c Data
	//返回函数局部变量地址
	return &c
}

//本函数测试入口参数和返回值情况
func dummy(b int) int {
	//声明一个变量c并赋值
	var c int
	c = b

	return c

}

//空函数,声明也不做
func void() {

}
