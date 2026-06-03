package types3gpp

//29503 Table 6.1.6.2.26
type SorInfo struct {
	// ackInd	AckInd	M	1
	// Contains the indication whether the acknowledgement from UE is needed.
	AckInd bool

	// sorMacIausf	SorMac	M	1	Contains the SoR-MAC-IAUSF.
	SorMacIausf string

	// countersor	CounterSor	M	1	Contains the CounterSoR.
	Countersor string

	// steeringContainer	SteeringContainer	C	1
	// When present, this information contains the information needed to
	// update the "Operator Controlled PLMN Selector with Access Technology" list stored in the USIM
	// either as an array of preferred PLMN/AccessTechnologies combinations in priority order.
	// The first entry in the array indicates the highest priority and the last entry indicates the lowest.
	// Or it may contain a secured packet.
	// If no change of the "Operator Controlled PLMN Selector with
	// Access Technology" list stored in the USIM is needed, then this attribute shall be absent.
	SteeringContainer SteeringContainer
}

type SteeringContainer struct {
	//array(SteeringInfo)	1..N	List of PLMN/AccessTechnologies combinations.
	SteeringInfoList []SteeringInfo

	//SecuredPacket	1
	// A Secured packet containing one or more APDUs commands dedicated to Remote File Management or
	// command responses (see ETSI TS 102.225 [28]).
	SecuredPacket string
}

// 29509 Table 6..6.2.4-1
type SteeringInfo struct {
	//plmnId	PlmnId	M	1	Contains a preferred PLMN identity.
	PlmnId PlmnID

	//accessTechList	Array(AccessTech)	C	1..N
	//When present it contains the 28 referred access technologies as listed in subclause 4.2.5 of 3GPP TS 31.102 [15].
	// If absent it means that all access technologies are equivalently preferred in this PLMN.
	AccessTechList []AccessTech
}

type AccessTech byte

const (
	NR AccessTech = iota
	EUTRAN_IN_WBS1_MODE_AND_NBS1_MODE
	EUTRAN_IN_NBS1_MODE_ONLY
	EUTRAN_IN_WBS1_MODE_ONLY
	UTRAN
	GSM_AND_ECGSM_IoT
	GSM_WITHOUT_ECGSM_IoT
	ECGSM_IoT_ONLY
	CDMA_1xRTT
	CDMA_HRPD
	GSM_COMPACT
)
