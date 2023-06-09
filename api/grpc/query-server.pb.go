// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: query-server.proto

package grpc

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

type StockAggregate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType string  `protobuf:"bytes,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	Average   float32 `protobuf:"fixed32,2,opt,name=average,proto3" json:"average,omitempty"`
	Min       float32 `protobuf:"fixed32,3,opt,name=min,proto3" json:"min,omitempty"`
	Max       float32 `protobuf:"fixed32,4,opt,name=max,proto3" json:"max,omitempty"`
	Start     float32 `protobuf:"fixed32,5,opt,name=start,proto3" json:"start,omitempty"`
	End       float32 `protobuf:"fixed32,6,opt,name=end,proto3" json:"end,omitempty"`
	StartTime int64   `protobuf:"varint,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   int64   `protobuf:"varint,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *StockAggregate) Reset() {
	*x = StockAggregate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockAggregate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockAggregate) ProtoMessage() {}

func (x *StockAggregate) ProtoReflect() protoreflect.Message {
	mi := &file_query_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockAggregate.ProtoReflect.Descriptor instead.
func (*StockAggregate) Descriptor() ([]byte, []int) {
	return file_query_server_proto_rawDescGZIP(), []int{0}
}

func (x *StockAggregate) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *StockAggregate) GetAverage() float32 {
	if x != nil {
		return x.Average
	}
	return 0
}

func (x *StockAggregate) GetMin() float32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *StockAggregate) GetMax() float32 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *StockAggregate) GetStart() float32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *StockAggregate) GetEnd() float32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *StockAggregate) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *StockAggregate) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type StockFetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockName string `protobuf:"bytes,1,opt,name=stock_name,json=stockName,proto3" json:"stock_name,omitempty"`
}

func (x *StockFetchRequest) Reset() {
	*x = StockFetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockFetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockFetchRequest) ProtoMessage() {}

func (x *StockFetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockFetchRequest.ProtoReflect.Descriptor instead.
func (*StockFetchRequest) Descriptor() ([]byte, []int) {
	return file_query_server_proto_rawDescGZIP(), []int{1}
}

func (x *StockFetchRequest) GetStockName() string {
	if x != nil {
		return x.StockName
	}
	return ""
}

type StockFetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stock []*StockAggregate `protobuf:"bytes,1,rep,name=stock,proto3" json:"stock,omitempty"`
}

func (x *StockFetchResponse) Reset() {
	*x = StockFetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockFetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockFetchResponse) ProtoMessage() {}

func (x *StockFetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockFetchResponse.ProtoReflect.Descriptor instead.
func (*StockFetchResponse) Descriptor() ([]byte, []int) {
	return file_query_server_proto_rawDescGZIP(), []int{2}
}

func (x *StockFetchResponse) GetStock() []*StockAggregate {
	if x != nil {
		return x.Stock
	}
	return nil
}

var File_query_server_proto protoreflect.FileDescriptor

var file_query_server_proto_rawDesc = []byte{
	0x0a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0xcf, 0x01, 0x0a, 0x0e, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x61, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x65, 0x6e,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x11, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x3f, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x32, 0x55, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x45, 0x0a, 0x0e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x41, 0x67,
	0x67, 0x73, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_query_server_proto_rawDescOnce sync.Once
	file_query_server_proto_rawDescData = file_query_server_proto_rawDesc
)

func file_query_server_proto_rawDescGZIP() []byte {
	file_query_server_proto_rawDescOnce.Do(func() {
		file_query_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_query_server_proto_rawDescData)
	})
	return file_query_server_proto_rawDescData
}

var file_query_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_query_server_proto_goTypes = []interface{}{
	(*StockAggregate)(nil),     // 0: api.StockAggregate
	(*StockFetchRequest)(nil),  // 1: api.StockFetchRequest
	(*StockFetchResponse)(nil), // 2: api.StockFetchResponse
}
var file_query_server_proto_depIdxs = []int32{
	0, // 0: api.StockFetchResponse.stock:type_name -> api.StockAggregate
	1, // 1: api.StockFetcher.FetchStockAggs:input_type -> api.StockFetchRequest
	2, // 2: api.StockFetcher.FetchStockAggs:output_type -> api.StockFetchResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_query_server_proto_init() }
func file_query_server_proto_init() {
	if File_query_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_query_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockAggregate); i {
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
		file_query_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockFetchRequest); i {
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
		file_query_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockFetchResponse); i {
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
			RawDescriptor: file_query_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_query_server_proto_goTypes,
		DependencyIndexes: file_query_server_proto_depIdxs,
		MessageInfos:      file_query_server_proto_msgTypes,
	}.Build()
	File_query_server_proto = out.File
	file_query_server_proto_rawDesc = nil
	file_query_server_proto_goTypes = nil
	file_query_server_proto_depIdxs = nil
}
