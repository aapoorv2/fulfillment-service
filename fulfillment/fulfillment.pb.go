// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: fulfillment.proto

package fulfillment

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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	City     string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AssignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	City    string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *AssignRequest) Reset() {
	*x = AssignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignRequest) ProtoMessage() {}

func (x *AssignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignRequest.ProtoReflect.Descriptor instead.
func (*AssignRequest) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{2}
}

func (x *AssignRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *AssignRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type AssignResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AssignResponse) Reset() {
	*x = AssignResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignResponse) ProtoMessage() {}

func (x *AssignResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignResponse.ProtoReflect.Descriptor instead.
func (*AssignResponse) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{3}
}

func (x *AssignResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{4}
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FetchDeliveriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchDeliveriesRequest) Reset() {
	*x = FetchDeliveriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchDeliveriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchDeliveriesRequest) ProtoMessage() {}

func (x *FetchDeliveriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchDeliveriesRequest.ProtoReflect.Descriptor instead.
func (*FetchDeliveriesRequest) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{6}
}

type Delivery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId uint64 `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	City    string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Status  string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Delivery) Reset() {
	*x = Delivery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Delivery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delivery) ProtoMessage() {}

func (x *Delivery) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Delivery.ProtoReflect.Descriptor instead.
func (*Delivery) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{7}
}

func (x *Delivery) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Delivery) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Delivery) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Delivery) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type FetchDeliveriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deliveries []*Delivery `protobuf:"bytes,1,rep,name=deliveries,proto3" json:"deliveries,omitempty"`
}

func (x *FetchDeliveriesResponse) Reset() {
	*x = FetchDeliveriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fulfillment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchDeliveriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchDeliveriesResponse) ProtoMessage() {}

func (x *FetchDeliveriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fulfillment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchDeliveriesResponse.ProtoReflect.Descriptor instead.
func (*FetchDeliveriesResponse) Descriptor() ([]byte, []int) {
	return file_fulfillment_proto_rawDescGZIP(), []int{8}
}

func (x *FetchDeliveriesResponse) GetDeliveries() []*Delivery {
	if x != nil {
		return x.Deliveries
	}
	return nil
}

var File_fulfillment_proto protoreflect.FileDescriptor

var file_fulfillment_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x5d, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x22,
	0x2c, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3e, 0x0a,
	0x0d, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x22, 0x2a, 0x0a,
	0x0e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2a, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x61, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x50, 0x0a, 0x17, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x32, 0xfa, 0x02, 0x0a, 0x0b, 0x46, 0x75, 0x6c, 0x46, 0x69, 0x6c,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x13, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x66,
	0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x75, 0x6c, 0x66, 0x69,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1c,
	0x2e, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x66,
	0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x1f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1a,
	0x2e, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x75, 0x6c,
	0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x1c, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x41, 0x6c, 0x6c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x46, 0x6f, 0x72,
	0x41, 0x6e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c,
	0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x66,
	0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x19, 0x5a, 0x17, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x66, 0x75, 0x6c, 0x66, 0x69, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fulfillment_proto_rawDescOnce sync.Once
	file_fulfillment_proto_rawDescData = file_fulfillment_proto_rawDesc
)

func file_fulfillment_proto_rawDescGZIP() []byte {
	file_fulfillment_proto_rawDescOnce.Do(func() {
		file_fulfillment_proto_rawDescData = protoimpl.X.CompressGZIP(file_fulfillment_proto_rawDescData)
	})
	return file_fulfillment_proto_rawDescData
}

var file_fulfillment_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_fulfillment_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),         // 0: fulfillment.RegisterRequest
	(*RegisterResponse)(nil),        // 1: fulfillment.RegisterResponse
	(*AssignRequest)(nil),           // 2: fulfillment.AssignRequest
	(*AssignResponse)(nil),          // 3: fulfillment.AssignResponse
	(*UpdateRequest)(nil),           // 4: fulfillment.UpdateRequest
	(*UpdateResponse)(nil),          // 5: fulfillment.UpdateResponse
	(*FetchDeliveriesRequest)(nil),  // 6: fulfillment.FetchDeliveriesRequest
	(*Delivery)(nil),                // 7: fulfillment.Delivery
	(*FetchDeliveriesResponse)(nil), // 8: fulfillment.FetchDeliveriesResponse
}
var file_fulfillment_proto_depIdxs = []int32{
	7, // 0: fulfillment.FetchDeliveriesResponse.deliveries:type_name -> fulfillment.Delivery
	2, // 1: fulfillment.FulFillment.AssignDeliveryAgent:input_type -> fulfillment.AssignRequest
	0, // 2: fulfillment.FulFillment.RegisterDeliveryAgent:input_type -> fulfillment.RegisterRequest
	4, // 3: fulfillment.FulFillment.UpdateDeliveryAgentAvailability:input_type -> fulfillment.UpdateRequest
	6, // 4: fulfillment.FulFillment.FetchAllDeliveriesForAnAgent:input_type -> fulfillment.FetchDeliveriesRequest
	3, // 5: fulfillment.FulFillment.AssignDeliveryAgent:output_type -> fulfillment.AssignResponse
	1, // 6: fulfillment.FulFillment.RegisterDeliveryAgent:output_type -> fulfillment.RegisterResponse
	5, // 7: fulfillment.FulFillment.UpdateDeliveryAgentAvailability:output_type -> fulfillment.UpdateResponse
	8, // 8: fulfillment.FulFillment.FetchAllDeliveriesForAnAgent:output_type -> fulfillment.FetchDeliveriesResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_fulfillment_proto_init() }
func file_fulfillment_proto_init() {
	if File_fulfillment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fulfillment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_fulfillment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_fulfillment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignRequest); i {
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
		file_fulfillment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignResponse); i {
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
		file_fulfillment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_fulfillment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_fulfillment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchDeliveriesRequest); i {
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
		file_fulfillment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Delivery); i {
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
		file_fulfillment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchDeliveriesResponse); i {
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
			RawDescriptor: file_fulfillment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fulfillment_proto_goTypes,
		DependencyIndexes: file_fulfillment_proto_depIdxs,
		MessageInfos:      file_fulfillment_proto_msgTypes,
	}.Build()
	File_fulfillment_proto = out.File
	file_fulfillment_proto_rawDesc = nil
	file_fulfillment_proto_goTypes = nil
	file_fulfillment_proto_depIdxs = nil
}
