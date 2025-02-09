// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: payment_messages.proto

package protobuf

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

type AddPaymentMethodRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	UserId          string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CardNumber      string                 `protobuf:"bytes,2,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	CardholderName  string                 `protobuf:"bytes,3,opt,name=cardholder_name,json=cardholderName,proto3" json:"cardholder_name,omitempty"`
	ExpirationMonth int32                  `protobuf:"varint,4,opt,name=expiration_month,json=expirationMonth,proto3" json:"expiration_month,omitempty"`
	ExpirationYear  int32                  `protobuf:"varint,5,opt,name=expiration_year,json=expirationYear,proto3" json:"expiration_year,omitempty"`
	Cvv             string                 `protobuf:"bytes,6,opt,name=cvv,proto3" json:"cvv,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *AddPaymentMethodRequest) Reset() {
	*x = AddPaymentMethodRequest{}
	mi := &file_payment_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddPaymentMethodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPaymentMethodRequest) ProtoMessage() {}

func (x *AddPaymentMethodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPaymentMethodRequest.ProtoReflect.Descriptor instead.
func (*AddPaymentMethodRequest) Descriptor() ([]byte, []int) {
	return file_payment_messages_proto_rawDescGZIP(), []int{0}
}

func (x *AddPaymentMethodRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddPaymentMethodRequest) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *AddPaymentMethodRequest) GetCardholderName() string {
	if x != nil {
		return x.CardholderName
	}
	return ""
}

func (x *AddPaymentMethodRequest) GetExpirationMonth() int32 {
	if x != nil {
		return x.ExpirationMonth
	}
	return 0
}

func (x *AddPaymentMethodRequest) GetExpirationYear() int32 {
	if x != nil {
		return x.ExpirationYear
	}
	return 0
}

func (x *AddPaymentMethodRequest) GetCvv() string {
	if x != nil {
		return x.Cvv
	}
	return ""
}

type AddPaymentMethodResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	PaymentMethodId string                 `protobuf:"bytes,1,opt,name=payment_method_id,json=paymentMethodId,proto3" json:"payment_method_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *AddPaymentMethodResponse) Reset() {
	*x = AddPaymentMethodResponse{}
	mi := &file_payment_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddPaymentMethodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPaymentMethodResponse) ProtoMessage() {}

func (x *AddPaymentMethodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPaymentMethodResponse.ProtoReflect.Descriptor instead.
func (*AddPaymentMethodResponse) Descriptor() ([]byte, []int) {
	return file_payment_messages_proto_rawDescGZIP(), []int{1}
}

func (x *AddPaymentMethodResponse) GetPaymentMethodId() string {
	if x != nil {
		return x.PaymentMethodId
	}
	return ""
}

type ProcessPaymentRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	UserId          string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PaymentMethodId string                 `protobuf:"bytes,2,opt,name=payment_method_id,json=paymentMethodId,proto3" json:"payment_method_id,omitempty"`
	Amount          float64                `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency        string                 `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ProcessPaymentRequest) Reset() {
	*x = ProcessPaymentRequest{}
	mi := &file_payment_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessPaymentRequest) ProtoMessage() {}

func (x *ProcessPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessPaymentRequest.ProtoReflect.Descriptor instead.
func (*ProcessPaymentRequest) Descriptor() ([]byte, []int) {
	return file_payment_messages_proto_rawDescGZIP(), []int{2}
}

func (x *ProcessPaymentRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ProcessPaymentRequest) GetPaymentMethodId() string {
	if x != nil {
		return x.PaymentMethodId
	}
	return ""
}

func (x *ProcessPaymentRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ProcessPaymentRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type ProcessPaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TransactionId string                 `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessPaymentResponse) Reset() {
	*x = ProcessPaymentResponse{}
	mi := &file_payment_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessPaymentResponse) ProtoMessage() {}

func (x *ProcessPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessPaymentResponse.ProtoReflect.Descriptor instead.
func (*ProcessPaymentResponse) Descriptor() ([]byte, []int) {
	return file_payment_messages_proto_rawDescGZIP(), []int{3}
}

func (x *ProcessPaymentResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *ProcessPaymentResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_payment_messages_proto protoreflect.FileDescriptor

var file_payment_messages_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xe2, 0x01, 0x0a, 0x17, 0x41, 0x64, 0x64,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x61, 0x72, 0x64, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x61, 0x72, 0x64, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x76, 0x76, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x76, 0x76, 0x22, 0x46, 0x0a,
	0x18, 0x41, 0x64, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x49, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x57, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0xc1, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x14, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6d, 0x61,
	0x71, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x61, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xa2, 0x02, 0x03,
	0x41, 0x56, 0x50, 0xaa, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0xca, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0xe2, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x10, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_messages_proto_rawDescOnce sync.Once
	file_payment_messages_proto_rawDescData = file_payment_messages_proto_rawDesc
)

func file_payment_messages_proto_rawDescGZIP() []byte {
	file_payment_messages_proto_rawDescOnce.Do(func() {
		file_payment_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_messages_proto_rawDescData)
	})
	return file_payment_messages_proto_rawDescData
}

var file_payment_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_payment_messages_proto_goTypes = []any{
	(*AddPaymentMethodRequest)(nil),  // 0: api.v1.payment.AddPaymentMethodRequest
	(*AddPaymentMethodResponse)(nil), // 1: api.v1.payment.AddPaymentMethodResponse
	(*ProcessPaymentRequest)(nil),    // 2: api.v1.payment.ProcessPaymentRequest
	(*ProcessPaymentResponse)(nil),   // 3: api.v1.payment.ProcessPaymentResponse
}
var file_payment_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_payment_messages_proto_init() }
func file_payment_messages_proto_init() {
	if File_payment_messages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payment_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payment_messages_proto_goTypes,
		DependencyIndexes: file_payment_messages_proto_depIdxs,
		MessageInfos:      file_payment_messages_proto_msgTypes,
	}.Build()
	File_payment_messages_proto = out.File
	file_payment_messages_proto_rawDesc = nil
	file_payment_messages_proto_goTypes = nil
	file_payment_messages_proto_depIdxs = nil
}
