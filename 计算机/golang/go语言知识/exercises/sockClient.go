package main

import (
	"net"
	"log"
	"fmt"
)

func main(){
	conn,err := net.Dial("tcp","localhost:5023")
	if err != nil{
		log.Fatal(err.Error())
	}
	buff := make([]byte,100)
	for{
		n,err:=conn.Read(buff)
		if err !=nil{
			log.Fatal(err.Error())
		}
		fmt.Printf("recv:%s\n",string(buff[:n]))
	}
}