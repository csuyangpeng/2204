//Package types3gpp define types, constant, varables in 3GPP specifications
package types3gpp

// Constant definition
const (
	IMSI_BCD_LEN = 8
	INV_IMSI64   = 0
)

// UeNgApID define type for ue ng ap id
type UeNgApID uint64

// Teid define type for teid
type Teid uint32

// string msisnd formate
type Msisdn string

// RAT Type enum
type RatType byte

const (
	RatType_NR      RatType = 1
	RatType_EUTRA   RatType = 2
	RatType_WLAN    RatType = 3
	RatType_VIRTUAL RatType = 4
	RatType_INV     RatType = 0
)

func (p *RatType) Store(v string) {
	switch v {
	case "EUTRA":
		*p = RatType_EUTRA
	case "NR":
		*p = RatType_NR
	case "WLAN":
		*p = RatType_WLAN
	case "VIRTUAL":
		*p = RatType_VIRTUAL
	default:
		*p = RatType_INV

	}
}

// CoreNetworkType
type CoreNetworkType byte

const (
	FiveGC CoreNetworkType = 1
	EPC    CoreNetworkType = 2
)
