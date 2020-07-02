// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: arith.proto

package service

import (
	context "context"
	fmt "fmt"
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

type ArithRequest struct {
	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (m *ArithRequest) Reset()         { *m = ArithRequest{} }
func (m *ArithRequest) String() string { return proto.CompactTextString(m) }
func (*ArithRequest) ProtoMessage()    {}
func (*ArithRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_211289c5d02710c5, []int{0}
}
func (m *ArithRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ArithRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ArithRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ArithRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArithRequest.Merge(m, src)
}
func (m *ArithRequest) XXX_Size() int {
	return m.Size()
}
func (m *ArithRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArithRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArithRequest proto.InternalMessageInfo

func (m *ArithRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *ArithRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type ArithResponse struct {
	Pro int32 `protobuf:"varint,1,opt,name=pro,proto3" json:"pro,omitempty"`
}

func (m *ArithResponse) Reset()         { *m = ArithResponse{} }
func (m *ArithResponse) String() string { return proto.CompactTextString(m) }
func (*ArithResponse) ProtoMessage()    {}
func (*ArithResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_211289c5d02710c5, []int{1}
}
func (m *ArithResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ArithResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ArithResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ArithResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArithResponse.Merge(m, src)
}
func (m *ArithResponse) XXX_Size() int {
	return m.Size()
}
func (m *ArithResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArithResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArithResponse proto.InternalMessageInfo

func (m *ArithResponse) GetPro() int32 {
	if m != nil {
		return m.Pro
	}
	return 0
}

func init() {
	proto.RegisterType((*ArithRequest)(nil), "ArithRequest")
	proto.RegisterType((*ArithResponse)(nil), "ArithResponse")
}

func init() { proto.RegisterFile("arith.proto", fileDescriptor_211289c5d02710c5) }

var fileDescriptor_211289c5d02710c5 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2c, 0xca, 0x2c,
	0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe2, 0xe2, 0x71, 0x04, 0x71, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x78, 0xb8, 0x18, 0x13, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58,
	0x83, 0x18, 0x13, 0x41, 0xbc, 0x24, 0x09, 0x26, 0x08, 0x2f, 0x49, 0x49, 0x91, 0x8b, 0x17, 0xaa,
	0xb6, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x80, 0x8b, 0xb9, 0xa0, 0x28, 0x1f, 0xaa, 0x1c,
	0xc4, 0x34, 0xb2, 0x84, 0x1a, 0x17, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xa4, 0xc9, 0xc5,
	0xe1, 0x5b, 0x9a, 0x53, 0x92, 0x59, 0x90, 0x53, 0x29, 0xc4, 0xab, 0x87, 0x6c, 0x93, 0x14, 0x9f,
	0x1e, 0x8a, 0x61, 0x4e, 0x2a, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xc5,
	0xa5, 0xa7, 0xa7, 0x5f, 0x0c, 0x31, 0x30, 0x89, 0x0d, 0xec, 0x6c, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe8, 0x7d, 0xeb, 0x0d, 0xc5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArithServiceClient is the client API for ArithService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArithServiceClient interface {
	Multiply(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*ArithResponse, error)
}

type arithServiceClient struct {
	cc *grpc.ClientConn
}

func NewArithServiceClient(cc *grpc.ClientConn) ArithServiceClient {
	return &arithServiceClient{cc}
}

func (c *arithServiceClient) Multiply(ctx context.Context, in *ArithRequest, opts ...grpc.CallOption) (*ArithResponse, error) {
	out := new(ArithResponse)
	err := c.cc.Invoke(ctx, "/ArithService/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithServiceServer is the server API for ArithService service.
type ArithServiceServer interface {
	Multiply(context.Context, *ArithRequest) (*ArithResponse, error)
}

// UnimplementedArithServiceServer can be embedded to have forward compatible implementations.
type UnimplementedArithServiceServer struct {
}

func (*UnimplementedArithServiceServer) Multiply(ctx context.Context, req *ArithRequest) (*ArithResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}

func RegisterArithServiceServer(s *grpc.Server, srv ArithServiceServer) {
	s.RegisterService(&_ArithService_serviceDesc, srv)
}

func _ArithService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArithRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ArithService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithServiceServer).Multiply(ctx, req.(*ArithRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArithService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ArithService",
	HandlerType: (*ArithServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Multiply",
			Handler:    _ArithService_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "arith.proto",
}

func (m *ArithRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ArithRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ArithRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.B != 0 {
		i = encodeVarintArith(dAtA, i, uint64(m.B))
		i--
		dAtA[i] = 0x10
	}
	if m.A != 0 {
		i = encodeVarintArith(dAtA, i, uint64(m.A))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ArithResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ArithResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ArithResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pro != 0 {
		i = encodeVarintArith(dAtA, i, uint64(m.Pro))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintArith(dAtA []byte, offset int, v uint64) int {
	offset -= sovArith(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ArithRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.A != 0 {
		n += 1 + sovArith(uint64(m.A))
	}
	if m.B != 0 {
		n += 1 + sovArith(uint64(m.B))
	}
	return n
}

func (m *ArithResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pro != 0 {
		n += 1 + sovArith(uint64(m.Pro))
	}
	return n
}

func sovArith(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozArith(x uint64) (n int) {
	return sovArith(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ArithRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArith
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
			return fmt.Errorf("proto: ArithRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ArithRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			m.A = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArith
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.A |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			m.B = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArith
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.B |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipArith(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthArith
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthArith
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
func (m *ArithResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArith
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
			return fmt.Errorf("proto: ArithResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ArithResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pro", wireType)
			}
			m.Pro = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArith
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pro |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipArith(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthArith
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthArith
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
func skipArith(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowArith
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
					return 0, ErrIntOverflowArith
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
					return 0, ErrIntOverflowArith
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
				return 0, ErrInvalidLengthArith
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupArith
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthArith
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthArith        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowArith          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupArith = fmt.Errorf("proto: unexpected end of group")
)
