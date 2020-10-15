// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language.proto

package pulumirpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetRequiredPluginsRequest struct {
	Project              string   `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Pwd                  string   `protobuf:"bytes,2,opt,name=pwd,proto3" json:"pwd,omitempty"`
	Program              string   `protobuf:"bytes,3,opt,name=program,proto3" json:"program,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequiredPluginsRequest) Reset()         { *m = GetRequiredPluginsRequest{} }
func (m *GetRequiredPluginsRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequiredPluginsRequest) ProtoMessage()    {}
func (*GetRequiredPluginsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e123c61d1ddd0892, []int{0}
}

func (m *GetRequiredPluginsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequiredPluginsRequest.Unmarshal(m, b)
}
func (m *GetRequiredPluginsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequiredPluginsRequest.Marshal(b, m, deterministic)
}
func (m *GetRequiredPluginsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequiredPluginsRequest.Merge(m, src)
}
func (m *GetRequiredPluginsRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequiredPluginsRequest.Size(m)
}
func (m *GetRequiredPluginsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequiredPluginsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequiredPluginsRequest proto.InternalMessageInfo

func (m *GetRequiredPluginsRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *GetRequiredPluginsRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *GetRequiredPluginsRequest) GetProgram() string {
	if m != nil {
		return m.Program
	}
	return ""
}

type GetRequiredPluginsResponse struct {
	Plugins              []*PluginDependency `protobuf:"bytes,1,rep,name=plugins,proto3" json:"plugins,omitempty"`
	Providers            []*PluginInfo       `protobuf:"bytes,2,rep,name=providers,proto3" json:"providers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetRequiredPluginsResponse) Reset()         { *m = GetRequiredPluginsResponse{} }
func (m *GetRequiredPluginsResponse) String() string { return proto.CompactTextString(m) }
func (*GetRequiredPluginsResponse) ProtoMessage()    {}
func (*GetRequiredPluginsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e123c61d1ddd0892, []int{1}
}

func (m *GetRequiredPluginsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequiredPluginsResponse.Unmarshal(m, b)
}
func (m *GetRequiredPluginsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequiredPluginsResponse.Marshal(b, m, deterministic)
}
func (m *GetRequiredPluginsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequiredPluginsResponse.Merge(m, src)
}
func (m *GetRequiredPluginsResponse) XXX_Size() int {
	return xxx_messageInfo_GetRequiredPluginsResponse.Size(m)
}
func (m *GetRequiredPluginsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequiredPluginsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequiredPluginsResponse proto.InternalMessageInfo

func (m *GetRequiredPluginsResponse) GetPlugins() []*PluginDependency {
	if m != nil {
		return m.Plugins
	}
	return nil
}

func (m *GetRequiredPluginsResponse) GetProviders() []*PluginInfo {
	if m != nil {
		return m.Providers
	}
	return nil
}

// RunRequest asks the interpreter to execute a program.
type RunRequest struct {
	Project              string            `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Stack                string            `protobuf:"bytes,2,opt,name=stack,proto3" json:"stack,omitempty"`
	Pwd                  string            `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	Program              string            `protobuf:"bytes,4,opt,name=program,proto3" json:"program,omitempty"`
	Args                 []string          `protobuf:"bytes,5,rep,name=args,proto3" json:"args,omitempty"`
	Config               map[string]string `protobuf:"bytes,6,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DryRun               bool              `protobuf:"varint,7,opt,name=dryRun,proto3" json:"dryRun,omitempty"`
	Parallel             int32             `protobuf:"varint,8,opt,name=parallel,proto3" json:"parallel,omitempty"`
	MonitorAddress       string            `protobuf:"bytes,9,opt,name=monitor_address,json=monitorAddress,proto3" json:"monitor_address,omitempty"`
	QueryMode            bool              `protobuf:"varint,10,opt,name=queryMode,proto3" json:"queryMode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RunRequest) Reset()         { *m = RunRequest{} }
func (m *RunRequest) String() string { return proto.CompactTextString(m) }
func (*RunRequest) ProtoMessage()    {}
func (*RunRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e123c61d1ddd0892, []int{2}
}

func (m *RunRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunRequest.Unmarshal(m, b)
}
func (m *RunRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunRequest.Marshal(b, m, deterministic)
}
func (m *RunRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunRequest.Merge(m, src)
}
func (m *RunRequest) XXX_Size() int {
	return xxx_messageInfo_RunRequest.Size(m)
}
func (m *RunRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunRequest proto.InternalMessageInfo

func (m *RunRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *RunRequest) GetStack() string {
	if m != nil {
		return m.Stack
	}
	return ""
}

func (m *RunRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *RunRequest) GetProgram() string {
	if m != nil {
		return m.Program
	}
	return ""
}

func (m *RunRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *RunRequest) GetConfig() map[string]string {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *RunRequest) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

func (m *RunRequest) GetParallel() int32 {
	if m != nil {
		return m.Parallel
	}
	return 0
}

func (m *RunRequest) GetMonitorAddress() string {
	if m != nil {
		return m.MonitorAddress
	}
	return ""
}

func (m *RunRequest) GetQueryMode() bool {
	if m != nil {
		return m.QueryMode
	}
	return false
}

// RunResponse is the response back from the interpreter/source back to the monitor.
type RunResponse struct {
	// An unhandled error if any occurred.
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// An error happened.  And it was reported to the user.  Work should stop immediately
	// with nothing further to print to the user.  This corresponds to a "result.Bail()"
	// value in the 'go' layer.
	Bail                 bool     `protobuf:"varint,2,opt,name=bail,proto3" json:"bail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunResponse) Reset()         { *m = RunResponse{} }
func (m *RunResponse) String() string { return proto.CompactTextString(m) }
func (*RunResponse) ProtoMessage()    {}
func (*RunResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e123c61d1ddd0892, []int{3}
}

func (m *RunResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunResponse.Unmarshal(m, b)
}
func (m *RunResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunResponse.Marshal(b, m, deterministic)
}
func (m *RunResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunResponse.Merge(m, src)
}
func (m *RunResponse) XXX_Size() int {
	return xxx_messageInfo_RunResponse.Size(m)
}
func (m *RunResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunResponse proto.InternalMessageInfo

func (m *RunResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *RunResponse) GetBail() bool {
	if m != nil {
		return m.Bail
	}
	return false
}

type StartProviderRequest struct {
	// The name of the provider to start.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The version of the provider to start.
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartProviderRequest) Reset()         { *m = StartProviderRequest{} }
func (m *StartProviderRequest) String() string { return proto.CompactTextString(m) }
func (*StartProviderRequest) ProtoMessage()    {}
func (*StartProviderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e123c61d1ddd0892, []int{4}
}

func (m *StartProviderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartProviderRequest.Unmarshal(m, b)
}
func (m *StartProviderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartProviderRequest.Marshal(b, m, deterministic)
}
func (m *StartProviderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartProviderRequest.Merge(m, src)
}
func (m *StartProviderRequest) XXX_Size() int {
	return xxx_messageInfo_StartProviderRequest.Size(m)
}
func (m *StartProviderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartProviderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartProviderRequest proto.InternalMessageInfo

func (m *StartProviderRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StartProviderRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type StartProviderResponse struct {
	// The address at which the provider service is listening.
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartProviderResponse) Reset()         { *m = StartProviderResponse{} }
func (m *StartProviderResponse) String() string { return proto.CompactTextString(m) }
func (*StartProviderResponse) ProtoMessage()    {}
func (*StartProviderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e123c61d1ddd0892, []int{5}
}

func (m *StartProviderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartProviderResponse.Unmarshal(m, b)
}
func (m *StartProviderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartProviderResponse.Marshal(b, m, deterministic)
}
func (m *StartProviderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartProviderResponse.Merge(m, src)
}
func (m *StartProviderResponse) XXX_Size() int {
	return xxx_messageInfo_StartProviderResponse.Size(m)
}
func (m *StartProviderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartProviderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartProviderResponse proto.InternalMessageInfo

func (m *StartProviderResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequiredPluginsRequest)(nil), "pulumirpc.GetRequiredPluginsRequest")
	proto.RegisterType((*GetRequiredPluginsResponse)(nil), "pulumirpc.GetRequiredPluginsResponse")
	proto.RegisterType((*RunRequest)(nil), "pulumirpc.RunRequest")
	proto.RegisterMapType((map[string]string)(nil), "pulumirpc.RunRequest.ConfigEntry")
	proto.RegisterType((*RunResponse)(nil), "pulumirpc.RunResponse")
	proto.RegisterType((*StartProviderRequest)(nil), "pulumirpc.StartProviderRequest")
	proto.RegisterType((*StartProviderResponse)(nil), "pulumirpc.StartProviderResponse")
}

func init() { proto.RegisterFile("language.proto", fileDescriptor_e123c61d1ddd0892) }

var fileDescriptor_e123c61d1ddd0892 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x6d, 0x4f, 0xd4, 0x40,
	0x10, 0xa6, 0x57, 0xee, 0xa5, 0x83, 0x80, 0xd9, 0x00, 0x59, 0x8b, 0x89, 0xb5, 0xd1, 0x78, 0x9f,
	0x4a, 0x84, 0xf8, 0x82, 0x9f, 0x34, 0x42, 0x88, 0x89, 0x26, 0x64, 0xf5, 0xbb, 0x59, 0xda, 0xa1,
	0xa9, 0xb4, 0xbb, 0x65, 0xbb, 0x3d, 0xd3, 0x7f, 0xe0, 0x2f, 0xf3, 0xdf, 0xf8, 0x1f, 0x4c, 0xb7,
	0xdb, 0xbb, 0x03, 0x8e, 0xf0, 0x6d, 0x9e, 0xe9, 0x33, 0xdb, 0xe7, 0x99, 0xd9, 0x59, 0xd8, 0xca,
	0xb9, 0x48, 0x6b, 0x9e, 0x62, 0x54, 0x2a, 0xa9, 0x25, 0xf1, 0xca, 0x3a, 0xaf, 0x8b, 0x4c, 0x95,
	0xb1, 0xff, 0xa8, 0xcc, 0xeb, 0x34, 0x13, 0xdd, 0x07, 0x7f, 0x3f, 0x95, 0x32, 0xcd, 0xf1, 0xc0,
	0xa0, 0x8b, 0xfa, 0xf2, 0x00, 0x8b, 0x52, 0x37, 0xdd, 0xc7, 0x90, 0xc3, 0x93, 0x33, 0xd4, 0x0c,
	0xaf, 0xeb, 0x4c, 0x61, 0x72, 0x6e, 0xea, 0xaa, 0x16, 0x62, 0xa5, 0x09, 0x85, 0x71, 0xa9, 0xe4,
	0x2f, 0x8c, 0x35, 0x75, 0x02, 0x67, 0xea, 0xb1, 0x1e, 0x92, 0xc7, 0xe0, 0x96, 0xbf, 0x13, 0x3a,
	0x30, 0xd9, 0x36, 0xb4, 0xdc, 0x54, 0xf1, 0x82, 0xba, 0x73, 0x6e, 0x0b, 0xc3, 0x3f, 0x0e, 0xf8,
	0xab, 0xfe, 0x51, 0x95, 0x52, 0x54, 0x48, 0xde, 0xc0, 0xb8, 0x93, 0x5b, 0x51, 0x27, 0x70, 0xa7,
	0x1b, 0x87, 0xfb, 0xd1, 0xdc, 0x49, 0xd4, 0x91, 0x4f, 0xb0, 0x44, 0x91, 0xa0, 0x88, 0x1b, 0xd6,
	0x73, 0xc9, 0x11, 0x78, 0xa5, 0x92, 0xb3, 0x2c, 0x41, 0x55, 0xd1, 0x81, 0x29, 0xdc, 0xbd, 0x53,
	0xf8, 0x45, 0x5c, 0x4a, 0xb6, 0xe0, 0x85, 0xff, 0x06, 0x00, 0xac, 0x16, 0x0f, 0xfb, 0xdb, 0x81,
	0x61, 0xa5, 0x79, 0x7c, 0x65, 0x1d, 0x76, 0xa0, 0x77, 0xed, 0xae, 0x74, 0xbd, 0x7e, 0xc3, 0x35,
	0x21, 0xb0, 0xce, 0x55, 0x5a, 0xd1, 0x61, 0xe0, 0x4e, 0x3d, 0x66, 0x62, 0x72, 0x0c, 0xa3, 0x58,
	0x8a, 0xcb, 0x2c, 0xa5, 0x23, 0x23, 0xf8, 0xf9, 0x92, 0xe0, 0x85, 0xac, 0xe8, 0xb3, 0xe1, 0x9c,
	0x0a, 0xad, 0x1a, 0x66, 0x0b, 0xc8, 0x1e, 0x8c, 0x12, 0xd5, 0xb0, 0x5a, 0xd0, 0x71, 0xe0, 0x4c,
	0x27, 0xcc, 0x22, 0xe2, 0xc3, 0xa4, 0xe4, 0x8a, 0xe7, 0x39, 0xe6, 0x74, 0x12, 0x38, 0xd3, 0x21,
	0x9b, 0x63, 0xf2, 0x0a, 0xb6, 0x0b, 0x29, 0x32, 0x2d, 0xd5, 0x4f, 0x9e, 0x24, 0x0a, 0xab, 0x8a,
	0x7a, 0x46, 0xe4, 0x96, 0x4d, 0x7f, 0xea, 0xb2, 0xe4, 0x29, 0x78, 0xd7, 0x35, 0xaa, 0xe6, 0x9b,
	0x4c, 0x90, 0x82, 0x39, 0x7f, 0x91, 0xf0, 0x8f, 0x61, 0x63, 0x49, 0x51, 0xdb, 0x84, 0x2b, 0x6c,
	0x6c, 0xc3, 0xda, 0xb0, 0x6d, 0xd6, 0x8c, 0xe7, 0x35, 0xf6, 0xcd, 0x32, 0xe0, 0xc3, 0xe0, 0xbd,
	0x13, 0xbe, 0x83, 0x0d, 0xe3, 0xcb, 0x8e, 0x7a, 0x07, 0x86, 0xa8, 0x94, 0x54, 0xb6, 0xb8, 0x03,
	0x6d, 0xa7, 0x2e, 0x78, 0x96, 0x9b, 0xea, 0x09, 0x33, 0x71, 0x78, 0x02, 0x3b, 0xdf, 0x35, 0x57,
	0xfa, 0xdc, 0x8e, 0xae, 0x9f, 0x18, 0x81, 0x75, 0xc1, 0x0b, 0xb4, 0x07, 0x98, 0xb8, 0x9d, 0xc1,
	0x0c, 0x55, 0x95, 0x49, 0x61, 0x05, 0xf4, 0x30, 0x7c, 0x0d, 0xbb, 0xb7, 0x4e, 0xb1, 0x42, 0x28,
	0x8c, 0xfb, 0x8e, 0xd8, 0xc1, 0x5b, 0x78, 0xf8, 0x77, 0x00, 0xdb, 0x5f, 0xed, 0x62, 0xb1, 0x5a,
	0xe8, 0xac, 0x40, 0x12, 0x03, 0xb9, 0x7b, 0x7f, 0xc9, 0x8b, 0xa5, 0xe1, 0xdd, 0xbb, 0x42, 0xfe,
	0xcb, 0x07, 0x58, 0x9d, 0xa0, 0x70, 0x8d, 0xbc, 0x05, 0xb7, 0x9d, 0xe7, 0xee, 0xca, 0x2b, 0xe1,
	0xef, 0xdd, 0x4e, 0xcf, 0xeb, 0x7e, 0xc0, 0xe6, 0x0d, 0x8f, 0xe4, 0xd9, 0x12, 0x75, 0x55, 0x0f,
	0xfd, 0xe0, 0x7e, 0xc2, 0xfc, 0xd4, 0x8f, 0xb0, 0x79, 0x86, 0x7a, 0xb1, 0x44, 0x64, 0x2f, 0xea,
	0x5e, 0x91, 0xa8, 0x7f, 0x45, 0xa2, 0xd3, 0xf6, 0x15, 0xf1, 0x57, 0xef, 0x5c, 0xb8, 0x76, 0x31,
	0x32, 0xc4, 0xa3, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x92, 0xec, 0xe0, 0x2c, 0xa7, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LanguageRuntimeClient is the client API for LanguageRuntime service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LanguageRuntimeClient interface {
	// GetRequiredPlugins computes the complete set of anticipated plugins required by a program.
	GetRequiredPlugins(ctx context.Context, in *GetRequiredPluginsRequest, opts ...grpc.CallOption) (*GetRequiredPluginsResponse, error)
	// Run executes a program and returns its result.
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error)
	// StartProvider starts a resource provider implemented by the program.
	StartProvider(ctx context.Context, in *StartProviderRequest, opts ...grpc.CallOption) (*StartProviderResponse, error)
	// GetPluginInfo returns generic information about this plugin, like its version.
	GetPluginInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PluginInfo, error)
}

type languageRuntimeClient struct {
	cc grpc.ClientConnInterface
}

func NewLanguageRuntimeClient(cc grpc.ClientConnInterface) LanguageRuntimeClient {
	return &languageRuntimeClient{cc}
}

func (c *languageRuntimeClient) GetRequiredPlugins(ctx context.Context, in *GetRequiredPluginsRequest, opts ...grpc.CallOption) (*GetRequiredPluginsResponse, error) {
	out := new(GetRequiredPluginsResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.LanguageRuntime/GetRequiredPlugins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageRuntimeClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.LanguageRuntime/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageRuntimeClient) StartProvider(ctx context.Context, in *StartProviderRequest, opts ...grpc.CallOption) (*StartProviderResponse, error) {
	out := new(StartProviderResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.LanguageRuntime/StartProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *languageRuntimeClient) GetPluginInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := c.cc.Invoke(ctx, "/pulumirpc.LanguageRuntime/GetPluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LanguageRuntimeServer is the server API for LanguageRuntime service.
type LanguageRuntimeServer interface {
	// GetRequiredPlugins computes the complete set of anticipated plugins required by a program.
	GetRequiredPlugins(context.Context, *GetRequiredPluginsRequest) (*GetRequiredPluginsResponse, error)
	// Run executes a program and returns its result.
	Run(context.Context, *RunRequest) (*RunResponse, error)
	// StartProvider starts a resource provider implemented by the program.
	StartProvider(context.Context, *StartProviderRequest) (*StartProviderResponse, error)
	// GetPluginInfo returns generic information about this plugin, like its version.
	GetPluginInfo(context.Context, *empty.Empty) (*PluginInfo, error)
}

// UnimplementedLanguageRuntimeServer can be embedded to have forward compatible implementations.
type UnimplementedLanguageRuntimeServer struct {
}

func (*UnimplementedLanguageRuntimeServer) GetRequiredPlugins(ctx context.Context, req *GetRequiredPluginsRequest) (*GetRequiredPluginsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequiredPlugins not implemented")
}
func (*UnimplementedLanguageRuntimeServer) Run(ctx context.Context, req *RunRequest) (*RunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (*UnimplementedLanguageRuntimeServer) StartProvider(ctx context.Context, req *StartProviderRequest) (*StartProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartProvider not implemented")
}
func (*UnimplementedLanguageRuntimeServer) GetPluginInfo(ctx context.Context, req *empty.Empty) (*PluginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginInfo not implemented")
}

func RegisterLanguageRuntimeServer(s *grpc.Server, srv LanguageRuntimeServer) {
	s.RegisterService(&_LanguageRuntime_serviceDesc, srv)
}

func _LanguageRuntime_GetRequiredPlugins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequiredPluginsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).GetRequiredPlugins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.LanguageRuntime/GetRequiredPlugins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).GetRequiredPlugins(ctx, req.(*GetRequiredPluginsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanguageRuntime_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.LanguageRuntime/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanguageRuntime_StartProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).StartProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.LanguageRuntime/StartProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).StartProvider(ctx, req.(*StartProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LanguageRuntime_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageRuntimeServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.LanguageRuntime/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageRuntimeServer).GetPluginInfo(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _LanguageRuntime_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pulumirpc.LanguageRuntime",
	HandlerType: (*LanguageRuntimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRequiredPlugins",
			Handler:    _LanguageRuntime_GetRequiredPlugins_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _LanguageRuntime_Run_Handler,
		},
		{
			MethodName: "StartProvider",
			Handler:    _LanguageRuntime_StartProvider_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _LanguageRuntime_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language.proto",
}
