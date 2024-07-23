//  (C) Copyright [2018-2022] Hewlett Packard Enterprise Development LP
//  Note: File corresponds to AOS 8.10.x.x release

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: aruba-iot-nb-ap-health-update.proto

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

// Indicates status of the Radio/USB device
type HealthStatus int32

const (
	HealthStatus_healthy     HealthStatus = 0
	HealthStatus_degraded    HealthStatus = 1
	HealthStatus_unavailable HealthStatus = 2
)

// Enum value maps for HealthStatus.
var (
	HealthStatus_name = map[int32]string{
		0: "healthy",
		1: "degraded",
		2: "unavailable",
	}
	HealthStatus_value = map[string]int32{
		"healthy":     0,
		"degraded":    1,
		"unavailable": 2,
	}
)

func (x HealthStatus) Enum() *HealthStatus {
	p := new(HealthStatus)
	*p = x
	return p
}

func (x HealthStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_aruba_iot_nb_ap_health_update_proto_enumTypes[0].Descriptor()
}

func (HealthStatus) Type() protoreflect.EnumType {
	return &file_aruba_iot_nb_ap_health_update_proto_enumTypes[0]
}

func (x HealthStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *HealthStatus) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = HealthStatus(num)
	return nil
}

// Deprecated: Use HealthStatus.Descriptor instead.
func (HealthStatus) EnumDescriptor() ([]byte, []int) {
	return file_aruba_iot_nb_ap_health_update_proto_rawDescGZIP(), []int{0}
}

// Indicates FW type running on the onboard IoT radio
type IotRadioFirmware int32

const (
	IotRadioFirmware_arubaDefault IotRadioFirmware = 0
)

// Enum value maps for IotRadioFirmware.
var (
	IotRadioFirmware_name = map[int32]string{
		0: "arubaDefault",
	}
	IotRadioFirmware_value = map[string]int32{
		"arubaDefault": 0,
	}
)

func (x IotRadioFirmware) Enum() *IotRadioFirmware {
	p := new(IotRadioFirmware)
	*p = x
	return p
}

func (x IotRadioFirmware) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IotRadioFirmware) Descriptor() protoreflect.EnumDescriptor {
	return file_aruba_iot_nb_ap_health_update_proto_enumTypes[1].Descriptor()
}

func (IotRadioFirmware) Type() protoreflect.EnumType {
	return &file_aruba_iot_nb_ap_health_update_proto_enumTypes[1]
}

func (x IotRadioFirmware) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *IotRadioFirmware) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = IotRadioFirmware(num)
	return nil
}

// Deprecated: Use IotRadioFirmware.Descriptor instead.
func (IotRadioFirmware) EnumDescriptor() ([]byte, []int) {
	return file_aruba_iot_nb_ap_health_update_proto_rawDescGZIP(), []int{1}
}

// Indicates type of the IoT radio HW
// gen1 is capable of BLE 4.0 only (3xx series minus 303P)
// gen2 is capable of BLE 5.0 and Zigbee  (303P, 5xx and 6xx series)
type IotRadioType int32

const (
	IotRadioType_gen1 IotRadioType = 0
	IotRadioType_gen2 IotRadioType = 1
)

// Enum value maps for IotRadioType.
var (
	IotRadioType_name = map[int32]string{
		0: "gen1",
		1: "gen2",
	}
	IotRadioType_value = map[string]int32{
		"gen1": 0,
		"gen2": 1,
	}
)

func (x IotRadioType) Enum() *IotRadioType {
	p := new(IotRadioType)
	*p = x
	return p
}

func (x IotRadioType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IotRadioType) Descriptor() protoreflect.EnumDescriptor {
	return file_aruba_iot_nb_ap_health_update_proto_enumTypes[2].Descriptor()
}

func (IotRadioType) Type() protoreflect.EnumType {
	return &file_aruba_iot_nb_ap_health_update_proto_enumTypes[2]
}

