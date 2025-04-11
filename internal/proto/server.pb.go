// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/server.proto

package bootstrapserver

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateBucketReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Bucket        string                 `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBucketReq) Reset() {
	*x = CreateBucketReq{}
	mi := &file_proto_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBucketReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBucketReq) ProtoMessage() {}

func (x *CreateBucketReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBucketReq.ProtoReflect.Descriptor instead.
func (*CreateBucketReq) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBucketReq) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *CreateBucketReq) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

type CreateBucketResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBucketResp) Reset() {
	*x = CreateBucketResp{}
	mi := &file_proto_server_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBucketResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBucketResp) ProtoMessage() {}

func (x *CreateBucketResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBucketResp.ProtoReflect.Descriptor instead.
func (*CreateBucketResp) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBucketResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetProducerConnectionReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Bucket        string                 `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProducerConnectionReq) Reset() {
	*x = GetProducerConnectionReq{}
	mi := &file_proto_server_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProducerConnectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProducerConnectionReq) ProtoMessage() {}

func (x *GetProducerConnectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProducerConnectionReq.ProtoReflect.Descriptor instead.
func (*GetProducerConnectionReq) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{2}
}

func (x *GetProducerConnectionReq) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *GetProducerConnectionReq) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

type GetProducerConnectionResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ProducerId    string                 `protobuf:"bytes,2,opt,name=producerId,proto3" json:"producerId,omitempty"`
	Success       bool                   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProducerConnectionResp) Reset() {
	*x = GetProducerConnectionResp{}
	mi := &file_proto_server_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProducerConnectionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProducerConnectionResp) ProtoMessage() {}

func (x *GetProducerConnectionResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProducerConnectionResp.ProtoReflect.Descriptor instead.
func (*GetProducerConnectionResp) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{3}
}

func (x *GetProducerConnectionResp) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetProducerConnectionResp) GetProducerId() string {
	if x != nil {
		return x.ProducerId
	}
	return ""
}

