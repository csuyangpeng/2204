#ifndef __pduSessionResourceReleaseResponseCodec_H__
#define __pduSessionResourceReleaseResponseCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <pduSessionResourceReleaseResponsePdu.h>

class PduSessResRelRespCodec
{
  public:
    PduSessResRelRespCodec();
    ~PduSessResRelRespCodec();

    void setRanUeNgapId(Uint32 id);
    Uint32 getRanUeNgapId();

    void setAmfUeNgapId(Uint64 id);
    Uint64 getAmfUeNgapId();

    void setUserLocationInfo(UserLocationInfo& info);
    UserLocationInfo& getUserLocationInfo();
    bool isUserLocInfoPrst();

    void addPduSessResRelRespList(PduSessResRelRespItem &val);
    PduSessResRelRespList& getPduSessResRelRespList();

    //bool encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer);
    MsgBuffer encode(NgapOssCtxt *ctxt);
    bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

    bool reset();
    std::string toString();

  private:

    PduSessResRelRespPdu msgPdu_m;

	Uint64 amfUeNgapId_m;
    Uint32 ranUeNgapId_m;

    UserLocationInfo userLocationInfo_m;
    bool userLocInfoPrst_m;

    PduSessResRelRespList pduSessResRelRespList_m;
};

#endif