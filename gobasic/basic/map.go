// Package basic 01_基础 map
package basic

import "fmt"

func mapType() {
	// 创建一个空的map
	m := make(map[string]int)
	fmt.Println(m) // map[]

	// 添加元素
	m["age"] = 18
	m["score"] = 100

	// 获取元素
	fmt.Println("age:", m["age"])
	fmt.Println("score:", m["score"])

	// 删除元素
	delete(m, "score")
	fmt.Println("删除 score:", m)

	// 检查元素是否存在
	value, exists := m["score"]
	fmt.Println("score exists?", exists, "value:", value)

	// 使用 map 字面量初始化
	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println("m2:", m2)

	// map 的长度
	fmt.Println("m2 长度:", len(m2))

	// 一个会报错的场景
	var m3 map[string]int
	// m3["a"] = 123 // panic: assignment to entry in nil map
	fmt.Println(m3)
}
