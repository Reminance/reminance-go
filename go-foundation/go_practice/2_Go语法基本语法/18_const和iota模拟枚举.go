/**
 * @File : 18_const和iota模拟枚举
 * @Author: Octal_H
 * @Date : 2019/11/13
 * @Desc :
 */

package main

import "fmt"

//声明芯片类型
type ChipType int

func main() {
	//int 定义为 Weapon 类型
	type Weapon int
	//const 下方的常量可以使用 Weapon 作为默认类型。该行使用 iota 进行常量值自动生成，iota 的起始值为 0
	const (
		Arrow Weapon = iota // 开始生成枚举值, 默认为0
		Shuriken
		SniperRifle
		Rifle
		Blower
	)
	// 输出所有枚举值
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
	// 使用枚举类型并赋初值
	var weapon Weapon = Blower
	fmt.Println(weapon)

	fmt.Println("-------------------------")
	//iota 不仅可以生成每次增加 1 的枚举值。还可以利用 iota 来做一些强大的枚举常量值生成器
	//生成标志位常量
	const (
		//iota 使用了一个移位操作，每次将上一次的值左移一位（二进制位），以得出每一位的常量值。
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		FlagBlue
	)
	//将枚举值按二进制格式输出，可以清晰地看到每一位的变化
	//将 3 个枚举按照常量输出，分别输出 2、4、8，都是将 1 每次左移一位的结果。
	fmt.Printf("%d %b %d %d\n", FlagNone, FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("%b %b %b %b\n", FlagNone, FlagRed, FlagGreen, FlagBlue)

	fmt.Println("-------------------------------------------")
	//将枚举值转换为字符串

	fmt.Printf("%s %d\n", CPU, CPU)
	fmt.Printf("%s %d\n", GPU, GPU)
}

const (
	//将 const 里定义的常量值设为 ChipType 类型，且从 0 开始，每行值加 1。
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"

}
