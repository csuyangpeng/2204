#include <iostream>
#include <sstream>
#include <string>

#include "errorIndicationCodec.h"

using namespace std;

ErrorIndicationCodec::ErrorIndicationCodec() 
{
    reset();
}

ErrorIndicationCodec::~ErrorIndicationCodec()
{
}

void ErrorIndicationCodec::setRanUeNgapId(Uint32 id) 
{
    ranUeNgapIdPrst_m = true;
    ranUeNgapId_m = id;
}

Uint32 ErrorIndicationCodec::getRanUeNgapId() 
{
    return ranUeNgapId_m;
}

bool ErrorIndicationCodec::isRanUeNgapIdPrst()
{
    return ranUeNgapIdPrst_m;
}

void ErrorIndicationCodec::setAmfUeNgapId(Uint64 id) 
{
    amfUeNgapIdPrst_m = true;
    amfUeNgapId_m = id;
}

Uint64 ErrorIndicationCodec::getAmfUeNgapId() 
{
    return amfUeNgapId_m;
}

bool ErrorIndicationCodec::isAmfUeNgapIdPrst()
{
    return amfUeNgapIdPrst_m;
}

void 
ErrorIndicationCodec::setCauseType(Uint16 cType)
{
    causeTypePrst_m = true;
	causeType_m = cType;
}

Uint16 
ErrorIndicationCodec::getCauseType()
{
	return causeType_m;
}

bool ErrorIndicationCodec::isCauseTypePrst()
{
    return causeTypePrst_m;
}

void 
ErrorIndicationCodec::setCauseValue(Uint16 value)
{
	causeValue_m = value;
}

Uint16 
ErrorIndicationCodec::getCauseValue()
{
	return causeValue_m;
}

bool ErrorIndicationCodec::reset() 
{
    msgPdu_m.reset_v();

    amfUeNgapId_m = 0;
    amfUeNgapIdPrst_m = false;

    ranUeNgapId_m = 0;
    ranUeNgapIdPrst_m = false;

	causeType_m = 0;
    causeTypePrst_m = false;
    causeValue_m = 0;

    return true;
}

//bool ErrorIndicationCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer ErrorIndicationCodec::encode(NgapOssCtxt* ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    // encode RAN UE NGAP ID
    if(ranUeNgapIdPrst_m)
    {
        RAN_UE_NGAP_ID* ranUeIdPtr = msgPdu_m.getRanUeNgapIdEncodePtr();
        *ranUeIdPtr = ranUeNgapId_m;
    }

    // encode AMF UE NGAP ID
    if(amfUeNgapIdPrst_m)
    {
        AMF_UE_NGAP_ID* amfUeIdPtr = msgPdu_m.getAmfUeNgapIdEncodePtr();
        *amfUeIdPtr = amfUeNgapId_m;
    }

   //encode Cause Type and Cause Value
    if(causeTypePrst_m)
    {
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
    }
    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool ErrorIndicationCodec::decode(NgapOssCtxt* ctxt,MsgBuffer msgBuf) 
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
    if(msgPdu_m.isRanUeNgapIdPresent())
    {
        ranUeNgapId_m = *(msgPdu_m.getRanUeNgapIdDecodedPtr());
        ranUeNgapIdPrst_m = true;
    }
    
    // decode amf ue ngap id
    if(msgPdu_m.isAmfUeNgapIdPresent())
    {
        amfUeNgapId_m = *(msgPdu_m.getAmfUeNgapIdDecodedPtr());
        amfUeNgapIdPrst_m = true;
    }

	//decode Cause Type and Cause Value
    if(msgPdu_m.isCausePresent())
    {
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
        causeTypePrst_m = true;
    }
	return true;
}

string ErrorIndicationCodec::toString()
{
	ostringstream oss;

    oss <<"ErrorIndicationCodec:" << endl
        << "1.AMF UE NGAP ID:(" << amfUeNgapId_m << ")" << endl
        << "2.RAN UE NGAP ID:(" << ranUeNgapId_m << ")" << endl
		<<"CauseType: "<<causeType_m
		<<"CauseValue: "<<causeValue_m;

	return oss.str();
}
