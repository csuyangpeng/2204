package gctxt

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/syncmap"
	"lite5gc/cmn/types"
)

var seidSmfN4CxtTable syncmap.SyncMap //key:seid,value:N4Cxt

func ValuesOfSmfN4InfoTbl(key KeyType) (CxtList []*N4SessionContext, err error) {
	switch key {
	case N4SessionIDCxtType:
		seidSmfN4CxtTable.Range(func(key, value interface{}) bool {
			//fmt.Println(key, value)
			ctxt, ok := value.(*N4SessionContext)
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

func CreateN4Context(key interface{}) error {
	rlogger.FuncEntry(types.ModuleSmfCtxt, nil)
	var err error
	switch key.(type) {
	case N4SessionIDKey:
		n4Ctxt := &N4SessionContext{}
		err = seidSmfN4CxtTable.Set(key, n4Ctxt) // seidSmfN4CxtTable 隐含取地址调用方法
		if err != nil {
			err = fmt.Errorf("failed to create the seidSmfN4CxtTable, error(%s)", err)
		}
	default:
		err = fmt.Errorf("invalid key")
	}

	return err
}

func AddIndexN4Context(key interface{}, ctxt *N4SessionContext) error {
	rlogger.FuncEntry(types.ModuleSmfCtxt, ctxt)

	var err error

	switch key.(type) {
	case N4SessionIDKey:
		err = seidSmfN4CxtTable.Set(key, ctxt)
		if err != nil {
			err = fmt.Errorf("failed to set key(%d),err(%s)", key.(N4SessionIDKey), err)
		}
	default:
		err = fmt.Errorf("invalid key type")
	}
	return err
}

func GetN4Context(key interface{}) (n4Ctxt *N4SessionContext, err error) {
	rlogger.FuncEntry(types.ModuleSmfCtxt, nil)

	switch key.(type) {
	case N4SessionIDKey:
		val := seidSmfN4CxtTable.Get(key)
		if val == nil {

			err = fmt.Errorf("failed to find N4SessionContext with SEID key(%d)", key.(N4SessionIDKey))
			//rlogger.Trace(types.ModuleSmfCtxt, rlogger.ERROR, nil, err)
			return
		}
		ctxt, ok := val.(*N4SessionContext)
		if !ok {
			err = fmt.Errorf("invalid n4 session context type")
			//rlogger.Trace(types.ModuleSmfCtxt, rlogger.ERROR, nil, err)
			return
		}
		n4Ctxt = ctxt

	default:
		err = fmt.Errorf("invalid key")
	}
	//rlogger.Trace(types.ModuleSmfCtxt, rlogger.ERROR, nil, err)
	return
}

func UpdateN4Context(key interface{}, n4Ctxt *N4SessionContext) error {
	rlogger.FuncEntry(types.ModuleSmfCtxt, n4Ctxt)
	if n4Ctxt == nil {
		return fmt.Errorf("invalid input parameter, nil n4Ctxt")
	}

	switch key.(type) {
	case N4SessionIDKey:
		seidSmfN4CxtTable.Update(key, n4Ctxt)
	default:
		return fmt.Errorf("invalid key")
	}

	return nil
}

func DeleteN4Context(key interface{}) error {
	rlogger.FuncEntry(types.ModuleSmfCtxt, nil)

	switch key.(type) {
	case N4SessionIDKey:
		seidSmfN4CxtTable.Del(key)
	default:
		return fmt.Errorf("invalid key")
	}

	return nil
}

func LengthOfN4ContextTbl(key KeyType) uint64 {
	var length uint64
	switch key {
	case N4SessionIDCxtType:
		length = seidSmfN4CxtTable.Length64()
	default:
		length = 0 //invalid
	}

	return length
}
