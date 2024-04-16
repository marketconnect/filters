// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: filter_message.proto

package proto

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

// filter values
type GetFilterValuesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilterName string `protobuf:"bytes,1,opt,name=filterName,proto3" json:"filterName,omitempty"`
}

func (x *GetFilterValuesReq) Reset() {
	*x = GetFilterValuesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilterValuesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilterValuesReq) ProtoMessage() {}

func (x *GetFilterValuesReq) ProtoReflect() protoreflect.Message {
	mi := &file_filter_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilterValuesReq.ProtoReflect.Descriptor instead.
func (*GetFilterValuesReq) Descriptor() ([]byte, []int) {
	return file_filter_message_proto_rawDescGZIP(), []int{0}
}

func (x *GetFilterValuesReq) GetFilterName() string {
	if x != nil {
		return x.FilterName
	}
	return ""
}

type GetFilterValuesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *GetFilterValuesResp) Reset() {
	*x = GetFilterValuesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilterValuesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilterValuesResp) ProtoMessage() {}

func (x *GetFilterValuesResp) ProtoReflect() protoreflect.Message {
	mi := &file_filter_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilterValuesResp.ProtoReflect.Descriptor instead.
func (*GetFilterValuesResp) Descriptor() ([]byte, []int) {
	return file_filter_message_proto_rawDescGZIP(), []int{1}
}

func (x *GetFilterValuesResp) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// search query
type GetSearchQueryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queries []string `protobuf:"bytes,1,rep,name=queries,proto3" json:"queries,omitempty"`
}

func (x *GetSearchQueryReq) Reset() {
	*x = GetSearchQueryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSearchQueryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSearchQueryReq) ProtoMessage() {}

func (x *GetSearchQueryReq) ProtoReflect() protoreflect.Message {
	mi := &file_filter_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSearchQueryReq.ProtoReflect.Descriptor instead.
func (*GetSearchQueryReq) Descriptor() ([]byte, []int) {
	return file_filter_message_proto_rawDescGZIP(), []int{2}
}

func (x *GetSearchQueryReq) GetQueries() []string {
	if x != nil {
		return x.Queries
	}
	return nil
}

type GetSearchQueryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frequencies []int32 `protobuf:"varint,1,rep,packed,name=frequencies,proto3" json:"frequencies,omitempty"`
}

func (x *GetSearchQueryResp) Reset() {
	*x = GetSearchQueryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSearchQueryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSearchQueryResp) ProtoMessage() {}

func (x *GetSearchQueryResp) ProtoReflect() protoreflect.Message {
	mi := &file_filter_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSearchQueryResp.ProtoReflect.Descriptor instead.
func (*GetSearchQueryResp) Descriptor() ([]byte, []int) {
	return file_filter_message_proto_rawDescGZIP(), []int{3}
}

func (x *GetSearchQueryResp) GetFrequencies() []int32 {
	if x != nil {
		return x.Frequencies
	}
	return nil
}

// keywords by filter
type GetKeywordsByFilterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilterID int64 `protobuf:"varint,1,opt,name=filterID,proto3" json:"filterID,omitempty"`
	Limit    int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset   int32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetKeywordsByFilterReq) Reset() {
	*x = GetKeywordsByFilterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeywordsByFilterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeywordsByFilterReq) ProtoMessage() {}

