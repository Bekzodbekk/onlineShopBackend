// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: protos/dashboard.proto

package dashboardpb

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

type Dashboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total       string `protobuf:"bytes,1,opt,name=total,proto3" json:"total,omitempty"`                                 //jami
	SoldOut     int32  `protobuf:"varint,2,opt,name=sold_out,json=soldOut,proto3" json:"sold_out,omitempty"`             //sotildi
	AddDebt     int32  `protobuf:"varint,3,opt,name=add_debt,json=addDebt,proto3" json:"add_debt,omitempty"`             // qarz berildi
	GivenDebt   int32  `protobuf:"varint,4,opt,name=given_debt,json=givenDebt,proto3" json:"given_debt,omitempty"`       // qarzni berishdi
	TotalClient int32  `protobuf:"varint,5,opt,name=total_client,json=totalClient,proto3" json:"total_client,omitempty"` //jami client
	Profit      int64  `protobuf:"varint,6,opt,name=profit,proto3" json:"profit,omitempty"`                              //foyda
	BodyPrice   int64  `protobuf:"varint,7,opt,name=body_price,json=bodyPrice,proto3" json:"body_price,omitempty"`       //tan narxi
	PaidDebt    int64  `protobuf:"varint,8,opt,name=paid_debt,json=paidDebt,proto3" json:"paid_debt,omitempty"`          //to'langan qarz
	CreatedAt   string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   int64  `protobuf:"varint,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Dashboard) Reset() {
	*x = Dashboard{}
	mi := &file_protos_dashboard_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Dashboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dashboard) ProtoMessage() {}

func (x *Dashboard) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dashboard_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dashboard.ProtoReflect.Descriptor instead.
func (*Dashboard) Descriptor() ([]byte, []int) {
	return file_protos_dashboard_proto_rawDescGZIP(), []int{0}
}

func (x *Dashboard) GetTotal() string {
	if x != nil {
		return x.Total
	}
	return ""
}

func (x *Dashboard) GetSoldOut() int32 {
	if x != nil {
		return x.SoldOut
	}
	return 0
}

func (x *Dashboard) GetAddDebt() int32 {
	if x != nil {
		return x.AddDebt
	}
	return 0
}

func (x *Dashboard) GetGivenDebt() int32 {
	if x != nil {
		return x.GivenDebt
	}
	return 0
}

func (x *Dashboard) GetTotalClient() int32 {
	if x != nil {
		return x.TotalClient
	}
	return 0
}

func (x *Dashboard) GetProfit() int64 {
	if x != nil {
		return x.Profit
	}
	return 0
}

func (x *Dashboard) GetBodyPrice() int64 {
	if x != nil {
		return x.BodyPrice
	}
	return 0
}

func (x *Dashboard) GetPaidDebt() int64 {
	if x != nil {
		return x.PaidDebt
	}
	return 0
}

func (x *Dashboard) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Dashboard) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Dashboard) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type GetDashboardInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDashboardInfoReq) Reset() {
	*x = GetDashboardInfoReq{}
	mi := &file_protos_dashboard_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDashboardInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardInfoReq) ProtoMessage() {}

func (x *GetDashboardInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dashboard_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardInfoReq.ProtoReflect.Descriptor instead.
func (*GetDashboardInfoReq) Descriptor() ([]byte, []int) {
	return file_protos_dashboard_proto_rawDescGZIP(), []int{1}
}

type GetDashboardInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    bool       `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message   string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Dashboard *Dashboard `protobuf:"bytes,3,opt,name=Dashboard,proto3" json:"Dashboard,omitempty"`
}

func (x *GetDashboardInfoResp) Reset() {
	*x = GetDashboardInfoResp{}
	mi := &file_protos_dashboard_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDashboardInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardInfoResp) ProtoMessage() {}

func (x *GetDashboardInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dashboard_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardInfoResp.ProtoReflect.Descriptor instead.
func (*GetDashboardInfoResp) Descriptor() ([]byte, []int) {
	return file_protos_dashboard_proto_rawDescGZIP(), []int{2}
}

func (x *GetDashboardInfoResp) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *GetDashboardInfoResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetDashboardInfoResp) GetDashboard() *Dashboard {
	if x != nil {
		return x.Dashboard
	}
	return nil
}

var File_protos_dashboard_proto protoreflect.FileDescriptor

var file_protos_dashboard_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x02, 0x0a, 0x09, 0x44, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x6f, 0x6c, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x73, 0x6f, 0x6c, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x5f, 0x64,
	0x65, 0x62, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x64, 0x64, 0x44, 0x65,
	0x62, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x5f, 0x64, 0x65, 0x62, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x44, 0x65, 0x62,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x6f, 0x64, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x62, 0x6f, 0x64, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x69, 0x64, 0x5f, 0x64, 0x65, 0x62, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x69, 0x64, 0x44, 0x65, 0x62, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x22, 0x72, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x44, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x09, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x32, 0x55, 0x0a, 0x10, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x15,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_dashboard_proto_rawDescOnce sync.Once
	file_protos_dashboard_proto_rawDescData = file_protos_dashboard_proto_rawDesc
)

func file_protos_dashboard_proto_rawDescGZIP() []byte {
	file_protos_dashboard_proto_rawDescOnce.Do(func() {
		file_protos_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_dashboard_proto_rawDescData)
	})
	return file_protos_dashboard_proto_rawDescData
}

var file_protos_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_dashboard_proto_goTypes = []any{
	(*Dashboard)(nil),            // 0: Dashboard
	(*GetDashboardInfoReq)(nil),  // 1: GetDashboardInfoReq
	(*GetDashboardInfoResp)(nil), // 2: GetDashboardInfoResp
}
var file_protos_dashboard_proto_depIdxs = []int32{
	0, // 0: GetDashboardInfoResp.Dashboard:type_name -> Dashboard
	1, // 1: DashboardService.GetDashboardInfo:input_type -> GetDashboardInfoReq
	2, // 2: DashboardService.GetDashboardInfo:output_type -> GetDashboardInfoResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_dashboard_proto_init() }
func file_protos_dashboard_proto_init() {
	if File_protos_dashboard_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_dashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_dashboard_proto_goTypes,
		DependencyIndexes: file_protos_dashboard_proto_depIdxs,
		MessageInfos:      file_protos_dashboard_proto_msgTypes,
	}.Build()
	File_protos_dashboard_proto = out.File
	file_protos_dashboard_proto_rawDesc = nil
	file_protos_dashboard_proto_goTypes = nil
	file_protos_dashboard_proto_depIdxs = nil
}
