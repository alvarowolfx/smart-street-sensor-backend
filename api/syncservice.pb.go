// Code generated by protoc-gen-go. DO NOT EDIT.
// source: syncservice.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	syncservice.proto

It has these top-level messages:
	RegisterDeviceMessage
	DeviceFilter
	UpdateStatusMessage
	LoginMessage
	Metric
	LoginResponse
	RegisterResponse
	MetricResponse
	StatusResponse
	Device
	DeviceResponse
*/
package api

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

type Status int32

const (
	Status_SUCCESS Status = 0
	Status_FAILED  Status = 1
)

var Status_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILED",
}
var Status_value = map[string]int32{
	"SUCCESS": 0,
	"FAILED":  1,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RegisterDeviceMessage struct {
	DeviceId string `protobuf:"bytes,1,opt,name=deviceId" json:"deviceId,omitempty"`
	Token    string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *RegisterDeviceMessage) Reset()                    { *m = RegisterDeviceMessage{} }
func (m *RegisterDeviceMessage) String() string            { return proto.CompactTextString(m) }
func (*RegisterDeviceMessage) ProtoMessage()               {}
func (*RegisterDeviceMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegisterDeviceMessage) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *RegisterDeviceMessage) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DeviceFilter struct {
	Page    int32 `protobuf:"varint,1,opt,name=page" json:"page,omitempty"`
	PerPage int32 `protobuf:"varint,2,opt,name=perPage" json:"perPage,omitempty"`
}

func (m *DeviceFilter) Reset()                    { *m = DeviceFilter{} }
func (m *DeviceFilter) String() string            { return proto.CompactTextString(m) }
func (*DeviceFilter) ProtoMessage()               {}
func (*DeviceFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeviceFilter) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *DeviceFilter) GetPerPage() int32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

type UpdateStatusMessage struct {
	DeviceId string `protobuf:"bytes,1,opt,name=deviceId" json:"deviceId,omitempty"`
	Status   bool   `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
}

func (m *UpdateStatusMessage) Reset()                    { *m = UpdateStatusMessage{} }
func (m *UpdateStatusMessage) String() string            { return proto.CompactTextString(m) }
func (*UpdateStatusMessage) ProtoMessage()               {}
func (*UpdateStatusMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateStatusMessage) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *UpdateStatusMessage) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type LoginMessage struct {
	Login    string `protobuf:"bytes,1,opt,name=login" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginMessage) Reset()                    { *m = LoginMessage{} }
func (m *LoginMessage) String() string            { return proto.CompactTextString(m) }
func (*LoginMessage) ProtoMessage()               {}
func (*LoginMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LoginMessage) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *LoginMessage) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Metric struct {
	Latitude    int32 `protobuf:"varint,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude   int32 `protobuf:"varint,2,opt,name=longitude" json:"longitude,omitempty"`
	Speed       int32 `protobuf:"varint,3,opt,name=speed" json:"speed,omitempty"`
	Altitude    int32 `protobuf:"varint,4,opt,name=altitude" json:"altitude,omitempty"`
	Ax          int32 `protobuf:"varint,5,opt,name=ax" json:"ax,omitempty"`
	Ay          int32 `protobuf:"varint,6,opt,name=ay" json:"ay,omitempty"`
	Az          int32 `protobuf:"varint,7,opt,name=az" json:"az,omitempty"`
	Gx          int32 `protobuf:"varint,8,opt,name=gx" json:"gx,omitempty"`
	Gy          int32 `protobuf:"varint,9,opt,name=gy" json:"gy,omitempty"`
	Gz          int32 `protobuf:"varint,10,opt,name=gz" json:"gz,omitempty"`
	Mx          int32 `protobuf:"varint,11,opt,name=mx" json:"mx,omitempty"`
	My          int32 `protobuf:"varint,12,opt,name=my" json:"my,omitempty"`
	Mz          int32 `protobuf:"varint,13,opt,name=mz" json:"mz,omitempty"`
	Temperature int32 `protobuf:"varint,14,opt,name=temperature" json:"temperature,omitempty"`
	Humidity    int32 `protobuf:"varint,15,opt,name=humidity" json:"humidity,omitempty"`
	Pressure    int32 `protobuf:"varint,16,opt,name=pressure" json:"pressure,omitempty"`
	Timestamp   int32 `protobuf:"varint,17,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Metric) Reset()                    { *m = Metric{} }
func (m *Metric) String() string            { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()               {}
func (*Metric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Metric) GetLatitude() int32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Metric) GetLongitude() int32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Metric) GetSpeed() int32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *Metric) GetAltitude() int32 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Metric) GetAx() int32 {
	if m != nil {
		return m.Ax
	}
	return 0
}

