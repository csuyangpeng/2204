package types3gpp

type DnnInfo struct {
	Dnn               Apn
	DefDnnInd         bool // true: the dnn is the default dnn
	LboRoamingAllowed bool // Allowed
	IwkEpsInd         bool // true subscribed with EPS
	LadnInd           bool // true the dnn is local area data network
}
