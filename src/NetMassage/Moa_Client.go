package NetMassage

import (
	"net"
)

//这里可以用一个连接池子，启动的时候自动创建N个连接，需要使用直接获取，使用过后释放
//先写一个简单的，每次都重新创建连接
func Conn2server() net.Conn {
	conn, err := net.Dial("tcp", "200.200.169.212:6800")
	checkError(err)
	return conn
}
