/*
Define types, constant, variable, enmu etc used in the application.
*/
package types

// logger default values defs
const (
	DefConfFileSys     string = "config/cm_sys_conf.yaml"
	DefConfFileAmf     string = "config/cm_amf_conf.yaml"
	DefConfFileSmf     string = "config/cm_smf_conf.yaml"
	DefConfFileUpf     string = "config/cm_upf_conf.yaml"
	DefConfFileUdm     string = "config/cm_udm_conf.yaml"
	DefConfFileSidfKey string = "config/pki.key"
)

type ModuleName string

const (
	SC    ModuleName = "sc"    //sc instance name
	NGAP  ModuleName = "ngap"  //ngap instance name
	DPE   ModuleName = "dpe"   //dpe instance name
	SmfSc ModuleName = "smfsc" //smf sc
)

type IdKey string

const (
	AMFUeNgapId IdKey = "amfUeNgapId" //amf ue ngap id
	TMSI        IdKey = "tmsi"        //amf tmsi
	SmfTEID     IdKey = "teid"        //smf teid
	SmfSEID     IdKey = "seid"        //smf seid
	SmCtxtId    IdKey = "smCtxtId"    //sm context id
	PDRID       IdKey = "pdrId"       //pdr id
)

// max goroutine numnber definition defs
const (
	MaxNgapInst     = 10000
	MaxScInst       = 40
	MaxDpeInst      = 3
	MaxNumAmfNgapId = 0x00FFFFFF
	MaxNumTmsi      = 0x0FFFFFFF
	MaxSmfScInst    = 40
	MaxNumSmfTeid   = 0xFFFFFFFF
	MaxNumSmfSeid   = 0x00FFFFFFFFFFFFFF
	MaxNumPdrid     = 10000
)

// SM Context ID struct
// | 8bit | 24 bits  |
// |sc id|sm ctxt id|
const (
	SmcScidAnd   = 0xFF000000
	SmcScidShift = 24
)

// TMSI
// | 4 bit | 28 bit|
// | restart coutner | tmsi id |
const (
	AmfTmsiFilter = 0x0FFFFFFF
)

// max resource limitation for System
const (
	BufSize8192     = 8192
	ChanBufSize1024 = 1024
)

type ByteOrder uint8

const (
	LittleEndian ByteOrder = iota
	BigEndian
)
