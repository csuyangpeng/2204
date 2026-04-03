package metrics

import "lite5gc/cmn/metric"

// upf 统计指标 meter
const (
	// forward
	totalPacketsPerSec string = "Total Packets/Sec (P-I)"
	totalBitsPerSec    string = "Total Bits/Sec (P-I)"

	// up link
	upLinkPacketsSentPerSec     string = "UpLink Packets Sent/Sec (P-I)"
	upLinkPacketsReceivedPerSec string = "UpLink Packets Received/Sec (P-I)"
	upLinkBitsSentPerSec        string = "UpLink Bits Sent/Sec (P-I)"
	upLinkBitsReceivedPerSec    string = "UpLink Bits Received/Sec (P-I)"

	// down link
	downLinkPacketsSentPerSec     string = "DownLink Packets Sent/Sec (P-I)"
	downLinkPacketsReceivedPerSec string = "DownLink Packets Received/Sec (P-I)"

	downLinkBitsSentPerSec     string = "DownLink Bits Sent/Sec (P-I)"
	downLinkBitsReceivedPerSec string = "DownLink Bits Received/Sec (P-I)"
)

// Session 统计
// UpfSessionMetrics const
const (
	UpLinkSessionPacketsSentPerSec       string = "UpLink Session Packets Sent/Sec (P-I)"
	UpLinkSessionPacketsReceivedPerSec          = "UpLink Session Packets Received/Sec (P-I)"
	DownLinkSessionPacketsSentPerSec            = "DownLink Session Packets Sent/Sec (P-I)"
	DownLinkSessionPacketsReceivedPerSec        = "DownLink Session Packets Received/Sec (P-I)"
	UpLinkSessionBitsReceivedPerSec             = "UpLink Session Bits Received/Sec (P-I)"
	UpLinkSessionBitsSentPerSec                 = "UpLink Session Bits Sent/Sec (P-I)"
	DownLinkSessionBitsReceivedPerSec           = "DownLink Session Bits Received/Sec (P-I)"
	DownLinkSessionBitsSentPerSec               = "DownLink Session Bits Sent/Sec (P-I)"
)

// 用于初始化
var UpfSessionMetricsMeter = []string{
	UpLinkSessionPacketsSentPerSec,
	UpLinkSessionPacketsReceivedPerSec,
	DownLinkSessionPacketsSentPerSec,
	DownLinkSessionPacketsReceivedPerSec,
	UpLinkSessionBitsReceivedPerSec,
	UpLinkSessionBitsSentPerSec,
	DownLinkSessionBitsReceivedPerSec,
	DownLinkSessionBitsSentPerSec,
}

// upf module 统计指标全局变量(区间值)
type UpfModulePI struct {
	// forward
	TotalPacketsPerSec metric.Meter
	TotalBitsPerSec    metric.Meter

	// uplink
	UpLinkPacketsSentPerSec     metric.Meter
	UpLinkPacketsReceivedPerSec metric.Meter

	UpLinkBitsSentPerSec     metric.Meter
	UpLinkBitsReceivedPerSec metric.Meter

	// downlink
	DownLinkPacketsSentPerSec     metric.Meter
	DownLinkPacketsReceivedPerSec metric.Meter

	DownLinkBitsSentPerSec     metric.Meter
	DownLinkBitsReceivedPerSec metric.Meter
}
