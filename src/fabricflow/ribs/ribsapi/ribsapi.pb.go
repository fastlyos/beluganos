// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ribsapi.proto

/*
Package ribsapi is a generated protocol buffer package.

It is generated from these files:
	ribsapi.proto

It has these top-level messages:
	GetRicsRequest
	GetNexthopsRequest
	GetNexthopMapRequest
	ModRibReply
	SyncRibReply
	SyncRibRequest
	MonitorRibRequest
	RibUpdate
	RicEntry
	Path
	Nexthop
	NexthopMap
*/
package ribsapi

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

type GetRicsRequest struct {
}

func (m *GetRicsRequest) Reset()                    { *m = GetRicsRequest{} }
func (m *GetRicsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRicsRequest) ProtoMessage()               {}
func (*GetRicsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetNexthopsRequest struct {
}

func (m *GetNexthopsRequest) Reset()                    { *m = GetNexthopsRequest{} }
func (m *GetNexthopsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNexthopsRequest) ProtoMessage()               {}
func (*GetNexthopsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetNexthopMapRequest struct {
}

func (m *GetNexthopMapRequest) Reset()                    { *m = GetNexthopMapRequest{} }
func (m *GetNexthopMapRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNexthopMapRequest) ProtoMessage()               {}
func (*GetNexthopMapRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ModRibReply struct {
}

func (m *ModRibReply) Reset()                    { *m = ModRibReply{} }
func (m *ModRibReply) String() string            { return proto.CompactTextString(m) }
func (*ModRibReply) ProtoMessage()               {}
func (*ModRibReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type SyncRibReply struct {
}

func (m *SyncRibReply) Reset()                    { *m = SyncRibReply{} }
func (m *SyncRibReply) String() string            { return proto.CompactTextString(m) }
func (*SyncRibReply) ProtoMessage()               {}
func (*SyncRibReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type SyncRibRequest struct {
	Rt string `protobuf:"bytes,1,opt,name=rt" json:"rt,omitempty"`
}

func (m *SyncRibRequest) Reset()                    { *m = SyncRibRequest{} }
func (m *SyncRibRequest) String() string            { return proto.CompactTextString(m) }
func (*SyncRibRequest) ProtoMessage()               {}
func (*SyncRibRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SyncRibRequest) GetRt() string {
	if m != nil {
		return m.Rt
	}
	return ""
}

type MonitorRibRequest struct {
	Rt  string `protobuf:"bytes,1,opt,name=rt" json:"rt,omitempty"`
	NId uint32 `protobuf:"varint,2,opt,name=n_id,json=nId" json:"n_id,omitempty"`
}

func (m *MonitorRibRequest) Reset()                    { *m = MonitorRibRequest{} }
func (m *MonitorRibRequest) String() string            { return proto.CompactTextString(m) }
func (*MonitorRibRequest) ProtoMessage()               {}
func (*MonitorRibRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MonitorRibRequest) GetRt() string {
	if m != nil {
		return m.Rt
	}
	return ""
}

func (m *MonitorRibRequest) GetNId() uint32 {
	if m != nil {
		return m.NId
	}
	return 0
}

type RibUpdate struct {
	Rt     string   `protobuf:"bytes,1,opt,name=rt" json:"rt,omitempty"`
	Prefix string   `protobuf:"bytes,2,opt,name=prefix" json:"prefix,omitempty"`
	Paths  [][]byte `protobuf:"bytes,3,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (m *RibUpdate) Reset()                    { *m = RibUpdate{} }
func (m *RibUpdate) String() string            { return proto.CompactTextString(m) }
func (*RibUpdate) ProtoMessage()               {}
func (*RibUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RibUpdate) GetRt() string {
	if m != nil {
		return m.Rt
	}
	return ""
}

func (m *RibUpdate) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *RibUpdate) GetPaths() [][]byte {
	if m != nil {
		return m.Paths
	}
	return nil
}

type RicEntry struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	NId   uint32 `protobuf:"varint,2,opt,name=n_id,json=nId" json:"n_id,omitempty"`
	Addr  string `protobuf:"bytes,3,opt,name=addr" json:"addr,omitempty"`
	Port  uint32 `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`
	Rt    string `protobuf:"bytes,5,opt,name=rt" json:"rt,omitempty"`
	Rd    string `protobuf:"bytes,6,opt,name=rd" json:"rd,omitempty"`
	Label uint32 `protobuf:"varint,7,opt,name=label" json:"label,omitempty"`
}

func (m *RicEntry) Reset()                    { *m = RicEntry{} }
func (m *RicEntry) String() string            { return proto.CompactTextString(m) }
func (*RicEntry) ProtoMessage()               {}
func (*RicEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RicEntry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RicEntry) GetNId() uint32 {
	if m != nil {
		return m.NId
	}
	return 0
}

func (m *RicEntry) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *RicEntry) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *RicEntry) GetRt() string {
	if m != nil {
		return m.Rt
	}
	return ""
}

func (m *RicEntry) GetRd() string {
	if m != nil {
		return m.Rd
	}
	return ""
}

func (m *RicEntry) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

type Path struct {
	Prefix  string `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
	Nexthop string `protobuf:"bytes,2,opt,name=nexthop" json:"nexthop,omitempty"`
}

func (m *Path) Reset()                    { *m = Path{} }
func (m *Path) String() string            { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()               {}
func (*Path) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Path) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Path) GetNexthop() string {
	if m != nil {
		return m.Nexthop
	}
	return ""
}

type Nexthop struct {
	Addr  string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Rt    string `protobuf:"bytes,2,opt,name=rt" json:"rt,omitempty"`
	SrcId string `protobuf:"bytes,3,opt,name=src_id,json=srcId" json:"src_id,omitempty"`
}

func (m *Nexthop) Reset()                    { *m = Nexthop{} }
func (m *Nexthop) String() string            { return proto.CompactTextString(m) }
func (*Nexthop) ProtoMessage()               {}
func (*Nexthop) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Nexthop) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Nexthop) GetRt() string {
	if m != nil {
		return m.Rt
	}
	return ""
}

func (m *Nexthop) GetSrcId() string {
	if m != nil {
		return m.SrcId
	}
	return ""
}

type NexthopMap struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=val" json:"val,omitempty"`
}

func (m *NexthopMap) Reset()                    { *m = NexthopMap{} }
func (m *NexthopMap) String() string            { return proto.CompactTextString(m) }
func (*NexthopMap) ProtoMessage()               {}
func (*NexthopMap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *NexthopMap) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *NexthopMap) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRicsRequest)(nil), "ribsapi.GetRicsRequest")
	proto.RegisterType((*GetNexthopsRequest)(nil), "ribsapi.GetNexthopsRequest")
	proto.RegisterType((*GetNexthopMapRequest)(nil), "ribsapi.GetNexthopMapRequest")
	proto.RegisterType((*ModRibReply)(nil), "ribsapi.ModRibReply")
	proto.RegisterType((*SyncRibReply)(nil), "ribsapi.SyncRibReply")
	proto.RegisterType((*SyncRibRequest)(nil), "ribsapi.SyncRibRequest")
	proto.RegisterType((*MonitorRibRequest)(nil), "ribsapi.MonitorRibRequest")
	proto.RegisterType((*RibUpdate)(nil), "ribsapi.RibUpdate")
	proto.RegisterType((*RicEntry)(nil), "ribsapi.RicEntry")
	proto.RegisterType((*Path)(nil), "ribsapi.Path")
	proto.RegisterType((*Nexthop)(nil), "ribsapi.Nexthop")
	proto.RegisterType((*NexthopMap)(nil), "ribsapi.NexthopMap")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RIBSApi service

