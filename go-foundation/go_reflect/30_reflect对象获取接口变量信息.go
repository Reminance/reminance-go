/**
 * @File : 30_reflect对象获取接口变量信息
 * @Author: Octal_H
 * @Date : 2019/11/12
 * @Desc :
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.23
	// "接口类型变量"-->"反射类型对象"
	value := reflect.ValueOf(num)
	fmt.Printf("%T,%v\n", value, value) //reflect.Value,1.23

	// "反射类型对象" --> "接口类型变量"
	convertValue := value.Interface().(float64)
	fmt.Printf("%T,%v\n", convertValue, convertValue) //float64,1.23

	/*
		反射类型对象-- >接口类型变量，理解为"强制转换"
		Golang对类型要求非常严格，类型一定要完全符合
		一个是*float64，一个float64，如果弄混，直接panic
	*/
	pointer := reflect.ValueOf(&num)
	convertPointer := pointer.Interface().(*float64)
	fmt.Printf("%T,%v\n", convertPointer, convertPointer) //*float64,0xc00000a0b8

}
