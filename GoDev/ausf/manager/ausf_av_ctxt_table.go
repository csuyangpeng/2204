package manager

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/syncmap"
	"lite5gc/cmn/types"
)

var ausfAvContextTable syncmap.SyncMap

func CreateAusfAvContext(imsi string) (error, *types.AusfAvContext) {
	rlogger.FuncEntry(types.ModuleAusf, &imsi)

	var err error

	avCtxt := &types.AusfAvContext{}
	avCtxt.Imsi = imsi

	err = ausfAvContextTable.Set(imsi, avCtxt)
	if err != nil {
		return fmt.Errorf("failed to create ausf av context in table, error(%s)", err), nil
	}

	return nil, avCtxt
}

func DeleteAusfAvContext(imsi string) {
	rlogger.FuncEntry(types.ModuleAusf, &imsi)

	ausfAvContextTable.Del(imsi)
}

func GetAusfAvContext(imsi string) (avCtxt *types.AusfAvContext, err error) {
	rlogger.FuncEntry(types.ModuleAusf, &imsi)

	val := ausfAvContextTable.Get(imsi)
	if val == nil {
		err = fmt.Errorf("failed to find AusfAvContext with Imsi (%s)", imsi)
		return nil, err
	}

	avCtxt, ok := val.(*types.AusfAvContext)
	if !ok {
		err = fmt.Errorf("invalid AusfAvontext type")
		return nil, err
	}

	return avCtxt, nil
}

func UpdateAusfAvContext(imsi string, avCtxt *types.AusfAvContext) error {
	rlogger.FuncEntry(types.ModuleAusf, &imsi)
	if avCtxt == nil {
		return fmt.Errorf("invalid input parameter, nil AusfAvCtxt")
	}

	ausfAvContextTable.Update(imsi, avCtxt)

	return nil
}

func LengthOfAusfAvCtxtTbl(imsi string) uint32 {
	return ausfAvContextTable.Length()
}
