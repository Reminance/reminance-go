/**
 * @File : 16_变量的生命周期
 * @Author: Octal_H
 * @Date : 2019/11/13
 * @Desc :
 */

package main

/*
Go语言的自动垃圾收集器是如何知道一个变量何时可以被回收的呢？
基本的实现思路是，遍历当前运行程序中变量（包括全局变量与局部变量）的指针和引用，
				查看是否可以通过某个指针或引用找到该变量，如果找不到则说明该变量已经不再被使用，可以被自动垃圾收集器回收。
*/
func main() {

}
