package main

import (
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:5023")
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	for {
		n, err := conn.Write([]byte(time.Now().Format("2006-01-02 15:04:05") + "  我是服务器"))
		if err != nil {
			log.Printf("写错误：%s  %v\n", err.Error(), n)
		}
		time.Sleep(time.Second * 2)
	}

}
