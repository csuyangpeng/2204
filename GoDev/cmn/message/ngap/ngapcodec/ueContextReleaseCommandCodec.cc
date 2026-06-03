#include <iostream>
#include <sstream>
#include <string>

#include "ueContextReleaseCommandCodec.h"

using namespace std;

UeContextReleaseCommandCodec::UeContextReleaseCommandCodec()
{
    reset();
}

UeContextReleaseCommandCodec::~UeContextReleaseCommandCodec()
{
}

bool
UeContextReleaseCommandCodec::reset()
{
    msgPdu_m.reset_v();

	causeType_m = 0;
    causeValue_m = 0;

    ueNgapIdsType_m = 0;
    amfNgapId_m = 0;
    ranNgapId_m = 0;

    return true;
}

void 
UeContextReleaseCommandCodec::setCauseType(Uint16 cType)
{
	causeType_m = cType;
}

Uint16 
UeContextReleaseCommandCodec::getCauseType()
{
	return causeType_m;
}

void 
UeContextReleaseCommandCodec::setCauseValue(Uint16 value)
{
	causeValue_m = value;
}

Uint16 
UeContextReleaseCommandCodec::getCauseValue()
{
	return causeValue_m;
}

void 
UeContextReleaseCommandCodec::setUeNgapIdsType(Uint16 idType)
{
	ueNgapIdsType_m = idType;
}

Uint16 
UeContextReleaseCommandCodec::getUeNgapIdsType()
{
	return ueNgapIdsType_m;
}

// void 
// UeContextReleaseCommandCodec::setUeNgapIdsValue(UeNgapIds &value)
// {
// 	ueNgapIdsValue_m = value;
// }

// UeNgapIds&
// UeContextReleaseCommandCodec::getUeNgapIdsValue()
// {
// 	return ueNgapIdsValue_m;
// }
void UeContextReleaseCommandCodec::setRanUeNgapId(Uint32 id) 
{
    ranNgapId_m = id;
}

Uint32 UeContextReleaseCommandCodec::getRanUeNgapId() 
{
    return ranNgapId_m;
}

void UeContextReleaseCommandCodec::setAmfUeNgapId(Uint64 id)
{
    amfNgapId_m = id;
}

Uint64 UeContextReleaseCommandCodec::getAmfUeNgapId()
{
    return amfNgapId_m;
}

//bool UeContextReleaseCommandCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer UeContextReleaseCommandCodec::encode(NgapOssCtxt *ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    //Ue_Ngap_Ids
    UE_NGAP_IDs* ueNgapIdsPtr=msgPdu_m.getUeNgapIdsEncodePtr();
    switch (ueNgapIdsType_m)
    {
        case uE_NGAP_ID_pair_chosen:
            ueNgapIdsPtr->choice = ueNgapIdsType_m;
            ueNgapIdsPtr->u.uE_NGAP_ID_pair.aMF_UE_NGAP_ID = amfNgapId_m;
            ueNgapIdsPtr->u.uE_NGAP_ID_pair.rAN_UE_NGAP_ID = ranNgapId_m;
            break;
        case aMF_UE_NGAP_ID_chosen:
            ueNgapIdsPtr->choice = ueNgapIdsType_m;
            ueNgapIdsPtr->u.aMF_UE_NGAP_ID = amfNgapId_m;
            break;

        default:
        break;
    }

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

bool UeContextReleaseCommandCodec::decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf)
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

	//decode Cause Type and Cause Value
    causeType_m = msgPdu_m.getCauseDecodedPtr()->choice;
	switch(msgPdu_m.getCauseDecodedPtr()->choice){
		case radioNetwork_chosen:
		causeValue_m = msgPdu_m.getCauseDecodedPtr()->u.radioNetwork;
		break;
		case transport_chosen:
		causeValue_m = msgPdu_m.getCauseDecodedPtr()->u.transport;
		break;
		case nas_chosen:
		causeValue_m = msgPdu_m.getCauseDecodedPtr()->u.nas;
		break;
		case protocol_chosen:
		causeValue_m = msgPdu_m.getCauseDecodedPtr()->u.protocol;
		break;
		case misc_chosen:
		causeValue_m = msgPdu_m.getCauseDecodedPtr()->u.misc;
		break;
	}

	//decode UeNgapIds
    ueNgapIdsType_m=msgPdu_m.getUeNgapIdsDecodedPtr()->choice;
	switch(ueNgapIdsType_m){
		case uE_NGAP_ID_pair_chosen:
		ranNgapId_m=(msgPdu_m.getUeNgapIdsDecodedPtr()->u).uE_NGAP_ID_pair.rAN_UE_NGAP_ID;
		amfNgapId_m=(msgPdu_m.getUeNgapIdsDecodedPtr()->u).uE_NGAP_ID_pair.aMF_UE_NGAP_ID;
		break;
		case aMF_UE_NGAP_ID_chosen:
		amfNgapId_m = (msgPdu_m.getUeNgapIdsDecodedPtr()->u).aMF_UE_NGAP_ID;
		break;

	}
	return true;
}

string UeContextReleaseCommandCodec::toString()
{
	ostringstream oss;

    oss <<"UeContextReleaseCommandCodec:" << endl
		<<"CauseType: "<<causeType_m << endl
		<<"CauseValue: "<<causeValue_m << endl
        <<"UeNgapIdsType: "<<ueNgapIdsType_m << endl
		<<"UeNgapIdsValue:  ranNgapId:  "<<ranNgapId_m <<"  amfNgapId:  "<<amfNgapId_m;

	return oss.str();
}
