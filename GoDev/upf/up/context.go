package up

import "lite5gc/upf/context/ipport"
import "lite5gc/upf/context/pdrcontext"

// n3 context
type HandlerContext struct {
	Ipport *ipport.IpPort
	Msgcxt *pdrcontext.DataFlowContext
}

func (hc HandlerContext) Copy() interface{} {

	cxt := HandlerContext{
		Ipport: hc.Ipport,
		Msgcxt: hc.Msgcxt, //&DataFlowContext{},
	}
	//*cxt.Msgcxt = *hc.Msgcxt
	//rlogger.Trace(moduleTag, types.WARN, nil, "hc copy %+v", hc.Msgcxt)
	//rlogger.Trace(moduleTag, types.WARN, nil, "hc copy %+v", cxt.Msgcxt)
	//rlogger.Trace(moduleTag, types.WARN, nil, "hc copy %s", debug.Stack())
	return cxt
}

func (hc HandlerContext) Delete() {
}
