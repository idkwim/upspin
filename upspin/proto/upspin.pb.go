// Code generated by protoc-gen-go.
// source: upspin.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	upspin.proto

It has these top-level messages:
	Endpoint
	Location
	AuthenticateRequest
	Signature
	AuthenticateResponse
	ConfigureRequest
	ConfigureResponse
	EndpointRequest
	EndpointResponse
	ServerUserNameRequest
	ServerUserNameResponse
	PingRequest
	PingResponse
	StoreGetRequest
	StoreGetResponse
	StorePutRequest
	StorePutResponse
	StoreDeleteRequest
	StoreDeleteResponse
	UserLookupRequest
	UserLookupResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto1.ProtoPackageIsVersion1

// Endpoint mirrors upspin.Endpoint.
type Endpoint struct {
	Transport int32  `protobuf:"varint,1,opt,name=transport" json:"transport,omitempty"`
	NetAddr   string `protobuf:"bytes,2,opt,name=net_addr,json=netAddr" json:"net_addr,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto1.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Location mirrors upspin.Location.
type Location struct {
	Endpoint  *Endpoint `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	Reference string    `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto1.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Location) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type AuthenticateRequest struct {
	UserName  string     `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Now       string     `protobuf:"bytes,2,opt,name=now" json:"now,omitempty"`
	Signature *Signature `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
}

func (m *AuthenticateRequest) Reset()                    { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string            { return proto1.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()               {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AuthenticateRequest) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Signature struct {
	R string `protobuf:"bytes,1,opt,name=r" json:"r,omitempty"`
	S string `protobuf:"bytes,2,opt,name=s" json:"s,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto1.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type AuthenticateResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *AuthenticateResponse) Reset()                    { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string            { return proto1.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()               {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ConfigureRequest struct {
	Options []string `protobuf:"bytes,1,rep,name=options" json:"options,omitempty"`
}

