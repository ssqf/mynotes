package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func handleReq(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/", handleReq)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Printf("serv listen error:%s\n", err.Error())
		return
	}
}
