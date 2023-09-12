// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: gotocompany/optimus/core/v1beta1/project.proto

package optimus

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type RegisterProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *ProjectSpecification `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *RegisterProjectRequest) Reset() {
	*x = RegisterProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterProjectRequest) ProtoMessage() {}

func (x *RegisterProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterProjectRequest.ProtoReflect.Descriptor instead.
func (*RegisterProjectRequest) Descriptor() ([]byte, []int) {
	return file_gotocompany_optimus_core_v1beta1_project_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterProjectRequest) GetProject() *ProjectSpecification {
	if x != nil {
		return x.Project
	}
	return nil
}

type RegisterProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterProjectResponse) Reset() {
	*x = RegisterProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterProjectResponse) ProtoMessage() {}

func (x *RegisterProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterProjectResponse.ProtoReflect.Descriptor instead.
func (*RegisterProjectResponse) Descriptor() ([]byte, []int) {
	return file_gotocompany_optimus_core_v1beta1_project_proto_rawDescGZIP(), []int{1}
}

type ListProjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProjectsRequest) Reset() {
	*x = ListProjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsRequest) ProtoMessage() {}

func (x *ListProjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsRequest.ProtoReflect.Descriptor instead.
func (*ListProjectsRequest) Descriptor() ([]byte, []int) {
	return file_gotocompany_optimus_core_v1beta1_project_proto_rawDescGZIP(), []int{2}
}

type ListProjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Projects []*ProjectSpecification `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *ListProjectsResponse) Reset() {
	*x = ListProjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsResponse) ProtoMessage() {}

func (x *ListProjectsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsResponse.ProtoReflect.Descriptor instead.
func (*ListProjectsResponse) Descriptor() ([]byte, []int) {
	return file_gotocompany_optimus_core_v1beta1_project_proto_rawDescGZIP(), []int{3}
}

func (x *ListProjectsResponse) GetProjects() []*ProjectSpecification {
	if x != nil {
		return x.Projects
	}
	return nil
}

type GetProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName string `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
}

func (x *GetProjectRequest) Reset() {
	*x = GetProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectRequest) ProtoMessage() {}

func (x *GetProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectRequest.ProtoReflect.Descriptor instead.
func (*GetProjectRequest) Descriptor() ([]byte, []int) {
	return file_gotocompany_optimus_core_v1beta1_project_proto_rawDescGZIP(), []int{4}
}

func (x *GetProjectRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

type GetProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project *ProjectSpecification `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *GetProjectResponse) Reset() {
	*x = GetProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectResponse) ProtoMessage() {}

func (x *GetProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectResponse.ProtoReflect.Descriptor instead.
func (*GetProjectResponse) Descriptor() ([]byte, []int) {
	return file_gotocompany_optimus_core_v1beta1_project_proto_rawDescGZIP(), []int{5}
}

func (x *GetProjectResponse) GetProject() *ProjectSpecification {
	if x != nil {
		return x.Project
	}
	return nil
}

type ProjectSpecification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string                                         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Config  map[string]string                              `protobuf:"bytes,2,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Presets map[string]*ProjectSpecification_ProjectPreset `protobuf:"bytes,4,rep,name=presets,proto3" json:"presets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ProjectSpecification) Reset() {
	*x = ProjectSpecification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectSpecification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectSpecification) ProtoMessage() {}

func (x *ProjectSpecification) ProtoReflect() protoreflect.Message {
	mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectSpecification.ProtoReflect.Descriptor instead.
func (*ProjectSpecification) Descriptor() ([]byte, []int) {
	return file_gotocompany_optimus_core_v1beta1_project_proto_rawDescGZIP(), []int{6}
}

func (x *ProjectSpecification) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectSpecification) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ProjectSpecification) GetPresets() map[string]*ProjectSpecification_ProjectPreset {
	if x != nil {
		return x.Presets
	}
	return nil
}

