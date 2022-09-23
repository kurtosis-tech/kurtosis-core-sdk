// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: executable_module_service.proto

package kurtosis_core_rpc_api_bindings

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ExecutableModuleServiceClient is the client API for ExecutableModuleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecutableModuleServiceClient interface {
	// Returns true if the executable module is available
	IsAvailable(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Runs the module's execute function
	Execute(ctx context.Context, in *ExecuteArgs, opts ...grpc.CallOption) (*ExecuteResponse, error)
}

type executableModuleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutableModuleServiceClient(cc grpc.ClientConnInterface) ExecutableModuleServiceClient {
	return &executableModuleServiceClient{cc}
}

func (c *executableModuleServiceClient) IsAvailable(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/module_api.ExecutableModuleService/IsAvailable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executableModuleServiceClient) Execute(ctx context.Context, in *ExecuteArgs, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, "/module_api.ExecutableModuleService/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutableModuleServiceServer is the server API for ExecutableModuleService service.
// All implementations must embed UnimplementedExecutableModuleServiceServer
// for forward compatibility
type ExecutableModuleServiceServer interface {
	// Returns true if the executable module is available
	IsAvailable(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// Runs the module's execute function
	Execute(context.Context, *ExecuteArgs) (*ExecuteResponse, error)
	mustEmbedUnimplementedExecutableModuleServiceServer()
}

// UnimplementedExecutableModuleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExecutableModuleServiceServer struct {
}

func (UnimplementedExecutableModuleServiceServer) IsAvailable(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAvailable not implemented")
}
func (UnimplementedExecutableModuleServiceServer) Execute(context.Context, *ExecuteArgs) (*ExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedExecutableModuleServiceServer) mustEmbedUnimplementedExecutableModuleServiceServer() {
}

// UnsafeExecutableModuleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecutableModuleServiceServer will
// result in compilation errors.
type UnsafeExecutableModuleServiceServer interface {
	mustEmbedUnimplementedExecutableModuleServiceServer()
}

func RegisterExecutableModuleServiceServer(s grpc.ServiceRegistrar, srv ExecutableModuleServiceServer) {
	s.RegisterService(&ExecutableModuleService_ServiceDesc, srv)
}

func _ExecutableModuleService_IsAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutableModuleServiceServer).IsAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/module_api.ExecutableModuleService/IsAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutableModuleServiceServer).IsAvailable(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExecutableModuleService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutableModuleServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/module_api.ExecutableModuleService/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutableModuleServiceServer).Execute(ctx, req.(*ExecuteArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// ExecutableModuleService_ServiceDesc is the grpc.ServiceDesc for ExecutableModuleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExecutableModuleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "module_api.ExecutableModuleService",
	HandlerType: (*ExecutableModuleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAvailable",
			Handler:    _ExecutableModuleService_IsAvailable_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _ExecutableModuleService_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "executable_module_service.proto",
}