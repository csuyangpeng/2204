#ifndef __initialContextSetupRequestCodec_H__
#define __initialContextSetupRequestCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <initialContextSetupRequestPdu.h>

class InitialContextSetupRequestCodec
{
public:
    InitialContextSetupRequestCodec();
    ~InitialContextSetupRequestCodec();

    void setRanUeNgapId(Uint32 id);
    Uint32 getRanUeNgapId();

    void setAmfUeNgapId(Uint64 id);
    Uint64 getAmfUeNgapId();
    
    void setNasPdu(std::string naspdu);
    std::string getNasPdu();
    bool isNasPduPrst();
    
    void setOldAmfName(std::string name);
    std::string getOldAmfName();
    bool isAmfNamePrst();
        
    void setUeAmbr(UeAmbr& ambr);
    UeAmbr& getUeAmbr();
    bool isUeAMBRPrst();

    void setGuami(Guami& val);
    Guami& getGuami();

    void setUeSecurityCap(UeSecurityCap& val);
    UeSecurityCap& getUeSecurityCap();

    void setIndexToRfsp(long val);
    long getIndexToRfsp();
    bool isIndexToRFSPPrst();

    void addAllowedSNssai(SNssai& val);
    AllowedNssaiList& getAllowedSNssaiList();

    void setSecurityKey(std::string val);
    std::string getSecurityKey();

    void setTraceAct(TraceAct& val);
    TraceAct& getTraceAct();
    bool isTraceActPrst();

    void setUERadioCapability(std::string val);
    std::string getUERadioCapability();
    bool isUeRadioCapPrst();

    void setMaskedImeiSv(std::string val);
    std::string getMaskedImeiSv();
    bool isMaskedImeiSvPrst();

    void setEmergeFBInd(EmergeFBReqInd & val);
    EmergeFBReqInd& getEmergeFBInd();
    bool isEmergeFallbackPrst();

    void setRRCITRR(Uint32 val);
    Uint32 getRRCITRR();
    bool isRRCReportPrst();

    void setUeRadioCapForPaging(UeRadioCapForPaging& val);
    UeRadioCapForPaging& getUeRadioCapForPaging();
    bool isUeRadioCapForPagingPrst();

    void setRedirForVoiceEpsFallback(Uint8 val);
    Uint8 getRedirForVoiceEpsFallback();
    bool isRedirForVoiceEpsFallbackPrst();

    void addPduSessResReqList(PduSessResSetupReqItem &val);
    PduSessResSetupReqList& getPduSessResReqList();

    MsgBuffer encode(NgapOssCtxt* ctxt);
    //bool encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer);
    bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

    bool reset();
    std::string toString();
    
private:

    InitialContextSetupRequestPdu msgPdu_m;

    Uint64 amfUeNgapId_m;
    Uint32 ranUeNgapId_m;

    std::string nasPdu_m;
    bool nasPduPrst_m;

    std::string oldAmfName_m;
    bool amfNamePrst_m;

    long indexToRfsp_m;
    bool indexToRFSPPrst_m;

    UeAmbr ueAmbr_m;
    bool ueAMBRPrst_m;
    
    AllowedNssaiList allowedSNssaiList_m;

    Guami guami_m;

    UeSecurityCap ueSecurityCap_m;

    std::string securityKey_m;

    TraceAct traceAct_m;
    bool traceActPrst_m;

    std::string ueRadioCap_m;
    bool ueRadioCapPrst_m;

    std::string maskImeiSv_m;
    bool maskedImeiSvPrst_m;

    EmergeFBReqInd emergeFBReqInd_m;
    bool emergeFBReqIndPrst_m;

    Uint32 rrcITRR_m;
    bool rrcReportPrst_m;

    UeRadioCapForPaging ueRadioCapForPaging_m;
    bool ueRadioCapForPagingPrst_m;

    Uint8 redirVoiceEpsFallback_m;
    bool redirVoiceEpsFallbackPrst_m;

    PduSessResSetupReqList pduSessResReqList_m;
};

#endif
