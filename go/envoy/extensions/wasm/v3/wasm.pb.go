// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/extensions/wasm/v3/wasm.proto

package envoy_extensions_wasm_v3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// [[#not-implemented-hide:]
// Configuration for a Wasm VM.
// [#next-free-field: 7]
type VmConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An ID which will be used along with a hash of the wasm code (or the name of the registered Null
	// VM plugin) to determine which VM will be used for the plugin. All plugins which use the same
	// *vm_id* and code will use the same VM. May be left blank. Sharing a VM between plugins can
	// reduce memory utilization and make sharing of data easier which may have security implications.
	// See ref: "TODO: add ref" for details.
	VmId string `protobuf:"bytes,1,opt,name=vm_id,json=vmId,proto3" json:"vm_id,omitempty"`
	// The Wasm runtime type (either "v8" or "null" for code compiled into Envoy).
	Runtime string `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`
	// The Wasm code that Envoy will execute.
	Code *v3.AsyncDataSource `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	// The Wasm configuration used in initialization of a new VM
	// (proxy_on_start). `google.protobuf.Struct` is serialized as JSON before
	// passing it to the plugin. `google.protobuf.BytesValue` and
	// `google.protobuf.StringValue` are passed directly without the wrapper.
	Configuration *any.Any `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// Allow the wasm file to include pre-compiled code on VMs which support it.
	// Warning: this should only be enable for trusted sources as the precompiled code is not
	// verified.
	AllowPrecompiled bool `protobuf:"varint,5,opt,name=allow_precompiled,json=allowPrecompiled,proto3" json:"allow_precompiled,omitempty"`
	// If true and the code needs to be remotely fetched and it is not in the cache then NACK the configuration
	// update and do a background fetch to fill the cache, otherwise fetch the code asynchronously and enter
	// warming state.
	NackOnCodeCacheMiss bool `protobuf:"varint,6,opt,name=nack_on_code_cache_miss,json=nackOnCodeCacheMiss,proto3" json:"nack_on_code_cache_miss,omitempty"`
}

func (x *VmConfig) Reset() {
	*x = VmConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VmConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VmConfig) ProtoMessage() {}

func (x *VmConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VmConfig.ProtoReflect.Descriptor instead.
func (*VmConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_wasm_v3_wasm_proto_rawDescGZIP(), []int{0}
}

func (x *VmConfig) GetVmId() string {
	if x != nil {
		return x.VmId
	}
	return ""
}

func (x *VmConfig) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *VmConfig) GetCode() *v3.AsyncDataSource {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *VmConfig) GetConfiguration() *any.Any {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *VmConfig) GetAllowPrecompiled() bool {
	if x != nil {
		return x.AllowPrecompiled
	}
	return false
}

func (x *VmConfig) GetNackOnCodeCacheMiss() bool {
	if x != nil {
		return x.NackOnCodeCacheMiss
	}
	return false
}

// [[#not-implemented-hide:]
// Base Configuration for Wasm Plugins e.g. filters and services.
// [#next-free-field: 6]
type PluginConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique name for a filters/services in a VM for use in identifying the filter/service if
	// multiple filters/services are handled by the same *vm_id* and *root_id* and for
	// logging/debugging.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A unique ID for a set of filters/services in a VM which will share a RootContext and Contexts
	// if applicable (e.g. an Wasm HttpFilter and an Wasm AccessLog). If left blank, all
	// filters/services with a blank root_id with the same *vm_id* will share Context(s).
	RootId string `protobuf:"bytes,2,opt,name=root_id,json=rootId,proto3" json:"root_id,omitempty"`
	// Configuration for finding or starting VM.
	//
	// Types that are assignable to VmConfig:
	//	*PluginConfig_InlineVmConfig
	VmConfig isPluginConfig_VmConfig `protobuf_oneof:"vm_config"`
	// Filter/service configuration used to configure or reconfigure a plugin
	// (proxy_on_configuration).
	// `google.protobuf.Struct` is serialized as JSON before
	// passing it to the plugin. `google.protobuf.BytesValue` and
	// `google.protobuf.StringValue` are passed directly without the wrapper.
	Configuration *any.Any `protobuf:"bytes,4,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// If there is a fatal error on the VM (e.g. exception, abort(), on_start or on_configure return false),
	// then all plugins associated with the VM will either fail closed (by default), e.g. by returning an HTTP 503 error,
	// or fail open (if 'fail_open' is set to true) by bypassing the filter. Note: when on_start or on_configure return false
	// during xDS updates the xDS configuration will be rejected and when on_start or on_configuration return false on initial
	// startup the proxy will not start.
	FailOpen bool `protobuf:"varint,5,opt,name=fail_open,json=failOpen,proto3" json:"fail_open,omitempty"`
}

func (x *PluginConfig) Reset() {
	*x = PluginConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginConfig) ProtoMessage() {}

func (x *PluginConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginConfig.ProtoReflect.Descriptor instead.
func (*PluginConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_wasm_v3_wasm_proto_rawDescGZIP(), []int{1}
}

func (x *PluginConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginConfig) GetRootId() string {
	if x != nil {
		return x.RootId
	}
	return ""
}

func (m *PluginConfig) GetVmConfig() isPluginConfig_VmConfig {
	if m != nil {
		return m.VmConfig
	}
	return nil
}

func (x *PluginConfig) GetInlineVmConfig() *VmConfig {
	if x, ok := x.GetVmConfig().(*PluginConfig_InlineVmConfig); ok {
		return x.InlineVmConfig
	}
	return nil
}

func (x *PluginConfig) GetConfiguration() *any.Any {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *PluginConfig) GetFailOpen() bool {
	if x != nil {
		return x.FailOpen
	}
	return false
}

type isPluginConfig_VmConfig interface {
	isPluginConfig_VmConfig()
}

type PluginConfig_InlineVmConfig struct {
	InlineVmConfig *VmConfig `protobuf:"bytes,3,opt,name=inline_vm_config,json=inlineVmConfig,proto3,oneof"` // In the future add referential VM configurations.
}

func (*PluginConfig_InlineVmConfig) isPluginConfig_VmConfig() {}

// [[#not-implemented-hide:]
// WasmService is configured as a built-in *envoy.wasm_service* :ref:`WasmService
// <config_wasm_service>` This opaque configuration will be used to create a Wasm Service.
type WasmService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// General plugin configuration.
	Config *PluginConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// If true, create a single VM rather than creating one VM per worker. Such a singleton can
	// not be used with filters.
	Singleton bool `protobuf:"varint,2,opt,name=singleton,proto3" json:"singleton,omitempty"`
}

