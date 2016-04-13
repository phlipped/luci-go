// Code generated by protoc-gen-go.
// source: messages.proto
// DO NOT EDIT!

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	BlobWithSHA1
	TagCache
	InstanceCache
*/
package messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// BlobWithSHA1 is a wrapper around a binary blob with SHA1 hash to verify
// its integrity.
type BlobWithSHA1 struct {
	Blob []byte `protobuf:"bytes,1,opt,name=blob,proto3" json:"blob,omitempty"`
	Sha1 []byte `protobuf:"bytes,2,opt,name=sha1,proto3" json:"sha1,omitempty"`
}

func (m *BlobWithSHA1) Reset()                    { *m = BlobWithSHA1{} }
func (m *BlobWithSHA1) String() string            { return proto.CompactTextString(m) }
func (*BlobWithSHA1) ProtoMessage()               {}
func (*BlobWithSHA1) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// TagCache stores a mapping (package name, tag) -> instance ID to speed up
// subsequence ResolveVersion calls when tags are used.
type TagCache struct {
	// Capped list of entries, most recently resolved is last.
	Entries []*TagCache_Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *TagCache) Reset()                    { *m = TagCache{} }
func (m *TagCache) String() string            { return proto.CompactTextString(m) }
func (*TagCache) ProtoMessage()               {}
func (*TagCache) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TagCache) GetEntries() []*TagCache_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type TagCache_Entry struct {
	Package    string `protobuf:"bytes,1,opt,name=package" json:"package,omitempty"`
	Tag        string `protobuf:"bytes,2,opt,name=tag" json:"tag,omitempty"`
	InstanceId string `protobuf:"bytes,3,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
}

func (m *TagCache_Entry) Reset()                    { *m = TagCache_Entry{} }
func (m *TagCache_Entry) String() string            { return proto.CompactTextString(m) }
func (*TagCache_Entry) ProtoMessage()               {}
func (*TagCache_Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// InstanceCache stores a list of instances in cache
// and their last access time.
type InstanceCache struct {
	// Entries is a map of {instance id -> information about instance}.
	Entries map[string]*InstanceCache_Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// LastSynced is timestamp when we synchronized Entires with actual
	// instance files.
	LastSynced *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=last_synced,json=lastSynced" json:"last_synced,omitempty"`
}

func (m *InstanceCache) Reset()                    { *m = InstanceCache{} }
func (m *InstanceCache) String() string            { return proto.CompactTextString(m) }
func (*InstanceCache) ProtoMessage()               {}
func (*InstanceCache) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *InstanceCache) GetEntries() map[string]*InstanceCache_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *InstanceCache) GetLastSynced() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastSynced
	}
	return nil
}

// Entry stores info about an instance.
type InstanceCache_Entry struct {
	// LastAccess is last time this instance was retrieved from or put to the
	// cache.
	LastAccess *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=last_access,json=lastAccess" json:"last_access,omitempty"`
}

func (m *InstanceCache_Entry) Reset()                    { *m = InstanceCache_Entry{} }
func (m *InstanceCache_Entry) String() string            { return proto.CompactTextString(m) }
func (*InstanceCache_Entry) ProtoMessage()               {}
func (*InstanceCache_Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *InstanceCache_Entry) GetLastAccess() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastAccess
	}
	return nil
}

func init() {
	proto.RegisterType((*BlobWithSHA1)(nil), "messages.BlobWithSHA1")
	proto.RegisterType((*TagCache)(nil), "messages.TagCache")
	proto.RegisterType((*TagCache_Entry)(nil), "messages.TagCache.Entry")
	proto.RegisterType((*InstanceCache)(nil), "messages.InstanceCache")
	proto.RegisterType((*InstanceCache_Entry)(nil), "messages.InstanceCache.Entry")
}

