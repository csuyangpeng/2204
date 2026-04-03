package types3gpp

import "fmt"

type Guami struct {
	PlmnId PlmnID
	AmfId  AmfIdentifier
}

func (p Guami) String() string {
	return fmt.Sprintf("PlmnId(%s),AmfId(%s)", &(p.PlmnId), &(p.AmfId))
}

// ServedGuami struct definition
//type ServedGuami struct {
//	PlmnID PlmnID
//	AmfID  AmfIdentifier
//}
//
//func (p *ServedGuami) String() string {
//	return fmt.Sprintf("PlmnId: %s, AmfIdentifier: %s",
//		&p.PlmnID, &p.AmfID)
//}
