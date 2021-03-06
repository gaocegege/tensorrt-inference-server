// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/core/api.proto

package nvidia_inferenceserver

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//@@  .. cpp:enum:: Flag
//@@
//@@     Flags that can be associated with an inference request.
//@@     All flags are packed bitwise into the 'flags' field and
//@@     so the value of each must be a power-of-2.
//@@
type InferRequestHeader_Flag int32

const (
	//@@    .. cpp:enumerator:: Flag::FLAG_NONE = 0
	//@@
	//@@       Value indicating no flags are enabled.
	//@@
	InferRequestHeader_FLAG_NONE InferRequestHeader_Flag = 0
	//@@    .. cpp:enumerator:: Flag::FLAG_SEQUENCE_START = 1 << 0
	//@@
	//@@       This request is the start of a related sequence of requests.
	//@@
	InferRequestHeader_FLAG_SEQUENCE_START InferRequestHeader_Flag = 1
	//@@    .. cpp:enumerator:: Flag::FLAG_SEQUENCE_END = 1 << 1
	//@@
	//@@       This request is the end of a related sequence of requests.
	//@@
	InferRequestHeader_FLAG_SEQUENCE_END InferRequestHeader_Flag = 2
)

var InferRequestHeader_Flag_name = map[int32]string{
	0: "FLAG_NONE",
	1: "FLAG_SEQUENCE_START",
	2: "FLAG_SEQUENCE_END",
}

var InferRequestHeader_Flag_value = map[string]int32{
	"FLAG_NONE":           0,
	"FLAG_SEQUENCE_START": 1,
	"FLAG_SEQUENCE_END":   2,
}

func (x InferRequestHeader_Flag) String() string {
	return proto.EnumName(InferRequestHeader_Flag_name, int32(x))
}

func (InferRequestHeader_Flag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f50cf56bc862a54, []int{0, 0}
}

//@@
//@@.. cpp:var:: message InferRequestHeader
//@@
//@@   Meta-data for an inferencing request. The actual input data is
//@@   delivered separate from this header, in the HTTP body for an HTTP
//@@   request, or in the :cpp:var:`InferRequest` message for a gRPC request.
//@@
type InferRequestHeader struct {
	//@@  .. cpp:var:: uint64 id
	//@@
	//@@     The ID of the inference request. The response of the request will
	//@@     have the same ID in InferResponseHeader. The request sender can use
	//@@     the ID to correlate the response to corresponding request if needed.
	//@@
	Id uint64 `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	//@@  .. cpp:var:: uint32 flags
	//@@
	//@@     The flags associated with this request. This field holds a bitwise-or
	//@@     of all flag values.
	//@@
	Flags uint32 `protobuf:"varint,6,opt,name=flags,proto3" json:"flags,omitempty"`
	//@@  .. cpp:var:: uint64 correlation_id
	//@@
	//@@     The correlation ID of the inference request. Default is 0, which
	//@@     indictes that the request has no correlation ID. The correlation ID
	//@@     is used to indicate two or more inference request are related to
	//@@     each other. How this relationship is handled by the inference
	//@@     server is determined by the model's scheduling policy.
	//@@
	CorrelationId uint64 `protobuf:"varint,4,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	//@@  .. cpp:var:: uint32 batch_size
	//@@
	//@@     The batch size of the inference request. This must be >= 1. For
	//@@     models that don't support batching, batch_size must be 1.
	//@@
	BatchSize uint32 `protobuf:"varint,1,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	//@@  .. cpp:var:: Input input (repeated)
	//@@
	//@@     The input meta-data for the inputs provided with the the inference
	//@@     request.
	//@@
	Input []*InferRequestHeader_Input `protobuf:"bytes,2,rep,name=input,proto3" json:"input,omitempty"`
	//@@  .. cpp:var:: Output output (repeated)
	//@@
	//@@     The output meta-data for the inputs provided with the the inference
	//@@     request.
	//@@
	Output               []*InferRequestHeader_Output `protobuf:"bytes,3,rep,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *InferRequestHeader) Reset()         { *m = InferRequestHeader{} }
func (m *InferRequestHeader) String() string { return proto.CompactTextString(m) }
func (*InferRequestHeader) ProtoMessage()    {}
func (*InferRequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f50cf56bc862a54, []int{0}
}

func (m *InferRequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferRequestHeader.Unmarshal(m, b)
}
func (m *InferRequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferRequestHeader.Marshal(b, m, deterministic)
}
func (m *InferRequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferRequestHeader.Merge(m, src)
}
func (m *InferRequestHeader) XXX_Size() int {
	return xxx_messageInfo_InferRequestHeader.Size(m)
}
func (m *InferRequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_InferRequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_InferRequestHeader proto.InternalMessageInfo

func (m *InferRequestHeader) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InferRequestHeader) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *InferRequestHeader) GetCorrelationId() uint64 {
	if m != nil {
		return m.CorrelationId
	}
	return 0
}

