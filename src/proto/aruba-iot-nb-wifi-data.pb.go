//  (C) Copyright [2018-2022] Hewlett Packard Enterprise Development LP
//  Note: File corresponds to AOS 8.10.x.x release

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: aruba-iot-nb-wifi-data.proto

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

// WiFi device data as recorded by AOS
type WiFiData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Advertisement address of the sender
	Mac []byte `protobuf:"bytes,1,opt,name=mac" json:"mac,omitempty"`
	// device sub type (rtls, assoc, unassoc, ...)
	DeviceClass []DeviceClassEnum `protobuf:"varint,2,rep,name=deviceClass,enum=aruba_telemetry.DeviceClassEnum" json:"deviceClass,omitempty"`
	// signal strength in dBm
	Rssi *int32 `protobuf:"zigzag32,3,opt,name=rssi" json:"rssi,omitempty"`
	// payload of rtls tag (passed opaquely to server)
	RtlsPayload []byte `protobuf:"bytes,4,opt,name=rtls_payload,json=rtlsPayload" json:"rtls_payload,omitempty"`
}

func (x *WiFiData) Reset() {
	*x = WiFiData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruba_iot_nb_wifi_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WiFiData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WiFiData) ProtoMessage() {}

func (x *WiFiData) ProtoReflect() protoreflect.Message {
	mi := &file_aruba_iot_nb_wifi_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WiFiData.ProtoReflect.Descriptor instead.
func (*WiFiData) Descriptor() ([]byte, []int) {
	return file_aruba_iot_nb_wifi_data_proto_rawDescGZIP(), []int{0}
}

func (x *WiFiData) GetMac() []byte {
	if x != nil {
		return x.Mac
	}
	return nil
}

func (x *WiFiData) GetDeviceClass() []DeviceClassEnum {
	if x != nil {
		return x.DeviceClass
	}
	return nil
}

func (x *WiFiData) GetRssi() int32 {
	if x != nil && x.Rssi != nil {
		return *x.Rssi
	}
	return 0
}

func (x *WiFiData) GetRtlsPayload() []byte {
	if x != nil {
		return x.RtlsPayload
	}
	return nil
}

var File_aruba_iot_nb_wifi_data_proto protoreflect.FileDescriptor

var file_aruba_iot_nb_wifi_data_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x72, 0x75, 0x62, 0x61, 0x2d, 0x69, 0x6f, 0x74, 0x2d, 0x6e, 0x62, 0x2d, 0x77,
	0x69, 0x66, 0x69, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x61, 0x72, 0x75, 0x62, 0x61, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x1a,
	0x15, 0x61, 0x72, 0x75, 0x62, 0x61, 0x2d, 0x69, 0x6f, 0x74, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x08, 0x57, 0x69, 0x46, 0x69, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x42, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x72, 0x75,
	0x62, 0x61, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0b, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x73, 0x73,
	0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x72, 0x73, 0x73, 0x69, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x74, 0x6c, 0x73, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0b, 0x72, 0x74, 0x6c, 0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x42, 0x1b, 0x5a, 0x19, 0x6c, 0x75, 0x62, 0x64, 0x75, 0x62, 0x2e, 0x6e, 0x6c, 0x2f, 0x61, 0x72,
	0x75, 0x62, 0x61, 0x2d, 0x69, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_aruba_iot_nb_wifi_data_proto_rawDescOnce sync.Once
	file_aruba_iot_nb_wifi_data_proto_rawDescData = file_aruba_iot_nb_wifi_data_proto_rawDesc
)

func file_aruba_iot_nb_wifi_data_proto_rawDescGZIP() []byte {
	file_aruba_iot_nb_wifi_data_proto_rawDescOnce.Do(func() {
		file_aruba_iot_nb_wifi_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_aruba_iot_nb_wifi_data_proto_rawDescData)
	})
	return file_aruba_iot_nb_wifi_data_proto_rawDescData
}

var file_aruba_iot_nb_wifi_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_aruba_iot_nb_wifi_data_proto_goTypes = []any{
	(*WiFiData)(nil),     // 0: aruba_telemetry.WiFiData
	(DeviceClassEnum)(0), // 1: aruba_telemetry.deviceClassEnum
}
var file_aruba_iot_nb_wifi_data_proto_depIdxs = []int32{
	1, // 0: aruba_telemetry.WiFiData.deviceClass:type_name -> aruba_telemetry.deviceClassEnum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aruba_iot_nb_wifi_data_proto_init() }
func file_aruba_iot_nb_wifi_data_proto_init() {
	if File_aruba_iot_nb_wifi_data_proto != nil {
		return
	}
	file_aruba_iot_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_aruba_iot_nb_wifi_data_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*WiFiData); i {
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
			RawDescriptor: file_aruba_iot_nb_wifi_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aruba_iot_nb_wifi_data_proto_goTypes,
		DependencyIndexes: file_aruba_iot_nb_wifi_data_proto_depIdxs,
		MessageInfos:      file_aruba_iot_nb_wifi_data_proto_msgTypes,
	}.Build()
	File_aruba_iot_nb_wifi_data_proto = out.File
	file_aruba_iot_nb_wifi_data_proto_rawDesc = nil
	file_aruba_iot_nb_wifi_data_proto_goTypes = nil
	file_aruba_iot_nb_wifi_data_proto_depIdxs = nil
}
