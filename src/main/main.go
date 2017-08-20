package main

import (
	"Cserver"
	"fmt"
	"strconv"
	"runtime"
	"log"
	"net/http"
	"net/rpc"
	"net"
)
//实现RPC，通过注册的RPC函数实现多并发，利用的是单核cpu，多核后面优化
type Num int
func (t *Num) Hi(args int, reply *string) error {
    *reply = "echo:" + strconv.Itoa(args)
	count := args
	for i := 0; i<count; i++{
		go foo(i)
		fmt.Printf("dangqian thread ne %d \n",runtime.NumGoroutine())				
	}
	//确保每一个线程执行完成，但是容错不够，这里还是主协程
	for i := 0; i<count; i++{
		<- ch
	}
	fmt.Printf("%d \n",args)
    return nil
}

var ch chan int = make(chan int)
func foo(id int){
	//sendmassage则是主要的函数实现
	Cserver.Sendmassage(strconv.Itoa(id))
	//当信道=15的时候，结束，防止主线程提前结束
	ch <- id
}
func main(){
	rpc.Register(new(Num))
    rpc.HandleHTTP()
    l, e := net.Listen("tcp", ":1234")
    if e != nil {
        log.Fatal("listen error:", e)
    }
    http.Serve(l, nil)
	
}