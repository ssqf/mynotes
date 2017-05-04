// 结构体

package main

import (
	"fmt"
)

const (
	UnknowSex = iota // UnknowSex 未知的性别
	Male             // Male 男性
	Female           // Female 女性
)

// Person 人类
type Person struct {
	Name    string
	Age     uint8
	Sex     uint8
	Stature uint8
	//Say     func(string)
}

func say(str string) {
	fmt.Println(str)
}
func (p Person) Say(str string) {
	fmt.Println(p.Name + ":" + str)
}
func main() {
	var jack Person = Person{"jack", 18, Male, 170}
	alice := Person{}
	小明 := new(Person)
	小明.Name = "小明"
	//小明.Say = say
	小明.Say("hello")
	jack.Say("你好")
	fmt.Println(*小明)
	fmt.Println(alice)
	fmt.Println(jack)
}
