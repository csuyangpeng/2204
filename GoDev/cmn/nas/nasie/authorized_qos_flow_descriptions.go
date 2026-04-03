package nasie

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

const (
	MinQFI = 0
	MaxQFI = 63
)

type OperationCodeIE byte

const (
	CreateNewQoSFlowDescription      OperationCodeIE = 1
	DeleteExistingQoSFlowDescription OperationCodeIE = 2
	ModifyExistingQoSFlowDescription OperationCodeIE = 3
)

func (p *OperationCodeIE) StoreWithString(val string) error {
	switch val {
	case "create new qos flow description":
		*p = CreateNewQoSFlowDescription
	case "delete existing qos flow description":
		*p = DeleteExistingQoSFlowDescription
	case "modify existing qos flow description":
		*p = ModifyExistingQoSFlowDescription
	default:
		return fmt.Errorf("invalid qos flow desc opr code(%s)", val)
	}
	return nil
}

// refers to 24.501 9.11.4.12
type QoSFlowsDesc struct {
	Descr []QoSFlowDescription
}

func (p QoSFlowsDesc) String() string {
	var rt string
	for i, v := range p.Descr {
		rt = rt + fmt.Sprintf("\nQosRule{%d : %s}", i+1, v)
	}
	rt = rt + fmt.Sprintf("\n")
	return rt
}

type QoSFlowDescription struct {
	QFI                byte
	OperationCode      OperationCodeIE
	E                  bool
	NumberOfParameters byte
	ParameterList      ParametersIEs
}

func (p QoSFlowDescription) String() string {
	return fmt.Sprintf(
		"QFI(%v),"+
			"OperationCode(%v),"+
			"E(%v),"+
			"NumberOfParameters(%v),"+
			"ParameterList(%v),",
		p.QFI,
		p.OperationCode,
		p.E,
		p.NumberOfParameters,
		p.ParameterList)
}

type ParametersIEs struct {
	ParmsList []ParametersIE
}

func (p ParametersIE) String() string {
	var s string
	s += fmt.Sprintf("Parameter ID(%v),", p.ParameterID)
	switch p.ParameterID {
	case FiveQI:
		s += fmt.Sprintf("QI5(%v);", p.QI5Content)
	case GFBRUplink:
		s += fmt.Sprintf("GFBR Uplink(%v);", p.GFBRUplinkContent)
	case GFBRDownlink:
		s += fmt.Sprintf("GFBR Downlink(%v);", p.GFBRDownlinkContent)
	case MFBRUplink:
		s += fmt.Sprintf("MFBR Uplink(%v);", p.MFBRUplinkContent)
	case MFBRDownlink:
		s += fmt.Sprintf("MFBR Downlink(%v);", p.MFBRDownlinkContent)
	}
	return s
}

type ParametersIE struct {
	ParameterID ParameterIdentifier
	//Parameter contents
	QI5Content QI5Contents
	//GBR
	GFBRUplinkContent   BitRate
	GFBRDownlinkContent BitRate
	MFBRUplinkContent   BitRate
	MFBRDownlinkContent BitRate
	//the averaging window for both uplink and downlink in milliseconds
	AveragingWindowContent uint16
	//38.413 9.3.1.82	Averaging Window
	//This IE indicates the Averaging Window for a QoS flow.
	//IE/Group Name	Presence	Range	IE type and reference	Semantics description
	//Averaging Window	M		INTEGER (0..4095, …)	Unit: ms.
	//The default value of the IE is 2000ms.

	//23.501 Table 5.7.4-1 when 5QI=9 AveragingWindowContent=N/A
}

type ParameterIdentifier byte

const (
	FiveQI = iota + 1
	GFBRUplink
	GFBRDownlink
	MFBRUplink
	MFBRDownlink
	AveragingWindow
	//EPSBearerIdentity  //not develop yet
)

func (p *ParameterIdentifier) StoreWithString(val string) error {
	switch val {
	case "5qi":
		*p = FiveQI
	case "GFBR uplink":
		*p = GFBRUplink
	case "GFBR downlink":
		*p = GFBRDownlink
	case "MFBR uplink":
		*p = MFBRUplink
	case "MFBR downlink":
		*p = MFBRDownlink
	case "averaging window":
		*p = AveragingWindow
	default:
		return fmt.Errorf("invalid qos flow desc param list id(%s)", val)
	}
	return nil
}

