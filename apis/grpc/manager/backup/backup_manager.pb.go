//
// Copyright (C) 2019 kpango (Yusuke Kato)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package backup

import (
	context "context"
	fmt "fmt"
	_ "github.com/danielvladco/go-proto-gql/pb"
	proto "github.com/gogo/protobuf/proto"
	payload "github.com/vdaas/vald/apis/grpc/payload"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("backup/backup_manager.proto", fileDescriptor_d3d7e5699810d1ca) }

var fileDescriptor_d3d7e5699810d1ca = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0x29, 0xdf, 0x17, 0x22, 0x23, 0x90, 0x30, 0x6e, 0x4c, 0x51, 0x48, 0xba, 0x64, 0x31,
	0x93, 0x28, 0x2b, 0x13, 0x5d, 0x10, 0xff, 0x84, 0x44, 0x22, 0x61, 0xc1, 0xc2, 0x8d, 0x99, 0xb6,
	0xe3, 0x58, 0x6d, 0x7b, 0xc7, 0xe9, 0x40, 0x42, 0x8c, 0x1b, 0x5f, 0x81, 0x17, 0xe1, 0x31, 0x5c,
	0x9a, 0xf8, 0x02, 0x84, 0xf8, 0x20, 0x86, 0x4e, 0x21, 0x51, 0x90, 0x55, 0xd3, 0x73, 0xe6, 0xfc,
	0xee, 0x9d, 0xcc, 0x41, 0x35, 0x97, 0x79, 0x4f, 0x43, 0x49, 0xcd, 0xe7, 0x2e, 0x62, 0x31, 0x13,
	0x5c, 0x11, 0xa9, 0x40, 0x03, 0xae, 0xfc, 0x54, 0xed, 0xb2, 0x64, 0xe3, 0x10, 0x98, 0x6f, 0x6c,
	0xfb, 0x40, 0x00, 0x88, 0x90, 0x53, 0x26, 0x03, 0xca, 0xe2, 0x18, 0x34, 0xd3, 0x01, 0xc4, 0x49,
	0xe6, 0x96, 0xa4, 0x4b, 0xc5, 0x73, 0x68, 0xfe, 0x8e, 0x26, 0xff, 0x50, 0xa1, 0x9d, 0xd2, 0xf0,
	0x29, 0x2a, 0x5e, 0x71, 0x3d, 0xe0, 0x9e, 0x06, 0x85, 0x31, 0x59, 0x32, 0x6f, 0xdc, 0x47, 0xee,
	0x69, 0xd2, 0x39, 0xb7, 0xed, 0xdf, 0x5a, 0x97, 0x6b, 0x66, 0xce, 0x3b, 0x39, 0xdc, 0x42, 0xc5,
	0x6b, 0xf0, 0xcc, 0xa8, 0x8d, 0xf1, 0xea, 0x4a, 0xeb, 0xc4, 0xf7, 0x40, 0x3a, 0xbd, 0xc4, 0xc9,
	0xe1, 0x1e, 0xda, 0xe9, 0x73, 0x11, 0x24, 0x9a, 0x2b, 0xbc, 0x85, 0x6f, 0x57, 0x56, 0xde, 0x45,
	0x24, 0xf5, 0xd8, 0xd9, 0x9f, 0xce, 0x1a, 0xd6, 0xdb, 0xe7, 0xd7, 0x24, 0x5f, 0x71, 0x8a, 0x54,
	0x65, 0x88, 0x13, 0xab, 0x89, 0xcf, 0x50, 0x79, 0x49, 0xec, 0x0e, 0x43, 0x1d, 0xe0, 0xda, 0xdf,
	0xd8, 0x64, 0x8d, 0x9b, 0xc3, 0x97, 0xa8, 0xd0, 0xe7, 0x11, 0x8c, 0xf8, 0xc6, 0x4b, 0x6c, 0xd9,
	0xa3, 0x59, 0xa2, 0x2a, 0x0d, 0xd2, 0x97, 0xc0, 0x7f, 0xc5, 0x2d, 0xb4, 0x6b, 0x38, 0x66, 0x8b,
	0xbd, 0x75, 0xd8, 0x86, 0xe9, 0xf6, 0xff, 0xe9, 0xac, 0x91, 0x6f, 0xcb, 0xf7, 0x79, 0xdd, 0xfa,
	0x98, 0xd7, 0xad, 0xd9, 0xbc, 0x6e, 0xa1, 0x43, 0x50, 0x82, 0x8c, 0x7c, 0xc6, 0x12, 0x32, 0x62,
	0xa1, 0x4f, 0x96, 0x65, 0x30, 0x2d, 0x68, 0x57, 0x07, 0x2c, 0xf4, 0xcd, 0x1b, 0x76, 0x8d, 0xd3,
	0xb3, 0x6e, 0x89, 0x08, 0xf4, 0xc3, 0xd0, 0x25, 0x1e, 0x44, 0x34, 0x8d, 0xd2, 0x45, 0x74, 0x51,
	0x89, 0x84, 0x0a, 0x25, 0x3d, 0x9a, 0x41, 0xb2, 0x82, 0xb9, 0x85, 0xb4, 0x0e, 0xc7, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x42, 0xd9, 0x71, 0x45, 0x78, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BackupClient is the client API for Backup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackupClient interface {
	GetVector(ctx context.Context, in *payload.Object_ID, opts ...grpc.CallOption) (*payload.Object_MetaVector, error)
	Locations(ctx context.Context, in *payload.Object_ID, opts ...grpc.CallOption) (*payload.Info_IPs, error)
	Register(ctx context.Context, in *payload.Object_MetaVector, opts ...grpc.CallOption) (*payload.Empty, error)
	RegisterMulti(ctx context.Context, in *payload.Object_MetaVectors, opts ...grpc.CallOption) (*payload.Empty, error)
	Remove(ctx context.Context, in *payload.Object_ID, opts ...grpc.CallOption) (*payload.Empty, error)
	RemoveMulti(ctx context.Context, in *payload.Object_IDs, opts ...grpc.CallOption) (*payload.Empty, error)
}

