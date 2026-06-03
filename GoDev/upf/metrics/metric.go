package metrics

import (
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/metric"
	"lite5gc/cmn/rlogger"
	"lite5gc/upf/context/n4context"
	"lite5gc/upf/cp/pdr"
	"time"
)

const TickDuration = 5

// open counter
var SessionCounterStart = true
var ModuleCounterStart = true

// upf session 统计使用get获取指标变量
// 每个session 一个统计表MetricItems
func SessionCounterInit() (metric.Registry, error) {
	MetricItems := metric.NewMetricsMap()
	metric.CountersCreate(UpfSessionMetricsCounter, MetricItems)
	metric.MetersCreate(UpfSessionMetricsMeter, MetricItems)

	return MetricItems, nil
}

// notes:Meter required release
func SessionMeterRelease(r metric.Registry) error {
	metric.MetersRelease(UpfSessionMetricsMeter, r)

	return nil
}

// upf session pps 计算
// UpfCounterInit 中启动 upf session pps 计算
func upfSessionPps() {
	tickChan := time.Tick(TickDuration * time.Second)
	for {
		select {
		case <-tickChan:
			// 遍历N4上下文，对每个session 计算 pps
			n4List, err := n4context.ValuesOfN4ContextTbl(n4context.N4SessionIDCxtType)
			if err != nil {
				rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Failed to get N4 Context:%s", err)
				return
			}
			for _, v := range n4List {
				// 在每个MetricItems中计算pps
				CalcPPS(v.MetricItems, v.MetricItemsSnapshot)
			}
		}
	}
}

// 查询时计算pps
func upfSessionPpsOne(seid uint64) metric.Registry {

	// 遍历N4上下文，对每个session 计算 pps
	n4Cxt, err := n4context.GetN4Context(n4context.N4SessionIDKey(seid))
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Failed to get N4 Context:%s", err)
		return nil
	}
	CalcPPS(n4Cxt.MetricItems, n4Cxt.MetricItemsSnapshot)
	tickChan := time.Tick(TickDuration * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case <-tickChan:
			CalcPPS(n4Cxt.MetricItems, n4Cxt.MetricItemsSnapshot)
		}
	}

	return n4Cxt.MetricItems
}

// 计算PPS 保存到MetricItems
func CalcPPS(MetricItems metric.Registry, Snapshot metric.Registry) {
	// 计算区间值P-I
	//  事件数Mark:当前计数 - 前一次计数
	//  session
	//  UpLinkSessionPacketsSentPerSec,
	IntervalIncrementPPS(UpLinkSessionSendPacket, UpLinkSessionPacketsSentPerSec, MetricItems, Snapshot)
	// UpLinkSessionPacketsReceivedPerSec
	IntervalIncrementPPS(UpLinkSessionReceivedPacket, UpLinkSessionPacketsReceivedPerSec, MetricItems, Snapshot)
	//	DownLinkSessionPacketsSentPerSec,
	IntervalIncrementPPS(DownLinkSessionSendPacket, DownLinkSessionPacketsSentPerSec, MetricItems, Snapshot)
	//	DownLinkSessionPacketsReceivedPerSec,
	IntervalIncrementPPS(DownLinkSessionReceivedPacket, DownLinkSessionPacketsReceivedPerSec, MetricItems, Snapshot)
	//	UpLinkSessionBitsReceivedPerSec,
	IntervalIncrementPPS(UpLinkSessionBitsReceived, UpLinkSessionBitsReceivedPerSec, MetricItems, Snapshot)
	//	UpLinkSessionBitsSentPerSec,
	IntervalIncrementPPS(UpLinkSessionBitsSent, UpLinkSessionBitsSentPerSec, MetricItems, Snapshot)
	//	DownLinkSessionBitsReceivedPerSec,
	IntervalIncrementPPS(DownLinkSessionBitsReceived, DownLinkSessionBitsReceivedPerSec, MetricItems, Snapshot)
	//	DownLinkSessionBitsSentPerSec,
	IntervalIncrementPPS(DownLinkSessionBitsSent, DownLinkSessionBitsSentPerSec, MetricItems, Snapshot)
}

// Calculated interval increment
func IntervalIncrementPPS(srcname string, dstname string,
	MetricItems metric.Registry, Snapshot metric.Registry) int64 {

	current := metric.Get(srcname, MetricItems).Count()
	PacketPI := current - metric.Get(srcname, Snapshot).Count()
	metric.GetMeter(dstname, MetricItems).Mark(PacketPI)
	Previous := metric.Get(srcname, Snapshot)
	Previous.Clear()
	Previous.Inc(current)

	return PacketPI
}

