// Package basic 01_基础 interface
package basic

import "fmt"

// 基础接口
type animal interface {
	speak() string
	move() string
}

type dog struct{}

// 实际开发中，大都不使用 d dog，一般使用 d *dog
func (d *dog) speak() string {
	return "woof"
}

func (d *dog) move() string {
	return "run"
}

// 空接口
func printValue(i interface{}) {
	// 打印任意类型的值
	fmt.Println(i)
}

// 嵌套接口
type reader interface {
	read(p []byte) (n int, err error)
}

type writer interface {
	write(p []byte) (n int, err error)
}

type readerWriter interface {
	reader
	writer
}

type file struct{}

func (f *file) read(p []byte) (n int, err error) {
	return len(p), nil
}

func (f *file) write(p []byte) (n int, err error) {
	return len(p), nil
}

// 空接口和类型断言
func typeAssert(i interface{}) {
	// 类型断言判断具体类型
	switch v := i.(type) {
	case int:
		fmt.Printf("integer: %d\n", v)
	case string:
		fmt.Printf("string: %s\n", v)
	default:
		fmt.Println("other type")
	}
}

func basicInterface() {
	var d animal = &dog{}
	fmt.Println(d.speak())
	fmt.Println(d.move())
}

func embeddedInterface() {
	var rw readerWriter = &file{}
	data := []byte("Hello, World!")
	rw.read(data)
	rw.write(data)
}

func interfaceType() {
	// 基础接口
	basicInterface()
	// 空接口
	printValue(20)             // int 类型
	printValue("Hello, Go!")   // string 类型
	printValue(3.14)           // float64 类型
	printValue([]int{1, 2, 3}) // 切片类型
	// 嵌套接口
	embeddedInterface()
	// 接口类型断言
	typeAssert(20)           // 输出 integer: 20
	typeAssert("Hello, Go!") // 输出 string: Hello, Go!
	typeAssert(3.14)         // 输出 other type
}