func (x *GetProducerConnectionResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetConsumerConnnectionReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Bucket        string                 `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetConsumerConnnectionReq) Reset() {
	*x = GetConsumerConnnectionReq{}
	mi := &file_proto_server_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetConsumerConnnectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsumerConnnectionReq) ProtoMessage() {}

func (x *GetConsumerConnnectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsumerConnnectionReq.ProtoReflect.Descriptor instead.
func (*GetConsumerConnnectionReq) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{4}
}

func (x *GetConsumerConnnectionReq) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *GetConsumerConnnectionReq) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

type GetConsumerConnectionResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ConsumerId    string                 `protobuf:"bytes,2,opt,name=consumerId,proto3" json:"consumerId,omitempty"`
	Success       bool                   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetConsumerConnectionResp) Reset() {
	*x = GetConsumerConnectionResp{}
	mi := &file_proto_server_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetConsumerConnectionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsumerConnectionResp) ProtoMessage() {}

func (x *GetConsumerConnectionResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsumerConnectionResp.ProtoReflect.Descriptor instead.
func (*GetConsumerConnectionResp) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{5}
}

func (x *GetConsumerConnectionResp) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetConsumerConnectionResp) GetConsumerId() string {
	if x != nil {
		return x.ConsumerId
	}
	return ""
}

func (x *GetConsumerConnectionResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_server_proto protoreflect.FileDescriptor

const file_proto_server_proto_rawDesc = "" +
	"\n" +
	"\x12proto/server.proto\x12\x0fbootstrapserver\"?\n" +
	"\x0fCreateBucketReq\x12\x14\n" +
	"\x05topic\x18\x01 \x01(\tR\x05topic\x12\x16\n" +
	"\x06bucket\x18\x02 \x01(\tR\x06bucket\",\n" +
	"\x10CreateBucketResp\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"H\n" +
	"\x18GetProducerConnectionReq\x12\x14\n" +
	"\x05topic\x18\x01 \x01(\tR\x05topic\x12\x16\n" +
	"\x06bucket\x18\x02 \x01(\tR\x06bucket\"o\n" +
	"\x19GetProducerConnectionResp\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\tR\aaddress\x12\x1e\n" +
	"\n" +
	"producerId\x18\x02 \x01(\tR\n" +
	"producerId\x12\x18\n" +
	"\asuccess\x18\x03 \x01(\bR\asuccess\"I\n" +
	"\x19GetConsumerConnnectionReq\x12\x14\n" +
	"\x05topic\x18\x01 \x01(\tR\x05topic\x12\x16\n" +
	"\x06bucket\x18\x02 \x01(\tR\x06bucket\"o\n" +
	"\x19GetConsumerConnectionResp\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\tR\aaddress\x12\x1e\n" +
	"\n" +
	"consumerId\x18\x02 \x01(\tR\n" +
	"consumerId\x12\x18\n" +
	"\asuccess\x18\x03 \x01(\bR\asuccess2\xc7\x02\n" +
	"\x0fBootstrapServer\x12S\n" +
	"\fCreateBucket\x12 .bootstrapserver.CreateBucketReq\x1a!.bootstrapserver.CreateBucketResp\x12n\n" +
	"\x15GetProducerConnection\x12).bootstrapserver.GetProducerConnectionReq\x1a*.bootstrapserver.GetProducerConnectionResp\x12o\n" +
	"\x15GetConsumerConnection\x12*.bootstrapserver.GetConsumerConnnectionReq\x1a*.bootstrapserver.GetConsumerConnectionRespB\"Z ./internal/proto;bootstrapserverb\x06proto3"

var (
	file_proto_server_proto_rawDescOnce sync.Once
	file_proto_server_proto_rawDescData []byte
)

func file_proto_server_proto_rawDescGZIP() []byte {
	file_proto_server_proto_rawDescOnce.Do(func() {
		file_proto_server_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_server_proto_rawDesc), len(file_proto_server_proto_rawDesc)))
	})
	return file_proto_server_proto_rawDescData
}

var file_proto_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_server_proto_goTypes = []any{
	(*CreateBucketReq)(nil),           // 0: bootstrapserver.CreateBucketReq
	(*CreateBucketResp)(nil),          // 1: bootstrapserver.CreateBucketResp
	(*GetProducerConnectionReq)(nil),  // 2: bootstrapserver.GetProducerConnectionReq
	(*GetProducerConnectionResp)(nil), // 3: bootstrapserver.GetProducerConnectionResp
	(*GetConsumerConnnectionReq)(nil), // 4: bootstrapserver.GetConsumerConnnectionReq
	(*GetConsumerConnectionResp)(nil), // 5: bootstrapserver.GetConsumerConnectionResp
}
var file_proto_server_proto_depIdxs = []int32{
	0, // 0: bootstrapserver.BootstrapServer.CreateBucket:input_type -> bootstrapserver.CreateBucketReq
	2, // 1: bootstrapserver.BootstrapServer.GetProducerConnection:input_type -> bootstrapserver.GetProducerConnectionReq
	4, // 2: bootstrapserver.BootstrapServer.GetConsumerConnection:input_type -> bootstrapserver.GetConsumerConnnectionReq
	1, // 3: bootstrapserver.BootstrapServer.CreateBucket:output_type -> bootstrapserver.CreateBucketResp
	3, // 4: bootstrapserver.BootstrapServer.GetProducerConnection:output_type -> bootstrapserver.GetProducerConnectionResp
	5, // 5: bootstrapserver.BootstrapServer.GetConsumerConnection:output_type -> bootstrapserver.GetConsumerConnectionResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_server_proto_init() }
func file_proto_server_proto_init() {
	if File_proto_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_server_proto_rawDesc), len(file_proto_server_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_server_proto_goTypes,
		DependencyIndexes: file_proto_server_proto_depIdxs,
		MessageInfos:      file_proto_server_proto_msgTypes,
	}.Build()
	File_proto_server_proto = out.File
	file_proto_server_proto_goTypes = nil
	file_proto_server_proto_depIdxs = nil
}
