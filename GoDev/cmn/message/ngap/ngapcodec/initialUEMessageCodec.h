#ifndef __initialUEMessageCodec_H__
#define __initialUEMessageCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <initialUEMessagePdu.h>

class InitialUEMessageCodec
{
public:
    InitialUEMessageCodec();
    ~InitialUEMessageCodec();

    void setRanUeNgapId(Uint32 id);
    Uint32 getRanUeNgapId();
    
    void setNasPdu(std::string naspdu);
    std::string getNasPdu();

    void setUserLocationInfo(UserLocationInfo& info);
    UserLocationInfo& getUserLocationInfo();
    
    void setRrcEstablishmentCause(Uint8 cause);
    Uint8 getRrcEstablishmentCause();

    void set5gStmsi(Stmsi5G& stmsi);
    Stmsi5G& get5gStmsi();
    bool is5gStmsiPresent();
        
    void setAmfSetId(Uint8* Id);
    Uint8* getAmfSetId();
    bool isAmfSetPresent();

    void setUeContextRequest(unsigned short ueCtxtReq);
    unsigned short getUeContextRequest();
    bool isUeContextRequestPresent();

    void addAllowedNssai(SNssai& nssai);
    AllowedNssaiList& getAllowedNssaiList();
    bool isAllowedNssaiPresent();

    MsgBuffer encode(NgapOssCtxt* ctxt);
    bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

    bool reset();
    std::string toString();
    
private:

    InitialUEMessagePdu msgPdu_m;

    Uint32 ranUeNgapId_m;

    std::string nasPdu_m;

    UserLocationInfo userLocationInfo_m;

    Uint8 rrcEstablishmentCause_m;
    
    Stmsi5G stmsi_m;
    bool isStmsi5gPrst_m;
    
    AmfSetId amfSetId_m;
    bool isAmfSetIdPrst_m;

    unsigned short ueContextRequest_m;
    bool isUeContextRequestPrst_m;

    AllowedNssaiList allowedNssaiList_m;
    bool isAllowedNssaiPrst_m;
};

#endif