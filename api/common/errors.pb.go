// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: common/errors.proto

package common

import (
	_ "github.com/go-keg/keg/third_party/response"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// code range 100001 - 109999
type CommonError int32

const (
	CommonError_Success CommonError = 0
	// *
	// 通用 - 基本错误 1000xx
	CommonError_ErrUnknown CommonError = 100001
	CommonError_ErrCommon  CommonError = 100002
	// 请求验证异常
	CommonError_ErrValidate CommonError = 100003
	// token 无效
	CommonError_ErrUnauthorized CommonError = 100100
	// token 无效
	CommonError_ErrTokenInvalid CommonError = 100101
	// token 过期
	CommonError_ErrTokenExpiration CommonError = 100102
	// 已登录，但是没有角色身份
	CommonError_ErrNoPartnerIdentity      CommonError = 100110
	CommonError_ErrNoMerchantIdentity     CommonError = 100111
	CommonError_ErrNoBankerIdentity       CommonError = 100112
	CommonError_ErrNoOrganizationIdentity CommonError = 100113
	// 权限验证不通过
	CommonError_ErrPermissionDenied CommonError = 100120 //
	// 资源不存在
	CommonError_ErrResourceNotFound      CommonError = 100201
	CommonError_ErrQueryConditionIsEmpty CommonError = 100202
)

// Enum value maps for CommonError.
var (
	CommonError_name = map[int32]string{
		0:      "Success",
		100001: "ErrUnknown",
		100002: "ErrCommon",
		100003: "ErrValidate",
		100100: "ErrUnauthorized",
		100101: "ErrTokenInvalid",
		100102: "ErrTokenExpiration",
		100110: "ErrNoPartnerIdentity",
		100111: "ErrNoMerchantIdentity",
		100112: "ErrNoBankerIdentity",
		100113: "ErrNoOrganizationIdentity",
		100120: "ErrPermissionDenied",
		100201: "ErrResourceNotFound",
		100202: "ErrQueryConditionIsEmpty",
	}
	CommonError_value = map[string]int32{
		"Success":                   0,
		"ErrUnknown":                100001,
		"ErrCommon":                 100002,
		"ErrValidate":               100003,
		"ErrUnauthorized":           100100,
		"ErrTokenInvalid":           100101,
		"ErrTokenExpiration":        100102,
		"ErrNoPartnerIdentity":      100110,
		"ErrNoMerchantIdentity":     100111,
		"ErrNoBankerIdentity":       100112,
		"ErrNoOrganizationIdentity": 100113,
		"ErrPermissionDenied":       100120,
		"ErrResourceNotFound":       100201,
		"ErrQueryConditionIsEmpty":  100202,
	}
)

func (x CommonError) Enum() *CommonError {
	p := new(CommonError)
	*p = x
	return p
}

func (x CommonError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommonError) Descriptor() protoreflect.EnumDescriptor {
	return file_common_errors_proto_enumTypes[0].Descriptor()
}

func (CommonError) Type() protoreflect.EnumType {
	return &file_common_errors_proto_enumTypes[0]
}

func (x CommonError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonError.Descriptor instead.
func (CommonError) EnumDescriptor() ([]byte, []int) {
	return file_common_errors_proto_rawDescGZIP(), []int{0}
}

var File_common_errors_proto protoreflect.FileDescriptor

var file_common_errors_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x1a, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xab, 0x06, 0x0a, 0x0b,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x07, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x1a, 0x0e, 0xc8, 0x4b, 0xc8, 0x01, 0xd2, 0x4b,
	0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x0a, 0x45, 0x72, 0x72, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0xa1, 0x8d, 0x06, 0x1a, 0x1c, 0xc8, 0x4b, 0xf4, 0x03,
	0xd2, 0x4b, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x20, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x10, 0xa2, 0x8d, 0x06, 0x1a, 0x1c, 0xc8, 0x4b, 0xf4, 0x03, 0xd2,
	0x4b, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x20, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x10, 0xa3, 0x8d, 0x06, 0x1a, 0x09, 0xc8, 0x4b, 0xc8, 0x01,
	0xd2, 0x4b, 0x02, 0x25, 0x73, 0x12, 0x2a, 0x0a, 0x0f, 0x45, 0x72, 0x72, 0x55, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x10, 0x84, 0x8e, 0x06, 0x1a, 0x13, 0xc8, 0x4b,
	0x91, 0x03, 0xd2, 0x4b, 0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x12, 0x2b, 0x0a, 0x0f, 0x45, 0x72, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x10, 0x85, 0x8e, 0x06, 0x1a, 0x14, 0xc8, 0x4b, 0x91, 0x03, 0xd2, 0x4b,
	0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x31,
	0x0a, 0x12, 0x45, 0x72, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x10, 0x86, 0x8e, 0x06, 0x1a, 0x17, 0xc8, 0x4b, 0xc8, 0x01, 0xd2, 0x4b,
	0x10, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x4d, 0x0a, 0x14, 0x45, 0x72, 0x72, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x10, 0x8e, 0x8e, 0x06, 0x1a, 0x31, 0xc8,
	0x4b, 0xc8, 0x01, 0xd2, 0x4b, 0x2a, 0x54, 0x68, 0x65, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x4f, 0x0a, 0x15, 0x45, 0x72, 0x72, 0x4e, 0x6f, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x10, 0x8f, 0x8e, 0x06, 0x1a, 0x32, 0xc8,
	0x4b, 0xc8, 0x01, 0xd2, 0x4b, 0x2b, 0x54, 0x68, 0x65, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x4b, 0x0a, 0x13, 0x45, 0x72, 0x72, 0x4e, 0x6f, 0x42, 0x61, 0x6e, 0x6b, 0x65, 0x72,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x10, 0x90, 0x8e, 0x06, 0x1a, 0x30, 0xc8, 0x4b,
	0xc8, 0x01, 0xd2, 0x4b, 0x29, 0x54, 0x68, 0x65, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x62,
	0x61, 0x6e, 0x6b, 0x65, 0x72, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x57,
	0x0a, 0x19, 0x45, 0x72, 0x72, 0x4e, 0x6f, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x10, 0x91, 0x8e, 0x06, 0x1a,
	0x36, 0xc8, 0x4b, 0xc8, 0x01, 0xd2, 0x4b, 0x2f, 0x54, 0x68, 0x65, 0x20, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68, 0x61, 0x76,
	0x65, 0x20, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x13, 0x45, 0x72, 0x72, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x10, 0x98,
	0x8e, 0x06, 0x1a, 0x18, 0xc8, 0x4b, 0xc8, 0x01, 0xd2, 0x4b, 0x11, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x44, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x13,
	0x45, 0x72, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x6f, 0x74, 0x46, 0x6f,
	0x75, 0x6e, 0x64, 0x10, 0xe9, 0x8e, 0x06, 0x1a, 0x19, 0xc8, 0x4b, 0x94, 0x03, 0xd2, 0x4b, 0x12,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x4e, 0x6f, 0x74, 0x20, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x12, 0x3f, 0x0a, 0x18, 0x45, 0x72, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x10, 0xea,
	0x8e, 0x06, 0x1a, 0x1f, 0xc8, 0x4b, 0xc8, 0x01, 0xd2, 0x4b, 0x18, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x20, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x49, 0x73, 0x20, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x04, 0xc0, 0x4b, 0xf4, 0x03, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x69, 0x69, 0x78, 0x79, 0x2f, 0x67, 0x6f,
	0x2d, 0x6b, 0x65, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_common_errors_proto_rawDescOnce sync.Once
	file_common_errors_proto_rawDescData = file_common_errors_proto_rawDesc
)

func file_common_errors_proto_rawDescGZIP() []byte {
	file_common_errors_proto_rawDescOnce.Do(func() {
		file_common_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_errors_proto_rawDescData)
	})
	return file_common_errors_proto_rawDescData
}

var file_common_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_errors_proto_goTypes = []interface{}{
	(CommonError)(0), // 0: apis.common.CommonError
}
var file_common_errors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_errors_proto_init() }
func file_common_errors_proto_init() {
	if File_common_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_errors_proto_goTypes,
		DependencyIndexes: file_common_errors_proto_depIdxs,
		EnumInfos:         file_common_errors_proto_enumTypes,
	}.Build()
	File_common_errors_proto = out.File
	file_common_errors_proto_rawDesc = nil
	file_common_errors_proto_goTypes = nil
	file_common_errors_proto_depIdxs = nil
}
