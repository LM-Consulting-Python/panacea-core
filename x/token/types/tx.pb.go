// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: token/v2/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
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

// this line is used by starport scaffolding # proto/tx/message
type MsgIssueToken struct {
	Name             string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ShortSymbol      string         `protobuf:"bytes,2,opt,name=shortSymbol,proto3" json:"shortSymbol,omitempty"`
	TotalSupplyMicro types.IntProto `protobuf:"bytes,3,opt,name=totalSupplyMicro,proto3" json:"totalSupplyMicro"`
	Mintable         bool           `protobuf:"varint,4,opt,name=mintable,proto3" json:"mintable,omitempty"`
	OwnerAddress     string         `protobuf:"bytes,5,opt,name=ownerAddress,proto3" json:"ownerAddress,omitempty"`
}

func (m *MsgIssueToken) Reset()         { *m = MsgIssueToken{} }
func (m *MsgIssueToken) String() string { return proto.CompactTextString(m) }
func (*MsgIssueToken) ProtoMessage()    {}
func (*MsgIssueToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0e3bc0da35d4a9a, []int{0}
}
func (m *MsgIssueToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIssueToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIssueToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIssueToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIssueToken.Merge(m, src)
}
func (m *MsgIssueToken) XXX_Size() int {
	return m.Size()
}
func (m *MsgIssueToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIssueToken.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIssueToken proto.InternalMessageInfo

func (m *MsgIssueToken) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgIssueToken) GetShortSymbol() string {
	if m != nil {
		return m.ShortSymbol
	}
	return ""
}

func (m *MsgIssueToken) GetTotalSupplyMicro() types.IntProto {
	if m != nil {
		return m.TotalSupplyMicro
	}
	return types.IntProto{}
}

func (m *MsgIssueToken) GetMintable() bool {
	if m != nil {
		return m.Mintable
	}
	return false
}

func (m *MsgIssueToken) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

type MsgIssueTokenResponse struct {
}

func (m *MsgIssueTokenResponse) Reset()         { *m = MsgIssueTokenResponse{} }
func (m *MsgIssueTokenResponse) String() string { return proto.CompactTextString(m) }
func (*MsgIssueTokenResponse) ProtoMessage()    {}
func (*MsgIssueTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0e3bc0da35d4a9a, []int{1}
}
func (m *MsgIssueTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIssueTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIssueTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIssueTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIssueTokenResponse.Merge(m, src)
}
func (m *MsgIssueTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgIssueTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIssueTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIssueTokenResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgIssueToken)(nil), "panacea.token.v2.MsgIssueToken")
	proto.RegisterType((*MsgIssueTokenResponse)(nil), "panacea.token.v2.MsgIssueTokenResponse")
}

func init() { proto.RegisterFile("token/v2/tx.proto", fileDescriptor_c0e3bc0da35d4a9a) }

