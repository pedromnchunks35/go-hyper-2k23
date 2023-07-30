// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: Weather_service.proto

package protofiles

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

type SearchDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryName string `protobuf:"bytes,1,opt,name=country_name,json=countryName,proto3" json:"country_name,omitempty"`
}

func (x *SearchDetails) Reset() {
	*x = SearchDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Weather_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchDetails) ProtoMessage() {}

func (x *SearchDetails) ProtoReflect() protoreflect.Message {
	mi := &file_Weather_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchDetails.ProtoReflect.Descriptor instead.
func (*SearchDetails) Descriptor() ([]byte, []int) {
	return file_Weather_service_proto_rawDescGZIP(), []int{0}
}

func (x *SearchDetails) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

type Information struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Current  *Current  `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
}

func (x *Information) Reset() {
	*x = Information{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Weather_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Information) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Information) ProtoMessage() {}

func (x *Information) ProtoReflect() protoreflect.Message {
	mi := &file_Weather_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Information.ProtoReflect.Descriptor instead.
func (*Information) Descriptor() ([]byte, []int) {
	return file_Weather_service_proto_rawDescGZIP(), []int{1}
}

func (x *Information) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Information) GetCurrent() *Current {
	if x != nil {
		return x.Current
	}
	return nil
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Region         string  `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	Country        string  `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Lat            float32 `protobuf:"fixed32,4,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon            float32 `protobuf:"fixed32,5,opt,name=lon,proto3" json:"lon,omitempty"`
	TzId           string  `protobuf:"bytes,6,opt,name=tz_id,json=tzId,proto3" json:"tz_id,omitempty"`
	LocaltimeEpoch int32   `protobuf:"varint,7,opt,name=localtime_epoch,json=localtimeEpoch,proto3" json:"localtime_epoch,omitempty"`
	Localtime      string  `protobuf:"bytes,8,opt,name=localtime,proto3" json:"localtime,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Weather_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_Weather_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_Weather_service_proto_rawDescGZIP(), []int{2}
}

func (x *Location) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Location) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Location) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Location) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Location) GetLon() float32 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *Location) GetTzId() string {
	if x != nil {
		return x.TzId
	}
	return ""
}

func (x *Location) GetLocaltimeEpoch() int32 {
	if x != nil {
		return x.LocaltimeEpoch
	}
	return 0
}

func (x *Location) GetLocaltime() string {
	if x != nil {
		return x.Localtime
	}
	return ""
}

type Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastUpdatedEpoch int32      `protobuf:"varint,1,opt,name=last_updated_epoch,json=lastUpdatedEpoch,proto3" json:"last_updated_epoch,omitempty"`
	LastUpdated      string     `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	TempC            float32    `protobuf:"fixed32,3,opt,name=temp_c,json=tempC,proto3" json:"temp_c,omitempty"`
	TempF            float32    `protobuf:"fixed32,4,opt,name=temp_f,json=tempF,proto3" json:"temp_f,omitempty"`
	IsDay            int32      `protobuf:"varint,5,opt,name=is_day,json=isDay,proto3" json:"is_day,omitempty"`
	Condition        *Condition `protobuf:"bytes,6,opt,name=condition,proto3" json:"condition,omitempty"`
	WindMph          float32    `protobuf:"fixed32,7,opt,name=wind_mph,json=windMph,proto3" json:"wind_mph,omitempty"`
	WindKph          float32    `protobuf:"fixed32,8,opt,name=wind_kph,json=windKph,proto3" json:"wind_kph,omitempty"`
	WindDegree       int32      `protobuf:"varint,9,opt,name=wind_degree,json=windDegree,proto3" json:"wind_degree,omitempty"`
	WindDir          string     `protobuf:"bytes,10,opt,name=wind_dir,json=windDir,proto3" json:"wind_dir,omitempty"`
	PressureMb       float32    `protobuf:"fixed32,11,opt,name=pressure_mb,json=pressureMb,proto3" json:"pressure_mb,omitempty"`
	PressureIn       float32    `protobuf:"fixed32,12,opt,name=pressure_in,json=pressureIn,proto3" json:"pressure_in,omitempty"`
	PrecipMm         float32    `protobuf:"fixed32,13,opt,name=precip_mm,json=precipMm,proto3" json:"precip_mm,omitempty"`
	PrecipIn         float32    `protobuf:"fixed32,14,opt,name=precip_in,json=precipIn,proto3" json:"precip_in,omitempty"`
	Humidity         int32      `protobuf:"varint,15,opt,name=humidity,proto3" json:"humidity,omitempty"`
	Cloud            int32      `protobuf:"varint,16,opt,name=cloud,proto3" json:"cloud,omitempty"`
	FeelslikeC       float32    `protobuf:"fixed32,17,opt,name=feelslike_c,json=feelslikeC,proto3" json:"feelslike_c,omitempty"`
	FeelsLikeF       float32    `protobuf:"fixed32,18,opt,name=feels_like_f,json=feelsLikeF,proto3" json:"feels_like_f,omitempty"`
	VisKm            float32    `protobuf:"fixed32,19,opt,name=vis_km,json=visKm,proto3" json:"vis_km,omitempty"`
	VisMiles         float32    `protobuf:"fixed32,20,opt,name=vis_miles,json=visMiles,proto3" json:"vis_miles,omitempty"`
	Uv               float32    `protobuf:"fixed32,21,opt,name=uv,proto3" json:"uv,omitempty"`
	GusMph           float32    `protobuf:"fixed32,22,opt,name=gus_mph,json=gusMph,proto3" json:"gus_mph,omitempty"`
	GusKph           float32    `protobuf:"fixed32,23,opt,name=gus_kph,json=gusKph,proto3" json:"gus_kph,omitempty"`
}