func (x *WasmService) Reset() {
	*x = WasmService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WasmService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WasmService) ProtoMessage() {}

func (x *WasmService) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WasmService.ProtoReflect.Descriptor instead.
func (*WasmService) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_wasm_v3_wasm_proto_rawDescGZIP(), []int{2}
}

func (x *WasmService) GetConfig() *PluginConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *WasmService) GetSingleton() bool {
	if x != nil {
		return x.Singleton
	}
	return false
}

var File_envoy_extensions_wasm_v3_wasm_proto protoreflect.FileDescriptor

var file_envoy_extensions_wasm_v3_wasm_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x2f, 0x76, 0x33, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x33, 0x1a,
	0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x02, 0x0a, 0x08, 0x56, 0x6d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x13, 0x0a, 0x05, 0x76, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x76, 0x6d, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x72, 0x65, 0x63,
	0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x12,
	0x34, 0x0a, 0x17, 0x6e, 0x61, 0x63, 0x6b, 0x5f, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x13, 0x6e, 0x61, 0x63, 0x6b, 0x4f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x4d, 0x69, 0x73, 0x73, 0x22, 0xf1, 0x01, 0x0a, 0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f,
	0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f,
	0x74, 0x49, 0x64, 0x12, 0x4e, 0x0a, 0x10, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x76, 0x6d,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x33, 0x2e, 0x56, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x48, 0x00, 0x52, 0x0e, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x56, 0x6d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x3a, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x4f, 0x70, 0x65, 0x6e, 0x42, 0x0b, 0x0a, 0x09,
	0x76, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x6b, 0x0a, 0x0b, 0x57, 0x61, 0x73,
	0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d,
	0x2e, 0x76, 0x33, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x42, 0x3d, 0x0a, 0x26, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x76, 0x33,
	0x42, 0x09, 0x57, 0x61, 0x73, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_wasm_v3_wasm_proto_rawDescOnce sync.Once
	file_envoy_extensions_wasm_v3_wasm_proto_rawDescData = file_envoy_extensions_wasm_v3_wasm_proto_rawDesc
)

func file_envoy_extensions_wasm_v3_wasm_proto_rawDescGZIP() []byte {
	file_envoy_extensions_wasm_v3_wasm_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_wasm_v3_wasm_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_wasm_v3_wasm_proto_rawDescData)
	})
	return file_envoy_extensions_wasm_v3_wasm_proto_rawDescData
}

var file_envoy_extensions_wasm_v3_wasm_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_extensions_wasm_v3_wasm_proto_goTypes = []interface{}{
	(*VmConfig)(nil),           // 0: envoy.extensions.wasm.v3.VmConfig
	(*PluginConfig)(nil),       // 1: envoy.extensions.wasm.v3.PluginConfig
	(*WasmService)(nil),        // 2: envoy.extensions.wasm.v3.WasmService
	(*v3.AsyncDataSource)(nil), // 3: envoy.config.core.v3.AsyncDataSource
	(*any.Any)(nil),            // 4: google.protobuf.Any
}
var file_envoy_extensions_wasm_v3_wasm_proto_depIdxs = []int32{
	3, // 0: envoy.extensions.wasm.v3.VmConfig.code:type_name -> envoy.config.core.v3.AsyncDataSource
	4, // 1: envoy.extensions.wasm.v3.VmConfig.configuration:type_name -> google.protobuf.Any
	0, // 2: envoy.extensions.wasm.v3.PluginConfig.inline_vm_config:type_name -> envoy.extensions.wasm.v3.VmConfig
	4, // 3: envoy.extensions.wasm.v3.PluginConfig.configuration:type_name -> google.protobuf.Any
	1, // 4: envoy.extensions.wasm.v3.WasmService.config:type_name -> envoy.extensions.wasm.v3.PluginConfig
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_wasm_v3_wasm_proto_init() }
func file_envoy_extensions_wasm_v3_wasm_proto_init() {
	if File_envoy_extensions_wasm_v3_wasm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VmConfig); i {
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
		file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginConfig); i {
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
		file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WasmService); i {
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
	file_envoy_extensions_wasm_v3_wasm_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PluginConfig_InlineVmConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_wasm_v3_wasm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_wasm_v3_wasm_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_wasm_v3_wasm_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_wasm_v3_wasm_proto_msgTypes,
	}.Build()
	File_envoy_extensions_wasm_v3_wasm_proto = out.File
	file_envoy_extensions_wasm_v3_wasm_proto_rawDesc = nil
	file_envoy_extensions_wasm_v3_wasm_proto_goTypes = nil
	file_envoy_extensions_wasm_v3_wasm_proto_depIdxs = nil
}
