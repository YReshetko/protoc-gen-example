// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ext.proto

package test_example

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

var E_Roles = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         1000,
	Name:          "test.example.roles",
	Tag:           "bytes,1000,opt,name=roles",
	Filename:      "proto/ext.proto",
}

var E_Addition = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         1001,
	Name:          "test.example.addition",
	Tag:           "bytes,1001,opt,name=addition",
	Filename:      "proto/ext.proto",
}

func init() {
	proto.RegisterExtension(E_Roles)
	proto.RegisterExtension(E_Addition)
}

func init() { proto.RegisterFile("proto/ext.proto", fileDescriptor_507b0bc5ae60d3fc) }

var fileDescriptor_507b0bc5ae60d3fc = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xad, 0x28, 0xd1, 0x03, 0xb3, 0x84, 0x78, 0x4a, 0x52, 0x8b, 0x4b, 0xf4, 0x52,
	0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52, 0xa5, 0x5c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92,
	0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0xc1, 0xca, 0x92, 0x4a, 0xd3, 0x20,
	0x8c, 0x64, 0xdd, 0xf4, 0xd4, 0x3c, 0xdd, 0xf4, 0x7c, 0xfd, 0x94, 0xd4, 0xe2, 0xe4, 0xa2, 0xcc,
	0x82, 0x92, 0xfc, 0x22, 0x24, 0x26, 0xc4, 0x4c, 0x2b, 0x53, 0x2e, 0xd6, 0xa2, 0xfc, 0x9c, 0xd4,
	0x62, 0x21, 0x39, 0xbd, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x3d, 0x98, 0x21, 0x7a, 0xbe, 0xa9,
	0x25, 0x19, 0xf9, 0x29, 0xfe, 0x05, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x12, 0x2f, 0xd8, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x20, 0xaa, 0xad, 0xac, 0xb9, 0x38, 0x12, 0x53, 0x52, 0x32, 0x41, 0x52, 0x04,
	0x75, 0xbe, 0x84, 0xe8, 0x84, 0x6b, 0x48, 0x62, 0x03, 0x2b, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xbb, 0xc8, 0x4e, 0x5e, 0xe1, 0x00, 0x00, 0x00,
}