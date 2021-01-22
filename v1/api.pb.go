// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: notes/v1/api.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Author    string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Text      string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	Private   bool   `protobuf:"varint,5,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_notes_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_notes_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *Note) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Note) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Note) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Note) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Note) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

type NoteFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids     []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Author  []string `protobuf:"bytes,2,rep,name=author,proto3" json:"author,omitempty"`
	Before  string   `protobuf:"bytes,3,opt,name=before,proto3" json:"before,omitempty"`
	After   string   `protobuf:"bytes,4,opt,name=after,proto3" json:"after,omitempty"`
	Private bool     `protobuf:"varint,5,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *NoteFilter) Reset() {
	*x = NoteFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteFilter) ProtoMessage() {}

func (x *NoteFilter) ProtoReflect() protoreflect.Message {
	mi := &file_notes_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteFilter.ProtoReflect.Descriptor instead.
func (*NoteFilter) Descriptor() ([]byte, []int) {
	return file_notes_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *NoteFilter) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *NoteFilter) GetAuthor() []string {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *NoteFilter) GetBefore() string {
	if x != nil {
		return x.Before
	}
	return ""
}

func (x *NoteFilter) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

func (x *NoteFilter) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

type Notebook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error  `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Notes []*Note `protobuf:"bytes,2,rep,name=notes,proto3" json:"notes,omitempty"`
}

func (x *Notebook) Reset() {
	*x = Notebook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notebook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notebook) ProtoMessage() {}

func (x *Notebook) ProtoReflect() protoreflect.Message {
	mi := &file_notes_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notebook.ProtoReflect.Descriptor instead.
func (*Notebook) Descriptor() ([]byte, []int) {
	return file_notes_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *Notebook) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Notebook) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notes_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_notes_v1_api_proto_msgTypes[3]
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
	return file_notes_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *Error) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_notes_v1_api_proto protoreflect.FileDescriptor

var file_notes_v1_api_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x04,
	0x4e, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x7e, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x57, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x65,
	0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65,
	0x73, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9f, 0x01, 0x0a, 0x0b, 0x4e, 0x6f, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x12, 0x14, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x15, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x12, 0x46, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x1a, 0x12, 0x2e, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b,
	0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0xe2, 0x01, 0x5a, 0x1d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x62, 0x65, 0x6e, 0x67, 0x66,
	0x6f, 0x72, 0x74, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x92, 0x41, 0xbf, 0x01,
	0x12, 0x94, 0x01, 0x0a, 0x05, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x09, 0x62, 0x62,
	0x65, 0x6e, 0x67, 0x66, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x62, 0x65, 0x6e,
	0x67, 0x66, 0x6f, 0x72, 0x74, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x1a, 0x11, 0x69, 0x6e, 0x66,
	0x6f, 0x40, 0x62, 0x65, 0x6e, 0x67, 0x66, 0x6f, 0x72, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x42,
	0x0a, 0x14, 0x42, 0x53, 0x44, 0x20, 0x33, 0x2d, 0x43, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x20, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x62, 0x65, 0x6e, 0x67,
	0x66, 0x6f, 0x72, 0x74, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e,
	0x53, 0x45, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notes_v1_api_proto_rawDescOnce sync.Once
	file_notes_v1_api_proto_rawDescData = file_notes_v1_api_proto_rawDesc
)

func file_notes_v1_api_proto_rawDescGZIP() []byte {
	file_notes_v1_api_proto_rawDescOnce.Do(func() {
		file_notes_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_notes_v1_api_proto_rawDescData)
	})
	return file_notes_v1_api_proto_rawDescData
}

var file_notes_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_notes_v1_api_proto_goTypes = []interface{}{
	(*Note)(nil),       // 0: notes.v1.Note
	(*NoteFilter)(nil), // 1: notes.v1.NoteFilter
	(*Notebook)(nil),   // 2: notes.v1.Notebook
	(*Error)(nil),      // 3: notes.v1.Error
}
var file_notes_v1_api_proto_depIdxs = []int32{
	3, // 0: notes.v1.Notebook.error:type_name -> notes.v1.Error
	0, // 1: notes.v1.Notebook.notes:type_name -> notes.v1.Note
	1, // 2: notes.v1.NoteService.Fetch:input_type -> notes.v1.NoteFilter
	0, // 3: notes.v1.NoteService.Create:input_type -> notes.v1.Note
	2, // 4: notes.v1.NoteService.Fetch:output_type -> notes.v1.Notebook
	2, // 5: notes.v1.NoteService.Create:output_type -> notes.v1.Notebook
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notes_v1_api_proto_init() }
func file_notes_v1_api_proto_init() {
	if File_notes_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notes_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
		file_notes_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteFilter); i {
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
		file_notes_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notebook); i {
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
		file_notes_v1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notes_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notes_v1_api_proto_goTypes,
		DependencyIndexes: file_notes_v1_api_proto_depIdxs,
		MessageInfos:      file_notes_v1_api_proto_msgTypes,
	}.Build()
	File_notes_v1_api_proto = out.File
	file_notes_v1_api_proto_rawDesc = nil
	file_notes_v1_api_proto_goTypes = nil
	file_notes_v1_api_proto_depIdxs = nil
}