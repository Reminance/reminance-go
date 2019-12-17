/**
 * @File : 04_append()为切片添加元素
 * @Author: Octal_H
 * @Date : 2019/11/14
 * @Desc :
 */

package main

import "fmt"

func main() {
	/*
		append() 可以为切片动态添加元素
			果空间不足以容纳足够多的元素，切片就会进行“扩容”，此时新切片的长度会发生改变。
	*/
	var a []int
	a = append(a, 1)                 // 追加1个元素
	a = append(a, 1, 2, 3)           // 追加多个元素, 手写解包方式
	a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包

	//切片在扩容时，容量的扩展规律是按容量的 2 倍数进行扩充，例如 1、2、4、8、16……
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}

	//在切片的开头添加元素
	a = append([]int{0}, a...)          // 在开头添加1个元素
	a = append([]int{-3, -2, -1}, a...) // 在开头添加1个切片

	//切片也支持链式操作，我们可以将多个 append 操作组合起来，实现在切片中间插入元素：
	//var a2 []int
	//a2 = append(a2[:i], append([]int{x}, a2[i:]...)...) // 在第i个位置插入x
	//a2 = append(a2[:i], append([]int{1,2,3}, a2[i:]...)...) // 在第i个位置插入切片
}
