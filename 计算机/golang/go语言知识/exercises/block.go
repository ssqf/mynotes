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
	resp, err := http.Get("http://192.168.10.112")
	if err == nil {
		resp.Body.Close()
	}

}
func c() {
	t := time.NewTicker(time.Second * 30)
	for {
		fmt.Printf("start time:%s\n", time.Now().Format("15:04:05"))
		bb = 0
		for i := 0; i < 1000; i++ {
			go a()
			if i%11 == 0 {
				time.Sleep(time.Millisecond * 300)
				fmt.Printf("i:%d go:%d\n", i, runtime.NumGoroutine())
			}
			//time.Sleep(time.Millisecond * 20)
		}
		<-t.C
		fmt.Printf("over time:%s\n", time.Now().Format("15:04:05"))
	}
}
func main() {
	go c()
	for {
		//time.Sleep(time.Second * 20)
	}
}
