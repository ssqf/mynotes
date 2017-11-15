package main

import "fmt"

func main() {
	data := make([]int, 5, 5)
	fmt.Printf("date addr:%p &data[0]:%p\n", &data, &data[0]) //date addr:0xc04204c3a0 &data[0]:0xc04206e030
	s1 := data[:0]
	s := append(s1, 1)
	fmt.Printf("s addr:%p s:%p &s[0]:%p\n", &s, s, &s[0]) //s addr:0xc04204c3e0 s:0xc04206e030 &s[0]:0xc04206e030
	fmt.Printf("%v,%v\n", data, s)                        //[1 0 0 0 0],[1]
}
