#ifndef __ueContextReleaseCommandCodec_H__
#define __ueContextReleaseCommandCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <ueContextReleaseCommandPdu.h>

class UeContextReleaseCommandCodec
{
public:
    UeContextReleaseCommandCodec();
    ~UeContextReleaseCommandCodec();

	void setUeNgapIdsType(Uint16 idType);
	Uint16 getUeNgapIdsType();

    void setRanUeNgapId(Uint32 id);
    Uint32 getRanUeNgapId();

    void setAmfUeNgapId(Uint64 id);
    Uint64 getAmfUeNgapId();

	void setCauseType(Uint16 cType);
	Uint16 getCauseType();

    void setCauseValue(Uint16 value);
	Uint16 getCauseValue();

	MsgBuffer encode(NgapOssCtxt *ctxt);
	bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

	bool reset();
	std::string toString();

private:
    UeContextReleaseCommandPdu msgPdu_m;

    Uint16 ueNgapIdsType_m;
	Uint32 ranNgapId_m;
	Uint64 amfNgapId_m;

	Uint16 causeType_m;
    Uint16 causeValue_m;

};

#endif