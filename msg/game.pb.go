// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: game.proto

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

type Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestInfo string `protobuf:"bytes,1,opt,name=testInfo,proto3" json:"testInfo,omitempty"`
}

func (x *Frame) Reset() {
	*x = Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frame) ProtoMessage() {}

func (x *Frame) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frame.ProtoReflect.Descriptor instead.
func (*Frame) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

func (x *Frame) GetTestInfo() string {
	if x != nil {
		return x.TestInfo
	}
	return ""
}

//一个操作
type OptionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId int32 `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"` //玩家id
	OptType  int32 `protobuf:"varint,2,opt,name=optType,proto3" json:"optType,omitempty"`   //操作类型
}

func (x *OptionEvent) Reset() {
	*x = OptionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionEvent) ProtoMessage() {}

func (x *OptionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionEvent.ProtoReflect.Descriptor instead.
func (*OptionEvent) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *OptionEvent) GetPlayerId() int32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *OptionEvent) GetOptType() int32 {
	if x != nil {
		return x.OptType
	}
	return 0
}

//一帧对应的操作
type FrameOpts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrameId int32          `protobuf:"varint,1,opt,name=frameId,proto3" json:"frameId,omitempty"`
	Opts    []*OptionEvent `protobuf:"bytes,2,rep,name=opts,proto3" json:"opts,omitempty"`
}

func (x *FrameOpts) Reset() {
	*x = FrameOpts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrameOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrameOpts) ProtoMessage() {}

func (x *FrameOpts) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrameOpts.ProtoReflect.Descriptor instead.
func (*FrameOpts) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2}
}

func (x *FrameOpts) GetFrameId() int32 {
	if x != nil {
		return x.FrameId
	}
	return 0
}

func (x *FrameOpts) GetOpts() []*OptionEvent {
	if x != nil {
		return x.Opts
	}
	return nil
}

//发送给客户端的同步消息，包含未同步的帧，frameid为最后一帧的id
type LogicFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrameId      int32        `protobuf:"varint,1,opt,name=frameId,proto3" json:"frameId,omitempty"`
	UnsyncFrames []*FrameOpts `protobuf:"bytes,2,rep,name=unsyncFrames,proto3" json:"unsyncFrames,omitempty"`
}

func (x *LogicFrame) Reset() {
	*x = LogicFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogicFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogicFrame) ProtoMessage() {}

func (x *LogicFrame) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogicFrame.ProtoReflect.Descriptor instead.
func (*LogicFrame) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{3}
}

func (x *LogicFrame) GetFrameId() int32 {
	if x != nil {
		return x.FrameId
	}
	return 0
}

func (x *LogicFrame) GetUnsyncFrames() []*FrameOpts {
	if x != nil {
		return x.UnsyncFrames
	}
	return nil
}

//接收来自客户端的同步消息，包含某房间某玩家某一帧的所有操作
type NextFrameOpts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrameId  int32          `protobuf:"varint,1,opt,name=frameId,proto3" json:"frameId,omitempty"`
	RoomId   int32          `protobuf:"varint,2,opt,name=roomId,proto3" json:"roomId,omitempty"`
	PlayerId int32          `protobuf:"varint,3,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Opts     []*OptionEvent `protobuf:"bytes,4,rep,name=opts,proto3" json:"opts,omitempty"`
}

func (x *NextFrameOpts) Reset() {
	*x = NextFrameOpts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextFrameOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextFrameOpts) ProtoMessage() {}

func (x *NextFrameOpts) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextFrameOpts.ProtoReflect.Descriptor instead.
func (*NextFrameOpts) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{4}
}

func (x *NextFrameOpts) GetFrameId() int32 {
	if x != nil {
		return x.FrameId
	}
	return 0
}

func (x *NextFrameOpts) GetRoomId() int32 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *NextFrameOpts) GetPlayerId() int32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *NextFrameOpts) GetOpts() []*OptionEvent {
	if x != nil {
		return x.Opts
	}
	return nil
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x23, 0x0a, 0x05, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x43, 0x0a, 0x0b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4b, 0x0a, 0x09, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x22, 0x5a, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69,
	0x63, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x64,
	0x12, 0x32, 0x0a, 0x0c, 0x75, 0x6e, 0x73, 0x79, 0x6e, 0x63, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x4f, 0x70, 0x74, 0x73, 0x52, 0x0c, 0x75, 0x6e, 0x73, 0x79, 0x6e, 0x63, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x0d, 0x4e, 0x65, 0x78, 0x74, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData = file_game_proto_rawDesc
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_rawDescData)
	})
	return file_game_proto_rawDescData
}

var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_game_proto_goTypes = []interface{}{
	(*Frame)(nil),         // 0: msg.Frame
	(*OptionEvent)(nil),   // 1: msg.OptionEvent
	(*FrameOpts)(nil),     // 2: msg.FrameOpts
	(*LogicFrame)(nil),    // 3: msg.LogicFrame
	(*NextFrameOpts)(nil), // 4: msg.NextFrameOpts
}
var file_game_proto_depIdxs = []int32{
	1, // 0: msg.FrameOpts.opts:type_name -> msg.OptionEvent
	2, // 1: msg.LogicFrame.unsyncFrames:type_name -> msg.FrameOpts
	1, // 2: msg.NextFrameOpts.opts:type_name -> msg.OptionEvent
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frame); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrameOpts); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogicFrame); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextFrameOpts); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_rawDesc = nil
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}
