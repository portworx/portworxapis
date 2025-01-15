// Please use the following editor setup for this file:
// Tab size=2; Tabs as spaces; Clean up trailing whitepsace
//
// In vim add: au FileType proto setl sw=2 ts=2 expandtab list
//
// In vscode install vscode-proto3 extension and add this to your settings.json:
//    "[proto3]": {
//        "editor.tabSize": 2,
//        "editor.insertSpaces": true,
//        "editor.rulers": [80],
//        "editor.detectIndentation": true,
//        "files.trimTrailingWhitespace": true
//    }
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: public/portworx/common/apiv1/sort.proto

package common

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

// Field names for sorting the list of resources.
type SortBy_Field int32

const (
	// Unspecified, do not use.
	SortBy_FIELD_UNSPECIFIED SortBy_Field = 0
	// Sorting based on the name of the resource.
	SortBy_NAME SortBy_Field = 1
	// Sorting on create time of the resource.
	SortBy_CREATED_AT SortBy_Field = 2
	// Sorting on update time of the resource.
	SortBy_UPDATED_AT SortBy_Field = 3
	// Sorting on phase of the resource.
	SortBy_PHASE SortBy_Field = 4
)

// Enum value maps for SortBy_Field.
var (
	SortBy_Field_name = map[int32]string{
		0: "FIELD_UNSPECIFIED",
		1: "NAME",
		2: "CREATED_AT",
		3: "UPDATED_AT",
		4: "PHASE",
	}
	SortBy_Field_value = map[string]int32{
		"FIELD_UNSPECIFIED": 0,
		"NAME":              1,
		"CREATED_AT":        2,
		"UPDATED_AT":        3,
		"PHASE":             4,
	}
)

func (x SortBy_Field) Enum() *SortBy_Field {
	p := new(SortBy_Field)
	*p = x
	return p
}

func (x SortBy_Field) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortBy_Field) Descriptor() protoreflect.EnumDescriptor {
	return file_public_portworx_common_apiv1_sort_proto_enumTypes[0].Descriptor()
}

func (SortBy_Field) Type() protoreflect.EnumType {
	return &file_public_portworx_common_apiv1_sort_proto_enumTypes[0]
}

func (x SortBy_Field) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortBy_Field.Descriptor instead.
func (SortBy_Field) EnumDescriptor() ([]byte, []int) {
	return file_public_portworx_common_apiv1_sort_proto_rawDescGZIP(), []int{1, 0}
}

// Value of sort order for the list of resources.
type SortOrder_Value int32

const (
	// Unspecified, do not use.
	SortOrder_VALUE_UNSPECIFIED SortOrder_Value = 0
	// Sort order ascending.
	SortOrder_ASC SortOrder_Value = 1
	// Sort order descending.
	SortOrder_DESC SortOrder_Value = 2
)

// Enum value maps for SortOrder_Value.
var (
	SortOrder_Value_name = map[int32]string{
		0: "VALUE_UNSPECIFIED",
		1: "ASC",
		2: "DESC",
	}
	SortOrder_Value_value = map[string]int32{
		"VALUE_UNSPECIFIED": 0,
		"ASC":               1,
		"DESC":              2,
	}
)

func (x SortOrder_Value) Enum() *SortOrder_Value {
	p := new(SortOrder_Value)
	*p = x
	return p
}

func (x SortOrder_Value) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder_Value) Descriptor() protoreflect.EnumDescriptor {
	return file_public_portworx_common_apiv1_sort_proto_enumTypes[1].Descriptor()
}

func (SortOrder_Value) Type() protoreflect.EnumType {
	return &file_public_portworx_common_apiv1_sort_proto_enumTypes[1]
}

func (x SortOrder_Value) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder_Value.Descriptor instead.
func (SortOrder_Value) EnumDescriptor() ([]byte, []int) {
	return file_public_portworx_common_apiv1_sort_proto_rawDescGZIP(), []int{2, 0}
}

// The details of the attribute for which the requested list of resource to be sorted.
type Sort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the attribute to sort results by.
	SortBy SortBy_Field `protobuf:"varint,1,opt,name=sort_by,json=sortBy,proto3,enum=public.portworx.common.v1.SortBy_Field" json:"sort_by,omitempty"`
	// Order of sorting to be applied on requested list.
	// If sort_by having some value and sort_order is not provided, by default ascending order will be used to sort the list.
	SortOrder SortOrder_Value `protobuf:"varint,2,opt,name=sort_order,json=sortOrder,proto3,enum=public.portworx.common.v1.SortOrder_Value" json:"sort_order,omitempty"`
}

func (x *Sort) Reset() {
	*x = Sort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_common_apiv1_sort_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sort) ProtoMessage() {}

func (x *Sort) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_common_apiv1_sort_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sort.ProtoReflect.Descriptor instead.
func (*Sort) Descriptor() ([]byte, []int) {
	return file_public_portworx_common_apiv1_sort_proto_rawDescGZIP(), []int{0}
}

