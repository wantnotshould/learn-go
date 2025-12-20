// Package basic 01_基础 function
package basic

import "fmt"

func addFunc(x, y int) int {
	return x + y
}

// 变长参数，可以接收多个 int 参数
func sumFunc(prefix string, values ...int) {
	total := 0
	for _, value := range values {
		total += value
	}
	fmt.Println(prefix, total)
}

// 返回多个值
func calculate(x, y int) (int, int) {
	return x + y, x * y
}

// 带命名返回值
func divide(x, y int) (result int, success bool) {
	if y == 0 {
		return 0, false
	}

	// return x / y, true
	result = x / y
	success = true
	return
}

// 返回一个函数
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// 函数作为一个参数
func operate(x int, y int, op func(int, int) int) int {
	return op(x, y)
}

func function() {
	fmt.Printf("addFunc(3, 4) = %d\n", addFunc(3, 4)) // addFunc(3, 4) = 7
	sumFunc("total sum(1~5):", 1, 2, 3, 4, 5)         // total sum(1~5): 15

	sumResult, product := calculate(3, 4)
	fmt.Println("sum:", sumResult, "product:", product) // sum: 7 product: 12

	result, success := divide(10, 2)
	if success {
		fmt.Println("division result:", result) // 输出 division result: 5
	} else {
		fmt.Println("division failed")
	}

	double := multiplier(2)
	fmt.Println("double(5):", double(5)) // 输出 double(5): 10

	sumFunc := func(x, y int) int {
		return x + y
	}
	fmt.Println("operate(3, 4, sumFunc):", operate(3, 4, sumFunc)) // 输出 operate(3, 4, sumFunc): 7
}
