// Code generated by proto-gen-gogo. DO NOT EDIT.
// source: github.com/vine-io/services/proto/service/apiserver/v1/apiserver.proto

package apiserverv1

import (
	context "context"
	ebinary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _ = ebinary.BigEndian

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

type Empty struct {
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3bb445ac4faf80b, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return m.XSize()
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type IPRsp struct {
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (m *IPRsp) Reset()         { *m = IPRsp{} }
func (m *IPRsp) String() string { return proto.CompactTextString(m) }
func (*IPRsp) ProtoMessage()    {}
func (*IPRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3bb445ac4faf80b, []int{1}
}
func (m *IPRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IPRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IPRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IPRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPRsp.Merge(m, src)
}
func (m *IPRsp) XXX_Size() int {
	return m.XSize()
}
func (m *IPRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_IPRsp.DiscardUnknown(m)
}

var xxx_messageInfo_IPRsp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Empty)(nil), "apiserverv1.Empty")
	proto.RegisterType((*IPRsp)(nil), "apiserverv1.IPRsp")
}

func init() {
	proto.RegisterFile("github.com/vine-io/services/proto/service/apiserver/v1/apiserver.proto", fileDescriptor_c3bb445ac4faf80b)
}

var fileDescriptor_c3bb445ac4faf80b = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xcb, 0xcc, 0x4b, 0xd5, 0xcd, 0xcc, 0xd7, 0x2f,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0x81, 0x73, 0xf5, 0x13,
	0x0b, 0x32, 0x41, 0xcc, 0xd4, 0x22, 0xfd, 0x32, 0x43, 0x04, 0x47, 0x0f, 0xac, 0x48, 0x88, 0x1b,
	0x2e, 0x50, 0x66, 0xa8, 0xc4, 0xce, 0xc5, 0xea, 0x9a, 0x5b, 0x50, 0x52, 0xa9, 0x24, 0xcd, 0xc5,
	0xea, 0x19, 0x10, 0x54, 0x5c, 0x20, 0x24, 0xc4, 0xc5, 0x92, 0x98, 0x92, 0x52, 0x24, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x1b, 0x55, 0x71, 0x09, 0x38, 0x06, 0x78, 0x06, 0x83, 0x35,
	0x05, 0x43, 0xec, 0x10, 0x32, 0xe6, 0x62, 0xf7, 0x48, 0x4d, 0xcc, 0x29, 0xc9, 0xa8, 0x12, 0x12,
	0xd2, 0x43, 0x32, 0x52, 0x0f, 0x6c, 0x9e, 0x14, 0x16, 0x31, 0x25, 0x06, 0x21, 0x43, 0x2e, 0x56,
	0xf7, 0xd4, 0x12, 0xcf, 0x00, 0x22, 0xb4, 0x80, 0x5d, 0xa3, 0xc4, 0xe0, 0x94, 0x70, 0xe2, 0xa1,
	0x1c, 0xc3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1,
	0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x39, 0x91, 0x17, 0x28,
	0xd6, 0x48, 0x36, 0x25, 0xb1, 0x81, 0xd5, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x8d,
	0x25, 0xf8, 0x61, 0x01, 0x00, 0x00,
}

func (m *Empty) XSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *IPRsp) XSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovApiserver(uint64(l))
	}
	return n
}

func sovApiserver(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func sozApiserver(x uint64) (n int) {
	return sovApiserver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.XSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	size := m.XSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Empty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *IPRsp) Marshal() (dAtA []byte, err error) {
	size := m.XSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IPRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.XSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IPRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintApiserver(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApiserver(dAtA []byte, offset int, v uint64) int {
	offset -= sovApiserver(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiserver
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
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipApiserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApiserver
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
func (m *IPRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiserver
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
			return fmt.Errorf("proto: IPRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IPRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiserver
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
				return ErrInvalidLengthApiserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiserver
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApiserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApiserver
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
func skipApiserver(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApiserver
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
					return 0, ErrIntOverflowApiserver
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
					return 0, ErrIntOverflowApiserver
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
				return 0, ErrInvalidLengthApiserver
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApiserver
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApiserver
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApiserver        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApiserver          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApiserver = fmt.Errorf("proto: unexpected end of group")
)

// APIServerServiceClient is the client API for APIServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIServerServiceClient interface {
	// +gen:get=/api/v1/healthz
	Healthz(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// +gen:get=/api/v1/getIP
	GetIP(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IPRsp, error)
}

type aPIServerServiceClient struct {
	cc *grpc.ClientConn
}

func NewAPIServerServiceClient(cc *grpc.ClientConn) APIServerServiceClient {
	return &aPIServerServiceClient{cc}
}

func (c *aPIServerServiceClient) Healthz(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/apiserverv1.APIServerService/Healthz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServerServiceClient) GetIP(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IPRsp, error) {
	out := new(IPRsp)
	err := c.cc.Invoke(ctx, "/apiserverv1.APIServerService/GetIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServerServiceServer is the server API for APIServerService service.
type APIServerServiceServer interface {
	// +gen:get=/api/v1/healthz
	Healthz(context.Context, *Empty) (*Empty, error)
	// +gen:get=/api/v1/getIP
	GetIP(context.Context, *Empty) (*IPRsp, error)
}

// UnimplementedAPIServerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServerServiceServer struct {
}

func (*UnimplementedAPIServerServiceServer) Healthz(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthz not implemented")
}
func (*UnimplementedAPIServerServiceServer) GetIP(ctx context.Context, req *Empty) (*IPRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIP not implemented")
}

func RegisterAPIServerServiceServer(s *grpc.Server, srv APIServerServiceServer) {
	s.RegisterService(&_APIServerService_serviceDesc, srv)
}

func _APIServerService_Healthz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServerServiceServer).Healthz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserverv1.APIServerService/Healthz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServerServiceServer).Healthz(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIServerService_GetIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServerServiceServer).GetIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserverv1.APIServerService/GetIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServerServiceServer).GetIP(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _APIServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apiserverv1.APIServerService",
	HandlerType: (*APIServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthz",
			Handler:    _APIServerService_Healthz_Handler,
		},
		{
			MethodName: "GetIP",
			Handler:    _APIServerService_GetIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/vine-io/services/proto/service/apiserver/v1/apiserver.proto",
}
