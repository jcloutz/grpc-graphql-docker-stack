// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

/*
Package todo is a generated protocol buffer package.

It is generated from these files:
	todo.proto

It has these top-level messages:
	TodoCreateRequest
	TodoFilter
	TodoUpdateRequest
	TodoResponse
*/
package todo

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

type TodoCreateRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Complete bool   `protobuf:"varint,2,opt,name=complete" json:"complete,omitempty"`
}

func (m *TodoCreateRequest) Reset()                    { *m = TodoCreateRequest{} }
func (m *TodoCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*TodoCreateRequest) ProtoMessage()               {}
func (*TodoCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TodoCreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TodoCreateRequest) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

type TodoFilter struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *TodoFilter) Reset()                    { *m = TodoFilter{} }
func (m *TodoFilter) String() string            { return proto.CompactTextString(m) }
func (*TodoFilter) ProtoMessage()               {}
func (*TodoFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TodoFilter) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type TodoUpdateRequest struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Complete bool   `protobuf:"varint,3,opt,name=complete" json:"complete,omitempty"`
}

func (m *TodoUpdateRequest) Reset()                    { *m = TodoUpdateRequest{} }
func (m *TodoUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*TodoUpdateRequest) ProtoMessage()               {}
func (*TodoUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TodoUpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TodoUpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TodoUpdateRequest) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

type TodoResponse struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Complete  bool   `protobuf:"varint,3,opt,name=complete" json:"complete,omitempty"`
	CreatedAt int64  `protobuf:"varint,4,opt,name=createdAt" json:"createdAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,5,opt,name=updatedAt" json:"updatedAt,omitempty"`
}

func (m *TodoResponse) Reset()                    { *m = TodoResponse{} }
func (m *TodoResponse) String() string            { return proto.CompactTextString(m) }
func (*TodoResponse) ProtoMessage()               {}
func (*TodoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TodoResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TodoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TodoResponse) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

