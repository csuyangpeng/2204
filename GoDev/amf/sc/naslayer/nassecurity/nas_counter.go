package nassecurity

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func IncreaseNasCounter(nasCounter uint32) (uint32, bool) {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "input nas counter (%d)", nasCounter)

	overflow := false
	counter := nasCounter + 1

	if counter > 0xffffff {
		// wrap around 24.501 4.4.3.5
		counter = 0
		overflow = true
	}

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "output nas counter (%d) over flow(%v)", counter, overflow)

	return counter, overflow
}

func GetUplinkNasCounter(nasCounter uint32) byte {
	return byte(nasCounter & 0x000000ff)
}

//24.501 4.4.3.1
func StoreUplinkNasCounter(nasSeqNum uint8, nasCounter *uint32) bool {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "input nas seq num(%d), nas counter (%d)", nasSeqNum, *nasCounter)

	NasCount := *nasCounter

	overflowPart := NasCount >> 8
	nasSeqPart := uint8(NasCount & 0xff)
	//if nasSeqNum < nasSeqPart {
	if nasSeqPart == 0xff {
		// overflow happended
		overflowPart += 1
	}
	nasSeqPart = nasSeqNum

	NasCount = overflowPart<<8 | uint32(nasSeqPart)

	if NasCount > 0xffffff {
		*nasCounter = 0
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "input nas seq num(%d),output nas counter(%x)", nasSeqNum, *nasCounter)
		return true
	} else {
		*nasCounter = NasCount & 0xffffff
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "input nas seq num(%d),output nas counter(%x)", nasSeqNum, *nasCounter)
		return false
	}
}
