// Code generated by protoc-gen-go.
// source: logs.proto
// DO NOT EDIT!

/*
Package logdog is a generated protocol buffer package.

It is generated from these files:
	logs.proto
	state.proto

It has these top-level messages:
	GetRequest
	TailRequest
	GetResponse
	QueryRequest
	QueryResponse
	ListRequest
	ListResponse
	LogStreamState
*/
package logdog

import prpccommon "github.com/luci/luci-go/common/prpc"
import prpc "github.com/luci/luci-go/server/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import logpb "github.com/luci/luci-go/common/proto/logdog/logpb"
import google_protobuf "github.com/luci/luci-go/common/proto/google"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Trinary represents a trinary value.
type QueryRequest_Trinary int32

const (
	// Both positive and negative results will be returned.
	QueryRequest_BOTH QueryRequest_Trinary = 0
	// Only positive results will be returned.
	QueryRequest_YES QueryRequest_Trinary = 1
	// Only negative results will be returned.
	QueryRequest_NO QueryRequest_Trinary = 2
)

var QueryRequest_Trinary_name = map[int32]string{
	0: "BOTH",
	1: "YES",
	2: "NO",
}
var QueryRequest_Trinary_value = map[string]int32{
	"BOTH": 0,
	"YES":  1,
	"NO":   2,
}

func (x QueryRequest_Trinary) String() string {
	return proto.EnumName(QueryRequest_Trinary_name, int32(x))
}
func (QueryRequest_Trinary) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type ListResponse_Component_Type int32

const (
	ListResponse_Component_PATH    ListResponse_Component_Type = 0
	ListResponse_Component_STREAM  ListResponse_Component_Type = 1
	ListResponse_Component_PROJECT ListResponse_Component_Type = 2
)

var ListResponse_Component_Type_name = map[int32]string{
	0: "PATH",
	1: "STREAM",
	2: "PROJECT",
}
var ListResponse_Component_Type_value = map[string]int32{
	"PATH":    0,
	"STREAM":  1,
	"PROJECT": 2,
}

func (x ListResponse_Component_Type) String() string {
	return proto.EnumName(ListResponse_Component_Type_name, int32(x))
}
func (ListResponse_Component_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0, 0}
}

