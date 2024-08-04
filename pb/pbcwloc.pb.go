// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.19.6
// source: pb/pbcwloc.proto

package pb

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

type MotionActivityType int32

const (
	MotionActivity_unknown    MotionActivityType = 0
	MotionActivity_stationary MotionActivityType = 1
	MotionActivity_walking    MotionActivityType = 2
	MotionActivity_running    MotionActivityType = 3
	MotionActivity_automotive MotionActivityType = 4
	MotionActivity_cycling    MotionActivityType = 5
)

// Enum value maps for MotionActivityType.
var (
	MotionActivityType_name = map[int32]string{
		0: "unknown",
		1: "stationary",
		2: "walking",
		3: "running",
		4: "automotive",
		5: "cycling",
	}
	MotionActivityType_value = map[string]int32{
		"unknown":    0,
		"stationary": 1,
		"walking":    2,
		"running":    3,
		"automotive": 4,
		"cycling":    5,
	}
)

func (x MotionActivityType) Enum() *MotionActivityType {
	p := new(MotionActivityType)
	*p = x
	return p
}

func (x MotionActivityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MotionActivityType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_pbcwloc_proto_enumTypes[0].Descriptor()
}

func (MotionActivityType) Type() protoreflect.EnumType {
	return &file_pb_pbcwloc_proto_enumTypes[0]
}

func (x MotionActivityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MotionActivityType.Descriptor instead.
func (MotionActivityType) EnumDescriptor() ([]byte, []int) {
	return file_pb_pbcwloc_proto_rawDescGZIP(), []int{2, 0}
}

// _CLPWifiAPLocationReadFrom
type PbcWifiEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bssid     string           `protobuf:"bytes,1,opt,name=bssid,proto3" json:"bssid,omitempty"`
	Channel   int32            `protobuf:"varint,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Rssi      int32            `protobuf:"varint,3,opt,name=rssi,proto3" json:"rssi,omitempty"`
	Location  *PbcWlocLocation `protobuf:"bytes,4,opt,name=location,proto3,oneof" json:"location,omitempty"`
	Hidden    int32            `protobuf:"varint,7,opt,name=hidden,proto3" json:"hidden,omitempty"`
	Timestamp float64          `protobuf:"fixed64,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ScanType  int32            `protobuf:"varint,9,opt,name=scan_type,json=scanType,proto3" json:"scan_type,omitempty"`
}

func (x *PbcWifiEntry) Reset() {
	*x = PbcWifiEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pbcwloc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbcWifiEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbcWifiEntry) ProtoMessage() {}

func (x *PbcWifiEntry) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pbcwloc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbcWifiEntry.ProtoReflect.Descriptor instead.
func (*PbcWifiEntry) Descriptor() ([]byte, []int) {
	return file_pb_pbcwloc_proto_rawDescGZIP(), []int{0}
}

func (x *PbcWifiEntry) GetBssid() string {
	if x != nil {
		return x.Bssid
	}
	return ""
}

func (x *PbcWifiEntry) GetChannel() int32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

func (x *PbcWifiEntry) GetRssi() int32 {
	if x != nil {
		return x.Rssi
	}
	return 0
}

func (x *PbcWifiEntry) GetLocation() *PbcWlocLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *PbcWifiEntry) GetHidden() int32 {
	if x != nil {
		return x.Hidden
	}
	return 0
}

func (x *PbcWifiEntry) GetTimestamp() float64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *PbcWifiEntry) GetScanType() int32 {
	if x != nil {
		return x.ScanType
	}
	return 0
}

type PbcWlocRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceInfo  *DeviceType     `protobuf:"bytes,163,opt,name=device_info,json=deviceInfo,proto3" json:"device_info,omitempty"`
	WifiEntries []*PbcWifiEntry `protobuf:"bytes,3,rep,name=wifi_entries,json=wifiEntries,proto3" json:"wifi_entries,omitempty"`
}

func (x *PbcWlocRequest) Reset() {
	*x = PbcWlocRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pbcwloc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbcWlocRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbcWlocRequest) ProtoMessage() {}

