// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: filter_service.proto

package proto

import (
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

var File_filter_service_proto protoreflect.FileDescriptor

var file_filter_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x14, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x57, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_filter_service_proto_goTypes = []interface{}{
	(*GetFilterValuesReq)(nil),  // 0: main.GetFilterValuesReq
	(*GetFilterValuesResp)(nil), // 1: main.GetFilterValuesResp
}
var file_filter_service_proto_depIdxs = []int32{
	0, // 0: main.FilterService.GetFilterValues:input_type -> main.GetFilterValuesReq
	1, // 1: main.FilterService.GetFilterValues:output_type -> main.GetFilterValuesResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_filter_service_proto_init() }
func file_filter_service_proto_init() {
	if File_filter_service_proto != nil {
		return
	}
	file_filter_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_filter_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_filter_service_proto_goTypes,
		DependencyIndexes: file_filter_service_proto_depIdxs,
	}.Build()
	File_filter_service_proto = out.File
	file_filter_service_proto_rawDesc = nil
	file_filter_service_proto_goTypes = nil
	file_filter_service_proto_depIdxs = nil
}
