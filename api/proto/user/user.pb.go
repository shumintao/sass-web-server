// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type RegistRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	NickName             string   `protobuf:"bytes,3,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	PasswordAgain        string   `protobuf:"bytes,5,opt,name=password_again,json=passwordAgain,proto3" json:"password_again,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistRequest) Reset()         { *m = RegistRequest{} }
func (m *RegistRequest) String() string { return proto.CompactTextString(m) }
func (*RegistRequest) ProtoMessage()    {}
func (*RegistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *RegistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistRequest.Unmarshal(m, b)
}
func (m *RegistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistRequest.Marshal(b, m, deterministic)
}
func (m *RegistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistRequest.Merge(m, src)
}
func (m *RegistRequest) XXX_Size() int {
	return xxx_messageInfo_RegistRequest.Size(m)
}
func (m *RegistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegistRequest proto.InternalMessageInfo

func (m *RegistRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RegistRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegistRequest) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *RegistRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegistRequest) GetPasswordAgain() string {
	if m != nil {
		return m.PasswordAgain
	}
	return ""
}

type RegistResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistResponse) Reset()         { *m = RegistResponse{} }
func (m *RegistResponse) String() string { return proto.CompactTextString(m) }
func (*RegistResponse) ProtoMessage()    {}
func (*RegistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *RegistResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistResponse.Unmarshal(m, b)
}
func (m *RegistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistResponse.Marshal(b, m, deterministic)
}
func (m *RegistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistResponse.Merge(m, src)
}
func (m *RegistResponse) XXX_Size() int {
	return xxx_messageInfo_RegistResponse.Size(m)
}
func (m *RegistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegistResponse proto.InternalMessageInfo

func (m *RegistResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegistResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type LoginRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Account struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Nickname             string   `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar               string   `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Status               int32    `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Roles                string   `protobuf:"bytes,7,opt,name=roles,proto3" json:"roles,omitempty"`
	Type                 int32    `protobuf:"varint,8,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Account) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Account) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Account) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Account) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Account) GetRoles() string {
	if m != nil {
		return m.Roles
	}
	return ""
}

func (m *Account) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type LoginResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *Account `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LoginResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LoginResponse) GetData() *Account {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*RegistRequest)(nil), "go.micro.srv.user.RegistRequest")
	proto.RegisterType((*RegistResponse)(nil), "go.micro.srv.user.RegistResponse")
	proto.RegisterType((*LoginRequest)(nil), "go.micro.srv.user.LoginRequest")
	proto.RegisterType((*Account)(nil), "go.micro.srv.user.Account")
	proto.RegisterType((*LoginResponse)(nil), "go.micro.srv.user.LoginResponse")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcf, 0x4a, 0xfb, 0x40,
	0x10, 0xfe, 0x25, 0x4d, 0xff, 0x4d, 0x7f, 0x2d, 0xb8, 0xa8, 0x2c, 0xf1, 0x60, 0x0d, 0x08, 0x9e,
	0x22, 0xd4, 0xbb, 0xd0, 0x93, 0x17, 0x15, 0x0c, 0x7a, 0x2e, 0x6b, 0xb2, 0x84, 0xc5, 0x26, 0x1b,
	0x77, 0x37, 0x15, 0x5f, 0xc7, 0x17, 0xf0, 0x19, 0x7c, 0x33, 0xd9, 0xd9, 0x84, 0xb6, 0x58, 0x45,
	0xbc, 0x84, 0xfd, 0x66, 0xe6, 0x9b, 0xf9, 0xe6, 0xdb, 0x0d, 0x1c, 0x54, 0x4a, 0x1a, 0x79, 0x5e,
	0x6b, 0xae, 0xf0, 0x13, 0x23, 0x26, 0x7b, 0xb9, 0x8c, 0x0b, 0x91, 0x2a, 0x19, 0x6b, 0xb5, 0x8a,
	0x6d, 0x22, 0x7a, 0xf3, 0x60, 0x9c, 0xf0, 0x5c, 0x68, 0x93, 0xf0, 0xe7, 0x9a, 0x6b, 0x43, 0x8e,
	0x60, 0x68, 0x33, 0x8b, 0x92, 0x15, 0x9c, 0x7a, 0x53, 0xef, 0x6c, 0x98, 0x0c, 0x6c, 0xe0, 0x96,
	0x15, 0x9c, 0xec, 0x43, 0x97, 0x17, 0x4c, 0x2c, 0xa9, 0x8f, 0x09, 0x07, 0x2c, 0xa5, 0x14, 0xe9,
	0x93, 0xa3, 0x74, 0x1c, 0xc5, 0x06, 0x90, 0x12, 0xc2, 0xa0, 0x62, 0x5a, 0xbf, 0x48, 0x95, 0xd1,
	0xc0, 0xe5, 0x5a, 0x4c, 0x4e, 0x61, 0xd2, 0x9e, 0x17, 0x2c, 0x67, 0xa2, 0xa4, 0x5d, 0xac, 0x18,
	0xb7, 0xd1, 0xb9, 0x0d, 0x46, 0x97, 0x30, 0x69, 0x35, 0xea, 0x4a, 0x96, 0x9a, 0x13, 0x02, 0x41,
	0x2a, 0x33, 0xa7, 0xaf, 0x9b, 0xe0, 0x99, 0x50, 0xe8, 0x17, 0x5c, 0x6b, 0x96, 0xf3, 0x46, 0x5d,
	0x0b, 0xa3, 0x2b, 0xf8, 0x7f, 0x2d, 0x73, 0x51, 0xfe, 0x6a, 0xc5, 0x4d, 0xbd, 0xfe, 0xb6, 0xde,
	0xe8, 0xc3, 0x83, 0xfe, 0x3c, 0x4d, 0x65, 0x5d, 0x1a, 0x32, 0x01, 0x5f, 0x64, 0x8d, 0x00, 0x5f,
	0x64, 0x96, 0x67, 0x7b, 0x60, 0x4f, 0x7f, 0xdd, 0xb3, 0xdc, 0xb2, 0xad, 0xb3, 0x69, 0x5b, 0x08,
	0xe8, 0x12, 0x32, 0x82, 0xb5, 0x6b, 0xc8, 0x38, 0x84, 0x1e, 0x5b, 0x31, 0xc3, 0x54, 0xe3, 0x48,
	0x83, 0x6c, 0x5c, 0x1b, 0x66, 0x6a, 0x4d, 0x7b, 0x38, 0xb9, 0x41, 0x76, 0x82, 0x92, 0x4b, 0xae,
	0x69, 0xdf, 0x4d, 0x40, 0x60, 0x6d, 0x32, 0xaf, 0x15, 0xa7, 0x03, 0x67, 0x93, 0x3d, 0x47, 0x05,
	0x8c, 0x1b, 0x33, 0xfe, 0xe2, 0x25, 0x89, 0x21, 0xc8, 0x98, 0x61, 0xb8, 0xc9, 0x68, 0x16, 0xc6,
	0x5f, 0x9e, 0x54, 0xdc, 0x18, 0x94, 0x60, 0xdd, 0xec, 0xdd, 0x83, 0xe0, 0x41, 0x73, 0x45, 0xee,
	0x61, 0x74, 0x63, 0x0b, 0xdd, 0x4d, 0x92, 0xe9, 0x0e, 0xe6, 0xd6, 0x43, 0x0c, 0x4f, 0x7e, 0xa8,
	0x70, 0xd2, 0xa3, 0x7f, 0xe4, 0x0e, 0x00, 0xbb, 0xe2, 0x4a, 0xe4, 0x78, 0x07, 0x65, 0xf3, 0xe6,
	0xc3, 0xe9, 0xf7, 0x05, 0x6d, 0xcb, 0xc7, 0x1e, 0xfe, 0x2c, 0x17, 0x9f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x79, 0x6f, 0x1c, 0x7d, 0x45, 0x03, 0x00, 0x00,
}