// Code generated by protoc-gen-go.
// source: walk_graph.proto
// DO NOT EDIT!

package dm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/luci/luci-go/common/proto/google"
import _ "github.com/luci/luci-go/common/proto/google"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Direction indicates that direction of dependencies that the request should
// walk.
type WalkGraphReq_Mode_Direction int32

const (
	WalkGraphReq_Mode_FORWARDS  WalkGraphReq_Mode_Direction = 0
	WalkGraphReq_Mode_BACKWARDS WalkGraphReq_Mode_Direction = 1
	WalkGraphReq_Mode_BOTH      WalkGraphReq_Mode_Direction = 2
)

var WalkGraphReq_Mode_Direction_name = map[int32]string{
	0: "FORWARDS",
	1: "BACKWARDS",
	2: "BOTH",
}
var WalkGraphReq_Mode_Direction_value = map[string]int32{
	"FORWARDS":  0,
	"BACKWARDS": 1,
	"BOTH":      2,
}

func (x WalkGraphReq_Mode_Direction) String() string {
	return proto.EnumName(WalkGraphReq_Mode_Direction_name, int32(x))
}
func (WalkGraphReq_Mode_Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor8, []int{0, 0, 0}
}

// WalkGraphReq allows you to walk from one or more Quests through their
// Attempt's forward dependencies.
//
//
// The handler will evaluate all of the queries, executing them in parallel.
// For each attempt or quest produced by the query, it will queue a walk
// operation for that node, respecting the options set (max_depth, etc.).
type WalkGraphReq struct {
	// optional. See Include.AttemptResult for restrictions.
	Auth *Execution_Auth `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	// Query specifies a list of queries to start the graph traversal on. The
	// traversal will occur as a union of the query results. Redundant
	// specification will not cause additional heavy work; every graph node will
	// be processed exactly once, regardless of how many times it appears in the
	// query results. However, redundancy in the queries will cause the server to
	// retrieve and discard more information.
	Query *GraphQuery         `protobuf:"bytes,2,opt,name=query" json:"query,omitempty"`
	Mode  *WalkGraphReq_Mode  `protobuf:"bytes,3,opt,name=mode" json:"mode,omitempty"`
	Limit *WalkGraphReq_Limit `protobuf:"bytes,4,opt,name=limit" json:"limit,omitempty"`
	// Include allows you to add additional information to the returned
	// GraphData which is typically medium-to-large sized.
	Include *WalkGraphReq_Include `protobuf:"bytes,5,opt,name=include" json:"include,omitempty"`
}

func (m *WalkGraphReq) Reset()                    { *m = WalkGraphReq{} }
func (m *WalkGraphReq) String() string            { return proto.CompactTextString(m) }
func (*WalkGraphReq) ProtoMessage()               {}
func (*WalkGraphReq) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *WalkGraphReq) GetAuth() *Execution_Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *WalkGraphReq) GetQuery() *GraphQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *WalkGraphReq) GetMode() *WalkGraphReq_Mode {
	if m != nil {
		return m.Mode
	}
	return nil
}

func (m *WalkGraphReq) GetLimit() *WalkGraphReq_Limit {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *WalkGraphReq) GetInclude() *WalkGraphReq_Include {
	if m != nil {
		return m.Include
	}
	return nil
}

type WalkGraphReq_Mode struct {
	// DFS sets whether this is a Depth-first (ish) or a Breadth-first (ish) load.
	// Since the load operation is multi-threaded, the search order is best
	// effort, but will actually be some hybrid between DFS and BFS. This setting
	// controls the bias direction of the hybrid loading algorithm.
	Dfs       bool                        `protobuf:"varint,1,opt,name=dfs" json:"dfs,omitempty"`
	Direction WalkGraphReq_Mode_Direction `protobuf:"varint,2,opt,name=direction,enum=dm.WalkGraphReq_Mode_Direction" json:"direction,omitempty"`
}

func (m *WalkGraphReq_Mode) Reset()                    { *m = WalkGraphReq_Mode{} }
func (m *WalkGraphReq_Mode) String() string            { return proto.CompactTextString(m) }
func (*WalkGraphReq_Mode) ProtoMessage()               {}
func (*WalkGraphReq_Mode) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0, 0} }

type WalkGraphReq_Limit struct {
	// MaxDepth sets the number of attempts to traverse; 0 means 'immediate'
	// (no dependencies), -1 means 'no limit', and >0 is a limit.
	//
	// Any negative value besides -1 is an error.
	MaxDepth int64 `protobuf:"varint,1,opt,name=max_depth,json=maxDepth" json:"max_depth,omitempty"`
	// MaxTime sets the maximum amount of time that the query processor should
	// take. Application of this deadline is 'best effort', which means the query
	// may take a bit longer than this timeout and still attempt to return data.
	//
	// This is different than the grpc timeout header, which will set a hard
	// deadline for the request.
	MaxTime *google_protobuf2.Duration `protobuf:"bytes,2,opt,name=max_time,json=maxTime" json:"max_time,omitempty"`
	// MaxDataSize sets the maximum amount of 'Data' (in bytes) that can be
	// returned, if include.quest_data, include.attempt_data, and/or
	// include.attempt_result are set. If this limit is hit, then the
	// appropriate 'partial' value will be set for that object, but the base
	// object would still be included in the result.
	//
	// If this limit is 0, a default limit of 16MB will be used. If this limit
	// exceeds 30MB, it will be reduced to 30MB.
	MaxDataSize uint32 `protobuf:"varint,3,opt,name=max_data_size,json=maxDataSize" json:"max_data_size,omitempty"`
}

func (m *WalkGraphReq_Limit) Reset()                    { *m = WalkGraphReq_Limit{} }
func (m *WalkGraphReq_Limit) String() string            { return proto.CompactTextString(m) }
func (*WalkGraphReq_Limit) ProtoMessage()               {}
func (*WalkGraphReq_Limit) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0, 1} }

func (m *WalkGraphReq_Limit) GetMaxTime() *google_protobuf2.Duration {
	if m != nil {
		return m.MaxTime
	}
	return nil
}

type WalkGraphReq_Include struct {
	// ObjectIds fills the 'Id' field of Quest, Attempt, and Execution objects.
	// If this is false, then those fields will be omitted.
	ObjectIds bool `protobuf:"varint,1,opt,name=object_ids,json=objectIds" json:"object_ids,omitempty"`
	// QuestData instructs the request to include the Data field for Quests in
	// GraphData.
	QuestData bool `protobuf:"varint,2,opt,name=quest_data,json=questData" json:"quest_data,omitempty"`
	// AttemptData instructs the request to include the Data field for Attempts
	// in GraphData.
	AttemptData bool `protobuf:"varint,3,opt,name=attempt_data,json=attemptData" json:"attempt_data,omitempty"`
	// AttemptResult will include the Attempt result payloads for any
	// Attempts that it returns. This option also implies AttemptData.
	//
	// If the requestor is an execution, the query logic will only include
	// result for an Attempt if the execution's Attempt depends on it, otherwise
	// it will be blank. To view an AttemptResult, the querying Attempt must
	// first depend on it.
	//
	// If the request would return more than limit.max_data_size of data, the
	// remaining attempt results will have their Partial.Data field set to true.
	AttemptResult bool `protobuf:"varint,4,opt,name=attempt_result,json=attemptResult" json:"attempt_result,omitempty"`
	// ExpiredAttempts allows you to view attempts which have expired results,
	// which are normally excluded from the graph.
	ExpiredAttempts bool `protobuf:"varint,5,opt,name=expired_attempts,json=expiredAttempts" json:"expired_attempts,omitempty"`
	// Executions is the number of Executions to include per Attempt. If this
	// is 0, then the execution data will be omitted completely.
	NumExecutions uint32 `protobuf:"varint,6,opt,name=num_executions,json=numExecutions" json:"num_executions,omitempty"`
	// FwdDeps instructs WalkGraph to include forward dependency information
	// from the result. This only changes the presence of information in the
	// result; if the query is walking forward attempt dependencies, that will
	// still occur even if this is false.
	FwdDeps bool `protobuf:"varint,7,opt,name=fwd_deps,json=fwdDeps" json:"fwd_deps,omitempty"`
	// BackDeps instructs WalkGraph to include the backwards dependency
	// information. This only changes the presence of information in the result;
	// if the query is walking backward attempt dependencies, that will still
	// occur even if this is false.
	BackDeps bool `protobuf:"varint,8,opt,name=back_deps,json=backDeps" json:"back_deps,omitempty"`
}

func (m *WalkGraphReq_Include) Reset()                    { *m = WalkGraphReq_Include{} }
func (m *WalkGraphReq_Include) String() string            { return proto.CompactTextString(m) }
func (*WalkGraphReq_Include) ProtoMessage()               {}
func (*WalkGraphReq_Include) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0, 2} }

func init() {
	proto.RegisterType((*WalkGraphReq)(nil), "dm.WalkGraphReq")
	proto.RegisterType((*WalkGraphReq_Mode)(nil), "dm.WalkGraphReq.Mode")
	proto.RegisterType((*WalkGraphReq_Limit)(nil), "dm.WalkGraphReq.Limit")
	proto.RegisterType((*WalkGraphReq_Include)(nil), "dm.WalkGraphReq.Include")
	proto.RegisterEnum("dm.WalkGraphReq_Mode_Direction", WalkGraphReq_Mode_Direction_name, WalkGraphReq_Mode_Direction_value)
}

var fileDescriptor8 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xdd, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0xe9, 0x77, 0x7a, 0xfb, 0x41, 0xb0, 0x04, 0xca, 0x82, 0x60, 0x50, 0x01, 0x02, 0x09,
	0x65, 0x52, 0xe1, 0x95, 0x87, 0x8e, 0xf2, 0x31, 0x01, 0x9a, 0xf0, 0x26, 0xed, 0x31, 0x72, 0x6b,
	0xb7, 0x0b, 0x4b, 0x9a, 0x2c, 0x76, 0xb4, 0x8e, 0x07, 0xfe, 0x01, 0x9e, 0x79, 0xe2, 0x9f, 0xc5,
	0xbe, 0x76, 0x06, 0x62, 0x7b, 0xab, 0xcf, 0xf9, 0xf5, 0xc4, 0xf7, 0xf8, 0x82, 0x7f, 0xc1, 0xd2,
	0xb3, 0x78, 0x5d, 0xb2, 0xe2, 0x34, 0x2a, 0xca, 0x5c, 0xe5, 0xa4, 0xc9, 0xb3, 0xf0, 0xe1, 0x3a,
	0xcf, 0xd7, 0xa9, 0xd8, 0x43, 0x65, 0x51, 0xad, 0xf6, 0x78, 0x55, 0x32, 0x95, 0xe4, 0x1b, 0xcb,
	0x84, 0xbb, 0xff, 0xfb, 0x2a, 0xc9, 0x84, 0x54, 0x2c, 0x2b, 0x1c, 0xe0, 0x63, 0x62, 0xcc, 0x99,
	0x62, 0x4e, 0xb9, 0x63, 0x95, 0xf3, 0x4a, 0x94, 0x97, 0x4e, 0x1a, 0xa8, 0xcb, 0x42, 0x48, 0x7b,
	0x98, 0xfc, 0xea, 0xc2, 0xf0, 0x44, 0xdf, 0xe5, 0x83, 0xc1, 0xa8, 0x38, 0x27, 0xcf, 0xa0, 0xcd,
	0x2a, 0x75, 0x1a, 0x34, 0x1e, 0x35, 0x9e, 0x0f, 0xa6, 0x24, 0xe2, 0x59, 0xf4, 0x6e, 0x2b, 0x96,
	0x15, 0x5e, 0x63, 0xa6, 0x1d, 0x8a, 0x3e, 0x79, 0x02, 0x1d, 0x0c, 0x0d, 0x9a, 0x08, 0x8e, 0x0d,
	0x88, 0x21, 0x5f, 0x8d, 0x4a, 0xad, 0x49, 0x5e, 0x40, 0x3b, 0xcb, 0xb9, 0x08, 0x5a, 0x08, 0xdd,
	0x35, 0xd0, 0xbf, 0x5f, 0x8b, 0xbe, 0x68, 0x93, 0x22, 0x42, 0x5e, 0x42, 0x27, 0x4d, 0xb2, 0x44,
	0x05, 0x6d, 0x64, 0xef, 0x5d, 0x63, 0x3f, 0x1b, 0x97, 0x5a, 0x88, 0x4c, 0xa1, 0x97, 0x6c, 0x96,
	0x69, 0xa5, 0xb3, 0x3b, 0xc8, 0x07, 0xd7, 0xf8, 0x03, 0xeb, 0xd3, 0x1a, 0x0c, 0x7f, 0x36, 0xa0,
	0x6d, 0x3e, 0x48, 0x7c, 0x68, 0xf1, 0x95, 0xc4, 0x11, 0x3d, 0x6a, 0x7e, 0x92, 0x37, 0xd0, 0xe7,
	0x49, 0x29, 0x96, 0x66, 0x4a, 0x9c, 0x68, 0x3c, 0xdd, 0xbd, 0xf1, 0xb2, 0xd1, 0xbc, 0xc6, 0xe8,
	0xdf, 0x7f, 0x4c, 0xa6, 0xd0, 0xbf, 0xd2, 0xc9, 0x10, 0xbc, 0xf7, 0x87, 0xf4, 0x64, 0x46, 0xe7,
	0x47, 0xfe, 0x2d, 0x32, 0x82, 0xfe, 0xfe, 0xec, 0xed, 0x27, 0x7b, 0x6c, 0x10, 0x0f, 0xda, 0xfb,
	0x87, 0xc7, 0x1f, 0xfd, 0x66, 0xf8, 0x03, 0x3a, 0x38, 0x11, 0xb9, 0x0f, 0xfd, 0x8c, 0x6d, 0x63,
	0x2e, 0x0a, 0x57, 0x7b, 0x8b, 0x7a, 0x5a, 0x98, 0x9b, 0x33, 0x79, 0x0d, 0xe6, 0x77, 0x6c, 0x1e,
	0xda, 0x35, 0xbd, 0x13, 0xd9, 0x2d, 0x88, 0xea, 0x2d, 0x88, 0xe6, 0x6e, 0x4b, 0x68, 0x4f, 0xa3,
	0xc7, 0x9a, 0x24, 0x13, 0x18, 0x61, 0xa4, 0xde, 0x83, 0x58, 0x26, 0xdf, 0x6d, 0xff, 0x23, 0x3a,
	0x30, 0xb1, 0x5a, 0x3b, 0xd2, 0x52, 0xf8, 0xbb, 0x09, 0x3d, 0x57, 0x11, 0x79, 0x00, 0x90, 0x2f,
	0xbe, 0xe9, 0xfb, 0xc7, 0x09, 0xaf, 0x7b, 0xe9, 0x5b, 0xe5, 0x80, 0x4b, 0x63, 0xeb, 0xe7, 0x94,
	0x0a, 0x03, 0xf1, 0x1a, 0xda, 0x46, 0xc5, 0xa4, 0x91, 0xc7, 0x30, 0x64, 0x4a, 0x89, 0xac, 0x70,
	0x40, 0x0b, 0x81, 0x81, 0xd3, 0x10, 0x79, 0x0a, 0xe3, 0x1a, 0x29, 0x85, 0xac, 0x52, 0xfb, 0xca,
	0x1e, 0x1d, 0x39, 0x95, 0xa2, 0xa8, 0xd7, 0xc5, 0x17, 0xdb, 0x42, 0x37, 0xc9, 0x63, 0x67, 0x48,
	0x7c, 0x5e, 0x8f, 0xde, 0x76, 0xfa, 0xcc, 0xc9, 0x26, 0x71, 0x53, 0x65, 0xb1, 0xa8, 0x77, 0x53,
	0x06, 0x5d, 0x9c, 0x71, 0xa4, 0xd5, 0xab, 0x85, 0x95, 0x64, 0x07, 0xbc, 0xd5, 0x05, 0x37, 0xe5,
	0xca, 0xa0, 0x87, 0x49, 0x3d, 0x7d, 0xd6, 0xdd, 0x4a, 0xd3, 0xfb, 0x82, 0x2d, 0xcf, 0xac, 0xe7,
	0xa1, 0xe7, 0x19, 0xc1, 0x98, 0x8b, 0x2e, 0xb6, 0xfb, 0xea, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x81, 0x23, 0xf5, 0x35, 0xa9, 0x03, 0x00, 0x00,
}
