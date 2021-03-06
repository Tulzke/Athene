// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api/property.proto

package api

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

type SearchParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inventory string   `protobuf:"bytes,1,opt,name=Inventory,proto3" json:"Inventory,omitempty"`
	Serial    string   `protobuf:"bytes,2,opt,name=Serial,proto3" json:"Serial,omitempty"`
	Action    uint32   `protobuf:"varint,3,opt,name=Action,proto3" json:"Action,omitempty"`
	Warehouse uint32   `protobuf:"varint,4,opt,name=Warehouse,proto3" json:"Warehouse,omitempty"`
	Offset    uint32   `protobuf:"varint,5,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit     uint32   `protobuf:"varint,6,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Order     string   `protobuf:"bytes,7,opt,name=Order,proto3" json:"Order,omitempty"`
	Name      string   `protobuf:"bytes,8,opt,name=Name,proto3" json:"Name,omitempty"`
	Groups    []string `protobuf:"bytes,9,rep,name=Groups,proto3" json:"Groups,omitempty"`
}

func (x *SearchParams) Reset() {
	*x = SearchParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_property_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchParams) ProtoMessage() {}

func (x *SearchParams) ProtoReflect() protoreflect.Message {
	mi := &file_api_property_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchParams.ProtoReflect.Descriptor instead.
func (*SearchParams) Descriptor() ([]byte, []int) {
	return file_api_property_proto_rawDescGZIP(), []int{0}
}

func (x *SearchParams) GetInventory() string {
	if x != nil {
		return x.Inventory
	}
	return ""
}

func (x *SearchParams) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *SearchParams) GetAction() uint32 {
	if x != nil {
		return x.Action
	}
	return 0
}

func (x *SearchParams) GetWarehouse() uint32 {
	if x != nil {
		return x.Warehouse
	}
	return 0
}

func (x *SearchParams) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SearchParams) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SearchParams) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

func (x *SearchParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SearchParams) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

type Properties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Properties []*Property `protobuf:"bytes,1,rep,name=Properties,proto3" json:"Properties,omitempty"`
}

func (x *Properties) Reset() {
	*x = Properties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_property_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Properties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Properties) ProtoMessage() {}

func (x *Properties) ProtoReflect() protoreflect.Message {
	mi := &file_api_property_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Properties.ProtoReflect.Descriptor instead.
func (*Properties) Descriptor() ([]byte, []int) {
	return file_api_property_proto_rawDescGZIP(), []int{1}
}

func (x *Properties) GetProperties() []*Property {
	if x != nil {
		return x.Properties
	}
	return nil
}

type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32                 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Inventory string                 `protobuf:"bytes,2,opt,name=Inventory,proto3" json:"Inventory,omitempty"`
	Serial    string                 `protobuf:"bytes,3,opt,name=Serial,proto3" json:"Serial,omitempty"`
	Name      string                 `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=Created_at,json=CreatedAt,proto3" json:"Created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=Updated_at,json=UpdatedAt,proto3" json:"Updated_at,omitempty"`
	Warehouse uint32                 `protobuf:"varint,7,opt,name=Warehouse,proto3" json:"Warehouse,omitempty"`
	State     uint32                 `protobuf:"varint,8,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_property_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_api_property_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_api_property_proto_rawDescGZIP(), []int{2}
}

func (x *Property) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Property) GetInventory() string {
	if x != nil {
		return x.Inventory
	}
	return ""
}

func (x *Property) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *Property) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Property) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Property) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Property) GetWarehouse() uint32 {
	if x != nil {
		return x.Warehouse
	}
	return 0
}

func (x *Property) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

type Warehouse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Warehouse) Reset() {
	*x = Warehouse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_property_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Warehouse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Warehouse) ProtoMessage() {}

func (x *Warehouse) ProtoReflect() protoreflect.Message {
	mi := &file_api_property_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Warehouse.ProtoReflect.Descriptor instead.
func (*Warehouse) Descriptor() ([]byte, []int) {
	return file_api_property_proto_rawDescGZIP(), []int{3}
}

func (x *Warehouse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Warehouse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Warehouses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Warehouse `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Warehouses) Reset() {
	*x = Warehouses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_property_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Warehouses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Warehouses) ProtoMessage() {}

func (x *Warehouses) ProtoReflect() protoreflect.Message {
	mi := &file_api_property_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Warehouses.ProtoReflect.Descriptor instead.
func (*Warehouses) Descriptor() ([]byte, []int) {
	return file_api_property_proto_rawDescGZIP(), []int{4}
}

func (x *Warehouses) GetData() []*Warehouse {
	if x != nil {
		return x.Data
	}
	return nil
}

type EmptyPropertyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyPropertyRequest) Reset() {
	*x = EmptyPropertyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_property_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyPropertyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyPropertyRequest) ProtoMessage() {}

func (x *EmptyPropertyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_property_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyPropertyRequest.ProtoReflect.Descriptor instead.
func (*EmptyPropertyRequest) Descriptor() ([]byte, []int) {
	return file_api_property_proto_rawDescGZIP(), []int{5}
}

type EmptyPropertyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyPropertyResponse) Reset() {
	*x = EmptyPropertyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_property_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyPropertyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyPropertyResponse) ProtoMessage() {}

