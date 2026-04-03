package configure

import (
	"fmt"
	"lite5gc/cmn/types3gpp"
)

// amf service configuration
type AmfService struct {
	AmfInstanceId string
	AmfName       string
	AmfRelCap     int
	AmfIdentifier types3gpp.AmfIdentifier
}

func (p AmfService) String() (strbuf string) {
	strbuf += fmt.Sprintf("AmfService Info:\n")
	strbuf += fmt.Sprintln("AmfInstanceId: ", p.AmfInstanceId)
	strbuf += fmt.Sprintln("AmfName: ", p.AmfName)
	strbuf += fmt.Sprintln("AmfIdentifier: ", p.AmfIdentifier)
	strbuf += fmt.Sprintln("AmfRelCap: ", p.AmfRelCap)
	return strbuf
}
