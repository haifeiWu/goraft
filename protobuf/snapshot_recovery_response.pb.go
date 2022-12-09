// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: protobuf/snapshot_recovery_response.proto

package protobuf

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

type SnapshotRecoveryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        *uint64 `protobuf:"varint,1,req,name=Term" json:"Term,omitempty"`
	Success     *bool   `protobuf:"varint,2,req,name=Success" json:"Success,omitempty"`
	CommitIndex *uint64 `protobuf:"varint,3,req,name=CommitIndex" json:"CommitIndex,omitempty"`
}

func (x *SnapshotRecoveryResponse) Reset() {
	*x = SnapshotRecoveryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_snapshot_recovery_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotRecoveryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotRecoveryResponse) ProtoMessage() {}

func (x *SnapshotRecoveryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_snapshot_recovery_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotRecoveryResponse.ProtoReflect.Descriptor instead.
func (*SnapshotRecoveryResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_snapshot_recovery_response_proto_rawDescGZIP(), []int{0}
}

func (x *SnapshotRecoveryResponse) GetTerm() uint64 {
	if x != nil && x.Term != nil {
		return *x.Term
	}
	return 0
}

func (x *SnapshotRecoveryResponse) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *SnapshotRecoveryResponse) GetCommitIndex() uint64 {
	if x != nil && x.CommitIndex != nil {
		return *x.CommitIndex
	}
	return 0
}

var File_protobuf_snapshot_recovery_response_proto protoreflect.FileDescriptor

var file_protobuf_snapshot_recovery_response_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x6a, 0x0a, 0x18, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x52,
	0x04, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x04, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_protobuf_snapshot_recovery_response_proto_rawDescOnce sync.Once
	file_protobuf_snapshot_recovery_response_proto_rawDescData = file_protobuf_snapshot_recovery_response_proto_rawDesc
)

func file_protobuf_snapshot_recovery_response_proto_rawDescGZIP() []byte {
	file_protobuf_snapshot_recovery_response_proto_rawDescOnce.Do(func() {
		file_protobuf_snapshot_recovery_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_snapshot_recovery_response_proto_rawDescData)
	})
	return file_protobuf_snapshot_recovery_response_proto_rawDescData
}

var file_protobuf_snapshot_recovery_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_snapshot_recovery_response_proto_goTypes = []interface{}{
	(*SnapshotRecoveryResponse)(nil), // 0: protobuf.SnapshotRecoveryResponse
}
var file_protobuf_snapshot_recovery_response_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_snapshot_recovery_response_proto_init() }
func file_protobuf_snapshot_recovery_response_proto_init() {
	if File_protobuf_snapshot_recovery_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_snapshot_recovery_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotRecoveryResponse); i {
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
			RawDescriptor: file_protobuf_snapshot_recovery_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_snapshot_recovery_response_proto_goTypes,
		DependencyIndexes: file_protobuf_snapshot_recovery_response_proto_depIdxs,
		MessageInfos:      file_protobuf_snapshot_recovery_response_proto_msgTypes,
	}.Build()
	File_protobuf_snapshot_recovery_response_proto = out.File
	file_protobuf_snapshot_recovery_response_proto_rawDesc = nil
	file_protobuf_snapshot_recovery_response_proto_goTypes = nil
	file_protobuf_snapshot_recovery_response_proto_depIdxs = nil
}
