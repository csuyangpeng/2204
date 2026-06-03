package nasie

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"net"
)

func (p PacketFilterContent) String() string {
	return fmt.Sprintf("id(%d)", p.PacketFilterContentID)
}

type PacketFilterContent struct {
	PacketFilterContentID []PacketFilterContentIdentifier
	//Match-all type, no value

	//Protocol identifier/Next header type     **landslide support**
	V4ProtocolIdOrV6NextHeader byte

	//IPv4 local address type
	IPV4LocalAddr net.IPNet

	//IPv4 remote address type                 **landslide support**
	IPV4RemoteAddr net.IPNet

	//Single local port type
	SinglePort uint16
	//Local port range type                    **landslide support**
	LocalPortRangeLowLimit  uint16
	LocalPortRangeHighLimit uint16

	//Single remote port type
	SingleRemotePort uint16
	//Remote port range type                   **landslide support**
	RemotePortRangeLowLimit  uint16
	RemotePortRangeHighLimit uint16

	//todo
	//IPv6 remote address/prefix length type
	IPV6RemoteAddr          net.IP
	IPV6RemoteAddrPrefixLen byte
	//IPv6 local address/prefix length type
	IPV6LocalAddr          net.IP
	IPV6LocalAddrPrefixLen byte
	//Security parameter index type            **landslide support**
	SecurityParamIndex uint32
	//Type of service/Traffic class type       **landslide support**
	TypeOfServiceTrafficClass     byte
	TypeOfServiceTrafficClassMask byte
	//Destination MAC address type
	DestinationMACAddr []byte
	//Source MAC address type
	SourceMACAddr []byte
	//Flow label type                     **landslide support**
	FlowLableType []byte
	//802.1Q C-TAG VID type todo
	//802.1Q S-TAG VID type todo
	//802.1Q C-TAG PCP/DEI type todo
	//802.1Q S-TAG PCP/DEI type todo
	//Ethertype type todo
	IeFlags bitset.BitSet
}

const (
	Ieid_MatchAlltype uint = iota
	Ieid_IPv4RemoteAddressType
	Ieid_IPv4LocalAddressType
	Ieid_IPv6RemoteAddressPrefixLengthType
	Ieid_IPv6LocalAddressPrefixLengthType
	Ieid_ProtocolIdentifierNextHeaderType
	Ieid_SingleLocalPortType
	Ieid_LocalPortRangeType
	Ieid_SingleRemotePortType
	Ieid_RemotePortRangeType
	Ieid_SecurityParameterIndexType
	Ieid_TypeOfServiceTrafficClassType
	Ieid_FlowLabelType
	Ieid_DestinationMACAddressType
	Ieid_SourceMACAddressType
	Ieid_Q8021CTAGVIDType
	Ieid_Q8021STAGVIDType
	Ieid_Q8021CTAGPCPDEIType
	Ieid_Q8021STAGPCPDEIType
	Ieid_EthertypeType
)

// refer to 24.501 9.11.4.13
//Packet filter component type identifier
//Bits
//8 7 6 5 4 3 2 1
//0 0 0 0 0 0 0 1	Match-all type
//0 0 0 1 0 0 0 0	IPv4 remote address type
//0 0 0 1 0 0 0 1	IPv4 local address type
//0 0 1 0 0 0 0 1	IPv6 remote address/prefix length type
//0 0 1 0 0 0 1 1	IPv6 local address/prefix length type
//0 0 1 1 0 0 0 0	Protocol identifier/Next header type
//0 1 0 0 0 0 0 0	Single local port type
//0 1 0 0 0 0 0 1	Local port range type
//0 1 0 1 0 0 0 0	Single remote port type
//0 1 0 1 0 0 0 1	Remote port range type
//0 1 1 0 0 0 0 0	Security parameter index type
//0 1 1 1 0 0 0 0	Type of service/Traffic class type
//1 0 0 0 0 0 0 0	Flow label type
//1 0 0 0 0 0 0 1	Destination MAC address type
//1 0 0 0 0 0 1 0	Source MAC address type
//1 0 0 0 0 0 1 1	802.1Q C-TAG VID type
//1 0 0 0 0 1 0 0	802.1Q S-TAG VID type
//1 0 0 0 0 1 0 1	802.1Q C-TAG PCP/DEI type
//1 0 0 0 0 1 1 0	802.1Q S-TAG PCP/DEI type
//1 0 0 0 0 1 1 1	Ethertype type
//All other values are reserved.

type PacketFilterContentIdentifier byte

