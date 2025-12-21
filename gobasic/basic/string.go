// Package basic 01_基础 string
package basic

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

// fmt中使用的格式化
// %d 格式化为十进制整数类型
// fmt.Sprintf("%5d", 42)   // 输出："   42"（总宽度5，右对齐，默认填充空格）
// fmt.Sprintf("%-5d", 42)  // 输出："42   "（总宽度5，左对齐，默认填充空格）
// fmt.Sprintf("%05d", 42)  // 输出："00042"（总宽度5，用0填充）
// %b 格式化为二进制
// fmt.Sprintf("%b", 42)    // 输出："101010"（整数42的二进制表示）
// %o 格式化为八进制
// fmt.Sprintf("%o", 42)    // 输出："52"（整数42的八进制表示）
// %x 格式化为十六进制小写
// fmt.Sprintf("%x", 42)    // 输出："2a"（整数42的十六进制表示，使用小写字母）
// %X 格式化为十六进制大写
// fmt.Sprintf("%X", 42)    // 输出："2A"（整数42的十六进制表示，使用大写字母）
// %s 格式化为字符串类型
// fmt.Sprintf("%s", "hello") // 输出："hello"（字符串格式化）
// %f 格式化为浮点数
// fmt.Sprintf("%f", 3.141592) // 输出："3.141593"（默认保留6位小数）
// fmt.Sprintf("%.2f", 3.141592) // 输出："3.14"（保留2位小数）
// %v 通用，对传入值的类型进行默认格式化
// %+v %#v
// fmt.Sprintf("%v", 42)      // 输出："42"（默认格式化为整数）
// fmt.Sprintf("%v", true)    // 输出："true"（默认格式化为布尔值）
// %e 格式化为科学计数法（小写）
// fmt.Sprintf("%e", 12345.6789) // 输出："1.234568e+04"（浮动点数以科学计数法表示）
// %E 格式化为科学计数法（大写）
// fmt.Sprintf("%E", 12345.6789) // 输出："1.234568E+04"（浮动点数以科学计数法表示，使用大写字母E）
// %g 格式化为最简短的浮动点数或科学计数法
// fmt.Sprintf("%g", 12345.6789)  // 输出："12345.6789"（浮动点数，以最简短格式表示）
// fmt.Sprintf("%g", 0.000001234) // 输出："1.234e-06"（浮动点数或科学计数法表示）
// %G 格式化为最简短的浮动点数或科学计数法（使用大写E）
// fmt.Sprintf("%G", 12345.6789)  // 输出："12345.6789"（浮动点数，以最简短格式表示，使用大写E）
// %p 格式化为指针的值
// fmt.Sprintf("%p", &myVar) // 输出类似："0xc000022070"（指针的内存地址）
// %t 格式化为布尔值
// fmt.Sprintf("%t", true)   // 输出："true"
// fmt.Sprintf("%t", false)  // 输出："false"
// %c 格式化为字符，格式化整数值为对应的 Unicode 字符
// fmt.Sprintf("%c", 65)     // 输出："A"（ASCII 65 对应的字符）

// 字符串转数字
func strToNum() {
	// 转为整数
	strNum := "12345"
	num, _ := strconv.Atoi(strNum)
	fmt.Printf("strNum := '%s', 转数字 strconv.Atoi(strNum) = %d\n", strNum, num)

	numInt64, _ := strconv.ParseInt(strNum, 10, 64) // 10 是进制 64 是位数
	fmt.Printf("strNum := '%s', 转数字 strconv.ParseInt(strNum) = %d\n", strNum, numInt64)

	// 有符号最大 2^63-1 9223372036854775807
	errStrNum := "9223372036854775808"
	_, err := strconv.Atoi(errStrNum)
	if err != nil {
		// value out of range
		fmt.Printf("字符串 %s 转数字失败:%v\n", errStrNum, err)

		// 无符号最大打印方式
		fmt.Println("math.MaxUint:", uint64(math.MaxUint))

		fmt.Printf("math.MaxInt: %d, 超过时需要使用 big.Int\n", math.MaxInt)
		bigNum := new(big.Int)
		_, ok := bigNum.SetString(errStrNum, 10)
		if !ok {
			fmt.Println("字符串转大整数失败")
		} else {
			fmt.Println("大整数转换成功:", bigNum)
		}
	}

	// 浮点数
	strFloat := "123.45"
	numFloat, _ := strconv.ParseFloat(strFloat, 64) // 64 是浮点数精度
	// %f 格式化为浮点数
	// %.2f 保留2位小数，会四舍五入
	fmt.Printf("strFloat := '%s', 转浮点数 strconv.ParseFloat(strFloat) = %.1f\n", strFloat, numFloat)
}

