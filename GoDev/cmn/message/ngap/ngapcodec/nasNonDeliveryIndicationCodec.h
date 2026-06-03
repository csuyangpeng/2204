#ifndef __nasNonDeliveryIndicationCodec_H__
#define __nasNonDeliveryIndicationCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <nasNonDeliveryIndicationPdu.h>

class NasNonDeliveryIndicationCodec
{
public:
    NasNonDeliveryIndicationCodec();
    ~NasNonDeliveryIndicationCodec();

    void setRanUeNgapId(Uint32 id);
    Uint32 getRanUeNgapId();

    void setAmfUeNgapId(Uint64 id);
    Uint64 getAmfUeNgapId();
	
    void setNasPdu(std::string naspdu);
    std::string getNasPdu();

	void setCauseType(Uint16 cType);
	Uint16 getCauseType();

    void setCauseValue(Uint16 value);
	Uint16 getCauseValue();
	
    MsgBuffer encode(NgapOssCtxt *ctxt);
    bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

    bool reset();
	
    std::string toString();
    
private:

    NasNonDeliveryIndicationPdu msgPdu_m;

	Uint64 amfUeNgapId_m;
	
    Uint32 ranUeNgapId_m;

    std::string nasPdu_m;

	Uint16 causeType_m;
    Uint16 causeValue_m;

};

#endif


