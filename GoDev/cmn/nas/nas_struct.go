package nas

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

//According to 24.007
//EPD value (octet 1, bit 1 to bit 8)
//Bits
//8	7	6	5	4	3	2	1
//0	0	0	0	1	1	1	0	reserved
//0	0	0	1	1	1	1	0	reserved
//0	0	1	0	1	1	1	0	5GS session management messages
//0	0	1	1	1	1	1	0	reserved
//0	1	0	0	1	1	1	0	reserved
//0	1	0	1	1	1	1	0	reserved
//0	1	1	0	1	1	1	0	reserved
//0	1	1	1	1	1	1	0	5GS mobility management messages
type Epd byte

const (
	Epd5gsMobMgntMsg  Epd = 0x7E
	Epd5gsSessMgntMsg Epd = 0x2E
)

func (p *Epd) IsValid() error {
	if Epd5gsMobMgntMsg != *p || Epd5gsSessMgntMsg != *p {
		return ErrInvalidEpd
	}
	return nil
}

// GetEpd, read the epd byte from NAS octet stream
func GetEpd(msgBuf *bytes.Reader) (Epd, error) {
	val, err := msgBuf.ReadByte()
	if err != nil {
		return 0, fmt.Errorf("failed to read a byte from reader, "+
			"error:%s", err)
	}
	epd := Epd(val)
	if epd != Epd5gsMobMgntMsg && epd != Epd5gsSessMgntMsg {
		return 0, fmt.Errorf("invalid epd from reader. epd: %d", epd)
	}
	return epd, nil
}

func (p Epd) String() string {
	switch p {
	case Epd5gsMobMgntMsg:
		return "MobMgntMsg"
	case Epd5gsSessMgntMsg:
		return "SessMgntMsg"
	default:
		return "unknown"
	}
}

//According to 24.501 Table 9.3.1
//Security header type (octet 1)
//Bits
//4	3	2	1
//0	0	0	0	Plain 5GS NAS message, not security protected
//Security protected 5GS NAS message:
//0	0	0	1	Integrity protected
//0	0	1	0	Integrity protected and ciphered
//0	0	1	1	Integrity protected with new 5G NAS security context (NOTE 1)
//0	1	0	0	Integrity protected and ciphered with new 5G NAS security context (NOTE 2)
type SecHeaderType byte

const (
	PlainNasMsg                    SecHeaderType = 0
	IntegrityPrtc                  SecHeaderType = 1
	IntegrityPrtctCipher           SecHeaderType = 2
	IntegrityPrtctNewSecCtxt       SecHeaderType = 3
	IntegrityPrtctCipherNewSecCtxt SecHeaderType = 4
	MaxSecHeaderTypeValue          SecHeaderType = 5
)

func (p *SecHeaderType) IsValid() error {
	if (*p & 0x0F) >= MaxSecHeaderTypeValue {
		return ErrInvalidSecHeaderType
	}
	return nil
}

// GetSecHeaderType, read the security header type byte from NAS octet stream
func GetSecHeaderType(msgBuf *bytes.Reader) (SecHeaderType, error) {
	val, err := msgBuf.ReadByte()
	if err != nil {
		return 0, fmt.Errorf("failed to read a byte from reader, "+
			"error:%s", err)
	}
	headerType := SecHeaderType(val)
	if headerType != PlainNasMsg &&
		headerType != IntegrityPrtc &&
		headerType != IntegrityPrtctCipher &&
		headerType != IntegrityPrtctNewSecCtxt &&
		headerType != IntegrityPrtctCipherNewSecCtxt {
		return 0, fmt.Errorf("invalid epd from reader. "+
			"security header type: %d", headerType)
	}
	return headerType, nil
}

