// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/app/v1alpha1/query.proto

package appv1alpha1

import (
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

// QueryConfigRequest is the Query/Config request type.
type QueryConfigRequest struct {
}

func (m *QueryConfigRequest) Reset()         { *m = QueryConfigRequest{} }
func (m *QueryConfigRequest) String() string { return proto.CompactTextString(m) }
func (*QueryConfigRequest) ProtoMessage()    {}
func (*QueryConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd84655af543ba53, []int{0}
}
func (m *QueryConfigRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryConfigRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryConfigRequest.Merge(m, src)
}
func (m *QueryConfigRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryConfigRequest proto.InternalMessageInfo

// QueryConfigRequest is the Query/Config response type.
type QueryConfigResponse struct {
	// config is the current app config.
	Config *Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *QueryConfigResponse) Reset()         { *m = QueryConfigResponse{} }
func (m *QueryConfigResponse) String() string { return proto.CompactTextString(m) }
func (*QueryConfigResponse) ProtoMessage()    {}
func (*QueryConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd84655af543ba53, []int{1}
}
func (m *QueryConfigResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryConfigResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryConfigResponse.Merge(m, src)
}
func (m *QueryConfigResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryConfigResponse proto.InternalMessageInfo

func (m *QueryConfigResponse) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryConfigRequest)(nil), "cosmos.app.v1alpha1.QueryConfigRequest")
	proto.RegisterType((*QueryConfigResponse)(nil), "cosmos.app.v1alpha1.QueryConfigResponse")
}

func init() { proto.RegisterFile("cosmos/app/v1alpha1/query.proto", fileDescriptor_bd84655af543ba53) }

var fileDescriptor_bd84655af543ba53 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0x2c, 0x28, 0xd0, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4,
	0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86, 0x28, 0xd0,
	0x4b, 0x2c, 0x28, 0xd0, 0x83, 0x29, 0x90, 0x52, 0xc0, 0xa6, 0x2b, 0x39, 0x3f, 0x2f, 0x2d, 0x33,
	0x1d, 0xa2, 0x4d, 0x49, 0x84, 0x4b, 0x28, 0x10, 0x64, 0x8a, 0x33, 0x58, 0x30, 0x28, 0xb5, 0xb0,
	0x34, 0xb5, 0xb8, 0x44, 0xc9, 0x8b, 0x4b, 0x18, 0x45, 0xb4, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55,
	0xc8, 0x98, 0x8b, 0x0d, 0xa2, 0x59, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x5a, 0x0f, 0x8b,
	0xa5, 0x7a, 0x50, 0x4d, 0x50, 0xa5, 0x46, 0x99, 0x5c, 0xac, 0x60, 0xb3, 0x84, 0x12, 0xb8, 0xd8,
	0x20, 0x52, 0x42, 0xea, 0x58, 0xf5, 0x61, 0xba, 0x43, 0x4a, 0x83, 0xb0, 0x42, 0x88, 0xd3, 0x94,
	0x98, 0x3b, 0x98, 0x18, 0x9d, 0xee, 0x32, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3,
	0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c,
	0x03, 0x97, 0x78, 0x72, 0x7e, 0x2e, 0x36, 0xc3, 0x9c, 0xb8, 0xc0, 0xa6, 0x05, 0x80, 0x02, 0x23,
	0x80, 0x31, 0xca, 0x00, 0xa2, 0xa4, 0x38, 0x25, 0x5b, 0x2f, 0x33, 0x5f, 0x3f, 0xb1, 0x20, 0x53,
	0x1f, 0x4b, 0x08, 0x5a, 0x27, 0x16, 0x14, 0xc0, 0xd8, 0x8b, 0x98, 0x98, 0x9d, 0x1d, 0x23, 0x56,
	0x31, 0x09, 0x3b, 0x43, 0xcc, 0x76, 0x2c, 0x28, 0xd0, 0x0b, 0x83, 0xca, 0x9d, 0x82, 0x89, 0xc6,
	0x38, 0x16, 0x14, 0xc4, 0xc0, 0x44, 0x1f, 0x31, 0xc9, 0x63, 0x11, 0x8d, 0x71, 0x0f, 0x70, 0xf2,
	0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0x7c, 0xc5, 0x24, 0x0a, 0x51, 0x61, 0x65, 0xe5, 0x58,
	0x50, 0x60, 0x65, 0x05, 0x53, 0x93, 0xc4, 0x06, 0x8e, 0x33, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc5, 0xa0, 0x81, 0x20, 0x0d, 0x02, 0x00, 0x00,
}

func (m *QueryConfigRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryConfigRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryConfigRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryConfigResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryConfigResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryConfigResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryConfigRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryConfigResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryConfigRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryConfigRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryConfigRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryConfigResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryConfigResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryConfigResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &Config{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
