// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.19.6
// source: pb/BSSIDApple.proto

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

type WifiDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bssid    string               `protobuf:"bytes,1,opt,name=bssid,proto3" json:"bssid,omitempty"`
	Location *WifiDevice_Location `protobuf:"bytes,2,opt,name=location,proto3,oneof" json:"location,omitempty"`
}

func (x *WifiDevice) Reset() {
	*x = WifiDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_BSSIDApple_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WifiDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WifiDevice) ProtoMessage() {}

func (x *WifiDevice) ProtoReflect() protoreflect.Message {
	mi := &file_pb_BSSIDApple_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WifiDevice.ProtoReflect.Descriptor instead.
func (*WifiDevice) Descriptor() ([]byte, []int) {
	return file_pb_BSSIDApple_proto_rawDescGZIP(), []int{0}
}

func (x *WifiDevice) GetBssid() string {
	if x != nil {
		return x.Bssid
	}
	return ""
}

func (x *WifiDevice) GetLocation() *WifiDevice_Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type AppleWLoc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnknownValue0 *int64        `protobuf:"varint,1,opt,name=unknown_value0,json=unknownValue0,proto3,oneof" json:"unknown_value0,omitempty"`
	WifiDevices   []*WifiDevice `protobuf:"bytes,2,rep,name=wifi_devices,json=wifiDevices,proto3" json:"wifi_devices,omitempty"`
	UnknownValue1 *int32        `protobuf:"varint,3,opt,name=unknown_value1,json=unknownValue1,proto3,oneof" json:"unknown_value1,omitempty"`
	NumResults    *int32        `protobuf:"varint,4,opt,name=num_results,json=numResults,proto3,oneof" json:"num_results,omitempty"`
	APIName       *string       `protobuf:"bytes,5,opt,name=APIName,proto3,oneof" json:"APIName,omitempty"`
	UnknownValue2 *string       `protobuf:"bytes,6,opt,name=unknown_value2,json=unknownValue2,proto3,oneof" json:"unknown_value2,omitempty"`
	UnknownValue7 *int64        `protobuf:"varint,7,opt,name=unknown_value7,json=unknownValue7,proto3,oneof" json:"unknown_value7,omitempty"`
	DeviceType    *DeviceType   `protobuf:"bytes,33,opt,name=device_type,json=deviceType,proto3,oneof" json:"device_type,omitempty"`
}

func (x *AppleWLoc) Reset() {
	*x = AppleWLoc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_BSSIDApple_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppleWLoc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppleWLoc) ProtoMessage() {}

func (x *AppleWLoc) ProtoReflect() protoreflect.Message {
	mi := &file_pb_BSSIDApple_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppleWLoc.ProtoReflect.Descriptor instead.
func (*AppleWLoc) Descriptor() ([]byte, []int) {
	return file_pb_BSSIDApple_proto_rawDescGZIP(), []int{1}
}

func (x *AppleWLoc) GetUnknownValue0() int64 {
	if x != nil && x.UnknownValue0 != nil {
		return *x.UnknownValue0
	}
	return 0
}

func (x *AppleWLoc) GetWifiDevices() []*WifiDevice {
	if x != nil {
		return x.WifiDevices
	}
	return nil
}

func (x *AppleWLoc) GetUnknownValue1() int32 {
	if x != nil && x.UnknownValue1 != nil {
		return *x.UnknownValue1
	}
	return 0
}

func (x *AppleWLoc) GetNumResults() int32 {
	if x != nil && x.NumResults != nil {
		return *x.NumResults
	}
	return 0
}

func (x *AppleWLoc) GetAPIName() string {
	if x != nil && x.APIName != nil {
		return *x.APIName
	}
	return ""
}

func (x *AppleWLoc) GetUnknownValue2() string {
	if x != nil && x.UnknownValue2 != nil {
		return *x.UnknownValue2
	}
	return ""
}

func (x *AppleWLoc) GetUnknownValue7() int64 {
	if x != nil && x.UnknownValue7 != nil {
		return *x.UnknownValue7
	}
	return 0
}

func (x *AppleWLoc) GetDeviceType() *DeviceType {
	if x != nil {
		return x.DeviceType
	}
	return nil
}

type DeviceType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperatingSystem string `protobuf:"bytes,1,opt,name=operating_system,json=operatingSystem,proto3" json:"operating_system,omitempty"`
	Model           string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *DeviceType) Reset() {
	*x = DeviceType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_BSSIDApple_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceType) ProtoMessage() {}

