/**
 * @File : 31_reflect对象设置实际变量的值2
 * @Author: Octal_H
 * @Date : 2019/11/12
 * @Desc :
 */

package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string
	Age    int
	School string
}

func main() {
	s1 := Student{"王二狗", 19, "宁大"}
	//通过反射,更改对象的数值,前提是数据可以被更改
	fmt.Printf("%T\n", s1) //main.Student
	p1 := &s1
	fmt.Printf("%T\n", p1) //*main.Student
	fmt.Println(s1.Name)
	fmt.Println((*p1).Name, p1.Name)

	//改变数值
	value := reflect.ValueOf(&s1)
	//判断是不是指针,如果是,继续修改
	if value.Kind() == reflect.Ptr {
		newValue := value.Elem()
		fmt.Println(newValue.CanSet()) //true

		f1 := newValue.FieldByName("Name")
		f1.SetString("王三")
		f3 := newValue.FieldByName("School")
		f3.SetString("中北")
		fmt.Println(s1)

	}

}
