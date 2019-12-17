/**
 * @File : 22_channel通道初识2
 * @Author: Octal_H
 * @Date : 2019/11/11
 * @Desc :
 */

package main

import "fmt"

func main() {
	var ch1 chan bool
	ch1 = make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine中,i:", i)
		}
		//循环结束后，向通道中写数据，表示要结束了。。
		ch1 <- true
		fmt.Println("结束...")
	}()
	data := <-ch1
	fmt.Println("main...data-->", data)
	fmt.Println("main...over...")
}