func (x *DeviceType) ProtoReflect() protoreflect.Message {
	mi := &file_pb_BSSIDApple_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceType.ProtoReflect.Descriptor instead.
func (*DeviceType) Descriptor() ([]byte, []int) {
	return file_pb_BSSIDApple_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceType) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

func (x *DeviceType) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

type WifiDevice_Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude                           *int64 `protobuf:"varint,1,opt,name=latitude,proto3,oneof" json:"latitude,omitempty"`
	Longitude                          *int64 `protobuf:"varint,2,opt,name=longitude,proto3,oneof" json:"longitude,omitempty"`
	HorizontalAccuracy                 *int64 `protobuf:"varint,3,opt,name=horizontal_accuracy,json=horizontalAccuracy,proto3,oneof" json:"horizontal_accuracy,omitempty"`
	UnknownValue4                      *int64 `protobuf:"varint,4,opt,name=unknown_value4,json=unknownValue4,proto3,oneof" json:"unknown_value4,omitempty"`
	Altitude                           *int64 `protobuf:"varint,5,opt,name=altitude,proto3,oneof" json:"altitude,omitempty"`
	VerticalAccuracy                   *int64 `protobuf:"varint,6,opt,name=vertical_accuracy,json=verticalAccuracy,proto3,oneof" json:"vertical_accuracy,omitempty"`
	Speed                              *int64 `protobuf:"varint,7,opt,name=speed,proto3,oneof" json:"speed,omitempty"`
	Course                             *int64 `protobuf:"varint,8,opt,name=course,proto3,oneof" json:"course,omitempty"`
	Timestamp                          *int64 `protobuf:"varint,9,opt,name=timestamp,proto3,oneof" json:"timestamp,omitempty"`
	UnknownContext                     *int64 `protobuf:"varint,10,opt,name=unknown_context,json=unknownContext,proto3,oneof" json:"unknown_context,omitempty"` // Unused
	MotionActivityType                 *int64 `protobuf:"varint,11,opt,name=motion_activity_type,json=motionActivityType,proto3,oneof" json:"motion_activity_type,omitempty"`
	MotionActivityConfidence           *int64 `protobuf:"varint,12,opt,name=motion_activity_confidence,json=motionActivityConfidence,proto3,oneof" json:"motion_activity_confidence,omitempty"`
	Provider                           *int64 `protobuf:"varint,13,opt,name=provider,proto3,oneof" json:"provider,omitempty"` // Unused from this point on
	Floor                              *int64 `protobuf:"varint,14,opt,name=floor,proto3,oneof" json:"floor,omitempty"`
	Unknown15                          *int64 `protobuf:"varint,15,opt,name=unknown15,proto3,oneof" json:"unknown15,omitempty"`
	MotionVehicleConnectedStateChanged *int64 `protobuf:"varint,16,opt,name=motion_vehicle_connected_state_changed,json=motionVehicleConnectedStateChanged,proto3,oneof" json:"motion_vehicle_connected_state_changed,omitempty"`
	MotionVehicleConnected             *int64 `protobuf:"varint,17,opt,name=motion_vehicle_connected,json=motionVehicleConnected,proto3,oneof" json:"motion_vehicle_connected,omitempty"`
	RawMotionActivity                  *int64 `protobuf:"varint,18,opt,name=raw_motion_activity,json=rawMotionActivity,proto3,oneof" json:"raw_motion_activity,omitempty"`
	MotionActivity                     *int64 `protobuf:"varint,19,opt,name=motion_activity,json=motionActivity,proto3,oneof" json:"motion_activity,omitempty"`
	DominantMotionActivity             *int64 `protobuf:"varint,20,opt,name=dominant_motion_activity,json=dominantMotionActivity,proto3,oneof" json:"dominant_motion_activity,omitempty"`
	CourseAccuracy                     *int64 `protobuf:"varint,21,opt,name=course_accuracy,json=courseAccuracy,proto3,oneof" json:"course_accuracy,omitempty"`
	SpeedAccuracy                      *int64 `protobuf:"varint,22,opt,name=speed_accuracy,json=speedAccuracy,proto3,oneof" json:"speed_accuracy,omitempty"`
	ModeIndicator                      *int64 `protobuf:"varint,23,opt,name=mode_indicator,json=modeIndicator,proto3,oneof" json:"mode_indicator,omitempty"`
	HorzUncSemiMaj                     *int64 `protobuf:"varint,24,opt,name=horzUncSemiMaj,proto3,oneof" json:"horzUncSemiMaj,omitempty"` // WTF is this
	HorzUncSemiMin                     *int64 `protobuf:"varint,25,opt,name=horzUncSemiMin,proto3,oneof" json:"horzUncSemiMin,omitempty"`
	HorzUncSemiMajAz                   *int64 `protobuf:"varint,26,opt,name=horzUncSemiMajAz,proto3,oneof" json:"horzUncSemiMajAz,omitempty"`
	SatelliteReport                    *int64 `protobuf:"varint,27,opt,name=satellite_report,json=satelliteReport,proto3,oneof" json:"satellite_report,omitempty"`
	IsFromLocationController           *int64 `protobuf:"varint,28,opt,name=is_from_location_controller,json=isFromLocationController,proto3,oneof" json:"is_from_location_controller,omitempty"`
	PipelineDiagnosticReport           *int64 `protobuf:"varint,29,opt,name=pipeline_diagnostic_report,json=pipelineDiagnosticReport,proto3,oneof" json:"pipeline_diagnostic_report,omitempty"`
	BaroCalibrationIndication          *int64 `protobuf:"varint,30,opt,name=baro_calibration_indication,json=baroCalibrationIndication,proto3,oneof" json:"baro_calibration_indication,omitempty"`
	ProcessingMetadata                 *int64 `protobuf:"varint,31,opt,name=processing_metadata,json=processingMetadata,proto3,oneof" json:"processing_metadata,omitempty"`
}

