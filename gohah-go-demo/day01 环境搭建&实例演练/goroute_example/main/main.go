package main

import(
	"github.com/reminance/reminance-go/gohah-go-demo/day01 环境搭建&实例演练/goroute_example/goroute"
	"fmt"
)


func main() {
	var pipe chan int
	pipe = make(chan int, 1)
	go goroute.Add(100, 300, pipe)

	sum := <- pipe
	fmt.Println("sum=", sum)
}