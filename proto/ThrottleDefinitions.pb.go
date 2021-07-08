// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: proto/ThrottleDefinitions.proto

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

// A set of operations which should be collectively throttled at a given milli-ops-per-second limit.
type ThrottleGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operations     []HederaFunctionality `protobuf:"varint,1,rep,packed,name=operations,proto3,enum=proto.HederaFunctionality" json:"operations,omitempty"` // The operations to be throttled
	MilliOpsPerSec uint64                `protobuf:"varint,2,opt,name=milliOpsPerSec,proto3" json:"milliOpsPerSec,omitempty"`                               // The number of total operations per second across the entire network, multiplied by 1000. So, to choose 3 operations per second (which on a network of 30 nodes is a tenth of an operation per second for each node), set milliOpsPerSec = 3000. And to choose 3.6 ops per second, use milliOpsPerSec = 3600. Minimum allowed value is 1, and maximum allowed value is 9223372.
}

func (x *ThrottleGroup) Reset() {
	*x = ThrottleGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ThrottleDefinitions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThrottleGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThrottleGroup) ProtoMessage() {}

func (x *ThrottleGroup) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ThrottleDefinitions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThrottleGroup.ProtoReflect.Descriptor instead.
func (*ThrottleGroup) Descriptor() ([]byte, []int) {
	return file_proto_ThrottleDefinitions_proto_rawDescGZIP(), []int{0}
}

func (x *ThrottleGroup) GetOperations() []HederaFunctionality {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *ThrottleGroup) GetMilliOpsPerSec() uint64 {
	if x != nil {
		return x.MilliOpsPerSec
	}
	return 0
}

// A list of throttle groups that should all compete for the same internal bucket.
type ThrottleBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                     // A name for this bucket (primarily for use in logs)
	BurstPeriodMs  uint64           `protobuf:"varint,2,opt,name=burstPeriodMs,proto3" json:"burstPeriodMs,omitempty"`  // The number of milliseconds required for this bucket to drain completely when full. The product of this number and the least common multiple of the milliOpsPerSec values in this bucket must not exceed 9223372036.
	ThrottleGroups []*ThrottleGroup `protobuf:"bytes,3,rep,name=throttleGroups,proto3" json:"throttleGroups,omitempty"` // The throttle groups competing for this bucket
}

func (x *ThrottleBucket) Reset() {
	*x = ThrottleBucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ThrottleDefinitions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThrottleBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThrottleBucket) ProtoMessage() {}

func (x *ThrottleBucket) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ThrottleDefinitions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThrottleBucket.ProtoReflect.Descriptor instead.
func (*ThrottleBucket) Descriptor() ([]byte, []int) {
	return file_proto_ThrottleDefinitions_proto_rawDescGZIP(), []int{1}
}

func (x *ThrottleBucket) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ThrottleBucket) GetBurstPeriodMs() uint64 {
	if x != nil {
		return x.BurstPeriodMs
	}
	return 0
}

func (x *ThrottleBucket) GetThrottleGroups() []*ThrottleGroup {
	if x != nil {
		return x.ThrottleGroups
	}
	return nil
}

// A list of throttle buckets which, simultaneously enforced, define the system's throttling policy.
//<ol>
//<li> When an operation appears in more than one throttling bucket, all its buckets must have room
//or it will be throttled.</li>
//<li>An operation assigned to no buckets is always throttled.</li>
//</ol>
type ThrottleDefinitions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThrottleBuckets []*ThrottleBucket `protobuf:"bytes,1,rep,name=throttleBuckets,proto3" json:"throttleBuckets,omitempty"`
}

func (x *ThrottleDefinitions) Reset() {
	*x = ThrottleDefinitions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ThrottleDefinitions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThrottleDefinitions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThrottleDefinitions) ProtoMessage() {}

func (x *ThrottleDefinitions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ThrottleDefinitions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThrottleDefinitions.ProtoReflect.Descriptor instead.
func (*ThrottleDefinitions) Descriptor() ([]byte, []int) {
	return file_proto_ThrottleDefinitions_proto_rawDescGZIP(), []int{2}
}

func (x *ThrottleDefinitions) GetThrottleBuckets() []*ThrottleBucket {
	if x != nil {
		return x.ThrottleBuckets
	}
	return nil
}

var File_proto_ThrottleDefinitions_proto protoreflect.FileDescriptor

var file_proto_ThrottleDefinitions_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x73, 0x0a, 0x0d, 0x54, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x3a, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a,
	0x0e, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x4f, 0x70, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x4f, 0x70, 0x73, 0x50,
	0x65, 0x72, 0x53, 0x65, 0x63, 0x22, 0x88, 0x01, 0x0a, 0x0e, 0x54, 0x68, 0x72, 0x6f, 0x74, 0x74,
	0x6c, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x62, 0x75, 0x72, 0x73, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x4d, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0d, 0x62, 0x75, 0x72, 0x73, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x4d, 0x73, 0x12, 0x3c, 0x0a, 0x0e, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x0e, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x22, 0x56, 0x0a, 0x13, 0x54, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x0f, 0x74, 0x68, 0x72, 0x6f, 0x74,
	0x74, 0x6c, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0f, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x4b, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ThrottleDefinitions_proto_rawDescOnce sync.Once
	file_proto_ThrottleDefinitions_proto_rawDescData = file_proto_ThrottleDefinitions_proto_rawDesc
)

func file_proto_ThrottleDefinitions_proto_rawDescGZIP() []byte {
	file_proto_ThrottleDefinitions_proto_rawDescOnce.Do(func() {
		file_proto_ThrottleDefinitions_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ThrottleDefinitions_proto_rawDescData)
	})
	return file_proto_ThrottleDefinitions_proto_rawDescData
}

var file_proto_ThrottleDefinitions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_ThrottleDefinitions_proto_goTypes = []interface{}{
	(*ThrottleGroup)(nil),       // 0: proto.ThrottleGroup
	(*ThrottleBucket)(nil),      // 1: proto.ThrottleBucket
	(*ThrottleDefinitions)(nil), // 2: proto.ThrottleDefinitions
	(HederaFunctionality)(0),    // 3: proto.HederaFunctionality
}
var file_proto_ThrottleDefinitions_proto_depIdxs = []int32{
	3, // 0: proto.ThrottleGroup.operations:type_name -> proto.HederaFunctionality
	0, // 1: proto.ThrottleBucket.throttleGroups:type_name -> proto.ThrottleGroup
	1, // 2: proto.ThrottleDefinitions.throttleBuckets:type_name -> proto.ThrottleBucket
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_ThrottleDefinitions_proto_init() }
func file_proto_ThrottleDefinitions_proto_init() {
	if File_proto_ThrottleDefinitions_proto != nil {
		return
	}
	file_proto_BasicTypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_ThrottleDefinitions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThrottleGroup); i {
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
		file_proto_ThrottleDefinitions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThrottleBucket); i {
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
		file_proto_ThrottleDefinitions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThrottleDefinitions); i {
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
			RawDescriptor: file_proto_ThrottleDefinitions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ThrottleDefinitions_proto_goTypes,
		DependencyIndexes: file_proto_ThrottleDefinitions_proto_depIdxs,
		MessageInfos:      file_proto_ThrottleDefinitions_proto_msgTypes,
	}.Build()
	File_proto_ThrottleDefinitions_proto = out.File
	file_proto_ThrottleDefinitions_proto_rawDesc = nil
	file_proto_ThrottleDefinitions_proto_goTypes = nil
	file_proto_ThrottleDefinitions_proto_depIdxs = nil
}
