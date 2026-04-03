package manager

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/syncmap"
	"lite5gc/cmn/types"
)

var ueAuthDataTable syncmap.SyncMap

func CreateUeSecContext(imsi string) (error, *types.UeSecContext) {
	rlogger.FuncEntry(types.ModuleUdm, nil)

	var err error

	ueSecCtxt := &types.UeSecContext{}
	ueSecCtxt.Imsi = imsi

	err = ueAuthDataTable.Set(imsi, ueSecCtxt)
	if err != nil {
		return fmt.Errorf("failed to create ue auth data in table, error(%s)", err), nil
	}

	return nil, ueSecCtxt
}

func DeleteUeSecContext(imsi string) {
	rlogger.FuncEntry(types.ModuleUdm, nil)
	ueAuthDataTable.Del(imsi)
}

func GetUeSecContext(imsi string) (ueSecCtxt *types.UeSecContext, err error) {
	rlogger.FuncEntry(types.ModuleUdm,nil)

	val := ueAuthDataTable.Get(imsi)
	if val == nil {
		err = fmt.Errorf("failed to find UeSecContext with Imsi (%s)", imsi)
		return nil, err
	}

	ueSecCtxt, ok := val.(*types.UeSecContext)
	if !ok {
		err = fmt.Errorf("invalid ue security context type")
		return nil, err
	}

	return ueSecCtxt, nil
}

func UpdateUeSecContext(imsi string, ueSecCtxt *types.UeSecContext) error {
	rlogger.FuncEntry(types.ModuleUdm,nil)
	if ueSecCtxt == nil {
		return fmt.Errorf("invalid input parameter, nil ueSecCtxt")
	}

	ueAuthDataTable.Update(imsi, ueSecCtxt)

	return nil
}

func LengthOfUeSecCtxtTbl(imsi string) uint32 {
	return ueAuthDataTable.Length()
}
