// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: deployment/APISync.proto

package deployment

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CommitObjectUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *CommitObjectUpdate) Reset() {
	*x = CommitObjectUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_APISync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitObjectUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitObjectUpdate) ProtoMessage() {}

func (x *CommitObjectUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_APISync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitObjectUpdate.ProtoReflect.Descriptor instead.
func (*CommitObjectUpdate) Descriptor() ([]byte, []int) {
	return file_deployment_APISync_proto_rawDescGZIP(), []int{0}
}

func (x *CommitObjectUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CommitObjectUpdate) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

type UpdateStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *UpdateStatus) Reset() {
	*x = UpdateStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_APISync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStatus) ProtoMessage() {}

func (x *UpdateStatus) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_APISync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStatus.ProtoReflect.Descriptor instead.
func (*UpdateStatus) Descriptor() ([]byte, []int) {
	return file_deployment_APISync_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateStatus) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateStatus) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

type DeleteStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteStatus) Reset() {
	*x = DeleteStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_APISync_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStatus) ProtoMessage() {}

func (x *DeleteStatus) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_APISync_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStatus.ProtoReflect.Descriptor instead.
func (*DeleteStatus) Descriptor() ([]byte, []int) {
	return file_deployment_APISync_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteStatus) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type APISyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*APISyncRequest_CommitObjectUpdate
	//	*APISyncRequest_UpdateStatus
	//	*APISyncRequest_DeleteStatus
	Request     isAPISyncRequest_Request `protobuf_oneof:"request"`
	PackageType string                   `protobuf:"bytes,4,opt,name=packageType,proto3" json:"packageType,omitempty"`
}

func (x *APISyncRequest) Reset() {
	*x = APISyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_APISync_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APISyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APISyncRequest) ProtoMessage() {}

func (x *APISyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_APISync_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APISyncRequest.ProtoReflect.Descriptor instead.
func (*APISyncRequest) Descriptor() ([]byte, []int) {
	return file_deployment_APISync_proto_rawDescGZIP(), []int{3}
}

func (m *APISyncRequest) GetRequest() isAPISyncRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *APISyncRequest) GetCommitObjectUpdate() *CommitObjectUpdate {
	if x, ok := x.GetRequest().(*APISyncRequest_CommitObjectUpdate); ok {
		return x.CommitObjectUpdate
	}
	return nil
}

func (x *APISyncRequest) GetUpdateStatus() *UpdateStatus {
	if x, ok := x.GetRequest().(*APISyncRequest_UpdateStatus); ok {
		return x.UpdateStatus
	}
	return nil
}

func (x *APISyncRequest) GetDeleteStatus() *DeleteStatus {
	if x, ok := x.GetRequest().(*APISyncRequest_DeleteStatus); ok {
		return x.DeleteStatus
	}
	return nil
}

func (x *APISyncRequest) GetPackageType() string {
	if x != nil {
		return x.PackageType
	}
	return ""
}

type isAPISyncRequest_Request interface {
	isAPISyncRequest_Request()
}

type APISyncRequest_CommitObjectUpdate struct {
	CommitObjectUpdate *CommitObjectUpdate `protobuf:"bytes,1,opt,name=commitObjectUpdate,proto3,oneof"`
}

type APISyncRequest_UpdateStatus struct {
	UpdateStatus *UpdateStatus `protobuf:"bytes,2,opt,name=updateStatus,proto3,oneof"`
}

type APISyncRequest_DeleteStatus struct {
	DeleteStatus *DeleteStatus `protobuf:"bytes,3,opt,name=deleteStatus,proto3,oneof"`
}

func (*APISyncRequest_CommitObjectUpdate) isAPISyncRequest_Request() {}

func (*APISyncRequest_UpdateStatus) isAPISyncRequest_Request() {}

func (*APISyncRequest_DeleteStatus) isAPISyncRequest_Request() {}

type ObjectUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *ObjectUpdated) Reset() {
	*x = ObjectUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_APISync_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectUpdated) ProtoMessage() {}

func (x *ObjectUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_APISync_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectUpdated.ProtoReflect.Descriptor instead.
func (*ObjectUpdated) Descriptor() ([]byte, []int) {
	return file_deployment_APISync_proto_rawDescGZIP(), []int{4}
}

func (x *ObjectUpdated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ObjectUpdated) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

type ObjectDeleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *ObjectDeleted) Reset() {
	*x = ObjectDeleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_APISync_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectDeleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectDeleted) ProtoMessage() {}