func (m *Metric) GetAy() int32 {
	if m != nil {
		return m.Ay
	}
	return 0
}

func (m *Metric) GetAz() int32 {
	if m != nil {
		return m.Az
	}
	return 0
}

func (m *Metric) GetGx() int32 {
	if m != nil {
		return m.Gx
	}
	return 0
}

func (m *Metric) GetGy() int32 {
	if m != nil {
		return m.Gy
	}
	return 0
}

func (m *Metric) GetGz() int32 {
	if m != nil {
		return m.Gz
	}
	return 0
}

func (m *Metric) GetMx() int32 {
	if m != nil {
		return m.Mx
	}
	return 0
}

func (m *Metric) GetMy() int32 {
	if m != nil {
		return m.My
	}
	return 0
}

func (m *Metric) GetMz() int32 {
	if m != nil {
		return m.Mz
	}
	return 0
}

func (m *Metric) GetTemperature() int32 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

func (m *Metric) GetHumidity() int32 {
	if m != nil {
		return m.Humidity
	}
	return 0
}

func (m *Metric) GetPressure() int32 {
	if m != nil {
		return m.Pressure
	}
	return 0
}

func (m *Metric) GetTimestamp() int32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type LoginResponse struct {
	Status Status `protobuf:"varint,1,opt,name=status,enum=api.Status" json:"status,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LoginResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_SUCCESS
}

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RegisterResponse struct {
	Status Status `protobuf:"varint,1,opt,name=status,enum=api.Status" json:"status,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RegisterResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_SUCCESS
}

