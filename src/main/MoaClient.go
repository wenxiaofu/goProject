package main

import (
	"bytes"
	"com.sangfor.moa.protobuf"
	"encoding/binary"
	//	"encoding/json" //转换成json需要
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

type Moamassage struct {
	srv_name string                                  `json:"srv_name"`
	srv_code int32                                   `json:"srv_code"`
	inf_name string                                  `json:"inf_name"`
	inf_code int32                                   `json:"inf_code"`
	pb       com_sangfor_moa_protobuf.PB_AthLoginReq `json:"pb"`
	srv_rsp  string                                  `json:"srv_rsp"`
}

func IntToBytes(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func main() {
	conn, err := net.Dial("tcp", "200.200.169.212:6800")
	checkError(err)
	//defer conn.Close()
	//构建认证pb
	pb_auth := com_sangfor_moa_protobuf.PB_AthLoginReq{
		DomainName:    proto.String("wxf"),
		LoginAccount:  proto.String("12200000001"),
		ClientVersion: proto.String("9.9.9"),
		//		Cdt:           proto.Int64(3),
		//	Did: proto.Int64(406429),
	}

	//构建请求格式
	mm := Moamassage{
		srv_name: "auth",
		srv_code: 2,
		inf_name: "PB_AthLoginReq",
		inf_code: 1,
		pb:       pb_auth,
		srv_rsp:  "PB_AthLoginRsp",
	}

	//构造操作码，用于第三部分的proto文件头
	Srvo := mm.srv_code<<8 | mm.inf_code
	//	fmt.Println(uint32(Srvo))
	heade := com_sangfor_moa_protobuf.PB_SrvHead{
		Srvop: proto.Uint32(uint32(Srvo)), //看能不能减少转换
		//		Session: proto.String("dddd"),
	}

	//构建二进制请求

	//第四部分,思路，pb转换成二进制，传输
	//第四部分也可以通过json构建
	fob, err := proto.Marshal(&pb_auth)

	if err != nil {
		fmt.Println("Error1: ", err)
		return
	}
	//	fmt.Println(fob) //byte[]转换成string 输出

	fobb := []byte(fob)
	//构建第三部分，第三部分是一个json格式的数据转换成二进制
	// tb, errs := json.Marshal(heade) //转换成JSON返回的是byte[]
	// if errs != nil {
	// 	fmt.Println(errs.Error())
	// }
	// fmt.Println(string(tb)) //byte[]转换成string 输出
	// ttb := []byte(tb)
	//第三部分也使用pb
	tb, err := proto.Marshal(&heade)

	if err != nil {
		fmt.Println("Error1: ", err)
		return
	}
	fmt.Println(string(tb)) //byte[]转换成string 输出
	fmt.Println("tb 的长度是 %d", len(tb))
	fmt.Println(tb)
	ttb := []byte(tb)
	//解开pb查看一下
	var umData com_sangfor_moa_protobuf.PB_SrvHead
	err = proto.Unmarshal(tb, &umData)

	if err != nil {
		fmt.Println("Error1: ", err)
		return
	}
	fmt.Println(*umData.Srvop)

	//构建第二部分,取第三部分的长度
	//sb := len(ttb)
	//	fmt.Println("第三部分的长度为: %d", sb)
	//构建第一部分
	//获取总长度
	var sum int
	sum = len(fobb) + len(ttb) + 4 + 16
	//sums := IntToBytes(sum)
	sunb := []byte(IntToBytes(sum))
	//	var arr []byte
	arr := make([]byte, 48) //这一个数组的长度多长没有关系
	//srv_code二进制

	//srv_codes := IntToBytes(int(mm.srv_code)) //这个2就是mm的srv_code,这里有多个转换
	scb := []byte(IntToBytes(int(mm.srv_code)))
	//加入第一部分,f服务码

	copy(arr[0:2], scb[2:4])

	//	arr = append(arr, scb)
	//加入第二部分，包的总长度
	copy(arr[4:], sunb)
	//	arr = append(arr, sunb)
	//测试额外加一部分
	//var addnum int
	addnumaa := []byte(IntToBytes(3))
	copy(arr[8:12], addnumaa[0:4])

	//加入第三部分，服务码
	copy(arr[12:], ttb)
	//arr = append(arr[8:], ttb)
	//加入第四部分
	copy(arr[15:], fobb)
	//arr = append(arr[21:], fobb)

	//最后的结果

	fmt.Println(arr)
	//	conn.Write([]byte(arr))
	//获取服务端的响应值
	//buf := make([]byte, 1024)
	//	conn.Read(buf)
	//构建第二个请求
	pwd := "12345"
	pwdb := []byte(pwd)
	var par com_sangfor_moa_protobuf.PB_AthAuthReq
	//par := com_sangfor_moa_protobuf.PB_AthAuthReq
	//下面肯定有错
	par.LoginAccount = proto.String("12200000001")
	par.EncryptData = pwdb
	fmt.Println(par.EncryptData)
	//先构建第四部分
	PBAthAuthReq, err := proto.Marshal(&par)

	if err != nil {
		fmt.Println("Error1: ", err)
		return
	}
	//	fmt.Println(fob) //byte[]转换成string 输出

	PBAthAuthReqb := []byte(PBAthAuthReq)
	//构建第三部分的pb
	SrvoT := 2<<8 | 3
	//	fmt.Println(uint32(Srvo))
	headeT := com_sangfor_moa_protobuf.PB_SrvHead{
		Srvop: proto.Uint32(uint32(SrvoT)), //看能不能减少转换
		//		Session: proto.String("dddd"),
	}
	//pb转成二进制
	tbt, err := proto.Marshal(&headeT)

	if err != nil {
		fmt.Println("Error1: ", err)
		return
	}
	fmt.Println(string(tbt)) //byte[]转换成string 输出
	fmt.Println("tb 的长度是 %d", len(tbt))
	fmt.Println(tbt)
	ttbt := []byte(tbt)
	//拼接二进制请求
	//获取总长度
	var sumt int
	sumt = len(PBAthAuthReqb) + len(ttbt) + 4 + 16
	//sums := IntToBytes(sum)
	sunbt := []byte(IntToBytes(sumt))
	//	var arr []byte
	arrt := make([]byte, 43) //这一个数组的长度多长没有关系
	//srv_code二进制

	//srv_codes := IntToBytes(int(mm.srv_code)) //这个2就是mm的srv_code,这里有多个转换
	scbt := []byte(IntToBytes(int(mm.srv_code)))
	//加入第一部分,f服务码

	copy(arrt[0:2], scbt[2:4])

	//	arr = append(arr, scb)
	//加入第二部分，包的总长度
	copy(arrt[4:], sunbt)
	//	arr = append(arr, sunb)
	//测试额外加一部分
	//var addnum int
	addnumaat := []byte(IntToBytes(3))
	copy(arrt[8:12], addnumaat[0:4])

	//加入第三部分，服务码
	copy(arrt[12:12+len(ttbt)], ttbt)
	//arr = append(arr[8:], ttb)
	//加入第四部分
	copy(arrt[12+len(ttbt):], PBAthAuthReqb)
	//arr = append(arr[21:], fobb)
	fmt.Println(arrt)
	//	conn.Write([]byte(arrt))
	//获取服务端的响应值
	//buft := make([]byte, 1024)
	//conn.Read(buft)
	//上面竟然意外的成功了，现在模拟创建一个工作汇报
	//构建第四部分
	var cworkreport com_sangfor_moa_protobuf.PB_WRCreateWorkReportReq
	var workreportg com_sangfor_moa_protobuf.PB_WorkReport
	workreportg.Content = proto.String("fromjmeterworkreport here is a test")
	workreportg.Wrdate = proto.Int64(1492704000000)
	workreportg.Type = proto.Int32(0)
	cworkreport.Workreport = &workreportg
	//解码生成二进制文件
	cworkreport_t, err := proto.Marshal(&cworkreport)

	if err != nil {
		fmt.Println("Error1: ", err)
		return
	}
	//	fmt.Println(fob) //byte[]转换成string 输出

	cworkreport_t_b := []byte(cworkreport_t)

	//构建第三部分，可能会有问题，操作码和服务码很大的情况，这个很容易解决
	wSrvo := 57<<8 | 1

	wheade := com_sangfor_moa_protobuf.PB_SrvHead{
		Srvop: proto.Uint32(uint32(wSrvo)), //看能不能减少转换
		//		Session: proto.String("dddd"),
	}
	//转换成二进制
	wheade_t, err := proto.Marshal(&wheade)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	wheade_t_b := []byte(wheade_t)
	//获取总长度
	var wsum int
	wsum = len(cworkreport_t_b) + len(wheade_t_b) + 4 + 16
	//sums := IntToBytes(sum)
	wsunb := []byte(IntToBytes(wsum))
	fmt.Println("数组的总长度是 %d", wsum)
	//	var arr []byte
	//开始构建请求格式
	//arrt_w := make([]byte, 71) 最开始调试以为没关系写死，结果很惨
	arrt_w := make([]byte, wsum) //如果是最后一个请求则没有关系，否则关系很大，后面的请求都会失败，这一个数组的长度多长没有关系
	//srv_code二进制

	//srv_codes := IntToBytes(int(mm.srv_code)) //这个2就是mm的srv_code,这里有多个转换
	scbt_w := []byte(IntToBytes(int(57)))
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
	copy(arrt_w[12:12+len(wheade_t_b)], wheade_t_b)
	//arr = append(arr[8:], ttb)
	//加入第四部分
	//copy(arrt_w[15:], cworkreport_t_b)
	copy(arrt_w[12+len(wheade_t_b):], cworkreport_t_b)
	//arr = append(arr[21:], fobb)
	conn.Write([]byte(arr))
	//获取服务端的响应值
	buf := make([]byte, 1024)
	conn.Read(buf)

	//发送第二个请求
	conn.Write([]byte(arrt))
	//获取服务端的响应值
	buft := make([]byte, 2048)
	conn.Read(buft)

	conn.Write([]byte(arrt_w))
	buft_w_w := make([]byte, 2048)
	conn.Read(buft_w_w)
	fmt.Println(buft_w_w)
	//conn.Write([]byte(arrt_w))
	conn.Write([]byte(arrt_w))
	//conn.Write([]byte(arrt_w))
	//fmt.Println(arrt_w)
	//fmt.Println(arrt_w)
	//获取服务端的响应值
	buft_w := make([]byte, 2048)
	conn.Read(buft_w)
	fmt.Println(buft_w)
	conn.Write([]byte(arrt_w))
	conn.Write([]byte(arrt_w))
	fmt.Println("send msg")
}
