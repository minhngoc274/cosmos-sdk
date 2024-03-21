// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/staking/module/v1/module.proto

package modulev1

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

// Module is the config object of the staking module.
type Module struct {
	// hooks_order specifies the order of staking hooks and should be a list
	// of module names which provide a staking hooks instance. If no order is
	// provided, then hooks will be applied in alphabetical order of module names.
	HooksOrder []string `protobuf:"bytes,1,rep,name=hooks_order,json=hooksOrder,proto3" json:"hooks_order,omitempty"`
	// authority defines the custom module authority. If not set, defaults to the governance module.
	Authority string `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	// bech32_prefix_validator is the bech32 validator prefix for the app.
	Bech32PrefixValidator string `protobuf:"bytes,3,opt,name=bech32_prefix_validator,json=bech32PrefixValidator,proto3" json:"bech32_prefix_validator,omitempty"`
	// bech32_prefix_consensus is the bech32 consensus node prefix for the app.
	Bech32PrefixConsensus string `protobuf:"bytes,4,opt,name=bech32_prefix_consensus,json=bech32PrefixConsensus,proto3" json:"bech32_prefix_consensus,omitempty"`
}

func (m *Module) Reset()         { *m = Module{} }
func (m *Module) String() string { return proto.CompactTextString(m) }
func (*Module) ProtoMessage()    {}
func (*Module) Descriptor() ([]byte, []int) {
	return fileDescriptor_28c71533429339ae, []int{0}
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

func (m *Module) GetHooksOrder() []string {
	if m != nil {
		return m.HooksOrder
	}
	return nil
}

func (m *Module) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *Module) GetBech32PrefixValidator() string {
	if m != nil {
		return m.Bech32PrefixValidator
	}
	return ""
}

func (m *Module) GetBech32PrefixConsensus() string {
	if m != nil {
		return m.Bech32PrefixConsensus
	}
	return ""
}

func init() {
	proto.RegisterType((*Module)(nil), "cosmos.staking.module.v1.Module")
}

func init() {
	proto.RegisterFile("cosmos/staking/module/v1/module.proto", fileDescriptor_28c71533429339ae)
}

var fileDescriptor_28c71533429339ae = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x2f, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0xcf, 0xcd, 0x4f, 0x29, 0xcd,
	0x49, 0xd5, 0x2f, 0x33, 0x84, 0xb2, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x24, 0x20, 0xca,
	0xf4, 0xa0, 0xca, 0xf4, 0xa0, 0x92, 0x65, 0x86, 0x52, 0x0a, 0x50, 0x03, 0x12, 0x0b, 0x0a, 0xf4,
	0xcb, 0x0c, 0x13, 0x73, 0x0a, 0x32, 0x12, 0x51, 0xf5, 0x2a, 0x5d, 0x67, 0xe4, 0x62, 0xf3, 0x05,
	0x0b, 0x08, 0xc9, 0x73, 0x71, 0x67, 0xe4, 0xe7, 0x67, 0x17, 0xc7, 0xe7, 0x17, 0xa5, 0xa4, 0x16,
	0x49, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x06, 0x71, 0x81, 0x85, 0xfc, 0x41, 0x22, 0x42, 0x32, 0x5c,
	0x9c, 0x89, 0xa5, 0x25, 0x19, 0xf9, 0x45, 0x99, 0x25, 0x95, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x08, 0x01, 0x21, 0x33, 0x2e, 0xf1, 0xa4, 0xd4, 0xe4, 0x0c, 0x63, 0xa3, 0xf8, 0x82, 0xa2,
	0xd4, 0xb4, 0xcc, 0x8a, 0xf8, 0xb2, 0xc4, 0x9c, 0xcc, 0x94, 0xc4, 0x92, 0xfc, 0x22, 0x09, 0x66,
	0xb0, 0x5a, 0x51, 0x88, 0x74, 0x00, 0x58, 0x36, 0x0c, 0x26, 0x89, 0xa9, 0x2f, 0x39, 0x3f, 0xaf,
	0x38, 0x35, 0xaf, 0xb8, 0xb4, 0x58, 0x82, 0x05, 0x53, 0x9f, 0x33, 0x4c, 0xd2, 0x4a, 0x6e, 0xd7,
	0x81, 0x69, 0xb7, 0x18, 0x25, 0xb8, 0xc4, 0x20, 0x7e, 0x2c, 0x4e, 0xc9, 0xd6, 0xcb, 0xcc, 0xd7,
	0xaf, 0x80, 0x05, 0x96, 0xd3, 0x2f, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63,
	0xe0, 0x92, 0x49, 0xce, 0xcf, 0xd5, 0xc3, 0x15, 0x68, 0x4e, 0xdc, 0x90, 0xf0, 0x08, 0x00, 0x85,
	0x4f, 0x00, 0x63, 0x94, 0x11, 0x8a, 0xf9, 0x89, 0x05, 0x99, 0xfa, 0xb8, 0x62, 0xc5, 0x1a, 0xc2,
	0x2a, 0x33, 0x5c, 0xc4, 0xc4, 0xec, 0x1c, 0xec, 0xbb, 0x8a, 0x49, 0xc2, 0x19, 0x62, 0x47, 0x30,
	0xd4, 0x0e, 0x88, 0xc1, 0x7a, 0x61, 0x86, 0xa7, 0x60, 0x52, 0x31, 0x50, 0xa9, 0x18, 0x88, 0x54,
	0x4c, 0x98, 0xe1, 0x23, 0x26, 0x15, 0x5c, 0x52, 0x31, 0xee, 0x01, 0x4e, 0xbe, 0xa9, 0x25, 0x89,
	0x29, 0x89, 0x25, 0x89, 0xaf, 0x98, 0xa4, 0x21, 0xca, 0xac, 0xac, 0xa0, 0xea, 0xac, 0xac, 0x20,
	0x0a, 0xad, 0xac, 0xc2, 0x0c, 0x93, 0xd8, 0xc0, 0xb1, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xff, 0xfc, 0x31, 0x94, 0x42, 0x02, 0x00, 0x00,
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
	if len(m.Bech32PrefixConsensus) > 0 {
		i -= len(m.Bech32PrefixConsensus)
		copy(dAtA[i:], m.Bech32PrefixConsensus)
		i = encodeVarintModule(dAtA, i, uint64(len(m.Bech32PrefixConsensus)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Bech32PrefixValidator) > 0 {
		i -= len(m.Bech32PrefixValidator)
		copy(dAtA[i:], m.Bech32PrefixValidator)
		i = encodeVarintModule(dAtA, i, uint64(len(m.Bech32PrefixValidator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintModule(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.HooksOrder) > 0 {
		for iNdEx := len(m.HooksOrder) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.HooksOrder[iNdEx])
			copy(dAtA[i:], m.HooksOrder[iNdEx])
			i = encodeVarintModule(dAtA, i, uint64(len(m.HooksOrder[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
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
	if len(m.HooksOrder) > 0 {
		for _, s := range m.HooksOrder {
			l = len(s)
			n += 1 + l + sovModule(uint64(l))
		}
	}
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovModule(uint64(l))
	}
	l = len(m.Bech32PrefixValidator)
	if l > 0 {
		n += 1 + l + sovModule(uint64(l))
	}
	l = len(m.Bech32PrefixConsensus)
	if l > 0 {
		n += 1 + l + sovModule(uint64(l))
	}
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HooksOrder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModule
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
				return ErrInvalidLengthModule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HooksOrder = append(m.HooksOrder, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModule
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
				return ErrInvalidLengthModule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bech32PrefixValidator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModule
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
				return ErrInvalidLengthModule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bech32PrefixValidator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bech32PrefixConsensus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModule
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
				return ErrInvalidLengthModule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bech32PrefixConsensus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
