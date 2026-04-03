package sctp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
	"syscall"
	"unsafe"
)

type NotificationHandler func([]byte, SctpConnIf) error

type EventSubscribe struct {
	DataIO          uint8
	Association     uint8
	Address         uint8
	SendFailure     uint8
	PeerError       uint8
	Shutdown        uint8
	PartialDelivery uint8
	AdaptationLayer uint8
	Authentication  uint8
	SenderDry       uint8
}

type InitMsg struct {
	NumOstreams    uint16
	MaxInstreams   uint16
	MaxAttempts    uint16
	MaxInitTimeout uint16
}

func (p *InitMsg) String() string {
	return fmt.Sprintf("NumOstreams(%d),MaxInstreams(%d),MaxAttemps(%d),MaxInitTimeout(%d)",
		p.NumOstreams, p.MaxInstreams, p.MaxAttempts, p.MaxInitTimeout)
}

type SndRcvInfo struct {
	Stream  uint16
	SSN     uint16
	Flags   uint16
	_       uint16
	PPID    uint32
	Context uint32
	TTL     uint32
	TSN     uint32
	CumTSN  uint32
	AssocID int32
}

type SndInfo struct {
	SID     uint16
	Flags   uint16
	PPID    uint32
	Context uint32
	AssocID int32
}

type GetAddrsOld struct {
	AssocID int32
	AddrNum int32
	Addrs   uintptr
}

type NotificationHeader struct {
	Type   uint16
	Flags  uint16
	Length uint32
}

var nativeEndian binary.ByteOrder
var sndRcvInfoSize uintptr

func init() {
	i := uint16(1)
	if *(*byte)(unsafe.Pointer(&i)) == 0 {
		nativeEndian = binary.BigEndian
	} else {
		nativeEndian = binary.LittleEndian
	}
	info := SndRcvInfo{}
	sndRcvInfoSize = unsafe.Sizeof(info)
}

func toBuf(v interface{}) []byte {
	var buf bytes.Buffer
	binary.Write(&buf, nativeEndian, v)
	return buf.Bytes()
}

func htons(h uint16) uint16 {
	if nativeEndian == binary.LittleEndian {
		return (h << 8 & 0xff00) | (h >> 8 & 0xff)
	}
	return h
}

var ntohs = htons

// setInitOpts sets options for an SCTP association initialization
// see https://tools.ietf.org/html/rfc4960#page-25
func setInitOpts(fd int, options InitMsg) error {
	optlen := unsafe.Sizeof(options)
	_, _, err := setsockopt(fd, SCTP_INITMSG, uintptr(unsafe.Pointer(&options)), uintptr(optlen))
	return err
}

func getInitOpts(fd int) (*InitMsg, error) {
	initMsg := &InitMsg{}

	msglen := unsafe.Sizeof(initMsg)

	_, _, err := getsockopt(fd, SCTP_INITMSG, uintptr(unsafe.Pointer(initMsg)), uintptr(unsafe.Pointer(&msglen)))
	if err != nil {
		return nil, err
	}
	return initMsg, nil
}

func setNoDelay(fd int, flag bool) error {
	var options uint32
	if flag == true {
		options = 1
	} else {
		options = 0
	}
	optlen := unsafe.Sizeof(options)
	_, _, err := setsockopt(fd, SCTP_NODELAY, uintptr(unsafe.Pointer(&options)), uintptr(optlen))
	return err
}

func ResolveSCTPAddr(network, addrs string) (*SCTPAddr, error) {
	tcpnet := ""
	switch network {
	case "", "sctp":
		tcpnet = "tcp"
	case "sctp4":
		tcpnet = "tcp4"
	case "sctp6":
		tcpnet = "tcp6"
	default:
		return nil, fmt.Errorf("invalid net: %s", network)
	}
	elems := strings.Split(addrs, "/")
	if len(elems) == 0 {
		return nil, fmt.Errorf("invalid input: %s", addrs)
	}
	ipaddrs := make([]net.IPAddr, 0, len(elems))
	for _, e := range elems[:len(elems)-1] {
		tcpa, err := net.ResolveTCPAddr(tcpnet, e+":")
		if err != nil {
			return nil, err
		}
		ipaddrs = append(ipaddrs, net.IPAddr{IP: tcpa.IP, Zone: tcpa.Zone})
	}
	tcpa, err := net.ResolveTCPAddr(tcpnet, elems[len(elems)-1])
	if err != nil {
		return nil, err
	}
	if tcpa.IP != nil {
		ipaddrs = append(ipaddrs, net.IPAddr{IP: tcpa.IP, Zone: tcpa.Zone})
	} else {
		ipaddrs = nil
	}
	return &SCTPAddr{
		IPAddrs: ipaddrs,
		Port:    tcpa.Port,
	}, nil
}

