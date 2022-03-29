// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: algo.proto

package encryption

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

type Algo int32

const (
	Algo_UnknownAlgo Algo = 0
	Algo_GCM         Algo = 1
)

// Enum value maps for Algo.
var (
	Algo_name = map[int32]string{
		0: "UnknownAlgo",
		1: "GCM",
	}
	Algo_value = map[string]int32{
		"UnknownAlgo": 0,
		"GCM":         1,
	}
)

func (x Algo) Enum() *Algo {
	p := new(Algo)
	*p = x
	return p
}

func (x Algo) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Algo) Descriptor() protoreflect.EnumDescriptor {
	return file_algo_proto_enumTypes[0].Descriptor()
}

func (Algo) Type() protoreflect.EnumType {
	return &file_algo_proto_enumTypes[0]
}

func (x Algo) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Algo.Descriptor instead.
func (Algo) EnumDescriptor() ([]byte, []int) {
	return file_algo_proto_rawDescGZIP(), []int{0}
}

var File_algo_proto protoreflect.FileDescriptor

var file_algo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x6c, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x20, 0x0a, 0x04,
	0x41, 0x6c, 0x67, 0x6f, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x41,
	0x6c, 0x67, 0x6f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x43, 0x4d, 0x10, 0x01, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x66, 0x6f,
	0x72, 0x64, 0x31, 0x30, 0x30, 0x30, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_algo_proto_rawDescOnce sync.Once
	file_algo_proto_rawDescData = file_algo_proto_rawDesc
)

func file_algo_proto_rawDescGZIP() []byte {
	file_algo_proto_rawDescOnce.Do(func() {
		file_algo_proto_rawDescData = protoimpl.X.CompressGZIP(file_algo_proto_rawDescData)
	})
	return file_algo_proto_rawDescData
}

var file_algo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_algo_proto_goTypes = []interface{}{
	(Algo)(0), // 0: Algo
}
var file_algo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_algo_proto_init() }
func file_algo_proto_init() {
	if File_algo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_algo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_algo_proto_goTypes,
		DependencyIndexes: file_algo_proto_depIdxs,
		EnumInfos:         file_algo_proto_enumTypes,
	}.Build()
	File_algo_proto = out.File
	file_algo_proto_rawDesc = nil
	file_algo_proto_goTypes = nil
	file_algo_proto_depIdxs = nil
}
