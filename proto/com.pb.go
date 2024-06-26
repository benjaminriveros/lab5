// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v5.26.1
// source: proto/com.proto

package proto

import (
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

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tipo  string `protobuf:"bytes,1,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Ns    string `protobuf:"bytes,2,opt,name=ns,proto3" json:"ns,omitempty"`
	Nb    string `protobuf:"bytes,3,opt,name=nb,proto3" json:"nb,omitempty"`
	Valor string `protobuf:"bytes,4,opt,name=valor,proto3" json:"valor,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_com_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_proto_com_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_proto_com_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *Command) GetNs() string {
	if x != nil {
		return x.Ns
	}
	return ""
}

func (x *Command) GetNb() string {
	if x != nil {
		return x.Nb
	}
	return ""
}

func (x *Command) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

type Ip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *Ip) Reset() {
	*x = Ip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_com_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ip) ProtoMessage() {}

func (x *Ip) ProtoReflect() protoreflect.Message {
	mi := &file_proto_com_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ip.ProtoReflect.Descriptor instead.
func (*Ip) Descriptor() ([]byte, []int) {
	return file_proto_com_proto_rawDescGZIP(), []int{1}
}

func (x *Ip) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

var File_proto_com_proto protoreflect.FileDescriptor

var file_proto_com_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x53, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x6e, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x62, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x6e, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x22, 0x14, 0x0a, 0x02,
	0x49, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x32, 0x33, 0x0a, 0x07, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x28, 0x0a,
	0x0b, 0x73, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0d, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x08, 0x2e, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x49, 0x70, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_com_proto_rawDescOnce sync.Once
	file_proto_com_proto_rawDescData = file_proto_com_proto_rawDesc
)

func file_proto_com_proto_rawDescGZIP() []byte {
	file_proto_com_proto_rawDescOnce.Do(func() {
		file_proto_com_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_com_proto_rawDescData)
	})
	return file_proto_com_proto_rawDescData
}

var file_proto_com_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_com_proto_goTypes = []interface{}{
	(*Command)(nil), // 0: code.Command
	(*Ip)(nil),      // 1: code.Ip
}
var file_proto_com_proto_depIdxs = []int32{
	0, // 0: code.General.sendCommand:input_type -> code.Command
	1, // 1: code.General.sendCommand:output_type -> code.Ip
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_com_proto_init() }
func file_proto_com_proto_init() {
	if File_proto_com_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_com_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_proto_com_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ip); i {
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
			RawDescriptor: file_proto_com_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_com_proto_goTypes,
		DependencyIndexes: file_proto_com_proto_depIdxs,
		MessageInfos:      file_proto_com_proto_msgTypes,
	}.Build()
	File_proto_com_proto = out.File
	file_proto_com_proto_rawDesc = nil
	file_proto_com_proto_goTypes = nil
	file_proto_com_proto_depIdxs = nil
}