//According to 24.501 Table 9.7.1
//Bits
//8	7	6	5	4	3	2	1
//
//0	1	-	-	-	-	-	-		5GS mobility management messages
//
//0	1	0	0	0	0	0	1		Registration request
//0	1	0	0	0	0	1	0		Registration accept
//0	1	0	0	0	0	1	1		Registration complete
//0	1	0	0	0	1	0	0		Registration reject
//0	1	0	0	0	1	0	1		Deregistration request (UE originating)
//0	1	0	0	0	1	1	0		Deregistration accept (UE originating)
//0	1	0	0	0	1	1	1		Deregistration request (UE terminated)
//0	1	0	0	1	0	0	0		Deregistration accept (UE terminated)
//
//0	1	0	0	1	1	0	0		Service request
//0	1	0	0	1	1	0	1		Service reject
//0	1	0	0	1	1	1	0		Service accept
//
//0	1	0	1	0	1	0	0		Configuration update command
//0	1	0	1	0	1	0	1		Configuration update complete
//0	1	0	1	0	1	1	0		Authentication request
//0	1	0	1	0	1	1	1		Authentication response
//0	1	0	1	1	0	0	0		Authentication reject
//0	1	0	1	1	0	0	1		Authentication failure
//0	1	0	1	1	0	1	0		Authentication result
//0	1	0	1	1	0	1	1		Identity request
//0	1	0	1	1	1	0	0		Identity response
//0	1	0	1	1	1	0	1		Security mode command
//0	1	0	1	1	1	1	0		Security mode complete
//0	1	0	1	1	1	1	1		Security mode reject
//
//0	1	1	0	0	1	0	0		5GMM status
//0	1	1	0	0	1	0	1		Notification
//0	1	1	0	0	1	1	0		Notification response
//0	1	1	0	0	1	1	1		UL NAS transport
//0	1	1	0	1	0	0	0		DL NAS transport

type MsgType byte

type MmMsgType byte

const (
	RegistrationRequest      MmMsgType = 0x41
	RegistrationAccept       MmMsgType = 0x42
	RegistrationComplete     MmMsgType = 0x43
	RegistrationReject       MmMsgType = 0x44
	DeregistrationRequestUe  MmMsgType = 0x45
	DeregistrationAcceptUe   MmMsgType = 0x46
	DeregistrationRequestUeT MmMsgType = 0x47
	DeregistrationAcceptUeT  MmMsgType = 0x48
	ServiceRequest           MmMsgType = 0x4C
	ServiceReject            MmMsgType = 0x4D
	ServiceAccept            MmMsgType = 0x4E
	ConfigUpdateCommand      MmMsgType = 0x54
	ConfigUpdateComplete     MmMsgType = 0x55
	AuthenticationRequest    MmMsgType = 0x56
	AuthenticationResponse   MmMsgType = 0x57
	AuthenticationReject     MmMsgType = 0x58
	AuthenticationFailure    MmMsgType = 0x59
	AuthenticationResult     MmMsgType = 0x5A
	IdentifyRequest          MmMsgType = 0x5B
	IdentifyResponse         MmMsgType = 0x5C
	SecurityModeCommand      MmMsgType = 0x5D
	SecurityModeComplete     MmMsgType = 0x5E
	SecurityModeReject       MmMsgType = 0x5F
	MmStatus5G               MmMsgType = 0x64
	Notification             MmMsgType = 0x65
	NotificationResponse     MmMsgType = 0x66
	ULNasTransport           MmMsgType = 0x67
	DLNasTransport           MmMsgType = 0x68

	//for network-anrelease command
	AnReleaseCmd MmMsgType = 0xFF
)

func (p *MmMsgType) IsValid() error {
	//TODO
	return nil
}

type MmNasMessageHeader struct {
	ExtendProtoDisc    Epd
	SecurityHeaderType SecHeaderType
	MessageType        MmMsgType
}

func (p *MmNasMessageHeader) IsValid() error {
	if (p.SecurityHeaderType.IsValid() != nil) &&
		(p.MessageType.IsValid() != nil) {
		return ErrInvalidMmNasMsgHeader
	}
	return nil
}

// GetEpd, read the epd byte from NAS octet stream
func (p *MmNasMessageHeader) Decode(msgBuf *bytes.Reader) error {

	err := binary.Read(msgBuf, binary.LittleEndian, p)
	if err != nil {
		return fmt.Errorf("failed to read mm nas msg header, "+
			"error:%s", err)
	}

	return nil
}

