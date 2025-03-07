// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: proto/vault/v1/service.proto

package vaultv1

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

type LivenessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LivenessRequest) Reset() {
	*x = LivenessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LivenessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LivenessRequest) ProtoMessage() {}

func (x *LivenessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LivenessRequest.ProtoReflect.Descriptor instead.
func (*LivenessRequest) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_service_proto_rawDescGZIP(), []int{0}
}

type LivenessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *LivenessResponse) Reset() {
	*x = LivenessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LivenessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LivenessResponse) ProtoMessage() {}

func (x *LivenessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LivenessResponse.ProtoReflect.Descriptor instead.
func (*LivenessResponse) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *LivenessResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type EncryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyring string `protobuf:"bytes,1,opt,name=keyring,proto3" json:"keyring,omitempty"`
	Data    string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EncryptRequest) Reset() {
	*x = EncryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptRequest) ProtoMessage() {}

func (x *EncryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptRequest.ProtoReflect.Descriptor instead.
func (*EncryptRequest) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *EncryptRequest) GetKeyring() string {
	if x != nil {
		return x.Keyring
	}
	return ""
}

func (x *EncryptRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type EncryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encrypted string `protobuf:"bytes,1,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	KeyId     string `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
}

func (x *EncryptResponse) Reset() {
	*x = EncryptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptResponse) ProtoMessage() {}

func (x *EncryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptResponse.ProtoReflect.Descriptor instead.
func (*EncryptResponse) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *EncryptResponse) GetEncrypted() string {
	if x != nil {
		return x.Encrypted
	}
	return ""
}

func (x *EncryptResponse) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

type EncryptBulkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyring string   `protobuf:"bytes,1,opt,name=keyring,proto3" json:"keyring,omitempty"`
	Data    []string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *EncryptBulkRequest) Reset() {
	*x = EncryptBulkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptBulkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptBulkRequest) ProtoMessage() {}

func (x *EncryptBulkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptBulkRequest.ProtoReflect.Descriptor instead.
func (*EncryptBulkRequest) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *EncryptBulkRequest) GetKeyring() string {
	if x != nil {
		return x.Keyring
	}
	return ""
}

func (x *EncryptBulkRequest) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

type EncryptBulkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encrypted []*EncryptResponse `protobuf:"bytes,1,rep,name=encrypted,proto3" json:"encrypted,omitempty"`
}

func (x *EncryptBulkResponse) Reset() {
	*x = EncryptBulkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptBulkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptBulkResponse) ProtoMessage() {}

func (x *EncryptBulkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptBulkResponse.ProtoReflect.Descriptor instead.
func (*EncryptBulkResponse) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *EncryptBulkResponse) GetEncrypted() []*EncryptResponse {
	if x != nil {
		return x.Encrypted
	}
	return nil
}

type DecryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyring   string `protobuf:"bytes,1,opt,name=keyring,proto3" json:"keyring,omitempty"`
	Encrypted string `protobuf:"bytes,2,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
}

func (x *DecryptRequest) Reset() {
	*x = DecryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptRequest) ProtoMessage() {}

func (x *DecryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptRequest.ProtoReflect.Descriptor instead.
func (*DecryptRequest) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *DecryptRequest) GetKeyring() string {
	if x != nil {
		return x.Keyring
	}
	return ""
}

func (x *DecryptRequest) GetEncrypted() string {
	if x != nil {
		return x.Encrypted
	}
	return ""
}

type DecryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plaintext string `protobuf:"bytes,1,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
}

func (x *DecryptResponse) Reset() {
	*x = DecryptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptResponse) ProtoMessage() {}

func (x *DecryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptResponse.ProtoReflect.Descriptor instead.
func (*DecryptResponse) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *DecryptResponse) GetPlaintext() string {
	if x != nil {
		return x.Plaintext
	}
	return ""
}

type CreateDEKRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyring string `protobuf:"bytes,1,opt,name=keyring,proto3" json:"keyring,omitempty"`
}

func (x *CreateDEKRequest) Reset() {
	*x = CreateDEKRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDEKRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDEKRequest) ProtoMessage() {}

func (x *CreateDEKRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDEKRequest.ProtoReflect.Descriptor instead.
func (*CreateDEKRequest) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *CreateDEKRequest) GetKeyring() string {
	if x != nil {
		return x.Keyring
	}
	return ""
}

type CreateDEKResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
}

func (x *CreateDEKResponse) Reset() {
	*x = CreateDEKResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDEKResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDEKResponse) ProtoMessage() {}

func (x *CreateDEKResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDEKResponse.ProtoReflect.Descriptor instead.
func (*CreateDEKResponse) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_service_proto_rawDescGZIP(), []int{9}
}

func (x *CreateDEKResponse) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

type ReEncryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyring   string `protobuf:"bytes,1,opt,name=keyring,proto3" json:"keyring,omitempty"`
	Encrypted string `protobuf:"bytes,2,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	// Specify the key_id to use for re-encryption. If not provided, the latest will be used
	KeyId *string `protobuf:"bytes,3,opt,name=key_id,json=keyId,proto3,oneof" json:"key_id,omitempty"`
}

func (x *ReEncryptRequest) Reset() {
	*x = ReEncryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReEncryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReEncryptRequest) ProtoMessage() {}

func (x *ReEncryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReEncryptRequest.ProtoReflect.Descriptor instead.
func (*ReEncryptRequest) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_service_proto_rawDescGZIP(), []int{10}
}

func (x *ReEncryptRequest) GetKeyring() string {
	if x != nil {
		return x.Keyring
	}
	return ""
}

func (x *ReEncryptRequest) GetEncrypted() string {
	if x != nil {
		return x.Encrypted
	}
	return ""
}

func (x *ReEncryptRequest) GetKeyId() string {
	if x != nil && x.KeyId != nil {
		return *x.KeyId
	}
	return ""
}

type ReEncryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Encrypted string `protobuf:"bytes,1,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	KeyId     string `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
}

func (x *ReEncryptResponse) Reset() {
	*x = ReEncryptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReEncryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReEncryptResponse) ProtoMessage() {}

func (x *ReEncryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReEncryptResponse.ProtoReflect.Descriptor instead.
func (*ReEncryptResponse) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_service_proto_rawDescGZIP(), []int{11}
}

func (x *ReEncryptResponse) GetEncrypted() string {
	if x != nil {
		return x.Encrypted
	}
	return ""
}

func (x *ReEncryptResponse) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

type ReEncryptDEKsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReEncryptDEKsRequest) Reset() {
	*x = ReEncryptDEKsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReEncryptDEKsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReEncryptDEKsRequest) ProtoMessage() {}

func (x *ReEncryptDEKsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReEncryptDEKsRequest.ProtoReflect.Descriptor instead.
func (*ReEncryptDEKsRequest) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_service_proto_rawDescGZIP(), []int{12}
}

type ReEncryptDEKsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReEncryptDEKsResponse) Reset() {
	*x = ReEncryptDEKsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_service_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReEncryptDEKsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReEncryptDEKsResponse) ProtoMessage() {}

func (x *ReEncryptDEKsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_service_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReEncryptDEKsResponse.ProtoReflect.Descriptor instead.
func (*ReEncryptDEKsResponse) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_service_proto_rawDescGZIP(), []int{13}
}

var File_proto_vault_v1_service_proto protoreflect.FileDescriptor

var file_proto_vault_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x76, 0x65,
	0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2a, 0x0a, 0x10, 0x4c,
	0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3e, 0x0a, 0x0e, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x46, 0x0a, 0x0f, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x22,
	0x42, 0x0a, 0x12, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x4e, 0x0a, 0x13, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x42, 0x75,
	0x6c, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x65, 0x64, 0x22, 0x48, 0x0a, 0x0e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x22, 0x2f, 0x0a,
	0x0f, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x2c,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x45, 0x4b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2a, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x45, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x22, 0x71, 0x0a, 0x10, 0x52, 0x65, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x11, 0x52,
	0x65, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6b, 0x65, 0x79, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x44, 0x45, 0x4b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a,
	0x15, 0x52, 0x65, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x45, 0x4b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x89, 0x04, 0x0a, 0x0c, 0x56, 0x61, 0x75, 0x6c, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x4c, 0x69, 0x76, 0x65, 0x6e,
	0x65, 0x73, 0x73, 0x12, 0x19, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x45, 0x4b, 0x12, 0x1a, 0x2e, 0x76, 0x61, 0x75, 0x6c,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x45, 0x4b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x45, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12,
	0x18, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x61, 0x75, 0x6c,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x42, 0x75, 0x6c, 0x6b, 0x12, 0x1c, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12,
	0x18, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x61, 0x75, 0x6c,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x09, 0x52, 0x65, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x12, 0x1a, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52,
	0x0a, 0x0d, 0x52, 0x65, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x45, 0x4b, 0x73, 0x12,
	0x1e, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x44, 0x45, 0x4b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x44, 0x45, 0x4b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x75, 0x6e, 0x6b, 0x65, 0x79, 0x65, 0x64, 0x2f, 0x75, 0x6e, 0x6b, 0x65, 0x79, 0x2f, 0x61,
	0x70, 0x70, 0x73, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x61, 0x75,
	0x6c, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_vault_v1_service_proto_rawDescOnce sync.Once
	file_proto_vault_v1_service_proto_rawDescData = file_proto_vault_v1_service_proto_rawDesc
)

func file_proto_vault_v1_service_proto_rawDescGZIP() []byte {
	file_proto_vault_v1_service_proto_rawDescOnce.Do(func() {
		file_proto_vault_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_vault_v1_service_proto_rawDescData)
	})
	return file_proto_vault_v1_service_proto_rawDescData
}

