#ifndef __pduSessionResourceReleaseCommandCodec_H__
#define __pduSessionResourceReleaseCommandCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <pduSessionResourceReleaseCommandPdu.h>

class PduSessResRelCmdCodec
{
  public:
    PduSessResRelCmdCodec();
    ~PduSessResRelCmdCodec();

    void setRanUeNgapId(Uint32 id);
    Uint32 getRanUeNgapId();

    void setAmfUeNgapId(Uint64 id);
    Uint64 getAmfUeNgapId();
	
    void setNasPdu(std::string naspdu);
    std::string getNasPdu();
    bool isNasPduPrst();

    void setRanPagingPriority(Uint16 ranPagingPrt);
    Uint16 getRanPagingPriority();
	bool isRanPagingPriorityPresent();

    void addPduSessResRelCmdList(PduSessResRelCmdItem &val);
    PduSessResRelCmdList& getPduSessResRelCmdList();

    MsgBuffer encode(NgapOssCtxt* ctxt);
    bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

    bool reset();
    std::string toString();

  private:
    PduSessResRelCmdPdu msgPdu_m;

	Uint64 amfUeNgapId_m;
    Uint32 ranUeNgapId_m;

    std::string nasPdu_m;
    bool nasPduPrst_m;

	Uint16 ranPagingPriority_m;
	bool ranPagingPriorityPresent_m;

    PduSessResRelCmdList pduSessResRelCmdList_m;
};

#endif