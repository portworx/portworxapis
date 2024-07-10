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
// source: public/portworx/pds/catalog/images/apiv1/images.proto

package image

import (
	apiv1 "github.com/portworx/portworxapis/sdk/golang/public/portworx/common/apiv1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enabled indicate either list all, only enabled or only disabled images.
type ListImagesRequest_Enabled int32

const (
	// List all images.
	ListImagesRequest_ENABLED_UNSPECIFIED ListImagesRequest_Enabled = 0
	// List only enabled images.
	ListImagesRequest_TRUE ListImagesRequest_Enabled = 1
	// List only disabled images.
	ListImagesRequest_FALSE ListImagesRequest_Enabled = 2
)

// Enum value maps for ListImagesRequest_Enabled.
var (
	ListImagesRequest_Enabled_name = map[int32]string{
		0: "ENABLED_UNSPECIFIED",
		1: "TRUE",
		2: "FALSE",
	}
	ListImagesRequest_Enabled_value = map[string]int32{
		"ENABLED_UNSPECIFIED": 0,
		"TRUE":                1,
		"FALSE":               2,
	}
)

func (x ListImagesRequest_Enabled) Enum() *ListImagesRequest_Enabled {
	p := new(ListImagesRequest_Enabled)
	*p = x
	return p
}

func (x ListImagesRequest_Enabled) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListImagesRequest_Enabled) Descriptor() protoreflect.EnumDescriptor {
	return file_public_portworx_pds_catalog_images_apiv1_images_proto_enumTypes[0].Descriptor()
}

func (ListImagesRequest_Enabled) Type() protoreflect.EnumType {
	return &file_public_portworx_pds_catalog_images_apiv1_images_proto_enumTypes[0]
}

func (x ListImagesRequest_Enabled) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListImagesRequest_Enabled.Descriptor instead.
func (ListImagesRequest_Enabled) EnumDescriptor() ([]byte, []int) {
	return file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDescGZIP(), []int{4, 0}
}

// Resource representing the data service image.
type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metadata of the resource.
	Meta *apiv1.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	// Information related to the data service image.
	Info *Info `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetMeta() *apiv1.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Image) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

// Information related to the data service image.
type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reference of the image.
	References *References `protobuf:"bytes,1,opt,name=references,proto3" json:"references,omitempty"`
	// Image registry where the image is stored.
	Registry string `protobuf:"bytes,2,opt,name=registry,proto3" json:"registry,omitempty"`
	// Image registry namespace where the image is stored.
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Tag associated with the image.
	Tag string `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	// Build version of the image.
	Build string `protobuf:"bytes,5,opt,name=build,proto3" json:"build,omitempty"`
	// Flag indicating if TLS is supported for a data service using this image.
	TlsSupport bool `protobuf:"varint,6,opt,name=tls_support,json=tlsSupport,proto3" json:"tls_support,omitempty"`
	// Capabilities associated with this image.
	Capabilities map[string]string `protobuf:"bytes,7,rep,name=capabilities,proto3" json:"capabilities,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Additional images associated with this data service image.
	AdditionalImages map[string]string `protobuf:"bytes,8,rep,name=additional_images,json=additionalImages,proto3" json:"additional_images,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDescGZIP(), []int{1}
}

func (x *Info) GetReferences() *References {
	if x != nil {
		return x.References
	}
	return nil
}

func (x *Info) GetRegistry() string {
	if x != nil {
		return x.Registry
	}
	return ""
}

func (x *Info) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Info) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Info) GetBuild() string {
	if x != nil {
		return x.Build
	}
	return ""
}

func (x *Info) GetTlsSupport() bool {
	if x != nil {
		return x.TlsSupport
	}
	return false
}

func (x *Info) GetCapabilities() map[string]string {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

func (x *Info) GetAdditionalImages() map[string]string {
	if x != nil {
		return x.AdditionalImages
	}
	return nil
}

// References to other resources.
type References struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UID of the Data service.
	DataServiceId string `protobuf:"bytes,1,opt,name=data_service_id,json=dataServiceId,proto3" json:"data_service_id,omitempty"`
	// UID of the Data service version.
	DataServiceVersionId string `protobuf:"bytes,2,opt,name=data_service_version_id,json=dataServiceVersionId,proto3" json:"data_service_version_id,omitempty"`
}

func (x *References) Reset() {
	*x = References{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *References) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*References) ProtoMessage() {}

func (x *References) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use References.ProtoReflect.Descriptor instead.
func (*References) Descriptor() ([]byte, []int) {
	return file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDescGZIP(), []int{2}
}

func (x *References) GetDataServiceId() string {
	if x != nil {
		return x.DataServiceId
	}
	return ""
}

func (x *References) GetDataServiceVersionId() string {
	if x != nil {
		return x.DataServiceVersionId
	}
	return ""
}

// Request to get the image details.
type GetImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UID of the image.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetImageRequest) Reset() {
	*x = GetImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageRequest) ProtoMessage() {}

