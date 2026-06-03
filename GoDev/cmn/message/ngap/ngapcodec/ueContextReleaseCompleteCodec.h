#ifndef __ueContextReleaseCompleteCodec_H__
#define __ueContextReleaseCompleteCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <ueContextReleaseCompletePdu.h>

class UeContextReleaseCompleteCodec
{
public:
    UeContextReleaseCompleteCodec();
    ~UeContextReleaseCompleteCodec();

    void setRanUeNgapId(Uint32 id);
    Uint32 getRanUeNgapId();

    void setAmfUeNgapId(Uint64 id);
    Uint64 getAmfUeNgapId();

    void setUserLocationInfo(UserLocationInfo& info);
    UserLocationInfo& getUserLocationInfo();
	bool isUliPrst();

    void setRecommandCellList(RecommandCellList& cell);
    RecommandCellList& getRecommandCellList();

    void setRecommandRanNodeList(RecommandRanNodeList& ranNode);
    RecommandRanNodeList& getRecommandRanNodeList();
	bool isRecommandInfoPrst();

    void addPduSessResListRelCpl(PduSessResRelCplItem& Item);
    PduSessResRelCplList &getPduSessResListRelCpl();
	
    MsgBuffer encode(NgapOssCtxt *ctxt);
    bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

    bool reset();
	
    std::string toString();

private:

    UeContextReleaseCompletePdu msgPdu_m;

	Uint64 amfUeNgapId_m;
    Uint32 ranUeNgapId_m;

    UserLocationInfo userLocationInfo_m;
	bool isUliPrst_m;
	
    RecommandCellList 		recommandCellList_m;
	RecommandRanNodeList	recommandRanNodeList_m;
	bool isRecommandInfoPrst_m;

    PduSessResRelCplList pduSessResListRelCpl_m;
};

#endif