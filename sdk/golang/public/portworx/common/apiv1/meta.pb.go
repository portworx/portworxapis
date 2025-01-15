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
// source: public/portworx/common/apiv1/meta.proto

package common

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Meta holds general resource metadata.
type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UID of the resource of the format <resource prefix>-<uuid>.
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// Name of the resource.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the resource.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// A string that identifies the version of this object that can be used by clients to determine when objects have changed.
	// This value must be passed unmodified back to the server by the client.
	ResourceVersion string `protobuf:"bytes,4,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty"`
	// Creation time of the object.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Update time of the object.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Labels to apply to the object.
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations for the object.
	Annotations map[string]string `protobuf:"bytes,8,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Reference to parent object of this resource.
	ParentReference *Reference `protobuf:"bytes,9,opt,name=parent_reference,json=parentReference,proto3" json:"parent_reference,omitempty"`
	// Resource names holds the mapping between the resource IDs and its display name which will be consumed by the frontend.
	ResourceNames map[string]string `protobuf:"bytes,10,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_common_apiv1_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_common_apiv1_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_public_portworx_common_apiv1_meta_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Meta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Meta) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Meta) GetResourceVersion() string {
	if x != nil {
		return x.ResourceVersion
	}
	return ""
}

func (x *Meta) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Meta) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Meta) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Meta) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Meta) GetParentReference() *Reference {
	if x != nil {
		return x.ParentReference
	}
	return nil
}

func (x *Meta) GetResourceNames() map[string]string {
	if x != nil {
		return x.ResourceNames
	}
	return nil
}

