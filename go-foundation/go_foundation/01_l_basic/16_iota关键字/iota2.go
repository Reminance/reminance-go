package main

import "fmt"

func main() {
	const (
		A = iota   //0
		B          //1
		C          //2
		D = "haha" //haha   iota=3
		E          //haha,  iota=4
		F = 100    //100, iota=5
		G          //100, iota=5
		H = iota   //7
		I          //8
	)
	const (
		J = iota //0
	)

	fmt.Println(A, B, C, D, E, F, G, H, I, J)
}
