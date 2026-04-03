/** Copyright(C),2020-2022
* Author: zmj
* Date: 12/17/20 1:20 PM
* Description:
 */
package gctxt

import (
	"fmt"
	"lite5gc/cmn/types3gpp"
	"testing"
)

func TestUpdateUeContext(t *testing.T) {

	ueContext := NewUeContextByAmfUeNgapID(types3gpp.UeNgApID(100))
	// save the amf_ue_ngap_id index with ue context
	err := AddIndexUeContext(AmfUeNgApId(ueContext.GetAmfUeNgapId()), ueContext)
	if err != nil {
		fmt.Println("failed to add index ue context table")
		return
	}
	DumpAmfUeIdUeCtxtTable()

	new_ueContext := NewUeContextByAmfUeNgapID(types3gpp.UeNgApID(200))
	UpdateUeContext(AmfUeNgApId(100), new_ueContext)
	DumpAmfUeIdUeCtxtTable()

	return
}