func (m *ConfigureRequest) Reset()                    { *m = ConfigureRequest{} }
func (m *ConfigureRequest) String() string            { return proto1.CompactTextString(m) }
func (*ConfigureRequest) ProtoMessage()               {}
func (*ConfigureRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ConfigureResponse struct {
}

func (m *ConfigureResponse) Reset()                    { *m = ConfigureResponse{} }
func (m *ConfigureResponse) String() string            { return proto1.CompactTextString(m) }
func (*ConfigureResponse) ProtoMessage()               {}
func (*ConfigureResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type EndpointRequest struct {
}

func (m *EndpointRequest) Reset()                    { *m = EndpointRequest{} }
func (m *EndpointRequest) String() string            { return proto1.CompactTextString(m) }
func (*EndpointRequest) ProtoMessage()               {}
func (*EndpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type EndpointResponse struct {
	Endpoint *Endpoint `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *EndpointResponse) Reset()                    { *m = EndpointResponse{} }
func (m *EndpointResponse) String() string            { return proto1.CompactTextString(m) }
func (*EndpointResponse) ProtoMessage()               {}
func (*EndpointResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *EndpointResponse) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type ServerUserNameRequest struct {
}

func (m *ServerUserNameRequest) Reset()                    { *m = ServerUserNameRequest{} }
func (m *ServerUserNameRequest) String() string            { return proto1.CompactTextString(m) }
func (*ServerUserNameRequest) ProtoMessage()               {}
func (*ServerUserNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type ServerUserNameResponse struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
}

func (m *ServerUserNameResponse) Reset()                    { *m = ServerUserNameResponse{} }
func (m *ServerUserNameResponse) String() string            { return proto1.CompactTextString(m) }
func (*ServerUserNameResponse) ProtoMessage()               {}
func (*ServerUserNameResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type PingRequest struct {
	PingSequence int32 `protobuf:"varint,1,opt,name=ping_sequence,json=pingSequence" json:"ping_sequence,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto1.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type PingResponse struct {
	PingSequence int32 `protobuf:"varint,1,opt,name=ping_sequence,json=pingSequence" json:"ping_sequence,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto1.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type StoreGetRequest struct {
	Reference string `protobuf:"bytes,1,opt,name=reference" json:"reference,omitempty"`
}

func (m *StoreGetRequest) Reset()                    { *m = StoreGetRequest{} }
func (m *StoreGetRequest) String() string            { return proto1.CompactTextString(m) }
func (*StoreGetRequest) ProtoMessage()               {}
func (*StoreGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type StoreGetResponse struct {
	Data      []byte      `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Locations []*Location `protobuf:"bytes,2,rep,name=locations" json:"locations,omitempty"`
}

func (m *StoreGetResponse) Reset()                    { *m = StoreGetResponse{} }
func (m *StoreGetResponse) String() string            { return proto1.CompactTextString(m) }
func (*StoreGetResponse) ProtoMessage()               {}
func (*StoreGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *StoreGetResponse) GetLocations() []*Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

type StorePutRequest struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *StorePutRequest) Reset()                    { *m = StorePutRequest{} }
func (m *StorePutRequest) String() string            { return proto1.CompactTextString(m) }
func (*StorePutRequest) ProtoMessage()               {}
func (*StorePutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

type StorePutResponse struct {
	Reference string `protobuf:"bytes,1,opt,name=reference" json:"reference,omitempty"`
}

func (m *StorePutResponse) Reset()                    { *m = StorePutResponse{} }
func (m *StorePutResponse) String() string            { return proto1.CompactTextString(m) }
func (*StorePutResponse) ProtoMessage()               {}
func (*StorePutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

type StoreDeleteRequest struct {
	Reference string `protobuf:"bytes,1,opt,name=reference" json:"reference,omitempty"`
}

func (m *StoreDeleteRequest) Reset()                    { *m = StoreDeleteRequest{} }
func (m *StoreDeleteRequest) String() string            { return proto1.CompactTextString(m) }
func (*StoreDeleteRequest) ProtoMessage()               {}
func (*StoreDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

type StoreDeleteResponse struct {
}

func (m *StoreDeleteResponse) Reset()                    { *m = StoreDeleteResponse{} }
func (m *StoreDeleteResponse) String() string            { return proto1.CompactTextString(m) }
func (*StoreDeleteResponse) ProtoMessage()               {}
func (*StoreDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

type UserLookupRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=UserName,json=userName" json:"UserName,omitempty"`
}

func (m *UserLookupRequest) Reset()                    { *m = UserLookupRequest{} }
func (m *UserLookupRequest) String() string            { return proto1.CompactTextString(m) }
func (*UserLookupRequest) ProtoMessage()               {}
func (*UserLookupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

type UserLookupResponse struct {
	Endpoints  []*Endpoint `protobuf:"bytes,1,rep,name=Endpoints,json=endpoints" json:"Endpoints,omitempty"`
	PublicKeys []string    `protobuf:"bytes,2,rep,name=PublicKeys,json=publicKeys" json:"PublicKeys,omitempty"`
}

func (m *UserLookupResponse) Reset()                    { *m = UserLookupResponse{} }
func (m *UserLookupResponse) String() string            { return proto1.CompactTextString(m) }
func (*UserLookupResponse) ProtoMessage()               {}
func (*UserLookupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *UserLookupResponse) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func init() {
	proto1.RegisterType((*Endpoint)(nil), "proto.Endpoint")
	proto1.RegisterType((*Location)(nil), "proto.Location")
	proto1.RegisterType((*AuthenticateRequest)(nil), "proto.AuthenticateRequest")
	proto1.RegisterType((*Signature)(nil), "proto.Signature")
	proto1.RegisterType((*AuthenticateResponse)(nil), "proto.AuthenticateResponse")
	proto1.RegisterType((*ConfigureRequest)(nil), "proto.ConfigureRequest")
	proto1.RegisterType((*ConfigureResponse)(nil), "proto.ConfigureResponse")
	proto1.RegisterType((*EndpointRequest)(nil), "proto.EndpointRequest")
	proto1.RegisterType((*EndpointResponse)(nil), "proto.EndpointResponse")
	proto1.RegisterType((*ServerUserNameRequest)(nil), "proto.ServerUserNameRequest")
	proto1.RegisterType((*ServerUserNameResponse)(nil), "proto.ServerUserNameResponse")
	proto1.RegisterType((*PingRequest)(nil), "proto.PingRequest")
	proto1.RegisterType((*PingResponse)(nil), "proto.PingResponse")
	proto1.RegisterType((*StoreGetRequest)(nil), "proto.StoreGetRequest")
	proto1.RegisterType((*StoreGetResponse)(nil), "proto.StoreGetResponse")
	proto1.RegisterType((*StorePutRequest)(nil), "proto.StorePutRequest")
	proto1.RegisterType((*StorePutResponse)(nil), "proto.StorePutResponse")
	proto1.RegisterType((*StoreDeleteRequest)(nil), "proto.StoreDeleteRequest")
	proto1.RegisterType((*StoreDeleteResponse)(nil), "proto.StoreDeleteResponse")
	proto1.RegisterType((*UserLookupRequest)(nil), "proto.UserLookupRequest")
	proto1.RegisterType((*UserLookupResponse)(nil), "proto.UserLookupResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Store service

type StoreClient interface {
	// Service methods:
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error)
	Endpoint(ctx context.Context, in *EndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error)
	ServerUserName(ctx context.Context, in *ServerUserNameRequest, opts ...grpc.CallOption) (*ServerUserNameResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Get(ctx context.Context, in *StoreGetRequest, opts ...grpc.CallOption) (*StoreGetResponse, error)
	Put(ctx context.Context, in *StorePutRequest, opts ...grpc.CallOption) (*StorePutResponse, error)
	Delete(ctx context.Context, in *StoreDeleteRequest, opts ...grpc.CallOption) (*StoreDeleteResponse, error)
}

type storeClient struct {
	cc *grpc.ClientConn
}

func NewStoreClient(cc *grpc.ClientConn) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := grpc.Invoke(ctx, "/proto.Store/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error) {
	out := new(ConfigureResponse)
	err := grpc.Invoke(ctx, "/proto.Store/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Endpoint(ctx context.Context, in *EndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error) {
	out := new(EndpointResponse)
	err := grpc.Invoke(ctx, "/proto.Store/Endpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) ServerUserName(ctx context.Context, in *ServerUserNameRequest, opts ...grpc.CallOption) (*ServerUserNameResponse, error) {
	out := new(ServerUserNameResponse)
	err := grpc.Invoke(ctx, "/proto.Store/ServerUserName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/proto.Store/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Get(ctx context.Context, in *StoreGetRequest, opts ...grpc.CallOption) (*StoreGetResponse, error) {
	out := new(StoreGetResponse)
	err := grpc.Invoke(ctx, "/proto.Store/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Put(ctx context.Context, in *StorePutRequest, opts ...grpc.CallOption) (*StorePutResponse, error) {
	out := new(StorePutResponse)
	err := grpc.Invoke(ctx, "/proto.Store/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Delete(ctx context.Context, in *StoreDeleteRequest, opts ...grpc.CallOption) (*StoreDeleteResponse, error) {
	out := new(StoreDeleteResponse)
	err := grpc.Invoke(ctx, "/proto.Store/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Store service

type StoreServer interface {
	// Service methods:
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	Configure(context.Context, *ConfigureRequest) (*ConfigureResponse, error)
	Endpoint(context.Context, *EndpointRequest) (*EndpointResponse, error)
	ServerUserName(context.Context, *ServerUserNameRequest) (*ServerUserNameResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Get(context.Context, *StoreGetRequest) (*StoreGetResponse, error)
	Put(context.Context, *StorePutRequest) (*StorePutResponse, error)
	Delete(context.Context, *StoreDeleteRequest) (*StoreDeleteResponse, error)
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Configure(ctx, req.(*ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Endpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Endpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/Endpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Endpoint(ctx, req.(*EndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_ServerUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerUserNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).ServerUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/ServerUserName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).ServerUserName(ctx, req.(*ServerUserNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Get(ctx, req.(*StoreGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorePutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Put(ctx, req.(*StorePutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Delete(ctx, req.(*StoreDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _Store_Authenticate_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Store_Configure_Handler,
		},
		{
			MethodName: "Endpoint",
			Handler:    _Store_Endpoint_Handler,
		},
		{
			MethodName: "ServerUserName",
			Handler:    _Store_ServerUserName_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Store_Ping_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Store_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Store_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Store_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for User service

type UserClient interface {
	// Service methods:
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error)
	Endpoint(ctx context.Context, in *EndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error)
	ServerUserName(ctx context.Context, in *ServerUserNameRequest, opts ...grpc.CallOption) (*ServerUserNameResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Lookup(ctx context.Context, in *UserLookupRequest, opts ...grpc.CallOption) (*UserLookupResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := grpc.Invoke(ctx, "/proto.User/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error) {
	out := new(ConfigureResponse)
	err := grpc.Invoke(ctx, "/proto.User/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Endpoint(ctx context.Context, in *EndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error) {
	out := new(EndpointResponse)
	err := grpc.Invoke(ctx, "/proto.User/Endpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ServerUserName(ctx context.Context, in *ServerUserNameRequest, opts ...grpc.CallOption) (*ServerUserNameResponse, error) {
	out := new(ServerUserNameResponse)
	err := grpc.Invoke(ctx, "/proto.User/ServerUserName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/proto.User/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Lookup(ctx context.Context, in *UserLookupRequest, opts ...grpc.CallOption) (*UserLookupResponse, error) {
	out := new(UserLookupResponse)
	err := grpc.Invoke(ctx, "/proto.User/Lookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserServer interface {
	// Service methods:
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	Configure(context.Context, *ConfigureRequest) (*ConfigureResponse, error)
	Endpoint(context.Context, *EndpointRequest) (*EndpointResponse, error)
	ServerUserName(context.Context, *ServerUserNameRequest) (*ServerUserNameResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Lookup(context.Context, *UserLookupRequest) (*UserLookupResponse, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.User/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.User/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Configure(ctx, req.(*ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Endpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Endpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.User/Endpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Endpoint(ctx, req.(*EndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ServerUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerUserNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ServerUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.User/ServerUserName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ServerUserName(ctx, req.(*ServerUserNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.User/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.User/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Lookup(ctx, req.(*UserLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _User_Authenticate_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _User_Configure_Handler,
		},
		{
			MethodName: "Endpoint",
			Handler:    _User_Endpoint_Handler,
		},
		{
			MethodName: "ServerUserName",
			Handler:    _User_ServerUserName_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _User_Ping_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _User_Lookup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xec, 0x54, 0xdd, 0x4f, 0x13, 0x41,
	0x10, 0xe7, 0x6c, 0x0b, 0xbd, 0xa1, 0x4a, 0xd9, 0xf2, 0x51, 0x0e, 0x34, 0x64, 0x8d, 0x91, 0x44,
	0x3e, 0xb4, 0xc4, 0xc4, 0x17, 0x83, 0x04, 0x0d, 0x31, 0x12, 0x6d, 0xae, 0xe1, 0x99, 0x1c, 0xed,
	0x52, 0x2f, 0xc0, 0xee, 0xb9, 0xbb, 0xa7, 0xf1, 0x9f, 0xf0, 0xd5, 0x7f, 0xd7, 0xec, 0xdd, 0xec,
	0x7d, 0xf5, 0x24, 0x9a, 0xf8, 0xe8, 0xd3, 0xdd, 0xce, 0xcc, 0xef, 0x37, 0xb3, 0x33, 0xbf, 0x59,
	0xe8, 0xc4, 0x91, 0x8a, 0x42, 0xbe, 0x1f, 0x49, 0xa1, 0x05, 0x69, 0x25, 0x1f, 0x7a, 0x02, 0xed,
	0x77, 0x7c, 0x12, 0x89, 0x90, 0x6b, 0xb2, 0x05, 0xae, 0x96, 0x01, 0x57, 0x91, 0x90, 0xba, 0xef,
	0x6c, 0x3b, 0x3b, 0x2d, 0x3f, 0x37, 0x90, 0x0d, 0x68, 0x73, 0xa6, 0x2f, 0x82, 0xc9, 0x44, 0xf6,
	0xef, 0x6d, 0x3b, 0x3b, 0xae, 0xbf, 0xc0, 0x99, 0x3e, 0x9e, 0x4c, 0x24, 0x3d, 0x87, 0xf6, 0x99,
	0x18, 0x07, 0x3a, 0x14, 0x9c, 0x3c, 0x83, 0x36, 0x43, 0xc2, 0x84, 0x63, 0x71, 0xb0, 0x94, 0x66,
	0xdc, 0xb7, 0x79, 0xfc, 0x2c, 0xc0, 0x64, 0x94, 0xec, 0x8a, 0x49, 0xc6, 0xc7, 0x0c, 0x49, 0x73,
	0x03, 0xd5, 0xd0, 0x3b, 0x8e, 0xf5, 0x67, 0xc6, 0x75, 0x38, 0x0e, 0x34, 0xf3, 0xd9, 0x97, 0x98,
	0x29, 0x4d, 0x36, 0xc1, 0x8d, 0x15, 0x93, 0x17, 0x3c, 0xb8, 0x65, 0x49, 0x0a, 0xd7, 0x6f, 0x1b,
	0xc3, 0xc7, 0xe0, 0x96, 0x91, 0x2e, 0x34, 0xb8, 0xf8, 0x86, 0x5c, 0xe6, 0x97, 0xec, 0x83, 0xab,
	0xc2, 0x29, 0x0f, 0x74, 0x2c, 0x59, 0xbf, 0x91, 0x54, 0xd4, 0xc5, 0x8a, 0x46, 0xd6, 0xee, 0xe7,
	0x21, 0xf4, 0x29, 0xb8, 0x99, 0x9d, 0x74, 0xc0, 0x91, 0x98, 0xc3, 0x91, 0xe6, 0xa4, 0x90, 0xda,
	0x51, 0x74, 0x17, 0x56, 0xca, 0xe5, 0xa9, 0x48, 0x70, 0xc5, 0xc8, 0x0a, 0xb4, 0xb4, 0xb8, 0x66,
	0x1c, 0x71, 0xe9, 0x81, 0xee, 0x42, 0xf7, 0x44, 0xf0, 0xab, 0x70, 0x6a, 0xd2, 0xe1, 0x4d, 0xfa,
	0xb0, 0x20, 0x22, 0xd3, 0x35, 0xd5, 0x77, 0xb6, 0x1b, 0xa6, 0xa3, 0x78, 0xa4, 0x3d, 0x58, 0x2e,
	0x44, 0xa7, 0xc4, 0x74, 0x19, 0x96, 0xb2, 0x1e, 0xa6, 0x0c, 0xf4, 0x08, 0xba, 0xb9, 0x09, 0xf3,
	0xff, 0xcd, 0x04, 0xe8, 0x3a, 0xac, 0x8e, 0x98, 0xfc, 0xca, 0xe4, 0x39, 0x76, 0xd0, 0x32, 0xbf,
	0x84, 0xb5, 0xaa, 0x03, 0xf9, 0xef, 0xea, 0x3f, 0x1d, 0xc0, 0xe2, 0x30, 0xe4, 0x53, 0x7b, 0xc3,
	0xc7, 0x70, 0x3f, 0x0a, 0xf9, 0xf4, 0x42, 0x99, 0xb3, 0x19, 0x72, 0x2a, 0xab, 0x8e, 0x31, 0x8e,
	0xd0, 0x46, 0x0f, 0xa1, 0x93, 0x62, 0x30, 0xc1, 0x1f, 0x81, 0x0e, 0x60, 0x69, 0xa4, 0x85, 0x64,
	0xa7, 0xcc, 0x36, 0xa3, 0xac, 0x26, 0xa7, 0xaa, 0xa6, 0x73, 0xe8, 0xe6, 0x00, 0xcc, 0x44, 0xa0,
	0x39, 0x09, 0x74, 0x90, 0x04, 0x77, 0xfc, 0xe4, 0x9f, 0xec, 0x81, 0x7b, 0x83, 0x62, 0x36, 0xc3,
	0x6e, 0x14, 0xfa, 0x67, 0x45, 0xee, 0xe7, 0x11, 0xf4, 0x09, 0xd6, 0x31, 0x8c, 0xb3, 0x3a, 0x6a,
	0x58, 0xe9, 0x73, 0xcc, 0x9e, 0x84, 0x61, 0xf6, 0xbb, 0xeb, 0x1d, 0x00, 0x49, 0x10, 0x6f, 0xd9,
	0x0d, 0xcb, 0xc5, 0x7f, 0x37, 0x66, 0x15, 0x7a, 0x25, 0x0c, 0x0a, 0xe7, 0x00, 0x96, 0xcd, 0x14,
	0xcf, 0x84, 0xb8, 0x8e, 0x23, 0xcb, 0xe4, 0x41, 0xdb, 0x8e, 0x76, 0x66, 0x8a, 0x63, 0x20, 0x45,
	0x00, 0xd6, 0xbb, 0x07, 0xae, 0x55, 0x50, 0x2a, 0xd8, 0x1a, 0x65, 0xb9, 0x56, 0x59, 0x8a, 0x3c,
	0x02, 0x18, 0xc6, 0x97, 0x37, 0xe1, 0xf8, 0x03, 0xfb, 0x9e, 0x76, 0xd2, 0xf5, 0x21, 0xca, 0x2c,
	0x83, 0x1f, 0x4d, 0x68, 0x25, 0xd5, 0x92, 0xf7, 0xd0, 0x29, 0x6e, 0x12, 0xf1, 0x90, 0xb5, 0x66,
	0xfb, 0xbd, 0xcd, 0x5a, 0x1f, 0x5e, 0x74, 0x8e, 0xbc, 0x01, 0x37, 0x5b, 0x1c, 0xb2, 0x8e, 0xb1,
	0xd5, 0xc5, 0xf3, 0xfa, 0xb3, 0x8e, 0x8c, 0xe1, 0x75, 0xe1, 0x45, 0x5c, 0xab, 0x5e, 0x0f, 0xf1,
	0xeb, 0x33, 0xf6, 0x0c, 0xfe, 0x09, 0x1e, 0x94, 0xf7, 0x86, 0x6c, 0xd9, 0xd7, 0xa6, 0x6e, 0xcf,
	0xbc, 0x87, 0xbf, 0xf1, 0x66, 0x84, 0x2f, 0xa0, 0x69, 0xb6, 0x83, 0x10, 0x0c, 0x2c, 0xac, 0x97,
	0xd7, 0x2b, 0xd9, 0x32, 0xc8, 0x2b, 0x68, 0x9c, 0xb2, 0xbc, 0xfa, 0xca, 0x9e, 0x64, 0xd5, 0x57,
	0xd7, 0x21, 0x45, 0x0e, 0xe3, 0x0a, 0x32, 0x57, 0x76, 0x19, 0x59, 0x90, 0x32, 0x9d, 0x23, 0xc7,
	0x30, 0x9f, 0xaa, 0x8e, 0x6c, 0x14, 0x83, 0x4a, 0xea, 0xf5, 0xbc, 0x3a, 0x97, 0xa5, 0x18, 0xfc,
	0x6c, 0x40, 0xd3, 0x34, 0xe0, 0xbf, 0x1e, 0xfe, 0xb1, 0x1e, 0x8e, 0x60, 0x3e, 0x5d, 0x65, 0x62,
	0x2f, 0x3a, 0xf3, 0x1c, 0x78, 0x1b, 0x35, 0x1e, 0x4b, 0x70, 0x39, 0x9f, 0xf8, 0x0e, 0x7f, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x18, 0x39, 0xe2, 0x4f, 0x43, 0x08, 0x00, 0x00,
}