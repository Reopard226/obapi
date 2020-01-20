// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/iam/service.proto

package iam

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// User object used to retrieve keys
type User struct {
	Auth0UserId          string   `protobuf:"bytes,1,opt,name=auth0UserId,proto3" json:"auth0UserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_b572537561b6c5b6, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetAuth0UserId() string {
	if m != nil {
		return m.Auth0UserId
	}
	return ""
}

type UserKey struct {
	Expires              int64    `protobuf:"varint,1,opt,name=expires,proto3" json:"expires,omitempty" bson:"expires,omitempty"`
	ApikeyId             string   `protobuf:"bytes,2,opt,name=apikey_id,json=apikeyId,proto3" json:"apikey_id,omitempty" bson:"apikey_id,omitempty"`
	KeyTag               string   `protobuf:"bytes,3,opt,name=key_tag,json=keyTag,proto3" json:"key_tag,omitempty" bson:"key_tag,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" bson:"user_id,omitempty"`
	SigningKeyId         string   `protobuf:"bytes,5,opt,name=signing_key_id,json=signingKeyId,proto3" json:"signing_key_id,omitempty" bson:"signing_key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserKey) Reset()         { *m = UserKey{} }
func (m *UserKey) String() string { return proto.CompactTextString(m) }
func (*UserKey) ProtoMessage()    {}
func (*UserKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_b572537561b6c5b6, []int{1}
}

func (m *UserKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserKey.Unmarshal(m, b)
}
func (m *UserKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserKey.Marshal(b, m, deterministic)
}
func (m *UserKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserKey.Merge(m, src)
}
func (m *UserKey) XXX_Size() int {
	return xxx_messageInfo_UserKey.Size(m)
}
func (m *UserKey) XXX_DiscardUnknown() {
	xxx_messageInfo_UserKey.DiscardUnknown(m)
}

var xxx_messageInfo_UserKey proto.InternalMessageInfo

func (m *UserKey) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (m *UserKey) GetApikeyId() string {
	if m != nil {
		return m.ApikeyId
	}
	return ""
}

func (m *UserKey) GetKeyTag() string {
	if m != nil {
		return m.KeyTag
	}
	return ""
}

func (m *UserKey) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserKey) GetSigningKeyId() string {
	if m != nil {
		return m.SigningKeyId
	}
	return ""
}

type UserKeys struct {
	NumberOfKeys         int64      `protobuf:"varint,1,opt,name=number_of_keys,json=numberOfKeys,proto3" json:"number_of_keys,omitempty"`
	Keys                 []*UserKey `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-" bson:"-"`
	XXX_unrecognized     []byte     `json:"-" bson:"-"`
	XXX_sizecache        int32      `json:"-" bson:"-"`
}

func (m *UserKeys) Reset()         { *m = UserKeys{} }
func (m *UserKeys) String() string { return proto.CompactTextString(m) }
func (*UserKeys) ProtoMessage()    {}
func (*UserKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_b572537561b6c5b6, []int{2}
}

func (m *UserKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserKeys.Unmarshal(m, b)
}
func (m *UserKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserKeys.Marshal(b, m, deterministic)
}
func (m *UserKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserKeys.Merge(m, src)
}
func (m *UserKeys) XXX_Size() int {
	return xxx_messageInfo_UserKeys.Size(m)
}
func (m *UserKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_UserKeys.DiscardUnknown(m)
}

var xxx_messageInfo_UserKeys proto.InternalMessageInfo

func (m *UserKeys) GetNumberOfKeys() int64 {
	if m != nil {
		return m.NumberOfKeys
	}
	return 0
}

func (m *UserKeys) GetKeys() []*UserKey {
	if m != nil {
		return m.Keys
	}
	return nil
}

