//即时聊天客户端

package main

//用户数据
type Users struct {
	id      int
	Name    string
	Sex     string
	Age     int
	State   int
	Friends map[int]interface{}
}