func (x *WifiDevice_Location) Reset() {
	*x = WifiDevice_Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_BSSIDApple_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WifiDevice_Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WifiDevice_Location) ProtoMessage() {}

func (x *WifiDevice_Location) ProtoReflect() protoreflect.Message {
	mi := &file_pb_BSSIDApple_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WifiDevice_Location.ProtoReflect.Descriptor instead.
func (*WifiDevice_Location) Descriptor() ([]byte, []int) {
	return file_pb_BSSIDApple_proto_rawDescGZIP(), []int{0, 0}
}

func (x *WifiDevice_Location) GetLatitude() int64 {
	if x != nil && x.Latitude != nil {
		return *x.Latitude
	}
	return 0
}

func (x *WifiDevice_Location) GetLongitude() int64 {
	if x != nil && x.Longitude != nil {
		return *x.Longitude
	}
	return 0
}

func (x *WifiDevice_Location) GetHorizontalAccuracy() int64 {
	if x != nil && x.HorizontalAccuracy != nil {
		return *x.HorizontalAccuracy
	}
	return 0
}

func (x *WifiDevice_Location) GetUnknownValue4() int64 {
	if x != nil && x.UnknownValue4 != nil {
		return *x.UnknownValue4
	}
	return 0
}

func (x *WifiDevice_Location) GetAltitude() int64 {
	if x != nil && x.Altitude != nil {
		return *x.Altitude
	}
	return 0
}

func (x *WifiDevice_Location) GetVerticalAccuracy() int64 {
	if x != nil && x.VerticalAccuracy != nil {
		return *x.VerticalAccuracy
	}
	return 0
}

func (x *WifiDevice_Location) GetSpeed() int64 {
	if x != nil && x.Speed != nil {
		return *x.Speed
	}
	return 0
}

func (x *WifiDevice_Location) GetCourse() int64 {
	if x != nil && x.Course != nil {
		return *x.Course
	}
	return 0
}

func (x *WifiDevice_Location) GetTimestamp() int64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *WifiDevice_Location) GetUnknownContext() int64 {
	if x != nil && x.UnknownContext != nil {
		return *x.UnknownContext
	}
	return 0
}

func (x *WifiDevice_Location) GetMotionActivityType() int64 {
	if x != nil && x.MotionActivityType != nil {
		return *x.MotionActivityType
	}
	return 0
}

