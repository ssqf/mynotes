package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	l1, err := net.Listen("tcp", "192.168.10.112:51001")
	if err != nil {
		log.Fatal(err.Error())
	}
	l2, err := net.Listen("tcp", "192.168.10.112:52001")
	if err != nil {
		log.Fatal(err.Error())
	}
	go accept(l1)
	go accept(l2)
	fmt.Printf("服务器监听开启\n")
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
	i := byte(0x00)
	for {
		//dat := []byte(time.Now().Format("2006-01-02 15:04:05") + "  我是服务器")
		//n, err := conn.Write(dat)
		n, err := conn.Write([]byte{i})
		i++
		fmt.Printf("发送数据：%x\n", i)
		if err != nil {
			log.Printf("写错误：%s  %v\n", err.Error(), n)
		}
		//time.Sleep(time.Second * 1)
		buff := make([]byte, 1024)
		n, err = conn.Read(buff)
		if err != nil {
			log.Printf("读错误：%s  %v\n", err.Error(), n)
		}
		fmt.Printf("读到数据：%v %s\n", buff[:n], string(buff[:n]))
		time.Sleep(time.Second * 10)
	}

}