// GetRequest is the request structure for the user Get endpoint.
//
// If the requested log stream exists, a valid GetRequest will succeed
// regardless of whether the requested log range was available.
//
// Note that this endpoint may return fewer logs than requested due to either
// availability or internal constraints.
type GetRequest struct {
	// The request project to request.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// The path of the log stream to get.
	//
	// This can either be a LogDog stream path or the SHA256 hash of a LogDog
	// stream path.
	//
	// Some utilities may find passing around a full LogDog path to be cumbersome
	// due to its length. They can opt to pass around the hash instead and
	// retrieve logs using it.
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	// If true, requests that the log stream's state is returned.
	State bool `protobuf:"varint,3,opt,name=state" json:"state,omitempty"`
	// The initial log stream index to retrieve.
	Index int64 `protobuf:"varint,4,opt,name=index" json:"index,omitempty"`
	// The maximum number of bytes to return. If non-zero, it is applied as a
	// constraint to limit the number of logs that are returned.
	//
	// This only returns complete logs. Assuming logs are available, it will
	// return at least one log (even if it violates the size constraint) and as
	// many additional logs as it can without exceeding this constraint.
	ByteCount int32 `protobuf:"varint,5,opt,name=byte_count,json=byteCount" json:"byte_count,omitempty"`
	// The maximum number of log records to request.
	//
	// If this value is zero, no count constraint will be applied. If this value
	// is less than zero, no log entries will be returned. This can be used to
	// fetch log stream descriptors without fetching any log records.
	LogCount int32 `protobuf:"varint,6,opt,name=log_count,json=logCount" json:"log_count,omitempty"`
	// If true, allows the range request to return non-contiguous records.
	//
	// A contiguous request (default) will iterate forwards from the supplied
	// Index and stop if either the end of stream is encountered or there is a
	// missing stream index. A NonContiguous request will remove the latter
	// condition.
	//
	// For example, say the log stream consists of:
	// [3, 4, 6, 7]
	//
	// A contiguous request with Index 3 will return: [3, 4], stopping because
	// 5 is missing. A non-contiguous request will return [3, 4, 6, 7].
	NonContiguous bool `protobuf:"varint,7,opt,name=non_contiguous,json=nonContiguous" json:"non_contiguous,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// TailRequest is the request structure for the user Tail endpoint. It returns
// the last log in a given log stream at the time of the request.
type TailRequest struct {
	// The request project to request.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// The path of the log stream to get.
	//
	// This can either be a LogDog stream path or the SHA256 hash of a LogDog
	// stream path.
	//
	// Some utilities may find passing around a full LogDog path to be cumbersome
	// due to its length. They can opt to pass around the hash instead and
	// retrieve logs using it.
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	// If true, requests that the log stream's state is returned.
	State bool `protobuf:"varint,3,opt,name=state" json:"state,omitempty"`
}

func (m *TailRequest) Reset()                    { *m = TailRequest{} }
func (m *TailRequest) String() string            { return proto.CompactTextString(m) }
func (*TailRequest) ProtoMessage()               {}
func (*TailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// GetResponse is the response structure for the user Get endpoint.
type GetResponse struct {
	// Project is the project name that these logs belong to.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// The log stream descriptor and state for this stream.
	//
	// It can be requested by setting the request's State field to true. If the
	// Proto field is true, the State's Descriptor field will not be included.
	State *LogStreamState `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	// The expanded LogStreamDescriptor protobuf. It is intended for JSON
	// consumption.
	//
	// If the GetRequest's Proto field is false, this will be populated;
	// otherwise, the serialized protobuf will be written to the DescriptorProto
	// field.
	Desc *logpb.LogStreamDescriptor `protobuf:"bytes,3,opt,name=desc" json:"desc,omitempty"`
	// Log represents the set of retrieved log records.
	Logs []*logpb.LogEntry `protobuf:"bytes,4,rep,name=logs" json:"logs,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetResponse) GetState() *LogStreamState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *GetResponse) GetDesc() *logpb.LogStreamDescriptor {
	if m != nil {
		return m.Desc
	}
	return nil
}

func (m *GetResponse) GetLogs() []*logpb.LogEntry {
	if m != nil {
		return m.Logs
	}
	return nil
}

// QueryRequest is the request structure for the user Query endpoint.
type QueryRequest struct {
	// The request project to request.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// The query parameter.
	//
	// The path expression may substitute a glob ("*") for a specific path
	// component. That is, any stream that matches the remaining structure qualifies
	// regardless of its value in that specific positional field.
	//
	// An unbounded wildcard may appear as a component at the end of both the
	// prefix and name query components. "**" matches all remaining components.
	//
	// If the supplied path query does not contain a path separator ("+"), it will
	// be treated as if the prefix is "**".
	//
	// Examples:
	//   - Empty ("") will return all streams.
	//   - **/+/** will return all streams.
	//   - foo/bar/** will return all streams with the "foo/bar" prefix.
	//   - foo/bar/**/+/baz will return all streams beginning with the "foo/bar"
	//     prefix and named "baz" (e.g., "foo/bar/qux/lol/+/baz")
	//   - foo/bar/+/** will return all streams with a "foo/bar" prefix.
	//   - foo/*/+/baz will return all streams with a two-component prefix whose
	//     first value is "foo" and whose name is "baz".
	//   - foo/bar will return all streams whose name is "foo/bar".
	//   - */* will return all streams with two-component names.
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	// If true, returns that the streams' full state is returned instead of just
	// its Path.
	State bool `protobuf:"varint,3,opt,name=state" json:"state,omitempty"`
	// If true, causes the requested state to be returned as serialized protobuf
	// data instead of deserialized JSON structures.
	Proto bool `protobuf:"varint,4,opt,name=proto" json:"proto,omitempty"`
	// Next, if not empty, indicates that this query should continue at the point
	// where the previous query left off.
	Next string `protobuf:"bytes,5,opt,name=next" json:"next,omitempty"`
	// MaxResults is the maximum number of query results to return.
	//
	// If MaxResults is zero, no upper bound will be indicated. However, the
	// returned result count is still be subject to internal constraints.
	MaxResults int32 `protobuf:"varint,6,opt,name=max_results,json=maxResults" json:"max_results,omitempty"`
	// ContentType, if not empty, restricts results to streams with the supplied
	// content type.
	ContentType string                         `protobuf:"bytes,10,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	StreamType  *QueryRequest_StreamTypeFilter `protobuf:"bytes,11,opt,name=stream_type,json=streamType" json:"stream_type,omitempty"`
	// Newer restricts results to streams created after the specified date.
	Newer *google_protobuf.Timestamp `protobuf:"bytes,12,opt,name=newer" json:"newer,omitempty"`
	// Older restricts results to streams created before the specified date.
	Older *google_protobuf.Timestamp `protobuf:"bytes,13,opt,name=older" json:"older,omitempty"`
	// If not empty, constrains the results to those whose protobuf version string
	// matches the supplied version.
	ProtoVersion string `protobuf:"bytes,14,opt,name=proto_version,json=protoVersion" json:"proto_version,omitempty"`
	// Tags is the set of tags to constrain the query with.
	//
	// A Tag entry may either be:
	// - A key/value query, in which case the results are constrained by logs
	//   whose tag includes that key/value pair.
	// - A key with an missing (nil) value, in which case the results are
	//   constraints by logs that have that tag key, regardless of its value.
	Tags map[string]string `protobuf:"bytes,15,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Purged restricts the query to streams that have or haven't been purged.
	Purged QueryRequest_Trinary `protobuf:"varint,16,opt,name=purged,enum=logdog.QueryRequest_Trinary" json:"purged,omitempty"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QueryRequest) GetStreamType() *QueryRequest_StreamTypeFilter {
	if m != nil {
		return m.StreamType
	}
	return nil
}

