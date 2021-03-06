// Code generated by protoc-gen-go.
// source: raft/raft.proto
// DO NOT EDIT!

/*
Package raft is a generated protocol buffer package.

It is generated from these files:
	raft/raft.proto

It has these top-level messages:
	EmptyMessage
	EntryRequest
	AppendEntriesRequest
	StateMachineCommand
	KeyStateCommand
	Node
	Configuration
	DemoCommand
	Entry
	LogEntry
	AppendEntriesResponse
	RequestVoteRequest
	RequestVoteResponse
	SnapshotRequest
	SnapshotResponse
*/
package raft

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

type KeyStateCommand_KSMType int32

const (
	KeyStateCommand_UpdateKeyPiece KeyStateCommand_KSMType = 0
	KeyStateCommand_NewGeneration  KeyStateCommand_KSMType = 1
	KeyStateCommand_OwnerComplete  KeyStateCommand_KSMType = 2
)

var KeyStateCommand_KSMType_name = map[int32]string{
	0: "UpdateKeyPiece",
	1: "NewGeneration",
	2: "OwnerComplete",
}
var KeyStateCommand_KSMType_value = map[string]int32{
	"UpdateKeyPiece": 0,
	"NewGeneration":  1,
	"OwnerComplete":  2,
}

func (x KeyStateCommand_KSMType) String() string {
	return proto.EnumName(KeyStateCommand_KSMType_name, int32(x))
}
func (KeyStateCommand_KSMType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type Configuration_ConfigurationType int32

const (
	Configuration_CurrentConfiguration Configuration_ConfigurationType = 0
	Configuration_FutureConfiguration  Configuration_ConfigurationType = 1
)

var Configuration_ConfigurationType_name = map[int32]string{
	0: "CurrentConfiguration",
	1: "FutureConfiguration",
}
var Configuration_ConfigurationType_value = map[string]int32{
	"CurrentConfiguration": 0,
	"FutureConfiguration":  1,
}

func (x Configuration_ConfigurationType) String() string {
	return proto.EnumName(Configuration_ConfigurationType_name, int32(x))
}
func (Configuration_ConfigurationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

type Entry_EntryType int32

const (
	Entry_StateMachineCommand Entry_EntryType = 0
	Entry_ConfigurationChange Entry_EntryType = 1
	Entry_Demo                Entry_EntryType = 2
	Entry_KeyStateCommand     Entry_EntryType = 3
)

var Entry_EntryType_name = map[int32]string{
	0: "StateMachineCommand",
	1: "ConfigurationChange",
	2: "Demo",
	3: "KeyStateCommand",
}
var Entry_EntryType_value = map[string]int32{
	"StateMachineCommand": 0,
	"ConfigurationChange": 1,
	"Demo":                2,
	"KeyStateCommand":     3,
}

func (x Entry_EntryType) String() string {
	return proto.EnumName(Entry_EntryType_name, int32(x))
}
func (Entry_EntryType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{8, 0} }

type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()                    { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()               {}
func (*EmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EntryRequest struct {
	SenderId string `protobuf:"bytes,1,opt,name=sender_id" json:"sender_id,omitempty"`
	Entry    *Entry `protobuf:"bytes,2,opt,name=entry" json:"entry,omitempty"`
}

func (m *EntryRequest) Reset()                    { *m = EntryRequest{} }
func (m *EntryRequest) String() string            { return proto.CompactTextString(m) }
func (*EntryRequest) ProtoMessage()               {}
func (*EntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EntryRequest) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type AppendEntriesRequest struct {
	Term         uint64   `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	LeaderId     string   `protobuf:"bytes,2,opt,name=leader_id" json:"leader_id,omitempty"`
	PrevLogIndex uint64   `protobuf:"varint,3,opt,name=prev_log_index" json:"prev_log_index,omitempty"`
	PrevLogTerm  uint64   `protobuf:"varint,4,opt,name=prev_log_term" json:"prev_log_term,omitempty"`
	Entries      []*Entry `protobuf:"bytes,5,rep,name=entries" json:"entries,omitempty"`
	LeaderCommit uint64   `protobuf:"varint,6,opt,name=leader_commit" json:"leader_commit,omitempty"`
}

func (m *AppendEntriesRequest) Reset()                    { *m = AppendEntriesRequest{} }
func (m *AppendEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesRequest) ProtoMessage()               {}
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AppendEntriesRequest) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type StateMachineCommand struct {
	Type uint32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	// Used for Write command
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Offset int64  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	// Used for Write and Truncate commands
	Length int64 `protobuf:"varint,5,opt,name=length" json:"length,omitempty"`
	// Used for Link and Rename commands
	OldPath string `protobuf:"bytes,6,opt,name=old_path" json:"old_path,omitempty"`
	NewPath string `protobuf:"bytes,7,opt,name=new_path" json:"new_path,omitempty"`
	// Used for Create, Chmod and Mkdir commands
	Mode uint32 `protobuf:"varint,8,opt,name=mode" json:"mode,omitempty"`
	// Used for Utimes command
	AccessSeconds     int64 `protobuf:"varint,9,opt,name=access_seconds" json:"access_seconds,omitempty"`
	AccessNanoseconds int64 `protobuf:"varint,10,opt,name=access_nanoseconds" json:"access_nanoseconds,omitempty"`
	ModifySeconds     int64 `protobuf:"varint,11,opt,name=modify_seconds" json:"modify_seconds,omitempty"`
	ModifyNanoseconds int64 `protobuf:"varint,12,opt,name=modify_nanoseconds" json:"modify_nanoseconds,omitempty"`
}

func (m *StateMachineCommand) Reset()                    { *m = StateMachineCommand{} }
func (m *StateMachineCommand) String() string            { return proto.CompactTextString(m) }
func (*StateMachineCommand) ProtoMessage()               {}
func (*StateMachineCommand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type KeyStateCommand struct {
	Type KeyStateCommand_KSMType `protobuf:"varint,1,opt,name=type,enum=raft.KeyStateCommand_KSMType" json:"type,omitempty"`
	// UpdateKeyPiece arguments
	KeyOwner   *Node `protobuf:"bytes,2,opt,name=key_owner" json:"key_owner,omitempty"`
	KeyHolder  *Node `protobuf:"bytes,3,opt,name=key_holder" json:"key_holder,omitempty"`
	Generation int64 `protobuf:"varint,4,opt,name=generation" json:"generation,omitempty"`
	// NewGeneration arguments
	NewNode string `protobuf:"bytes,5,opt,name=new_node" json:"new_node,omitempty"`
	// OwnerCompete arguments
	OwnerComplete string `protobuf:"bytes,6,opt,name=owner_complete" json:"owner_complete,omitempty"`
}

func (m *KeyStateCommand) Reset()                    { *m = KeyStateCommand{} }
func (m *KeyStateCommand) String() string            { return proto.CompactTextString(m) }
func (*KeyStateCommand) ProtoMessage()               {}
func (*KeyStateCommand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *KeyStateCommand) GetKeyOwner() *Node {
	if m != nil {
		return m.KeyOwner
	}
	return nil
}

func (m *KeyStateCommand) GetKeyHolder() *Node {
	if m != nil {
		return m.KeyHolder
	}
	return nil
}

type Node struct {
	Ip         string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port       string `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
	CommonName string `protobuf:"bytes,3,opt,name=common_name" json:"common_name,omitempty"`
	NodeId     string `protobuf:"bytes,4,opt,name=node_id" json:"node_id,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Configuration struct {
	Type  Configuration_ConfigurationType `protobuf:"varint,1,opt,name=type,enum=raft.Configuration_ConfigurationType" json:"type,omitempty"`
	Nodes []*Node                         `protobuf:"bytes,2,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *Configuration) Reset()                    { *m = Configuration{} }
func (m *Configuration) String() string            { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()               {}
func (*Configuration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Configuration) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type DemoCommand struct {
	Number uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *DemoCommand) Reset()                    { *m = DemoCommand{} }
func (m *DemoCommand) String() string            { return proto.CompactTextString(m) }
func (*DemoCommand) ProtoMessage()               {}
func (*DemoCommand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Entry struct {
	Type       Entry_EntryType      `protobuf:"varint,1,opt,name=type,enum=raft.Entry_EntryType" json:"type,omitempty"`
	Uuid       string               `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
	Command    *StateMachineCommand `protobuf:"bytes,3,opt,name=command" json:"command,omitempty"`
	Config     *Configuration       `protobuf:"bytes,4,opt,name=config" json:"config,omitempty"`
	Demo       *DemoCommand         `protobuf:"bytes,5,opt,name=demo" json:"demo,omitempty"`
	KeyCommand *KeyStateCommand     `protobuf:"bytes,6,opt,name=key_command" json:"key_command,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Entry) GetCommand() *StateMachineCommand {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Entry) GetConfig() *Configuration {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Entry) GetDemo() *DemoCommand {
	if m != nil {
		return m.Demo
	}
	return nil
}

func (m *Entry) GetKeyCommand() *KeyStateCommand {
	if m != nil {
		return m.KeyCommand
	}
	return nil
}

type LogEntry struct {
	Term  uint64 `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	Entry *Entry `protobuf:"bytes,2,opt,name=entry" json:"entry,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *LogEntry) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type AppendEntriesResponse struct {
	Term      uint64 `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	NextIndex uint64 `protobuf:"varint,2,opt,name=next_index" json:"next_index,omitempty"`
	Success   bool   `protobuf:"varint,3,opt,name=success" json:"success,omitempty"`
}

func (m *AppendEntriesResponse) Reset()                    { *m = AppendEntriesResponse{} }
func (m *AppendEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesResponse) ProtoMessage()               {}
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type RequestVoteRequest struct {
	Term         uint64 `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	CandidateId  string `protobuf:"bytes,2,opt,name=candidate_id" json:"candidate_id,omitempty"`
	LastLogIndex uint64 `protobuf:"varint,3,opt,name=last_log_index" json:"last_log_index,omitempty"`
	LastLogTerm  uint64 `protobuf:"varint,4,opt,name=last_log_term" json:"last_log_term,omitempty"`
}

func (m *RequestVoteRequest) Reset()                    { *m = RequestVoteRequest{} }
func (m *RequestVoteRequest) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteRequest) ProtoMessage()               {}
func (*RequestVoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type RequestVoteResponse struct {
	Term        uint64 `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	VoteGranted bool   `protobuf:"varint,2,opt,name=vote_granted" json:"vote_granted,omitempty"`
}

func (m *RequestVoteResponse) Reset()                    { *m = RequestVoteResponse{} }
func (m *RequestVoteResponse) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteResponse) ProtoMessage()               {}
func (*RequestVoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type SnapshotRequest struct {
	Term              uint64 `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	LeaderId          string `protobuf:"bytes,2,opt,name=leader_id" json:"leader_id,omitempty"`
	LastIncludedIndex uint64 `protobuf:"varint,3,opt,name=last_included_index" json:"last_included_index,omitempty"`
	LastIncludedTerm  uint64 `protobuf:"varint,4,opt,name=last_included_term" json:"last_included_term,omitempty"`
	Offset            uint64 `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
	Data              []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	Done              bool   `protobuf:"varint,7,opt,name=done" json:"done,omitempty"`
}

func (m *SnapshotRequest) Reset()                    { *m = SnapshotRequest{} }
func (m *SnapshotRequest) String() string            { return proto.CompactTextString(m) }
func (*SnapshotRequest) ProtoMessage()               {}
func (*SnapshotRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type SnapshotResponse struct {
	Term uint64 `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
}

func (m *SnapshotResponse) Reset()                    { *m = SnapshotResponse{} }
func (m *SnapshotResponse) String() string            { return proto.CompactTextString(m) }
func (*SnapshotResponse) ProtoMessage()               {}
func (*SnapshotResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func init() {
	proto.RegisterType((*EmptyMessage)(nil), "raft.EmptyMessage")
	proto.RegisterType((*EntryRequest)(nil), "raft.EntryRequest")
	proto.RegisterType((*AppendEntriesRequest)(nil), "raft.AppendEntriesRequest")
	proto.RegisterType((*StateMachineCommand)(nil), "raft.StateMachineCommand")
	proto.RegisterType((*KeyStateCommand)(nil), "raft.KeyStateCommand")
	proto.RegisterType((*Node)(nil), "raft.Node")
	proto.RegisterType((*Configuration)(nil), "raft.Configuration")
	proto.RegisterType((*DemoCommand)(nil), "raft.DemoCommand")
	proto.RegisterType((*Entry)(nil), "raft.Entry")
	proto.RegisterType((*LogEntry)(nil), "raft.LogEntry")
	proto.RegisterType((*AppendEntriesResponse)(nil), "raft.AppendEntriesResponse")
	proto.RegisterType((*RequestVoteRequest)(nil), "raft.RequestVoteRequest")
	proto.RegisterType((*RequestVoteResponse)(nil), "raft.RequestVoteResponse")
	proto.RegisterType((*SnapshotRequest)(nil), "raft.SnapshotRequest")
	proto.RegisterType((*SnapshotResponse)(nil), "raft.SnapshotResponse")
	proto.RegisterEnum("raft.KeyStateCommand_KSMType", KeyStateCommand_KSMType_name, KeyStateCommand_KSMType_value)
	proto.RegisterEnum("raft.Configuration_ConfigurationType", Configuration_ConfigurationType_name, Configuration_ConfigurationType_value)
	proto.RegisterEnum("raft.Entry_EntryType", Entry_EntryType_name, Entry_EntryType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for RaftNetwork service

type RaftNetworkClient interface {
	AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error)
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error)
	ClientToLeaderRequest(ctx context.Context, in *EntryRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
	InstallSnapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotResponse, error)
}

type raftNetworkClient struct {
	cc *grpc.ClientConn
}

func NewRaftNetworkClient(cc *grpc.ClientConn) RaftNetworkClient {
	return &raftNetworkClient{cc}
}

func (c *raftNetworkClient) AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error) {
	out := new(AppendEntriesResponse)
	err := grpc.Invoke(ctx, "/raft.RaftNetwork/AppendEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftNetworkClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error) {
	out := new(RequestVoteResponse)
	err := grpc.Invoke(ctx, "/raft.RaftNetwork/RequestVote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftNetworkClient) ClientToLeaderRequest(ctx context.Context, in *EntryRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := grpc.Invoke(ctx, "/raft.RaftNetwork/ClientToLeaderRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftNetworkClient) InstallSnapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotResponse, error) {
	out := new(SnapshotResponse)
	err := grpc.Invoke(ctx, "/raft.RaftNetwork/InstallSnapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RaftNetwork service

type RaftNetworkServer interface {
	AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error)
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error)
	ClientToLeaderRequest(context.Context, *EntryRequest) (*EmptyMessage, error)
	InstallSnapshot(context.Context, *SnapshotRequest) (*SnapshotResponse, error)
}

func RegisterRaftNetworkServer(s *grpc.Server, srv RaftNetworkServer) {
	s.RegisterService(&_RaftNetwork_serviceDesc, srv)
}

func _RaftNetwork_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftNetworkServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftNetwork/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftNetworkServer).AppendEntries(ctx, req.(*AppendEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftNetwork_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftNetworkServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftNetwork/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftNetworkServer).RequestVote(ctx, req.(*RequestVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftNetwork_ClientToLeaderRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftNetworkServer).ClientToLeaderRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftNetwork/ClientToLeaderRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftNetworkServer).ClientToLeaderRequest(ctx, req.(*EntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftNetwork_InstallSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftNetworkServer).InstallSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.RaftNetwork/InstallSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftNetworkServer).InstallSnapshot(ctx, req.(*SnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RaftNetwork_serviceDesc = grpc.ServiceDesc{
	ServiceName: "raft.RaftNetwork",
	HandlerType: (*RaftNetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppendEntries",
			Handler:    _RaftNetwork_AppendEntries_Handler,
		},
		{
			MethodName: "RequestVote",
			Handler:    _RaftNetwork_RequestVote_Handler,
		},
		{
			MethodName: "ClientToLeaderRequest",
			Handler:    _RaftNetwork_ClientToLeaderRequest_Handler,
		},
		{
			MethodName: "InstallSnapshot",
			Handler:    _RaftNetwork_InstallSnapshot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 948 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x56, 0x4b, 0x8f, 0xdb, 0x54,
	0x14, 0x6e, 0xde, 0xc9, 0xc9, 0x73, 0x6e, 0x26, 0x6d, 0x26, 0x65, 0x60, 0xe4, 0x0a, 0x09, 0x15,
	0x69, 0x90, 0x52, 0x36, 0x2c, 0x58, 0x94, 0x94, 0x22, 0xda, 0xce, 0x80, 0x3a, 0x85, 0x15, 0x52,
	0xe4, 0x3a, 0x27, 0x89, 0x55, 0xe7, 0x5e, 0x63, 0x5f, 0x77, 0x9a, 0x1d, 0xff, 0x82, 0x05, 0x7b,
	0xf8, 0x07, 0xfc, 0x37, 0x76, 0x9c, 0x7b, 0xae, 0x9d, 0xc4, 0x19, 0x0b, 0xb1, 0x89, 0xe6, 0x9e,
	0xc7, 0x77, 0xbe, 0xf3, 0xf4, 0x40, 0x3f, 0x72, 0x97, 0xfa, 0x0b, 0xf3, 0x73, 0x19, 0x46, 0x4a,
	0x2b, 0x51, 0x35, 0x7f, 0x3b, 0x3d, 0xe8, 0x7c, 0xbb, 0x09, 0xf5, 0xf6, 0x0a, 0xe3, 0xd8, 0x5d,
	0xa1, 0xf3, 0x35, 0xbd, 0xa5, 0x8e, 0xb6, 0xaf, 0xf1, 0xd7, 0x04, 0x63, 0x2d, 0x4e, 0xa0, 0x15,
	0xa3, 0x5c, 0x60, 0x34, 0xf7, 0x17, 0xe3, 0xd2, 0x45, 0xe9, 0xb3, 0x96, 0x98, 0x40, 0x0d, 0x8d,
	0xc9, 0xb8, 0x4c, 0xcf, 0xf6, 0xb4, 0x7d, 0xc9, 0xa0, 0xec, 0xe5, 0xfc, 0x51, 0x82, 0xd3, 0xa7,
	0x61, 0x48, 0x1e, 0xe6, 0xed, 0x63, 0x9c, 0xe1, 0x74, 0xa0, 0xaa, 0x31, 0xda, 0x30, 0x44, 0xd5,
	0xa0, 0x06, 0xe8, 0xa6, 0xa8, 0x65, 0x46, 0xbd, 0x0f, 0xbd, 0x30, 0xc2, 0xf7, 0xf3, 0x40, 0xad,
	0xe6, 0x3e, 0x05, 0xfc, 0x30, 0xae, 0xb0, 0xe9, 0x08, 0xba, 0x3b, 0x39, 0x23, 0x54, 0x59, 0xfc,
	0x11, 0x34, 0xd0, 0x46, 0x18, 0xd7, 0x2e, 0x2a, 0x47, 0x34, 0x8c, 0x53, 0x8a, 0xef, 0xa9, 0xcd,
	0xc6, 0xd7, 0xe3, 0xba, 0x71, 0x72, 0xfe, 0x29, 0xc1, 0xf0, 0x46, 0xbb, 0x1a, 0xaf, 0x5c, 0x6f,
	0xed, 0x4b, 0x9c, 0x91, 0xd2, 0x95, 0x0b, 0x26, 0xb7, 0x0d, 0x91, 0xc9, 0x75, 0xcd, 0x2b, 0x74,
	0xf5, 0x3a, 0xe5, 0x45, 0xaf, 0x85, 0xab, 0x5d, 0x66, 0xd3, 0x11, 0x3d, 0xa8, 0xab, 0xe5, 0x32,
	0x46, 0xcd, 0x34, 0x2a, 0xe6, 0x1d, 0xa0, 0x5c, 0x91, 0x75, 0x8d, 0xdf, 0x03, 0x68, 0xaa, 0x60,
	0x31, 0x67, 0xff, 0x3a, 0xfb, 0x93, 0x44, 0xe2, 0xad, 0x95, 0x34, 0x32, 0xc4, 0x8d, 0x5a, 0xe0,
	0xb8, 0xc9, 0xd1, 0x28, 0x6f, 0xd7, 0xf3, 0xa8, 0xfa, 0xf3, 0x18, 0x3d, 0x25, 0x17, 0xf1, 0xb8,
	0xc5, 0x48, 0x13, 0x10, 0xa9, 0x5c, 0xba, 0x52, 0x65, 0x3a, 0x60, 0x1d, 0xf9, 0x10, 0x82, 0xbf,
	0xdc, 0xee, 0x7c, 0xda, 0x99, 0x4f, 0x2a, 0x3f, 0xf4, 0xe9, 0x18, 0x9d, 0xf3, 0x5b, 0x19, 0xfa,
	0x2f, 0x71, 0xcb, 0xe9, 0x67, 0x79, 0x7f, 0x7e, 0x90, 0x77, 0x6f, 0x7a, 0x6e, 0x2b, 0x78, 0x64,
	0x74, 0xf9, 0xf2, 0xe6, 0xea, 0x0d, 0x19, 0x89, 0x73, 0x68, 0xbd, 0xc3, 0xed, 0x5c, 0xdd, 0x4a,
	0x8c, 0xd2, 0xd6, 0x83, 0xf5, 0xb8, 0xa6, 0x6c, 0xc4, 0xc7, 0x00, 0x46, 0xbd, 0xa6, 0xf4, 0x49,
	0x5f, 0xb9, 0xa3, 0x17, 0x00, 0x2b, 0x24, 0x57, 0x57, 0xfb, 0x4a, 0xa6, 0xd5, 0x4b, 0x6b, 0x23,
	0x4d, 0x35, 0x6a, 0xd9, 0x14, 0x70, 0x00, 0xd3, 0xb7, 0x30, 0x40, 0x8d, 0xb6, 0x8a, 0xce, 0x0c,
	0x1a, 0x19, 0x0f, 0x01, 0xbd, 0x9f, 0x42, 0x6a, 0x09, 0x12, 0xd1, 0x1f, 0x7d, 0xf4, 0x70, 0x70,
	0x8f, 0xe6, 0xa9, 0x7b, 0x8d, 0xb7, 0xdf, 0xed, 0xf0, 0x07, 0x25, 0x23, 0xfa, 0xc1, 0x20, 0xcd,
	0x52, 0xa0, 0x41, 0xd9, 0x79, 0x0e, 0x55, 0xa6, 0x02, 0x50, 0xf6, 0xc3, 0x74, 0x98, 0x4d, 0xb3,
	0x55, 0xa4, 0xd3, 0x66, 0x0f, 0xa1, 0x6d, 0x06, 0x46, 0x49, 0x2a, 0xe0, 0x06, 0x39, 0x8b, 0x96,
	0xe8, 0x43, 0xc3, 0x30, 0x34, 0xa3, 0x5a, 0x65, 0x32, 0x7f, 0x95, 0xa0, 0x3b, 0x53, 0x72, 0xe9,
	0xaf, 0x12, 0x1b, 0x4e, 0x3c, 0xc9, 0x15, 0xf2, 0x53, 0x9b, 0x76, 0xce, 0x24, 0xff, 0xe2, 0x44,
	0xce, 0xa0, 0x66, 0x70, 0x63, 0x8a, 0x5d, 0xc9, 0x17, 0x8b, 0x98, 0x9e, 0xdc, 0xb5, 0x1f, 0xc3,
	0xe9, 0x2c, 0x89, 0x22, 0x1a, 0xfb, 0x9c, 0x8e, 0xd2, 0x7f, 0x00, 0xc3, 0xe7, 0x89, 0x4e, 0x22,
	0xcc, 0x2b, 0x4a, 0xce, 0x39, 0xb4, 0x9f, 0xe1, 0x46, 0x65, 0xfd, 0xa6, 0x69, 0x95, 0xc9, 0xe6,
	0x2d, 0xf5, 0x87, 0xd7, 0xd0, 0xf9, 0xbb, 0x0c, 0x35, 0xbb, 0x30, 0x8f, 0x72, 0x09, 0x8c, 0x0e,
	0x76, 0xc9, 0xfe, 0x32, 0x01, 0xaa, 0x55, 0x92, 0xec, 0x16, 0xf6, 0x31, 0x34, 0x3c, 0x8b, 0x9b,
	0x76, 0xfb, 0xcc, 0x7a, 0x15, 0x2d, 0xd8, 0x23, 0xa8, 0x7b, 0x4c, 0x8d, 0x2b, 0xd8, 0x9e, 0x0e,
	0x0b, 0x2a, 0x24, 0x3e, 0xa1, 0x4d, 0x23, 0xb2, 0x3c, 0x09, 0xed, 0xe9, 0x89, 0x35, 0x39, 0xa4,
	0xff, 0x18, 0xda, 0x66, 0xc4, 0xb2, 0xa8, 0x75, 0xb6, 0x1b, 0x15, 0x4e, 0xad, 0xf3, 0x0b, 0xb4,
	0xf6, 0xc4, 0x1f, 0x14, 0xae, 0xbd, 0x2d, 0x5c, 0x8e, 0xc3, 0x6c, 0xed, 0xca, 0x15, 0xd2, 0xf4,
	0x34, 0xa1, 0x6a, 0x22, 0x0f, 0xca, 0x34, 0x12, 0xc7, 0x6b, 0x33, 0xa8, 0x38, 0x5f, 0x42, 0xf3,
	0x95, 0x5a, 0xd9, 0xd2, 0xe5, 0x2f, 0xdb, 0x7f, 0x1d, 0xc7, 0x17, 0x30, 0x3a, 0xba, 0x8d, 0x71,
	0xa8, 0x64, 0x8c, 0x47, 0x10, 0x34, 0x9f, 0x12, 0x3f, 0xe8, 0xf4, 0x0a, 0x96, 0x59, 0x46, 0x33,
	0x18, 0x27, 0x7c, 0x0e, 0xb8, 0xd8, 0x4d, 0x67, 0x05, 0x22, 0x3d, 0xad, 0x3f, 0x2b, 0x8d, 0xc5,
	0x57, 0xf6, 0x14, 0x3a, 0x1e, 0xf1, 0xf5, 0xcd, 0xb2, 0xe4, 0x0e, 0x6d, 0xe0, 0xc6, 0xba, 0xe8,
	0xd0, 0xee, 0xe4, 0xfb, 0x43, 0xeb, 0x7c, 0x05, 0xc3, 0x5c, 0xa0, 0x42, 0xca, 0x14, 0xe9, 0x3d,
	0x69, 0xe7, 0xab, 0xc8, 0x95, 0x1a, 0x6d, 0xa4, 0xa6, 0xf3, 0x7b, 0x09, 0xfa, 0x37, 0xd2, 0x0d,
	0xe3, 0xb5, 0xd2, 0xff, 0xfb, 0x3b, 0xf0, 0x10, 0x86, 0x4c, 0xc3, 0x97, 0x5e, 0x90, 0x2c, 0x70,
	0x91, 0xe3, 0x48, 0x07, 0x2e, 0xaf, 0x3c, 0xf8, 0x22, 0xec, 0x4f, 0x73, 0x8d, 0xdf, 0xd9, 0xe1,
	0xae, 0xf3, 0xe1, 0x36, 0x2f, 0x25, 0x91, 0x4f, 0x70, 0xd3, 0xb9, 0x80, 0xc1, 0x9e, 0x58, 0x51,
	0x46, 0xd3, 0x3f, 0xcb, 0xd0, 0x7e, 0x4d, 0xad, 0xbb, 0x46, 0x7d, 0xab, 0xa2, 0x77, 0xe2, 0x05,
	0x74, 0x73, 0xbd, 0x13, 0x13, 0xdb, 0xd9, 0xa2, 0x8f, 0xdd, 0xe4, 0x61, 0xa1, 0xce, 0xc6, 0x71,
	0xee, 0x89, 0x67, 0x04, 0xbd, 0x2f, 0xa9, 0x18, 0x5b, 0xeb, 0xbb, 0xed, 0x9c, 0x9c, 0x15, 0x68,
	0x76, 0x28, 0x4f, 0x61, 0x34, 0x0b, 0x7c, 0x1a, 0xb6, 0x37, 0xea, 0x15, 0xd7, 0x30, 0x2b, 0xb1,
	0x38, 0x98, 0xb9, 0x0c, 0x29, 0x93, 0x1d, 0x7e, 0xea, 0xef, 0x89, 0x6f, 0xa0, 0xff, 0xbd, 0x8c,
	0xb5, 0x1b, 0x04, 0x59, 0x35, 0x44, 0xba, 0x4e, 0x47, 0x6d, 0x9b, 0xdc, 0x3f, 0x16, 0x67, 0x34,
	0xde, 0xd6, 0xf9, 0xbf, 0x89, 0x27, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x49, 0xe8, 0x27, 0xd2,
	0x60, 0x08, 0x00, 0x00,
}
