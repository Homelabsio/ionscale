// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: ionscale/v1/ionscale.proto

package ionscalev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_ionscale_v1_ionscale_proto protoreflect.FileDescriptor

var file_ionscale_v1_ionscale_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f,
	0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6f,
	0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x69, 0x6f, 0x6e, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x69,
	0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x69, 0x6c, 0x6e,
	0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x69, 0x6f, 0x6e, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x69, 0x6f, 0x6e, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x69, 0x6f,
	0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x69, 0x6f, 0x6e, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x65, 0x72, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa1, 0x12, 0x0a, 0x0f, 0x49, 0x6f,
	0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x69, 0x6f,
	0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6f,
	0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b,
	0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x22,
	0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4f, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x44, 0x45, 0x52, 0x50, 0x4d, 0x61, 0x70, 0x12, 0x1e, 0x2e, 0x69, 0x6f, 0x6e, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x45, 0x52, 0x50, 0x4d,
	0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6f, 0x6e, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x45, 0x52, 0x50, 0x4d,
	0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a,
	0x53, 0x65, 0x74, 0x44, 0x45, 0x52, 0x50, 0x4d, 0x61, 0x70, 0x12, 0x1e, 0x2e, 0x69, 0x6f, 0x6e,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x45, 0x52, 0x50,
	0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6f, 0x6e,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x45, 0x52, 0x50,
	0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x24, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x5e, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x73, 0x12, 0x23, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x58, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65,
	0x74, 0x12, 0x21, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x69, 0x6f,
	0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61,
	0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69,
	0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x58, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65,
	0x74, 0x12, 0x21, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x20, 0x2e, 0x69, 0x6f, 0x6e,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69,
	0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x4e,
	0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x55, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x20, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49,
	0x41, 0x4d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x20, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x41, 0x4d, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6f, 0x6e,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x41, 0x4d, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x55, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x49, 0x41, 0x4d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x20, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x49, 0x41, 0x4d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x49, 0x41, 0x4d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x20, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a,
	0x0c, 0x53, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x20, 0x2e,
	0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x41,
	0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4b,
	0x65, 0x79, 0x12, 0x1e, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x69, 0x6f, 0x6e, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x58, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79,
	0x12, 0x21, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x20, 0x2e, 0x69, 0x6f, 0x6e, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6f,
	0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x55, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73,
	0x12, 0x20, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x21, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x69, 0x6f,
	0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x6a, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4b,
	0x65, 0x79, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x27, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x4b, 0x65, 0x79, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4b, 0x65, 0x79, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x21,
	0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x69, 0x6f,
	0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x10, 0x53, 0x65,
	0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x24,
	0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3d, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x73, 0x69, 0x65,
	0x62, 0x65, 0x6e, 0x73, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_ionscale_v1_ionscale_proto_goTypes = []interface{}{
	(*GetVersionRequest)(nil),           // 0: ionscale.v1.GetVersionRequest
	(*AuthenticationRequest)(nil),       // 1: ionscale.v1.AuthenticationRequest
	(*GetDERPMapRequest)(nil),           // 2: ionscale.v1.GetDERPMapRequest
	(*SetDERPMapRequest)(nil),           // 3: ionscale.v1.SetDERPMapRequest
	(*CreateAuthMethodRequest)(nil),     // 4: ionscale.v1.CreateAuthMethodRequest
	(*ListAuthMethodsRequest)(nil),      // 5: ionscale.v1.ListAuthMethodsRequest
	(*CreateTailnetRequest)(nil),        // 6: ionscale.v1.CreateTailnetRequest
	(*GetTailnetRequest)(nil),           // 7: ionscale.v1.GetTailnetRequest
	(*ListTailnetRequest)(nil),          // 8: ionscale.v1.ListTailnetRequest
	(*DeleteTailnetRequest)(nil),        // 9: ionscale.v1.DeleteTailnetRequest
	(*GetDNSConfigRequest)(nil),         // 10: ionscale.v1.GetDNSConfigRequest
	(*SetDNSConfigRequest)(nil),         // 11: ionscale.v1.SetDNSConfigRequest
	(*GetIAMPolicyRequest)(nil),         // 12: ionscale.v1.GetIAMPolicyRequest
	(*SetIAMPolicyRequest)(nil),         // 13: ionscale.v1.SetIAMPolicyRequest
	(*GetACLPolicyRequest)(nil),         // 14: ionscale.v1.GetACLPolicyRequest
	(*SetACLPolicyRequest)(nil),         // 15: ionscale.v1.SetACLPolicyRequest
	(*GetAuthKeyRequest)(nil),           // 16: ionscale.v1.GetAuthKeyRequest
	(*CreateAuthKeyRequest)(nil),        // 17: ionscale.v1.CreateAuthKeyRequest
	(*DeleteAuthKeyRequest)(nil),        // 18: ionscale.v1.DeleteAuthKeyRequest
	(*ListAuthKeysRequest)(nil),         // 19: ionscale.v1.ListAuthKeysRequest
	(*ListMachinesRequest)(nil),         // 20: ionscale.v1.ListMachinesRequest
	(*ExpireMachineRequest)(nil),        // 21: ionscale.v1.ExpireMachineRequest
	(*SetMachineKeyExpiryRequest)(nil),  // 22: ionscale.v1.SetMachineKeyExpiryRequest
	(*DeleteMachineRequest)(nil),        // 23: ionscale.v1.DeleteMachineRequest
	(*GetMachineRoutesRequest)(nil),     // 24: ionscale.v1.GetMachineRoutesRequest
	(*SetMachineRoutesRequest)(nil),     // 25: ionscale.v1.SetMachineRoutesRequest
	(*GetVersionResponse)(nil),          // 26: ionscale.v1.GetVersionResponse
	(*AuthenticationResponse)(nil),      // 27: ionscale.v1.AuthenticationResponse
	(*GetDERPMapResponse)(nil),          // 28: ionscale.v1.GetDERPMapResponse
	(*SetDERPMapResponse)(nil),          // 29: ionscale.v1.SetDERPMapResponse
	(*CreateAuthMethodResponse)(nil),    // 30: ionscale.v1.CreateAuthMethodResponse
	(*ListAuthMethodsResponse)(nil),     // 31: ionscale.v1.ListAuthMethodsResponse
	(*CreateTailnetResponse)(nil),       // 32: ionscale.v1.CreateTailnetResponse
	(*GetTailnetResponse)(nil),          // 33: ionscale.v1.GetTailnetResponse
	(*ListTailnetResponse)(nil),         // 34: ionscale.v1.ListTailnetResponse
	(*DeleteTailnetResponse)(nil),       // 35: ionscale.v1.DeleteTailnetResponse
	(*GetDNSConfigResponse)(nil),        // 36: ionscale.v1.GetDNSConfigResponse
	(*SetDNSConfigResponse)(nil),        // 37: ionscale.v1.SetDNSConfigResponse
	(*GetIAMPolicyResponse)(nil),        // 38: ionscale.v1.GetIAMPolicyResponse
	(*SetIAMPolicyResponse)(nil),        // 39: ionscale.v1.SetIAMPolicyResponse
	(*GetACLPolicyResponse)(nil),        // 40: ionscale.v1.GetACLPolicyResponse
	(*SetACLPolicyResponse)(nil),        // 41: ionscale.v1.SetACLPolicyResponse
	(*GetAuthKeyResponse)(nil),          // 42: ionscale.v1.GetAuthKeyResponse
	(*CreateAuthKeyResponse)(nil),       // 43: ionscale.v1.CreateAuthKeyResponse
	(*DeleteAuthKeyResponse)(nil),       // 44: ionscale.v1.DeleteAuthKeyResponse
	(*ListAuthKeysResponse)(nil),        // 45: ionscale.v1.ListAuthKeysResponse
	(*ListMachinesResponse)(nil),        // 46: ionscale.v1.ListMachinesResponse
	(*ExpireMachineResponse)(nil),       // 47: ionscale.v1.ExpireMachineResponse
	(*SetMachineKeyExpiryResponse)(nil), // 48: ionscale.v1.SetMachineKeyExpiryResponse
	(*DeleteMachineResponse)(nil),       // 49: ionscale.v1.DeleteMachineResponse
	(*GetMachineRoutesResponse)(nil),    // 50: ionscale.v1.GetMachineRoutesResponse
}
var file_ionscale_v1_ionscale_proto_depIdxs = []int32{
	0,  // 0: ionscale.v1.IonscaleService.GetVersion:input_type -> ionscale.v1.GetVersionRequest
	1,  // 1: ionscale.v1.IonscaleService.Authenticate:input_type -> ionscale.v1.AuthenticationRequest
	2,  // 2: ionscale.v1.IonscaleService.GetDERPMap:input_type -> ionscale.v1.GetDERPMapRequest
	3,  // 3: ionscale.v1.IonscaleService.SetDERPMap:input_type -> ionscale.v1.SetDERPMapRequest
	4,  // 4: ionscale.v1.IonscaleService.CreateAuthMethod:input_type -> ionscale.v1.CreateAuthMethodRequest
	5,  // 5: ionscale.v1.IonscaleService.ListAuthMethods:input_type -> ionscale.v1.ListAuthMethodsRequest
	6,  // 6: ionscale.v1.IonscaleService.CreateTailnet:input_type -> ionscale.v1.CreateTailnetRequest
	7,  // 7: ionscale.v1.IonscaleService.GetTailnet:input_type -> ionscale.v1.GetTailnetRequest
	8,  // 8: ionscale.v1.IonscaleService.ListTailnets:input_type -> ionscale.v1.ListTailnetRequest
	9,  // 9: ionscale.v1.IonscaleService.DeleteTailnet:input_type -> ionscale.v1.DeleteTailnetRequest
	10, // 10: ionscale.v1.IonscaleService.GetDNSConfig:input_type -> ionscale.v1.GetDNSConfigRequest
	11, // 11: ionscale.v1.IonscaleService.SetDNSConfig:input_type -> ionscale.v1.SetDNSConfigRequest
	12, // 12: ionscale.v1.IonscaleService.GetIAMPolicy:input_type -> ionscale.v1.GetIAMPolicyRequest
	13, // 13: ionscale.v1.IonscaleService.SetIAMPolicy:input_type -> ionscale.v1.SetIAMPolicyRequest
	14, // 14: ionscale.v1.IonscaleService.GetACLPolicy:input_type -> ionscale.v1.GetACLPolicyRequest
	15, // 15: ionscale.v1.IonscaleService.SetACLPolicy:input_type -> ionscale.v1.SetACLPolicyRequest
	16, // 16: ionscale.v1.IonscaleService.GetAuthKey:input_type -> ionscale.v1.GetAuthKeyRequest
	17, // 17: ionscale.v1.IonscaleService.CreateAuthKey:input_type -> ionscale.v1.CreateAuthKeyRequest
	18, // 18: ionscale.v1.IonscaleService.DeleteAuthKey:input_type -> ionscale.v1.DeleteAuthKeyRequest
	19, // 19: ionscale.v1.IonscaleService.ListAuthKeys:input_type -> ionscale.v1.ListAuthKeysRequest
	20, // 20: ionscale.v1.IonscaleService.ListMachines:input_type -> ionscale.v1.ListMachinesRequest
	21, // 21: ionscale.v1.IonscaleService.ExpireMachine:input_type -> ionscale.v1.ExpireMachineRequest
	22, // 22: ionscale.v1.IonscaleService.SetMachineKeyExpiry:input_type -> ionscale.v1.SetMachineKeyExpiryRequest
	23, // 23: ionscale.v1.IonscaleService.DeleteMachine:input_type -> ionscale.v1.DeleteMachineRequest
	24, // 24: ionscale.v1.IonscaleService.GetMachineRoutes:input_type -> ionscale.v1.GetMachineRoutesRequest
	25, // 25: ionscale.v1.IonscaleService.SetMachineRoutes:input_type -> ionscale.v1.SetMachineRoutesRequest
	26, // 26: ionscale.v1.IonscaleService.GetVersion:output_type -> ionscale.v1.GetVersionResponse
	27, // 27: ionscale.v1.IonscaleService.Authenticate:output_type -> ionscale.v1.AuthenticationResponse
	28, // 28: ionscale.v1.IonscaleService.GetDERPMap:output_type -> ionscale.v1.GetDERPMapResponse
	29, // 29: ionscale.v1.IonscaleService.SetDERPMap:output_type -> ionscale.v1.SetDERPMapResponse
	30, // 30: ionscale.v1.IonscaleService.CreateAuthMethod:output_type -> ionscale.v1.CreateAuthMethodResponse
	31, // 31: ionscale.v1.IonscaleService.ListAuthMethods:output_type -> ionscale.v1.ListAuthMethodsResponse
	32, // 32: ionscale.v1.IonscaleService.CreateTailnet:output_type -> ionscale.v1.CreateTailnetResponse
	33, // 33: ionscale.v1.IonscaleService.GetTailnet:output_type -> ionscale.v1.GetTailnetResponse
	34, // 34: ionscale.v1.IonscaleService.ListTailnets:output_type -> ionscale.v1.ListTailnetResponse
	35, // 35: ionscale.v1.IonscaleService.DeleteTailnet:output_type -> ionscale.v1.DeleteTailnetResponse
	36, // 36: ionscale.v1.IonscaleService.GetDNSConfig:output_type -> ionscale.v1.GetDNSConfigResponse
	37, // 37: ionscale.v1.IonscaleService.SetDNSConfig:output_type -> ionscale.v1.SetDNSConfigResponse
	38, // 38: ionscale.v1.IonscaleService.GetIAMPolicy:output_type -> ionscale.v1.GetIAMPolicyResponse
	39, // 39: ionscale.v1.IonscaleService.SetIAMPolicy:output_type -> ionscale.v1.SetIAMPolicyResponse
	40, // 40: ionscale.v1.IonscaleService.GetACLPolicy:output_type -> ionscale.v1.GetACLPolicyResponse
	41, // 41: ionscale.v1.IonscaleService.SetACLPolicy:output_type -> ionscale.v1.SetACLPolicyResponse
	42, // 42: ionscale.v1.IonscaleService.GetAuthKey:output_type -> ionscale.v1.GetAuthKeyResponse
	43, // 43: ionscale.v1.IonscaleService.CreateAuthKey:output_type -> ionscale.v1.CreateAuthKeyResponse
	44, // 44: ionscale.v1.IonscaleService.DeleteAuthKey:output_type -> ionscale.v1.DeleteAuthKeyResponse
	45, // 45: ionscale.v1.IonscaleService.ListAuthKeys:output_type -> ionscale.v1.ListAuthKeysResponse
	46, // 46: ionscale.v1.IonscaleService.ListMachines:output_type -> ionscale.v1.ListMachinesResponse
	47, // 47: ionscale.v1.IonscaleService.ExpireMachine:output_type -> ionscale.v1.ExpireMachineResponse
	48, // 48: ionscale.v1.IonscaleService.SetMachineKeyExpiry:output_type -> ionscale.v1.SetMachineKeyExpiryResponse
	49, // 49: ionscale.v1.IonscaleService.DeleteMachine:output_type -> ionscale.v1.DeleteMachineResponse
	50, // 50: ionscale.v1.IonscaleService.GetMachineRoutes:output_type -> ionscale.v1.GetMachineRoutesResponse
	50, // 51: ionscale.v1.IonscaleService.SetMachineRoutes:output_type -> ionscale.v1.GetMachineRoutesResponse
	26, // [26:52] is the sub-list for method output_type
	0,  // [0:26] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_ionscale_v1_ionscale_proto_init() }
func file_ionscale_v1_ionscale_proto_init() {
	if File_ionscale_v1_ionscale_proto != nil {
		return
	}
	file_ionscale_v1_version_proto_init()
	file_ionscale_v1_auth_proto_init()
	file_ionscale_v1_tailnets_proto_init()
	file_ionscale_v1_auth_methods_proto_init()
	file_ionscale_v1_auth_keys_proto_init()
	file_ionscale_v1_machines_proto_init()
	file_ionscale_v1_routes_proto_init()
	file_ionscale_v1_dns_proto_init()
	file_ionscale_v1_iam_proto_init()
	file_ionscale_v1_acl_proto_init()
	file_ionscale_v1_derp_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ionscale_v1_ionscale_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ionscale_v1_ionscale_proto_goTypes,
		DependencyIndexes: file_ionscale_v1_ionscale_proto_depIdxs,
	}.Build()
	File_ionscale_v1_ionscale_proto = out.File
	file_ionscale_v1_ionscale_proto_rawDesc = nil
	file_ionscale_v1_ionscale_proto_goTypes = nil
	file_ionscale_v1_ionscale_proto_depIdxs = nil
}