type RIBSApiClient interface {
	GetRics(ctx context.Context, in *GetRicsRequest, opts ...grpc.CallOption) (RIBSApi_GetRicsClient, error)
	GetNexthops(ctx context.Context, in *GetNexthopsRequest, opts ...grpc.CallOption) (RIBSApi_GetNexthopsClient, error)
	GetNexthopMap(ctx context.Context, in *GetNexthopMapRequest, opts ...grpc.CallOption) (RIBSApi_GetNexthopMapClient, error)
}

type rIBSApiClient struct {
	cc *grpc.ClientConn
}

func NewRIBSApiClient(cc *grpc.ClientConn) RIBSApiClient {
	return &rIBSApiClient{cc}
}

func (c *rIBSApiClient) GetRics(ctx context.Context, in *GetRicsRequest, opts ...grpc.CallOption) (RIBSApi_GetRicsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RIBSApi_serviceDesc.Streams[0], c.cc, "/ribsapi.RIBSApi/GetRics", opts...)
	if err != nil {
		return nil, err
	}
	x := &rIBSApiGetRicsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RIBSApi_GetRicsClient interface {
	Recv() (*RicEntry, error)
	grpc.ClientStream
}

type rIBSApiGetRicsClient struct {
	grpc.ClientStream
}

func (x *rIBSApiGetRicsClient) Recv() (*RicEntry, error) {
	m := new(RicEntry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rIBSApiClient) GetNexthops(ctx context.Context, in *GetNexthopsRequest, opts ...grpc.CallOption) (RIBSApi_GetNexthopsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RIBSApi_serviceDesc.Streams[1], c.cc, "/ribsapi.RIBSApi/GetNexthops", opts...)
	if err != nil {
		return nil, err
	}
	x := &rIBSApiGetNexthopsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RIBSApi_GetNexthopsClient interface {
	Recv() (*Nexthop, error)
	grpc.ClientStream
}

type rIBSApiGetNexthopsClient struct {
	grpc.ClientStream
}

func (x *rIBSApiGetNexthopsClient) Recv() (*Nexthop, error) {
	m := new(Nexthop)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rIBSApiClient) GetNexthopMap(ctx context.Context, in *GetNexthopMapRequest, opts ...grpc.CallOption) (RIBSApi_GetNexthopMapClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RIBSApi_serviceDesc.Streams[2], c.cc, "/ribsapi.RIBSApi/GetNexthopMap", opts...)
	if err != nil {
		return nil, err
	}
	x := &rIBSApiGetNexthopMapClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RIBSApi_GetNexthopMapClient interface {
	Recv() (*NexthopMap, error)
	grpc.ClientStream
}

type rIBSApiGetNexthopMapClient struct {
	grpc.ClientStream
}

func (x *rIBSApiGetNexthopMapClient) Recv() (*NexthopMap, error) {
	m := new(NexthopMap)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RIBSApi service

type RIBSApiServer interface {
	GetRics(*GetRicsRequest, RIBSApi_GetRicsServer) error
	GetNexthops(*GetNexthopsRequest, RIBSApi_GetNexthopsServer) error
	GetNexthopMap(*GetNexthopMapRequest, RIBSApi_GetNexthopMapServer) error
}

func RegisterRIBSApiServer(s *grpc.Server, srv RIBSApiServer) {
	s.RegisterService(&_RIBSApi_serviceDesc, srv)
}

func _RIBSApi_GetRics_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRicsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RIBSApiServer).GetRics(m, &rIBSApiGetRicsServer{stream})
}

