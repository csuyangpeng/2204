// +build linux,!386

package sctp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"syscall"
	"unsafe"
)

func setsockopt(fd int, optname, optval, optlen uintptr) (uintptr, uintptr, error) {
	// FIXME: syscall.SYS_SETSOCKOPT is undefined on 386
	r0, r1, errno := syscall.Syscall6(syscall.SYS_SETSOCKOPT,
		uintptr(fd),
		SOL_SCTP,
		optname,
		optval,
		optlen,
		0)
	if errno != 0 {
		return r0, r1, errno
	}
	return r0, r1, nil
}

func getsockopt(fd int, optname, optval, optlen uintptr) (uintptr, uintptr, error) {
	// FIXME: syscall.SYS_GETSOCKOPT is undefined on 386
	r0, r1, errno := syscall.Syscall6(syscall.SYS_GETSOCKOPT,
		uintptr(fd),
		SOL_SCTP,
		optname,
		optval,
		optlen,
		0)
	if errno != 0 {
		return r0, r1, errno
	}
	return r0, r1, nil
}

func parseSndRcvInfo(b []byte) (*SndRcvInfo, error) {
	msgs, err := syscall.ParseSocketControlMessage(b)
	if err != nil {
		return nil, err
	}
	for _, m := range msgs {
		if m.Header.Level == syscall.IPPROTO_SCTP {
			switch m.Header.Type {
			case SCTP_CMSG_SNDRCV:
				return (*SndRcvInfo)(unsafe.Pointer(&m.Data[0])), nil
			}
		}
	}
	return nil, nil
}

func ParseSctpNotifyHeader(b []byte) (*SctpNotifyHeader, error) {
	//get sctp notify type
	header := &SctpNotifyHeader{}
	err := header.Read(b)
	if err != nil {
		return nil, err
	}
	return header, nil
}

func ParseSctpNotify_PeerAddrChange(b []byte) (*SctpPeerAddrChange, error) {
	//get sctp peer addr change event
	data := &SctpPeerAddrChange{}
	err := data.Read(b)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ParseSctpNotify_AssocChange(b []byte) (*SctpAssocChange, error) {
	//get sctp association change event
	data := &SctpAssocChange{}
	err := data.Read(b)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ListenSCTP - start listener on specified address/port
func ListenSCTP(net string, laddr *SCTPAddr) (*SCTPListener, error) {
	return ListenSCTPExt(net, laddr, InitMsg{NumOstreams: SCTP_MAX_STREAM})
}

// ListenSCTPExt - start listener on specified address/port with given SCTP options
func ListenSCTPExt(network string, laddr *SCTPAddr, options InitMsg) (*SCTPListener, error) {
	af, ipv6only := favoriteAddrFamily(network, laddr, nil, "listen")
	sock, err := syscall.Socket(
		af,
		syscall.SOCK_STREAM,
		syscall.IPPROTO_SCTP,
	)
	if err != nil {
		return nil, err
	}

	// close socket on error
	defer func() {
		if err != nil {
			syscall.Close(sock)
		}
	}()
	if err = setDefaultSockopts(sock, af, ipv6only); err != nil {
		return nil, err
	}
	err = setInitOpts(sock, options)
	if err != nil {
		return nil, err
	}

	if laddr != nil {
		// If IP address and/or port was not provided so far, let's use the unspecified IPv4 or IPv6 address
		if len(laddr.IPAddrs) == 0 {
			if af == syscall.AF_INET {
				laddr.IPAddrs = append(laddr.IPAddrs, net.IPAddr{IP: net.IPv4zero})
			} else if af == syscall.AF_INET6 {
				laddr.IPAddrs = append(laddr.IPAddrs, net.IPAddr{IP: net.IPv6zero})
			}
		}
		err := SCTPBind(sock, laddr, SCTP_BINDX_ADD_ADDR)
		if err != nil {
			return nil, err
		}
	}

	err = syscall.Listen(sock, syscall.SOMAXCONN)
	if err != nil {
		return nil, err
	}
	return &SCTPListener{
		fd: sock,
	}, nil
}

// DialSCTP - bind socket to laddr (if given) and connect to raddr
func DialSCTP(net string, laddr, raddr *SCTPAddr) (*SCTPConn, error) {
	return DialSCTPExt(net, laddr, raddr, InitMsg{NumOstreams: SCTP_MAX_STREAM})
}

// DialSCTPExt - same as DialSCTP but with given SCTP options
func DialSCTPExt(network string, laddr, raddr *SCTPAddr, options InitMsg) (*SCTPConn, error) {
	af, ipv6only := favoriteAddrFamily(network, laddr, raddr, "dial")
	sock, err := syscall.Socket(
		af,
		syscall.SOCK_STREAM,
		syscall.IPPROTO_SCTP,
	)
	if err != nil {
		return nil, err
	}

	// close socket on error
	defer func() {
		if err != nil {
			syscall.Close(sock)
		}
	}()
	if err = setDefaultSockopts(sock, af, ipv6only); err != nil {
		return nil, err
	}
	err = setInitOpts(sock, options)
	if err != nil {
		return nil, err
	}
	if laddr != nil {
		// If IP address and/or port was not provided so far, let's use the unspecified IPv4 or IPv6 address
		if len(laddr.IPAddrs) == 0 {
			if af == syscall.AF_INET {
				laddr.IPAddrs = append(laddr.IPAddrs, net.IPAddr{IP: net.IPv4zero})
			} else if af == syscall.AF_INET6 {
				laddr.IPAddrs = append(laddr.IPAddrs, net.IPAddr{IP: net.IPv6zero})
			}
		}
		err := SCTPBind(sock, laddr, SCTP_BINDX_ADD_ADDR)
		if err != nil {
			return nil, err
		}
	}
	_, err = SCTPConnect(sock, raddr)
	if err != nil {
		return nil, err
	}
	return NewSCTPConn(sock), nil
}

//struct {
//__u16 sn_type;             /* Notification type. */
//__u16 sn_flags;
//__u32 sn_length;
//} sn_header;
type SctpNotifyHeader struct {
	SnType   uint16
	SnFlags  uint16
	SnLength uint32
}

func (p *SctpNotifyHeader) Read(b []byte) error {
	// sn type
	r := bytes.NewReader(b)
	dataU16 := make([]byte, 2)
	err := binary.Read(r, binary.LittleEndian, &dataU16)
	if err != nil {
		return fmt.Errorf("failed to read sctp notify type, error(%s)", err)
	}
	p.SnType = binary.LittleEndian.Uint16(dataU16)

	// sn flags
	err = binary.Read(r, binary.LittleEndian, &dataU16)
	if err != nil {
		return fmt.Errorf("failed to read sctp notify flags, error(%s)", err)
	}
	p.SnFlags = binary.LittleEndian.Uint16(dataU16)

	// sn length
	dataU32 := make([]byte, 4)
	err = binary.Read(r, binary.LittleEndian, &dataU32)
	if err != nil {
		return fmt.Errorf("failed to read sctp notify length, error(%s)", err)
	}
	p.SnLength = binary.LittleEndian.Uint32(dataU32)

	return nil
}

// struct sctp_paddr_change {
//	__u16 spc_type;
//	__u16 spc_flags;
//	__u32 spc_length;
//	struct sockaddr_storage spc_aaddr;
//	int spc_state;
//	int spc_error;
//	sctp_assoc_t spc_assoc_id;
//} __attribute__((packed, aligned(4)));
//type SctpPaddrChangeNotify struct{
//	spc_aaddr [128]byte
//	spc_state
//}
