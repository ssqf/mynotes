package main

//查看web服务性能 在http://localhost:8012/debug/pprof/中打开，
//即可查看Goroutine 堆栈等信息
import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func say(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("input:%v", r)
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", say)
	http.ListenAndServe("localhost:8012", nil)
}
