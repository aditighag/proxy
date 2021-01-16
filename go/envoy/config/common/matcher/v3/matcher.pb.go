// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/config/common/matcher/v3/matcher.proto

package envoy_config_common_matcher_v3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/route/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Match configuration. This is a recursive structure which allows complex nested match
// configurations to be built using various logical operators.
// [#next-free-field: 11]
type MatchPredicate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Rule:
	//	*MatchPredicate_OrMatch
	//	*MatchPredicate_AndMatch
	//	*MatchPredicate_NotMatch
	//	*MatchPredicate_AnyMatch
	//	*MatchPredicate_HttpRequestHeadersMatch
	//	*MatchPredicate_HttpRequestTrailersMatch
	//	*MatchPredicate_HttpResponseHeadersMatch
	//	*MatchPredicate_HttpResponseTrailersMatch
	//	*MatchPredicate_HttpRequestGenericBodyMatch
	//	*MatchPredicate_HttpResponseGenericBodyMatch
	Rule isMatchPredicate_Rule `protobuf_oneof:"rule"`
}

func (x *MatchPredicate) Reset() {
	*x = MatchPredicate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchPredicate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchPredicate) ProtoMessage() {}

func (x *MatchPredicate) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchPredicate.ProtoReflect.Descriptor instead.
func (*MatchPredicate) Descriptor() ([]byte, []int) {
	return file_envoy_config_common_matcher_v3_matcher_proto_rawDescGZIP(), []int{0}
}

func (m *MatchPredicate) GetRule() isMatchPredicate_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (x *MatchPredicate) GetOrMatch() *MatchPredicate_MatchSet {
	if x, ok := x.GetRule().(*MatchPredicate_OrMatch); ok {
		return x.OrMatch
	}
	return nil
}

func (x *MatchPredicate) GetAndMatch() *MatchPredicate_MatchSet {
	if x, ok := x.GetRule().(*MatchPredicate_AndMatch); ok {
		return x.AndMatch
	}
	return nil
}

func (x *MatchPredicate) GetNotMatch() *MatchPredicate {
	if x, ok := x.GetRule().(*MatchPredicate_NotMatch); ok {
		return x.NotMatch
	}
	return nil
}

func (x *MatchPredicate) GetAnyMatch() bool {
	if x, ok := x.GetRule().(*MatchPredicate_AnyMatch); ok {
		return x.AnyMatch
	}
	return false
}

func (x *MatchPredicate) GetHttpRequestHeadersMatch() *HttpHeadersMatch {
	if x, ok := x.GetRule().(*MatchPredicate_HttpRequestHeadersMatch); ok {
		return x.HttpRequestHeadersMatch
	}
	return nil
}

func (x *MatchPredicate) GetHttpRequestTrailersMatch() *HttpHeadersMatch {
	if x, ok := x.GetRule().(*MatchPredicate_HttpRequestTrailersMatch); ok {
		return x.HttpRequestTrailersMatch
	}
	return nil
}

func (x *MatchPredicate) GetHttpResponseHeadersMatch() *HttpHeadersMatch {
	if x, ok := x.GetRule().(*MatchPredicate_HttpResponseHeadersMatch); ok {
		return x.HttpResponseHeadersMatch
	}
	return nil
}

func (x *MatchPredicate) GetHttpResponseTrailersMatch() *HttpHeadersMatch {
	if x, ok := x.GetRule().(*MatchPredicate_HttpResponseTrailersMatch); ok {
		return x.HttpResponseTrailersMatch
	}
	return nil
}

func (x *MatchPredicate) GetHttpRequestGenericBodyMatch() *HttpGenericBodyMatch {
	if x, ok := x.GetRule().(*MatchPredicate_HttpRequestGenericBodyMatch); ok {
		return x.HttpRequestGenericBodyMatch
	}
	return nil
}

func (x *MatchPredicate) GetHttpResponseGenericBodyMatch() *HttpGenericBodyMatch {
	if x, ok := x.GetRule().(*MatchPredicate_HttpResponseGenericBodyMatch); ok {
		return x.HttpResponseGenericBodyMatch
	}
	return nil
}

type isMatchPredicate_Rule interface {
	isMatchPredicate_Rule()
}