var fileDescriptor0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x90, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0xe9, 0xf6, 0xee, 0xdd, 0x76, 0x3a, 0x45, 0x72, 0x55, 0x0a, 0x32, 0x29, 0x5e, 0xec,
	0x2a, 0x63, 0x1d, 0x88, 0x28, 0x08, 0xf3, 0x0f, 0xb8, 0xdb, 0xac, 0x20, 0x5e, 0x8d, 0xb4, 0x8d,
	0x6d, 0x59, 0xd7, 0x96, 0x25, 0x13, 0xfa, 0x3d, 0xfc, 0x1a, 0x7e, 0x47, 0x93, 0xb4, 0x51, 0xe7,
	0x85, 0x78, 0x77, 0xce, 0x93, 0x5f, 0xce, 0x79, 0xce, 0x03, 0xc7, 0x5b, 0xc6, 0x39, 0x4d, 0x18,
	0xc7, 0xd5, 0xae, 0x14, 0x25, 0x1a, 0x98, 0xde, 0x1d, 0x27, 0x65, 0x99, 0xe4, 0x6c, 0xaa, 0xf5,
	0x70, 0xff, 0x32, 0x15, 0x99, 0x7c, 0x13, 0x74, 0x5b, 0x35, 0xa8, 0x77, 0x01, 0xa3, 0xdb, 0xbc,
	0x0c, 0x9f, 0x32, 0x91, 0xae, 0x1e, 0x17, 0x33, 0x84, 0xe0, 0x5f, 0x28, 0x7b, 0xc7, 0x3a, 0xb3,
	0x26, 0x23, 0xa2, 0x6b, 0xa5, 0xf1, 0x94, 0xce, 0x9c, 0x4e, 0xa3, 0xa9, 0xda, 0x7b, 0xb3, 0x60,
	0x10, 0xd0, 0xe4, 0x8e, 0x46, 0x29, 0x43, 0x3e, 0xf4, 0x59, 0x21, 0x76, 0x19, 0xe3, 0xf2, 0x5f,
	0x77, 0x62, 0xfb, 0x0e, 0xfe, 0x74, 0x64, 0x20, 0xfc, 0x20, 0x89, 0x9a, 0x18, 0xd0, 0x0d, 0xa0,
	0xa7, 0x15, 0xe4, 0x40, 0xbf, 0xa2, 0xd1, 0x46, 0xc2, 0x7a, 0xe9, 0x90, 0x98, 0x16, 0x9d, 0x40,
	0x57, 0xd0, 0x44, 0xaf, 0x1d, 0x12, 0x55, 0xa2, 0x31, 0xd8, 0x59, 0x21, 0xed, 0x17, 0x11, 0x5b,
	0x67, 0xb1, 0xd3, 0xd5, 0x2f, 0x60, 0xa4, 0x65, 0xec, 0xbd, 0x77, 0xe0, 0x68, 0xd9, 0xb6, 0x8d,
	0xb7, 0x9b, 0x9f, 0xde, 0xce, 0xbf, 0xbc, 0x1d, 0x90, 0xda, 0xa0, 0xc4, 0x0e, 0x7d, 0xa2, 0x6b,
	0xb0, 0x73, 0xca, 0xc5, 0x9a, 0xd7, 0x12, 0x8c, 0xb5, 0x19, 0xdb, 0x77, 0x71, 0x93, 0x2b, 0x36,
	0xb9, 0xe2, 0xc0, 0xe4, 0x4a, 0x40, 0xe1, 0x2b, 0x4d, 0xbb, 0xf7, 0xe6, 0x48, 0x33, 0x85, 0x46,
	0x91, 0x5c, 0xfe, 0xd7, 0x29, 0x0b, 0x4d, 0xbb, 0xcf, 0x30, 0xfa, 0xee, 0x4d, 0xe5, 0xb2, 0x61,
	0x75, 0x9b, 0x96, 0x2a, 0xd1, 0x1c, 0x7a, 0xaf, 0x34, 0xdf, 0xb3, 0x76, 0xf0, 0xe9, 0x6f, 0x27,
	0xd6, 0xa4, 0x61, 0xaf, 0x3a, 0x97, 0x56, 0xf8, 0x5f, 0xaf, 0x9e, 0x7f, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xac, 0x56, 0x49, 0xa3, 0x42, 0x02, 0x00, 0x00,
}
