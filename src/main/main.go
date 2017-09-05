package main

import (
	"Cserver"
	//"Mysqldb"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"runtime"
	"strconv"
)

//实现RPC，通过注册的RPC函数实现多并发，利用的是单核cpu，多核后面优化
type Num int

func (t *Num) Hi(users int, reply *string) error {
	*reply = "echo:" + strconv.Itoa(users)
	count := users
	//获取数据库中的模板数据，只一次读取
	//先把数据连接放到这，好修改
	db, err := sql.Open("mysql", "moatest:moatest@tcp(200.200.169.215:3306)/moatest?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	//关闭数据库，db会被多个goroutine共享，可以不调用
	defer db.Close()
	//获取jcontext和Pcontext数据
	//jcontext, pcontext := Mysqldb.GetOrgData(db, serv_op, contrl_op)
	//获取到所有的用户名和密码信息
	//useranpwd := Cserver.GetUserbyID(db, "200.200.169.212")
	user, pwd := Cserver.GetUsersByIp(db, "200.200.169.212")

	for i := 0; i < count; i++ {
		//需要发送用户名和密码实现登录，已经业务请求
		go foo(i, user[i], pwd[i])
		//	go Cserver.Send2MoaS(user[i], pwd[i]) //发送业务请求
		fmt.Printf("dangqian thread ne %d \n", runtime.NumGoroutine())
	}
	//确保每一个线程执行完成，但是容错不够，这里还是主协程
	for i := 0; i < count; i++ {
		<-ch
	}
	fmt.Printf("%d \n", users)
	return nil
}

/*
func (t *Num) Hi(args int, reply *string) error {
	*reply = "echo:" + strconv.Itoa(args)
	count := args
	for i := 0; i < count; i++ {
		go foo(i) //发送业务请求
		fmt.Printf("dangqian thread ne %d \n", runtime.NumGoroutine())
	}
	//确保每一个线程执行完成，但是容错不够，这里还是主协程
	for i := 0; i < count; i++ {
		<-ch
	}
	fmt.Printf("%d \n", args)
	return nil
}
*/
var ch chan int = make(chan int)

func foo(id int, user string, pwd string) {
	//sendmassage则是主要的函数实现
	Cserver.Send2MoaS(user, pwd)
	//当信道=15的时候，结束，防止主线程提前结束
	ch <- id
}
func main() {
	rpc.Register(new(Num))
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)

}
