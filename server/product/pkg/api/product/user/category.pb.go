// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: category.proto

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

var File_category_proto protoreflect.FileDescriptor

var file_category_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x1a, 0x17, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xba, 0x05, 0x0a, 0x0f, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x86, 0x01, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x85, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x8b, 0x01,
	0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x32, 0x1d, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x88, 0x01, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x26,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x2a, 0x1d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7d, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0xc0, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x0d, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6d, 0x61, 0x71, 0x61,
	0x73, 0x65, 0x2f, 0x73, 0x61, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xa2, 0x02, 0x03, 0x41,
	0x56, 0x43, 0xaa, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0xca, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0xe2, 0x02, 0x1b, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_category_proto_goTypes = []any{
	(*CreateCategoryRequest)(nil),   // 0: api.v1.category.CreateCategoryRequest
	(*GetCategoryByIdRequest)(nil),  // 1: api.v1.category.GetCategoryByIdRequest
	(*UpdateCategoryRequest)(nil),   // 2: api.v1.category.UpdateCategoryRequest
	(*DeleteCategoryRequest)(nil),   // 3: api.v1.category.DeleteCategoryRequest
	(*ListCategoriesRequest)(nil),   // 4: api.v1.category.ListCategoriesRequest
	(*CreateCategoryResponse)(nil),  // 5: api.v1.category.CreateCategoryResponse
	(*GetCategoryByIdResponse)(nil), // 6: api.v1.category.GetCategoryByIdResponse
	(*UpdateCategoryResponse)(nil),  // 7: api.v1.category.UpdateCategoryResponse
	(*DeleteCategoryResponse)(nil),  // 8: api.v1.category.DeleteCategoryResponse
	(*ListCategoriesResponse)(nil),  // 9: api.v1.category.ListCategoriesResponse
}
var file_category_proto_depIdxs = []int32{
	0, // 0: api.v1.category.CategoryService.CreateCategory:input_type -> api.v1.category.CreateCategoryRequest
	1, // 1: api.v1.category.CategoryService.GetCategoryById:input_type -> api.v1.category.GetCategoryByIdRequest
	2, // 2: api.v1.category.CategoryService.UpdateCategory:input_type -> api.v1.category.UpdateCategoryRequest
	3, // 3: api.v1.category.CategoryService.DeleteCategory:input_type -> api.v1.category.DeleteCategoryRequest
	4, // 4: api.v1.category.CategoryService.ListCategories:input_type -> api.v1.category.ListCategoriesRequest
	5, // 5: api.v1.category.CategoryService.CreateCategory:output_type -> api.v1.category.CreateCategoryResponse
	6, // 6: api.v1.category.CategoryService.GetCategoryById:output_type -> api.v1.category.GetCategoryByIdResponse
	7, // 7: api.v1.category.CategoryService.UpdateCategory:output_type -> api.v1.category.UpdateCategoryResponse
	8, // 8: api.v1.category.CategoryService.DeleteCategory:output_type -> api.v1.category.DeleteCategoryResponse
	9, // 9: api.v1.category.CategoryService.ListCategories:output_type -> api.v1.category.ListCategoriesResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_category_proto_init() }
func file_category_proto_init() {
	if File_category_proto != nil {
		return
	}
	file_category_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_category_proto_goTypes,
		DependencyIndexes: file_category_proto_depIdxs,
	}.Build()
	File_category_proto = out.File
	file_category_proto_rawDesc = nil
	file_category_proto_goTypes = nil
	file_category_proto_depIdxs = nil
}
