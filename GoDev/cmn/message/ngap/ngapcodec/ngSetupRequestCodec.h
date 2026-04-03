#ifndef __ngSetupRequestCodec_H__
#define __ngSetupRequestCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <ngSetupRequestPdu.h>

class NgSetupRequestCodec
{
public:
    NgSetupRequestCodec();
    ~NgSetupRequestCodec();

    void setRanNodeName(std::string name);
    bool getRNodeNamePrst();
    std::string getRanNodeName();

    void setPagingDrx(Uint8 pdrx);
    Uint8 getPagingDrx();

    void setGgnbId(GGnbId& ggnbid);
    GGnbId& getGgnbId();

    void addSupTAList(SupTAItem& stai);
    SupTAList& getSupTAList();

    void setUeRetentionInfo(Uint8 ueRetentionInfo);
    bool getUeRetentionInfoPrst();
    Uint8 getUeRetentionInfo();

    void dumpMessage();

    MsgBuffer encode(NgapOssCtxt* ctxt);
    bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

    bool reset();

private:

    NGSetupRequestPdu msgPdu_m;

    bool rNodeNamePresent_m;
    std::string ranNodeName_m;

    Uint8 pagingDrx_m;

    GGnbId gGnbId_m;

    SupTAList supTAList_m;

    bool  UeRetentionInfoPrst_m;
    Uint8 UeRetentionInfo_m;
};
#endif