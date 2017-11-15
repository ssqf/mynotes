package main

import "fmt"

func main() {
	s := "我们"
	s[0] = '你'
	b := []byte(s)
	b[0] = 'n'
	fmt.Println(s, b)
}
