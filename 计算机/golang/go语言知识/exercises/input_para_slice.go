package main

import "fmt"

func testSlice(s []int) {
	//s = append(s, 1, 2)
	s[7] = 7
}

func main() {
	//slice := []int{}
	slice := make([]int, 5, 10)

	slice[0] = 1
	testSlice(slice)
	fmt.Println(slice)
}

/*
slice 还是传值进去
*/
