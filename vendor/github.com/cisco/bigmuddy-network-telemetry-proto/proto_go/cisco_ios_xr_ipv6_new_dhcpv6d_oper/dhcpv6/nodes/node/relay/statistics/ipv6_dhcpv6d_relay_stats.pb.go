// Code generated by protoc-gen-go.
// source: ipv6_dhcpv6d_relay_stats.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_statistics is a generated protocol buffer package.

It is generated from these files:
	ipv6_dhcpv6d_relay_stats.proto

It has these top-level messages:
	Ipv6Dhcpv6DRelayStats_KEYS
	Ipv6Dhcpv6DRelayStats
	Ipv6Dhcpv6DRelayStatsItem
	Ipv6Dhcpv6DFilteredStats
*/
package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_statistics

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// DHCPv6 relay statistics
type Ipv6Dhcpv6DRelayStats_KEYS struct {
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
}

func (m *Ipv6Dhcpv6DRelayStats_KEYS) Reset()                    { *m = Ipv6Dhcpv6DRelayStats_KEYS{} }
func (m *Ipv6Dhcpv6DRelayStats_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DRelayStats_KEYS) ProtoMessage()               {}
func (*Ipv6Dhcpv6DRelayStats_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv6Dhcpv6DRelayStats_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type Ipv6Dhcpv6DRelayStats struct {
	Ipv6Dhcpv6DRelayStats []*Ipv6Dhcpv6DRelayStatsItem `protobuf:"bytes,50,rep,name=ipv6_dhcpv6d_relay_stats,json=ipv6Dhcpv6dRelayStats" json:"ipv6_dhcpv6d_relay_stats,omitempty"`
}

func (m *Ipv6Dhcpv6DRelayStats) Reset()                    { *m = Ipv6Dhcpv6DRelayStats{} }
func (m *Ipv6Dhcpv6DRelayStats) String() string            { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DRelayStats) ProtoMessage()               {}
func (*Ipv6Dhcpv6DRelayStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6Dhcpv6DRelayStats) GetIpv6Dhcpv6DRelayStats() []*Ipv6Dhcpv6DRelayStatsItem {
	if m != nil {
		return m.Ipv6Dhcpv6DRelayStats
	}
	return nil
}

type Ipv6Dhcpv6DRelayStatsItem struct {
	// DHCPv6 L3 VRF name
	VrfName string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	// Relay statistics
	Statistics *Ipv6Dhcpv6DFilteredStats `protobuf:"bytes,2,opt,name=statistics" json:"statistics,omitempty"`
}

func (m *Ipv6Dhcpv6DRelayStatsItem) Reset()                    { *m = Ipv6Dhcpv6DRelayStatsItem{} }
func (m *Ipv6Dhcpv6DRelayStatsItem) String() string            { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DRelayStatsItem) ProtoMessage()               {}
func (*Ipv6Dhcpv6DRelayStatsItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ipv6Dhcpv6DRelayStatsItem) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6Dhcpv6DRelayStatsItem) GetStatistics() *Ipv6Dhcpv6DFilteredStats {
	if m != nil {
		return m.Statistics
	}
	return nil
}

// DHCPv6 filtered statistics
type Ipv6Dhcpv6DFilteredStats struct {
	// Received packets
	ReceivedPackets uint64 `protobuf:"varint,1,opt,name=received_packets,json=receivedPackets" json:"received_packets,omitempty"`
	// Transmitted packets
	TransmittedPackets uint64 `protobuf:"varint,2,opt,name=transmitted_packets,json=transmittedPackets" json:"transmitted_packets,omitempty"`
	// Dropped packets
	DroppedPackets uint64 `protobuf:"varint,3,opt,name=dropped_packets,json=droppedPackets" json:"dropped_packets,omitempty"`
}

func (m *Ipv6Dhcpv6DFilteredStats) Reset()                    { *m = Ipv6Dhcpv6DFilteredStats{} }
func (m *Ipv6Dhcpv6DFilteredStats) String() string            { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DFilteredStats) ProtoMessage()               {}
func (*Ipv6Dhcpv6DFilteredStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Ipv6Dhcpv6DFilteredStats) GetReceivedPackets() uint64 {
	if m != nil {
		return m.ReceivedPackets
	}
	return 0
}

func (m *Ipv6Dhcpv6DFilteredStats) GetTransmittedPackets() uint64 {
	if m != nil {
		return m.TransmittedPackets
	}
	return 0
}

