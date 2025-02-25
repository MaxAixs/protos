// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: analytics.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskDoneItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	ItemId        int32                  `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskDoneItem) Reset() {
	*x = TaskDoneItem{}
	mi := &file_analytics_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskDoneItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDoneItem) ProtoMessage() {}

func (x *TaskDoneItem) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDoneItem.ProtoReflect.Descriptor instead.
func (*TaskDoneItem) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{0}
}

func (x *TaskDoneItem) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TaskDoneItem) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *TaskDoneItem) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *TaskDoneItem) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type TaskDoneItems struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*TaskDoneItem        `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskDoneItems) Reset() {
	*x = TaskDoneItems{}
	mi := &file_analytics_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskDoneItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDoneItems) ProtoMessage() {}

func (x *TaskDoneItems) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDoneItems.ProtoReflect.Descriptor instead.
func (*TaskDoneItems) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *TaskDoneItems) GetItems() []*TaskDoneItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type ServiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServiceResponse) Reset() {
	*x = ServiceResponse{}
	mi := &file_analytics_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceResponse) ProtoMessage() {}

func (x *ServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceResponse.ProtoReflect.Descriptor instead.
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CompletedTask struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	ItemId        int32                  `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Count         int32                  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompletedTask) Reset() {
	*x = CompletedTask{}
	mi := &file_analytics_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompletedTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletedTask) ProtoMessage() {}

func (x *CompletedTask) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletedTask.ProtoReflect.Descriptor instead.
func (*CompletedTask) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{3}
}

func (x *CompletedTask) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CompletedTask) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CompletedTask) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *CompletedTask) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type WeeklyCompletedTasksResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*CompletedTask       `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WeeklyCompletedTasksResponse) Reset() {
	*x = WeeklyCompletedTasksResponse{}
	mi := &file_analytics_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WeeklyCompletedTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeeklyCompletedTasksResponse) ProtoMessage() {}

func (x *WeeklyCompletedTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeeklyCompletedTasksResponse.ProtoReflect.Descriptor instead.
func (*WeeklyCompletedTasksResponse) Descriptor() ([]byte, []int) {
	return file_analytics_proto_rawDescGZIP(), []int{4}
}

func (x *WeeklyCompletedTasksResponse) GetTasks() []*CompletedTask {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_analytics_proto protoreflect.FileDescriptor

var file_analytics_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x0c, 0x54,
	0x61, 0x73, 0x6b, 0x44, 0x6f, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3e,
	0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x6f, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44,
	0x6f, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2b,
	0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6d, 0x0a, 0x0d, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4e, 0x0a, 0x1c, 0x57, 0x65,
	0x65, 0x6b, 0x6c, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x61, 0x73,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x32, 0xb3, 0x01, 0x0a, 0x0d, 0x41,
	0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x45, 0x0a, 0x0d,
	0x53, 0x61, 0x76, 0x65, 0x44, 0x6f, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x18, 0x2e,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x6f,
	0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x1a, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x18, 0x46, 0x65, 0x74, 0x63, 0x68, 0x57, 0x65, 0x65, 0x6b,
	0x6c, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x27, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x09, 0x5a, 0x07, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_analytics_proto_rawDescOnce sync.Once
	file_analytics_proto_rawDescData []byte
)

func file_analytics_proto_rawDescGZIP() []byte {
	file_analytics_proto_rawDescOnce.Do(func() {
		file_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_analytics_proto_rawDesc), len(file_analytics_proto_rawDesc)))
	})
	return file_analytics_proto_rawDescData
}

var file_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_analytics_proto_goTypes = []any{
	(*TaskDoneItem)(nil),                 // 0: analytics.TaskDoneItem
	(*TaskDoneItems)(nil),                // 1: analytics.TaskDoneItems
	(*ServiceResponse)(nil),              // 2: analytics.ServiceResponse
	(*CompletedTask)(nil),                // 3: analytics.CompletedTask
	(*WeeklyCompletedTasksResponse)(nil), // 4: analytics.WeeklyCompletedTasksResponse
	(*timestamppb.Timestamp)(nil),        // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                // 6: google.protobuf.Empty
}
var file_analytics_proto_depIdxs = []int32{
	5, // 0: analytics.TaskDoneItem.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: analytics.TaskDoneItems.items:type_name -> analytics.TaskDoneItem
	3, // 2: analytics.WeeklyCompletedTasksResponse.tasks:type_name -> analytics.CompletedTask
	1, // 3: analytics.AnalyticsData.SaveDoneTasks:input_type -> analytics.TaskDoneItems
	6, // 4: analytics.AnalyticsData.FetchWeeklyCompletedTask:input_type -> google.protobuf.Empty
	2, // 5: analytics.AnalyticsData.SaveDoneTasks:output_type -> analytics.ServiceResponse
	4, // 6: analytics.AnalyticsData.FetchWeeklyCompletedTask:output_type -> analytics.WeeklyCompletedTasksResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_analytics_proto_init() }
func file_analytics_proto_init() {
	if File_analytics_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_analytics_proto_rawDesc), len(file_analytics_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analytics_proto_goTypes,
		DependencyIndexes: file_analytics_proto_depIdxs,
		MessageInfos:      file_analytics_proto_msgTypes,
	}.Build()
	File_analytics_proto = out.File
	file_analytics_proto_goTypes = nil
	file_analytics_proto_depIdxs = nil
}
