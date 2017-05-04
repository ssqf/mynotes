// 区别new 和 make

package main

import (
	"fmt"
)

func main() {
	var n = new(int)
	fmt.Printf("%T %v %v\n", n, n, *n) //输出 *int 0xc042008230 0
}
