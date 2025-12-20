// Package basic 01_基础 slice
package basic

import (
	"fmt"
	"slices"
)

func createSlice() {
	// 下标，数组/切片的一部分
	// 数组定义（参考array.go）
	arr := [...]int{1, 2, 3, 4, 5}
	slice := arr[0:]
	slice2 := slice[2:]
	fmt.Printf("arr = %v\n", arr)
	fmt.Printf("slice1 := arr[0:], 值：%v\n", slice)
	fmt.Printf("slice2 := slice1[2:], 值：%v\n", slice2)

	// make 关键字
	// make([]T, length, capacity)
	// capacity >= length
	// 可以省略，省略时 length = capacity
	slice3 := make([]int, 3) // 创建一个长度为 3，容量为 3 的 int 切片
	fmt.Println(slice3, len(slice3), cap(slice3))

	// make 创建一个空的映射
	m := make(map[string]int)
	m["a"] = 1
	fmt.Println(len(m))

	// 扩展下 new
	// 作用是分配内存，并返回指向该类型零值的指针
	arr2 := new([100]int) // 创建一个指向 [100]int 类型数组的指针
	slice4 := arr2[0:50]  // 切片引用前 50 个元素
	fmt.Println(slice4)   // 输出 [0, 0, 0, ..., 0]（50 个零）
}

func modifySlice(slice []int) {
	// 通过值传递修改切片中的值
	slice[0] = 100
	fmt.Println("修改后的切片:", slice) // 输出: [100 2 3 4 5]
}

func modifySlicePointer(slice *[]int) {
	// 通过指针传递修改切片中的值
	(*slice)[0] = 200
	fmt.Println("通过指针修改后的切片:", *slice) // 输出: [200 2 3 4 5]
}

func slice() {
	createSlice()

	// 追加元素
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	fmt.Println("切片添加元素后:", slice) // 输出: [1 2 3 4 5]

	// 复制
	slice2 := make([]int, len(slice))
	copy(slice2, slice)
	fmt.Println("复制后的切片:", slice2) // 输出: [1 2 3 4 5]

	// 切片元素检查
	target := 2
	found := false
	for _, v := range slice {
		if v == target {
			found = true
			break
		}
	}

	if found {
		fmt.Println("元素存在:", target) // 输出: 元素存在: 2
	} else {
		fmt.Println("元素不存在")
	}

	// 还有一个方法也可以检查 go1.21+
	fmt.Printf("slices.Contains(%v, %d) = %t\n", slice, target, slices.Contains(slice, target))

	// 切片遍历
	for i, v := range slice {
		fmt.Printf("索引: %d, 元素: %d\n", i, v)
	}

	// 多维切片
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("矩阵:", matrix) // 输出: [[1 2 3] [4 5 6] [7 8 9]]

	// 修改前
	fmt.Printf("slice 修改前 %v\n", slice)
	slice3 := slice[0:]
	fmt.Printf("slice3 %v\n", slice3)
	// 修改切片
	modifySlice(slice3) // 传递切片值
	fmt.Printf("modifySlice(slice3), slice3 = %v\n", slice3)
	fmt.Printf("slice %v\n", slice)

	// 修改切片指针
	modifySlicePointer(&slice3) // 传递切片指针
	fmt.Printf("modifySlicePointer(&slice3), slice3 = %v\n", slice3)
	fmt.Printf("slice %v\n", slice)
}
