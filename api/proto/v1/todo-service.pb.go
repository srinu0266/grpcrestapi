// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/v1/todo-service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ToDo struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Remainder            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=remainder,proto3" json:"remainder,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ToDo) Reset()         { *m = ToDo{} }
func (m *ToDo) String() string { return proto.CompactTextString(m) }
func (*ToDo) ProtoMessage()    {}
func (*ToDo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a190534404bbf0f9, []int{0}
}

func (m *ToDo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToDo.Unmarshal(m, b)
}
func (m *ToDo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToDo.Marshal(b, m, deterministic)
}
func (m *ToDo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToDo.Merge(m, src)
}
func (m *ToDo) XXX_Size() int {
	return xxx_messageInfo_ToDo.Size(m)
}
func (m *ToDo) XXX_DiscardUnknown() {
	xxx_messageInfo_ToDo.DiscardUnknown(m)
}

var xxx_messageInfo_ToDo proto.InternalMessageInfo

func (m *ToDo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ToDo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ToDo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ToDo) GetRemainder() *timestamp.Timestamp {
	if m != nil {
		return m.Remainder
	}
	return nil
}

type CreateRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	ToDo                 *ToDo    `protobuf:"bytes,2,opt,name=toDo,proto3" json:"toDo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a190534404bbf0f9, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetToDo() *ToDo {
	if m != nil {
		return m.ToDo
	}
	return nil
}

type CreateResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a190534404bbf0f9, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*ToDo)(nil), "v1.ToDo")
	proto.RegisterType((*CreateRequest)(nil), "v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "v1.CreateResponse")
}

func init() {
	proto.RegisterFile("api/proto/v1/todo-service.proto", fileDescriptor_a190534404bbf0f9)
}

var fileDescriptor_a190534404bbf0f9 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4e, 0x84, 0x30,
	0x14, 0x85, 0xc3, 0x8f, 0x13, 0xb9, 0xc4, 0x89, 0x36, 0x2e, 0x08, 0x31, 0x19, 0xc2, 0x8a, 0x8d,
	0x25, 0xe0, 0xc6, 0x95, 0x2e, 0x9c, 0x27, 0xa8, 0xf3, 0x02, 0xcc, 0x70, 0x9d, 0x34, 0x19, 0xb8,
	0xb5, 0xed, 0xf0, 0x0c, 0x3e, 0xb6, 0xa1, 0x0d, 0x71, 0x4c, 0x66, 0x07, 0xa7, 0x27, 0x5f, 0xbf,
	0x53, 0xd8, 0x74, 0x4a, 0xd6, 0x4a, 0x93, 0xa5, 0x7a, 0x6a, 0x6a, 0x4b, 0x3d, 0x3d, 0x1b, 0xd4,
	0x93, 0x3c, 0x20, 0x77, 0x29, 0x0b, 0xa7, 0x26, 0xdf, 0x1c, 0x89, 0x8e, 0x27, 0xf4, 0xbd, 0xfd,
	0xf9, 0xab, 0xb6, 0x72, 0x40, 0x63, 0xbb, 0x41, 0xf9, 0x52, 0xf9, 0x13, 0x40, 0xbc, 0xa3, 0x2d,
	0xb1, 0x35, 0x84, 0xb2, 0xcf, 0x82, 0x22, 0xa8, 0x22, 0x11, 0xca, 0x9e, 0x3d, 0xc2, 0x8d, 0x95,
	0xf6, 0x84, 0x59, 0x58, 0x04, 0x55, 0x22, 0xfc, 0x0f, 0x2b, 0x20, 0xed, 0xd1, 0x1c, 0xb4, 0x54,
	0x56, 0xd2, 0x98, 0x45, 0xee, 0xec, 0x32, 0x62, 0xaf, 0x90, 0x68, 0x1c, 0x3a, 0x39, 0xf6, 0xa8,
	0xb3, 0xb8, 0x08, 0xaa, 0xb4, 0xcd, 0xb9, 0xb7, 0xe0, 0x8b, 0x05, 0xdf, 0x2d, 0x16, 0xe2, 0xaf,
	0x5c, 0xbe, 0xc3, 0xdd, 0x87, 0xc6, 0xce, 0xa2, 0xc0, 0xef, 0x33, 0x1a, 0xcb, 0xee, 0x21, 0xea,
	0x94, 0x74, 0x4e, 0x89, 0x98, 0x3f, 0xd9, 0x13, 0xc4, 0x96, 0xb6, 0xe4, 0x9c, 0xd2, 0xf6, 0x96,
	0x4f, 0x0d, 0x9f, 0xe5, 0x85, 0x4b, 0xcb, 0x16, 0xd6, 0x0b, 0xc0, 0x28, 0x1a, 0x0d, 0x5e, 0x21,
	0xf8, 0x99, 0xe1, 0x32, 0xb3, 0x7d, 0x83, 0x74, 0x26, 0x7c, 0xfa, 0x97, 0x63, 0x35, 0xac, 0x3c,
	0x82, 0x3d, 0xcc, 0xf0, 0x7f, 0x3e, 0x39, 0xbb, 0x8c, 0xfc, 0x0d, 0xfb, 0x95, 0xdb, 0xf4, 0xf2,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x30, 0xa7, 0xa3, 0xc9, 0x8e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ToDoServiceClient is the client API for ToDoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ToDoServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type toDoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewToDoServiceClient(cc grpc.ClientConnInterface) ToDoServiceClient {
	return &toDoServiceClient{cc}
}

func (c *toDoServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/v1.ToDoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToDoServiceServer is the server API for ToDoService service.
type ToDoServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
}

// UnimplementedToDoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedToDoServiceServer struct {
}

func (*UnimplementedToDoServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterToDoServiceServer(s *grpc.Server, srv ToDoServiceServer) {
	s.RegisterService(&_ToDoService_serviceDesc, srv)
}

func _ToDoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ToDoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ToDoService",
	HandlerType: (*ToDoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ToDoService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/todo-service.proto",
}