func baseString() {
	// 字符串转数字
	// strToNum()

	// 单行
	str := "Hello, Go!"
	// 多行
	multilineStr := `Hello,
	Go!
	Welcome go Go string study.
	`
	fmt.Println(str, multilineStr)

	// 字符串大小写
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))

	// 字符串长度
	cnStr := "你好，世界！hello, world!"
	// len(cnStr) 计算字节数量，而不是字符数
	fmt.Printf("len(%s) = %d\n", cnStr, len(cnStr))

	// Go 字符串是 UTF-8 编码的，因此它支持 Unicode 字符
	// 可以使用 rune 来处理单个字符（而不是单字节字符）
	for _, runeValue := range cnStr {
		fmt.Printf("%c ", runeValue) // 输出每个字符
	}

	fmt.Println("")

	// 字符串索引、切片
	// 会打印相应的ASCII码
	fmt.Printf("cnStr = %s, cnStr[0] = %v\n", cnStr, cnStr[0])

	subStr := cnStr[0:3]
	fmt.Printf("cnStr = %s, cnStr[0:3] = %s\n", cnStr, subStr)

	// 字符串拼接
	str1 := "Hello"
	str2 := "Go"
	str3 := str1 + " " + str2
	fmt.Println("str3 = ", str3)

	// 替换
	// 第三个参数 >0 限制替换的次数 =-1/<0 替换所有 =0 不进行替换
	newStr3 := strings.Replace(str3, "Go", "Golang", 1)
	fmt.Println(newStr3)

	// 连接多个
	github := []string{"github.com", "wantnotshould", "learn-go"}
	githubURL := strings.Join(github, "/")
	fmt.Println("Github:", githubURL)

	// 还可以使用 strings.Replace 中用到的 strings.Builder 拼接
	var b strings.Builder
	b.WriteString("Use")
	b.WriteString(" ")
	b.WriteString("strings.Builder")
	result := b.String()
	fmt.Println(result)

	// rune 切片
	fmt.Printf("rune[str3] = %v\n", []rune(str3)) // [72 101 108 108 111 32 71 111]

	// 字节数组转字符串
	fmt.Println(string([]byte{72, 101, 108, 108, 111, 32, 71, 111})) // 不可以超过255
	fmt.Println(string([]rune{72, 101, 108, 108, 111, 32, 71, 111}))

	// 判断是否包含某个字符串
	fmt.Println(strings.Contains(githubURL, "learn-go"))

	// 查找位置
	fmt.Println(strings.Index(githubURL, "wantnotshould"))

	// 字符串按分隔符分隔
	fruitsStr := "apple,orange,bnana"
	fruits := strings.Split(fruitsStr, ",")
	fmt.Printf("strings.Split(\"%s\", \",\") = %v\n", fruitsStr, fruits)

	// 字符串按空白符分隔
	languageStr := "go php vue "
	language := strings.Fields(languageStr)
	fmt.Println(language)

	// 字符串修剪
	fmt.Println(strings.TrimSpace(" Hello, Go   "))
	fmt.Println(strings.Trim("!!!Hello Go!!!", "!"))

	// 字符串格式化
	name := "Go"
	version := 1.25
	tips := fmt.Sprintf("Welcome to %s version %v", name, version)
	fmt.Println(tips)
}