func (x *Current) Reset() {
	*x = Current{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Weather_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Current) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Current) ProtoMessage() {}

func (x *Current) ProtoReflect() protoreflect.Message {
	mi := &file_Weather_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Current.ProtoReflect.Descriptor instead.
func (*Current) Descriptor() ([]byte, []int) {
	return file_Weather_service_proto_rawDescGZIP(), []int{3}
}

func (x *Current) GetLastUpdatedEpoch() int32 {
	if x != nil {
		return x.LastUpdatedEpoch
	}
	return 0
}

func (x *Current) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

func (x *Current) GetTempC() float32 {
	if x != nil {
		return x.TempC
	}
	return 0
}

func (x *Current) GetTempF() float32 {
	if x != nil {
		return x.TempF
	}
	return 0
}

func (x *Current) GetIsDay() int32 {
	if x != nil {
		return x.IsDay
	}
	return 0
}

func (x *Current) GetCondition() *Condition {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *Current) GetWindMph() float32 {
	if x != nil {
		return x.WindMph
	}
	return 0
}

func (x *Current) GetWindKph() float32 {
	if x != nil {
		return x.WindKph
	}
	return 0
}

func (x *Current) GetWindDegree() int32 {
	if x != nil {
		return x.WindDegree
	}
	return 0
}

func (x *Current) GetWindDir() string {
	if x != nil {
		return x.WindDir
	}
	return ""
}

func (x *Current) GetPressureMb() float32 {
	if x != nil {
		return x.PressureMb
	}
	return 0
}

func (x *Current) GetPressureIn() float32 {
	if x != nil {
		return x.PressureIn
	}
	return 0
}

func (x *Current) GetPrecipMm() float32 {
	if x != nil {
		return x.PrecipMm
	}
	return 0
}

func (x *Current) GetPrecipIn() float32 {
	if x != nil {
		return x.PrecipIn
	}
	return 0
}

func (x *Current) GetHumidity() int32 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *Current) GetCloud() int32 {
	if x != nil {
		return x.Cloud
	}
	return 0
}

func (x *Current) GetFeelslikeC() float32 {
	if x != nil {
		return x.FeelslikeC
	}
	return 0
}

func (x *Current) GetFeelsLikeF() float32 {
	if x != nil {
		return x.FeelsLikeF
	}
	return 0
}

func (x *Current) GetVisKm() float32 {
	if x != nil {
		return x.VisKm
	}
	return 0
}

func (x *Current) GetVisMiles() float32 {
	if x != nil {
		return x.VisMiles
	}
	return 0
}

func (x *Current) GetUv() float32 {
	if x != nil {
		return x.Uv
	}
	return 0
}

func (x *Current) GetGusMph() float32 {
	if x != nil {
		return x.GusMph
	}
	return 0
}

func (x *Current) GetGusKph() float32 {
	if x != nil {
		return x.GusKph
	}
	return 0
}

