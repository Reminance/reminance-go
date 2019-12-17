/**
 * @File : demo01_error
 * @Author: Octal_H
 * @Date : 2019/11/8
 * @Desc :
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("06_l_error/96_错误error/test.txt")
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err) //open test.txt: no such file or directory
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println("1.Op:", ins.Op)
			fmt.Println("2.Path:", ins.Path)
			fmt.Println("3.Err:", ins.Err)
		}
		return
	}
	fmt.Println(f.Name(), "打开文件成功。。")

}
