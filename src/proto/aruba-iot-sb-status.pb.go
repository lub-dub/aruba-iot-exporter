//  (C) Copyright [2018-2022] Hewlett Packard Enterprise Development LP
//  Note: File corresponds to AOS 8.10.x.x release

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: aruba-iot-sb-status.proto

package proto

import (
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

// Server Response to the IoT transport profile's connection
type ConnectCode int32

const (
	// connect ok
	ConnectCode_statusOK ConnectCode = 0
	// Token has expired
	ConnectCode_tokenExpire ConnectCode = 1
)

// Enum value maps for ConnectCode.
var (
	ConnectCode_name = map[int32]string{
		0: "statusOK",
		1: "tokenExpire",
	}
	ConnectCode_value = map[string]int32{
		"statusOK":    0,
		"tokenExpire": 1,
	}
)

func (x ConnectCode) Enum() *ConnectCode {
	p := new(ConnectCode)
	*p = x
	return p
}

func (x ConnectCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectCode) Descriptor() protoreflect.EnumDescriptor {
	return file_aruba_iot_sb_status_proto_enumTypes[0].Descriptor()
}

func (ConnectCode) Type() protoreflect.EnumType {
	return &file_aruba_iot_sb_status_proto_enumTypes[0]
}

func (x ConnectCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ConnectCode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ConnectCode(num)
	return nil
}

// Deprecated: Use ConnectCode.Descriptor instead.
func (ConnectCode) EnumDescriptor() ([]byte, []int) {
	return file_aruba_iot_sb_status_proto_rawDescGZIP(), []int{0}
}

// This topic can be used to let server respond on the connection with AP/Controllers
type ConnectStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The server response connection code
	ConnectCode *ConnectCode `protobuf:"varint,1,opt,name=connectCode,enum=aruba_telemetry.ConnectCode" json:"connectCode,omitempty"`
	// The Server response description for connection code
	ConnectDescription *string `protobuf:"bytes,2,opt,name=connectDescription" json:"connectDescription,omitempty"`
}

func (x *ConnectStatus) Reset() {
	*x = ConnectStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruba_iot_sb_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectStatus) ProtoMessage() {}

func (x *ConnectStatus) ProtoReflect() protoreflect.Message {
	mi := &file_aruba_iot_sb_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectStatus.ProtoReflect.Descriptor instead.
func (*ConnectStatus) Descriptor() ([]byte, []int) {
	return file_aruba_iot_sb_status_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectStatus) GetConnectCode() ConnectCode {
	if x != nil && x.ConnectCode != nil {
		return *x.ConnectCode
	}
	return ConnectCode_statusOK
}

func (x *ConnectStatus) GetConnectDescription() string {
	if x != nil && x.ConnectDescription != nil {
		return *x.ConnectDescription
	}
	return ""
}

var File_aruba_iot_sb_status_proto protoreflect.FileDescriptor

var file_aruba_iot_sb_status_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x72, 0x75, 0x62, 0x61, 0x2d, 0x69, 0x6f, 0x74, 0x2d, 0x73, 0x62, 0x2d, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x72, 0x75,
	0x62, 0x61, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x22, 0x7f, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x72, 0x75, 0x62, 0x61, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a,
	0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x2c, 0x0a,
	0x0b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0c, 0x0a, 0x08,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x10, 0x01, 0x42, 0x1b, 0x5a, 0x19, 0x6c,
	0x75, 0x62, 0x64, 0x75, 0x62, 0x2e, 0x6e, 0x6c, 0x2f, 0x61, 0x72, 0x75, 0x62, 0x61, 0x2d, 0x69,
	0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_aruba_iot_sb_status_proto_rawDescOnce sync.Once
	file_aruba_iot_sb_status_proto_rawDescData = file_aruba_iot_sb_status_proto_rawDesc
)

func file_aruba_iot_sb_status_proto_rawDescGZIP() []byte {
	file_aruba_iot_sb_status_proto_rawDescOnce.Do(func() {
		file_aruba_iot_sb_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_aruba_iot_sb_status_proto_rawDescData)
	})
	return file_aruba_iot_sb_status_proto_rawDescData
}

var file_aruba_iot_sb_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aruba_iot_sb_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_aruba_iot_sb_status_proto_goTypes = []any{
	(ConnectCode)(0),      // 0: aruba_telemetry.ConnectCode
	(*ConnectStatus)(nil), // 1: aruba_telemetry.ConnectStatus
}
var file_aruba_iot_sb_status_proto_depIdxs = []int32{
	0, // 0: aruba_telemetry.ConnectStatus.connectCode:type_name -> aruba_telemetry.ConnectCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aruba_iot_sb_status_proto_init() }
func file_aruba_iot_sb_status_proto_init() {
	if File_aruba_iot_sb_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aruba_iot_sb_status_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ConnectStatus); i {
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
			RawDescriptor: file_aruba_iot_sb_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aruba_iot_sb_status_proto_goTypes,
		DependencyIndexes: file_aruba_iot_sb_status_proto_depIdxs,
		EnumInfos:         file_aruba_iot_sb_status_proto_enumTypes,
		MessageInfos:      file_aruba_iot_sb_status_proto_msgTypes,
	}.Build()
	File_aruba_iot_sb_status_proto = out.File
	file_aruba_iot_sb_status_proto_rawDesc = nil
	file_aruba_iot_sb_status_proto_goTypes = nil
	file_aruba_iot_sb_status_proto_depIdxs = nil
}
