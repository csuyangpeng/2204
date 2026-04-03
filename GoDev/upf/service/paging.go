package service

import (
	"container/list"
	"errors"
	"fmt"
	"github.com/intel-go/nff-go/packet"
	flowT "github.com/intel-go/nff-go/types"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/upf/adapter"
	"lite5gc/upf/context/ipport"
	"lite5gc/upf/context/pdrcontext"
	"lite5gc/upf/cp/n4layer"
	"lite5gc/upf/cp/pdr"
	"lite5gc/upf/utils"
)

var PagingServer *rawServer

func init() {
	//var upfCtxt *types.AppContext
	//StartBufferServer(upfCtxt)

}

// StartBufferServer
func StartBufferServer(upfCtxt *types.AppContext) error {
	rlogger.FuncEntry(moduleTag, nil)
	// create a receive message server
	PagingServer = NewRawServer()
	if PagingServer == nil {
		panic("Failed to apply for memory")
		return errors.New("Failed to apply for memory")
	}

	// paging cache to n4Ctxt
	go StartDPERawPagingHandle(upfCtxt, PagingServer)

	// paging support,paging server
	go StartBufferListHandle(upfCtxt, &n4layer.SendingList)

	// paging support,send dl data
	go StartPagingBufferHandle(upfCtxt, &n4layer.UpfN4Layer, PagingServer)

	return nil
}

// StartDPERawPagingHandle goroutine开启Rawsocket收包,paging
func StartDPERawPagingHandle(upfCtxt *types.AppContext, server *rawServer) {
	fmt.Printf("start raw DPE Paging\n")
	rlogger.FuncEntry(moduleTag, nil)
	rlogger.Trace(moduleTag, rlogger.INFO, nil, "raw dpe Paging server routine start")
	upfCtxt.Wg.Add(1)
	defer upfCtxt.Wg.Done()

	for {
		select {
		case <-upfCtxt.Ctx.Done():
			rlogger.Trace(moduleTag, rlogger.ERROR, nil, "raw dpe Paging server routine exit")
			return

		case BuffMsg := <-server.MsgBuffChan:
			rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "Receive N6 Message buff :%+v \n", *BuffMsg.Msgcxt)
			// 缓存消息
			n4layer.N4SessionDeactivationBuffering(BuffMsg.Msgbuf, BuffMsg.Msgcxt)
		}
	}
}

func StartBufferListHandle(upfCtxt *types.AppContext, sl *n4layer.SendList) {
	rlogger.FuncEntry(moduleTag, nil)
	rlogger.Trace(moduleTag, rlogger.INFO, nil, "Paging Buffer List routine start")
	upfCtxt.Wg.Add(1)
	defer upfCtxt.Wg.Done()

	for {
		if sl.SendList.Len() > 0 {
			rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "PagingSend StartBufferListHandle 1 sl.SendList.Len:%d", sl.SendList.Len())
			sl.Rw.RLock()
			msgbufList := sl.SendList.Front() //n4cxt.Buffer
			if msgbufList == nil {
				continue
			}
			sl.SendList.Remove(msgbufList)
			sl.Rw.RUnlock()
			bufList := msgbufList.Value.(*list.List)
			rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "PagingSend StartBufferListHandle 2 sl.SendList.Len:%d", sl.SendList.Len())
			// 阻塞发送
			n4layer.SendBufferMsgObstructive(bufList)
		} else {
			// 阻塞等待
			select {
			case <-upfCtxt.Ctx.Done():
				rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Paging Buffer List routine exit")
				return

			case <-sl.State:
				rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "PagingSend StartBufferListHandle 0 SendingList receive task to work")
				continue

			}
		}
	}

}

