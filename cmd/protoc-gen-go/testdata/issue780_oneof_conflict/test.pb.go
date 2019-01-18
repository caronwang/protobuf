// Code generated by protoc-gen-go. DO NOT EDIT.
// source: issue780_oneof_conflict/test.proto

package oneoftest

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

type Foo struct {
	// Types that are valid to be assigned to Bar:
	//	*Foo_GetBar
	Bar                  isFoo_Bar `protobuf_oneof:"bar"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Foo) ProtoReflect() protoreflect.Message {
	return xxx_Test_protoFile_messageTypes[0].MessageOf(m)
}
func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_48462cafc802a68e_gzipped, []int{0}
}

func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

type isFoo_Bar interface {
	isFoo_Bar()
}

type Foo_GetBar struct {
	GetBar string `protobuf:"bytes,1,opt,name=get_bar,json=getBar,oneof"`
}

func (*Foo_GetBar) isFoo_Bar() {}

func (m *Foo) GetBar() isFoo_Bar {
	if m != nil {
		return m.Bar
	}
	return nil
}

func (m *Foo) GetGetBar() string {
	if x, ok := m.GetBar().(*Foo_GetBar); ok {
		return x.GetBar
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Foo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Foo_GetBar)(nil),
	}
}

func init() {
	proto.RegisterFile("issue780_oneof_conflict/test.proto", fileDescriptor_48462cafc802a68e_gzipped)
	proto.RegisterType((*Foo)(nil), "oneoftest.Foo")
}

var fileDescriptor_48462cafc802a68e = []byte{
	// 88 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x22, 0x69, 0x73, 0x73, 0x75, 0x65, 0x37, 0x38, 0x30, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x74, 0x65, 0x73, 0x74, 0x22,
	0x27, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x19, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x61,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x67, 0x65, 0x74, 0x42, 0x61,
	0x72, 0x42, 0x05, 0x0a, 0x03, 0x62, 0x61, 0x72,
}

var fileDescriptor_48462cafc802a68e_gzipped = func() []byte {
	bb := new(bytes.Buffer)
	zw, _ := gzip.NewWriterLevel(bb, gzip.NoCompression)
	zw.Write(fileDescriptor_48462cafc802a68e)
	zw.Close()
	return bb.Bytes()
}()

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var Test_protoFile protoreflect.FileDescriptor

var xxx_Test_protoFile_messageTypes [1]protoimpl.MessageType
var xxx_Test_protoFile_goTypes = []interface{}{
	(*Foo)(nil), // 0: oneoftest.Foo
}
var xxx_Test_protoFile_depIdxs = []int32{}

func init() {
	var messageTypes [1]protoreflect.MessageType
	Test_protoFile = protoimpl.FileBuilder{
		RawDescriptor:      fileDescriptor_48462cafc802a68e,
		GoTypes:            xxx_Test_protoFile_goTypes,
		DependencyIndexes:  xxx_Test_protoFile_depIdxs,
		MessageOutputTypes: messageTypes[:],
	}.Init()
	messageGoTypes := xxx_Test_protoFile_goTypes[0:][:1]
	for i, mt := range messageTypes[:] {
		xxx_Test_protoFile_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_Test_protoFile_messageTypes[i].PBType = mt
	}
	xxx_Test_protoFile_goTypes = nil
	xxx_Test_protoFile_depIdxs = nil
}
