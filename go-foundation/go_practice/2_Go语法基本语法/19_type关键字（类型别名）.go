/**
 * @File : 19_type关键字（类型别名）
 * @Author: Octal_H
 * @Date : 2019/11/13
 * @Desc :
 */

package main

import (
	"fmt"
	"reflect"
)

// 将NewInt定义为int类型
//将 NewInt 定义为 int 类型，这是常见的定义类型的方法，通过 type 关键字的定义，NewInt 会形成一种新的类型，NewInt 本身依然具备 int 类型的特性。
type NewInt int

//将int取一个别名叫IntAlias,将 IntAlias 设置为 int 的一个别名，使用 IntAlias 与 int 等效。
type IntAlias = int

func main() {

	/*
		区分类型别名与类型定义
	*/
	// 将a声明为NewInt类型
	var a NewInt
	// 查看a的类型名
	fmt.Printf("a type:%T\n", a) //main.NewInt

	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名
	fmt.Printf("a type:%T\n", a2) //int

	fmt.Println("----------------------------")
	/*
		在结构体成员嵌入时使用别名
	*/
	// 声明变量b为车辆类型
	var b Vehicle

	// 指定调用FakeBrand的Show
	b.FakeBrand.show()
	// 取a的类型反射对象
	ta := reflect.TypeOf(b)
	// 遍历a的所有成员
	for i := 0; i < ta.NumField(); i++ {
		// a的成员信息
		f := ta.Field(i)
		// 打印成员的字段名和类型
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.Name())
	}

}

// 定义商标结构
type Brand struct {
}

// 为商标结构添加Show()方法
func (t Brand) show() {

}

// 为Brand定义一个别名FakeBrand
type FakeBrand = Brand

// 定义车辆结构
type Vehicle struct {
	// 嵌入两个结构
	FakeBrand
	Brand
}

// 定义time.Duration的别名为MyDuration
//type MyDuration = time.Duration
//// 为MyDuration添加一个函数
//func (m MyDuration) EasySet(a string) {		//cannot define new methods on non-local type time.Duration
//}
