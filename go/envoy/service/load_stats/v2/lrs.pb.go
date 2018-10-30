// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/load_stats/v2/lrs.proto

package v2

import (
	fmt "fmt"
	core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
	endpoint "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/endpoint"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/lyft/protoc-gen-validate/validate"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A load report Envoy sends to the management server.
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
type LoadStatsRequest struct {
	// Node identifier for Envoy instance.
	Node *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// A list of load stats to report.
	ClusterStats         []*endpoint.ClusterStats `protobuf:"bytes,2,rep,name=cluster_stats,json=clusterStats,proto3" json:"cluster_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *LoadStatsRequest) Reset()         { *m = LoadStatsRequest{} }
func (m *LoadStatsRequest) String() string { return proto.CompactTextString(m) }
func (*LoadStatsRequest) ProtoMessage()    {}
func (*LoadStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7332d279836518, []int{0}
}

func (m *LoadStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadStatsRequest.Unmarshal(m, b)
}
func (m *LoadStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadStatsRequest.Marshal(b, m, deterministic)
}
func (m *LoadStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsRequest.Merge(m, src)
}
func (m *LoadStatsRequest) XXX_Size() int {
	return xxx_messageInfo_LoadStatsRequest.Size(m)
}
func (m *LoadStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsRequest proto.InternalMessageInfo

func (m *LoadStatsRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *LoadStatsRequest) GetClusterStats() []*endpoint.ClusterStats {
	if m != nil {
		return m.ClusterStats
	}
	return nil
}

// The management server sends envoy a LoadStatsResponse with all clusters it
// is interested in learning load stats about.
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
type LoadStatsResponse struct {
	// Clusters to report stats for.
	Clusters []string `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	// The minimum interval of time to collect stats over. This is only a minimum for two reasons:
	// 1. There may be some delay from when the timer fires until stats sampling occurs.
	// 2. For clusters that were already feature in the previous *LoadStatsResponse*, any traffic
	//    that is observed in between the corresponding previous *LoadStatsRequest* and this
	//    *LoadStatsResponse* will also be accumulated and billed to the cluster. This avoids a period
	//    of inobservability that might otherwise exists between the messages. New clusters are not
	//    subject to this consideration.
	LoadReportingInterval *duration.Duration `protobuf:"bytes,2,opt,name=load_reporting_interval,json=loadReportingInterval,proto3" json:"load_reporting_interval,omitempty"`
	// Set to *true* if the management server supports endpoint granularity
	// report.
	ReportEndpointGranularity bool     `protobuf:"varint,3,opt,name=report_endpoint_granularity,json=reportEndpointGranularity,proto3" json:"report_endpoint_granularity,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *LoadStatsResponse) Reset()         { *m = LoadStatsResponse{} }
func (m *LoadStatsResponse) String() string { return proto.CompactTextString(m) }
func (*LoadStatsResponse) ProtoMessage()    {}
func (*LoadStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd7332d279836518, []int{1}
}

func (m *LoadStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadStatsResponse.Unmarshal(m, b)
}
func (m *LoadStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadStatsResponse.Marshal(b, m, deterministic)
}
func (m *LoadStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsResponse.Merge(m, src)
}
func (m *LoadStatsResponse) XXX_Size() int {
	return xxx_messageInfo_LoadStatsResponse.Size(m)
}
func (m *LoadStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsResponse proto.InternalMessageInfo

func (m *LoadStatsResponse) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *LoadStatsResponse) GetLoadReportingInterval() *duration.Duration {
	if m != nil {
		return m.LoadReportingInterval
	}
	return nil
}

func (m *LoadStatsResponse) GetReportEndpointGranularity() bool {
	if m != nil {
		return m.ReportEndpointGranularity
	}
	return false
}

func init() {
	proto.RegisterType((*LoadStatsRequest)(nil), "envoy.service.load_stats.v2.LoadStatsRequest")
	proto.RegisterType((*LoadStatsResponse)(nil), "envoy.service.load_stats.v2.LoadStatsResponse")
}

func init() {
	proto.RegisterFile("envoy/service/load_stats/v2/lrs.proto", fileDescriptor_cd7332d279836518)
}

var fileDescriptor_cd7332d279836518 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0xca, 0xd4, 0x30,
	0x14, 0xc5, 0x4d, 0xeb, 0x9f, 0x31, 0x9f, 0xa2, 0x06, 0x65, 0xfa, 0xcd, 0x88, 0x94, 0x11, 0xb5,
	0x20, 0x26, 0x52, 0xf7, 0x2e, 0x46, 0x45, 0x85, 0x41, 0xb0, 0xb3, 0x73, 0x53, 0x32, 0xed, 0xb5,
	0x04, 0x6a, 0x52, 0x93, 0x34, 0x30, 0x6f, 0xa0, 0x1b, 0x17, 0x3e, 0x8e, 0x2b, 0x9f, 0xc0, 0xf7,
	0xf0, 0x2d, 0xa4, 0x4d, 0x3b, 0xd3, 0x71, 0x21, 0xee, 0x1a, 0xee, 0xef, 0xdc, 0x9c, 0x73, 0x1a,
	0xfc, 0x00, 0xa4, 0x53, 0x7b, 0x66, 0x40, 0x3b, 0x51, 0x00, 0xab, 0x15, 0x2f, 0x73, 0x63, 0xb9,
	0x35, 0xcc, 0xa5, 0xac, 0xd6, 0x86, 0x36, 0x5a, 0x59, 0x45, 0x96, 0x3d, 0x46, 0x07, 0x8c, 0x1e,
	0x31, 0xea, 0xd2, 0xc5, 0x5d, 0xbf, 0x83, 0x37, 0xa2, 0x13, 0x15, 0x4a, 0x03, 0xdb, 0x71, 0x03,
	0x5e, 0xba, 0x78, 0x74, 0x32, 0x05, 0x59, 0x36, 0x4a, 0x48, 0xeb, 0x6f, 0xd2, 0xd0, 0x28, 0x6d,
	0x07, 0xf0, 0x5e, 0xa5, 0x54, 0x55, 0x03, 0xeb, 0x4f, 0xbb, 0xf6, 0x23, 0x2b, 0x5b, 0xcd, 0xad,
	0x50, 0x72, 0x98, 0xcf, 0x1d, 0xaf, 0x45, 0xc9, 0x2d, 0xb0, 0xf1, 0xc3, 0x0f, 0x56, 0x5f, 0x11,
	0xbe, 0xb9, 0x51, 0xbc, 0xdc, 0x76, 0x86, 0x32, 0xf8, 0xdc, 0x82, 0xb1, 0xe4, 0x31, 0xbe, 0x28,
	0x55, 0x09, 0x11, 0x8a, 0x51, 0x72, 0x96, 0xce, 0xa9, 0x0f, 0xc0, 0x1b, 0x41, 0x5d, 0x4a, 0x3b,
	0x8f, 0xf4, 0x9d, 0x2a, 0x21, 0xeb, 0x21, 0xf2, 0x06, 0x5f, 0x2f, 0xea, 0xd6, 0x58, 0xd0, 0x3e,
	0x55, 0x14, 0xc4, 0x61, 0x72, 0x96, 0xde, 0x3f, 0x55, 0x8d, 0xde, 0xe9, 0x0b, 0xcf, 0xfa, 0xfb,
	0xae, 0x15, 0x93, 0xd3, 0xea, 0x17, 0xc2, 0xb7, 0x26, 0x5e, 0x4c, 0xa3, 0xa4, 0x01, 0xf2, 0x10,
	0xcf, 0x06, 0xca, 0x44, 0x28, 0x0e, 0x93, 0xab, 0x6b, 0xfc, 0xe3, 0xf7, 0xcf, 0xf0, 0xd2, 0x77,
	0x14, 0xcc, 0x50, 0x76, 0x98, 0x91, 0xf7, 0x78, 0x3e, 0xe9, 0x45, 0xc8, 0x2a, 0x17, 0xd2, 0x82,
	0x76, 0xbc, 0x8e, 0x82, 0x3e, 0xc7, 0x39, 0xf5, 0x25, 0xd1, 0xb1, 0x24, 0xfa, 0x72, 0x28, 0x29,
	0xbb, 0xd3, 0x29, 0xb3, 0x51, 0xf8, 0x76, 0xd0, 0x91, 0xe7, 0x78, 0xe9, 0xb7, 0xe5, 0xa3, 0xfd,
	0xbc, 0xd2, 0x5c, 0xb6, 0x35, 0xd7, 0xc2, 0xee, 0xa3, 0x30, 0x46, 0xc9, 0x2c, 0x3b, 0xf7, 0xc8,
	0xab, 0x81, 0x78, 0x7d, 0x04, 0xd2, 0x6f, 0x08, 0xdf, 0xde, 0x4c, 0x37, 0x6f, 0xfd, 0x1b, 0x20,
	0x0e, 0xdf, 0xd8, 0x5a, 0x0d, 0xfc, 0xd3, 0x21, 0x2e, 0x79, 0x42, 0xff, 0xf1, 0x4c, 0xe8, 0xdf,
	0xbf, 0x68, 0x41, 0xff, 0x17, 0xf7, 0x2d, 0xae, 0x2e, 0x24, 0xe8, 0x29, 0x5a, 0x5f, 0xf9, 0x10,
	0xb8, 0xf4, 0x0b, 0x42, 0xbb, 0xcb, 0x7d, 0x07, 0xcf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb6,
	0xeb, 0x58, 0x9b, 0xc3, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoadReportingServiceClient is the client API for LoadReportingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadReportingServiceClient interface {
	// Advanced API to allow for multi-dimensional load balancing by remote
	// server. For receiving LB assignments, the steps are:
	// 1, The management server is configured with per cluster/zone/load metric
	//    capacity configuration. The capacity configuration definition is
	//    outside of the scope of this document.
	// 2. Envoy issues a standard {Stream,Fetch}Endpoints request for the clusters
	//    to balance.
	//
	// Independently, Envoy will initiate a StreamLoadStats bidi stream with a
	// management server:
	// 1. Once a connection establishes, the management server publishes a
	//    LoadStatsResponse for all clusters it is interested in learning load
	//    stats about.
	// 2. For each cluster, Envoy load balances incoming traffic to upstream hosts
	//    based on per-zone weights and/or per-instance weights (if specified)
	//    based on intra-zone LbPolicy. This information comes from the above
	//    {Stream,Fetch}Endpoints.
	// 3. When upstream hosts reply, they optionally add header <define header
	//    name> with ASCII representation of EndpointLoadMetricStats.
	// 4. Envoy aggregates load reports over the period of time given to it in
	//    LoadStatsResponse.load_reporting_interval. This includes aggregation
	//    stats Envoy maintains by itself (total_requests, rpc_errors etc.) as
	//    well as load metrics from upstream hosts.
	// 5. When the timer of load_reporting_interval expires, Envoy sends new
	//    LoadStatsRequest filled with load reports for each cluster.
	// 6. The management server uses the load reports from all reported Envoys
	//    from around the world, computes global assignment and prepares traffic
	//    assignment destined for each zone Envoys are located in. Goto 2.
	StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error)
}

type loadReportingServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoadReportingServiceClient(cc *grpc.ClientConn) LoadReportingServiceClient {
	return &loadReportingServiceClient{cc}
}

func (c *loadReportingServiceClient) StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LoadReportingService_serviceDesc.Streams[0], "/envoy.service.load_stats.v2.LoadReportingService/StreamLoadStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &loadReportingServiceStreamLoadStatsClient{stream}
	return x, nil
}

type LoadReportingService_StreamLoadStatsClient interface {
	Send(*LoadStatsRequest) error
	Recv() (*LoadStatsResponse, error)
	grpc.ClientStream
}

type loadReportingServiceStreamLoadStatsClient struct {
	grpc.ClientStream
}

func (x *loadReportingServiceStreamLoadStatsClient) Send(m *LoadStatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsClient) Recv() (*LoadStatsResponse, error) {
	m := new(LoadStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoadReportingServiceServer is the server API for LoadReportingService service.
type LoadReportingServiceServer interface {
	// Advanced API to allow for multi-dimensional load balancing by remote
	// server. For receiving LB assignments, the steps are:
	// 1, The management server is configured with per cluster/zone/load metric
	//    capacity configuration. The capacity configuration definition is
	//    outside of the scope of this document.
	// 2. Envoy issues a standard {Stream,Fetch}Endpoints request for the clusters
	//    to balance.
	//
	// Independently, Envoy will initiate a StreamLoadStats bidi stream with a
	// management server:
	// 1. Once a connection establishes, the management server publishes a
	//    LoadStatsResponse for all clusters it is interested in learning load
	//    stats about.
	// 2. For each cluster, Envoy load balances incoming traffic to upstream hosts
	//    based on per-zone weights and/or per-instance weights (if specified)
	//    based on intra-zone LbPolicy. This information comes from the above
	//    {Stream,Fetch}Endpoints.
	// 3. When upstream hosts reply, they optionally add header <define header
	//    name> with ASCII representation of EndpointLoadMetricStats.
	// 4. Envoy aggregates load reports over the period of time given to it in
	//    LoadStatsResponse.load_reporting_interval. This includes aggregation
	//    stats Envoy maintains by itself (total_requests, rpc_errors etc.) as
	//    well as load metrics from upstream hosts.
	// 5. When the timer of load_reporting_interval expires, Envoy sends new
	//    LoadStatsRequest filled with load reports for each cluster.
	// 6. The management server uses the load reports from all reported Envoys
	//    from around the world, computes global assignment and prepares traffic
	//    assignment destined for each zone Envoys are located in. Goto 2.
	StreamLoadStats(LoadReportingService_StreamLoadStatsServer) error
}

func RegisterLoadReportingServiceServer(s *grpc.Server, srv LoadReportingServiceServer) {
	s.RegisterService(&_LoadReportingService_serviceDesc, srv)
}

func _LoadReportingService_StreamLoadStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadReportingServiceServer).StreamLoadStats(&loadReportingServiceStreamLoadStatsServer{stream})
}

type LoadReportingService_StreamLoadStatsServer interface {
	Send(*LoadStatsResponse) error
	Recv() (*LoadStatsRequest, error)
	grpc.ServerStream
}

type loadReportingServiceStreamLoadStatsServer struct {
	grpc.ServerStream
}

func (x *loadReportingServiceStreamLoadStatsServer) Send(m *LoadStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsServer) Recv() (*LoadStatsRequest, error) {
	m := new(LoadStatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LoadReportingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.load_stats.v2.LoadReportingService",
	HandlerType: (*LoadReportingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLoadStats",
			Handler:       _LoadReportingService_StreamLoadStats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/load_stats/v2/lrs.proto",
}