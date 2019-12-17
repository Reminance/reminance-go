/**
 * @File : demo04_pointerparamter
 * @Author: Octal_H
 * @Date : 2019/11/6
 * @Desc :
 */

package main

import "fmt"

func main() {
	/*
		指针作为参数:

		参数的传递:值传递,引用传递

	*/

	a := 10
	fmt.Println("fun1()函数调用前,a:", a)
	fun1(a)
	fmt.Println("fun1()函数调用后,a:", a)

	fun2(&a)
	fmt.Println("fun2()函数调用后,a:", a)

	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("fun3()函数调用前:", arr1)
	fun3(arr1)
	fmt.Println("fun3()函数调用后:", arr1)

	fun4(&arr1)
	fmt.Println("fun4()函数调用后:", arr1)

}

func fun4(p2 *[4]int) {
	fmt.Println("fun4()函数中的数组指针:", p2)
	p2[0] = 200
	fmt.Println("fun4()函数中的数组指针:", p2)
}

func fun3(arr2 [4]int) {
	fmt.Println("fun3()函数中数组的:", arr2)
	arr2[0] = 100
	fmt.Println("fun3()函数中修改数组:", arr2)
}

func fun1(num int) { //值传递:num = a = 10
	fmt.Println("fun1()函数中,num的值:", num)
	num = 100
	fmt.Println("fun1()函数中修改num的值:", num)
}

func fun2(p1 *int) { //传递a的地址,引用传递
	fmt.Println("fun2()函数中,p1的值:", *p1)
	*p1 = 200
	fmt.Println("fun2()函数中修改*p1的值:", *p1)
}
