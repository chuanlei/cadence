// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: uber/cadence/api/v1/domain.proto

package apiv1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type DomainStatus int32

const (
	DomainStatus_DOMAIN_STATUS_INVALID    DomainStatus = 0
	DomainStatus_DOMAIN_STATUS_REGISTERED DomainStatus = 1
	DomainStatus_DOMAIN_STATUS_DEPRECATED DomainStatus = 2
	DomainStatus_DOMAIN_STATUS_DELETED    DomainStatus = 3
)

// Enum value maps for DomainStatus.
var (
	DomainStatus_name = map[int32]string{
		0: "DOMAIN_STATUS_INVALID",
		1: "DOMAIN_STATUS_REGISTERED",
		2: "DOMAIN_STATUS_DEPRECATED",
		3: "DOMAIN_STATUS_DELETED",
	}
	DomainStatus_value = map[string]int32{
		"DOMAIN_STATUS_INVALID":    0,
		"DOMAIN_STATUS_REGISTERED": 1,
		"DOMAIN_STATUS_DEPRECATED": 2,
		"DOMAIN_STATUS_DELETED":    3,
	}
)

func (x DomainStatus) Enum() *DomainStatus {
	p := new(DomainStatus)
	*p = x
	return p
}

func (x DomainStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DomainStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_uber_cadence_api_v1_domain_proto_enumTypes[0].Descriptor()
}

func (DomainStatus) Type() protoreflect.EnumType {
	return &file_uber_cadence_api_v1_domain_proto_enumTypes[0]
}

func (x DomainStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DomainStatus.Descriptor instead.
func (DomainStatus) EnumDescriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_domain_proto_rawDescGZIP(), []int{0}
}

type ArchivalStatus int32

const (
	ArchivalStatus_ARCHIVAL_STATUS_INVALID  ArchivalStatus = 0
	ArchivalStatus_ARCHIVAL_STATUS_DISABLED ArchivalStatus = 1
	ArchivalStatus_ARCHIVAL_STATUS_ENABLED  ArchivalStatus = 2
)

// Enum value maps for ArchivalStatus.
var (
	ArchivalStatus_name = map[int32]string{
		0: "ARCHIVAL_STATUS_INVALID",
		1: "ARCHIVAL_STATUS_DISABLED",
		2: "ARCHIVAL_STATUS_ENABLED",
	}
	ArchivalStatus_value = map[string]int32{
		"ARCHIVAL_STATUS_INVALID":  0,
		"ARCHIVAL_STATUS_DISABLED": 1,
		"ARCHIVAL_STATUS_ENABLED":  2,
	}
)

func (x ArchivalStatus) Enum() *ArchivalStatus {
	p := new(ArchivalStatus)
	*p = x
	return p
}

func (x ArchivalStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArchivalStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_uber_cadence_api_v1_domain_proto_enumTypes[1].Descriptor()
}

func (ArchivalStatus) Type() protoreflect.EnumType {
	return &file_uber_cadence_api_v1_domain_proto_enumTypes[1]
}

func (x ArchivalStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArchivalStatus.Descriptor instead.
func (ArchivalStatus) EnumDescriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_domain_proto_rawDescGZIP(), []int{1}
}

type Domain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                               string                             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                             string                             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status                           DomainStatus                       `protobuf:"varint,3,opt,name=status,proto3,enum=uber.cadence.api.v1.DomainStatus" json:"status,omitempty"`
	Description                      string                             `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	OwnerEmail                       string                             `protobuf:"bytes,5,opt,name=owner_email,json=ownerEmail,proto3" json:"owner_email,omitempty"`
	Data                             map[string]string                  `protobuf:"bytes,6,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	WorkflowExecutionRetentionPeriod *durationpb.Duration               `protobuf:"bytes,7,opt,name=workflow_execution_retention_period,json=workflowExecutionRetentionPeriod,proto3" json:"workflow_execution_retention_period,omitempty"`
	BadBinaries                      *BadBinaries                       `protobuf:"bytes,8,opt,name=bad_binaries,json=badBinaries,proto3" json:"bad_binaries,omitempty"`
	HistoryArchivalStatus            ArchivalStatus                     `protobuf:"varint,9,opt,name=history_archival_status,json=historyArchivalStatus,proto3,enum=uber.cadence.api.v1.ArchivalStatus" json:"history_archival_status,omitempty"`
	HistoryArchivalUri               string                             `protobuf:"bytes,10,opt,name=history_archival_uri,json=historyArchivalUri,proto3" json:"history_archival_uri,omitempty"`
	VisibilityArchivalStatus         ArchivalStatus                     `protobuf:"varint,11,opt,name=visibility_archival_status,json=visibilityArchivalStatus,proto3,enum=uber.cadence.api.v1.ArchivalStatus" json:"visibility_archival_status,omitempty"`
	VisibilityArchivalUri            string                             `protobuf:"bytes,12,opt,name=visibility_archival_uri,json=visibilityArchivalUri,proto3" json:"visibility_archival_uri,omitempty"`
	ActiveClusterName                string                             `protobuf:"bytes,13,opt,name=active_cluster_name,json=activeClusterName,proto3" json:"active_cluster_name,omitempty"`
	Clusters                         []*ClusterReplicationConfiguration `protobuf:"bytes,14,rep,name=clusters,proto3" json:"clusters,omitempty"`
	FailoverVersion                  int64                              `protobuf:"varint,15,opt,name=failover_version,json=failoverVersion,proto3" json:"failover_version,omitempty"`
	IsGlobalDomain                   bool                               `protobuf:"varint,16,opt,name=is_global_domain,json=isGlobalDomain,proto3" json:"is_global_domain,omitempty"`
}

