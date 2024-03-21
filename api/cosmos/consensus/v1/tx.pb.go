// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/consensus/v1/tx.proto

package consensusv1

import (
	types "buf.build/gen/go/tendermint/tendermint/protocolbuffers/go/tendermint/types"
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

// MsgUpdateParams is the Msg/UpdateParams request type.
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params defines the x/consensus parameters to update.
	// VersionsParams is not included in this Msg because it is tracked
	// separarately in x/upgrade.
	//
	// NOTE: All parameters must be supplied.
	Block     *types.BlockParams     `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	Evidence  *types.EvidenceParams  `protobuf:"bytes,3,opt,name=evidence,proto3" json:"evidence,omitempty"`
	Validator *types.ValidatorParams `protobuf:"bytes,4,opt,name=validator,proto3" json:"validator,omitempty"`
	// Since: cosmos-sdk 0.50
	Abci *types.ABCIParams `protobuf:"bytes,5,opt,name=abci,proto3" json:"abci,omitempty"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2135c60575ab504d, []int{0}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetBlock() *types.BlockParams {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *MsgUpdateParams) GetEvidence() *types.EvidenceParams {
	if m != nil {
		return m.Evidence
	}
	return nil
}

func (m *MsgUpdateParams) GetValidator() *types.ValidatorParams {
	if m != nil {
		return m.Validator
	}
	return nil
}

func (m *MsgUpdateParams) GetAbci() *types.ABCIParams {
	if m != nil {
		return m.Abci
	}
	return nil
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2135c60575ab504d, []int{1}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateParams)(nil), "cosmos.consensus.v1.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "cosmos.consensus.v1.MsgUpdateParamsResponse")
}

func init() { proto.RegisterFile("cosmos/consensus/v1/tx.proto", fileDescriptor_2135c60575ab504d) }

