// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: shop_product.proto

package shop_product

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     uint64                 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	mi := &file_shop_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_shop_product_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductRequest) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type ShopDetail struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShopDetail) Reset() {
	*x = ShopDetail{}
	mi := &file_shop_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShopDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopDetail) ProtoMessage() {}

func (x *ShopDetail) ProtoReflect() protoreflect.Message {
	mi := &file_shop_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopDetail.ProtoReflect.Descriptor instead.
func (*ShopDetail) Descriptor() ([]byte, []int) {
	return file_shop_product_proto_rawDescGZIP(), []int{1}
}

func (x *ShopDetail) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ShopDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Stock         uint64                 `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	Shop          *ShopDetail            `protobuf:"bytes,6,opt,name=shop,proto3" json:"shop,omitempty"`
	Image         string                 `protobuf:"bytes,7,opt,name=image,proto3" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductResponse) Reset() {
	*x = GetProductResponse{}
	mi := &file_shop_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductResponse) ProtoMessage() {}

func (x *GetProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductResponse.ProtoReflect.Descriptor instead.
func (*GetProductResponse) Descriptor() ([]byte, []int) {
	return file_shop_product_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProductResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetProductResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetProductResponse) GetStock() uint64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *GetProductResponse) GetShop() *ShopDetail {
	if x != nil {
		return x.Shop
	}
	return nil
}

func (x *GetProductResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type GetProductsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductIds    []uint64               `protobuf:"varint,1,rep,packed,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductsRequest) Reset() {
	*x = GetProductsRequest{}
	mi := &file_shop_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsRequest) ProtoMessage() {}

func (x *GetProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_product_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsRequest.ProtoReflect.Descriptor instead.
func (*GetProductsRequest) Descriptor() ([]byte, []int) {
	return file_shop_product_proto_rawDescGZIP(), []int{3}
}

func (x *GetProductsRequest) GetProductIds() []uint64 {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

type GetProductsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*GetProductResponse  `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductsResponse) Reset() {
	*x = GetProductsResponse{}
	mi := &file_shop_product_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsResponse) ProtoMessage() {}

func (x *GetProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_product_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsResponse.ProtoReflect.Descriptor instead.
func (*GetProductsResponse) Descriptor() ([]byte, []int) {
	return file_shop_product_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductsResponse) GetProducts() []*GetProductResponse {
	if x != nil {
		return x.Products
	}
	return nil
}

type CheckProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     uint64                 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckProductRequest) Reset() {
	*x = CheckProductRequest{}
	mi := &file_shop_product_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckProductRequest) ProtoMessage() {}

func (x *CheckProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_product_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckProductRequest.ProtoReflect.Descriptor instead.
func (*CheckProductRequest) Descriptor() ([]byte, []int) {
	return file_shop_product_proto_rawDescGZIP(), []int{5}
}

func (x *CheckProductRequest) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type CheckProductReponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsExists      bool                   `protobuf:"varint,1,opt,name=is_exists,json=isExists,proto3" json:"is_exists,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckProductReponse) Reset() {
	*x = CheckProductReponse{}
	mi := &file_shop_product_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckProductReponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckProductReponse) ProtoMessage() {}

func (x *CheckProductReponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_product_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckProductReponse.ProtoReflect.Descriptor instead.
func (*CheckProductReponse) Descriptor() ([]byte, []int) {
	return file_shop_product_proto_rawDescGZIP(), []int{6}
}

func (x *CheckProductReponse) GetIsExists() bool {
	if x != nil {
		return x.IsExists
	}
	return false
}

type CheckShopRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShopId        uint64                 `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckShopRequest) Reset() {
	*x = CheckShopRequest{}
	mi := &file_shop_product_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckShopRequest) ProtoMessage() {}

