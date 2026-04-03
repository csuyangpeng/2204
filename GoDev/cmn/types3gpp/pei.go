package types3gpp

type PeiType uint8

const (
	IMEI   PeiType = 0
	IMEISV PeiType = 1

	LenImei   uint8 = 15
	LenImeiSv uint8 = 16
)

// 23.003 3.2.1
// IMEI    15 digits    TAC  +  SNR  +  CD/SD
// IMEISV  16 digits    TAC  +  SNR  +  SVN

type Pei struct {
	PeiType PeiType
	Imei    [LenImei]byte
	ImeiSv  [LenImeiSv]byte
	//tac     [8]uint8
	//snr     [6]uint8
	//svn     [2]uint8
	//cdsd    uint8
}
