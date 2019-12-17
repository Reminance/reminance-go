/**
 * @File : demo03_error
 * @Author: Octal_H
 * @Date : 2019/11/8
 * @Desc :
 */

package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupHost("www.ocpro.cn")
	fmt.Print(err)
	if ins, ok := err.(*net.DNSError); ok {
		if ins.Timeout() {
			fmt.Println("操作超时。。")
		} else if ins.Temporary() {
			fmt.Println("临时性错误。。")
		} else {
			fmt.Println("通常错误。。")
		}
	}
	fmt.Println(addr)
}