func (m *InferRequestHeader) GetBatchSize() uint32 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

func (m *InferRequestHeader) GetInput() []*InferRequestHeader_Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *InferRequestHeader) GetOutput() []*InferRequestHeader_Output {
	if m != nil {
		return m.Output
	}
	return nil
}

//@@  .. cpp:var:: message Input
//@@
//@@     Meta-data for an input tensor provided as part of an inferencing
//@@     request.
//@@
type InferRequestHeader_Input struct {
	//@@    .. cpp:var:: string name
	//@@
	//@@       The name of the input tensor.
	//@@
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//@@    .. cpp:var:: int64 dims (repeated)
	//@@
	//@@       The shape of the input tensor, not including the batch dimension.
	//@@       Optional if the model configuration for this input explicitly
	//@@       specifies all dimensions of the shape. Required if the model
	//@@       configuration for this input has any wildcard dimensions (-1).
	//@@
	Dims []int64 `protobuf:"varint,2,rep,packed,name=dims,proto3" json:"dims,omitempty"`
	//@@    .. cpp:var:: uint64 batch_byte_size
	//@@
	//@@       The size of the full batch of the input tensor, in bytes.
	//@@       Optional for tensors with fixed-sized datatypes. Required
	//@@       for tensors with a non-fixed-size datatype (like STRING).
	//@@
	BatchByteSize        uint64   `protobuf:"varint,3,opt,name=batch_byte_size,json=batchByteSize,proto3" json:"batch_byte_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InferRequestHeader_Input) Reset()         { *m = InferRequestHeader_Input{} }
func (m *InferRequestHeader_Input) String() string { return proto.CompactTextString(m) }
func (*InferRequestHeader_Input) ProtoMessage()    {}
func (*InferRequestHeader_Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f50cf56bc862a54, []int{0, 0}
}

func (m *InferRequestHeader_Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferRequestHeader_Input.Unmarshal(m, b)
}
func (m *InferRequestHeader_Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferRequestHeader_Input.Marshal(b, m, deterministic)
}
func (m *InferRequestHeader_Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferRequestHeader_Input.Merge(m, src)
}
func (m *InferRequestHeader_Input) XXX_Size() int {
	return xxx_messageInfo_InferRequestHeader_Input.Size(m)
}
func (m *InferRequestHeader_Input) XXX_DiscardUnknown() {
	xxx_messageInfo_InferRequestHeader_Input.DiscardUnknown(m)
}

var xxx_messageInfo_InferRequestHeader_Input proto.InternalMessageInfo

func (m *InferRequestHeader_Input) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InferRequestHeader_Input) GetDims() []int64 {
	if m != nil {
		return m.Dims
	}
	return nil
}

func (m *InferRequestHeader_Input) GetBatchByteSize() uint64 {
	if m != nil {
		return m.BatchByteSize
	}
	return 0
}

//@@  .. cpp:var:: message Output
//@@
//@@     Meta-data for a requested output tensor as part of an inferencing
//@@     request.
//@@
type InferRequestHeader_Output struct {
	//@@    .. cpp:var:: string name
	//@@
	//@@       The name of the output tensor.
	//@@
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//@@    .. cpp:var:: Class cls
	//@@
	//@@       Optional. If defined return this output as a classification
	//@@       instead of raw data. The output tensor will be interpreted as
	//@@       probabilities and the classifications associated with the
	//@@       highest probabilities will be returned.
	//@@
	Cls                  *InferRequestHeader_Output_Class `protobuf:"bytes,3,opt,name=cls,proto3" json:"cls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *InferRequestHeader_Output) Reset()         { *m = InferRequestHeader_Output{} }
func (m *InferRequestHeader_Output) String() string { return proto.CompactTextString(m) }
func (*InferRequestHeader_Output) ProtoMessage()    {}
func (*InferRequestHeader_Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f50cf56bc862a54, []int{0, 1}
}

func (m *InferRequestHeader_Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferRequestHeader_Output.Unmarshal(m, b)
}
func (m *InferRequestHeader_Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferRequestHeader_Output.Marshal(b, m, deterministic)
}
func (m *InferRequestHeader_Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferRequestHeader_Output.Merge(m, src)
}
func (m *InferRequestHeader_Output) XXX_Size() int {
	return xxx_messageInfo_InferRequestHeader_Output.Size(m)
}
func (m *InferRequestHeader_Output) XXX_DiscardUnknown() {
	xxx_messageInfo_InferRequestHeader_Output.DiscardUnknown(m)
}

var xxx_messageInfo_InferRequestHeader_Output proto.InternalMessageInfo

func (m *InferRequestHeader_Output) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InferRequestHeader_Output) GetCls() *InferRequestHeader_Output_Class {
	if m != nil {
		return m.Cls
	}
	return nil
}