type MatchPredicate_OrMatch struct {
	// A set that describes a logical OR. If any member of the set matches, the match configuration
	// matches.
	OrMatch *MatchPredicate_MatchSet `protobuf:"bytes,1,opt,name=or_match,json=orMatch,proto3,oneof"`
}

type MatchPredicate_AndMatch struct {
	// A set that describes a logical AND. If all members of the set match, the match configuration
	// matches.
	AndMatch *MatchPredicate_MatchSet `protobuf:"bytes,2,opt,name=and_match,json=andMatch,proto3,oneof"`
}

type MatchPredicate_NotMatch struct {
	// A negation match. The match configuration will match if the negated match condition matches.
	NotMatch *MatchPredicate `protobuf:"bytes,3,opt,name=not_match,json=notMatch,proto3,oneof"`
}

type MatchPredicate_AnyMatch struct {
	// The match configuration will always match.
	AnyMatch bool `protobuf:"varint,4,opt,name=any_match,json=anyMatch,proto3,oneof"`
}

type MatchPredicate_HttpRequestHeadersMatch struct {
	// HTTP request headers match configuration.
	HttpRequestHeadersMatch *HttpHeadersMatch `protobuf:"bytes,5,opt,name=http_request_headers_match,json=httpRequestHeadersMatch,proto3,oneof"`
}

type MatchPredicate_HttpRequestTrailersMatch struct {
	// HTTP request trailers match configuration.
	HttpRequestTrailersMatch *HttpHeadersMatch `protobuf:"bytes,6,opt,name=http_request_trailers_match,json=httpRequestTrailersMatch,proto3,oneof"`
}

type MatchPredicate_HttpResponseHeadersMatch struct {
	// HTTP response headers match configuration.
	HttpResponseHeadersMatch *HttpHeadersMatch `protobuf:"bytes,7,opt,name=http_response_headers_match,json=httpResponseHeadersMatch,proto3,oneof"`
}

type MatchPredicate_HttpResponseTrailersMatch struct {
	// HTTP response trailers match configuration.
	HttpResponseTrailersMatch *HttpHeadersMatch `protobuf:"bytes,8,opt,name=http_response_trailers_match,json=httpResponseTrailersMatch,proto3,oneof"`
}

type MatchPredicate_HttpRequestGenericBodyMatch struct {
	// HTTP request generic body match configuration.
	HttpRequestGenericBodyMatch *HttpGenericBodyMatch `protobuf:"bytes,9,opt,name=http_request_generic_body_match,json=httpRequestGenericBodyMatch,proto3,oneof"`
}

type MatchPredicate_HttpResponseGenericBodyMatch struct {
	// HTTP response generic body match configuration.
	HttpResponseGenericBodyMatch *HttpGenericBodyMatch `protobuf:"bytes,10,opt,name=http_response_generic_body_match,json=httpResponseGenericBodyMatch,proto3,oneof"`
}

func (*MatchPredicate_OrMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_AndMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_NotMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_AnyMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpRequestHeadersMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpRequestTrailersMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpResponseHeadersMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpResponseTrailersMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpRequestGenericBodyMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpResponseGenericBodyMatch) isMatchPredicate_Rule() {}