//type Contents struct {
//	Unit UnitForSession
//	Rate uint16
//}

type QI5Contents byte

const (
	QI1   QI5Contents = 1
	QI2   QI5Contents = 2
	QI3   QI5Contents = 3
	QI4   QI5Contents = 4
	QI5   QI5Contents = 5
	QI6   QI5Contents = 6
	QI7   QI5Contents = 7
	QI8   QI5Contents = 8
	QI9   QI5Contents = 9
	QI65  QI5Contents = 65
	QI66  QI5Contents = 66
	QI67  QI5Contents = 67
	QI69  QI5Contents = 69
	QI70  QI5Contents = 70
	QI75  QI5Contents = 75
	QI79  QI5Contents = 79
	QI80  QI5Contents = 80
	QI82  QI5Contents = 82
	QI83  QI5Contents = 83
	QI84  QI5Contents = 84
	QI85  QI5Contents = 85
	QIMax QI5Contents = 255
)

func (p *QI5Contents) StoreWithInt(val int) error {
	if !(val > 0 && val <= int(QIMax)) {
		return fmt.Errorf("invalid 5qi (%d)", val)
	}
	switch QI5Contents(val) {
	case QI1:
	case QI2:
	case QI3:
	case QI4:
	case QI5:
	case QI6:
	case QI7:
	case QI8:
	case QI9:
	case QI65:
	case QI66:
	case QI67:
	case QI69:
	case QI70:
	case QI75:
	case QI79:
	case QI80:
	case QI82:
	case QI83:
	case QI84:
	case QI85:
	default:
		return fmt.Errorf("invalid 5qi (%d)", val)
	}
	*p = QI5Contents(val)
	return nil
}

func (p *QoSFlowDescription) Decode(msgBuf *bytes.Reader) (err error, len int) {
	rlogger.FuncEntry(types.ModCmn, nil)
	len = 0 // length of total byte
	//octet4
	octet4, err := msgBuf.ReadByte()
	if err != nil {
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "fail to read byte")
		return fmt.Errorf("fail to read byte"), len
	}
	p.QFI, _ = utils.GetBitsValue(octet4, 1, 6)
	len += 1
	//fmt.Println("p.QFI",p.QFI)

	//octet5
	octet5, err := msgBuf.ReadByte()
	if err != nil {
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "no more bytes")
		return fmt.Errorf("fail to read byte"), len
	}
	opCode, _ := utils.GetBitsValue(octet5, 6, 8)
	p.OperationCode = OperationCodeIE(opCode >> 5)
	len += 1
	//fmt.Println("p.OperationCode",p.OperationCode)

	//octet6
	octet6, err := msgBuf.ReadByte()
	if err != nil {
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "no more bytes")
		return fmt.Errorf("fail to read byte"), len
	}
	p.NumberOfParameters, _ = utils.GetBitsValue(octet6, 1, 6)
	p.E, _ = utils.GetBitValue(octet6, 7)
	len += 1
	//fmt.Println("p.NumberOfParameters",p.NumberOfParameters)
	//fmt.Println("p.E",p.E)

	p.ParameterList.ParmsList = []ParametersIE{}
	for i := 0; i < int(p.NumberOfParameters); i++ {
		plist := ParametersIE{}
		//octet7
		octet7, err := msgBuf.ReadByte()
		if err != nil {
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "no more bytes")
			return fmt.Errorf("fail to read byte"), len
		}
		plist.ParameterID = ParameterIdentifier(octet7)
		len += 1
		//fmt.Println("plist.ParameterID",plist.ParameterID)

		switch plist.ParameterID {
		case FiveQI: //2 bytes
			//length
			msgBuf.ReadByte()
			//value
			qi5, err := msgBuf.ReadByte()
			if err != nil {
				rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "no more bytes")
				return fmt.Errorf("fail to read byte"), len
			}
			plist.QI5Content = QI5Contents(qi5)
			len += 2
			//fmt.Println("plist.QI5Content",plist.QI5Content)
		case GFBRUplink: //4 bytes
			//length
			msgBuf.ReadByte()
			//value
			plist.GFBRUplinkContent.Decode(msgBuf)
			len += 4
		case GFBRDownlink: //4 bytes
			//length
			msgBuf.ReadByte()
			//value
			plist.GFBRDownlinkContent.Decode(msgBuf)
			len += 4
		case MFBRUplink: //4 bytes
			//length
			msgBuf.ReadByte()
			//value
			plist.MFBRUplinkContent.Decode(msgBuf)
			len += 4
		case MFBRDownlink: //4 bytes
			//length
			msgBuf.ReadByte()
			//value
			plist.MFBRDownlinkContent.Decode(msgBuf)
			len += 4
		case AveragingWindow: //3 bytes
			//length
			msgBuf.ReadByte()

			//value
			valueBytes := make([]byte, 2)
			binary.Read(msgBuf, binary.BigEndian, valueBytes)
			plist.AveragingWindowContent = binary.BigEndian.Uint16(valueBytes)
			len += 3
		}
		p.ParameterList.ParmsList = append(p.ParameterList.ParmsList, plist)
	}

	return nil, len
}

