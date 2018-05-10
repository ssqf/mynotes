package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	// Comma-separated list; last entry is empty.
	input := []byte{'i', 0x02, '+', '1', '2', '3', '4', '5', '6', 0x2, 0x2, 0xe, 0x3, 0x5, 0x9, 0x02, '-', '6', '5', '3', '4', '5', '6', 0x1, 0x2, 0x8, 0x3, 0x5, 0x33}
	scanner := bufio.NewScanner(bytes.NewReader(input))
	// Define a split function that separates on commas.
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
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
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
	}
	scanner.Split(onComma)
	// Scan.
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

func xorVerify(dat []byte) byte {
	xor := byte(0x0)
	for _, v := range dat {
		xor = xor ^ v
	}
	return xor
}
