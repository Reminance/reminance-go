/**
 * @File : util
 * @Author: Octal_H
 * @Date : 2019/11/10
 * @Desc :
 */

package utils

import "fmt"

func Count() {
	fmt.Println("utils包下的Count函数...")
}

func init() {
	fmt.Println("utils另一个init函数")
}

func init() {
	fmt.Println("utils包下的init函数,用于初始化一些信息")
}
