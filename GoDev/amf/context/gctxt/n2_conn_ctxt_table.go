package gctxt

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/syncmap"
	"lite5gc/cmn/types"
)

var amfNgApIdN2ConnCtxtTable syncmap.SyncMap

func DumpN2ConnCtxtTable() (buf string) {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)
	buf += fmt.Sprintln("============== amfNgApIdN2ConnCtxtTable ==============")
	amfNgApIdN2ConnCtxtTable.Range(func(key, value interface{}) bool {
		buf += fmt.Sprintln("amfNgApId:", key, " N2ConnCtxt:", value)
		return true
	})
	return buf
}

func ValuesOfN2ConnCtxtTbl(key KeyType) (CxtList []*N2ConnCtxt, err error) {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)
	switch key {
	case AmfUeNgApIdType:
		amfNgApIdN2ConnCtxtTable.Range(func(key, value interface{}) bool {
			ctxt, ok := value.(*N2ConnCtxt)
			if !ok {
				err = fmt.Errorf("invalid n2 connection context type")
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

func CreateN2ConnCtxt(key interface{}) (*N2ConnCtxt, error) {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)
	switch key.(type) {
	case AmfUeNgApId:
		n2Conn := &N2ConnCtxt{}
		err := amfNgApIdN2ConnCtxtTable.Set(key, n2Conn)
		if err != nil {
			err = fmt.Errorf("failed to create "+
				"AmfNgApIdN2ConnCtxtTable, error(%s)", err)
		}
		rlogger.Trace(types.ModuleAmfNgap, rlogger.INFO, nil, "add amf-ngap-id(%d) index in n2 context",key)
		return n2Conn, err
	default:
		err := fmt.Errorf("invalid key")
		return nil, err
	}
}

func GetN2ConnContext(key interface{}) (connCtxt *N2ConnCtxt, err error) {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)
	switch keyVal := key.(type) {
	case AmfUeNgApId:
		val := amfNgApIdN2ConnCtxtTable.Get(key)
		if val == nil {
			err = fmt.Errorf("failed to find N2ConnContext "+
				"with AmfUeNgApId (%d)", keyVal)
			return
		}
		ctxt, ok := val.(*N2ConnCtxt)
		if !ok {
			err = fmt.Errorf("invalid ue context type")
			return
		}
		connCtxt = ctxt
	default:
		err = fmt.Errorf("invalid key")
	}
	return
}

func DeleteN2ConnContext(key interface{}) error {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)
	switch key.(type) {
	case AmfUeNgApId:
		amfNgApIdN2ConnCtxtTable.Del(key)
	default:
		return fmt.Errorf("invalid key")
	}
	return nil
}

func AddIndexN2ConnContext(key interface{}, ctxt *N2ConnCtxt) error {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)
	var err error
	switch key.(type) {
	case AmfUeNgApId:
		err = amfNgApIdN2ConnCtxtTable.Set(key, ctxt)
		if err != nil {
			err = fmt.Errorf("failed to set key(%d),err(%s)", key.(AmfUeNgApId), err)
		}
	default:
		err = fmt.Errorf("invalid key type")
	}
	return err
}

func GetGnbInstIdByAmfUeNgapId(id uint64) (gnbInstId uint32, e error) {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)
	n2Conn, err := GetN2ConnContext(AmfUeNgApId(id))
	if err != nil {
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.ERROR, nil, "failed to get n2 connection data")
		return 0, err
	}
	return n2Conn.GnbInfo.GnbInstId, nil
}