func (x *WifiDevice_Location) GetMotionActivityConfidence() int64 {
	if x != nil && x.MotionActivityConfidence != nil {
		return *x.MotionActivityConfidence
	}
	return 0
}

func (x *WifiDevice_Location) GetProvider() int64 {
	if x != nil && x.Provider != nil {
		return *x.Provider
	}
	return 0
}

func (x *WifiDevice_Location) GetFloor() int64 {
	if x != nil && x.Floor != nil {
		return *x.Floor
	}
	return 0
}

func (x *WifiDevice_Location) GetUnknown15() int64 {
	if x != nil && x.Unknown15 != nil {
		return *x.Unknown15
	}
	return 0
}

func (x *WifiDevice_Location) GetMotionVehicleConnectedStateChanged() int64 {
	if x != nil && x.MotionVehicleConnectedStateChanged != nil {
		return *x.MotionVehicleConnectedStateChanged
	}
	return 0
}

func (x *WifiDevice_Location) GetMotionVehicleConnected() int64 {
	if x != nil && x.MotionVehicleConnected != nil {
		return *x.MotionVehicleConnected
	}
	return 0
}

func (x *WifiDevice_Location) GetRawMotionActivity() int64 {
	if x != nil && x.RawMotionActivity != nil {
		return *x.RawMotionActivity
	}
	return 0
}

func (x *WifiDevice_Location) GetMotionActivity() int64 {
	if x != nil && x.MotionActivity != nil {
		return *x.MotionActivity
	}
	return 0
}

func (x *WifiDevice_Location) GetDominantMotionActivity() int64 {
	if x != nil && x.DominantMotionActivity != nil {
		return *x.DominantMotionActivity
	}
	return 0
}

func (x *WifiDevice_Location) GetCourseAccuracy() int64 {
	if x != nil && x.CourseAccuracy != nil {
		return *x.CourseAccuracy
	}
	return 0
}

func (x *WifiDevice_Location) GetSpeedAccuracy() int64 {
	if x != nil && x.SpeedAccuracy != nil {
		return *x.SpeedAccuracy
	}
	return 0
}

func (x *WifiDevice_Location) GetModeIndicator() int64 {
	if x != nil && x.ModeIndicator != nil {
		return *x.ModeIndicator
	}
	return 0
}

func (x *WifiDevice_Location) GetHorzUncSemiMaj() int64 {
	if x != nil && x.HorzUncSemiMaj != nil {
		return *x.HorzUncSemiMaj
	}
	return 0
}

func (x *WifiDevice_Location) GetHorzUncSemiMin() int64 {
	if x != nil && x.HorzUncSemiMin != nil {
		return *x.HorzUncSemiMin
	}
	return 0
}

func (x *WifiDevice_Location) GetHorzUncSemiMajAz() int64 {
	if x != nil && x.HorzUncSemiMajAz != nil {
		return *x.HorzUncSemiMajAz
	}
	return 0
}

func (x *WifiDevice_Location) GetSatelliteReport() int64 {
	if x != nil && x.SatelliteReport != nil {
		return *x.SatelliteReport
	}
	return 0
}

func (x *WifiDevice_Location) GetIsFromLocationController() int64 {
	if x != nil && x.IsFromLocationController != nil {
		return *x.IsFromLocationController
	}
	return 0
}

func (x *WifiDevice_Location) GetPipelineDiagnosticReport() int64 {
	if x != nil && x.PipelineDiagnosticReport != nil {
		return *x.PipelineDiagnosticReport
	}
	return 0
}

func (x *WifiDevice_Location) GetBaroCalibrationIndication() int64 {
	if x != nil && x.BaroCalibrationIndication != nil {
		return *x.BaroCalibrationIndication
	}
	return 0
}

func (x *WifiDevice_Location) GetProcessingMetadata() int64 {
	if x != nil && x.ProcessingMetadata != nil {
		return *x.ProcessingMetadata
	}
	return 0
}

var File_pb_BSSIDApple_proto protoreflect.FileDescriptor

