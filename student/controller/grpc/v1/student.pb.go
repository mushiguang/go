package v1

import (
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

// ListStudentsRequest defines ListStudents request struct.
type ListStudentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset *int64 `protobuf:"varint,1,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
	Limit  *int64 `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *ListStudentsRequest) Reset() {
	*x = ListStudentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStudentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStudentsRequest) ProtoMessage() {}

func (x *ListStudentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStudentsRequest.ProtoReflect.Descriptor instead.
func (*ListStudentsRequest) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{0}
}

func (x *ListStudentsRequest) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *ListStudentsRequest) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

// StudentInfo contains student details.
type StudentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Nickname    string `protobuf:"bytes,2,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	Password    string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone       string `protobuf:"bytes,5,opt,name=Phone,proto3" json:"Phone,omitempty"`
	IsAdmin     int64  `protobuf:"varint,6,opt,name=IsAdmin,proto3" json:"IsAdmin,omitempty"`
	TotalPolicy string `protobuf:"bytes,7,opt,name=TotalPolicy,proto3" json:"TotalPolicy,omitempty"`
	LoginedAt   string `protobuf:"bytes,8,opt,name=LoginedAt,proto3" json:"LoginedAt,omitempty"`
}

func (x *StudentInfo) Reset() {
	*x = StudentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentInfo) ProtoMessage() {}

func (x *StudentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentInfo.ProtoReflect.Descriptor instead.
func (*StudentInfo) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{1}
}

func (x *StudentInfo) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *StudentInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *StudentInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *StudentInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *StudentInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *StudentInfo) GetIsAdmin() int64 {
	if x != nil {
		return x.IsAdmin
	}
	return 0
}

func (x *StudentInfo) GetTotalPolicy() string {
	if x != nil {
		return x.TotalPolicy
	}
	return ""
}

func (x *StudentInfo) GetLoginedAt() string {
	if x != nil {
		return x.LoginedAt
	}
	return ""
}

// ListStudentsResponse defines ListStudents response struct.
type ListStudentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64       `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Items      []*StudentInfo `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListStudentsResponse) Reset() {
	*x = ListStudentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStudentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStudentsResponse) ProtoMessage() {}

func (x *ListStudentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStudentsResponse.ProtoReflect.Descriptor instead.
func (*ListStudentsResponse) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{2}
}

func (x *ListStudentsResponse) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListStudentsResponse) GetItems() []*StudentInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_student_proto protoreflect.FileDescriptor

var file_student_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0xe0, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x49, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5b, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x32, 0x48, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x09,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x58,
	0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x6d, 0x6f, 0x6e, 0x6b, 0x65, 0x79, 0x2f, 0x67, 0x6f, 0x2f, 0x35, 0x30,
	0x5f, 0x77, 0x65, 0x62, 0x2f, 0x33, 0x30, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x38, 0x30, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_student_proto_rawDescOnce sync.Once
	file_student_proto_rawDescData = file_student_proto_rawDesc
)

func file_student_proto_rawDescGZIP() []byte {
	file_student_proto_rawDescOnce.Do(func() {
		file_student_proto_rawDescData = protoimpl.X.CompressGZIP(file_student_proto_rawDescData)
	})
	return file_student_proto_rawDescData
}

var file_student_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_student_proto_goTypes = []interface{}{
	(*ListStudentsRequest)(nil),  // 0: proto.ListStudentsRequest
	(*StudentInfo)(nil),          // 1: proto.StudentInfo
	(*ListStudentsResponse)(nil), // 2: proto.ListStudentsResponse
}
var file_student_proto_depIdxs = []int32{
	1, // 0: proto.ListStudentsResponse.items:type_name -> proto.StudentInfo
	0, // 1: proto.Student.ListStudents:input_type -> proto.ListStudentsRequest
	2, // 2: proto.Student.ListStudents:output_type -> proto.ListStudentsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_student_proto_init() }
func file_student_proto_init() {
	if File_student_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_student_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStudentsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_student_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_student_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStudentsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_student_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_student_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_student_proto_goTypes,
		DependencyIndexes: file_student_proto_depIdxs,
		MessageInfos:      file_student_proto_msgTypes,
	}.Build()
	File_student_proto = out.File
	file_student_proto_rawDesc = nil
	file_student_proto_goTypes = nil
	file_student_proto_depIdxs = nil
}