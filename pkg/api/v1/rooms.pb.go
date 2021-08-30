// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/v1/rooms.proto

package v1

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The ping request.
type UpdateRoomWithPingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Target scheduler name.
	// NOTE: On http protocol, this operates as a path param.
	SchedulerName string `protobuf:"bytes,1,opt,name=scheduler_name,json=schedulerName,proto3" json:"scheduler_name,omitempty"`
	// Target room name.
	// NOTE: On http protocol, this operates as a path param.
	RoomName string `protobuf:"bytes,2,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	// Target room metadata.
	Metadata *_struct.Struct `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Indicates the room status.
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// Timestamp of the ping event.
	Timestamp int64 `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *UpdateRoomWithPingRequest) Reset() {
	*x = UpdateRoomWithPingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_rooms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoomWithPingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoomWithPingRequest) ProtoMessage() {}

func (x *UpdateRoomWithPingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rooms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoomWithPingRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoomWithPingRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_rooms_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateRoomWithPingRequest) GetSchedulerName() string {
	if x != nil {
		return x.SchedulerName
	}
	return ""
}

func (x *UpdateRoomWithPingRequest) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *UpdateRoomWithPingRequest) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UpdateRoomWithPingRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateRoomWithPingRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// The ping response.
type UpdateRoomWithPingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UpdateRoomWithPingResponse) Reset() {
	*x = UpdateRoomWithPingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_rooms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoomWithPingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoomWithPingResponse) ProtoMessage() {}

func (x *UpdateRoomWithPingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_rooms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoomWithPingResponse.ProtoReflect.Descriptor instead.
func (*UpdateRoomWithPingResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_rooms_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRoomWithPingResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_v1_rooms_proto protoreflect.FileDescriptor

var file_api_v1_rooms_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x36, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xaf, 0x01,
	0x0a, 0x0c, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9e,
	0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x57, 0x69, 0x74,
	0x68, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x57, 0x69, 0x74, 0x68,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x3b, 0x22, 0x36, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x2f, 0x7b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x2a, 0x7d, 0x2f, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2f, 0x7b, 0x72, 0x6f, 0x6f, 0x6d, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x2a, 0x7d, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x3a, 0x01, 0x2a, 0x42,
	0x51, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6f, 0x70, 0x66, 0x72, 0x65, 0x65, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x70, 0x66, 0x72, 0x65, 0x65, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f,
	0x6d, 0x61, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_rooms_proto_rawDescOnce sync.Once
	file_api_v1_rooms_proto_rawDescData = file_api_v1_rooms_proto_rawDesc
)

func file_api_v1_rooms_proto_rawDescGZIP() []byte {
	file_api_v1_rooms_proto_rawDescOnce.Do(func() {
		file_api_v1_rooms_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_rooms_proto_rawDescData)
	})
	return file_api_v1_rooms_proto_rawDescData
}

var file_api_v1_rooms_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_rooms_proto_goTypes = []interface{}{
	(*UpdateRoomWithPingRequest)(nil),  // 0: api.v1.UpdateRoomWithPingRequest
	(*UpdateRoomWithPingResponse)(nil), // 1: api.v1.UpdateRoomWithPingResponse
	(*_struct.Struct)(nil),             // 2: google.protobuf.Struct
}
var file_api_v1_rooms_proto_depIdxs = []int32{
	2, // 0: api.v1.UpdateRoomWithPingRequest.metadata:type_name -> google.protobuf.Struct
	0, // 1: api.v1.RoomsService.UpdateRoomWithPing:input_type -> api.v1.UpdateRoomWithPingRequest
	1, // 2: api.v1.RoomsService.UpdateRoomWithPing:output_type -> api.v1.UpdateRoomWithPingResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_rooms_proto_init() }
func file_api_v1_rooms_proto_init() {
	if File_api_v1_rooms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_rooms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoomWithPingRequest); i {
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
		file_api_v1_rooms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoomWithPingResponse); i {
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
			RawDescriptor: file_api_v1_rooms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_rooms_proto_goTypes,
		DependencyIndexes: file_api_v1_rooms_proto_depIdxs,
		MessageInfos:      file_api_v1_rooms_proto_msgTypes,
	}.Build()
	File_api_v1_rooms_proto = out.File
	file_api_v1_rooms_proto_rawDesc = nil
	file_api_v1_rooms_proto_goTypes = nil
	file_api_v1_rooms_proto_depIdxs = nil
}