type ProjectSpecification_ProjectPreset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	TruncateTo  string `protobuf:"bytes,3,opt,name=truncate_to,json=truncateTo,proto3" json:"truncate_to,omitempty"`
	Offset      string `protobuf:"bytes,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Size        string `protobuf:"bytes,5,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ProjectSpecification_ProjectPreset) Reset() {
	*x = ProjectSpecification_ProjectPreset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectSpecification_ProjectPreset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectSpecification_ProjectPreset) ProtoMessage() {}

func (x *ProjectSpecification_ProjectPreset) ProtoReflect() protoreflect.Message {
	mi := &file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectSpecification_ProjectPreset.ProtoReflect.Descriptor instead.
func (*ProjectSpecification_ProjectPreset) Descriptor() ([]byte, []int) {
	return file_gotocompany_optimus_core_v1beta1_project_proto_rawDescGZIP(), []int{6, 1}
}

func (x *ProjectSpecification_ProjectPreset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectSpecification_ProjectPreset) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProjectSpecification_ProjectPreset) GetTruncateTo() string {
	if x != nil {
		return x.TruncateTo
	}
	return ""
}

func (x *ProjectSpecification_ProjectPreset) GetOffset() string {
	if x != nil {
		return x.Offset
	}
	return ""
}

func (x *ProjectSpecification_ProjectPreset) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

var File_gotocompany_optimus_core_v1beta1_project_proto protoreflect.FileDescriptor

var file_gotocompany_optimus_core_v1beta1_project_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6d, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x67, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6d, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x70, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75,
	0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4a, 0x04, 0x08, 0x02,
	0x10, 0x03, 0x22, 0x25, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x04, 0x08,
	0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x6a, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x36, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x66, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75,
	0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0xbe, 0x04, 0x0a,
	0x14, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x5a, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5d, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x92, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x75, 0x6e, 0x63,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72,
	0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x1a, 0x80, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x5a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x32, 0xf3, 0x03,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xa3, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x38, 0x2e, 0x67, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39,
	0x2e, 0x67, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6d, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x97, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x35, 0x2e, 0x67, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36,
	0x2e, 0x67, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x6d, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0xa0, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x33, 0x2e, 0x67, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6d, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x67, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x21, 0x12, 0x1f, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x42, 0x97, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x6f,
	0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x42, 0x15, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x01, 0x5a,
	0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x92,
	0x41, 0x3b, 0x12, 0x05, 0x32, 0x03, 0x30, 0x2e, 0x31, 0x1a, 0x0e, 0x31, 0x32, 0x37, 0x2e, 0x30,
	0x2e, 0x30, 0x2e, 0x31, 0x3a, 0x39, 0x31, 0x30, 0x30, 0x22, 0x04, 0x2f, 0x61, 0x70, 0x69, 0x2a,
	0x01, 0x01, 0x72, 0x19, 0x0a, 0x17, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x20, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gotocompany_optimus_core_v1beta1_project_proto_rawDescOnce sync.Once
	file_gotocompany_optimus_core_v1beta1_project_proto_rawDescData = file_gotocompany_optimus_core_v1beta1_project_proto_rawDesc
)

func file_gotocompany_optimus_core_v1beta1_project_proto_rawDescGZIP() []byte {
	file_gotocompany_optimus_core_v1beta1_project_proto_rawDescOnce.Do(func() {
		file_gotocompany_optimus_core_v1beta1_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_gotocompany_optimus_core_v1beta1_project_proto_rawDescData)
	})
	return file_gotocompany_optimus_core_v1beta1_project_proto_rawDescData
}

