package NetMassage

import (
	"com.sangfor.moa.protobuf"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func Authbyte(user string) []byte {
	pb_auth := com_sangfor_moa_protobuf.PB_AthLoginReq{
		DomainName:    proto.String("wxf"),
		LoginAccount:  proto.String(user),
		ClientVersion: proto.String("9.9.9"),
		//		Cdt:           proto.Int64(3),
		//	Did: proto.Int64(406429),
	}
	return BuildMassage(&pb_auth, 2, 1)

}

func Auth2Pb(user string) []byte {
	pb_auth := com_sangfor_moa_protobuf.PB_AthLoginReq{
		DomainName:    proto.String("wxf"),
		LoginAccount:  proto.String(user),
		ClientVersion: proto.String("9.9.9"),
		//		Cdt:           proto.Int64(3),
		//	Did: proto.Int64(406429),
	}

	return protoMarshal2byte(&pb_auth)
}

func Auth2json(user string) string {
	pb_auth := com_sangfor_moa_protobuf.PB_AthLoginReq{
		DomainName:    proto.String("wxf"),
		LoginAccount:  proto.String(user),
		ClientVersion: proto.String("9.9.9"),
		//		Cdt:           proto.Int64(3),
		//	Did: proto.Int64(406429),
	}
	jsons, errs := json.Marshal(pb_auth) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println(string(jsons)) //byte[]转换成string 输出
	return string(jsons)
}

func LoginMassage(user string, pwd string) []byte {
	pwdb := []byte(pwd)
	var par com_sangfor_moa_protobuf.PB_AthAuthReq
	par.LoginAccount = proto.String(user)
	par.EncryptData = pwdb
	fmt.Println(par.EncryptData)
	return BuildMassage(&par, 2, 3)

}