// StartPagingBufferHandle goroutine开启发送缓存数据到UE,paging
func StartPagingBufferHandle(upfCtxt *types.AppContext, server *n4layer.N4Layer, RawServer *rawServer) {
	rlogger.Trace(moduleTag, rlogger.INFO, nil, "Paging Buffer server routine start")
	upfCtxt.Wg.Add(1)
	defer upfCtxt.Wg.Done()

	for {
		select {
		case <-upfCtxt.Ctx.Done():
			rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Paging Buffer server routine exit")
			return

		case Msgbuf := <-server.BufferMsg:
			rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "Receive N6 buffer Message :%v \n", Msgbuf)
			// 发送缓存消息到UE
			fiveTuple, err := utils.IpFiveTuple(Msgbuf)
			if err != nil {
				//upfmetric.UpfmoduleSet.DownLinkTotalInvalidPackets.Inc(1)
				rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Receive N6 buffer Message :%s \n", err)
				continue
			}
			// 下行方向
			fiveTuple.Direction = nasie.DownlinkOnly
			//应用过滤规则
			rule, err := pdr.LookupPDRs(fiveTuple)
			if err != nil || rule == nil {
				// 没有匹配规则的包丢弃
				//upfmetric.UpfmoduleSet.DownLinkTotalInvalidPackets.Inc(1)
				rlogger.Trace(moduleTag, rlogger.WARN, nil, "Receive N6 buffer Message :%s \n", err)
				continue
			}

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
					continue
				case pdr.BUFF:
					//	todo：
					continue
				case pdr.DUPL:
					//	todo：
					continue
				case pdr.NOCPBUFF: // buff and notify,deactivate up DL
					rlogger.Trace(moduleTag, rlogger.WARN, &flowCxt, "applicable forwarding action:DL deactivate")

					// paging support
					flowCxt.SEID = rule.SEID
					flowCxt.RuleID = rule.RuleID
					RawServer.ReceiveBuffMsg(Msgbuf, &flowCxt)

					continue
				default:
					// 默认转发 到N3
					flowCxt.DP.QFI = rule.PDI.QFI
					flowCxt.GnbTEID = 0
					flowCxt.GnbIP = RawServer.GNbIp
					flowCxt.OuterHeaderDesc = pfcp.IEOuterHCreation_GTPU_UDP_IPv4
					rlogger.Trace(moduleTag, rlogger.ERROR, &flowCxt, "No applicable forwarding action")
					rlogger.Trace(moduleTag, rlogger.ERROR, &flowCxt, "No applicable forwarding action %d,%+v", rule.SEID, rule)
				}
			} else {
				rlogger.Trace(moduleTag, rlogger.ERROR, &flowCxt, "No applicable forwarding rules")

			}

			// buffer 发送到到N3处理
			recvRet := nffDPERawPagingHandle(Msgbuf, &flowCxt)
			if recvRet != true {
				rlogger.Trace(moduleTag, rlogger.ERROR, &flowCxt, "The buffer is full, discarding the message")
			}

			rlogger.Trace(moduleTag, rlogger.DEBUG, &flowCxt, "Receive N6Msg from DN buffer\n")

		}
	}
}

