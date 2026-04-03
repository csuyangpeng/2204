package gctxt

import (
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/syncmap"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

var imsiSmfUeCtxtTable syncmap.SyncMap

func DumpSMFUeCtxtTable() {
	fmt.Println("============== DumpSMFUeCtxtTable ==============")
	fmt.Println("- imsiSmfUeCtxtTable -")
	imsiSmfUeCtxtTable.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})
}

func ValuesOfSmfUeInfoTbl(key KeyType) (CxtList []*UeContext, err error) {
	switch key {
	case ImsiType:
		imsiSmfUeCtxtTable.Range(func(key, value interface{}) bool {
			ctxt, ok := value.(*UeContext)
			if !ok {
				err = fmt.Errorf("invalid smf session info type")
				return false
			}
			CxtList = append(CxtList, ctxt)
			return true
		})
	default:
		err = fmt.Errorf("invalid key")
		return
	}
	return
}

func AddIndexUeContext(key interface{}, ctxt *UeContext) error {
	rlogger.FuncEntry(types.ModuleSmfCtxt, nil)
	var err error
	switch key.(type) {
	case ImsiKey:
		err = imsiSmfUeCtxtTable.Set(key, ctxt)
		if err != nil {
			err = fmt.Errorf("failed to set key(%d),err(%s)", key.(ImsiKey), err)
		}
	default:
		err = fmt.Errorf("invalid key type")
	}
	return err
}

func GetUeContext(key interface{}) (ueCtxt *UeContext, err error) {
	rlogger.FuncEntry(types.ModuleSmfCtxt, ueCtxt)
	switch key.(type) {
	case ImsiKey:
		val := imsiSmfUeCtxtTable.Get(key)
		if val == nil {
			err = fmt.Errorf("failed to find UeContext with Imsi key(%d)", key.(ImsiKey))
			return
		}
		ctxt, ok := val.(*UeContext)
		if !ok {
			err = fmt.Errorf("invalid ue context type")
			return
		}
		ueCtxt = ctxt
	default:
		err = fmt.Errorf("invalid key")
	}
	return
}

func UpdateUeContext(key interface{}, ueCtxt *UeContext) error {
	rlogger.FuncEntry(types.ModuleSmfCtxt, ueCtxt)
	if ueCtxt == nil {
		return fmt.Errorf("invalid input parameter, nil ueCtxt")
	}
	switch key.(type) {
	case ImsiKey:
		imsiSmfUeCtxtTable.Update(key, ueCtxt)
	default:
		return fmt.Errorf("invalid key")
	}

	return nil
}

func DeleteUeContext(key interface{}) error {
	rlogger.FuncEntry(types.ModuleSmfCtxt, nil)
	switch key.(type) {
	case ImsiKey:
		imsiSmfUeCtxtTable.Del(key)
	default:
		return fmt.Errorf("invalid key")
	}
	return nil
}

func LengthOfUeCtxtTbl(key KeyType) uint32 {
	var length uint32
	switch key {
	case ImsiType:
		length = imsiSmfUeCtxtTable.Length()
	default:
		length = 0 //invalid
	}
	return length
}

func RemoveSmfUeContext(imsi types3gpp.Imsi) error {
	rlogger.FuncEntry(types.ModuleSmfCtxt, nil)
	// get *UeContext for imsiSmfUeCtxtTable
	sessionCtxt, err := GetUeContext(ImsiKey(imsi.GetValue()))
	if err != nil {
		rlogger.Trace(types.ModuleSmfCtxt, rlogger.INFO, nil,
			"failed to find smf ueCtxt by index(imsi:%s) ,error(%s)",
			imsi, err)
		return fmt.Errorf("failed to find smf ueCtxt by index(imsi:%s) ,error(%s)", imsi, err)
	}
	rlogger.Trace(types.ModuleSmfCtxt, rlogger.INFO, nil, sessionCtxt)
	// delete the imsi key for imsiSmfUeCtxtTable
	err = DeleteUeContext(ImsiKey(imsi.GetValue()))
	if err != nil {
		rlogger.Trace(types.ModuleSmfCtxt, rlogger.INFO, nil,
			"failed to delete smf ueCtxt by index(imsi:%s) error(%s)",
			imsi, err)
		return fmt.Errorf("failed to delete smf ueCtxt by index(imsi:%s) ,error(%s)", imsi, err)
	}
	return nil
}

func RemoveSmfSessionContext(imsi types3gpp.Imsi,psi nas.PduSessID) error {
	rlogger.FuncEntry(types.ModuleSmfCtxt, nil)

	ueCtxt, err := GetUeContext(ImsiKey(imsi.GetValue()))
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas,rlogger.ERROR,nil,
			"failed to get ue context with imsi(%s)",
			imsi.String())
		return types.ErrFailFindUeCtxt
	}

	err = DeleteSessionContext(SeidKey(ueCtxt.GetPduSessCtxt(psi).SEID))
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "fail to del session ctxt")
		return types.ErrFailDelSessionCtxt
	}
	err = ueCtxt.DeletePduSessCtxt(psi)
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "fail to del session ctxt")
		return types.ErrFailDelSessionCtxt
	}
	rlogger.Trace(types.ModuleSmfCtxt, rlogger.DEBUG, nil, "delete pdu session (%d)", psi)

	//如果该ue在smf没有任何会话信息了，就删掉ue上下文
	if len(ueCtxt.PduSessCtxts) == 0 {
		// delete the imsi key for imsiSmfUeCtxtTable
		err = DeleteUeContext(ImsiKey(imsi.GetValue()))
		if err != nil {
			rlogger.Trace(types.ModuleSmfCtxt, rlogger.INFO, nil,
				"failed to delete smf ueCtxt by index(imsi:%s) error(%s)",
				imsi, err)
			return fmt.Errorf("failed to delete smf ueCtxt by index(imsi:%s) ,error(%s)", imsi, err)
		}
	}
	return nil
}

func RemoveSmfAllUeSessionContext(imsi types3gpp.Imsi) error {
	rlogger.FuncEntry(types.ModuleSmfCtxt, nil)

	ueCtxt, err := GetUeContext(ImsiKey(imsi.GetValue()))
	if err != nil {
		return fmt.Errorf("failed to get ue context with imsi(%s)", imsi.String())
	}

	for psi, v := range ueCtxt.PduSessCtxts {
		err = DeleteSessionContext(SeidKey(v.SEID))
		if err != nil {
			rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "fail to del session ctxt")
			return fmt.Errorf("fail to del session ctxt")
		}
		rlogger.Trace(types.ModuleSmfCtxt, rlogger.DEBUG, nil, "delete pdu session (%d)", psi)
	}

	// delete the imsi key for imsiSmfUeCtxtTable
	err = DeleteUeContext(ImsiKey(imsi.GetValue()))
	if err != nil {
		rlogger.Trace(types.ModuleSmfCtxt, rlogger.INFO, nil,
			"failed to delete smf ueCtxt by index(imsi:%s) error(%s)",
			imsi, err)
		return fmt.Errorf("failed to delete smf ueCtxt by index(imsi:%s) ,error(%s)", imsi, err)
	}
	return nil
}
