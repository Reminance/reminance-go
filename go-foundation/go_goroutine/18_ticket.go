/**
 * @File : 18_ticket
 * @Author: Octal_H
 * @Date : 2019/11/11
 * @Desc :
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//全局变量
var ticket = 10 //10张票

func main() {
	/*
		4个goroutine,模拟4个售票口
	*/

	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	time.Sleep(3 * time.Second)

}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出:", ticket)
			ticket--
		} else {
			fmt.Println(name, "售完,没有票了...")
			break
		}
	}
}