var file_pb_BSSIDApple_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x2f, 0x42, 0x53, 0x53, 0x49, 0x44, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x11, 0x0a, 0x0a, 0x57, 0x69, 0x66, 0x69, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x73, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x73, 0x73, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x57,
	0x69, 0x66, 0x69, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x1a, 0xf1, 0x10, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x21, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x34, 0x0a, 0x13, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c,
	0x5f, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x02, 0x52, 0x12, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x41, 0x63, 0x63,
	0x75, 0x72, 0x61, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x75, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x03, 0x52, 0x0d, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x34, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x04, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x5f, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x05, 0x52, 0x10, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x75,
	0x72, 0x61, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x06, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x07, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x21, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x08, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x88,
	0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x48, 0x09, 0x52, 0x0e, 0x75,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x35, 0x0a, 0x14, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x48, 0x0a,
	0x52, 0x12, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x1a, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x48, 0x0b, 0x52, 0x18, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x48, 0x0c, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x66,
	0x6c, 0x6f, 0x6f, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x48, 0x0d, 0x52, 0x05, 0x66, 0x6c,
	0x6f, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x31, 0x35, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x48, 0x0e, 0x52, 0x09, 0x75, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x31, 0x35, 0x88, 0x01, 0x01, 0x12, 0x57, 0x0a, 0x26, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x48, 0x0f, 0x52, 0x22, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x3d, 0x0a, 0x18, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x10, 0x52, 0x16, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x33, 0x0a, 0x13, 0x72, 0x61, 0x77, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x48, 0x11,
	0x52, 0x11, 0x72, 0x61, 0x77, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x12, 0x52, 0x0e, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x18, 0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x74,
	0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x48, 0x13, 0x52, 0x16, 0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x61,
	0x6e, 0x74, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x61, 0x63,
	0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x48, 0x14, 0x52, 0x0e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x41, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x88, 0x01,
	0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x75, 0x72,
	0x61, 0x63, 0x79, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x48, 0x15, 0x52, 0x0d, 0x73, 0x70, 0x65,
	0x65, 0x64, 0x41, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a,
	0x0e, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x03, 0x48, 0x16, 0x52, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x68, 0x6f, 0x72,
	0x7a, 0x55, 0x6e, 0x63, 0x53, 0x65, 0x6d, 0x69, 0x4d, 0x61, 0x6a, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x17, 0x52, 0x0e, 0x68, 0x6f, 0x72, 0x7a, 0x55, 0x6e, 0x63, 0x53, 0x65, 0x6d, 0x69,
	0x4d, 0x61, 0x6a, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x68, 0x6f, 0x72, 0x7a, 0x55, 0x6e,
	0x63, 0x53, 0x65, 0x6d, 0x69, 0x4d, 0x69, 0x6e, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x48, 0x18,
	0x52, 0x0e, 0x68, 0x6f, 0x72, 0x7a, 0x55, 0x6e, 0x63, 0x53, 0x65, 0x6d, 0x69, 0x4d, 0x69, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x10, 0x68, 0x6f, 0x72, 0x7a, 0x55, 0x6e, 0x63, 0x53, 0x65,
	0x6d, 0x69, 0x4d, 0x61, 0x6a, 0x41, 0x7a, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x03, 0x48, 0x19, 0x52,
	0x10, 0x68, 0x6f, 0x72, 0x7a, 0x55, 0x6e, 0x63, 0x53, 0x65, 0x6d, 0x69, 0x4d, 0x61, 0x6a, 0x41,
	0x7a, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74,
	0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03, 0x48, 0x1a,
	0x52, 0x0f, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x1b, 0x69, 0x73, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x03, 0x48, 0x1b, 0x52, 0x18, 0x69, 0x73, 0x46,
	0x72, 0x6f, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x1a, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x5f,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x03, 0x48, 0x1c, 0x52, 0x18,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74,
	0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x43, 0x0a, 0x1b, 0x62,
	0x61, 0x72, 0x6f, 0x5f, 0x63, 0x61, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x1d, 0x52, 0x19, 0x62, 0x61, 0x72, 0x6f, 0x43, 0x61, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x34, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x03, 0x48, 0x1e, 0x52,
	0x12, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c,
	0x5f, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x75, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x34, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x76, 0x65,
	0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42,
	0x1d, 0x0a, 0x1b, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x31, 0x35, 0x42, 0x29, 0x0a, 0x27, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x42, 0x1b,
	0x0a, 0x19, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x16, 0x0a, 0x14, 0x5f,
	0x72, 0x61, 0x77, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x42, 0x1b, 0x0a, 0x19, 0x5f, 0x64, 0x6f, 0x6d, 0x69,
	0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x70, 0x65,
	0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x42, 0x11, 0x0a, 0x0f, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x11,
	0x0a, 0x0f, 0x5f, 0x68, 0x6f, 0x72, 0x7a, 0x55, 0x6e, 0x63, 0x53, 0x65, 0x6d, 0x69, 0x4d, 0x61,
	0x6a, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x68, 0x6f, 0x72, 0x7a, 0x55, 0x6e, 0x63, 0x53, 0x65, 0x6d,
	0x69, 0x4d, 0x69, 0x6e, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x68, 0x6f, 0x72, 0x7a, 0x55, 0x6e, 0x63,
	0x53, 0x65, 0x6d, 0x69, 0x4d, 0x61, 0x6a, 0x41, 0x7a, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x73, 0x61,
	0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x1e,
	0x0a, 0x1c, 0x5f, 0x69, 0x73, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x42, 0x1d,
	0x0a, 0x1b, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x69, 0x61, 0x67,
	0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x1e, 0x0a,
	0x1c, 0x5f, 0x62, 0x61, 0x72, 0x6f, 0x5f, 0x63, 0x61, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x16, 0x0a,
	0x14, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xdb, 0x03, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x57, 0x4c, 0x6f, 0x63,
	0x12, 0x2a, 0x0a, 0x0e, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x30, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0d, 0x75, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x30, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x0c,
	0x77, 0x69, 0x66, 0x69, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x57, 0x69, 0x66, 0x69, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x0b, 0x77, 0x69, 0x66, 0x69, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x0e,
	0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x0d, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x31, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52,
	0x0a, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1d,
	0x0a, 0x07, 0x41, 0x50, 0x49, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x03, 0x52, 0x07, 0x41, 0x50, 0x49, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a,
	0x0e, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0d, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x75, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x37, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x05, 0x52, 0x0d, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x37, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x48, 0x06, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x75, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x30, 0x42, 0x11, 0x0a, 0x0f, 0x5f,
	0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x41, 0x50, 0x49, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x75,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x42, 0x11, 0x0a,
	0x0f, 0x5f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x37,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x4d, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_BSSIDApple_proto_rawDescOnce sync.Once
	file_pb_BSSIDApple_proto_rawDescData = file_pb_BSSIDApple_proto_rawDesc
)

