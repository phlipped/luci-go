// Code generated by protoc-gen-go.
// source: claim_execution.proto
// DO NOT EDIT!

package dm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ClaimExecutionRsp struct {
	Quest *Quest `protobuf:"bytes,1,opt,name=quest" json:"quest,omitempty"`
	// Auth is the auth with an Activation Token to be used with the
	// ActivateExecution rpc.
	Auth *Execution_Auth `protobuf:"bytes,2,opt,name=auth" json:"auth,omitempty"`
}

func (m *ClaimExecutionRsp) Reset()                    { *m = ClaimExecutionRsp{} }
func (m *ClaimExecutionRsp) String() string            { return proto.CompactTextString(m) }
func (*ClaimExecutionRsp) ProtoMessage()               {}
func (*ClaimExecutionRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ClaimExecutionRsp) GetQuest() *Quest {
	if m != nil {
		return m.Quest
	}
	return nil
}

func (m *ClaimExecutionRsp) GetAuth() *Execution_Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func init() {
	proto.RegisterType((*ClaimExecutionRsp)(nil), "dm.ClaimExecutionRsp")
}

var fileDescriptor1 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0x49, 0xcc,
	0xcc, 0x8d, 0x4f, 0xad, 0x48, 0x4d, 0x2e, 0x2d, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x4a, 0xc9, 0x95, 0x12, 0x48, 0x2f, 0x4a, 0x2c, 0xc8, 0x88, 0x4f, 0x49, 0x2c,
	0x49, 0x84, 0x88, 0x2a, 0xc5, 0x70, 0x09, 0x3a, 0x83, 0x94, 0xbb, 0xc2, 0x54, 0x07, 0x15, 0x17,
	0x08, 0xc9, 0x73, 0xb1, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b,
	0x71, 0xea, 0xa5, 0xe4, 0xea, 0x05, 0x82, 0x04, 0x82, 0x20, 0xe2, 0x42, 0x6a, 0x5c, 0x2c, 0x89,
	0xa5, 0x25, 0x19, 0x12, 0x4c, 0x60, 0x79, 0x21, 0x90, 0x3c, 0xdc, 0x00, 0x3d, 0x47, 0xa0, 0x4c,
	0x10, 0x58, 0x3e, 0x89, 0x0d, 0x6c, 0x89, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x70, 0xd5, 0x63,
	0x12, 0x93, 0x00, 0x00, 0x00,
}
