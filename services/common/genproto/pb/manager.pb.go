// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: manager.proto

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

type ServiceIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ServiceIdRequest) Reset() {
	*x = ServiceIdRequest{}
	mi := &file_manager_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceIdRequest) ProtoMessage() {}

func (x *ServiceIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceIdRequest.ProtoReflect.Descriptor instead.
func (*ServiceIdRequest) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceIdRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ServiceFullInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status   bool                   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Location *ServiceLocation       `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Current  *ServiceCurrentWeather `protobuf:"bytes,5,opt,name=current,proto3" json:"current,omitempty"`
}

func (x *ServiceFullInfo) Reset() {
	*x = ServiceFullInfo{}
	mi := &file_manager_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceFullInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceFullInfo) ProtoMessage() {}

func (x *ServiceFullInfo) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceFullInfo.ProtoReflect.Descriptor instead.
func (*ServiceFullInfo) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceFullInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ServiceFullInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceFullInfo) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ServiceFullInfo) GetLocation() *ServiceLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *ServiceFullInfo) GetCurrent() *ServiceCurrentWeather {
	if x != nil {
		return x.Current
	}
	return nil
}

type ServiceLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	City    string                 `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	Region  string                 `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	Current *ServiceCurrentWeather `protobuf:"bytes,3,opt,name=current,proto3" json:"current,omitempty"`
}

func (x *ServiceLocation) Reset() {
	*x = ServiceLocation{}
	mi := &file_manager_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceLocation) ProtoMessage() {}

func (x *ServiceLocation) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceLocation.ProtoReflect.Descriptor instead.
func (*ServiceLocation) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceLocation) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *ServiceLocation) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ServiceLocation) GetCurrent() *ServiceCurrentWeather {
	if x != nil {
		return x.Current
	}
	return nil
}

type ServiceCurrentWeather struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Condition   *ServiceConditionWeather `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	LastUpdated string                   `protobuf:"bytes,2,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	WindDir     string                   `protobuf:"bytes,3,opt,name=windDir,proto3" json:"windDir,omitempty"`
	FeelTemp    float32                  `protobuf:"fixed32,4,opt,name=feelTemp,proto3" json:"feelTemp,omitempty"`
	Temp        float32                  `protobuf:"fixed32,5,opt,name=temp,proto3" json:"temp,omitempty"`
	WindVel     float32                  `protobuf:"fixed32,6,opt,name=windVel,proto3" json:"windVel,omitempty"`
	IsDay       uint32                   `protobuf:"varint,7,opt,name=isDay,proto3" json:"isDay,omitempty"`
}

func (x *ServiceCurrentWeather) Reset() {
	*x = ServiceCurrentWeather{}
	mi := &file_manager_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceCurrentWeather) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceCurrentWeather) ProtoMessage() {}

func (x *ServiceCurrentWeather) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceCurrentWeather.ProtoReflect.Descriptor instead.
func (*ServiceCurrentWeather) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{3}
}

func (x *ServiceCurrentWeather) GetCondition() *ServiceConditionWeather {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *ServiceCurrentWeather) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

func (x *ServiceCurrentWeather) GetWindDir() string {
	if x != nil {
		return x.WindDir
	}
	return ""
}

func (x *ServiceCurrentWeather) GetFeelTemp() float32 {
	if x != nil {
		return x.FeelTemp
	}
	return 0
}

func (x *ServiceCurrentWeather) GetTemp() float32 {
	if x != nil {
		return x.Temp
	}
	return 0
}

func (x *ServiceCurrentWeather) GetWindVel() float32 {
	if x != nil {
		return x.WindVel
	}
	return 0
}

func (x *ServiceCurrentWeather) GetIsDay() uint32 {
	if x != nil {
		return x.IsDay
	}
	return 0
}

type ServiceConditionWeather struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Icon string `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (x *ServiceConditionWeather) Reset() {
	*x = ServiceConditionWeather{}
	mi := &file_manager_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceConditionWeather) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConditionWeather) ProtoMessage() {}

func (x *ServiceConditionWeather) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConditionWeather.ProtoReflect.Descriptor instead.
func (*ServiceConditionWeather) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{4}
}

func (x *ServiceConditionWeather) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ServiceConditionWeather) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

type RaspiService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status bool             `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Data   *ServiceCardData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RaspiService) Reset() {
	*x = RaspiService{}
	mi := &file_manager_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RaspiService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaspiService) ProtoMessage() {}

func (x *RaspiService) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaspiService.ProtoReflect.Descriptor instead.
func (*RaspiService) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{5}
}

func (x *RaspiService) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RaspiService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RaspiService) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *RaspiService) GetData() *ServiceCardData {
	if x != nil {
		return x.Data
	}
	return nil
}

type ServiceCardData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Icon        string `protobuf:"bytes,1,opt,name=icon,proto3" json:"icon,omitempty"`
	DataText    string `protobuf:"bytes,2,opt,name=dataText,proto3" json:"dataText,omitempty"`
	LastUpdated string `protobuf:"bytes,3,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
}

func (x *ServiceCardData) Reset() {
	*x = ServiceCardData{}
	mi := &file_manager_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceCardData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceCardData) ProtoMessage() {}

func (x *ServiceCardData) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceCardData.ProtoReflect.Descriptor instead.
func (*ServiceCardData) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{6}
}