func (x IotRadioType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *IotRadioType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = IotRadioType(num)
	return nil
}

// Deprecated: Use IotRadioType.Descriptor instead.
func (IotRadioType) EnumDescriptor() ([]byte, []int) {
	return file_aruba_iot_nb_ap_health_update_proto_rawDescGZIP(), []int{2}
}

// Iot Radio sub message include information about the
// onboard/USB Aruba IoT radio
type IotRadio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// MAC address of the device
	Mac []byte `protobuf:"bytes,1,opt,name=mac" json:"mac,omitempty"`
	// Hardware type of the Aruba radio
	Hardware *IotRadioType `protobuf:"varint,2,opt,name=hardware,enum=aruba_telemetry.IotRadioType" json:"hardware,omitempty"`
	// FW type running on the Aruba IoT radio
	Firmware *IotRadioFirmware `protobuf:"varint,3,opt,name=firmware,enum=aruba_telemetry.IotRadioFirmware" json:"firmware,omitempty"`
	// Indicates status of the Radio/USB device
	Health *HealthStatus `protobuf:"varint,4,opt,name=health,enum=aruba_telemetry.HealthStatus" json:"health,omitempty"`
	// Value will be true in case of the AP-USB-ZB dongle, false otherwise
	External *bool `protobuf:"varint,5,opt,name=external" json:"external,omitempty"`
}

func (x *IotRadio) Reset() {
	*x = IotRadio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruba_iot_nb_ap_health_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IotRadio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IotRadio) ProtoMessage() {}

func (x *IotRadio) ProtoReflect() protoreflect.Message {
	mi := &file_aruba_iot_nb_ap_health_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IotRadio.ProtoReflect.Descriptor instead.
func (*IotRadio) Descriptor() ([]byte, []int) {
	return file_aruba_iot_nb_ap_health_update_proto_rawDescGZIP(), []int{0}
}

func (x *IotRadio) GetMac() []byte {
	if x != nil {
		return x.Mac
	}
	return nil
}

func (x *IotRadio) GetHardware() IotRadioType {
	if x != nil && x.Hardware != nil {
		return *x.Hardware
	}
	return IotRadioType_gen1
}

func (x *IotRadio) GetFirmware() IotRadioFirmware {
	if x != nil && x.Firmware != nil {
		return *x.Firmware
	}
	return IotRadioFirmware_arubaDefault
}

func (x *IotRadio) GetHealth() HealthStatus {
	if x != nil && x.Health != nil {
		return *x.Health
	}
	return HealthStatus_healthy
}

func (x *IotRadio) GetExternal() bool {
	if x != nil && x.External != nil {
		return *x.External
	}
	return false
}

// USB device sub message include information about
// supported partner USB devices (IoT gateways, sensors, etc.)
type UsbDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is a unique string to identify the attached USB device
	Identifier *string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	// Indicates status of the Radio/USB device
	Health *HealthStatus `protobuf:"varint,2,opt,name=health,enum=aruba_telemetry.HealthStatus" json:"health,omitempty"`
}

func (x *UsbDevice) Reset() {
	*x = UsbDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruba_iot_nb_ap_health_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsbDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsbDevice) ProtoMessage() {}

func (x *UsbDevice) ProtoReflect() protoreflect.Message {
	mi := &file_aruba_iot_nb_ap_health_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsbDevice.ProtoReflect.Descriptor instead.
func (*UsbDevice) Descriptor() ([]byte, []int) {
	return file_aruba_iot_nb_ap_health_update_proto_rawDescGZIP(), []int{1}
}

func (x *UsbDevice) GetIdentifier() string {
	if x != nil && x.Identifier != nil {
		return *x.Identifier
	}
	return ""
}

func (x *UsbDevice) GetHealth() HealthStatus {
	if x != nil && x.Health != nil {
		return *x.Health
	}
	return HealthStatus_healthy
}

// Indicates health of the various IoT devices onboard/attached
// to the AP
type ApHealthUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApStatus *HealthStatus `protobuf:"varint,1,opt,name=apStatus,enum=aruba_telemetry.HealthStatus" json:"apStatus,omitempty"`
	Radio    []*IotRadio   `protobuf:"bytes,2,rep,name=radio" json:"radio,omitempty"`
	Usb      []*UsbDevice  `protobuf:"bytes,3,rep,name=usb" json:"usb,omitempty"`
}

