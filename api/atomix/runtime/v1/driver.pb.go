// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/runtime/v1/driver.proto

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type InstallDriverRequest struct {
	// Types that are valid to be assigned to Driver:
	//	*InstallDriverRequest_Header
	//	*InstallDriverRequest_Chunk
	//	*InstallDriverRequest_Trailer
	Driver isInstallDriverRequest_Driver `protobuf_oneof:"driver"`
}

func (m *InstallDriverRequest) Reset()         { *m = InstallDriverRequest{} }
func (m *InstallDriverRequest) String() string { return proto.CompactTextString(m) }
func (*InstallDriverRequest) ProtoMessage()    {}
func (*InstallDriverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2691b25646b3550, []int{0}
}
func (m *InstallDriverRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InstallDriverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InstallDriverRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InstallDriverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallDriverRequest.Merge(m, src)
}
func (m *InstallDriverRequest) XXX_Size() int {
	return m.Size()
}
func (m *InstallDriverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallDriverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InstallDriverRequest proto.InternalMessageInfo

type isInstallDriverRequest_Driver interface {
	isInstallDriverRequest_Driver()
	MarshalTo([]byte) (int, error)
	Size() int
}

type InstallDriverRequest_Header struct {
	Header *DriverHeader `protobuf:"bytes,1,opt,name=header,proto3,oneof" json:"header,omitempty"`
}
type InstallDriverRequest_Chunk struct {
	Chunk *DriverChunk `protobuf:"bytes,2,opt,name=chunk,proto3,oneof" json:"chunk,omitempty"`
}
type InstallDriverRequest_Trailer struct {
	Trailer *DriverTrailer `protobuf:"bytes,3,opt,name=trailer,proto3,oneof" json:"trailer,omitempty"`
}

func (*InstallDriverRequest_Header) isInstallDriverRequest_Driver()  {}
func (*InstallDriverRequest_Chunk) isInstallDriverRequest_Driver()   {}
func (*InstallDriverRequest_Trailer) isInstallDriverRequest_Driver() {}

func (m *InstallDriverRequest) GetDriver() isInstallDriverRequest_Driver {
	if m != nil {
		return m.Driver
	}
	return nil
}

func (m *InstallDriverRequest) GetHeader() *DriverHeader {
	if x, ok := m.GetDriver().(*InstallDriverRequest_Header); ok {
		return x.Header
	}
	return nil
}

func (m *InstallDriverRequest) GetChunk() *DriverChunk {
	if x, ok := m.GetDriver().(*InstallDriverRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

func (m *InstallDriverRequest) GetTrailer() *DriverTrailer {
	if x, ok := m.GetDriver().(*InstallDriverRequest_Trailer); ok {
		return x.Trailer
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*InstallDriverRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*InstallDriverRequest_Header)(nil),
		(*InstallDriverRequest_Chunk)(nil),
		(*InstallDriverRequest_Trailer)(nil),
	}
}

type InstallDriverResponse struct {
}

func (m *InstallDriverResponse) Reset()         { *m = InstallDriverResponse{} }
func (m *InstallDriverResponse) String() string { return proto.CompactTextString(m) }
func (*InstallDriverResponse) ProtoMessage()    {}
func (*InstallDriverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2691b25646b3550, []int{1}
}
func (m *InstallDriverResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InstallDriverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InstallDriverResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InstallDriverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallDriverResponse.Merge(m, src)
}
func (m *InstallDriverResponse) XXX_Size() int {
	return m.Size()
}
func (m *InstallDriverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallDriverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InstallDriverResponse proto.InternalMessageInfo

type DriverHeader struct {
	DriverID DriverId `protobuf:"bytes,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id"`
}

func (m *DriverHeader) Reset()         { *m = DriverHeader{} }
func (m *DriverHeader) String() string { return proto.CompactTextString(m) }
func (*DriverHeader) ProtoMessage()    {}
func (*DriverHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2691b25646b3550, []int{2}
}
func (m *DriverHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DriverHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DriverHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DriverHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverHeader.Merge(m, src)
}
func (m *DriverHeader) XXX_Size() int {
	return m.Size()
}
func (m *DriverHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverHeader.DiscardUnknown(m)
}

var xxx_messageInfo_DriverHeader proto.InternalMessageInfo

func (m *DriverHeader) GetDriverID() DriverId {
	if m != nil {
		return m.DriverID
	}
	return DriverId{}
}

type DriverChunk struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *DriverChunk) Reset()         { *m = DriverChunk{} }
func (m *DriverChunk) String() string { return proto.CompactTextString(m) }
func (*DriverChunk) ProtoMessage()    {}
func (*DriverChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2691b25646b3550, []int{3}
}
func (m *DriverChunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DriverChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DriverChunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DriverChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverChunk.Merge(m, src)
}
func (m *DriverChunk) XXX_Size() int {
	return m.Size()
}
func (m *DriverChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverChunk.DiscardUnknown(m)
}

var xxx_messageInfo_DriverChunk proto.InternalMessageInfo

func (m *DriverChunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type DriverTrailer struct {
	Checksum string `protobuf:"bytes,1,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (m *DriverTrailer) Reset()         { *m = DriverTrailer{} }
func (m *DriverTrailer) String() string { return proto.CompactTextString(m) }
func (*DriverTrailer) ProtoMessage()    {}
func (*DriverTrailer) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2691b25646b3550, []int{4}
}
func (m *DriverTrailer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DriverTrailer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DriverTrailer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DriverTrailer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverTrailer.Merge(m, src)
}
func (m *DriverTrailer) XXX_Size() int {
	return m.Size()
}
func (m *DriverTrailer) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverTrailer.DiscardUnknown(m)
}

var xxx_messageInfo_DriverTrailer proto.InternalMessageInfo

func (m *DriverTrailer) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

type DriverId struct {
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *DriverId) Reset()         { *m = DriverId{} }
func (m *DriverId) String() string { return proto.CompactTextString(m) }
func (*DriverId) ProtoMessage()    {}
func (*DriverId) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2691b25646b3550, []int{5}
}
func (m *DriverId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DriverId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DriverId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DriverId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverId.Merge(m, src)
}
func (m *DriverId) XXX_Size() int {
	return m.Size()
}
func (m *DriverId) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverId.DiscardUnknown(m)
}

var xxx_messageInfo_DriverId proto.InternalMessageInfo

func (m *DriverId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DriverId) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*InstallDriverRequest)(nil), "atomix.runtime.v1.InstallDriverRequest")
	proto.RegisterType((*InstallDriverResponse)(nil), "atomix.runtime.v1.InstallDriverResponse")
	proto.RegisterType((*DriverHeader)(nil), "atomix.runtime.v1.DriverHeader")
	proto.RegisterType((*DriverChunk)(nil), "atomix.runtime.v1.DriverChunk")
	proto.RegisterType((*DriverTrailer)(nil), "atomix.runtime.v1.DriverTrailer")
	proto.RegisterType((*DriverId)(nil), "atomix.runtime.v1.DriverId")
}

func init() { proto.RegisterFile("atomix/runtime/v1/driver.proto", fileDescriptor_d2691b25646b3550) }

var fileDescriptor_d2691b25646b3550 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x6b, 0xe2, 0x40,
	0x14, 0xc0, 0x93, 0xfd, 0xa3, 0xf1, 0xa9, 0xb0, 0x3b, 0xb8, 0x6c, 0x70, 0x61, 0x74, 0x73, 0xa9,
	0x50, 0x48, 0xd0, 0x42, 0x69, 0xa1, 0x27, 0xeb, 0x21, 0x16, 0x7a, 0x19, 0x7a, 0xea, 0xa5, 0x4c,
	0x9d, 0x41, 0x83, 0x26, 0xb1, 0x93, 0x3f, 0xf4, 0x63, 0xf4, 0x63, 0x79, 0xb4, 0xb7, 0x9e, 0xa4,
	0xc4, 0x2f, 0x52, 0x9c, 0x49, 0x8a, 0xb6, 0x4a, 0x6f, 0x6f, 0xe6, 0xfd, 0x7e, 0xef, 0xcd, 0x7b,
	0x0c, 0x60, 0x1a, 0x87, 0xbe, 0xf7, 0xe8, 0x88, 0x24, 0x88, 0x3d, 0x9f, 0x3b, 0x69, 0xd7, 0x61,
	0xc2, 0x4b, 0xb9, 0xb0, 0xe7, 0x22, 0x8c, 0x43, 0xf4, 0x5b, 0xe5, 0xed, 0x3c, 0x6f, 0xa7, 0xdd,
	0x66, 0x63, 0x1c, 0x8e, 0x43, 0x99, 0x75, 0x36, 0x91, 0x02, 0xad, 0x67, 0x1d, 0x1a, 0xc3, 0x20,
	0x8a, 0xe9, 0x6c, 0x36, 0x90, 0x05, 0x08, 0x7f, 0x48, 0x78, 0x14, 0xa3, 0x73, 0x28, 0x4d, 0x38,
	0x65, 0x5c, 0x98, 0x7a, 0x5b, 0xef, 0x54, 0x7b, 0x2d, 0xfb, 0x53, 0x49, 0x5b, 0x19, 0xae, 0xc4,
	0x5c, 0x8d, 0xe4, 0x02, 0x3a, 0x85, 0x9f, 0xa3, 0x49, 0x12, 0x4c, 0xcd, 0x6f, 0xd2, 0xc4, 0x07,
	0xcd, 0xcb, 0x0d, 0xe5, 0x6a, 0x44, 0xe1, 0xe8, 0x02, 0xca, 0xb1, 0xa0, 0xde, 0x8c, 0x0b, 0xf3,
	0xbb, 0x34, 0xdb, 0x07, 0xcd, 0x1b, 0xc5, 0xb9, 0x1a, 0x29, 0x94, 0xbe, 0x01, 0x25, 0xb5, 0x02,
	0xeb, 0x2f, 0xfc, 0xf9, 0x30, 0x52, 0x34, 0x0f, 0x83, 0x88, 0x5b, 0xb7, 0x50, 0xdb, 0x7e, 0x32,
	0xba, 0x82, 0x8a, 0x52, 0xee, 0x3c, 0x96, 0x8f, 0xf9, 0xef, 0x60, 0xcb, 0x21, 0xeb, 0xff, 0x5a,
	0xac, 0x5a, 0x5a, 0xb6, 0x6a, 0x19, 0xf9, 0xcd, 0x80, 0x18, 0x2c, 0xcf, 0x59, 0xff, 0xa1, 0xba,
	0x35, 0x14, 0x42, 0xf0, 0x83, 0xd1, 0x98, 0xca, 0xaa, 0x35, 0x22, 0x63, 0xeb, 0x18, 0xea, 0x3b,
	0xaf, 0x47, 0x4d, 0x30, 0x46, 0x13, 0x3e, 0x9a, 0x46, 0x89, 0x2f, 0xc1, 0x0a, 0x79, 0x3f, 0x5b,
	0x67, 0x50, 0x74, 0x61, 0x9b, 0x62, 0x01, 0xf5, 0x79, 0xce, 0xc8, 0x18, 0x99, 0x50, 0x4e, 0xb9,
	0x88, 0xbc, 0x30, 0x90, 0x6b, 0xae, 0x90, 0xe2, 0xd8, 0x4b, 0x8a, 0x36, 0xd7, 0x34, 0xa0, 0x63,
	0x2e, 0x10, 0x83, 0xfa, 0xce, 0x3e, 0xd0, 0xd1, 0x9e, 0x21, 0xf7, 0x7d, 0x82, 0x66, 0xe7, 0x6b,
	0x50, 0xad, 0xb6, 0xa3, 0xf7, 0xcd, 0x45, 0x86, 0xf5, 0x65, 0x86, 0xf5, 0xd7, 0x0c, 0xeb, 0x4f,
	0x6b, 0xac, 0x2d, 0xd7, 0x58, 0x7b, 0x59, 0x63, 0xed, 0xbe, 0x24, 0xbf, 0xda, 0xc9, 0x5b, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x48, 0x2e, 0x39, 0xc9, 0xb5, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DriverManagerClient is the client API for DriverManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DriverManagerClient interface {
	InstallDriver(ctx context.Context, opts ...grpc.CallOption) (DriverManager_InstallDriverClient, error)
}

type driverManagerClient struct {
	cc *grpc.ClientConn
}

func NewDriverManagerClient(cc *grpc.ClientConn) DriverManagerClient {
	return &driverManagerClient{cc}
}

func (c *driverManagerClient) InstallDriver(ctx context.Context, opts ...grpc.CallOption) (DriverManager_InstallDriverClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DriverManager_serviceDesc.Streams[0], "/atomix.runtime.v1.DriverManager/InstallDriver", opts...)
	if err != nil {
		return nil, err
	}
	x := &driverManagerInstallDriverClient{stream}
	return x, nil
}

type DriverManager_InstallDriverClient interface {
	Send(*InstallDriverRequest) error
	CloseAndRecv() (*InstallDriverResponse, error)
	grpc.ClientStream
}

type driverManagerInstallDriverClient struct {
	grpc.ClientStream
}

func (x *driverManagerInstallDriverClient) Send(m *InstallDriverRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *driverManagerInstallDriverClient) CloseAndRecv() (*InstallDriverResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(InstallDriverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DriverManagerServer is the server API for DriverManager service.
type DriverManagerServer interface {
	InstallDriver(DriverManager_InstallDriverServer) error
}

// UnimplementedDriverManagerServer can be embedded to have forward compatible implementations.
type UnimplementedDriverManagerServer struct {
}

func (*UnimplementedDriverManagerServer) InstallDriver(srv DriverManager_InstallDriverServer) error {
	return status.Errorf(codes.Unimplemented, "method InstallDriver not implemented")
}

func RegisterDriverManagerServer(s *grpc.Server, srv DriverManagerServer) {
	s.RegisterService(&_DriverManager_serviceDesc, srv)
}

func _DriverManager_InstallDriver_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DriverManagerServer).InstallDriver(&driverManagerInstallDriverServer{stream})
}

type DriverManager_InstallDriverServer interface {
	SendAndClose(*InstallDriverResponse) error
	Recv() (*InstallDriverRequest, error)
	grpc.ServerStream
}

type driverManagerInstallDriverServer struct {
	grpc.ServerStream
}

func (x *driverManagerInstallDriverServer) SendAndClose(m *InstallDriverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *driverManagerInstallDriverServer) Recv() (*InstallDriverRequest, error) {
	m := new(InstallDriverRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DriverManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.runtime.v1.DriverManager",
	HandlerType: (*DriverManagerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InstallDriver",
			Handler:       _DriverManager_InstallDriver_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "atomix/runtime/v1/driver.proto",
}

func (m *InstallDriverRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstallDriverRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InstallDriverRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Driver != nil {
		{
			size := m.Driver.Size()
			i -= size
			if _, err := m.Driver.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *InstallDriverRequest_Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InstallDriverRequest_Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDriver(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *InstallDriverRequest_Chunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InstallDriverRequest_Chunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Chunk != nil {
		{
			size, err := m.Chunk.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDriver(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *InstallDriverRequest_Trailer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InstallDriverRequest_Trailer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Trailer != nil {
		{
			size, err := m.Trailer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDriver(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *InstallDriverResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstallDriverResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InstallDriverResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *DriverHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DriverHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DriverHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.DriverID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDriver(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *DriverChunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DriverChunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DriverChunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintDriver(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DriverTrailer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DriverTrailer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DriverTrailer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Checksum) > 0 {
		i -= len(m.Checksum)
		copy(dAtA[i:], m.Checksum)
		i = encodeVarintDriver(dAtA, i, uint64(len(m.Checksum)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DriverId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DriverId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DriverId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintDriver(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDriver(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDriver(dAtA []byte, offset int, v uint64) int {
	offset -= sovDriver(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InstallDriverRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Driver != nil {
		n += m.Driver.Size()
	}
	return n
}

func (m *InstallDriverRequest_Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovDriver(uint64(l))
	}
	return n
}
func (m *InstallDriverRequest_Chunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Chunk != nil {
		l = m.Chunk.Size()
		n += 1 + l + sovDriver(uint64(l))
	}
	return n
}
func (m *InstallDriverRequest_Trailer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Trailer != nil {
		l = m.Trailer.Size()
		n += 1 + l + sovDriver(uint64(l))
	}
	return n
}
func (m *InstallDriverResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *DriverHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DriverID.Size()
	n += 1 + l + sovDriver(uint64(l))
	return n
}

func (m *DriverChunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovDriver(uint64(l))
	}
	return n
}

func (m *DriverTrailer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Checksum)
	if l > 0 {
		n += 1 + l + sovDriver(uint64(l))
	}
	return n
}

func (m *DriverId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDriver(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovDriver(uint64(l))
	}
	return n
}

func sovDriver(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDriver(x uint64) (n int) {
	return sovDriver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InstallDriverRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDriver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InstallDriverRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstallDriverRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDriver
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDriver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DriverHeader{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Driver = &InstallDriverRequest_Header{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDriver
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDriver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DriverChunk{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Driver = &InstallDriverRequest_Chunk{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trailer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDriver
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDriver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DriverTrailer{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Driver = &InstallDriverRequest_Trailer{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDriver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDriver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InstallDriverResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDriver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InstallDriverResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstallDriverResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDriver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDriver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DriverHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDriver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DriverHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DriverHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DriverID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDriver
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDriver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DriverID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDriver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDriver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DriverChunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDriver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DriverChunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DriverChunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDriver
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDriver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDriver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDriver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DriverTrailer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDriver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DriverTrailer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DriverTrailer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDriver
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDriver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checksum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDriver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDriver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DriverId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDriver
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DriverId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DriverId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDriver
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDriver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDriver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDriver
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDriver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDriver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDriver
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDriver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDriver
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDriver
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDriver
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDriver
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDriver
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDriver
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDriver        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDriver          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDriver = fmt.Errorf("proto: unexpected end of group")
)