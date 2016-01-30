// Code generated by protoc-gen-go.
// source: state.proto
// DO NOT EDIT!

package logs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// LogStreamState is a bidirectional state value used in UpdateStream calls.
//
// LogStreamState is embeddable in Endpoints request/response structs.
type LogStreamState struct {
	// ProtoVersion is the protobuf version for this stream.
	ProtoVersion string `protobuf:"bytes,1,opt,name=proto_version" json:"proto_version,omitempty"`
	// The time when the log stream was registered with the Coordinator.
	Created *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
	// The time when the log stream's state was last updated.
	Updated *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=updated" json:"updated,omitempty"`
	// The stream index of the log stream's terminal message. If the value is -1,
	// the log is still streaming.
	TerminalIndex int64 `protobuf:"varint,4,opt,name=terminal_index" json:"terminal_index,omitempty"`
	// If non-nil, the log stream is archived, and this field contains archival
	// details.
	Archive *LogStreamState_ArchiveInfo `protobuf:"bytes,5,opt,name=archive" json:"archive,omitempty"`
	// Indicates the purged state of a log. A log that has been purged is only
	// acknowledged to administrative clients.
	Purged bool `protobuf:"varint,6,opt,name=purged" json:"purged,omitempty"`
}

func (m *LogStreamState) Reset()                    { *m = LogStreamState{} }
func (m *LogStreamState) String() string            { return proto.CompactTextString(m) }
func (*LogStreamState) ProtoMessage()               {}
func (*LogStreamState) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *LogStreamState) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *LogStreamState) GetUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *LogStreamState) GetArchive() *LogStreamState_ArchiveInfo {
	if m != nil {
		return m.Archive
	}
	return nil
}

// ArchiveInfo contains archive details for the log stream.
type LogStreamState_ArchiveInfo struct {
	// The Google Storage URL where the log stream's index is archived.
	IndexUrl string `protobuf:"bytes,1,opt,name=index_url" json:"index_url,omitempty"`
	// The Google Storage URL where the log stream's raw stream data is archived.
	StreamUrl string `protobuf:"bytes,2,opt,name=stream_url" json:"stream_url,omitempty"`
	// The Google Storage URL where the log stream's assembled data is archived.
	DataUrl string `protobuf:"bytes,3,opt,name=data_url" json:"data_url,omitempty"`
}

func (m *LogStreamState_ArchiveInfo) Reset()                    { *m = LogStreamState_ArchiveInfo{} }
func (m *LogStreamState_ArchiveInfo) String() string            { return proto.CompactTextString(m) }
func (*LogStreamState_ArchiveInfo) ProtoMessage()               {}
func (*LogStreamState_ArchiveInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func init() {
	proto.RegisterType((*LogStreamState)(nil), "logs.LogStreamState")
	proto.RegisterType((*LogStreamState_ArchiveInfo)(nil), "logs.LogStreamState.ArchiveInfo")
}

var fileDescriptor1 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x69, 0x5a, 0xd3, 0x66, 0x82, 0x41, 0x17, 0x94, 0x90, 0x8b, 0xc5, 0x93, 0x20, 0x6c,
	0x51, 0x9f, 0xc0, 0x8b, 0x20, 0x78, 0xab, 0xf7, 0xb0, 0x6d, 0xa6, 0xeb, 0x42, 0x92, 0x09, 0x9b,
	0xdd, 0xe2, 0x03, 0xf9, 0xa0, 0xee, 0x4e, 0x2c, 0xe8, 0xc9, 0xeb, 0x3f, 0xdf, 0x7c, 0xf3, 0x0f,
	0xe4, 0xa3, 0x53, 0x0e, 0xe5, 0x60, 0xc9, 0x91, 0x58, 0xb4, 0xa4, 0xc7, 0xea, 0x46, 0x13, 0xe9,
	0x16, 0x37, 0x9c, 0xed, 0xfc, 0x61, 0xe3, 0x4c, 0x87, 0x01, 0xeb, 0x86, 0x09, 0xbb, 0xfd, 0x4a,
	0xa0, 0x78, 0x23, 0xbd, 0x75, 0x16, 0x55, 0xb7, 0x8d, 0xfb, 0xe2, 0x0a, 0xce, 0x79, 0x56, 0x1f,
	0xd1, 0x8e, 0x86, 0xfa, 0x72, 0xb6, 0x9e, 0xdd, 0x65, 0xe2, 0x1e, 0x96, 0xfb, 0x00, 0x39, 0x6c,
	0xca, 0x24, 0x04, 0xf9, 0x63, 0x25, 0x27, 0xb9, 0x3c, 0xc9, 0xe5, 0xfb, 0x49, 0x1e, 0x61, 0x3f,
	0x34, 0x0c, 0xcf, 0xff, 0x85, 0xaf, 0xa1, 0x70, 0x68, 0x3b, 0xd3, 0xab, 0xb6, 0x36, 0x7d, 0x83,
	0x9f, 0xe5, 0x22, 0xec, 0xcc, 0xc5, 0x03, 0x2c, 0x95, 0xdd, 0x7f, 0x98, 0x23, 0x96, 0x67, 0x2c,
	0x59, 0xcb, 0xf8, 0x94, 0xfc, 0xdb, 0x57, 0x3e, 0x4f, 0xcc, 0x6b, 0x7f, 0x20, 0x51, 0x40, 0x3a,
	0x78, 0xab, 0xc3, 0xd9, 0x34, 0x6c, 0xac, 0xaa, 0x17, 0xc8, 0x7f, 0x8f, 0x2f, 0x21, 0xe3, 0x03,
	0xb5, 0xb7, 0xed, 0xcf, 0x5b, 0x02, 0x60, 0x64, 0x19, 0x67, 0x09, 0x67, 0x17, 0xb0, 0x0a, 0xdd,
	0x15, 0x27, 0xb1, 0x7e, 0xb6, 0x4b, 0xb9, 0xf6, 0xd3, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f,
	0xb2, 0x4c, 0xe6, 0x63, 0x01, 0x00, 0x00,
}