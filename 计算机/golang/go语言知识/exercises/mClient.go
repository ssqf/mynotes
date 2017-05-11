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
	}
	fmt.Printf("连接：%s成功\n", addr)

	for {
		time.Sleep(time.Second * 1)
		buff := make([]byte, 1024)
		n, ok := conn.Read(buff)
		if ok != nil {
			fmt.Printf("读取：%s 失败  %s\n", addr, ok.Error())
			continue
		}
		fmt.Printf("收到：%v \n", buff[:n])
		conn.Write([]byte{69, 3, 4, 187, 245, 0, 9, 10, 231})
	}
}

func main() {
	port := []string{"51001", "61001", "61002"}
	for _, v := range port {
		go connServ(v)
	}
	for {

	}
}
