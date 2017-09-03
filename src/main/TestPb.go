// main.go
package main

import (
	//t "./test"
	"com.sangfor.moa.protobuf"
	"encoding/json" //转换成json需要
	"fmt"
	"github.com/golang/protobuf/proto"
	"reflect"
)

//func ListName(stu interface()) map[string]string
type header struct {
	Encryption  string `json:"encryption"`
	Timestamp   int64  `json:"timestamp"`
	Key         string `json:"key"`
	Partnercode int    `json:"partnercode"`
}

func main() {

	//测试反射

	//转换成JSON字符串
	headerO1 := header{
		Encryption:  "sha",
		Timestamp:   1482463793,
		Key:         "2342874840784a81d4d9e335aaf76260",
		Partnercode: 10025,
	}
	jsons, errs := json.Marshal(headerO1) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println(string(jsons)) //byte[]转换成string 输出

	// 创建一个对象, 并填充字段, 可以使用proto中的类型函数来处理例如Int32(XXX)
	//	var hw proto.Message
	hw := com_sangfor_moa_protobuf.MyMsg{
		Id:  proto.Int32(1),
		Str: proto.String("iyviasbjasdv"),
		Opt: proto.Int32(2),
	}
	fmt.Print("dddddddddd-------%s", reflect.TypeOf(hw))
	//将数据转换成json格式
	jsonM, errs := json.Marshal(hw)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println(string(jsonM)) //byte[]转换成string 输出

	// 对数据进行编码, 注意参数是message指针
	mData, err := proto.Marshal(&hw)

	if err != nil {
		fmt.Println("Error1: ", err)
		return
	}
	fmt.Println(string(mData)) //byte[]转换成string 输出
	// 下面进行解码, 注意参数
	var umData com_sangfor_moa_protobuf.MyMsg
	err = proto.Unmarshal(mData, &umData)

	if err != nil {
		fmt.Println("Error2: ", err)
		return
	}

	// 输出结果
	fmt.Println(*umData.Id, "  ", *umData.Str, "  ", *umData.Opt)
}
