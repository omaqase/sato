// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: user.proto

package protobuf

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd7, 0x0e, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x07,
	0x53, 0x65, 0x6e, 0x64, 0x4f, 0x54, 0x50, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x65, 0x6e, 0x64,
	0x2d, 0x6f, 0x74, 0x70, 0x12, 0x6e, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54,
	0x50, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x2d, 0x6f, 0x74, 0x70, 0x12, 0x5e, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x4f, 0x75, 0x74, 0x12, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67,
	0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x4f, 0x75, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22,
	0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x12, 0x77, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1c, 0x12, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x62, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x3a, 0x01, 0x2a, 0x32, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x12, 0xb7, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x64, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x54,
	0x6f, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x64, 0x54, 0x6f, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x29, 0x12, 0x27, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x12, 0x47, 0x0a, 0x08, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x4a, 0x57, 0x54, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x4a, 0x57, 0x54, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x0a, 0x4d, 0x61, 0x6b, 0x65, 0x53, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x6d, 0x61, 0x6b, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x7c, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x54, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x12, 0x22, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x54,
	0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x41, 0x64, 0x64, 0x54, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01,
	0x2a, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x12, 0x88, 0x01, 0x0a, 0x13, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x73, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46,
	0x72, 0x6f, 0x6d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x73, 0x12, 0x73, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x18, 0x12, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x12, 0x5f, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x12, 0x68, 0x0a, 0x09, 0x41, 0x64,
	0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01,
	0x2a, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x63, 0x61, 0x72, 0x74, 0x12, 0x84, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x32, 0x1e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x81, 0x01, 0x0a, 0x0e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x12, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x2a,
	0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x61,
	0x72, 0x74, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x65, 0x0a, 0x09, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x42, 0xa1, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x09, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6d, 0x61, 0x71, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x61, 0x74, 0x6f,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xa2, 0x02,
	0x03, 0x41, 0x56, 0x55, 0xaa, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0xca, 0x02, 0x0b, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x55, 0x73, 0x65, 0x72,
	0xe2, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_user_proto_goTypes = []any{
	(*SendOTPRequest)(nil),                        // 0: api.v1.user.SendOTPRequest
	(*VerifyOTPRequest)(nil),                      // 1: api.v1.user.VerifyOTPRequest
	(*LogOutRequest)(nil),                         // 2: api.v1.user.LogOutRequest
	(*RefreshTokenRequest)(nil),                   // 3: api.v1.user.RefreshTokenRequest
	(*UpdateUserRequest)(nil),                     // 4: api.v1.user.UpdateUserRequest
	(*GetSubscribedToPromotionUsersRequest)(nil),  // 5: api.v1.user.GetSubscribedToPromotionUsersRequest
	(*ParseJWTRequest)(nil),                       // 6: api.v1.user.ParseJWTRequest
	(*MakeSellerRequest)(nil),                     // 7: api.v1.user.MakeSellerRequest
	(*AddToFavoritesRequest)(nil),                 // 8: api.v1.user.AddToFavoritesRequest
	(*RemoveFromFavoritesRequest)(nil),            // 9: api.v1.user.RemoveFromFavoritesRequest
	(*GetFavoritesRequest)(nil),                   // 10: api.v1.user.GetFavoritesRequest
	(*GetCartRequest)(nil),                        // 11: api.v1.user.GetCartRequest
	(*AddToCartRequest)(nil),                      // 12: api.v1.user.AddToCartRequest
	(*UpdateCartItemRequest)(nil),                 // 13: api.v1.user.UpdateCartItemRequest
	(*RemoveFromCartRequest)(nil),                 // 14: api.v1.user.RemoveFromCartRequest
	(*ClearCartRequest)(nil),                      // 15: api.v1.user.ClearCartRequest
	(*SendOTPResponse)(nil),                       // 16: api.v1.user.SendOTPResponse
	(*VerifyOTPResponse)(nil),                     // 17: api.v1.user.VerifyOTPResponse
	(*LogOutResponse)(nil),                        // 18: api.v1.user.LogOutResponse
	(*RefreshTokenResponse)(nil),                  // 19: api.v1.user.RefreshTokenResponse
	(*UpdateUserResponse)(nil),                    // 20: api.v1.user.UpdateUserResponse
	(*GetSubscribedToPromotionUsersResponse)(nil), // 21: api.v1.user.GetSubscribedToPromotionUsersResponse
	(*ParseJWTResponse)(nil),                      // 22: api.v1.user.ParseJWTResponse
	(*MakeSellerResponse)(nil),                    // 23: api.v1.user.MakeSellerResponse
	(*AddToFavoritesResponse)(nil),                // 24: api.v1.user.AddToFavoritesResponse
	(*RemoveFromFavoritesResponse)(nil),           // 25: api.v1.user.RemoveFromFavoritesResponse
	(*GetFavoritesResponse)(nil),                  // 26: api.v1.user.GetFavoritesResponse
	(*GetCartResponse)(nil),                       // 27: api.v1.user.GetCartResponse
	(*AddToCartResponse)(nil),                     // 28: api.v1.user.AddToCartResponse
	(*UpdateCartItemResponse)(nil),                // 29: api.v1.user.UpdateCartItemResponse
	(*RemoveFromCartResponse)(nil),                // 30: api.v1.user.RemoveFromCartResponse
	(*ClearCartResponse)(nil),                     // 31: api.v1.user.ClearCartResponse
}
var file_user_proto_depIdxs = []int32{
	0,  // 0: api.v1.user.UserService.SendOTP:input_type -> api.v1.user.SendOTPRequest
	1,  // 1: api.v1.user.UserService.VerifyOTP:input_type -> api.v1.user.VerifyOTPRequest
	2,  // 2: api.v1.user.UserService.LogOut:input_type -> api.v1.user.LogOutRequest
	3,  // 3: api.v1.user.UserService.RefreshToken:input_type -> api.v1.user.RefreshTokenRequest
	4,  // 4: api.v1.user.UserService.Update:input_type -> api.v1.user.UpdateUserRequest
	5,  // 5: api.v1.user.UserService.GetSubscribedToPromotionUsers:input_type -> api.v1.user.GetSubscribedToPromotionUsersRequest
	6,  // 6: api.v1.user.UserService.ParseJWT:input_type -> api.v1.user.ParseJWTRequest
	7,  // 7: api.v1.user.UserService.MakeSeller:input_type -> api.v1.user.MakeSellerRequest
	8,  // 8: api.v1.user.UserService.AddToFavorites:input_type -> api.v1.user.AddToFavoritesRequest
	9,  // 9: api.v1.user.UserService.RemoveFromFavorites:input_type -> api.v1.user.RemoveFromFavoritesRequest
	10, // 10: api.v1.user.UserService.GetFavorites:input_type -> api.v1.user.GetFavoritesRequest
	11, // 11: api.v1.user.UserService.GetCart:input_type -> api.v1.user.GetCartRequest
	12, // 12: api.v1.user.UserService.AddToCart:input_type -> api.v1.user.AddToCartRequest
	13, // 13: api.v1.user.UserService.UpdateCartItem:input_type -> api.v1.user.UpdateCartItemRequest
	14, // 14: api.v1.user.UserService.RemoveFromCart:input_type -> api.v1.user.RemoveFromCartRequest
	15, // 15: api.v1.user.UserService.ClearCart:input_type -> api.v1.user.ClearCartRequest
	16, // 16: api.v1.user.UserService.SendOTP:output_type -> api.v1.user.SendOTPResponse
	17, // 17: api.v1.user.UserService.VerifyOTP:output_type -> api.v1.user.VerifyOTPResponse
	18, // 18: api.v1.user.UserService.LogOut:output_type -> api.v1.user.LogOutResponse
	19, // 19: api.v1.user.UserService.RefreshToken:output_type -> api.v1.user.RefreshTokenResponse
	20, // 20: api.v1.user.UserService.Update:output_type -> api.v1.user.UpdateUserResponse
	21, // 21: api.v1.user.UserService.GetSubscribedToPromotionUsers:output_type -> api.v1.user.GetSubscribedToPromotionUsersResponse
	22, // 22: api.v1.user.UserService.ParseJWT:output_type -> api.v1.user.ParseJWTResponse
	23, // 23: api.v1.user.UserService.MakeSeller:output_type -> api.v1.user.MakeSellerResponse
	24, // 24: api.v1.user.UserService.AddToFavorites:output_type -> api.v1.user.AddToFavoritesResponse
	25, // 25: api.v1.user.UserService.RemoveFromFavorites:output_type -> api.v1.user.RemoveFromFavoritesResponse
	26, // 26: api.v1.user.UserService.GetFavorites:output_type -> api.v1.user.GetFavoritesResponse
	27, // 27: api.v1.user.UserService.GetCart:output_type -> api.v1.user.GetCartResponse
	28, // 28: api.v1.user.UserService.AddToCart:output_type -> api.v1.user.AddToCartResponse
	29, // 29: api.v1.user.UserService.UpdateCartItem:output_type -> api.v1.user.UpdateCartItemResponse
	30, // 30: api.v1.user.UserService.RemoveFromCart:output_type -> api.v1.user.RemoveFromCartResponse
	31, // 31: api.v1.user.UserService.ClearCart:output_type -> api.v1.user.ClearCartResponse
	16, // [16:32] is the sub-list for method output_type
	0,  // [0:16] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_user_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
