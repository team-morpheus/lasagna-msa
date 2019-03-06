// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal-recipes-service.proto

package lasagna_srv_internal_recipes_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Recipe struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Recipe) Reset()         { *m = Recipe{} }
func (m *Recipe) String() string { return proto.CompactTextString(m) }
func (*Recipe) ProtoMessage()    {}
func (*Recipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_459e21869af81ec2, []int{0}
}

func (m *Recipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Recipe.Unmarshal(m, b)
}
func (m *Recipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Recipe.Marshal(b, m, deterministic)
}
func (m *Recipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Recipe.Merge(m, src)
}
func (m *Recipe) XXX_Size() int {
	return xxx_messageInfo_Recipe.Size(m)
}
func (m *Recipe) XXX_DiscardUnknown() {
	xxx_messageInfo_Recipe.DiscardUnknown(m)
}

var xxx_messageInfo_Recipe proto.InternalMessageInfo

func (m *Recipe) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Recipe) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Recipe) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_459e21869af81ec2, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Response struct {
	Created              bool      `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Recipe               *Recipe   `protobuf:"bytes,2,opt,name=recipe,proto3" json:"recipe,omitempty"`
	Recipes              []*Recipe `protobuf:"bytes,3,rep,name=recipes,proto3" json:"recipes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_459e21869af81ec2, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetRecipe() *Recipe {
	if m != nil {
		return m.Recipe
	}
	return nil
}

func (m *Response) GetRecipes() []*Recipe {
	if m != nil {
		return m.Recipes
	}
	return nil
}

func init() {
	proto.RegisterType((*Recipe)(nil), "lasagna.srv.internal.recipes.service.Recipe")
	proto.RegisterType((*GetRequest)(nil), "lasagna.srv.internal.recipes.service.GetRequest")
	proto.RegisterType((*Response)(nil), "lasagna.srv.internal.recipes.service.Response")
}

func init() { proto.RegisterFile("internal-recipes-service.proto", fileDescriptor_459e21869af81ec2) }

var fileDescriptor_459e21869af81ec2 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x6d, 0x2b, 0xdd, 0x75, 0x76, 0xf1, 0x30, 0x07, 0x09, 0x1e, 0xa4, 0x14, 0x0f, 0x1e,
	0xdc, 0x20, 0xeb, 0x23, 0x28, 0x8a, 0x17, 0x0f, 0xf1, 0x09, 0x62, 0x3b, 0x48, 0x60, 0x4d, 0x63,
	0x26, 0xf6, 0xc1, 0x7c, 0x3c, 0x4f, 0x42, 0x92, 0xe2, 0x1e, 0xdb, 0x5b, 0x32, 0xf3, 0xfd, 0xf9,
	0x7f, 0x7e, 0x02, 0x57, 0xc6, 0x06, 0xf2, 0x56, 0x1f, 0x76, 0x9e, 0x3a, 0xe3, 0x88, 0x77, 0x4c,
	0x7e, 0x34, 0x1d, 0x49, 0xe7, 0x87, 0x30, 0xe0, 0xf5, 0x41, 0xb3, 0xfe, 0xb0, 0x5a, 0xb2, 0x1f,
	0xe5, 0xc4, 0xca, 0xcc, 0xca, 0xcc, 0xb6, 0xaf, 0x50, 0xab, 0x38, 0xc2, 0x73, 0x28, 0x4d, 0x2f,
	0x8a, 0xa6, 0xb8, 0x39, 0x53, 0xa5, 0xe9, 0x11, 0xe1, 0xd4, 0xea, 0x4f, 0x12, 0x65, 0x9c, 0xc4,
	0x33, 0x36, 0xb0, 0xe9, 0x89, 0x3b, 0x6f, 0x5c, 0x30, 0x83, 0x15, 0x55, 0x5c, 0x1d, 0x8f, 0xda,
	0x2d, 0xc0, 0x33, 0x05, 0x45, 0x5f, 0xdf, 0xc4, 0xa1, 0xfd, 0x29, 0x60, 0xad, 0x88, 0xdd, 0x60,
	0x99, 0x50, 0xc0, 0xaa, 0xf3, 0xa4, 0x03, 0x25, 0x97, 0xb5, 0x9a, 0xae, 0xf8, 0x08, 0x75, 0xca,
	0x15, 0xcd, 0x36, 0xfb, 0x5b, 0x39, 0x27, 0xbb, 0x4c, 0xc1, 0x55, 0xd6, 0xe2, 0x13, 0xac, 0x32,
	0x21, 0xaa, 0xa6, 0x5a, 0xfc, 0xcc, 0x24, 0xde, 0xff, 0x16, 0x70, 0xf1, 0x92, 0xe1, 0xb4, 0xe3,
	0xb7, 0x84, 0xa2, 0x85, 0xed, 0x43, 0xcc, 0x9c, 0x3b, 0x5b, 0xe4, 0x70, 0x29, 0xe7, 0xd2, 0xa9,
	0xb0, 0xf6, 0x04, 0x5d, 0x6e, 0x33, 0x02, 0x78, 0x37, 0x4f, 0xff, 0xdf, 0xff, 0x72, 0xc7, 0xf7,
	0x3a, 0x7e, 0x9e, 0xfb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0x42, 0x04, 0xa1, 0x5e, 0x02,
	0x00, 0x00,
}
