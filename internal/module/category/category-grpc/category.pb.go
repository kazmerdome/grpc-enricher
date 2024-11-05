// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: internal/module/category/category-grpc/category.proto

package category_grpc

import (
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

type CategoryStatus int32

const (
	CategoryStatus_STATUS_ACTIVE   CategoryStatus = 0
	CategoryStatus_STATUS_PENDING  CategoryStatus = 1
	CategoryStatus_STATUS_ARCHIVED CategoryStatus = 2
)

// Enum value maps for CategoryStatus.
var (
	CategoryStatus_name = map[int32]string{
		0: "STATUS_ACTIVE",
		1: "STATUS_PENDING",
		2: "STATUS_ARCHIVED",
	}
	CategoryStatus_value = map[string]int32{
		"STATUS_ACTIVE":   0,
		"STATUS_PENDING":  1,
		"STATUS_ARCHIVED": 2,
	}
)

func (x CategoryStatus) Enum() *CategoryStatus {
	p := new(CategoryStatus)
	*p = x
	return p
}

func (x CategoryStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CategoryStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_module_category_category_grpc_category_proto_enumTypes[0].Descriptor()
}

func (CategoryStatus) Type() protoreflect.EnumType {
	return &file_internal_module_category_category_grpc_category_proto_enumTypes[0]
}

func (x CategoryStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CategoryStatus.Descriptor instead.
func (CategoryStatus) EnumDescriptor() ([]byte, []int) {
	return file_internal_module_category_category_grpc_category_proto_rawDescGZIP(), []int{0}
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Slug      string                 `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Status    CategoryStatus         `protobuf:"varint,4,opt,name=status,proto3,enum=category.CategoryStatus" json:"status,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_category_category_grpc_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_category_category_grpc_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_internal_module_category_category_grpc_category_proto_rawDescGZIP(), []int{0}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Category) GetStatus() CategoryStatus {
	if x != nil {
		return x.Status
	}
	return CategoryStatus_STATUS_ACTIVE
}

func (x *Category) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Category) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Enrich
// This is for enriching the category if has any fields to enrich
type CategoryEnrichParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnrichAllFields    *bool `protobuf:"varint,1,opt,name=enrich_all_fields,json=enrichAllFields,proto3,oneof" json:"enrich_all_fields,omitempty"`
	EnrichAllRelations *bool `protobuf:"varint,2,opt,name=enrich_all_relations,json=enrichAllRelations,proto3,oneof" json:"enrich_all_relations,omitempty"`
	Id                 *bool `protobuf:"varint,3,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name               *bool `protobuf:"varint,4,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Slug               *bool `protobuf:"varint,5,opt,name=slug,proto3,oneof" json:"slug,omitempty"`
	Status             *bool `protobuf:"varint,6,opt,name=status,proto3,oneof" json:"status,omitempty"`
	CreatedAt          *bool `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt          *bool `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
}

func (x *CategoryEnrichParams) Reset() {
	*x = CategoryEnrichParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_module_category_category_grpc_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryEnrichParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryEnrichParams) ProtoMessage() {}