//Encode
func (p *MmNasMessageHeader) Encode() []byte {
	var encBuf []byte
	encBuf = append(encBuf, byte(p.ExtendProtoDisc))
	encBuf = append(encBuf, byte(p.SecurityHeaderType))
	encBuf = append(encBuf, byte(p.MessageType))

	return encBuf
}

// security protected nas 5gs message
type MmNasSecMessageHeader struct {
	ExtendProtoDisc    Epd
	SecurityHeaderType SecHeaderType
	Mac                [4]byte
	Sqn                uint8
}

// Encode
func (p *MmNasSecMessageHeader) Encode() []byte {
	var encBuf []byte
	encBuf = append(encBuf, byte(p.ExtendProtoDisc))
	encBuf = append(encBuf, byte(p.SecurityHeaderType))
	encBuf = append(encBuf, p.Mac[:]...)
	encBuf = append(encBuf, p.Sqn)
	return encBuf
}

// Decode
func (p *MmNasSecMessageHeader) Decode(msgBuf *bytes.Reader) error {
	err := binary.Read(msgBuf, binary.BigEndian, p)
	if err != nil {
		return fmt.Errorf("failed to read mm nas msg header, "+
			"error:%s", err)
	}
	return nil
}

//According to 24.501 Table 9.7.2
//Bits
//8	7	6	5	4	3	2	1
//
//1	1	-	-	-	-	-	-		5GS session management messages
//
//1	1	0	0	0	0	0	1		PDU session establishment request
//1	1	0	0	0	0	1	0		PDU session establishment accept
//1	1	0	0	0	0	1	1		PDU session establishment reject
//
//1	1	0	0	0	1	0	1		PDU session authentication command
//1	1	0	0	0	1	1	0		PDU session authentication complete
//1	1	0	0	0	1	1	1		PDU session authentication result
//
//1	1	0	0	1	0	0	1		PDU session modification request
//1	1	0	0	1	0	1	0		PDU session modification reject
//1	1	0	0	1	0	1	1		PDU session modification command
//1	1	0	0	1	1	0	0		PDU session modification complete
//1	1	0	0	1	1	0	1		PDU session modification command reject
//
//1	1	0	1	0	0	0	1		PDU session release request
//1	1	0	1	0	0	1	0		PDU session release reject
//1	1	0	1	0	0	1	1		PDU session release command
//1	1	0	1	0	1	0	0		PDU session release complete
//
//1	1	0	1	0	1	1	0		5GSM status
type SmMsgType byte

const (
	PduSessEstabishRequest     SmMsgType = 0xC1
	PduSessEstabishAccept      SmMsgType = 0xC2
	PduSessEstabishReject      SmMsgType = 0xC3
	PduSessionAuthCommand      SmMsgType = 0xC5
	PduSessionAuthComplete     SmMsgType = 0xC6
	PduSessionAuthResult       SmMsgType = 0xC7
	PduSessionModRequest       SmMsgType = 0xC9
	PduSessionModReject        SmMsgType = 0xCA
	PduSessionModCommand       SmMsgType = 0xCB
	PduSessionModComplete      SmMsgType = 0xCC
	PduSessionModCommandReject SmMsgType = 0xCD
	PduSessionRelRequest       SmMsgType = 0xD1
	PduSessionRelReject        SmMsgType = 0xD2
	PduSessionRelCommand       SmMsgType = 0xD3
	PduSessionRelComplete      SmMsgType = 0xD4
	SMStatus5G                 SmMsgType = 0xD6
)