func (m *QueryRequest) GetNewer() *google_protobuf.Timestamp {
	if m != nil {
		return m.Newer
	}
	return nil
}

func (m *QueryRequest) GetOlder() *google_protobuf.Timestamp {
	if m != nil {
		return m.Older
	}
	return nil
}

func (m *QueryRequest) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// The stream type to filter on.
type QueryRequest_StreamTypeFilter struct {
	// The StreamType value to filter on.
	Value logpb.StreamType `protobuf:"varint,1,opt,name=value,enum=logpb.StreamType" json:"value,omitempty"`
}

func (m *QueryRequest_StreamTypeFilter) Reset()         { *m = QueryRequest_StreamTypeFilter{} }
func (m *QueryRequest_StreamTypeFilter) String() string { return proto.CompactTextString(m) }
func (*QueryRequest_StreamTypeFilter) ProtoMessage()    {}
func (*QueryRequest_StreamTypeFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

// QueryResponse is the response structure for the user Query endpoint.
type QueryResponse struct {
	// Project is the project name that all responses belong to.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// The list of streams that were identified as the result of the query.
	Streams []*QueryResponse_Stream `protobuf:"bytes,2,rep,name=streams" json:"streams,omitempty"`
	// If not empty, indicates that there are more query results available.
	// These results can be requested by repeating the Query request with the
	// same Path field and supplying this value in the Next field.
	Next string `protobuf:"bytes,3,opt,name=next" json:"next,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *QueryResponse) GetStreams() []*QueryResponse_Stream {
	if m != nil {
		return m.Streams
	}
	return nil
}

// Stream represents a single query response stream.
type QueryResponse_Stream struct {
	// Path is the log stream path.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// State is the log stream descriptor and state for this stream.
	//
	// It can be requested by setting the request's State field to true. If the
	// Proto field is true, the State's Descriptor field will not be included.
	State *LogStreamState `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	// The JSON-packed log stream descriptor protobuf.
	//
	// A Descriptor entry corresponds to the Path with the same index.
	//
	// If the query request's State field is set, the descriptor will be
	// populated. If the Proto field is false, Descriptor will be populated;
	// otherwise, DescriptorProto will be populated with the serialized descriptor
	// protobuf.
	Desc *logpb.LogStreamDescriptor `protobuf:"bytes,3,opt,name=desc" json:"desc,omitempty"`
	// The serialized log stream Descriptor protobuf.
	DescProto []byte `protobuf:"bytes,4,opt,name=desc_proto,json=descProto,proto3" json:"desc_proto,omitempty"`
}

func (m *QueryResponse_Stream) Reset()                    { *m = QueryResponse_Stream{} }
func (m *QueryResponse_Stream) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse_Stream) ProtoMessage()               {}
func (*QueryResponse_Stream) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *QueryResponse_Stream) GetState() *LogStreamState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *QueryResponse_Stream) GetDesc() *logpb.LogStreamDescriptor {
	if m != nil {
		return m.Desc
	}
	return nil
}

