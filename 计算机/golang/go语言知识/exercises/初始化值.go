package main

import "fmt"

func main() {
	var slice []int
	//slice[0] = 1 //runtime error: index out of range
	slice = append(slice, 1, 2, 3, 4, 5)
	println(slice)
	if slice == nil {
		print("slice == nil\n")
	}

	// var m map[int]int
	// m[1] = 1
	// fmt.Println(m)

	var ch chan int
	ch <- 1
	fmt.Println(<-ch)
}
