// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: cell.proto

package cell

import (
	encryptable_value "github.com/gford1000-go/protobuf/types/encryptable_value"
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

type Cell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to C:
	//	*Cell_Token
	//	*Cell_Scalar
	C isCell_C `protobuf_oneof:"c"`
}

func (x *Cell) Reset() {
	*x = Cell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cell_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cell) ProtoMessage() {}

func (x *Cell) ProtoReflect() protoreflect.Message {
	mi := &file_cell_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cell.ProtoReflect.Descriptor instead.
func (*Cell) Descriptor() ([]byte, []int) {
	return file_cell_proto_rawDescGZIP(), []int{0}
}

func (m *Cell) GetC() isCell_C {
	if m != nil {
		return m.C
	}
	return nil
}

func (x *Cell) GetToken() uint64 {
	if x, ok := x.GetC().(*Cell_Token); ok {
		return x.Token
	}
	return 0
}

func (x *Cell) GetScalar() *encryptable_value.EncryptableValue {
	if x, ok := x.GetC().(*Cell_Scalar); ok {
		return x.Scalar
	}
	return nil
}

type isCell_C interface {
	isCell_C()
}

type Cell_Token struct {
	Token uint64 `protobuf:"varint,1,opt,name=token,proto3,oneof"`
}

type Cell_Scalar struct {
	Scalar *encryptable_value.EncryptableValue `protobuf:"bytes,2,opt,name=scalar,proto3,oneof"`
}

func (*Cell_Token) isCell_C() {}

func (*Cell_Scalar) isCell_C() {}

type CollectedCells struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenToCellMap map[uint64]*Cell `protobuf:"bytes,1,rep,name=tokenToCellMap,proto3" json:"tokenToCellMap,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CollectedCells) Reset() {
	*x = CollectedCells{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cell_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectedCells) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectedCells) ProtoMessage() {}

func (x *CollectedCells) ProtoReflect() protoreflect.Message {
	mi := &file_cell_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectedCells.ProtoReflect.Descriptor instead.
func (*CollectedCells) Descriptor() ([]byte, []int) {
	return file_cell_proto_rawDescGZIP(), []int{1}
}

func (x *CollectedCells) GetTokenToCellMap() map[uint64]*Cell {
	if x != nil {
		return x.TokenToCellMap
	}
	return nil
}

type TokenMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HashToTokenMap map[string]uint64 `protobuf:"bytes,1,rep,name=hashToTokenMap,proto3" json:"hashToTokenMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *TokenMap) Reset() {
	*x = TokenMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cell_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenMap) ProtoMessage() {}

func (x *TokenMap) ProtoReflect() protoreflect.Message {
	mi := &file_cell_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenMap.ProtoReflect.Descriptor instead.
func (*TokenMap) Descriptor() ([]byte, []int) {
	return file_cell_proto_rawDescGZIP(), []int{2}
}

func (x *TokenMap) GetHashToTokenMap() map[string]uint64 {
	if x != nil {
		return x.HashToTokenMap
	}
	return nil
}

var File_cell_proto protoreflect.FileDescriptor

var file_cell_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x04, 0x43, 0x65, 0x6c, 0x6c, 0x12, 0x16, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x06, 0x73, 0x63, 0x61, 0x6c,
	0x61, 0x72, 0x42, 0x03, 0x0a, 0x01, 0x63, 0x22, 0xa7, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x65, 0x6c, 0x6c, 0x73, 0x12, 0x4b, 0x0a, 0x0e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x54, 0x6f, 0x43, 0x65, 0x6c, 0x6c, 0x4d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x65,
	0x6c, 0x6c, 0x73, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x6f, 0x43, 0x65, 0x6c, 0x6c, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x6f,
	0x43, 0x65, 0x6c, 0x6c, 0x4d, 0x61, 0x70, 0x1a, 0x48, 0x0a, 0x13, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x6f, 0x43, 0x65, 0x6c, 0x6c, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x1b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x94, 0x01, 0x0a, 0x08, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x61, 0x70, 0x12, 0x45,
	0x0a, 0x0e, 0x68, 0x61, 0x73, 0x68, 0x54, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x61, 0x70,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x61,
	0x70, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x54, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x68, 0x54, 0x6f, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x4d, 0x61, 0x70, 0x1a, 0x41, 0x0a, 0x13, 0x48, 0x61, 0x73, 0x68, 0x54, 0x6f, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x66, 0x6f, 0x72, 0x64, 0x31, 0x30, 0x30, 0x30,
	0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x63, 0x65, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cell_proto_rawDescOnce sync.Once
	file_cell_proto_rawDescData = file_cell_proto_rawDesc
)

func file_cell_proto_rawDescGZIP() []byte {
	file_cell_proto_rawDescOnce.Do(func() {
		file_cell_proto_rawDescData = protoimpl.X.CompressGZIP(file_cell_proto_rawDescData)
	})
	return file_cell_proto_rawDescData
}

var file_cell_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cell_proto_goTypes = []interface{}{
	(*Cell)(nil),           // 0: Cell
	(*CollectedCells)(nil), // 1: CollectedCells
	(*TokenMap)(nil),       // 2: TokenMap
	nil,                    // 3: CollectedCells.TokenToCellMapEntry
	nil,                    // 4: TokenMap.HashToTokenMapEntry
	(*encryptable_value.EncryptableValue)(nil), // 5: EncryptableValue
}
var file_cell_proto_depIdxs = []int32{
	5, // 0: Cell.scalar:type_name -> EncryptableValue
	3, // 1: CollectedCells.tokenToCellMap:type_name -> CollectedCells.TokenToCellMapEntry
	4, // 2: TokenMap.hashToTokenMap:type_name -> TokenMap.HashToTokenMapEntry
	0, // 3: CollectedCells.TokenToCellMapEntry.value:type_name -> Cell
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cell_proto_init() }
func file_cell_proto_init() {
	if File_cell_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cell_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cell); i {
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
		file_cell_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectedCells); i {
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
		file_cell_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenMap); i {
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
	file_cell_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Cell_Token)(nil),
		(*Cell_Scalar)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cell_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cell_proto_goTypes,
		DependencyIndexes: file_cell_proto_depIdxs,
		MessageInfos:      file_cell_proto_msgTypes,
	}.Build()
	File_cell_proto = out.File
	file_cell_proto_rawDesc = nil
	file_cell_proto_goTypes = nil
	file_cell_proto_depIdxs = nil
}
