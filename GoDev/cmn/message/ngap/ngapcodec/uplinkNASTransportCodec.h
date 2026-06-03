#ifndef __UplinkNASTransportCodec_H__
#define __UplinkNASTransportCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <uplinkNASTransportPdu.h>

class UplinkNASTransportCodec
{
public:
    UplinkNASTransportCodec();
    ~UplinkNASTransportCodec();

    void setRanUeNgapId(Uint32 id);
    Uint32 getRanUeNgapId();

    void setAmfUeNgapId(Uint64 id);
    Uint64 getAmfUeNgapId();
	
    void setNasPdu(std::string naspdu);
    std::string getNasPdu();

    void setUserLocationInfo(UserLocationInfo& info);
    UserLocationInfo& getUserLocationInfo();
	
    //bool encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer);
    MsgBuffer encode(NgapOssCtxt *ctxt);
    bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

    bool reset();
	
    std::string toString();
    
private:

    UplinkNASTransportPdu msgPdu_m;

	Uint64 amfUeNgapId_m;
	
    Uint32 ranUeNgapId_m;

    std::string nasPdu_m;

    UserLocationInfo userLocationInfo_m;

};

#endif

