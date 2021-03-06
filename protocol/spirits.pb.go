// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spirits.proto

/*
Package gogrpc is a generated protocol buffer package.

It is generated from these files:
	spirits.proto

It has these top-level messages:
	Spirit
	SpiritCreationRequest
	SpiritUpdateRequest
	SpiritDeleteRequest
	SpiritSearchRequest
	SpiritCreationResponse
	SpiritUpdateResponse
	SpiritDeleteResponse
	SpiritSearchResponse
*/
package gogrpc

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

type Spirit_SpiritType int32

const (
	Spirit_TypeUnspecified Spirit_SpiritType = 0
	Spirit_TypeRhum        Spirit_SpiritType = 1
	Spirit_TypeWine        Spirit_SpiritType = 2
	Spirit_TypeWhisky      Spirit_SpiritType = 3
	Spirit_TypeChampagne   Spirit_SpiritType = 4
	Spirit_TypeGin         Spirit_SpiritType = 5
	Spirit_TypeBeer        Spirit_SpiritType = 6
)

var Spirit_SpiritType_name = map[int32]string{
	0: "TypeUnspecified",
	1: "TypeRhum",
	2: "TypeWine",
	3: "TypeWhisky",
	4: "TypeChampagne",
	5: "TypeGin",
	6: "TypeBeer",
}
var Spirit_SpiritType_value = map[string]int32{
	"TypeUnspecified": 0,
	"TypeRhum":        1,
	"TypeWine":        2,
	"TypeWhisky":      3,
	"TypeChampagne":   4,
	"TypeGin":         5,
	"TypeBeer":        6,
}

