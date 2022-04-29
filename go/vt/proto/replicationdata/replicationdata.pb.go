//
//Copyright 2019 The Vitess Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// This file defines the replication related structures we use.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: replicationdata.proto

package replicationdata

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

// StopReplicationMode is used to provide controls over how replication is stopped.
type StopReplicationMode int32

const (
	StopReplicationMode_IOANDSQLTHREAD StopReplicationMode = 0
	StopReplicationMode_IOTHREADONLY   StopReplicationMode = 1
)

// Enum value maps for StopReplicationMode.
var (
	StopReplicationMode_name = map[int32]string{
		0: "IOANDSQLTHREAD",
		1: "IOTHREADONLY",
	}
	StopReplicationMode_value = map[string]int32{
		"IOANDSQLTHREAD": 0,
		"IOTHREADONLY":   1,
	}
)

func (x StopReplicationMode) Enum() *StopReplicationMode {
	p := new(StopReplicationMode)
	*p = x
	return p
}

func (x StopReplicationMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StopReplicationMode) Descriptor() protoreflect.EnumDescriptor {
	return file_replicationdata_proto_enumTypes[0].Descriptor()
}

func (StopReplicationMode) Type() protoreflect.EnumType {
	return &file_replicationdata_proto_enumTypes[0]
}

func (x StopReplicationMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StopReplicationMode.Descriptor instead.
func (StopReplicationMode) EnumDescriptor() ([]byte, []int) {
	return file_replicationdata_proto_rawDescGZIP(), []int{0}
}

// Status is the replication status for MySQL/MariaDB/File-based. Returned by a
// flavor-specific command and parsed into a Position and fields.
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position string `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	// These fields should be removed in Vitess 15+ and fully replaced by the io_state and sql_state fields
	// reserved 2, 3;
	// reserved "io_thread_running", "sql_thread_running";
	IoThreadRunning       bool   `protobuf:"varint,2,opt,name=io_thread_running,json=ioThreadRunning,proto3" json:"io_thread_running,omitempty"`
	SqlThreadRunning      bool   `protobuf:"varint,3,opt,name=sql_thread_running,json=sqlThreadRunning,proto3" json:"sql_thread_running,omitempty"`
	ReplicationLagSeconds uint32 `protobuf:"varint,4,opt,name=replication_lag_seconds,json=replicationLagSeconds,proto3" json:"replication_lag_seconds,omitempty"`
	SourceHost            string `protobuf:"bytes,5,opt,name=source_host,json=sourceHost,proto3" json:"source_host,omitempty"`
	SourcePort            int32  `protobuf:"varint,6,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty"`
	ConnectRetry          int32  `protobuf:"varint,7,opt,name=connect_retry,json=connectRetry,proto3" json:"connect_retry,omitempty"`
	// RelayLogPosition will be empty for flavors that do not support returning the full GTIDSet from the relay log, such as MariaDB.
	RelayLogPosition     string `protobuf:"bytes,8,opt,name=relay_log_position,json=relayLogPosition,proto3" json:"relay_log_position,omitempty"`
	FilePosition         string `protobuf:"bytes,9,opt,name=file_position,json=filePosition,proto3" json:"file_position,omitempty"`
	FileRelayLogPosition string `protobuf:"bytes,10,opt,name=file_relay_log_position,json=fileRelayLogPosition,proto3" json:"file_relay_log_position,omitempty"`
	SourceServerId       uint32 `protobuf:"varint,11,opt,name=source_server_id,json=sourceServerId,proto3" json:"source_server_id,omitempty"`
	SourceUuid           string `protobuf:"bytes,12,opt,name=source_uuid,json=sourceUuid,proto3" json:"source_uuid,omitempty"`
	IoState              int32  `protobuf:"varint,13,opt,name=io_state,json=ioState,proto3" json:"io_state,omitempty"`
	LastIoError          string `protobuf:"bytes,14,opt,name=last_io_error,json=lastIoError,proto3" json:"last_io_error,omitempty"`
	SqlState             int32  `protobuf:"varint,15,opt,name=sql_state,json=sqlState,proto3" json:"sql_state,omitempty"`
	LastSqlError         string `protobuf:"bytes,16,opt,name=last_sql_error,json=lastSqlError,proto3" json:"last_sql_error,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replicationdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_replicationdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_replicationdata_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *Status) GetIoThreadRunning() bool {
	if x != nil {
		return x.IoThreadRunning
	}
	return false
}

func (x *Status) GetSqlThreadRunning() bool {
	if x != nil {
		return x.SqlThreadRunning
	}
	return false
}

func (x *Status) GetReplicationLagSeconds() uint32 {
	if x != nil {
		return x.ReplicationLagSeconds
	}
	return 0
}

func (x *Status) GetSourceHost() string {
	if x != nil {
		return x.SourceHost
	}
	return ""
}

func (x *Status) GetSourcePort() int32 {
	if x != nil {
		return x.SourcePort
	}
	return 0
}

func (x *Status) GetConnectRetry() int32 {
	if x != nil {
		return x.ConnectRetry
	}
	return 0
}

func (x *Status) GetRelayLogPosition() string {
	if x != nil {
		return x.RelayLogPosition
	}
	return ""
}

func (x *Status) GetFilePosition() string {
	if x != nil {
		return x.FilePosition
	}
	return ""
}