// nff 处理n6 buff 消息
func nffDPERawPagingHandle(buff []byte, cxt *pdrcontext.DataFlowContext) bool {
	rlogger.FuncEntry(moduleTag, nil)

	N6Cxt := cxt
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "NffN6Handler hc %p", &N6Cxt)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "NffN6Handler hc %+v", N6Cxt)

	conutPacket := 0 //test count
	var EncodeMsg []byte

	// 创建包头指示
	if N6Cxt.OuterHeaderDesc == pfcp.IEOuterHCreation_GTPU_UDP_IPv4 {
		GtpMsg, err := adapter.SendN3MsgHandleExt(buff, N6Cxt)
		rlogger.Trace(moduleTag, rlogger.DEBUG, N6Cxt, "Receive N6 Message len=%d value: %#x\n", len(GtpMsg), GtpMsg)
		if err != nil {
			//fmt.Println(err)
			rlogger.Trace(moduleTag, rlogger.ERROR, N6Cxt, "Failed to N6 message encode!")
			// N3EncodeMsg failed,discard message
			return false
		}
		EncodeMsg = GtpMsg
		// 设置转发目的 IP
		//server.GNbIp = N6Cxt.Msgcxt.GnbIP
		rlogger.Trace(moduleTag, rlogger.DEBUG, N6Cxt, "Receive N6 Message MsgCxt value: %+v\n", *N6Cxt)
	} else {
		EncodeMsg = buff
		rlogger.Trace(moduleTag, rlogger.ERROR, N6Cxt, "Receive N6 Message OuterHeaderDesc fault:%d", N6Cxt.OuterHeaderDesc)
	}

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
	rlogger.Trace(moduleTag, rlogger.DEBUG, N6Cxt, "add gtp head,etherLen-ipv4MinLen-UDPLen-GTPMinLen", etherLen, ipv4MinLen, UDPLen, GTPMinLen)
	//length := currentPacket.GetPacketLen() - flowT.EtherLen
	// 创建 packet
	currentPacket, err := packet.NewPacket()
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, N6Cxt, "failed to NewPacket")
		return false
	}
	ok := packet.InitEmptyIPv4UDPPacket(currentPacket, uint(len(EncodeMsg)))
	if !ok {
		rlogger.Trace(moduleTag, rlogger.ERROR, N6Cxt, "failed to InitEmptyIPv4UDPPacket")
		return false
	}

	currentPacket.ParseL3()
	//非固定头填充
	currentPacket.PacketBytesChange(etherLen+ipv4MinLen+UDPLen, EncodeMsg)
	rlogger.Trace(moduleTag, rlogger.DEBUG, N6Cxt, "change N6 Message len=%d value: %#x\n", len(EncodeMsg), EncodeMsg)

	// 固定头填充
	//gtp.HeaderType = 0x34          // 001 - GTPv1, 1 - not GTP', 0 - reserved, 000 - no optional fields
	//gtp.MessageType = packet.G_PDU // encapsulated user message
	//gtp.MessageLength = packet.SwapBytesUint16(uint16(length))
	//gtp.TEID = packet.SwapBytesUint32(uint32(N6Cxt.Msgcxt.GnbTEID))
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "n6cxt %+v", N6Cxt)

	// Fill new IPv4 header with addresses according to context
	ipv4 := currentPacket.GetIPv4NoCheck() //(*packet.IPv4Hdr)(currentPacket.L3)

	// 源IP，目的IP
	if len(ipport.IpPorts) <= int(ipport.N3Outport) {
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, "len of IpPorts ", len(ipport.IpPorts), "N3Outport ", ipport.N3Outport)
	}
	Ipport := ipport.IpPorts[ipport.N3Outport]
	srcAddr := Ipport.Subnet.IPv4.Addr
	dstAddr := flowT.SliceToIPv4(N6Cxt.GnbIP.To4())
	ipv4.SrcAddr = srcAddr
	ipv4.DstAddr = dstAddr
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "ipv4", ipv4.SrcAddr, ipv4.DstAddr)

	// Fill L2
	currentPacket.Ether.EtherType = flowT.SwapIPV4Number
	currentPacket.Ether.SAddr = Ipport.MacAddress
	if Ipport.StaticARP {
		currentPacket.Ether.DAddr = Ipport.DstMacAddress
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "ipv4 Static currentPacket.Ether.DAddr", currentPacket.Ether.DAddr)
	} else {
		// Find l2 addresses for new destionation IP in ARP cache
		targetIP := ipv4.DstAddr
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "ipv4 targetIP", targetIP)
		targetMAC, found := Ipport.NeighCache.LookupMACForIPv4(targetIP)
		if !found {
			// fmt.Println("Not found MAC address for IP", targetIP.String())
			Ipport.NeighCache.SendARPRequestForIPv4(targetIP, ipv4.SrcAddr, 0)
			rlogger.Trace(moduleTag, rlogger.ERROR, nil, "ipv4 targetIP ", targetIP, "ipv4 SrcAddr", ipv4.SrcAddr)
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

			packet.GArpMutex.Lock()
			packet.GArpBuffers.PushBack(arpPacket)
			packet.GArpMutex.Unlock()
			return false
		}
		currentPacket.Ether.DAddr = targetMAC
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "ipv4 currentPacket.Ether.DAddr", currentPacket.Ether.DAddr)

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

	// Send to network ，发送到N3对端
	currentPacket.SendPacket(Ipport.Index)
	return true
}
