package main


import(
	"github.com/reminance/reminance-go/gohah-go-demo/day01 环境搭建&实例演练/package_example/calc"
	"fmt"
)

func main() {
	sum := calc.Add(100, 300)
	sub := calc.Sub(100, 300)

	fmt.Println("sum=",sum)
	fmt.Println("sub=", sub)
}