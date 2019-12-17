/**
 * @File : 1_数组详解
 * @Author: Octal_H
 * @Date : 2019/11/13
 * @Desc :
 */

package main

import "fmt"

func main() {
	/*
		Go语言数组的声明
			var 数组变量名 [元素数量]Type
	*/
	var a [3]int             //定义三个整数的数组
	fmt.Println(a[0])        // 打印第一个元素
	fmt.Println(a[len(a)-1]) // 打印最后一个元素

	//打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// 仅打印元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	//使用数组字面值语法，用一组值来初始化数组：
	//var q [3]int = [3]int{1,2,3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	//在数组长度的位置出现“...”省略号，则表示数组的长度是根据初始化值的个数来计算
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	/*
		比较两个数组是否相等
	*/
	a1 := [2]int{1, 2}
	b1 := [...]int{1, 2}
	c1 := [2]int{1, 3}
	fmt.Println(a1 == b1, a1 == c1, b1 == c1) // "true false false"
	//d1 := [3]int{1, 2}
	//fmt.Println(a1 == d1) // 编译错误：无法比较 [2]int == [3]i

	/*
		遍历数组——访问每一个数组元素
	*/
	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"
	for k, v := range team {
		fmt.Println(k, v)
	}

}
