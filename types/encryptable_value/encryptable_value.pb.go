// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: encryptable_value.proto

package encryptable_value

import (
	encryption "github.com/gford1000-go/protobuf/types/encryption"
	value "github.com/gford1000-go/protobuf/types/value"
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

type EncryptableValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to C:
	//	*EncryptableValue_E
	//	*EncryptableValue_V
	C isEncryptableValue_C `protobuf_oneof:"c"`
}

func (x *EncryptableValue) Reset() {
	*x = EncryptableValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encryptable_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptableValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptableValue) ProtoMessage() {}

func (x *EncryptableValue) ProtoReflect() protoreflect.Message {
	mi := &file_encryptable_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptableValue.ProtoReflect.Descriptor instead.
func (*EncryptableValue) Descriptor() ([]byte, []int) {
	return file_encryptable_value_proto_rawDescGZIP(), []int{0}
}

func (m *EncryptableValue) GetC() isEncryptableValue_C {
	if m != nil {
		return m.C
	}
	return nil
}

func (x *EncryptableValue) GetE() *encryption.EncryptedObject {
	if x, ok := x.GetC().(*EncryptableValue_E); ok {
		return x.E
	}
	return nil
}

func (x *EncryptableValue) GetV() *value.Value {
	if x, ok := x.GetC().(*EncryptableValue_V); ok {
		return x.V
	}
	return nil
}

type isEncryptableValue_C interface {
	isEncryptableValue_C()
}

type EncryptableValue_E struct {
	E *encryption.EncryptedObject `protobuf:"bytes,2,opt,name=e,proto3,oneof"`
}

type EncryptableValue_V struct {
	V *value.Value `protobuf:"bytes,3,opt,name=v,proto3,oneof"`
}

func (*EncryptableValue_E) isEncryptableValue_C() {}

func (*EncryptableValue_V) isEncryptableValue_C() {}

var File_encryptable_value_proto protoreflect.FileDescriptor

var file_encryptable_value_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51,
	0x0a, 0x10, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x20, 0x0a, 0x01, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x48,
	0x00, 0x52, 0x01, 0x65, 0x12, 0x16, 0x0a, 0x01, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x01, 0x76, 0x42, 0x03, 0x0a, 0x01,
	0x63, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x66, 0x6f, 0x72, 0x64, 0x31, 0x30, 0x30, 0x30, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_encryptable_value_proto_rawDescOnce sync.Once
	file_encryptable_value_proto_rawDescData = file_encryptable_value_proto_rawDesc
)

func file_encryptable_value_proto_rawDescGZIP() []byte {
	file_encryptable_value_proto_rawDescOnce.Do(func() {
		file_encryptable_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_encryptable_value_proto_rawDescData)
	})
	return file_encryptable_value_proto_rawDescData
}

var file_encryptable_value_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_encryptable_value_proto_goTypes = []interface{}{
	(*EncryptableValue)(nil),           // 0: EncryptableValue
	(*encryption.EncryptedObject)(nil), // 1: EncryptedObject
	(*value.Value)(nil),                // 2: Value
}
var file_encryptable_value_proto_depIdxs = []int32{
	1, // 0: EncryptableValue.e:type_name -> EncryptedObject
	2, // 1: EncryptableValue.v:type_name -> Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_encryptable_value_proto_init() }
func file_encryptable_value_proto_init() {
	if File_encryptable_value_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_encryptable_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptableValue); i {
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
	file_encryptable_value_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*EncryptableValue_E)(nil),
		(*EncryptableValue_V)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_encryptable_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_encryptable_value_proto_goTypes,
		DependencyIndexes: file_encryptable_value_proto_depIdxs,
		MessageInfos:      file_encryptable_value_proto_msgTypes,
	}.Build()
	File_encryptable_value_proto = out.File
	file_encryptable_value_proto_rawDesc = nil
	file_encryptable_value_proto_goTypes = nil
	file_encryptable_value_proto_depIdxs = nil
}
