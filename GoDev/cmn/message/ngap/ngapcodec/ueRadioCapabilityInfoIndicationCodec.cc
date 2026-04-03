#include <iostream>
#include <sstream>
#include <string>

#include "ueRadioCapabilityInfoIndicationCodec.h"

using namespace std;

UeRadioCapaInfoIndicationCodec::UeRadioCapaInfoIndicationCodec() 
{
    reset();
}

UeRadioCapaInfoIndicationCodec::~UeRadioCapaInfoIndicationCodec()
{
}

void UeRadioCapaInfoIndicationCodec::setRanUeNgapId(Uint32 id) 
{
    ranUeNgapId_m = id;
}

Uint32 UeRadioCapaInfoIndicationCodec::getRanUeNgapId() 
{
    return ranUeNgapId_m;
}

void UeRadioCapaInfoIndicationCodec::setAmfUeNgapId(Uint64 id) 
{
    amfUeNgapId_m = id;
}

Uint64 UeRadioCapaInfoIndicationCodec::getAmfUeNgapId() 
{
    return amfUeNgapId_m;
}

std::string UeRadioCapaInfoIndicationCodec::getUERadioCapability()
{
    return ueRadioCap_m;
}

void UeRadioCapaInfoIndicationCodec::setUERadioCapability(std::string val)
{

    ueRadioCap_m = val;
}

UeRadioCapForPaging& UeRadioCapaInfoIndicationCodec::getUeRadioCapForPaging()
{
    return ueRadioCapForPaging_m;
}

void UeRadioCapaInfoIndicationCodec::setUeRadioCapForPaging(UeRadioCapForPaging& val)
{
    ueRadioCapForPagingPrst_m=true;
    ueRadioCapForPaging_m = val;
}

bool UeRadioCapaInfoIndicationCodec::isUeRadioCapForPagingPrst()
{
    return ueRadioCapForPagingPrst_m;
}

bool UeRadioCapaInfoIndicationCodec::reset() 
{
    msgPdu_m.reset_v();

    amfUeNgapId_m = 0;
    ranUeNgapId_m = 0;

    ueRadioCap_m.clear();

    ueRadioCapForPaging_m.ueRadioCapForPagingNr.clear();
    ueRadioCapForPaging_m.ueRadioCapForPagingEutra.clear();
    ueRadioCapForPagingPrst_m=false;
}

//bool UeRadioCapaInfoIndicationCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer UeRadioCapaInfoIndicationCodec::encode(NgapOssCtxt *ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    // encode RAN UE NGAP ID
    RAN_UE_NGAP_ID* ranUeIdPtr = msgPdu_m.getRanUeNgapIdEncodePtr();
    *ranUeIdPtr = ranUeNgapId_m;

    // encode AMF UE NGAP ID
    AMF_UE_NGAP_ID* amfUeIdPtr = msgPdu_m.getAmfUeNgapIdEncodePtr();
    *amfUeIdPtr = amfUeNgapId_m;

    // encode UE Radio Capability
    UERadioCapability* ueRadioCapPtr = msgPdu_m.getUeRadioCapabilityEncodePtr();
    ueRadioCapPtr->length = ueRadioCap_m.length();
    memcpy(ueRadioCapPtr->value, ueRadioCap_m.c_str(), ueRadioCapPtr->length);

    // encode UE Radio Capability For Paging
    if(ueRadioCapForPagingPrst_m) {
        UERadioCapabilityForPaging* ueRadioCapForPagingPtr = 
            msgPdu_m.getUeRadioCapabilityForPagingEncodePtr();
		if(ueRadioCapForPaging_m.ueRadioCapForPagingNrPrst) {
			ueRadioCapForPagingPtr->bit_mask |= uERadioCapabilityForPagingOfNR_present;
            ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfNR
			    .length = ueRadioCapForPaging_m.ueRadioCapForPagingNr.length();
            memcpy(ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfNR.value, 
                   ueRadioCapForPaging_m.ueRadioCapForPagingNr.c_str(), 
                   ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfNR.length);
		}
		if(ueRadioCapForPaging_m.ueRadioCapForPagingEutraPrst) {
			ueRadioCapForPagingPtr->bit_mask |= uERadioCapabilityForPagingOfEUTRA_present;
            ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfEUTRA
			    .length = ueRadioCapForPaging_m.ueRadioCapForPagingEutra.length();
            memcpy(ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfEUTRA.value, 
                   ueRadioCapForPaging_m.ueRadioCapForPagingEutra.c_str(), 
                   ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfEUTRA.length);
		}
    }

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool UeRadioCapaInfoIndicationCodec::decode(NgapOssCtxt* ctxt,MsgBuffer msgBuf) 
{
    reset();

    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;

    msgPdu_m.setPerBufRef_v(ossbuf);

    if (0 != msgPdu_m.decodePerToOss(ctxt)) {
        cout << "Failed to decode Per to OSS" << endl;
        return false;
    }
    // decode RAN UE NGAP ID
    ranUeNgapId_m = *(msgPdu_m.getRanUeNgapIdDecodedPtr());

    // decode amf ue ngap id
    amfUeNgapId_m = *(msgPdu_m.getAmfUeNgapIdDecodedPtr());

    // decode UE Radio Capability
    UERadioCapability* ueRadioCapPtr = msgPdu_m.getUeRadioCapabilityDecodedPtr();
    ueRadioCap_m.assign((char*)(ueRadioCapPtr->value), ueRadioCapPtr->length);

    // decode UE Radio Capability For Paging
    if(msgPdu_m.isUeRadioCapabilityForPagingPresent()) {
        if(msgPdu_m.getUeRadioCapabilityForPagingDecodedPtr()->bit_mask & uERadioCapabilityForPagingOfNR_present)
       	{
            UERadioCapabilityForPaging* ueRadioCapForPagingPtr = 
                msgPdu_m.getUeRadioCapabilityForPagingDecodedPtr();
			ueRadioCapForPaging_m.ueRadioCapForPagingNrPrst = true;
            ueRadioCapForPaging_m.ueRadioCapForPagingNr
				.assign((char*)(ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfNR.value), 
                                ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfNR.length);
        }
        if(msgPdu_m.getUeRadioCapabilityForPagingDecodedPtr()->bit_mask & uERadioCapabilityForPagingOfEUTRA_present)
       	{
            UERadioCapabilityForPaging* ueRadioCapForPagingPtr = 
                msgPdu_m.getUeRadioCapabilityForPagingDecodedPtr();
			ueRadioCapForPaging_m.ueRadioCapForPagingEutraPrst = true;
            ueRadioCapForPaging_m.ueRadioCapForPagingEutra
				.assign((char*)(ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfEUTRA.value), 
                                ueRadioCapForPagingPtr->uERadioCapabilityForPagingOfEUTRA.length);
        }

        ueRadioCapForPagingPrst_m=true;
    }
  return true;
}

string UeRadioCapaInfoIndicationCodec::toString(){
   ostringstream oss;

   oss << "UeRadioCapaInfoIndicationCodec:" << endl
       << "1.AMF UE NGAP ID:(" << amfUeNgapId_m << ")" << endl
       << "2.RAN UE NGAP ID:(" << ranUeNgapId_m << ")" << endl
       << "3.UERadioCapability:(" << ueRadioCap_m << ")" << endl
	   << "UeRadioCapabilityForPaging:Prst" <<ueRadioCapForPagingPrst_m<< endl
	   << "Nr:Prst" << ueRadioCapForPaging_m.ueRadioCapForPagingNrPrst << endl
	   << " value " << ueRadioCapForPaging_m.ueRadioCapForPagingNr <<endl
	   << "Eutra:Prst" << ueRadioCapForPaging_m.ueRadioCapForPagingEutraPrst << endl
	   << " value " << ueRadioCapForPaging_m.ueRadioCapForPagingEutra <<endl;

   return oss.str(); 
}