// Reference identifies the resource type, version of the uid and the resource.
type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// API group of the resource.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Version of the API.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// UID of the resource.
	Uid string `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_portworx_common_apiv1_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_public_portworx_common_apiv1_meta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_public_portworx_common_apiv1_meta_proto_rawDescGZIP(), []int{1}
}

func (x *Reference) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Reference) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Reference) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

var File_public_portworx_common_apiv1_meta_proto protoreflect.FileDescriptor

var file_public_portworx_common_apiv1_meta_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72,
	0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x12,
	0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xe2, 0x8c, 0xcf, 0xd7, 0x08, 0x02, 0x08, 0x01, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x48, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x34, 0xba, 0x48, 0x31, 0x72, 0x2f, 0x10, 0x03, 0x18, 0x3f, 0x32, 0x29, 0x5e, 0x5b, 0x61,
	0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a,
	0x30, 0x2d, 0x39, 0x5f, 0x5c, 0x2d, 0x5c, 0x20, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d,
	0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x18, 0x64, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0xcb, 0x07, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x85, 0x07, 0xba, 0x48, 0x81, 0x07, 0x9a, 0x01, 0xfd, 0x06, 0x22,
	0x89, 0x05, 0xba, 0x01, 0xff, 0x04, 0x12, 0xc1, 0x02, 0x61, 0x20, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x63,
	0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x6e,
	0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x73, 0x2c, 0x20, 0x27, 0x2d, 0x27, 0x2c, 0x20, 0x27, 0x5f, 0x27, 0x20, 0x6f, 0x72, 0x20, 0x27,
	0x2e, 0x27, 0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20,
	0x61, 0x6e, 0x20, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x20,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x20, 0x28, 0x65, 0x2e, 0x67, 0x2e, 0x20,
	0x27, 0x4d, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x27, 0x2c, 0x20, 0x20, 0x6f, 0x72, 0x20, 0x27, 0x6d,
	0x79, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x27, 0x2c, 0x20, 0x20, 0x6f, 0x72, 0x20, 0x27, 0x31, 0x32,
	0x33, 0x2d, 0x61, 0x62, 0x63, 0x27, 0x2c, 0x20, 0x72, 0x65, 0x67, 0x65, 0x78, 0x20, 0x75, 0x73,
	0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x69, 0x73, 0x20, 0x27, 0x28, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d,
	0x39, 0x5d, 0x5b, 0x2d, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x2e, 0x5d,
	0x2a, 0x29, 0x3f, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x27, 0x29,
	0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x6e, 0x20, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x20, 0x44, 0x4e, 0x53, 0x20, 0x73, 0x75, 0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x20,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x27, 0x2f, 0x27, 0x20, 0x28,
	0x65, 0x2e, 0x67, 0x2e, 0x20, 0x27, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x27, 0x29, 0x1a, 0xb8, 0x02, 0x28, 0x73, 0x69,
	0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x28, 0x27, 0x2f,
	0x27, 0x29, 0x29, 0x20, 0x3d, 0x3d, 0x20, 0x31, 0x20, 0x26, 0x26, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x28, 0x27, 0x2f, 0x27, 0x29, 0x5b, 0x30, 0x5d, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x28, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d,
	0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x5b, 0x2d, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5f, 0x2e, 0x5d, 0x2a, 0x29, 0x3f, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5d, 0x24, 0x27, 0x29, 0x29, 0x20, 0x7c, 0x7c, 0x20, 0x28, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74,
	0x68, 0x69, 0x73, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x28, 0x27, 0x2f, 0x27, 0x29, 0x29, 0x20,
	0x3d, 0x3d, 0x20, 0x32, 0x20, 0x26, 0x26, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x70, 0x6c,
	0x69, 0x74, 0x28, 0x27, 0x2f, 0x27, 0x29, 0x5b, 0x30, 0x5d, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x5b, 0x2d,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d,
	0x29, 0x3f, 0x28, 0x5c, 0x5c, 0x2e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x5b,
	0x2d, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5d, 0x29, 0x3f, 0x29, 0x2a, 0x24, 0x27, 0x29, 0x20, 0x26, 0x26, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x28, 0x27, 0x2f, 0x27, 0x29, 0x5b, 0x31, 0x5d, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x28, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d,
	0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x5b, 0x2d, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5f, 0x2e, 0x5d, 0x2a, 0x29, 0x3f, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5d, 0x24, 0x27, 0x29, 0x29, 0x72, 0x04, 0x10, 0x03, 0x18, 0x3f, 0x2a, 0xee, 0x01, 0xba, 0x01,
	0xe6, 0x01, 0x12, 0x99, 0x01, 0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65,
	0x20, 0x61, 0x6e, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x20, 0x6f, 0x72, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x20, 0x63, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x2c, 0x20, 0x27, 0x2d, 0x27, 0x2c, 0x20, 0x27, 0x5f, 0x27,
	0x20, 0x6f, 0x72, 0x20, 0x27, 0x2e, 0x27, 0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x6d, 0x75, 0x73,
	0x74, 0x20, 0x73, 0x74, 0x61, 0x72, 0x74, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x65, 0x6e, 0x64, 0x20,
	0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x6e, 0x75, 0x6d,
	0x65, 0x72, 0x69, 0x63, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x1a, 0x48,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x3d, 0x3d, 0x20, 0x27, 0x27, 0x20, 0x7c, 0x7c, 0x20, 0x74, 0x68,
	0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x28, 0x5b, 0x41,
	0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x5b, 0x2d, 0x41, 0x2d, 0x5a, 0x61, 0x2d,
	0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x2e, 0x5d, 0x2a, 0x29, 0x3f, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d,
	0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x27, 0x29, 0x72, 0x02, 0x18, 0x3f, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0xe9, 0x05, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x94, 0x05, 0xba,
	0x48, 0x90, 0x05, 0x9a, 0x01, 0x8c, 0x05, 0x22, 0x89, 0x05, 0xba, 0x01, 0xff, 0x04, 0x12, 0xc1,
	0x02, 0x61, 0x20, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x20, 0x6e, 0x61, 0x6d,
	0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x20, 0x6f,
	0x66, 0x20, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x20, 0x63,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x2c, 0x20, 0x27, 0x2d, 0x27, 0x2c, 0x20,
	0x27, 0x5f, 0x27, 0x20, 0x6f, 0x72, 0x20, 0x27, 0x2e, 0x27, 0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20,
	0x6d, 0x75, 0x73, 0x74, 0x20, 0x73, 0x74, 0x61, 0x72, 0x74, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x65,
	0x6e, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x20, 0x28, 0x65, 0x2e, 0x67, 0x2e, 0x20, 0x27, 0x4d, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x27,
	0x2c, 0x20, 0x20, 0x6f, 0x72, 0x20, 0x27, 0x6d, 0x79, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x27, 0x2c,
	0x20, 0x20, 0x6f, 0x72, 0x20, 0x27, 0x31, 0x32, 0x33, 0x2d, 0x61, 0x62, 0x63, 0x27, 0x2c, 0x20,
	0x72, 0x65, 0x67, 0x65, 0x78, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x73, 0x20, 0x27, 0x28, 0x5b,
	0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x5b, 0x2d, 0x41, 0x2d, 0x5a, 0x61,
	0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x2e, 0x5d, 0x2a, 0x29, 0x3f, 0x5b, 0x41, 0x2d, 0x5a, 0x61,
	0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x27, 0x29, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x6e,
	0x20, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x20, 0x44, 0x4e, 0x53, 0x20, 0x73, 0x75,
	0x62, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x20, 0x61,
	0x6e, 0x64, 0x20, 0x27, 0x2f, 0x27, 0x20, 0x28, 0x65, 0x2e, 0x67, 0x2e, 0x20, 0x27, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x27, 0x29, 0x1a, 0xb8, 0x02, 0x28, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e,
	0x73, 0x70, 0x6c, 0x69, 0x74, 0x28, 0x27, 0x2f, 0x27, 0x29, 0x29, 0x20, 0x3d, 0x3d, 0x20, 0x31,
	0x20, 0x26, 0x26, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x28, 0x27,
	0x2f, 0x27, 0x29, 0x5b, 0x30, 0x5d, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27,
	0x5e, 0x28, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x5b, 0x2d, 0x41,
	0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x2e, 0x5d, 0x2a, 0x29, 0x3f, 0x5b, 0x41,
	0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x27, 0x29, 0x29, 0x20, 0x7c, 0x7c,
	0x20, 0x28, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x70, 0x6c, 0x69,
	0x74, 0x28, 0x27, 0x2f, 0x27, 0x29, 0x29, 0x20, 0x3d, 0x3d, 0x20, 0x32, 0x20, 0x26, 0x26, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x28, 0x27, 0x2f, 0x27, 0x29, 0x5b,
	0x30, 0x5d, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d,
	0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x5b, 0x2d, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2a,
	0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x28, 0x5c, 0x5c, 0x2e, 0x5b, 0x61,
	0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x5b, 0x2d, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d,
	0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x29, 0x2a, 0x24, 0x27, 0x29,
	0x20, 0x26, 0x26, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x28, 0x27,
	0x2f, 0x27, 0x29, 0x5b, 0x31, 0x5d, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27,
	0x5e, 0x28, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x5b, 0x2d, 0x41,
	0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x2e, 0x5d, 0x2a, 0x29, 0x3f, 0x5b, 0x41,
	0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x27, 0x29, 0x29, 0x72, 0x04, 0x10,
	0x03, 0x18, 0x3f, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x54, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x40, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x55, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xe2, 0x8c, 0xcf,
	0xd7, 0x08, 0x02, 0x08, 0x01, 0x52, 0x03, 0x75, 0x69, 0x64, 0x42, 0x69, 0x0a, 0x1d, 0x63, 0x6f,
	0x6d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72,
	0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4d, 0x65, 0x74,
	0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x70, 0x78, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x6f, 0x72,
	0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x3b, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_public_portworx_common_apiv1_meta_proto_rawDescOnce sync.Once
	file_public_portworx_common_apiv1_meta_proto_rawDescData = file_public_portworx_common_apiv1_meta_proto_rawDesc
)

func file_public_portworx_common_apiv1_meta_proto_rawDescGZIP() []byte {
	file_public_portworx_common_apiv1_meta_proto_rawDescOnce.Do(func() {
		file_public_portworx_common_apiv1_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_portworx_common_apiv1_meta_proto_rawDescData)
	})
	return file_public_portworx_common_apiv1_meta_proto_rawDescData
}

var file_public_portworx_common_apiv1_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_public_portworx_common_apiv1_meta_proto_goTypes = []any{
	(*Meta)(nil),                  // 0: public.portworx.common.v1.Meta
	(*Reference)(nil),             // 1: public.portworx.common.v1.Reference
	nil,                           // 2: public.portworx.common.v1.Meta.LabelsEntry
	nil,                           // 3: public.portworx.common.v1.Meta.AnnotationsEntry
	nil,                           // 4: public.portworx.common.v1.Meta.ResourceNamesEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_public_portworx_common_apiv1_meta_proto_depIdxs = []int32{
	5, // 0: public.portworx.common.v1.Meta.create_time:type_name -> google.protobuf.Timestamp
	5, // 1: public.portworx.common.v1.Meta.update_time:type_name -> google.protobuf.Timestamp
	2, // 2: public.portworx.common.v1.Meta.labels:type_name -> public.portworx.common.v1.Meta.LabelsEntry
	3, // 3: public.portworx.common.v1.Meta.annotations:type_name -> public.portworx.common.v1.Meta.AnnotationsEntry
	1, // 4: public.portworx.common.v1.Meta.parent_reference:type_name -> public.portworx.common.v1.Reference
	4, // 5: public.portworx.common.v1.Meta.resource_names:type_name -> public.portworx.common.v1.Meta.ResourceNamesEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_public_portworx_common_apiv1_meta_proto_init() }
func file_public_portworx_common_apiv1_meta_proto_init() {
	if File_public_portworx_common_apiv1_meta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_portworx_common_apiv1_meta_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Meta); i {
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
		file_public_portworx_common_apiv1_meta_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Reference); i {
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
			RawDescriptor: file_public_portworx_common_apiv1_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_public_portworx_common_apiv1_meta_proto_goTypes,
		DependencyIndexes: file_public_portworx_common_apiv1_meta_proto_depIdxs,
		MessageInfos:      file_public_portworx_common_apiv1_meta_proto_msgTypes,
	}.Build()
	File_public_portworx_common_apiv1_meta_proto = out.File
	file_public_portworx_common_apiv1_meta_proto_rawDesc = nil
	file_public_portworx_common_apiv1_meta_proto_goTypes = nil
	file_public_portworx_common_apiv1_meta_proto_depIdxs = nil
}
