#include <iostream>
#include <sstream>
#include <string>

#include "nasNonDeliveryIndicationCodec.h"

using namespace std;

NasNonDeliveryIndicationCodec::NasNonDeliveryIndicationCodec()
{
    reset();
}

NasNonDeliveryIndicationCodec::~NasNonDeliveryIndicationCodec()
{
}

void 
NasNonDeliveryIndicationCodec::setRanUeNgapId(Uint32 id)
{
  ranUeNgapId_m = id;
}

Uint32 
NasNonDeliveryIndicationCodec::getRanUeNgapId()
{
  return ranUeNgapId_m;
}

void 
NasNonDeliveryIndicationCodec::setAmfUeNgapId(Uint64 id)
{
  amfUeNgapId_m = id;
}

Uint64 
NasNonDeliveryIndicationCodec::getAmfUeNgapId()
{
  return amfUeNgapId_m;
}


void 
NasNonDeliveryIndicationCodec::setNasPdu(std::string naspdu)
{
  nasPdu_m = naspdu;
}

std::string
NasNonDeliveryIndicationCodec::getNasPdu()
{
  return nasPdu_m;
}

void 
NasNonDeliveryIndicationCodec::setCauseType(Uint16 cType)
{
	causeType_m = cType;
}

Uint16 
NasNonDeliveryIndicationCodec::getCauseType()
{
	return causeType_m;
}

void 
NasNonDeliveryIndicationCodec::setCauseValue(Uint16 value)
{
	causeValue_m = value;
}

Uint16 
NasNonDeliveryIndicationCodec::getCauseValue()
{
	return causeValue_m;
}

bool NasNonDeliveryIndicationCodec::reset()
{
  msgPdu_m.reset_v();

  amfUeNgapId_m = 0;
  ranUeNgapId_m = 0;
  nasPdu_m.clear();
  
  causeType_m = 0;
  causeValue_m = 0;
}

MsgBuffer NasNonDeliveryIndicationCodec::encode(NgapOssCtxt *ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    //encode RAN UE NGAP ID
    RAN_UE_NGAP_ID* ranUeIdPtr=msgPdu_m.getRanUeNgapIdEncodePtr();
   *ranUeIdPtr= ranUeNgapId_m;

    //encode AMF UE NGAP ID
    AMF_UE_NGAP_ID* amfUeIdPtr=msgPdu_m.getAmfUeNgapIdEncodePtr();
   *amfUeIdPtr= amfUeNgapId_m;
  
    //encode NAS PDU
    NAS_PDU* nasPduPtr=msgPdu_m.getNasPduEncodePtr();
    nasPduPtr->length = nasPdu_m.length();
    memcpy(nasPduPtr->value, nasPdu_m.c_str(), nasPduPtr->length); 

	//encode Cause Type and Cause Value
	Cause* causePtr = msgPdu_m.getCauseEncodePtr();
	 switch (causeType_m)
	{
		 case radioNetwork_chosen:
			 causePtr->choice = causeType_m;
			 causePtr->u.radioNetwork = (CauseRadioNetwork)causeValue_m;
			 break;
		 case transport_chosen:
			 causePtr->choice = causeType_m;
			 causePtr->u.transport = (CauseTransport)causeValue_m;
			 break;
		 case nas_chosen:
			 causePtr->choice = causeType_m;
			 causePtr->u.nas = (CauseNas)causeValue_m;
			 break;
		 case protocol_chosen:
			 causePtr->choice = causeType_m;
			 causePtr->u.protocol = (CauseProtocol)causeValue_m;
			 break;
		 case misc_chosen:
			 causePtr->choice = causeType_m;
			 causePtr->u.misc = (CauseMisc)causeValue_m;
			 break;
	
		 default:
		 break;
	}

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool NasNonDeliveryIndicationCodec::decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf)
{
    reset();

    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;
    
    msgPdu_m.setPerBufRef_v(ossbuf);

    if (0 != msgPdu_m.decodePerToOss(ctxt)){
       cout << "Failed to decode Per to OSS" << endl;
       return false;
    }

    //decode RAN UE NGAP ID
    ranUeNgapId_m=*(msgPdu_m.getRanUeNgapIdDecodedPtr());

    //decode amf ue ngap id
    amfUeNgapId_m=*(msgPdu_m.getAmfUeNgapIdDecodedPtr());

    //decode NAS PDU
    NAS_PDU* nasPduPtr=msgPdu_m.getNasPduDecodedPtr();
    nasPdu_m.assign((char*)(nasPduPtr->value),nasPduPtr->length);

	//decode Cause Type and Cause Value
    causeType_m=(msgPdu_m.getCauseDecodedPtr())->choice;
	switch((msgPdu_m.getCauseDecodedPtr())->choice){
		case radioNetwork_chosen:
		causeValue_m=((msgPdu_m.getCauseDecodedPtr())->u).radioNetwork;
		break;
		case transport_chosen:
		causeValue_m=((msgPdu_m.getCauseDecodedPtr())->u).transport;
		break;
		case nas_chosen:
		causeValue_m=((msgPdu_m.getCauseDecodedPtr())->u).nas;
		break;
		case protocol_chosen:
		causeValue_m=((msgPdu_m.getCauseDecodedPtr())->u).protocol;
		break;
		case misc_chosen:
		causeValue_m=((msgPdu_m.getCauseDecodedPtr())->u).misc;
		break;
	}
    return true;
}

string NasNonDeliveryIndicationCodec::toString()
{
    ostringstream oss;

    oss <<"NasNonDeliveryIndicationCodec:" << endl
        <<"RAN UE NGAP ID:(" << ranUeNgapId_m << ")"<<endl
        <<"AMF UE NGAP ID:(" << amfUeNgapId_m << ")"<<endl
        <<"NAS PDU:("<<nasPdu_m <<")"<< endl
		<<" CauseType: "<<causeType_m
		<<" CauseValue: "<<causeValue_m << endl;
   return oss.str();
}


