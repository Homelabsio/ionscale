// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/ionscale.proto

package api

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

var File_api_ionscale_proto protoreflect.FileDescriptor

var file_api_ionscale_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x61, 0x70, 0x69, 0x2f,
	0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe2, 0x06, 0x0a, 0x08, 0x49, 0x6f, 0x6e,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x73,
	0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x69, 0x6c, 0x6e,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x69, 0x6c, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c,
	0x53, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74,
	0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x53, 0x65,
	0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x43,
	0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b,
	0x65, 0x79, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65,
	0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x73, 0x69, 0x65,
	0x62, 0x65, 0x6e, 0x73, 0x2f, 0x69, 0x6f, 0x6e, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_api_ionscale_proto_goTypes = []interface{}{
	(*GetVersionRequest)(nil),     // 0: api.GetVersionRequest
	(*CreateTailnetRequest)(nil),  // 1: api.CreateTailnetRequest
	(*ListTailnetRequest)(nil),    // 2: api.ListTailnetRequest
	(*GetDNSConfigRequest)(nil),   // 3: api.GetDNSConfigRequest
	(*SetDNSConfigRequest)(nil),   // 4: api.SetDNSConfigRequest
	(*GetACLPolicyRequest)(nil),   // 5: api.GetACLPolicyRequest
	(*SetACLPolicyRequest)(nil),   // 6: api.SetACLPolicyRequest
	(*CreateAuthKeyRequest)(nil),  // 7: api.CreateAuthKeyRequest
	(*DeleteAuthKeyRequest)(nil),  // 8: api.DeleteAuthKeyRequest
	(*ListAuthKeysRequest)(nil),   // 9: api.ListAuthKeysRequest
	(*ListMachinesRequest)(nil),   // 10: api.ListMachinesRequest
	(*DeleteMachineRequest)(nil),  // 11: api.DeleteMachineRequest
	(*GetVersionResponse)(nil),    // 12: api.GetVersionResponse
	(*CreateTailnetResponse)(nil), // 13: api.CreateTailnetResponse
	(*ListTailnetResponse)(nil),   // 14: api.ListTailnetResponse
	(*GetDNSConfigResponse)(nil),  // 15: api.GetDNSConfigResponse
	(*SetDNSConfigResponse)(nil),  // 16: api.SetDNSConfigResponse
	(*GetACLPolicyResponse)(nil),  // 17: api.GetACLPolicyResponse
	(*SetACLPolicyResponse)(nil),  // 18: api.SetACLPolicyResponse
	(*CreateAuthKeyResponse)(nil), // 19: api.CreateAuthKeyResponse
	(*DeleteAuthKeyResponse)(nil), // 20: api.DeleteAuthKeyResponse
	(*ListAuthKeysResponse)(nil),  // 21: api.ListAuthKeysResponse
	(*ListMachinesResponse)(nil),  // 22: api.ListMachinesResponse
	(*DeleteMachineResponse)(nil), // 23: api.DeleteMachineResponse
}
var file_api_ionscale_proto_depIdxs = []int32{
	0,  // 0: api.Ionscale.GetVersion:input_type -> api.GetVersionRequest
	1,  // 1: api.Ionscale.CreateTailnet:input_type -> api.CreateTailnetRequest
	2,  // 2: api.Ionscale.ListTailnets:input_type -> api.ListTailnetRequest
	3,  // 3: api.Ionscale.GetDNSConfig:input_type -> api.GetDNSConfigRequest
	4,  // 4: api.Ionscale.SetDNSConfig:input_type -> api.SetDNSConfigRequest
	5,  // 5: api.Ionscale.GetACLPolicy:input_type -> api.GetACLPolicyRequest
	6,  // 6: api.Ionscale.SetACLPolicy:input_type -> api.SetACLPolicyRequest
	7,  // 7: api.Ionscale.CreateAuthKey:input_type -> api.CreateAuthKeyRequest
	8,  // 8: api.Ionscale.DeleteAuthKey:input_type -> api.DeleteAuthKeyRequest
	9,  // 9: api.Ionscale.ListAuthKeys:input_type -> api.ListAuthKeysRequest
	10, // 10: api.Ionscale.ListMachines:input_type -> api.ListMachinesRequest
	11, // 11: api.Ionscale.DeleteMachine:input_type -> api.DeleteMachineRequest
	12, // 12: api.Ionscale.GetVersion:output_type -> api.GetVersionResponse
	13, // 13: api.Ionscale.CreateTailnet:output_type -> api.CreateTailnetResponse
	14, // 14: api.Ionscale.ListTailnets:output_type -> api.ListTailnetResponse
	15, // 15: api.Ionscale.GetDNSConfig:output_type -> api.GetDNSConfigResponse
	16, // 16: api.Ionscale.SetDNSConfig:output_type -> api.SetDNSConfigResponse
	17, // 17: api.Ionscale.GetACLPolicy:output_type -> api.GetACLPolicyResponse
	18, // 18: api.Ionscale.SetACLPolicy:output_type -> api.SetACLPolicyResponse
	19, // 19: api.Ionscale.CreateAuthKey:output_type -> api.CreateAuthKeyResponse
	20, // 20: api.Ionscale.DeleteAuthKey:output_type -> api.DeleteAuthKeyResponse
	21, // 21: api.Ionscale.ListAuthKeys:output_type -> api.ListAuthKeysResponse
	22, // 22: api.Ionscale.ListMachines:output_type -> api.ListMachinesResponse
	23, // 23: api.Ionscale.DeleteMachine:output_type -> api.DeleteMachineResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_ionscale_proto_init() }
func file_api_ionscale_proto_init() {
	if File_api_ionscale_proto != nil {
		return
	}
	file_api_version_proto_init()
	file_api_tailnets_proto_init()
	file_api_auth_keys_proto_init()
	file_api_machines_proto_init()
	file_api_dns_proto_init()
	file_api_acl_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_ionscale_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ionscale_proto_goTypes,
		DependencyIndexes: file_api_ionscale_proto_depIdxs,
	}.Build()
	File_api_ionscale_proto = out.File
	file_api_ionscale_proto_rawDesc = nil
	file_api_ionscale_proto_goTypes = nil
	file_api_ionscale_proto_depIdxs = nil
}
