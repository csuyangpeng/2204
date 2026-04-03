package configure

import (
	"github.com/satori/go.uuid"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"sync"
)

//AMF configuration
type AmfConfig struct {
	Version    VerConfig
	Logger     LoggerConfig
	N2         IpAddr
	Sctp       SctpParameter
	Service    AmfService
	NAS        NasConfig
	PlmnList   types3gpp.PlmnList
	TaiLists   []types3gpp.TAIList
	Nssai      []types3gpp.Snssai
	Sbi        SBI
	ConfRWlock sync.RWMutex
}

func (p *AmfConfig) ConvertAmfConfig(cmData *CmAmfConfig) error {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)
	// check the input parameter
	if cmData == nil {
		return types.ErrInputParaNil
	}

	// version
	p.Version = cmData.Version

	// logger
	p.Logger = cmData.Logger

	// n2 if
	p.N2 = cmData.N2if

	// sctp
	p.Sctp = cmData.Sctp

	// SBI
	p.Sbi = cmData.Sbi

	// amf service
	u := uuid.NewV4()
	p.Service.AmfInstanceId = u.String()
	p.Service.AmfName = cmData.Service.AmfName
	p.Service.AmfRelCap = cmData.Service.RelativeAmfCapacity
	p.Service.AmfIdentifier.Pointer = cmData.Service.AmfIdentifier.Pointer
	p.Service.AmfIdentifier.RegionId = cmData.Service.AmfIdentifier.RegionId
	p.Service.AmfIdentifier.SetAmfSetIdU16(cmData.Service.AmfIdentifier.SetId)

	// amf nas
	err := p.NAS.SecCap.StoreNrIntAlgo(cmData.Nas.IntegrityAlgorithm)
	if err != nil {
		rlogger.Trace(types.ModuleCmnConf, rlogger.ERROR, nil, "failed to load integrity protect algo, error(%s)", err)
		return err
	}
	err = p.NAS.SecCap.StoreNrEncAlgo(cmData.Nas.EncryptAlgorithm)
	if err != nil {
		rlogger.Trace(types.ModuleCmnConf, rlogger.ERROR, nil, "failed to load cipher protect algo, error(%s)", err)
		return err
	}
	p.NAS.T3512min = cmData.Nas.T3512min
	p.NAS.T3513Sec = cmData.Nas.T3513Sec
	p.NAS.T3502min = cmData.Nas.T3502min
	p.NAS.T3550sec = cmData.Nas.T3550sec
	p.NAS.T3560sec = cmData.Nas.T3560sec
	p.NAS.T3570sec = cmData.Nas.T3570sec
	p.NAS.T3522sec = cmData.Nas.T3522sec
	p.NAS.T3555sec = cmData.Nas.T3555sec
	p.NAS.T3565sec = cmData.Nas.T3565sec

	// plmn list
	for _, v := range cmData.Plmnlist {
		var plmn types3gpp.PlmnID
		err := plmn.SetString(v)
		if err != nil {
			rlogger.Trace(types.ModuleCmnConf, rlogger.ERROR, nil,
				"failed to set plmn string(%s)", v)
			continue
		}

		p.PlmnList.AddPlmn(plmn)
	}

	// tai lists
	for _, v := range cmData.TaiLists {
		var tailist types3gpp.TAIList
		var plmn types3gpp.PlmnID
		err := plmn.SetString(v.Plmn)
		if err != nil {
			rlogger.Trace(types.ModuleCmnConf, rlogger.ERROR, nil,
				"failed to set plmn string(%s)", v)
			continue
		}
		for _, tacv := range v.Tac {
			tailist.MultiplePLMN = false // todo current unique plmn supported
			var tai types3gpp.TAI
			tai.Plmn = plmn
			tai.Tac.SetTac(tacv)
			tailist.Tais = append(tailist.Tais, tai)
		}

		p.TaiLists = append(p.TaiLists, tailist)
	}

	//nssai
	for _, v := range cmData.Nssai {
		var snssai types3gpp.Snssai
		snssai.Sst = v.Sst
		if v.Sd != 0 {
			snssai.Sd = v.Sd
			snssai.SdPrst = true
		} else {
			snssai.SdPrst = false
		}
		p.Nssai = append(p.Nssai, snssai)
	}

	return nil
}
