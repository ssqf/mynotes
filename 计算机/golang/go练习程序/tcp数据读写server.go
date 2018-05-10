package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"strconv"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:10000")
	if err != nil {
		log.Fatalf("net listen error: % v ", err)
	}
	fmt.Printf("listen tcp 0.0.0.0:10000")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("listen accept error: % v ", err)
			continue
		}
		fmt.Printf("recv client request:%v\n", conn.RemoteAddr)
		go serv(conn)
	}

}

func serv(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	scanner.Split(splitYiDongDiBangA9)
	for scanner.Scan() {
		data := scanner.Bytes()
		if len(data) != 8 {
			log.Fatalf("不为8位:%v\n", data)
			continue
		}
		num, _ := strconv.ParseInt(string(data[:7]), 10, 32)
		pos := int(data[7])
		fmt.Printf("data:%v,num:%d,pos:%d\n", data, num, pos)
		result := float64(num) / math.Pow10(pos)
		fmt.Printf("重量:%f\n", result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

func splitYiDongDiBangA9(data []byte, atEOF bool) (advance int, token []byte, err error) {
	fmt.Printf("recv data:%v\n", data)
	for i := 0; i < len(data); i++ {
		if data[i] == 0x2 { //第一位是0x2
			//return i + 1, data[:i], nil
			if len(data)-i < 11 {
				return i, nil, nil
			}

			if data[i+11] != 0x3 { //第十二位结束是0x3
				continue
			}

			x := xorVerify(data[i : i+8])
			if (x>>4 != data[i+9]) || (x&0xF != data[i+10]) {
				log.Printf("数据校验失败:%x %x %x\n", x, data[i+9], data[i+10])
				return i + 11, nil, nil
			}

			return i + 11, data[i+1 : i+9], nil
		}
	}
	return len(data), nil, nil
}

func xorVerify(dat []byte) byte {
	xor := byte(0x0)
	for _, v := range dat {
		xor = xor ^ v
	}
	return xor
}
