package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func a() {
	b()
}

var bb = 0

func b() {
	fmt.Printf("b:%d\n", bb)
	bb++
	resp, err := http.Get("http://www.baidu.com")
	if err == nil {
		resp.Body.Close()
	}

}

func main() {
	for {
		for {
		go b()
		time.Sleep(time.Millisecond * 10)
		fmt.Printf("go:%d\n", runtime.NumGoroutine())
	}
}