type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Icon string `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"`
	Code int32  `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Weather_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_Weather_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_Weather_service_proto_rawDescGZIP(), []int{4}
}

func (x *Condition) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Condition) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Condition) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_Weather_service_proto protoreflect.FileDescriptor

var file_Weather_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x22, 0x32, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6e, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x07,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x22, 0xd0, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x13,
	0x0a, 0x05, 0x74, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x7a, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x1c, 0x0a, 0x09,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xad, 0x05, 0x0a, 0x07, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x45,
	0x70, 0x6f, 0x63, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x65, 0x6d, 0x70, 0x5f,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x65, 0x6d, 0x70, 0x43, 0x12, 0x15,
	0x0a, 0x06, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x74, 0x65, 0x6d, 0x70, 0x46, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x64, 0x61, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x73, 0x44, 0x61, 0x79, 0x12, 0x33, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x69, 0x6e, 0x64, 0x5f, 0x6d, 0x70, 0x68, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x4d, 0x70, 0x68, 0x12, 0x19, 0x0a, 0x08,
	0x77, 0x69, 0x6e, 0x64, 0x5f, 0x6b, 0x70, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x77, 0x69, 0x6e, 0x64, 0x4b, 0x70, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x69, 0x6e, 0x64, 0x5f,
	0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x77, 0x69,
	0x6e, 0x64, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x69, 0x6e, 0x64,
	0x5f, 0x64, 0x69, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x64,
	0x44, 0x69, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x5f,
	0x6d, 0x62, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75,
	0x72, 0x65, 0x4d, 0x62, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65,
	0x5f, 0x69, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x75, 0x72, 0x65, 0x49, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x70, 0x5f,
	0x6d, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x70, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x4d, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x70, 0x5f, 0x69, 0x6e, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x70, 0x72, 0x65, 0x63, 0x69, 0x70, 0x49, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x65, 0x65, 0x6c, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x63,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x65, 0x65, 0x6c, 0x73, 0x6c, 0x69, 0x6b,
	0x65, 0x43, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x65, 0x65, 0x6c, 0x73, 0x5f, 0x6c, 0x69, 0x6b, 0x65,
	0x5f, 0x66, 0x18, 0x12, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x65, 0x65, 0x6c, 0x73, 0x4c,
	0x69, 0x6b, 0x65, 0x46, 0x12, 0x15, 0x0a, 0x06, 0x76, 0x69, 0x73, 0x5f, 0x6b, 0x6d, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x69, 0x73, 0x4b, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x76,
	0x69, 0x73, 0x5f, 0x6d, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x76, 0x69, 0x73, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x75, 0x76, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x02, 0x75, 0x76, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x75, 0x73, 0x5f,
	0x6d, 0x70, 0x68, 0x18, 0x16, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x67, 0x75, 0x73, 0x4d, 0x70,
	0x68, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x75, 0x73, 0x5f, 0x6b, 0x70, 0x68, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x06, 0x67, 0x75, 0x73, 0x4b, 0x70, 0x68, 0x22, 0x47, 0x0a, 0x09, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x58, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x14, 0x5a,
	0x12, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Weather_service_proto_rawDescOnce sync.Once
	file_Weather_service_proto_rawDescData = file_Weather_service_proto_rawDesc
)

func file_Weather_service_proto_rawDescGZIP() []byte {
	file_Weather_service_proto_rawDescOnce.Do(func() {
		file_Weather_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_Weather_service_proto_rawDescData)
	})
	return file_Weather_service_proto_rawDescData
}

var file_Weather_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Weather_service_proto_goTypes = []interface{}{
	(*SearchDetails)(nil), // 0: protofiles.SearchDetails
	(*Information)(nil),   // 1: protofiles.Information
	(*Location)(nil),      // 2: protofiles.Location
	(*Current)(nil),       // 3: protofiles.Current
	(*Condition)(nil),     // 4: protofiles.Condition
}
var file_Weather_service_proto_depIdxs = []int32{
	2, // 0: protofiles.Information.location:type_name -> protofiles.Location
	3, // 1: protofiles.Information.current:type_name -> protofiles.Current
	4, // 2: protofiles.Current.condition:type_name -> protofiles.Condition
	0, // 3: protofiles.WeatherService.RequestWeather:input_type -> protofiles.SearchDetails
	1, // 4: protofiles.WeatherService.RequestWeather:output_type -> protofiles.Information
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Weather_service_proto_init() }
func file_Weather_service_proto_init() {
	if File_Weather_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Weather_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchDetails); i {
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
		file_Weather_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Information); i {
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
		file_Weather_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_Weather_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Current); i {
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
		file_Weather_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
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
			RawDescriptor: file_Weather_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Weather_service_proto_goTypes,
		DependencyIndexes: file_Weather_service_proto_depIdxs,
		MessageInfos:      file_Weather_service_proto_msgTypes,
	}.Build()
	File_Weather_service_proto = out.File
	file_Weather_service_proto_rawDesc = nil
	file_Weather_service_proto_goTypes = nil
	file_Weather_service_proto_depIdxs = nil
}