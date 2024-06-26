// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.3
// source: api/report.proto

package api

import (
	_ "github.com/gogo/protobuf/gogoproto"
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

// See https://github.com/botuniverse/onebot-11/blob/master/event/README.md
type ReportEventReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 事件发生的时间戳
	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	// 收到事件的机器人 QQ 号
	SelfId int64 `protobuf:"varint,2,opt,name=self_id,json=selfId,proto3" json:"self_id,omitempty"`
	// 上报类型，可能的值为 message、
	PostType string `protobuf:"bytes,3,opt,name=post_type,json=postType,proto3" json:"post_type,omitempty"`
	// 消息类型， 可能的值为 private、group
	MessageType string `protobuf:"bytes,4,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	// 消息子类型
	// 在 message_type 取 private 时，如果是好友则是 friend，如果是群临时会话则是 group
	// 在 message_type 取 group 时，正常消息是 normal，匿名消息是 anonymous，系统提示（如「管理员已禁止群内匿名聊天」）是 notice
	SubType string `protobuf:"bytes,5,opt,name=sub_type,json=subType,proto3" json:"sub_type,omitempty"`
	// 消息 ID
	MessageId int64 `protobuf:"varint,6,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// 群号
	// 仅当 message_type 取 group 时有效
	GroupId int64 `protobuf:"varint,7,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// 发送者 QQ 号
	UserId int64 `protobuf:"varint,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 匿名信息，如果不是匿名消息则为 null
	// 仅当 message_type 取 group 时有效
	Anonymous *ReportEventReq_Anonymous `protobuf:"bytes,9,opt,name=anonymous,proto3" json:"anonymous,omitempty"`
	// 消息内容
	Message []*Segment `protobuf:"bytes,10,rep,name=message,proto3" json:"message,omitempty"`
	// 原始消息内容
	RawMessage string `protobuf:"bytes,11,opt,name=raw_message,json=rawMessage,proto3" json:"raw_message,omitempty"`
	// 字体
	Font int64 `protobuf:"varint,12,opt,name=font,proto3" json:"font,omitempty"`
	// 发送人信息，各字段尽最大努力提供
	Sender *ReportEventReq_Sender `protobuf:"bytes,13,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (x *ReportEventReq) Reset() {
	*x = ReportEventReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportEventReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportEventReq) ProtoMessage() {}

func (x *ReportEventReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportEventReq.ProtoReflect.Descriptor instead.
func (*ReportEventReq) Descriptor() ([]byte, []int) {
	return file_api_report_proto_rawDescGZIP(), []int{0}
}

func (x *ReportEventReq) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *ReportEventReq) GetSelfId() int64 {
	if x != nil {
		return x.SelfId
	}
	return 0
}

func (x *ReportEventReq) GetPostType() string {
	if x != nil {
		return x.PostType
	}
	return ""
}

func (x *ReportEventReq) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *ReportEventReq) GetSubType() string {
	if x != nil {
		return x.SubType
	}
	return ""
}

func (x *ReportEventReq) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *ReportEventReq) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *ReportEventReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReportEventReq) GetAnonymous() *ReportEventReq_Anonymous {
	if x != nil {
		return x.Anonymous
	}
	return nil
}

func (x *ReportEventReq) GetMessage() []*Segment {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *ReportEventReq) GetRawMessage() string {
	if x != nil {
		return x.RawMessage
	}
	return ""
}

func (x *ReportEventReq) GetFont() int64 {
	if x != nil {
		return x.Font
	}
	return 0
}

func (x *ReportEventReq) GetSender() *ReportEventReq_Sender {
	if x != nil {
		return x.Sender
	}
	return nil
}

type ReportEventResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 要回复的内容
	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	// 消息内容是否作为纯文本发送（即不解析 CQ 码），只在 reply 字段是字符串时有效
	AutoEscape bool `protobuf:"varint,2,opt,name=auto_escape,json=autoEscape,proto3" json:"auto_escape,omitempty"`
	// 是否要在回复开头 at 发送者（自动添加），发送者是匿名用户时无效，默认 at 发送者
	// 仅当 message_type 取 group 时有效
	AtSender bool `protobuf:"varint,3,opt,name=at_sender,json=atSender,proto3" json:"at_sender,omitempty"`
	// 撤回该条消息，默认不撤回
	// 仅当 message_type 取 group 时有效
	Delete bool `protobuf:"varint,4,opt,name=delete,proto3" json:"delete,omitempty"`
	// 把发送者踢出群组（需要登录号权限足够），不拒绝此人后续加群请求，发送者是匿名用户时无效，默认不踢
	// 仅当 message_type 取 group 时有效
	Kick bool `protobuf:"varint,5,opt,name=kick,proto3" json:"kick,omitempty"`
	// 把发送者禁言 ban_duration 指定时长，对匿名用户也有效，默认不禁言
	// 仅当 message_type 取 group 时有效
	Ban bool `protobuf:"varint,6,opt,name=ban,proto3" json:"ban,omitempty"`
	// 禁言时长，默认30分钟
	// 仅当 message_type 取 group 时有效
	BanDuration int64 `protobuf:"varint,7,opt,name=ban_duration,json=banDuration,proto3" json:"ban_duration,omitempty"` // ---- message quick reply end ----
}

func (x *ReportEventResp) Reset() {
	*x = ReportEventResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportEventResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportEventResp) ProtoMessage() {}

func (x *ReportEventResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportEventResp.ProtoReflect.Descriptor instead.
func (*ReportEventResp) Descriptor() ([]byte, []int) {
	return file_api_report_proto_rawDescGZIP(), []int{1}
}

func (x *ReportEventResp) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

func (x *ReportEventResp) GetAutoEscape() bool {
	if x != nil {
		return x.AutoEscape
	}
	return false
}

func (x *ReportEventResp) GetAtSender() bool {
	if x != nil {
		return x.AtSender
	}
	return false
}

func (x *ReportEventResp) GetDelete() bool {
	if x != nil {
		return x.Delete
	}
	return false
}

func (x *ReportEventResp) GetKick() bool {
	if x != nil {
		return x.Kick
	}
	return false
}

func (x *ReportEventResp) GetBan() bool {
	if x != nil {
		return x.Ban
	}
	return false
}

func (x *ReportEventResp) GetBanDuration() int64 {
	if x != nil {
		return x.BanDuration
	}
	return 0
}

type ReportEventReq_Sender struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 发送者 QQ 号
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 昵称
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// 性别，male 或 female 或 unknown
	Sex int64 `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`
	// 年龄
	Age int64 `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	// 群名片／备注
	// 仅当 message_type 取 group 时有效
	Card string `protobuf:"bytes,5,opt,name=card,proto3" json:"card,omitempty"`
	// 地区
	// 仅当 message_type 取 group 时有效
	Area string `protobuf:"bytes,6,opt,name=area,proto3" json:"area,omitempty"`
	// 成员等级
	// 仅当 message_type 取 group 时有效
	Level string `protobuf:"bytes,7,opt,name=level,proto3" json:"level,omitempty"`
	// 角色，owner 或 admin 或 member
	// 仅当 message_type 取 group 时有效
	Role string `protobuf:"bytes,8,opt,name=role,proto3" json:"role,omitempty"`
	// 专属头衔
	// 仅当 message_type 取 group 时有效
	Title string `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *ReportEventReq_Sender) Reset() {
	*x = ReportEventReq_Sender{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportEventReq_Sender) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportEventReq_Sender) ProtoMessage() {}

func (x *ReportEventReq_Sender) ProtoReflect() protoreflect.Message {
	mi := &file_api_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportEventReq_Sender.ProtoReflect.Descriptor instead.
func (*ReportEventReq_Sender) Descriptor() ([]byte, []int) {
	return file_api_report_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ReportEventReq_Sender) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReportEventReq_Sender) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ReportEventReq_Sender) GetSex() int64 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *ReportEventReq_Sender) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *ReportEventReq_Sender) GetCard() string {
	if x != nil {
		return x.Card
	}
	return ""
}

func (x *ReportEventReq_Sender) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *ReportEventReq_Sender) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *ReportEventReq_Sender) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ReportEventReq_Sender) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ReportEventReq_Anonymous struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 匿名用户 ID
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 匿名用户名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 匿名用户 flag，在调用禁言 API 时需要传入
	Flag string `protobuf:"bytes,3,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *ReportEventReq_Anonymous) Reset() {
	*x = ReportEventReq_Anonymous{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_report_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportEventReq_Anonymous) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportEventReq_Anonymous) ProtoMessage() {}

