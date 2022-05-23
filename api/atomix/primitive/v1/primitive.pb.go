// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/primitive/v1/primitive.proto

package v1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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

type Primitive struct {
	PrimitiveMeta `protobuf:"bytes,1,opt,name=meta,proto3,embedded=meta" json:"meta"`
}

func (m *Primitive) Reset()         { *m = Primitive{} }
func (m *Primitive) String() string { return proto.CompactTextString(m) }
func (*Primitive) ProtoMessage()    {}
func (*Primitive) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce01ee29cb809ef2, []int{0}
}
func (m *Primitive) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Primitive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Primitive.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Primitive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Primitive.Merge(m, src)
}
func (m *Primitive) XXX_Size() int {
	return m.Size()
}
func (m *Primitive) XXX_DiscardUnknown() {
	xxx_messageInfo_Primitive.DiscardUnknown(m)
}

var xxx_messageInfo_Primitive proto.InternalMessageInfo

type PrimitiveMeta struct {
	ID   PrimitiveId `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Type string      `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *PrimitiveMeta) Reset()         { *m = PrimitiveMeta{} }
func (m *PrimitiveMeta) String() string { return proto.CompactTextString(m) }
func (*PrimitiveMeta) ProtoMessage()    {}
func (*PrimitiveMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce01ee29cb809ef2, []int{1}
}
func (m *PrimitiveMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrimitiveMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrimitiveMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrimitiveMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveMeta.Merge(m, src)
}
func (m *PrimitiveMeta) XXX_Size() int {
	return m.Size()
}
func (m *PrimitiveMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveMeta.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveMeta proto.InternalMessageInfo

func (m *PrimitiveMeta) GetID() PrimitiveId {
	if m != nil {
		return m.ID
	}
	return PrimitiveId{}
}

func (m *PrimitiveMeta) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type PrimitiveId struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *PrimitiveId) Reset()         { *m = PrimitiveId{} }
func (m *PrimitiveId) String() string { return proto.CompactTextString(m) }
func (*PrimitiveId) ProtoMessage()    {}
func (*PrimitiveId) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce01ee29cb809ef2, []int{2}
}
func (m *PrimitiveId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrimitiveId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrimitiveId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrimitiveId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveId.Merge(m, src)
}
func (m *PrimitiveId) XXX_Size() int {
	return m.Size()
}
func (m *PrimitiveId) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveId.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveId proto.InternalMessageInfo

func (m *PrimitiveId) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *PrimitiveId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Primitive)(nil), "atomix.primitive.v1.Primitive")
	proto.RegisterType((*PrimitiveMeta)(nil), "atomix.primitive.v1.PrimitiveMeta")
	proto.RegisterType((*PrimitiveId)(nil), "atomix.primitive.v1.PrimitiveId")
}

func init() {
	proto.RegisterFile("atomix/primitive/v1/primitive.proto", fileDescriptor_ce01ee29cb809ef2)
}

var fileDescriptor_ce01ee29cb809ef2 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x2c, 0xc9, 0xcf,
	0xcd, 0xac, 0xd0, 0x2f, 0x28, 0xca, 0xcc, 0xcd, 0x2c, 0xc9, 0x2c, 0x4b, 0xd5, 0x2f, 0x33, 0x44,
	0x70, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84, 0x21, 0x8a, 0xf4, 0x10, 0xe2, 0x65, 0x86,
	0x52, 0x92, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0x25, 0x49, 0xa5, 0x69, 0xfa, 0x89,
	0x79, 0x95, 0x10, 0xf5, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xa6, 0x3e, 0x88, 0x05, 0x11,
	0x55, 0xf2, 0xe5, 0xe2, 0x0c, 0x80, 0x19, 0x20, 0xe4, 0xc0, 0xc5, 0x92, 0x9b, 0x5a, 0x92, 0x28,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xa4, 0x87, 0xc5, 0x06, 0x3d, 0xb8, 0x6a, 0xdf, 0xd4,
	0x92, 0x44, 0x27, 0x8e, 0x13, 0xf7, 0xe4, 0x19, 0x2e, 0xdc, 0x93, 0x67, 0x0c, 0x02, 0xeb, 0x54,
	0x4a, 0xe4, 0xe2, 0x45, 0x51, 0x20, 0x64, 0xc3, 0xc5, 0x94, 0x99, 0x02, 0x35, 0x50, 0x01, 0xbf,
	0x81, 0x9e, 0x29, 0x4e, 0x5c, 0x20, 0xe3, 0x1e, 0xdd, 0x93, 0x67, 0xf2, 0x74, 0x09, 0x62, 0xca,
	0x4c, 0x11, 0x12, 0xe2, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c,
	0x02, 0xb3, 0x95, 0x5c, 0xb9, 0xb8, 0x91, 0xb4, 0x08, 0xc9, 0x70, 0x71, 0xe6, 0x25, 0xe6, 0xa6,
	0x16, 0x17, 0x24, 0x26, 0xa7, 0x82, 0xed, 0xe1, 0x0c, 0x42, 0x08, 0x80, 0x0c, 0x00, 0x71, 0x60,
	0x06, 0x80, 0xd8, 0x56, 0x2c, 0x2f, 0x16, 0xc8, 0x33, 0x3a, 0x49, 0x9c, 0x78, 0x24, 0xc7, 0x78,
	0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7,
	0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x12, 0x1b, 0x38, 0x64, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x47, 0x65, 0x68, 0x74, 0x86, 0x01, 0x00, 0x00,
}

func (this *PrimitiveId) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PrimitiveId)
	if !ok {
		that2, ok := that.(PrimitiveId)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (m *Primitive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Primitive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Primitive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PrimitiveMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPrimitive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PrimitiveMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrimitiveMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrimitiveMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintPrimitive(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPrimitive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PrimitiveId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrimitiveId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrimitiveId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPrimitive(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintPrimitive(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPrimitive(dAtA []byte, offset int, v uint64) int {
	offset -= sovPrimitive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Primitive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PrimitiveMeta.Size()
	n += 1 + l + sovPrimitive(uint64(l))
	return n
}

func (m *PrimitiveMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovPrimitive(uint64(l))
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovPrimitive(uint64(l))
	}
	return n
}

func (m *PrimitiveId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovPrimitive(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPrimitive(uint64(l))
	}
	return n
}

func sovPrimitive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPrimitive(x uint64) (n int) {
	return sovPrimitive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Primitive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrimitive
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
			return fmt.Errorf("proto: Primitive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Primitive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimitiveMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
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
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PrimitiveMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrimitive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrimitive
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
func (m *PrimitiveMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrimitive
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
			return fmt.Errorf("proto: PrimitiveMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrimitiveMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
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
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
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
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrimitive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrimitive
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
func (m *PrimitiveId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPrimitive
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
			return fmt.Errorf("proto: PrimitiveId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrimitiveId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
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
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPrimitive
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
				return ErrInvalidLengthPrimitive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPrimitive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPrimitive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPrimitive
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
func skipPrimitive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPrimitive
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
					return 0, ErrIntOverflowPrimitive
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
					return 0, ErrIntOverflowPrimitive
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
				return 0, ErrInvalidLengthPrimitive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPrimitive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPrimitive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPrimitive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPrimitive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPrimitive = fmt.Errorf("proto: unexpected end of group")
)