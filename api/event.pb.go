// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.3
// source: api/event.proto

package api

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

type EventId int32

const (
	// ---- message event ----
	EventId_MessageEventAll            EventId = 0
	EventId_MessageEventPrivateMessage EventId = 1
	EventId_MessageEventGroupMessage   EventId = 2
	// ---- notice event ----
	EventId_NoticeEventAll EventId = 100
)

// Enum value maps for EventId.
var (
	EventId_name = map[int32]string{
		0:   "MessageEventAll",
		1:   "MessageEventPrivateMessage",
		2:   "MessageEventGroupMessage",
		100: "NoticeEventAll",
	}
	EventId_value = map[string]int32{
		"MessageEventAll":            0,
		"MessageEventPrivateMessage": 1,
		"MessageEventGroupMessage":   2,
		"NoticeEventAll":             100,
	}
)

func (x EventId) Enum() *EventId {
	p := new(EventId)
	*p = x
	return p
}

func (x EventId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventId) Descriptor() protoreflect.EnumDescriptor {
	return file_api_event_proto_enumTypes[0].Descriptor()
}

func (EventId) Type() protoreflect.EnumType {
	return &file_api_event_proto_enumTypes[0]
}

func (x EventId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventId.Descriptor instead.
func (EventId) EnumDescriptor() ([]byte, []int) {
	return file_api_event_proto_rawDescGZIP(), []int{0}
}

var File_api_event_proto protoreflect.FileDescriptor

var file_api_event_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x71, 0x71, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2a, 0x70, 0x0a, 0x07,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x10, 0x64, 0x42, 0x0b,
	0x5a, 0x09, 0x71, 0x71, 0x62, 0x6f, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_event_proto_rawDescOnce sync.Once
	file_api_event_proto_rawDescData = file_api_event_proto_rawDesc
)

func file_api_event_proto_rawDescGZIP() []byte {
	file_api_event_proto_rawDescOnce.Do(func() {
		file_api_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_event_proto_rawDescData)
	})
	return file_api_event_proto_rawDescData
}

var file_api_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_event_proto_goTypes = []interface{}{
	(EventId)(0), // 0: qqbot.api.EventId
}
var file_api_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_event_proto_init() }
func file_api_event_proto_init() {
	if File_api_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_event_proto_goTypes,
		DependencyIndexes: file_api_event_proto_depIdxs,
		EnumInfos:         file_api_event_proto_enumTypes,
	}.Build()
	File_api_event_proto = out.File
	file_api_event_proto_rawDesc = nil
	file_api_event_proto_goTypes = nil
	file_api_event_proto_depIdxs = nil
}