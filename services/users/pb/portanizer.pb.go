// Code generated by protoc-gen-go. DO NOT EDIT.
// source: portanizer.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type CreateAccountRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Pwd                  string   `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountRequest) Reset()         { *m = CreateAccountRequest{} }
func (m *CreateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRequest) ProtoMessage()    {}
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_966fb747b99153fb, []int{0}
}

func (m *CreateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRequest.Unmarshal(m, b)
}
func (m *CreateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRequest.Marshal(b, m, deterministic)
}
func (m *CreateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRequest.Merge(m, src)
}
func (m *CreateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRequest.Size(m)
}
func (m *CreateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRequest proto.InternalMessageInfo

func (m *CreateAccountRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateAccountRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type CreateAccountReply struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountReply) Reset()         { *m = CreateAccountReply{} }
func (m *CreateAccountReply) String() string { return proto.CompactTextString(m) }
func (*CreateAccountReply) ProtoMessage()    {}
func (*CreateAccountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_966fb747b99153fb, []int{1}
}

func (m *CreateAccountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountReply.Unmarshal(m, b)
}
func (m *CreateAccountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountReply.Marshal(b, m, deterministic)
}
func (m *CreateAccountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountReply.Merge(m, src)
}
func (m *CreateAccountReply) XXX_Size() int {
	return xxx_messageInfo_CreateAccountReply.Size(m)
}
func (m *CreateAccountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountReply proto.InternalMessageInfo

func (m *CreateAccountReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateAccountReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_966fb747b99153fb, []int{2}
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateAccountRequest)(nil), "pb.CreateAccountRequest")
	proto.RegisterType((*CreateAccountReply)(nil), "pb.CreateAccountReply")
	proto.RegisterType((*User)(nil), "pb.User")
}

func init() { proto.RegisterFile("portanizer.proto", fileDescriptor_966fb747b99153fb) }

var fileDescriptor_966fb747b99153fb = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xc8, 0x2f, 0x2a,
	0x49, 0xcc, 0xcb, 0xac, 0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x52, 0xb2, 0xe3, 0x12, 0x71, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x75, 0x4c, 0x4e, 0xce, 0x2f, 0xcd,
	0x2b, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x4d, 0xcd, 0x4d, 0xcc,
	0xcc, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0x04, 0xb8, 0x98, 0x0b, 0xca,
	0x53, 0x24, 0x98, 0xc0, 0x62, 0x20, 0xa6, 0x92, 0x0b, 0x97, 0x10, 0x9a, 0xfe, 0x82, 0x9c, 0x4a,
	0x21, 0x19, 0x2e, 0x96, 0xd2, 0xe2, 0xd4, 0x22, 0xb0, 0x66, 0x6e, 0x23, 0x0e, 0xbd, 0x82, 0x24,
	0xbd, 0xd0, 0xe2, 0xd4, 0xa2, 0x20, 0xb0, 0x28, 0xc8, 0x94, 0xd4, 0xa2, 0x22, 0x98, 0x29, 0xa9,
	0x45, 0x45, 0x4a, 0x3a, 0x5c, 0x2c, 0x20, 0x79, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0xa8, 0x95,
	0x4c, 0x99, 0x29, 0x08, 0x57, 0x30, 0x21, 0xb9, 0xc2, 0xc8, 0x87, 0x8b, 0x15, 0xa4, 0xba, 0x58,
	0xc8, 0x99, 0x8b, 0x17, 0xc5, 0x72, 0x21, 0x09, 0x90, 0x4d, 0xd8, 0xfc, 0x23, 0x25, 0x86, 0x45,
	0xa6, 0x20, 0xa7, 0x52, 0x89, 0x21, 0x89, 0x0d, 0x1c, 0x18, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x9e, 0x01, 0xfd, 0xfa, 0x20, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountReply, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountReply, error) {
	out := new(CreateAccountReply)
	err := c.cc.Invoke(ctx, "/pb.Users/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountReply, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Users/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Users_CreateAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portanizer.proto",
}