func (x *PbcWlocRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pbcwloc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbcWlocRequest.ProtoReflect.Descriptor instead.
func (*PbcWlocRequest) Descriptor() ([]byte, []int) {
	return file_pb_pbcwloc_proto_rawDescGZIP(), []int{1}
}

func (x *PbcWlocRequest) GetDeviceInfo() *DeviceType {
	if x != nil {
		return x.DeviceInfo
	}
	return nil
}

func (x *PbcWlocRequest) GetWifiEntries() []*PbcWifiEntry {
	if x != nil {
		return x.WifiEntries
	}
	return nil
}

// _CLPMotionActivityReadFrom
type MotionActivity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Confidence uint32             `protobuf:"varint,1,opt,name=confidence,proto3" json:"confidence,omitempty"`
	Activity   MotionActivityType `protobuf:"varint,2,opt,name=activity,proto3,enum=MotionActivityType" json:"activity,omitempty"`
}

func (x *MotionActivity) Reset() {
	*x = MotionActivity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pbcwloc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MotionActivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MotionActivity) ProtoMessage() {}

func (x *MotionActivity) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pbcwloc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MotionActivity.ProtoReflect.Descriptor instead.
func (*MotionActivity) Descriptor() ([]byte, []int) {
	return file_pb_pbcwloc_proto_rawDescGZIP(), []int{2}
}

func (x *MotionActivity) GetConfidence() uint32 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

func (x *MotionActivity) GetActivity() MotionActivityType {
	if x != nil {
		return x.Activity
	}
	return MotionActivity_unknown
}

type PbcWlocLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude                           float64         `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude                          float64         `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	HorizontalAccuracy                 float32         `protobuf:"fixed32,3,opt,name=horizontal_accuracy,json=horizontalAccuracy,proto3" json:"horizontal_accuracy,omitempty"`
	Altitude                           float32         `protobuf:"fixed32,5,opt,name=altitude,proto3" json:"altitude,omitempty"`
	VerticalAccuracy                   float32         `protobuf:"fixed32,6,opt,name=vertical_accuracy,json=verticalAccuracy,proto3" json:"vertical_accuracy,omitempty"`
	Speed                              *float32        `protobuf:"fixed32,7,opt,name=speed,proto3,oneof" json:"speed,omitempty"`
	Course                             *float32        `protobuf:"fixed32,8,opt,name=course,proto3,oneof" json:"course,omitempty"`
	Timestamp                          float64         `protobuf:"fixed64,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Provider                           int32           `protobuf:"varint,13,opt,name=provider,proto3" json:"provider,omitempty"` // Only shows up if your phone is tied to a cell provider
	MotionVehicleConnectedStateChanged int32           `protobuf:"varint,16,opt,name=motion_vehicle_connected_state_changed,json=motionVehicleConnectedStateChanged,proto3" json:"motion_vehicle_connected_state_changed,omitempty"`
	MotionVehicleConnected             int32           `protobuf:"varint,17,opt,name=motion_vehicle_connected,json=motionVehicleConnected,proto3" json:"motion_vehicle_connected,omitempty"`
	RawMotionActivity                  *MotionActivity `protobuf:"bytes,18,opt,name=raw_motion_activity,json=rawMotionActivity,proto3" json:"raw_motion_activity,omitempty"`
	MotionActivity                     *MotionActivity `protobuf:"bytes,19,opt,name=motion_activity,json=motionActivity,proto3" json:"motion_activity,omitempty"`
	DominantMotionActivity             *MotionActivity `protobuf:"bytes,20,opt,name=dominant_motion_activity,json=dominantMotionActivity,proto3" json:"dominant_motion_activity,omitempty"`
	CourseAccuracy                     *float32        `protobuf:"fixed32,21,opt,name=course_accuracy,json=courseAccuracy,proto3,oneof" json:"course_accuracy,omitempty"`
	SpeedAccuracy                      *float32        `protobuf:"fixed32,22,opt,name=speed_accuracy,json=speedAccuracy,proto3,oneof" json:"speed_accuracy,omitempty"`
}

func (x *PbcWlocLocation) Reset() {
	*x = PbcWlocLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pbcwloc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbcWlocLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbcWlocLocation) ProtoMessage() {}

func (x *PbcWlocLocation) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pbcwloc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbcWlocLocation.ProtoReflect.Descriptor instead.
func (*PbcWlocLocation) Descriptor() ([]byte, []int) {
	return file_pb_pbcwloc_proto_rawDescGZIP(), []int{3}
}

func (x *PbcWlocLocation) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *PbcWlocLocation) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *PbcWlocLocation) GetHorizontalAccuracy() float32 {
	if x != nil {
		return x.HorizontalAccuracy
	}
	return 0
}

func (x *PbcWlocLocation) GetAltitude() float32 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *PbcWlocLocation) GetVerticalAccuracy() float32 {
	if x != nil {
		return x.VerticalAccuracy
	}
	return 0
}

func (x *PbcWlocLocation) GetSpeed() float32 {
	if x != nil && x.Speed != nil {
		return *x.Speed
	}
	return 0
}

func (x *PbcWlocLocation) GetCourse() float32 {
	if x != nil && x.Course != nil {
		return *x.Course
	}
	return 0
}

func (x *PbcWlocLocation) GetTimestamp() float64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *PbcWlocLocation) GetProvider() int32 {
	if x != nil {
		return x.Provider
	}
	return 0
}

func (x *PbcWlocLocation) GetMotionVehicleConnectedStateChanged() int32 {
	if x != nil {
		return x.MotionVehicleConnectedStateChanged
	}
	return 0
}

func (x *PbcWlocLocation) GetMotionVehicleConnected() int32 {
	if x != nil {
		return x.MotionVehicleConnected
	}
	return 0
}

func (x *PbcWlocLocation) GetRawMotionActivity() *MotionActivity {
	if x != nil {
		return x.RawMotionActivity
	}
	return nil
}

func (x *PbcWlocLocation) GetMotionActivity() *MotionActivity {
	if x != nil {
		return x.MotionActivity
	}
	return nil
}

func (x *PbcWlocLocation) GetDominantMotionActivity() *MotionActivity {
	if x != nil {
		return x.DominantMotionActivity
	}
	return nil
}

func (x *PbcWlocLocation) GetCourseAccuracy() float32 {
	if x != nil && x.CourseAccuracy != nil {
		return *x.CourseAccuracy
	}
	return 0
}

func (x *PbcWlocLocation) GetSpeedAccuracy() float32 {
	if x != nil && x.SpeedAccuracy != nil {
		return *x.SpeedAccuracy
	}
	return 0
}

var File_pb_pbcwloc_proto protoreflect.FileDescriptor

var file_pb_pbcwloc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x63, 0x77, 0x6c, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x70, 0x62, 0x2f, 0x42, 0x53, 0x53, 0x49, 0x44, 0x41, 0x70, 0x70, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01, 0x0a, 0x0c, 0x50, 0x62, 0x63, 0x57,
	0x69, 0x66, 0x69, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x73, 0x73, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x73, 0x73, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x73, 0x73, 0x69,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x73, 0x73, 0x69, 0x12, 0x31, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x50, 0x62, 0x63, 0x57, 0x6c, 0x6f, 0x63, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x63, 0x61, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x71, 0x0a, 0x0e, 0x50, 0x62, 0x63, 0x57, 0x6c, 0x6f, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0xa3, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x30, 0x0a, 0x0c, 0x77, 0x69, 0x66, 0x69, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x50, 0x62, 0x63, 0x57, 0x69, 0x66, 0x69,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x77, 0x69, 0x66, 0x69, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x22, 0xbe, 0x01, 0x0a, 0x0e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x52, 0x08, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0x5a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x76, 0x65, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x63, 0x79, 0x63, 0x6c, 0x69, 0x6e,
	0x67, 0x10, 0x05, 0x22, 0xa1, 0x06, 0x0a, 0x0f, 0x50, 0x62, 0x63, 0x57, 0x6c, 0x6f, 0x63, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x2f, 0x0a, 0x13, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x5f,
	0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x75, 0x72, 0x61,
	0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x2b,
	0x0a, 0x11, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x75, 0x72,
	0x61, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x63, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x12, 0x19, 0x0a, 0x05, 0x73,
	0x70, 0x65, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x05, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x48, 0x01, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x52, 0x0a,
	0x26, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x22, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x12, 0x38, 0x0a, 0x18, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x16, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x13, 0x72,
	0x61, 0x77, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x11, 0x72, 0x61, 0x77, 0x4d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x0f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x0e, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x49, 0x0a, 0x18, 0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x61,
	0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x16, 0x64, 0x6f, 0x6d, 0x69, 0x6e,
	0x61, 0x6e, 0x74, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0x2c, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x75,
	0x72, 0x61, 0x63, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x02, 0x48, 0x02, 0x52, 0x0e, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x41, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12,
	0x2a, 0x0a, 0x0e, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63,
	0x79, 0x18, 0x16, 0x20, 0x01, 0x28, 0x02, 0x48, 0x03, 0x52, 0x0d, 0x73, 0x70, 0x65, 0x65, 0x64,
	0x41, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x73, 0x70, 0x65, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x75,
	0x72, 0x61, 0x63, 0x79, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x61,
	0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_pbcwloc_proto_rawDescOnce sync.Once
	file_pb_pbcwloc_proto_rawDescData = file_pb_pbcwloc_proto_rawDesc
)

func file_pb_pbcwloc_proto_rawDescGZIP() []byte {
	file_pb_pbcwloc_proto_rawDescOnce.Do(func() {
		file_pb_pbcwloc_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_pbcwloc_proto_rawDescData)
	})
	return file_pb_pbcwloc_proto_rawDescData
}

var file_pb_pbcwloc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_pbcwloc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_pbcwloc_proto_goTypes = []interface{}{
	(MotionActivityType)(0), // 0: MotionActivity.type
	(*PbcWifiEntry)(nil),    // 1: PbcWifiEntry
	(*PbcWlocRequest)(nil),  // 2: PbcWlocRequest
	(*MotionActivity)(nil),  // 3: MotionActivity
	(*PbcWlocLocation)(nil), // 4: PbcWlocLocation
	(*DeviceType)(nil),      // 5: DeviceType
}
var file_pb_pbcwloc_proto_depIdxs = []int32{
	4, // 0: PbcWifiEntry.location:type_name -> PbcWlocLocation
	5, // 1: PbcWlocRequest.device_info:type_name -> DeviceType
	1, // 2: PbcWlocRequest.wifi_entries:type_name -> PbcWifiEntry
	0, // 3: MotionActivity.activity:type_name -> MotionActivity.type
	3, // 4: PbcWlocLocation.raw_motion_activity:type_name -> MotionActivity
	3, // 5: PbcWlocLocation.motion_activity:type_name -> MotionActivity
	3, // 6: PbcWlocLocation.dominant_motion_activity:type_name -> MotionActivity
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_pb_pbcwloc_proto_init() }
func file_pb_pbcwloc_proto_init() {
	if File_pb_pbcwloc_proto != nil {
		return
	}
	file_pb_BSSIDApple_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_pbcwloc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbcWifiEntry); i {
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
		file_pb_pbcwloc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbcWlocRequest); i {
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
		file_pb_pbcwloc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MotionActivity); i {
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
		file_pb_pbcwloc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbcWlocLocation); i {
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
	file_pb_pbcwloc_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pb_pbcwloc_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_pbcwloc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_pbcwloc_proto_goTypes,
		DependencyIndexes: file_pb_pbcwloc_proto_depIdxs,
		EnumInfos:         file_pb_pbcwloc_proto_enumTypes,
		MessageInfos:      file_pb_pbcwloc_proto_msgTypes,
	}.Build()
	File_pb_pbcwloc_proto = out.File
	file_pb_pbcwloc_proto_rawDesc = nil
	file_pb_pbcwloc_proto_goTypes = nil
	file_pb_pbcwloc_proto_depIdxs = nil
}
