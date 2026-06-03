/** Copyright(C),2020-2022
* Author: zmj
* Date: 12/22/20 4:58 PM
* Description:
 */
package ngaplayer

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func (p *LayerMgr) handleUeRadioCapInfoIndMsg(ctx context.Context, msgBuf []byte) error {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	message := ngapmsg.NewUeRadioCapInfoIndMsg()
	message.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
	err := message.Decode(msgBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"failed to decode ue radio capability info indication message")
		return err
	}

	ueCtxt, err := gctxt.GetUeContext(gctxt.AmfUeNgApId(message.AmfUeNgapId))
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"failed to find the ue context with AmfUeNGAPId(%d)",
			message.AmfUeNgapId)
		return err
	}

	//save the ue capability in ue context
	ueCtxt.UeRadioCapInfo = message.UeRadioCapbility

	// todo save more info here
	
	return nil
}
