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
// source: error.proto

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

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Custom      *structpb.Struct `protobuf:"bytes,1,opt,name=custom,proto3" json:"custom,omitempty"`
	Exception   *Exception       `protobuf:"bytes,2,opt,name=exception,proto3" json:"exception,omitempty"`
	Log         *ErrorLog        `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Id          string           `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	GroupingKey string           `protobuf:"bytes,5,opt,name=grouping_key,json=groupingKey,proto3" json:"grouping_key,omitempty"`
	Culprit     string           `protobuf:"bytes,6,opt,name=culprit,proto3" json:"culprit,omitempty"`
	StackTrace  string           `protobuf:"bytes,7,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
	Message     string           `protobuf:"bytes,8,opt,name=message,proto3" json:"message,omitempty"`
	Type        string           `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCustom() *structpb.Struct {
	if x != nil {
		return x.Custom
	}
	return nil
}

func (x *Error) GetException() *Exception {
	if x != nil {
		return x.Exception
	}
	return nil
}

func (x *Error) GetLog() *ErrorLog {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *Error) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Error) GetGroupingKey() string {
	if x != nil {
		return x.GroupingKey
	}
	return ""
}

func (x *Error) GetCulprit() string {
	if x != nil {
		return x.Culprit
	}
	return ""
}

func (x *Error) GetStackTrace() string {
	if x != nil {
		return x.StackTrace
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Exception struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message    string             `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Module     string             `protobuf:"bytes,2,opt,name=module,proto3" json:"module,omitempty"`
	Code       string             `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Attributes *structpb.Struct   `protobuf:"bytes,4,opt,name=attributes,proto3" json:"attributes,omitempty"`
	Stacktrace []*StacktraceFrame `protobuf:"bytes,5,rep,name=stacktrace,proto3" json:"stacktrace,omitempty"`
	Type       string             `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Handled    *bool              `protobuf:"varint,7,opt,name=handled,proto3,oneof" json:"handled,omitempty"`
	Cause      []*Exception       `protobuf:"bytes,8,rep,name=cause,proto3" json:"cause,omitempty"`
}

func (x *Exception) Reset() {
	*x = Exception{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exception) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exception) ProtoMessage() {}

func (x *Exception) ProtoReflect() protoreflect.Message {
	mi := &file_error_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exception.ProtoReflect.Descriptor instead.
func (*Exception) Descriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{1}
}

func (x *Exception) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Exception) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *Exception) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Exception) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Exception) GetStacktrace() []*StacktraceFrame {
	if x != nil {
		return x.Stacktrace
	}
	return nil
}

func (x *Exception) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Exception) GetHandled() bool {
	if x != nil && x.Handled != nil {
		return *x.Handled
	}
	return false
}

func (x *Exception) GetCause() []*Exception {
	if x != nil {
		return x.Cause
	}
	return nil
}

type ErrorLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      string             `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Level        string             `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	ParamMessage string             `protobuf:"bytes,3,opt,name=param_message,json=paramMessage,proto3" json:"param_message,omitempty"`
	LoggerName   string             `protobuf:"bytes,4,opt,name=logger_name,json=loggerName,proto3" json:"logger_name,omitempty"`
	Stacktrace   []*StacktraceFrame `protobuf:"bytes,5,rep,name=stacktrace,proto3" json:"stacktrace,omitempty"`
}

func (x *ErrorLog) Reset() {
	*x = ErrorLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorLog) ProtoMessage() {}

func (x *ErrorLog) ProtoReflect() protoreflect.Message {
	mi := &file_error_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorLog.ProtoReflect.Descriptor instead.
func (*ErrorLog) Descriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorLog) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorLog) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *ErrorLog) GetParamMessage() string {
	if x != nil {
		return x.ParamMessage
	}
	return ""
}

func (x *ErrorLog) GetLoggerName() string {
	if x != nil {
		return x.LoggerName
	}
	return ""
}

func (x *ErrorLog) GetStacktrace() []*StacktraceFrame {
	if x != nil {
		return x.Stacktrace
	}
	return nil
}

var File_error_proto protoreflect.FileDescriptor

var file_error_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65,
	0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x02,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x65,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6c,
	0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63,
	0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x6c, 0x70, 0x72, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x75, 0x6c, 0x70, 0x72, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xbb, 0x02, 0x0a, 0x09, 0x45, 0x78,
	0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x37, 0x0a,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6c, 0x61,
	0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x05, 0x63, 0x61,
	0x75, 0x73, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6c, 0x61, 0x73,
	0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x65, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x63, 0x61, 0x75, 0x73, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x22, 0xc1, 0x01, 0x0a, 0x08, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2f, 0x61, 0x70, 0x6d, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_error_proto_rawDescOnce sync.Once
	file_error_proto_rawDescData = file_error_proto_rawDesc
)

func file_error_proto_rawDescGZIP() []byte {
	file_error_proto_rawDescOnce.Do(func() {
		file_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_proto_rawDescData)
	})
	return file_error_proto_rawDescData
}

var file_error_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_error_proto_goTypes = []interface{}{
	(*Error)(nil),           // 0: elastic.apm.v1.Error
	(*Exception)(nil),       // 1: elastic.apm.v1.Exception
	(*ErrorLog)(nil),        // 2: elastic.apm.v1.ErrorLog
	(*structpb.Struct)(nil), // 3: google.protobuf.Struct
	(*StacktraceFrame)(nil), // 4: elastic.apm.v1.StacktraceFrame
}
var file_error_proto_depIdxs = []int32{
	3, // 0: elastic.apm.v1.Error.custom:type_name -> google.protobuf.Struct
	1, // 1: elastic.apm.v1.Error.exception:type_name -> elastic.apm.v1.Exception
	2, // 2: elastic.apm.v1.Error.log:type_name -> elastic.apm.v1.ErrorLog
	3, // 3: elastic.apm.v1.Exception.attributes:type_name -> google.protobuf.Struct
	4, // 4: elastic.apm.v1.Exception.stacktrace:type_name -> elastic.apm.v1.StacktraceFrame
	1, // 5: elastic.apm.v1.Exception.cause:type_name -> elastic.apm.v1.Exception
	4, // 6: elastic.apm.v1.ErrorLog.stacktrace:type_name -> elastic.apm.v1.StacktraceFrame
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_error_proto_init() }
func file_error_proto_init() {
	if File_error_proto != nil {
		return
	}
	file_stacktrace_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_error_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exception); i {
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
		file_error_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorLog); i {
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
	file_error_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_proto_goTypes,
		DependencyIndexes: file_error_proto_depIdxs,
		MessageInfos:      file_error_proto_msgTypes,
	}.Build()
	File_error_proto = out.File
	file_error_proto_rawDesc = nil
	file_error_proto_goTypes = nil
	file_error_proto_depIdxs = nil
}
