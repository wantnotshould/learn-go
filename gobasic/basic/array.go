// Package basic 01_基础 array
package basic

import "fmt"

func modifyArray(arr [3]int) {
	arr[0] = 99
}

func modifyArrayPointer(arr *[3]int) {
	arr[0] = 99
}

func array() {
	// 初始化
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3}
	fmt.Println(arr1 == arr2) // true

	// 访问
	fmt.Println(arr1[0]) // 1
	arr1[1] = 10

	// 数组长度
	fmt.Println(arr1) // [1 10 3]

	// 数组传递
	modifyArray(arr1)
	fmt.Println(arr1) // [1 10 3] modifyArray 未修改数组

	// 使用数组指针传递
	modifyArrayPointer(&arr1)
	fmt.Println(arr1) // 输出 [99, 10, 3]，已被修改

	// 长度
	fmt.Printf("len(%v) = %d\n", arr1, len(arr1))

	// 多维数组
	arr3 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr3) // 输出 [[1 2 3] [4 5 6]]

	// 切片初体验
	slice := arr1[1:3]
	slice[1] = 100
	fmt.Println(slice, arr1) // 输出 [10 100] [99 10 100]
}
