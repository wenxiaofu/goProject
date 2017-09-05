package main

import (
	"fmt"
	"log"
	"net/rpc"
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
	//用来接收多少并发,通过RPC发送个加压器server,并发用户数要跟数据库里面已经有的用户数对应，不能超过，超过会out of the index
	//需要处理一下，用户循环使用，可以使用sql批量添加用户，需要写一个脚本
	var args = 4
	var reply string
	//使用call调用
	err = client.Call("Num.Hi", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: =%d\n", reply)
}
