package sctp

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// kernel include/net/sctp/user.h
//660  * 7.2.1 Association Status (SCTP_STATUS)
//661  *
//662  *   Applications can retrieve current status information about an
//663  *   association, including association state, peer receiver window size,
//664  *   number of unacked data chunks, and number of data chunks pending
//665  *   receipt.  This information is read-only.  The following structure is
//666  *   used to access this information:
//667  */
//668 struct sctp_status {
//669     sctp_assoc_t        sstat_assoc_id;
//670     __s32           sstat_state;
//671     __u32           sstat_rwnd;
//672     __u16           sstat_unackdata;
//673     __u16           sstat_penddata;
//674     __u16           sstat_instrms;
//675     __u16           sstat_outstrms;
//676     __u32           sstat_fragmentation_point;
//677     struct sctp_paddrinfo   sstat_primary;
//678 };
type SctpStatus struct {
	Sstat_assoc_id            uint32
	Sstat_state               int32
	Sstat_rwnd                uint32
	Sstat_unackdata           uint16
	Sstat_penddata            uint16
	Sstat_instrms             uint16
	Sstat_outstrms            uint16
	Sstat_fragmentation_point uint32
	Sstat_primary             [152]byte
}

//enum sctp_sstat_state {
//	SCTP_EMPTY                = 0,
//	SCTP_CLOSED               = 1,
//	SCTP_COOKIE_WAIT          = 2,
//	SCTP_COOKIE_ECHOED        = 3,
//	SCTP_ESTABLISHED          = 4,
//	SCTP_SHUTDOWN_PENDING     = 5,
//	SCTP_SHUTDOWN_SENT        = 6,
//	SCTP_SHUTDOWN_RECEIVED    = 7,
//	SCTP_SHUTDOWN_ACK_SENT    = 8,
//};
type SctpConnStatus int32

const (
	SCTP_EMPTY SctpConnStatus = iota
	SCTP_CLOSED
	SCTP_COOKIE_WAIT
	SCTP_COOKIE_ECHOED
	SCTP_ESTABLISHED
	SCTP_SHUTDOWN_PENDING
	SCTP_SHUTDOWN_SENT
	SCTP_SHUTDOWN_RECEIVED
	SCTP_SHUTDOWN_ACK_SENT
)

func (p SctpConnStatus) String() string {
	var str string
	switch p {
	case SCTP_EMPTY:
		str = "sctp_empty"
	case SCTP_CLOSED:
		str = "sctp_closed"
	case SCTP_COOKIE_WAIT:
		str = "sctp_cookie_wait"
	case SCTP_COOKIE_ECHOED:
		str = "sctp_cookie_echoed"
	case SCTP_ESTABLISHED:
		str = "sctp_established"
	case SCTP_SHUTDOWN_PENDING:
		str = "sctp_shutdown_pending"
	case SCTP_SHUTDOWN_SENT:
		str = "sctp_shutdown_send"
	case SCTP_SHUTDOWN_RECEIVED:
		str = "sctp_shutdown_received"
	case SCTP_SHUTDOWN_ACK_SENT:
		str = "sctp_shutdown_ack_send"
	default:
		str = "unknown_status"
	}
	return str
}

//0 struct sctp_paddrparams {
//  541     sctp_assoc_t        spp_assoc_id;
//  542     struct sockaddr_storage spp_address;
//  543     __u32           spp_hbinterval;
//  544     __u16           spp_pathmaxrxt;
//  545     __u32           spp_pathmtu;
//  546     __u32           spp_sackdelay;
//  547     __u32           spp_flags;
//  548 } __attribute__((packed, aligned(4)));
type SctpPaddrParams struct {
	SppAssocId    uint32
	SppAddr       [128]byte
	SppHbInterval uint32
	SppPathMaxRxt uint16
	SppPathMtu    uint32
	SppSackDelay  uint32
	SppFlags      uint32
}