func (x *CategoryEnrichParams) ProtoReflect() protoreflect.Message {
	mi := &file_internal_module_category_category_grpc_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryEnrichParams.ProtoReflect.Descriptor instead.
func (*CategoryEnrichParams) Descriptor() ([]byte, []int) {
	return file_internal_module_category_category_grpc_category_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryEnrichParams) GetEnrichAllFields() bool {
	if x != nil && x.EnrichAllFields != nil {
		return *x.EnrichAllFields
	}
	return false
}

func (x *CategoryEnrichParams) GetEnrichAllRelations() bool {
	if x != nil && x.EnrichAllRelations != nil {
		return *x.EnrichAllRelations
	}
	return false
}

func (x *CategoryEnrichParams) GetId() bool {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return false
}

func (x *CategoryEnrichParams) GetName() bool {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return false
}

func (x *CategoryEnrichParams) GetSlug() bool {
	if x != nil && x.Slug != nil {
		return *x.Slug
	}
	return false
}

func (x *CategoryEnrichParams) GetStatus() bool {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return false
}

func (x *CategoryEnrichParams) GetCreatedAt() bool {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return false
}

func (x *CategoryEnrichParams) GetUpdatedAt() bool {
	if x != nil && x.UpdatedAt != nil {
		return *x.UpdatedAt
	}
	return false
}

var File_internal_module_category_category_grpc_category_proto protoreflect.FileDescriptor

var file_internal_module_category_category_grpc_category_proto_rawDesc = []byte{
	0x0a, 0x35, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xea, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x9b, 0x03, 0x0a, 0x14, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x45, 0x6e, 0x72, 0x69,
	0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2f, 0x0a, 0x11, 0x65, 0x6e, 0x72, 0x69,
	0x63, 0x68, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x41, 0x6c, 0x6c,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x14, 0x65, 0x6e, 0x72,
	0x69, 0x63, 0x68, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x12, 0x65, 0x6e, 0x72, 0x69, 0x63,
	0x68, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x02,
	0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17,
	0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x04,
	0x73, 0x6c, 0x75, 0x67, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x06, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x48, 0x07, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x14, 0x0a, 0x12,
	0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x5f, 0x61, 0x6c,
	0x6c, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x5f,
	0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x73, 0x6c, 0x75, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x2a, 0x4c, 0x0a,
	0x0e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x45, 0x44, 0x10, 0x02, 0x42, 0x47, 0x5a, 0x45, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x7a, 0x6d, 0x65, 0x72,
	0x64, 0x6f, 0x6d, 0x65, 0x2f, 0x65, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_module_category_category_grpc_category_proto_rawDescOnce sync.Once
	file_internal_module_category_category_grpc_category_proto_rawDescData = file_internal_module_category_category_grpc_category_proto_rawDesc
)

func file_internal_module_category_category_grpc_category_proto_rawDescGZIP() []byte {
	file_internal_module_category_category_grpc_category_proto_rawDescOnce.Do(func() {
		file_internal_module_category_category_grpc_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_module_category_category_grpc_category_proto_rawDescData)
	})
	return file_internal_module_category_category_grpc_category_proto_rawDescData
}

var file_internal_module_category_category_grpc_category_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_module_category_category_grpc_category_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_module_category_category_grpc_category_proto_goTypes = []any{
	(CategoryStatus)(0),           // 0: category.CategoryStatus
	(*Category)(nil),              // 1: category.Category
	(*CategoryEnrichParams)(nil),  // 2: category.CategoryEnrichParams
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_internal_module_category_category_grpc_category_proto_depIdxs = []int32{
	0, // 0: category.Category.status:type_name -> category.CategoryStatus
	3, // 1: category.Category.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: category.Category.updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_module_category_category_grpc_category_proto_init() }
func file_internal_module_category_category_grpc_category_proto_init() {
	if File_internal_module_category_category_grpc_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_module_category_category_grpc_category_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Category); i {
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
		file_internal_module_category_category_grpc_category_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CategoryEnrichParams); i {
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
	file_internal_module_category_category_grpc_category_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_module_category_category_grpc_category_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_module_category_category_grpc_category_proto_goTypes,
		DependencyIndexes: file_internal_module_category_category_grpc_category_proto_depIdxs,
		EnumInfos:         file_internal_module_category_category_grpc_category_proto_enumTypes,
		MessageInfos:      file_internal_module_category_category_grpc_category_proto_msgTypes,
	}.Build()
	File_internal_module_category_category_grpc_category_proto = out.File
	file_internal_module_category_category_grpc_category_proto_rawDesc = nil
	file_internal_module_category_category_grpc_category_proto_goTypes = nil
	file_internal_module_category_category_grpc_category_proto_depIdxs = nil
}
