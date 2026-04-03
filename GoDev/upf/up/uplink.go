package up

import (
	"github.com/intel-go/nff-go/flow"
	"github.com/intel-go/nff-go/packet"
	flowT "github.com/intel-go/nff-go/types"
	"lite5gc/cmn/message/gtpv1u"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/metric"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types3gpp"
	"lite5gc/upf/context/n4context"
	"lite5gc/upf/context/pdrcontext"
	"lite5gc/upf/cp/pdr"
	. "lite5gc/upf/defs"
	"lite5gc/upf/metrics"
	"lite5gc/upf/utils"
	"net"
	"os"
)

// n3 消息保留，n3 返回true
func NffN3Handler(currentPacket *packet.Packet, context flow.UserContext) bool {
	//true 保留在f0
	//fmt.Println("----------------------------------------------------------")
	//	packet.L4 udp 2152
	//fmt.Println("NffN3Handler")
	currentPacket.ParseL3()
	ipv4 := currentPacket.GetIPv4()
	//fmt.Printf("NffN3Handler1 %v", ipv4)

	if ipv4 != nil {
		currentPacket.ParseL4ForIPv4()
		if currentPacket.GetUDPForIPv4() != nil {
			/*		fmt.Printf("NffN3Handler2 %v", currentPacket.GetUDPForIPv4())
					fmt.Printf("NffN3Handler3 %v\n", currentPacket.GetUDPNoCheck().DstPort)
					fmt.Println("NffN3Handler3:", packet.SwapUDPPortGTPU)
			*/
			if currentPacket.GetUDPNoCheck().DstPort == packet.SwapUDPPortGTPU {
				return true
			}
		}
	}
	//fmt.Println("----------------------------------------------------------")
	//NffStartDPEHandle()
	return false
}

func NffStartDPEHandleVectorV1(currentPackets []*packet.Packet, mask *[vecSize]bool, notDrop *[vecSize]bool, cxt flow.UserContext) {
	rlogger.FuncEntry(moduleTag, nil)
	n := uint(0)
	for i := uint(0); i < vecSize; i++ {
		if (*mask)[i] == true {
			isDrop := NffStartDPEHandleV1(currentPackets[i], cxt)
			if !isDrop {
				notDrop[i] = true
				n++
			}
		}
	}
}

