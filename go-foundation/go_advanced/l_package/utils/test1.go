/**
 * @File : test1
 * @Author: Octal_H
 * @Date : 2019/11/10
 * @Desc :
 */

package utils

import "fmt"

//package main 	//Multiple packages in directory: utils, main more...

func MyTest2() {
	Count()
}

func init() {
	fmt.Println("utils包下的test1.go文件中的init()函数")
}
