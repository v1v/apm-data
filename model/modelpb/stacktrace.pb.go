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
// 	protoc        v3.21.12
// source: stacktrace.proto

package modelpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StacktraceFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vars                *structpb.Struct `protobuf:"bytes,1,opt,name=vars,proto3" json:"vars,omitempty"`
	Lineno              *uint32          `protobuf:"varint,2,opt,name=lineno,proto3,oneof" json:"lineno,omitempty"`
	Colno               *uint32          `protobuf:"varint,3,opt,name=colno,proto3,oneof" json:"colno,omitempty"`
	Filename            string           `protobuf:"bytes,4,opt,name=filename,proto3" json:"filename,omitempty"`
	Classname           string           `protobuf:"bytes,5,opt,name=classname,proto3" json:"classname,omitempty"`
	ContextLine         string           `protobuf:"bytes,6,opt,name=context_line,json=contextLine,proto3" json:"context_line,omitempty"`
	Module              string           `protobuf:"bytes,7,opt,name=module,proto3" json:"module,omitempty"`
	Function            string           `protobuf:"bytes,8,opt,name=function,proto3" json:"function,omitempty"`
	AbsPath             string           `protobuf:"bytes,9,opt,name=abs_path,json=absPath,proto3" json:"abs_path,omitempty"`
	SourcemapError      string           `protobuf:"bytes,10,opt,name=sourcemap_error,json=sourcemapError,proto3" json:"sourcemap_error,omitempty"`
	Original            *Original        `protobuf:"bytes,11,opt,name=original,proto3" json:"original,omitempty"`
	PreContext          []string         `protobuf:"bytes,12,rep,name=pre_context,json=preContext,proto3" json:"pre_context,omitempty"`
	PostContext         []string         `protobuf:"bytes,13,rep,name=post_context,json=postContext,proto3" json:"post_context,omitempty"`
	LibraryFrame        bool             `protobuf:"varint,14,opt,name=library_frame,json=libraryFrame,proto3" json:"library_frame,omitempty"`
	SourcemapUpdated    bool             `protobuf:"varint,15,opt,name=sourcemap_updated,json=sourcemapUpdated,proto3" json:"sourcemap_updated,omitempty"`
	ExcludeFromGrouping bool             `protobuf:"varint,16,opt,name=exclude_from_grouping,json=excludeFromGrouping,proto3" json:"exclude_from_grouping,omitempty"`
}

func (x *StacktraceFrame) Reset() {
	*x = StacktraceFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stacktrace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StacktraceFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StacktraceFrame) ProtoMessage() {}

func (x *StacktraceFrame) ProtoReflect() protoreflect.Message {
	mi := &file_stacktrace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StacktraceFrame.ProtoReflect.Descriptor instead.
func (*StacktraceFrame) Descriptor() ([]byte, []int) {
	return file_stacktrace_proto_rawDescGZIP(), []int{0}
}

func (x *StacktraceFrame) GetVars() *structpb.Struct {
	if x != nil {
		return x.Vars
	}
	return nil
}

func (x *StacktraceFrame) GetLineno() uint32 {
	if x != nil && x.Lineno != nil {
		return *x.Lineno
	}
	return 0
}

func (x *StacktraceFrame) GetColno() uint32 {
	if x != nil && x.Colno != nil {
		return *x.Colno
	}
	return 0
}

func (x *StacktraceFrame) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *StacktraceFrame) GetClassname() string {
	if x != nil {
		return x.Classname
	}
	return ""
}

func (x *StacktraceFrame) GetContextLine() string {
	if x != nil {
		return x.ContextLine
	}
	return ""
}

func (x *StacktraceFrame) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *StacktraceFrame) GetFunction() string {
	if x != nil {
		return x.Function
	}
	return ""
}

func (x *StacktraceFrame) GetAbsPath() string {
	if x != nil {
		return x.AbsPath
	}
	return ""
}

func (x *StacktraceFrame) GetSourcemapError() string {
	if x != nil {
		return x.SourcemapError
	}
	return ""
}

func (x *StacktraceFrame) GetOriginal() *Original {
	if x != nil {
		return x.Original
	}
	return nil
}

func (x *StacktraceFrame) GetPreContext() []string {
	if x != nil {
		return x.PreContext
	}
	return nil
}

func (x *StacktraceFrame) GetPostContext() []string {
	if x != nil {
		return x.PostContext
	}
	return nil
}

func (x *StacktraceFrame) GetLibraryFrame() bool {
	if x != nil {
		return x.LibraryFrame
	}
	return false
}

func (x *StacktraceFrame) GetSourcemapUpdated() bool {
	if x != nil {
		return x.SourcemapUpdated
	}
	return false
}

