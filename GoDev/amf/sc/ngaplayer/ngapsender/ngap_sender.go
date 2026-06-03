/** Copyright(C),2020-2022
* Author: zmj
* Date: 11/25/20 1:46 PM
* Description:
 */
package ngapsender

import (
	"context"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
)

type NgapSender struct {
	scId    uint32
	ossCtxt codec.OssCtxt
	ctxt    context.Context
}

func NewNgapSender(id uint32, ossCtxt codec.OssCtxt) *NgapSender {
	return &NgapSender{
		scId:    id,
		ossCtxt: ossCtxt,
	}
}
