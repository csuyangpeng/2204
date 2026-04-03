#ifndef __downlinkNASTransportCodec_H__
#define __downlinkNASTransportCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <downlinkNASTransportPdu.h>

class DownlinkNASTransportCodec
{
public:
    DownlinkNASTransportCodec();
    ~DownlinkNASTransportCodec();

    void setRanUeNgapId(Uint32 id);
    Uint32 getRanUeNgapId();

    void setAmfUeNgapId(Uint64 id);
    Uint64 getAmfUeNgapId();
	
    void setNasPdu(std::string naspdu);
    std::string getNasPdu();

    void setRanPagingPriority(Uint16 ranPagingPrt);
    Uint16 getRanPagingPriority();
	bool isRanPagingPriorityPresent();
    
    void setOldAmfName(std::string name);
    std::string getOldAmfName();
    bool isOldAmfNamePresent();

    void setIndexToRfsp(long index);
    long getIndexToRfsp();
    bool isIndexToRfspPresent();
        
    void setUeAmbr(UeAmbr& ambr);
    UeAmbr& getUeAmbr();
    bool isUeAmbrPresent();

    void addAllowedNssai(SNssai& nssai);
    AllowedNssaiList& getAllowedNssaiList();
    bool isAllowedNssaiPresent();

    MsgBuffer encode(NgapOssCtxt* ctxt);
    bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

    bool reset();
    std::string toString();
    
private:

    DownlinkNASTransportPdu msgPdu_m;

	Uint64 amfUeNgapId_m;
    Uint32 ranUeNgapId_m;

    std::string nasPdu_m;

    std::string oldAmfName_m;
	bool oldAmfNamePresent_m;

	Uint16 ranPagingPriority_m;
	bool ranPagingPriorityPresent_m;

	long indexToRfsp_m;
	bool indexToRfspPresent_m;

	UeAmbr ueAmbr_m;
	bool ueAmbrPresent_m;
	
    AllowedNssaiList allowedNssaiList_m;
    bool allowedNssaiPrst_m;
};

#endif

