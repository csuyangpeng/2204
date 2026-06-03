// Package prcdctxt is the procedure context
// which store the information during the procedure
package prcdctxt

import T "lite5gc/cmn/types3gpp"

// Base is interface for procedure context
type Base interface {
	GetCurrentState() string
}

//BaseCtxt is basic context for all procedure context
type BaseCtxt struct {
	imsi T.Imsi
	//current procedure status
	status string
}

//GetCurrentState return the current state for a procedure
func (p BaseCtxt) GetCurrentState() string {
	return p.status
}
