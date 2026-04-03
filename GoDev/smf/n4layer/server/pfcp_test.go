package server

import (
	"context"
	"lite5gc/cmn/types/configure"
	"sync"
	"testing"

	"lite5gc/cmn/types"
)

func TestSmfN4ResponseMsg(t *testing.T) {
	smfCtxt := &types.AppContext{
		Name:     "upf",
		Wg:       &sync.WaitGroup{},
		ConfPath: types.DefConfFileSmf,
	}
	smfCtxt.Ctx, smfCtxt.Cancel = context.WithCancel(context.Background())
	//load smf configuration
	err := configure.LoadConfigSmf("ini", smfCtxt.ConfPath)
	if err != nil {
		panic("load smf config file failed")
	}
	// start N4 tcp server
	errN4 := StartN4Server(smfCtxt)
	if errN4 != nil {
		panic("Failed to start n4 socket server ")
	}

	smfCtxt.Wg.Wait()
}
