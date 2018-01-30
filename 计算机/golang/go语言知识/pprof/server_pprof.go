package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	//访问http://localhost:6060/debug/pprof/ 即可查看Goroutine 和堆栈信息
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	addr := "127.0.0.1:8756"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen failed:%v\n", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept err:%v\n", err)
			continue
		}
		go doService(conn)
	}
}

func doService(conn net.Conn) {
	for {
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Printf("conn read error:%v\n", err)
			return
		}
		fmt.Printf("read connent:\n%s\n", string(buff[0:n]))
	}
}
