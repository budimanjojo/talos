// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: security/security.proto

package security

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SecurityService_Certificate_FullMethodName = "/securityapi.SecurityService/Certificate"
)

// SecurityServiceClient is the client API for SecurityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecurityServiceClient interface {
	Certificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type securityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecurityServiceClient(cc grpc.ClientConnInterface) SecurityServiceClient {
	return &securityServiceClient{cc}
}

func (c *securityServiceClient) Certificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, SecurityService_Certificate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityServiceServer is the server API for SecurityService service.
// All implementations must embed UnimplementedSecurityServiceServer
// for forward compatibility
type SecurityServiceServer interface {
	Certificate(context.Context, *CertificateRequest) (*CertificateResponse, error)
	mustEmbedUnimplementedSecurityServiceServer()
}

// UnimplementedSecurityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecurityServiceServer struct {
}

func (UnimplementedSecurityServiceServer) Certificate(context.Context, *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Certificate not implemented")
}
func (UnimplementedSecurityServiceServer) mustEmbedUnimplementedSecurityServiceServer() {}

// UnsafeSecurityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecurityServiceServer will
// result in compilation errors.
type UnsafeSecurityServiceServer interface {
	mustEmbedUnimplementedSecurityServiceServer()
}

func RegisterSecurityServiceServer(s grpc.ServiceRegistrar, srv SecurityServiceServer) {
	s.RegisterService(&SecurityService_ServiceDesc, srv)
}

func _SecurityService_Certificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).Certificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityService_Certificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).Certificate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecurityService_ServiceDesc is the grpc.ServiceDesc for SecurityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecurityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "securityapi.SecurityService",
	HandlerType: (*SecurityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Certificate",
			Handler:    _SecurityService_Certificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "security/security.proto",
}
