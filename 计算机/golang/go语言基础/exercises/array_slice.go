//数组和切片

package main

import (
	"fmt"
)

func main() {
	var xx int// 错误
	var arr [5]int         //定义数组
	var arr2 [3][5]float64 //定义二维数组
	arr3 := [4]int{1,2,3}
	var arr4 [2][3]int = [2][3]int{{1,2,3},{4,5,6}}
	xx = 6
	fmt.Println(xx)
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
}
