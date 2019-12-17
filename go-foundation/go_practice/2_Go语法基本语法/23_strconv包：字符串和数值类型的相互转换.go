/**
 * @File : 23_strconv包：字符串和数值类型的相互转换
 * @Author: Octal_H
 * @Date : 2019/11/13
 * @Desc :
 */

package main

import (
	"fmt"
	"strconv"
)

/*
strconv 包里提供了很多函数，大概分为以下几类：
	字符串转 int：Atoi()；
	int 转字符串：Itoa()；
	Parse 类函数将 string 转换为指定的数值类型：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。
因为 string 类型转数值类型可能会失败，所以这些函数都有第二个返回值表示转换是否成功；
	Format 类函数将指定的数值类型转 string 类型：FormatBool()、FormatFloat()、FormatInt()、FormatUint()；
	Append 类函数用于将指定的数值类型转换成字符串后追加到一个切片中：AppendBool()、AppendFloat()、AppendInt()、AppendUint()。
*/

func main() {
	/*
		string 和 int 的转换
	*/
	//1) int 转换为字符串：Itoa()
	fmt.Println("a" + strconv.Itoa(32)) //a32

	//2) string 转换为 int：Atoi()
	i, _ := strconv.Atoi("3")
	fmt.Println(3 + i) //6
	//转换失败
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println("err:", err) //strconv.Atoi: parsing "a": invalid syntax
	}

	/*
		Parse 类函数
			Parse 类函数用于将字符串转换为指定的数值类型，其中包含 ParseBool()、ParseFloat()、ParseInt()、ParseUint() 等方法。
		参数:
			s 表示需要转换的字符串
			base 表示以什么进制的方式去解析给定的字符串，有效值为 0 和 2 ~ 36。
				当 base=0 的时候，表示根据 string 的前缀来判断以什么进制去解析：0x 开头表示以 16 进制的方式去解析，
				0 开头的以 8 进制方式去解析，其它的以 10 进制方式解析。
			bitSize 表示转换为多少位的 int 或 uint，有效值为 0、8、16、32、64，当值为 0 的时候，表示转换为 int 或 uint 类型。
				例如 bitSize=8 表示转换后的值的类型为 int8 或 uint8。
	*/
	b, err := strconv.ParseBool("bool")
	f, err := strconv.ParseFloat("3.1414", 64)
	ii, err := strconv.ParseInt("-42", 10, 64)
	u, err := strconv.ParseUint("42", 10, 32)

	fmt.Printf("%T, %v\n", b, b)
	fmt.Printf("%T, %v\n", f, f)
	fmt.Printf("%T, %v\n", ii, ii)
	fmt.Printf("%T, %v\n", u, u)

	fmt.Println("--------------------------")
	/*
		Format 类函数
			将指定数值类型转换为 string 类型，其中包含 FormatBool()、FormatFloat()、FormatInt()、FormatUint() 等方法。
		参数:
		bitSize 表示 f 的来源类型（32 指 float32、64 指 float64），会据此进行四舍五入。
		fmt 表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则就用'f'格式）、'G'（指数很大时用'E'格式，否则就用'f'格式）。
		prec 控制精度（排除指数部分）：对 'f'、'e'、'E'，它表示小数点后的数字个数，对'g'、'G'，它控制总的数字个数。如果 prec 为 -1，则代表使用最少但又必需的数字来表示 f。
	*/
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-42, 16)
	s4 := strconv.FormatUint(42, 16)

	fmt.Printf("%T, %v\n", s1, s1)
	fmt.Printf("%T, %v\n", s2, s2)
	fmt.Printf("%T, %v\n", s3, s3)
	fmt.Printf("%T, %v\n", s4, s4)

	/*
		Append 类函数
			Append 类函数用于将指定类型转换成字符串后追加到一个切片中，其中包含 AppendBool()、AppendFloat()、AppendInt()、AppendUint() 等方法。
	*/
	//声明一个slice
	b10 := []byte("int (base 10):")

	//将转化为10进制的string,追加到slice中
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10)) //-42

	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16)) //-2a
}
