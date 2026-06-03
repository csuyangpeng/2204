package nasie

// 24.501 9.11.3.5
//	Table 9.11.3.5.1: 5GS network feature support information element
//	IMS voice over PS session over 3GPP access indicator (IMS-VoPS-3GPP) (octet 3, bit 1)
//	This bit indicates the support of IMS voice over PS session over 3GPP access
//	(see NOTE 1)
//	Bit
//	1
//	0				IMS voice over PS session not supported over 3GPP access
//	1				IMS voice over PS session supported over 3GPP access
//
//	IMS voice over PS session over non-3GPP access indicator (IMS-VoPS-N3GPP) (octet 3, bit 2)
//	This bit indicates the support of IMS voice over PS session over non-3GPP access
//	Bit
//	2
//	0				IMS voice over PS session not supported over non-3GPP access
//	1				IMS voice over PS session supported over non-3GPP access
//
//	Emergency service support indicator for 3GPP access (EMC) (octet 3, bit 3 and bit 4)
//	This bit indicates the support of emergency services in 5GS for 3GPP access (see NOTE 2)
//	Bits
//	4	3
//	0	0			Emergency services not supported
//	0	1			Emergency services supported in NR connected to 5GCN only
//	1	0			Emergency services supported in E-UTRA connected to 5GCN only
//	1	1			Emergency services supported in NR connected to 5GCN and E-UTRA connected to 5GCN
//
//	Emergency service fallback indicator for 3GPP access (EMF) (octet 3, bit 5 and bit 6)
//	This bit indicates the support of emergency services fallback for 3GPP access (see NOTE 2)
//	Bits
//	6	5
//	0	0			Emergency services fallback not supported
//	0	1			Emergency services fallback supported in NR connected to 5GCN only
//	1	0			Emergency services fallback supported in E-UTRA connected to 5GCN only
//	1	1			Emergency services fallback supported in NR connected to 5GCN and E-UTRA connected to 5GCN
//
//	Interworking without N26 interface indicator (IWK N26) (octet 3, bit 7)
//	This bit indicates whether interworking without N26 interface is supported
//	Bit
//	7
//	0				Interworking without N26 interface not supported
//	1				Interworking without N26 interface supported
//
//	MPS indicator (MPSI) (octet 3, bit 8)
//	This bit indicates the support of MPS in the RPLMN or equivalent PLMN.
//	Bit
//	8
//	0				Access identity 1 not valid in RPLMN or equivalent PLMN
//	1				Access identity 1 valid in RPLMN or equivalent PLMN
//
//	Emergency service support for non-3GPP access indicator (EMCN3) (octet 4, bit 1)
//	This bit indicates the support of emergency services in 5GS for non-3GPP access
//	Bit (see NOTE 3)
//	1
//	0				Emergency services not supported over non-3GPP access
//	1				Emergency services supported over non-3GPP access
//
//	MCS indicator (MCSI) (octet 4, bit 2)
//	This bit indicates the support of MCS in the RPLMN or equivalent PLMN.
//	Bit
//	2
//	0				Access identity 2 not valid in RPLMN or equivalent PLMN
//	1				Access identity 2 valid in RPLMN or equivalent PLMN
//
//	Bits 3 to 8 in octets 4 and all bits in octet 5 are spare and shall be coded as zero,
// if the respective octet is included in the information element.
//
//	NOTE 1:	For a registration procedure over non-3GPP access, bit 1 of octet 3 is ignored.
//	NOTE 2:	For a registration procedure over 3GPP access, bit 1 of octet 4 is ignored.
//	NOTE 3:	For a registration procedure over non-3GPP access, bits 3 to 6 of octet 3 are ignored.
type Emc byte

const (
	EmcSerNotSupport   Emc = 0
	EmcSerSupNROnly    Emc = 1
	EmcSerSupEutraOnly Emc = 2
	EmcSerSupEutraNR   Emc = 3
)

type Emf byte

const (
	EmfNotSupport   Emf = 0
	EmfSupNROnly    Emf = 1
	EmfSupEutraOnly Emf = 2
	EmfSupEutraNR   Emf = 3
)

type NetworkFeatureSupport struct {
	ImsVoPs3gpp    bool
	ImsVoPsNon3gpp bool
	Emc            Emc
	Emf            Emf
	IwkN26         bool // false - inter-working without N26 not supported
	Mpsi           bool // false - access identity 1 not valid in RPLMN or equivalent PLMNc
	Emcw           bool // false - Emergency services not supported over non-3GPP access
}
