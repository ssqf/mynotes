package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func sendData(conn net.Conn) {
	for {
		fmt.Println("请写入数据:")
		line, _, err := bufio.NewReader(os.Stdin).ReadLine()
		if err != nil {
			log.Fatalln("标准输入读取数据失败", err.Error())
		}
		s := strings.Fields(string(line))
		fmt.Printf("%v\n", s)
		d := make([]byte, len(s))
		for i, v := range s {
			t, _ := strconv.ParseUint(v, 16, 8)
			d[i] = byte(t)
		}
		//line = append(line)
		conn.Write(d)
		//conn.Write([]byte("\n"))
		fmt.Printf("发送数据：%X\n", d)
	}
}
func recvData(conn net.Conn) {
	for {
		fmt.Println("等待读取数据")
		data := make([]byte, 1024)
		//data, err := bufio.NewReader(conn).ReadBytes('\n')
		//data, err := bufio.NewReader(conn).ReadBytes('\n')
		n, err := conn.Read(data)
		if err != nil {
			conn.Close()
			log.Println("读取数据失败", err.Error())
			return
		}
		fmt.Printf("读到的数据:%v %s\n", data[:n], string(data[:n]))
	}

}
func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[Error:]", r)
		}
	}()
	fmt.Println("Service Start!!! 输入监听端口")
	listenPort := ""
	fmt.Scanf("%s", &listenPort)
	log.Println("开始监听：", listenPort)
	l, err := net.Listen("tcp", "0.0.0.0:"+listenPort)
	if err != nil {
		log.Fatalln("监听失败", err.Error())
	}
	conn, err := l.Accept() //接受连接
	if err != nil {
		log.Fatalln("建立连接失败", err.Error())
	}
	log.Printf("连接建立成功\n")
	go sendData(conn)
	go recvData(conn)
	for {
	}

}