func (x *ReportEventReq_Anonymous) ProtoReflect() protoreflect.Message {
	mi := &file_api_report_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportEventReq_Anonymous.ProtoReflect.Descriptor instead.
func (*ReportEventReq_Anonymous) Descriptor() ([]byte, []int) {
	return file_api_report_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ReportEventReq_Anonymous) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReportEventReq_Anonymous) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReportEventReq_Anonymous) GetFlag() string {
	if x != nil {
		return x.Flag
	}
	return ""
}

var File_api_report_proto protoreflect.FileDescriptor

var file_api_report_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x71, 0x71, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x2d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x05,
	0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x66, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x09, 0x61,
	0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x71, 0x71, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x2e, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x6f, 0x75, 0x73, 0x52, 0x09, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x12, 0x2c,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x71, 0x71, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x61, 0x77, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x61, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x6f, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x6f, 0x6e,
	0x74, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x71, 0x71, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x1a, 0xc9, 0x01, 0x0a, 0x06,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x61, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x1a, 0x43, 0x0a, 0x09, 0x41, 0x6e, 0x6f, 0x6e, 0x79,
	0x6d, 0x6f, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x22, 0xc6, 0x01, 0x0a,
	0x0f, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x65,
	0x73, 0x63, 0x61, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x75, 0x74,
	0x6f, 0x45, 0x73, 0x63, 0x61, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x74, 0x5f, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x74, 0x53, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6b, 0x69, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6b, 0x69, 0x63, 0x6b,
	0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x62,
	0x61, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x6e, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x61, 0x6e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x62, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x58, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19,
	0x2e, 0x71, 0x71, 0x62, 0x6f, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x71, 0x71, 0x62, 0x6f,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x07, 0x2f,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x0b, 0x5a, 0x09, 0x71, 0x71, 0x62,
	0x6f, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_report_proto_rawDescOnce sync.Once
	file_api_report_proto_rawDescData = file_api_report_proto_rawDesc
)

