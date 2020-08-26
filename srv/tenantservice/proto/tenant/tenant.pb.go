// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: proto/tenant/tenant.proto

package srv_tenant

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Tenant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: gorm:"primary_key"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	// @inject_tag: gorm:"type:varchar(255);not null;comment:'名称'"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" gorm:"type:varchar(255);not null;comment:'名称'"`
	// @inject_tag: gorm:"type:varchar(100);not null;comment:'助记码'"
	MnemonicCode string `protobuf:"bytes,3,opt,name=mnemonic_code,json=mnemonicCode,proto3" json:"mnemonic_code,omitempty" gorm:"type:varchar(100);not null;comment:'助记码'"`
	// @inject_tag: gorm:"type:varchar(100);not null;comment:'联系人'"
	Contact string `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty" gorm:"type:varchar(100);not null;comment:'联系人'"`
	// @inject_tag: gorm:"type:varchar(100);not null;comment:'联系人电话'"
	ContactPhone string `protobuf:"bytes,5,opt,name=contact_phone,json=contactPhone,proto3" json:"contact_phone,omitempty" gorm:"type:varchar(100);not null;comment:'联系人电话'"`
	// @inject_tag: gorm:"type:varchar(100);not null;comment:'联系地址'"
	ContactAddress string `protobuf:"bytes,6,opt,name=contact_address,json=contactAddress,proto3" json:"contact_address,omitempty" gorm:"type:varchar(100);not null;comment:'联系地址'"`
	// @inject_tag: gorm:"type:int(5);not null;comment:'规模'"
	Size int32 `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty" gorm:"type:int(5);not null;comment:'规模'"`
	// @inject_tag: gorm:"type:int(11);not null;comment:'管理员'"
	AdministratorId int32 `protobuf:"varint,8,opt,name=administrator_id,json=administratorId,proto3" json:"administrator_id,omitempty" gorm:"type:int(11);not null;comment:'管理员'"`
	// @inject_tag: gorm:"type:tinyint(1);not null;default:1;comment:'启用'"
	Enable bool `protobuf:"varint,9,opt,name=enable,proto3" json:"enable,omitempty" gorm:"type:tinyint(1);not null;default:1;comment:'启用'"`
	// @inject_tag: gorm:"type:date;default:null;comment:'合约开始时间'"
	ContractStart string `protobuf:"bytes,10,opt,name=contract_start,json=contractStart,proto3" json:"contract_start,omitempty" gorm:"type:date;default:null;comment:'合约开始时间'"`
	// @inject_tag: gorm:"type:date;default:null;comment:'合约结束时间'"
	ContractEnd string `protobuf:"bytes,11,opt,name=contract_end,json=contractEnd,proto3" json:"contract_end,omitempty" gorm:"type:date;default:null;comment:'合约结束时间'"`
	// @inject_tag: gorm:"type:datetime;default:null"
	CreatedAt string `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" gorm:"type:datetime;default:null"`
	// @inject_tag: gorm:"type:datetime;default:null"
	UpdatedAt string `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" gorm:"type:datetime;default:null"`
	// @inject_tag: gorm:"type:datetime;default:null"
	DeletedAt string `protobuf:"bytes,14,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty" gorm:"type:datetime;default:null"`
}

func (x *Tenant) Reset() {
	*x = Tenant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tenant_tenant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tenant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tenant) ProtoMessage() {}

func (x *Tenant) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tenant_tenant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tenant.ProtoReflect.Descriptor instead.
func (*Tenant) Descriptor() ([]byte, []int) {
	return file_proto_tenant_tenant_proto_rawDescGZIP(), []int{0}
}

func (x *Tenant) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tenant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tenant) GetMnemonicCode() string {
	if x != nil {
		return x.MnemonicCode
	}
	return ""
}

func (x *Tenant) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *Tenant) GetContactPhone() string {
	if x != nil {
		return x.ContactPhone
	}
	return ""
}

func (x *Tenant) GetContactAddress() string {
	if x != nil {
		return x.ContactAddress
	}
	return ""
}

func (x *Tenant) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Tenant) GetAdministratorId() int32 {
	if x != nil {
		return x.AdministratorId
	}
	return 0
}

func (x *Tenant) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *Tenant) GetContractStart() string {
	if x != nil {
		return x.ContractStart
	}
	return ""
}

func (x *Tenant) GetContractEnd() string {
	if x != nil {
		return x.ContractEnd
	}
	return ""
}

func (x *Tenant) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Tenant) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Tenant) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type GetTenantPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: form:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" form:"name"`
	// @inject_tag: form:"current_page"
	CurrentPage int32 `protobuf:"varint,2,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty" form:"current_page"`
	// @inject_tag: form:"page_size"
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" form:"page_size"`
}

