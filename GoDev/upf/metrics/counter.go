package metrics

import (
	"lite5gc/cmn/metric"
	"net"
)

// upf 统计指标 counter
const (
	// forward
	totalPackets          string = "Total Packets"
	totalBits             string = "Total Bits"
	totalDiscardedPackets string = "Total Discarded Packets"

	// up link
	upLinkTotalPacketsSent     string = "UpLink Total Packets Sent"
	upLinkTotalPacketsReceived string = "UpLink Total Packets Received"

	upLinkTotalBitsSent                 string = "UpLink Total Bits Sent"
	upLinkTotalBitsReceived             string = "UpLink Total Bits Received"
	upLinkTotalReceivedPacketsDiscarded string = "UpLink Total Received Packets Discarded"
	upLinkTotalSentPacketsDiscarded     string = "UpLink Total Sent Packets Discarded"
	upLinkTotalInvalidPackets           string = "UpLink Total Invalid Packets"
	upLinkTotalDiscardedPackets         string = "UpLink Total Discarded Packets"

	// down link
	downLinkTotalPacketsSent     string = "DownLink Total Packets Sent"
	downLinkTotalPacketsReceived string = "DownLink Total Packets Received"

	downLinkTotalBitsSent                 string = "DownLink Total Bits Sent"
	downLinkTotalBitsReceived             string = "DownLink Total Bits Received"
	downLinkTotalReceivedPacketsDiscarded string = "DownLink Total Received Packets Discarded"
	downLinkTotalSentPacketsDiscarded     string = "DownLink Total Sent Packets Discarded"
	downLinkTotalInvalidPackets           string = "DownLink Total Invalid Packets"
	downLinkTotalDiscardedPackets         string = "DownLink Total Discarded Packets"
)

// Session 统计
// UpfSessionMetrics const
const (
	UpLinkSessionReceivedPacket   string = "UpLink Session Received Packet"
	UpLinkSessionSendPacket       string = "UpLink Session Send Packet"
	DownLinkSessionReceivedPacket string = "DownLink Session Received Packet"
	DownLinkSessionSendPacket     string = "DownLink Session Send Packet"

	UpLinkSessionBitsReceived   = "UpLink Session Bits Received"
	UpLinkSessionBitsSent       = "UpLink Session Bits Sent"
	DownLinkSessionBitsReceived = "DownLink Session Bits Received"
	DownLinkSessionBitsSent     = "DownLink Session Bits Sent"

	DownLinkSessionPacketsDiscarded = "DownLink Session Packets Discarded"
	UpLinkSessionPacketsDiscarded   = "UpLink Session Packets Discarded"

	// 当前及时值
	UpfSessionID                                  = "Upf Session ID"
	DownLinkSessionRuleID                         = "DownLink Session Rule ID"
	DownLinkSessionUEIPAddress                    = "DownLink Session UE IP Address"
	DownLinkSessionQoSFlowID                      = "DownLink Session QoS Flow ID"
	DownLinkSessionFlowDescription                = "DownLink Session Flow Description"
	DownLinkSessionForwardingAction               = "DownLink Session Forwarding Action"
	DownLinkSessionOuterHeaderCreationTEID        = "DownLink Session Outer Header Creation TEID"
	DownLinkSessionOuterHeaderCreationIPv4Address = "DownLink Session Outer Header Creation IPv4 Address"

	UpLinkSessionRuleID                        = "UpLink Session Rule ID"
	UpLinkSessionTrafficEndpointID             = "UpLink Session Traffic Endpoint ID"
	UpLinkSessionQoSFlowID                     = "UpLink Session QoS Flow ID"
	UpLinkSessionOuterHeaderRemovalDescription = "UpLink Session Outer Header Removal Description"
	UpLinkSessionForwardingAction              = "UpLink Session Forwarding Action"
	UpLinkSessionNetworkInstance               = "UpLink Session Network Instance"
	UpLinkSessionDNGatewayIPAddress            = "UpLink Session DN Gateway IP Address"
)