func (m *Ipv6Dhcpv6DFilteredStats) GetDroppedPackets() uint64 {
	if m != nil {
		return m.DroppedPackets
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DRelayStats_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.relay.statistics.ipv6_dhcpv6d_relay_stats_KEYS")
	proto.RegisterType((*Ipv6Dhcpv6DRelayStats)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.relay.statistics.ipv6_dhcpv6d_relay_stats")
	proto.RegisterType((*Ipv6Dhcpv6DRelayStatsItem)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.relay.statistics.ipv6_dhcpv6d_relay_stats_item")
	proto.RegisterType((*Ipv6Dhcpv6DFilteredStats)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.relay.statistics.ipv6_dhcpv6d_filtered_stats")
}

func init() { proto.RegisterFile("ipv6_dhcpv6d_relay_stats.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x3f, 0x4b, 0x3b, 0x31,
	0x18, 0xc7, 0x49, 0xfb, 0xe3, 0x67, 0x9b, 0x82, 0x95, 0x88, 0x70, 0x52, 0x94, 0x72, 0x8b, 0x75,
	0x89, 0x50, 0xa1, 0x93, 0xa3, 0x9d, 0x04, 0x91, 0xeb, 0xe4, 0x14, 0xd2, 0xe4, 0x39, 0x0c, 0xf6,
	0x2e, 0x21, 0x09, 0xa7, 0xae, 0xbe, 0x0f, 0x17, 0x5f, 0x86, 0x8b, 0x6f, 0x4d, 0x92, 0x6b, 0xf5,
	0x96, 0xbb, 0x49, 0x97, 0x83, 0xfb, 0xfe, 0x79, 0x9e, 0xe7, 0x03, 0xc1, 0xa7, 0xca, 0x54, 0x0b,
	0x26, 0x1f, 0x84, 0xa9, 0x16, 0x92, 0x59, 0xd8, 0xf0, 0x17, 0xe6, 0x3c, 0xf7, 0x8e, 0x1a, 0xab,
	0xbd, 0x26, 0x4b, 0xa1, 0x9c, 0xd0, 0x4c, 0x69, 0xc7, 0x9e, 0x2d, 0x8b, 0xe1, 0x12, 0x9e, 0xbe,
	0x0b, 0xda, 0x80, 0xa5, 0xf5, 0x0f, 0x2d, 0xb5, 0x04, 0x17, 0xbf, 0x34, 0xce, 0xa1, 0x61, 0x8e,
	0x72, 0x5e, 0x09, 0x97, 0x5e, 0xe1, 0x93, 0xb6, 0x45, 0xec, 0x66, 0x79, 0xbf, 0x22, 0x13, 0x3c,
	0x0c, 0x4d, 0x56, 0xf2, 0x02, 0x12, 0x34, 0x45, 0xb3, 0x61, 0x36, 0x08, 0xc2, 0x2d, 0x2f, 0x20,
	0xfd, 0x40, 0x38, 0x69, 0xab, 0x93, 0xb7, 0x0e, 0x33, 0x99, 0x4f, 0xfb, 0xb3, 0xd1, 0x5c, 0xd2,
	0x5f, 0xa1, 0xa0, 0xad, 0x08, 0xca, 0x43, 0x91, 0x1d, 0x05, 0xfb, 0xba, 0x76, 0xb3, 0x60, 0xae,
	0x82, 0x97, 0x7e, 0xa2, 0x0e, 0xf6, 0x50, 0x24, 0xc7, 0x78, 0x50, 0xd9, 0xbc, 0x89, 0xbe, 0x57,
	0xd9, 0x3c, 0x90, 0x93, 0x57, 0x84, 0xf1, 0xcf, 0x01, 0x49, 0x6f, 0x8a, 0x66, 0xa3, 0xf9, 0xfa,
	0x2f, 0x70, 0x72, 0xb5, 0xf1, 0x60, 0x41, 0xd6, 0x87, 0x65, 0x8d, 0xad, 0xe9, 0x3b, 0xc2, 0x93,
	0x8e, 0x2c, 0x39, 0xc7, 0x07, 0x16, 0x04, 0xa8, 0x0a, 0x24, 0x33, 0x5c, 0x3c, 0x82, 0x77, 0x91,
	0xe3, 0x5f, 0x36, 0xde, 0xe9, 0x77, 0xb5, 0x4c, 0x2e, 0xf0, 0xa1, 0xb7, 0xbc, 0x74, 0x85, 0xf2,
	0xbe, 0x91, 0xee, 0xc5, 0x34, 0x69, 0x58, 0xbb, 0xc2, 0x19, 0x1e, 0x4b, 0xab, 0x8d, 0x69, 0x84,
	0xfb, 0x31, 0xbc, 0xbf, 0x95, 0xb7, 0xc1, 0xf5, 0xff, 0xf8, 0x5e, 0x2f, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xf3, 0x61, 0x14, 0xa2, 0xd1, 0x02, 0x00, 0x00,
}
