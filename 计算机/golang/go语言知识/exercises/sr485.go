package main

import (
	"fmt"
	"log"
	"net"
)

var ch chan bool
var ch1 chan bool

func main() {
	ch = make(chan bool)
	ch1 = make(chan bool)
	l1, err := net.Listen("tcp", "192.168.10.112:51001")
	if err != nil {
		log.Fatal(err.Error())
	}
	l2, err := net.Listen("tcp", "192.168.10.112:52001")
	if err != nil {
		log.Fatal(err.Error())
	}
	go accept(l1)
	go accept1(l2)
	fmt.Printf("服务器监听开启\n")
	ch1 <- true
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
func accept1(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("建立连接成功1：%v %v\n", l, conn)
		go handleConn1(conn)
	}
}
func handleConn(conn net.Conn) {
	i := byte(0x00)
	for {
		//dat := []byte(time.Now().Format("2006-01-02 15:04:05") + "  我是服务器")
		//n, err := conn.Write(dat)
		//fmt.Println("写")
		if <-ch1 {
			fmt.Printf("发送数据：%x\n", i)
			n, err := conn.Write([]byte{i})
			i++
			if err != nil {
				log.Printf("写错误：%s  %v\n", err.Error(), n)
			}
			ch <- true
		}
	}

}
func handleConn1(conn net.Conn) {
	for {
		//fmt.Println("读")
		//dat := []byte(time.Now().Format("2006-01-02 15:04:05") + "  我是服务器")
		//n, err := conn.Write(dat)
		//time.Sleep(time.Second * 1)
		if <-ch {
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				log.Printf("读错误：%s  %v\n", err.Error(), n)
			}
			fmt.Printf("读到数据：%v %s\n", buff[:n], string(buff[:n]))
			//time.Sleep(time.Second * 5)
			ch1 <- true
		}

	}
}
