package confsc

import (
	"fmt"
	"lite5gc/cmn/observer"
)

const (
	CSCLogConsoleUpdate observer.MsgType = "CSCLogConsoleUpdate"
	CSCLogFileUpdate    observer.MsgType = "CSCLogFileUpdate"
	CSCLogLemsUpdate    observer.MsgType = "CSCLogLemsUpdate"
	CSCLogLevelUpdate   observer.MsgType = "CSCLogLevelUpdate"
	CSCLogPathUpdate    observer.MsgType = "CSCLogPathUpdate"
	CSCLogNfNoUpdate    observer.MsgType = "CSCLogNfNoUpdate"

	CSCConfLogConsoleUpdate         observer.MsgType = "CSCConfLogConsoleUpdate"
	CSCConfLogFileUpdate            observer.MsgType = "CSCConfLogFileUpdate"
	CSCConfLogLemsUpdate            observer.MsgType = "CSCConfLogLemsUpdate"
	CSCConfLogLevelUpdate           observer.MsgType = "CSCConfLogLevelUpdate"
	CSCConfLogPathUpdate            observer.MsgType = "CSCConfLogPathUpdate"
	CSCConfModuleCountSwitchUpdate  observer.MsgType = "CSCConfModuleCountSwitchUpdate"
	CSCConfSessionCountSwitchUpdate observer.MsgType = "CSCConfSessionCountSwitchUpdate"
	CSCConfVPPUPFN3IPAddressUpdate  observer.MsgType = "CSCConfVPPUPFN3IPAddressUpdate"
	CSCConfVPPUPFN4IPAddressUpdate  observer.MsgType = "CSCConfVPPUPFN4IPAddressUpdate"
	CSCConfVPPUPFDnnIPAddrUpdate    observer.MsgType = "CSCConfVPPUPFDnnIPAddrUpdate"
	CSCConfVPPIfStatsEnableUpdate   observer.MsgType = "CSCConfVPPIfStatsEnableUpdate"

	CSCDbmgrDbDriverNameUpdate observer.MsgType = "CSCDbmgrDbDriverNameUpdate"
	CSCDbmgrDbSourcePathUpdate observer.MsgType = "CSCDbmgrDbSourcePathUpdate"

	CSCConfDbDriverNameUpdate observer.MsgType = "CSCConfDbDriverNameUpdate"
	CSCConfDbSourcePathUpdate observer.MsgType = "CSCConfDbSourcePathUpdate"

	CSCPMModuleCountSwitchUpdate  observer.MsgType = "CSCPMModuleCountSwitchUpdate"
	CSCPMSessionCountSwitchUpdate observer.MsgType = "CSCPMSessionCountSwitchUpdate"

	CSCVPPUPFN3IPAddressUpdate observer.MsgType = "CSCVPPUPFN3IPAddressUpdate"
	CSCVPPUPFN3PortUpdate      observer.MsgType = "CSCVPPUPFN3PortUpdate"
	CSCVPPUPFN4IPAddressUpdate observer.MsgType = "CSCVPPUPFN4IPAddressUpdate"
	CSCVPPUPFN4PortUpdate      observer.MsgType = "CSCVPPUPFN4PortUpdate"
	CSCVPPUPFDnnIPAddrUpdate   observer.MsgType = "CSCVPPUPFDnnIPAddrUpdate"
	CSCVPPIfStatsEnableUpdate  observer.MsgType = "CSCVPPIfStatsEnableUpdate"
)

const (
	ESCVPPUPFStatsUpdate  observer.MsgType = "ESCVPPUPFStatsUpdate"
	ESCVPPConnEventUpdate observer.MsgType = "ESCVPPConnEventUpdate"
)

type CSCCliLogConsoleUpdateParam struct {
	Swih  bool
	Level string
}

type CSCCliLogFileUpdateParam struct {
	Swih  bool
	Level string
	Path  string
}

type CSCCliLogLemsUpdateParam struct {
	Swih  bool
	Level string
}

type CSCLemsBaiUPFIPAddressUpdateParam struct {
	NetIntf int
	IfName  string
	IPAddr  string
}

type ESCVPPUPFStatsUpdateParam struct {
	IfRxPackets uint64
	IfRxBytes   uint64
	IfTxPackets uint64
	IfTxBytes   uint64

	UpfUlRxPackets   uint64
	UpfUlRxBytes     uint64
	UpfUlTxPackets   uint64
	UpfUlTxBytes     uint64
	UpfUlDropPackets uint64
	UpfUlDropBytes   uint64
	UpfDlRxPackets   uint64
	UpfDlRxBytes     uint64
	UpfDlTxPackets   uint64
	UpfDlTxBytes     uint64
	UpfDlDropPackets uint64
	UpfDlDropBytes   uint64
}

type ESCVPPConnEventUpdateParam int

const (
	VPPConnEventConnected ESCVPPConnEventUpdateParam = iota
	VPPConnEventDisconnected
	VPPConnEventFailed
)

func (s ESCVPPConnEventUpdateParam) String() string {
	switch s {
	case VPPConnEventConnected:
		return "VPPConnEventConnected"
	case VPPConnEventDisconnected:
		return "VPPConnEventDisconnected"
	case VPPConnEventFailed:
		return "VPPConnEventFailed"
	default:
		return fmt.Sprintf("UnknownState(%d)", s)
	}
}