// 用于初始化
var UpfSessionMetricsCounter = []string{
	UpLinkSessionReceivedPacket,
	UpLinkSessionSendPacket,
	DownLinkSessionReceivedPacket,
	DownLinkSessionSendPacket,

	UpLinkSessionBitsReceived,
	UpLinkSessionBitsSent,
	DownLinkSessionBitsReceived,
	DownLinkSessionBitsSent,

	DownLinkSessionPacketsDiscarded,
	UpLinkSessionPacketsDiscarded,
	// 当前及时值
	UpfSessionID,
	DownLinkSessionRuleID,
	DownLinkSessionUEIPAddress,
	DownLinkSessionQoSFlowID,
	DownLinkSessionFlowDescription,
	DownLinkSessionForwardingAction,
	DownLinkSessionOuterHeaderCreationTEID,
	DownLinkSessionOuterHeaderCreationIPv4Address,

	UpLinkSessionRuleID,
	UpLinkSessionTrafficEndpointID,
	UpLinkSessionQoSFlowID,
	UpLinkSessionOuterHeaderRemovalDescription,
	UpLinkSessionForwardingAction,
	UpLinkSessionNetworkInstance,
	UpLinkSessionDNGatewayIPAddress,
}

type UpfSessionInfo struct {
	UpfSessionID                                  uint64 `json:SEID",omitempty"`
	DownLinkSessionRuleID                         uint16
	DownLinkSessionUEIPAddress                    net.IP `json:UE IP",omitempty"`
	DownLinkSessionQoSFlowID                      uint8
	DownLinkSessionFlowDescription                string
	DownLinkSessionForwardingAction               uint8
	DownLinkSessionOuterHeaderCreationTEID        uint32
	DownLinkSessionOuterHeaderCreationIPv4Address net.IP

	UpLinkSessionRuleID                        uint16
	UpLinkSessionTrafficEndpointID             uint32 `json:TEID",omitempty"`
	UpLinkSessionQoSFlowID                     uint8
	UpLinkSessionOuterHeaderRemovalDescription uint8
	UpLinkSessionForwardingAction              uint8
	UpLinkSessionNetworkInstance               string
	UpLinkSessionDNGatewayIPAddress            net.IP
}

// upf module 统计指标全局变量
type UpfModule struct {
	// forward
	TotalPackets                  metric.Counter
	TotalPacketsPrevious          int64
	TotalBits                     metric.Counter
	TotalBitsPrevious             int64
	TotalDiscardedPackets         metric.Counter
	TotalDiscardedPacketsPrevious int64

	// uplink
	UpLinkTotalPacketsSent             metric.Counter
	UpLinkTotalPacketsSentPrevious     int64
	UpLinkTotalPacketsReceived         metric.Counter
	UpLinkTotalPacketsReceivedPrevious int64

	UpLinkTotalBitsSent             metric.Counter
	UpLinkTotalBitsSentPrevious     int64
	UpLinkTotalBitsReceived         metric.Counter
	UpLinkTotalBitsReceivedPrevious int64

	UpLinkTotalReceivedPacketsDiscarded metric.Counter
	UpLinkTotalSentPacketsDiscarded     metric.Counter
	UpLinkTotalInvalidPackets           metric.Counter
	UpLinkTotalDiscardedPackets         metric.Counter

	// downlink
	DownLinkTotalPacketsSent             metric.Counter
	DownLinkTotalPacketsSentPrevious     int64
	DownLinkTotalPacketsReceived         metric.Counter
	DownLinkTotalPacketsReceivedPrevious int64

	DownLinkTotalBitsSent             metric.Counter
	DownLinkTotalBitsSentPrevious     int64
	DownLinkTotalBitsReceived         metric.Counter
	DownLinkTotalBitsReceivedPrevious int64

	DownLinkTotalReceivedPacketsDiscarded metric.Counter
	DownLinkTotalSentPacketsDiscarded     metric.Counter
	DownLinkTotalInvalidPackets           metric.Counter
	DownLinkTotalDiscardedPackets         metric.Counter
}
