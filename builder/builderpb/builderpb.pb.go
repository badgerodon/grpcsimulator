// Code generated by protoc-gen-go. DO NOT EDIT.
// source: builderpb.proto

/*
Package builderpb is a generated protocol buffer package.

It is generated from these files:
	builderpb.proto

It has these top-level messages:
	BuildRequest
	BuildResponse
*/
package builderpb

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

type BuildRequest struct {
	ImportPath string `protobuf:"bytes,1,opt,name=import_path,json=importPath" json:"import_path,omitempty"`
	Branch     string `protobuf:"bytes,2,opt,name=branch" json:"branch,omitempty"`
}

func (m *BuildRequest) Reset()                    { *m = BuildRequest{} }
func (m *BuildRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildRequest) ProtoMessage()               {}
func (*BuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BuildRequest) GetImportPath() string {
	if m != nil {
		return m.ImportPath
	}
	return ""
}

func (m *BuildRequest) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

type BuildResponse struct {
	Location string `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
}

func (m *BuildResponse) Reset()                    { *m = BuildResponse{} }
func (m *BuildResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildResponse) ProtoMessage()               {}
func (*BuildResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BuildResponse) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func init() {
	proto.RegisterType((*BuildRequest)(nil), "BuildRequest")
	proto.RegisterType((*BuildResponse)(nil), "BuildResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Service service

type ServiceClient interface {
	Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error) {
	out := new(BuildResponse)
	err := grpc.Invoke(ctx, "/Service/Build", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceServer interface {
	Build(context.Context, *BuildRequest) (*BuildResponse, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Service/Build",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Build(ctx, req.(*BuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Build",
			Handler:    _Service_Build_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "builderpb.proto",
}

func init() { proto.RegisterFile("builderpb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2a, 0xcd, 0xcc,
	0x49, 0x49, 0x2d, 0x2a, 0x48, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x72, 0xe7, 0xe2, 0x71,
	0x02, 0x09, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xc9, 0x73, 0x71, 0x67, 0xe6, 0x16,
	0xe4, 0x17, 0x95, 0xc4, 0x17, 0x24, 0x96, 0x64, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x71,
	0x41, 0x84, 0x02, 0x12, 0x4b, 0x32, 0x84, 0xc4, 0xb8, 0xd8, 0x92, 0x8a, 0x12, 0xf3, 0x92, 0x33,
	0x24, 0x98, 0xc0, 0x72, 0x50, 0x9e, 0x92, 0x36, 0x17, 0x2f, 0xd4, 0xa0, 0xe2, 0x82, 0xfc, 0xbc,
	0xe2, 0x54, 0x21, 0x29, 0x2e, 0x8e, 0x9c, 0xfc, 0xe4, 0xc4, 0x92, 0xcc, 0xfc, 0x3c, 0xa8, 0x31,
	0x70, 0xbe, 0x91, 0x21, 0x17, 0x7b, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x1a, 0x17,
	0x2b, 0x58, 0x9f, 0x10, 0xaf, 0x1e, 0xb2, 0x43, 0xa4, 0xf8, 0xf4, 0x50, 0x8c, 0x4b, 0x62, 0x03,
	0xbb, 0xd7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xc5, 0xf1, 0x3e, 0xc2, 0x00, 0x00, 0x00,
}
