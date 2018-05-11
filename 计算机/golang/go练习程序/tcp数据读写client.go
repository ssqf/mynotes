package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatalf(" dial tcp 127.0.0.1:12345 error:%v ", err)
	}
	writer := bufio.NewWriter(conn)
	for {
		rand.Seed(time.Now().UnixNano())
		num := byte(rand.Int() % 10)
		for i := 0; i < int(num); i++ {
			slat := byte(rand.Uint32()%0xF0) + byte(3)
			writer.WriteByte(slat)
			time.Sleep(time.Duration(rand.Int63()%100) * time.Millisecond)
		}
		data := createData()
		for _, b := range data {
			writer.WriteByte(b)
			time.Sleep(time.Duration(rand.Int63()%100) * time.Millisecond)
		}
		rand.Seed(time.Now().UnixNano())
		num = byte(rand.Int() % 10)
		for i := 0; i < int(num); i++ {
			slat := byte(rand.Uint32()%0xF0) + byte(3)
			writer.WriteByte(slat)
			time.Sleep(time.Duration(rand.Int63()%100) * time.Millisecond)
		}
		writer.Flush()
		time.Sleep(time.Duration(rand.Int63()%5) * time.Second)
	}
}

func createData() [12]byte {
	data := [12]byte{}
	data[0] = 0x2
	data[11] = 0x3
	rand.Seed(time.Now().UnixNano())
	pos := byte(rand.Int() % 6)
	weight := int(rand.Float64() * 100000.0)
	sign := 1
	if rand.Int()%2 == 0 {
		sign = 1
	} else {
		sign = -1
	}
	numStr := fmt.Sprintf("%+07d", weight*sign)
	fmt.Printf("weight:%s\n", numStr)
	numByte := []byte(numStr)
	if len(numByte) != 7 {
		log.Fatal("num !=7")
	}
	for i := 0; i < 7; i++ {
		data[1+i] = numByte[i]
	}
	data[8] = pos
	x := xorVerify(data[1:9])
	data[9] = x >> 4
	data[10] = x & 0xF
	return data
}

func xorVerify(dat []byte) byte {
	xor := byte(0x0)
	for _, v := range dat {
		xor = xor ^ v
	}
	return xor
}