func NffStartDPEHandleV1(currentPacket *packet.Packet, cxt flow.UserContext) bool {
	rlogger.FuncEntry(moduleTag, nil)
	msgCxt := MsgCxt{}
	msgCxt.Msgbuf, _ = currentPacket.GetPacketPayload()
	msgCxt.Msgcxt = &pdrcontext.DataFlowContext{}

	N3Cxt := cxt.(HandlerContext)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "NffN6Handler hc %p", &N3Cxt)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "NffN6Handler hc %+v", N3Cxt)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "NffN6Handler hc %+v, %+v", N3Cxt.Ipport, N3Cxt.Msgcxt)
	//N3Cxt.Msgcxt = msgCxt.Msgcxt // 没有上行消息都申请一个消息上下文

	// gtp解码，返回gtp的长度和填充gtp信息到上下文中
	// 失败返回 false，做drop处理
	gtpbody, err := RecvN3MsgHandle(msgCxt)
	if err != nil {
		//fmt.Println(err)
		rlogger.Trace(moduleTag, rlogger.ERROR, msgCxt.Msgcxt, "Failed to N3 message decode:%s", err)
		// N3DecodeMsg failed,discard message
		// todo count
		metrics.UpfmoduleSet.UpLinkTotalInvalidPackets.Inc(1)
		return false
	}

	//应用过滤规则
	teid := uint32(msgCxt.Msgcxt.UpfTEID) // (Msgcxt为地址类型）
	rlogger.Trace(moduleTag, rlogger.INFO, msgCxt.Msgcxt, "Received TEID:%v!", teid)
	// 1.1 upgraded version 201812
	fiveTuple, err := utils.IpFiveTuple(gtpbody)
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, msgCxt.Msgcxt, "Failed to N3 message receive IpFiveTuple :%s, upf teid %d", err, teid)
		// todo count
		metrics.UpfmoduleSet.UpLinkTotalInvalidPackets.Inc(1)
		return false
	}
	// mac padding delete
	n := int(fiveTuple.Length)
	if n > len(gtpbody) {
		rlogger.Trace(moduleTag, rlogger.ERROR, msgCxt.Msgcxt, "Failed to ip length:%d,body length:%d", n, len(gtpbody))
		return false
	}
	gtpbody = gtpbody[:n]
	// 上行方向
	fiveTuple.Direction = nasie.UplinkOnly
	rule, err := pdr.LookupULPDRs(teid, fiveTuple)
	if err != nil || rule == nil {
		// 无效TEID的消息丢弃
		rlogger.Trace(moduleTag, rlogger.ERROR, msgCxt.Msgcxt, "Failed to N3 message PDR matching :%s", err)
		// todo count
		metrics.UpfmoduleSet.UpLinkTotalInvalidPackets.Inc(1)
		return false
	}
	rlogger.Trace(moduleTag, rlogger.DEBUG, msgCxt.Msgcxt, "N3 message TEID matching succeeded,"+
		"Received upf teid %d,Ue msg addr(src ip %s,dst ip %s),rule %s",
		teid, fiveTuple.SrcIp, fiveTuple.DstIp, rule)

	// session counter
	var n4Cxt *n4context.N4SessionContext
	var errN4 error
	if metrics.SessionCounterStart == true {
		n4Cxt, errN4 = n4context.GetN4Context(n4context.N4SessionIDKey(rule.SEID))
		if errN4 != nil {
			rlogger.Trace(moduleTag, rlogger.ERROR, n4Cxt, "Failed to get N4 Context:%s", errN4)
		}
	}

	// 应用QFI规则
	qfiIgnore := os.Getenv("QFI_IGNORE")
	if qfiIgnore == "" {
		if false {
			if msgCxt.Msgcxt.UP.QFI != rule.PDI.QFI {
				// 收到的QFI与本地保存的QFI不相等,消息丢弃
				rlogger.Trace(moduleTag, rlogger.ERROR, msgCxt.Msgcxt, "Failed to N3 message QFI matching")
				// todo count
				metrics.UpfmoduleSet.UpLinkTotalInvalidPackets.Inc(1)
				// session counter
				if metrics.SessionCounterStart == true {
					metric.MetricItemAdd(metrics.UpLinkSessionPacketsDiscarded, n4Cxt.MetricItems)
				}
				return false

			}
		}
	}
	// 收到的合法消息统计
	if metrics.SessionCounterStart == true && n4Cxt != nil {
		metric.MetricItemAdd(metrics.UpLinkSessionReceivedPacket, n4Cxt.MetricItems)
		metric.MetricItemAddX(metrics.UpLinkSessionBitsReceived, n4Cxt.MetricItems, int64(len(gtpbody)))
	}

	// 应用转发规则
	if rule.FarI.Far != nil {
		rlogger.Trace(moduleTag, rlogger.DEBUG, msgCxt.Msgcxt, "Action:%v", rule.FarI.Far.Action)
		switch rule.FarI.Far.Action {
		case pdr.FORW:
			if rule.FarI.Far.DstInterface == pdr.DestinationInterface_Core &&
				rule.OuterHeaderRemoval == uint8(pfcp.IEOuterHRemoval_GTPU_UDP_IPv4) {
				//if rule.FarI.Far.OuterHeaderCreation != nil {
				//	copy(server.DnGwIp[:], rule.FarI.Far.OuterHeaderCreation.IPv4Addr.To4())
				//	rlogger.Trace(moduleTag, rlogger.INFO, nil,  "Matching to forwarding rules,dst ip:%s", net.IP(server.DnGwIp[:]))
				//}
				// DNN support
				if rule.FarI.Far.NetworkInstance != "" {
					dnn := rule.FarI.Far.NetworkInstance
					msgCxt.Msgcxt.DnIp = pdr.GetDnnGWIp(dnn)
					rlogger.Trace(moduleTag, rlogger.DEBUG, msgCxt.Msgcxt, "Matching to forwarding rules,dst ip:%v", msgCxt.Msgcxt.DnIp)
				} else {
					dnn := "cmnet"
					msgCxt.Msgcxt.DnIp = pdr.GetDnnGWIp(dnn)
					rlogger.Trace(moduleTag, rlogger.DEBUG, msgCxt.Msgcxt, "Matching to forwarding rules,default dst ip:%v", msgCxt.Msgcxt.DnIp)
				}
			}
			// 发送到N6
		case pdr.DROP: //dropped
			//todo counter
			metrics.UpfmoduleSet.UpLinkTotalSentPacketsDiscarded.Inc(1)
			// session counter
			if metrics.SessionCounterStart == true && n4Cxt != nil {
				metric.MetricItemAdd(metrics.UpLinkSessionPacketsDiscarded, n4Cxt.MetricItems)
			}
			rlogger.Trace(moduleTag, rlogger.ERROR, msgCxt.Msgcxt, "DROP:applicable forwarding action")
			return false
		case pdr.BUFF:
			//	todo：
			rlogger.Trace(moduleTag, rlogger.ERROR, msgCxt.Msgcxt, "No applicable forwarding action")
			return false
		case pdr.DUPL:
			//	todo：
			rlogger.Trace(moduleTag, rlogger.ERROR, msgCxt.Msgcxt, "No applicable forwarding action")
			return false
		case pdr.NOCPBUFF: // buff and notify,deactivate up UL
			rlogger.Trace(moduleTag, rlogger.ERROR, msgCxt.Msgcxt, "applicable forwarding action:UL deactivate")
			//todo counter
			metrics.UpfmoduleSet.UpLinkTotalSentPacketsDiscarded.Inc(1)
			// session counter
			if metrics.SessionCounterStart == true && n4Cxt != nil {
				metric.MetricItemAdd(metrics.UpLinkSessionPacketsDiscarded, n4Cxt.MetricItems)
			}
			return false
		default:
			// todo 默认转发 到N6
			rlogger.Trace(moduleTag, rlogger.ERROR, msgCxt.Msgcxt, "No applicable forwarding action")
			return false
		}
	} else {
		rlogger.Trace(moduleTag, rlogger.ERROR, msgCxt.Msgcxt, "No applicable forwarding rules")
		//todo counter
		metrics.UpfmoduleSet.UpLinkTotalSentPacketsDiscarded.Inc(1)
		// session counter
		if metrics.SessionCounterStart == true && n4Cxt != nil {
			metric.MetricItemAdd(metrics.UpLinkSessionPacketsDiscarded, n4Cxt.MetricItems)
		}
		return false
	}

	//todo counter
	// session counter
	if metrics.SessionCounterStart == true && n4Cxt != nil {
		metric.MetricItemAdd(metrics.UpLinkSessionSendPacket, n4Cxt.MetricItems)
		metric.MetricItemAddX(metrics.UpLinkSessionBitsSent, n4Cxt.MetricItems, int64(len(gtpbody)))
	}

	// 去包头
	// No checking to extension headers. Assume GTP length = 8
	EtherLen := uint(14)
	IPv4MinLen := uint(20)
	// ip length - udp length
	IPv4MinLen = uint(packet.SwapBytesUint16(currentPacket.GetIPv4NoCheck().TotalLength) - packet.SwapBytesUint16(currentPacket.GetUDPNoCheck().DgramLen))
	rlogger.Trace(moduleTag, rlogger.DEBUG, msgCxt.Msgcxt, "length of ip-TotalLength-DgramLen head:%d:%d:%d",
		IPv4MinLen,
		packet.SwapBytesUint16(currentPacket.GetIPv4NoCheck().TotalLength),
		packet.SwapBytesUint16(currentPacket.GetUDPNoCheck().DgramLen))
	UDPLen := uint(8)
	GTPMinLen := uint(16)
	GTPCurrentLen := GTPMinLen
	// gtp length : udp length - UDP head length - gtpbody length
	GTPCurrentLen = uint(packet.SwapBytesUint16(currentPacket.GetUDPNoCheck().DgramLen)-uint16(len(gtpbody))) - UDPLen
	rlogger.Trace(moduleTag, rlogger.DEBUG, msgCxt.Msgcxt, "length of gtpbody %d", len(gtpbody))
	rlogger.Trace(moduleTag, rlogger.DEBUG, msgCxt.Msgcxt, "length of ip-udp-gtp head:%d:%d:%d", IPv4MinLen, UDPLen, GTPCurrentLen)

	N3Cxt.Msgcxt.Rw.Lock()
	N3Cxt.Msgcxt.DnIp = msgCxt.Msgcxt.DnIp
	N3Cxt.Msgcxt.UpfTEID = msgCxt.Msgcxt.UpfTEID
	N3Cxt.Msgcxt.Rw.Unlock()
	// 去掉 IPv4MinLen+UDPLen+GTPMinLen
	currentPacket.DecapsulateHead(EtherLen, IPv4MinLen+UDPLen+GTPCurrentLen)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "NffN6Handler hc %p", &N3Cxt)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "n3cxt %+v", N3Cxt)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "n3cxt %+v", N3Cxt.Ipport)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "n3cxt %+v", N3Cxt.Msgcxt)

	return NffN6SendHandlerV1(currentPacket, N3Cxt)
}

