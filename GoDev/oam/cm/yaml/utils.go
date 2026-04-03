package yaml

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func checkError(err error) {
	if err != nil {
		rlogger.Trace(types.ModuleOamCm, rlogger.ERROR, nil, "package yaml occur %v", err)
		panic(err)	// todo return error
	}
}