func (x *GetTenantPageRequest) Reset() {
	*x = GetTenantPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tenant_tenant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantPageRequest) ProtoMessage() {}

func (x *GetTenantPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tenant_tenant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantPageRequest.ProtoReflect.Descriptor instead.
func (*GetTenantPageRequest) Descriptor() ([]byte, []int) {
	return file_proto_tenant_tenant_proto_rawDescGZIP(), []int{1}
}

func (x *GetTenantPageRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetTenantPageRequest) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *GetTenantPageRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetTenantInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTenantInfoRequest) Reset() {
	*x = GetTenantInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tenant_tenant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantInfoRequest) ProtoMessage() {}

func (x *GetTenantInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tenant_tenant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantInfoRequest.ProtoReflect.Descriptor instead.
func (*GetTenantInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_tenant_tenant_proto_rawDescGZIP(), []int{2}
}

func (x *GetTenantInfoRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTenantInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Tenant `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetTenantInfoResponse) Reset() {
	*x = GetTenantInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tenant_tenant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantInfoResponse) ProtoMessage() {}

func (x *GetTenantInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tenant_tenant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantInfoResponse.ProtoReflect.Descriptor instead.
func (*GetTenantInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_tenant_tenant_proto_rawDescGZIP(), []int{3}
}

func (x *GetTenantInfoResponse) GetData() *Tenant {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetTenantPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*Tenant `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Total int32     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetTenantPageResponse) Reset() {
	*x = GetTenantPageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tenant_tenant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantPageResponse) ProtoMessage() {}

func (x *GetTenantPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tenant_tenant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantPageResponse.ProtoReflect.Descriptor instead.
func (*GetTenantPageResponse) Descriptor() ([]byte, []int) {
	return file_proto_tenant_tenant_proto_rawDescGZIP(), []int{4}
}

func (x *GetTenantPageResponse) GetData() []*Tenant {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetTenantPageResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type StoreTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Form *Tenant `protobuf:"bytes,1,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *StoreTenantRequest) Reset() {
	*x = StoreTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tenant_tenant_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreTenantRequest) ProtoMessage() {}

func (x *StoreTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tenant_tenant_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreTenantRequest.ProtoReflect.Descriptor instead.
func (*StoreTenantRequest) Descriptor() ([]byte, []int) {
	return file_proto_tenant_tenant_proto_rawDescGZIP(), []int{5}
}

func (x *StoreTenantRequest) GetForm() *Tenant {
	if x != nil {
		return x.Form
	}
	return nil
}

type StoreTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Tenant `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StoreTenantResponse) Reset() {
	*x = StoreTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tenant_tenant_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreTenantResponse) ProtoMessage() {}

func (x *StoreTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tenant_tenant_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreTenantResponse.ProtoReflect.Descriptor instead.
func (*StoreTenantResponse) Descriptor() ([]byte, []int) {
	return file_proto_tenant_tenant_proto_rawDescGZIP(), []int{6}
}

func (x *StoreTenantResponse) GetData() *Tenant {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Form *Tenant `protobuf:"bytes,1,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *UpdateTenantRequest) Reset() {
	*x = UpdateTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tenant_tenant_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantRequest) ProtoMessage() {}

func (x *UpdateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tenant_tenant_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantRequest.ProtoReflect.Descriptor instead.
func (*UpdateTenantRequest) Descriptor() ([]byte, []int) {
	return file_proto_tenant_tenant_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTenantRequest) GetForm() *Tenant {
	if x != nil {
		return x.Form
	}
	return nil
}

type UpdateTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Tenant `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateTenantResponse) Reset() {
	*x = UpdateTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tenant_tenant_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantResponse) ProtoMessage() {}

func (x *UpdateTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tenant_tenant_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantResponse.ProtoReflect.Descriptor instead.
func (*UpdateTenantResponse) Descriptor() ([]byte, []int) {
	return file_proto_tenant_tenant_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateTenantResponse) GetData() *Tenant {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_tenant_tenant_proto protoreflect.FileDescriptor

var file_proto_tenant_tenant_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x72, 0x76,
	0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0xb7, 0x03, 0x0a, 0x06, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e,
	0x69, 0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d,
	0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x65, 0x6e,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x45, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x6a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x26, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x72, 0x76, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x55, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x73, 0x72, 0x76, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x3c, 0x0a,
	0x12, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x3d, 0x0a, 0x13, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3d, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x3e, 0x0a, 0x14, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xe6, 0x02, 0x0a, 0x0d, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x2e, 0x73,
	0x72, 0x76, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x73, 0x72, 0x76, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0b, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x73, 0x72, 0x76,
	0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x72, 0x76,
	0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a,
	0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1f, 0x2e,
	0x73, 0x72, 0x76, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x73, 0x72, 0x76, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_tenant_tenant_proto_rawDescOnce sync.Once
	file_proto_tenant_tenant_proto_rawDescData = file_proto_tenant_tenant_proto_rawDesc
)

func file_proto_tenant_tenant_proto_rawDescGZIP() []byte {
	file_proto_tenant_tenant_proto_rawDescOnce.Do(func() {
		file_proto_tenant_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_tenant_tenant_proto_rawDescData)
	})
	return file_proto_tenant_tenant_proto_rawDescData
}

var file_proto_tenant_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_tenant_tenant_proto_goTypes = []interface{}{
	(*Tenant)(nil),                // 0: srv.tenant.Tenant
	(*GetTenantPageRequest)(nil),  // 1: srv.tenant.GetTenantPageRequest
	(*GetTenantInfoRequest)(nil),  // 2: srv.tenant.GetTenantInfoRequest
	(*GetTenantInfoResponse)(nil), // 3: srv.tenant.GetTenantInfoResponse
	(*GetTenantPageResponse)(nil), // 4: srv.tenant.GetTenantPageResponse
	(*StoreTenantRequest)(nil),    // 5: srv.tenant.StoreTenantRequest
	(*StoreTenantResponse)(nil),   // 6: srv.tenant.StoreTenantResponse
	(*UpdateTenantRequest)(nil),   // 7: srv.tenant.UpdateTenantRequest
	(*UpdateTenantResponse)(nil),  // 8: srv.tenant.UpdateTenantResponse
}
var file_proto_tenant_tenant_proto_depIdxs = []int32{
	0,  // 0: srv.tenant.GetTenantInfoResponse.data:type_name -> srv.tenant.Tenant
	0,  // 1: srv.tenant.GetTenantPageResponse.data:type_name -> srv.tenant.Tenant
	0,  // 2: srv.tenant.StoreTenantRequest.form:type_name -> srv.tenant.Tenant
	0,  // 3: srv.tenant.StoreTenantResponse.data:type_name -> srv.tenant.Tenant
	0,  // 4: srv.tenant.UpdateTenantRequest.form:type_name -> srv.tenant.Tenant
	0,  // 5: srv.tenant.UpdateTenantResponse.data:type_name -> srv.tenant.Tenant
	2,  // 6: srv.tenant.TenantService.GetTenantInfo:input_type -> srv.tenant.GetTenantInfoRequest
	1,  // 7: srv.tenant.TenantService.GetTenantPage:input_type -> srv.tenant.GetTenantPageRequest
	5,  // 8: srv.tenant.TenantService.StoreTenant:input_type -> srv.tenant.StoreTenantRequest
	7,  // 9: srv.tenant.TenantService.UpdateTenant:input_type -> srv.tenant.UpdateTenantRequest
	3,  // 10: srv.tenant.TenantService.GetTenantInfo:output_type -> srv.tenant.GetTenantInfoResponse
	4,  // 11: srv.tenant.TenantService.GetTenantPage:output_type -> srv.tenant.GetTenantPageResponse
	6,  // 12: srv.tenant.TenantService.StoreTenant:output_type -> srv.tenant.StoreTenantResponse
	8,  // 13: srv.tenant.TenantService.UpdateTenant:output_type -> srv.tenant.UpdateTenantResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_tenant_tenant_proto_init() }
func file_proto_tenant_tenant_proto_init() {
	if File_proto_tenant_tenant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_tenant_tenant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tenant); i {
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
		file_proto_tenant_tenant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTenantPageRequest); i {
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
		file_proto_tenant_tenant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTenantInfoRequest); i {
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
		file_proto_tenant_tenant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTenantInfoResponse); i {
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
		file_proto_tenant_tenant_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTenantPageResponse); i {
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
		file_proto_tenant_tenant_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreTenantRequest); i {
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
		file_proto_tenant_tenant_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreTenantResponse); i {
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
		file_proto_tenant_tenant_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTenantRequest); i {
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
		file_proto_tenant_tenant_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTenantResponse); i {
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
			RawDescriptor: file_proto_tenant_tenant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_tenant_tenant_proto_goTypes,
		DependencyIndexes: file_proto_tenant_tenant_proto_depIdxs,
		MessageInfos:      file_proto_tenant_tenant_proto_msgTypes,
	}.Build()
	File_proto_tenant_tenant_proto = out.File
	file_proto_tenant_tenant_proto_rawDesc = nil
	file_proto_tenant_tenant_proto_goTypes = nil
	file_proto_tenant_tenant_proto_depIdxs = nil
}
