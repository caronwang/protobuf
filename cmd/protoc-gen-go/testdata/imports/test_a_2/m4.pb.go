// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_a_2/m4.proto

package test_a_2

import (
	bytes "bytes"
	gzip "compress/gzip"
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	reflect "reflect"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type M4 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M4) ProtoReflect() protoreflect.Message {
	return xxx_M4_protoFile_messageTypes[0].MessageOf(m)
}
func (m *M4) Reset()         { *m = M4{} }
func (m *M4) String() string { return proto.CompactTextString(m) }
func (*M4) ProtoMessage()    {}
func (*M4) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdd24f82f6c5a786_gzipped, []int{0}
}

func (m *M4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M4.Unmarshal(m, b)
}
func (m *M4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M4.Marshal(b, m, deterministic)
}
func (m *M4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M4.Merge(m, src)
}
func (m *M4) XXX_Size() int {
	return xxx_messageInfo_M4.Size(m)
}
func (m *M4) XXX_DiscardUnknown() {
	xxx_messageInfo_M4.DiscardUnknown(m)
}

var xxx_messageInfo_M4 proto.InternalMessageInfo

func init() {
	proto.RegisterFile("imports/test_a_2/m4.proto", fileDescriptor_fdd24f82f6c5a786_gzipped)
	proto.RegisterType((*M4)(nil), "test.a.M4")
}

var fileDescriptor_fdd24f82f6c5a786 = []byte{
	// 119 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x19, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61,
	0x5f, 0x32, 0x2f, 0x6d, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x61, 0x22, 0x04, 0x0a, 0x02, 0x4d, 0x34, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x5f, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var fileDescriptor_fdd24f82f6c5a786_gzipped = func() []byte {
	bb := new(bytes.Buffer)
	zw, _ := gzip.NewWriterLevel(bb, gzip.NoCompression)
	zw.Write(fileDescriptor_fdd24f82f6c5a786)
	zw.Close()
	return bb.Bytes()
}()

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var M4_protoFile protoreflect.FileDescriptor

var xxx_M4_protoFile_messageTypes [1]protoimpl.MessageType
var xxx_M4_protoFile_goTypes = []interface{}{
	(*M4)(nil), // 0: test.a.M4
}
var xxx_M4_protoFile_depIdxs = []int32{}

func init() {
	var messageTypes [1]protoreflect.MessageType
	M4_protoFile = protoimpl.FileBuilder{
		RawDescriptor:      fileDescriptor_fdd24f82f6c5a786,
		GoTypes:            xxx_M4_protoFile_goTypes,
		DependencyIndexes:  xxx_M4_protoFile_depIdxs,
		MessageOutputTypes: messageTypes[:],
	}.Init()
	messageGoTypes := xxx_M4_protoFile_goTypes[0:][:1]
	for i, mt := range messageTypes[:] {
		xxx_M4_protoFile_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_M4_protoFile_messageTypes[i].PBType = mt
	}
	xxx_M4_protoFile_goTypes = nil
	xxx_M4_protoFile_depIdxs = nil
}
