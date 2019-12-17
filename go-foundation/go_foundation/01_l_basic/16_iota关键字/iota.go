package main

import "fmt"

func main() {
	/*
		iota：特殊常量，可以认为是一个可以被编译器修改的常量
			每当定义一个const，iota的初始值为0
			每当定义一个常量，就会自动累加1
			直到下一个const出现，清零
	*/
	const (
		a = iota //0
		b = iota //1
		c = iota //2
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	const (
		d = iota //0
		e        //1
	)
	fmt.Println(d)
	fmt.Println(e)

	//枚举中
	const (
		MALE = iota
		FEMALE
		UNKNOW
	)
	fmt.Println(MALE, FEMALE, UNKNOW)

}