var fileDescriptor_2135c60575ab504d = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x9b, 0x74, 0xab, 0x76, 0x14, 0xc4, 0xac, 0xd2, 0x6c, 0xd9, 0x8d, 0x75, 0x11, 0x59,
	0x8a, 0x9b, 0xd9, 0xec, 0x82, 0x60, 0x14, 0x64, 0x13, 0x44, 0x3c, 0x14, 0x4a, 0xd5, 0x22, 0x52,
	0x90, 0x69, 0x32, 0xc4, 0xa1, 0x9b, 0x4c, 0xc8, 0xcc, 0x86, 0xee, 0x4d, 0x3c, 0x7a, 0xf2, 0x33,
	0x78, 0xf4, 0xd4, 0x83, 0x1f, 0x42, 0x3c, 0x2d, 0x9e, 0xc4, 0x93, 0xb4, 0x87, 0xc2, 0x7e, 0x0a,
	0xe9, 0xcc, 0x34, 0xd5, 0xb6, 0x07, 0x2f, 0x81, 0x79, 0xff, 0xdf, 0xef, 0x3d, 0x32, 0xf3, 0xc0,
	0x76, 0x40, 0x59, 0x4c, 0x19, 0x0c, 0x68, 0xc2, 0x70, 0xc2, 0x4e, 0x19, 0xcc, 0x1d, 0xc8, 0x87,
	0x76, 0x9a, 0x51, 0x4e, 0x8d, 0x4d, 0x99, 0xda, 0x45, 0x6a, 0xe7, 0x4e, 0xfd, 0x06, 0x8a, 0x49,
	0x42, 0xa1, 0xf8, 0x4a, 0xae, 0xbe, 0x25, 0xb9, 0xb7, 0xe2, 0x04, 0x95, 0x24, 0xa3, 0x9a, 0x1a,
	0x10, 0xb3, 0x68, 0xd6, 0x3a, 0x66, 0x91, 0x0a, 0x76, 0x38, 0x4e, 0x42, 0x9c, 0xc5, 0x24, 0xe1,
	0x90, 0x9f, 0xa5, 0x98, 0xc1, 0x14, 0x65, 0x28, 0x56, 0xde, 0xee, 0x85, 0x0e, 0xae, 0xb7, 0x58,
	0xf4, 0x2a, 0x0d, 0x11, 0xc7, 0x6d, 0x91, 0x18, 0x0f, 0x40, 0x15, 0x9d, 0xf2, 0x77, 0x34, 0x23,
	0xfc, 0xcc, 0xd4, 0x1a, 0xda, 0x5e, 0xd5, 0x33, 0x7f, 0x7c, 0xdd, 0xbf, 0xa9, 0x06, 0x1e, 0x87,
	0x61, 0x86, 0x19, 0x7b, 0xc1, 0x33, 0x92, 0x44, 0x9d, 0x05, 0x6a, 0x1c, 0x81, 0x4a, 0xff, 0x84,
	0x06, 0x03, 0x53, 0x6f, 0x68, 0x7b, 0x57, 0x0f, 0x77, 0xec, 0xc5, 0x68, 0x5b, 0x8c, 0xb6, 0xbd,
	0x59, 0x2c, 0xa7, 0x74, 0x24, 0x6b, 0x3c, 0x06, 0x57, 0x70, 0x4e, 0x42, 0x9c, 0x04, 0xd8, 0x2c,
	0x0b, 0xaf, 0xb1, 0xea, 0x3d, 0x55, 0x84, 0x52, 0x0b, 0xc3, 0x78, 0x02, 0xaa, 0x39, 0x3a, 0x21,
	0x21, 0xe2, 0x34, 0x33, 0x37, 0x84, 0x7e, 0x67, 0x55, 0xef, 0xce, 0x11, 0xe5, 0x2f, 0x1c, 0xe3,
	0x00, 0x6c, 0xa0, 0x7e, 0x40, 0xcc, 0x8a, 0x70, 0xb7, 0x57, 0xdd, 0x63, 0xcf, 0x7f, 0xae, 0x34,
	0x41, 0xba, 0x0f, 0x3f, 0x4c, 0x47, 0xcd, 0xc5, 0x5f, 0x7f, 0x9c, 0x8e, 0x9a, 0xf7, 0xe4, 0xcd,
	0xec, 0xb3, 0x70, 0x00, 0x87, 0x7f, 0xbd, 0xf1, 0xd2, 0xc5, 0xee, 0x6e, 0x81, 0xda, 0x52, 0xa9,
	0x83, 0x59, 0x3a, 0xc3, 0x0f, 0x53, 0x50, 0x6e, 0xb1, 0xc8, 0xe8, 0x83, 0x6b, 0xff, 0x3c, 0xc5,
	0x5d, 0x7b, 0xcd, 0x6a, 0xd8, 0x4b, 0x4d, 0xea, 0xf7, 0xff, 0x87, 0x9a, 0x8f, 0xaa, 0x57, 0xde,
	0x4f, 0x47, 0x4d, 0xcd, 0xfb, 0xa5, 0x7d, 0x1b, 0x5b, 0xda, 0xf9, 0xd8, 0xd2, 0x7e, 0x8f, 0x2d,
	0xed, 0xd3, 0xc4, 0x2a, 0x9d, 0x4f, 0xac, 0xd2, 0xcf, 0x89, 0x55, 0x02, 0xb5, 0x80, 0xc6, 0xeb,
	0x5a, 0x7a, 0x97, 0x5f, 0x0e, 0xdb, 0xb3, 0xb5, 0x69, 0x6b, 0x6f, 0x0e, 0x64, 0xce, 0xc2, 0x81,
	0x4d, 0x28, 0x44, 0x29, 0x81, 0x6b, 0x56, 0xfc, 0x51, 0x71, 0xc8, 0x9d, 0xcf, 0x7a, 0xd9, 0xf7,
	0x5f, 0x7f, 0xd1, 0x37, 0x7d, 0xd9, 0xd8, 0x2f, 0x1a, 0x77, 0x9d, 0xef, 0xf3, 0x6a, 0xaf, 0xa8,
	0xf6, 0xba, 0xce, 0x58, 0xbf, 0xbd, 0xa6, 0xda, 0x7b, 0xd6, 0xf6, 0x5a, 0x98, 0xa3, 0x10, 0x71,
	0x74, 0xa1, 0xdf, 0x92, 0x84, 0xeb, 0x16, 0x88, 0xeb, 0x76, 0x9d, 0xfe, 0x25, 0xb1, 0xdd, 0x47,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x12, 0x6e, 0x4e, 0x46, 0x78, 0x03, 0x00, 0x00,
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Abci != nil {
		{
			size, err := m.Abci.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Validator != nil {
		{
			size, err := m.Validator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Evidence != nil {
		{
			size, err := m.Evidence.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Evidence != nil {
		l = m.Evidence.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Validator != nil {
		l = m.Validator.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Abci != nil {
		l = m.Abci.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Block == nil {
				m.Block = &types.BlockParams{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Evidence == nil {
				m.Evidence = &types.EvidenceParams{}
			}
			if err := m.Evidence.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Validator == nil {
				m.Validator = &types.ValidatorParams{}
			}
			if err := m.Validator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abci", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Abci == nil {
				m.Abci = &types.ABCIParams{}
			}
			if err := m.Abci.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