const (
	MatchAlltype                      PacketFilterContentIdentifier = 1
	IPv4RemoteAddressType             PacketFilterContentIdentifier = 0x10 //
	IPv4LocalAddressType              PacketFilterContentIdentifier = 0x11
	IPv6RemoteAddressPrefixLengthType PacketFilterContentIdentifier = 0x21
	IPv6LocalAddressPrefixLengthType  PacketFilterContentIdentifier = 0x23
	ProtocolIdentifierNextHeaderType  PacketFilterContentIdentifier = 0x30 //
	SingleLocalPortType               PacketFilterContentIdentifier = 0x40
	LocalPortRangeType                PacketFilterContentIdentifier = 0x41 //
	SingleRemotePortType              PacketFilterContentIdentifier = 0x50
	RemotePortRangeType               PacketFilterContentIdentifier = 0x51 //
	SecurityParameterIndexType        PacketFilterContentIdentifier = 0x60 //
	TypeOfServiceTrafficClassType     PacketFilterContentIdentifier = 0x70 //
	FlowLabelType                     PacketFilterContentIdentifier = 0x80 //
	DestinationMACAddressType         PacketFilterContentIdentifier = 0x81
	SourceMACAddressType              PacketFilterContentIdentifier = 0x82
	Q8021CTAGVIDType                  PacketFilterContentIdentifier = 0x83
	Q8021STAGVIDType                  PacketFilterContentIdentifier = 0x84
	Q8021CTAGPCPDEIType               PacketFilterContentIdentifier = 0x85
	Q8021STAGPCPDEIType               PacketFilterContentIdentifier = 0x86
	EthertypeType                     PacketFilterContentIdentifier = 0x87
)

// 在一个Packet Filter中，二者只能选其一的有：
// IPv4RemoteAddressType 和 IPv6RemoteAddressPrefixLengthType
// IPv4LocalAddressType  和 IPv6LocalAddressPrefixLengthType
// SingleLocalPortType   和 LocalPortRangeType
// SingleRemotePortType  和 RemotePortRangeType
func (p *PacketFilterContent) Encode() []byte {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte
	for i := 0; i < len(p.PacketFilterContentID); i++ {
		//title
		encBuf = append(encBuf, byte(p.PacketFilterContentID[i]))
		//value
		switch p.PacketFilterContentID[i] {
		case IPv4RemoteAddressType:
			encBuf = append(encBuf, p.IPV4RemoteAddr.IP[:]...)
			encBuf = append(encBuf, p.IPV4RemoteAddr.Mask[:]...)
		case IPv4LocalAddressType:
			encBuf = append(encBuf, p.IPV4LocalAddr.IP[:]...)
			encBuf = append(encBuf, p.IPV4LocalAddr.Mask[:]...)
		case IPv6RemoteAddressPrefixLengthType:
			encBuf = append(encBuf, p.IPV6RemoteAddr[:]...)
			encBuf = append(encBuf, p.IPV6RemoteAddrPrefixLen)
		case IPv6LocalAddressPrefixLengthType:
			encBuf = append(encBuf, p.IPV6LocalAddr[:]...)
			encBuf = append(encBuf, p.IPV6LocalAddrPrefixLen)
		case ProtocolIdentifierNextHeaderType:
			encBuf = append(encBuf, p.V4ProtocolIdOrV6NextHeader)
		case SingleLocalPortType, SingleRemotePortType:
			bytes := make([]byte, 2)
			binary.BigEndian.PutUint16(bytes, p.SinglePort)
			encBuf = append(encBuf, bytes[:]...)
		case LocalPortRangeType:
			lowbytes := make([]byte, 2)
			binary.BigEndian.PutUint16(lowbytes, p.LocalPortRangeHighLimit)
			encBuf = append(encBuf, lowbytes[:]...)
			highbytes := make([]byte, 2)
			binary.BigEndian.PutUint16(highbytes, p.LocalPortRangeHighLimit)
			encBuf = append(encBuf, highbytes[:]...)
		case RemotePortRangeType:
			lowbytes := make([]byte, 2)
			binary.BigEndian.PutUint16(lowbytes, p.RemotePortRangeHighLimit)
			encBuf = append(encBuf, lowbytes[:]...)
			highbytes := make([]byte, 2)
			binary.BigEndian.PutUint16(highbytes, p.RemotePortRangeHighLimit)
			encBuf = append(encBuf, highbytes[:]...)
		case SecurityParameterIndexType:
			securitybytes := make([]byte, 4)
			binary.BigEndian.PutUint32(securitybytes, p.SecurityParamIndex)
			encBuf = append(encBuf, securitybytes[:]...)
		case TypeOfServiceTrafficClassType:
			encBuf = append(encBuf, p.TypeOfServiceTrafficClass)
			encBuf = append(encBuf, p.TypeOfServiceTrafficClassMask)
		case DestinationMACAddressType:
			encBuf = append(encBuf, p.DestinationMACAddr[:]...)
		case SourceMACAddressType:
			encBuf = append(encBuf, p.SourceMACAddr[:]...)
		case FlowLabelType:
			encBuf = append(encBuf, p.FlowLableType[:]...)
		case Q8021CTAGVIDType:
		case Q8021STAGVIDType:
		case Q8021CTAGPCPDEIType:
		case Q8021STAGPCPDEIType:
		case EthertypeType:
		default:
		}
	}

	return encBuf
}

