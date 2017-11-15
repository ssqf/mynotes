package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/astaxie/beego/httplib"
)

func main() {
	req := httplib.Get("https://ii6895liii.com")

	req.SetTimeout(time.Second*3, time.Second*1)
	resp, err := req.Response()
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	s, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp:%s\n", string(s))
}
