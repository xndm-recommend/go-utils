// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.5.1
// source: Tracing.proto

//*
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pb

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

//Used to pass through the information necessary to continue
//a trace after an RPC is made. All we need is the traceid
//(so we know the overarching trace this message is a part of), and
//the id of the current span when this message was sent, so we know
//what span caused the new span we will create when this message is received.
type RPCTInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId  *int64 `protobuf:"varint,1,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	ParentId *int64 `protobuf:"varint,2,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
}

func (x *RPCTInfo) Reset() {
	*x = RPCTInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tracing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCTInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCTInfo) ProtoMessage() {}

func (x *RPCTInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Tracing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCTInfo.ProtoReflect.Descriptor instead.
func (*RPCTInfo) Descriptor() ([]byte, []int) {
	return file_Tracing_proto_rawDescGZIP(), []int{0}
}

func (x *RPCTInfo) GetTraceId() int64 {
	if x != nil && x.TraceId != nil {
		return *x.TraceId
	}
	return 0
}

func (x *RPCTInfo) GetParentId() int64 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

var File_Tracing_proto protoreflect.FileDescriptor

var file_Tracing_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x42, 0x0a, 0x08, 0x52, 0x50, 0x43, 0x54, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x40, 0x0a, 0x2a, 0x6f, 0x72, 0x67, 0x2e, 0x61,
	0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01, 0xa0, 0x01, 0x01,
}

var (
	file_Tracing_proto_rawDescOnce sync.Once
	file_Tracing_proto_rawDescData = file_Tracing_proto_rawDesc
)

func file_Tracing_proto_rawDescGZIP() []byte {
	file_Tracing_proto_rawDescOnce.Do(func() {
		file_Tracing_proto_rawDescData = protoimpl.X.CompressGZIP(file_Tracing_proto_rawDescData)
	})
	return file_Tracing_proto_rawDescData
}

var file_Tracing_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Tracing_proto_goTypes = []interface{}{
	(*RPCTInfo)(nil), // 0: pb.RPCTInfo
}
var file_Tracing_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Tracing_proto_init() }
func file_Tracing_proto_init() {
	if File_Tracing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Tracing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCTInfo); i {
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
			RawDescriptor: file_Tracing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Tracing_proto_goTypes,
		DependencyIndexes: file_Tracing_proto_depIdxs,
		MessageInfos:      file_Tracing_proto_msgTypes,
	}.Build()
	File_Tracing_proto = out.File
	file_Tracing_proto_rawDesc = nil
	file_Tracing_proto_goTypes = nil
	file_Tracing_proto_depIdxs = nil
}