var file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_gotocompany_optimus_core_v1beta1_project_proto_goTypes = []interface{}{
	(*RegisterProjectRequest)(nil),             // 0: gotocompany.optimus.core.v1beta1.RegisterProjectRequest
	(*RegisterProjectResponse)(nil),            // 1: gotocompany.optimus.core.v1beta1.RegisterProjectResponse
	(*ListProjectsRequest)(nil),                // 2: gotocompany.optimus.core.v1beta1.ListProjectsRequest
	(*ListProjectsResponse)(nil),               // 3: gotocompany.optimus.core.v1beta1.ListProjectsResponse
	(*GetProjectRequest)(nil),                  // 4: gotocompany.optimus.core.v1beta1.GetProjectRequest
	(*GetProjectResponse)(nil),                 // 5: gotocompany.optimus.core.v1beta1.GetProjectResponse
	(*ProjectSpecification)(nil),               // 6: gotocompany.optimus.core.v1beta1.ProjectSpecification
	nil,                                        // 7: gotocompany.optimus.core.v1beta1.ProjectSpecification.ConfigEntry
	(*ProjectSpecification_ProjectPreset)(nil), // 8: gotocompany.optimus.core.v1beta1.ProjectSpecification.ProjectPreset
	nil, // 9: gotocompany.optimus.core.v1beta1.ProjectSpecification.PresetsEntry
}
var file_gotocompany_optimus_core_v1beta1_project_proto_depIdxs = []int32{
	6, // 0: gotocompany.optimus.core.v1beta1.RegisterProjectRequest.project:type_name -> gotocompany.optimus.core.v1beta1.ProjectSpecification
	6, // 1: gotocompany.optimus.core.v1beta1.ListProjectsResponse.projects:type_name -> gotocompany.optimus.core.v1beta1.ProjectSpecification
	6, // 2: gotocompany.optimus.core.v1beta1.GetProjectResponse.project:type_name -> gotocompany.optimus.core.v1beta1.ProjectSpecification
	7, // 3: gotocompany.optimus.core.v1beta1.ProjectSpecification.config:type_name -> gotocompany.optimus.core.v1beta1.ProjectSpecification.ConfigEntry
	9, // 4: gotocompany.optimus.core.v1beta1.ProjectSpecification.presets:type_name -> gotocompany.optimus.core.v1beta1.ProjectSpecification.PresetsEntry
	8, // 5: gotocompany.optimus.core.v1beta1.ProjectSpecification.PresetsEntry.value:type_name -> gotocompany.optimus.core.v1beta1.ProjectSpecification.ProjectPreset
	0, // 6: gotocompany.optimus.core.v1beta1.ProjectService.RegisterProject:input_type -> gotocompany.optimus.core.v1beta1.RegisterProjectRequest
	2, // 7: gotocompany.optimus.core.v1beta1.ProjectService.ListProjects:input_type -> gotocompany.optimus.core.v1beta1.ListProjectsRequest
	4, // 8: gotocompany.optimus.core.v1beta1.ProjectService.GetProject:input_type -> gotocompany.optimus.core.v1beta1.GetProjectRequest
	1, // 9: gotocompany.optimus.core.v1beta1.ProjectService.RegisterProject:output_type -> gotocompany.optimus.core.v1beta1.RegisterProjectResponse
	3, // 10: gotocompany.optimus.core.v1beta1.ProjectService.ListProjects:output_type -> gotocompany.optimus.core.v1beta1.ListProjectsResponse
	5, // 11: gotocompany.optimus.core.v1beta1.ProjectService.GetProject:output_type -> gotocompany.optimus.core.v1beta1.GetProjectResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_gotocompany_optimus_core_v1beta1_project_proto_init() }
func file_gotocompany_optimus_core_v1beta1_project_proto_init() {
	if File_gotocompany_optimus_core_v1beta1_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterProjectRequest); i {
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
		file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterProjectResponse); i {
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
		file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsRequest); i {
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
		file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsResponse); i {
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
		file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectRequest); i {
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
		file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectResponse); i {
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
		file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectSpecification); i {
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
		file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectSpecification_ProjectPreset); i {
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
			RawDescriptor: file_gotocompany_optimus_core_v1beta1_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gotocompany_optimus_core_v1beta1_project_proto_goTypes,
		DependencyIndexes: file_gotocompany_optimus_core_v1beta1_project_proto_depIdxs,
		MessageInfos:      file_gotocompany_optimus_core_v1beta1_project_proto_msgTypes,
	}.Build()
	File_gotocompany_optimus_core_v1beta1_project_proto = out.File
	file_gotocompany_optimus_core_v1beta1_project_proto_rawDesc = nil
	file_gotocompany_optimus_core_v1beta1_project_proto_goTypes = nil
	file_gotocompany_optimus_core_v1beta1_project_proto_depIdxs = nil
}