func (x *GetImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageRequest.ProtoReflect.Descriptor instead.
func (*GetImageRequest) Descriptor() ([]byte, []int) {
	return file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDescGZIP(), []int{3}
}

func (x *GetImageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request to list images for specified data service, version.
type ListImagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Only include the latest image for each data service version id.
	Latest bool `protobuf:"varint,1,opt,name=latest,proto3" json:"latest,omitempty"`
	// UID of the data service.
	DataServiceId string `protobuf:"bytes,2,opt,name=data_service_id,json=dataServiceId,proto3" json:"data_service_id,omitempty"`
	// UID of the data service version.
	DataServiceVersionId string `protobuf:"bytes,3,opt,name=data_service_version_id,json=dataServiceVersionId,proto3" json:"data_service_version_id,omitempty"`
	// Sorting details using which requested list of data service images to be sorted.
	Sort *apiv1.Sort `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	// Pagination parameters for listing data service images.
	Pagination *apiv1.PageBasedPaginationRequest `protobuf:"bytes,5,opt,name=pagination,proto3" json:"pagination,omitempty"`
	// Filter images based on enabled flag.
	Enabled ListImagesRequest_Enabled `protobuf:"varint,7,opt,name=enabled,proto3,enum=public.portworx.pds.image.v1.ListImagesRequest_Enabled" json:"enabled,omitempty"`
}

func (x *ListImagesRequest) Reset() {
	*x = ListImagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImagesRequest) ProtoMessage() {}

func (x *ListImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListImagesRequest.ProtoReflect.Descriptor instead.
func (*ListImagesRequest) Descriptor() ([]byte, []int) {
	return file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDescGZIP(), []int{4}
}

func (x *ListImagesRequest) GetLatest() bool {
	if x != nil {
		return x.Latest
	}
	return false
}

func (x *ListImagesRequest) GetDataServiceId() string {
	if x != nil {
		return x.DataServiceId
	}
	return ""
}

func (x *ListImagesRequest) GetDataServiceVersionId() string {
	if x != nil {
		return x.DataServiceVersionId
	}
	return ""
}

func (x *ListImagesRequest) GetSort() *apiv1.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *ListImagesRequest) GetPagination() *apiv1.PageBasedPaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListImagesRequest) GetEnabled() ListImagesRequest_Enabled {
	if x != nil {
		return x.Enabled
	}
	return ListImagesRequest_ENABLED_UNSPECIFIED
}

// Response to list images request.
type ListImagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of images.
	Images []*Image `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
	// Pagination metadata for this response.
	// (-- api-linter: core::0132::response-unknown-fields=disabled
	//
	//	aip.dev/not-precedent: We need this field for pagination. --)
	Pagination *apiv1.PageBasedPaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListImagesResponse) Reset() {
	*x = ListImagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImagesResponse) ProtoMessage() {}

func (x *ListImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListImagesResponse.ProtoReflect.Descriptor instead.
func (*ListImagesResponse) Descriptor() ([]byte, []int) {
	return file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDescGZIP(), []int{5}
}

func (x *ListImagesResponse) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *ListImagesResponse) GetPagination() *apiv1.PageBasedPaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_public_portworx_pds_catalog_images_apiv1_images_proto protoreflect.FileDescriptor

var file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDesc = []byte{
	0x0a, 0x35, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72,
	0x78, 0x2f, 0x70, 0x64, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x6f, 0x72,
	0x74, 0x77, 0x6f, 0x72, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x33,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77,
	0x6f, 0x72, 0x78, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x9a, 0x04, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x48, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x74, 0x6c, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x58, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x43, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x63,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x65, 0x0a, 0x11, 0x61,
	0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x10, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x1a, 0x3f, 0x0a, 0x11, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x43, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6b, 0x0a, 0x0a, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x35,
	0x0a, 0x17, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xac, 0x03, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x35,
	0x0a, 0x17, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x5a, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x42,
	0x61, 0x73, 0x65, 0x64, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x37,
	0x0a, 0x07, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x4e, 0x41,
	0x42, 0x4c, 0x45, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x52, 0x55, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x46, 0x41, 0x4c, 0x53, 0x45, 0x10, 0x02, 0x22, 0xa9, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78,
	0x2e, 0x70, 0x64, 0x73, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72,
	0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x42, 0x61, 0x73, 0x65, 0x64, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x32, 0xa6, 0x02, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x2d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77,
	0x6f, 0x72, 0x78, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f,
	0x72, 0x78, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f,
	0x70, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x8f, 0x01, 0x0a, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x2f, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x70, 0x64, 0x73, 0x2e,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x70, 0x64, 0x73,
	0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x70, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x79, 0x0a, 0x20,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77,
	0x6f, 0x72, 0x78, 0x2e, 0x70, 0x64, 0x73, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x0b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x72, 0x65,
	0x2d, 0x70, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2f, 0x70, 0x64, 0x73, 0x2f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x3b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDescOnce sync.Once
	file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDescData = file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDesc
)

func file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDescGZIP() []byte {
	file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDescOnce.Do(func() {
		file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDescData)
	})
	return file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDescData
}

