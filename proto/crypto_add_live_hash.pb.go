// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/crypto_add_live_hash.proto

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

//*
// A hash---presumably of some kind of credential or certificate---along with a list of keys, each
// of which may be either a primitive or a threshold key.
type LiveHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The account to which the livehash is attached
	AccountId *AccountID `protobuf:"bytes,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	//*
	// The SHA-384 hash of a credential or certificate
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	//*
	// A list of keys (primitive or threshold), all of which must sign to attach the livehash to an account, and any one of which can later delete it.
	Keys *KeyList `protobuf:"bytes,3,opt,name=keys,proto3" json:"keys,omitempty"`
	//*
	// The duration for which the livehash will remain valid
	Duration *Duration `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *LiveHash) Reset() {
	*x = LiveHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crypto_add_live_hash_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveHash) ProtoMessage() {}

func (x *LiveHash) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crypto_add_live_hash_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveHash.ProtoReflect.Descriptor instead.
func (*LiveHash) Descriptor() ([]byte, []int) {
	return file_proto_crypto_add_live_hash_proto_rawDescGZIP(), []int{0}
}

func (x *LiveHash) GetAccountId() *AccountID {
	if x != nil {
		return x.AccountId
	}
	return nil
}

func (x *LiveHash) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *LiveHash) GetKeys() *KeyList {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *LiveHash) GetDuration() *Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

//*
// At consensus, attaches the given livehash to the given account.  The hash can be deleted by the
// key controlling the account, or by any of the keys associated to the livehash.  Hence livehashes
// provide a revocation service for their implied credentials; for example, when an authority grants
// a credential to the account, the account owner will cosign with the authority (or authorities) to
// attach a hash of the credential to the account---hence proving the grant. If the credential is
// revoked, then any of the authorities may delete it (or the account owner). In this way, the
// livehash mechanism acts as a revocation service.  An account cannot have two identical livehashes
// associated. To modify the list of keys in a livehash, the livehash should first be deleted, then
// recreated with a new list of keys.
type CryptoAddLiveHashTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// A hash of some credential or certificate, along with the keys of the entities that asserted it validity
	LiveHash *LiveHash `protobuf:"bytes,3,opt,name=liveHash,proto3" json:"liveHash,omitempty"`
}

func (x *CryptoAddLiveHashTransactionBody) Reset() {
	*x = CryptoAddLiveHashTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crypto_add_live_hash_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoAddLiveHashTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoAddLiveHashTransactionBody) ProtoMessage() {}

func (x *CryptoAddLiveHashTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crypto_add_live_hash_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoAddLiveHashTransactionBody.ProtoReflect.Descriptor instead.
func (*CryptoAddLiveHashTransactionBody) Descriptor() ([]byte, []int) {
	return file_proto_crypto_add_live_hash_proto_rawDescGZIP(), []int{1}
}

func (x *CryptoAddLiveHashTransactionBody) GetLiveHash() *LiveHash {
	if x != nil {
		return x.LiveHash
	}
	return nil
}

var File_proto_crypto_add_live_hash_proto protoreflect.FileDescriptor

var file_proto_crypto_add_live_hash_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x61,
	0x64, 0x64, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x76,
	0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x2b, 0x0a,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4f, 0x0a, 0x20, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x2b,
	0x0a, 0x08, 0x6c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73,
	0x68, 0x52, 0x08, 0x6c, 0x69, 0x76, 0x65, 0x48, 0x61, 0x73, 0x68, 0x42, 0x4b, 0x0a, 0x1a, 0x63,
	0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f,
	0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_crypto_add_live_hash_proto_rawDescOnce sync.Once
	file_proto_crypto_add_live_hash_proto_rawDescData = file_proto_crypto_add_live_hash_proto_rawDesc
)

func file_proto_crypto_add_live_hash_proto_rawDescGZIP() []byte {
	file_proto_crypto_add_live_hash_proto_rawDescOnce.Do(func() {
		file_proto_crypto_add_live_hash_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_crypto_add_live_hash_proto_rawDescData)
	})
	return file_proto_crypto_add_live_hash_proto_rawDescData
}

var file_proto_crypto_add_live_hash_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_crypto_add_live_hash_proto_goTypes = []interface{}{
	(*LiveHash)(nil),                         // 0: proto.LiveHash
	(*CryptoAddLiveHashTransactionBody)(nil), // 1: proto.CryptoAddLiveHashTransactionBody
	(*AccountID)(nil),                        // 2: proto.AccountID
	(*KeyList)(nil),                          // 3: proto.KeyList
	(*Duration)(nil),                         // 4: proto.Duration
}
var file_proto_crypto_add_live_hash_proto_depIdxs = []int32{
	2, // 0: proto.LiveHash.accountId:type_name -> proto.AccountID
	3, // 1: proto.LiveHash.keys:type_name -> proto.KeyList
	4, // 2: proto.LiveHash.duration:type_name -> proto.Duration
	0, // 3: proto.CryptoAddLiveHashTransactionBody.liveHash:type_name -> proto.LiveHash
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_crypto_add_live_hash_proto_init() }
func file_proto_crypto_add_live_hash_proto_init() {
	if File_proto_crypto_add_live_hash_proto != nil {
		return
	}
	file_proto_basic_types_proto_init()
	file_proto_duration_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_crypto_add_live_hash_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveHash); i {
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
		file_proto_crypto_add_live_hash_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoAddLiveHashTransactionBody); i {
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
			RawDescriptor: file_proto_crypto_add_live_hash_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_crypto_add_live_hash_proto_goTypes,
		DependencyIndexes: file_proto_crypto_add_live_hash_proto_depIdxs,
		MessageInfos:      file_proto_crypto_add_live_hash_proto_msgTypes,
	}.Build()
	File_proto_crypto_add_live_hash_proto = out.File
	file_proto_crypto_add_live_hash_proto_rawDesc = nil
	file_proto_crypto_add_live_hash_proto_goTypes = nil
	file_proto_crypto_add_live_hash_proto_depIdxs = nil
}