var fileDescriptor_c0e3bc0da35d4a9a = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x3b, 0x17, 0xee, 0x0d, 0x77, 0xb8, 0x37, 0xc1, 0x89, 0xc6, 0xa6, 0x89, 0xa5, 0x61,
	0x23, 0x1b, 0x67, 0x42, 0x7d, 0x02, 0x59, 0xc9, 0x82, 0x68, 0x8a, 0x71, 0x61, 0xe2, 0x62, 0x3a,
	0x4c, 0x4a, 0x63, 0x3b, 0xa7, 0xe9, 0x0c, 0x08, 0x6f, 0xe1, 0x63, 0xb1, 0x64, 0x65, 0x5c, 0x19,
	0x03, 0x2f, 0x62, 0xda, 0xa2, 0x01, 0x5d, 0xb8, 0x3b, 0x73, 0xfe, 0x73, 0xfe, 0x7c, 0x67, 0x7e,
	0x7c, 0x60, 0xe0, 0x41, 0x2a, 0x36, 0xf3, 0x99, 0x99, 0xd3, 0x2c, 0x07, 0x03, 0xa4, 0x95, 0x71,
	0xc5, 0x85, 0xe4, 0xb4, 0x94, 0xe8, 0xcc, 0x77, 0x0e, 0x23, 0x88, 0xa0, 0x14, 0x59, 0x51, 0x55,
	0x73, 0x8e, 0x2b, 0x40, 0xa7, 0xa0, 0x59, 0xc8, 0xb5, 0x64, 0xb3, 0x5e, 0x28, 0x0d, 0xef, 0x31,
	0x01, 0xb1, 0xaa, 0xf4, 0xce, 0x33, 0xc2, 0xff, 0x87, 0x3a, 0x1a, 0x68, 0x3d, 0x95, 0x37, 0x85,
	0x15, 0x21, 0xb8, 0xae, 0x78, 0x2a, 0x6d, 0xe4, 0xa1, 0xee, 0xdf, 0xa0, 0xac, 0x89, 0x87, 0x9b,
	0x7a, 0x02, 0xb9, 0x19, 0x2d, 0xd2, 0x10, 0x12, 0xfb, 0x57, 0x29, 0xed, 0xb6, 0xc8, 0x15, 0x6e,
	0x19, 0x30, 0x3c, 0x19, 0x4d, 0xb3, 0x2c, 0x59, 0x0c, 0x63, 0x91, 0x83, 0x5d, 0xf3, 0x50, 0xb7,
	0xe9, 0x9f, 0xd0, 0x0a, 0x81, 0x16, 0x08, 0x74, 0x8b, 0x40, 0x07, 0xca, 0x5c, 0x17, 0x00, 0xfd,
	0xfa, 0xf2, 0xb5, 0x6d, 0x05, 0xdf, 0x96, 0x89, 0x83, 0x1b, 0x69, 0xac, 0x0c, 0x0f, 0x13, 0x69,
	0xd7, 0x3d, 0xd4, 0x6d, 0x04, 0x9f, 0x6f, 0xd2, 0xc1, 0xff, 0xe0, 0x51, 0xc9, 0xfc, 0x62, 0x3c,
	0xce, 0xa5, 0xd6, 0xf6, 0xef, 0x92, 0x67, 0xaf, 0xd7, 0x39, 0xc6, 0x47, 0x7b, 0x77, 0x05, 0x52,
	0x67, 0xa0, 0xb4, 0xf4, 0xef, 0x71, 0x6d, 0xa8, 0x23, 0x72, 0x8b, 0xf1, 0xce, 0xd1, 0x6d, 0xfa,
	0xf5, 0x3f, 0xe9, 0xde, 0xb6, 0x73, 0xfa, 0xc3, 0xc0, 0x87, 0x7d, 0xff, 0x72, 0xb9, 0x76, 0xd1,
	0x6a, 0xed, 0xa2, 0xb7, 0xb5, 0x8b, 0x9e, 0x36, 0xae, 0xb5, 0xda, 0xb8, 0xd6, 0xcb, 0xc6, 0xb5,
	0xee, 0x68, 0x14, 0x9b, 0xc9, 0x34, 0xa4, 0x02, 0x52, 0x96, 0xca, 0x71, 0x1c, 0x26, 0x20, 0xd8,
	0xd6, 0xf5, 0x4c, 0x40, 0x2e, 0xd9, 0x9c, 0x55, 0x41, 0x9b, 0x45, 0x26, 0x75, 0xf8, 0xa7, 0x4c,
	0xe8, 0xfc, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x82, 0x67, 0x73, 0x6f, 0xfe, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	IssueToken(ctx context.Context, in *MsgIssueToken, opts ...grpc.CallOption) (*MsgIssueTokenResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) IssueToken(ctx context.Context, in *MsgIssueToken, opts ...grpc.CallOption) (*MsgIssueTokenResponse, error) {
	out := new(MsgIssueTokenResponse)
	err := c.cc.Invoke(ctx, "/panacea.token.v2.Msg/IssueToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	IssueToken(context.Context, *MsgIssueToken) (*MsgIssueTokenResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) IssueToken(ctx context.Context, req *MsgIssueToken) (*MsgIssueTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueToken not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_IssueToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgIssueToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).IssueToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/panacea.token.v2.Msg/IssueToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).IssueToken(ctx, req.(*MsgIssueToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "panacea.token.v2.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueToken",
			Handler:    _Msg_IssueToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "token/v2/tx.proto",
}

func (m *MsgIssueToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIssueToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIssueToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Mintable {
		i--
		if m.Mintable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.TotalSupplyMicro.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ShortSymbol) > 0 {
		i -= len(m.ShortSymbol)
		copy(dAtA[i:], m.ShortSymbol)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ShortSymbol)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgIssueTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIssueTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIssueTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgIssueToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ShortSymbol)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.TotalSupplyMicro.Size()
	n += 1 + l + sovTx(uint64(l))
	if m.Mintable {
		n += 2
	}
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgIssueTokenResponse) Size() (n int) {
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
func (m *MsgIssueToken) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgIssueToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIssueToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShortSymbol", wireType)
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
			m.ShortSymbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupplyMicro", wireType)
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
			if err := m.TotalSupplyMicro.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mintable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			m.Mintable = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
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
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
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
func (m *MsgIssueTokenResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgIssueTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIssueTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
