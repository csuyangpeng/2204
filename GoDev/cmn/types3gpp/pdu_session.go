package types3gpp

import "fmt"

type PduSession struct {
	PduSessId int
	Dnn       Apn
	SmfInstId string
	PlmnId    PlmnID
}

func (p *PduSession) String() string {
	return fmt.Sprintf("psi(%d),dnn(%s),SmfInstId(%s),PlmnId(%s)",
		p.PduSessId, p.Dnn.String(), p.SmfInstId, p.PlmnId)
}

type PduSessStatus byte

const (
	SessActived    PduSessStatus = 1
	SessDeactive   PduSessStatus = 2
	SessActivating PduSessStatus = 3
)

func (p PduSessStatus) String() string {
	switch p {
	case SessDeactive:
		return "deactive"
	case SessActivating:
		return "activating"
	case SessActived:
		return "actived"
	default:
		return "invalid"
	}
}
