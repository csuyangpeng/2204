package udmdata

import (
	"lite5gc/cmn/nas/nasie"
	t3 "lite5gc/cmn/types3gpp"
)

type Nssai struct {
	SupFeatures uint64
	DefSnssai   nasie.SNssai
	SnssaiList  []nasie.SNssai
}

//29503 Table 6.1.6.2.4-1
type AccMobSubscribeData struct {
	SupFeatures        string
	MsIsdn             []t3.Msisdn
	InternalGrpIds     []string
	SubsUeAmbr         t3.Ambr
	Nssai              Nssai
	RatRestrictions    []t3.RatType
	ForbiddenAreas     []t3.Area
	SerAreaRestictions t3.SerAreaRstrc
	CnTypeRestrictions []t3.CoreNetworkType
	RfspIndex          uint16 //[1,256]
	SubsRegTimer       uint   // second
	UeUsageType        int
	MpsPriority        bool
	McsPriority        bool
	ActivateTime       uint //second
	DlPacketCount      int  //minimum: -1
	SorInfo            t3.SorInfo
	MicoAllowed        bool //Indicates whether MICO mode is allowed for the UE.
	// The following values are defined:
	//0: "Extended DL Data Buffering NOT REQUESTED"
	//-1: "Extended DL Data Buffering REQUESTED, without a suggested number of packets"
	//n>0: "Extended DL Data Buffering REQUESTED, with a suggested number of n packets"
	SharedAmDataIds []string
	DlPktCount      int
}
