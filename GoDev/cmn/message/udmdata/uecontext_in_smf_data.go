package udmdata

import "lite5gc/cmn/types3gpp"

type UeContextInSmfData struct {
	PduSessions map[int]*types3gpp.PduSession
}
