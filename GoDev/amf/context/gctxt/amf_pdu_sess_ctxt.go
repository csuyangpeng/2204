package gctxt

import (
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/types3gpp"
)

type AmfPduSessCtxt struct {
	Psi            nas.PduSessID
	OldPsi         nas.PduSessID
	ReqType        nasie.PduSessRequestType
	SNssai         nasie.SNssai
	Dnn            types3gpp.Apn
	AdditionalInfo []byte
	Status         types3gpp.PduSessStatus
	IsPSIExist     bool //主要用于标志会话拒绝后要不要删掉信息
}

func NewAmfPduSessCtxt(id nas.PduSessID) *AmfPduSessCtxt {
	p := &AmfPduSessCtxt{}
	p.Psi = id
	return p
}
