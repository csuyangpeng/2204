package types3gpp

type RestrictionType uint8

const (
	AllowedAreas    RestrictionType = 1
	NotAllowedAreas RestrictionType = 2
)

type SerAreaRstrc struct {
	//string
	//"ALLOWED_AREAS" or "NOT_ALLOWED_AREAS"
	//shall be present if and only if the areas attribute is present
	IsAreaRstrcPrst bool
	RType           RestrictionType
	Areas           []Area
	//0 means not present
	MaxNumofTAs uint
}
