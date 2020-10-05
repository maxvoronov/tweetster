// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tweets.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Procedures
type PostsGetListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostsGetListRequest) Reset()         { *m = PostsGetListRequest{} }
func (m *PostsGetListRequest) String() string { return proto.CompactTextString(m) }
func (*PostsGetListRequest) ProtoMessage()    {}
func (*PostsGetListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8cfa2ff8b4ecacd, []int{0}
}

func (m *PostsGetListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostsGetListRequest.Unmarshal(m, b)
}
func (m *PostsGetListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostsGetListRequest.Marshal(b, m, deterministic)
}
func (m *PostsGetListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostsGetListRequest.Merge(m, src)
}
func (m *PostsGetListRequest) XXX_Size() int {
	return xxx_messageInfo_PostsGetListRequest.Size(m)
}
func (m *PostsGetListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostsGetListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostsGetListRequest proto.InternalMessageInfo

type PostsGetListResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostsGetListResponse) Reset()         { *m = PostsGetListResponse{} }
func (m *PostsGetListResponse) String() string { return proto.CompactTextString(m) }
func (*PostsGetListResponse) ProtoMessage()    {}
func (*PostsGetListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8cfa2ff8b4ecacd, []int{1}
}

func (m *PostsGetListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostsGetListResponse.Unmarshal(m, b)
}
func (m *PostsGetListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostsGetListResponse.Marshal(b, m, deterministic)
}
func (m *PostsGetListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostsGetListResponse.Merge(m, src)
}
func (m *PostsGetListResponse) XXX_Size() int {
	return xxx_messageInfo_PostsGetListResponse.Size(m)
}
func (m *PostsGetListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostsGetListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostsGetListResponse proto.InternalMessageInfo

func (m *PostsGetListResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

type PostsGetByIdRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostsGetByIdRequest) Reset()         { *m = PostsGetByIdRequest{} }
func (m *PostsGetByIdRequest) String() string { return proto.CompactTextString(m) }
func (*PostsGetByIdRequest) ProtoMessage()    {}
func (*PostsGetByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8cfa2ff8b4ecacd, []int{2}
}

func (m *PostsGetByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostsGetByIdRequest.Unmarshal(m, b)
}
func (m *PostsGetByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostsGetByIdRequest.Marshal(b, m, deterministic)
}
func (m *PostsGetByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostsGetByIdRequest.Merge(m, src)
}
func (m *PostsGetByIdRequest) XXX_Size() int {
	return xxx_messageInfo_PostsGetByIdRequest.Size(m)
}
func (m *PostsGetByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostsGetByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostsGetByIdRequest proto.InternalMessageInfo

func (m *PostsGetByIdRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PostsGetByIdResponse struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostsGetByIdResponse) Reset()         { *m = PostsGetByIdResponse{} }
func (m *PostsGetByIdResponse) String() string { return proto.CompactTextString(m) }
func (*PostsGetByIdResponse) ProtoMessage()    {}
func (*PostsGetByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8cfa2ff8b4ecacd, []int{3}
}

func (m *PostsGetByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostsGetByIdResponse.Unmarshal(m, b)
}
func (m *PostsGetByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostsGetByIdResponse.Marshal(b, m, deterministic)
}
func (m *PostsGetByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostsGetByIdResponse.Merge(m, src)
}
func (m *PostsGetByIdResponse) XXX_Size() int {
	return xxx_messageInfo_PostsGetByIdResponse.Size(m)
}
func (m *PostsGetByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostsGetByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostsGetByIdResponse proto.InternalMessageInfo

func (m *PostsGetByIdResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

// Models
type Post struct {
	Id                   uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId             uint64               `protobuf:"varint,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Content              string               `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8cfa2ff8b4ecacd, []int{4}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Post) GetAuthorId() uint64 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Post) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*PostsGetListRequest)(nil), "pb.PostsGetListRequest")
	proto.RegisterType((*PostsGetListResponse)(nil), "pb.PostsGetListResponse")
	proto.RegisterType((*PostsGetByIdRequest)(nil), "pb.PostsGetByIdRequest")
	proto.RegisterType((*PostsGetByIdResponse)(nil), "pb.PostsGetByIdResponse")
	proto.RegisterType((*Post)(nil), "pb.Post")
}

func init() { proto.RegisterFile("tweets.proto", fileDescriptor_a8cfa2ff8b4ecacd) }