func (x *Domain) Reset() {
	*x = Domain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain) ProtoMessage() {}

func (x *Domain) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain.ProtoReflect.Descriptor instead.
func (*Domain) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_domain_proto_rawDescGZIP(), []int{0}
}

func (x *Domain) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Domain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Domain) GetStatus() DomainStatus {
	if x != nil {
		return x.Status
	}
	return DomainStatus_DOMAIN_STATUS_INVALID
}

func (x *Domain) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Domain) GetOwnerEmail() string {
	if x != nil {
		return x.OwnerEmail
	}
	return ""
}

func (x *Domain) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Domain) GetWorkflowExecutionRetentionPeriod() *durationpb.Duration {
	if x != nil {
		return x.WorkflowExecutionRetentionPeriod
	}
	return nil
}

func (x *Domain) GetBadBinaries() *BadBinaries {
	if x != nil {
		return x.BadBinaries
	}
	return nil
}

func (x *Domain) GetHistoryArchivalStatus() ArchivalStatus {
	if x != nil {
		return x.HistoryArchivalStatus
	}
	return ArchivalStatus_ARCHIVAL_STATUS_INVALID
}

func (x *Domain) GetHistoryArchivalUri() string {
	if x != nil {
		return x.HistoryArchivalUri
	}
	return ""
}

func (x *Domain) GetVisibilityArchivalStatus() ArchivalStatus {
	if x != nil {
		return x.VisibilityArchivalStatus
	}
	return ArchivalStatus_ARCHIVAL_STATUS_INVALID
}

func (x *Domain) GetVisibilityArchivalUri() string {
	if x != nil {
		return x.VisibilityArchivalUri
	}
	return ""
}

func (x *Domain) GetActiveClusterName() string {
	if x != nil {
		return x.ActiveClusterName
	}
	return ""
}

func (x *Domain) GetClusters() []*ClusterReplicationConfiguration {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *Domain) GetFailoverVersion() int64 {
	if x != nil {
		return x.FailoverVersion
	}
	return 0
}

func (x *Domain) GetIsGlobalDomain() bool {
	if x != nil {
		return x.IsGlobalDomain
	}
	return false
}

type ClusterReplicationConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
}

func (x *ClusterReplicationConfiguration) Reset() {
	*x = ClusterReplicationConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_domain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterReplicationConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterReplicationConfiguration) ProtoMessage() {}

func (x *ClusterReplicationConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_domain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterReplicationConfiguration.ProtoReflect.Descriptor instead.
func (*ClusterReplicationConfiguration) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_domain_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterReplicationConfiguration) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

