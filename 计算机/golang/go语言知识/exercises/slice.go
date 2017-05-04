// 切片相关说明

package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	var slice = []int{5, 4, 3}
	n := copy(slice, array[0:5])
	fmt.Printf("type:%T len:%v cap:%v connect:%v\n", slice, n, cap(slice), slice)
}
