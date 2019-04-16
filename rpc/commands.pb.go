// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commands.proto

package rpc // import "github.com/arduino/arduino-cli/rpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Configuration contains information to instantiate an Arduino Platform Service
type Configuration struct {
	// dataDir represents the current root of the arduino tree (defaulted to
	// `$HOME/.arduino15` on linux).
	DataDir string `protobuf:"bytes,1,opt,name=dataDir,proto3" json:"dataDir,omitempty"`
	// sketchbookDir represents the current root of the sketchbooks tree
	// (defaulted to `$HOME/Arduino`).
	SketchbookDir string `protobuf:"bytes,2,opt,name=sketchbookDir,proto3" json:"sketchbookDir,omitempty"`
	// ArduinoIDEDirectory is the directory of the Arduino IDE if the CLI runs
	// together with it.
	DownloadsDir string `protobuf:"bytes,3,opt,name=downloadsDir,proto3" json:"downloadsDir,omitempty"`
	// BoardManagerAdditionalUrls contains the additional URL for 3rd party
	// packages
	BoardManagerAdditionalUrls []string `protobuf:"bytes,4,rep,name=boardManagerAdditionalUrls,proto3" json:"boardManagerAdditionalUrls,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_595e725b92afdfd7, []int{0}
}
func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (dst *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(dst, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetDataDir() string {
	if m != nil {
		return m.DataDir
	}
	return ""
}

func (m *Configuration) GetSketchbookDir() string {
	if m != nil {
		return m.SketchbookDir
	}
	return ""
}

func (m *Configuration) GetDownloadsDir() string {
	if m != nil {
		return m.DownloadsDir
	}
	return ""
}

func (m *Configuration) GetBoardManagerAdditionalUrls() []string {
	if m != nil {
		return m.BoardManagerAdditionalUrls
	}
	return nil
}

type InitReq struct {
	Configuration        *Configuration `protobuf:"bytes,1,opt,name=configuration,proto3" json:"configuration,omitempty"`
	LibraryManagerOnly   bool           `protobuf:"varint,2,opt,name=library_manager_only,json=libraryManagerOnly,proto3" json:"library_manager_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *InitReq) Reset()         { *m = InitReq{} }
func (m *InitReq) String() string { return proto.CompactTextString(m) }
func (*InitReq) ProtoMessage()    {}
func (*InitReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_595e725b92afdfd7, []int{1}
}
func (m *InitReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitReq.Unmarshal(m, b)
}
func (m *InitReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitReq.Marshal(b, m, deterministic)
}
func (dst *InitReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitReq.Merge(dst, src)
}
func (m *InitReq) XXX_Size() int {
	return xxx_messageInfo_InitReq.Size(m)
}
func (m *InitReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InitReq.DiscardUnknown(m)
}

var xxx_messageInfo_InitReq proto.InternalMessageInfo

func (m *InitReq) GetConfiguration() *Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *InitReq) GetLibraryManagerOnly() bool {
	if m != nil {
		return m.LibraryManagerOnly
	}
	return false
}

type InitResp struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *InitResp) Reset()         { *m = InitResp{} }
func (m *InitResp) String() string { return proto.CompactTextString(m) }
func (*InitResp) ProtoMessage()    {}
func (*InitResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_595e725b92afdfd7, []int{2}
}
func (m *InitResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitResp.Unmarshal(m, b)
}
func (m *InitResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitResp.Marshal(b, m, deterministic)
}
func (dst *InitResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitResp.Merge(dst, src)
}
func (m *InitResp) XXX_Size() int {
	return xxx_messageInfo_InitResp.Size(m)
}
func (m *InitResp) XXX_DiscardUnknown() {
	xxx_messageInfo_InitResp.DiscardUnknown(m)
}

var xxx_messageInfo_InitResp proto.InternalMessageInfo

func (m *InitResp) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

type DestroyReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DestroyReq) Reset()         { *m = DestroyReq{} }
func (m *DestroyReq) String() string { return proto.CompactTextString(m) }
func (*DestroyReq) ProtoMessage()    {}
func (*DestroyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_595e725b92afdfd7, []int{3}
}
func (m *DestroyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyReq.Unmarshal(m, b)
}
func (m *DestroyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyReq.Marshal(b, m, deterministic)
}
func (dst *DestroyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyReq.Merge(dst, src)
}
func (m *DestroyReq) XXX_Size() int {
	return xxx_messageInfo_DestroyReq.Size(m)
}
func (m *DestroyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyReq.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyReq proto.InternalMessageInfo

func (m *DestroyReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

type DestroyResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DestroyResp) Reset()         { *m = DestroyResp{} }
func (m *DestroyResp) String() string { return proto.CompactTextString(m) }
func (*DestroyResp) ProtoMessage()    {}
func (*DestroyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_595e725b92afdfd7, []int{4}
}
func (m *DestroyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DestroyResp.Unmarshal(m, b)
}
func (m *DestroyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DestroyResp.Marshal(b, m, deterministic)
}
func (dst *DestroyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DestroyResp.Merge(dst, src)
}
func (m *DestroyResp) XXX_Size() int {
	return xxx_messageInfo_DestroyResp.Size(m)
}
func (m *DestroyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DestroyResp.DiscardUnknown(m)
}

var xxx_messageInfo_DestroyResp proto.InternalMessageInfo

type RescanReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RescanReq) Reset()         { *m = RescanReq{} }
func (m *RescanReq) String() string { return proto.CompactTextString(m) }
func (*RescanReq) ProtoMessage()    {}
func (*RescanReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_595e725b92afdfd7, []int{5}
}
func (m *RescanReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RescanReq.Unmarshal(m, b)
}
func (m *RescanReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RescanReq.Marshal(b, m, deterministic)
}
func (dst *RescanReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RescanReq.Merge(dst, src)
}
func (m *RescanReq) XXX_Size() int {
	return xxx_messageInfo_RescanReq.Size(m)
}
func (m *RescanReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RescanReq.DiscardUnknown(m)
}

var xxx_messageInfo_RescanReq proto.InternalMessageInfo

func (m *RescanReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

type RescanResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RescanResp) Reset()         { *m = RescanResp{} }
func (m *RescanResp) String() string { return proto.CompactTextString(m) }
func (*RescanResp) ProtoMessage()    {}
func (*RescanResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_595e725b92afdfd7, []int{6}
}
func (m *RescanResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RescanResp.Unmarshal(m, b)
}
func (m *RescanResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RescanResp.Marshal(b, m, deterministic)
}
func (dst *RescanResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RescanResp.Merge(dst, src)
}
func (m *RescanResp) XXX_Size() int {
	return xxx_messageInfo_RescanResp.Size(m)
}
func (m *RescanResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RescanResp.DiscardUnknown(m)
}

var xxx_messageInfo_RescanResp proto.InternalMessageInfo

type UpdateIndexReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateIndexReq) Reset()         { *m = UpdateIndexReq{} }
func (m *UpdateIndexReq) String() string { return proto.CompactTextString(m) }
func (*UpdateIndexReq) ProtoMessage()    {}
func (*UpdateIndexReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_595e725b92afdfd7, []int{7}
}
func (m *UpdateIndexReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateIndexReq.Unmarshal(m, b)
}
func (m *UpdateIndexReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateIndexReq.Marshal(b, m, deterministic)
}
func (dst *UpdateIndexReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateIndexReq.Merge(dst, src)
}
func (m *UpdateIndexReq) XXX_Size() int {
	return xxx_messageInfo_UpdateIndexReq.Size(m)
}
func (m *UpdateIndexReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateIndexReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateIndexReq proto.InternalMessageInfo

func (m *UpdateIndexReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

type UpdateIndexResp struct {
	DownloadProgress     *DownloadProgress `protobuf:"bytes,1,opt,name=download_progress,json=downloadProgress,proto3" json:"download_progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdateIndexResp) Reset()         { *m = UpdateIndexResp{} }
func (m *UpdateIndexResp) String() string { return proto.CompactTextString(m) }
func (*UpdateIndexResp) ProtoMessage()    {}
func (*UpdateIndexResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_595e725b92afdfd7, []int{8}
}
func (m *UpdateIndexResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateIndexResp.Unmarshal(m, b)
}
func (m *UpdateIndexResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateIndexResp.Marshal(b, m, deterministic)
}
func (dst *UpdateIndexResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateIndexResp.Merge(dst, src)
}
func (m *UpdateIndexResp) XXX_Size() int {
	return xxx_messageInfo_UpdateIndexResp.Size(m)
}
func (m *UpdateIndexResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateIndexResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateIndexResp proto.InternalMessageInfo

func (m *UpdateIndexResp) GetDownloadProgress() *DownloadProgress {
	if m != nil {
		return m.DownloadProgress
	}
	return nil
}

func init() {
	proto.RegisterType((*Configuration)(nil), "arduino.Configuration")
	proto.RegisterType((*InitReq)(nil), "arduino.InitReq")
	proto.RegisterType((*InitResp)(nil), "arduino.InitResp")
	proto.RegisterType((*DestroyReq)(nil), "arduino.DestroyReq")
	proto.RegisterType((*DestroyResp)(nil), "arduino.DestroyResp")
	proto.RegisterType((*RescanReq)(nil), "arduino.RescanReq")
	proto.RegisterType((*RescanResp)(nil), "arduino.RescanResp")
	proto.RegisterType((*UpdateIndexReq)(nil), "arduino.UpdateIndexReq")
	proto.RegisterType((*UpdateIndexResp)(nil), "arduino.UpdateIndexResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArduinoCoreClient is the client API for ArduinoCore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArduinoCoreClient interface {
	// Start a new instance of the Arduino Core Service
	Init(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*InitResp, error)
	// Destroy an instance of the Arduino Core Service
	Destroy(ctx context.Context, in *DestroyReq, opts ...grpc.CallOption) (*DestroyResp, error)
	// Rescan instance of the Arduino Core Service
	Rescan(ctx context.Context, in *RescanReq, opts ...grpc.CallOption) (*RescanResp, error)
	// Update package index of the Arduino Core Service
	UpdateIndex(ctx context.Context, in *UpdateIndexReq, opts ...grpc.CallOption) (ArduinoCore_UpdateIndexClient, error)
	// Requests details about a board
	BoardDetails(ctx context.Context, in *BoardDetailsReq, opts ...grpc.CallOption) (*BoardDetailsResp, error)
	Compile(ctx context.Context, in *CompileReq, opts ...grpc.CallOption) (ArduinoCore_CompileClient, error)
	PlatformInstall(ctx context.Context, in *PlatformInstallReq, opts ...grpc.CallOption) (ArduinoCore_PlatformInstallClient, error)
	PlatformDownload(ctx context.Context, in *PlatformDownloadReq, opts ...grpc.CallOption) (ArduinoCore_PlatformDownloadClient, error)
	PlatformUninstall(ctx context.Context, in *PlatformUninstallReq, opts ...grpc.CallOption) (ArduinoCore_PlatformUninstallClient, error)
	PlatformUpgrade(ctx context.Context, in *PlatformUpgradeReq, opts ...grpc.CallOption) (ArduinoCore_PlatformUpgradeClient, error)
	Upload(ctx context.Context, in *UploadReq, opts ...grpc.CallOption) (ArduinoCore_UploadClient, error)
	PlatformSearch(ctx context.Context, in *PlatformSearchReq, opts ...grpc.CallOption) (*PlatformSearchResp, error)
	PlatformList(ctx context.Context, in *PlatformListReq, opts ...grpc.CallOption) (*PlatformListResp, error)
}

type arduinoCoreClient struct {
	cc *grpc.ClientConn
}

func NewArduinoCoreClient(cc *grpc.ClientConn) ArduinoCoreClient {
	return &arduinoCoreClient{cc}
}

func (c *arduinoCoreClient) Init(ctx context.Context, in *InitReq, opts ...grpc.CallOption) (*InitResp, error) {
	out := new(InitResp)
	err := c.cc.Invoke(ctx, "/arduino.ArduinoCore/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arduinoCoreClient) Destroy(ctx context.Context, in *DestroyReq, opts ...grpc.CallOption) (*DestroyResp, error) {
	out := new(DestroyResp)
	err := c.cc.Invoke(ctx, "/arduino.ArduinoCore/Destroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arduinoCoreClient) Rescan(ctx context.Context, in *RescanReq, opts ...grpc.CallOption) (*RescanResp, error) {
	out := new(RescanResp)
	err := c.cc.Invoke(ctx, "/arduino.ArduinoCore/Rescan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arduinoCoreClient) UpdateIndex(ctx context.Context, in *UpdateIndexReq, opts ...grpc.CallOption) (ArduinoCore_UpdateIndexClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ArduinoCore_serviceDesc.Streams[0], "/arduino.ArduinoCore/UpdateIndex", opts...)
	if err != nil {
		return nil, err
	}
	x := &arduinoCoreUpdateIndexClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArduinoCore_UpdateIndexClient interface {
	Recv() (*UpdateIndexResp, error)
	grpc.ClientStream
}

type arduinoCoreUpdateIndexClient struct {
	grpc.ClientStream
}

func (x *arduinoCoreUpdateIndexClient) Recv() (*UpdateIndexResp, error) {
	m := new(UpdateIndexResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *arduinoCoreClient) BoardDetails(ctx context.Context, in *BoardDetailsReq, opts ...grpc.CallOption) (*BoardDetailsResp, error) {
	out := new(BoardDetailsResp)
	err := c.cc.Invoke(ctx, "/arduino.ArduinoCore/BoardDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arduinoCoreClient) Compile(ctx context.Context, in *CompileReq, opts ...grpc.CallOption) (ArduinoCore_CompileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ArduinoCore_serviceDesc.Streams[1], "/arduino.ArduinoCore/Compile", opts...)
	if err != nil {
		return nil, err
	}
	x := &arduinoCoreCompileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArduinoCore_CompileClient interface {
	Recv() (*CompileResp, error)
	grpc.ClientStream
}

type arduinoCoreCompileClient struct {
	grpc.ClientStream
}

func (x *arduinoCoreCompileClient) Recv() (*CompileResp, error) {
	m := new(CompileResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *arduinoCoreClient) PlatformInstall(ctx context.Context, in *PlatformInstallReq, opts ...grpc.CallOption) (ArduinoCore_PlatformInstallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ArduinoCore_serviceDesc.Streams[2], "/arduino.ArduinoCore/PlatformInstall", opts...)
	if err != nil {
		return nil, err
	}
	x := &arduinoCorePlatformInstallClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArduinoCore_PlatformInstallClient interface {
	Recv() (*PlatformInstallResp, error)
	grpc.ClientStream
}

type arduinoCorePlatformInstallClient struct {
	grpc.ClientStream
}

func (x *arduinoCorePlatformInstallClient) Recv() (*PlatformInstallResp, error) {
	m := new(PlatformInstallResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *arduinoCoreClient) PlatformDownload(ctx context.Context, in *PlatformDownloadReq, opts ...grpc.CallOption) (ArduinoCore_PlatformDownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ArduinoCore_serviceDesc.Streams[3], "/arduino.ArduinoCore/PlatformDownload", opts...)
	if err != nil {
		return nil, err
	}
	x := &arduinoCorePlatformDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArduinoCore_PlatformDownloadClient interface {
	Recv() (*PlatformDownloadResp, error)
	grpc.ClientStream
}

type arduinoCorePlatformDownloadClient struct {
	grpc.ClientStream
}

func (x *arduinoCorePlatformDownloadClient) Recv() (*PlatformDownloadResp, error) {
	m := new(PlatformDownloadResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *arduinoCoreClient) PlatformUninstall(ctx context.Context, in *PlatformUninstallReq, opts ...grpc.CallOption) (ArduinoCore_PlatformUninstallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ArduinoCore_serviceDesc.Streams[4], "/arduino.ArduinoCore/PlatformUninstall", opts...)
	if err != nil {
		return nil, err
	}
	x := &arduinoCorePlatformUninstallClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArduinoCore_PlatformUninstallClient interface {
	Recv() (*PlatformUninstallResp, error)
	grpc.ClientStream
}

type arduinoCorePlatformUninstallClient struct {
	grpc.ClientStream
}

func (x *arduinoCorePlatformUninstallClient) Recv() (*PlatformUninstallResp, error) {
	m := new(PlatformUninstallResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *arduinoCoreClient) PlatformUpgrade(ctx context.Context, in *PlatformUpgradeReq, opts ...grpc.CallOption) (ArduinoCore_PlatformUpgradeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ArduinoCore_serviceDesc.Streams[5], "/arduino.ArduinoCore/PlatformUpgrade", opts...)
	if err != nil {
		return nil, err
	}
	x := &arduinoCorePlatformUpgradeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArduinoCore_PlatformUpgradeClient interface {
	Recv() (*PlatformUpgradeResp, error)
	grpc.ClientStream
}

type arduinoCorePlatformUpgradeClient struct {
	grpc.ClientStream
}

func (x *arduinoCorePlatformUpgradeClient) Recv() (*PlatformUpgradeResp, error) {
	m := new(PlatformUpgradeResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *arduinoCoreClient) Upload(ctx context.Context, in *UploadReq, opts ...grpc.CallOption) (ArduinoCore_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ArduinoCore_serviceDesc.Streams[6], "/arduino.ArduinoCore/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &arduinoCoreUploadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArduinoCore_UploadClient interface {
	Recv() (*UploadResp, error)
	grpc.ClientStream
}

type arduinoCoreUploadClient struct {
	grpc.ClientStream
}

func (x *arduinoCoreUploadClient) Recv() (*UploadResp, error) {
	m := new(UploadResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *arduinoCoreClient) PlatformSearch(ctx context.Context, in *PlatformSearchReq, opts ...grpc.CallOption) (*PlatformSearchResp, error) {
	out := new(PlatformSearchResp)
	err := c.cc.Invoke(ctx, "/arduino.ArduinoCore/PlatformSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arduinoCoreClient) PlatformList(ctx context.Context, in *PlatformListReq, opts ...grpc.CallOption) (*PlatformListResp, error) {
	out := new(PlatformListResp)
	err := c.cc.Invoke(ctx, "/arduino.ArduinoCore/PlatformList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArduinoCoreServer is the server API for ArduinoCore service.
type ArduinoCoreServer interface {
	// Start a new instance of the Arduino Core Service
	Init(context.Context, *InitReq) (*InitResp, error)
	// Destroy an instance of the Arduino Core Service
	Destroy(context.Context, *DestroyReq) (*DestroyResp, error)
	// Rescan instance of the Arduino Core Service
	Rescan(context.Context, *RescanReq) (*RescanResp, error)
	// Update package index of the Arduino Core Service
	UpdateIndex(*UpdateIndexReq, ArduinoCore_UpdateIndexServer) error
	// Requests details about a board
	BoardDetails(context.Context, *BoardDetailsReq) (*BoardDetailsResp, error)
	Compile(*CompileReq, ArduinoCore_CompileServer) error
	PlatformInstall(*PlatformInstallReq, ArduinoCore_PlatformInstallServer) error
	PlatformDownload(*PlatformDownloadReq, ArduinoCore_PlatformDownloadServer) error
	PlatformUninstall(*PlatformUninstallReq, ArduinoCore_PlatformUninstallServer) error
	PlatformUpgrade(*PlatformUpgradeReq, ArduinoCore_PlatformUpgradeServer) error
	Upload(*UploadReq, ArduinoCore_UploadServer) error
	PlatformSearch(context.Context, *PlatformSearchReq) (*PlatformSearchResp, error)
	PlatformList(context.Context, *PlatformListReq) (*PlatformListResp, error)
}

func RegisterArduinoCoreServer(s *grpc.Server, srv ArduinoCoreServer) {
	s.RegisterService(&_ArduinoCore_serviceDesc, srv)
}

func _ArduinoCore_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArduinoCoreServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arduino.ArduinoCore/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArduinoCoreServer).Init(ctx, req.(*InitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArduinoCore_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArduinoCoreServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arduino.ArduinoCore/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArduinoCoreServer).Destroy(ctx, req.(*DestroyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArduinoCore_Rescan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RescanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArduinoCoreServer).Rescan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arduino.ArduinoCore/Rescan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArduinoCoreServer).Rescan(ctx, req.(*RescanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArduinoCore_UpdateIndex_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UpdateIndexReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArduinoCoreServer).UpdateIndex(m, &arduinoCoreUpdateIndexServer{stream})
}

type ArduinoCore_UpdateIndexServer interface {
	Send(*UpdateIndexResp) error
	grpc.ServerStream
}

type arduinoCoreUpdateIndexServer struct {
	grpc.ServerStream
}

func (x *arduinoCoreUpdateIndexServer) Send(m *UpdateIndexResp) error {
	return x.ServerStream.SendMsg(m)
}

func _ArduinoCore_BoardDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoardDetailsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArduinoCoreServer).BoardDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arduino.ArduinoCore/BoardDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArduinoCoreServer).BoardDetails(ctx, req.(*BoardDetailsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArduinoCore_Compile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CompileReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArduinoCoreServer).Compile(m, &arduinoCoreCompileServer{stream})
}

type ArduinoCore_CompileServer interface {
	Send(*CompileResp) error
	grpc.ServerStream
}

type arduinoCoreCompileServer struct {
	grpc.ServerStream
}

func (x *arduinoCoreCompileServer) Send(m *CompileResp) error {
	return x.ServerStream.SendMsg(m)
}

func _ArduinoCore_PlatformInstall_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PlatformInstallReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArduinoCoreServer).PlatformInstall(m, &arduinoCorePlatformInstallServer{stream})
}

type ArduinoCore_PlatformInstallServer interface {
	Send(*PlatformInstallResp) error
	grpc.ServerStream
}

type arduinoCorePlatformInstallServer struct {
	grpc.ServerStream
}

func (x *arduinoCorePlatformInstallServer) Send(m *PlatformInstallResp) error {
	return x.ServerStream.SendMsg(m)
}

func _ArduinoCore_PlatformDownload_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PlatformDownloadReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArduinoCoreServer).PlatformDownload(m, &arduinoCorePlatformDownloadServer{stream})
}

type ArduinoCore_PlatformDownloadServer interface {
	Send(*PlatformDownloadResp) error
	grpc.ServerStream
}

type arduinoCorePlatformDownloadServer struct {
	grpc.ServerStream
}

func (x *arduinoCorePlatformDownloadServer) Send(m *PlatformDownloadResp) error {
	return x.ServerStream.SendMsg(m)
}

func _ArduinoCore_PlatformUninstall_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PlatformUninstallReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArduinoCoreServer).PlatformUninstall(m, &arduinoCorePlatformUninstallServer{stream})
}

type ArduinoCore_PlatformUninstallServer interface {
	Send(*PlatformUninstallResp) error
	grpc.ServerStream
}

type arduinoCorePlatformUninstallServer struct {
	grpc.ServerStream
}

func (x *arduinoCorePlatformUninstallServer) Send(m *PlatformUninstallResp) error {
	return x.ServerStream.SendMsg(m)
}

func _ArduinoCore_PlatformUpgrade_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PlatformUpgradeReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArduinoCoreServer).PlatformUpgrade(m, &arduinoCorePlatformUpgradeServer{stream})
}

type ArduinoCore_PlatformUpgradeServer interface {
	Send(*PlatformUpgradeResp) error
	grpc.ServerStream
}

type arduinoCorePlatformUpgradeServer struct {
	grpc.ServerStream
}

func (x *arduinoCorePlatformUpgradeServer) Send(m *PlatformUpgradeResp) error {
	return x.ServerStream.SendMsg(m)
}

func _ArduinoCore_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UploadReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArduinoCoreServer).Upload(m, &arduinoCoreUploadServer{stream})
}

type ArduinoCore_UploadServer interface {
	Send(*UploadResp) error
	grpc.ServerStream
}

type arduinoCoreUploadServer struct {
	grpc.ServerStream
}

func (x *arduinoCoreUploadServer) Send(m *UploadResp) error {
	return x.ServerStream.SendMsg(m)
}

func _ArduinoCore_PlatformSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlatformSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArduinoCoreServer).PlatformSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arduino.ArduinoCore/PlatformSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArduinoCoreServer).PlatformSearch(ctx, req.(*PlatformSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArduinoCore_PlatformList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlatformListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArduinoCoreServer).PlatformList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arduino.ArduinoCore/PlatformList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArduinoCoreServer).PlatformList(ctx, req.(*PlatformListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArduinoCore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "arduino.ArduinoCore",
	HandlerType: (*ArduinoCoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _ArduinoCore_Init_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _ArduinoCore_Destroy_Handler,
		},
		{
			MethodName: "Rescan",
			Handler:    _ArduinoCore_Rescan_Handler,
		},
		{
			MethodName: "BoardDetails",
			Handler:    _ArduinoCore_BoardDetails_Handler,
		},
		{
			MethodName: "PlatformSearch",
			Handler:    _ArduinoCore_PlatformSearch_Handler,
		},
		{
			MethodName: "PlatformList",
			Handler:    _ArduinoCore_PlatformList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateIndex",
			Handler:       _ArduinoCore_UpdateIndex_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Compile",
			Handler:       _ArduinoCore_Compile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PlatformInstall",
			Handler:       _ArduinoCore_PlatformInstall_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PlatformDownload",
			Handler:       _ArduinoCore_PlatformDownload_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PlatformUninstall",
			Handler:       _ArduinoCore_PlatformUninstall_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PlatformUpgrade",
			Handler:       _ArduinoCore_PlatformUpgrade_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Upload",
			Handler:       _ArduinoCore_Upload_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "commands.proto",
}

func init() { proto.RegisterFile("commands.proto", fileDescriptor_commands_595e725b92afdfd7) }

var fileDescriptor_commands_595e725b92afdfd7 = []byte{
	// 644 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0x5d, 0x4f, 0xdb, 0x30,
	0x14, 0x5d, 0x07, 0xa2, 0x70, 0xdb, 0x02, 0x35, 0x68, 0x2b, 0x61, 0x9b, 0x50, 0xc4, 0xc3, 0x5e,
	0x28, 0x08, 0xa4, 0x49, 0xfb, 0xd0, 0x26, 0x68, 0x35, 0xa9, 0xd2, 0x3e, 0x58, 0xb7, 0x3e, 0x6c,
	0x2f, 0xc8, 0x4d, 0x4c, 0xb1, 0x48, 0xed, 0xcc, 0x4e, 0xb5, 0xf5, 0x77, 0xed, 0xdf, 0xed, 0x69,
	0xb6, 0x63, 0x9b, 0xa4, 0x85, 0x4a, 0xe3, 0x29, 0xf5, 0x39, 0xe7, 0x9e, 0xfb, 0xe1, 0xdb, 0x04,
	0xd6, 0x23, 0x3e, 0x1e, 0x63, 0x16, 0xcb, 0x76, 0x2a, 0x78, 0xc6, 0x51, 0x15, 0x8b, 0x78, 0x42,
	0x19, 0x0f, 0xea, 0x9a, 0xe0, 0x2c, 0x87, 0x83, 0xda, 0x90, 0x2b, 0xc2, 0x1e, 0x1a, 0x8a, 0x4a,
	0x69, 0x42, 0xec, 0x11, 0x22, 0x2e, 0xdc, 0xef, 0xfa, 0x24, 0x4d, 0x38, 0xb6, 0xc2, 0xf0, 0x4f,
	0x05, 0x1a, 0x1d, 0xce, 0x2e, 0xe9, 0x68, 0x22, 0x70, 0x46, 0x39, 0x43, 0x2d, 0xa8, 0xc6, 0x38,
	0xc3, 0x5d, 0x2a, 0x5a, 0x95, 0xbd, 0xca, 0xf3, 0xb5, 0xbe, 0x3b, 0xa2, 0x7d, 0x68, 0xc8, 0x6b,
	0x92, 0x45, 0x57, 0x43, 0xce, 0xaf, 0x35, 0xff, 0xd0, 0xf0, 0x65, 0x10, 0x85, 0x50, 0x8f, 0xf9,
	0x2f, 0xa6, 0x73, 0x48, 0x2d, 0x5a, 0x32, 0xa2, 0x12, 0x86, 0xde, 0x42, 0x60, 0xaa, 0xfd, 0x88,
	0x19, 0x1e, 0x11, 0x71, 0x1a, 0xc7, 0x54, 0xe7, 0xc6, 0xc9, 0x40, 0x24, 0xb2, 0xb5, 0xbc, 0xb7,
	0xa4, 0x22, 0x16, 0x28, 0xc2, 0x29, 0x54, 0x7b, 0x8c, 0x66, 0x7d, 0xf2, 0x13, 0xbd, 0x01, 0xd5,
	0x6b, 0xa1, 0x7e, 0x53, 0x74, 0xed, 0xf8, 0x51, 0xdb, 0x4e, 0xa9, 0x5d, 0xea, 0xae, 0x5f, 0x16,
	0xa3, 0x23, 0xd8, 0x4e, 0xe8, 0x50, 0x60, 0x31, 0xbd, 0x18, 0xe7, 0x99, 0x2e, 0x38, 0x4b, 0xa6,
	0xa6, 0xb3, 0xd5, 0x3e, 0xb2, 0x9c, 0x2d, 0xe2, 0xb3, 0x62, 0xc2, 0x97, 0xb0, 0x9a, 0xa7, 0x96,
	0x29, 0x3a, 0x80, 0x55, 0xca, 0x64, 0x86, 0x59, 0x44, 0x6c, 0xda, 0xa6, 0x4f, 0xdb, 0xb3, 0x44,
	0xdf, 0x4b, 0xc2, 0xd7, 0x00, 0x5d, 0x22, 0x33, 0xc1, 0xa7, 0xba, 0xf0, 0xff, 0x0c, 0x6e, 0x40,
	0xcd, 0x07, 0xcb, 0x34, 0x7c, 0x05, 0x6b, 0xea, 0x19, 0x61, 0x76, 0x0f, 0xab, 0x3a, 0x80, 0x8b,
	0x55, 0x4e, 0xef, 0x60, 0x7d, 0x90, 0xaa, 0x2b, 0x26, 0x3d, 0x16, 0x93, 0xdf, 0xf7, 0xb0, 0xfb,
	0x0e, 0x1b, 0x25, 0x03, 0x35, 0x98, 0xf7, 0xd0, 0x74, 0xf7, 0x7d, 0xa1, 0xf6, 0x6c, 0x24, 0x88,
	0x94, 0xd6, 0x6a, 0xc7, 0x5b, 0x75, 0xad, 0xe2, 0xdc, 0x0a, 0xfa, 0x9b, 0xf1, 0x0c, 0x72, 0xfc,
	0x77, 0x05, 0x6a, 0xa7, 0xb9, 0xbc, 0xa3, 0x36, 0x58, 0x55, 0xb6, 0xac, 0x87, 0x8f, 0x36, 0x0b,
	0xf5, 0x98, 0x35, 0x08, 0x9a, 0x33, 0x88, 0x6a, 0xec, 0x01, 0x7a, 0x01, 0x55, 0x3b, 0x33, 0xb4,
	0x75, 0x93, 0xd6, 0x5f, 0x41, 0xb0, 0x3d, 0x0f, 0x9a, 0xb8, 0x13, 0x58, 0xc9, 0x07, 0x84, 0x90,
	0x57, 0xf8, 0x69, 0x07, 0x5b, 0x73, 0x98, 0x09, 0xea, 0x42, 0xad, 0x30, 0x06, 0xf4, 0xd8, 0xab,
	0xca, 0xd3, 0x0d, 0x5a, 0xb7, 0x13, 0xda, 0xe3, 0xa8, 0x82, 0x3a, 0x50, 0x3f, 0xd3, 0x7b, 0xdf,
	0x25, 0x19, 0xa6, 0x89, 0x44, 0x37, 0xea, 0x22, 0xac, 0x7d, 0x76, 0xee, 0x60, 0xd4, 0xf8, 0x55,
	0xdf, 0x9d, 0xfc, 0xff, 0x5f, 0xe8, 0xdb, 0x22, 0xe5, 0xbe, 0x3d, 0x28, 0x53, 0x95, 0xfc, 0x13,
	0x6c, 0x9c, 0x27, 0x38, 0xbb, 0xe4, 0x62, 0x6c, 0xee, 0x39, 0x49, 0xd0, 0xae, 0x97, 0xce, 0x30,
	0xda, 0xe7, 0xc9, 0xdd, 0xa4, 0xf1, 0xfb, 0x02, 0x9b, 0x8e, 0x70, 0x97, 0x8d, 0xe6, 0x63, 0x1c,
	0xa5, 0x1d, 0x9f, 0x2e, 0x60, 0x8d, 0xe5, 0x37, 0x68, 0x3a, 0x66, 0xc0, 0xa8, 0x2d, 0x72, 0x3e,
	0xca, 0x73, 0xda, 0xf4, 0xd9, 0x22, 0x7a, 0xb6, 0xf1, 0x41, 0x3a, 0x12, 0x38, 0x26, 0xb7, 0x34,
	0x6e, 0x99, 0xdb, 0x1b, 0xf7, 0xa4, 0xf1, 0x53, 0x0b, 0x34, 0x30, 0x6f, 0xd9, 0xc2, 0x02, 0xe5,
	0x40, 0x79, 0x81, 0x1c, 0x66, 0x82, 0x7a, 0xb0, 0xee, 0xdc, 0xbe, 0x12, 0x2c, 0xa2, 0x2b, 0x14,
	0xcc, 0xa5, 0xc9, 0x09, 0x6d, 0xb2, 0x7b, 0x27, 0xa7, 0x16, 0x40, 0x6d, 0x91, 0x43, 0x3f, 0x50,
	0x99, 0x15, 0xb6, 0xa8, 0x08, 0x97, 0xb7, 0xa8, 0xcc, 0xc8, 0xf4, 0x6c, 0xff, 0x47, 0x38, 0xa2,
	0xd9, 0xd5, 0x64, 0xd8, 0x56, 0x1f, 0x93, 0x43, 0x2b, 0x73, 0xcf, 0x83, 0x28, 0xa1, 0x87, 0x22,
	0x8d, 0x86, 0x2b, 0xe6, 0x3b, 0x72, 0xf2, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xdc, 0xdd, 0xfe,
	0xa6, 0x06, 0x00, 0x00,
}