// session info
func UpfSessionInfoGet(seid uint64) *UpfSessionInfo {
	info := &UpfSessionInfo{}
	n4Cxt, err := n4context.GetN4Context(n4context.N4SessionIDKey(seid))
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Failed to get N4 Context:%s", err)
		return nil
	}
	info.UpfSessionID = n4Cxt.SEID

	// down link is pfcp.Core
	// todo 上下行只是读取一个PDR
	var farId uint32
	for _, pdrinfo := range n4Cxt.PDRs {
		if pdrinfo != nil {
			if pdrinfo.PDI.SourceInterface.InterfaceValue == pfcp.Core {
				info.DownLinkSessionRuleID, _ = pdrinfo.PDRID.Get()
				if pdrinfo.PDI.UEIPaddress != nil {
					info.DownLinkSessionUEIPAddress = pdrinfo.PDI.UEIPaddress.IPv4Addr.To4()
				}
				if pdrinfo.PDI.QFIs != nil && pdrinfo.PDI.QFIs[0] != nil {
					info.DownLinkSessionQoSFlowID, _ = pdrinfo.PDI.QFIs[0].Get()
				}
				if pdrinfo.PDI.SDFFilters != nil && pdrinfo.PDI.SDFFilters[0] != nil {
					info.DownLinkSessionFlowDescription = string(pdrinfo.PDI.SDFFilters[0].FlowDescription)
				}
				farId, _ = pdrinfo.FARID.Get()

				for _, far := range n4Cxt.FARs {
					if far.FARID.Value == farId {
						info.DownLinkSessionForwardingAction = far.ApplyAction.Flag
						if far.ForwardingParameters != nil && far.ForwardingParameters.OuterHeaderCreation != nil {
							info.DownLinkSessionOuterHeaderCreationTEID = uint32(far.ForwardingParameters.OuterHeaderCreation.TEID)
							info.DownLinkSessionOuterHeaderCreationIPv4Address = far.ForwardingParameters.OuterHeaderCreation.IPv4Addr
						}
					}
				}
			}
			if pdrinfo.PDI.SourceInterface.InterfaceValue == pfcp.Access {
				info.UpLinkSessionRuleID, _ = pdrinfo.PDRID.Get()
				if pdrinfo.PDI.LocalFTEID != nil {
					info.UpLinkSessionTrafficEndpointID = uint32(pdrinfo.PDI.LocalFTEID.TEID)
				}
				if pdrinfo.PDI.QFIs != nil && pdrinfo.PDI.QFIs[0] != nil {
					info.UpLinkSessionQoSFlowID, _ = pdrinfo.PDI.QFIs[0].Get()
				}
				if pdrinfo.OuterHeaderRemoval != nil {
					info.UpLinkSessionOuterHeaderRemovalDescription = pdrinfo.OuterHeaderRemoval.Description
				}
				farId, _ = pdrinfo.FARID.Get()

				for _, far := range n4Cxt.FARs {
					if far.FARID.Value == farId {
						info.UpLinkSessionForwardingAction = far.ApplyAction.Flag
						if far.ForwardingParameters != nil && far.ForwardingParameters.NetworkInstance != nil {
							info.UpLinkSessionNetworkInstance = far.ForwardingParameters.NetworkInstance.NetworkInstance
						}
						info.UpLinkSessionDNGatewayIPAddress = pdr.GetDnnGWIp(info.UpLinkSessionNetworkInstance)
					}
				}
			}
		}

	}

	return info
}

// upf module 统计使用全局指标变量
var UpfmoduleSet UpfModule
var UpfmodulePISet UpfModulePI

