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
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: public/portworx/pds/common/apiv1/pdsapplicationresource.proto

package common

import (
	apiv1 "github.com/portworx/portworxapis/sdk/golang/public/portworx/common/apiv1"
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

// Type of the resource can be associated to the project.
type PDSApplicationResource_Type int32

const (
	// Unspecified, do not use.
	PDSApplicationResource_TYPE_UNSPECIFIED PDSApplicationResource_Type = 0
	// List of supported PDS application resources.
	// data service resource.
	PDSApplicationResource_DATA_SERVICE PDSApplicationResource_Type = 1
)

// Enum value maps for PDSApplicationResource_Type.
var (
	PDSApplicationResource_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "DATA_SERVICE",
	}
	PDSApplicationResource_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"DATA_SERVICE":     1,
	}
)

func (x PDSApplicationResource_Type) Enum() *PDSApplicationResource_Type {
	p := new(PDSApplicationResource_Type)
	*p = x
	return p
}

func (x PDSApplicationResource_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PDSApplicationResource_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_enumTypes[0].Descriptor()
}

func (PDSApplicationResource_Type) Type() protoreflect.EnumType {
	return &file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_enumTypes[0]
}

func (x PDSApplicationResource_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PDSApplicationResource_Type.Descriptor instead.
func (PDSApplicationResource_Type) EnumDescriptor() ([]byte, []int) {
	return file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_rawDescGZIP(), []int{0, 0}
}

// The resource type of Data Service
type PDSApplicationResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PDSApplicationResource) Reset() {
	*x = PDSApplicationResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PDSApplicationResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PDSApplicationResource) ProtoMessage() {}

func (x *PDSApplicationResource) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PDSApplicationResource.ProtoReflect.Descriptor instead.
func (*PDSApplicationResource) Descriptor() ([]byte, []int) {
	return file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_rawDescGZIP(), []int{0}
}

// PDSApplicationSelector is used to query deployments using the associated application reesources.
type PDSApplicationSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Application_filters is the list of all filters that should be applied to fetch data related to deployment.
	// Each filter will have AND relationship.
	ApplicationFilters []*PDSApplicationSelector_PDSApplicationFilter `protobuf:"bytes,1,rep,name=application_filters,json=applicationFilters,proto3" json:"application_filters,omitempty"`
}

func (x *PDSApplicationSelector) Reset() {
	*x = PDSApplicationSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PDSApplicationSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PDSApplicationSelector) ProtoMessage() {}

func (x *PDSApplicationSelector) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PDSApplicationSelector.ProtoReflect.Descriptor instead.
func (*PDSApplicationSelector) Descriptor() ([]byte, []int) {
	return file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_rawDescGZIP(), []int{1}
}

func (x *PDSApplicationSelector) GetApplicationFilters() []*PDSApplicationSelector_PDSApplicationFilter {
	if x != nil {
		return x.ApplicationFilters
	}
	return nil
}

