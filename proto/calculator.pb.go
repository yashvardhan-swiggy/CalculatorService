// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: proto/calculator.proto

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

type GetSumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int64 `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 int64 `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *GetSumRequest) Reset() {
	*x = GetSumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSumRequest) ProtoMessage() {}

func (x *GetSumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSumRequest.ProtoReflect.Descriptor instead.
func (*GetSumRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *GetSumRequest) GetNum1() int64 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *GetSumRequest) GetNum2() int64 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type GetSumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int64 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *GetSumResponse) Reset() {
	*x = GetSumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSumResponse) ProtoMessage() {}

func (x *GetSumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSumResponse.ProtoReflect.Descriptor instead.
func (*GetSumResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *GetSumResponse) GetSum() int64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type GetPrimeNumbersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *GetPrimeNumbersRequest) Reset() {
	*x = GetPrimeNumbersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrimeNumbersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrimeNumbersRequest) ProtoMessage() {}

func (x *GetPrimeNumbersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrimeNumbersRequest.ProtoReflect.Descriptor instead.
func (*GetPrimeNumbersRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *GetPrimeNumbersRequest) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type GetPrimeNumbersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prime int64 `protobuf:"varint,1,opt,name=prime,proto3" json:"prime,omitempty"`
}

func (x *GetPrimeNumbersResponse) Reset() {
	*x = GetPrimeNumbersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrimeNumbersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrimeNumbersResponse) ProtoMessage() {}

func (x *GetPrimeNumbersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrimeNumbersResponse.ProtoReflect.Descriptor instead.
func (*GetPrimeNumbersResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *GetPrimeNumbersResponse) GetPrime() int64 {
	if x != nil {
		return x.Prime
	}
	return 0
}

type GetAverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *GetAverageRequest) Reset() {
	*x = GetAverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAverageRequest) ProtoMessage() {}

func (x *GetAverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAverageRequest.ProtoReflect.Descriptor instead.
func (*GetAverageRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{4}
}

func (x *GetAverageRequest) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type GetAverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Average int64 `protobuf:"varint,1,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *GetAverageResponse) Reset() {
	*x = GetAverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAverageResponse) ProtoMessage() {}

func (x *GetAverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAverageResponse.ProtoReflect.Descriptor instead.
func (*GetAverageResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{5}
}

func (x *GetAverageResponse) GetAverage() int64 {
	if x != nil {
		return x.Average
	}
	return 0
}

type GetMaximumNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *GetMaximumNumberRequest) Reset() {
	*x = GetMaximumNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMaximumNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMaximumNumberRequest) ProtoMessage() {}

func (x *GetMaximumNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMaximumNumberRequest.ProtoReflect.Descriptor instead.
func (*GetMaximumNumberRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{6}
}

func (x *GetMaximumNumberRequest) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type GetMaximumNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaximumNum int64 `protobuf:"varint,1,opt,name=maximumNum,proto3" json:"maximumNum,omitempty"`
}

func (x *GetMaximumNumberResponse) Reset() {
	*x = GetMaximumNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMaximumNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMaximumNumberResponse) ProtoMessage() {}

func (x *GetMaximumNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMaximumNumberResponse.ProtoReflect.Descriptor instead.
func (*GetMaximumNumberResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{7}
}

func (x *GetMaximumNumberResponse) GetMaximumNum() int64 {
	if x != nil {
		return x.MaximumNum
	}
	return 0
}

var File_proto_calculator_proto protoreflect.FileDescriptor

var file_proto_calculator_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x75, 0x6d,
	0x32, 0x22, 0x22, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x73, 0x75, 0x6d, 0x22, 0x2a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6d,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75,
	0x6d, 0x22, 0x2f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x6d, 0x65, 0x22, 0x25, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x2b, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x3a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x78,
	0x69, 0x6d, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x4e, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x4e,
	0x75, 0x6d, 0x32, 0x8e, 0x02, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x6d, 0x12, 0x0e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6d,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x39, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x12, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x47, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x78, 0x69, 0x6d, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_calculator_proto_rawDescOnce sync.Once
	file_proto_calculator_proto_rawDescData = file_proto_calculator_proto_rawDesc
)

func file_proto_calculator_proto_rawDescGZIP() []byte {
	file_proto_calculator_proto_rawDescOnce.Do(func() {
		file_proto_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_calculator_proto_rawDescData)
	})
	return file_proto_calculator_proto_rawDescData
}

var file_proto_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_calculator_proto_goTypes = []interface{}{
	(*GetSumRequest)(nil),            // 0: GetSumRequest
	(*GetSumResponse)(nil),           // 1: GetSumResponse
	(*GetPrimeNumbersRequest)(nil),   // 2: GetPrimeNumbersRequest
	(*GetPrimeNumbersResponse)(nil),  // 3: GetPrimeNumbersResponse
	(*GetAverageRequest)(nil),        // 4: GetAverageRequest
	(*GetAverageResponse)(nil),       // 5: GetAverageResponse
	(*GetMaximumNumberRequest)(nil),  // 6: GetMaximumNumberRequest
	(*GetMaximumNumberResponse)(nil), // 7: GetMaximumNumberResponse
}
var file_proto_calculator_proto_depIdxs = []int32{
	0, // 0: CalculatorService.GetSum:input_type -> GetSumRequest
	2, // 1: CalculatorService.GetPrimeNumbers:input_type -> GetPrimeNumbersRequest
	4, // 2: CalculatorService.GetAverage:input_type -> GetAverageRequest
	6, // 3: CalculatorService.GetMaximum:input_type -> GetMaximumNumberRequest
	1, // 4: CalculatorService.GetSum:output_type -> GetSumResponse
	3, // 5: CalculatorService.GetPrimeNumbers:output_type -> GetPrimeNumbersResponse
	5, // 6: CalculatorService.GetAverage:output_type -> GetAverageResponse
	7, // 7: CalculatorService.GetMaximum:output_type -> GetMaximumNumberResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_calculator_proto_init() }
func file_proto_calculator_proto_init() {
	if File_proto_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSumRequest); i {
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
		file_proto_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSumResponse); i {
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
		file_proto_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrimeNumbersRequest); i {
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
		file_proto_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrimeNumbersResponse); i {
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
		file_proto_calculator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAverageRequest); i {
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
		file_proto_calculator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAverageResponse); i {
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
		file_proto_calculator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMaximumNumberRequest); i {
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
		file_proto_calculator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMaximumNumberResponse); i {
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
			RawDescriptor: file_proto_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_calculator_proto_goTypes,
		DependencyIndexes: file_proto_calculator_proto_depIdxs,
		MessageInfos:      file_proto_calculator_proto_msgTypes,
	}.Build()
	File_proto_calculator_proto = out.File
	file_proto_calculator_proto_rawDesc = nil
	file_proto_calculator_proto_goTypes = nil
	file_proto_calculator_proto_depIdxs = nil
}
