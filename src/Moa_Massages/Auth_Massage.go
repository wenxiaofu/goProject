package Moa_Massages

import (
	"NetMassage"
	"com.sangfor.moa.protobuf"
	"github.com/golang/protobuf/proto"
	"net"
)

func form_auth_massage(user string) []byte {
	var pb_auth com_sangfor_moa_protobuf.PB_AthLoginReq
	pb_auth.DomainName = proto.String("wxf")
	pb_auth.LoginAccount = proto.String(user)
	pb_auth.ClientVersion = proto.String("9.9.9")
	return NetMassage.BuildMassage(pb_auth, 2, 1)
}
