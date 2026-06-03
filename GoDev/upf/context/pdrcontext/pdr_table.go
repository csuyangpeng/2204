package pdrcontext

import (
	"lite5gc/cmn/syncmap"
)

// Core PDR Table
var IptoPDRTable syncmap.SyncMap //key: UE IP ,value:CorePDR

// Access PDR Table
var TeidtoPDRTable syncmap.SyncMap //key: TEID ,value:AccessPDR
