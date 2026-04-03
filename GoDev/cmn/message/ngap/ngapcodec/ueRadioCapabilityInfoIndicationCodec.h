#ifndef __ueRadioCapabilityInfoIndicationCodec_H__
#define __ueRadioCapabilityInfoIndicationCodec_H__

#include <string>
#include <ngapMsgTypes.h>
#include <ngapOssCtxt.h>
#include <ueRadioCapabilityInfoIndicationPdu.h>

class UeRadioCapaInfoIndicationCodec
{
public:
    UeRadioCapaInfoIndicationCodec();
    ~UeRadioCapaInfoIndicationCodec();

    void setRanUeNgapId(Uint32 id);
    Uint32 getRanUeNgapId();

    void setAmfUeNgapId(Uint64 id);
    Uint64 getAmfUeNgapId();

    void setUERadioCapability(std::string val);
    std::string getUERadioCapability();

    void setUeRadioCapForPaging(UeRadioCapForPaging& val);
    UeRadioCapForPaging& getUeRadioCapForPaging();
    bool isUeRadioCapForPagingPrst();

    //bool encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer);
    MsgBuffer encode(NgapOssCtxt *ctxt);
    bool decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf);

    bool reset();
    std::string toString();

private:
    UeRadioCapaInfoIndicationPdu msgPdu_m;

	Uint64 amfUeNgapId_m;
    Uint32 ranUeNgapId_m;

    std::string ueRadioCap_m;

    UeRadioCapForPaging ueRadioCapForPaging_m;
    bool ueRadioCapForPagingPrst_m;
};

#endif