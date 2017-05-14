package main

import (
	"fmt"
	"net"
	"time"
)

func compareSlice(s1 []byte, s2 []byte) bool {
	if len(s1) != len(s1) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}
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
		switch {
		case compareSlice(buff[:n], []byte{1, 3, 0, 16, 0, 1, 133, 207}):
			conn.Write([]byte{177, 3, 2, 9, 96, 255, 230}) //24
		case compareSlice(buff[:n], []byte{1, 3, 0, 17, 0, 1, 212, 15}):
			conn.Write([]byte{83, 3, 2, 3, 32, 0, 160}) //800
		case compareSlice(buff[:n], []byte{1, 3, 0, 0, 0, 2, 196, 11}):
			conn.Write([]byte{69, 3, 4, 187, 245, 0, 9, 10, 231}) //6379.41
		case compareSlice(buff[:n], []byte{1, 3, 0, 8, 0, 2, 69, 201}):
			conn.Write([]byte{138, 3, 4, 187, 9, 0, 0, 53, 221}) //47.881
		}
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
