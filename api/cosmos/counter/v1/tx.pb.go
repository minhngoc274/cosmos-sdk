// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/counter/v1/tx.proto

package counterv1

import (
	_ "cosmossdk.io/api/amino"
	_ "cosmossdk.io/api/cosmos/msg/v1"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgIncreaseCounter defines a count Msg service counter.
type MsgIncreaseCounter struct {
	// signer is the address that controls the module (defaults to x/gov unless overwritten).
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	// count is the number of times to increment the counter.
	Count int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *MsgIncreaseCounter) Reset()         { *m = MsgIncreaseCounter{} }
func (m *MsgIncreaseCounter) String() string { return proto.CompactTextString(m) }
func (*MsgIncreaseCounter) ProtoMessage()    {}
func (*MsgIncreaseCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_16b86a09a73a7eb5, []int{0}
}
func (m *MsgIncreaseCounter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIncreaseCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIncreaseCounter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIncreaseCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIncreaseCounter.Merge(m, src)
}
func (m *MsgIncreaseCounter) XXX_Size() int {
	return m.Size()
}
func (m *MsgIncreaseCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIncreaseCounter.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIncreaseCounter proto.InternalMessageInfo

func (m *MsgIncreaseCounter) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgIncreaseCounter) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// MsgIncreaseCountResponse is the Msg/Counter response type.
type MsgIncreaseCountResponse struct {
	// new_count is the number of times the counter was incremented.
	NewCount int64 `protobuf:"varint,1,opt,name=new_count,json=newCount,proto3" json:"new_count,omitempty"`
}

func (m *MsgIncreaseCountResponse) Reset()         { *m = MsgIncreaseCountResponse{} }
func (m *MsgIncreaseCountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgIncreaseCountResponse) ProtoMessage()    {}
func (*MsgIncreaseCountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_16b86a09a73a7eb5, []int{1}
}
func (m *MsgIncreaseCountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIncreaseCountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIncreaseCountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIncreaseCountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIncreaseCountResponse.Merge(m, src)
}
func (m *MsgIncreaseCountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgIncreaseCountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIncreaseCountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIncreaseCountResponse proto.InternalMessageInfo

func (m *MsgIncreaseCountResponse) GetNewCount() int64 {
	if m != nil {
		return m.NewCount
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgIncreaseCounter)(nil), "cosmos.counter.v1.MsgIncreaseCounter")
	proto.RegisterType((*MsgIncreaseCountResponse)(nil), "cosmos.counter.v1.MsgIncreaseCountResponse")
}

func init() { proto.RegisterFile("cosmos/counter/v1/tx.proto", fileDescriptor_16b86a09a73a7eb5) }

var fileDescriptor_16b86a09a73a7eb5 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0xce, 0x2f, 0xcd, 0x2b, 0x49, 0x2d, 0xd2, 0x2f, 0x33, 0xd4, 0x2f, 0xa9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0xc8, 0xe9, 0x41, 0xe5, 0xf4, 0xca, 0x0c,
	0xa5, 0x24, 0x21, 0x42, 0xf1, 0x60, 0x05, 0xfa, 0x50, 0x79, 0x30, 0x47, 0x4a, 0x1c, 0x6a, 0x52,
	0x6e, 0x71, 0x3a, 0xc8, 0x94, 0xdc, 0xe2, 0x74, 0xa8, 0x84, 0x60, 0x62, 0x6e, 0x66, 0x5e, 0xbe,
	0x3e, 0x98, 0x84, 0x08, 0x29, 0x75, 0x32, 0x72, 0x09, 0xf9, 0x16, 0xa7, 0x7b, 0xe6, 0x25, 0x17,
	0xa5, 0x26, 0x16, 0xa7, 0x3a, 0x43, 0x2c, 0x10, 0x32, 0xe0, 0x62, 0x2b, 0xce, 0x4c, 0xcf, 0x4b,
	0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0x92, 0xb8, 0xb4, 0x45, 0x57, 0x04, 0x6a, 0x89,
	0x63, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0x71, 0x70, 0x49, 0x51, 0x66, 0x5e, 0x7a, 0x10, 0x54, 0x9d,
	0x90, 0x08, 0x17, 0x2b, 0xd8, 0x75, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x10, 0x8e, 0x95,
	0x76, 0xd3, 0xf3, 0x0d, 0x5a, 0x50, 0x25, 0x5d, 0xcf, 0x37, 0x68, 0x49, 0x43, 0xcc, 0xd0, 0x2d,
	0x4e, 0xc9, 0xd6, 0xcf, 0x84, 0xda, 0x19, 0x0f, 0xf5, 0x95, 0x92, 0x39, 0x97, 0x04, 0xba, 0x53,
	0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xa4, 0xb9, 0x38, 0xf3, 0x52, 0xcb, 0x21,
	0x4a, 0xc1, 0x6e, 0x62, 0x0e, 0xe2, 0xc8, 0x4b, 0x2d, 0x07, 0x2b, 0x32, 0x2a, 0xe4, 0x62, 0xf6,
	0x2d, 0x4e, 0x17, 0x4a, 0xe6, 0xe2, 0x45, 0xd1, 0x2c, 0xa4, 0xaa, 0x87, 0x11, 0x6e, 0x7a, 0x98,
	0x9e, 0x95, 0xd2, 0x26, 0x42, 0x19, 0xcc, 0x21, 0x52, 0xac, 0x0d, 0xcf, 0x37, 0x68, 0x31, 0x3a,
	0x9d, 0x61, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27,
	0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x06, 0x2e, 0xd1, 0xe4, 0xfc,
	0x5c, 0x4c, 0x13, 0x9d, 0xd8, 0x43, 0x2a, 0x02, 0x40, 0x41, 0x1e, 0xc0, 0x18, 0xa5, 0x03, 0x91,
	0x2d, 0x4e, 0xc9, 0xd6, 0xcb, 0xcc, 0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0xc7, 0x88, 0x7b, 0x6b, 0x28,
	0xb3, 0xcc, 0x70, 0x11, 0x13, 0xb3, 0xb3, 0x73, 0xc4, 0x2a, 0x26, 0x41, 0x67, 0x88, 0x91, 0x50,
	0x57, 0xeb, 0x85, 0x19, 0x9e, 0x82, 0x89, 0xc5, 0x40, 0xc5, 0x62, 0xc2, 0x0c, 0x1f, 0x31, 0xc9,
	0x62, 0x88, 0xc5, 0xb8, 0x07, 0x38, 0xf9, 0xa6, 0x96, 0x24, 0xa6, 0x24, 0x96, 0x24, 0xbe, 0x62,
	0x12, 0x86, 0xc8, 0x5b, 0x59, 0x41, 0x15, 0x58, 0x59, 0x85, 0x19, 0x26, 0xb1, 0x81, 0x53, 0x83,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x26, 0xda, 0x1f, 0x7c, 0x85, 0x02, 0x00, 0x00,
}

func (m *MsgIncreaseCounter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIncreaseCounter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIncreaseCounter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgIncreaseCountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIncreaseCountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIncreaseCountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NewCount != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.NewCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgIncreaseCounter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovTx(uint64(m.Count))
	}
	return n
}

func (m *MsgIncreaseCountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NewCount != 0 {
		n += 1 + sovTx(uint64(m.NewCount))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgIncreaseCounter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgIncreaseCounter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIncreaseCounter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgIncreaseCountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgIncreaseCountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIncreaseCountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewCount", wireType)
			}
			m.NewCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
