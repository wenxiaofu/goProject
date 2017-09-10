package NetMassage

import (
	"com.sangfor.moa.protobuf"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func Get_auth_massage(json_s string) proto.Message {
	var pb_auth com_sangfor_moa_protobuf.PB_AthLoginReq
	//先将json转换成go对象

	err := json.Unmarshal([]byte(json_s), &pb_auth)
	if err != nil {
		fmt.Println(err)
	}
	return &pb_auth

}

//增加请求路由，用于解析json数据，这样前端就不用自己
func Get_massage(modelop int32, contrlop int32, json_s string) []byte {
	switch modelop {
	case 2:
		//认证模块
		switch contrlop {
		case 1:
			return BuildMassage(Get_auth_massage(json_s), modelop, contrlop)
		default:
			fmt.Printf("Default")
			return nil
		}
	case 3:
		switch contrlop {
		case 1:
			fmt.Println("12223")
			return nil
		}

	default:
		fmt.Printf("Default")
		return nil
	}
	return nil
}
