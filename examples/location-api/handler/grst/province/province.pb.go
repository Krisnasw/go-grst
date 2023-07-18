// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: province.proto

package province

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/krisnasw/go-grst/protoc-gen-cdd/ext/cddapis/cdd/api"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,json=id,proto3" json:"id,omitempty" validate:"required" default:"1"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_province_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_province_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_province_proto_rawDescGZIP(), []int{0}
}

func (x *GetReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Province struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=Id,json=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"name,omitempty"`
}

func (x *Province) Reset() {
	*x = Province{}
	if protoimpl.UnsafeEnabled {
		mi := &file_province_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Province) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Province) ProtoMessage() {}

func (x *Province) ProtoReflect() protoreflect.Message {
	mi := &file_province_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Province.ProtoReflect.Descriptor instead.
func (*Province) Descriptor() ([]byte, []int) {
	return file_province_proto_rawDescGZIP(), []int{1}
}

func (x *Province) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Province) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Provinces struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provinces []*Province `protobuf:"bytes,1,rep,name=Provinces,json=provinces,proto3" json:"provinces,omitempty"`
}

func (x *Provinces) Reset() {
	*x = Provinces{}
	if protoimpl.UnsafeEnabled {
		mi := &file_province_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provinces) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provinces) ProtoMessage() {}

func (x *Provinces) ProtoReflect() protoreflect.Message {
	mi := &file_province_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provinces.ProtoReflect.Descriptor instead.
func (*Provinces) Descriptor() ([]byte, []int) {
	return file_province_proto_rawDescGZIP(), []int{2}
}

func (x *Provinces) GetProvinces() []*Province {
	if x != nil {
		return x.Provinces
	}
	return nil
}

var File_province_proto protoreflect.FileDescriptor

var file_province_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x64, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x64, 0x64, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x11, 0xc2, 0x8a, 0x3b, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x82,
	0xc9, 0x3b, 0x01, 0x31, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x0e, 0x82, 0xcc, 0x3a, 0x0a, 0x8a, 0xcc, 0x3a, 0x02, 0x69, 0x64, 0x90, 0xcc, 0x3a, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x32, 0xc8, 0x01, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x22, 0x45, 0x82, 0xbd, 0x3f, 0x19, 0x88, 0xbd, 0x3f, 0x01, 0x9a, 0xbd, 0x3f, 0x05, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x9a, 0xbd, 0x3f, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x5a, 0x10, 0x22, 0x0e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x2f, 0x7b, 0x49, 0x64, 0x7d, 0x12, 0x0e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x2f, 0x7b, 0x49, 0x64, 0x7d, 0x12, 0x48, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x22,
	0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_province_proto_rawDescOnce sync.Once
	file_province_proto_rawDescData = file_province_proto_rawDesc
)

func file_province_proto_rawDescGZIP() []byte {
	file_province_proto_rawDescOnce.Do(func() {
		file_province_proto_rawDescData = protoimpl.X.CompressGZIP(file_province_proto_rawDescData)
	})
	return file_province_proto_rawDescData
}

var file_province_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_province_proto_goTypes = []interface{}{
	(*GetReq)(nil),      // 0: province.GetReq
	(*Province)(nil),    // 1: province.Province
	(*Provinces)(nil),   // 2: province.Provinces
	(*empty.Empty)(nil), // 3: google.protobuf.Empty
}
var file_province_proto_depIdxs = []int32{
	1, // 0: province.Provinces.Provinces:type_name -> province.Province
	0, // 1: province.province.Get:input_type -> province.GetReq
	3, // 2: province.province.GetAll:input_type -> google.protobuf.Empty
	1, // 3: province.province.Get:output_type -> province.Province
	2, // 4: province.province.GetAll:output_type -> province.Provinces
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_province_proto_init() }
func file_province_proto_init() {
	if File_province_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_province_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
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
		file_province_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Province); i {
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
		file_province_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provinces); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_province_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_province_proto_goTypes,
		DependencyIndexes: file_province_proto_depIdxs,
		MessageInfos:      file_province_proto_msgTypes,
	}.Build()
	File_province_proto = out.File
	file_province_proto_rawDesc = nil
	file_province_proto_goTypes = nil
	file_province_proto_depIdxs = nil
}
