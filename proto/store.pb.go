// Code generated by protoc-gen-go. DO NOT EDIT.
// source: store.proto

package store

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Mutation struct {
	Field int32 `protobuf:"varint,1,opt,name=field,proto3" json:"field,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Mutation_StringValue
	//	*Mutation_Int64Value
	Value                isMutation_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Mutation) Reset()         { *m = Mutation{} }
func (m *Mutation) String() string { return proto.CompactTextString(m) }
func (*Mutation) ProtoMessage()    {}
func (*Mutation) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{0}
}

func (m *Mutation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mutation.Unmarshal(m, b)
}
func (m *Mutation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mutation.Marshal(b, m, deterministic)
}
func (m *Mutation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mutation.Merge(m, src)
}
func (m *Mutation) XXX_Size() int {
	return xxx_messageInfo_Mutation.Size(m)
}
func (m *Mutation) XXX_DiscardUnknown() {
	xxx_messageInfo_Mutation.DiscardUnknown(m)
}

var xxx_messageInfo_Mutation proto.InternalMessageInfo

func (m *Mutation) GetField() int32 {
	if m != nil {
		return m.Field
	}
	return 0
}

type isMutation_Value interface {
	isMutation_Value()
}

type Mutation_StringValue struct {
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Mutation_Int64Value struct {
	Int64Value int64 `protobuf:"varint,3,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

func (*Mutation_StringValue) isMutation_Value() {}

func (*Mutation_Int64Value) isMutation_Value() {}

func (m *Mutation) GetValue() isMutation_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Mutation) GetStringValue() string {
	if x, ok := m.GetValue().(*Mutation_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Mutation) GetInt64Value() int64 {
	if x, ok := m.GetValue().(*Mutation_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Mutation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Mutation_OneofMarshaler, _Mutation_OneofUnmarshaler, _Mutation_OneofSizer, []interface{}{
		(*Mutation_StringValue)(nil),
		(*Mutation_Int64Value)(nil),
	}
}

func _Mutation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Mutation)
	// value
	switch x := m.Value.(type) {
	case *Mutation_StringValue:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *Mutation_Int64Value:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int64Value))
	case nil:
	default:
		return fmt.Errorf("Mutation.Value has unexpected type %T", x)
	}
	return nil
}

