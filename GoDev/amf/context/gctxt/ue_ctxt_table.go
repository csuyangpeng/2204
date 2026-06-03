package gctxt

import (
	"fmt"
	"gopkg.in/ffmt.v1"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/syncmap"
	"lite5gc/cmn/types"
	"strings"
)

var imsiUeCtxtTable syncmap.SyncMap
var amfUeIdUeCtxtTable syncmap.SyncMap
var gutiUeCtxtTable syncmap.SyncMap
var stmsiUeCtxtTable syncmap.SyncMap

func AddIndexUeContext(key interface{}, ctxt *UeContext) error {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)
	var err error
	switch key.(type) {
	case ImsiKey:
		err = imsiUeCtxtTable.Set(key, ctxt)
		if err != nil {
			err = fmt.Errorf("failed to set key(%d),err(%s)", key.(ImsiKey), err)
		}
	case AmfUeNgApId:
		err = amfUeIdUeCtxtTable.Set(key, ctxt)
		if err != nil {
			err = fmt.Errorf("failed to set key(%d),err(%s)", key.(AmfUeNgApId), err)
		}
	case GutiKey:
		err = gutiUeCtxtTable.Set(key, ctxt)
		if err != nil {
			err = fmt.Errorf("failed to set key(%s),err(%s)", key.(GutiKey), err)
		}
	case StmsiKey:
		err = stmsiUeCtxtTable.Set(key, ctxt)
		if err != nil {
			err = fmt.Errorf("failed to set key(%s),err(%s)", key.(StmsiKey), err)
		}
	default:
		err = fmt.Errorf("invalid key type")
	}
	return err
}

func DeleteUeContext(key interface{}, ueCtxt *UeContext) error {
	rlogger.FuncEntry(types.ModuleAmfCtxt, ueCtxt.GetImsiPtr())
	if ueCtxt == nil {
		return fmt.Errorf("invalid input parameter, nil ueCtxt")
	}

	switch key.(type) {
	case AmfUeNgApId:
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, ueCtxt, "delete recored with key(%v) in AmfUeId-UeContext Table", key)
		amfUeIdUeCtxtTable.Del(key)
	case GutiKey:
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, ueCtxt, "delete recored with key(%v) in Guti-UeContext Table", key)
		gutiUeCtxtTable.Del(key)
	case StmsiKey:
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, ueCtxt, "delete recored with key(%v) in Stmsi-UeContext Table", key)
		stmsiUeCtxtTable.Del(key)
	case ImsiKey:
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, ueCtxt, "delete recored with key(%v) in imsi-UeContext Table", key)
		imsiUeCtxtTable.Del(key)
	default:
		return fmt.Errorf("invalid key type")
	}

	return nil
}