func (x *ApHealthUpdate) Reset() {
	*x = ApHealthUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruba_iot_nb_ap_health_update_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApHealthUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApHealthUpdate) ProtoMessage() {}

func (x *ApHealthUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_aruba_iot_nb_ap_health_update_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApHealthUpdate.ProtoReflect.Descriptor instead.
func (*ApHealthUpdate) Descriptor() ([]byte, []int) {
	return file_aruba_iot_nb_ap_health_update_proto_rawDescGZIP(), []int{2}
}

func (x *ApHealthUpdate) GetApStatus() HealthStatus {
	if x != nil && x.ApStatus != nil {
		return *x.ApStatus
	}
	return HealthStatus_healthy
}

func (x *ApHealthUpdate) GetRadio() []*IotRadio {
	if x != nil {
		return x.Radio
	}
	return nil
}

func (x *ApHealthUpdate) GetUsb() []*UsbDevice {
	if x != nil {
		return x.Usb
	}
	return nil
}

var File_aruba_iot_nb_ap_health_update_proto protoreflect.FileDescriptor

var file_aruba_iot_nb_ap_health_update_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x72, 0x75, 0x62, 0x61, 0x2d, 0x69, 0x6f, 0x74, 0x2d, 0x6e, 0x62, 0x2d, 0x61,
	0x70, 0x2d, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x72, 0x75, 0x62, 0x61, 0x5f, 0x74, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x22, 0xe9, 0x01, 0x0a, 0x08, 0x49, 0x6f, 0x74, 0x52, 0x61,
	0x64, 0x69, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x39, 0x0a, 0x08, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x61, 0x72, 0x75, 0x62, 0x61, 0x5f,
	0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x49, 0x6f, 0x74, 0x52, 0x61, 0x64,
	0x69, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65,
	0x12, 0x3d, 0x0a, 0x08, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x21, 0x2e, 0x61, 0x72, 0x75, 0x62, 0x61, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x49, 0x6f, 0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x46, 0x69, 0x72,
	0x6d, 0x77, 0x61, 0x72, 0x65, 0x52, 0x08, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x12,
	0x35, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x61, 0x72, 0x75, 0x62, 0x61, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x22, 0x62, 0x0a, 0x09, 0x55, 0x73, 0x62, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x35, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x61, 0x72, 0x75, 0x62, 0x61, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0xaa, 0x01, 0x0a, 0x0e, 0x41, 0x70, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x61, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x61, 0x72,
	0x75, 0x62, 0x61, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x08, 0x61, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x05, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x72, 0x75, 0x62, 0x61, 0x5f, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x49, 0x6f, 0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x52, 0x05,
	0x72, 0x61, 0x64, 0x69, 0x6f, 0x12, 0x2c, 0x0a, 0x03, 0x75, 0x73, 0x62, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x72, 0x75, 0x62, 0x61, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x55, 0x73, 0x62, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x03,
	0x75, 0x73, 0x62, 0x2a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x64, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x02, 0x2a,
	0x24, 0x0a, 0x10, 0x49, 0x6f, 0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x46, 0x69, 0x72, 0x6d, 0x77,
	0x61, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x61, 0x72, 0x75, 0x62, 0x61, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x10, 0x00, 0x2a, 0x22, 0x0a, 0x0c, 0x49, 0x6f, 0x74, 0x52, 0x61, 0x64, 0x69,
	0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x67, 0x65, 0x6e, 0x31, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x67, 0x65, 0x6e, 0x32, 0x10, 0x01, 0x42, 0x1b, 0x5a, 0x19, 0x6c, 0x75, 0x62,
	0x64, 0x75, 0x62, 0x2e, 0x6e, 0x6c, 0x2f, 0x61, 0x72, 0x75, 0x62, 0x61, 0x2d, 0x69, 0x6f, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_aruba_iot_nb_ap_health_update_proto_rawDescOnce sync.Once
	file_aruba_iot_nb_ap_health_update_proto_rawDescData = file_aruba_iot_nb_ap_health_update_proto_rawDesc
)

