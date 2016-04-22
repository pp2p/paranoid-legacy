// Code generated by protoc-gen-go.

// source: api.proto

// DO NOT EDIT!

/*

Package fileserver is a generated protocol buffer package.

It is generated from these files:

  api.proto

It has these top-level messages:

  ServeResponse

  ServeRequest

  UnServeRequest

  ListServeRequest

  ListServeResponse

  ServedFile

*/

package fileserver

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

const _ = proto.ProtoPackageIsVersion1

type ServeResponse struct {
	ServeResponse string `protobuf:"bytes,1,opt,name=ServeResponse,json=serveResponse" json:"ServeResponse,omitempty"`

	ServerPort string `protobuf:"bytes,2,opt,name=ServerPort,json=serverPort" json:"ServerPort,omitempty"`
}

func (m *ServeResponse) Reset() { *m = ServeResponse{} }

func (m *ServeResponse) String() string { return proto.CompactTextString(m) }

func (*ServeResponse) ProtoMessage() {}

func (*ServeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServeRequest struct {
	Uuid string `protobuf:"bytes,1,opt,name=Uuid,json=uuid" json:"Uuid,omitempty"`

	Pool string `protobuf:"bytes,2,opt,name=Pool,json=pool" json:"Pool,omitempty"`

	FilePath string `protobuf:"bytes,3,opt,name=FilePath,json=filePath" json:"FilePath,omitempty"`

	FileData []byte `protobuf:"bytes,4,opt,name=FileData,json=fileData,proto3" json:"FileData,omitempty"`

	Timeout int32 `protobuf:"varint,5,opt,name=Timeout,json=timeout" json:"Timeout,omitempty"`

	Limit int32 `protobuf:"varint,6,opt,name=Limit,json=limit" json:"Limit,omitempty"`
}

func (m *ServeRequest) Reset() { *m = ServeRequest{} }

func (m *ServeRequest) String() string { return proto.CompactTextString(m) }

func (*ServeRequest) ProtoMessage() {}

func (*ServeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type UnServeRequest struct {
	Uuid string `protobuf:"bytes,1,opt,name=Uuid,json=uuid" json:"Uuid,omitempty"`

	Pool string `protobuf:"bytes,2,opt,name=Pool,json=pool" json:"Pool,omitempty"`

	FilePath string `protobuf:"bytes,3,opt,name=FilePath,json=filePath" json:"FilePath,omitempty"`
}

func (m *UnServeRequest) Reset() { *m = UnServeRequest{} }

func (m *UnServeRequest) String() string { return proto.CompactTextString(m) }

func (*UnServeRequest) ProtoMessage() {}

func (*UnServeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ListServeRequest struct {
	Uuid string `protobuf:"bytes,1,opt,name=Uuid,json=uuid" json:"Uuid,omitempty"`

	Pool string `protobuf:"bytes,2,opt,name=Pool,json=pool" json:"Pool,omitempty"`
}

func (m *ListServeRequest) Reset() { *m = ListServeRequest{} }

func (m *ListServeRequest) String() string { return proto.CompactTextString(m) }

func (*ListServeRequest) ProtoMessage() {}

func (*ListServeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ListServeResponse struct {
	ServedFiles []*ServedFile `protobuf:"bytes,1,rep,name=ServedFiles,json=servedFiles" json:"ServedFiles,omitempty"`
}

func (m *ListServeResponse) Reset() { *m = ListServeResponse{} }

func (m *ListServeResponse) String() string { return proto.CompactTextString(m) }

func (*ListServeResponse) ProtoMessage() {}

func (*ListServeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListServeResponse) GetServedFiles() []*ServedFile {

	if m != nil {

		return m.ServedFiles

	}

	return nil

}

type ServedFile struct {
	FilePath string `protobuf:"bytes,1,opt,name=FilePath,json=filePath" json:"FilePath,omitempty"`

	FileHash string `protobuf:"bytes,2,opt,name=FileHash,json=fileHash" json:"FileHash,omitempty"`

	AccessLimit int32 `protobuf:"varint,3,opt,name=AccessLimit,json=accessLimit" json:"AccessLimit,omitempty"`

	ExpirationTime string `protobuf:"bytes,4,opt,name=ExpirationTime,json=expirationTime" json:"ExpirationTime,omitempty"`
}

func (m *ServedFile) Reset() { *m = ServedFile{} }

func (m *ServedFile) String() string { return proto.CompactTextString(m) }

func (*ServedFile) ProtoMessage() {}

func (*ServedFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {

	proto.RegisterType((*ServeResponse)(nil), "fileserver.ServeResponse")

	proto.RegisterType((*ServeRequest)(nil), "fileserver.ServeRequest")

	proto.RegisterType((*UnServeRequest)(nil), "fileserver.UnServeRequest")

	proto.RegisterType((*ListServeRequest)(nil), "fileserver.ListServeRequest")

	proto.RegisterType((*ListServeResponse)(nil), "fileserver.ListServeResponse")

	proto.RegisterType((*ServedFile)(nil), "fileserver.ServedFile")

}

// Reference imports to suppress errors if they are not otherwise used.

var _ context.Context

var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file

// is compatible with the grpc package it is being compiled against.

const _ = grpc.SupportPackageIsVersion2

// Client API for Fileserver service

type FileserverClient interface {

	// File Server Calls

	ServeFile(ctx context.Context, in *ServeRequest, opts ...grpc.CallOption) (*ServeResponse, error)

	UnServeFile(ctx context.Context, in *UnServeRequest, opts ...grpc.CallOption) (*ServeResponse, error)

	ListServer(ctx context.Context, in *ListServeRequest, opts ...grpc.CallOption) (*ListServeResponse, error)
}

type fileserverClient struct {
	cc *grpc.ClientConn
}

func NewFileserverClient(cc *grpc.ClientConn) FileserverClient {

	return &fileserverClient{cc}

}

func (c *fileserverClient) ServeFile(ctx context.Context, in *ServeRequest, opts ...grpc.CallOption) (*ServeResponse, error) {

	out := new(ServeResponse)

	err := grpc.Invoke(ctx, "/fileserver.Fileserver/ServeFile", in, out, c.cc, opts...)

	if err != nil {

		return nil, err

	}

	return out, nil

}

func (c *fileserverClient) UnServeFile(ctx context.Context, in *UnServeRequest, opts ...grpc.CallOption) (*ServeResponse, error) {

	out := new(ServeResponse)

	err := grpc.Invoke(ctx, "/fileserver.Fileserver/UnServeFile", in, out, c.cc, opts...)

	if err != nil {

		return nil, err

	}

	return out, nil

}

func (c *fileserverClient) ListServer(ctx context.Context, in *ListServeRequest, opts ...grpc.CallOption) (*ListServeResponse, error) {

	out := new(ListServeResponse)

	err := grpc.Invoke(ctx, "/fileserver.Fileserver/ListServer", in, out, c.cc, opts...)

	if err != nil {

		return nil, err

	}

	return out, nil

}

// Server API for Fileserver service

type FileserverServer interface {

	// File Server Calls

	ServeFile(context.Context, *ServeRequest) (*ServeResponse, error)

	UnServeFile(context.Context, *UnServeRequest) (*ServeResponse, error)

	ListServer(context.Context, *ListServeRequest) (*ListServeResponse, error)
}

func RegisterFileserverServer(s *grpc.Server, srv FileserverServer) {

	s.RegisterService(&_Fileserver_serviceDesc, srv)

}

func _Fileserver_ServeFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {

	in := new(ServeRequest)

	if err := dec(in); err != nil {

		return nil, err

	}

	if interceptor == nil {

		return srv.(FileserverServer).ServeFile(ctx, in)

	}

	info := &grpc.UnaryServerInfo{

		Server: srv,

		FullMethod: "/fileserver.Fileserver/ServeFile",
	}

	handler := func(ctx context.Context, req interface{}) (interface{}, error) {

		return srv.(FileserverServer).ServeFile(ctx, req.(*ServeRequest))

	}

	return interceptor(ctx, in, info, handler)

}

func _Fileserver_UnServeFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {

	in := new(UnServeRequest)

	if err := dec(in); err != nil {

		return nil, err

	}

	if interceptor == nil {

		return srv.(FileserverServer).UnServeFile(ctx, in)

	}

	info := &grpc.UnaryServerInfo{

		Server: srv,

		FullMethod: "/fileserver.Fileserver/UnServeFile",
	}

	handler := func(ctx context.Context, req interface{}) (interface{}, error) {

		return srv.(FileserverServer).UnServeFile(ctx, req.(*UnServeRequest))

	}

	return interceptor(ctx, in, info, handler)

}

func _Fileserver_ListServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {

	in := new(ListServeRequest)

	if err := dec(in); err != nil {

		return nil, err

	}

	if interceptor == nil {

		return srv.(FileserverServer).ListServer(ctx, in)

	}

	info := &grpc.UnaryServerInfo{

		Server: srv,

		FullMethod: "/fileserver.Fileserver/ListServer",
	}

	handler := func(ctx context.Context, req interface{}) (interface{}, error) {

		return srv.(FileserverServer).ListServer(ctx, req.(*ListServeRequest))

	}

	return interceptor(ctx, in, info, handler)

}

var _Fileserver_serviceDesc = grpc.ServiceDesc{

	ServiceName: "fileserver.Fileserver",

	HandlerType: (*FileserverServer)(nil),

	Methods: []grpc.MethodDesc{

		{

			MethodName: "ServeFile",

			Handler: _Fileserver_ServeFile_Handler,
		},

		{

			MethodName: "UnServeFile",

			Handler: _Fileserver_UnServeFile_Handler,
		},

		{

			MethodName: "ListServer",

			Handler: _Fileserver_ListServer_Handler,
		},
	},

	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{

	// 381 bytes of a gzipped FileDescriptorProto

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0xd1, 0x4e, 0xf2, 0x30,

	0x18, 0xfd, 0xf7, 0xb3, 0x01, 0xfb, 0x06, 0xe4, 0xff, 0x1b, 0x63, 0x2a, 0x51, 0x43, 0x16, 0x63,

	0xb8, 0xe2, 0x02, 0x6f, 0x8c, 0x77, 0x1a, 0x25, 0x26, 0x62, 0x42, 0x26, 0x3c, 0xc0, 0x84, 0x12,

	0x9a, 0x00, 0x9d, 0x6d, 0x67, 0x7c, 0x0a, 0x1f, 0xc1, 0xb7, 0xf3, 0x3d, 0x6c, 0xbb, 0xc2, 0x18,

	0xa2, 0x17, 0xc6, 0xbb, 0x7e, 0xe7, 0x7c, 0x3d, 0x3b, 0x3d, 0x27, 0x03, 0x3f, 0x4e, 0x68, 0x27,

	0xe1, 0x4c, 0x32, 0x04, 0x53, 0x3a, 0x27, 0x82, 0xf0, 0x67, 0xc2, 0xc3, 0x11, 0xd4, 0x1f, 0xf4,

	0x29, 0x22, 0x22, 0x61, 0x4b, 0x41, 0xd0, 0xc9, 0x16, 0x80, 0x9d, 0x96, 0xd3, 0xf6, 0xa3, 0xba,

	0x28, 0x6c, 0x1d, 0x03, 0x98, 0x2d, 0x3e, 0x60, 0x5c, 0xe2, 0xbf, 0x66, 0x05, 0xc4, 0x1a, 0x09,

	0xdf, 0x1c, 0xa8, 0x59, 0x99, 0xa7, 0x94, 0x08, 0x89, 0x10, 0xb8, 0xa3, 0x94, 0x4e, 0xac, 0x9a,

	0x9b, 0xaa, 0xb3, 0xc6, 0x06, 0x8c, 0xcd, 0xed, 0x75, 0x37, 0x51, 0x67, 0xd4, 0x84, 0x6a, 0x4f,

	0xb9, 0x1b, 0xc4, 0x72, 0x86, 0x4b, 0x06, 0xaf, 0x4e, 0xed, 0xbc, 0xe2, 0xae, 0x63, 0x19, 0x63,

	0x57, 0x71, 0xb5, 0x8c, 0xd3, 0x33, 0xc2, 0x50, 0x19, 0xd2, 0x05, 0x61, 0xa9, 0xc4, 0x9e, 0xa2,

	0xbc, 0xa8, 0x22, 0xb3, 0x11, 0xed, 0x81, 0xd7, 0xa7, 0x0b, 0x2a, 0x71, 0xd9, 0xe0, 0xde, 0x5c,

	0x0f, 0xe1, 0x10, 0x1a, 0xa3, 0xe5, 0x6f, 0x3b, 0x0c, 0x2f, 0xe0, 0x5f, 0x9f, 0x0a, 0xf9, 0x13,

	0xdd, 0xf0, 0x1e, 0xfe, 0x6f, 0xdc, 0xb5, 0x39, 0x9f, 0x43, 0x60, 0x80, 0x89, 0xfe, 0xa4, 0x50,

	0x1a, 0xa5, 0x76, 0xd0, 0xdd, 0xef, 0xe4, 0x05, 0x76, 0x72, 0x3a, 0x0a, 0x44, 0xbe, 0x1a, 0xbe,

	0x3a, 0xb6, 0x22, 0x33, 0x17, 0x5c, 0x3b, 0xbb, 0x73, 0xbd, 0x8d, 0xc5, 0xcc, 0x3a, 0x32, 0x9c,

	0x9e, 0x51, 0x0b, 0x82, 0xcb, 0xf1, 0x98, 0x08, 0x91, 0x65, 0x58, 0x32, 0x19, 0x06, 0x71, 0x0e,

	0xa1, 0x53, 0x68, 0xdc, 0xbc, 0x24, 0x94, 0xc7, 0x92, 0xb2, 0xa5, 0xee, 0xc0, 0x74, 0xe3, 0x47,

	0x0d, 0x52, 0x40, 0xbb, 0xef, 0xca, 0x50, 0x6f, 0xed, 0x1b, 0x5d, 0x81, 0x6f, 0xec, 0x19, 0x77,

	0xf8, 0xd3, 0x8b, 0x6c, 0x7a, 0xcd, 0x83, 0x1d, 0x4c, 0x96, 0x4d, 0xf8, 0x07, 0xf5, 0x20, 0xb0,

	0x25, 0x66, 0x6f, 0xdc, 0xdc, 0x2d, 0xb6, 0xfb, 0xbd, 0xce, 0x1d, 0xc0, 0x3a, 0x7a, 0x8e, 0x0e,

	0x37, 0x57, 0xb7, 0xeb, 0x6c, 0x1e, 0x7d, 0xc1, 0xae, 0xc4, 0x1e, 0xcb, 0xe6, 0x27, 0x3b, 0xfb,

	0x08, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x05, 0x3e, 0xbd, 0x71, 0x03, 0x00, 0x00,
}
