#ifndef __initialContextSetupFailureCodec_H__
#define __initialContextSetupFailureCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <initialContextSetupFailurePdu.h>

class InitialContextSetupFailureCodec
{
public:
    InitialContextSetupFailureCodec();
    ~InitialContextSetupFailureCodec();

    void setRanUeNgapId(Uint32 id);
    Uint32 getRanUeNgapId();

    void setAmfUeNgapId(Uint64 id);
    Uint64 getAmfUeNgapId();

	void setCauseType(Uint16 cType);
	Uint16 getCauseType();
    void setCauseValue(Uint16 value);
	Uint16 getCauseValue();

    void addPduSessResFailToSetupList(PduSessResFailedSetupItem &val);
    PduSessResFailedSetupList& getPduSessResFailToSetupList();

    MsgBuffer encode(NgapOssCtxt* ctxt);
    bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

    bool reset();
    std::string toString();
    
private:

    InitialContextSetupFailurePdu msgPdu_m;

	Uint64 amfUeNgapId_m;
    Uint32 ranUeNgapId_m;

	Uint16 causeType_m;
    Uint16 causeValue_m;

    PduSessResFailedSetupList pduSessResFailToSetupList_m;
};

#endif