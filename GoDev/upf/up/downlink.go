package up

import (
	"github.com/intel-go/nff-go/flow"
	"github.com/intel-go/nff-go/packet"
	flowT "github.com/intel-go/nff-go/types"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/metric"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/upf/adapter"
	"lite5gc/upf/context/n4context"
	"lite5gc/upf/context/pdrcontext"
	"lite5gc/upf/cp/pdr"
	"lite5gc/upf/defs"
	"lite5gc/upf/metrics"
	"lite5gc/upf/service"
	"lite5gc/upf/utils"
	"sync/atomic"
)

func NffStartDPERawHandleVector(currentPackets []*packet.Packet, mask *[vecSize]bool, notDrop *[vecSize]bool, cxt flow.UserContext) {
	rlogger.FuncEntry(moduleTag, nil)
	n := uint(0)
	for i := uint(0); i < vecSize; i++ {
		if (*mask)[i] == true {
			isDrop := NffStartDPERawHandle(currentPackets[i], cxt)
			if !isDrop {
				notDrop[i] = true
				n++
			}
		}
	}
}

// nff 处理n6消息
func NffStartDPERawHandle(currentPacket *packet.Packet, cxt flow.UserContext) bool {
	retVal := true

	rlogger.FuncEntry(moduleTag, nil)

	result := NffN6Handler(currentPacket, cxt)
	if !result {
		return false
	}

	N6Cxt := cxt.(HandlerContext)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "NffN6Handler hc %p", &N6Cxt)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "NffN6Handler hc %+v", N6Cxt)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "NffN6Handler hc %+v,\n%+v", N6Cxt.Ipport, N6Cxt.Msgcxt)

	conutPacket := 0 //test count
	var EncodeMsg []byte
	// processing N6 interface Rawsocket messages
	// N3 编码后，发送消息到RAN,N3接口对端
	RawMsg := currentPacket.GetRawPacketBytes()
	// 创建包头指示
	N6Cxt.Msgcxt.Rw.RLock()
	if N6Cxt.Msgcxt.OuterHeaderDesc == pfcp.IEOuterHCreation_GTPU_UDP_IPv4 {
		N6Cxt.Msgcxt.Rw.RUnlock()
		GtpMsg, err := adapter.SendN3MsgHandleExt(RawMsg[flowT.EtherLen:], N6Cxt.Msgcxt)
		rlogger.Trace(moduleTag, rlogger.DEBUG, N6Cxt.Msgcxt, "Receive N6 Message len=%d value: %#x\n", len(GtpMsg), GtpMsg)
		if err != nil {
			//fmt.Println(err)
			rlogger.Trace(moduleTag, rlogger.ERROR, N6Cxt.Msgcxt, "Failed to N6 message encode!")
			// N3EncodeMsg failed,discard message
			return false
		}
		EncodeMsg = GtpMsg
		// 设置转发目的 IP
		//server.GNbIp = N6Cxt.Msgcxt.GnbIP
		rlogger.Trace(moduleTag, rlogger.DEBUG, N6Cxt.Msgcxt, "Receive N6 Message MsgCxt value: %+v\n", *N6Cxt.Msgcxt)
	} else {
		N6Cxt.Msgcxt.Rw.RUnlock()
		EncodeMsg = RawMsg
		rlogger.Trace(moduleTag, rlogger.ERROR, N6Cxt.Msgcxt, "Receive N6 Message OuterHeaderDesc fault:%d", N6Cxt.Msgcxt.OuterHeaderDesc)
	}

	// Send to network ，发送到N3对端
	//SendN3UdpMsg(msg, server)
	//performance improvement
	//SendN3UdpMsgExtConn(msg, server)

	conutPacket++
	//fmt.Printf("Send n3 DPENo[%d]:%d \n", DPENo, conutPacket)
	//rlogger.Trace(moduleTag, rlogger.DEBUG, RawMsg.Msgcxt, "Send n3 DPENo[%d]:%d \n", DPENo, conutPacket)

	// 加gtp包头
	//currentPacket.EncapsulateIPv4GTP()
	etherLen := uint(flowT.EtherLen)
	ipv4MinLen := uint(flowT.IPv4MinLen)
	UDPLen := uint(flowT.UDPLen)
	GTPMinLen := uint(flowT.GTPMinLen)
	GTPMinLen = 20
	rlogger.Trace(moduleTag, rlogger.DEBUG, N6Cxt.Msgcxt, "add gtp head,etherLen-ipv4MinLen-UDPLen-GTPMinLen", etherLen, ipv4MinLen, UDPLen, GTPMinLen)
	//length := currentPacket.GetPacketLen() - flowT.EtherLen
	if !currentPacket.EncapsulateHead(etherLen, ipv4MinLen+UDPLen+GTPMinLen) {
		rlogger.Trace(moduleTag, rlogger.ERROR, N6Cxt.Msgcxt, "add gtp head ,failed to EncapsulateHead")
		return false
	}
	currentPacket.ParseL3()
	//gtp := (*packet.GTPHdr)(unsafe.Pointer(uintptr(currentPacket.L3) + flowT.IPv4MinLen + flowT.UDPLen))
	//fill gtp
	//非固定头填充
	currentPacket.PacketBytesChange(etherLen+ipv4MinLen+UDPLen, EncodeMsg)
	rlogger.Trace(moduleTag, rlogger.DEBUG, N6Cxt.Msgcxt, "change N6 Message len=%d value: %#x\n", len(EncodeMsg), EncodeMsg)

	// 固定头填充
	//gtp.HeaderType = 0x34          // 001 - GTPv1, 1 - not GTP', 0 - reserved, 000 - no optional fields
	//gtp.MessageType = packet.G_PDU // encapsulated user message
	//gtp.MessageLength = packet.SwapBytesUint16(uint16(length))
	//gtp.TEID = packet.SwapBytesUint32(uint32(N6Cxt.Msgcxt.GnbTEID))
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "n6cxt %+v", N6Cxt)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "n6cxt %+v", N6Cxt.Ipport)

	// Fill new IPv4 header with addresses according to context
	ipv4 := currentPacket.GetIPv4NoCheck() //(*packet.IPv4Hdr)(currentPacket.L3)
	// 源IP，目的IP
	srcAddr := N6Cxt.Ipport.Subnet.IPv4.Addr
	N6Cxt.Msgcxt.Rw.RLock()
	dstAddr := flowT.SliceToIPv4(N6Cxt.Msgcxt.GnbIP.To4())
	N6Cxt.Msgcxt.Rw.RUnlock()
	ipv4.SrcAddr = srcAddr
	ipv4.DstAddr = dstAddr
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "ipv4", ipv4.SrcAddr, ipv4.DstAddr)

	// Fill L2
	currentPacket.Ether.EtherType = flowT.SwapIPV4Number
	currentPacket.Ether.SAddr = N6Cxt.Ipport.MacAddress
	if N6Cxt.Ipport.StaticARP {
		currentPacket.Ether.DAddr = N6Cxt.Ipport.DstMacAddress
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "ipv4 Static currentPacket.Ether.DAddr", currentPacket.Ether.DAddr)
	} else {
		// Find l2 addresses for new destionation IP in ARP cache
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "n6 arp table %s", N6Cxt.Ipport.NeighCache.PrintArpTable())
		targetIP := ipv4.DstAddr
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "lookup ipv4 targetIP %s", dstAddr)
		targetMAC, found := N6Cxt.Ipport.NeighCache.LookupMACForIPv4(targetIP)
		if !found {
			// fmt.Println("Not found MAC address for IP", targetIP.String())
			N6Cxt.Ipport.NeighCache.SendARPRequestForIPv4(dstAddr, srcAddr, 0)
			rlogger.Trace(moduleTag, rlogger.ERROR, nil, "found n3dstIp %s,n3srcIp %s", dstAddr, srcAddr)
			rlogger.Trace(moduleTag, rlogger.ERROR, nil, "found n3src MAC %s", N6Cxt.Ipport.MacAddress)

			retVal = false
			//return false
		}
		currentPacket.Ether.DAddr = targetMAC
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "lookup ipv4 targetMAC %s", targetMAC)
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "ipv4 currentPacket.Ether.DAddr %s", currentPacket.Ether.DAddr)

	}

	// fill ip
	// Fill up l3
	ipv4.VersionIhl = 0x45
	ipv4.TypeOfService = 0
	ipv4.PacketID = 0xe803
	ipv4.FragmentOffset = 0
	ipv4.TimeToLive = 64

	length := currentPacket.GetPacketLen()
	ipv4.TotalLength = packet.SwapBytesUint16(uint16(length - flowT.EtherLen))
	ipv4.NextProtoID = flowT.UDPNumber
	ipv4.HdrChecksum = packet.SwapBytesUint16(packet.CalculateIPv4Checksum(ipv4))
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "ipv4 head %+v", ipv4)

	// fill udp
	// Fill up L4
	currentPacket.ParseL4ForIPv4()
	udp := currentPacket.GetUDPForIPv4()
	udp.SrcPort = packet.SwapUDPPortGTPU
	udp.DstPort = packet.SwapUDPPortGTPU
	udp.DgramLen = packet.SwapBytesUint16(uint16(length - flowT.EtherLen - flowT.IPv4MinLen))
	currentPacket.ParseL7(flowT.UDPNumber)
	// Calculate checksums
	ipv4.HdrChecksum = packet.SwapBytesUint16(packet.CalculateIPv4Checksum(ipv4))
	udp.DgramCksum = packet.SwapBytesUint16(packet.CalculateIPv4UDPChecksum(ipv4, udp, currentPacket.Data))
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "ipv4 udp head %+v", udp)

	if (!retVal) {
		// 没有find, 将报文先保存起来，如果缓存中的报文数量超过200个，将旧的释放掉
		if (packet.GArpBuffers.Len() > 200){
			packet.GArpMutex.Lock()
			bufferHead := packet.GArpBuffers.Front()
			packet.GArpBuffers.Remove(bufferHead)
			packet.GArpMutex.Unlock()
			oldPacket := bufferHead.Value.(* packet.Packet)
			oldPacket.FreePacket()
		}
		rlogger.Trace(moduleTag, rlogger.INFO, nil, "[2] save packet.Ether", currentPacket.Ether)
		srcbuf := currentPacket.GetRawPacketBytes()
		// 从arp buf pool里申请一个mbuf内存缓存起来
		arpPacket, _ := packet.NewArpBufPacket()		
		arpbuf := (*[2 << 10]byte)(arpPacket.StartAtOffset(0))
		copy(arpbuf[0:], srcbuf)
		arpPacket.PacketSetDataLen(uint(len(srcbuf)))
		arpPacket.ParseL3()

		packet.GArpMutex.Lock()
		packet.GArpBuffers.PushBack(arpPacket)
		packet.GArpMutex.Unlock()
	}

	return retVal
}