//@@    .. cpp:var:: message Class
//@@
//@@       Options for an output returned as a classification.
//@@
type InferRequestHeader_Output_Class struct {
	//@@      .. cpp:var:: uint32 count
	//@@
	//@@         Indicates how many classification values should be returned
	//@@         for the output. The 'count' highest priority values are
	//@@         returned.
	//@@
	Count                uint32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InferRequestHeader_Output_Class) Reset()         { *m = InferRequestHeader_Output_Class{} }
func (m *InferRequestHeader_Output_Class) String() string { return proto.CompactTextString(m) }
func (*InferRequestHeader_Output_Class) ProtoMessage()    {}
func (*InferRequestHeader_Output_Class) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f50cf56bc862a54, []int{0, 1, 0}
}

func (m *InferRequestHeader_Output_Class) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferRequestHeader_Output_Class.Unmarshal(m, b)
}
func (m *InferRequestHeader_Output_Class) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferRequestHeader_Output_Class.Marshal(b, m, deterministic)
}
func (m *InferRequestHeader_Output_Class) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferRequestHeader_Output_Class.Merge(m, src)
}
func (m *InferRequestHeader_Output_Class) XXX_Size() int {
	return xxx_messageInfo_InferRequestHeader_Output_Class.Size(m)
}
func (m *InferRequestHeader_Output_Class) XXX_DiscardUnknown() {
	xxx_messageInfo_InferRequestHeader_Output_Class.DiscardUnknown(m)
}

var xxx_messageInfo_InferRequestHeader_Output_Class proto.InternalMessageInfo

func (m *InferRequestHeader_Output_Class) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