func RecvN3MsgHandle(msgCxt MsgCxt) ([]byte, error) {
	//N3DecodeMsg := &gtpv1_u.N3MsgParser{Msgbuf:[]byte{}}
	Msgbuf := msgCxt.Msgbuf
	N3DecodeMsg := &gtpv1u.N3MsgParser{}
	N3DecodeMsg.Msgbuf = Msgbuf
	err := N3DecodeMsg.DecodeMsg()
	if err != nil {
		//fmt.Println(err)
		rlogger.Trace(moduleTag, rlogger.WARN, msgCxt.Msgcxt, "Failed to N3 message decode!")
		// N3DecodeMsg failed,discard message
		return nil, err
	}
	// 获取gNB带来的teid与QFI
	if N3DecodeMsg.GtpuMsg.MessageType == gtpv1u.Msg_Type_G_PDU {
		if N3DecodeMsg.GPDUSessionUL != nil {
			msgCxt.Msgcxt.UpfTEID = types3gpp.Teid(N3DecodeMsg.GPDUSessionUL.TEID)
			//msgCxt.Msgcxt.UpfIP = N3DecodeMsg.GPDUSessionUL
			msgCxt.Msgcxt.UP = N3DecodeMsg.GPDUSessionUL.PDUSessionContainer.ULPDU
			return N3DecodeMsg.GtpuMsg.Gtpbody, nil
		} else {
			//TODO pass when fail // No extension header
			msgCxt.Msgcxt.UpfTEID = types3gpp.Teid(N3DecodeMsg.GtpuMsg.TEID)
			return N3DecodeMsg.GtpuMsg.Gtpbody, nil
		}

	}
	// 其他node消息
	return N3DecodeMsg.GtpuMsg.Gtpbody, nil
}