func (x Spirit_SpiritType) String() string {
	return proto.EnumName(Spirit_SpiritType_name, int32(x))
}
func (Spirit_SpiritType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Spirit is the basic spirit representation
type Spirit struct {
	Id          string            `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Distiller   string            `protobuf:"bytes,3,opt,name=Distiller" json:"Distiller,omitempty"`
	Bottler     string            `protobuf:"bytes,4,opt,name=Bottler" json:"Bottler,omitempty"`
	Country     string            `protobuf:"bytes,5,opt,name=Country" json:"Country,omitempty"`
	Region      string            `protobuf:"bytes,6,opt,name=Region" json:"Region,omitempty"`
	Composition string            `protobuf:"bytes,7,opt,name=Composition" json:"Composition,omitempty"`
	Type        Spirit_SpiritType `protobuf:"varint,8,opt,name=type,enum=gogrpc.Spirit_SpiritType" json:"type,omitempty"`
	Age         int32             `protobuf:"varint,9,opt,name=Age" json:"Age,omitempty"`
	// google.protobuf.Timestamp BottlingDate = 10;
	BottlingDate int64   `protobuf:"varint,10,opt,name=BottlingDate" json:"BottlingDate,omitempty"`
	Score        float32 `protobuf:"fixed32,11,opt,name=Score" json:"Score,omitempty"`
	Comment      string  `protobuf:"bytes,12,opt,name=Comment" json:"Comment,omitempty"`
}

func (m *Spirit) Reset()                    { *m = Spirit{} }
func (m *Spirit) String() string            { return proto.CompactTextString(m) }
func (*Spirit) ProtoMessage()               {}
func (*Spirit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Spirit) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Spirit) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Spirit) GetDistiller() string {
	if m != nil {
		return m.Distiller
	}
	return ""
}

func (m *Spirit) GetBottler() string {
	if m != nil {
		return m.Bottler
	}
	return ""
}

func (m *Spirit) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Spirit) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Spirit) GetComposition() string {
	if m != nil {
		return m.Composition
	}
	return ""
}

func (m *Spirit) GetType() Spirit_SpiritType {
	if m != nil {
		return m.Type
	}
	return Spirit_TypeUnspecified
}

func (m *Spirit) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Spirit) GetBottlingDate() int64 {
	if m != nil {
		return m.BottlingDate
	}
	return 0
}

func (m *Spirit) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Spirit) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type SpiritCreationRequest struct {
	Spirit *Spirit `protobuf:"bytes,1,opt,name=spirit" json:"spirit,omitempty"`
}

func (m *SpiritCreationRequest) Reset()                    { *m = SpiritCreationRequest{} }
func (m *SpiritCreationRequest) String() string            { return proto.CompactTextString(m) }
func (*SpiritCreationRequest) ProtoMessage()               {}
func (*SpiritCreationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SpiritCreationRequest) GetSpirit() *Spirit {
	if m != nil {
		return m.Spirit
	}
	return nil
}

type SpiritUpdateRequest struct {
	Spirit *Spirit `protobuf:"bytes,1,opt,name=spirit" json:"spirit,omitempty"`
}

func (m *SpiritUpdateRequest) Reset()                    { *m = SpiritUpdateRequest{} }
func (m *SpiritUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*SpiritUpdateRequest) ProtoMessage()               {}
func (*SpiritUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SpiritUpdateRequest) GetSpirit() *Spirit {
	if m != nil {
		return m.Spirit
	}
	return nil
}

type SpiritDeleteRequest struct {
	Spirit *Spirit `protobuf:"bytes,1,opt,name=spirit" json:"spirit,omitempty"`
}

func (m *SpiritDeleteRequest) Reset()                    { *m = SpiritDeleteRequest{} }
func (m *SpiritDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*SpiritDeleteRequest) ProtoMessage()               {}
func (*SpiritDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SpiritDeleteRequest) GetSpirit() *Spirit {
	if m != nil {
		return m.Spirit
	}
	return nil
}

type SpiritSearchRequest struct {
	Spirit *Spirit `protobuf:"bytes,1,opt,name=spirit" json:"spirit,omitempty"`
}

func (m *SpiritSearchRequest) Reset()                    { *m = SpiritSearchRequest{} }
func (m *SpiritSearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SpiritSearchRequest) ProtoMessage()               {}
func (*SpiritSearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SpiritSearchRequest) GetSpirit() *Spirit {
	if m != nil {
		return m.Spirit
	}
	return nil
}

type SpiritCreationResponse struct {
	Spirit *Spirit `protobuf:"bytes,1,opt,name=spirit" json:"spirit,omitempty"`
}

func (m *SpiritCreationResponse) Reset()                    { *m = SpiritCreationResponse{} }
func (m *SpiritCreationResponse) String() string            { return proto.CompactTextString(m) }
func (*SpiritCreationResponse) ProtoMessage()               {}
func (*SpiritCreationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SpiritCreationResponse) GetSpirit() *Spirit {
	if m != nil {
		return m.Spirit
	}
	return nil
}

type SpiritUpdateResponse struct {
	Spirit *Spirit `protobuf:"bytes,1,opt,name=spirit" json:"spirit,omitempty"`
}

func (m *SpiritUpdateResponse) Reset()                    { *m = SpiritUpdateResponse{} }
func (m *SpiritUpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*SpiritUpdateResponse) ProtoMessage()               {}
func (*SpiritUpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SpiritUpdateResponse) GetSpirit() *Spirit {
	if m != nil {
		return m.Spirit
	}
	return nil
}

type SpiritDeleteResponse struct {
	Deleted bool `protobuf:"varint,1,opt,name=deleted" json:"deleted,omitempty"`
}

func (m *SpiritDeleteResponse) Reset()                    { *m = SpiritDeleteResponse{} }
func (m *SpiritDeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*SpiritDeleteResponse) ProtoMessage()               {}
func (*SpiritDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SpiritDeleteResponse) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

type SpiritSearchResponse struct {
	Spirit []*Spirit `protobuf:"bytes,1,rep,name=spirit" json:"spirit,omitempty"`
}

func (m *SpiritSearchResponse) Reset()                    { *m = SpiritSearchResponse{} }
func (m *SpiritSearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SpiritSearchResponse) ProtoMessage()               {}
func (*SpiritSearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SpiritSearchResponse) GetSpirit() []*Spirit {
	if m != nil {
		return m.Spirit
	}
	return nil
}

func init() {
	proto.RegisterType((*Spirit)(nil), "gogrpc.Spirit")
	proto.RegisterType((*SpiritCreationRequest)(nil), "gogrpc.SpiritCreationRequest")
	proto.RegisterType((*SpiritUpdateRequest)(nil), "gogrpc.SpiritUpdateRequest")
	proto.RegisterType((*SpiritDeleteRequest)(nil), "gogrpc.SpiritDeleteRequest")
	proto.RegisterType((*SpiritSearchRequest)(nil), "gogrpc.SpiritSearchRequest")
	proto.RegisterType((*SpiritCreationResponse)(nil), "gogrpc.SpiritCreationResponse")
	proto.RegisterType((*SpiritUpdateResponse)(nil), "gogrpc.SpiritUpdateResponse")
	proto.RegisterType((*SpiritDeleteResponse)(nil), "gogrpc.SpiritDeleteResponse")
	proto.RegisterType((*SpiritSearchResponse)(nil), "gogrpc.SpiritSearchResponse")
	proto.RegisterEnum("gogrpc.Spirit_SpiritType", Spirit_SpiritType_name, Spirit_SpiritType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SpiritService service

type SpiritServiceClient interface {
	// CreateSpirit creates a spirit and returns the created spirit
	CreateSpirit(ctx context.Context, in *SpiritCreationRequest, opts ...grpc.CallOption) (*SpiritCreationResponse, error)
	// UpdateSpirit updates an existing spirit and returns the updated spirit
	UpdateSpirit(ctx context.Context, in *SpiritUpdateRequest, opts ...grpc.CallOption) (*SpiritUpdateResponse, error)
	// DeleteSpirit deletes an existing spirit and returns true if deleted
	DeleteSpirit(ctx context.Context, in *SpiritDeleteRequest, opts ...grpc.CallOption) (*SpiritDeleteResponse, error)
	// SearchSpirit searches for spirits matching the requests and streams the found spirits
	SearchSpirit(ctx context.Context, in *SpiritSearchRequest, opts ...grpc.CallOption) (SpiritService_SearchSpiritClient, error)
}

type spiritServiceClient struct {
	cc *grpc.ClientConn
}

func NewSpiritServiceClient(cc *grpc.ClientConn) SpiritServiceClient {
	return &spiritServiceClient{cc}
}

func (c *spiritServiceClient) CreateSpirit(ctx context.Context, in *SpiritCreationRequest, opts ...grpc.CallOption) (*SpiritCreationResponse, error) {
	out := new(SpiritCreationResponse)
	err := grpc.Invoke(ctx, "/gogrpc.SpiritService/CreateSpirit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spiritServiceClient) UpdateSpirit(ctx context.Context, in *SpiritUpdateRequest, opts ...grpc.CallOption) (*SpiritUpdateResponse, error) {
	out := new(SpiritUpdateResponse)
	err := grpc.Invoke(ctx, "/gogrpc.SpiritService/UpdateSpirit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spiritServiceClient) DeleteSpirit(ctx context.Context, in *SpiritDeleteRequest, opts ...grpc.CallOption) (*SpiritDeleteResponse, error) {
	out := new(SpiritDeleteResponse)
	err := grpc.Invoke(ctx, "/gogrpc.SpiritService/DeleteSpirit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spiritServiceClient) SearchSpirit(ctx context.Context, in *SpiritSearchRequest, opts ...grpc.CallOption) (SpiritService_SearchSpiritClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SpiritService_serviceDesc.Streams[0], c.cc, "/gogrpc.SpiritService/SearchSpirit", opts...)
	if err != nil {
		return nil, err
	}
	x := &spiritServiceSearchSpiritClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SpiritService_SearchSpiritClient interface {
	Recv() (*SpiritSearchResponse, error)
	grpc.ClientStream
}

type spiritServiceSearchSpiritClient struct {
	grpc.ClientStream
}

func (x *spiritServiceSearchSpiritClient) Recv() (*SpiritSearchResponse, error) {
	m := new(SpiritSearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SpiritService service

type SpiritServiceServer interface {
	// CreateSpirit creates a spirit and returns the created spirit
	CreateSpirit(context.Context, *SpiritCreationRequest) (*SpiritCreationResponse, error)
	// UpdateSpirit updates an existing spirit and returns the updated spirit
	UpdateSpirit(context.Context, *SpiritUpdateRequest) (*SpiritUpdateResponse, error)
	// DeleteSpirit deletes an existing spirit and returns true if deleted
	DeleteSpirit(context.Context, *SpiritDeleteRequest) (*SpiritDeleteResponse, error)
	// SearchSpirit searches for spirits matching the requests and streams the found spirits
	SearchSpirit(*SpiritSearchRequest, SpiritService_SearchSpiritServer) error
}

func RegisterSpiritServiceServer(s *grpc.Server, srv SpiritServiceServer) {
	s.RegisterService(&_SpiritService_serviceDesc, srv)
}

func _SpiritService_CreateSpirit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpiritCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpiritServiceServer).CreateSpirit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gogrpc.SpiritService/CreateSpirit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpiritServiceServer).CreateSpirit(ctx, req.(*SpiritCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpiritService_UpdateSpirit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpiritUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpiritServiceServer).UpdateSpirit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gogrpc.SpiritService/UpdateSpirit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpiritServiceServer).UpdateSpirit(ctx, req.(*SpiritUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpiritService_DeleteSpirit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpiritDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpiritServiceServer).DeleteSpirit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gogrpc.SpiritService/DeleteSpirit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpiritServiceServer).DeleteSpirit(ctx, req.(*SpiritDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpiritService_SearchSpirit_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SpiritSearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SpiritServiceServer).SearchSpirit(m, &spiritServiceSearchSpiritServer{stream})
}

type SpiritService_SearchSpiritServer interface {
	Send(*SpiritSearchResponse) error
	grpc.ServerStream
}

type spiritServiceSearchSpiritServer struct {
	grpc.ServerStream
}

func (x *spiritServiceSearchSpiritServer) Send(m *SpiritSearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _SpiritService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gogrpc.SpiritService",
	HandlerType: (*SpiritServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSpirit",
			Handler:    _SpiritService_CreateSpirit_Handler,
		},
		{
			MethodName: "UpdateSpirit",
			Handler:    _SpiritService_UpdateSpirit_Handler,
		},
		{
			MethodName: "DeleteSpirit",
			Handler:    _SpiritService_DeleteSpirit_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchSpirit",
			Handler:       _SpiritService_SearchSpirit_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spirits.proto",
}

func init() { proto.RegisterFile("spirits.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5f, 0x8b, 0xd3, 0x4e,
	0x14, 0x6d, 0x92, 0x36, 0x6d, 0x6f, 0xd3, 0xfe, 0xf2, 0xbb, 0xbb, 0x2e, 0xe3, 0xba, 0x4a, 0xc8,
	0x83, 0xf4, 0xc5, 0xb2, 0xd4, 0x67, 0xff, 0xb5, 0x05, 0x59, 0x44, 0x85, 0xd4, 0x45, 0xf0, 0x2d,
	0x36, 0xd7, 0x74, 0xb0, 0xf9, 0x63, 0x32, 0x15, 0x8a, 0x9f, 0x50, 0x3f, 0x95, 0x64, 0x26, 0xa1,
	0x3b, 0xa5, 0x0b, 0xae, 0x3e, 0x25, 0xe7, 0x9e, 0x9c, 0x73, 0xe7, 0xdc, 0xdc, 0x04, 0x86, 0x65,
	0xce, 0x0b, 0x2e, 0xca, 0x49, 0x5e, 0x64, 0x22, 0x43, 0x3b, 0xce, 0xe2, 0x22, 0x5f, 0xf9, 0xbf,
	0x2c, 0xb0, 0x97, 0x92, 0xc1, 0x11, 0x98, 0x57, 0x11, 0x33, 0x3c, 0x63, 0xdc, 0x0f, 0xcc, 0xab,
	0x08, 0x11, 0xda, 0xef, 0xc2, 0x84, 0x98, 0x29, 0x2b, 0xf2, 0x1e, 0x2f, 0xa0, 0xbf, 0xe0, 0xa5,
	0xe0, 0x9b, 0x0d, 0x15, 0xcc, 0x92, 0xc4, 0xbe, 0x80, 0x0c, 0xba, 0xb3, 0x4c, 0x88, 0x8a, 0x6b,
	0x4b, 0xae, 0x81, 0x15, 0x33, 0xcf, 0xb6, 0xa9, 0x28, 0x76, 0xac, 0xa3, 0x98, 0x1a, 0xe2, 0x19,
	0xd8, 0x01, 0xc5, 0x3c, 0x4b, 0x99, 0x2d, 0x89, 0x1a, 0xa1, 0x07, 0x83, 0x79, 0x96, 0xe4, 0x59,
	0xc9, 0x45, 0x45, 0x76, 0x25, 0x79, 0xb3, 0x84, 0x4f, 0xa0, 0x2d, 0x76, 0x39, 0xb1, 0x9e, 0x67,
	0x8c, 0x47, 0xd3, 0xfb, 0x13, 0x95, 0x68, 0xa2, 0xd2, 0xd4, 0x97, 0x0f, 0xbb, 0x9c, 0x02, 0xf9,
	0x18, 0xba, 0x60, 0xbd, 0x8a, 0x89, 0xf5, 0x3d, 0x63, 0xdc, 0x09, 0xaa, 0x5b, 0xf4, 0xc1, 0x91,
	0xe7, 0xe3, 0x69, 0xbc, 0x08, 0x05, 0x31, 0xf0, 0x8c, 0xb1, 0x15, 0x68, 0x35, 0x3c, 0x85, 0xce,
	0x72, 0x95, 0x15, 0xc4, 0x06, 0x9e, 0x31, 0x36, 0x03, 0x05, 0x54, 0x9c, 0x24, 0xa1, 0x54, 0x30,
	0xa7, 0x89, 0x23, 0xa1, 0xff, 0x03, 0x60, 0xdf, 0x19, 0x4f, 0xe0, 0xbf, 0xea, 0x7a, 0x9d, 0x96,
	0x39, 0xad, 0xf8, 0x17, 0x4e, 0x91, 0xdb, 0x42, 0x07, 0x7a, 0xf2, 0x58, 0xeb, 0x6d, 0xe2, 0x1a,
	0x0d, 0xfa, 0xc8, 0x53, 0x72, 0x4d, 0x1c, 0x01, 0x48, 0xb4, 0xe6, 0xe5, 0xd7, 0x9d, 0x6b, 0xe1,
	0xff, 0x30, 0xac, 0xf0, 0x7c, 0x1d, 0x26, 0x79, 0x18, 0xa7, 0xe4, 0xb6, 0x71, 0x00, 0xdd, 0xaa,
	0xf4, 0x9a, 0xa7, 0x6e, 0xa7, 0x51, 0xcf, 0x88, 0x0a, 0xd7, 0xf6, 0x5f, 0xc0, 0x3d, 0xd5, 0x7c,
	0x5e, 0x50, 0x58, 0xcd, 0x28, 0xa0, 0x6f, 0x5b, 0x2a, 0x05, 0x3e, 0x06, 0x5b, 0xbd, 0x7e, 0xf9,
	0x7a, 0x07, 0xd3, 0x91, 0x3e, 0xac, 0xa0, 0x66, 0xfd, 0x67, 0x70, 0xa2, 0x2a, 0xd7, 0x79, 0x14,
	0x0a, 0xfa, 0x6b, 0xf9, 0x82, 0x36, 0xf4, 0x0f, 0xf2, 0x25, 0x85, 0xc5, 0x6a, 0x7d, 0x57, 0xf9,
	0x4b, 0x38, 0x3b, 0x4c, 0x5f, 0xe6, 0x59, 0x5a, 0xd2, 0x1f, 0x3b, 0x3c, 0x87, 0x53, 0x3d, 0xfe,
	0x1d, 0xf5, 0x97, 0x8d, 0xbe, 0xc9, 0x5f, 0xeb, 0x19, 0x74, 0x23, 0x59, 0x51, 0x9f, 0x57, 0x2f,
	0x68, 0xe0, 0xbe, 0x63, 0x13, 0xf9, 0x48, 0x47, 0xeb, 0xf6, 0x8e, 0xd3, 0x9f, 0x26, 0x0c, 0x1b,
	0x83, 0xe2, 0x3b, 0x5f, 0x11, 0xbe, 0x07, 0x47, 0xe6, 0xa7, 0xfa, 0xab, 0x7e, 0xa8, 0x2b, 0x0f,
	0x36, 0xe3, 0xfc, 0xd1, 0x6d, 0xb4, 0x3a, 0x88, 0xdf, 0xc2, 0x37, 0xe0, 0xa8, 0x71, 0xd4, 0x86,
	0x0f, 0x74, 0x85, 0xb6, 0x29, 0xe7, 0x17, 0xc7, 0xc9, 0x9b, 0x66, 0x6a, 0x36, 0xc7, 0xcd, 0xb4,
	0xbd, 0x39, 0x34, 0xd3, 0x87, 0xea, 0xb7, 0xf0, 0x2d, 0x38, 0x6a, 0x6c, 0xc7, 0xcd, 0xb4, 0x2d,
	0x3a, 0x34, 0xd3, 0xe7, 0xed, 0xb7, 0x2e, 0x8d, 0x59, 0xef, 0x53, 0xfd, 0x53, 0xfc, 0x6c, 0xcb,
	0x7f, 0xe4, 0xd3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x21, 0x0f, 0x29, 0x34, 0x05, 0x00,
	0x00,
}