//enum  sctp_spp_flags {
//	SPP_HB_ENABLE = 1<<0,		/*Enable heartbeats*/
//	SPP_HB_DISABLE = 1<<1,		/*Disable heartbeats*/
//	SPP_HB = SPP_HB_ENABLE | SPP_HB_DISABLE,
//	SPP_HB_DEMAND = 1<<2,		/*Send heartbeat immediately*/
//	SPP_PMTUD_ENABLE = 1<<3,	/*Enable PMTU discovery*/
//	SPP_PMTUD_DISABLE = 1<<4,	/*Disable PMTU discovery*/
//	SPP_PMTUD = SPP_PMTUD_ENABLE | SPP_PMTUD_DISABLE,
//	SPP_SACKDELAY_ENABLE = 1<<5,	/*Enable SACK*/
//	SPP_SACKDELAY_DISABLE = 1<<6,	/*Disable SACK*/
//	SPP_SACKDELAY = SPP_SACKDELAY_ENABLE | SPP_SACKDELAY_DISABLE,
//	SPP_HB_TIME_IS_ZERO = 1<<7,	/* Set HB delay to 0 */
//};
const (
	SPP_HB_ENABLE         = 1 << 0
	SPP_HB_DISABLE        = 1 << 1
	SPP_HB                = SPP_HB_ENABLE | SPP_HB_DISABLE
	SPP_HB_DEMAND         = 1 << 2
	SPP_PMTUD_ENABLE      = 1 << 3
	SPP_PMTUD_DISABLE     = 1 << 4
	SPP_PMTUD             = SPP_PMTUD_ENABLE | SPP_PMTUD_DISABLE
	SPP_SACKDELAY_ENABLE  = 1 << 5
	SPP_SACKDELAY_DISABLE = 1 << 6
	SPP_SACKDELAY         = SPP_SACKDELAY_ENABLE | SPP_SACKDELAY_DISABLE
	SPP_HB_TIME_IS_ZERO   = 1 << 7
)

// /*
// * 5.3.1.2 SCTP_PEER_ADDR_CHANGE
// *
// *   When a destination address on a multi-homed peer encounters a change
// *   an interface details event is sent.  The information has the
// *   following structure:
// */
//struct sctp_paddr_change {
//	__u16 spc_type;
//	__u16 spc_flags;
//	__u32 spc_length;
//	struct sockaddr_storage spc_aaddr;
//	int spc_state;
//	int spc_error;
//	sctp_assoc_t spc_assoc_id;
//} __attribute__((packed, aligned(4)));
type SctpPeerAddrChange struct {
	SnType     uint16
	SnFlags    uint16
	SnLength   uint32
	SpcAddr    [128]byte
	SpcState   uint32
	SpcError   uint32
	SpcAssocId uint32
}

func (p *SctpPeerAddrChange) Read(b []byte) error {

	r := bytes.NewReader(b)
	// sn type
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

	// spc addr
	data128 := make([]byte, 128)
	err = binary.Read(r, binary.LittleEndian, &data128)
	if err != nil {
		return fmt.Errorf("failed to read sctp spc addr, error(%s)", err)
	}
	for i, v := range data128 {
		p.SpcAddr[i] = v
	}

	// Spc State
	dataU32 = make([]byte, 4)
	err = binary.Read(r, binary.LittleEndian, &dataU32)
	if err != nil {
		return fmt.Errorf("failed to read sctp spc state, error(%s)", err)
	}
	p.SpcState = binary.LittleEndian.Uint32(dataU32)

	// Spc Error
	dataU32 = make([]byte, 4)
	err = binary.Read(r, binary.LittleEndian, &dataU32)
	if err != nil {
		return fmt.Errorf("failed to read sctp spc error, error(%s)", err)
	}
	p.SpcError = binary.LittleEndian.Uint32(dataU32)

	// SpcAssocId
	dataU32 = make([]byte, 4)
	err = binary.Read(r, binary.LittleEndian, &dataU32)
	if err != nil {
		return fmt.Errorf("failed to read sctp assoc id, error(%s)", err)
	}
	p.SpcAssocId = binary.LittleEndian.Uint32(dataU32)
	return nil
}

//enum sctp_spc_state {
//	SCTP_ADDR_AVAILABLE,
//	SCTP_ADDR_UNREACHABLE,
//	SCTP_ADDR_REMOVED,
//	SCTP_ADDR_ADDED,
//	SCTP_ADDR_MADE_PRIM,
//	SCTP_ADDR_CONFIRMED,
//};
type SctpSpcState uint32

const (
	SCTP_ADDR_AVAILABLE SctpSpcState = iota
	SCTP_ADDR_UNREACHABLE
	SCTP_ADDR_REMOVED
	SCTP_ADDR_ADDED
	SCTP_ADDR_MADE_PRIM
	SCTP_ADDR_CONFIRMED
)