var file_proto_vault_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_proto_vault_v1_service_proto_goTypes = []any{
	(*LivenessRequest)(nil),       // 0: vault.v1.LivenessRequest
	(*LivenessResponse)(nil),      // 1: vault.v1.LivenessResponse
	(*EncryptRequest)(nil),        // 2: vault.v1.EncryptRequest
	(*EncryptResponse)(nil),       // 3: vault.v1.EncryptResponse
	(*EncryptBulkRequest)(nil),    // 4: vault.v1.EncryptBulkRequest
	(*EncryptBulkResponse)(nil),   // 5: vault.v1.EncryptBulkResponse
	(*DecryptRequest)(nil),        // 6: vault.v1.DecryptRequest
	(*DecryptResponse)(nil),       // 7: vault.v1.DecryptResponse
	(*CreateDEKRequest)(nil),      // 8: vault.v1.CreateDEKRequest
	(*CreateDEKResponse)(nil),     // 9: vault.v1.CreateDEKResponse
	(*ReEncryptRequest)(nil),      // 10: vault.v1.ReEncryptRequest
	(*ReEncryptResponse)(nil),     // 11: vault.v1.ReEncryptResponse
	(*ReEncryptDEKsRequest)(nil),  // 12: vault.v1.ReEncryptDEKsRequest
	(*ReEncryptDEKsResponse)(nil), // 13: vault.v1.ReEncryptDEKsResponse
}
var file_proto_vault_v1_service_proto_depIdxs = []int32{
	3,  // 0: vault.v1.EncryptBulkResponse.encrypted:type_name -> vault.v1.EncryptResponse
	0,  // 1: vault.v1.VaultService.Liveness:input_type -> vault.v1.LivenessRequest
	8,  // 2: vault.v1.VaultService.CreateDEK:input_type -> vault.v1.CreateDEKRequest
	2,  // 3: vault.v1.VaultService.Encrypt:input_type -> vault.v1.EncryptRequest
	4,  // 4: vault.v1.VaultService.EncryptBulk:input_type -> vault.v1.EncryptBulkRequest
	6,  // 5: vault.v1.VaultService.Decrypt:input_type -> vault.v1.DecryptRequest
	10, // 6: vault.v1.VaultService.ReEncrypt:input_type -> vault.v1.ReEncryptRequest
	12, // 7: vault.v1.VaultService.ReEncryptDEKs:input_type -> vault.v1.ReEncryptDEKsRequest
	1,  // 8: vault.v1.VaultService.Liveness:output_type -> vault.v1.LivenessResponse
	9,  // 9: vault.v1.VaultService.CreateDEK:output_type -> vault.v1.CreateDEKResponse
	3,  // 10: vault.v1.VaultService.Encrypt:output_type -> vault.v1.EncryptResponse
	5,  // 11: vault.v1.VaultService.EncryptBulk:output_type -> vault.v1.EncryptBulkResponse
	7,  // 12: vault.v1.VaultService.Decrypt:output_type -> vault.v1.DecryptResponse
	11, // 13: vault.v1.VaultService.ReEncrypt:output_type -> vault.v1.ReEncryptResponse
	13, // 14: vault.v1.VaultService.ReEncryptDEKs:output_type -> vault.v1.ReEncryptDEKsResponse
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_vault_v1_service_proto_init() }
func file_proto_vault_v1_service_proto_init() {
	if File_proto_vault_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_vault_v1_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LivenessRequest); i {
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
		file_proto_vault_v1_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LivenessResponse); i {
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
		file_proto_vault_v1_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EncryptRequest); i {
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
		file_proto_vault_v1_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*EncryptResponse); i {
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
		file_proto_vault_v1_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*EncryptBulkRequest); i {
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
		file_proto_vault_v1_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*EncryptBulkResponse); i {
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
		file_proto_vault_v1_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DecryptRequest); i {
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
		file_proto_vault_v1_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DecryptResponse); i {
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
		file_proto_vault_v1_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*CreateDEKRequest); i {
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
		file_proto_vault_v1_service_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*CreateDEKResponse); i {
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
		file_proto_vault_v1_service_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ReEncryptRequest); i {
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
		file_proto_vault_v1_service_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*ReEncryptResponse); i {
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
		file_proto_vault_v1_service_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*ReEncryptDEKsRequest); i {
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
		file_proto_vault_v1_service_proto_msgTypes[13].Exporter = func(v any, i int) any {
			switch v := v.(*ReEncryptDEKsResponse); i {
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
	file_proto_vault_v1_service_proto_msgTypes[10].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_vault_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_vault_v1_service_proto_goTypes,
		DependencyIndexes: file_proto_vault_v1_service_proto_depIdxs,
		MessageInfos:      file_proto_vault_v1_service_proto_msgTypes,
	}.Build()
	File_proto_vault_v1_service_proto = out.File
	file_proto_vault_v1_service_proto_rawDesc = nil
	file_proto_vault_v1_service_proto_goTypes = nil
	file_proto_vault_v1_service_proto_depIdxs = nil
}