//@@
//@@.. cpp:var:: message InferResponseHeader
//@@
//@@   Meta-data for the response to an inferencing request. The actual output
//@@   data is delivered separate from this header, in the HTTP body for an HTTP
//@@   request, or in the :cpp:var:`InferResponse` message for a gRPC request.
//@@
type InferResponseHeader struct {
	//@@  .. cpp:var:: uint64 id
	//@@
	//@@     The ID of the inference response. The response will have the same ID
	//@@     as the ID of its originated request. The request sender can use
	//@@     the ID to correlate the response to corresponding request if needed.
	//@@
	Id uint64 `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	//@@  .. cpp:var:: string model_name
	//@@
	//@@     The name of the model that produced the outputs.
	//@@
	ModelName string `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	//@@  .. cpp:var:: int64 model_version
	//@@
	//@@     The version of the model that produced the outputs.
	//@@
	ModelVersion int64 `protobuf:"varint,2,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"`
	//@@  .. cpp:var:: uint32 batch_size
	//@@
	//@@     The batch size of the outputs. This will always be equal to the
	//@@     batch size of the inputs. For models that don't support
	//@@     batching the batch_size will be 1.
	//@@
	BatchSize uint32 `protobuf:"varint,3,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	//@@  .. cpp:var:: Output output (repeated)
	//@@
	//@@     The outputs, in the same order as they were requested in
	//@@     :cpp:var:`InferRequestHeader`.
	//@@
	Output               []*InferResponseHeader_Output `protobuf:"bytes,4,rep,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *InferResponseHeader) Reset()         { *m = InferResponseHeader{} }
func (m *InferResponseHeader) String() string { return proto.CompactTextString(m) }
func (*InferResponseHeader) ProtoMessage()    {}
func (*InferResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f50cf56bc862a54, []int{1}
}

func (m *InferResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferResponseHeader.Unmarshal(m, b)
}
func (m *InferResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferResponseHeader.Marshal(b, m, deterministic)
}
func (m *InferResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferResponseHeader.Merge(m, src)
}
func (m *InferResponseHeader) XXX_Size() int {
	return xxx_messageInfo_InferResponseHeader.Size(m)
}
func (m *InferResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_InferResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_InferResponseHeader proto.InternalMessageInfo

func (m *InferResponseHeader) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InferResponseHeader) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *InferResponseHeader) GetModelVersion() int64 {
	if m != nil {
		return m.ModelVersion
	}
	return 0
}

func (m *InferResponseHeader) GetBatchSize() uint32 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

func (m *InferResponseHeader) GetOutput() []*InferResponseHeader_Output {
	if m != nil {
		return m.Output
	}
	return nil
}

//@@  .. cpp:var:: message Output
//@@
//@@     Meta-data for an output tensor requested as part of an inferencing
//@@     request.
//@@
type InferResponseHeader_Output struct {
	//@@    .. cpp:var:: string name
	//@@
	//@@       The name of the output tensor.
	//@@
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//@@    .. cpp:var:: Raw raw
	//@@
	//@@       If specified deliver results for this output as raw tensor data.
	//@@       The actual output data is delivered in the HTTP body for an HTTP
	//@@       request, or in the :cpp:var:`InferResponse` message for a gRPC
	//@@       request. Only one of 'raw' and 'batch_classes' may be specified.
	//@@
	Raw *InferResponseHeader_Output_Raw `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	//@@    .. cpp:var:: Classes batch_classes (repeated)
	//@@
	//@@       If specified deliver results for this output as classifications.
	//@@       There is one :cpp:var:`Classes` object for each batch entry in
	//@@       the output. Only one of 'raw' and 'batch_classes' may be
	//@@       specified.
	//@@
	BatchClasses         []*InferResponseHeader_Output_Classes `protobuf:"bytes,3,rep,name=batch_classes,json=batchClasses,proto3" json:"batch_classes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *InferResponseHeader_Output) Reset()         { *m = InferResponseHeader_Output{} }
func (m *InferResponseHeader_Output) String() string { return proto.CompactTextString(m) }
func (*InferResponseHeader_Output) ProtoMessage()    {}
func (*InferResponseHeader_Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f50cf56bc862a54, []int{1, 0}
}

func (m *InferResponseHeader_Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferResponseHeader_Output.Unmarshal(m, b)
}
func (m *InferResponseHeader_Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferResponseHeader_Output.Marshal(b, m, deterministic)
}
func (m *InferResponseHeader_Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferResponseHeader_Output.Merge(m, src)
}
func (m *InferResponseHeader_Output) XXX_Size() int {
	return xxx_messageInfo_InferResponseHeader_Output.Size(m)
}
func (m *InferResponseHeader_Output) XXX_DiscardUnknown() {
	xxx_messageInfo_InferResponseHeader_Output.DiscardUnknown(m)
}

var xxx_messageInfo_InferResponseHeader_Output proto.InternalMessageInfo

func (m *InferResponseHeader_Output) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InferResponseHeader_Output) GetRaw() *InferResponseHeader_Output_Raw {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *InferResponseHeader_Output) GetBatchClasses() []*InferResponseHeader_Output_Classes {
	if m != nil {
		return m.BatchClasses
	}
	return nil
}

