#ifndef __pduSessionResourceModifyRequestCodec_H__
#define __pduSessionResourceModifyRequestCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <pduSessionResourceModifyRequestPdu.h>

class PduSessResModReqCodec
{
  public:
    PduSessResModReqCodec();
    ~PduSessResModReqCodec();

    void setRanUeNgapId(Uint32 id);
    Uint32 getRanUeNgapId();

    void setAmfUeNgapId(Uint64 id);
    Uint64 getAmfUeNgapId();

    void setRanPagingPriority(Uint16 ranPagingPrt);
    Uint16 getRanPagingPriority();
    bool isRanPagingPriorityPresent();

    void addPduSessResModReqList(PduSessResModReqItem &val);
    PduSessResModReqList& getPduSessResModReqList();

    MsgBuffer encode(NgapOssCtxt* ctxt);
    bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

    bool reset();
    std::string toString();

  private:
    PduSessResModReqPdu msgPdu_m;

    Uint64 amfUeNgapId_m;
    Uint32 ranUeNgapId_m;

    Uint16 ranPagingPriority_m;
    bool ranPagingPriorityPresent_m;

    PduSessResModReqList pduSessResModReqList_m;
};

#endif
