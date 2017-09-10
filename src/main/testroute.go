package main

import (
	"NetMassage"
	"com.sangfor.moa.protobuf"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	pb_client_type := com_sangfor_moa_protobuf.PB_ClientDevType_PBCDT_SERVER
	//pb_client_type = com_sangfor_moa_protobuf.PB_ClientDevType_value["PBCDT_SERVER"]
	pb_auth := com_sangfor_moa_protobuf.PB_AthLoginReq{
		DomainName:    proto.String("wxf"),
		LoginAccount:  proto.String("12200000001"),
		ClientVersion: proto.String("9.9.9"),
	}
	pb_auth.Cdt = &pb_client_type
	//生成json
	jsonM, _ := json.Marshal(pb_auth)
	fmt.Println(NetMassage.Get_massage(2, 1, string(jsonM)))

	fmt.Println(NetMassage.Authbyte("12200000001"))

}
