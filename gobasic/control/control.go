// Package control 02_流程控制
package control

import "fmt"

// 打印一个三角形 *
func printTriangleStar(row int) {
	for i := 0; i < row; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

// 打印偶数
func printEvenNumbers() {
	fmt.Println("打印前20个数的偶数：")
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println("")
}

// 类型断言
func typeAssert(a any) {
	switch v := a.(type) {
	case int:
		fmt.Printf("integer : %d\n", v)
	case string:
		fmt.Printf("string: %s\n", v)
	default:
		fmt.Println("other type")
	}
}

func Control() {
	printTriangleStar(10)
	printEvenNumbers()

	typeAssert(20)           // 输出 integer: 20
	typeAssert("Hello, Go!") // 输出 string: Hello, Go!
	typeAssert(3.14)         // 输出 other type
}