func file_pb_BSSIDApple_proto_rawDescGZIP() []byte {
	file_pb_BSSIDApple_proto_rawDescOnce.Do(func() {
		file_pb_BSSIDApple_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_BSSIDApple_proto_rawDescData)
	})
	return file_pb_BSSIDApple_proto_rawDescData
}

var file_pb_BSSIDApple_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_BSSIDApple_proto_goTypes = []interface{}{
	(*WifiDevice)(nil),          // 0: WifiDevice
	(*AppleWLoc)(nil),           // 1: AppleWLoc
	(*DeviceType)(nil),          // 2: DeviceType
	(*WifiDevice_Location)(nil), // 3: WifiDevice.Location
}
var file_pb_BSSIDApple_proto_depIdxs = []int32{
	3, // 0: WifiDevice.location:type_name -> WifiDevice.Location
	0, // 1: AppleWLoc.wifi_devices:type_name -> WifiDevice
	2, // 2: AppleWLoc.device_type:type_name -> DeviceType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_BSSIDApple_proto_init() }
func file_pb_BSSIDApple_proto_init() {
	if File_pb_BSSIDApple_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_BSSIDApple_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WifiDevice); i {
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
		file_pb_BSSIDApple_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppleWLoc); i {
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
		file_pb_BSSIDApple_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceType); i {
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
		file_pb_BSSIDApple_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WifiDevice_Location); i {
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
	file_pb_BSSIDApple_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pb_BSSIDApple_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_pb_BSSIDApple_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_BSSIDApple_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_BSSIDApple_proto_goTypes,
		DependencyIndexes: file_pb_BSSIDApple_proto_depIdxs,
		MessageInfos:      file_pb_BSSIDApple_proto_msgTypes,
	}.Build()
	File_pb_BSSIDApple_proto = out.File
	file_pb_BSSIDApple_proto_rawDesc = nil
	file_pb_BSSIDApple_proto_goTypes = nil
	file_pb_BSSIDApple_proto_depIdxs = nil
}
