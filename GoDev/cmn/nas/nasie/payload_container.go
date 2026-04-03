package nasie

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

// use this data struct only when PayloadType == nasie.MultiplePayload
// if PayloadType
type PayloadContainerIE struct {
	NumberOfEntry         byte
	PayloadContainerEntry []PayloadContainerENTRY
}

type PayloadContainerENTRY struct {
	NumberOfOptIEs       byte
	PayloadContainerType PayloadContainerType
	OptionalIEs          []OptionalIE
	ContainerContents    []byte
}

type OptionalIE struct {
	IEType   TypeOfOptIE
	IELength byte
	IEValue  []byte
}

func (p *PayloadContainerIE) Reset() {
	p.NumberOfEntry = 0
	for i := 0; i < len(p.PayloadContainerEntry); i++ {
		p.PayloadContainerEntry[i].Reset()
	}
}

func (p *PayloadContainerENTRY) Reset() {
	p.NumberOfOptIEs = 0
	p.PayloadContainerType = 0
	for i := 0; i < len(p.OptionalIEs); i++ {
		p.OptionalIEs[i].Reset()
	}
	p.ContainerContents = []byte{}
}

func (p *OptionalIE) Reset() {
	p.IEType = 0
	p.IELength = 0
	p.IEValue = []byte{}
}

//This field contains the IEI of the optional IE entry and is 1 octet in length.
//IEI 	Optional IE name	Optional IE reference
//12	PDU session ID	PDU session identity 2 (see subclause 9.11.3.41)
//24	Additional information	Additional information (see subclause 9.11.2.1)
//58	5GMM cause	5GMM cause (see subclause 9.11.3.2)
//37	Back-off timer value	GPRS timer 3 (see subclause 9.11.2.5)
//59	Old PDU session ID	PDU session identity (see subclause 2 9.11.3.41)
//8-	Request type	Request type (see subclause 9.11.3.47)
//22	S-NSSAI	S-NSSAI (see subclause 9.11.2.8)
//25	DNN	DNN (see subclause 9.11.3.1A)
type TypeOfOptIE byte

const (
	PDU_Session_ID TypeOfOptIE = iota
	Additional_Information
	FiveGMM_cause
	Back_Off_Timer_Value
	Old_PDU_Session_ID
	Request_Type
	S_NSSAI
	DNN_Type
)

func (p *PayloadContainerIE) Encode() []byte {
	rlogger.FuncEntry(types.ModCmn, nil)
	//no Length, start from Value
	var encBuf []byte

	encBuf = append(encBuf, p.NumberOfEntry)

	for i := 0; i < int(p.NumberOfEntry); i++ {
		entryI := p.PayloadContainerEntry[i]

		octet2 := byte(entryI.NumberOfOptIEs) << 4
		octet2 |= byte(entryI.PayloadContainerType)

		var optBuf []byte
		for j := 0; j < int(entryI.NumberOfOptIEs); j++ {
			optIEJ := entryI.OptionalIEs[j]
			optBuf = append(optBuf, byte(optIEJ.IEType))
			optBuf = append(optBuf, byte(optIEJ.IELength))
			optBuf = append(optBuf, optIEJ.IEValue...)
		}

		var eBuf []byte
		eBuf = append(eBuf, byte(1+len(optBuf)+len(entryI.ContainerContents)))
		eBuf = append(eBuf, octet2)
		eBuf = append(eBuf, optBuf...)
		eBuf = append(eBuf, entryI.ContainerContents...)

		encBuf = append(encBuf, eBuf...)
	}

	return encBuf
}

func (p *PayloadContainerIE) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	//start from Length

	//Length of payload container contents
	lenContainer := make([]byte, 2)
	err := binary.Read(msgBuf, binary.BigEndian, &lenContainer)
	if err != nil {
		rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "fail to read length of payLoad container")
		return fmt.Errorf("fail to read length of payLoad container")
	}

	//Number of entries
	p.NumberOfEntry, _ = msgBuf.ReadByte()

	for i := 0; i < int(p.NumberOfEntry); i++ {
		entryI := p.PayloadContainerEntry[i]

		//Length of Payload container entry
		//lenOfEntry, _ := msgBuf.ReadByte()
		lenOfEntry := make([]byte, 2)
		err := binary.Read(msgBuf, binary.BigEndian, &lenOfEntry)
		if err != nil {
			rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "fail to read length of payLoad container entry")
			return fmt.Errorf("fail to read length of payLoad container entry")
		}
		lenEntry := binary.BigEndian.Uint16(lenOfEntry)
		//Number of optional IEs && Payload container type
		octet2, _ := msgBuf.ReadByte()
		typeByte, _ := utils.GetBitsValue(octet2, 1, 4)
		entryI.PayloadContainerType = PayloadContainerType(typeByte)
		entryI.NumberOfOptIEs, _ = utils.GetBitsValue(octet2, 5, 8)

		var lenOfOptIE int
		for j := 0; j < int(entryI.NumberOfOptIEs); j++ {
			optIEJ := entryI.OptionalIEs[j]
			//Type of optional IE
			ieTypeByte, _ := msgBuf.ReadByte()
			optIEJ.IEType = TypeOfOptIE(ieTypeByte)
			//Length of optional IE
			optIEJ.IELength, _ = msgBuf.ReadByte()
			//Value of optional IE
			optIEJ.IEValue = make([]byte, optIEJ.IELength)
			err = binary.Read(msgBuf, binary.BigEndian, &optIEJ.IEValue)
			if err != nil {
				rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "fail to read value of optional IE")
				return fmt.Errorf("fail to read value of optional IE")
			}
			lenOfOptIE += int(optIEJ.IELength) + 2
		}
		//Payload container contents
		entryI.ContainerContents = make([]byte, int(lenEntry)-1-lenOfOptIE)
		err = binary.Read(msgBuf, binary.BigEndian, &entryI.ContainerContents)
		if err != nil {
			rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "fail to read contents of Payload container entry")
			return fmt.Errorf("fail to read contents of Payload container entry")
		}
	}

	return nil
}
