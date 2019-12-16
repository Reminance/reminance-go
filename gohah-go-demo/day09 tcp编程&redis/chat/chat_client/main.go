package main

import (
	"fmt"
	"github.com/reminance/reminance-go/gohah-go-demo/day09 tcp编程&redis/chat/proto"
	"net"
)

var userId int
var passwd string
var msgChan chan proto.UserRecvMessageReq

func init() {
	msgChan = make(chan proto.UserRecvMessageReq, 1000)
}

func main() {
	//var userId int
	//var passwd string

	fmt.Scanf("%d %s\n", &userId, &passwd)

	conn, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	err = login(conn, userId, passwd)
	if err != nil {
		fmt.Println("login failed, err:", err)
		return
	}

	go processServerMessage(conn)
	for {
		logic(conn)
	}
}