type CreateKeyRequest struct {
	Expires              int64    `protobuf:"varint,1,opt,name=expires,proto3" json:"expires,omitempty"`
	KeyTag               string   `protobuf:"bytes,2,opt,name=key_tag,json=keyTag,proto3" json:"key_tag,omitempty"`
	UserId               string   `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *CreateKeyRequest) Reset()         { *m = CreateKeyRequest{} }
func (m *CreateKeyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateKeyRequest) ProtoMessage()    {}
func (*CreateKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b572537561b6c5b6, []int{3}
}

func (m *CreateKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateKeyRequest.Unmarshal(m, b)
}
func (m *CreateKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateKeyRequest.Marshal(b, m, deterministic)
}
func (m *CreateKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateKeyRequest.Merge(m, src)
}
func (m *CreateKeyRequest) XXX_Size() int {
	return xxx_messageInfo_CreateKeyRequest.Size(m)
}
func (m *CreateKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateKeyRequest proto.InternalMessageInfo

func (m *CreateKeyRequest) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (m *CreateKeyRequest) GetKeyTag() string {
	if m != nil {
		return m.KeyTag
	}
	return ""
}

func (m *CreateKeyRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type UserKeyWithSecret struct {
	Expires              int64    `protobuf:"varint,1,opt,name=expires,proto3" json:"expires,omitempty"`
	ApikeyId             string   `protobuf:"bytes,2,opt,name=apikey_id,json=apikeyId,proto3" json:"apikey_id,omitempty"`
	KeyTag               string   `protobuf:"bytes,3,opt,name=key_tag,json=keyTag,proto3" json:"key_tag,omitempty"`
	ApikeySecret         string   `protobuf:"bytes,4,opt,name=apikey_secret,json=apikeySecret,proto3" json:"apikey_secret,omitempty"`
	UserId               string   `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *UserKeyWithSecret) Reset()         { *m = UserKeyWithSecret{} }
func (m *UserKeyWithSecret) String() string { return proto.CompactTextString(m) }
func (*UserKeyWithSecret) ProtoMessage()    {}
func (*UserKeyWithSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_b572537561b6c5b6, []int{4}
}

func (m *UserKeyWithSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserKeyWithSecret.Unmarshal(m, b)
}
func (m *UserKeyWithSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserKeyWithSecret.Marshal(b, m, deterministic)
}
func (m *UserKeyWithSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserKeyWithSecret.Merge(m, src)
}
func (m *UserKeyWithSecret) XXX_Size() int {
	return xxx_messageInfo_UserKeyWithSecret.Size(m)
}
func (m *UserKeyWithSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_UserKeyWithSecret.DiscardUnknown(m)
}

var xxx_messageInfo_UserKeyWithSecret proto.InternalMessageInfo

func (m *UserKeyWithSecret) GetExpires() int64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (m *UserKeyWithSecret) GetApikeyId() string {
	if m != nil {
		return m.ApikeyId
	}
	return ""
}

func (m *UserKeyWithSecret) GetKeyTag() string {
	if m != nil {
		return m.KeyTag
	}
	return ""
}

func (m *UserKeyWithSecret) GetApikeySecret() string {
	if m != nil {
		return m.ApikeySecret
	}
	return ""
}

