package NetMassage

import (
	"bytes"
	"com.sangfor.moa.protobuf"
	"encoding/binary"
	//	"encoding/json" //转换成json需要
	"fmt"
	"github.com/golang/protobuf/proto"
	//"strconv"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

//不知道是否需要这个转换，Marshal好像直接返回的就是二进制,pb里面提供了转json的方法
func protoMarshal2byte(pb proto.Message) []byte {
	cworkreport_t, _ := proto.Marshal(pb)

	//	fmt.Println(fob) //byte[]转换成string 输出

	return []byte(cworkreport_t)

}

//将pb转换成json数据
func proto2json() {

}

//int2bytes
func IntToBytes(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

//16进制转10进制,现在最大是两位，直接转换
func sixteen2t(num int32) int32 {
	return num/10*16 + num%10
}

//生成操作码的pb2byte
func form_pb_serop_b(modelop, contrlop int32) []byte {
	wheade := com_sangfor_moa_protobuf.PB_SrvHead{
		Srvop: proto.Uint32(uint32(sixteen2t(modelop)<<8 | contrlop)), //看能不能减少转换
		//		Session: proto.String("dddd"),
	}
	return protoMarshal2byte(&wheade)
}

func BuildMassage(pb_context proto.Message, modelop int32, contrlop int32) []byte {
	//解码生成二进制文件
	pb_context_b := protoMarshal2byte(pb_context)
	//构建第三部分，可能会有问题，操作码和服务码很大的情况，这个很容易解决
	pb_serop_b := form_pb_serop_b(modelop, contrlop)
	//转换成二进制

	//获取总长度
	wsum := len(pb_context_b) + len(pb_serop_b) + 4 + 16
	wsunb := []byte(IntToBytes(wsum))
	fmt.Println("数组的总长度是 %d", wsum)
	//	var arr []byte
	//开始构建请求格式
	//arrt_w := make([]byte, 71) 最开始调试以为没关系写死，结果很惨
	arrt_w := make([]byte, wsum) //如果是最后一个请求则没有关系，否则关系很大，后面的请求都会失败，这一个数组的长度多长没有关系
	//srv_code二进制

	//srv_codes := IntToBytes(int(mm.srv_code)) //这个2就是mm的srv_code,这里有多个转换,后面再优化，这一步性能可能很差
	scbt_w := []byte(IntToBytes(int(sixteen2t(modelop))))
	//加入第一部分,f服务码
	fmt.Println(scbt_w)
	copy(arrt_w[0:2], scbt_w[2:4])
	//copy(arrt_w[0:4], scbt_w)
	//	arr = append(arr, scb)
	//加入第二部分，包的总长度
	copy(arrt_w[4:], wsunb)
	//	arr = append(arr, sunb)
	//测试额外加一部分
	//var addnum int
	addnumaatt := []byte(IntToBytes(3))
	copy(arrt_w[8:12], addnumaatt[0:4])

	//加入第三部分，服务码
	copy(arrt_w[12:12+len(pb_serop_b)], pb_serop_b)
	//arr = append(arr[8:], ttb)
	//加入第四部分
	//copy(arrt_w[15:], cworkreport_t_b)
	copy(arrt_w[12+len(pb_context_b):], pb_context_b)
	return arrt_w
}