func (x *StacktraceFrame) GetExcludeFromGrouping() bool {
	if x != nil {
		return x.ExcludeFromGrouping
	}
	return false
}

type Original struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbsPath      string  `protobuf:"bytes,1,opt,name=abs_path,json=absPath,proto3" json:"abs_path,omitempty"`
	Filename     string  `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Classname    string  `protobuf:"bytes,3,opt,name=classname,proto3" json:"classname,omitempty"`
	Lineno       *uint32 `protobuf:"varint,4,opt,name=lineno,proto3,oneof" json:"lineno,omitempty"`
	Colno        *uint32 `protobuf:"varint,5,opt,name=colno,proto3,oneof" json:"colno,omitempty"`
	Function     string  `protobuf:"bytes,6,opt,name=function,proto3" json:"function,omitempty"`
	LibraryFrame bool    `protobuf:"varint,7,opt,name=library_frame,json=libraryFrame,proto3" json:"library_frame,omitempty"`
}

func (x *Original) Reset() {
	*x = Original{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stacktrace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Original) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Original) ProtoMessage() {}

func (x *Original) ProtoReflect() protoreflect.Message {
	mi := &file_stacktrace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Original.ProtoReflect.Descriptor instead.
func (*Original) Descriptor() ([]byte, []int) {
	return file_stacktrace_proto_rawDescGZIP(), []int{1}
}

func (x *Original) GetAbsPath() string {
	if x != nil {
		return x.AbsPath
	}
	return ""
}

func (x *Original) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Original) GetClassname() string {
	if x != nil {
		return x.Classname
	}
	return ""
}

func (x *Original) GetLineno() uint32 {
	if x != nil && x.Lineno != nil {
		return *x.Lineno
	}
	return 0
}

func (x *Original) GetColno() uint32 {
	if x != nil && x.Colno != nil {
		return *x.Colno
	}
	return 0
}

func (x *Original) GetFunction() string {
	if x != nil {
		return x.Function
	}
	return ""
}

func (x *Original) GetLibraryFrame() bool {
	if x != nil {
		return x.LibraryFrame
	}
	return false
}

var File_stacktrace_proto protoreflect.FileDescriptor

var file_stacktrace_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe0, 0x04, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x76, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x76, 0x61, 0x72,
	0x73, 0x12, 0x1b, 0x0a, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x00, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x6e, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52,
	0x05, 0x63, 0x6f, 0x6c, 0x6e, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x62,
	0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x62,
	0x73, 0x50, 0x61, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d,
	0x61, 0x70, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x70, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x34,
	0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6f, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a,
	0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x70, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6d, 0x61, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x69, 0x6e, 0x67, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x65, 0x78, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x6e, 0x6f, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f,
	0x6c, 0x6e, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x08, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x62, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x62, 0x73, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x6e, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x6e, 0x6f, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x01, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6e, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x6e, 0x6f, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f,
	0x6c, 0x6e, 0x6f, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x6d, 0x2d, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stacktrace_proto_rawDescOnce sync.Once
	file_stacktrace_proto_rawDescData = file_stacktrace_proto_rawDesc
)

func file_stacktrace_proto_rawDescGZIP() []byte {
	file_stacktrace_proto_rawDescOnce.Do(func() {
		file_stacktrace_proto_rawDescData = protoimpl.X.CompressGZIP(file_stacktrace_proto_rawDescData)
	})
	return file_stacktrace_proto_rawDescData
}

var file_stacktrace_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stacktrace_proto_goTypes = []interface{}{
	(*StacktraceFrame)(nil), // 0: elastic.apm.v1.StacktraceFrame
	(*Original)(nil),        // 1: elastic.apm.v1.Original
	(*structpb.Struct)(nil), // 2: google.protobuf.Struct
}
var file_stacktrace_proto_depIdxs = []int32{
	2, // 0: elastic.apm.v1.StacktraceFrame.vars:type_name -> google.protobuf.Struct
	1, // 1: elastic.apm.v1.StacktraceFrame.original:type_name -> elastic.apm.v1.Original
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stacktrace_proto_init() }
func file_stacktrace_proto_init() {
	if File_stacktrace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stacktrace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StacktraceFrame); i {
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
		file_stacktrace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Original); i {
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
	file_stacktrace_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_stacktrace_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stacktrace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stacktrace_proto_goTypes,
		DependencyIndexes: file_stacktrace_proto_depIdxs,
		MessageInfos:      file_stacktrace_proto_msgTypes,
	}.Build()
	File_stacktrace_proto = out.File
	file_stacktrace_proto_rawDesc = nil
	file_stacktrace_proto_goTypes = nil
	file_stacktrace_proto_depIdxs = nil
}