func NffN6SendHandlerV1(currentPacket *packet.Packet, cxt flow.UserContext) bool {
	rlogger.FuncEntry(moduleTag, nil)

	N3Cxt := cxt.(HandlerContext)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "NffN6SendHandlerV1 hc %p", &N3Cxt)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "NffN6SendHandlerV1 hc %+v", N3Cxt)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "N3Cxt %+v", N3Cxt.Msgcxt)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "N3Cxt %+v", N3Cxt.Ipport)

	// 查找对端mac地址

	//	修改对应的mac
	currentPacket.Ether.SAddr, currentPacket.Ether.DAddr = currentPacket.Ether.DAddr, currentPacket.Ether.SAddr
	//currentPacket.Ether.SAddr = flow.GetPortMACAddress(0)
	//52:54:00:cf:cf:79
	//MacAddr, _ := flowT.StringToMACAddress("52:54:00:cf:cf:79")
	//currentPacket.Ether.DAddr = MacAddr

	//	从context中获取对端IP，根据对端ip 发起arp，获得mac地址
	//	从IP映射到mac，填充对端mac
	currentPacket.ParseL3()

	// Fill L2
	currentPacket.Ether.SAddr = N3Cxt.Ipport.MacAddress
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "currentPacket.Ether %v", currentPacket.Ether)

	if N3Cxt.Ipport.StaticARP {
		currentPacket.Ether.DAddr = N3Cxt.Ipport.DstMacAddress
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, "N3Cxt.Ipport.StaticARP %v", N3Cxt.Ipport.StaticARP)
	} else {
		// Set destination MAC address to Dnn GW port,
		// so it effectively routes a packet to Dnn GW IP address
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "n3 arp table %s", N3Cxt.Ipport.NeighCache.PrintArpTable())
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "arp lookup %v", N3Cxt.Msgcxt.DnIp)
		N3Cxt.Msgcxt.Rw.RLock()
		if len(N3Cxt.Msgcxt.DnIp) != net.IPv4len {
			rlogger.Trace(moduleTag, rlogger.ERROR, nil, "N3Cxt.Msgcxt.DnIp %v", N3Cxt.Msgcxt.DnIp)
			N3Cxt.Msgcxt.Rw.RUnlock()
			return false
		}
		dstIp := flowT.SliceToIPv4(N3Cxt.Msgcxt.DnIp)
		srcIp := flowT.SliceToIPv4(net.ParseIP(configure.UpfConf.N6.Ipv4).To4())

		N3Cxt.Msgcxt.Rw.RUnlock()
		targetMAC, found := N3Cxt.Ipport.NeighCache.LookupMACForIPv4(dstIp)
		if !found {
			// fmt.Println("Not found MAC address for IP", targetIP.String())
			N3Cxt.Ipport.NeighCache.SendARPRequestForIPv4(dstIp, srcIp, 0)
			rlogger.Trace(moduleTag, rlogger.ERROR, nil, "found dstIp %v,srcIp %v", dstIp, srcIp)
			rlogger.Trace(moduleTag, rlogger.ERROR, nil, "found src MAC %v", N3Cxt.Ipport.MacAddress)
			// 将报文先保存起来
			if (packet.GArpBuffers.Len() > 200){
				packet.GArpMutex.Lock()
				bufferHead := packet.GArpBuffers.Front()
				packet.GArpBuffers.Remove(bufferHead)
				packet.GArpMutex.Unlock()
				oldPacket := bufferHead.Value.(* packet.Packet)
				oldPacket.FreePacket()
			}
			srcbuf := currentPacket.GetRawPacketBytes()
			// 从arp buf pool里申请一个mbuf内存缓存起来
			arpPacket, _ := packet.NewArpBufPacket()
			arpbuf := (*[2 << 10]byte)(arpPacket.StartAtOffset(0))
			copy(arpbuf[0:], srcbuf)
			arpPacket.ParseL3()
			arpPacket.PacketSetDataLen(uint(len(srcbuf)))

			packet.GArpMutex.Lock()
			packet.GArpBuffers.PushBack(arpPacket)
			packet.GArpMutex.Unlock()
			rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "[1] save packet.Ether %v", currentPacket.Ether)
			return false
		}
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "dnnIp %v ,dstIp %v ,srcIp %v:", dstIp, currentPacket.GetIPv4NoCheck().DstAddr, currentPacket.GetIPv4NoCheck().SrcAddr)
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "src MAC %v", N3Cxt.Ipport.MacAddress)
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "targetMAC %v", targetMAC)
		currentPacket.Ether.DAddr = targetMAC
	}
	return true

}