//encode QoSFlowDescription to nas octet stream
func (p *QoSFlowDescription) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte
	//octet 4 QFI
	encBuf = append(encBuf, p.QFI)
	//octet 5 OperationCode
	encBuf = append(encBuf, byte(p.OperationCode)<<5)
	//octet 6: E & NumberOfParameters
	octet6 := utils.BoolToByte(p.E) << 6
	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "E(%v),v(%x)", p.E, octet6)
	octet6 |= p.NumberOfParameters

	encBuf = append(encBuf, octet6)

	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "ParameterList, number(%d), list(%v) ", p.NumberOfParameters, p.ParameterList.ParmsList)

	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "Encode QoSFlowDescription before ParameterList(%x)", encBuf)
	//ParameterList 24.501 Table 9.11.4.12.1
	for i := 0; i < len(p.ParameterList.ParmsList); i++ {
		switch p.ParameterList.ParmsList[i].ParameterID {
		case FiveQI:
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "encode for 5QI,ParameterID(%d),5QI(%d)",
				p.ParameterList.ParmsList[i].ParameterID,
				p.ParameterList.ParmsList[i].QI5Content)
			encBuf = append(encBuf, byte(p.ParameterList.ParmsList[i].ParameterID))
			encBuf = append(encBuf, byte(1))
			encBuf = append(encBuf, byte(p.ParameterList.ParmsList[i].QI5Content))
		case GFBRUplink:
			encBuf = append(encBuf, byte(p.ParameterList.ParmsList[i].ParameterID))
			valueBytes := p.ParameterList.ParmsList[i].GFBRUplinkContent.Encode()
			encBuf = append(encBuf, byte(len(valueBytes)))
			encBuf = append(encBuf, valueBytes...)
		case GFBRDownlink:
			encBuf = append(encBuf, byte(p.ParameterList.ParmsList[i].ParameterID))
			valueBytes := p.ParameterList.ParmsList[i].GFBRDownlinkContent.Encode()
			encBuf = append(encBuf, byte(len(valueBytes)))
			encBuf = append(encBuf, valueBytes...)
		case MFBRUplink:
			encBuf = append(encBuf, byte(p.ParameterList.ParmsList[i].ParameterID))
			valueBytes := p.ParameterList.ParmsList[i].MFBRUplinkContent.Encode()
			encBuf = append(encBuf, byte(len(valueBytes)))
			encBuf = append(encBuf, valueBytes...)
		case MFBRDownlink:
			encBuf = append(encBuf, byte(p.ParameterList.ParmsList[i].ParameterID))
			valueBytes := p.ParameterList.ParmsList[i].MFBRDownlinkContent.Encode()
			encBuf = append(encBuf, byte(len(valueBytes)))
			encBuf = append(encBuf, valueBytes...)
		case AveragingWindow:
			encBuf = append(encBuf, byte(p.ParameterList.ParmsList[i].ParameterID))
			encBuf = append(encBuf, byte(2))
			lengthBuf := make([]byte, 2)
			binary.BigEndian.PutUint16(lengthBuf, uint16(p.ParameterList.ParmsList[i].AveragingWindowContent))
			encBuf = append(encBuf, lengthBuf[:]...)
		}
	}
	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "Encode QoSFlowDescription (%x)", encBuf)
	return encBuf, nil
}
