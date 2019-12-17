/**
 * @File : 18_临界资源安全问题
 * @Author: Octal_H
 * @Date : 2019/11/11
 * @Desc :
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		临界资源：

	*/
	a := 1
	go func() {
		a = 2
		fmt.Println("goroutine中...", a)
	}()

	a = 3
	time.Sleep(1)
	fmt.Println("main goroutine中...", a)

}
