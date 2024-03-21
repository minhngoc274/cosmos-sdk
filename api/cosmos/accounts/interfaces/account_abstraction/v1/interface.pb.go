// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accounts/interfaces/account_abstraction/v1/interface.proto

package account_abstractionv1

import (
	v1beta1 "cosmossdk.io/api/cosmos/tx/v1beta1"
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

// MsgAuthenticate is a message that an x/account account abstraction implementer
// must handle to authenticate a transaction. Always ensure the caller is the Accounts module.
type MsgAuthenticate struct {
	// bundler defines the address of the bundler that sent the operation.
	// NOTE: in case the operation was sent directly by the user, this field will reflect
	// the user address.
	Bundler string `protobuf:"bytes,1,opt,name=bundler,proto3" json:"bundler,omitempty"`
	// raw_tx defines the raw version of the tx, this is useful to compute the signature quickly.
	RawTx *v1beta1.TxRaw `protobuf:"bytes,2,opt,name=raw_tx,json=rawTx,proto3" json:"raw_tx,omitempty"`
	// tx defines the decoded version of the tx, coming from raw_tx.
	Tx *v1beta1.Tx `protobuf:"bytes,3,opt,name=tx,proto3" json:"tx,omitempty"`
	// signer_index defines the index of the signer in the tx.
	// Specifically this can be used to extract the signature at the correct
	// index.
	SignerIndex uint32 `protobuf:"varint,4,opt,name=signer_index,json=signerIndex,proto3" json:"signer_index,omitempty"`
}

func (m *MsgAuthenticate) Reset()         { *m = MsgAuthenticate{} }
func (m *MsgAuthenticate) String() string { return proto.CompactTextString(m) }
func (*MsgAuthenticate) ProtoMessage()    {}
func (*MsgAuthenticate) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{0}
}
func (m *MsgAuthenticate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAuthenticate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAuthenticate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAuthenticate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAuthenticate.Merge(m, src)
}
func (m *MsgAuthenticate) XXX_Size() int {
	return m.Size()
}
func (m *MsgAuthenticate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAuthenticate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAuthenticate proto.InternalMessageInfo

func (m *MsgAuthenticate) GetBundler() string {
	if m != nil {
		return m.Bundler
	}
	return ""
}

func (m *MsgAuthenticate) GetRawTx() *v1beta1.TxRaw {
	if m != nil {
		return m.RawTx
	}
	return nil
}

func (m *MsgAuthenticate) GetTx() *v1beta1.Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *MsgAuthenticate) GetSignerIndex() uint32 {
	if m != nil {
		return m.SignerIndex
	}
	return 0
}

// MsgAuthenticateResponse is the response to MsgAuthenticate.
// The authentication either fails or succeeds, this is why
// there are no auxiliary fields to the response.
type MsgAuthenticateResponse struct {
}

func (m *MsgAuthenticateResponse) Reset()         { *m = MsgAuthenticateResponse{} }
func (m *MsgAuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAuthenticateResponse) ProtoMessage()    {}
func (*MsgAuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{1}
}
func (m *MsgAuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAuthenticateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAuthenticateResponse.Merge(m, src)
}
func (m *MsgAuthenticateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAuthenticateResponse proto.InternalMessageInfo

// QueryAuthenticationMethods is a query that an x/account account abstraction implementer
// must handle to return the authentication methods that the account supports.
type QueryAuthenticationMethods struct {
}

func (m *QueryAuthenticationMethods) Reset()         { *m = QueryAuthenticationMethods{} }
func (m *QueryAuthenticationMethods) String() string { return proto.CompactTextString(m) }
func (*QueryAuthenticationMethods) ProtoMessage()    {}
func (*QueryAuthenticationMethods) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{2}
}
func (m *QueryAuthenticationMethods) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuthenticationMethods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuthenticationMethods.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuthenticationMethods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuthenticationMethods.Merge(m, src)
}
func (m *QueryAuthenticationMethods) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuthenticationMethods) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuthenticationMethods.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuthenticationMethods proto.InternalMessageInfo

// QueryAuthenticationMethodsResponse is the response to QueryAuthenticationMethods.
type QueryAuthenticationMethodsResponse struct {
	// authentication_methods are the authentication methods that the account supports.
	AuthenticationMethods []string `protobuf:"bytes,1,rep,name=authentication_methods,json=authenticationMethods,proto3" json:"authentication_methods,omitempty"`
}