var file_public_portworx_pds_catalog_images_apiv1_images_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_public_portworx_pds_catalog_images_apiv1_images_proto_goTypes = []interface{}{
	(ListImagesRequest_Enabled)(0),           // 0: public.portworx.pds.image.v1.ListImagesRequest.Enabled
	(*Image)(nil),                            // 1: public.portworx.pds.image.v1.Image
	(*Info)(nil),                             // 2: public.portworx.pds.image.v1.Info
	(*References)(nil),                       // 3: public.portworx.pds.image.v1.References
	(*GetImageRequest)(nil),                  // 4: public.portworx.pds.image.v1.GetImageRequest
	(*ListImagesRequest)(nil),                // 5: public.portworx.pds.image.v1.ListImagesRequest
	(*ListImagesResponse)(nil),               // 6: public.portworx.pds.image.v1.ListImagesResponse
	nil,                                      // 7: public.portworx.pds.image.v1.Info.CapabilitiesEntry
	nil,                                      // 8: public.portworx.pds.image.v1.Info.AdditionalImagesEntry
	(*apiv1.Meta)(nil),                       // 9: public.portworx.common.v1.Meta
	(*apiv1.Sort)(nil),                       // 10: public.portworx.common.v1.Sort
	(*apiv1.PageBasedPaginationRequest)(nil), // 11: public.portworx.common.v1.PageBasedPaginationRequest
	(*apiv1.PageBasedPaginationResponse)(nil), // 12: public.portworx.common.v1.PageBasedPaginationResponse
}
var file_public_portworx_pds_catalog_images_apiv1_images_proto_depIdxs = []int32{
	9,  // 0: public.portworx.pds.image.v1.Image.meta:type_name -> public.portworx.common.v1.Meta
	2,  // 1: public.portworx.pds.image.v1.Image.info:type_name -> public.portworx.pds.image.v1.Info
	3,  // 2: public.portworx.pds.image.v1.Info.references:type_name -> public.portworx.pds.image.v1.References
	7,  // 3: public.portworx.pds.image.v1.Info.capabilities:type_name -> public.portworx.pds.image.v1.Info.CapabilitiesEntry
	8,  // 4: public.portworx.pds.image.v1.Info.additional_images:type_name -> public.portworx.pds.image.v1.Info.AdditionalImagesEntry
	10, // 5: public.portworx.pds.image.v1.ListImagesRequest.sort:type_name -> public.portworx.common.v1.Sort
	11, // 6: public.portworx.pds.image.v1.ListImagesRequest.pagination:type_name -> public.portworx.common.v1.PageBasedPaginationRequest
	0,  // 7: public.portworx.pds.image.v1.ListImagesRequest.enabled:type_name -> public.portworx.pds.image.v1.ListImagesRequest.Enabled
	1,  // 8: public.portworx.pds.image.v1.ListImagesResponse.images:type_name -> public.portworx.pds.image.v1.Image
	12, // 9: public.portworx.pds.image.v1.ListImagesResponse.pagination:type_name -> public.portworx.common.v1.PageBasedPaginationResponse
	4,  // 10: public.portworx.pds.image.v1.ImageService.GetImage:input_type -> public.portworx.pds.image.v1.GetImageRequest
	5,  // 11: public.portworx.pds.image.v1.ImageService.ListImages:input_type -> public.portworx.pds.image.v1.ListImagesRequest
	1,  // 12: public.portworx.pds.image.v1.ImageService.GetImage:output_type -> public.portworx.pds.image.v1.Image
	6,  // 13: public.portworx.pds.image.v1.ImageService.ListImages:output_type -> public.portworx.pds.image.v1.ListImagesResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_public_portworx_pds_catalog_images_apiv1_images_proto_init() }
func file_public_portworx_pds_catalog_images_apiv1_images_proto_init() {
	if File_public_portworx_pds_catalog_images_apiv1_images_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*References); i {
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
		file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImageRequest); i {
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
		file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListImagesRequest); i {
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
		file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListImagesResponse); i {
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
			RawDescriptor: file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_public_portworx_pds_catalog_images_apiv1_images_proto_goTypes,
		DependencyIndexes: file_public_portworx_pds_catalog_images_apiv1_images_proto_depIdxs,
		EnumInfos:         file_public_portworx_pds_catalog_images_apiv1_images_proto_enumTypes,
		MessageInfos:      file_public_portworx_pds_catalog_images_apiv1_images_proto_msgTypes,
	}.Build()
	File_public_portworx_pds_catalog_images_apiv1_images_proto = out.File
	file_public_portworx_pds_catalog_images_apiv1_images_proto_rawDesc = nil
	file_public_portworx_pds_catalog_images_apiv1_images_proto_goTypes = nil
	file_public_portworx_pds_catalog_images_apiv1_images_proto_depIdxs = nil
}