func (x *ServiceCardData) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *ServiceCardData) GetDataText() string {
	if x != nil {
		return x.DataText
	}
	return ""
}

func (x *ServiceCardData) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

type GetServicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetServicesRequest) Reset() {
	*x = GetServicesRequest{}
	mi := &file_manager_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicesRequest) ProtoMessage() {}

func (x *GetServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicesRequest.ProtoReflect.Descriptor instead.
func (*GetServicesRequest) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{7}
}

type GetServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*RaspiService `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *GetServicesResponse) Reset() {
	*x = GetServicesResponse{}
	mi := &file_manager_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicesResponse) ProtoMessage() {}

func (x *GetServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicesResponse.ProtoReflect.Descriptor instead.
func (*GetServicesResponse) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{8}
}

func (x *GetServicesResponse) GetServices() []*RaspiService {
	if x != nil {
		return x.Services
	}
	return nil
}

type ConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigResponse) Reset() {
	*x = ConfigResponse{}
	mi := &file_manager_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigResponse) ProtoMessage() {}

func (x *ConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigResponse.ProtoReflect.Descriptor instead.
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{9}
}

var File_manager_proto protoreflect.FileDescriptor

var file_manager_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0xbd, 0x01, 0x0a,
	0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x77, 0x0a, 0x0f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x07, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x22, 0xf3, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12,
	0x3e, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x44, 0x69, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x64, 0x44, 0x69, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x65, 0x65, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x66,
	0x65, 0x65, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x6d, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x74, 0x65, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x77,
	0x69, 0x6e, 0x64, 0x56, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x77, 0x69,
	0x6e, 0x64, 0x56, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x44, 0x61, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x73, 0x44, 0x61, 0x79, 0x22, 0x41, 0x0a, 0x17, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x57,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x78,
	0x0a, 0x0c, 0x52, 0x61, 0x73, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x61, 0x72, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x63, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x14, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x73, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x10, 0x0a,
	0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xec, 0x02, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x1b, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42,
	0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x19,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x73, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x00, 0x12, 0x41, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x73, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3a,
	0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e,
	0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x72, 0x61, 0x73, 0x70, 0x69, 0x2d, 0x68, 0x74, 0x6d, 0x78, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_manager_proto_rawDescOnce sync.Once
	file_manager_proto_rawDescData = file_manager_proto_rawDesc
)

func file_manager_proto_rawDescGZIP() []byte {
	file_manager_proto_rawDescOnce.Do(func() {
		file_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_manager_proto_rawDescData)
	})
	return file_manager_proto_rawDescData
}

var file_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_manager_proto_goTypes = []any{
	(*ServiceIdRequest)(nil),        // 0: manager.ServiceIdRequest
	(*ServiceFullInfo)(nil),         // 1: manager.ServiceFullInfo
	(*ServiceLocation)(nil),         // 2: manager.ServiceLocation
	(*ServiceCurrentWeather)(nil),   // 3: manager.ServiceCurrentWeather
	(*ServiceConditionWeather)(nil), // 4: manager.ServiceConditionWeather
	(*RaspiService)(nil),            // 5: manager.RaspiService
	(*ServiceCardData)(nil),         // 6: manager.ServiceCardData
	(*GetServicesRequest)(nil),      // 7: manager.GetServicesRequest
	(*GetServicesResponse)(nil),     // 8: manager.GetServicesResponse
	(*ConfigResponse)(nil),          // 9: manager.ConfigResponse
}
var file_manager_proto_depIdxs = []int32{
	2,  // 0: manager.ServiceFullInfo.location:type_name -> manager.ServiceLocation
	3,  // 1: manager.ServiceFullInfo.current:type_name -> manager.ServiceCurrentWeather
	3,  // 2: manager.ServiceLocation.current:type_name -> manager.ServiceCurrentWeather
	4,  // 3: manager.ServiceCurrentWeather.condition:type_name -> manager.ServiceConditionWeather
	6,  // 4: manager.RaspiService.data:type_name -> manager.ServiceCardData
	5,  // 5: manager.GetServicesResponse.services:type_name -> manager.RaspiService
	7,  // 6: manager.ServiceManager.GetServices:input_type -> manager.GetServicesRequest
	0,  // 7: manager.ServiceManager.StartService:input_type -> manager.ServiceIdRequest
	0,  // 8: manager.ServiceManager.StopService:input_type -> manager.ServiceIdRequest
	0,  // 9: manager.ServiceManager.GetFullInfo:input_type -> manager.ServiceIdRequest
	0,  // 10: manager.ServiceManager.GetConfig:input_type -> manager.ServiceIdRequest
	8,  // 11: manager.ServiceManager.GetServices:output_type -> manager.GetServicesResponse
	5,  // 12: manager.ServiceManager.StartService:output_type -> manager.RaspiService
	5,  // 13: manager.ServiceManager.StopService:output_type -> manager.RaspiService
	1,  // 14: manager.ServiceManager.GetFullInfo:output_type -> manager.ServiceFullInfo
	9,  // 15: manager.ServiceManager.GetConfig:output_type -> manager.ConfigResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_manager_proto_init() }
func file_manager_proto_init() {
	if File_manager_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_manager_proto_goTypes,
		DependencyIndexes: file_manager_proto_depIdxs,
		MessageInfos:      file_manager_proto_msgTypes,
	}.Build()
	File_manager_proto = out.File
	file_manager_proto_rawDesc = nil
	file_manager_proto_goTypes = nil
	file_manager_proto_depIdxs = nil
}