func file_api_report_proto_rawDescGZIP() []byte {
	file_api_report_proto_rawDescOnce.Do(func() {
		file_api_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_report_proto_rawDescData)
	})
	return file_api_report_proto_rawDescData
}

var file_api_report_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_report_proto_goTypes = []interface{}{
	(*ReportEventReq)(nil),           // 0: qqbot.api.ReportEventReq
	(*ReportEventResp)(nil),          // 1: qqbot.api.ReportEventResp
	(*ReportEventReq_Sender)(nil),    // 2: qqbot.api.ReportEventReq.Sender
	(*ReportEventReq_Anonymous)(nil), // 3: qqbot.api.ReportEventReq.Anonymous
	(*Segment)(nil),                  // 4: qqbot.api.Segment
}
var file_api_report_proto_depIdxs = []int32{
	3, // 0: qqbot.api.ReportEventReq.anonymous:type_name -> qqbot.api.ReportEventReq.Anonymous
	4, // 1: qqbot.api.ReportEventReq.message:type_name -> qqbot.api.Segment
	2, // 2: qqbot.api.ReportEventReq.sender:type_name -> qqbot.api.ReportEventReq.Sender
	0, // 3: qqbot.api.Report.ReportEvent:input_type -> qqbot.api.ReportEventReq
	1, // 4: qqbot.api.Report.ReportEvent:output_type -> qqbot.api.ReportEventResp
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_report_proto_init() }
func file_api_report_proto_init() {
	if File_api_report_proto != nil {
		return
	}
	file_api_segment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportEventReq); i {
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
		file_api_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportEventResp); i {
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
		file_api_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportEventReq_Sender); i {
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
		file_api_report_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportEventReq_Anonymous); i {
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
			RawDescriptor: file_api_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_report_proto_goTypes,
		DependencyIndexes: file_api_report_proto_depIdxs,
		MessageInfos:      file_api_report_proto_msgTypes,
	}.Build()
	File_api_report_proto = out.File
	file_api_report_proto_rawDesc = nil
	file_api_report_proto_goTypes = nil
	file_api_report_proto_depIdxs = nil
}
