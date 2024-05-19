// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: graph_request.proto

package __

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

type GraphRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeNames []string `protobuf:"bytes,1,rep,name=node_names,json=nodeNames,proto3" json:"node_names,omitempty"`
}

func (x *GraphRequest) Reset() {
	*x = GraphRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graph_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphRequest) ProtoMessage() {}

func (x *GraphRequest) ProtoReflect() protoreflect.Message {
	mi := &file_graph_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphRequest.ProtoReflect.Descriptor instead.
func (*GraphRequest) Descriptor() ([]byte, []int) {
	return file_graph_request_proto_rawDescGZIP(), []int{0}
}

func (x *GraphRequest) GetNodeNames() []string {
	if x != nil {
		return x.NodeNames
	}
	return nil
}

var File_graph_request_proto protoreflect.FileDescriptor

var file_graph_request_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x22, 0x2d, 0x0a, 0x0c,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x42, 0x03, 0x5a, 0x01, 0x2e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_graph_request_proto_rawDescOnce sync.Once
	file_graph_request_proto_rawDescData = file_graph_request_proto_rawDesc
)

func file_graph_request_proto_rawDescGZIP() []byte {
	file_graph_request_proto_rawDescOnce.Do(func() {
		file_graph_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_graph_request_proto_rawDescData)
	})
	return file_graph_request_proto_rawDescData
}

var file_graph_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_graph_request_proto_goTypes = []interface{}{
	(*GraphRequest)(nil), // 0: graph.GraphRequest
}
var file_graph_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_graph_request_proto_init() }
func file_graph_request_proto_init() {
	if File_graph_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_graph_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphRequest); i {
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
			RawDescriptor: file_graph_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_graph_request_proto_goTypes,
		DependencyIndexes: file_graph_request_proto_depIdxs,
		MessageInfos:      file_graph_request_proto_msgTypes,
	}.Build()
	File_graph_request_proto = out.File
	file_graph_request_proto_rawDesc = nil
	file_graph_request_proto_goTypes = nil
	file_graph_request_proto_depIdxs = nil
}
