// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: proxy/wireguard/config.proto

package wireguard

import (
	net "github.com/v2fly/v2ray-core/v5/common/net"
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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address       *net.IPOrDomain `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port          uint32          `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Network       net.Network     `protobuf:"varint,3,opt,name=network,proto3,enum=v2ray.core.common.net.Network" json:"network,omitempty"`
	LocalAddress  []string        `protobuf:"bytes,4,rep,name=localAddress,proto3" json:"localAddress,omitempty"`
	PrivateKey    string          `protobuf:"bytes,5,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	PeerPublicKey string          `protobuf:"bytes,6,opt,name=peer_public_key,json=peerPublicKey,proto3" json:"peer_public_key,omitempty"`
	PreSharedKey  string          `protobuf:"bytes,7,opt,name=pre_shared_key,json=preSharedKey,proto3" json:"pre_shared_key,omitempty"`
	Mtu           uint32          `protobuf:"varint,8,opt,name=mtu,proto3" json:"mtu,omitempty"`
	UserLevel     uint32          `protobuf:"varint,9,opt,name=user_level,json=userLevel,proto3" json:"user_level,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_wireguard_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_wireguard_config_proto_msgTypes[0]
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
	return file_proxy_wireguard_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetAddress() *net.IPOrDomain {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Config) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Config) GetNetwork() net.Network {
	if x != nil {
		return x.Network
	}
	return net.Network(0)
}

func (x *Config) GetLocalAddress() []string {
	if x != nil {
		return x.LocalAddress
	}
	return nil
}

func (x *Config) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *Config) GetPeerPublicKey() string {
	if x != nil {
		return x.PeerPublicKey
	}
	return ""
}

func (x *Config) GetPreSharedKey() string {
	if x != nil {
		return x.PreSharedKey
	}
	return ""
}

func (x *Config) GetMtu() uint32 {
	if x != nil {
		return x.Mtu
	}
	return 0
}

func (x *Config) GetUserLevel() uint32 {
	if x != nil {
		return x.UserLevel
	}
	return 0
}

var File_proxy_wireguard_config_proto protoreflect.FileDescriptor

var file_proxy_wireguard_config_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72,
	0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e,
	0x65, 0x74, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf6, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x4f, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x38, 0x0a, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x70,
	0x65, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x65, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x65,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x74, 0x75,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6d, 0x74, 0x75, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x3a, 0x1d, 0x82, 0xb5, 0x18, 0x0a,
	0x0a, 0x08, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x82, 0xb5, 0x18, 0x0b, 0x12, 0x09,
	0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x42, 0x6f, 0x0a, 0x1e, 0x63, 0x6f, 0x6d,
	0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x50, 0x01, 0x5a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0xaa, 0x02, 0x1a,
	0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x57, 0x69, 0x72, 0x65, 0x47, 0x75, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proxy_wireguard_config_proto_rawDescOnce sync.Once
	file_proxy_wireguard_config_proto_rawDescData = file_proxy_wireguard_config_proto_rawDesc
)

func file_proxy_wireguard_config_proto_rawDescGZIP() []byte {
	file_proxy_wireguard_config_proto_rawDescOnce.Do(func() {
		file_proxy_wireguard_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_wireguard_config_proto_rawDescData)
	})
	return file_proxy_wireguard_config_proto_rawDescData
}

var file_proxy_wireguard_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proxy_wireguard_config_proto_goTypes = []interface{}{
	(*Config)(nil),         // 0: v2ray.core.proxy.wireguard.Config
	(*net.IPOrDomain)(nil), // 1: v2ray.core.common.net.IPOrDomain
	(net.Network)(0),       // 2: v2ray.core.common.net.Network
}
var file_proxy_wireguard_config_proto_depIdxs = []int32{
	1, // 0: v2ray.core.proxy.wireguard.Config.address:type_name -> v2ray.core.common.net.IPOrDomain
	2, // 1: v2ray.core.proxy.wireguard.Config.network:type_name -> v2ray.core.common.net.Network
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proxy_wireguard_config_proto_init() }
func file_proxy_wireguard_config_proto_init() {
	if File_proxy_wireguard_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_wireguard_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_proxy_wireguard_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_wireguard_config_proto_goTypes,
		DependencyIndexes: file_proxy_wireguard_config_proto_depIdxs,
		MessageInfos:      file_proxy_wireguard_config_proto_msgTypes,
	}.Build()
	File_proxy_wireguard_config_proto = out.File
	file_proxy_wireguard_config_proto_rawDesc = nil
	file_proxy_wireguard_config_proto_goTypes = nil
	file_proxy_wireguard_config_proto_depIdxs = nil
}
