// Package basic 01_基础 struct
package basic

import (
	"fmt"
	"reflect"
	"strings"
)

func structTypes() {
	type profile struct {
		Github string `json:"github" description:"github"`
		Name   string `json:"name" description:"name" custom:"自定义信息" binding:"required,max=26"`
	}

	myProfile := profile{
		Github: "github.com/wantnotshould",
		Name:   "Perry He",
	}

	fmt.Printf("%v\n", myProfile)  // {github.com/wantnotshould Perry He}
	fmt.Printf("%+v\n", myProfile) // {Github:github.com/wantnotshould Name:Perry He}
	fmt.Printf("%#v\n", myProfile) // basic.profile{Github:"github.com/wantnotshould", Name:"Perry He"}

	// 获取下字段类型后 `` 中的信息
	// 获取类型
	t := reflect.TypeFor[profile]() // go1.22+
	v := reflect.ValueOf(myProfile)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) // 获取字段信息
		fmt.Printf("字段名: %s, tag: %s\n", field.Name, field.Tag)

		// tag信息
		if jsonTag, ok := field.Tag.Lookup("json"); ok {
			fmt.Printf("  -> json tag: %s\n", jsonTag)
		}

		if descriptionTag, ok := field.Tag.Lookup("description"); ok {
			fmt.Printf("  -> description tag: %s\n", descriptionTag)
		}

		if customTag, ok := field.Tag.Lookup("custom"); ok {
			fmt.Printf("  -> custom tag: %s\n", customTag)
		}

		tagStr := string(field.Tag)
		tagArr := strings.Fields(tagStr)
		fmt.Println(tagArr)

		// 如果自定义写一些自定义验证时，可以使用.Get方法获取，然后Split拆分下
		// 获取字段的值
		fmt.Println("Type:", v.Type()) // 输出结构体类型
		fmt.Println("Kind:", v.Kind()) // 输出结构体的 Kind
		fieldVal := v.Field(i)
		fmt.Printf("fieldVal: %+v\n", fieldVal) // fieldVal: Perry He
		bindingTag := field.Tag.Get("binding")
		fmt.Println(bindingTag) // required,max=26

		// 一些简单的验证操作可以扩展实现
		bindingTagArr := strings.SplitSeq(bindingTag, ",")
		for valid := range bindingTagArr {
			if valid == "required" {
				if fieldVal.Type().String() == "" {
					fmt.Println("a ...any")
				}
			} else {
				validArr := strings.Split(valid, "=")
				if len(validArr) == 2 {
					fmt.Println("validArr[0]:", validArr[0])
					fmt.Println("validArr[0]:", validArr[1])
				}
			}
		}
	}
}
