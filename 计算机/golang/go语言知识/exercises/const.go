// 常量

package main

import (
	"fmt"
)

const (
	a = iota // iota 从0开始， 之后每行自动加一
	b
	c
	d = 9
	e        //延续上一个变量的值
	f = iota // 继续按行增加
	g
	h
)

func main() {
	const xx int = 4
	//const ddf int  //错误，常量必须有初识值
	const pp, qp = 99, 100
	const (
		aa = 5 //aa := 5 错误 := 只能定义变量
		bb = "ddd"
		cc = 'k'
		dd = xx
	)
	fmt.Println(a, b, c, d, e, f, g, h) // 0 1 2 9 9 5 6 7
	fmt.Println(xx)                     // 4
	fmt.Println(aa, bb, cc, dd)         // 5 ddd 107 4
}