//@@    .. cpp:var:: message Raw
//@@
//@@       Meta-data for an output tensor being returned as raw data.
//@@
type InferResponseHeader_Output_Raw struct {
	//@@      .. cpp:var:: int64 dims (repeated)
	//@@
	//@@         The shape of the output tensor, not including the batch
	//@@         dimension.
	//@@
	Dims []int64 `protobuf:"varint,1,rep,packed,name=dims,proto3" json:"dims,omitempty"`
	//@@      .. cpp:var:: uint64 batch_byte_size
	//@@
	//@@         The full size of the output tensor, in bytes. For a
	//@@         batch output, this is the size of the entire batch.
	//@@
	BatchByteSize        uint64   `protobuf:"varint,2,opt,name=batch_byte_size,json=batchByteSize,proto3" json:"batch_byte_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InferResponseHeader_Output_Raw) Reset()         { *m = InferResponseHeader_Output_Raw{} }
func (m *InferResponseHeader_Output_Raw) String() string { return proto.CompactTextString(m) }
func (*InferResponseHeader_Output_Raw) ProtoMessage()    {}
func (*InferResponseHeader_Output_Raw) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f50cf56bc862a54, []int{1, 0, 0}
}

func (m *InferResponseHeader_Output_Raw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferResponseHeader_Output_Raw.Unmarshal(m, b)
}
func (m *InferResponseHeader_Output_Raw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferResponseHeader_Output_Raw.Marshal(b, m, deterministic)
}
func (m *InferResponseHeader_Output_Raw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferResponseHeader_Output_Raw.Merge(m, src)
}
func (m *InferResponseHeader_Output_Raw) XXX_Size() int {
	return xxx_messageInfo_InferResponseHeader_Output_Raw.Size(m)
}
func (m *InferResponseHeader_Output_Raw) XXX_DiscardUnknown() {
	xxx_messageInfo_InferResponseHeader_Output_Raw.DiscardUnknown(m)
}

var xxx_messageInfo_InferResponseHeader_Output_Raw proto.InternalMessageInfo

func (m *InferResponseHeader_Output_Raw) GetDims() []int64 {
	if m != nil {
		return m.Dims
	}
	return nil
}

func (m *InferResponseHeader_Output_Raw) GetBatchByteSize() uint64 {
	if m != nil {
		return m.BatchByteSize
	}
	return 0
}

//@@    .. cpp:var:: message Class
//@@
//@@       Information about each classification for this output.
//@@
type InferResponseHeader_Output_Class struct {
	//@@      .. cpp:var:: int32 idx
	//@@
	//@@         The classification index.
	//@@
	Idx int32 `protobuf:"varint,1,opt,name=idx,proto3" json:"idx,omitempty"`
	//@@      .. cpp:var:: float value
	//@@
	//@@         The classification value as a float (typically a
	//@@         probability).
	//@@
	Value float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	//@@      .. cpp:var:: string label
	//@@
	//@@         The label for the class (optional, only available if provided
	//@@         by the model).
	//@@
	Label                string   `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InferResponseHeader_Output_Class) Reset()         { *m = InferResponseHeader_Output_Class{} }
func (m *InferResponseHeader_Output_Class) String() string { return proto.CompactTextString(m) }
func (*InferResponseHeader_Output_Class) ProtoMessage()    {}
func (*InferResponseHeader_Output_Class) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f50cf56bc862a54, []int{1, 0, 1}
}

