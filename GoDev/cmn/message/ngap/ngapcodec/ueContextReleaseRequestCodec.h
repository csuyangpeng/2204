#ifndef __ueContextReleaseRequestCodec_H__
#define __ueContextReleaseRequestCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <ueContextReleaseRequestPdu.h>

class UeContextReleaseRequestCodec
{
public:
    UeContextReleaseRequestCodec();
    ~UeContextReleaseRequestCodec();

    void setRanUeNgapId(Uint32 id);
    Uint32 getRanUeNgapId();

    void setAmfUeNgapId(Uint64 id);
    Uint64 getAmfUeNgapId();

	void setCauseType(Uint16 cType);
	Uint16 getCauseType();

    void setCauseValue(Uint16 value);
	Uint16 getCauseValue();

    void addPduSessResListRelReq(unsigned short pduSessResRelReq);
    std::vector<unsigned short>& getPduSessResListRelReq();
	bool isPduSessResListRelReqPrst();

	//bool encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer);
	MsgBuffer encode(NgapOssCtxt *ctxt);
	bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

	bool reset();
	std::string toString();

private:

    UeContextReleaseRequestPdu msgPdu_m;

	Uint64 amfUeNgapId_m;
    Uint32 ranUeNgapId_m;

	Uint16 causeType_m;
    Uint16 causeValue_m;

	std::vector<unsigned short> pduSessResListRelReq_m;
	bool pduSessResListRelReqPrst_m;
};

#endif