func (m *TodoResponse) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *TodoResponse) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*TodoCreateRequest)(nil), "todo.TodoCreateRequest")
	proto.RegisterType((*TodoFilter)(nil), "todo.TodoFilter")
	proto.RegisterType((*TodoUpdateRequest)(nil), "todo.TodoUpdateRequest")
	proto.RegisterType((*TodoResponse)(nil), "todo.TodoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Todo service

type TodoClient interface {
	GetTodos(ctx context.Context, in *TodoFilter, opts ...grpc.CallOption) (Todo_GetTodosClient, error)
	CreateTodo(ctx context.Context, in *TodoCreateRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	UpdateTodo(ctx context.Context, in *TodoUpdateRequest, opts ...grpc.CallOption) (*TodoResponse, error)
}

type todoClient struct {
	cc *grpc.ClientConn
}

func NewTodoClient(cc *grpc.ClientConn) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) GetTodos(ctx context.Context, in *TodoFilter, opts ...grpc.CallOption) (Todo_GetTodosClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Todo_serviceDesc.Streams[0], c.cc, "/todo.Todo/GetTodos", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoGetTodosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Todo_GetTodosClient interface {
	Recv() (*TodoResponse, error)
	grpc.ClientStream
}

type todoGetTodosClient struct {
	grpc.ClientStream
}

func (x *todoGetTodosClient) Recv() (*TodoResponse, error) {
	m := new(TodoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *todoClient) CreateTodo(ctx context.Context, in *TodoCreateRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := grpc.Invoke(ctx, "/todo.Todo/CreateTodo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) UpdateTodo(ctx context.Context, in *TodoUpdateRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := grpc.Invoke(ctx, "/todo.Todo/UpdateTodo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Todo service

type TodoServer interface {
	GetTodos(*TodoFilter, Todo_GetTodosServer) error
	CreateTodo(context.Context, *TodoCreateRequest) (*TodoResponse, error)
	UpdateTodo(context.Context, *TodoUpdateRequest) (*TodoResponse, error)
}

func RegisterTodoServer(s *grpc.Server, srv TodoServer) {
	s.RegisterService(&_Todo_serviceDesc, srv)
}

func _Todo_GetTodos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TodoFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoServer).GetTodos(m, &todoGetTodosServer{stream})
}

type Todo_GetTodosServer interface {
	Send(*TodoResponse) error
	grpc.ServerStream
}

type todoGetTodosServer struct {
	grpc.ServerStream
}

func (x *todoGetTodosServer) Send(m *TodoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Todo_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Todo/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).CreateTodo(ctx, req.(*TodoCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Todo/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).UpdateTodo(ctx, req.(*TodoUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Todo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _Todo_CreateTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _Todo_UpdateTodo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTodos",
			Handler:       _Todo_GetTodos_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "todo.proto",
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xb1, 0x4e, 0xc3, 0x40,
	0x0c, 0x86, 0xb9, 0x34, 0xa0, 0xd4, 0x42, 0x08, 0xbc, 0x10, 0x55, 0x1d, 0xa2, 0x4c, 0x99, 0x2a,
	0x04, 0x6c, 0x4c, 0xa8, 0x12, 0xec, 0x07, 0x3c, 0x40, 0xe9, 0x79, 0x88, 0xd4, 0xd6, 0x21, 0x71,
	0x5f, 0x82, 0x07, 0xe2, 0xf9, 0x90, 0xef, 0x20, 0x47, 0xa3, 0x30, 0x75, 0x73, 0x6c, 0xff, 0xff,
	0x6f, 0x7d, 0x39, 0x00, 0x61, 0xc7, 0x8b, 0xa6, 0x65, 0x61, 0x4c, 0xb5, 0x2e, 0x97, 0x70, 0xf5,
	0xca, 0x8e, 0x97, 0x2d, 0xad, 0x84, 0x2c, 0x7d, 0xec, 0xa9, 0x13, 0x44, 0x48, 0x77, 0xab, 0x2d,
	0xe5, 0xa6, 0x30, 0xd5, 0xd4, 0xfa, 0x1a, 0x67, 0x90, 0xad, 0x79, 0xdb, 0x6c, 0x48, 0x28, 0x4f,
	0x0a, 0x53, 0x65, 0xb6, 0xff, 0x2e, 0xe7, 0x00, 0x6a, 0xf2, 0x54, 0x6f, 0x84, 0x5a, 0xbc, 0x80,
	0xa4, 0x76, 0x3f, 0xda, 0xa4, 0x76, 0xe5, 0x4b, 0x88, 0x78, 0x6b, 0xdc, 0x9f, 0x88, 0xc1, 0x52,
	0x1f, 0x99, 0xfc, 0x13, 0x39, 0x19, 0x44, 0x7e, 0x1a, 0x38, 0x57, 0x57, 0x4b, 0x5d, 0xc3, 0xbb,
	0x8e, 0x8e, 0x35, 0xc4, 0x39, 0x4c, 0xd7, 0x1e, 0x82, 0x7b, 0x94, 0x3c, 0x2d, 0x4c, 0x35, 0xb1,
	0xb1, 0xa1, 0xd3, 0xbd, 0xbf, 0x5f, 0xa7, 0xa7, 0x61, 0xda, 0x37, 0x6e, 0xbf, 0x0c, 0xa4, 0x7a,
	0x0c, 0xde, 0x43, 0xf6, 0x4c, 0xa2, 0x65, 0x87, 0x97, 0x0b, 0x0f, 0x3b, 0x82, 0x99, 0x61, 0xec,
	0xfc, 0x9e, 0x5d, 0x9e, 0xdc, 0x18, 0x7c, 0x00, 0x08, 0xfc, 0xbd, 0xc7, 0x75, 0xdc, 0x3a, 0xf8,
	0x2b, 0xe3, 0x72, 0x15, 0x07, 0xb2, 0x43, 0xf1, 0x01, 0xef, 0x71, 0xf1, 0xfb, 0x99, 0x7f, 0x0a,
	0x77, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xd0, 0x04, 0x44, 0x18, 0x02, 0x00, 0x00,
}