// HTTP headers match configuration.
type HttpHeadersMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HTTP headers to match.
	Headers []*v3.HeaderMatcher `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
}

func (x *HttpHeadersMatch) Reset() {
	*x = HttpHeadersMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpHeadersMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpHeadersMatch) ProtoMessage() {}

func (x *HttpHeadersMatch) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpHeadersMatch.ProtoReflect.Descriptor instead.
func (*HttpHeadersMatch) Descriptor() ([]byte, []int) {
	return file_envoy_config_common_matcher_v3_matcher_proto_rawDescGZIP(), []int{1}
}

func (x *HttpHeadersMatch) GetHeaders() []*v3.HeaderMatcher {
	if x != nil {
		return x.Headers
	}
	return nil
}

// HTTP generic body match configuration.
// List of text strings and hex strings to be located in HTTP body.
// All specified strings must be found in the HTTP body for positive match.
// The search may be limited to specified number of bytes from the body start.
//
// .. attention::
//
//   Searching for patterns in HTTP body is potentially cpu intensive. For each specified pattern, http body is scanned byte by byte to find a match.
//   If multiple patterns are specified, the process is repeated for each pattern. If location of a pattern is known, ``bytes_limit`` should be specified
//   to scan only part of the http body.
type HttpGenericBodyMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Limits search to specified number of bytes - default zero (no limit - match entire captured buffer).
	BytesLimit uint32 `protobuf:"varint,1,opt,name=bytes_limit,json=bytesLimit,proto3" json:"bytes_limit,omitempty"`
	// List of patterns to match.
	Patterns []*HttpGenericBodyMatch_GenericTextMatch `protobuf:"bytes,2,rep,name=patterns,proto3" json:"patterns,omitempty"`
}

func (x *HttpGenericBodyMatch) Reset() {
	*x = HttpGenericBodyMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpGenericBodyMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpGenericBodyMatch) ProtoMessage() {}

func (x *HttpGenericBodyMatch) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpGenericBodyMatch.ProtoReflect.Descriptor instead.
func (*HttpGenericBodyMatch) Descriptor() ([]byte, []int) {
	return file_envoy_config_common_matcher_v3_matcher_proto_rawDescGZIP(), []int{2}
}

func (x *HttpGenericBodyMatch) GetBytesLimit() uint32 {
	if x != nil {
		return x.BytesLimit
	}
	return 0
}

func (x *HttpGenericBodyMatch) GetPatterns() []*HttpGenericBodyMatch_GenericTextMatch {
	if x != nil {
		return x.Patterns
	}
	return nil
}

// A set of match configurations used for logical operations.
type MatchPredicate_MatchSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of rules that make up the set.
	Rules []*MatchPredicate `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *MatchPredicate_MatchSet) Reset() {
	*x = MatchPredicate_MatchSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchPredicate_MatchSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchPredicate_MatchSet) ProtoMessage() {}

func (x *MatchPredicate_MatchSet) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchPredicate_MatchSet.ProtoReflect.Descriptor instead.
func (*MatchPredicate_MatchSet) Descriptor() ([]byte, []int) {
	return file_envoy_config_common_matcher_v3_matcher_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MatchPredicate_MatchSet) GetRules() []*MatchPredicate {
	if x != nil {
		return x.Rules
	}
	return nil
}

type HttpGenericBodyMatch_GenericTextMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Rule:
	//	*HttpGenericBodyMatch_GenericTextMatch_StringMatch
	//	*HttpGenericBodyMatch_GenericTextMatch_BinaryMatch
	Rule isHttpGenericBodyMatch_GenericTextMatch_Rule `protobuf_oneof:"rule"`
}

func (x *HttpGenericBodyMatch_GenericTextMatch) Reset() {
	*x = HttpGenericBodyMatch_GenericTextMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpGenericBodyMatch_GenericTextMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpGenericBodyMatch_GenericTextMatch) ProtoMessage() {}

func (x *HttpGenericBodyMatch_GenericTextMatch) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpGenericBodyMatch_GenericTextMatch.ProtoReflect.Descriptor instead.
func (*HttpGenericBodyMatch_GenericTextMatch) Descriptor() ([]byte, []int) {
	return file_envoy_config_common_matcher_v3_matcher_proto_rawDescGZIP(), []int{2, 0}
}

func (m *HttpGenericBodyMatch_GenericTextMatch) GetRule() isHttpGenericBodyMatch_GenericTextMatch_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (x *HttpGenericBodyMatch_GenericTextMatch) GetStringMatch() string {
	if x, ok := x.GetRule().(*HttpGenericBodyMatch_GenericTextMatch_StringMatch); ok {
		return x.StringMatch
	}
	return ""
}

func (x *HttpGenericBodyMatch_GenericTextMatch) GetBinaryMatch() []byte {
	if x, ok := x.GetRule().(*HttpGenericBodyMatch_GenericTextMatch_BinaryMatch); ok {
		return x.BinaryMatch
	}
	return nil
}

type isHttpGenericBodyMatch_GenericTextMatch_Rule interface {
	isHttpGenericBodyMatch_GenericTextMatch_Rule()
}

type HttpGenericBodyMatch_GenericTextMatch_StringMatch struct {
	// Text string to be located in HTTP body.
	StringMatch string `protobuf:"bytes,1,opt,name=string_match,json=stringMatch,proto3,oneof"`
}

type HttpGenericBodyMatch_GenericTextMatch_BinaryMatch struct {
	// Sequence of bytes to be located in HTTP body.
	BinaryMatch []byte `protobuf:"bytes,2,opt,name=binary_match,json=binaryMatch,proto3,oneof"`
}