func (m *RegisterResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type MetricResponse struct {
	Status          Status `protobuf:"varint,1,opt,name=status,enum=api.Status" json:"status,omitempty"`
	InsertedMetrics int32  `protobuf:"varint,2,opt,name=insertedMetrics" json:"insertedMetrics,omitempty"`
}

func (m *MetricResponse) Reset()                    { *m = MetricResponse{} }
func (m *MetricResponse) String() string            { return proto.CompactTextString(m) }
func (*MetricResponse) ProtoMessage()               {}
func (*MetricResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MetricResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_SUCCESS
}

func (m *MetricResponse) GetInsertedMetrics() int32 {
	if m != nil {
		return m.InsertedMetrics
	}
	return 0
}

type StatusResponse struct {
	Status  Status `protobuf:"varint,1,opt,name=status,enum=api.Status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *StatusResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_SUCCESS
}

func (m *StatusResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Device struct {
	DeviceId   string  `protobuf:"bytes,1,opt,name=deviceId" json:"deviceId,omitempty"`
	LastMetric *Metric `protobuf:"bytes,2,opt,name=lastMetric" json:"lastMetric,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Device) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *Device) GetLastMetric() *Metric {
	if m != nil {
		return m.LastMetric
	}
	return nil
}

type DeviceResponse struct {
	Status  Status    `protobuf:"varint,1,opt,name=status,enum=api.Status" json:"status,omitempty"`
	Total   int32     `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
	Devices []*Device `protobuf:"bytes,3,rep,name=devices" json:"devices,omitempty"`
}

func (m *DeviceResponse) Reset()                    { *m = DeviceResponse{} }
func (m *DeviceResponse) String() string            { return proto.CompactTextString(m) }
func (*DeviceResponse) ProtoMessage()               {}
func (*DeviceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeviceResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_SUCCESS
}

func (m *DeviceResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *DeviceResponse) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterDeviceMessage)(nil), "api.RegisterDeviceMessage")
	proto.RegisterType((*DeviceFilter)(nil), "api.DeviceFilter")
	proto.RegisterType((*UpdateStatusMessage)(nil), "api.UpdateStatusMessage")
	proto.RegisterType((*LoginMessage)(nil), "api.LoginMessage")
	proto.RegisterType((*Metric)(nil), "api.Metric")
	proto.RegisterType((*LoginResponse)(nil), "api.LoginResponse")
	proto.RegisterType((*RegisterResponse)(nil), "api.RegisterResponse")
	proto.RegisterType((*MetricResponse)(nil), "api.MetricResponse")
	proto.RegisterType((*StatusResponse)(nil), "api.StatusResponse")
	proto.RegisterType((*Device)(nil), "api.Device")
	proto.RegisterType((*DeviceResponse)(nil), "api.DeviceResponse")
	proto.RegisterEnum("api.Status", Status_name, Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuthService service

type AuthServiceClient interface {
	// Register a device and return a token
	RegisterDevice(ctx context.Context, in *RegisterDeviceMessage, opts ...grpc.CallOption) (*RegisterResponse, error)
	// Login user with password and return a token
	Login(ctx context.Context, in *LoginMessage, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) RegisterDevice(ctx context.Context, in *RegisterDeviceMessage, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := grpc.Invoke(ctx, "/api.AuthService/RegisterDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginMessage, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/api.AuthService/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceServer interface {
	// Register a device and return a token
	RegisterDevice(context.Context, *RegisterDeviceMessage) (*RegisterResponse, error)
	// Login user with password and return a token
	Login(context.Context, *LoginMessage) (*LoginResponse, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_RegisterDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDeviceMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RegisterDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AuthService/RegisterDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RegisterDevice(ctx, req.(*RegisterDeviceMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDevice",
			Handler:    _AuthService_RegisterDevice_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "syncservice.proto",
}

// Client API for DeviceHandlingService service

type DeviceHandlingServiceClient interface {
	// Enable and disable devices
	UpdateStatus(ctx context.Context, in *UpdateStatusMessage, opts ...grpc.CallOption) (*StatusResponse, error)
	ListDevices(ctx context.Context, in *DeviceFilter, opts ...grpc.CallOption) (*DeviceResponse, error)
}

type deviceHandlingServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeviceHandlingServiceClient(cc *grpc.ClientConn) DeviceHandlingServiceClient {
	return &deviceHandlingServiceClient{cc}
}

func (c *deviceHandlingServiceClient) UpdateStatus(ctx context.Context, in *UpdateStatusMessage, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/api.DeviceHandlingService/UpdateStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceHandlingServiceClient) ListDevices(ctx context.Context, in *DeviceFilter, opts ...grpc.CallOption) (*DeviceResponse, error) {
	out := new(DeviceResponse)
	err := grpc.Invoke(ctx, "/api.DeviceHandlingService/ListDevices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeviceHandlingService service

type DeviceHandlingServiceServer interface {
	// Enable and disable devices
	UpdateStatus(context.Context, *UpdateStatusMessage) (*StatusResponse, error)
	ListDevices(context.Context, *DeviceFilter) (*DeviceResponse, error)
}

func RegisterDeviceHandlingServiceServer(s *grpc.Server, srv DeviceHandlingServiceServer) {
	s.RegisterService(&_DeviceHandlingService_serviceDesc, srv)
}

func _DeviceHandlingService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceHandlingServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceHandlingService/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceHandlingServiceServer).UpdateStatus(ctx, req.(*UpdateStatusMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceHandlingService_ListDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceHandlingServiceServer).ListDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceHandlingService/ListDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceHandlingServiceServer).ListDevices(ctx, req.(*DeviceFilter))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceHandlingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DeviceHandlingService",
	HandlerType: (*DeviceHandlingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateStatus",
			Handler:    _DeviceHandlingService_UpdateStatus_Handler,
		},
		{
			MethodName: "ListDevices",
			Handler:    _DeviceHandlingService_ListDevices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "syncservice.proto",
}

// Client API for MetricService service

type MetricServiceClient interface {
	// Insert new packets
	SendMetric(ctx context.Context, opts ...grpc.CallOption) (MetricService_SendMetricClient, error)
}

type metricServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetricServiceClient(cc *grpc.ClientConn) MetricServiceClient {
	return &metricServiceClient{cc}
}

func (c *metricServiceClient) SendMetric(ctx context.Context, opts ...grpc.CallOption) (MetricService_SendMetricClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MetricService_serviceDesc.Streams[0], c.cc, "/api.MetricService/SendMetric", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricServiceSendMetricClient{stream}
	return x, nil
}

type MetricService_SendMetricClient interface {
	Send(*Metric) error
	CloseAndRecv() (*MetricResponse, error)
	grpc.ClientStream
}

type metricServiceSendMetricClient struct {
	grpc.ClientStream
}

func (x *metricServiceSendMetricClient) Send(m *Metric) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricServiceSendMetricClient) CloseAndRecv() (*MetricResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MetricResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MetricService service

type MetricServiceServer interface {
	// Insert new packets
	SendMetric(MetricService_SendMetricServer) error
}

func RegisterMetricServiceServer(s *grpc.Server, srv MetricServiceServer) {
	s.RegisterService(&_MetricService_serviceDesc, srv)
}

func _MetricService_SendMetric_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricServiceServer).SendMetric(&metricServiceSendMetricServer{stream})
}

type MetricService_SendMetricServer interface {
	SendAndClose(*MetricResponse) error
	Recv() (*Metric, error)
	grpc.ServerStream
}

type metricServiceSendMetricServer struct {
	grpc.ServerStream
}

func (x *metricServiceSendMetricServer) SendAndClose(m *MetricResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricServiceSendMetricServer) Recv() (*Metric, error) {
	m := new(Metric)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MetricService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.MetricService",
	HandlerType: (*MetricServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMetric",
			Handler:       _MetricService_SendMetric_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "syncservice.proto",
}
var MetricService_serviceDesc = _MetricService_serviceDesc

func init() { proto.RegisterFile("syncservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 664 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x4d, 0x6f, 0x13, 0x3d,
	0x10, 0xc7, 0xf3, 0xd2, 0x6c, 0xda, 0xd9, 0x74, 0x9b, 0xba, 0xed, 0x23, 0x2b, 0x7a, 0x0e, 0x65,
	0x11, 0x52, 0x05, 0x52, 0x85, 0xc2, 0x81, 0x0b, 0x12, 0x54, 0x7d, 0x11, 0x41, 0xad, 0x80, 0x44,
	0x3d, 0x23, 0xd3, 0xb5, 0xb6, 0x16, 0xfb, 0xa6, 0xb5, 0x03, 0x49, 0x3e, 0x00, 0x77, 0xbe, 0x2f,
	0x07, 0xb4, 0x33, 0x76, 0xb2, 0x41, 0x05, 0x51, 0xc4, 0x2d, 0xbf, 0x99, 0xf1, 0xdf, 0x33, 0x9e,
	0xd9, 0x09, 0xec, 0xea, 0x79, 0x76, 0xa3, 0x65, 0xf9, 0x59, 0xdd, 0xc8, 0xe3, 0xa2, 0xcc, 0x4d,
	0xce, 0xda, 0xa2, 0x50, 0xe1, 0x08, 0x0e, 0xc6, 0x32, 0x56, 0xda, 0xc8, 0xf2, 0x4c, 0x56, 0xce,
	0x2b, 0xa9, 0xb5, 0x88, 0x25, 0x1b, 0xc0, 0x66, 0x84, 0x86, 0x51, 0xc4, 0x9b, 0x87, 0xcd, 0xa3,
	0xad, 0xf1, 0x92, 0xd9, 0x3e, 0x74, 0x4c, 0xfe, 0x49, 0x66, 0xbc, 0x85, 0x0e, 0x82, 0xf0, 0x05,
	0xf4, 0x48, 0xe2, 0x42, 0x25, 0x46, 0x96, 0x8c, 0xc1, 0x46, 0x21, 0x62, 0x89, 0xa7, 0x3b, 0x63,
	0xfc, 0xcd, 0x38, 0x74, 0x0b, 0x59, 0xbe, 0xab, 0xcc, 0x2d, 0x34, 0x3b, 0x0c, 0x47, 0xb0, 0x77,
	0x5d, 0x44, 0xc2, 0xc8, 0x89, 0x11, 0x66, 0xaa, 0xff, 0x24, 0x8d, 0xff, 0xc0, 0xd3, 0x18, 0x8c,
	0x5a, 0x9b, 0x63, 0x4b, 0xe1, 0x2b, 0xe8, 0x5d, 0xe6, 0xb1, 0xca, 0x9c, 0xc6, 0x3e, 0x74, 0x92,
	0x8a, 0xad, 0x00, 0x41, 0xa5, 0x5c, 0x08, 0xad, 0xbf, 0xe4, 0x65, 0x64, 0xeb, 0x58, 0x72, 0xf8,
	0xbd, 0x05, 0xde, 0x95, 0x34, 0xa5, 0xba, 0xa9, 0xc2, 0x12, 0x61, 0x94, 0x99, 0x46, 0xae, 0x92,
	0x25, 0xb3, 0xff, 0x61, 0x2b, 0xc9, 0xb3, 0x98, 0x9c, 0x54, 0xcf, 0xca, 0x50, 0x5d, 0xab, 0x0b,
	0x29, 0x23, 0xde, 0x46, 0x0f, 0x41, 0xa5, 0x27, 0x12, 0xab, 0xb7, 0x41, 0x7a, 0x8e, 0x59, 0x00,
	0x2d, 0x31, 0xe3, 0x1d, 0xb4, 0xb6, 0xc4, 0x0c, 0x79, 0xce, 0x3d, 0xcb, 0x73, 0xe4, 0x05, 0xef,
	0x5a, 0x5e, 0x54, 0x1c, 0xcf, 0xf8, 0x26, 0x71, 0x8c, 0xf1, 0xf1, 0x9c, 0x6f, 0x59, 0xc6, 0xf8,
	0x78, 0xc1, 0xc1, 0x32, 0xc6, 0xa7, 0x33, 0xee, 0x13, 0xa7, 0x18, 0x9f, 0xce, 0x79, 0xcf, 0x32,
	0xc6, 0xa7, 0x0b, 0xbe, 0x6d, 0x79, 0xc1, 0x0e, 0xc1, 0x37, 0x32, 0x2d, 0x64, 0x29, 0xcc, 0xb4,
	0x94, 0x3c, 0x40, 0x47, 0xdd, 0x54, 0x55, 0x73, 0x3b, 0x4d, 0x55, 0xa4, 0xcc, 0x9c, 0xef, 0x50,
	0x35, 0x8e, 0xf1, 0x81, 0x4b, 0xa9, 0x75, 0x75, 0xb4, 0x4f, 0x3e, 0xc7, 0xd5, 0xcb, 0x19, 0x95,
	0x4a, 0x6d, 0x44, 0x5a, 0xf0, 0x5d, 0x7a, 0xb9, 0xa5, 0x21, 0x7c, 0x03, 0xdb, 0xd8, 0xc0, 0xb1,
	0xd4, 0x45, 0x9e, 0x69, 0xc9, 0x1e, 0x2e, 0x3b, 0x5d, 0xb5, 0x20, 0x18, 0xfa, 0xc7, 0xa2, 0x50,
	0xc7, 0x34, 0x29, 0xae, 0xed, 0xbf, 0x98, 0xca, 0x2b, 0xe8, 0xbb, 0x01, 0xff, 0x17, 0x72, 0x1f,
	0x20, 0xa0, 0xc1, 0xb8, 0x9f, 0xd8, 0x11, 0xec, 0xa8, 0x4c, 0xcb, 0xd2, 0xc8, 0x88, 0x8e, 0x6b,
	0x3b, 0x2f, 0x3f, 0x9b, 0xc3, 0xb7, 0x10, 0xd8, 0xb3, 0xf7, 0xba, 0x80, 0x43, 0x37, 0xa5, 0x71,
	0xb7, 0xf9, 0x3a, 0x0c, 0xdf, 0x83, 0x47, 0x9f, 0xe5, 0x6f, 0xbf, 0xa5, 0x27, 0x00, 0x89, 0xd0,
	0x86, 0xb2, 0x40, 0x09, 0xdf, 0x5e, 0x64, 0xcb, 0xad, 0xb9, 0xc3, 0x12, 0x02, 0x92, 0xfc, 0x8b,
	0x17, 0x35, 0x22, 0xb1, 0xa5, 0x13, 0xb0, 0x47, 0xd0, 0xa5, 0x2c, 0x34, 0x6f, 0x1f, 0xb6, 0x97,
	0xd7, 0xda, 0x0b, 0x9c, 0xef, 0xf1, 0x03, 0xf0, 0x48, 0x8e, 0xf9, 0xd0, 0x9d, 0x5c, 0x9f, 0x9e,
	0x9e, 0x4f, 0x26, 0xfd, 0x06, 0x03, 0xf0, 0x2e, 0x4e, 0x46, 0x97, 0xe7, 0x67, 0xfd, 0xe6, 0xf0,
	0x6b, 0x13, 0xfc, 0x93, 0xa9, 0xb9, 0x9d, 0xd0, 0x9a, 0x63, 0xe7, 0x10, 0xac, 0xef, 0x36, 0x36,
	0x40, 0xe9, 0x3b, 0x17, 0xde, 0xe0, 0x60, 0xcd, 0xe7, 0x2a, 0x0b, 0x1b, 0xec, 0x29, 0x74, 0x70,
	0x1a, 0xd9, 0x2e, 0x46, 0xd4, 0x57, 0xcb, 0x80, 0xad, 0x4c, 0xab, 0x13, 0xc3, 0x6f, 0x4d, 0x38,
	0x20, 0xf1, 0xd7, 0x22, 0x8b, 0x12, 0x95, 0xc5, 0x2e, 0xa5, 0x97, 0xd0, 0xab, 0x6f, 0x39, 0xc6,
	0xf1, 0xfc, 0x1d, 0x8b, 0x6f, 0xb0, 0x57, 0x7f, 0xc1, 0x55, 0x32, 0xcf, 0xc1, 0xbf, 0x54, 0xda,
	0x90, 0xba, 0xb6, 0x29, 0xd5, 0xd7, 0xae, 0x3d, 0xb8, 0xde, 0x9f, 0xb0, 0x31, 0x3c, 0x85, 0x6d,
	0xea, 0x9e, 0x4b, 0x65, 0x08, 0x30, 0x91, 0x99, 0x9d, 0x3b, 0x56, 0xef, 0xb5, 0x95, 0x58, 0x9f,
	0xf3, 0xb0, 0x71, 0xd4, 0xfc, 0xe8, 0xe1, 0x3f, 0xc7, 0xb3, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xbb, 0x66, 0xcc, 0x96, 0x4e, 0x06, 0x00, 0x00,
}
