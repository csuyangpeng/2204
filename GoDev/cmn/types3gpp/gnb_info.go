package types3gpp

import (
	"fmt"
)

type GnbIdType uint8

const (
	GlobalGNB   GnbIdType = 1
	GlobalNgENB GnbIdType = 2
)

type GnbIdInfo struct {
	IdType GnbIdType
	Plmn   PlmnID
	GnbId  uint32
}

func (p GnbIdInfo) String() string {
	return fmt.Sprintf("plmn-%s-type-%d-gnbid-%d", p.Plmn, p.IdType, p.GnbId)
}

type GnbInfo struct {
	GnbInstId uint32
	GnbId     GnbIdInfo
	GnbIP     string
}

const (
	InvalidAmfNgapId uint64 = 0xFFFFFFFFFFFFFFFF
	InvalidInstId    uint32 = 0xFFFFFFFF
)

type Gnb2AmfScMsg struct {
	GnbInfo  GnbInfo
	MsgType  MsgType
	PrcdCode ProcedureCode
	NgapMsg  []byte
}

func (p *Gnb2AmfScMsg) String() string {
	return fmt.Sprintf("msg type(%s),prcd code(%s),ngap msg(%x)", p.MsgType, p.PrcdCode, p.NgapMsg)
}

// gnb information stored on AMF
type GnbInformation struct {
	GnbInstId    uint32    `json:"gnb instance id"`
	GnbId        uint32    `json:"gnb id"`
	GnbPlmn      string    `json:"gnb plmn"`
	GnbIp        string    `json:"gnb ip address"`
	RanNodeName  string    `json:"gnb name"`
	DefPagingDrx PagingDRX `json:"paging drx"`
	SprtTaList   []SprtTa  `json:"supported ta list"`
}
type Bplmn struct {
	Plmn      string
	SliceList []Snssai
}

func (p Bplmn) String() string {
	var str string
	for _, v := range p.SliceList {
		str += v.String()
	}

	return fmt.Sprintf("%s[%s]", p.Plmn, str)
}

type SprtTa struct {
	Tac    TAC
	BPlmns []Bplmn `json:"broadcast plmn"`
}

func (p SprtTa) String() string {
	var str string
	str = fmt.Sprintf("%s", p.Tac)
	str += ",bplmns{"
	for _, v := range p.BPlmns {
		str += fmt.Sprintf("%s,", v)
	}
	str += "}"
	return str
}

func (p GnbInformation) String() string {
	var str string

	str = fmt.Sprintf("gnb_inst_id(%d),gnb_id(%d),gnb_plmn(%s),gnb_ip(%s),gnb_name(%s),paging_drx(%d),",
		p.GnbInstId, p.GnbId, p.GnbPlmn, p.GnbIp, p.RanNodeName, p.DefPagingDrx)
	str += "supported ta list{"
	for i, v := range p.SprtTaList {
		str += fmt.Sprintf("%d-{%s}", i, v)
	}
	str += "}"
	return str
}
