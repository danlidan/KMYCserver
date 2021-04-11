// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: protoId.proto

package msg

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ProtoId int32

const (
	ProtoId_RegisterReqId    ProtoId = 0
	ProtoId_RegisterRspId    ProtoId = 1
	ProtoId_LoginReqId       ProtoId = 2
	ProtoId_LoginRspId       ProtoId = 3
	ProtoId_MatchReqId       ProtoId = 4
	ProtoId_MatchRspId       ProtoId = 5
	ProtoId_MatchCancelReqId ProtoId = 6
	ProtoId_MatchCancelRspId ProtoId = 7
)

// Enum value maps for ProtoId.
var (
	ProtoId_name = map[int32]string{
		0: "RegisterReqId",
		1: "RegisterRspId",
		2: "LoginReqId",
		3: "LoginRspId",
		4: "MatchReqId",
		5: "MatchRspId",
		6: "MatchCancelReqId",
		7: "MatchCancelRspId",
	}
	ProtoId_value = map[string]int32{
		"RegisterReqId":    0,
		"RegisterRspId":    1,
		"LoginReqId":       2,
		"LoginRspId":       3,
		"MatchReqId":       4,
		"MatchRspId":       5,
		"MatchCancelReqId": 6,
		"MatchCancelRspId": 7,
	}
)

func (x ProtoId) Enum() *ProtoId {
	p := new(ProtoId)
	*p = x
	return p
}

func (x ProtoId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtoId) Descriptor() protoreflect.EnumDescriptor {
	return file_protoId_proto_enumTypes[0].Descriptor()
}

func (ProtoId) Type() protoreflect.EnumType {
	return &file_protoId_proto_enumTypes[0]
}

func (x ProtoId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProtoId.Descriptor instead.
func (ProtoId) EnumDescriptor() ([]byte, []int) {
	return file_protoId_proto_rawDescGZIP(), []int{0}
}

type OptionType int32

const (
	OptionType_MoveId   OptionType = 0
	OptionType_ShootId  OptionType = 1
	OptionType_ReloadId OptionType = 2
)

// Enum value maps for OptionType.
var (
	OptionType_name = map[int32]string{
		0: "MoveId",
		1: "ShootId",
		2: "ReloadId",
	}
	OptionType_value = map[string]int32{
		"MoveId":   0,
		"ShootId":  1,
		"ReloadId": 2,
	}
)

func (x OptionType) Enum() *OptionType {
	p := new(OptionType)
	*p = x
	return p
}

func (x OptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_protoId_proto_enumTypes[1].Descriptor()
}

func (OptionType) Type() protoreflect.EnumType {
	return &file_protoId_proto_enumTypes[1]
}

func (x OptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OptionType.Descriptor instead.
func (OptionType) EnumDescriptor() ([]byte, []int) {
	return file_protoId_proto_rawDescGZIP(), []int{1}
}

var File_protoId_proto protoreflect.FileDescriptor

var file_protoId_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x49, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6d, 0x73, 0x67, 0x2a, 0x9b, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x49, 0x64,
	0x12, 0x11, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x49,
	0x64, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x73, 0x70, 0x49, 0x64, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x49, 0x64, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x73, 0x70, 0x49, 0x64, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x49, 0x64, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x73, 0x70, 0x49, 0x64, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x49, 0x64, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x73, 0x70, 0x49, 0x64,
	0x10, 0x07, 0x2a, 0x33, 0x0a, 0x0a, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x6f, 0x76, 0x65, 0x49, 0x64, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x68, 0x6f, 0x6f, 0x74, 0x49, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x6c,
	0x6f, 0x61, 0x64, 0x49, 0x64, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoId_proto_rawDescOnce sync.Once
	file_protoId_proto_rawDescData = file_protoId_proto_rawDesc
)

func file_protoId_proto_rawDescGZIP() []byte {
	file_protoId_proto_rawDescOnce.Do(func() {
		file_protoId_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoId_proto_rawDescData)
	})
	return file_protoId_proto_rawDescData
}

var file_protoId_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protoId_proto_goTypes = []interface{}{
	(ProtoId)(0),    // 0: msg.ProtoId
	(OptionType)(0), // 1: msg.OptionType
}
var file_protoId_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protoId_proto_init() }
func file_protoId_proto_init() {
	if File_protoId_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protoId_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protoId_proto_goTypes,
		DependencyIndexes: file_protoId_proto_depIdxs,
		EnumInfos:         file_protoId_proto_enumTypes,
	}.Build()
	File_protoId_proto = out.File
	file_protoId_proto_rawDesc = nil
	file_protoId_proto_goTypes = nil
	file_protoId_proto_depIdxs = nil
}
