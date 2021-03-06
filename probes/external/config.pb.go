// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/probes/external/config.proto

/*
Package external is a generated protocol buffer package.

It is generated from these files:
	github.com/google/cloudprober/probes/external/config.proto

It has these top-level messages:
	ProbeConf
	OutputMetricsOptions
*/
package external

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import cloudprober_metrics "github.com/google/cloudprober/metrics"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// External probes support two mode: ONCE and SERVER. In ONCE mode, external
// command is re-executed for each probe run, while in SERVER mode, command
// is run in server mode, re-executed only if not running already.
type ProbeConf_Mode int32

const (
	ProbeConf_ONCE   ProbeConf_Mode = 0
	ProbeConf_SERVER ProbeConf_Mode = 1
)

var ProbeConf_Mode_name = map[int32]string{
	0: "ONCE",
	1: "SERVER",
}
var ProbeConf_Mode_value = map[string]int32{
	"ONCE":   0,
	"SERVER": 1,
}

func (x ProbeConf_Mode) Enum() *ProbeConf_Mode {
	p := new(ProbeConf_Mode)
	*p = x
	return p
}
func (x ProbeConf_Mode) String() string {
	return proto.EnumName(ProbeConf_Mode_name, int32(x))
}
func (x *ProbeConf_Mode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProbeConf_Mode_value, data, "ProbeConf_Mode")
	if err != nil {
		return err
	}
	*x = ProbeConf_Mode(value)
	return nil
}
func (ProbeConf_Mode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// MetricsKind specifies whether to treat output metrics as GAUGE or
// CUMULATIVE. If left unspecified, metrics from ONCE mode probes are treated
// as GAUGE and metrics from SERVER mode probes are treated as CUMULATIVE.
type OutputMetricsOptions_MetricsKind int32

const (
	OutputMetricsOptions_UNDEFINED  OutputMetricsOptions_MetricsKind = 0
	OutputMetricsOptions_GAUGE      OutputMetricsOptions_MetricsKind = 1
	OutputMetricsOptions_CUMULATIVE OutputMetricsOptions_MetricsKind = 2
)

var OutputMetricsOptions_MetricsKind_name = map[int32]string{
	0: "UNDEFINED",
	1: "GAUGE",
	2: "CUMULATIVE",
}
var OutputMetricsOptions_MetricsKind_value = map[string]int32{
	"UNDEFINED":  0,
	"GAUGE":      1,
	"CUMULATIVE": 2,
}

func (x OutputMetricsOptions_MetricsKind) Enum() *OutputMetricsOptions_MetricsKind {
	p := new(OutputMetricsOptions_MetricsKind)
	*p = x
	return p
}
func (x OutputMetricsOptions_MetricsKind) String() string {
	return proto.EnumName(OutputMetricsOptions_MetricsKind_name, int32(x))
}
func (x *OutputMetricsOptions_MetricsKind) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OutputMetricsOptions_MetricsKind_value, data, "OutputMetricsOptions_MetricsKind")
	if err != nil {
		return err
	}
	*x = OutputMetricsOptions_MetricsKind(value)
	return nil
}
func (OutputMetricsOptions_MetricsKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type ProbeConf struct {
	Mode *ProbeConf_Mode `protobuf:"varint,1,opt,name=mode,enum=cloudprober.probes.external.ProbeConf_Mode,def=0" json:"mode,omitempty"`
	// Command.  For ONCE probes, arguments are processed for the following field
	// substitutions:
	// @probe@    Name of the probe
	// @target@   Hostname of the target
	// @address@  IP address of the target
	//
	// For example, for target ig-us-central1-a, /tools/recreate_vm -vm @target@
	// will get converted to: /tools/recreate_vm -vm ig-us-central1-a
	Command *string             `protobuf:"bytes,2,req,name=command" json:"command,omitempty"`
	Options []*ProbeConf_Option `protobuf:"bytes,3,rep,name=options" json:"options,omitempty"`
	// Export output as metrics, where output is the output returned by the
	// external probe process, over stdout for ONCE probes, and through ProbeReply
	// for SERVER probes. Cloudprober expects variables to be in the following
	// format in the output:
	// var1 value1 (for example: total_errors 589)
	OutputAsMetrics      *bool                 `protobuf:"varint,4,opt,name=output_as_metrics,json=outputAsMetrics,def=1" json:"output_as_metrics,omitempty"`
	OutputMetricsOptions *OutputMetricsOptions `protobuf:"bytes,5,opt,name=output_metrics_options,json=outputMetricsOptions" json:"output_metrics_options,omitempty"`
	// IP version: For target resolution
	IpVersion        *int32 `protobuf:"varint,100,opt,name=ip_version,json=ipVersion,def=4" json:"ip_version,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ProbeConf) Reset()                    { *m = ProbeConf{} }
func (m *ProbeConf) String() string            { return proto.CompactTextString(m) }
func (*ProbeConf) ProtoMessage()               {}
func (*ProbeConf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_ProbeConf_Mode ProbeConf_Mode = ProbeConf_ONCE
const Default_ProbeConf_OutputAsMetrics bool = true
const Default_ProbeConf_IpVersion int32 = 4

func (m *ProbeConf) GetMode() ProbeConf_Mode {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return Default_ProbeConf_Mode
}

func (m *ProbeConf) GetCommand() string {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return ""
}

func (m *ProbeConf) GetOptions() []*ProbeConf_Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *ProbeConf) GetOutputAsMetrics() bool {
	if m != nil && m.OutputAsMetrics != nil {
		return *m.OutputAsMetrics
	}
	return Default_ProbeConf_OutputAsMetrics
}

func (m *ProbeConf) GetOutputMetricsOptions() *OutputMetricsOptions {
	if m != nil {
		return m.OutputMetricsOptions
	}
	return nil
}

func (m *ProbeConf) GetIpVersion() int32 {
	if m != nil && m.IpVersion != nil {
		return *m.IpVersion
	}
	return Default_ProbeConf_IpVersion
}

// Options for the SERVER mode probe requests. These options are passed on the
// external probe server as part of the ProbeRequest. Values are substituted
// similar to command arguments for the ONCE mode probes.
type ProbeConf_Option struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProbeConf_Option) Reset()                    { *m = ProbeConf_Option{} }
func (m *ProbeConf_Option) String() string            { return proto.CompactTextString(m) }
func (*ProbeConf_Option) ProtoMessage()               {}
func (*ProbeConf_Option) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *ProbeConf_Option) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ProbeConf_Option) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type OutputMetricsOptions struct {
	MetricsKind *OutputMetricsOptions_MetricsKind `protobuf:"varint,1,opt,name=metrics_kind,json=metricsKind,enum=cloudprober.probes.external.OutputMetricsOptions_MetricsKind" json:"metrics_kind,omitempty"`
	// Additional labels (comma-separated) to attach to the output metrics, e.g.
	// "region=us-east1,zone=us-east1-d". ptype="external" and probe="<probeName>"
	// are attached automatically.
	AdditionalLabels *string `protobuf:"bytes,2,opt,name=additional_labels,json=additionalLabels" json:"additional_labels,omitempty"`
	// Whether to aggregate metrics in Cloudprober. If enabled, Cloudprober
	// aggregates the metrics returned by the external probe process -- external
	// probe process should return metrics only since the last probe run.
	// Note that this option is mutually exclusive with GAUGE metrics and
	// cloudprober will fail during initialization if both options are enabled.
	AggregateInCloudprober *bool `protobuf:"varint,3,opt,name=aggregate_in_cloudprober,json=aggregateInCloudprober,def=0" json:"aggregate_in_cloudprober,omitempty"`
	// Metrics that should be treated as distributions. These metrics are exported
	// by the external probe program as comma-separated list of values, for
	// example: "op_latency 4.7,5.6,5.9,6.1,4.9". To be able to build distribution
	// from these values, these metrics should be pre-configured in external
	// probe:
	// dist_metric {
	//   key: "op_latency"
	//   value {
	//     explicit_buckets: "1,2,4,8,16,32,64,128,256"
	//   }
	// }
	DistMetric       map[string]*cloudprober_metrics.Dist `protobuf:"bytes,4,rep,name=dist_metric,json=distMetric" json:"dist_metric,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *OutputMetricsOptions) Reset()                    { *m = OutputMetricsOptions{} }
func (m *OutputMetricsOptions) String() string            { return proto.CompactTextString(m) }
func (*OutputMetricsOptions) ProtoMessage()               {}
func (*OutputMetricsOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

const Default_OutputMetricsOptions_AggregateInCloudprober bool = false

func (m *OutputMetricsOptions) GetMetricsKind() OutputMetricsOptions_MetricsKind {
	if m != nil && m.MetricsKind != nil {
		return *m.MetricsKind
	}
	return OutputMetricsOptions_UNDEFINED
}

func (m *OutputMetricsOptions) GetAdditionalLabels() string {
	if m != nil && m.AdditionalLabels != nil {
		return *m.AdditionalLabels
	}
	return ""
}

func (m *OutputMetricsOptions) GetAggregateInCloudprober() bool {
	if m != nil && m.AggregateInCloudprober != nil {
		return *m.AggregateInCloudprober
	}
	return Default_OutputMetricsOptions_AggregateInCloudprober
}

func (m *OutputMetricsOptions) GetDistMetric() map[string]*cloudprober_metrics.Dist {
	if m != nil {
		return m.DistMetric
	}
	return nil
}

func init() {
	proto.RegisterType((*ProbeConf)(nil), "cloudprober.probes.external.ProbeConf")
	proto.RegisterType((*ProbeConf_Option)(nil), "cloudprober.probes.external.ProbeConf.Option")
	proto.RegisterType((*OutputMetricsOptions)(nil), "cloudprober.probes.external.OutputMetricsOptions")
	proto.RegisterEnum("cloudprober.probes.external.ProbeConf_Mode", ProbeConf_Mode_name, ProbeConf_Mode_value)
	proto.RegisterEnum("cloudprober.probes.external.OutputMetricsOptions_MetricsKind", OutputMetricsOptions_MetricsKind_name, OutputMetricsOptions_MetricsKind_value)
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/probes/external/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xdd, 0x6a, 0xdb, 0x4c,
	0x10, 0x8d, 0x2c, 0x39, 0x89, 0x46, 0xdf, 0x97, 0x28, 0x4b, 0x08, 0xaa, 0xdb, 0x0b, 0xe1, 0x2b,
	0x41, 0xa8, 0x94, 0x9a, 0x42, 0x8b, 0xa1, 0x14, 0x63, 0xab, 0xc6, 0x34, 0xb6, 0xcb, 0xb6, 0x36,
	0xbd, 0x53, 0x65, 0x6b, 0xad, 0x2e, 0x91, 0x76, 0x85, 0x7e, 0x42, 0xf3, 0x7e, 0x7d, 0x8d, 0xbe,
	0x4b, 0xf1, 0x6a, 0xfd, 0xd3, 0x12, 0x4c, 0x73, 0x25, 0xcd, 0x99, 0x73, 0x76, 0xcf, 0xcc, 0x59,
	0xe8, 0xc6, 0xb4, 0xfc, 0x5e, 0x2d, 0xdc, 0x25, 0x4f, 0xbd, 0x98, 0xf3, 0x38, 0x21, 0xde, 0x32,
	0xe1, 0x55, 0x94, 0xe5, 0x7c, 0x41, 0x72, 0x4f, 0x7c, 0x0a, 0x8f, 0xfc, 0x28, 0x49, 0xce, 0xc2,
	0xc4, 0x5b, 0x72, 0xb6, 0xa2, 0xb1, 0x9b, 0xe5, 0xbc, 0xe4, 0xe8, 0xf9, 0x1e, 0xd3, 0xad, 0x99,
	0xee, 0x86, 0xd9, 0xba, 0x39, 0x7c, 0x70, 0x4a, 0xca, 0x9c, 0x2e, 0x0b, 0x2f, 0xa2, 0x45, 0x59,
	0x1f, 0xd7, 0xfe, 0xa9, 0x82, 0xfe, 0x69, 0xdd, 0xed, 0x73, 0xb6, 0x42, 0x3e, 0x68, 0x29, 0x8f,
	0x88, 0xa5, 0xd8, 0x8a, 0x73, 0xd6, 0xb9, 0x76, 0x0f, 0xdc, 0xe5, 0x6e, 0x55, 0xee, 0x98, 0x47,
	0xa4, 0xab, 0x4d, 0x27, 0x7d, 0x1f, 0x0b, 0x39, 0xb2, 0xe0, 0x64, 0xc9, 0xd3, 0x34, 0x64, 0x91,
	0xd5, 0xb0, 0x1b, 0x8e, 0x8e, 0x37, 0x25, 0x1a, 0xc2, 0x09, 0xcf, 0x4a, 0xca, 0x59, 0x61, 0xa9,
	0xb6, 0xea, 0x18, 0x9d, 0x97, 0xff, 0x78, 0xc7, 0x54, 0xa8, 0xf0, 0x46, 0x8d, 0x6e, 0xe0, 0x82,
	0x57, 0x65, 0x56, 0x95, 0x41, 0x58, 0x04, 0x72, 0x2e, 0x4b, 0xb3, 0x15, 0xe7, 0xb4, 0xab, 0x95,
	0x79, 0x45, 0xf0, 0x79, 0xdd, 0xee, 0x15, 0xe3, 0xba, 0x89, 0x62, 0xb8, 0x92, 0x0a, 0x49, 0x0f,
	0x36, 0x4e, 0x9a, 0xb6, 0xe2, 0x18, 0x9d, 0x57, 0x07, 0x9d, 0x4c, 0x85, 0x54, 0x9e, 0x55, 0x9b,
	0x29, 0xf0, 0x25, 0x7f, 0x04, 0x45, 0x36, 0x00, 0xcd, 0x82, 0x7b, 0x92, 0x17, 0x94, 0x33, 0x2b,
	0xb2, 0x15, 0xa7, 0xd9, 0x55, 0x5e, 0x63, 0x9d, 0x66, 0xf3, 0x1a, 0x6b, 0x75, 0xe0, 0xb8, 0x26,
	0x23, 0x04, 0x1a, 0x0b, 0xd3, 0x7a, 0xe1, 0x3a, 0x16, 0xff, 0xe8, 0x12, 0x9a, 0xf7, 0x61, 0x52,
	0x11, 0xab, 0x21, 0xc0, 0xba, 0x68, 0xbf, 0x00, 0x6d, 0xbd, 0x67, 0x74, 0x0a, 0x62, 0xd3, 0xe6,
	0x11, 0x02, 0x38, 0xfe, 0xec, 0xe3, 0xb9, 0x8f, 0x4d, 0xa5, 0xfd, 0x4b, 0x85, 0xcb, 0xc7, 0x2c,
	0xa2, 0x6f, 0xf0, 0xdf, 0x66, 0xdc, 0x3b, 0xca, 0x22, 0x99, 0xec, 0xbb, 0x27, 0xcf, 0xea, 0xca,
	0xf2, 0x23, 0x65, 0x11, 0x36, 0xd2, 0x5d, 0x81, 0xae, 0xe1, 0x22, 0x8c, 0x22, 0xba, 0x66, 0x85,
	0x49, 0x90, 0x84, 0x0b, 0x92, 0x14, 0xd2, 0xba, 0xb9, 0x6b, 0xdc, 0x0a, 0x1c, 0xbd, 0x07, 0x2b,
	0x8c, 0xe3, 0x9c, 0xc4, 0x61, 0x49, 0x02, 0xca, 0x82, 0x3d, 0x1b, 0x96, 0x2a, 0xd2, 0x6b, 0xae,
	0xc2, 0xa4, 0x20, 0xf8, 0x6a, 0x4b, 0x1b, 0xb1, 0xfe, 0x8e, 0x84, 0x16, 0x60, 0xac, 0x5f, 0xaf,
	0xcc, 0xd0, 0xd2, 0xc4, 0x23, 0xea, 0x3d, 0x7d, 0x9c, 0x01, 0x2d, 0x24, 0xe4, 0xb3, 0x32, 0x7f,
	0xc0, 0x10, 0x6d, 0x81, 0xd6, 0x57, 0x38, 0xff, 0xab, 0x8d, 0x4c, 0x50, 0xef, 0xc8, 0x83, 0x8c,
	0x69, 0xfd, 0x8b, 0xbc, 0xfd, 0x94, 0x8c, 0xce, 0xb3, 0x3f, 0x2c, 0xc8, 0xfd, 0x88, 0x5b, 0x64,
	0x80, 0xdd, 0xc6, 0x5b, 0xa5, 0xfd, 0x06, 0x8c, 0xbd, 0x3d, 0xa2, 0xff, 0x41, 0x9f, 0x4d, 0x06,
	0xfe, 0x87, 0xd1, 0xc4, 0x1f, 0x98, 0x47, 0x48, 0x87, 0xe6, 0xb0, 0x37, 0x1b, 0xfa, 0xa6, 0x82,
	0xce, 0x00, 0xfa, 0xb3, 0xf1, 0xec, 0xb6, 0xf7, 0x65, 0x34, 0xf7, 0xcd, 0xc6, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x24, 0xa5, 0x97, 0x08, 0x32, 0x04, 0x00, 0x00,
}
