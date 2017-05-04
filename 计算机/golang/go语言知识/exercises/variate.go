// 变量的练习
//基本数据类型
package main

import "fmt"

func main() {
	var c, d int                 // 定义变量多个变量 有默认值0
	var e, f float64 = 9, 8      // 定义变量并初始化
	var dd int = 5               //定义变量并初始化  提示这个类型可推导
	var b uint32 = 3             // 定义变量 并初始化
	var k = 4                    // 定义变量，不指定类型，golang会自动推导变量类型
	a, b := 5, 8                 // 使用 := 定义变量，自动推导类型, b被重新赋值
	var str string = "hello，世界！" //定义一个字符串
	var ch rune = '走'            //定义一个字符

	x := 10 + 10i // 定义一个复数
	fmt.Println(c, f, k, b, a, e, d, dd)
	fmt.Print(x)
	fmt.Println(x)
	fmt.Println(str)
	fmt.Println(ch) //输出 36208 ‘走’的UTF8的编码
}