// ListRequest is the request structure for the user List endpoint.
//
// The List endpoint enables a directory-tree style traversal of the project
// and log stream space.
//
// For example, if project "myproj" had streams "a/+/baz", a listing would
// return:
// - Request: project="", path="", Response: {myproj} (Project names)
// - Request: project="myproj", path="", Response: {a} (Path components)
// - Request: project="myproj", path="a", Response: {+} (Path component)
// - Request: project="myproj", path="a/+", Response: {baz} (Stream component)
type ListRequest struct {
	// The project to query.
	//
	// If this is empty, the list results will show all projects that the user can
	// access.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// The path base to query.
	//
	// For example, log streams:
	// - foo/bar/+/baz
	// - foo/+/qux
	//
	// A query to path_base "foo" will return "bar" and "+", both path
	// components.
	PathBase string `protobuf:"bytes,2,opt,name=path_base,json=pathBase" json:"path_base,omitempty"`
	// State, if true, returns that the streams' full state instead of just its
	// Path.
	State bool `protobuf:"varint,3,opt,name=state" json:"state,omitempty"`
	// If not empty, indicates that this query should continue at the point where
	// the previous query left off.
	Next string `protobuf:"bytes,4,opt,name=next" json:"next,omitempty"`
	// If true, will return only streams. Otherwise, intermediate path components
	// will also be returned.
	StreamOnly bool `protobuf:"varint,5,opt,name=stream_only,json=streamOnly" json:"stream_only,omitempty"`
	// If true, indicates that purged streams should show up in the listing. It is
	// an error if a non-admin user requests this option.
	IncludePurged bool `protobuf:"varint,6,opt,name=include_purged,json=includePurged" json:"include_purged,omitempty"`
	// Offset, if >= 0, instructs the list operation to skip the supplied number
	// of results. This can be used for pagination.
	Offset int32 `protobuf:"varint,7,opt,name=offset" json:"offset,omitempty"`
	// The maximum number of componts to return.
	//
	// If <= 0, no upper bound will be indicated. However, the returned result
	// count is still be subject to internal constraints.
	MaxResults int32 `protobuf:"varint,8,opt,name=max_results,json=maxResults" json:"max_results,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// ListResponse is the response structure for the user List endpoint.
type ListResponse struct {
	// The project that the streams belong to.
	//
	// If the request was for the project tier of the list hierarchy, this will
	// be empty, and the components list will contain project elements.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// The hierarchy base that was requested.
	PathBase string `protobuf:"bytes,2,opt,name=path_base,json=pathBase" json:"path_base,omitempty"`
	// If not empty, indicates that there are more list results available.
	// These results can be requested by repeating the List request with the
	// same Path field and supplying this value in the Next field.
	Next       string                    `protobuf:"bytes,3,opt,name=next" json:"next,omitempty"`
	Components []*ListResponse_Component `protobuf:"bytes,4,rep,name=components" json:"components,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListResponse) GetComponents() []*ListResponse_Component {
	if m != nil {
		return m.Components
	}
	return nil
}

// The set of listed stream components.
type ListResponse_Component struct {
	// Name is the name of this path component.  When combined with the
	// response Base, this will form the full stream path.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The type of the component.
	Type ListResponse_Component_Type `protobuf:"varint,2,opt,name=type,enum=logdog.ListResponse_Component_Type" json:"type,omitempty"`
	// State is the log stream descriptor and state for this stream. It will
	// only be filled if this is a STREAM component.
	//
	// It can be requested by setting the request's State field to true. If the
	// Proto field is true, the State's Descriptor field will not be included.
	State *LogStreamState `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	// Descriptor is the JSON-packed log stream descriptor protobuf. It will
	// only be filled if this is a STREAM component.
	//
	// A Descriptor entry corresponds to the Path with the same index.
	//
	// If the list request's State field is set, the descriptor will be
	// populated.
	Desc *logpb.LogStreamDescriptor `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
}