var fileDescriptor_a8cfa2ff8b4ecacd = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0x87, 0xd7, 0xae, 0xff, 0xbf, 0xdb, 0xeb, 0xf4, 0x10, 0x15, 0xc3, 0x14, 0x1d, 0x01, 0x61,
	0xa7, 0x0c, 0xa6, 0x08, 0x1e, 0xd5, 0x83, 0x0c, 0x3c, 0x48, 0xdd, 0x7d, 0xb4, 0xcb, 0xeb, 0x0c,
	0xb8, 0x25, 0x36, 0xef, 0x14, 0xbf, 0x81, 0x5f, 0xc1, 0x6f, 0x2b, 0x4d, 0x56, 0xd6, 0xcd, 0x63,
	0x9e, 0xfe, 0x78, 0x9a, 0x3c, 0xd0, 0xa1, 0x4f, 0x44, 0x72, 0xd2, 0x16, 0x86, 0x0c, 0x8b, 0x6d,
	0xde, 0x3d, 0x9f, 0x19, 0x33, 0x7b, 0xc3, 0x81, 0x27, 0xf9, 0xf2, 0x65, 0x40, 0x7a, 0x8e, 0x8e,
	0xb2, 0xb9, 0x0d, 0x23, 0x71, 0x04, 0x07, 0x4f, 0xc6, 0x91, 0x7b, 0x40, 0x7a, 0xd4, 0x8e, 0x52,
	0x7c, 0x5f, 0xa2, 0x23, 0x71, 0x0d, 0x87, 0x9b, 0xd8, 0x59, 0xb3, 0x70, 0xc8, 0xce, 0xe0, 0x9f,
	0x2d, 0x39, 0x8f, 0x7a, 0xcd, 0xfe, 0xee, 0xb0, 0x25, 0x6d, 0x2e, 0xcb, 0x61, 0x1a, 0xb0, 0xb8,
	0x58, 0xeb, 0xee, 0xbe, 0x46, 0x6a, 0xa5, 0x63, 0xfb, 0x10, 0x6b, 0xc5, 0xa3, 0x5e, 0xd4, 0x4f,
	0xd2, 0x58, 0x2b, 0x71, 0xb5, 0xd6, 0x87, 0xd9, 0x4a, 0x7f, 0x0a, 0x49, 0xe9, 0xf1, 0xcb, 0xba,
	0xdd, 0x53, 0xf1, 0x1d, 0x41, 0x52, 0x1e, 0xb7, 0x75, 0xec, 0x04, 0xda, 0xd9, 0x92, 0x5e, 0x4d,
	0x31, 0xd1, 0x8a, 0xc7, 0x1e, 0xb7, 0x02, 0x18, 0x29, 0xc6, 0x61, 0x67, 0x6a, 0x16, 0x84, 0x0b,
	0xe2, 0xcd, 0x5e, 0xd4, 0x6f, 0xa7, 0xd5, 0x91, 0xdd, 0x00, 0x4c, 0x0b, 0xcc, 0x08, 0xd5, 0x24,
	0x23, 0x9e, 0xf8, 0x7f, 0x76, 0x65, 0x28, 0x26, 0xab, 0x62, 0x72, 0x5c, 0x15, 0x4b, 0xdb, 0xab,
	0xf5, 0x2d, 0x0d, 0x7f, 0x22, 0xd8, 0x1b, 0xfb, 0xd8, 0xcf, 0x58, 0x7c, 0xe8, 0x29, 0xb2, 0x7b,
	0xe8, 0xd4, 0x8b, 0xb1, 0xe3, 0xea, 0xf2, 0x5b, 0x69, 0xbb, 0xfc, 0xef, 0x87, 0xf0, 0x7a, 0xd1,
	0xa8, 0x4b, 0xca, 0x2e, 0x9b, 0x92, 0x5a, 0xd0, 0x4d, 0x49, 0x3d, 0xa1, 0x68, 0xe4, 0xff, 0xfd,
	0xd5, 0x2f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x7d, 0x34, 0x38, 0x0e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TweetsServiceClient is the client API for TweetsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TweetsServiceClient interface {
	PostsGetList(ctx context.Context, in *PostsGetListRequest, opts ...grpc.CallOption) (*PostsGetListResponse, error)
	PostsGetById(ctx context.Context, in *PostsGetByIdRequest, opts ...grpc.CallOption) (*PostsGetByIdResponse, error)
}

type tweetsServiceClient struct {
	cc *grpc.ClientConn
}

func NewTweetsServiceClient(cc *grpc.ClientConn) TweetsServiceClient {
	return &tweetsServiceClient{cc}
}

func (c *tweetsServiceClient) PostsGetList(ctx context.Context, in *PostsGetListRequest, opts ...grpc.CallOption) (*PostsGetListResponse, error) {
	out := new(PostsGetListResponse)
	err := c.cc.Invoke(ctx, "/pb.TweetsService/PostsGetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetsServiceClient) PostsGetById(ctx context.Context, in *PostsGetByIdRequest, opts ...grpc.CallOption) (*PostsGetByIdResponse, error) {
	out := new(PostsGetByIdResponse)
	err := c.cc.Invoke(ctx, "/pb.TweetsService/PostsGetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TweetsServiceServer is the server API for TweetsService service.
type TweetsServiceServer interface {
	PostsGetList(context.Context, *PostsGetListRequest) (*PostsGetListResponse, error)
	PostsGetById(context.Context, *PostsGetByIdRequest) (*PostsGetByIdResponse, error)
}

// UnimplementedTweetsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTweetsServiceServer struct {
}

func (*UnimplementedTweetsServiceServer) PostsGetList(ctx context.Context, req *PostsGetListRequest) (*PostsGetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostsGetList not implemented")
}
func (*UnimplementedTweetsServiceServer) PostsGetById(ctx context.Context, req *PostsGetByIdRequest) (*PostsGetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostsGetById not implemented")
}

func RegisterTweetsServiceServer(s *grpc.Server, srv TweetsServiceServer) {
	s.RegisterService(&_TweetsService_serviceDesc, srv)
}

func _TweetsService_PostsGetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostsGetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServiceServer).PostsGetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TweetsService/PostsGetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServiceServer).PostsGetList(ctx, req.(*PostsGetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetsService_PostsGetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostsGetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServiceServer).PostsGetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TweetsService/PostsGetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServiceServer).PostsGetById(ctx, req.(*PostsGetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TweetsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TweetsService",
	HandlerType: (*TweetsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostsGetList",
			Handler:    _TweetsService_PostsGetList_Handler,
		},
		{
			MethodName: "PostsGetById",
			Handler:    _TweetsService_PostsGetById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tweets.proto",
}
