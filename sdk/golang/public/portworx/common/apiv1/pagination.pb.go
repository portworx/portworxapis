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
// source: public/portworx/common/apiv1/pagination.proto

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

// Request parameters for page-based pagination.
type PageBasedPaginationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Page number is the page number to return based on the size.
	PageNumber int64 `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	// Page size is the maximum number of records to include per page.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *PageBasedPaginationRequest) Reset() {
	*x = PageBasedPaginationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_common_apiv1_pagination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageBasedPaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageBasedPaginationRequest) ProtoMessage() {}

func (x *PageBasedPaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_common_apiv1_pagination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageBasedPaginationRequest.ProtoReflect.Descriptor instead.
func (*PageBasedPaginationRequest) Descriptor() ([]byte, []int) {
	return file_public_portworx_common_apiv1_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *PageBasedPaginationRequest) GetPageNumber() int64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *PageBasedPaginationRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// Metadata related to page based pagination for paginated API responses.
type PageBasedPaginationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total records without pagination available in the dataset.
	TotalRecords int64 `protobuf:"varint,1,opt,name=total_records,json=totalRecords,proto3" json:"total_records,omitempty"`
	// Current page number for this paginated response.
	CurrentPage int64 `protobuf:"varint,2,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	// Page size used for pagination.
	PageSize int64 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Total pages based on the page_size provided in the request.
	TotalPages int64 `protobuf:"varint,4,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	// Next page if available in the dataset, -1 if unavailable.
	NextPage int64 `protobuf:"varint,5,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
	// Previous page if available in the dataset, -1 if unavailable.
	PrevPage int64 `protobuf:"varint,6,opt,name=prev_page,json=prevPage,proto3" json:"prev_page,omitempty"`
}

func (x *PageBasedPaginationResponse) Reset() {
	*x = PageBasedPaginationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_common_apiv1_pagination_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageBasedPaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageBasedPaginationResponse) ProtoMessage() {}

func (x *PageBasedPaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_common_apiv1_pagination_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageBasedPaginationResponse.ProtoReflect.Descriptor instead.
func (*PageBasedPaginationResponse) Descriptor() ([]byte, []int) {
	return file_public_portworx_common_apiv1_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *PageBasedPaginationResponse) GetTotalRecords() int64 {
	if x != nil {
		return x.TotalRecords
	}
	return 0
}

func (x *PageBasedPaginationResponse) GetCurrentPage() int64 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *PageBasedPaginationResponse) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageBasedPaginationResponse) GetTotalPages() int64 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *PageBasedPaginationResponse) GetNextPage() int64 {
	if x != nil {
		return x.NextPage
	}
	return 0
}

func (x *PageBasedPaginationResponse) GetPrevPage() int64 {
	if x != nil {
		return x.PrevPage
	}
	return 0
}

var File_public_portworx_common_apiv1_pagination_proto protoreflect.FileDescriptor

var file_public_portworx_common_apiv1_pagination_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72,
	0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x5a, 0x0a, 0x1a, 0x50, 0x61,
	0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x64, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xdd, 0x01, 0x0a, 0x1b, 0x50, 0x61, 0x67, 0x65, 0x42,
	0x61, 0x73, 0x65, 0x64, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65,
	0x76, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72,
	0x65, 0x76, 0x50, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x70, 0x78, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x77,
	0x6f, 0x72, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_public_portworx_common_apiv1_pagination_proto_rawDescOnce sync.Once
	file_public_portworx_common_apiv1_pagination_proto_rawDescData = file_public_portworx_common_apiv1_pagination_proto_rawDesc
)

func file_public_portworx_common_apiv1_pagination_proto_rawDescGZIP() []byte {
	file_public_portworx_common_apiv1_pagination_proto_rawDescOnce.Do(func() {
		file_public_portworx_common_apiv1_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_portworx_common_apiv1_pagination_proto_rawDescData)
	})
	return file_public_portworx_common_apiv1_pagination_proto_rawDescData
}

var file_public_portworx_common_apiv1_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_public_portworx_common_apiv1_pagination_proto_goTypes = []any{
	(*PageBasedPaginationRequest)(nil),  // 0: public.portworx.common.v1.PageBasedPaginationRequest
	(*PageBasedPaginationResponse)(nil), // 1: public.portworx.common.v1.PageBasedPaginationResponse
}
var file_public_portworx_common_apiv1_pagination_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_public_portworx_common_apiv1_pagination_proto_init() }
func file_public_portworx_common_apiv1_pagination_proto_init() {
	if File_public_portworx_common_apiv1_pagination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_portworx_common_apiv1_pagination_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PageBasedPaginationRequest); i {
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
		file_public_portworx_common_apiv1_pagination_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PageBasedPaginationResponse); i {
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
			RawDescriptor: file_public_portworx_common_apiv1_pagination_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_public_portworx_common_apiv1_pagination_proto_goTypes,
		DependencyIndexes: file_public_portworx_common_apiv1_pagination_proto_depIdxs,
		MessageInfos:      file_public_portworx_common_apiv1_pagination_proto_msgTypes,
	}.Build()
	File_public_portworx_common_apiv1_pagination_proto = out.File
	file_public_portworx_common_apiv1_pagination_proto_rawDesc = nil
	file_public_portworx_common_apiv1_pagination_proto_goTypes = nil
	file_public_portworx_common_apiv1_pagination_proto_depIdxs = nil
}
