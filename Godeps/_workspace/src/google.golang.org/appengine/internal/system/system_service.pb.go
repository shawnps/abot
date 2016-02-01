// Code generated by protoc-gen-go.
// source: google.golang.org/appengine/internal/system/system_service.proto
// DO NOT EDIT!

/*
Package system is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/appengine/internal/system/system_service.proto

It has these top-level messages:
	SystemServiceError
	SystemStat
	GetSystemStatsRequest
	GetSystemStatsResponse
	StartBackgroundRequestRequest
	StartBackgroundRequestResponse
*/
package system

import proto "github.com/avabot/ava/Godeps/_workspace/src/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SystemServiceError_ErrorCode int32

const (
	SystemServiceError_OK               SystemServiceError_ErrorCode = 0
	SystemServiceError_INTERNAL_ERROR   SystemServiceError_ErrorCode = 1
	SystemServiceError_BACKEND_REQUIRED SystemServiceError_ErrorCode = 2
	SystemServiceError_LIMIT_REACHED    SystemServiceError_ErrorCode = 3
)

var SystemServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "INTERNAL_ERROR",
	2: "BACKEND_REQUIRED",
	3: "LIMIT_REACHED",
}
var SystemServiceError_ErrorCode_value = map[string]int32{
	"OK":               0,
	"INTERNAL_ERROR":   1,
	"BACKEND_REQUIRED": 2,
	"LIMIT_REACHED":    3,
}

func (x SystemServiceError_ErrorCode) Enum() *SystemServiceError_ErrorCode {
	p := new(SystemServiceError_ErrorCode)
	*p = x
	return p
}
func (x SystemServiceError_ErrorCode) String() string {
	return proto.EnumName(SystemServiceError_ErrorCode_name, int32(x))
}
func (x *SystemServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SystemServiceError_ErrorCode_value, data, "SystemServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = SystemServiceError_ErrorCode(value)
	return nil
}

type SystemServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SystemServiceError) Reset()         { *m = SystemServiceError{} }
func (m *SystemServiceError) String() string { return proto.CompactTextString(m) }
func (*SystemServiceError) ProtoMessage()    {}

type SystemStat struct {
	// Instaneous value of this stat.
	Current *float64 `protobuf:"fixed64,1,opt,name=current" json:"current,omitempty"`
	// Average over time, if this stat has an instaneous value.
	Average1M  *float64 `protobuf:"fixed64,3,opt,name=average1m" json:"average1m,omitempty"`
	Average10M *float64 `protobuf:"fixed64,4,opt,name=average10m" json:"average10m,omitempty"`
	// Total value, if the stat accumulates over time.
	Total *float64 `protobuf:"fixed64,2,opt,name=total" json:"total,omitempty"`
	// Rate over time, if this stat accumulates.
	Rate1M           *float64 `protobuf:"fixed64,5,opt,name=rate1m" json:"rate1m,omitempty"`
	Rate10M          *float64 `protobuf:"fixed64,6,opt,name=rate10m" json:"rate10m,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SystemStat) Reset()         { *m = SystemStat{} }
func (m *SystemStat) String() string { return proto.CompactTextString(m) }
func (*SystemStat) ProtoMessage()    {}

func (m *SystemStat) GetCurrent() float64 {
	if m != nil && m.Current != nil {
		return *m.Current
	}
	return 0
}

func (m *SystemStat) GetAverage1M() float64 {
	if m != nil && m.Average1M != nil {
		return *m.Average1M
	}
	return 0
}

func (m *SystemStat) GetAverage10M() float64 {
	if m != nil && m.Average10M != nil {
		return *m.Average10M
	}
	return 0
}

func (m *SystemStat) GetTotal() float64 {
	if m != nil && m.Total != nil {
		return *m.Total
	}
	return 0
}

func (m *SystemStat) GetRate1M() float64 {
	if m != nil && m.Rate1M != nil {
		return *m.Rate1M
	}
	return 0
}

func (m *SystemStat) GetRate10M() float64 {
	if m != nil && m.Rate10M != nil {
		return *m.Rate10M
	}
	return 0
}

type GetSystemStatsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetSystemStatsRequest) Reset()         { *m = GetSystemStatsRequest{} }
func (m *GetSystemStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetSystemStatsRequest) ProtoMessage()    {}

type GetSystemStatsResponse struct {
	// CPU used by this instance, in mcycles.
	Cpu *SystemStat `protobuf:"bytes,1,opt,name=cpu" json:"cpu,omitempty"`
	// Physical memory (RAM) used by this instance, in megabytes.
	Memory           *SystemStat `protobuf:"bytes,2,opt,name=memory" json:"memory,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *GetSystemStatsResponse) Reset()         { *m = GetSystemStatsResponse{} }
func (m *GetSystemStatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetSystemStatsResponse) ProtoMessage()    {}

func (m *GetSystemStatsResponse) GetCpu() *SystemStat {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *GetSystemStatsResponse) GetMemory() *SystemStat {
	if m != nil {
		return m.Memory
	}
	return nil
}

type StartBackgroundRequestRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *StartBackgroundRequestRequest) Reset()         { *m = StartBackgroundRequestRequest{} }
func (m *StartBackgroundRequestRequest) String() string { return proto.CompactTextString(m) }
func (*StartBackgroundRequestRequest) ProtoMessage()    {}

type StartBackgroundRequestResponse struct {
	// Every /_ah/background request will have an X-AppEngine-BackgroundRequest
	// header, whose value will be equal to this parameter, the request_id.
	RequestId        *string `protobuf:"bytes,1,opt,name=request_id" json:"request_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StartBackgroundRequestResponse) Reset()         { *m = StartBackgroundRequestResponse{} }
func (m *StartBackgroundRequestResponse) String() string { return proto.CompactTextString(m) }
func (*StartBackgroundRequestResponse) ProtoMessage()    {}

func (m *StartBackgroundRequestResponse) GetRequestId() string {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return ""
}

func init() {
}
