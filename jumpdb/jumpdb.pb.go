// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jumpdb.proto

/*
Package jumpdb is a generated protocol buffer package.

It is generated from these files:
	jumpdb.proto

It has these top-level messages:
	Database
*/
package jumpdb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Database struct {
	Version    int32            `protobuf:"varint,1,opt,name=Version,json=version" json:"Version,omitempty"`
	PathWeight map[string]int64 `protobuf:"bytes,2,rep,name=PathWeight,json=pathWeight" json:"PathWeight,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *Database) Reset()                    { *m = Database{} }
func (m *Database) String() string            { return proto.CompactTextString(m) }
func (*Database) ProtoMessage()               {}
func (*Database) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Database) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Database) GetPathWeight() map[string]int64 {
	if m != nil {
		return m.PathWeight
	}
	return nil
}

func init() {
	proto.RegisterType((*Database)(nil), "Database")
}

func init() { proto.RegisterFile("jumpdb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x2a, 0xcd, 0x2d,
	0x48, 0x49, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x9a, 0xc7, 0xc8, 0xc5, 0xe1, 0x92, 0x58,
	0x92, 0x98, 0x94, 0x58, 0x9c, 0x2a, 0x24, 0xc1, 0xc5, 0x1e, 0x96, 0x5a, 0x54, 0x9c, 0x99, 0x9f,
	0x27, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1a, 0xc4, 0x5e, 0x06, 0xe1, 0x0a, 0x59, 0x72, 0x71, 0x05,
	0x24, 0x96, 0x64, 0x84, 0xa7, 0x66, 0xa6, 0x67, 0x94, 0x48, 0x30, 0x29, 0x30, 0x6b, 0x70, 0x1b,
	0x49, 0xea, 0xc1, 0x34, 0xea, 0x21, 0xe4, 0x5c, 0xf3, 0x4a, 0x8a, 0x2a, 0x83, 0xb8, 0x0a, 0xe0,
	0x02, 0x52, 0xb6, 0x5c, 0xfc, 0x68, 0xd2, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x60, 0x3b,
	0x38, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0x05,
	0x46, 0x0d, 0xe6, 0x20, 0x08, 0xc7, 0x8a, 0xc9, 0x82, 0x31, 0x89, 0x0d, 0xec, 0x4e, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x84, 0x11, 0x6b, 0xb7, 0x00, 0x00, 0x00,
}
