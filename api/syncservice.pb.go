// Code generated by protoc-gen-go. DO NOT EDIT.
// source: syncservice.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	syncservice.proto

It has these top-level messages:
	Location
	AccelerometerData
	GyroscopeData
	MagnetometerData
	ClimateData
	SyncPacket
	RequestSummary
	SyncResponse
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

type Location struct {
	Latitude  int32 `protobuf:"varint,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude" json:"longitude,omitempty"`
	Speed     int32 `protobuf:"varint,3,opt,name=speed" json:"speed,omitempty"`
	Altitude  int32 `protobuf:"varint,4,opt,name=altitude" json:"altitude,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Location) GetLatitude() int32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetLongitude() int32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Location) GetSpeed() int32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *Location) GetAltitude() int32 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

type AccelerometerData struct {
	Ax int32 `protobuf:"varint,1,opt,name=ax" json:"ax,omitempty"`
	Ay int32 `protobuf:"varint,2,opt,name=ay" json:"ay,omitempty"`
	Az int32 `protobuf:"varint,3,opt,name=az" json:"az,omitempty"`
}

func (m *AccelerometerData) Reset()                    { *m = AccelerometerData{} }
func (m *AccelerometerData) String() string            { return proto.CompactTextString(m) }
func (*AccelerometerData) ProtoMessage()               {}
func (*AccelerometerData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AccelerometerData) GetAx() int32 {
	if m != nil {
		return m.Ax
	}
	return 0
}

func (m *AccelerometerData) GetAy() int32 {
	if m != nil {
		return m.Ay
	}
	return 0
}

func (m *AccelerometerData) GetAz() int32 {
	if m != nil {
		return m.Az
	}
	return 0
}

type GyroscopeData struct {
	Gx int32 `protobuf:"varint,1,opt,name=gx" json:"gx,omitempty"`
	Gy int32 `protobuf:"varint,2,opt,name=gy" json:"gy,omitempty"`
	Gz int32 `protobuf:"varint,3,opt,name=gz" json:"gz,omitempty"`
}

func (m *GyroscopeData) Reset()                    { *m = GyroscopeData{} }
func (m *GyroscopeData) String() string            { return proto.CompactTextString(m) }
func (*GyroscopeData) ProtoMessage()               {}
func (*GyroscopeData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GyroscopeData) GetGx() int32 {
	if m != nil {
		return m.Gx
	}
	return 0
}

func (m *GyroscopeData) GetGy() int32 {
	if m != nil {
		return m.Gy
	}
	return 0
}

func (m *GyroscopeData) GetGz() int32 {
	if m != nil {
		return m.Gz
	}
	return 0
}

type MagnetometerData struct {
	Mx int32 `protobuf:"varint,1,opt,name=mx" json:"mx,omitempty"`
	My int32 `protobuf:"varint,2,opt,name=my" json:"my,omitempty"`
	Mz int32 `protobuf:"varint,3,opt,name=mz" json:"mz,omitempty"`
}

func (m *MagnetometerData) Reset()                    { *m = MagnetometerData{} }
func (m *MagnetometerData) String() string            { return proto.CompactTextString(m) }
func (*MagnetometerData) ProtoMessage()               {}
func (*MagnetometerData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MagnetometerData) GetMx() int32 {
	if m != nil {
		return m.Mx
	}
	return 0
}

func (m *MagnetometerData) GetMy() int32 {
	if m != nil {
		return m.My
	}
	return 0
}

func (m *MagnetometerData) GetMz() int32 {
	if m != nil {
		return m.Mz
	}
	return 0
}

type ClimateData struct {
	Temperature int32 `protobuf:"varint,1,opt,name=temperature" json:"temperature,omitempty"`
	Humidity    int32 `protobuf:"varint,2,opt,name=humidity" json:"humidity,omitempty"`
	Pressure    int32 `protobuf:"varint,3,opt,name=pressure" json:"pressure,omitempty"`
}

func (m *ClimateData) Reset()                    { *m = ClimateData{} }
func (m *ClimateData) String() string            { return proto.CompactTextString(m) }
func (*ClimateData) ProtoMessage()               {}
func (*ClimateData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ClimateData) GetTemperature() int32 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

func (m *ClimateData) GetHumidity() int32 {
	if m != nil {
		return m.Humidity
	}
	return 0
}

func (m *ClimateData) GetPressure() int32 {
	if m != nil {
		return m.Pressure
	}
	return 0
}

type SyncPacket struct {
	Location    *Location          `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	AccData     *AccelerometerData `protobuf:"bytes,2,opt,name=accData" json:"accData,omitempty"`
	GyroData    *GyroscopeData     `protobuf:"bytes,3,opt,name=gyroData" json:"gyroData,omitempty"`
	MagData     *MagnetometerData  `protobuf:"bytes,4,opt,name=magData" json:"magData,omitempty"`
	ClimateData *ClimateData       `protobuf:"bytes,5,opt,name=climateData" json:"climateData,omitempty"`
	Timestamp   int32              `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *SyncPacket) Reset()                    { *m = SyncPacket{} }
func (m *SyncPacket) String() string            { return proto.CompactTextString(m) }
func (*SyncPacket) ProtoMessage()               {}
func (*SyncPacket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SyncPacket) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *SyncPacket) GetAccData() *AccelerometerData {
	if m != nil {
		return m.AccData
	}
	return nil
}

func (m *SyncPacket) GetGyroData() *GyroscopeData {
	if m != nil {
		return m.GyroData
	}
	return nil
}

func (m *SyncPacket) GetMagData() *MagnetometerData {
	if m != nil {
		return m.MagData
	}
	return nil
}

func (m *SyncPacket) GetClimateData() *ClimateData {
	if m != nil {
		return m.ClimateData
	}
	return nil
}

func (m *SyncPacket) GetTimestamp() int32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type RequestSummary struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *RequestSummary) Reset()                    { *m = RequestSummary{} }
func (m *RequestSummary) String() string            { return proto.CompactTextString(m) }
func (*RequestSummary) ProtoMessage()               {}
func (*RequestSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RequestSummary) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SyncResponse struct {
	Status          Status `protobuf:"varint,1,opt,name=status,enum=api.Status" json:"status,omitempty"`
	InsertedPackets int32  `protobuf:"varint,2,opt,name=insertedPackets" json:"insertedPackets,omitempty"`
}

func (m *SyncResponse) Reset()                    { *m = SyncResponse{} }
func (m *SyncResponse) String() string            { return proto.CompactTextString(m) }
func (*SyncResponse) ProtoMessage()               {}
func (*SyncResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SyncResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_SUCCESS
}

func (m *SyncResponse) GetInsertedPackets() int32 {
	if m != nil {
		return m.InsertedPackets
	}
	return 0
}

func init() {
	proto.RegisterType((*Location)(nil), "api.Location")
	proto.RegisterType((*AccelerometerData)(nil), "api.AccelerometerData")
	proto.RegisterType((*GyroscopeData)(nil), "api.GyroscopeData")
	proto.RegisterType((*MagnetometerData)(nil), "api.MagnetometerData")
	proto.RegisterType((*ClimateData)(nil), "api.ClimateData")
	proto.RegisterType((*SyncPacket)(nil), "api.SyncPacket")
	proto.RegisterType((*RequestSummary)(nil), "api.RequestSummary")
	proto.RegisterType((*SyncResponse)(nil), "api.SyncResponse")
	proto.RegisterEnum("api.Status", Status_name, Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SyncService service

type SyncServiceClient interface {
	// Insert new packets
	SendTelemetry(ctx context.Context, opts ...grpc.CallOption) (SyncService_SendTelemetryClient, error)
	// Get last job summary
	GetLastJobSummary(ctx context.Context, in *RequestSummary, opts ...grpc.CallOption) (*SyncResponse, error)
}

type syncServiceClient struct {
	cc *grpc.ClientConn
}

func NewSyncServiceClient(cc *grpc.ClientConn) SyncServiceClient {
	return &syncServiceClient{cc}
}

func (c *syncServiceClient) SendTelemetry(ctx context.Context, opts ...grpc.CallOption) (SyncService_SendTelemetryClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SyncService_serviceDesc.Streams[0], c.cc, "/api.SyncService/SendTelemetry", opts...)
	if err != nil {
		return nil, err
	}
	x := &syncServiceSendTelemetryClient{stream}
	return x, nil
}

type SyncService_SendTelemetryClient interface {
	Send(*SyncPacket) error
	CloseAndRecv() (*SyncResponse, error)
	grpc.ClientStream
}

type syncServiceSendTelemetryClient struct {
	grpc.ClientStream
}

func (x *syncServiceSendTelemetryClient) Send(m *SyncPacket) error {
	return x.ClientStream.SendMsg(m)
}

func (x *syncServiceSendTelemetryClient) CloseAndRecv() (*SyncResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SyncResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *syncServiceClient) GetLastJobSummary(ctx context.Context, in *RequestSummary, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := grpc.Invoke(ctx, "/api.SyncService/GetLastJobSummary", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SyncService service

type SyncServiceServer interface {
	// Insert new packets
	SendTelemetry(SyncService_SendTelemetryServer) error
	// Get last job summary
	GetLastJobSummary(context.Context, *RequestSummary) (*SyncResponse, error)
}

func RegisterSyncServiceServer(s *grpc.Server, srv SyncServiceServer) {
	s.RegisterService(&_SyncService_serviceDesc, srv)
}

func _SyncService_SendTelemetry_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SyncServiceServer).SendTelemetry(&syncServiceSendTelemetryServer{stream})
}

type SyncService_SendTelemetryServer interface {
	SendAndClose(*SyncResponse) error
	Recv() (*SyncPacket, error)
	grpc.ServerStream
}

type syncServiceSendTelemetryServer struct {
	grpc.ServerStream
}

func (x *syncServiceSendTelemetryServer) SendAndClose(m *SyncResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *syncServiceSendTelemetryServer) Recv() (*SyncPacket, error) {
	m := new(SyncPacket)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SyncService_GetLastJobSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSummary)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetLastJobSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncService/GetLastJobSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetLastJobSummary(ctx, req.(*RequestSummary))
	}
	return interceptor(ctx, in, info, handler)
}

var _SyncService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.SyncService",
	HandlerType: (*SyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLastJobSummary",
			Handler:    _SyncService_GetLastJobSummary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendTelemetry",
			Handler:       _SyncService_SendTelemetry_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "syncservice.proto",
}

func init() { proto.RegisterFile("syncservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x86, 0x1b, 0x37, 0x5f, 0x1d, 0x93, 0x34, 0x59, 0x3e, 0x14, 0x55, 0x1c, 0x8a, 0xe1, 0x10,
	0x38, 0x04, 0x14, 0x0e, 0x9c, 0x10, 0x2a, 0x69, 0xa9, 0x40, 0x41, 0x42, 0x36, 0x1c, 0x39, 0x4c,
	0x9d, 0x91, 0xb1, 0xf0, 0xda, 0x66, 0x77, 0x5d, 0xd5, 0xb9, 0xf3, 0x2f, 0xf8, 0xb1, 0xc8, 0x63,
	0x7b, 0x9b, 0x14, 0x6e, 0x99, 0x99, 0x77, 0x9e, 0x99, 0xbc, 0xb3, 0x09, 0x4c, 0x75, 0x99, 0x86,
	0x9a, 0xd4, 0x75, 0x1c, 0xd2, 0x22, 0x57, 0x99, 0xc9, 0xc4, 0x21, 0xe6, 0xb1, 0x77, 0x0d, 0xc3,
	0x75, 0x16, 0xa2, 0x89, 0xb3, 0x54, 0x9c, 0xc0, 0x30, 0x41, 0x13, 0x9b, 0x62, 0x43, 0xb3, 0xce,
	0x69, 0x67, 0xde, 0xf3, 0x6d, 0x2c, 0x1e, 0xc3, 0x51, 0x92, 0xa5, 0x51, 0x5d, 0x74, 0xb8, 0x78,
	0x9b, 0x10, 0x0f, 0xa0, 0xa7, 0x73, 0xa2, 0xcd, 0xec, 0x90, 0x2b, 0x75, 0x50, 0xf1, 0x30, 0x69,
	0x78, 0xdd, 0x9a, 0xd7, 0xc6, 0xde, 0x0a, 0xa6, 0x67, 0x61, 0x48, 0x09, 0xa9, 0x4c, 0x92, 0x21,
	0x75, 0x8e, 0x06, 0xc5, 0x18, 0x1c, 0xbc, 0x69, 0x46, 0x3b, 0x78, 0xc3, 0x71, 0xd9, 0x4c, 0x73,
	0xb0, 0xe4, 0x78, 0xdb, 0xcc, 0x70, 0x70, 0xeb, 0xbd, 0x83, 0xd1, 0x65, 0xa9, 0x32, 0x1d, 0x66,
	0x39, 0xb5, 0x80, 0xc8, 0x02, 0x22, 0x06, 0x44, 0x16, 0x10, 0x31, 0x20, 0xb2, 0x80, 0x68, 0xeb,
	0xbd, 0x87, 0xc9, 0x67, 0x8c, 0x52, 0x32, 0xfb, 0x4b, 0x48, 0xcb, 0x90, 0xcc, 0x90, 0x96, 0x21,
	0x99, 0x21, 0x2d, 0x43, 0x6e, 0xbd, 0x08, 0xdc, 0x55, 0x12, 0x4b, 0x34, 0xf5, 0x0a, 0xa7, 0xe0,
	0x1a, 0x92, 0x39, 0x29, 0x34, 0x85, 0x6a, 0x7d, 0xdc, 0x4d, 0x55, 0xb6, 0xfc, 0x28, 0x64, 0xbc,
	0x89, 0x4d, 0x8b, 0xb5, 0x71, 0x55, 0xcb, 0x15, 0x69, 0x5d, 0xb5, 0xd6, 0x23, 0x6c, 0xec, 0xfd,
	0x71, 0x00, 0x82, 0x32, 0x0d, 0xbf, 0x60, 0xf8, 0x93, 0x8c, 0x78, 0x0e, 0xc3, 0xa4, 0xb9, 0x1c,
	0x4f, 0x71, 0x97, 0xa3, 0x05, 0xe6, 0xf1, 0xa2, 0x3d, 0xa7, 0x6f, 0xcb, 0xe2, 0x15, 0x0c, 0x30,
	0x0c, 0xab, 0xf5, 0x78, 0xa0, 0xbb, 0x7c, 0xc4, 0xca, 0x7f, 0x0e, 0xe0, 0xb7, 0x32, 0xb1, 0x80,
	0x61, 0x54, 0xaa, 0x8c, 0x5b, 0x0e, 0xb9, 0x45, 0x70, 0xcb, 0x9e, 0xdd, 0xbe, 0xd5, 0x88, 0x97,
	0x30, 0x90, 0x18, 0xb1, 0xbc, 0xcb, 0xf2, 0x87, 0x2c, 0xbf, 0x6b, 0xae, 0xdf, 0xaa, 0xc4, 0x12,
	0xdc, 0xf0, 0xd6, 0xb5, 0x59, 0x8f, 0x9b, 0x26, 0xdc, 0xb4, 0xe3, 0xa6, 0xbf, 0x2b, 0xaa, 0xde,
	0xa0, 0x89, 0x25, 0x69, 0x83, 0x32, 0x9f, 0xf5, 0xeb, 0x37, 0x68, 0x13, 0xde, 0x33, 0x18, 0xfb,
	0xf4, 0xab, 0x20, 0x6d, 0x82, 0x42, 0x4a, 0x54, 0xa5, 0x10, 0xd0, 0x4d, 0x51, 0xd6, 0x37, 0x38,
	0xf2, 0xf9, 0xb3, 0xf7, 0x1d, 0xee, 0x55, 0x1e, 0xfa, 0xa4, 0xf3, 0x2c, 0xd5, 0x24, 0x9e, 0x42,
	0x5f, 0x1b, 0x34, 0x85, 0x66, 0xd5, 0x78, 0xe9, 0xf2, 0x0a, 0x01, 0xa7, 0xfc, 0xa6, 0x24, 0xe6,
	0x70, 0x1c, 0xa7, 0x9a, 0x94, 0xa1, 0x4d, 0x6d, 0xbe, 0x6e, 0x0e, 0x77, 0x37, 0xfd, 0xe2, 0x09,
	0xf4, 0xeb, 0x5e, 0xe1, 0xc2, 0x20, 0xf8, 0xb6, 0x5a, 0x5d, 0x04, 0xc1, 0xe4, 0x40, 0x00, 0xf4,
	0x3f, 0x9c, 0x7d, 0x5c, 0x5f, 0x9c, 0x4f, 0x3a, 0xcb, 0xdf, 0x1d, 0x70, 0xab, 0x15, 0x82, 0xfa,
	0xc7, 0x28, 0xde, 0xc0, 0x28, 0xa0, 0x74, 0xf3, 0x95, 0x12, 0x92, 0x64, 0x54, 0x29, 0x8e, 0xeb,
	0x15, 0xec, 0xa5, 0x4f, 0xa6, 0x36, 0xd1, 0xae, 0xed, 0x1d, 0xcc, 0x3b, 0xe2, 0x2d, 0x4c, 0x2f,
	0xc9, 0xac, 0x51, 0x9b, 0x4f, 0xd9, 0x55, 0xfb, 0x9d, 0xef, 0xb3, 0x76, 0xdf, 0x88, 0xff, 0x02,
	0xae, 0xfa, 0xfc, 0x2f, 0xf0, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x32, 0xc0, 0x2f,
	0x1a, 0x04, 0x00, 0x00,
}