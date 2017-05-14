package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	l1, err := net.Listen("tcp", "localhost:51001")
	if err != nil {
		log.Fatal(err.Error())
	}
	l2, err := net.Listen("tcp", "localhost:52001")
	if err != nil {
		log.Fatal(err.Error())
	}
	go accept(l1)
	go accept(l2)
	for {
	}
}
func accept(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("建立连接成功：%v %v\n", l, conn)
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	for {
		n, err := conn.Write([]byte(time.Now().Format("2006-01-02 15:04:05") + "  我是服务器"))
		if err != nil {
			log.Printf("写错误：%s  %v\n", err.Error(), n)
		}
		buff := make([]byte, 1024)
		n, err = conn.Read(buff)
		if err != nil {
			log.Printf("读错误：%s  %v\n", err.Error(), n)
		}
		fmt.Printf("读到数据：%v %s\n", buff[:n], string(buff[:n]))
		time.Sleep(time.Second * 2)
	}

}
