package main

import (
	"net/rpc"
	"log"
	"fmt"
)

  
//func Conrpc(pro,ipport string) (*Client, error) {
	//client, err := rpc.DialHTTP(pro, ipport)
   // if err != nil {
     //   log.Fatal("dialing:", err)
    //}
	//return client
//}


func main() {
	//连接RPC服务器
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	  if err != nil {
        log.Fatal("dialing:", err)
    }
//用来接收多少并发,通过RPC发送个加压器server
    var args = 10000	
    var reply string
	//使用call调用
    err = client.Call("Num.Hi", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    fmt.Printf("Arith: =%d\n", reply)
}