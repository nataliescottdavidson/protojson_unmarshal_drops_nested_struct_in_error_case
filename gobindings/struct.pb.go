// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.1
// source: struct.proto

package gobindings

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

type NestedStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bang              string `protobuf:"bytes,1,opt,name=bang,proto3" json:"bang,omitempty"`
	IntButWeGetString int32  `protobuf:"varint,2,opt,name=intButWeGetString,proto3" json:"intButWeGetString,omitempty"`
}

func (x *NestedStruct) Reset() {
	*x = NestedStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_struct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedStruct) ProtoMessage() {}

func (x *NestedStruct) ProtoReflect() protoreflect.Message {
	mi := &file_struct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedStruct.ProtoReflect.Descriptor instead.
func (*NestedStruct) Descriptor() ([]byte, []int) {
	return file_struct_proto_rawDescGZIP(), []int{0}
}

func (x *NestedStruct) GetBang() string {
	if x != nil {
		return x.Bang
	}
	return ""
}

func (x *NestedStruct) GetIntButWeGetString() int32 {
	if x != nil {
		return x.IntButWeGetString
	}
	return 0
}

type Struct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Baz           string        `protobuf:"bytes,1,opt,name=baz,proto3" json:"baz,omitempty"`
	ContainsError *NestedStruct `protobuf:"bytes,2,opt,name=containsError,proto3" json:"containsError,omitempty"`
}

func (x *Struct) Reset() {
	*x = Struct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_struct_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Struct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Struct) ProtoMessage() {}

func (x *Struct) ProtoReflect() protoreflect.Message {
	mi := &file_struct_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Struct.ProtoReflect.Descriptor instead.
func (*Struct) Descriptor() ([]byte, []int) {
	return file_struct_proto_rawDescGZIP(), []int{1}
}

func (x *Struct) GetBaz() string {
	if x != nil {
		return x.Baz
	}
	return ""
}

func (x *Struct) GetContainsError() *NestedStruct {
	if x != nil {
		return x.ContainsError
	}
	return nil
}

var File_struct_proto protoreflect.FileDescriptor

var file_struct_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x6d, 0x79, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x50, 0x0a, 0x0c, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x6e, 0x67, 0x12, 0x2c,
	0x0a, 0x11, 0x69, 0x6e, 0x74, 0x42, 0x75, 0x74, 0x57, 0x65, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x69, 0x6e, 0x74, 0x42, 0x75,
	0x74, 0x57, 0x65, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x5a, 0x0a, 0x06,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x7a, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x61, 0x7a, 0x12, 0x3e, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x6d, 0x79, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x67, 0x6f,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_struct_proto_rawDescOnce sync.Once
	file_struct_proto_rawDescData = file_struct_proto_rawDesc
)

func file_struct_proto_rawDescGZIP() []byte {
	file_struct_proto_rawDescOnce.Do(func() {
		file_struct_proto_rawDescData = protoimpl.X.CompressGZIP(file_struct_proto_rawDescData)
	})
	return file_struct_proto_rawDescData
}

var file_struct_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_struct_proto_goTypes = []any{
	(*NestedStruct)(nil), // 0: my.package.NestedStruct
	(*Struct)(nil),       // 1: my.package.Struct
}
var file_struct_proto_depIdxs = []int32{
	0, // 0: my.package.Struct.containsError:type_name -> my.package.NestedStruct
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_struct_proto_init() }
func file_struct_proto_init() {
	if File_struct_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_struct_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NestedStruct); i {
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
		file_struct_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Struct); i {
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
			RawDescriptor: file_struct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_struct_proto_goTypes,
		DependencyIndexes: file_struct_proto_depIdxs,
		MessageInfos:      file_struct_proto_msgTypes,
	}.Build()
	File_struct_proto = out.File
	file_struct_proto_rawDesc = nil
	file_struct_proto_goTypes = nil
	file_struct_proto_depIdxs = nil
}