func (x *Status) GetFileRelayLogPosition() string {
	if x != nil {
		return x.FileRelayLogPosition
	}
	return ""
}

func (x *Status) GetSourceServerId() uint32 {
	if x != nil {
		return x.SourceServerId
	}
	return 0
}

func (x *Status) GetSourceUuid() string {
	if x != nil {
		return x.SourceUuid
	}
	return ""
}

func (x *Status) GetIoState() int32 {
	if x != nil {
		return x.IoState
	}
	return 0
}

func (x *Status) GetLastIoError() string {
	if x != nil {
		return x.LastIoError
	}
	return ""
}

func (x *Status) GetSqlState() int32 {
	if x != nil {
		return x.SqlState
	}
	return 0
}

func (x *Status) GetLastSqlError() string {
	if x != nil {
		return x.LastSqlError
	}
	return ""
}

// StopReplicationStatus represents the replication status before calling StopReplication, and the replication status collected immediately after
// calling StopReplication.
type StopReplicationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Before *Status `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	After  *Status `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *StopReplicationStatus) Reset() {
	*x = StopReplicationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replicationdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopReplicationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopReplicationStatus) ProtoMessage() {}

func (x *StopReplicationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_replicationdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopReplicationStatus.ProtoReflect.Descriptor instead.
func (*StopReplicationStatus) Descriptor() ([]byte, []int) {
	return file_replicationdata_proto_rawDescGZIP(), []int{1}
}

func (x *StopReplicationStatus) GetBefore() *Status {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *StopReplicationStatus) GetAfter() *Status {
	if x != nil {
		return x.After
	}
	return nil
}

// PrimaryStatus is the replication status for a MySQL primary (returned by 'show master status').
type PrimaryStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position     string `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	FilePosition string `protobuf:"bytes,2,opt,name=file_position,json=filePosition,proto3" json:"file_position,omitempty"`
}

func (x *PrimaryStatus) Reset() {
	*x = PrimaryStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replicationdata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimaryStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimaryStatus) ProtoMessage() {}

func (x *PrimaryStatus) ProtoReflect() protoreflect.Message {
	mi := &file_replicationdata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimaryStatus.ProtoReflect.Descriptor instead.
func (*PrimaryStatus) Descriptor() ([]byte, []int) {
	return file_replicationdata_proto_rawDescGZIP(), []int{2}
}

func (x *PrimaryStatus) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *PrimaryStatus) GetFilePosition() string {
	if x != nil {
		return x.FilePosition
	}
	return ""
}

var File_replicationdata_proto protoreflect.FileDescriptor

var file_replicationdata_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x22, 0xf4, 0x04, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2a, 0x0a, 0x11, 0x69, 0x6f, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x6f, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x12, 0x73,
	0x71, 0x6c, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x73, 0x71, 0x6c, 0x54, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x36, 0x0a, 0x17, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61, 0x67, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x67, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x6f,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x72,
	0x65, 0x74, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x74, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x17, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x69, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x69, 0x6f, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x71, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x73, 0x71, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x73, 0x71, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x71, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x77, 0x0a, 0x15, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0x50, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69,
	0x6c, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x3b, 0x0a, 0x13, 0x53, 0x74,
	0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4f, 0x41, 0x4e, 0x44, 0x53, 0x51, 0x4c, 0x54, 0x48, 0x52,
	0x45, 0x41, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4f, 0x54, 0x48, 0x52, 0x45, 0x41,
	0x44, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x76,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_replicationdata_proto_rawDescOnce sync.Once
	file_replicationdata_proto_rawDescData = file_replicationdata_proto_rawDesc
)

func file_replicationdata_proto_rawDescGZIP() []byte {
	file_replicationdata_proto_rawDescOnce.Do(func() {
		file_replicationdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_replicationdata_proto_rawDescData)
	})
	return file_replicationdata_proto_rawDescData
}

var file_replicationdata_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_replicationdata_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_replicationdata_proto_goTypes = []interface{}{
	(StopReplicationMode)(0),      // 0: replicationdata.StopReplicationMode
	(*Status)(nil),                // 1: replicationdata.Status
	(*StopReplicationStatus)(nil), // 2: replicationdata.StopReplicationStatus
	(*PrimaryStatus)(nil),         // 3: replicationdata.PrimaryStatus
}
var file_replicationdata_proto_depIdxs = []int32{
	1, // 0: replicationdata.StopReplicationStatus.before:type_name -> replicationdata.Status
	1, // 1: replicationdata.StopReplicationStatus.after:type_name -> replicationdata.Status
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_replicationdata_proto_init() }
func file_replicationdata_proto_init() {
	if File_replicationdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_replicationdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_replicationdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopReplicationStatus); i {
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
		file_replicationdata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimaryStatus); i {
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
			RawDescriptor: file_replicationdata_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_replicationdata_proto_goTypes,
		DependencyIndexes: file_replicationdata_proto_depIdxs,
		EnumInfos:         file_replicationdata_proto_enumTypes,
		MessageInfos:      file_replicationdata_proto_msgTypes,
	}.Build()
	File_replicationdata_proto = out.File
	file_replicationdata_proto_rawDesc = nil
	file_replicationdata_proto_goTypes = nil
	file_replicationdata_proto_depIdxs = nil
}
