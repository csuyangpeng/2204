package naslayer

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/cmn/idmgr"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func RegisterFailHandler(ctx context.Context, scInstId uint32,
	n2connData *gctxt.N2ConnCtxt, cause nas.Mm5gCause) error {
	// need to send registration reject msg to ran
	var ueCtxt *gctxt.UeContext
	//check ue context, if no . create a new one
	ueCtxt, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "Failed to find ue context."+
			" create a new one to send reg reject msg")
		// allocated  and save amfUeNgapId
		amfUeNgapId, err := idmgr.GetInst().BorrowID(string(types.AMFUeNgapId))
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to borrow amf ue ngap id, err(%s)", err)
			return fmt.Errorf("failed to borrow amf ue ngap id")
		}
		//set the sc instance id in amf id
		amfUeNgapId = (amfUeNgapId & 0x00FFFFFF) | (scInstId << 24)
		ueCtxt = &gctxt.UeContext{}
		ueCtxt.SetAmfUeNgapID(uint64(amfUeNgapId))
		ueCtxt.SetRanUeNgapId(n2connData.GnbConnID)
	}

	err := gctxt.AddIndexN2ConnContext((gctxt.AmfUeNgApId)(ueCtxt.GetAmfUeNgapId()), n2connData)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to add index to n2connData "+
			"with AmfUeNgApId %s", ueCtxt.GetAmfUeNgapId())
		return fmt.Errorf("failed to add index to n2connData of AmfUeNgApId %s",
			ueCtxt.GetAmfUeNgapId())

	}
	ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)
	//send reject msg
	err = mmsender.SendRegisterRejectMsg(ctx, cause)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to send register reject")
		return fmt.Errorf("failed to send register reject")
	}
	return nil
}
