package NetMassage

import (
	"com.sangfor.moa.protobuf"
	"github.com/golang/protobuf/proto"
)

func CreateW() []byte {
	var cworkreport com_sangfor_moa_protobuf.PB_WRCreateWorkReportReq
	var workreportg com_sangfor_moa_protobuf.PB_WorkReport
	workreportg.Content = proto.String("fromjmeterworkreport here is a test")
	workreportg.Wrdate = proto.Int64(1492704000000)
	workreportg.Type = proto.Int32(0)
	cworkreport.Workreport = &workreportg
	return BuildMassage(&cworkreport, 39, 1)
}
