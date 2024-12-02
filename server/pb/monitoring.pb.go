// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: monitoring.proto

package pb

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

type MonitoringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartMonitoring string `protobuf:"bytes,1,opt,name=start_monitoring,json=startMonitoring,proto3" json:"start_monitoring,omitempty"`
}

func (x *MonitoringRequest) Reset() {
	*x = MonitoringRequest{}
	mi := &file_monitoring_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MonitoringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitoringRequest) ProtoMessage() {}

func (x *MonitoringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitoringRequest.ProtoReflect.Descriptor instead.
func (*MonitoringRequest) Descriptor() ([]byte, []int) {
	return file_monitoring_proto_rawDescGZIP(), []int{0}
}

func (x *MonitoringRequest) GetStartMonitoring() string {
	if x != nil {
		return x.StartMonitoring
	}
	return ""
}

type MonitoringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Market              string  `protobuf:"bytes,1,opt,name=market,proto3" json:"market,omitempty"`
	TradeDate           string  `protobuf:"bytes,2,opt,name=trade_date,json=tradeDate,proto3" json:"trade_date,omitempty"`
	TradeTime           string  `protobuf:"bytes,3,opt,name=trade_time,json=tradeTime,proto3" json:"trade_time,omitempty"`
	TradeDateKst        string  `protobuf:"bytes,4,opt,name=trade_date_kst,json=tradeDateKst,proto3" json:"trade_date_kst,omitempty"`
	TradeTimeKst        string  `protobuf:"bytes,5,opt,name=trade_time_kst,json=tradeTimeKst,proto3" json:"trade_time_kst,omitempty"`
	TradeTimestamp      int64   `protobuf:"varint,6,opt,name=trade_timestamp,json=tradeTimestamp,proto3" json:"trade_timestamp,omitempty"`
	OpeningPrice        int64   `protobuf:"varint,7,opt,name=opening_price,json=openingPrice,proto3" json:"opening_price,omitempty"`
	HighPrice           int64   `protobuf:"varint,8,opt,name=high_price,json=highPrice,proto3" json:"high_price,omitempty"`
	LowPrice            int64   `protobuf:"varint,9,opt,name=low_price,json=lowPrice,proto3" json:"low_price,omitempty"`
	TradePrice          int64   `protobuf:"varint,10,opt,name=trade_price,json=tradePrice,proto3" json:"trade_price,omitempty"`
	PrevClosingPrice    int64   `protobuf:"varint,11,opt,name=prev_closing_price,json=prevClosingPrice,proto3" json:"prev_closing_price,omitempty"`
	Change              string  `protobuf:"bytes,12,opt,name=change,proto3" json:"change,omitempty"` // "FALL" 또는 "RISE"
	ChangePrice         int64   `protobuf:"varint,13,opt,name=change_price,json=changePrice,proto3" json:"change_price,omitempty"`
	ChangeRate          float64 `protobuf:"fixed64,14,opt,name=change_rate,json=changeRate,proto3" json:"change_rate,omitempty"`
	SignedChangePrice   int64   `protobuf:"varint,15,opt,name=signed_change_price,json=signedChangePrice,proto3" json:"signed_change_price,omitempty"`
	SignedChangeRate    float64 `protobuf:"fixed64,16,opt,name=signed_change_rate,json=signedChangeRate,proto3" json:"signed_change_rate,omitempty"`
	TradeVolume         float64 `protobuf:"fixed64,17,opt,name=trade_volume,json=tradeVolume,proto3" json:"trade_volume,omitempty"`
	AccTradePrice       float64 `protobuf:"fixed64,18,opt,name=acc_trade_price,json=accTradePrice,proto3" json:"acc_trade_price,omitempty"`
	AccTradePrice_24H   float64 `protobuf:"fixed64,19,opt,name=acc_trade_price_24h,json=accTradePrice24h,proto3" json:"acc_trade_price_24h,omitempty"`
	AccTradeVolume      float64 `protobuf:"fixed64,20,opt,name=acc_trade_volume,json=accTradeVolume,proto3" json:"acc_trade_volume,omitempty"`
	AccTradeVolume_24H  float64 `protobuf:"fixed64,21,opt,name=acc_trade_volume_24h,json=accTradeVolume24h,proto3" json:"acc_trade_volume_24h,omitempty"`
	Highest_52WeekPrice int64   `protobuf:"varint,22,opt,name=highest_52_week_price,json=highest52WeekPrice,proto3" json:"highest_52_week_price,omitempty"`
	Highest_52WeekDate  string  `protobuf:"bytes,23,opt,name=highest_52_week_date,json=highest52WeekDate,proto3" json:"highest_52_week_date,omitempty"`
	Lowest_52WeekPrice  int64   `protobuf:"varint,24,opt,name=lowest_52_week_price,json=lowest52WeekPrice,proto3" json:"lowest_52_week_price,omitempty"`
	Lowest_52WeekDate   string  `protobuf:"bytes,25,opt,name=lowest_52_week_date,json=lowest52WeekDate,proto3" json:"lowest_52_week_date,omitempty"`
	Timestamp           int64   `protobuf:"varint,26,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *MonitoringResponse) Reset() {
	*x = MonitoringResponse{}
	mi := &file_monitoring_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MonitoringResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitoringResponse) ProtoMessage() {}

func (x *MonitoringResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monitoring_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitoringResponse.ProtoReflect.Descriptor instead.
func (*MonitoringResponse) Descriptor() ([]byte, []int) {
	return file_monitoring_proto_rawDescGZIP(), []int{1}
}

func (x *MonitoringResponse) GetMarket() string {
	if x != nil {
		return x.Market
	}
	return ""
}

func (x *MonitoringResponse) GetTradeDate() string {
	if x != nil {
		return x.TradeDate
	}
	return ""
}

func (x *MonitoringResponse) GetTradeTime() string {
	if x != nil {
		return x.TradeTime
	}
	return ""
}

func (x *MonitoringResponse) GetTradeDateKst() string {
	if x != nil {
		return x.TradeDateKst
	}
	return ""
}

func (x *MonitoringResponse) GetTradeTimeKst() string {
	if x != nil {
		return x.TradeTimeKst
	}
	return ""
}

func (x *MonitoringResponse) GetTradeTimestamp() int64 {
	if x != nil {
		return x.TradeTimestamp
	}
	return 0
}

func (x *MonitoringResponse) GetOpeningPrice() int64 {
	if x != nil {
		return x.OpeningPrice
	}
	return 0
}

func (x *MonitoringResponse) GetHighPrice() int64 {
	if x != nil {
		return x.HighPrice
	}
	return 0
}

func (x *MonitoringResponse) GetLowPrice() int64 {
	if x != nil {
		return x.LowPrice
	}
	return 0
}

func (x *MonitoringResponse) GetTradePrice() int64 {
	if x != nil {
		return x.TradePrice
	}
	return 0
}

func (x *MonitoringResponse) GetPrevClosingPrice() int64 {
	if x != nil {
		return x.PrevClosingPrice
	}
	return 0
}

func (x *MonitoringResponse) GetChange() string {
	if x != nil {
		return x.Change
	}
	return ""
}

func (x *MonitoringResponse) GetChangePrice() int64 {
	if x != nil {
		return x.ChangePrice
	}
	return 0
}

func (x *MonitoringResponse) GetChangeRate() float64 {
	if x != nil {
		return x.ChangeRate
	}
	return 0
}

func (x *MonitoringResponse) GetSignedChangePrice() int64 {
	if x != nil {
		return x.SignedChangePrice
	}
	return 0
}

func (x *MonitoringResponse) GetSignedChangeRate() float64 {
	if x != nil {
		return x.SignedChangeRate
	}
	return 0
}

func (x *MonitoringResponse) GetTradeVolume() float64 {
	if x != nil {
		return x.TradeVolume
	}
	return 0
}

func (x *MonitoringResponse) GetAccTradePrice() float64 {
	if x != nil {
		return x.AccTradePrice
	}
	return 0
}

func (x *MonitoringResponse) GetAccTradePrice_24H() float64 {
	if x != nil {
		return x.AccTradePrice_24H
	}
	return 0
}

func (x *MonitoringResponse) GetAccTradeVolume() float64 {
	if x != nil {
		return x.AccTradeVolume
	}
	return 0
}

func (x *MonitoringResponse) GetAccTradeVolume_24H() float64 {
	if x != nil {
		return x.AccTradeVolume_24H
	}
	return 0
}

func (x *MonitoringResponse) GetHighest_52WeekPrice() int64 {
	if x != nil {
		return x.Highest_52WeekPrice
	}
	return 0
}

func (x *MonitoringResponse) GetHighest_52WeekDate() string {
	if x != nil {
		return x.Highest_52WeekDate
	}
	return ""
}

func (x *MonitoringResponse) GetLowest_52WeekPrice() int64 {
	if x != nil {
		return x.Lowest_52WeekPrice
	}
	return 0
}

func (x *MonitoringResponse) GetLowest_52WeekDate() string {
	if x != nil {
		return x.Lowest_52WeekDate
	}
	return ""
}

func (x *MonitoringResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_monitoring_proto protoreflect.FileDescriptor

var file_monitoring_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x3f,
	0x0a, 0x12, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x22,
	0x81, 0x08, 0x0a, 0x13, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x64, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x0e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x73, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x4b, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x6b, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4b, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x6e, 0x69,
	0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x68, 0x69, 0x67,
	0x68, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x6f, 0x77, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x63, 0x6c, 0x6f,
	0x73, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x70, 0x72, 0x65, 0x76, 0x43, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2e,
	0x0a, 0x13, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2c,
	0x0a, 0x12, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x61, 0x63, 0x63, 0x5f, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x32, 0x34, 0x68, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x61, 0x63, 0x63, 0x54, 0x72, 0x61, 0x64, 0x65, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x32, 0x34, 0x68, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x63, 0x63, 0x5f, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0e, 0x61, 0x63, 0x63, 0x54, 0x72, 0x61, 0x64, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x2f, 0x0a, 0x14, 0x61, 0x63, 0x63, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x32, 0x34, 0x68, 0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11,
	0x61, 0x63, 0x63, 0x54, 0x72, 0x61, 0x64, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x32, 0x34,
	0x68, 0x12, 0x31, 0x0a, 0x15, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x35, 0x32, 0x5f,
	0x77, 0x65, 0x65, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x12, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x35, 0x32, 0x57, 0x65, 0x65, 0x6b, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x5f,
	0x35, 0x32, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x35, 0x32, 0x57, 0x65, 0x65,
	0x6b, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x5f,
	0x35, 0x32, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x11, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x35, 0x32, 0x57, 0x65, 0x65,
	0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74,
	0x5f, 0x35, 0x32, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x19, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x6c, 0x6f, 0x77, 0x65, 0x73, 0x74, 0x35, 0x32, 0x57, 0x65, 0x65,
	0x6b, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x32, 0x63, 0x0a, 0x12, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x08, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1e, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monitoring_proto_rawDescOnce sync.Once
	file_monitoring_proto_rawDescData = file_monitoring_proto_rawDesc
)

func file_monitoring_proto_rawDescGZIP() []byte {
	file_monitoring_proto_rawDescOnce.Do(func() {
		file_monitoring_proto_rawDescData = protoimpl.X.CompressGZIP(file_monitoring_proto_rawDescData)
	})
	return file_monitoring_proto_rawDescData
}

var file_monitoring_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_monitoring_proto_goTypes = []any{
	(*MonitoringRequest)(nil),  // 0: monitoring.Monitoring_request
	(*MonitoringResponse)(nil), // 1: monitoring.Monitoring_response
}
var file_monitoring_proto_depIdxs = []int32{
	0, // 0: monitoring.Monitoring_Service.SayHello:input_type -> monitoring.Monitoring_request
	1, // 1: monitoring.Monitoring_Service.SayHello:output_type -> monitoring.Monitoring_response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_monitoring_proto_init() }
func file_monitoring_proto_init() {
	if File_monitoring_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_monitoring_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_monitoring_proto_goTypes,
		DependencyIndexes: file_monitoring_proto_depIdxs,
		MessageInfos:      file_monitoring_proto_msgTypes,
	}.Build()
	File_monitoring_proto = out.File
	file_monitoring_proto_rawDesc = nil
	file_monitoring_proto_goTypes = nil
	file_monitoring_proto_depIdxs = nil
}