type BadBinaries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Binaries map[string]*BadBinaryInfo `protobuf:"bytes,1,rep,name=binaries,proto3" json:"binaries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BadBinaries) Reset() {
	*x = BadBinaries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_domain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadBinaries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadBinaries) ProtoMessage() {}

func (x *BadBinaries) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_domain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadBinaries.ProtoReflect.Descriptor instead.
func (*BadBinaries) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_domain_proto_rawDescGZIP(), []int{2}
}

func (x *BadBinaries) GetBinaries() map[string]*BadBinaryInfo {
	if x != nil {
		return x.Binaries
	}
	return nil
}

type BadBinaryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason      string                 `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
	Operator    string                 `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
}

func (x *BadBinaryInfo) Reset() {
	*x = BadBinaryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uber_cadence_api_v1_domain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadBinaryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadBinaryInfo) ProtoMessage() {}

func (x *BadBinaryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_uber_cadence_api_v1_domain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadBinaryInfo.ProtoReflect.Descriptor instead.
func (*BadBinaryInfo) Descriptor() ([]byte, []int) {
	return file_uber_cadence_api_v1_domain_proto_rawDescGZIP(), []int{3}
}

func (x *BadBinaryInfo) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *BadBinaryInfo) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *BadBinaryInfo) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

var File_uber_cadence_api_v1_domain_proto protoreflect.FileDescriptor

var file_uber_cadence_api_v1_domain_proto_rawDesc = []byte{
	0x0a, 0x20, 0x75, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x07, 0x0a, 0x06, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63,
	0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x68, 0x0a, 0x23, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x20, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x74, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x43, 0x0a, 0x0c, 0x62, 0x61,
	0x64, 0x5f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x64, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x0b, 0x62, 0x61, 0x64, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x5b, 0x0a, 0x17, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x23, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x15, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a, 0x14,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x55, 0x72, 0x69, 0x12, 0x61,
	0x0a, 0x1a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x23, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x18, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x36, 0x0a, 0x17, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x15, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x55, 0x72, 0x69, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x08, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x75, 0x62,
	0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x66,
	0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x69, 0x73, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x44, 0x0a, 0x1f, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0xba, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x64, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x4a, 0x0a, 0x08, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x64, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x69, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x5f, 0x0a, 0x0d, 0x42,
	0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x64, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x82, 0x01, 0x0a,
	0x0d, 0x42, 0x61, 0x64, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x2a, 0x80, 0x01, 0x0a, 0x0c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a,
	0x18, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52,
	0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x44,
	0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x50,
	0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x4f, 0x4d,
	0x41, 0x49, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54,
	0x45, 0x44, 0x10, 0x03, 0x2a, 0x68, 0x0a, 0x0e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56,
	0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x41, 0x4c, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x41, 0x4c, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x42, 0x56,
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x62, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x75, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x2e,
	0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uber_cadence_api_v1_domain_proto_rawDescOnce sync.Once
	file_uber_cadence_api_v1_domain_proto_rawDescData = file_uber_cadence_api_v1_domain_proto_rawDesc
)

func file_uber_cadence_api_v1_domain_proto_rawDescGZIP() []byte {
	file_uber_cadence_api_v1_domain_proto_rawDescOnce.Do(func() {
		file_uber_cadence_api_v1_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_uber_cadence_api_v1_domain_proto_rawDescData)
	})
	return file_uber_cadence_api_v1_domain_proto_rawDescData
}

var file_uber_cadence_api_v1_domain_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_uber_cadence_api_v1_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_uber_cadence_api_v1_domain_proto_goTypes = []interface{}{
	(DomainStatus)(0),                       // 0: uber.cadence.api.v1.DomainStatus
	(ArchivalStatus)(0),                     // 1: uber.cadence.api.v1.ArchivalStatus
	(*Domain)(nil),                          // 2: uber.cadence.api.v1.Domain
	(*ClusterReplicationConfiguration)(nil), // 3: uber.cadence.api.v1.ClusterReplicationConfiguration
	(*BadBinaries)(nil),                     // 4: uber.cadence.api.v1.BadBinaries
	(*BadBinaryInfo)(nil),                   // 5: uber.cadence.api.v1.BadBinaryInfo
	nil,                                     // 6: uber.cadence.api.v1.Domain.DataEntry
	nil,                                     // 7: uber.cadence.api.v1.BadBinaries.BinariesEntry
	(*durationpb.Duration)(nil),             // 8: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil),           // 9: google.protobuf.Timestamp
}
var file_uber_cadence_api_v1_domain_proto_depIdxs = []int32{
	0,  // 0: uber.cadence.api.v1.Domain.status:type_name -> uber.cadence.api.v1.DomainStatus
	6,  // 1: uber.cadence.api.v1.Domain.data:type_name -> uber.cadence.api.v1.Domain.DataEntry
	8,  // 2: uber.cadence.api.v1.Domain.workflow_execution_retention_period:type_name -> google.protobuf.Duration
	4,  // 3: uber.cadence.api.v1.Domain.bad_binaries:type_name -> uber.cadence.api.v1.BadBinaries
	1,  // 4: uber.cadence.api.v1.Domain.history_archival_status:type_name -> uber.cadence.api.v1.ArchivalStatus
	1,  // 5: uber.cadence.api.v1.Domain.visibility_archival_status:type_name -> uber.cadence.api.v1.ArchivalStatus
	3,  // 6: uber.cadence.api.v1.Domain.clusters:type_name -> uber.cadence.api.v1.ClusterReplicationConfiguration
	7,  // 7: uber.cadence.api.v1.BadBinaries.binaries:type_name -> uber.cadence.api.v1.BadBinaries.BinariesEntry
	9,  // 8: uber.cadence.api.v1.BadBinaryInfo.created_time:type_name -> google.protobuf.Timestamp
	5,  // 9: uber.cadence.api.v1.BadBinaries.BinariesEntry.value:type_name -> uber.cadence.api.v1.BadBinaryInfo
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_uber_cadence_api_v1_domain_proto_init() }
func file_uber_cadence_api_v1_domain_proto_init() {
	if File_uber_cadence_api_v1_domain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_uber_cadence_api_v1_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Domain); i {
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
		file_uber_cadence_api_v1_domain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterReplicationConfiguration); i {
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
		file_uber_cadence_api_v1_domain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadBinaries); i {
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
		file_uber_cadence_api_v1_domain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadBinaryInfo); i {
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
			RawDescriptor: file_uber_cadence_api_v1_domain_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_uber_cadence_api_v1_domain_proto_goTypes,
		DependencyIndexes: file_uber_cadence_api_v1_domain_proto_depIdxs,
		EnumInfos:         file_uber_cadence_api_v1_domain_proto_enumTypes,
		MessageInfos:      file_uber_cadence_api_v1_domain_proto_msgTypes,
	}.Build()
	File_uber_cadence_api_v1_domain_proto = out.File
	file_uber_cadence_api_v1_domain_proto_rawDesc = nil
	file_uber_cadence_api_v1_domain_proto_goTypes = nil
	file_uber_cadence_api_v1_domain_proto_depIdxs = nil
}
