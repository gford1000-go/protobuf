// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: row_partition.proto

package row

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

type RowPartition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowId           int64            `protobuf:"varint,1,opt,name=rowId,proto3" json:"rowId,omitempty"`
	AttributeValues map[int32][]byte `protobuf:"bytes,2,rep,name=attributeValues,proto3" json:"attributeValues,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RowPartition) Reset() {
	*x = RowPartition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_row_partition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RowPartition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RowPartition) ProtoMessage() {}

func (x *RowPartition) ProtoReflect() protoreflect.Message {
	mi := &file_row_partition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RowPartition.ProtoReflect.Descriptor instead.
func (*RowPartition) Descriptor() ([]byte, []int) {
	return file_row_partition_proto_rawDescGZIP(), []int{0}
}

func (x *RowPartition) GetRowId() int64 {
	if x != nil {
		return x.RowId
	}
	return 0
}

func (x *RowPartition) GetAttributeValues() map[int32][]byte {
	if x != nil {
		return x.AttributeValues
	}
	return nil
}

var File_row_partition_proto protoreflect.FileDescriptor

var file_row_partition_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x6f, 0x77, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x0c, 0x52, 0x6f, 0x77, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x77, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x0f,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x52, 0x6f, 0x77, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x42, 0x0a, 0x14, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x66, 0x6f,
	0x72, 0x64, 0x31, 0x30, 0x30, 0x30, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_row_partition_proto_rawDescOnce sync.Once
	file_row_partition_proto_rawDescData = file_row_partition_proto_rawDesc
)

func file_row_partition_proto_rawDescGZIP() []byte {
	file_row_partition_proto_rawDescOnce.Do(func() {
		file_row_partition_proto_rawDescData = protoimpl.X.CompressGZIP(file_row_partition_proto_rawDescData)
	})
	return file_row_partition_proto_rawDescData
}

var file_row_partition_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_row_partition_proto_goTypes = []interface{}{
	(*RowPartition)(nil), // 0: RowPartition
	nil,                  // 1: RowPartition.AttributeValuesEntry
}
var file_row_partition_proto_depIdxs = []int32{
	1, // 0: RowPartition.attributeValues:type_name -> RowPartition.AttributeValuesEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_row_partition_proto_init() }
func file_row_partition_proto_init() {
	if File_row_partition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_row_partition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RowPartition); i {
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
			RawDescriptor: file_row_partition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_row_partition_proto_goTypes,
		DependencyIndexes: file_row_partition_proto_depIdxs,
		MessageInfos:      file_row_partition_proto_msgTypes,
	}.Build()
	File_row_partition_proto = out.File
	file_row_partition_proto_rawDesc = nil
	file_row_partition_proto_goTypes = nil
	file_row_partition_proto_depIdxs = nil
}