func (m *UserKeyWithSecret) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type DeleteKeyRequest struct {
	ApikeyId             string   `protobuf:"bytes,1,opt,name=apikey_id,json=apikeyId,proto3" json:"apikey_id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *DeleteKeyRequest) Reset()         { *m = DeleteKeyRequest{} }
func (m *DeleteKeyRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteKeyRequest) ProtoMessage()    {}
func (*DeleteKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b572537561b6c5b6, []int{5}
}

func (m *DeleteKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteKeyRequest.Unmarshal(m, b)
}
func (m *DeleteKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteKeyRequest.Marshal(b, m, deterministic)
}
func (m *DeleteKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteKeyRequest.Merge(m, src)
}
func (m *DeleteKeyRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteKeyRequest.Size(m)
}
func (m *DeleteKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteKeyRequest proto.InternalMessageInfo

func (m *DeleteKeyRequest) GetApikeyId() string {
	if m != nil {
		return m.ApikeyId
	}
	return ""
}

func (m *DeleteKeyRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type KeyDeletedResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *KeyDeletedResponse) Reset()         { *m = KeyDeletedResponse{} }
func (m *KeyDeletedResponse) String() string { return proto.CompactTextString(m) }
func (*KeyDeletedResponse) ProtoMessage()    {}
func (*KeyDeletedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b572537561b6c5b6, []int{6}
}

func (m *KeyDeletedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyDeletedResponse.Unmarshal(m, b)
}
func (m *KeyDeletedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyDeletedResponse.Marshal(b, m, deterministic)
}
func (m *KeyDeletedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyDeletedResponse.Merge(m, src)
}
func (m *KeyDeletedResponse) XXX_Size() int {
	return xxx_messageInfo_KeyDeletedResponse.Size(m)
}
func (m *KeyDeletedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyDeletedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KeyDeletedResponse proto.InternalMessageInfo

func (m *KeyDeletedResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "oceanbolt.com.iam.User")
	proto.RegisterType((*UserKey)(nil), "oceanbolt.com.iam.UserKey")
	proto.RegisterType((*UserKeys)(nil), "oceanbolt.com.iam.UserKeys")
	proto.RegisterType((*CreateKeyRequest)(nil), "oceanbolt.com.iam.CreateKeyRequest")
	proto.RegisterType((*UserKeyWithSecret)(nil), "oceanbolt.com.iam.UserKeyWithSecret")
	proto.RegisterType((*DeleteKeyRequest)(nil), "oceanbolt.com.iam.DeleteKeyRequest")
	proto.RegisterType((*KeyDeletedResponse)(nil), "oceanbolt.com.iam.KeyDeletedResponse")
}

func init() { proto.RegisterFile("rpc/iam/service.proto", fileDescriptor_b572537561b6c5b6) }

var fileDescriptor_b572537561b6c5b6 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4b, 0x6b, 0x13, 0x51,
	0x14, 0x26, 0x99, 0xe6, 0x75, 0x1a, 0x4b, 0x7b, 0x45, 0x12, 0xa6, 0x48, 0xc3, 0x34, 0x42, 0x16,
	0x32, 0x91, 0x0a, 0x0a, 0xe2, 0x42, 0xa3, 0x0b, 0x6b, 0x04, 0x61, 0x7c, 0x14, 0x5c, 0x18, 0x6f,
	0x92, 0xd3, 0xe9, 0xa5, 0x9d, 0x87, 0xf7, 0xde, 0x88, 0xd9, 0xfb, 0x2b, 0x5c, 0x08, 0xfe, 0x0f,
	0xff, 0x99, 0x1b, 0xb9, 0x8f, 0x4c, 0x66, 0x32, 0x89, 0xab, 0x64, 0xce, 0xf9, 0x1e, 0x67, 0xbe,
	0x0f, 0x06, 0xee, 0xf0, 0x74, 0x36, 0x64, 0x34, 0x1a, 0x0a, 0xe4, 0xdf, 0xd8, 0x0c, 0xfd, 0x94,
	0x27, 0x32, 0x21, 0x47, 0xc9, 0x0c, 0x69, 0x3c, 0x4d, 0x6e, 0xa4, 0x3f, 0x4b, 0x22, 0x9f, 0xd1,
	0xc8, 0xbd, 0x2d, 0x69, 0x18, 0x22, 0x1f, 0x9a, 0x1f, 0x83, 0xf3, 0x06, 0xb0, 0xf7, 0x41, 0x20,
	0x27, 0x3d, 0xd8, 0xa7, 0x0b, 0x79, 0xf5, 0x40, 0x3d, 0x9c, 0xcf, 0xbb, 0x95, 0x5e, 0x65, 0xd0,
	0x0a, 0xf2, 0x23, 0xef, 0x4f, 0x15, 0x1a, 0xea, 0xef, 0x18, 0x97, 0xe4, 0x31, 0x34, 0xf0, 0x7b,
	0xca, 0x38, 0x0a, 0x8d, 0x74, 0x46, 0x77, 0x7f, 0xfe, 0xf8, 0xe5, 0x74, 0xa7, 0x22, 0x89, 0x9f,
	0x78, 0x76, 0x73, 0x3f, 0x89, 0x98, 0xc4, 0x28, 0x95, 0x4b, 0x2f, 0x58, 0xa1, 0xc9, 0x53, 0x68,
	0xd1, 0x94, 0x5d, 0xe3, 0x72, 0xc2, 0xe6, 0xdd, 0xaa, 0x32, 0x19, 0x9d, 0x28, 0xaa, 0x6b, 0xa8,
	0xd9, 0x2e, 0x4f, 0x6e, 0x9a, 0xe9, 0xf9, 0x9c, 0x3c, 0x82, 0x86, 0x5a, 0x4b, 0x1a, 0x76, 0x1d,
	0xcd, 0xcd, 0xdb, 0xda, 0x4d, 0x9e, 0x59, 0xbf, 0xc6, 0xe5, 0x7b, 0x1a, 0x2a, 0xde, 0x42, 0x20,
	0x57, 0x9e, 0x7b, 0x25, 0x9e, 0xdd, 0x14, 0x78, 0x0b, 0xfd, 0xca, 0xe4, 0x35, 0x1c, 0x08, 0x16,
	0xc6, 0x2c, 0x0e, 0x27, 0xf6, 0xe4, 0x9a, 0xa6, 0xf7, 0x15, 0xfd, 0xc4, 0xd0, 0x8b, 0x80, 0xbc,
	0x4a, 0xdb, 0xae, 0xc6, 0xea, 0x76, 0xef, 0x0b, 0x34, 0x6d, 0x7a, 0x82, 0xf4, 0xe1, 0x20, 0x5e,
	0x44, 0x53, 0xe4, 0x93, 0xe4, 0x52, 0x11, 0x6d, 0x8a, 0x41, 0xdb, 0x4c, 0xdf, 0x5e, 0x6a, 0x94,
	0x0f, 0x7b, 0x7a, 0x57, 0xed, 0x39, 0x83, 0xfd, 0x33, 0xd7, 0x2f, 0x35, 0xea, 0x5b, 0xc1, 0x40,
	0xe3, 0xbc, 0xcf, 0x70, 0xf8, 0x82, 0x23, 0x95, 0xa8, 0x46, 0xf8, 0x75, 0x81, 0x42, 0x92, 0xee,
	0x46, 0x51, 0xeb, 0x26, 0x3a, 0xeb, 0x2c, 0x75, 0x0f, 0x59, 0x58, 0x9d, 0x75, 0x58, 0x8e, 0x59,
	0x98, 0x34, 0xbc, 0xdf, 0x15, 0x38, 0xb2, 0x8e, 0x17, 0x4c, 0x5e, 0xbd, 0xc3, 0x19, 0xc7, 0xff,
	0x39, 0x1c, 0x97, 0xba, 0xce, 0x55, 0xd9, 0xd9, 0xa8, 0x32, 0xb3, 0x3f, 0x85, 0x5b, 0x96, 0x25,
	0xb4, 0x81, 0x69, 0x2c, 0x68, 0x9b, 0xa1, 0x35, 0xcd, 0xdd, 0x58, 0x2b, 0xdc, 0xf8, 0x0a, 0x0e,
	0x5f, 0xe2, 0x0d, 0x16, 0x32, 0x28, 0xdc, 0x51, 0x29, 0xdf, 0xb1, 0x52, 0xaa, 0x16, 0x94, 0x7c,
	0x20, 0x63, 0x5c, 0x1a, 0xb1, 0x79, 0x80, 0x22, 0x4d, 0x62, 0x81, 0xea, 0x6d, 0x23, 0x14, 0x82,
	0x86, 0x68, 0x95, 0x56, 0x8f, 0x67, 0x7f, 0x2b, 0x50, 0x7f, 0xae, 0x55, 0xc9, 0x33, 0x68, 0xbe,
	0x61, 0x42, 0xea, 0x12, 0x3b, 0x3b, 0x6a, 0x73, 0x8f, 0x77, 0xf7, 0x29, 0xc8, 0x47, 0x68, 0x65,
	0x55, 0x92, 0xd3, 0x2d, 0xc8, 0xcd, 0xa2, 0xdd, 0xfe, 0x6e, 0xb9, 0x5c, 0x59, 0x17, 0xd0, 0xca,
	0xe2, 0xd9, 0xaa, 0xbb, 0x19, 0x9e, 0x7b, 0x6f, 0x0b, 0xa8, 0x9c, 0xcb, 0xa8, 0xf6, 0xc9, 0x61,
	0x34, 0x9a, 0xd6, 0xf5, 0x47, 0xe5, 0xe1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0xb8, 0x70,
	0x20, 0x95, 0x04, 0x00, 0x00,
}
