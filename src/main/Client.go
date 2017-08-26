package main

import (
	"fmt"
	"net"
)

func main() {
	currency := 20 //并发数,记住，一个连接数是打开一个端口号，window和linux的端口号都是有限制的
	count := 10    //每条连接发送多少次连接
	for i := 0; i < currency; i++ {
		go func() {
			for j := 0; j < count; j++ {
				sendMessage()
			}
		}()
	}
	select {}
}

func sendMessage() {
	conn, err := net.Dial("tcp", "200.200.169.212:6800")
	if err != nil {
		panic("error")
	}
	header := "GET / HTTP/1.0\r\n\r\n"
	fmt.Fprintf(conn, header)
}