func _Mutation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Mutation)
	switch tag {
	case 2: // value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Mutation_StringValue{x}
		return true, err
	case 3: // value.int64_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Mutation_Int64Value{int64(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Mutation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Mutation)
	// value
	switch x := m.Value.(type) {
	case *Mutation_StringValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *Mutation_Int64Value:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Int64Value))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Key struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Version              int32    `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{1}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Key) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type WriteRequest struct {
	Key                  *Key        `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Mutations            []*Mutation `protobuf:"bytes,2,rep,name=mutations,proto3" json:"mutations,omitempty"`
	Init                 *any.Any    `protobuf:"bytes,3,opt,name=init,proto3" json:"init,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{2}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *WriteRequest) GetMutations() []*Mutation {
	if m != nil {
		return m.Mutations
	}
	return nil
}

func (m *WriteRequest) GetInit() *any.Any {
	if m != nil {
		return m.Init
	}
	return nil
}

type WriteResponse struct {
	Key                  *Key     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98bbca36ef968dfc, []int{3}
}

func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponse.Unmarshal(m, b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
}
func (m *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(m, src)
}
func (m *WriteResponse) XXX_Size() int {
	return xxx_messageInfo_WriteResponse.Size(m)
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

func (m *WriteResponse) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*Mutation)(nil), "store.Mutation")
	proto.RegisterType((*Key)(nil), "store.Key")
	proto.RegisterType((*WriteRequest)(nil), "store.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "store.WriteResponse")
}

func init() { proto.RegisterFile("store.proto", fileDescriptor_98bbca36ef968dfc) }

var fileDescriptor_98bbca36ef968dfc = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x6b, 0xab, 0x40,
	0x10, 0xc6, 0x63, 0x8c, 0x2f, 0x2f, 0x63, 0x1e, 0x0f, 0xb6, 0x39, 0x48, 0xe8, 0xc1, 0xda, 0x8b,
	0x97, 0x28, 0x98, 0xd0, 0x7b, 0x4b, 0x0f, 0x81, 0xd0, 0xcb, 0x06, 0xda, 0x63, 0xd1, 0x74, 0x62,
	0x96, 0x9a, 0x5d, 0xe3, 0xae, 0x01, 0xff, 0x80, 0xfe, 0xdf, 0xc5, 0xd1, 0xd0, 0xf6, 0xd2, 0xc3,
	0xc2, 0xce, 0x37, 0xbf, 0x61, 0xbe, 0x6f, 0xc0, 0xd5, 0x46, 0x55, 0x18, 0x95, 0x95, 0x32, 0x8a,
	0x39, 0x54, 0xcc, 0xe3, 0x5c, 0x98, 0x43, 0x9d, 0x45, 0x3b, 0x75, 0x8c, 0x73, 0x55, 0xa4, 0x32,
	0x8f, 0xa9, 0x9f, 0xd5, 0xfb, 0xb8, 0x34, 0x4d, 0x89, 0x3a, 0x4e, 0x65, 0xd3, 0xbe, 0x6e, 0x2e,
	0x38, 0xc1, 0xdf, 0xa7, 0xda, 0xa4, 0x46, 0x28, 0xc9, 0x66, 0xe0, 0xec, 0x05, 0x16, 0x6f, 0x9e,
	0xe5, 0x5b, 0xa1, 0xc3, 0xbb, 0x82, 0xdd, 0xc2, 0x54, 0x9b, 0x4a, 0xc8, 0xfc, 0xf5, 0x9c, 0x16,
	0x35, 0x7a, 0x43, 0xdf, 0x0a, 0x27, 0xeb, 0x01, 0x77, 0x3b, 0xf5, 0xb9, 0x15, 0xd9, 0x0d, 0xb8,
	0x42, 0x9a, 0xbb, 0x55, 0xcf, 0xd8, 0xbe, 0x15, 0xda, 0xeb, 0x01, 0x07, 0x12, 0x09, 0x79, 0x18,
	0x83, 0x43, 0xcd, 0x60, 0x09, 0xf6, 0x06, 0x1b, 0xc6, 0x60, 0x54, 0xa6, 0xe6, 0x40, 0xcb, 0x26,
	0x9c, 0xfe, 0xcc, 0x83, 0xf1, 0x19, 0x2b, 0x2d, 0x94, 0xa4, 0x35, 0x0e, 0xbf, 0x94, 0xc1, 0x87,
	0x05, 0xd3, 0x97, 0x4a, 0x18, 0xe4, 0x78, 0xaa, 0x51, 0x1b, 0x76, 0x0d, 0xf6, 0x3b, 0x36, 0x34,
	0xed, 0x26, 0x10, 0x75, 0xb7, 0xd8, 0x60, 0xc3, 0x5b, 0x99, 0x2d, 0x60, 0x72, 0xec, 0x63, 0x69,
	0x6f, 0xe8, 0xdb, 0xa1, 0x9b, 0xfc, 0xef, 0x99, 0x4b, 0x5c, 0xfe, 0x45, 0xb0, 0x10, 0x46, 0x42,
	0x0a, 0x43, 0xbe, 0xdd, 0x64, 0x16, 0xe5, 0x4a, 0xe5, 0x45, 0x7f, 0xda, 0xac, 0xde, 0x47, 0xf7,
	0xb2, 0xe1, 0x44, 0x04, 0x0b, 0xf8, 0xd7, 0xdb, 0xd0, 0xa5, 0x92, 0x1a, 0x7f, 0xf7, 0x91, 0x3c,
	0xc2, 0x74, 0xdb, 0x2a, 0x5b, 0xac, 0xce, 0x62, 0x87, 0x6c, 0x05, 0x0e, 0x8d, 0xb3, 0xab, 0x9e,
	0xfc, 0x9e, 0x69, 0x3e, 0xfb, 0x29, 0x76, 0x1b, 0x82, 0x41, 0xf6, 0x87, 0x8c, 0x2c, 0x3f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x2c, 0xaa, 0xa7, 0x41, 0xf2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoreServiceClient is the client API for StoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreServiceClient interface {
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
}

type storeServiceClient struct {
	cc *grpc.ClientConn
}

func NewStoreServiceClient(cc *grpc.ClientConn) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/store.StoreService/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServiceServer is the server API for StoreService service.
type StoreServiceServer interface {
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
}

func RegisterStoreServiceServer(s *grpc.Server, srv StoreServiceServer) {
	s.RegisterService(&_StoreService_serviceDesc, srv)
}

func _StoreService_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/store.StoreService/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "store.StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _StoreService_Write_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store.proto",
}