func file_aruba_iot_nb_ap_health_update_proto_rawDescGZIP() []byte {
	file_aruba_iot_nb_ap_health_update_proto_rawDescOnce.Do(func() {
		file_aruba_iot_nb_ap_health_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_aruba_iot_nb_ap_health_update_proto_rawDescData)
	})
	return file_aruba_iot_nb_ap_health_update_proto_rawDescData
}

var file_aruba_iot_nb_ap_health_update_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_aruba_iot_nb_ap_health_update_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_aruba_iot_nb_ap_health_update_proto_goTypes = []any{
	(HealthStatus)(0),      // 0: aruba_telemetry.HealthStatus
	(IotRadioFirmware)(0),  // 1: aruba_telemetry.IotRadioFirmware
	(IotRadioType)(0),      // 2: aruba_telemetry.IotRadioType
	(*IotRadio)(nil),       // 3: aruba_telemetry.IotRadio
	(*UsbDevice)(nil),      // 4: aruba_telemetry.UsbDevice
	(*ApHealthUpdate)(nil), // 5: aruba_telemetry.ApHealthUpdate
}
var file_aruba_iot_nb_ap_health_update_proto_depIdxs = []int32{
	2, // 0: aruba_telemetry.IotRadio.hardware:type_name -> aruba_telemetry.IotRadioType
	1, // 1: aruba_telemetry.IotRadio.firmware:type_name -> aruba_telemetry.IotRadioFirmware
	0, // 2: aruba_telemetry.IotRadio.health:type_name -> aruba_telemetry.HealthStatus
	0, // 3: aruba_telemetry.UsbDevice.health:type_name -> aruba_telemetry.HealthStatus
	0, // 4: aruba_telemetry.ApHealthUpdate.apStatus:type_name -> aruba_telemetry.HealthStatus
	3, // 5: aruba_telemetry.ApHealthUpdate.radio:type_name -> aruba_telemetry.IotRadio
	4, // 6: aruba_telemetry.ApHealthUpdate.usb:type_name -> aruba_telemetry.UsbDevice
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_aruba_iot_nb_ap_health_update_proto_init() }
func file_aruba_iot_nb_ap_health_update_proto_init() {
	if File_aruba_iot_nb_ap_health_update_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aruba_iot_nb_ap_health_update_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IotRadio); i {
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
		file_aruba_iot_nb_ap_health_update_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UsbDevice); i {
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
		file_aruba_iot_nb_ap_health_update_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ApHealthUpdate); i {
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
			RawDescriptor: file_aruba_iot_nb_ap_health_update_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aruba_iot_nb_ap_health_update_proto_goTypes,
		DependencyIndexes: file_aruba_iot_nb_ap_health_update_proto_depIdxs,
		EnumInfos:         file_aruba_iot_nb_ap_health_update_proto_enumTypes,
		MessageInfos:      file_aruba_iot_nb_ap_health_update_proto_msgTypes,
	}.Build()
	File_aruba_iot_nb_ap_health_update_proto = out.File
	file_aruba_iot_nb_ap_health_update_proto_rawDesc = nil
	file_aruba_iot_nb_ap_health_update_proto_goTypes = nil
	file_aruba_iot_nb_ap_health_update_proto_depIdxs = nil
}
