package main

import "fmt"

const pi = 3.1415926

var r = 2

func say(str string) {
	fmt.Printf("str addr:%p\n", &str)
}

func main() {
	p := "hello"
	fmt.Printf("say addr:%p\n", say)
	fmt.Printf("p addr:%p,PI addr:%p,r addr:%p\n", &p, pi, &r)
}

// 常量不能获取地址，常量在常量区
// 每次执行他们的地址都一样，程序在编译时变量的地址就确定了，在操作系统上的进程都是使用虚拟内存
