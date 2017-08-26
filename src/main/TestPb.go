// main.go
package main

import (
	//t "./test"
	"fmt"
	"github.com/golang/protobuf/proto"
	"protos"
)

func main() {
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

	// 下面进行解码, 注意参数
	var umData protos.MyMsg
	err = proto.Unmarshal(mData, &umData)

	if err != nil {
		fmt.Println("Error2: ", err)
		return
	}

	// 输出结果
	fmt.Println(*umData.Id, "  ", *umData.Str, "  ", *umData.Opt)
}