func (p SmMsgType) String() string {
	switch p {
	case PduSessEstabishRequest:
		return "PduSessEstabishRequest"
	case PduSessEstabishAccept:
		return "PduSessEstabishAccept"
	case PduSessEstabishReject:
		return "PduSessEstabishReject"
	case PduSessionAuthCommand:
		return "PduSessionAuthCommand"
	case PduSessionAuthComplete:
		return "PduSessionAuthComplete"
	case PduSessionAuthResult:
		return "PduSessionAuthResult"
	case PduSessionModRequest:
		return "PduSessionModRequest"
	case PduSessionModReject:
		return "PduSessionModReject"
	case PduSessionModCommand:
		return "PduSessionModCommand"
	case PduSessionModComplete:
		return "PduSessionModComplete"
	case PduSessionModCommandReject:
		return "PduSessionModCommandReject"
	case PduSessionRelRequest:
		return "PduSessionRelRequest"
	case PduSessionRelReject:
		return "PduSessionRelReject"
	case PduSessionRelCommand:
		return "PduSessionRelCommand"
	case PduSessionRelComplete:
		return "PduSessionRelComplete"
	case SMStatus5G:
		return "SMStatus5G"
	default:
		return "unsupported sm message type"
	}
}
func (p SmMsgType) IsValid() error {
	switch p {
	case PduSessEstabishRequest:
	case PduSessEstabishAccept:
	case PduSessEstabishReject:
	case PduSessionAuthCommand:
	case PduSessionAuthComplete:
	case PduSessionAuthResult:
	case PduSessionModRequest:
	case PduSessionModReject:
	case PduSessionModCommand:
	case PduSessionModComplete:
	case PduSessionModCommandReject:
	case PduSessionRelRequest:
	case PduSessionRelReject:
	case PduSessionRelCommand:
	case PduSessionRelComplete:
	case SMStatus5G:
	default:
		return fmt.Errorf("invalid sm message type(%x)", p)
	}
	return nil
}

//PDU Session Identity value according to 24.007
type PrcdTransID byte

const (
	MinPTI PrcdTransID = 0x01
	MaxPTI PrcdTransID = 0xFE
)

func (p *PrcdTransID) IsValid() error {
	if *p > MaxPTI ||
		*p < MinPTI {
		return ErrInvalidPtiType
	}
	return nil
}

//PDU Session Identity according to 24.007
//the rest values are reserved
type PduSessID byte

const (
	MinPSI PduSessID = 1
	MaxPSI PduSessID = 15
)

func (p *PduSessID) IsValid() error {
	if *p > MaxPSI ||
		*p < MinPSI {
		return ErrInvalidPsiType
	}
	return nil
}

type SmNasMessageHeader struct {
	PduSessionID      PduSessID
	PrcdTransactionID PrcdTransID
	MessageType       SmMsgType
}

func (p *SmNasMessageHeader) IsValid() error {
	if (p.PduSessionID.IsValid() != nil) &&
		(p.PrcdTransactionID.IsValid() != nil) &&
		(p.MessageType.IsValid() != nil) {
		return ErrInvalidSmNasMsgHeader
	}
	return nil
}

// GetEpd, read the epd byte from NAS octet stream
func (p *SmNasMessageHeader) Decode(msgBuf *bytes.Reader) error {
	// PDU session ID
	PduSessionID, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read sm nas msg header psi, "+
			"error:%s", err)
	}
	p.PduSessionID = PduSessID(PduSessionID)

	// PTI,Procedure transaction identity
	PrcdTransactionID, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read sm nas msg header pti, "+
			"error:%s", err)
	}
	p.PrcdTransactionID = PrcdTransID(PrcdTransactionID)

	// Message type
	MessageType, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read sm nas msg header msg type, "+
			"error:%s", err)
	}
	p.MessageType = SmMsgType(MessageType)

	return nil
}

//encode SmNasMessageHeader to nas octet stream
func (p *SmNasMessageHeader) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte
	encBuf = append(encBuf, byte(p.PduSessionID))
	encBuf = append(encBuf, byte(p.PrcdTransactionID))
	encBuf = append(encBuf, byte(p.MessageType))

	return encBuf, nil
}
func (p *SmNasMessageHeader) String() string {
	return fmt.Sprintf("nas header: pdu session id(%d), pti(%d), msg type(%s)",
		p.PduSessionID, p.PrcdTransactionID, p.MessageType)
}

type UeUsage uint8

const (
	VoiceCentric UeUsage = 0
	DataCentric  UeUsage = 1
)
