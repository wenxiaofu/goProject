// Code generated by protoc-gen-go. DO NOT EDIT.
// source: attachment.proto

/*
Package com_sangfor_moa_protobuf is a generated protocol buffer package.

It is generated from these files:
	attachment.proto
	auth.proto
	common.proto
	contact.proto
	encrypt.proto
	head.proto
	srvhead.proto
	test.proto
	workreport.proto

It has these top-level messages:
	PB_Attachment
	PB_EntryIp
	PB_EntryIpPool
	PB_AthLoginReq
	PB_AthLoginRsp
	PB_AthAuthReq
	PB_AthAuthRsp
	PB_AthAuthCaptchaReq
	PB_AthAuthCaptchaRsp
	PB_AthFindPasswordReq
	PB_AthFindPasswordDomainInfo
	PB_AthFindPasswordRsp
	PB_AthModifyPasswordReq
	PB_AthModifyPasswordRsp
	PB_AthVerifyTicketReq
	PB_AthVerifyTicketRsp
	PB_AthLogoutReq
	PB_AthLogoutRsp
	PB_AthGetPidsResult
	PB_AthGetPidsReq
	PB_AthGetPidsRsp
	PB_AthVersionUpdateReq
	PB_AthVersionUpdateRsp
	PB_AthWebStatusPush
	PB_AthKickWebReq
	PB_AthKickWebRsp
	PB_AthActiveMsg
	PB_AthCheckPasswordReq
	PB_AthCheckPasswordRsp
	PB_AthPwdForceModifyReq
	PB_AthPwdForceModifyRsp
	PB_AthCaptchaSendReq
	PB_AthCaptchaSendRsp
	PB_AthCancelPush
	PB_AthVersionPush
	PB_AthIpsPush
	PB_MapPostion
	PB_Form
	PB_PhoneNumber
	PB_MailInfo
	PB_PersonExtendData
	PB_Department
	PB_Person
	PB_PersonContral
	PB_Group
	PB_GroupContral
	PB_PersonGroupInfo
	PB_PersonGroupInfoContral
	PB_PersonPersonInfo
	PB_PersonPersonInfoContral
	PB_Domain
	PB_PersonDataStatusDetail
	PB_PersonDataStatus
	PB_DomainSet
	PB_DomainSetContral
	PB_AthCryptReq
	PB_AthCryptRsp
	Head
	PB_ModMsgId
	PB_SrvHead
	MyMsg
	PB_WorkReport
	PB_WRReadStatus
	PB_WRCreateWorkReportReq
	PB_WRCreateWorkReportRsp
	PB_WRGetWorkReportReq
	PB_WRGetWorkReportRsp
	PB_WRGetWorkReportListIdReq
	PB_WRGetWorkReportListIdRsp
	PB_WRGetWorkReportsReq
	PB_WRGetWorkReportsRsp
	PB_WRGetManagementWorkReportsNumReq
	PB_WRGetManagementWorkReportsNumRsp
	PB_WRGetManagementWorkReportListReq
	PB_WRGetManagementWorkReportListRsp
	PB_WRGetManagementGroupWorkReportsNumReq
	PB_WRGetManagementGroupWorkReportsNumRsp
	PB_WRDelWorkReportsReq
	PB_WRDelWorkReportsRsp
	PB_WRModifyWorkReportReq
	PB_WRModifyWorkReportRsp
	PB_PersonStatistics
	PB_WRGetManagementPeriodTimeStatisticsReq
	PB_WRGetManagementPeriodTimeStatisticsRsp
	PB_WRGetManagementPeriodTimeGroupWorkReportsNumReq
	PB_WRGetManagementPeriodTimeGroupWorkReportsNumRsp
	PB_WRGetManagementGroupWorkReportsStatisticsReq
	PB_WRGetManagementGroupWorkReportsStatisticsRsp
	PB_WRGetManagementGroupWorkReportsStatisticsDetailReq
	PB_PersonWrCreateTime
	PB_WRGetManagementGroupWorkReportsStatisticsDetailRsp
	PB_WRGetManagementPeriodTimeDateReq
	PB_WRGetManagementPeriodTimeDateRsp
	PB_WRGetManagementPeriodTimeWorkReportsReq
	PB_WRGetManagementPeriodTimeWorkReportsRsp
	PB_FormModel
	PB_WRSetFormModelReq
	PB_WRSetFormModelRsp
	PB_WRGetFormModelReq
	PB_WRGetFormModelRsp
	PB_WRFormModelPush
	PB_WRWebGetWorkReportsReq
	PB_WRWebGetWorkReportsRsp
	PB_WRWebGetManagementWorkReportsNumReq
	PB_WRWebGetManagementWorkReportsNumRsp
	PB_WRExportWorkReportsReq
	PB_WRExportWorkReportsRsp
	PB_WRWorkReportCountReq
	PB_WRWorkReportCountRsp
	PB_WRSearchReq
	PB_WRSearchRsp
	PB_WRSetAggagretionInfoReq
	PB_WRSetAggagretionInfoRsp
	PB_WRGetAggagretionInfoReq
	PB_WRGetAggagretionInfoRsp
	PB_WRAggagretionPush
	PB_WRNotifyGetListReq
	PB_WRNotifyGetListRsp
	PB_WRNotifySetReplyReq
	PB_WRNotifySetReplyRsp
	PB_WRNotifySetLookReq
	PB_WRNotifySetLookRsp
*/
package com_sangfor_moa_protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PB_Attachment struct {
	Id               *int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type             *int32  `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	Size             *int64  `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
	Value            []byte  `protobuf:"bytes,5,opt,name=value" json:"value,omitempty"`
	Typeinfo         []byte  `protobuf:"bytes,6,opt,name=typeinfo" json:"typeinfo,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PB_Attachment) Reset()                    { *m = PB_Attachment{} }
func (m *PB_Attachment) String() string            { return proto.CompactTextString(m) }
func (*PB_Attachment) ProtoMessage()               {}
func (*PB_Attachment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PB_Attachment) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PB_Attachment) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *PB_Attachment) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *PB_Attachment) GetSize() int64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *PB_Attachment) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PB_Attachment) GetTypeinfo() []byte {
	if m != nil {
		return m.Typeinfo
	}
	return nil
}

func init() {
	proto.RegisterType((*PB_Attachment)(nil), "com.sangfor.moa.protobuf.PB_Attachment")
}

func init() { proto.RegisterFile("attachment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2c, 0x29, 0x49,
	0x4c, 0xce, 0xc8, 0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x48, 0xce,
	0xcf, 0xd5, 0x2b, 0x4e, 0xcc, 0x4b, 0x4f, 0xcb, 0x2f, 0xd2, 0xcb, 0xcd, 0x4f, 0x84, 0x08, 0x27,
	0x95, 0xa6, 0x29, 0xf5, 0x32, 0x72, 0xf1, 0x06, 0x38, 0xc5, 0x3b, 0xc2, 0x75, 0x08, 0xf1, 0x71,
	0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71,
	0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x20, 0xb1,
	0x92, 0xca, 0x82, 0x54, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x30, 0x1b, 0x24, 0x56, 0x9c,
	0x59, 0x95, 0x2a, 0xc1, 0x02, 0xd6, 0x09, 0x66, 0x0b, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94,
	0xa6, 0x4a, 0xb0, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x41, 0x38, 0x42, 0x52, 0x5c, 0x1c, 0x20, 0x1d,
	0x99, 0x79, 0x69, 0xf9, 0x12, 0x6c, 0x60, 0x09, 0x38, 0x1f, 0x10, 0x00, 0x00, 0xff, 0xff, 0x78,
	0xd9, 0xf6, 0x04, 0xbc, 0x00, 0x00, 0x00,
}
