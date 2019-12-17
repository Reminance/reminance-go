/**
 * @File : 14_Go语言指针
 * @Author: Octal_H
 * @Date : 2019/11/12
 * @Desc :
 */

package main

import (
	"flag"
	"fmt"
)

// 定义命令行参数
var mode = flag.String("mode", "", "process mode")

func main() {

	//指针实际用法
	//var cat int = 1
	//var str string = "hello go"
	//fmt.Printf("%p, %p",&cat,&str)

	//从指针获取指针指向的值
	//准备一个字符串类型
	var house string = "Malibu Point 10880, 90265"
	//对字符串取地址,ptr类型为*string
	ptr := &house
	//打印ptr的类型
	fmt.Printf("ptr type:%T\n", ptr) //*string
	//打印ptr的指针地址
	fmt.Printf("address:%p\n", ptr) //0xc0000401f0
	// 对指针进行取值操作
	value := *ptr
	//取值后的类型
	fmt.Printf("value type:%T\n", value) //string
	//指针取值后就是指向变量的值
	fmt.Printf("value:%s\n", value) //Malibu Point 10880, 90265

	fmt.Println("---------------------------")
	//使用指针修改值
	//准备两个变量
	x, y := 1, 2
	//交换变量值
	swap(&x, &y)
	//输出变量值
	fmt.Println(x, y)

	fmt.Println("---------------------------")
	//使用指针变量获取命令行的输入信息
	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	fmt.Println(*mode)

}

func swap(a, b *int) {
	//取a指针的值,赋给临时变量t
	t := *a
	//取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t
}