func (*HttpGenericBodyMatch_GenericTextMatch_StringMatch) isHttpGenericBodyMatch_GenericTextMatch_Rule() {
}

func (*HttpGenericBodyMatch_GenericTextMatch_BinaryMatch) isHttpGenericBodyMatch_GenericTextMatch_Rule() {
}

var File_envoy_config_common_matcher_v3_matcher_proto protoreflect.FileDescriptor

var file_envoy_config_common_matcher_v3_matcher_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33,
	0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x2c,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x08, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x54, 0x0a, 0x08, 0x6f, 0x72,
	0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x07, 0x6f, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x56, 0x0a, 0x09, 0x61, 0x6e, 0x64, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x08,
	0x61, 0x6e, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x4d, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x5f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x08, 0x6e,
	0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x26, 0x0a, 0x09, 0x61, 0x6e, 0x79, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x6a,
	0x02, 0x08, 0x01, 0x48, 0x00, 0x52, 0x08, 0x61, 0x6e, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x6f, 0x0a, 0x1a, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x48, 0x00, 0x52, 0x17, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x71, 0x0a, 0x1b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x48, 0x00, 0x52, 0x18, 0x68, 0x74, 0x74, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x71, 0x0a, 0x1b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x48, 0x00, 0x52, 0x18, 0x68, 0x74,
	0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x73, 0x0a, 0x1c, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73,
	0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x48, 0x00,
	0x52, 0x19, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72,
	0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x7c, 0x0a, 0x1f, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x42, 0x6f, 0x64, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x48, 0x00, 0x52, 0x1b, 0x68, 0x74,
	0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x42, 0x6f, 0x64, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x7e, 0x0a, 0x20, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x42, 0x6f, 0x64, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x48, 0x00, 0x52, 0x1c, 0x68, 0x74, 0x74,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x42, 0x6f, 0x64, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x5a, 0x0a, 0x08, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x65, 0x74, 0x12, 0x4e, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x02, 0x52, 0x05,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x42, 0x0b, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x12, 0x03, 0xf8,
	0x42, 0x01, 0x22, 0x52, 0x0a, 0x10, 0x48, 0x74, 0x74, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x3e, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x22, 0xa1, 0x02, 0x0a, 0x14, 0x48, 0x74, 0x74, 0x70, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x42, 0x6f, 0x64, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x1f, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x6b, 0x0a, 0x08, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x45, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x42,
	0x6f, 0x64, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x54, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01,
	0x02, 0x08, 0x01, 0x52, 0x08, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x1a, 0x7b, 0x0a,
	0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x2c, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x2c, 0x0a, 0x0c, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x7a, 0x02, 0x10, 0x01, 0x48, 0x00,
	0x52, 0x0b, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x0b, 0x0a,
	0x04, 0x72, 0x75, 0x6c, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x46, 0x0a, 0x2c, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x0c, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_common_matcher_v3_matcher_proto_rawDescOnce sync.Once
	file_envoy_config_common_matcher_v3_matcher_proto_rawDescData = file_envoy_config_common_matcher_v3_matcher_proto_rawDesc
)

func file_envoy_config_common_matcher_v3_matcher_proto_rawDescGZIP() []byte {
	file_envoy_config_common_matcher_v3_matcher_proto_rawDescOnce.Do(func() {
		file_envoy_config_common_matcher_v3_matcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_common_matcher_v3_matcher_proto_rawDescData)
	})
	return file_envoy_config_common_matcher_v3_matcher_proto_rawDescData
}