func (m *InferResponseHeader_Output_Class) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferResponseHeader_Output_Class.Unmarshal(m, b)
}
func (m *InferResponseHeader_Output_Class) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferResponseHeader_Output_Class.Marshal(b, m, deterministic)
}
func (m *InferResponseHeader_Output_Class) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferResponseHeader_Output_Class.Merge(m, src)
}
func (m *InferResponseHeader_Output_Class) XXX_Size() int {
	return xxx_messageInfo_InferResponseHeader_Output_Class.Size(m)
}
func (m *InferResponseHeader_Output_Class) XXX_DiscardUnknown() {
	xxx_messageInfo_InferResponseHeader_Output_Class.DiscardUnknown(m)
}

var xxx_messageInfo_InferResponseHeader_Output_Class proto.InternalMessageInfo

func (m *InferResponseHeader_Output_Class) GetIdx() int32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

func (m *InferResponseHeader_Output_Class) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *InferResponseHeader_Output_Class) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

//@@    .. cpp:var:: message Classes
//@@
//@@       Meta-data for an output tensor being returned as classifications.
//@@
type InferResponseHeader_Output_Classes struct {
	//@@      .. cpp:var:: Class cls (repeated)
	//@@
	//@@         The topk classes for this output.
	//@@
	Cls                  []*InferResponseHeader_Output_Class `protobuf:"bytes,1,rep,name=cls,proto3" json:"cls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *InferResponseHeader_Output_Classes) Reset()         { *m = InferResponseHeader_Output_Classes{} }
func (m *InferResponseHeader_Output_Classes) String() string { return proto.CompactTextString(m) }
func (*InferResponseHeader_Output_Classes) ProtoMessage()    {}
func (*InferResponseHeader_Output_Classes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f50cf56bc862a54, []int{1, 0, 2}
}

func (m *InferResponseHeader_Output_Classes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferResponseHeader_Output_Classes.Unmarshal(m, b)
}
func (m *InferResponseHeader_Output_Classes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferResponseHeader_Output_Classes.Marshal(b, m, deterministic)
}
func (m *InferResponseHeader_Output_Classes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferResponseHeader_Output_Classes.Merge(m, src)
}
func (m *InferResponseHeader_Output_Classes) XXX_Size() int {
	return xxx_messageInfo_InferResponseHeader_Output_Classes.Size(m)
}
func (m *InferResponseHeader_Output_Classes) XXX_DiscardUnknown() {
	xxx_messageInfo_InferResponseHeader_Output_Classes.DiscardUnknown(m)
}

var xxx_messageInfo_InferResponseHeader_Output_Classes proto.InternalMessageInfo

func (m *InferResponseHeader_Output_Classes) GetCls() []*InferResponseHeader_Output_Class {
	if m != nil {
		return m.Cls
	}
	return nil
}

func init() {
	proto.RegisterEnum("nvidia.inferenceserver.InferRequestHeader_Flag", InferRequestHeader_Flag_name, InferRequestHeader_Flag_value)
	proto.RegisterType((*InferRequestHeader)(nil), "nvidia.inferenceserver.InferRequestHeader")
	proto.RegisterType((*InferRequestHeader_Input)(nil), "nvidia.inferenceserver.InferRequestHeader.Input")
	proto.RegisterType((*InferRequestHeader_Output)(nil), "nvidia.inferenceserver.InferRequestHeader.Output")
	proto.RegisterType((*InferRequestHeader_Output_Class)(nil), "nvidia.inferenceserver.InferRequestHeader.Output.Class")
	proto.RegisterType((*InferResponseHeader)(nil), "nvidia.inferenceserver.InferResponseHeader")
	proto.RegisterType((*InferResponseHeader_Output)(nil), "nvidia.inferenceserver.InferResponseHeader.Output")
	proto.RegisterType((*InferResponseHeader_Output_Raw)(nil), "nvidia.inferenceserver.InferResponseHeader.Output.Raw")
	proto.RegisterType((*InferResponseHeader_Output_Class)(nil), "nvidia.inferenceserver.InferResponseHeader.Output.Class")
	proto.RegisterType((*InferResponseHeader_Output_Classes)(nil), "nvidia.inferenceserver.InferResponseHeader.Output.Classes")
}

func init() { proto.RegisterFile("src/core/api.proto", fileDescriptor_6f50cf56bc862a54) }

var fileDescriptor_6f50cf56bc862a54 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0x12, 0x41,
	0x18, 0x75, 0xff, 0x30, 0x7c, 0x2d, 0x15, 0xa7, 0xfe, 0x6c, 0x48, 0x48, 0x08, 0x46, 0xc3, 0xd5,
	0x56, 0x31, 0x51, 0xe3, 0x1d, 0xe2, 0x62, 0x69, 0x0c, 0x8d, 0x43, 0xab, 0x97, 0x9b, 0x61, 0x77,
	0x5a, 0x27, 0x59, 0x76, 0x71, 0x66, 0x01, 0xdb, 0x07, 0xf0, 0x59, 0x7c, 0x08, 0xdf, 0x4d, 0x33,
	0xdf, 0xac, 0x0d, 0x22, 0xd6, 0xd0, 0xbb, 0xf9, 0xce, 0xe6, 0x3b, 0x73, 0xce, 0x9c, 0x93, 0x05,
	0xa2, 0x64, 0x7c, 0x10, 0xe7, 0x92, 0x1f, 0xb0, 0x99, 0x08, 0x66, 0x32, 0x2f, 0x72, 0xf2, 0x20,
	0x5b, 0x88, 0x44, 0xb0, 0x40, 0x64, 0x67, 0x5c, 0xf2, 0x2c, 0xe6, 0x8a, 0xcb, 0x05, 0x97, 0xed,
	0x1f, 0x2e, 0x90, 0xa1, 0xc6, 0x28, 0xff, 0x32, 0xe7, 0xaa, 0x38, 0xe4, 0x2c, 0xe1, 0x92, 0xec,
	0x81, 0x2d, 0x12, 0xdf, 0x6b, 0x59, 0x1d, 0x97, 0xda, 0x22, 0x21, 0xf7, 0xc0, 0x3b, 0x4b, 0xd9,
	0xb9, 0xf2, 0x2b, 0x2d, 0xab, 0x53, 0xa3, 0x66, 0x20, 0x8f, 0x61, 0x2f, 0xce, 0xa5, 0xe4, 0x29,
	0x2b, 0x44, 0x9e, 0x45, 0x22, 0xf1, 0x5d, 0xdc, 0xa8, 0xad, 0xa0, 0xc3, 0x84, 0x34, 0x01, 0x26,
	0xac, 0x88, 0x3f, 0x47, 0x4a, 0x5c, 0x72, 0xdf, 0x42, 0x86, 0x2a, 0x22, 0x63, 0x71, 0xc9, 0xc9,
	0x00, 0x3c, 0x91, 0xcd, 0xe6, 0x85, 0x6f, 0xb7, 0x9c, 0xce, 0x4e, 0xf7, 0x69, 0xb0, 0x59, 0x6a,
	0xf0, 0xb7, 0xcc, 0x60, 0xa8, 0xf7, 0xa8, 0x59, 0x27, 0x43, 0xa8, 0xe4, 0xf3, 0x42, 0x13, 0x39,
	0x48, 0xf4, 0x6c, 0x0b, 0xa2, 0x63, 0x5c, 0xa4, 0x25, 0x41, 0xe3, 0x13, 0x78, 0x48, 0x4d, 0x08,
	0xb8, 0x19, 0x9b, 0x1a, 0xd1, 0x55, 0x8a, 0x67, 0x8d, 0x25, 0x62, 0xaa, 0x50, 0xae, 0x43, 0xf1,
	0x4c, 0x9e, 0xc0, 0x1d, 0x63, 0x71, 0x72, 0x51, 0x70, 0xe3, 0xd3, 0x31, 0x4f, 0x81, 0xf0, 0x9b,
	0x8b, 0x82, 0x6b, 0xaf, 0x8d, 0x6f, 0x16, 0x54, 0xcc, 0x5d, 0x1b, 0xa9, 0x87, 0xe0, 0xc4, 0xa9,
	0xc2, 0xd5, 0x9d, 0xee, 0xcb, 0xad, 0xf5, 0x07, 0xfd, 0x94, 0x29, 0x45, 0x35, 0x47, 0xa3, 0x09,
	0x1e, 0x4e, 0x3a, 0xba, 0x38, 0x9f, 0x67, 0x45, 0xf9, 0xf0, 0x66, 0x68, 0x87, 0xe0, 0x0e, 0x52,
	0x76, 0x4e, 0x6a, 0x50, 0x1d, 0xbc, 0xef, 0xbd, 0x8b, 0x46, 0xc7, 0xa3, 0xb0, 0x7e, 0x8b, 0x3c,
	0x84, 0x7d, 0x1c, 0xc7, 0xe1, 0x87, 0xd3, 0x70, 0xd4, 0x0f, 0xa3, 0xf1, 0x49, 0x8f, 0x9e, 0xd4,
	0x2d, 0x72, 0x1f, 0xee, 0xfe, 0xf9, 0x21, 0x1c, 0xbd, 0xad, 0xdb, 0xed, 0x9f, 0x2e, 0xec, 0x97,
	0x72, 0xd4, 0x2c, 0xcf, 0x14, 0xff, 0x47, 0x7f, 0x9a, 0x00, 0xd3, 0x3c, 0xe1, 0x69, 0xb4, 0x62,
	0xb9, 0x8a, 0xc8, 0x48, 0xfb, 0x7e, 0x04, 0x35, 0xf3, 0x79, 0xc1, 0xa5, 0x12, 0x79, 0xe6, 0xdb,
	0x2d, 0xab, 0xe3, 0xd0, 0x5d, 0x04, 0x3f, 0x1a, 0x6c, 0xad, 0x46, 0xce, 0x7a, 0x8d, 0x8e, 0xae,
	0xe2, 0x77, 0x31, 0xfe, 0xee, 0x7f, 0x9e, 0x6f, 0x55, 0xef, 0x7a, 0xfe, 0xdf, 0x9d, 0x6b, 0x63,
	0x3a, 0x04, 0x47, 0xb2, 0x25, 0x8a, 0xdc, 0xe9, 0xbe, 0xd8, 0xfe, 0x9e, 0x80, 0xb2, 0x25, 0xd5,
	0x14, 0x24, 0x02, 0x53, 0x90, 0x28, 0xd6, 0x59, 0x71, 0x55, 0x56, 0xf7, 0xf5, 0x0d, 0x38, 0xfb,
	0x86, 0x81, 0xee, 0x22, 0x61, 0x39, 0x35, 0x7a, 0xe0, 0x50, 0xb6, 0xbc, 0xea, 0xac, 0x75, 0x7d,
	0x67, 0xed, 0x4d, 0x9d, 0x0d, 0x7f, 0x37, 0xa9, 0x0e, 0x8e, 0x48, 0xbe, 0xe2, 0x4b, 0x78, 0x54,
	0x1f, 0x75, 0xb7, 0x16, 0x2c, 0x9d, 0x9b, 0x45, 0x9b, 0x9a, 0x41, 0xa3, 0x29, 0x9b, 0xf0, 0x14,
	0x33, 0xaa, 0x52, 0x33, 0x34, 0x4e, 0xe1, 0x76, 0x29, 0x8a, 0x1c, 0x99, 0x9a, 0x5b, 0xe8, 0xf5,
	0xd5, 0x4d, 0xbd, 0x62, 0xcf, 0x27, 0x15, 0xfc, 0xbf, 0x3d, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff,
	0x9e, 0x39, 0x3c, 0x63, 0xf5, 0x04, 0x00, 0x00,
}
