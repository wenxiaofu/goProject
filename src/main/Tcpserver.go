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

func recvConnMsg(conn net.Conn) {
	//  var buf [50]byte
	buf := make([]byte, 50)

	defer conn.Close()

	for {
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("conn closed")
			return
		}

		// 下面进行解码, 注意参数
		var umData protos.MyMsg
		err = proto.Unmarshal(buf[0:n], &umData)

		if err != nil {
			fmt.Println("Error2: ", err)
			return
		}
		// 输出结果
		fmt.Println(*umData.Id, "  ", *umData.Str, "  ", *umData.Opt)
		//fmt.Println("recv msg:", buf[0:n])
		fmt.Println("recv msg:", string(buf[0:n]))
	}
}

func main() {
	listen_sock, err := net.Listen("tcp", "localhost:10000")
	checkError(err)
	defer listen_sock.Close()

	for {
		new_conn, err := listen_sock.Accept()
		if err != nil {
			continue
		}

		go recvConnMsg(new_conn)
	}

}