// n6消息保留，返回true
func NffN6Handler(currentPacket *packet.Packet, cxt flow.UserContext) bool {
	hc := cxt.(HandlerContext)
	//true 保留在f0
	currentPacket.ParseL3()
	ipv4 := currentPacket.GetIPv4()
	//fmt.Printf("NffN3Handler1 %v", ipv4)

	if ipv4 != nil {
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "NffN6Handler hc %p", &hc)
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "NffN6Handler hc %+v", hc)
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "NffN6Handler hc %+v,%+v", hc.Ipport, hc.Msgcxt)
		currentPacket.ParseL4ForIPv4()
		/*		fmt.Printf("NffN6Handler %v", currentPacket.GetUDPForIPv4())
				fmt.Printf("NffN6Handler %v\n", currentPacket.GetUDPNoCheck().DstPort)
				fmt.Println("NffN6Handler:", packet.SwapUDPPortGTPU)
		*/
		buf := currentPacket.GetRawPacketBytes()
		return nffN6ReceiveHandler(buf, hc.Msgcxt)
	}
	//fmt.Println("----------------------------------------------------------")
	//NffStartDPEHandle()
	return false
}

// nff n6 handle
func nffN6ReceiveHandler(buf []byte, msgCxt *pdrcontext.DataFlowContext) bool {
	rlogger.FuncEntry(moduleTag, nil)
	//true 保留在f0
	countPacket := 0
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "nffN6ReceiveHandler msgCxt %p", msgCxt)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "nffN6ReceiveHandler msgCxt %v", msgCxt)
	// performance improvement pprof
	//buf := make([]byte, RAW_BUFFER)
	//var buf [RAW_BUFFER]byte
	// performance improvement pprof end
	n := len(buf)
	if n < 14 {
		//fmt.Println("the ethernet header length < 14")
		rlogger.Trace(moduleTag, rlogger.WARN, nil, "the ethernet header length < 14")
		//todo counter
		metrics.UpfmoduleSet.DownLinkTotalInvalidPackets.Inc(1)
		return false
	}
	atomic.AddUint64(&defs.N6CountReceivePacket, 1)

	// 获取IP五元组
	//fiveTuple, err := utils.IpFiveTuple(buf[14:])
	// performance improvement performance improvement
	fiveTuple, err := utils.IpFiveTuple(buf[14:])
	//fiveTuple, err := utils.IpFiveTuplePerf(buf[14:])
	// performance improvement performance improvement end
	if err != nil {
		//todo counter
		metrics.UpfmoduleSet.DownLinkTotalInvalidPackets.Inc(1)
		return false
	}
	// mac padding delete
	iplength := int(fiveTuple.Length)
	n = 14 + iplength
	// 下行方向
	fiveTuple.Direction = nasie.DownlinkOnly
	//应用过滤规则
	//1.1 upgraded version 201812
	rule, err := pdr.LookupPDRs(fiveTuple)
	if err != nil || rule == nil {
		// 没有匹配规则的包丢弃
		//todo counter
		metrics.UpfmoduleSet.DownLinkTotalInvalidPackets.Inc(1)
		return false
	}
	//todo session counter
	var n4Cxt *n4context.N4SessionContext
	var errN4 error
	if metrics.SessionCounterStart == true {

		n4Cxt, errN4 = n4context.GetN4Context(n4context.N4SessionIDKey(rule.SEID))
		if errN4 != nil {
			rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Failed to get N4 Context:%s", errN4)
		}
		if n4Cxt != nil {
			metric.MetricItemAdd(metrics.DownLinkSessionReceivedPacket, n4Cxt.MetricItems)
			metric.MetricItemAddX(metrics.DownLinkSessionBitsReceived, n4Cxt.MetricItems, int64(n-14))
		}
	}
	//todo counter
	metrics.UpfmoduleSet.DownLinkTotalPacketsReceived.Inc(1)
	metrics.UpfmoduleSet.DownLinkTotalBitsReceived.Inc(int64(n - 14))

	var flowCxt pdrcontext.DataFlowContext
	flowCxt.SEID = rule.SEID
	flowCxt.RuleID = rule.RuleID
	flowCxt.UEIP = fiveTuple.DstIp
	flowCxt.UEPort = fiveTuple.DstPort

	// 应用转发规则
	if rule.FarI.Far != nil {
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "Action:%v", rule.FarI.Far.Action)
		switch rule.FarI.Far.Action {
		case pdr.FORW:
			flowCxt.DP.QFI = rule.PDI.QFI
			if rule.FarI.Far.OuterHeaderCreation != nil {
				flowCxt.GnbTEID = rule.FarI.Far.OuterHeaderCreation.TEID
				flowCxt.GnbIP = rule.FarI.Far.OuterHeaderCreation.IPv4Addr
				flowCxt.OuterHeaderDesc = rule.FarI.Far.OuterHeaderCreation.Description
			} else {
				rlogger.Trace(moduleTag, rlogger.ERROR, &flowCxt, "No applicable forwarding OuterHeaderCreation")
				rlogger.Trace(moduleTag, rlogger.ERROR, &flowCxt, "No applicable forwarding OuterHeaderCreation %v", rule.PDI.SourceInterface) //DownlinkOnly  PacketFilterDirection = 1
				rlogger.Trace(moduleTag, rlogger.ERROR, &flowCxt, "No applicable forwarding OuterHeaderCreation %v", rule.FarI.FarID)
				rlogger.Trace(moduleTag, rlogger.ERROR, &flowCxt, "No applicable forwarding OuterHeaderCreation %d,%s", rule.SEID, rule.PDI.UEIPAddress)
				rlogger.Trace(moduleTag, rlogger.ERROR, &flowCxt, "No applicable forwarding OuterHeaderCreation %d,%+v", rule.SEID, rule)

			}
			// 发送到N3 chan
		case pdr.DROP: //dropped
			//todo counter
			metrics.UpfmoduleSet.DownLinkTotalSentPacketsDiscarded.Inc(1)
			if metrics.SessionCounterStart == true && n4Cxt != nil {
				metric.MetricItemAdd(metrics.DownLinkSessionPacketsDiscarded, n4Cxt.MetricItems)
			}
			return false
		case pdr.BUFF:
			//	todo：
			return false
		case pdr.DUPL:
			//	todo：
			return false
		case pdr.NOCPBUFF: // buff and notify,deactivate up DL
			rlogger.Trace(moduleTag, rlogger.WARN, &flowCxt, "applicable forwarding action:DL deactivate")
			// paging support
			flowCxt.DP.QFI = rule.PDI.QFI
			flowCxt.SEID = rule.SEID
			flowCxt.RuleID = rule.RuleID
			//todo nff support buffer
			service.PagingServer.ReceiveBuffMsg(buf[14:n], &flowCxt)

			metrics.UpfmoduleSet.DownLinkTotalSentPacketsDiscarded.Inc(1)
			if metrics.SessionCounterStart == true && n4Cxt != nil {
				metric.MetricItemAdd(metrics.DownLinkSessionPacketsDiscarded, n4Cxt.MetricItems)
			}
			return false
		default:
			// 默认转发 到N3
			flowCxt.DP.QFI = rule.PDI.QFI
			flowCxt.GnbTEID = 0
			//todo nff no support
			//flowCxt.GnbIP = RawServer.GNbIp
			flowCxt.OuterHeaderDesc = pfcp.IEOuterHCreation_GTPU_UDP_IPv4
			rlogger.Trace(moduleTag, rlogger.ERROR, &flowCxt, "No applicable forwarding action")
			rlogger.Trace(moduleTag, rlogger.ERROR, &flowCxt, "No applicable forwarding action %d,%+v", rule.SEID, rule)
			return false
		}
	} else {
		rlogger.Trace(moduleTag, rlogger.ERROR, &flowCxt, "No applicable forwarding rules")
		//todo counter
		metrics.UpfmoduleSet.DownLinkTotalSentPacketsDiscarded.Inc(1)
		if metrics.SessionCounterStart == true && n4Cxt != nil {
			metric.MetricItemAdd(metrics.DownLinkSessionPacketsDiscarded, n4Cxt.MetricItems)
		}
		return false
	}

	//rlogger.Trace(moduleTag, rlogger.INFO, nil,  "Receive N6Msg from DN : <%s>--><%s>: %#x\n", fiveTuple.SrcIp.String(), fiveTuple.DstIp.String(), buf[14:n])
	countPacket++
	rlogger.Trace(moduleTag, rlogger.DEBUG, &flowCxt, "Receive N6Msg from DN %d\n", countPacket)
	//todo session counter
	if metrics.SessionCounterStart == true && n4Cxt != nil {
		metric.MetricItemAdd(metrics.DownLinkSessionSendPacket, n4Cxt.MetricItems)
		metric.MetricItemAddX(metrics.DownLinkSessionBitsSent, n4Cxt.MetricItems, int64(n-14))
	}
	msgCxt.Rw.Lock()
	msgCxt.SEID = flowCxt.SEID
	msgCxt.RuleID = flowCxt.RuleID
	msgCxt.UEIP = flowCxt.UEIP
	msgCxt.UEPort = flowCxt.UEPort
	msgCxt.GnbTEID = flowCxt.GnbTEID
	msgCxt.GnbIP = flowCxt.GnbIP
	msgCxt.OuterHeaderDesc = flowCxt.OuterHeaderDesc
	msgCxt.DP = flowCxt.DP

	msgCxt.Rw.Unlock()

	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "nffN6ReceiveHandler msgCxt %p", msgCxt)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "nffN6ReceiveHandler msgCxt", msgCxt)
	return true
}