type RIBSApi_GetRicsServer interface {
	Send(*RicEntry) error
	grpc.ServerStream
}

type rIBSApiGetRicsServer struct {
	grpc.ServerStream
}

func (x *rIBSApiGetRicsServer) Send(m *RicEntry) error {
	return x.ServerStream.SendMsg(m)
}

func _RIBSApi_GetNexthops_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetNexthopsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RIBSApiServer).GetNexthops(m, &rIBSApiGetNexthopsServer{stream})
}

type RIBSApi_GetNexthopsServer interface {
	Send(*Nexthop) error
	grpc.ServerStream
}

type rIBSApiGetNexthopsServer struct {
	grpc.ServerStream
}

func (x *rIBSApiGetNexthopsServer) Send(m *Nexthop) error {
	return x.ServerStream.SendMsg(m)
}

func _RIBSApi_GetNexthopMap_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetNexthopMapRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RIBSApiServer).GetNexthopMap(m, &rIBSApiGetNexthopMapServer{stream})
}

type RIBSApi_GetNexthopMapServer interface {
	Send(*NexthopMap) error
	grpc.ServerStream
}

type rIBSApiGetNexthopMapServer struct {
	grpc.ServerStream
}

func (x *rIBSApiGetNexthopMapServer) Send(m *NexthopMap) error {
	return x.ServerStream.SendMsg(m)
}