func (x *EmptyPropertyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_property_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyPropertyResponse.ProtoReflect.Descriptor instead.
func (*EmptyPropertyResponse) Descriptor() ([]byte, []int) {
	return file_api_property_proto_rawDescGZIP(), []int{6}
}

type Count struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number uint64 `protobuf:"varint,1,opt,name=Number,proto3" json:"Number,omitempty"`
}

func (x *Count) Reset() {
	*x = Count{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_property_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Count) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Count) ProtoMessage() {}

func (x *Count) ProtoReflect() protoreflect.Message {
	mi := &file_api_property_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Count.ProtoReflect.Descriptor instead.
func (*Count) Descriptor() ([]byte, []int) {
	return file_api_property_proto_rawDescGZIP(), []int{7}
}

func (x *Count) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type GetActionsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actions []*OneAction `protobuf:"bytes,1,rep,name=Actions,proto3" json:"Actions,omitempty"`
}

func (x *GetActionsListResponse) Reset() {
	*x = GetActionsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_property_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActionsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActionsListResponse) ProtoMessage() {}

func (x *GetActionsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_property_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActionsListResponse.ProtoReflect.Descriptor instead.
func (*GetActionsListResponse) Descriptor() ([]byte, []int) {
	return file_api_property_proto_rawDescGZIP(), []int{8}
}

func (x *GetActionsListResponse) GetActions() []*OneAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

type OneAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Action string `protobuf:"bytes,3,opt,name=Action,proto3" json:"Action,omitempty"`
}

func (x *OneAction) Reset() {
	*x = OneAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_property_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneAction) ProtoMessage() {}

func (x *OneAction) ProtoReflect() protoreflect.Message {
	mi := &file_api_property_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneAction.ProtoReflect.Descriptor instead.
func (*OneAction) Descriptor() ([]byte, []int) {
	return file_api_property_proto_rawDescGZIP(), []int{9}
}

func (x *OneAction) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OneAction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OneAction) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type GetOnePropertyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetOnePropertyRequest) Reset() {
	*x = GetOnePropertyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_property_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOnePropertyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOnePropertyRequest) ProtoMessage() {}

