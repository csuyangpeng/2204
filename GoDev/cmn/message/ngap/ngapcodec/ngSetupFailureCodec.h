#ifndef __ngSetupFailureCodec_H__
#define __ngSetupFailureCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <ngSetupFailurePdu.h>

class NgSetupFailureCodec
{
public:

	NgSetupFailureCodec();
	~NgSetupFailureCodec();

	void setCauseType(Uint16 cType);
	Uint16 getCauseType();

    void setCauseValue(Uint16 value);
	Uint16 getCauseValue();

    void setRelativeTimeToWait(Uint32 time);
    Uint32 getRelativeTimeToWait();
	bool isTimeToWaitPrst();

	MsgBuffer encode(NgapOssCtxt* ctxt);
	bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

	bool reset();
	std::string toString();
	
private:

    NGSetupFailurePdu msgPdu_m;

	Uint16 causeType_m;

    Uint16 causeValue_m;

    Uint32 relativeTimeToWait_m;
	bool timeToWaitPrst_m;

}; 
#endif