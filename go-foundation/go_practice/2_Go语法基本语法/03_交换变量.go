/**
 * @File : 3_交换变量
 * @Author: Octal_H
 * @Date : 2019/11/12
 * @Desc :
 */

package main

import "fmt"

func main() {
	/*var a int = 100
	var b int = 200
	var t int

	t = a
	a = b
	b = t
	fmt.Println(a,b)*/

	var a int = 100
	var b int = 200

	/*a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Println(a,b)*/

	b, a = a, b
	fmt.Println(a, b)

}