func UpfCounterInit() (metric.Registry, error) {
	// Dot switch
	//SessionCounterStart = configure.UpfConf.PM.SessionCounterStart
	//ModuleCounterStart = configure.UpfConf.PM.ModuleCounteStart

	MetricItems := metric.NewMetricsMap()
	// forward
	UpfmoduleSet.TotalPackets = metric.NewCounter(totalPackets, MetricItems)
	UpfmoduleSet.TotalBits = metric.NewCounter(totalBits, MetricItems)
	UpfmoduleSet.TotalDiscardedPackets = metric.NewCounter(totalDiscardedPackets, MetricItems)

	// up link
	UpfmoduleSet.UpLinkTotalPacketsReceived = metric.NewCounter(upLinkTotalPacketsReceived, MetricItems)
	UpfmoduleSet.UpLinkTotalPacketsSent = metric.NewCounter(upLinkTotalPacketsSent, MetricItems)
	UpfmoduleSet.UpLinkTotalBitsSent = metric.NewCounter(upLinkTotalBitsSent, MetricItems)
	UpfmoduleSet.UpLinkTotalBitsReceived = metric.NewCounter(upLinkTotalBitsReceived, MetricItems)
	UpfmoduleSet.UpLinkTotalReceivedPacketsDiscarded = metric.NewCounter(upLinkTotalReceivedPacketsDiscarded, MetricItems)
	UpfmoduleSet.UpLinkTotalSentPacketsDiscarded = metric.NewCounter(upLinkTotalSentPacketsDiscarded, MetricItems)
	UpfmoduleSet.UpLinkTotalInvalidPackets = metric.NewCounter(upLinkTotalInvalidPackets, MetricItems)
	UpfmoduleSet.UpLinkTotalDiscardedPackets = metric.NewCounter(upLinkTotalDiscardedPackets, MetricItems)

	// down link
	UpfmoduleSet.DownLinkTotalPacketsReceived = metric.NewCounter(downLinkTotalPacketsReceived, MetricItems)
	UpfmoduleSet.DownLinkTotalPacketsSent = metric.NewCounter(downLinkTotalPacketsSent, MetricItems)
	UpfmoduleSet.DownLinkTotalBitsSent = metric.NewCounter(downLinkTotalBitsSent, MetricItems)
	UpfmoduleSet.DownLinkTotalBitsReceived = metric.NewCounter(downLinkTotalBitsReceived, MetricItems)
	UpfmoduleSet.DownLinkTotalReceivedPacketsDiscarded = metric.NewCounter(downLinkTotalReceivedPacketsDiscarded, MetricItems)
	UpfmoduleSet.DownLinkTotalSentPacketsDiscarded = metric.NewCounter(downLinkTotalSentPacketsDiscarded, MetricItems)
	UpfmoduleSet.DownLinkTotalInvalidPackets = metric.NewCounter(downLinkTotalInvalidPackets, MetricItems)
	UpfmoduleSet.DownLinkTotalDiscardedPackets = metric.NewCounter(downLinkTotalDiscardedPackets, MetricItems)

	// pps
	// forward
	UpfmodulePISet.TotalPacketsPerSec = metric.NewMeter(totalPacketsPerSec, MetricItems)
	UpfmodulePISet.TotalBitsPerSec = metric.NewMeter(totalBitsPerSec, MetricItems)
	// up link
	UpfmodulePISet.UpLinkPacketsReceivedPerSec = metric.NewMeter(upLinkPacketsReceivedPerSec, MetricItems)
	UpfmodulePISet.UpLinkPacketsSentPerSec = metric.NewMeter(upLinkPacketsSentPerSec, MetricItems)
	UpfmodulePISet.UpLinkBitsSentPerSec = metric.NewMeter(upLinkBitsSentPerSec, MetricItems)
	UpfmodulePISet.UpLinkBitsReceivedPerSec = metric.NewMeter(upLinkBitsReceivedPerSec, MetricItems)
	// down link
	UpfmodulePISet.DownLinkPacketsReceivedPerSec = metric.NewMeter(downLinkPacketsReceivedPerSec, MetricItems)
	UpfmodulePISet.DownLinkPacketsSentPerSec = metric.NewMeter(downLinkPacketsSentPerSec, MetricItems)
	UpfmodulePISet.DownLinkBitsSentPerSec = metric.NewMeter(downLinkBitsSentPerSec, MetricItems)
	UpfmodulePISet.DownLinkBitsReceivedPerSec = metric.NewMeter(downLinkBitsReceivedPerSec, MetricItems)

	tickChan := time.Tick(TickDuration * time.Second)

	go func() {
		var PacketPI int64 // Value per interval

		for {
			select {
			case <-tickChan:
				// 计算区间值P-I
				// 事件数Mark:当前计数 - 前一次计数

				// uplink
				PacketPI = IntervalIncrement(
					UpfmoduleSet.UpLinkTotalPacketsReceived.Count(),
					&UpfmoduleSet.UpLinkTotalPacketsReceivedPrevious)
				UpfmodulePISet.UpLinkPacketsReceivedPerSec.Mark(PacketPI)

				PacketPI = IntervalIncrement(
					UpfmoduleSet.UpLinkTotalPacketsSent.Count(),
					&UpfmoduleSet.UpLinkTotalPacketsSentPrevious)
				UpfmodulePISet.UpLinkPacketsSentPerSec.Mark(PacketPI)

				PacketPI = IntervalIncrement(
					UpfmoduleSet.UpLinkTotalBitsSent.Count(),
					&UpfmoduleSet.UpLinkTotalBitsSentPrevious)
				UpfmodulePISet.UpLinkBitsSentPerSec.Mark(PacketPI * 8)

				PacketPI = IntervalIncrement(
					UpfmoduleSet.UpLinkTotalBitsReceived.Count(),
					&UpfmoduleSet.UpLinkTotalBitsReceivedPrevious)
				UpfmodulePISet.UpLinkBitsReceivedPerSec.Mark(PacketPI * 8)

				// downlink
				PacketPI = IntervalIncrement(
					UpfmoduleSet.DownLinkTotalPacketsReceived.Count(),
					&UpfmoduleSet.DownLinkTotalPacketsReceivedPrevious)
				UpfmodulePISet.DownLinkPacketsReceivedPerSec.Mark(PacketPI)

				PacketPI = IntervalIncrement(
					UpfmoduleSet.DownLinkTotalPacketsSent.Count(),
					&UpfmoduleSet.DownLinkTotalPacketsSentPrevious)
				UpfmodulePISet.DownLinkPacketsSentPerSec.Mark(PacketPI)

				PacketPI = IntervalIncrement(
					UpfmoduleSet.DownLinkTotalBitsSent.Count(),
					&UpfmoduleSet.DownLinkTotalBitsSentPrevious)
				UpfmodulePISet.DownLinkBitsSentPerSec.Mark(PacketPI * 8)

				PacketPI = IntervalIncrement(
					UpfmoduleSet.DownLinkTotalBitsReceived.Count(),
					&UpfmoduleSet.DownLinkTotalBitsReceivedPrevious)
				UpfmodulePISet.DownLinkBitsReceivedPerSec.Mark(PacketPI * 8)

				// forward
				// 计算转发总量
				UpfmoduleSet.TotalPackets.Clear()
				UpfmoduleSet.TotalPackets.Inc(UpfmoduleSet.UpLinkTotalPacketsSentPrevious + UpfmoduleSet.DownLinkTotalPacketsSentPrevious)

				UpfmoduleSet.TotalBits.Clear()
				UpfmoduleSet.TotalBits.Inc(UpfmoduleSet.UpLinkTotalBitsSentPrevious + UpfmoduleSet.DownLinkTotalBitsSentPrevious)

				UpfmoduleSet.UpLinkTotalDiscardedPackets.Clear()
				UpfmoduleSet.UpLinkTotalDiscardedPackets.Inc(
					UpfmoduleSet.UpLinkTotalReceivedPacketsDiscarded.Count() +
						UpfmoduleSet.UpLinkTotalSentPacketsDiscarded.Count() +
						UpfmoduleSet.UpLinkTotalInvalidPackets.Count())

				UpfmoduleSet.DownLinkTotalDiscardedPackets.Clear()
				UpfmoduleSet.DownLinkTotalDiscardedPackets.Inc(UpfmoduleSet.DownLinkTotalReceivedPacketsDiscarded.Count() +
					UpfmoduleSet.DownLinkTotalSentPacketsDiscarded.Count() +
					UpfmoduleSet.DownLinkTotalInvalidPackets.Count())

				UpfmoduleSet.TotalDiscardedPackets.Clear()
				UpfmoduleSet.TotalDiscardedPackets.Inc(UpfmoduleSet.UpLinkTotalDiscardedPackets.Count() + UpfmoduleSet.DownLinkTotalDiscardedPackets.Count())

				PacketPI = IntervalIncrement(
					UpfmoduleSet.TotalPackets.Count(),
					&UpfmoduleSet.TotalPacketsPrevious)
				UpfmodulePISet.TotalPacketsPerSec.Mark(PacketPI)

				PacketPI = IntervalIncrement(
					UpfmoduleSet.TotalBits.Count(),
					&UpfmoduleSet.TotalBitsPrevious)
				UpfmodulePISet.TotalBitsPerSec.Mark(PacketPI * 8)

			}
		}
	}()
	// session pps
	go upfSessionPps()
	// 上报到lems
	// StartPMReport()

	return MetricItems, nil
}

