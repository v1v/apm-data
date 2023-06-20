// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: host.proto

package modelpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Host struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Os           *OS      `protobuf:"bytes,1,opt,name=os,proto3" json:"os,omitempty"`
	Hostname     string   `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Name         string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Id           string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Architecture string   `protobuf:"bytes,5,opt,name=architecture,proto3" json:"architecture,omitempty"`
	Type         string   `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Ip           []string `protobuf:"bytes,7,rep,name=ip,proto3" json:"ip,omitempty"`
}

func (x *Host) Reset() {
	*x = Host{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_host_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_host_proto_rawDescGZIP(), []int{0}
}

func (x *Host) GetOs() *OS {
	if x != nil {
		return x.Os
	}
	return nil
}

func (x *Host) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Host) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Host) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Host) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

func (x *Host) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Host) GetIp() []string {
	if x != nil {
		return x.Ip
	}
	return nil
}

var File_host_proto protoreflect.FileDescriptor

var file_host_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x6c,
	0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x08, 0x6f, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6c,
	0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x53, 0x52,
	0x02, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x70, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2f, 0x61, 0x70, 0x6d, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_host_proto_rawDescOnce sync.Once
	file_host_proto_rawDescData = file_host_proto_rawDesc
)

func file_host_proto_rawDescGZIP() []byte {
	file_host_proto_rawDescOnce.Do(func() {
		file_host_proto_rawDescData = protoimpl.X.CompressGZIP(file_host_proto_rawDescData)
	})
	return file_host_proto_rawDescData
}

var file_host_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_host_proto_goTypes = []interface{}{
	(*Host)(nil), // 0: elastic.apm.v1.Host
	(*OS)(nil),   // 1: elastic.apm.v1.OS
}
var file_host_proto_depIdxs = []int32{
	1, // 0: elastic.apm.v1.Host.os:type_name -> elastic.apm.v1.OS
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_host_proto_init() }
func file_host_proto_init() {
	if File_host_proto != nil {
		return
	}
	file_os_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_host_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Host); i {
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
			RawDescriptor: file_host_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_host_proto_goTypes,
		DependencyIndexes: file_host_proto_depIdxs,
		MessageInfos:      file_host_proto_msgTypes,
	}.Build()
	File_host_proto = out.File
	file_host_proto_rawDesc = nil
	file_host_proto_goTypes = nil
	file_host_proto_depIdxs = nil
}
