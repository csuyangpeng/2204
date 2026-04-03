package pfcpv1

import (
	"bytes"
	"fmt"
	"lite5gc/cmn/message/pfcp"
)

type Message struct {
	Header pfcp.PfcpHeader
	Body   interface{}
}

func (m *Message) String() string {
	w := bytes.NewBuffer(nil)
	fmt.Fprintf(w, "Version: %v\n", m.Header.Version)
	fmt.Fprintf(w, "MP: %v\n", m.Header.MPFlag)
	fmt.Fprintf(w, "S: %v\n", m.Header.SFlag)
	fmt.Fprintf(w, "Message Type: %v\n", m.Header.MessageType)
	fmt.Fprintf(w, "Length: %v\n", m.Header.Length)
	if m.Header.SFlag == 1 {
		fmt.Fprintf(w, "SEID: %v\n", m.Header.SEID)
	}
	fmt.Fprintf(w, "SequenceNumber: %v\n", m.Header.SequenceNumber)
	if m.Header.MPFlag == 1 {
		fmt.Fprintf(w, "MessagePriority: %v\n", m.Header.MessagePriority)
	}

	return w.String()
}

func (m *Message) HeaderSet(pfcpHead interface{}) error {
	switch h := pfcpHead.(type) {
	case pfcp.PfcpHeaderforSession:
		m.Header.Version = h.Version
		m.Header.SFlag = h.SFlag
		m.Header.MessageType = h.MessageType
		m.Header.SEID = h.SEID
		m.Header.SequenceNumber = h.SequenceNumber

	case pfcp.PfcpHeaderforNode:
		m.Header.Version = h.Version
		m.Header.MessageType = h.MessageType
		m.Header.SequenceNumber = h.SequenceNumber
	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}
func (m *Message) HeaderFillToMsg(pfcpHead interface{}) error {
	switch h := pfcpHead.(type) {
	case *pfcp.PfcpHeaderforSession:
		h.Version = m.Header.Version
		h.SFlag = m.Header.SFlag
		h.MessageType = m.Header.MessageType
		h.SEID = m.Header.SEID
		h.SequenceNumber = m.Header.SequenceNumber

	case *pfcp.PfcpHeaderforNode:
		h.Version = m.Header.Version
		h.MessageType = m.Header.MessageType
		h.SequenceNumber = m.Header.SequenceNumber
	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}
func (m *Message) BodySet(pfcpMsg interface{}) error {
	m.Body = pfcpMsg
	return nil
}
