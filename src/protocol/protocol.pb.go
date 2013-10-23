// Code generated by protoc-gen-go.
// source: src/protocol/protocol.proto
// DO NOT EDIT!

package protocol

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request_Type int32

const (
	Request_QUERY       Request_Type = 1
	Request_WRITE       Request_Type = 2
	Request_GET_SERVERS Request_Type = 3
)

var Request_Type_name = map[int32]string{
	1: "QUERY",
	2: "WRITE",
	3: "GET_SERVERS",
}
var Request_Type_value = map[string]int32{
	"QUERY":       1,
	"WRITE":       2,
	"GET_SERVERS": 3,
}

func (x Request_Type) Enum() *Request_Type {
	p := new(Request_Type)
	*p = x
	return p
}
func (x Request_Type) String() string {
	return proto.EnumName(Request_Type_name, int32(x))
}
func (x Request_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Request_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Request_Type_value, data, "Request_Type")
	if err != nil {
		return err
	}
	*x = Request_Type(value)
	return nil
}

type Response_Type int32

const (
	Response_QUERY      Response_Type = 1
	Response_WRITE_OK   Response_Type = 2
	Response_END_STREAM Response_Type = 3
)

var Response_Type_name = map[int32]string{
	1: "QUERY",
	2: "WRITE_OK",
	3: "END_STREAM",
}
var Response_Type_value = map[string]int32{
	"QUERY":      1,
	"WRITE_OK":   2,
	"END_STREAM": 3,
}

func (x Response_Type) Enum() *Response_Type {
	p := new(Response_Type)
	*p = x
	return p
}
func (x Response_Type) String() string {
	return proto.EnumName(Response_Type_name, int32(x))
}
func (x Response_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Response_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Response_Type_value, data, "Response_Type")
	if err != nil {
		return err
	}
	*x = Response_Type(value)
	return nil
}

type FieldValue struct {
	StringValue      *string  `protobuf:"bytes,1,opt,name=string_value" json:"string_value,omitempty"`
	DoubleValue      *float64 `protobuf:"fixed64,3,opt,name=double_value" json:"double_value,omitempty"`
	BoolValue        *bool    `protobuf:"varint,4,opt,name=bool_value" json:"bool_value,omitempty"`
	Int64Value       *int64   `protobuf:"varint,5,opt,name=int64_value" json:"int64_value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *FieldValue) Reset()         { *m = FieldValue{} }
func (m *FieldValue) String() string { return proto.CompactTextString(m) }
func (*FieldValue) ProtoMessage()    {}

func (m *FieldValue) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *FieldValue) GetDoubleValue() float64 {
	if m != nil && m.DoubleValue != nil {
		return *m.DoubleValue
	}
	return 0
}

func (m *FieldValue) GetBoolValue() bool {
	if m != nil && m.BoolValue != nil {
		return *m.BoolValue
	}
	return false
}

func (m *FieldValue) GetInt64Value() int64 {
	if m != nil && m.Int64Value != nil {
		return *m.Int64Value
	}
	return 0
}

type Point struct {
	Values           []*FieldValue `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	Timestamp        *int64        `protobuf:"varint,2,req,name=timestamp" json:"timestamp,omitempty"`
	SequenceNumber   *uint32       `protobuf:"varint,3,req,name=sequence_number" json:"sequence_number,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}

func (m *Point) GetValues() []*FieldValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Point) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Point) GetSequenceNumber() uint32 {
	if m != nil && m.SequenceNumber != nil {
		return *m.SequenceNumber
	}
	return 0
}

type Series struct {
	Points           []*Point `protobuf:"bytes,1,rep,name=points" json:"points,omitempty"`
	Name             *string  `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Fields           []string `protobuf:"bytes,3,rep,name=fields" json:"fields,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Series) Reset()         { *m = Series{} }
func (m *Series) String() string { return proto.CompactTextString(m) }
func (*Series) ProtoMessage()    {}

func (m *Series) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *Series) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Series) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

type Request struct {
	Id               *int32        `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Type             *Request_Type `protobuf:"varint,2,req,name=type,enum=protocol.Request_Type" json:"type,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Request) GetType() Request_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Request_QUERY
}

type Response struct {
	Id               *int32   `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Series           *Series  `protobuf:"bytes,2,opt,name=series" json:"series,omitempty"`
	Servers          []string `protobuf:"bytes,3,rep,name=servers" json:"servers,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Response) GetSeries() *Series {
	if m != nil {
		return m.Series
	}
	return nil
}

func (m *Response) GetServers() []string {
	if m != nil {
		return m.Servers
	}
	return nil
}

func init() {
	proto.RegisterEnum("protocol.Request_Type", Request_Type_name, Request_Type_value)
	proto.RegisterEnum("protocol.Response_Type", Response_Type_name, Response_Type_value)
}