func GetUeContext(key interface{}) (ueCtxt *UeContext, err error) {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)

	switch key.(type) {
	case ImsiKey:
		val := imsiUeCtxtTable.Get(key)
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

	case AmfUeNgApId:
		val := amfUeIdUeCtxtTable.Get(key)
		if val == nil {
			err = fmt.Errorf("failed to find UeContext with AmfUeNgApId key(%d)", key.(AmfUeNgApId))
			return
		}
		ctxt, ok := val.(*UeContext)
		if !ok {
			err = fmt.Errorf("invalid ue context type")
			return
		}
		ueCtxt = ctxt

	case GutiKey:
		val := gutiUeCtxtTable.Get(key)
		if val == nil {
			err = fmt.Errorf("failed to find UeContext with 5GGuti key(%s)", key.(GutiKey))
			return
		}
		ctxt, ok := val.(*UeContext)
		if !ok {
			err = fmt.Errorf("invalid ue context type")
			return
		}
		ueCtxt = ctxt

	case StmsiKey:
		val := stmsiUeCtxtTable.Get(key)
		if val == nil {
			err = fmt.Errorf("failed to find UeContext with STmsi key(%s)", key.(StmsiKey))
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
	rlogger.FuncEntry(types.ModuleAmfCtxt, ueCtxt.GetImsiPtr())
	if ueCtxt == nil {
		return fmt.Errorf("invalid input parameter, nil ueCtxt")
	}

	switch key.(type) {
	case ImsiKey:
		imsiUeCtxtTable.Update(key, ueCtxt)
	case AmfUeNgApId:
		amfUeIdUeCtxtTable.Update(key, ueCtxt)
	case GutiKey:
		gutiUeCtxtTable.Update(key, ueCtxt)
	case StmsiKey:
		stmsiUeCtxtTable.Update(key, ueCtxt)
	default:
		return fmt.Errorf("invalid key")
	}

	return nil
}

func LengthOfUeCtxtTbl(key KeyType) uint32 {
	var length uint32
	switch key {
	case ImsiType:
		length = imsiUeCtxtTable.Length()
	case AmfUeNgApIdType:
		length = amfUeIdUeCtxtTable.Length()
	case GutiType:
		length = gutiUeCtxtTable.Length()
	case StmsiType:
		length = stmsiUeCtxtTable.Length()
	default:
		length = 0 //invalid
	}

	return length
}

func DumpAllUeCtxt() string {
	var outstr []string
	outstr = append(outstr, "=== dump all ue context with imsi-uectxt-table ===\n")

	imsiUeCtxtTable.Range(func(key, value interface{}) bool {
		uectxt, ok := value.(*UeContext)
		if !ok {
			return false
		}
		varstr := ffmt.Spjson(*uectxt)
		outstr = append(outstr, fmt.Sprintf("%v : %s\n", key, varstr))
		return true
	})

	outstr = append(outstr, "=== dump all ue context with amfUeId-UeCtxt-Table ===\n")

	amfUeIdUeCtxtTable.Range(func(key, value interface{}) bool {
		uectxt, ok := value.(*UeContext)
		if !ok {
			return false
		}
		varstr := ffmt.Spjson(*uectxt)
		outstr = append(outstr, fmt.Sprintf("%v : %s\n", key, varstr))
		return true
	})

	outstr = append(outstr, "=== dump all ue context with stmsi-UeCtxt-Table ===\n")

	stmsiUeCtxtTable.Range(func(key, value interface{}) bool {
		uectxt, ok := value.(*UeContext)
		if !ok {
			return false
		}
		varstr := ffmt.Spjson(*uectxt)
		outstr = append(outstr, fmt.Sprintf("%v : %s\n", key, varstr))
		return true
	})

	return strings.Join(outstr, " ")
}

func DumpUeCtxtTable() {
	fmt.Println("============== DumpUeCtxtTable ==============")
	fmt.Println("- imsiUeCtxtTable -")
	imsiUeCtxtTable.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})

	fmt.Println("- amfUeIdUeCtxtTable -")
	amfUeIdUeCtxtTable.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})

	fmt.Println("- gutiUeCtxtTable -")
	gutiUeCtxtTable.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})

	fmt.Println("- stmsiUeCtxtTable -")
	stmsiUeCtxtTable.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})
}

func ValuesOfUEContextTbl(key KeyType) (CxtList []*UeContext, err error) {
	switch key {
	case ImsiType:
		imsiUeCtxtTable.Range(func(key, value interface{}) bool {
			ctxt, ok := value.(*UeContext)
			if !ok {
				err = fmt.Errorf("invalid ue session context type")
				return false
			}
			CxtList = append(CxtList, ctxt)
			return true
		})
	case AmfUeNgApIdType:
		amfUeIdUeCtxtTable.Range(func(key, value interface{}) bool {
			ctxt, ok := value.(*UeContext)
			if !ok {
				err = fmt.Errorf("invalid ue session context type")
				return false
			}
			CxtList = append(CxtList, ctxt)
			return true
		})
	case GutiType:
		gutiUeCtxtTable.Range(func(key, value interface{}) bool {
			ctxt, ok := value.(*UeContext)
			if !ok {
				err = fmt.Errorf("invalid ue session context type")
				return false
			}
			CxtList = append(CxtList, ctxt)
			return true
		})
	case StmsiType:
		stmsiUeCtxtTable.Range(func(key, value interface{}) bool {
			ctxt, ok := value.(*UeContext)
			if !ok {
				err = fmt.Errorf("invalid ue session context type")
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

func DumpAmfUeIdUeCtxtTable() string {
	var outstr []string
	amfUeIdUeCtxtTable.Range(func(key, value interface{}) bool {
		uectxt, ok := value.(*UeContext)
		if !ok {
			return false
		}
		varstr := fmt.Sprintf("%x", uectxt.GetAmfUeNgapId())
		keyval, ok := key.(AmfUeNgApId)
		if !ok {
			return false
		}

		outstr = append(outstr, fmt.Sprintf("(%x : %s) ", keyval, varstr))
		return true
	})

	var out string
	for _, v := range outstr {
		out += fmt.Sprintf("AmfUeIdUeCtxtTable {%s}", v)
	}

	return out
}
