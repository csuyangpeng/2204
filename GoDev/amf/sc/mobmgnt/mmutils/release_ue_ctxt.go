package mmutils

import (
	"lite5gc/amf/context/gctxt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/udm/arpf/manager"
)

func ReleaseUECtxt(ueCtxt *gctxt.UeContext) error {
	rlogger.FuncEntry(types.ModuleAmfMM, nil)

	//set the procedure ctxt to nil
	ueCtxt.SetProcCtxt(nil)

	//release imsi info
	imsi := ueCtxt.GetImsi()

	manager.DeleteUeSecContext(imsi.String())

	err := gctxt.DeleteUeContext(gctxt.ImsiKey(imsi.GetValue()), ueCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"nothing to release or failed to release ueCtxt by index(imsi:%s), error(%s)", imsi.String(), err)
		return types.ErrFailDelUeCtxt
	}

	err = gctxt.DeleteUeContext(gctxt.AmfUeNgApId(ueCtxt.GetAmfUeNgapId()), ueCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to release ueCtxt by index(AmfUeNgApId:%d) ,error(%s)", ueCtxt.GetAmfUeNgapId(), err)
		return types.ErrFailDelUeCtxt
	}

	if ueCtxt.Guti5g != nil {
		err = gctxt.DeleteUeContext(gctxt.GutiKey(ueCtxt.Guti5g.String()), ueCtxt)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to release ueCtxt by index(Guti5g:%d), error(%s)", ueCtxt.Guti5g.String(), err)
			return types.ErrFailDelUeCtxt
		}

		err = gctxt.DeleteUeContext(gctxt.StmsiKey(ueCtxt.Guti5g.GetStmsiKey()), ueCtxt)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to release ueCtxt by index(Stmsi:%d), error(%s)", ueCtxt.Guti5g.GetStmsiKey(), err)
			return types.ErrFailDelUeCtxt
		}
	}

	return nil
}
