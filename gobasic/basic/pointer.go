// Package basic 01_基础 pointer
package basic

import "fmt"

func pointer() {
	// 创建一个整型变量
	x := 58

	// 创建指向 x 的指针
	var p *int
	p = &x // &x 获取变量 x 的地址

	// 输出指针 p 和它指向的值
	fmt.Println("p:", p)   // 输出指针地址
	fmt.Println("*p:", *p) // 输出指针 p 所指向的值，即 58

	// 修改指针指向的值
	*p = 100
	fmt.Println("x after modification:", x) // 输出修改后的 x，即 100

	// 数组指针示例
	arr := [4]int{1, 2, 3, 4}
	var arrPointer *[4]int
	arrPointer = &arr // &arr 获取数组的地址

	// 输出数组指针的地址和值
	fmt.Println("arrPointer:", arrPointer)       // 输出数组指针的地址
	fmt.Println("arrPointer[0]:", arrPointer[0]) // 输出数组的第一个元素，即 1
}
