// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package marshal_test is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	TestOne
*/
package marshal_test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TestOne struct {
	StringFoo  string `protobuf:"bytes,1,opt,name=stringFoo" json:"stringFoo,omitempty"`
	IntFoo     int32  `protobuf:"varint,2,opt,name=intFoo" json:"intFoo,omitempty"`
	StringFooB string `protobuf:"bytes,3,opt,name=stringFooB" json:"stringFooB,omitempty"`
}

func (m *TestOne) Reset()                    { *m = TestOne{} }
func (m *TestOne) String() string            { return proto.CompactTextString(m) }
func (*TestOne) ProtoMessage()               {}
func (*TestOne) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*TestOne)(nil), "marshal_test.TestOne")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc9, 0x4d, 0x2c, 0x2a, 0xce, 0x48, 0xcc, 0x89,
	0x07, 0x89, 0x29, 0xc5, 0x73, 0xb1, 0x87, 0x00, 0x69, 0xff, 0xbc, 0x54, 0x21, 0x19, 0x2e, 0xce,
	0xe2, 0x92, 0xa2, 0xcc, 0xbc, 0x74, 0xb7, 0xfc, 0x7c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x84, 0x80, 0x90, 0x18, 0x17, 0x5b, 0x66, 0x5e, 0x09, 0x48, 0x8a, 0x09, 0x28, 0xc5, 0x1a, 0x04,
	0xe5, 0x09, 0xc9, 0x71, 0x71, 0xc1, 0x15, 0x39, 0x49, 0x30, 0x83, 0xb5, 0x21, 0x89, 0x24, 0xb1,
	0x81, 0x6d, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x60, 0x36, 0xb5, 0x46, 0x83, 0x00, 0x00,
	0x00,
}
