// Package basic 01_基础 operators
package basic

import "fmt"

func arithmeticOperator() {
	fmt.Println("算术运算符：")
	a := 10
	b := 3
	fmt.Println(a + b) // 加法
	fmt.Println(a - b) // 减法
	fmt.Println(a * b) // 乘法
	fmt.Println(a / b) // 除法
	fmt.Println(a % b) // 取余
}

func comparisonOperator() {
	fmt.Println("\n关系运算符：")
	a := 10
	b := 5
	fmt.Println(a == b) // 等于
	fmt.Println(a != b) // 不等于
	fmt.Println(a > b)  // 大于
	fmt.Println(a < b)  // 小于
	fmt.Println(a >= b) // 大于或等于
	fmt.Println(a <= b) // 小于或等于
}

func logicalOperator() {
	fmt.Println("\n逻辑运算符：")
	a := true
	b := false
	fmt.Println(a && b) // 与（AND）
	fmt.Println(a || b) // 或（OR）
	fmt.Println(!a)     // 非（NOT）
}

func bitwiseOperator() {
	fmt.Println("\n位运算符：")
	a := 5              // 二进制：0101
	b := 3              // 二进制：0011
	fmt.Println(a & b)  // 输出 1  (0101 & 0011 = 0001)
	fmt.Println(a | b)  // 输出 7  (0101 | 0011 = 0111)
	fmt.Println(a ^ b)  // 输出 6  (0101 ^ 0011 = 0110)
	fmt.Println(a &^ b) // 输出 4  (0101 &^ 0011 = 0100)

	fmt.Println(a << 1) // 输出 10  (左移一位，0101 << 1 = 1010)
	fmt.Println(a >> 1) // 输出 2   (右移一位，0101 >> 1 = 0010)
}

func assignmentOperator() {
	fmt.Println("\n赋值运算符：")
	a := 5
	a += 3         // a = a + 3
	fmt.Println(a) // 输出 8

	a *= 2         // a = a * 2
	fmt.Println(a) // 输出 16

	a >>= 2        // a = a >> 2 (右移 2 位)
	fmt.Println(a) // 输出 4
}

func otherOperator() {
	fmt.Println("\n其他运算符：")
	a := 5
	a++ // 自增
	fmt.Println(a)

	b := 10
	c := 20
	b, c = c, b // 多重赋值
	fmt.Println(b, c)
}

func conditionalOperator() {
	fmt.Println("\n条件运算符：")
	a := 10
	result := 0
	if a > 5 {
		result = 1
	} else {
		result = 0
	}
	fmt.Println(result)
}

func pointerOperator() {
	fmt.Println("\n指针运算符：")
	a := 10
	p := &a         // p 是指向 a 的指针
	fmt.Println(p)  // 输出 a 的地址
	fmt.Println(*p) // 输出 a 的值，即 10
	*p = 20         // 修改 a 的值，通过指针 p
	fmt.Println(a)  // 输出 20
}

func operators() {
	arithmeticOperator()
	comparisonOperator()
	logicalOperator()
	bitwiseOperator()
	assignmentOperator()
	otherOperator()
	conditionalOperator()
	pointerOperator()
}
