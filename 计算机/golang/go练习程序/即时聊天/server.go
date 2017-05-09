//即时聊天客户端

package main

import (
	"encoding/json"
	"fmt"
	"net"
)

// Users 用户数据
type Users struct {
	id      int
	Name    string
	Pwd     string
	Sex     string
	Age     int
	State   int
	Email   string
	Friends map[int]interface{}
}

// UserList 存储用户列表
var UserList = make(map[int]Users, 500)

// ConnList 当前用户连接列表
var ConnList = make(map[int]Users, 500)

// FindUser 通过id查找用户
func (ul *UserList) FindUser(id int) Users {

}

// AddUser  添加用户
func (ul *UserList) AddUser(user Users) {

}

// DelUser 删除用户
func (ul *UserList) DelUser(id int, pwd string) {

}

// Login 用户登录
func (ul *UserList) Login(id int, pwd string) {

}

// Logout 用户下线
func (ul *UserList) Logout(id int) {

}

// CreateUserID 创建一个用户的唯一ID
func CreateUserID(curID int) int {
	id := curID
	return func() int {
		id++
		return id
	}()
}

// InitServer 初始化服务器
func InitServer() net.Listener {
	listen, err := net.Listen("tcp", ":7908")
	if err != nil {
		panic("监听\":7908\"失败")
	}
	fmt.Println("监听\":7908\"成功")
	return listen
}

func handleConn(l net.Listener) (id int, conn net.Conn) {
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("连接失败:", err.Error())
			continue
		}
		go handleUserConn(conn)
	}

}

func handleUserConn(conn net.Conn) {
	userid := 0
	defer conn.Close()
	buff := make([]byte, 512)
	conn.Write([]byte("你好，请登录或注册"))
	for {
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println("读取数据失败:", err.Error())
			continue
		}
		//解析收到的json数据
		var dat map[string]interface{}
		err = json.Unmarshal(buff[:n], &dat)
		if err != nil {
			fmt.Println("json数据解析失败:", err.Error())
			continue
		}
		switch dat["action"] {
		case "adduser":
		case "login":
		case "logout":
		case "send":
		case "find":
		case "list":
		default:

		}

	}
}
func main() {

}