func (p SctpSpcState) String() string {
	var str string
	switch p {
	case SCTP_ADDR_AVAILABLE:
		str = "sctp_addr_available"
	case SCTP_ADDR_UNREACHABLE:
		str = "sctp_addr_unreachable"
	case SCTP_ADDR_REMOVED:
		str = "sctp_addr_removed"
	case SCTP_ADDR_ADDED:
		str = "sctp_addr_added"
	case SCTP_ADDR_MADE_PRIM:
		str = "sctp_addr_made_prim"
	case SCTP_ADDR_CONFIRMED:
		str = "sctp_addr_confirmed"
	default:
		str = "unknown_status"
	}
	return str
}

// /*
// * 5.3.1.1 SCTP_ASSOC_CHANGE
// *
// *   Communication notifications inform the ULP that an SCTP association
// *   has either begun or ended. The identifier for a new association is
// *   provided by this notificaion. The notification information has the
// *   following format:
// *
// */
//struct sctp_assoc_change {
//	__u16 sac_type;
//	__u16 sac_flags;
//	__u32 sac_length;
//	__u16 sac_state;
//	__u16 sac_error;
//	__u16 sac_outbound_streams;
//	__u16 sac_inbound_streams;
//	sctp_assoc_t sac_assoc_id;
//	__u8 sac_info[0];
//};
type SctpAssocChange struct {
	SnType             uint16
	SnFlags            uint16
	SnLength           uint32
	SacState           uint16
	SacError           uint16
	SacOutboundStreams uint16
	SacInboundStreams  uint16
	SacAssocId         uint32
	SacInfo            uint8
}

func (p *SctpAssocChange) Read(b []byte) error {

	r := bytes.NewReader(b)
	// sn type
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

	// Sac State
	data16 := make([]byte, 2)
	err = binary.Read(r, binary.LittleEndian, &data16)
	if err != nil {
		return fmt.Errorf("failed to read sctp sac state, error(%s)", err)
	}
	p.SacState = binary.LittleEndian.Uint16(data16)

	// Sac error
	data16 = make([]byte, 2)
	err = binary.Read(r, binary.LittleEndian, &data16)
	if err != nil {
		return fmt.Errorf("failed to read sctp sac error, error(%s)", err)
	}
	p.SacError = binary.LittleEndian.Uint16(data16)

	// Sac Outbound Streams
	data16 = make([]byte, 2)
	err = binary.Read(r, binary.LittleEndian, &data16)
	if err != nil {
		return fmt.Errorf("failed to read sctp sac outbound streams, error(%s)", err)
	}
	p.SacOutboundStreams = binary.LittleEndian.Uint16(data16)

	// Sac Inbound Streams
	data16 = make([]byte, 2)
	err = binary.Read(r, binary.LittleEndian, &data16)
	if err != nil {
		return fmt.Errorf("failed to read sctp sac inbound streams, error(%s)", err)
	}
	p.SacInboundStreams = binary.LittleEndian.Uint16(data16)

	// SpcAssocId
	dataU32 = make([]byte, 4)
	err = binary.Read(r, binary.LittleEndian, &dataU32)
	if err != nil {
		return fmt.Errorf("failed to read sctp assoc id, error(%s)", err)
	}
	p.SacAssocId = binary.LittleEndian.Uint32(dataU32)

	return nil
}

//enum sctp_sac_state {
//	SCTP_COMM_UP,
//	SCTP_COMM_LOST,
//	SCTP_RESTART,
//	SCTP_SHUTDOWN_COMP,
//	SCTP_CANT_STR_ASSOC,
//};
type SctpSacState uint16

const (
	SCTP_COMM_UP SctpSacState = iota
	SCTP_COMM_LOST
	SCTP_RESTART
	SCTP_SHUTDOWN_COMP
	SCTP_CANT_STR_ASSOC
)

func (p SctpSacState) String() string {
	var rtStr string
	switch p {
	case SCTP_COMM_UP:
		rtStr = "sctp_comm_up"
	case SCTP_COMM_LOST:
		rtStr = "sctp_comm_lost"
	case SCTP_RESTART:
		rtStr = "sctp_restart"
	case SCTP_SHUTDOWN_COMP:
		rtStr = "sctp_shutdown_complete"
	case SCTP_CANT_STR_ASSOC:
		rtStr = "sctp_cant_str_accoc"
	default:
		rtStr = "unknown state"
	}
	return rtStr
}