func (x *CheckShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_product_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckShopRequest.ProtoReflect.Descriptor instead.
func (*CheckShopRequest) Descriptor() ([]byte, []int) {
	return file_shop_product_proto_rawDescGZIP(), []int{7}
}

func (x *CheckShopRequest) GetShopId() uint64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

type CheckShopReponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsExists      bool                   `protobuf:"varint,1,opt,name=is_exists,json=isExists,proto3" json:"is_exists,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckShopReponse) Reset() {
	*x = CheckShopReponse{}
	mi := &file_shop_product_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckShopReponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckShopReponse) ProtoMessage() {}

func (x *CheckShopReponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_product_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckShopReponse.ProtoReflect.Descriptor instead.
func (*CheckShopReponse) Descriptor() ([]byte, []int) {
	return file_shop_product_proto_rawDescGZIP(), []int{8}
}

func (x *CheckShopReponse) GetIsExists() bool {
	if x != nil {
		return x.IsExists
	}
	return false
}

var File_shop_product_proto protoreflect.FileDescriptor

var file_shop_product_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x70, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xca, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x2c, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53,
	0x68, 0x6f, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x73, 0x22, 0x53, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x22, 0x34, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x69, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0x2b, 0x0a, 0x10, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x32, 0xde, 0x02, 0x0a, 0x12, 0x53, 0x68,
	0x6f, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1f,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x52, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x12, 0x20, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x21, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0a,
	0x53, 0x68, 0x6f, 0x70, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53,
	0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53,
	0x68, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_shop_product_proto_rawDescOnce sync.Once
	file_shop_product_proto_rawDescData []byte
)

func file_shop_product_proto_rawDescGZIP() []byte {
	file_shop_product_proto_rawDescOnce.Do(func() {
		file_shop_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_shop_product_proto_rawDesc), len(file_shop_product_proto_rawDesc)))
	})
	return file_shop_product_proto_rawDescData
}

var file_shop_product_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_shop_product_proto_goTypes = []any{
	(*GetProductRequest)(nil),   // 0: shop_product.GetProductRequest
	(*ShopDetail)(nil),          // 1: shop_product.ShopDetail
	(*GetProductResponse)(nil),  // 2: shop_product.GetProductResponse
	(*GetProductsRequest)(nil),  // 3: shop_product.GetProductsRequest
	(*GetProductsResponse)(nil), // 4: shop_product.GetProductsResponse
	(*CheckProductRequest)(nil), // 5: shop_product.CheckProductRequest
	(*CheckProductReponse)(nil), // 6: shop_product.CheckProductReponse
	(*CheckShopRequest)(nil),    // 7: shop_product.CheckShopRequest
	(*CheckShopReponse)(nil),    // 8: shop_product.CheckShopReponse
}
var file_shop_product_proto_depIdxs = []int32{
	1, // 0: shop_product.GetProductResponse.shop:type_name -> shop_product.ShopDetail
	2, // 1: shop_product.GetProductsResponse.products:type_name -> shop_product.GetProductResponse
	0, // 2: shop_product.ShopProductService.GetProduct:input_type -> shop_product.GetProductRequest
	3, // 3: shop_product.ShopProductService.GetProducts:input_type -> shop_product.GetProductsRequest
	5, // 4: shop_product.ShopProductService.ProductExists:input_type -> shop_product.CheckProductRequest
	7, // 5: shop_product.ShopProductService.ShopExists:input_type -> shop_product.CheckShopRequest
	2, // 6: shop_product.ShopProductService.GetProduct:output_type -> shop_product.GetProductResponse
	4, // 7: shop_product.ShopProductService.GetProducts:output_type -> shop_product.GetProductsResponse
	6, // 8: shop_product.ShopProductService.ProductExists:output_type -> shop_product.CheckProductReponse
	8, // 9: shop_product.ShopProductService.ShopExists:output_type -> shop_product.CheckShopReponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_shop_product_proto_init() }
func file_shop_product_proto_init() {
	if File_shop_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_shop_product_proto_rawDesc), len(file_shop_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shop_product_proto_goTypes,
		DependencyIndexes: file_shop_product_proto_depIdxs,
		MessageInfos:      file_shop_product_proto_msgTypes,
	}.Build()
	File_shop_product_proto = out.File
	file_shop_product_proto_goTypes = nil
	file_shop_product_proto_depIdxs = nil
}
