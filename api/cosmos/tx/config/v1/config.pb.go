// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/tx/config/v1/config.proto

package configv1

import (
	_ "cosmossdk.io/api/cosmos/app/v1alpha1"
	fmt "fmt"
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

// Config is the config object of the x/auth/tx package.
type Config struct {
	// skip_ante_handler defines whether the ante handler registration should be skipped in case an app wants to override
	// this functionality.
	SkipAnteHandler bool `protobuf:"varint,1,opt,name=skip_ante_handler,json=skipAnteHandler,proto3" json:"skip_ante_handler,omitempty"`
	// skip_post_handler defines whether the post handler registration should be skipped in case an app wants to override
	// this functionality.
	SkipPostHandler bool `protobuf:"varint,2,opt,name=skip_post_handler,json=skipPostHandler,proto3" json:"skip_post_handler,omitempty"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd08cfbfebda1e04, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Config.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return m.Size()
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetSkipAnteHandler() bool {
	if m != nil {
		return m.SkipAnteHandler
	}
	return false
}

func (m *Config) GetSkipPostHandler() bool {
	if m != nil {
		return m.SkipPostHandler
	}
	return false
}

func init() {
	proto.RegisterType((*Config)(nil), "cosmos.tx.config.v1.Config")
}

func init() { proto.RegisterFile("cosmos/tx/config/v1/config.proto", fileDescriptor_dd08cfbfebda1e04) }

var fileDescriptor_dd08cfbfebda1e04 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x2f, 0xa9, 0xd0, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x33, 0x84,
	0xb2, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84, 0x21, 0x2a, 0xf4, 0x4a, 0x2a, 0xf4, 0xa0,
	0xe2, 0x65, 0x86, 0x52, 0x30, 0x6d, 0x89, 0x05, 0x05, 0xfa, 0x65, 0x86, 0x89, 0x39, 0x05, 0x19,
	0x89, 0x86, 0xfa, 0xb9, 0xf9, 0x29, 0xa5, 0x39, 0xa9, 0x10, 0x6d, 0x4a, 0x0d, 0x8c, 0x5c, 0x6c,
	0xce, 0x60, 0xf5, 0x42, 0x5a, 0x5c, 0x82, 0xc5, 0xd9, 0x99, 0x05, 0xf1, 0x89, 0x79, 0x25, 0xa9,
	0xf1, 0x19, 0x89, 0x79, 0x29, 0x39, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0xfc,
	0x20, 0x09, 0xc7, 0xbc, 0x92, 0x54, 0x0f, 0x88, 0x30, 0x5c, 0x6d, 0x41, 0x7e, 0x71, 0x09, 0x5c,
	0x2d, 0x13, 0x42, 0x6d, 0x40, 0x7e, 0x71, 0x09, 0x54, 0xad, 0x95, 0xdc, 0xae, 0x03, 0xd3, 0x6e,
	0x31, 0x4a, 0x70, 0x89, 0x41, 0x1c, 0x53, 0x9c, 0x92, 0xad, 0x97, 0x99, 0xaf, 0x5f, 0xa1, 0x9f,
	0x58, 0x5a, 0x92, 0xa1, 0x5f, 0x52, 0xe1, 0x74, 0x87, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f,
	0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b,
	0x8f, 0xe5, 0x18, 0xb8, 0xc4, 0x93, 0xf3, 0x73, 0xf5, 0xb0, 0x78, 0xcc, 0x89, 0x1b, 0xe2, 0xe6,
	0x00, 0x90, 0x1f, 0x02, 0x18, 0xa3, 0x74, 0x51, 0x8c, 0x4e, 0x2c, 0xc8, 0xd4, 0xc7, 0x12, 0x5e,
	0xd6, 0x10, 0x56, 0x99, 0xe1, 0x22, 0x26, 0x66, 0xe7, 0x10, 0xe7, 0x55, 0x4c, 0xc2, 0xce, 0x10,
	0x93, 0x43, 0x2a, 0xf4, 0x20, 0xc6, 0xe9, 0x85, 0x19, 0x9e, 0x82, 0x89, 0xc6, 0x84, 0x54, 0xc4,
	0x40, 0x44, 0x63, 0xc2, 0x0c, 0x1f, 0x31, 0xc9, 0x63, 0x11, 0x8d, 0x71, 0x0f, 0x70, 0xf2, 0x4d,
	0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0x7c, 0xc5, 0x24, 0x06, 0x51, 0x61, 0x65, 0x15, 0x52, 0x61,
	0x65, 0x05, 0x51, 0x63, 0x65, 0x15, 0x66, 0x98, 0xc4, 0x06, 0x0e, 0x68, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe1, 0x9c, 0x45, 0x33, 0xc3, 0x01, 0x00, 0x00,
}

func (m *Config) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Config) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SkipPostHandler {
		i--
		if m.SkipPostHandler {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.SkipAnteHandler {
		i--
		if m.SkipAnteHandler {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Config) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SkipAnteHandler {
		n += 2
	}
	if m.SkipPostHandler {
		n += 2
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Config) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SkipAnteHandler", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SkipAnteHandler = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SkipPostHandler", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SkipPostHandler = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)
