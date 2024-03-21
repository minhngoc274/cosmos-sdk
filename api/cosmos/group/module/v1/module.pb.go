// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/group/module/v1/module.proto

package modulev1

import (
	_ "cosmossdk.io/api/amino"
	_ "cosmossdk.io/api/cosmos/app/v1alpha1"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Module is the config object of the group module.
type Module struct {
	// max_execution_period defines the max duration after a proposal's voting period ends that members can send a MsgExec
	// to execute the proposal.
	MaxExecutionPeriod time.Duration `protobuf:"bytes,1,opt,name=max_execution_period,json=maxExecutionPeriod,proto3,stdduration" json:"max_execution_period"`
	// MaxMetadataLen defines the max chars allowed in all
	// messages that allows creating or updating a group
	// with a metadata field
	// Defaults to 255 if not explicitly set.
	MaxMetadataLen uint64 `protobuf:"varint,2,opt,name=max_metadata_len,json=maxMetadataLen,proto3" json:"max_metadata_len,omitempty"`
	// MaxProposalTitleLen defines the max chars allowed
	// in string for the MsgSubmitProposal and Proposal
	// summary field
	// Defaults to 255 if not explicitly set.
	MaxProposalTitleLen uint64 `protobuf:"varint,3,opt,name=max_proposal_title_len,json=maxProposalTitleLen,proto3" json:"max_proposal_title_len,omitempty"`
	// MaxProposalSummaryLen defines the max chars allowed
	// in string for the MsgSubmitProposal and Proposal
	// summary field
	// Defaults to 10200 if not explicitly set.
	MaxProposalSummaryLen uint64 `protobuf:"varint,4,opt,name=max_proposal_summary_len,json=maxProposalSummaryLen,proto3" json:"max_proposal_summary_len,omitempty"`
}

func (m *Module) Reset()         { *m = Module{} }
func (m *Module) String() string { return proto.CompactTextString(m) }
func (*Module) ProtoMessage()    {}
func (*Module) Descriptor() ([]byte, []int) {
	return fileDescriptor_4429688cb3ccbd92, []int{0}
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

func (m *Module) GetMaxExecutionPeriod() time.Duration {
	if m != nil {
		return m.MaxExecutionPeriod
	}
	return 0
}

func (m *Module) GetMaxMetadataLen() uint64 {
	if m != nil {
		return m.MaxMetadataLen
	}
	return 0
}

func (m *Module) GetMaxProposalTitleLen() uint64 {
	if m != nil {
		return m.MaxProposalTitleLen
	}
	return 0
}

func (m *Module) GetMaxProposalSummaryLen() uint64 {
	if m != nil {
		return m.MaxProposalSummaryLen
	}
	return 0
}

func init() {
	proto.RegisterType((*Module)(nil), "cosmos.group.module.v1.Module")
}

func init() {
	proto.RegisterFile("cosmos/group/module/v1/module.proto", fileDescriptor_4429688cb3ccbd92)
}

var fileDescriptor_4429688cb3ccbd92 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x3f, 0x8b, 0xd4, 0x40,
	0x18, 0xc6, 0x77, 0x72, 0xc7, 0x15, 0x73, 0x28, 0x1a, 0xd7, 0x65, 0x6f, 0x91, 0xdc, 0x72, 0x36,
	0x8b, 0xc5, 0x8c, 0xf1, 0x0a, 0x21, 0x76, 0x7b, 0xca, 0x36, 0xb7, 0x10, 0x4e, 0xd9, 0xe2, 0x58,
	0x08, 0x73, 0x9b, 0x31, 0x06, 0x33, 0x79, 0x87, 0xfc, 0x23, 0x7e, 0x0b, 0x2b, 0x11, 0x4b, 0x2b,
	0xb1, 0xb2, 0xb6, 0xb2, 0x3c, 0xac, 0xae, 0x14, 0x0b, 0x95, 0x6c, 0x21, 0x58, 0xf8, 0x19, 0x24,
	0x33, 0x13, 0x70, 0x61, 0x9b, 0xf0, 0xe6, 0x7d, 0x9e, 0xe7, 0x57, 0x3c, 0xef, 0xe0, 0xbb, 0x2b,
	0xc8, 0x05, 0xe4, 0x34, 0xca, 0xa0, 0x94, 0x54, 0x40, 0x58, 0x26, 0x9c, 0x56, 0xae, 0x99, 0x88,
	0xcc, 0xa0, 0x00, 0x7b, 0xa0, 0x4d, 0x44, 0x99, 0x88, 0x91, 0x2a, 0x77, 0x34, 0x36, 0x61, 0x26,
	0x25, 0xad, 0x5c, 0x96, 0xc8, 0x17, 0x6c, 0x33, 0x39, 0xea, 0x47, 0x10, 0x81, 0x1a, 0x69, 0x3b,
	0x99, 0xad, 0x13, 0x01, 0x44, 0x09, 0xa7, 0xea, 0xef, 0xa2, 0x7c, 0x4e, 0xc3, 0x32, 0x63, 0x45,
	0x0c, 0xa9, 0xd1, 0x6f, 0x32, 0x11, 0xa7, 0x40, 0xd5, 0x57, 0xaf, 0x8e, 0xde, 0x59, 0x78, 0x6f,
	0xae, 0xc8, 0xf6, 0x39, 0xee, 0x0b, 0x56, 0x07, 0xbc, 0xe6, 0xab, 0xb2, 0x0d, 0x05, 0x92, 0x67,
	0x31, 0x84, 0x43, 0x34, 0x46, 0x93, 0xfd, 0x07, 0x07, 0x44, 0xc3, 0x49, 0x07, 0x27, 0x8f, 0x0d,
	0x7c, 0x7a, 0xed, 0xf2, 0xc7, 0x61, 0xef, 0xed, 0xcf, 0x43, 0xf4, 0xe1, 0xf7, 0xa7, 0x7b, 0xe8,
	0xcc, 0x16, 0xac, 0x7e, 0xd2, 0x41, 0x7c, 0xc5, 0xb0, 0x27, 0xf8, 0x46, 0xcb, 0x16, 0xbc, 0x60,
	0x21, 0x2b, 0x58, 0x90, 0xf0, 0x74, 0x68, 0x8d, 0xd1, 0x64, 0xf7, 0xec, 0xba, 0x60, 0xf5, 0xdc,
	0xac, 0x4f, 0x79, 0x6a, 0x1f, 0xe3, 0x41, 0xeb, 0x94, 0x19, 0x48, 0xc8, 0x59, 0x12, 0x14, 0x71,
	0x91, 0x70, 0xe5, 0xdf, 0x51, 0xfe, 0x5b, 0x82, 0xd5, 0xbe, 0x11, 0x9f, 0xb5, 0x5a, 0x1b, 0x7a,
	0x88, 0x87, 0x1b, 0xa1, 0xbc, 0x14, 0x82, 0x65, 0xaf, 0x54, 0x6c, 0x57, 0xc5, 0x6e, 0xff, 0x17,
	0x7b, 0xaa, 0xd5, 0x53, 0x9e, 0x7a, 0x77, 0x3e, 0x7f, 0x79, 0xf3, 0x1d, 0x0d, 0x70, 0x5f, 0x37,
	0x9e, 0x87, 0x2f, 0x49, 0x0c, 0xb4, 0xd6, 0x67, 0x9b, 0xfe, 0x45, 0x97, 0x8d, 0x83, 0xae, 0x1a,
	0x07, 0xfd, 0x6a, 0x1c, 0xf4, 0x7a, 0xed, 0xf4, 0xae, 0xd6, 0x4e, 0xef, 0xdb, 0xda, 0xe9, 0xe1,
	0xd1, 0x0a, 0x04, 0xd9, 0x7e, 0xbe, 0xe9, 0xbe, 0x2e, 0xd4, 0x6f, 0x8b, 0xf2, 0xd1, 0xf9, 0xfd,
	0x0d, 0x36, 0x93, 0x31, 0xdd, 0xfe, 0x36, 0x1e, 0xe9, 0xa9, 0x72, 0xdf, 0x5b, 0x3b, 0x27, 0xb3,
	0xf9, 0x47, 0x6b, 0x70, 0xa2, 0xf9, 0x33, 0xc5, 0xd7, 0x50, 0xb2, 0x70, 0xbf, 0x76, 0xc2, 0x52,
	0x09, 0x4b, 0x2d, 0x2c, 0x17, 0x6e, 0x63, 0x1d, 0x6d, 0x17, 0x96, 0x33, 0x7f, 0xda, 0x75, 0xfc,
	0xc7, 0x3a, 0xd0, 0x26, 0xcf, 0x53, 0x2e, 0xcf, 0xd3, 0x36, 0xcf, 0x5b, 0xb8, 0x17, 0x7b, 0xea,
	0xb8, 0xc7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x8f, 0xf4, 0xa0, 0xbe, 0x02, 0x00, 0x00,
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
	if m.MaxProposalSummaryLen != 0 {
		i = encodeVarintModule(dAtA, i, uint64(m.MaxProposalSummaryLen))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxProposalTitleLen != 0 {
		i = encodeVarintModule(dAtA, i, uint64(m.MaxProposalTitleLen))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxMetadataLen != 0 {
		i = encodeVarintModule(dAtA, i, uint64(m.MaxMetadataLen))
		i--
		dAtA[i] = 0x10
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.MaxExecutionPeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaxExecutionPeriod):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintModule(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
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
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaxExecutionPeriod)
	n += 1 + l + sovModule(uint64(l))
	if m.MaxMetadataLen != 0 {
		n += 1 + sovModule(uint64(m.MaxMetadataLen))
	}
	if m.MaxProposalTitleLen != 0 {
		n += 1 + sovModule(uint64(m.MaxProposalTitleLen))
	}
	if m.MaxProposalSummaryLen != 0 {
		n += 1 + sovModule(uint64(m.MaxProposalSummaryLen))
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
				return fmt.Errorf("proto: wrong wireType = %d for field MaxExecutionPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModule
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
				return ErrInvalidLengthModule
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.MaxExecutionPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMetadataLen", wireType)
			}
			m.MaxMetadataLen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxMetadataLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxProposalTitleLen", wireType)
			}
			m.MaxProposalTitleLen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxProposalTitleLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxProposalSummaryLen", wireType)
			}
			m.MaxProposalSummaryLen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxProposalSummaryLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
