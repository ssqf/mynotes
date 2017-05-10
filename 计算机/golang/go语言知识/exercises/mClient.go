package main

import (
	"fmt"
	"net"
	"time"
)

func connServ(addr string) {
	conn, err := net.Dial("tcp", "localhost:"+addr)
	if err != nil {
		fmt.Printf("连接：%s失败  %s\n", addr, err.Error())
		return
	} else {
		fmt.Printf("连接：%s成功\n", addr)
	}

	for {
		buff := make([]byte, 1024)
		n, ok := conn.Read(buff)
		if ok != nil {
			fmt.Printf("读取：%s 失败  %s\n", addr, ok.Error())
		}
		fmt.Printf("收到：%v %s\n", buff[:n], string(buff[:n]))
		conn.Write([]byte(time.Now().Format("2006-01-02 15:04:05") + "  我是" + addr))
		time.Sleep(time.Second * 3)
	}

}

func main() {
	port := []string{"51001", "60001", "61001"}
	for _, v := range port {
		go connServ(v)
	}
	for {

	}
}