var file_envoy_config_common_matcher_v3_matcher_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_envoy_config_common_matcher_v3_matcher_proto_goTypes = []interface{}{
	(*MatchPredicate)(nil),                        // 0: envoy.config.common.matcher.v3.MatchPredicate
	(*HttpHeadersMatch)(nil),                      // 1: envoy.config.common.matcher.v3.HttpHeadersMatch
	(*HttpGenericBodyMatch)(nil),                  // 2: envoy.config.common.matcher.v3.HttpGenericBodyMatch
	(*MatchPredicate_MatchSet)(nil),               // 3: envoy.config.common.matcher.v3.MatchPredicate.MatchSet
	(*HttpGenericBodyMatch_GenericTextMatch)(nil), // 4: envoy.config.common.matcher.v3.HttpGenericBodyMatch.GenericTextMatch
	(*v3.HeaderMatcher)(nil),                      // 5: envoy.config.route.v3.HeaderMatcher
}
var file_envoy_config_common_matcher_v3_matcher_proto_depIdxs = []int32{
	3,  // 0: envoy.config.common.matcher.v3.MatchPredicate.or_match:type_name -> envoy.config.common.matcher.v3.MatchPredicate.MatchSet
	3,  // 1: envoy.config.common.matcher.v3.MatchPredicate.and_match:type_name -> envoy.config.common.matcher.v3.MatchPredicate.MatchSet
	0,  // 2: envoy.config.common.matcher.v3.MatchPredicate.not_match:type_name -> envoy.config.common.matcher.v3.MatchPredicate
	1,  // 3: envoy.config.common.matcher.v3.MatchPredicate.http_request_headers_match:type_name -> envoy.config.common.matcher.v3.HttpHeadersMatch
	1,  // 4: envoy.config.common.matcher.v3.MatchPredicate.http_request_trailers_match:type_name -> envoy.config.common.matcher.v3.HttpHeadersMatch
	1,  // 5: envoy.config.common.matcher.v3.MatchPredicate.http_response_headers_match:type_name -> envoy.config.common.matcher.v3.HttpHeadersMatch
	1,  // 6: envoy.config.common.matcher.v3.MatchPredicate.http_response_trailers_match:type_name -> envoy.config.common.matcher.v3.HttpHeadersMatch
	2,  // 7: envoy.config.common.matcher.v3.MatchPredicate.http_request_generic_body_match:type_name -> envoy.config.common.matcher.v3.HttpGenericBodyMatch
	2,  // 8: envoy.config.common.matcher.v3.MatchPredicate.http_response_generic_body_match:type_name -> envoy.config.common.matcher.v3.HttpGenericBodyMatch
	5,  // 9: envoy.config.common.matcher.v3.HttpHeadersMatch.headers:type_name -> envoy.config.route.v3.HeaderMatcher
	4,  // 10: envoy.config.common.matcher.v3.HttpGenericBodyMatch.patterns:type_name -> envoy.config.common.matcher.v3.HttpGenericBodyMatch.GenericTextMatch
	0,  // 11: envoy.config.common.matcher.v3.MatchPredicate.MatchSet.rules:type_name -> envoy.config.common.matcher.v3.MatchPredicate
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_envoy_config_common_matcher_v3_matcher_proto_init() }
func file_envoy_config_common_matcher_v3_matcher_proto_init() {
	if File_envoy_config_common_matcher_v3_matcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchPredicate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpHeadersMatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpGenericBodyMatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchPredicate_MatchSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpGenericBodyMatch_GenericTextMatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MatchPredicate_OrMatch)(nil),
		(*MatchPredicate_AndMatch)(nil),
		(*MatchPredicate_NotMatch)(nil),
		(*MatchPredicate_AnyMatch)(nil),
		(*MatchPredicate_HttpRequestHeadersMatch)(nil),
		(*MatchPredicate_HttpRequestTrailersMatch)(nil),
		(*MatchPredicate_HttpResponseHeadersMatch)(nil),
		(*MatchPredicate_HttpResponseTrailersMatch)(nil),
		(*MatchPredicate_HttpRequestGenericBodyMatch)(nil),
		(*MatchPredicate_HttpResponseGenericBodyMatch)(nil),
	}
	file_envoy_config_common_matcher_v3_matcher_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*HttpGenericBodyMatch_GenericTextMatch_StringMatch)(nil),
		(*HttpGenericBodyMatch_GenericTextMatch_BinaryMatch)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_common_matcher_v3_matcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_common_matcher_v3_matcher_proto_goTypes,
		DependencyIndexes: file_envoy_config_common_matcher_v3_matcher_proto_depIdxs,
		MessageInfos:      file_envoy_config_common_matcher_v3_matcher_proto_msgTypes,
	}.Build()
	File_envoy_config_common_matcher_v3_matcher_proto = out.File
	file_envoy_config_common_matcher_v3_matcher_proto_rawDesc = nil
	file_envoy_config_common_matcher_v3_matcher_proto_goTypes = nil
	file_envoy_config_common_matcher_v3_matcher_proto_depIdxs = nil
}
