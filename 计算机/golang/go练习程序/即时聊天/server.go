//即时聊天客户端

package main

import (
	"bytes"
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
type UserList struct {
	User     Users
	NextUser *UserList
}

// Ulist 全局的用户列表
var Ulist *UserList

// InitUserList 初始化用户列表
func InitUserList() *UserList {
	//从文件中读取用户列表
}

// ShowUserList 显示所有用户
func (ul *UserList) ShowUserList() {
	for ul != nil {
		fmt.Printf("id:%d\tname:%s\tSex:%s\tAge:%d\tState:%d\n", ul.User.id, ul.User.Name, ul.User.Sex, ul.User.Age, ul.User.State)
		ul = ul.NextUser
	}

}

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

	}

}

func handleUserConn(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 512)
	conn.Write(bytes.NewBufferString("你好，请登录或注册").Bytes())
	for {
		n, err := conn.Read(buff)
        if err != nil{
            fmt.Println("读取数据失败:", err.Error())
			continue
        }
	}
}
func main() {

}
