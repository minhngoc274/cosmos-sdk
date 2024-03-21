// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/orm/module/v1alpha1/module.proto

package modulev1alpha1

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

// Module defines the ORM module which adds providers to the app container for
// ORM ModuleDB's and in the future will automatically register query
// services for modules that use the ORM.
type Module struct {
}

func (m *Module) Reset()         { *m = Module{} }
func (m *Module) String() string { return proto.CompactTextString(m) }
func (*Module) ProtoMessage()    {}
func (*Module) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd217c5a884c508, []int{0}
}
func (m *Module) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Module) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Module.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Module) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Module.Merge(m, src)
}
func (m *Module) XXX_Size() int {
	return m.Size()
}
func (m *Module) XXX_DiscardUnknown() {
	xxx_messageInfo_Module.DiscardUnknown(m)
}

var xxx_messageInfo_Module proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Module)(nil), "cosmos.orm.module.v1alpha1.Module")
}

func init() {
	proto.RegisterFile("cosmos/orm/module/v1alpha1/module.proto", fileDescriptor_3cd217c5a884c508)
}

var fileDescriptor_3cd217c5a884c508 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0xcf, 0x2f, 0xca, 0xd5, 0xcf, 0xcd, 0x4f, 0x29, 0xcd, 0x49, 0xd5, 0x2f, 0x33,
	0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0x84, 0xf2, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xa4,
	0x20, 0x0a, 0xf5, 0xf2, 0x8b, 0x72, 0xf5, 0xa0, 0x12, 0x30, 0x85, 0x52, 0x0a, 0x50, 0x43, 0x12,
	0x0b, 0x0a, 0xb0, 0xeb, 0x56, 0x52, 0xe2, 0x62, 0xf3, 0x05, 0xf3, 0xad, 0x24, 0x76, 0x1d, 0x98,
	0x76, 0x8b, 0x51, 0x88, 0x4b, 0x00, 0xa2, 0xa7, 0x38, 0x25, 0x5b, 0x2f, 0x33, 0x1f, 0x64, 0xbd,
	0x53, 0x0f, 0xd3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38,
	0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x70, 0xc9, 0x25, 0xe7,
	0xe7, 0xea, 0xe1, 0x76, 0x80, 0x13, 0x37, 0xc4, 0xf0, 0x00, 0x90, 0x5d, 0x01, 0x8c, 0x51, 0x56,
	0x28, 0x66, 0x27, 0x16, 0x64, 0xea, 0xe3, 0xf6, 0xa5, 0x35, 0x84, 0x0f, 0xe3, 0x2e, 0x62, 0x62,
	0x76, 0xf6, 0xf7, 0x5d, 0xc5, 0x24, 0xe5, 0x0c, 0xb1, 0xcd, 0xbf, 0x28, 0x57, 0x0f, 0x62, 0xb8,
	0x5e, 0x18, 0x54, 0xc9, 0x29, 0x98, 0x64, 0x8c, 0x7f, 0x51, 0x6e, 0x0c, 0x44, 0x32, 0x06, 0x26,
	0xf9, 0x88, 0x49, 0x0d, 0xb7, 0x64, 0x8c, 0x7b, 0x80, 0x93, 0x6f, 0x6a, 0x49, 0x62, 0x4a, 0x62,
	0x49, 0xe2, 0x2b, 0x26, 0x59, 0x88, 0x42, 0x2b, 0x2b, 0xff, 0xa2, 0x5c, 0x2b, 0x2b, 0x68, 0xf0,
	0x58, 0xc1, 0xd4, 0x26, 0xb1, 0x81, 0x43, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xff, 0x37,
	0xd9, 0x3e, 0xa2, 0x01, 0x00, 0x00,
}

func (m *Module) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Module) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Module) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintModule(dAtA []byte, offset int, v uint64) int {
	offset -= sovModule(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Module) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovModule(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModule(x uint64) (n int) {
	return sovModule(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Module) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModule
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
			return fmt.Errorf("proto: Module: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Module: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipModule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModule
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
func skipModule(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModule
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
					return 0, ErrIntOverflowModule
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
					return 0, ErrIntOverflowModule
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
				return 0, ErrInvalidLengthModule
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModule
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModule
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModule        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModule          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModule = fmt.Errorf("proto: unexpected end of group")
)