func (x *ObjectDeleted) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_APISync_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectDeleted.ProtoReflect.Descriptor instead.
func (*ObjectDeleted) Descriptor() ([]byte, []int) {
	return file_deployment_APISync_proto_rawDescGZIP(), []int{5}
}

func (x *ObjectDeleted) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ObjectDeleted) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

type APISyncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*APISyncResponse_ObjectUpdated
	//	*APISyncResponse_ObjectDeleted
	Response    isAPISyncResponse_Response `protobuf_oneof:"response"`
	Error       string                     `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	PackageType string                     `protobuf:"bytes,4,opt,name=packageType,proto3" json:"packageType,omitempty"`
}

func (x *APISyncResponse) Reset() {
	*x = APISyncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_APISync_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APISyncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APISyncResponse) ProtoMessage() {}

func (x *APISyncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_APISync_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APISyncResponse.ProtoReflect.Descriptor instead.
func (*APISyncResponse) Descriptor() ([]byte, []int) {
	return file_deployment_APISync_proto_rawDescGZIP(), []int{6}
}

func (m *APISyncResponse) GetResponse() isAPISyncResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *APISyncResponse) GetObjectUpdated() *ObjectUpdated {
	if x, ok := x.GetResponse().(*APISyncResponse_ObjectUpdated); ok {
		return x.ObjectUpdated
	}
	return nil
}

func (x *APISyncResponse) GetObjectDeleted() *ObjectDeleted {
	if x, ok := x.GetResponse().(*APISyncResponse_ObjectDeleted); ok {
		return x.ObjectDeleted
	}
	return nil
}

func (x *APISyncResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *APISyncResponse) GetPackageType() string {
	if x != nil {
		return x.PackageType
	}
	return ""
}

type isAPISyncResponse_Response interface {
	isAPISyncResponse_Response()
}

type APISyncResponse_ObjectUpdated struct {
	ObjectUpdated *ObjectUpdated `protobuf:"bytes,1,opt,name=objectUpdated,proto3,oneof"`
}

type APISyncResponse_ObjectDeleted struct {
	ObjectDeleted *ObjectDeleted `protobuf:"bytes,2,opt,name=objectDeleted,proto3,oneof"`
}

func (*APISyncResponse_ObjectUpdated) isAPISyncResponse_Response() {}

func (*APISyncResponse_ObjectDeleted) isAPISyncResponse_Response() {}

var File_deployment_APISync_proto protoreflect.FileDescriptor

var file_deployment_APISync_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x41, 0x50, 0x49,
	0x53, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x12, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x36, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x22, 0x1e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xee, 0x01, 0x0a, 0x0e, 0x41, 0x50, 0x49, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48,
	0x00, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x33, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x37, 0x0a, 0x0d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x37, 0x0a, 0x0d, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x22, 0xc5, 0x01, 0x0a, 0x0f, 0x41, 0x50, 0x49, 0x53, 0x79, 0x6e, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x36, 0x0a, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x38, 0x0a, 0x07, 0x41,
	0x50, 0x49, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x2d, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x0f,
	0x2e, 0x41, 0x50, 0x49, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x41, 0x50, 0x49, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x65, 0x72, 0x74, 0x61, 0x41, 0x49, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_deployment_APISync_proto_rawDescOnce sync.Once
	file_deployment_APISync_proto_rawDescData = file_deployment_APISync_proto_rawDesc
)

func file_deployment_APISync_proto_rawDescGZIP() []byte {
	file_deployment_APISync_proto_rawDescOnce.Do(func() {
		file_deployment_APISync_proto_rawDescData = protoimpl.X.CompressGZIP(file_deployment_APISync_proto_rawDescData)
	})
	return file_deployment_APISync_proto_rawDescData
}

var file_deployment_APISync_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_deployment_APISync_proto_goTypes = []interface{}{
	(*CommitObjectUpdate)(nil), // 0: CommitObjectUpdate
	(*UpdateStatus)(nil),       // 1: UpdateStatus
	(*DeleteStatus)(nil),       // 2: DeleteStatus
	(*APISyncRequest)(nil),     // 3: APISyncRequest
	(*ObjectUpdated)(nil),      // 4: ObjectUpdated
	(*ObjectDeleted)(nil),      // 5: ObjectDeleted
	(*APISyncResponse)(nil),    // 6: APISyncResponse
}
var file_deployment_APISync_proto_depIdxs = []int32{
	0, // 0: APISyncRequest.commitObjectUpdate:type_name -> CommitObjectUpdate
	1, // 1: APISyncRequest.updateStatus:type_name -> UpdateStatus
	2, // 2: APISyncRequest.deleteStatus:type_name -> DeleteStatus
	4, // 3: APISyncResponse.objectUpdated:type_name -> ObjectUpdated
	5, // 4: APISyncResponse.objectDeleted:type_name -> ObjectDeleted
	3, // 5: APISync.Sync:input_type -> APISyncRequest
	6, // 6: APISync.Sync:output_type -> APISyncResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_deployment_APISync_proto_init() }
func file_deployment_APISync_proto_init() {
	if File_deployment_APISync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deployment_APISync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitObjectUpdate); i {
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
		file_deployment_APISync_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStatus); i {
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
		file_deployment_APISync_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStatus); i {
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
		file_deployment_APISync_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APISyncRequest); i {
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
		file_deployment_APISync_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectUpdated); i {
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
		file_deployment_APISync_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectDeleted); i {
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
		file_deployment_APISync_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APISyncResponse); i {
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
	file_deployment_APISync_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*APISyncRequest_CommitObjectUpdate)(nil),
		(*APISyncRequest_UpdateStatus)(nil),
		(*APISyncRequest_DeleteStatus)(nil),
	}
	file_deployment_APISync_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*APISyncResponse_ObjectUpdated)(nil),
		(*APISyncResponse_ObjectDeleted)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_deployment_APISync_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deployment_APISync_proto_goTypes,
		DependencyIndexes: file_deployment_APISync_proto_depIdxs,
		MessageInfos:      file_deployment_APISync_proto_msgTypes,
	}.Build()
	File_deployment_APISync_proto = out.File
	file_deployment_APISync_proto_rawDesc = nil
	file_deployment_APISync_proto_goTypes = nil
	file_deployment_APISync_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// APISyncClient is the client API for APISync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APISyncClient interface {
	Sync(ctx context.Context, opts ...grpc.CallOption) (APISync_SyncClient, error)
}

type aPISyncClient struct {
	cc grpc.ClientConnInterface
}

func NewAPISyncClient(cc grpc.ClientConnInterface) APISyncClient {
	return &aPISyncClient{cc}
}

func (c *aPISyncClient) Sync(ctx context.Context, opts ...grpc.CallOption) (APISync_SyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &_APISync_serviceDesc.Streams[0], "/APISync/Sync", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPISyncSyncClient{stream}
	return x, nil
}

type APISync_SyncClient interface {
	Send(*APISyncRequest) error
	Recv() (*APISyncResponse, error)
	grpc.ClientStream
}

type aPISyncSyncClient struct {
	grpc.ClientStream
}

func (x *aPISyncSyncClient) Send(m *APISyncRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aPISyncSyncClient) Recv() (*APISyncResponse, error) {
	m := new(APISyncResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// APISyncServer is the server API for APISync service.
type APISyncServer interface {
	Sync(APISync_SyncServer) error
}

// UnimplementedAPISyncServer can be embedded to have forward compatible implementations.
type UnimplementedAPISyncServer struct {
}

func (*UnimplementedAPISyncServer) Sync(APISync_SyncServer) error {
	return status.Errorf(codes.Unimplemented, "method Sync not implemented")
}

func RegisterAPISyncServer(s *grpc.Server, srv APISyncServer) {
	s.RegisterService(&_APISync_serviceDesc, srv)
}

func _APISync_Sync_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(APISyncServer).Sync(&aPISyncSyncServer{stream})
}

type APISync_SyncServer interface {
	Send(*APISyncResponse) error
	Recv() (*APISyncRequest, error)
	grpc.ServerStream
}

type aPISyncSyncServer struct {
	grpc.ServerStream
}

func (x *aPISyncSyncServer) Send(m *APISyncResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aPISyncSyncServer) Recv() (*APISyncRequest, error) {
	m := new(APISyncRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _APISync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "APISync",
	HandlerType: (*APISyncServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sync",
			Handler:       _APISync_Sync_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "deployment/APISync.proto",
}
