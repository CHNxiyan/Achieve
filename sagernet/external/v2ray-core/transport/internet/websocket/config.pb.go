// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: transport/internet/websocket/config.proto

package websocket

import (
	_ "github.com/v2fly/v2ray-core/v5/common/protoext"
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

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_websocket_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_websocket_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_transport_internet_websocket_config_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Header) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL path to the WebSocket service. Empty value means root(/).
	Path                 string    `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Header               []*Header `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty"`
	AcceptProxyProtocol  bool      `protobuf:"varint,4,opt,name=accept_proxy_protocol,json=acceptProxyProtocol,proto3" json:"accept_proxy_protocol,omitempty"`
	MaxEarlyData         int32     `protobuf:"varint,5,opt,name=max_early_data,json=maxEarlyData,proto3" json:"max_early_data,omitempty"`
	UseBrowserForwarding bool      `protobuf:"varint,6,opt,name=use_browser_forwarding,json=useBrowserForwarding,proto3" json:"use_browser_forwarding,omitempty"`
	EarlyDataHeaderName  string    `protobuf:"bytes,7,opt,name=early_data_header_name,json=earlyDataHeaderName,proto3" json:"early_data_header_name,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_websocket_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_websocket_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_transport_internet_websocket_config_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Config) GetHeader() []*Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Config) GetAcceptProxyProtocol() bool {
	if x != nil {
		return x.AcceptProxyProtocol
	}
	return false
}

func (x *Config) GetMaxEarlyData() int32 {
	if x != nil {
		return x.MaxEarlyData
	}
	return 0
}

func (x *Config) GetUseBrowserForwarding() bool {
	if x != nil {
		return x.UseBrowserForwarding
	}
	return false
}

func (x *Config) GetEarlyDataHeaderName() string {
	if x != nil {
		return x.EarlyDataHeaderName
	}
	return ""
}

var File_transport_internet_websocket_config_proto protoreflect.FileDescriptor

var file_transport_internet_websocket_config_proto_rawDesc = []byte{
	0x0a, 0x29, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xda, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x47, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x32, 0x0a, 0x15, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x13, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x61, 0x72, 0x6c,
	0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61,
	0x78, 0x45, 0x61, 0x72, 0x6c, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x16, 0x75, 0x73,
	0x65, 0x5f, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x75, 0x73, 0x65, 0x42,
	0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x33, 0x0a, 0x16, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x44, 0x61, 0x74, 0x61, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x28, 0x82, 0xb5, 0x18, 0x0b, 0x0a, 0x09, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x82, 0xb5, 0x18, 0x04, 0x12, 0x02, 0x77, 0x73, 0x82, 0xb5,
	0x18, 0x0d, 0x8a, 0xff, 0x29, 0x09, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4a,
	0x04, 0x08, 0x01, 0x10, 0x02, 0x42, 0x96, 0x01, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32,
	0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x77, 0x65, 0x62, 0x73,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0xaa, 0x02, 0x27, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72,
	0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2e, 0x57, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_internet_websocket_config_proto_rawDescOnce sync.Once
	file_transport_internet_websocket_config_proto_rawDescData = file_transport_internet_websocket_config_proto_rawDesc
)

func file_transport_internet_websocket_config_proto_rawDescGZIP() []byte {
	file_transport_internet_websocket_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_websocket_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_internet_websocket_config_proto_rawDescData)
	})
	return file_transport_internet_websocket_config_proto_rawDescData
}

var file_transport_internet_websocket_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transport_internet_websocket_config_proto_goTypes = []interface{}{
	(*Header)(nil), // 0: v2ray.core.transport.internet.websocket.Header
	(*Config)(nil), // 1: v2ray.core.transport.internet.websocket.Config
}
var file_transport_internet_websocket_config_proto_depIdxs = []int32{
	0, // 0: v2ray.core.transport.internet.websocket.Config.header:type_name -> v2ray.core.transport.internet.websocket.Header
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_transport_internet_websocket_config_proto_init() }
func file_transport_internet_websocket_config_proto_init() {
	if File_transport_internet_websocket_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_internet_websocket_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_transport_internet_websocket_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_transport_internet_websocket_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_websocket_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_websocket_config_proto_depIdxs,
		MessageInfos:      file_transport_internet_websocket_config_proto_msgTypes,
	}.Build()
	File_transport_internet_websocket_config_proto = out.File
	file_transport_internet_websocket_config_proto_rawDesc = nil
	file_transport_internet_websocket_config_proto_goTypes = nil
	file_transport_internet_websocket_config_proto_depIdxs = nil
}