// PDSApplicationFilter is filter for a given resource type.
type PDSApplicationSelector_PDSApplicationFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key of key,value pair against which filtering needs to be performs based on associated data service.
	ResourceType PDSApplicationResource_Type `protobuf:"varint,101,opt,name=resource_type,json=resourceType,proto3,enum=public.portworx.pds.common.v1.PDSApplicationResource_Type" json:"resource_type,omitempty"`
	// Op provides the relationship between the key,value pair in the resp element(s).
	Op apiv1.Selector_Operator `protobuf:"varint,102,opt,name=op,proto3,enum=public.portworx.common.v1.Selector_Operator" json:"op,omitempty"`
	// Value of key,value pair against which filtering needs to be performs.
	Values []string `protobuf:"bytes,103,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *PDSApplicationSelector_PDSApplicationFilter) Reset() {
	*x = PDSApplicationSelector_PDSApplicationFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PDSApplicationSelector_PDSApplicationFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PDSApplicationSelector_PDSApplicationFilter) ProtoMessage() {}

func (x *PDSApplicationSelector_PDSApplicationFilter) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PDSApplicationSelector_PDSApplicationFilter.ProtoReflect.Descriptor instead.
func (*PDSApplicationSelector_PDSApplicationFilter) Descriptor() ([]byte, []int) {
	return file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PDSApplicationSelector_PDSApplicationFilter) GetResourceType() PDSApplicationResource_Type {
	if x != nil {
		return x.ResourceType
	}
	return PDSApplicationResource_TYPE_UNSPECIFIED
}

func (x *PDSApplicationSelector_PDSApplicationFilter) GetOp() apiv1.Selector_Operator {
	if x != nil {
		return x.Op
	}
	return apiv1.Selector_Operator(0)
}

func (x *PDSApplicationSelector_PDSApplicationFilter) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_public_portworx_pds_common_apiv1_pdsapplicationresource_proto protoreflect.FileDescriptor

var file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72,
	0x78, 0x2f, 0x70, 0x64, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x2f, 0x70, 0x64, 0x73, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78,
	0x2e, 0x70, 0x64, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x2b,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x16, 0x50,
	0x44, 0x53, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x2e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x43, 0x45, 0x10, 0x01, 0x22, 0xe5, 0x02, 0x0a, 0x16, 0x50, 0x44, 0x53, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x7b, 0x0a, 0x13, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4a, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e,
	0x70, 0x64, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x44,
	0x53, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x44, 0x53, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x12, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0xcd, 0x01,
	0x0a, 0x14, 0x50, 0x44, 0x53, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x5f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e,
	0x70, 0x64, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x44,
	0x53, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x66, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x67, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x83, 0x01,
	0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x42, 0x1b, 0x50, 0x44, 0x53, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x75, 0x72, 0x65, 0x2d, 0x70, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2f, 0x70, 0x64, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_rawDescOnce sync.Once
	file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_rawDescData = file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_rawDesc
)

func file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_rawDescGZIP() []byte {
	file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_rawDescOnce.Do(func() {
		file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_rawDescData)
	})
	return file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_rawDescData
}

var file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_goTypes = []interface{}{
	(PDSApplicationResource_Type)(0),                    // 0: public.portworx.pds.common.v1.PDSApplicationResource.Type
	(*PDSApplicationResource)(nil),                      // 1: public.portworx.pds.common.v1.PDSApplicationResource
	(*PDSApplicationSelector)(nil),                      // 2: public.portworx.pds.common.v1.PDSApplicationSelector
	(*PDSApplicationSelector_PDSApplicationFilter)(nil), // 3: public.portworx.pds.common.v1.PDSApplicationSelector.PDSApplicationFilter
	(apiv1.Selector_Operator)(0),                        // 4: public.portworx.common.v1.Selector.Operator
}
var file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_depIdxs = []int32{
	3, // 0: public.portworx.pds.common.v1.PDSApplicationSelector.application_filters:type_name -> public.portworx.pds.common.v1.PDSApplicationSelector.PDSApplicationFilter
	0, // 1: public.portworx.pds.common.v1.PDSApplicationSelector.PDSApplicationFilter.resource_type:type_name -> public.portworx.pds.common.v1.PDSApplicationResource.Type
	4, // 2: public.portworx.pds.common.v1.PDSApplicationSelector.PDSApplicationFilter.op:type_name -> public.portworx.common.v1.Selector.Operator
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_init() }
func file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_init() {
	if File_public_portworx_pds_common_apiv1_pdsapplicationresource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PDSApplicationResource); i {
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
		file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PDSApplicationSelector); i {
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
		file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PDSApplicationSelector_PDSApplicationFilter); i {
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
			RawDescriptor: file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_goTypes,
		DependencyIndexes: file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_depIdxs,
		EnumInfos:         file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_enumTypes,
		MessageInfos:      file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_msgTypes,
	}.Build()
	File_public_portworx_pds_common_apiv1_pdsapplicationresource_proto = out.File
	file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_rawDesc = nil
	file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_goTypes = nil
	file_public_portworx_pds_common_apiv1_pdsapplicationresource_proto_depIdxs = nil
}