var _RIBSApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ribsapi.RIBSApi",
	HandlerType: (*RIBSApiServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRics",
			Handler:       _RIBSApi_GetRics_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetNexthops",
			Handler:       _RIBSApi_GetNexthops_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetNexthopMap",
			Handler:       _RIBSApi_GetNexthopMap_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ribsapi.proto",
}

// Client API for RIBSCoreApi service

type RIBSCoreApiClient interface {
	ModRib(ctx context.Context, in *RibUpdate, opts ...grpc.CallOption) (*ModRibReply, error)
	MonitorRib(ctx context.Context, in *MonitorRibRequest, opts ...grpc.CallOption) (RIBSCoreApi_MonitorRibClient, error)
	SyncRib(ctx context.Context, in *SyncRibRequest, opts ...grpc.CallOption) (*SyncRibReply, error)
}

type rIBSCoreApiClient struct {
	cc *grpc.ClientConn
}

func NewRIBSCoreApiClient(cc *grpc.ClientConn) RIBSCoreApiClient {
	return &rIBSCoreApiClient{cc}
}

func (c *rIBSCoreApiClient) ModRib(ctx context.Context, in *RibUpdate, opts ...grpc.CallOption) (*ModRibReply, error) {
	out := new(ModRibReply)
	err := grpc.Invoke(ctx, "/ribsapi.RIBSCoreApi/ModRib", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rIBSCoreApiClient) MonitorRib(ctx context.Context, in *MonitorRibRequest, opts ...grpc.CallOption) (RIBSCoreApi_MonitorRibClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RIBSCoreApi_serviceDesc.Streams[0], c.cc, "/ribsapi.RIBSCoreApi/MonitorRib", opts...)
	if err != nil {
		return nil, err
	}
	x := &rIBSCoreApiMonitorRibClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RIBSCoreApi_MonitorRibClient interface {
	Recv() (*RibUpdate, error)
	grpc.ClientStream
}

type rIBSCoreApiMonitorRibClient struct {
	grpc.ClientStream
}

func (x *rIBSCoreApiMonitorRibClient) Recv() (*RibUpdate, error) {
	m := new(RibUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rIBSCoreApiClient) SyncRib(ctx context.Context, in *SyncRibRequest, opts ...grpc.CallOption) (*SyncRibReply, error) {
	out := new(SyncRibReply)
	err := grpc.Invoke(ctx, "/ribsapi.RIBSCoreApi/SyncRib", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RIBSCoreApi service

type RIBSCoreApiServer interface {
	ModRib(context.Context, *RibUpdate) (*ModRibReply, error)
	MonitorRib(*MonitorRibRequest, RIBSCoreApi_MonitorRibServer) error
	SyncRib(context.Context, *SyncRibRequest) (*SyncRibReply, error)
}

func RegisterRIBSCoreApiServer(s *grpc.Server, srv RIBSCoreApiServer) {
	s.RegisterService(&_RIBSCoreApi_serviceDesc, srv)
}

func _RIBSCoreApi_ModRib_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RibUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RIBSCoreApiServer).ModRib(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ribsapi.RIBSCoreApi/ModRib",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RIBSCoreApiServer).ModRib(ctx, req.(*RibUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _RIBSCoreApi_MonitorRib_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MonitorRibRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RIBSCoreApiServer).MonitorRib(m, &rIBSCoreApiMonitorRibServer{stream})
}

type RIBSCoreApi_MonitorRibServer interface {
	Send(*RibUpdate) error
	grpc.ServerStream
}

type rIBSCoreApiMonitorRibServer struct {
	grpc.ServerStream
}

func (x *rIBSCoreApiMonitorRibServer) Send(m *RibUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func _RIBSCoreApi_SyncRib_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRibRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RIBSCoreApiServer).SyncRib(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ribsapi.RIBSCoreApi/SyncRib",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RIBSCoreApiServer).SyncRib(ctx, req.(*SyncRibRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RIBSCoreApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ribsapi.RIBSCoreApi",
	HandlerType: (*RIBSCoreApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ModRib",
			Handler:    _RIBSCoreApi_ModRib_Handler,
		},
		{
			MethodName: "SyncRib",
			Handler:    _RIBSCoreApi_SyncRib_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorRib",
			Handler:       _RIBSCoreApi_MonitorRib_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ribsapi.proto",
}

func init() { proto.RegisterFile("ribsapi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x9d, 0x34, 0x6d, 0x42, 0x6f, 0xa7, 0x55, 0x7b, 0xe9, 0x0c, 0x51, 0x10, 0x52, 0xe5, 0x55,
	0x57, 0xa3, 0xd1, 0x80, 0x10, 0x88, 0xcd, 0xf0, 0x52, 0xd5, 0x45, 0x11, 0xf2, 0x88, 0x35, 0x72,
	0x62, 0xa3, 0x5a, 0x54, 0x89, 0x71, 0x0c, 0x9a, 0x7e, 0x04, 0x3f, 0xc5, 0x0f, 0xf0, 0x4b, 0x28,
	0x8e, 0xeb, 0xa6, 0x9d, 0xb2, 0xf3, 0xb9, 0x2f, 0x9f, 0x63, 0x9f, 0x0b, 0x43, 0x2d, 0xb3, 0x8a,
	0x29, 0x79, 0xa5, 0x74, 0x69, 0x4a, 0x8c, 0x1d, 0x24, 0x63, 0x18, 0x2d, 0x84, 0xa1, 0x32, 0xaf,
	0xa8, 0xf8, 0xf1, 0x53, 0x54, 0x86, 0x4c, 0x01, 0x17, 0xc2, 0x7c, 0x12, 0xf7, 0x66, 0x5d, 0x2a,
	0x1f, 0xbd, 0x84, 0xe9, 0x3e, 0xba, 0x62, 0x6a, 0x17, 0x1f, 0xc2, 0x60, 0x55, 0x72, 0x2a, 0x33,
	0x2a, 0xd4, 0x66, 0x4b, 0x46, 0x70, 0x7e, 0xb7, 0x2d, 0x72, 0x8f, 0x67, 0x30, 0xf2, 0xd8, 0x36,
	0xe0, 0x08, 0x3a, 0xda, 0x24, 0xc1, 0x2c, 0x98, 0xf7, 0x69, 0x47, 0x1b, 0xf2, 0x12, 0x26, 0xab,
	0xb2, 0x90, 0xa6, 0xd4, 0xff, 0x2f, 0xc2, 0x09, 0x74, 0x8b, 0xaf, 0x92, 0x27, 0x9d, 0x59, 0x30,
	0x1f, 0xd2, 0xb0, 0x58, 0x72, 0xb2, 0x84, 0x3e, 0x95, 0xd9, 0x17, 0xc5, 0x99, 0x11, 0x0f, 0xea,
	0x2f, 0x21, 0x52, 0x5a, 0x7c, 0x93, 0xf7, 0xb6, 0xa3, 0x4f, 0x1d, 0xc2, 0x29, 0xf4, 0x14, 0x33,
	0xeb, 0x2a, 0x09, 0x67, 0xe1, 0xfc, 0x9c, 0x36, 0x80, 0xfc, 0x0e, 0xe0, 0x11, 0x95, 0xf9, 0xc7,
	0xc2, 0xe8, 0x2d, 0x8e, 0x21, 0xfc, 0x2e, 0xb6, 0x6e, 0x56, 0x7d, 0x3c, 0x71, 0x39, 0x22, 0x74,
	0x19, 0xe7, 0x3a, 0x09, 0x6d, 0x95, 0x3d, 0xd7, 0x31, 0x55, 0x6a, 0x93, 0x74, 0x6d, 0x99, 0x3d,
	0x3b, 0x5e, 0x3d, 0xcf, 0xab, 0xc6, 0x3c, 0x89, 0x1c, 0xe6, 0x35, 0x9f, 0x0d, 0xcb, 0xc4, 0x26,
	0x89, 0x6d, 0x53, 0x03, 0xc8, 0x2b, 0xe8, 0x7e, 0x66, 0x66, 0xdd, 0x52, 0x11, 0x1c, 0xa8, 0x48,
	0x20, 0x2e, 0x9a, 0x8f, 0x70, 0xf2, 0x76, 0x90, 0x7c, 0x80, 0xd8, 0x7d, 0x91, 0xa7, 0x18, 0xb4,
	0x28, 0x36, 0x74, 0x3a, 0x9e, 0xce, 0x05, 0x44, 0x95, 0xce, 0x6b, 0x6d, 0x8d, 0x90, 0x5e, 0xa5,
	0xf3, 0x25, 0x27, 0xd7, 0x00, 0xfb, 0x8f, 0x3e, 0xf1, 0x20, 0x63, 0x08, 0x7f, 0xb1, 0x8d, 0x9b,
	0x53, 0x1f, 0x6f, 0xfe, 0x06, 0x10, 0xd3, 0xe5, 0xbb, 0xbb, 0xb7, 0x4a, 0xe2, 0x6b, 0x88, 0x9d,
	0xa3, 0xf0, 0xc9, 0xd5, 0xce, 0x75, 0x87, 0x1e, 0x4b, 0x27, 0x3e, 0xb1, 0x7b, 0x77, 0x72, 0x76,
	0x1d, 0xe0, 0x2d, 0x0c, 0x5a, 0xd6, 0xc3, 0xa7, 0xed, 0xf6, 0x23, 0x43, 0xa6, 0x63, 0x9f, 0x74,
	0x19, 0x3b, 0x61, 0x01, 0xc3, 0x03, 0x9b, 0xe2, 0xb3, 0x13, 0x33, 0xf6, 0xf6, 0x4d, 0x1f, 0x1f,
	0x4f, 0x59, 0x31, 0x3b, 0xe8, 0xe6, 0x4f, 0x00, 0x83, 0x5a, 0xd1, 0xfb, 0x52, 0x8b, 0x5a, 0xd5,
	0x0b, 0x88, 0x1a, 0x9f, 0x23, 0xb6, 0xb8, 0x3b, 0xff, 0xa5, 0x53, 0x1f, 0x6b, 0x2f, 0xc3, 0x19,
	0xde, 0x02, 0xec, 0xcd, 0x8d, 0x69, 0xab, 0xea, 0xc8, 0xf1, 0xe9, 0x89, 0xa9, 0x56, 0xd0, 0x1b,
	0x88, 0xdd, 0x02, 0xb5, 0x5e, 0xf3, 0x70, 0xa5, 0xd2, 0x8b, 0x87, 0x09, 0x7b, 0x7d, 0x16, 0xd9,
	0x65, 0x7f, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x2a, 0x70, 0x46, 0xfd, 0x03, 0x00, 0x00,
}