func (m *ListResponse_Component) Reset()                    { *m = ListResponse_Component{} }
func (m *ListResponse_Component) String() string            { return proto.CompactTextString(m) }
func (*ListResponse_Component) ProtoMessage()               {}
func (*ListResponse_Component) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *ListResponse_Component) GetState() *LogStreamState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *ListResponse_Component) GetDesc() *logpb.LogStreamDescriptor {
	if m != nil {
		return m.Desc
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "logdog.GetRequest")
	proto.RegisterType((*TailRequest)(nil), "logdog.TailRequest")
	proto.RegisterType((*GetResponse)(nil), "logdog.GetResponse")
	proto.RegisterType((*QueryRequest)(nil), "logdog.QueryRequest")
	proto.RegisterType((*QueryRequest_StreamTypeFilter)(nil), "logdog.QueryRequest.StreamTypeFilter")
	proto.RegisterType((*QueryResponse)(nil), "logdog.QueryResponse")
	proto.RegisterType((*QueryResponse_Stream)(nil), "logdog.QueryResponse.Stream")
	proto.RegisterType((*ListRequest)(nil), "logdog.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "logdog.ListResponse")
	proto.RegisterType((*ListResponse_Component)(nil), "logdog.ListResponse.Component")
	proto.RegisterEnum("logdog.QueryRequest_Trinary", QueryRequest_Trinary_name, QueryRequest_Trinary_value)
	proto.RegisterEnum("logdog.ListResponse_Component_Type", ListResponse_Component_Type_name, ListResponse_Component_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Logs service

type LogsClient interface {
	// Get returns state and log data for a single log stream.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Tail returns the last log in the log stream at the time of the request.
	Tail(ctx context.Context, in *TailRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Query returns log stream paths that match the requested query.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// List returns log stream paths rooted under the path hierarchy.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}
type logsPRPCClient struct {
	client *prpccommon.Client
}

func NewLogsPRPCClient(client *prpccommon.Client) LogsClient {
	return &logsPRPCClient{client}
}

func (c *logsPRPCClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.client.Call(ctx, "logdog.Logs", "Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logsPRPCClient) Tail(ctx context.Context, in *TailRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.client.Call(ctx, "logdog.Logs", "Tail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logsPRPCClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.client.Call(ctx, "logdog.Logs", "Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logsPRPCClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.client.Call(ctx, "logdog.Logs", "List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type logsClient struct {
	cc *grpc.ClientConn
}

func NewLogsClient(cc *grpc.ClientConn) LogsClient {
	return &logsClient{cc}
}

func (c *logsClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/logdog.Logs/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logsClient) Tail(ctx context.Context, in *TailRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/logdog.Logs/Tail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logsClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := grpc.Invoke(ctx, "/logdog.Logs/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logsClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/logdog.Logs/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Logs service

type LogsServer interface {
	// Get returns state and log data for a single log stream.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Tail returns the last log in the log stream at the time of the request.
	Tail(context.Context, *TailRequest) (*GetResponse, error)
	// Query returns log stream paths that match the requested query.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	// List returns log stream paths rooted under the path hierarchy.
	List(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterLogsServer(s prpc.Registrar, srv LogsServer) {
	s.RegisterService(&_Logs_serviceDesc, srv)
}

func _Logs_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logdog.Logs/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logs_Tail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServer).Tail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logdog.Logs/Tail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServer).Tail(ctx, req.(*TailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logs_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logdog.Logs/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logs_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logdog.Logs/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Logs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logdog.Logs",
	HandlerType: (*LogsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Logs_Get_Handler,
		},
		{
			MethodName: "Tail",
			Handler:    _Logs_Tail_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Logs_Query_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Logs_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 954 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xc6, 0x89, 0xf3, 0x77, 0x9c, 0xdd, 0x86, 0x61, 0xa9, 0x2c, 0xf3, 0x53, 0x48, 0xa9, 0x00,
	0x09, 0x9c, 0x12, 0x2a, 0x8a, 0xa8, 0x84, 0xd4, 0x2e, 0x5b, 0x10, 0x2a, 0xec, 0x76, 0x36, 0x42,
	0xe2, 0x2a, 0x72, 0x92, 0x59, 0xd7, 0xe0, 0x78, 0x82, 0x3d, 0x2e, 0x9b, 0xc7, 0xe0, 0x0a, 0xf1,
	0x0a, 0xbc, 0x09, 0x97, 0x3c, 0x02, 0x4f, 0xc0, 0x13, 0x20, 0x71, 0xe6, 0xcc, 0x38, 0x49, 0xa3,
	0xb0, 0x5b, 0xa4, 0xed, 0xcd, 0x66, 0xe6, 0x3b, 0x3f, 0x33, 0x73, 0xce, 0xf7, 0x1d, 0x2f, 0x40,
	0x2a, 0xe3, 0x22, 0x5c, 0xe4, 0x52, 0x49, 0xd6, 0xc4, 0xf5, 0x4c, 0xc6, 0x81, 0x57, 0xa8, 0x48,
	0x09, 0x03, 0x06, 0xf7, 0xe2, 0x44, 0x3d, 0x29, 0x27, 0xe1, 0x54, 0xce, 0x07, 0x69, 0x39, 0x4d,
	0xe8, 0xcf, 0x87, 0xb1, 0x1c, 0x20, 0x30, 0x97, 0xd9, 0x80, 0xbc, 0x06, 0x26, 0x52, 0xff, 0x2c,
	0x26, 0xfa, 0xaf, 0x0d, 0xbe, 0x11, 0x4b, 0x19, 0xa7, 0xc2, 0x38, 0x4d, 0xca, 0xb3, 0x81, 0x4a,
	0xe6, 0x02, 0xb3, 0xcf, 0x17, 0xc6, 0xa1, 0xff, 0x87, 0x03, 0xf0, 0xa5, 0x50, 0x5c, 0xfc, 0x54,
	0x22, 0xce, 0x7c, 0x68, 0x21, 0xfe, 0x83, 0x98, 0x2a, 0xdf, 0x79, 0xcb, 0x79, 0xaf, 0xc3, 0xab,
	0x2d, 0x63, 0xe0, 0x2e, 0x22, 0xf5, 0xc4, 0xaf, 0x11, 0x4c, 0x6b, 0x76, 0x00, 0x0d, 0xba, 0xa9,
	0x5f, 0x47, 0xb0, 0xcd, 0xcd, 0x46, 0xa3, 0x49, 0x36, 0x13, 0xe7, 0xbe, 0x8b, 0x68, 0x9d, 0x9b,
	0x0d, 0x7b, 0x03, 0x60, 0xb2, 0x54, 0x62, 0x3c, 0x95, 0x65, 0xa6, 0xfc, 0x06, 0x9a, 0x1a, 0xbc,
	0xa3, 0x91, 0x43, 0x0d, 0xb0, 0xd7, 0xa0, 0x83, 0xb7, 0xb6, 0xd6, 0x26, 0x59, 0xdb, 0x08, 0x18,
	0xe3, 0x2d, 0xd8, 0xcf, 0x64, 0x86, 0xc6, 0x4c, 0x25, 0x71, 0x29, 0xcb, 0xc2, 0x6f, 0xd1, 0x81,
	0x7b, 0x88, 0x1e, 0xae, 0xc0, 0xfe, 0x63, 0xf0, 0x46, 0x51, 0x92, 0x5e, 0xe1, 0x5b, 0xfa, 0xbf,
	0x3b, 0xe0, 0x51, 0x79, 0x8a, 0x85, 0xcc, 0x0a, 0x71, 0x41, 0xce, 0x0f, 0xaa, 0x78, 0x9d, 0xd4,
	0x1b, 0x5e, 0x0f, 0x4d, 0x47, 0xc2, 0x47, 0x32, 0x3e, 0x55, 0xb9, 0x88, 0xe6, 0xa7, 0xda, 0x5a,
	0xd5, 0x28, 0x04, 0x77, 0x26, 0x8a, 0x29, 0x1d, 0xe6, 0x0d, 0x83, 0x90, 0xfa, 0xb6, 0xf6, 0xfd,
	0x02, 0x6d, 0x79, 0xb2, 0x50, 0x32, 0xe7, 0xe4, 0xc7, 0x6e, 0x82, 0xab, 0x79, 0x82, 0x25, 0xad,
	0xa3, 0xff, 0xb5, 0xb5, 0xff, 0x51, 0xa6, 0xf2, 0x25, 0x27, 0x63, 0xff, 0xd7, 0x06, 0x74, 0x1f,
	0x97, 0x02, 0xf7, 0x57, 0xdb, 0x4d, 0x62, 0x0a, 0x75, 0x13, 0x51, 0xc3, 0x54, 0x8c, 0xcf, 0xc4,
	0xb9, 0xe9, 0x23, 0xc6, 0xeb, 0x35, 0xbb, 0x01, 0xde, 0x3c, 0x3a, 0x1f, 0xe7, 0xa2, 0x28, 0x53,
	0x55, 0xd8, 0x26, 0x02, 0x42, 0xdc, 0x20, 0xec, 0x6d, 0xe8, 0xea, 0x16, 0x8a, 0x4c, 0x8d, 0xd5,
	0x72, 0x21, 0x7c, 0xa0, 0x60, 0xcf, 0x62, 0x23, 0x84, 0xd8, 0x43, 0x40, 0xee, 0xeb, 0x0a, 0x18,
	0x0f, 0x8f, 0xca, 0x73, 0xab, 0xaa, 0xe5, 0xe6, 0xe3, 0x42, 0x53, 0x29, 0x1d, 0xf5, 0x30, 0x49,
	0x95, 0xc8, 0x39, 0x14, 0x2b, 0x84, 0xdd, 0x86, 0x46, 0x26, 0x7e, 0x16, 0xb9, 0xdf, 0xb5, 0x05,
	0x36, 0x3a, 0x08, 0x2b, 0x1d, 0x84, 0xa3, 0x4a, 0x07, 0xdc, 0x38, 0xea, 0x08, 0x99, 0xce, 0x30,
	0x62, 0xef, 0xf2, 0x08, 0x72, 0xc4, 0x9e, 0xec, 0x91, 0x71, 0xfc, 0x54, 0xe4, 0x45, 0x22, 0x33,
	0x7f, 0x9f, 0xde, 0xd3, 0x25, 0xf0, 0x3b, 0x83, 0xb1, 0x21, 0xb8, 0x2a, 0xc2, 0xc6, 0x5d, 0xa3,
	0xc6, 0xbd, 0xb9, 0xf3, 0x25, 0x23, 0x74, 0xb0, 0x7d, 0xd4, 0xbe, 0xec, 0x0e, 0x34, 0x17, 0x65,
	0x1e, 0x8b, 0x99, 0xdf, 0xc3, 0x8c, 0xfb, 0xc3, 0xd7, 0x77, 0x47, 0xe5, 0x49, 0x16, 0xe1, 0xd6,
	0xfa, 0x06, 0xf7, 0xa0, 0xb7, 0x5d, 0x12, 0xf6, 0x2e, 0x34, 0x9e, 0x46, 0x69, 0x29, 0xa8, 0xfd,
	0xfb, 0xc3, 0x97, 0x2d, 0x6f, 0xd6, 0x7e, 0xdc, 0xd8, 0x83, 0xbb, 0xd0, 0x59, 0xdd, 0x82, 0xf5,
	0xa0, 0xfe, 0xa3, 0x58, 0x5a, 0xca, 0xe8, 0xa5, 0x26, 0x81, 0xc9, 0x63, 0xf8, 0x62, 0x36, 0x9f,
	0xd5, 0x3e, 0x75, 0xfa, 0xef, 0x40, 0xcb, 0x5e, 0x84, 0xb5, 0xc1, 0x7d, 0x70, 0x3c, 0xfa, 0xaa,
	0xf7, 0x12, 0x6b, 0x41, 0xfd, 0xfb, 0xa3, 0xd3, 0x9e, 0xc3, 0x9a, 0x50, 0xfb, 0xf6, 0xb8, 0x57,
	0xeb, 0xff, 0x52, 0x83, 0x3d, 0x7b, 0xf9, 0x4b, 0x85, 0xf4, 0x09, 0xb4, 0x4c, 0x23, 0x0b, 0x3c,
	0x4d, 0x17, 0x6d, 0xfb, 0xf9, 0x26, 0x83, 0x7d, 0x04, 0xaf, 0x9c, 0x57, 0x94, 0xac, 0xaf, 0x29,
	0x19, 0xfc, 0xe6, 0x40, 0xd3, 0xf8, 0xad, 0x18, 0xef, 0x6c, 0x30, 0xfe, 0xc5, 0x6a, 0x16, 0x27,
	0x9e, 0xfe, 0x1d, 0xaf, 0xe5, 0xd3, 0xe5, 0x1d, 0x8d, 0x9c, 0xd0, 0xe4, 0xfd, 0x1b, 0x47, 0xcb,
	0xa3, 0xa4, 0x78, 0x8e, 0xd1, 0x8b, 0xb3, 0x51, 0x5f, 0x77, 0x3c, 0x89, 0x8a, 0xaa, 0x03, 0x6d,
	0x0d, 0x3c, 0xc0, 0xfd, 0x7f, 0xa8, 0xb6, 0x2a, 0x86, 0xfb, 0xac, 0x3e, 0xad, 0xb6, 0x64, 0x96,
	0x2e, 0x49, 0xba, 0xed, 0x4a, 0x34, 0xc7, 0x88, 0xe8, 0x31, 0x9b, 0x64, 0xd3, 0xb4, 0x9c, 0x89,
	0xb1, 0xe5, 0x5f, 0xd3, 0x8c, 0x59, 0x8b, 0x9e, 0x10, 0xc8, 0xae, 0x43, 0x53, 0x9e, 0x9d, 0x15,
	0x42, 0xd1, 0x14, 0x6e, 0x70, 0xbb, 0xdb, 0xd6, 0x7f, 0x7b, 0x5b, 0xff, 0xfd, 0x7f, 0x6a, 0xd0,
	0x35, 0x2f, 0xbe, 0x94, 0x04, 0x17, 0x3e, 0x79, 0x47, 0xa7, 0xd9, 0xe7, 0x00, 0xf8, 0x3d, 0xc4,
	0xb4, 0x38, 0x49, 0xaa, 0x31, 0xb9, 0x52, 0xdb, 0xe6, 0xa1, 0xe1, 0x61, 0xe5, 0xc6, 0x37, 0x22,
	0x82, 0xbf, 0x1c, 0xe8, 0xac, 0x2c, 0x74, 0x42, 0x34, 0x17, 0x15, 0x59, 0xf4, 0x9a, 0xdd, 0x45,
	0x25, 0xeb, 0x99, 0x54, 0x23, 0x29, 0xdd, 0xbc, 0x38, 0x77, 0x48, 0xe2, 0xa2, 0x80, 0x35, 0xcb,
	0xea, 0xff, 0x87, 0x65, 0xee, 0xf3, 0xb1, 0xac, 0xff, 0x3e, 0xb8, 0x34, 0xf1, 0x50, 0x7d, 0x27,
	0xf7, 0x49, 0x7d, 0x80, 0x9c, 0x1f, 0xf1, 0xa3, 0xfb, 0xdf, 0xa0, 0x00, 0x3d, 0x68, 0x9d, 0xf0,
	0xe3, 0xaf, 0x8f, 0x0e, 0x47, 0xbd, 0xda, 0xf0, 0x4f, 0x07, 0x5c, 0x4c, 0x54, 0xe0, 0x19, 0x75,
	0xfc, 0xa8, 0x31, 0x56, 0xdd, 0x64, 0xfd, 0x0f, 0x40, 0xf0, 0xca, 0x33, 0x98, 0xed, 0xd3, 0x6d,
	0x3c, 0x03, 0x3f, 0xac, 0x6c, 0x65, 0xdc, 0xf8, 0xcc, 0xee, 0x8e, 0xb8, 0x03, 0x0d, 0x52, 0x2b,
	0x3b, 0xd8, 0x35, 0xbb, 0x82, 0x57, 0x77, 0x4a, 0x9a, 0x7d, 0x84, 0xf7, 0xc3, 0x72, 0xae, 0xcf,
	0xd9, 0xd0, 0x47, 0x70, 0xb0, 0xab, 0xe2, 0x93, 0x26, 0xa9, 0xeb, 0xe3, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xca, 0x2d, 0x50, 0x3e, 0x47, 0x09, 0x00, 0x00,
}
