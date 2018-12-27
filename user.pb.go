// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user // import "github.com/taeho-io/user"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserType int32

const (
	UserType_EMAIL UserType = 0
)

var UserType_name = map[int32]string{
	0: "EMAIL",
}
var UserType_value = map[string]int32{
	"EMAIL": 0,
}

func (x UserType) String() string {
	return proto.EnumName(UserType_name, int32(x))
}
func (UserType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_2559444f57daed46, []int{0}
}

type RegisterRequest struct {
	UserType             UserType `protobuf:"varint,1,opt,name=user_type,json=userType,proto3,enum=user.UserType" json:"user_type,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_2559444f57daed46, []int{0}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(dst, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUserType() UserType {
	if m != nil {
		return m.UserType
	}
	return UserType_EMAIL
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RegisterResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_2559444f57daed46, []int{1}
}
func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(dst, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *RegisterResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type LogInRequest struct {
	UserType             UserType `protobuf:"varint,1,opt,name=user_type,json=userType,proto3,enum=user.UserType" json:"user_type,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInRequest) Reset()         { *m = LogInRequest{} }
func (m *LogInRequest) String() string { return proto.CompactTextString(m) }
func (*LogInRequest) ProtoMessage()    {}
func (*LogInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_2559444f57daed46, []int{2}
}
func (m *LogInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInRequest.Unmarshal(m, b)
}
func (m *LogInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInRequest.Marshal(b, m, deterministic)
}
func (dst *LogInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInRequest.Merge(dst, src)
}
func (m *LogInRequest) XXX_Size() int {
	return xxx_messageInfo_LogInRequest.Size(m)
}
func (m *LogInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogInRequest proto.InternalMessageInfo

func (m *LogInRequest) GetUserType() UserType {
	if m != nil {
		return m.UserType
	}
	return UserType_EMAIL
}

func (m *LogInRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *LogInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LogInResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresAt            int64    `protobuf:"varint,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInResponse) Reset()         { *m = LogInResponse{} }
func (m *LogInResponse) String() string { return proto.CompactTextString(m) }
func (*LogInResponse) ProtoMessage()    {}
func (*LogInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_2559444f57daed46, []int{3}
}
func (m *LogInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInResponse.Unmarshal(m, b)
}
func (m *LogInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInResponse.Marshal(b, m, deterministic)
}
func (dst *LogInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInResponse.Merge(dst, src)
}
func (m *LogInResponse) XXX_Size() int {
	return xxx_messageInfo_LogInResponse.Size(m)
}
func (m *LogInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogInResponse proto.InternalMessageInfo

func (m *LogInResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *LogInResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *LogInResponse) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_2559444f57daed46, []int{4}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type GetResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_2559444f57daed46, []int{5}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "user.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "user.RegisterResponse")
	proto.RegisterType((*LogInRequest)(nil), "user.LogInRequest")
	proto.RegisterType((*LogInResponse)(nil), "user.LogInResponse")
	proto.RegisterType((*GetRequest)(nil), "user.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "user.GetResponse")
	proto.RegisterEnum("user.UserType", UserType_name, UserType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/user.User/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error) {
	out := new(LogInResponse)
	err := c.cc.Invoke(ctx, "/user.User/LogIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/user.User/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	LogIn(context.Context, *LogInRequest) (*LogInResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/LogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LogIn(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "LogIn",
			Handler:    _User_LogIn_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _User_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_2559444f57daed46) }

var fileDescriptor_user_2559444f57daed46 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0x71, 0x6d, 0xa8, 0x3d, 0x18, 0x4a, 0xa7, 0xa5, 0xb5, 0x2c, 0x55, 0xa2, 0xee, 0x05,
	0xb5, 0x0d, 0x48, 0xe4, 0x98, 0x13, 0x91, 0x22, 0x84, 0x44, 0x2e, 0x16, 0xb9, 0x70, 0x41, 0x06,
	0x26, 0x60, 0x25, 0x78, 0x9d, 0xdd, 0x75, 0x12, 0xee, 0xf9, 0x2a, 0xf9, 0x9e, 0x91, 0x77, 0xcd,
	0x9f, 0x70, 0xc8, 0x21, 0xca, 0x6d, 0xe6, 0x37, 0x6f, 0x77, 0xdf, 0xdb, 0x5d, 0x80, 0x4c, 0x10,
	0xef, 0xa4, 0x9c, 0x49, 0x86, 0x56, 0x5e, 0x07, 0x4f, 0x06, 0x7c, 0x09, 0x69, 0x19, 0x0b, 0x49,
	0x3c, 0xa4, 0xbb, 0x8c, 0x84, 0xc4, 0x7f, 0xe0, 0xe4, 0xb3, 0xa9, 0xdc, 0xa4, 0xe4, 0x19, 0x2d,
	0xa3, 0x5d, 0xef, 0xd5, 0x3b, 0x6a, 0xe5, 0x95, 0x20, 0x3e, 0xde, 0xa4, 0x14, 0xda, 0x59, 0x51,
	0xe1, 0x77, 0x28, 0xd3, 0x3a, 0x8a, 0x6f, 0xbd, 0x4f, 0x2d, 0xa3, 0xed, 0x84, 0xba, 0x41, 0x1f,
	0xec, 0x34, 0x12, 0xe2, 0x81, 0xf1, 0x85, 0x67, 0xaa, 0xc1, 0xae, 0x47, 0x04, 0x2b, 0x89, 0xd6,
	0xe4, 0x59, 0x8a, 0xab, 0x3a, 0x98, 0x40, 0x63, 0xef, 0x42, 0xa4, 0x2c, 0x11, 0x84, 0xbf, 0xc1,
	0x8d, 0xe6, 0x73, 0x12, 0x62, 0x2a, 0xd9, 0x0d, 0x25, 0xca, 0x89, 0x13, 0x56, 0x35, 0x1b, 0xe7,
	0x08, 0xff, 0x40, 0x8d, 0xd3, 0x35, 0x27, 0xb1, 0x2a, 0x34, 0xda, 0x84, 0x5b, 0x40, 0x25, 0x0a,
	0x52, 0x70, 0x47, 0x6c, 0x39, 0x4c, 0xde, 0x15, 0xef, 0x27, 0x7c, 0x56, 0xe2, 0x78, 0xa1, 0xf6,
	0x36, 0xc3, 0x4a, 0xde, 0x0e, 0x17, 0x6f, 0x25, 0x0c, 0xee, 0xa1, 0x56, 0x9c, 0xf8, 0xb1, 0x51,
	0xf0, 0x17, 0x00, 0x3d, 0xa6, 0x31, 0x27, 0x31, 0x8d, 0xa4, 0x3a, 0xd6, 0x0c, 0x9d, 0x82, 0xf4,
	0x65, 0xe0, 0x02, 0x0c, 0x48, 0x16, 0x39, 0x83, 0x1a, 0x54, 0x55, 0xa7, 0x3d, 0xfc, 0x6d, 0x82,
	0xbd, 0xcd, 0x87, 0x0e, 0x94, 0x2f, 0x2e, 0xfb, 0xc3, 0x51, 0xa3, 0xd4, 0x7b, 0x36, 0xc0, 0xca,
	0x39, 0x9e, 0x81, 0xbd, 0x7d, 0x02, 0x6c, 0xea, 0xfb, 0x38, 0xfa, 0x18, 0xfe, 0x8f, 0x63, 0xac,
	0xb7, 0x0e, 0x4a, 0xd8, 0x83, 0xb2, 0x4a, 0x8c, 0xa8, 0x25, 0x87, 0x17, 0xee, 0x7f, 0x7b, 0xc5,
	0x76, 0x6b, 0xfe, 0x83, 0x39, 0x20, 0x89, 0x0d, 0x3d, 0xdd, 0x1b, 0xf7, 0xbf, 0x1e, 0x90, 0xad,
	0xfa, 0xdc, 0x9f, 0x78, 0xcb, 0x58, 0xae, 0xb2, 0x59, 0x67, 0xce, 0xd6, 0x5d, 0x19, 0xd1, 0x8a,
	0x9d, 0xc4, 0xac, 0x9b, 0x2b, 0x67, 0x15, 0xf5, 0xa3, 0x4f, 0x5f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xe6, 0x97, 0xd1, 0x7a, 0xdf, 0x02, 0x00, 0x00,
}
