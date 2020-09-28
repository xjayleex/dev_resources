// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: image.proto

package imageproto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TransStatCode int32

const (
	TransStatCode_Unknown TransStatCode = 0
	TransStatCode_Ok      TransStatCode = 1
	TransStatCode_Failed  TransStatCode = 2
)

// Enum value maps for TransStatCode.
var (
	TransStatCode_name = map[int32]string{
		0: "Unknown",
		1: "Ok",
		2: "Failed",
	}
	TransStatCode_value = map[string]int32{
		"Unknown": 0,
		"Ok":      1,
		"Failed":  2,
	}
)

func (x TransStatCode) Enum() *TransStatCode {
	p := new(TransStatCode)
	*p = x
	return p
}

func (x TransStatCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransStatCode) Descriptor() protoreflect.EnumDescriptor {
	return file_image_proto_enumTypes[0].Descriptor()
}

func (TransStatCode) Type() protoreflect.EnumType {
	return &file_image_proto_enumTypes[0]
}

func (x TransStatCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransStatCode.Descriptor instead.
func (TransStatCode) EnumDescriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{0}
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type TransferStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string        `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	StatusCode TransStatCode `protobuf:"varint,2,opt,name=StatusCode,proto3,enum=imageproto.TransStatCode" json:"StatusCode,omitempty"`
}

func (x *TransferStatus) Reset() {
	*x = TransferStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferStatus) ProtoMessage() {}

func (x *TransferStatus) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferStatus.ProtoReflect.Descriptor instead.
func (*TransferStatus) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{1}
}

func (x *TransferStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TransferStatus) GetStatusCode() TransStatCode {
	if x != nil {
		return x.StatusCode
	}
	return TransStatCode_Unknown
}

var File_image_proto protoreflect.FileDescriptor

var file_image_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x05, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x65, 0x0a, 0x0e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x2a, 0x30, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10,
	0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x10, 0x02, 0x32, 0x4f, 0x0a, 0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x11, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x1a, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_image_proto_rawDescOnce sync.Once
	file_image_proto_rawDescData = file_image_proto_rawDesc
)

func file_image_proto_rawDescGZIP() []byte {
	file_image_proto_rawDescOnce.Do(func() {
		file_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_image_proto_rawDescData)
	})
	return file_image_proto_rawDescData
}

var file_image_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_image_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_image_proto_goTypes = []interface{}{
	(TransStatCode)(0),     // 0: imageproto.TransStatCode
	(*Chunk)(nil),          // 1: imageproto.Chunk
	(*TransferStatus)(nil), // 2: imageproto.TransferStatus
}
var file_image_proto_depIdxs = []int32{
	0, // 0: imageproto.TransferStatus.StatusCode:type_name -> imageproto.TransStatCode
	1, // 1: imageproto.ImageTransfer.SendImage:input_type -> imageproto.Chunk
	2, // 2: imageproto.ImageTransfer.SendImage:output_type -> imageproto.TransferStatus
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_image_proto_init() }
func file_image_proto_init() {
	if File_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_image_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_image_proto_goTypes,
		DependencyIndexes: file_image_proto_depIdxs,
		EnumInfos:         file_image_proto_enumTypes,
		MessageInfos:      file_image_proto_msgTypes,
	}.Build()
	File_image_proto = out.File
	file_image_proto_rawDesc = nil
	file_image_proto_goTypes = nil
	file_image_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ImageTransferClient is the client API for ImageTransfer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImageTransferClient interface {
	SendImage(ctx context.Context, opts ...grpc.CallOption) (ImageTransfer_SendImageClient, error)
}

type imageTransferClient struct {
	cc grpc.ClientConnInterface
}

func NewImageTransferClient(cc grpc.ClientConnInterface) ImageTransferClient {
	return &imageTransferClient{cc}
}

func (c *imageTransferClient) SendImage(ctx context.Context, opts ...grpc.CallOption) (ImageTransfer_SendImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ImageTransfer_serviceDesc.Streams[0], "/imageproto.ImageTransfer/SendImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageTransferSendImageClient{stream}
	return x, nil
}

type ImageTransfer_SendImageClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*TransferStatus, error)
	grpc.ClientStream
}

type imageTransferSendImageClient struct {
	grpc.ClientStream
}

func (x *imageTransferSendImageClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imageTransferSendImageClient) CloseAndRecv() (*TransferStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TransferStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageTransferServer is the server API for ImageTransfer service.
type ImageTransferServer interface {
	SendImage(ImageTransfer_SendImageServer) error
}

// UnimplementedImageTransferServer can be embedded to have forward compatible implementations.
type UnimplementedImageTransferServer struct {
}

func (*UnimplementedImageTransferServer) SendImage(ImageTransfer_SendImageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendImage not implemented")
}

func RegisterImageTransferServer(s *grpc.Server, srv ImageTransferServer) {
	s.RegisterService(&_ImageTransfer_serviceDesc, srv)
}

func _ImageTransfer_SendImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImageTransferServer).SendImage(&imageTransferSendImageServer{stream})
}

type ImageTransfer_SendImageServer interface {
	SendAndClose(*TransferStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type imageTransferSendImageServer struct {
	grpc.ServerStream
}

func (x *imageTransferSendImageServer) SendAndClose(m *TransferStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imageTransferSendImageServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ImageTransfer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "imageproto.ImageTransfer",
	HandlerType: (*ImageTransferServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendImage",
			Handler:       _ImageTransfer_SendImage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "image.proto",
}