func (x *GetOnePropertyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_property_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOnePropertyRequest.ProtoReflect.Descriptor instead.
func (*GetOnePropertyRequest) Descriptor() ([]byte, []int) {
	return file_api_property_proto_rawDescGZIP(), []int{10}
}

func (x *GetOnePropertyRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IsInWarehouseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WarehouseId uint32 `protobuf:"varint,1,opt,name=WarehouseId,proto3" json:"WarehouseId,omitempty"`
	PropertyId  uint32 `protobuf:"varint,2,opt,name=PropertyId,proto3" json:"PropertyId,omitempty"`
}

func (x *IsInWarehouseReq) Reset() {
	*x = IsInWarehouseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_property_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsInWarehouseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsInWarehouseReq) ProtoMessage() {}

func (x *IsInWarehouseReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_property_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsInWarehouseReq.ProtoReflect.Descriptor instead.
func (*IsInWarehouseReq) Descriptor() ([]byte, []int) {
	return file_api_property_proto_rawDescGZIP(), []int{11}
}

func (x *IsInWarehouseReq) GetWarehouseId() uint32 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *IsInWarehouseReq) GetPropertyId() uint32 {
	if x != nil {
		return x.PropertyId
	}
	return 0
}

type PropStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok      bool   `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *PropStatus) Reset() {
	*x = PropStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_property_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropStatus) ProtoMessage() {}

func (x *PropStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_property_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropStatus.ProtoReflect.Descriptor instead.
func (*PropStatus) Descriptor() ([]byte, []int) {
	return file_api_property_proto_rawDescGZIP(), []int{12}
}

func (x *PropStatus) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *PropStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendToWarhouseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WarehouseId  uint32   `protobuf:"varint,1,opt,name=WarehouseId,proto3" json:"WarehouseId,omitempty"`
	PropertiesId []uint32 `protobuf:"varint,2,rep,packed,name=PropertiesId,proto3" json:"PropertiesId,omitempty"`
}

func (x *SendToWarhouseReq) Reset() {
	*x = SendToWarhouseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_property_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendToWarhouseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendToWarhouseReq) ProtoMessage() {}

func (x *SendToWarhouseReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_property_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendToWarhouseReq.ProtoReflect.Descriptor instead.
func (*SendToWarhouseReq) Descriptor() ([]byte, []int) {
	return file_api_property_proto_rawDescGZIP(), []int{13}
}

func (x *SendToWarhouseReq) GetWarehouseId() uint32 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *SendToWarhouseReq) GetPropertiesId() []uint32 {
	if x != nil {
		return x.PropertiesId
	}
	return nil
}

var File_api_property_proto protoreflect.FileDescriptor

var file_api_property_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x22, 0x37, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x29, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52,
	0x0a, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x8e, 0x02, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x57, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x2f, 0x0a, 0x09,
	0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a,
	0x0a, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x57, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x16, 0x0a, 0x14, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x0a, 0x05,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3e, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4f, 0x6e, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x47, 0x0a,
	0x09, 0x4f, 0x6e, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x22,
	0x54, 0x0a, 0x10, 0x49, 0x73, 0x49, 0x6e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x49, 0x64, 0x22, 0x36, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x70, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x4f, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x59, 0x0a,
	0x11, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x57, 0x61, 0x72, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x49, 0x64, 0x32, 0xe9, 0x02, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x12, 0x0d, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x33, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x73, 0x12, 0x15, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x57, 0x61, 0x72, 0x65, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x0d, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x1a, 0x06, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4f, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x16, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12,
	0x2f, 0x0a, 0x0d, 0x49, 0x73, 0x4f, 0x6e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x12, 0x11, 0x2e, 0x49, 0x73, 0x49, 0x6e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x32, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x12, 0x12, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x57, 0x61, 0x72, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_property_proto_rawDescOnce sync.Once
	file_api_property_proto_rawDescData = file_api_property_proto_rawDesc
)

func file_api_property_proto_rawDescGZIP() []byte {
	file_api_property_proto_rawDescOnce.Do(func() {
		file_api_property_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_property_proto_rawDescData)
	})
	return file_api_property_proto_rawDescData
}

var file_api_property_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_property_proto_goTypes = []interface{}{
	(*SearchParams)(nil),           // 0: SearchParams
	(*Properties)(nil),             // 1: Properties
	(*Property)(nil),               // 2: Property
	(*Warehouse)(nil),              // 3: Warehouse
	(*Warehouses)(nil),             // 4: Warehouses
	(*EmptyPropertyRequest)(nil),   // 5: EmptyPropertyRequest
	(*EmptyPropertyResponse)(nil),  // 6: EmptyPropertyResponse
	(*Count)(nil),                  // 7: Count
	(*GetActionsListResponse)(nil), // 8: GetActionsListResponse
	(*OneAction)(nil),              // 9: OneAction
	(*GetOnePropertyRequest)(nil),  // 10: GetOnePropertyRequest
	(*IsInWarehouseReq)(nil),       // 11: IsInWarehouseReq
	(*PropStatus)(nil),             // 12: PropStatus
	(*SendToWarhouseReq)(nil),      // 13: SendToWarhouseReq
	(*timestamppb.Timestamp)(nil),  // 14: google.protobuf.Timestamp
}
var file_api_property_proto_depIdxs = []int32{
	2,  // 0: Properties.Properties:type_name -> Property
	14, // 1: Property.Created_at:type_name -> google.protobuf.Timestamp
	14, // 2: Property.Updated_at:type_name -> google.protobuf.Timestamp
	3,  // 3: Warehouses.Data:type_name -> Warehouse
	9,  // 4: GetActionsListResponse.Actions:type_name -> OneAction
	0,  // 5: property.GetProperty:input_type -> SearchParams
	5,  // 6: property.GetWarehouses:input_type -> EmptyPropertyRequest
	0,  // 7: property.GetCount:input_type -> SearchParams
	5,  // 8: property.GetActionsList:input_type -> EmptyPropertyRequest
	10, // 9: property.GetOneProperty:input_type -> GetOnePropertyRequest
	11, // 10: property.IsOnWarehouse:input_type -> IsInWarehouseReq
	13, // 11: property.SendToWarehouse:input_type -> SendToWarhouseReq
	1,  // 12: property.GetProperty:output_type -> Properties
	4,  // 13: property.GetWarehouses:output_type -> Warehouses
	7,  // 14: property.GetCount:output_type -> Count
	8,  // 15: property.GetActionsList:output_type -> GetActionsListResponse
	2,  // 16: property.GetOneProperty:output_type -> Property
	12, // 17: property.IsOnWarehouse:output_type -> PropStatus
	12, // 18: property.SendToWarehouse:output_type -> PropStatus
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_property_proto_init() }
func file_api_property_proto_init() {
	if File_api_property_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_property_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchParams); i {
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
		file_api_property_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Properties); i {
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
		file_api_property_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
		file_api_property_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Warehouse); i {
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
		file_api_property_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Warehouses); i {
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
		file_api_property_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyPropertyRequest); i {
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
		file_api_property_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyPropertyResponse); i {
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
		file_api_property_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Count); i {
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
		file_api_property_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActionsListResponse); i {
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
		file_api_property_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneAction); i {
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
		file_api_property_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOnePropertyRequest); i {
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
		file_api_property_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsInWarehouseReq); i {
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
		file_api_property_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropStatus); i {
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
		file_api_property_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendToWarhouseReq); i {
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
			RawDescriptor: file_api_property_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_property_proto_goTypes,
		DependencyIndexes: file_api_property_proto_depIdxs,
		MessageInfos:      file_api_property_proto_msgTypes,
	}.Build()
	File_api_property_proto = out.File
	file_api_property_proto_rawDesc = nil
	file_api_property_proto_goTypes = nil
	file_api_property_proto_depIdxs = nil
}
