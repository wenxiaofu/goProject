package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"os"
	"protos"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	checkError(err)
	defer conn.Close()
	// 创建一个对象, 并填充字段, 可以使用proto中的类型函数来处理例如Int32(XXX)
	hw := protos.MyMsg{
		Id:  proto.Int32(1),
		Str: proto.String("iyviasbjasdv"),
		Opt: proto.Int32(2),
	}

	// 对数据进行编码, 注意参数是message指针
	mData, err := proto.Marshal(&hw)

	if err != nil {
		fmt.Println("Error1: ", err)
		return
	}

	conn.Write([]byte(mData))
	conn.Write([]byte("Hello world!w43444www\n"))
	conn.Write([]byte("Hello world!wwwwwwwwwwwwwwwwww\n"))

	fmt.Println("send msg")
}