func (x *GetKeywordsByFilterReq) ProtoReflect() protoreflect.Message {
	mi := &file_filter_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeywordsByFilterReq.ProtoReflect.Descriptor instead.
func (*GetKeywordsByFilterReq) Descriptor() ([]byte, []int) {
	return file_filter_message_proto_rawDescGZIP(), []int{4}
}

func (x *GetKeywordsByFilterReq) GetFilterID() int64 {
	if x != nil {
		return x.FilterID
	}
	return 0
}

func (x *GetKeywordsByFilterReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetKeywordsByFilterReq) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type KeywordByFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Normquery   string `protobuf:"bytes,1,opt,name=normquery,proto3" json:"normquery,omitempty"`
	Frequency   int32  `protobuf:"varint,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Competition int32  `protobuf:"varint,3,opt,name=competition,proto3" json:"competition,omitempty"`
	Count       int32  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *KeywordByFilter) Reset() {
	*x = KeywordByFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeywordByFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordByFilter) ProtoMessage() {}

func (x *KeywordByFilter) ProtoReflect() protoreflect.Message {
	mi := &file_filter_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordByFilter.ProtoReflect.Descriptor instead.
func (*KeywordByFilter) Descriptor() ([]byte, []int) {
	return file_filter_message_proto_rawDescGZIP(), []int{5}
}

func (x *KeywordByFilter) GetNormquery() string {
	if x != nil {
		return x.Normquery
	}
	return ""
}

func (x *KeywordByFilter) GetFrequency() int32 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *KeywordByFilter) GetCompetition() int32 {
	if x != nil {
		return x.Competition
	}
	return 0
}

func (x *KeywordByFilter) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetKeywordsByFilterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keywords []*KeywordByFilter `protobuf:"bytes,1,rep,name=keywords,proto3" json:"keywords,omitempty"`
}

func (x *GetKeywordsByFilterResp) Reset() {
	*x = GetKeywordsByFilterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeywordsByFilterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeywordsByFilterResp) ProtoMessage() {}

func (x *GetKeywordsByFilterResp) ProtoReflect() protoreflect.Message {
	mi := &file_filter_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeywordsByFilterResp.ProtoReflect.Descriptor instead.
func (*GetKeywordsByFilterResp) Descriptor() ([]byte, []int) {
	return file_filter_message_proto_rawDescGZIP(), []int{6}
}

func (x *GetKeywordsByFilterResp) GetKeywords() []*KeywordByFilter {
	if x != nil {
		return x.Keywords
	}
	return nil
}

var File_filter_message_proto protoreflect.FileDescriptor

var file_filter_message_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x34, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x22, 0x2d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x22, 0x36, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x22, 0x62, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x85, 0x01, 0x0a,
	0x0f, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x6f, 0x72, 0x6d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x6f, 0x72, 0x6d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x31, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x42, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filter_message_proto_rawDescOnce sync.Once
	file_filter_message_proto_rawDescData = file_filter_message_proto_rawDesc
)

func file_filter_message_proto_rawDescGZIP() []byte {
	file_filter_message_proto_rawDescOnce.Do(func() {
		file_filter_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_filter_message_proto_rawDescData)
	})
	return file_filter_message_proto_rawDescData
}

var file_filter_message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_filter_message_proto_goTypes = []interface{}{
	(*GetFilterValuesReq)(nil),      // 0: main.GetFilterValuesReq
	(*GetFilterValuesResp)(nil),     // 1: main.GetFilterValuesResp
	(*GetSearchQueryReq)(nil),       // 2: main.GetSearchQueryReq
	(*GetSearchQueryResp)(nil),      // 3: main.GetSearchQueryResp
	(*GetKeywordsByFilterReq)(nil),  // 4: main.GetKeywordsByFilterReq
	(*KeywordByFilter)(nil),         // 5: main.KeywordByFilter
	(*GetKeywordsByFilterResp)(nil), // 6: main.GetKeywordsByFilterResp
}
var file_filter_message_proto_depIdxs = []int32{
	5, // 0: main.GetKeywordsByFilterResp.keywords:type_name -> main.KeywordByFilter
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_filter_message_proto_init() }
func file_filter_message_proto_init() {
	if File_filter_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filter_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFilterValuesReq); i {
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
		file_filter_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFilterValuesResp); i {
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
		file_filter_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSearchQueryReq); i {
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
		file_filter_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSearchQueryResp); i {
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
		file_filter_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeywordsByFilterReq); i {
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
		file_filter_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeywordByFilter); i {
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
		file_filter_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeywordsByFilterResp); i {
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
			RawDescriptor: file_filter_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_filter_message_proto_goTypes,
		DependencyIndexes: file_filter_message_proto_depIdxs,
		MessageInfos:      file_filter_message_proto_msgTypes,
	}.Build()
	File_filter_message_proto = out.File
	file_filter_message_proto_rawDesc = nil
	file_filter_message_proto_goTypes = nil
	file_filter_message_proto_depIdxs = nil
}