// Calculated interval increment
func IntervalIncrement(Current int64, Previous *int64) int64 {
	PacketPI := Current - *Previous
	*Previous = Current
	return PacketPI
}

func Log(r metric.Registry) {
	//metric.Log(r, 5e9)
	//r.Get(metric.UpLinkReceivedPacketPerS).(metrics.Meter).Count()
	//
	/*pms := webTypes.NewNfPerformanceData()
	pms.NfNo = configure.LemsUpfAgentConf.NfNo.UpfNfNo
	pms.Params = append(pms.Params, webTypes.Param{111, 100})
	*/ //agent.StartPerformanceReport(5e9)
	//time.Sleep(10e9)
	//agent.StopPerformanceReport()
}

func SessionLog(seid uint64) {

	if SessionCounterStart == true {

		n4cxt, err := n4context.GetN4Context(n4context.N4SessionIDKey(seid))
		if err != nil {
			rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Failed to get N4 Context:%s", err)
			return
		}
		metric.Log(n4cxt.MetricItems, 7e9)
	}
}

func GetSeidFromIp(ip string) (uint64, error) {

	return pdr.GetSeidFromIp(ip)
}

func GetSeidFromTeid(teid uint32) (uint64, error) {

	return pdr.GetSeidFromTeid(teid)
}
