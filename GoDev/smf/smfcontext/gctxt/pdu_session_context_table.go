package gctxt

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/syncmap"
	"lite5gc/cmn/types"
)

var seidPduSessCtxtTable syncmap.SyncMap

func ValuesOfSmfSessionCtxtInfoTbl(key KeyType) (CxtList []*PduSessContext, err error) {
	rlogger.FuncEntry(types.ModuleSmfCtxt, nil)
	switch key {
	case SeidType:
		seidPduSessCtxtTable.Range(func(key, value interface{}) bool {
			ctxt, ok := value.(*PduSessContext)
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

func AddIndexSessContext(key interface{}, sessCtxt *PduSessContext) error {
	rlogger.FuncEntry(types.ModuleSmfCtxt, sessCtxt)
	var err error
	switch key.(type) {
	case SeidKey:
		err = seidPduSessCtxtTable.Set(key, sessCtxt)
		if err != nil {
			err = fmt.Errorf("failed to set key(%d),err(%s)", key.(SeidKey), err)
		}
	default:
		err = fmt.Errorf("invalid key type")
	}
	return err
}

func GetSessContext(key interface{}) (sessCtxt *PduSessContext, err error) {
	rlogger.FuncEntry(types.ModuleSmfCtxt, nil)

	switch key.(type) {
	case SeidKey:
		val := seidPduSessCtxtTable.Get(key)
		if val == nil {
			err = fmt.Errorf("failed to find UeContext with SmfSmCtxtId key(%d)", key.(SeidKey))
			return
		}
		ctxt, ok := val.(*PduSessContext)
		if !ok {
			err = fmt.Errorf("invalid ue context type")
			return
		}
		sessCtxt = ctxt

	default:
		err = fmt.Errorf("invalid key")
	}
	return
}

func DeleteSessionContext(key interface{}) error {
	rlogger.FuncEntry(types.ModuleSmfCtxt, nil)

	switch key.(type) {
	case SeidKey:
		seidPduSessCtxtTable.Del(key)
	default:
		return fmt.Errorf("invalid key")
	}

	return nil
}

func UpdateSessionContext(key interface{}, sessCtxt *PduSessContext) error {
	rlogger.FuncEntry(types.ModuleSmfCtxt, sessCtxt)
	if sessCtxt == nil {
		return fmt.Errorf("invalid input parameter, nil sessCtxt")
	}

	switch key.(type) {
	case SeidKey:
		seidPduSessCtxtTable.Update(key, sessCtxt)
	default:
		return fmt.Errorf("invalid key")
	}

	return nil
}