func SCTPConnect(fd int, addr *SCTPAddr) (int, error) {
	buf := addr.ToRawSockAddrBuf()
	param := GetAddrsOld{
		AddrNum: int32(len(buf)),
		Addrs:   uintptr(uintptr(unsafe.Pointer(&buf[0]))),
	}
	optlen := unsafe.Sizeof(param)
	_, _, err := getsockopt(fd, SCTP_SOCKOPT_CONNECTX3, uintptr(unsafe.Pointer(&param)), uintptr(unsafe.Pointer(&optlen)))
	if err == nil {
		return int(param.AssocID), nil
	} else if err != syscall.ENOPROTOOPT {
		return 0, err
	}
	r0, _, err := setsockopt(fd, SCTP_SOCKOPT_CONNECTX, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	return int(r0), err
}

func SCTPBind(fd int, addr *SCTPAddr, flags int) error {
	var option uintptr
	switch flags {
	case SCTP_BINDX_ADD_ADDR:
		option = SCTP_SOCKOPT_BINDX_ADD
	case SCTP_BINDX_REM_ADDR:
		option = SCTP_SOCKOPT_BINDX_REM
	default:
		return syscall.EINVAL
	}

	buf := addr.ToRawSockAddrBuf()
	_, _, err := setsockopt(fd, option, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	return err
}

func resolveFromRawAddr(ptr unsafe.Pointer, n int) (*SCTPAddr, error) {
	addr := &SCTPAddr{
		IPAddrs: make([]net.IPAddr, n),
	}

	switch family := (*(*syscall.RawSockaddrAny)(ptr)).Addr.Family; family {
	case syscall.AF_INET:
		addr.Port = int(ntohs(uint16((*(*syscall.RawSockaddrInet4)(ptr)).Port)))
		tmp := syscall.RawSockaddrInet4{}
		size := unsafe.Sizeof(tmp)
		for i := 0; i < n; i++ {
			a := *(*syscall.RawSockaddrInet4)(unsafe.Pointer(
				uintptr(ptr) + size*uintptr(i)))
			addr.IPAddrs[i] = net.IPAddr{IP: a.Addr[:]}
		}
	case syscall.AF_INET6:
		addr.Port = int(ntohs(uint16((*(*syscall.RawSockaddrInet4)(ptr)).Port)))
		tmp := syscall.RawSockaddrInet6{}
		size := unsafe.Sizeof(tmp)
		for i := 0; i < n; i++ {
			a := *(*syscall.RawSockaddrInet6)(unsafe.Pointer(
				uintptr(ptr) + size*uintptr(i)))
			var zone string
			ifi, err := net.InterfaceByIndex(int(a.Scope_id))
			if err == nil {
				zone = ifi.Name
			}
			addr.IPAddrs[i] = net.IPAddr{IP: a.Addr[:], Zone: zone}
		}
	default:
		return nil, fmt.Errorf("unknown address family: %d", family)
	}
	return addr, nil
}

func sctpGetAddrs(fd, id, optname int) (*SCTPAddr, error) {

	type getaddrs struct {
		assocId int32
		addrNum uint32
		addrs   [4096]byte
	}
	param := getaddrs{
		assocId: int32(id),
	}
	optlen := unsafe.Sizeof(param)
	_, _, err := getsockopt(fd, uintptr(optname), uintptr(unsafe.Pointer(&param)), uintptr(unsafe.Pointer(&optlen)))
	if err != nil {
		return nil, err
	}
	return resolveFromRawAddr(unsafe.Pointer(&param.addrs), int(param.addrNum))
}
