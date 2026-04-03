#ifndef __errorIndicationCodec_H__
#define __errorIndicationCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <errorIndicationPdu.h>

class ErrorIndicationCodec
{
public:
    ErrorIndicationCodec();
    ~ErrorIndicationCodec();

    void setRanUeNgapId(Uint32 id);
    Uint32 getRanUeNgapId();
    bool isRanUeNgapIdPrst();

    void setAmfUeNgapId(Uint64 id);
    Uint64 getAmfUeNgapId();
    bool isAmfUeNgapIdPrst();

	void setCauseType(Uint16 cType);
	Uint16 getCauseType();
    bool isCauseTypePrst();

    void setCauseValue(Uint16 value);
	Uint16 getCauseValue();

    MsgBuffer encode(NgapOssCtxt* ctxt);
    //bool encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer);
    bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

    bool reset();
    std::string toString();

private:

    ErrorIndicationPdu msgPdu_m;

	Uint64 amfUeNgapId_m;
    bool amfUeNgapIdPrst_m;
    Uint32 ranUeNgapId_m;
    bool ranUeNgapIdPrst_m;

	Uint16 causeType_m;
    Uint16 causeValue_m;
    bool causeTypePrst_m;
};
#endif