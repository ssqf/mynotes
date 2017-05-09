package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func sendData(conn net.Conn) {
	for {
		fmt.Println("请写入数据:")
		line, _, err := bufio.NewReader(os.Stdin).ReadLine()
		if err != nil {
			log.Fatalln("标准输入读取数据失败", err.Error())
		}
		line = append(line, 10)
		conn.Write(line)
		//conn.Write([]byte("\n"))
		fmt.Println("发送数据：", line)
	}

}
func main() {
	fmt.Println("Service Start!!!")
	l, err := net.Listen("tcp", ":7010")
	if err != nil {
		log.Fatalln("监听7010失败", err.Error())
	}
	conn, err := l.Accept()
	if err != nil {
		log.Fatalln("建立连接失败", err.Error())
	}
	go sendData(conn)
	for {
		fmt.Println("等待读取数据")
		data, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			conn.Close()
			log.Fatalln("读取数据失败", err.Error())
		}
		fmt.Println("读到的数据:", data)
	}
}
