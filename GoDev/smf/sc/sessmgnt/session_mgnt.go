package sessmgnt

import (
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

type SessMGMT struct {
	scId    uint32
	ossCtxt codec.OssCtxt
}

func NewSessMgmt(scId uint32) *SessMGMT {
	return &SessMGMT{
		scId: scId,
	}
}

func (p *SessMGMT) Initialize() {
	p.ossCtxt = codec.NewOssCtxt()
}

func (p *SessMGMT) Cleanup() {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)
	//delete oss ctxt
	codec.DeleteOssCtxt(p.ossCtxt)
}

func (p *SessMGMT) GetOssCtxt() codec.OssCtxt {
	return p.ossCtxt
}

func (p *SessMGMT) GetScId() uint32 {
	return p.scId
}
