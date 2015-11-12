// Code generated by protoc-gen-go.
// source: paranoidnetwork/paranoid.proto
// DO NOT EDIT!

/*
Package paranoid is a generated protocol buffer package.

It is generated from these files:
	paranoidnetwork/paranoid.proto

It has these top-level messages:
	EmptyMessage
	PingRequest
	CreatRequest
	WriteRequest
	WriteResponse
	LinkRequest
	UnlinkRequest
	RenameRequest
	TruncateRequest
	UtimesRequest
	ChmodRequest
*/
package paranoid

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

type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()         { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()    {}

type PingRequest struct {
	Ip   string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}

type CreatRequest struct {
	Path        string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Permissions uint32 `protobuf:"varint,2,opt,name=permissions" json:"permissions,omitempty"`
}

func (m *CreatRequest) Reset()         { *m = CreatRequest{} }
func (m *CreatRequest) String() string { return proto.CompactTextString(m) }
func (*CreatRequest) ProtoMessage()    {}

type WriteRequest struct {
	Path   string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Offset uint64 `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	Length uint64 `protobuf:"varint,4,opt,name=length" json:"length,omitempty"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}

type WriteResponse struct {
	BytesWritten uint64 `protobuf:"varint,1,opt,name=bytes_written" json:"bytes_written,omitempty"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}

type LinkRequest struct {
	OldPath string `protobuf:"bytes,1,opt,name=old_path" json:"old_path,omitempty"`
	NewPath string `protobuf:"bytes,2,opt,name=new_path" json:"new_path,omitempty"`
}

func (m *LinkRequest) Reset()         { *m = LinkRequest{} }
func (m *LinkRequest) String() string { return proto.CompactTextString(m) }
func (*LinkRequest) ProtoMessage()    {}

type UnlinkRequest struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *UnlinkRequest) Reset()         { *m = UnlinkRequest{} }
func (m *UnlinkRequest) String() string { return proto.CompactTextString(m) }
func (*UnlinkRequest) ProtoMessage()    {}

type RenameRequest struct {
	OldPath string `protobuf:"bytes,1,opt,name=old_path" json:"old_path,omitempty"`
	NewPath string `protobuf:"bytes,2,opt,name=new_path" json:"new_path,omitempty"`
}

func (m *RenameRequest) Reset()         { *m = RenameRequest{} }
func (m *RenameRequest) String() string { return proto.CompactTextString(m) }
func (*RenameRequest) ProtoMessage()    {}

type TruncateRequest struct {
	Path   string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Length uint64 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
}

func (m *TruncateRequest) Reset()         { *m = TruncateRequest{} }
func (m *TruncateRequest) String() string { return proto.CompactTextString(m) }
func (*TruncateRequest) ProtoMessage()    {}

// TODO(terry): Update this message to use microseconds by the end of
// sprint 2.
type UtimesRequest struct {
	Path               string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	AccessSeconds      uint64 `protobuf:"varint,2,opt,name=access_seconds" json:"access_seconds,omitempty"`
	AccessMicroseconds uint64 `protobuf:"varint,3,opt,name=access_microseconds" json:"access_microseconds,omitempty"`
	ModifySeconds      uint64 `protobuf:"varint,4,opt,name=modify_seconds" json:"modify_seconds,omitempty"`
	ModifyMicroseconds uint64 `protobuf:"varint,5,opt,name=modify_microseconds" json:"modify_microseconds,omitempty"`
}

func (m *UtimesRequest) Reset()         { *m = UtimesRequest{} }
func (m *UtimesRequest) String() string { return proto.CompactTextString(m) }
func (*UtimesRequest) ProtoMessage()    {}

type ChmodRequest struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Mode uint32 `protobuf:"varint,2,opt,name=mode" json:"mode,omitempty"`
}

func (m *ChmodRequest) Reset()         { *m = ChmodRequest{} }
func (m *ChmodRequest) String() string { return proto.CompactTextString(m) }
func (*ChmodRequest) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for ParanoidNetwork service

type ParanoidNetworkClient interface {
	// Used for health checking and discovery. Sends the IP and port of the
	// PFSD instance running on the client.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
	// Filesystem calls
	Creat(ctx context.Context, in *CreatRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
	Unlink(ctx context.Context, in *UnlinkRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
	Rename(ctx context.Context, in *RenameRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
	Truncate(ctx context.Context, in *TruncateRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
	Utimes(ctx context.Context, in *UtimesRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
	Chmod(ctx context.Context, in *ChmodRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
}

type paranoidNetworkClient struct {
	cc *grpc.ClientConn
}

func NewParanoidNetworkClient(cc *grpc.ClientConn) ParanoidNetworkClient {
	return &paranoidNetworkClient{cc}
}

func (c *paranoidNetworkClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/paranoid.ParanoidNetwork/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paranoidNetworkClient) Creat(ctx context.Context, in *CreatRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/paranoid.ParanoidNetwork/Creat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paranoidNetworkClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/paranoid.ParanoidNetwork/Write", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paranoidNetworkClient) Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/paranoid.ParanoidNetwork/Link", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paranoidNetworkClient) Unlink(ctx context.Context, in *UnlinkRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/paranoid.ParanoidNetwork/Unlink", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paranoidNetworkClient) Rename(ctx context.Context, in *RenameRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/paranoid.ParanoidNetwork/Rename", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paranoidNetworkClient) Truncate(ctx context.Context, in *TruncateRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/paranoid.ParanoidNetwork/Truncate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paranoidNetworkClient) Utimes(ctx context.Context, in *UtimesRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/paranoid.ParanoidNetwork/Utimes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paranoidNetworkClient) Chmod(ctx context.Context, in *ChmodRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/paranoid.ParanoidNetwork/Chmod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ParanoidNetwork service

type ParanoidNetworkServer interface {
	// Used for health checking and discovery. Sends the IP and port of the
	// PFSD instance running on the client.
	Ping(context.Context, *PingRequest) (*EmptyMessage, error)
	// Filesystem calls
	Creat(context.Context, *CreatRequest) (*EmptyMessage, error)
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	Link(context.Context, *LinkRequest) (*EmptyMessage, error)
	Unlink(context.Context, *UnlinkRequest) (*EmptyMessage, error)
	Rename(context.Context, *RenameRequest) (*EmptyMessage, error)
	Truncate(context.Context, *TruncateRequest) (*EmptyMessage, error)
	Utimes(context.Context, *UtimesRequest) (*EmptyMessage, error)
	Chmod(context.Context, *ChmodRequest) (*EmptyMessage, error)
}

func RegisterParanoidNetworkServer(s *grpc.Server, srv ParanoidNetworkServer) {
	s.RegisterService(&_ParanoidNetwork_serviceDesc, srv)
}

func _ParanoidNetwork_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ParanoidNetworkServer).Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ParanoidNetwork_Creat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ParanoidNetworkServer).Creat(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ParanoidNetwork_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ParanoidNetworkServer).Write(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ParanoidNetwork_Link_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(LinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ParanoidNetworkServer).Link(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ParanoidNetwork_Unlink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(UnlinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ParanoidNetworkServer).Unlink(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ParanoidNetwork_Rename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(RenameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ParanoidNetworkServer).Rename(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ParanoidNetwork_Truncate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(TruncateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ParanoidNetworkServer).Truncate(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ParanoidNetwork_Utimes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(UtimesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ParanoidNetworkServer).Utimes(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ParanoidNetwork_Chmod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ChmodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ParanoidNetworkServer).Chmod(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _ParanoidNetwork_serviceDesc = grpc.ServiceDesc{
	ServiceName: "paranoid.ParanoidNetwork",
	HandlerType: (*ParanoidNetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ParanoidNetwork_Ping_Handler,
		},
		{
			MethodName: "Creat",
			Handler:    _ParanoidNetwork_Creat_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _ParanoidNetwork_Write_Handler,
		},
		{
			MethodName: "Link",
			Handler:    _ParanoidNetwork_Link_Handler,
		},
		{
			MethodName: "Unlink",
			Handler:    _ParanoidNetwork_Unlink_Handler,
		},
		{
			MethodName: "Rename",
			Handler:    _ParanoidNetwork_Rename_Handler,
		},
		{
			MethodName: "Truncate",
			Handler:    _ParanoidNetwork_Truncate_Handler,
		},
		{
			MethodName: "Utimes",
			Handler:    _ParanoidNetwork_Utimes_Handler,
		},
		{
			MethodName: "Chmod",
			Handler:    _ParanoidNetwork_Chmod_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
