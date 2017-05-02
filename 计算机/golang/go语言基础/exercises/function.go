package main

import (
	"fmt"
	"strings"
)

// 创建一个添加后缀的函数
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	fmt.Println(addBmp("filename"))   //输出filename.bmp
	fmt.Println(addJpeg("filename1")) //输出filename1.jpeg
}
