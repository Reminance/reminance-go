/**
 * @File : 06_从切片中删除元素
 * @Author: Octal_H
 * @Date : 2019/11/14
 * @Desc :
 */

package main

import "fmt"

func main() {
	/*
		从开头位置删除
	*/
	//删除开头的元素可以直接移动数据指针
	a := []int{1, 2, 3}
	a = a[1:] // 删除开头1个元素
	//a = a[N:] // 删除开头N个元素
	fmt.Println(a)

	fmt.Println("-------------------")
	//也可以不移动数据指针，但是将后面的数据向开头移动，可以用 append 原地完成（所谓原地完成是指
	// 		在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化
	b := []int{1, 2, 3}
	b = b[:copy(b, b[1:])] // 删除开头1个元素
	//b = b[:copy(b, b[N:])] // 删除开头N个元素
	fmt.Println(b)

	fmt.Println("-------------------")
	/*
		从中间位置删除
			对于删除中间的元素，需要对剩余的元素进行一次整体挪动，同样可以用 append 或 copy 原地完成
	*/
	//a = []int{1, 2, 3, ...}
	//a = append(a[:i], a[i+1:]...) // 删除中间1个元素
	//a = append(a[:i], a[i+N:]...) // 删除中间N个元素
	//a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
	//a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素

	/*
		从尾部删除
	*/
	a = []int{1, 2, 3}
	a = a[:len(a)-1] // 删除尾部1个元素
	//a = a[:len(a)-N] // 删除尾部N个元素

	fmt.Println("---------------------")
	seq := []string{"a", "b", "c", "d", "e"}
	// 指定删除位置
	index := 2
	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index], seq[index+1:])
	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)

	fmt.Println(seq)
}
