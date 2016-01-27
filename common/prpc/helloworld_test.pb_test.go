// Code generated by protoc-gen-go.
// source: helloworld_test.proto
// DO NOT EDIT!

/*
Package prpc is a generated protocol buffer package.

It is generated from these files:
	helloworld_test.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package prpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HelloRequest)(nil), "prpc.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "prpc.HelloReply")
}

var fileDescriptor0 = []byte{
	// 106 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0x89, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x29, 0x28, 0x2a, 0x48, 0x56, 0x92, 0xe1, 0xe2, 0xf1, 0x00, 0x49, 0x07, 0xa5,
	0x16, 0x96, 0x02, 0xe5, 0x84, 0x78, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x95, 0x64, 0xb9, 0xb8, 0xa0, 0xb2, 0x05, 0x39, 0x95, 0x42, 0xfc, 0x5c, 0xec, 0xb9,
	0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0x50, 0xe9, 0x24, 0x36, 0xb0, 0x49, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x73, 0xfc, 0x73, 0xf6, 0x62, 0x00, 0x00, 0x00,
}