type backupClient struct {
	cc *grpc.ClientConn
}

func NewBackupClient(cc *grpc.ClientConn) BackupClient {
	return &backupClient{cc}
}

func (c *backupClient) GetVector(ctx context.Context, in *payload.Object_ID, opts ...grpc.CallOption) (*payload.Object_MetaVector, error) {
	out := new(payload.Object_MetaVector)
	err := c.cc.Invoke(ctx, "/backup_manager.Backup/GetVector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupClient) Locations(ctx context.Context, in *payload.Object_ID, opts ...grpc.CallOption) (*payload.Info_IPs, error) {
	out := new(payload.Info_IPs)
	err := c.cc.Invoke(ctx, "/backup_manager.Backup/Locations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupClient) Register(ctx context.Context, in *payload.Object_MetaVector, opts ...grpc.CallOption) (*payload.Empty, error) {
	out := new(payload.Empty)
	err := c.cc.Invoke(ctx, "/backup_manager.Backup/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupClient) RegisterMulti(ctx context.Context, in *payload.Object_MetaVectors, opts ...grpc.CallOption) (*payload.Empty, error) {
	out := new(payload.Empty)
	err := c.cc.Invoke(ctx, "/backup_manager.Backup/RegisterMulti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupClient) Remove(ctx context.Context, in *payload.Object_ID, opts ...grpc.CallOption) (*payload.Empty, error) {
	out := new(payload.Empty)
	err := c.cc.Invoke(ctx, "/backup_manager.Backup/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupClient) RemoveMulti(ctx context.Context, in *payload.Object_IDs, opts ...grpc.CallOption) (*payload.Empty, error) {
	out := new(payload.Empty)
	err := c.cc.Invoke(ctx, "/backup_manager.Backup/RemoveMulti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupServer is the server API for Backup service.
type BackupServer interface {
	GetVector(context.Context, *payload.Object_ID) (*payload.Object_MetaVector, error)
	Locations(context.Context, *payload.Object_ID) (*payload.Info_IPs, error)
	Register(context.Context, *payload.Object_MetaVector) (*payload.Empty, error)
	RegisterMulti(context.Context, *payload.Object_MetaVectors) (*payload.Empty, error)
	Remove(context.Context, *payload.Object_ID) (*payload.Empty, error)
	RemoveMulti(context.Context, *payload.Object_IDs) (*payload.Empty, error)
}

// UnimplementedBackupServer can be embedded to have forward compatible implementations.
type UnimplementedBackupServer struct {
}

func (*UnimplementedBackupServer) GetVector(ctx context.Context, req *payload.Object_ID) (*payload.Object_MetaVector, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVector not implemented")
}
func (*UnimplementedBackupServer) Locations(ctx context.Context, req *payload.Object_ID) (*payload.Info_IPs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Locations not implemented")
}
func (*UnimplementedBackupServer) Register(ctx context.Context, req *payload.Object_MetaVector) (*payload.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedBackupServer) RegisterMulti(ctx context.Context, req *payload.Object_MetaVectors) (*payload.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterMulti not implemented")
}
func (*UnimplementedBackupServer) Remove(ctx context.Context, req *payload.Object_ID) (*payload.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (*UnimplementedBackupServer) RemoveMulti(ctx context.Context, req *payload.Object_IDs) (*payload.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMulti not implemented")
}

func RegisterBackupServer(s *grpc.Server, srv BackupServer) {
	s.RegisterService(&_Backup_serviceDesc, srv)
}

func _Backup_GetVector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).GetVector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup_manager.Backup/GetVector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).GetVector(ctx, req.(*payload.Object_ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backup_Locations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).Locations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup_manager.Backup/Locations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).Locations(ctx, req.(*payload.Object_ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backup_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_MetaVector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup_manager.Backup/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).Register(ctx, req.(*payload.Object_MetaVector))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backup_RegisterMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_MetaVectors)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).RegisterMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup_manager.Backup/RegisterMulti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).RegisterMulti(ctx, req.(*payload.Object_MetaVectors))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backup_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup_manager.Backup/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).Remove(ctx, req.(*payload.Object_ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backup_RemoveMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payload.Object_IDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).RemoveMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup_manager.Backup/RemoveMulti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).RemoveMulti(ctx, req.(*payload.Object_IDs))
	}
	return interceptor(ctx, in, info, handler)
}

var _Backup_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backup_manager.Backup",
	HandlerType: (*BackupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVector",
			Handler:    _Backup_GetVector_Handler,
		},
		{
			MethodName: "Locations",
			Handler:    _Backup_Locations_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Backup_Register_Handler,
		},
		{
			MethodName: "RegisterMulti",
			Handler:    _Backup_RegisterMulti_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Backup_Remove_Handler,
		},
		{
			MethodName: "RemoveMulti",
			Handler:    _Backup_RemoveMulti_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backup/backup_manager.proto",
}
