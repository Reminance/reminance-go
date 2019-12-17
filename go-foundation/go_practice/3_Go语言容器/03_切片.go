/**
 * @File : 03_切片
 * @Author: Octal_H
 * @Date : 2019/11/14
 * @Desc :
 */

package main

import "fmt"

func main() {
	/*
		1) 从指定范围中生成切片
		切片和数组密不可分，如果将数组理解为一栋办公楼，那么切片就是把不同的连续楼层出租给使用者，
		出租的过程需要选择开始楼层和结束楼层，这个过程就会生成切片
	*/
	var highRiseBuilding [30]int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	// 区间
	fmt.Println(highRiseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])
	// 开头到中间指定位置的所有元素
	fmt.Println(highRiseBuilding[:2])

	fmt.Println("---------------------------")
	/*
		2) 表示原有的切片
		生成切片的格式中，当开始和结束位置都被忽略时，生成的切片将表示和原切片一致的切片，并且生成的切片与原切片在数据内容上也是一致的
	*/
	a := []int{1, 2, 3}
	fmt.Println(a[:])

	fmt.Println("---------------------------")
	/*
		3) 重置切片，清空拥有的元素
		把切片的开始和结束位置都设为 0 时，生成的切片将变空
	*/
	fmt.Println(a[0:0])

	fmt.Println("---------------------------")
	/*
		直接声明新的切片
	*/
	//声明字符串切片
	var strList []string
	//声明整型切片
	var numList []int
	//声明一个空切片
	var numListEmpty = []int{}

	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)

	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))

	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)

	fmt.Println("----------------------------")
	/*
		使用 make() 函数构造切片
	*/
	a1 := make([]int, 2)
	a2 := make([]int, 2, 10)
	fmt.Println(a1, a2)
	fmt.Println(len(a1), len(a2))

}