func (x *Sort) GetSortBy() SortBy_Field {
	if x != nil {
		return x.SortBy
	}
	return SortBy_FIELD_UNSPECIFIED
}

func (x *Sort) GetSortOrder() SortOrder_Value {
	if x != nil {
		return x.SortOrder
	}
	return SortOrder_VALUE_UNSPECIFIED
}

// Supported fields for sorting the requested list of resources.
type SortBy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SortBy) Reset() {
	*x = SortBy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_common_apiv1_sort_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortBy) ProtoMessage() {}

func (x *SortBy) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_common_apiv1_sort_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortBy.ProtoReflect.Descriptor instead.
func (*SortBy) Descriptor() ([]byte, []int) {
	return file_public_portworx_common_apiv1_sort_proto_rawDescGZIP(), []int{1}
}

// Sort orders for the requested list of resources.
type SortOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SortOrder) Reset() {
	*x = SortOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_common_apiv1_sort_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortOrder) ProtoMessage() {}

func (x *SortOrder) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_common_apiv1_sort_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortOrder.ProtoReflect.Descriptor instead.
func (*SortOrder) Descriptor() ([]byte, []int) {
	return file_public_portworx_common_apiv1_sort_proto_rawDescGZIP(), []int{2}
}

var File_public_portworx_common_apiv1_sort_proto protoreflect.FileDescriptor

var file_public_portworx_common_apiv1_sort_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72,
	0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73,
	0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x22, 0x93, 0x01, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x40, 0x0a,
	0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x42,
	0x79, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12,
	0x49, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x09, 0x73, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x5d, 0x0a, 0x06, 0x53, 0x6f,
	0x72, 0x74, 0x42, 0x79, 0x22, 0x53, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x15, 0x0a,
	0x11, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x10, 0x02, 0x12, 0x0e,
	0x0a, 0x0a, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x10, 0x03, 0x12, 0x09,
	0x0a, 0x05, 0x50, 0x48, 0x41, 0x53, 0x45, 0x10, 0x04, 0x22, 0x3e, 0x0a, 0x09, 0x53, 0x6f, 0x72,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x01, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x02, 0x42, 0x69, 0x0a, 0x1d, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x53, 0x6f, 0x72, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x70, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x3b, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_public_portworx_common_apiv1_sort_proto_rawDescOnce sync.Once
	file_public_portworx_common_apiv1_sort_proto_rawDescData = file_public_portworx_common_apiv1_sort_proto_rawDesc
)

func file_public_portworx_common_apiv1_sort_proto_rawDescGZIP() []byte {
	file_public_portworx_common_apiv1_sort_proto_rawDescOnce.Do(func() {
		file_public_portworx_common_apiv1_sort_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_portworx_common_apiv1_sort_proto_rawDescData)
	})
	return file_public_portworx_common_apiv1_sort_proto_rawDescData
}

var file_public_portworx_common_apiv1_sort_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_public_portworx_common_apiv1_sort_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_public_portworx_common_apiv1_sort_proto_goTypes = []any{
	(SortBy_Field)(0),    // 0: public.portworx.common.v1.SortBy.Field
	(SortOrder_Value)(0), // 1: public.portworx.common.v1.SortOrder.Value
	(*Sort)(nil),         // 2: public.portworx.common.v1.Sort
	(*SortBy)(nil),       // 3: public.portworx.common.v1.SortBy
	(*SortOrder)(nil),    // 4: public.portworx.common.v1.SortOrder
}
var file_public_portworx_common_apiv1_sort_proto_depIdxs = []int32{
	0, // 0: public.portworx.common.v1.Sort.sort_by:type_name -> public.portworx.common.v1.SortBy.Field
	1, // 1: public.portworx.common.v1.Sort.sort_order:type_name -> public.portworx.common.v1.SortOrder.Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_public_portworx_common_apiv1_sort_proto_init() }
func file_public_portworx_common_apiv1_sort_proto_init() {
	if File_public_portworx_common_apiv1_sort_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_portworx_common_apiv1_sort_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Sort); i {
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
		file_public_portworx_common_apiv1_sort_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SortBy); i {
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
		file_public_portworx_common_apiv1_sort_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SortOrder); i {
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
			RawDescriptor: file_public_portworx_common_apiv1_sort_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_public_portworx_common_apiv1_sort_proto_goTypes,
		DependencyIndexes: file_public_portworx_common_apiv1_sort_proto_depIdxs,
		EnumInfos:         file_public_portworx_common_apiv1_sort_proto_enumTypes,
		MessageInfos:      file_public_portworx_common_apiv1_sort_proto_msgTypes,
	}.Build()
	File_public_portworx_common_apiv1_sort_proto = out.File
	file_public_portworx_common_apiv1_sort_proto_rawDesc = nil
	file_public_portworx_common_apiv1_sort_proto_goTypes = nil
	file_public_portworx_common_apiv1_sort_proto_depIdxs = nil
}