func (p *PacketFilterContent) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	//octet 9
	packetFilterContentsLen, err := msgBuf.ReadByte()
	if err != nil {
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "no more bytes")
		return fmt.Errorf("fail to read byte")
	}
	rlogger.Trace(types.ModCmn, rlogger.INFO, nil, "PacketFilterContent length", packetFilterContentsLen)
	j := 0
	p.PacketFilterContentID = make([]PacketFilterContentIdentifier, packetFilterContentsLen)
	for i := packetFilterContentsLen; i > 0; j++ {
		//id
		contentID, err := msgBuf.ReadByte()
		if err != nil {
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "no more bytes")
			return fmt.Errorf("fail to read byte")
		}
		p.PacketFilterContentID[j] = PacketFilterContentIdentifier(contentID)
		i -= 1
		switch p.PacketFilterContentID[j] {
		case IPv4RemoteAddressType:
			//ip
			ip := make([]byte, 4)
			binary.Read(msgBuf, binary.BigEndian, ip)
			p.IPV4RemoteAddr.IP = ip
			//mask
			mask := make([]byte, 4)
			binary.Read(msgBuf, binary.BigEndian, mask)
			p.IPV4RemoteAddr.Mask = mask
			i -= 8
		case IPv4LocalAddressType:
			//ip
			ip := make([]byte, 4)
			binary.Read(msgBuf, binary.BigEndian, ip)
			p.IPV4LocalAddr.IP = ip
			//mask
			mask := make([]byte, 4)
			binary.Read(msgBuf, binary.BigEndian, mask)
			p.IPV4LocalAddr.Mask = mask
			i -= 8
		case IPv6RemoteAddressPrefixLengthType:
			ip := make([]byte, 16)
			binary.Read(msgBuf, binary.BigEndian, ip)
			p.IPV6RemoteAddr = ip
			p.IPV6RemoteAddrPrefixLen, _ = msgBuf.ReadByte()
			i -= 16
		case IPv6LocalAddressPrefixLengthType:
			ip := make([]byte, 16)
			binary.Read(msgBuf, binary.BigEndian, ip)
			p.IPV6LocalAddr = ip
			p.IPV6LocalAddrPrefixLen, _ = msgBuf.ReadByte()
			i -= 16
		case ProtocolIdentifierNextHeaderType:
			p.V4ProtocolIdOrV6NextHeader, _ = msgBuf.ReadByte()
			i -= 1
		case SingleLocalPortType, SingleRemotePortType:
			port := make([]byte, 2)
			binary.Read(msgBuf, binary.BigEndian, port)
			p.SinglePort = binary.BigEndian.Uint16(port)
			i -= 2
		case LocalPortRangeType:
			lowport := make([]byte, 2)
			binary.Read(msgBuf, binary.BigEndian, lowport)
			p.LocalPortRangeLowLimit = binary.BigEndian.Uint16(lowport)
			highport := make([]byte, 2)
			binary.Read(msgBuf, binary.BigEndian, highport)
			p.LocalPortRangeHighLimit = binary.BigEndian.Uint16(highport)
			i -= 4
		case RemotePortRangeType:
			lowport := make([]byte, 2)
			binary.Read(msgBuf, binary.BigEndian, lowport)
			p.RemotePortRangeLowLimit = binary.BigEndian.Uint16(lowport)
			highport := make([]byte, 2)
			binary.Read(msgBuf, binary.BigEndian, highport)
			p.RemotePortRangeHighLimit = binary.BigEndian.Uint16(highport)
			i -= 4
		case SecurityParameterIndexType:
			index := make([]byte, 4)
			binary.Read(msgBuf, binary.BigEndian, index)
			p.SecurityParamIndex = binary.BigEndian.Uint32(index)
			i -= 4
		case TypeOfServiceTrafficClassType:
			p.TypeOfServiceTrafficClass, _ = msgBuf.ReadByte()
			p.TypeOfServiceTrafficClassMask, _ = msgBuf.ReadByte()
			i -= 2
		case DestinationMACAddressType:
			desMAC := make([]byte, 16)
			binary.Read(msgBuf, binary.BigEndian, desMAC)
			p.DestinationMACAddr = desMAC
			i -= 16
		case SourceMACAddressType:
			sourMAC := make([]byte, 16)
			binary.Read(msgBuf, binary.BigEndian, sourMAC)
			p.SourceMACAddr = sourMAC
			i -= 16
		case FlowLabelType:
			bytes := make([]byte, 3)
			binary.Read(msgBuf, binary.BigEndian, bytes)
			p.FlowLableType = bytes
			i -= 3
		case Q8021CTAGVIDType:
		case Q8021STAGVIDType:
		case Q8021CTAGPCPDEIType:
		case Q8021STAGPCPDEIType:
		case EthertypeType:
		default:
		}
	}
	p.PacketFilterContentID = p.PacketFilterContentID[:j]
	//fmt.Println("pft.PacketFilterContentID",p.PacketFilterContentID)
	//fmt.Println("pft.V4ProtocolIdOrV6NextHeader",p.V4ProtocolIdOrV6NextHeader)
	return nil
}
