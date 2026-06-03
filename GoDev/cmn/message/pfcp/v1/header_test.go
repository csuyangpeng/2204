package pfcpv1

import (
	"testing"
	"zj-test/pfcpRef/pfcpv1"
)

const (
	NodeRelatedHeaderHex    = "2005000000000100"
	SessionRelatedHeaderHex = "233400000000000000000001000000F0"
)

var NodeRelatedHeader = pfcpv1.Header{
	Version:        1,
	S:              0,
	MessageType:    pfcpv1.PFCP_ASSOCIATION_SETUP_REQUEST,
	MessageLength:  0,
	SequenceNumber: 1,
}

var SessionRelatedHeader = pfcpv1.Header{
	Version:         1,
	MP:              1,
	S:               1,
	MessageType:     pfcpv1.PFCP_SESSION_MODIFICATION_REQUEST,
	MessageLength:   0,
	SEID:            1,
	SequenceNumber:  0,
	MessagePriority: 15,
}

func TestPFCPHeader_MarshalBinary(t *testing.T) {}

func TestPFCPHeader_UnmarshalBinary(t *testing.T) {}