func (m *QueryAuthenticationMethodsResponse) Reset()         { *m = QueryAuthenticationMethodsResponse{} }
func (m *QueryAuthenticationMethodsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAuthenticationMethodsResponse) ProtoMessage()    {}
func (*QueryAuthenticationMethodsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56b360422260e9d1, []int{3}
}
func (m *QueryAuthenticationMethodsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuthenticationMethodsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuthenticationMethodsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuthenticationMethodsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuthenticationMethodsResponse.Merge(m, src)
}
func (m *QueryAuthenticationMethodsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuthenticationMethodsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuthenticationMethodsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuthenticationMethodsResponse proto.InternalMessageInfo

func (m *QueryAuthenticationMethodsResponse) GetAuthenticationMethods() []string {
	if m != nil {
		return m.AuthenticationMethods
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgAuthenticate)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.MsgAuthenticate")
	proto.RegisterType((*MsgAuthenticateResponse)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.MsgAuthenticateResponse")
	proto.RegisterType((*QueryAuthenticationMethods)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.QueryAuthenticationMethods")
	proto.RegisterType((*QueryAuthenticationMethodsResponse)(nil), "cosmos.accounts.interfaces.account_abstraction.v1.QueryAuthenticationMethodsResponse")
}

func init() {
	proto.RegisterFile("cosmos/accounts/interfaces/account_abstraction/v1/interface.proto", fileDescriptor_56b360422260e9d1)
}

var fileDescriptor_56b360422260e9d1 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x8b, 0xd3, 0x40,
	0x1c, 0x86, 0x3b, 0xe9, 0xba, 0xb2, 0xb3, 0xfe, 0x81, 0xc0, 0x6a, 0x2c, 0x12, 0x6a, 0x40, 0xe8,
	0x69, 0xe2, 0xa8, 0x7b, 0x89, 0x5e, 0xd2, 0x3d, 0x48, 0x0f, 0x85, 0x1a, 0x96, 0x22, 0x1a, 0x28,
	0xd3, 0x64, 0xdc, 0x0d, 0xda, 0x99, 0x32, 0xf3, 0x6b, 0x3b, 0x7e, 0x0b, 0x3f, 0x81, 0x07, 0x0f,
	0x1e, 0xfc, 0x24, 0xe2, 0x69, 0x8f, 0x1e, 0xa5, 0xbd, 0xf9, 0x29, 0x64, 0x37, 0x9d, 0xa6, 0x6a,
	0x3c, 0xf4, 0x98, 0x77, 0xde, 0xe7, 0xc9, 0x0b, 0x33, 0x38, 0xce, 0xa4, 0x9e, 0x48, 0x1d, 0xb2,
	0x2c, 0x93, 0x33, 0x01, 0x3a, 0x2c, 0x04, 0x70, 0xf5, 0x96, 0x65, 0x7c, 0x93, 0x8d, 0xd8, 0x58,
	0x83, 0x62, 0x19, 0x14, 0x52, 0x84, 0x73, 0x5a, 0x35, 0xc8, 0x54, 0x49, 0x90, 0x2e, 0x2d, 0x15,
	0xc4, 0x2a, 0x48, 0xa5, 0x20, 0x35, 0x0a, 0x32, 0xa7, 0xad, 0xd6, 0xfa, 0xaf, 0x60, 0xc2, 0x39,
	0x1d, 0x73, 0x60, 0x34, 0x04, 0x53, 0xea, 0x82, 0x2f, 0x08, 0xdf, 0xee, 0xeb, 0xb3, 0x78, 0x06,
	0xe7, 0x5c, 0x40, 0x91, 0x31, 0xe0, 0xae, 0x87, 0xaf, 0x8f, 0x67, 0x22, 0x7f, 0xcf, 0x95, 0x87,
	0xda, 0xa8, 0x73, 0x90, 0xd8, 0x4f, 0x37, 0xc4, 0xfb, 0x8a, 0x2d, 0x46, 0x60, 0x3c, 0xa7, 0x8d,
	0x3a, 0x87, 0x8f, 0x3d, 0xb2, 0x5e, 0x03, 0x86, 0xac, 0xd5, 0xe4, 0xd4, 0x24, 0x6c, 0x91, 0x5c,
	0x53, 0x6c, 0x71, 0x6a, 0xdc, 0x87, 0xd8, 0x01, 0xe3, 0x35, 0xaf, 0xca, 0x47, 0xf5, 0x65, 0x07,
	0x8c, 0xfb, 0x00, 0xdf, 0xd0, 0xc5, 0x99, 0xe0, 0x6a, 0x54, 0x88, 0x9c, 0x1b, 0x6f, 0xaf, 0x8d,
	0x3a, 0x37, 0x93, 0xc3, 0x32, 0xeb, 0x5d, 0x46, 0xc1, 0x3d, 0x7c, 0xf7, 0xaf, 0x9d, 0x09, 0xd7,
	0x53, 0x29, 0x34, 0x0f, 0xee, 0xe3, 0xd6, 0xcb, 0x19, 0x57, 0x1f, 0xb6, 0x0e, 0x0b, 0x29, 0xfa,
	0x1c, 0xce, 0x65, 0xae, 0x83, 0x37, 0x38, 0xf8, 0xff, 0xa9, 0x75, 0xb8, 0xc7, 0xf8, 0x0e, 0xfb,
	0xa3, 0x30, 0x9a, 0x94, 0x0d, 0x0f, 0xb5, 0x9b, 0x9d, 0x83, 0xe4, 0x88, 0xd5, 0xe1, 0xdd, 0x4f,
	0xcd, 0x6f, 0x4b, 0x1f, 0x5d, 0x2c, 0x7d, 0xf4, 0x73, 0xe9, 0xa3, 0x8f, 0x2b, 0xbf, 0x71, 0xb1,
	0xf2, 0x1b, 0x3f, 0x56, 0x7e, 0x03, 0x1f, 0x67, 0x72, 0x42, 0x76, 0xbe, 0xac, 0xee, 0xad, 0x9e,
	0xad, 0x0c, 0x2e, 0x2f, 0x68, 0x80, 0x5e, 0xbf, 0x2a, 0x25, 0x3a, 0x7f, 0x47, 0x0a, 0x19, 0xb2,
	0x69, 0x11, 0xee, 0xfc, 0x8a, 0x9e, 0xd5, 0xc4, 0x73, 0xfa, 0xd9, 0xd9, 0x3b, 0x89, 0x7b, 0xf1,
	0x57, 0xe7, 0xd1, 0x49, 0x39, 0x33, 0xb6, 0x33, 0x7b, 0xd5, 0xcc, 0x75, 0x16, 0x6f, 0xad, 0x1c,
	0xd2, 0xef, 0x16, 0x49, 0x2d, 0x92, 0x56, 0x48, 0xfa, 0x2f, 0x92, 0x0e, 0xe9, 0xd2, 0x79, 0xbe,
	0x2b, 0x92, 0xbe, 0x18, 0x74, 0xfb, 0x1c, 0x58, 0xce, 0x80, 0xfd, 0x72, 0x9e, 0x96, 0x78, 0x14,
	0x59, 0x3e, 0x8a, 0x2a, 0xc1, 0x26, 0xdd, 0x32, 0x44, 0xd1, 0x90, 0x8e, 0xf7, 0xaf, 0x9e, 0xf9,
	0x93, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x9f, 0xf0, 0x5c, 0x7a, 0x03, 0x00, 0x00,
}

func (m *MsgAuthenticate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAuthenticate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAuthenticate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SignerIndex != 0 {
		i = encodeVarintInterface(dAtA, i, uint64(m.SignerIndex))
		i--
		dAtA[i] = 0x20
	}
	if m.Tx != nil {
		{
			size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInterface(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.RawTx != nil {
		{
			size, err := m.RawTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInterface(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Bundler) > 0 {
		i -= len(m.Bundler)
		copy(dAtA[i:], m.Bundler)
		i = encodeVarintInterface(dAtA, i, uint64(len(m.Bundler)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAuthenticateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAuthenticateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAuthenticateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAuthenticationMethods) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuthenticationMethods) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuthenticationMethods) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAuthenticationMethodsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuthenticationMethodsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuthenticationMethodsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AuthenticationMethods) > 0 {
		for iNdEx := len(m.AuthenticationMethods) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AuthenticationMethods[iNdEx])
			copy(dAtA[i:], m.AuthenticationMethods[iNdEx])
			i = encodeVarintInterface(dAtA, i, uint64(len(m.AuthenticationMethods[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintInterface(dAtA []byte, offset int, v uint64) int {
	offset -= sovInterface(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAuthenticate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bundler)
	if l > 0 {
		n += 1 + l + sovInterface(uint64(l))
	}
	if m.RawTx != nil {
		l = m.RawTx.Size()
		n += 1 + l + sovInterface(uint64(l))
	}
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovInterface(uint64(l))
	}
	if m.SignerIndex != 0 {
		n += 1 + sovInterface(uint64(m.SignerIndex))
	}
	return n
}

func (m *MsgAuthenticateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAuthenticationMethods) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAuthenticationMethodsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AuthenticationMethods) > 0 {
		for _, s := range m.AuthenticationMethods {
			l = len(s)
			n += 1 + l + sovInterface(uint64(l))
		}
	}
	return n
}

func sovInterface(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInterface(x uint64) (n int) {
	return sovInterface(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAuthenticate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: MsgAuthenticate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAuthenticate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bundler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bundler = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RawTx == nil {
				m.RawTx = &v1beta1.TxRaw{}
			}
			if err := m.RawTx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tx == nil {
				m.Tx = &v1beta1.Tx{}
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignerIndex", wireType)
			}
			m.SignerIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignerIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func (m *MsgAuthenticateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: MsgAuthenticateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAuthenticateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func (m *QueryAuthenticationMethods) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: QueryAuthenticationMethods: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuthenticationMethods: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func (m *QueryAuthenticationMethodsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterface
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
			return fmt.Errorf("proto: QueryAuthenticationMethodsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuthenticationMethodsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthenticationMethods", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterface
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
				return ErrInvalidLengthInterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterface
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthenticationMethods = append(m.AuthenticationMethods, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterface
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
func skipInterface(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInterface
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
					return 0, ErrIntOverflowInterface
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
					return 0, ErrIntOverflowInterface
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
				return 0, ErrInvalidLengthInterface
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInterface
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInterface
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInterface        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInterface          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInterface = fmt.Errorf("proto: unexpected end of group")
)
