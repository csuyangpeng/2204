#include <iostream>
#include <sstream>
#include <string>

#include "ueContextReleaseRequestCodec.h"

using namespace std;

UeContextReleaseRequestCodec::UeContextReleaseRequestCodec() 
{
    reset();
}

UeContextReleaseRequestCodec::~UeContextReleaseRequestCodec() 
{
}

void UeContextReleaseRequestCodec::setRanUeNgapId(Uint32 id) 
{
    ranUeNgapId_m = id;
}

Uint32 UeContextReleaseRequestCodec::getRanUeNgapId() 
{
    return ranUeNgapId_m;
}

void UeContextReleaseRequestCodec::setAmfUeNgapId(Uint64 id) 
{
    amfUeNgapId_m = id;
}

Uint64 UeContextReleaseRequestCodec::getAmfUeNgapId() 
{
    return amfUeNgapId_m;
}

void 
UeContextReleaseRequestCodec::setCauseType(Uint16 cType)
{
	causeType_m = cType;
}

Uint16 
UeContextReleaseRequestCodec::getCauseType()
{
	return causeType_m;
}

void 
UeContextReleaseRequestCodec::setCauseValue(Uint16 value)
{
	causeValue_m = value;
}

Uint16 
UeContextReleaseRequestCodec::getCauseValue()
{
	return causeValue_m;
}

void UeContextReleaseRequestCodec::addPduSessResListRelReq(unsigned short pduSessResRelReq) 
{
	pduSessResListRelReqPrst_m = true;
    pduSessResListRelReq_m.push_back(pduSessResRelReq);
}

std::vector<unsigned short>& UeContextReleaseRequestCodec::getPduSessResListRelReq() 
{
    return pduSessResListRelReq_m;
}

bool 
UeContextReleaseRequestCodec::isPduSessResListRelReqPrst()
{
    return pduSessResListRelReqPrst_m;
}


bool UeContextReleaseRequestCodec::reset() 
{
    msgPdu_m.reset_v();

    amfUeNgapId_m = 0;
    ranUeNgapId_m = 0;
	causeType_m = 0;
    causeValue_m = 0;

	pduSessResListRelReq_m.clear();
	pduSessResListRelReqPrst_m = false;

	return true;
}

//bool UeContextReleaseRequestCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer UeContextReleaseRequestCodec::encode(NgapOssCtxt *ctxt)
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

    //encode PDUSessionResourceListCxtRelReq
    if(pduSessResListRelReqPrst_m)
    {
        Uint32 num = pduSessResListRelReq_m.size();
        msgPdu_m.setNumofPduSessResListItemCxtRelReq_v(num);
    
        for (int i = 0; i < num; i++) 
        {
            PDUSessionResourceItemCxtRelReq *pduPtr = 
                msgPdu_m.getPduSessResItemCxtRelReqEncodePtr(i);
            if (NULL == pduPtr) 
            {
                return msgBuffer;
            }
            pduPtr->pDUSessionID = pduSessResListRelReq_m[i];
        }
    }

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool 
UeContextReleaseRequestCodec::decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf)
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
    // decode RAN UE NGAP ID
    ranUeNgapId_m = *(msgPdu_m.getRanUeNgapIdDecodedPtr());

    // decode amf ue ngap id
    amfUeNgapId_m = *(msgPdu_m.getAmfUeNgapIdDecodedPtr());

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

    //decode PDUSessionResourceListCxtRelReq
    if(msgPdu_m.isPduSessResListCxtRelReqPresent())
    {
        for (int i = 0; i < msgPdu_m.getNumofPduSessResListItemCxtRelReq();i++) 
        {
            PDUSessionResourceItemCxtRelReq *pduPtr =
                msgPdu_m.getPduSessResItemCxtRelReqDecodedPtr(i);
            if (NULL == pduPtr) 
            {
                return false;
            }
    
	    	pduSessResListRelReq_m.push_back(pduPtr->pDUSessionID);
        }
		pduSessResListRelReqPrst_m = true;
    }

	return true;
}

string UeContextReleaseRequestCodec::toString()
{
	ostringstream oss;

    oss <<"UeContextReleaseRequestCodec:" << endl
        << "AMF UE NGAP ID:(" << amfUeNgapId_m << ")" << endl
        << "RAN UE NGAP ID:(" << ranUeNgapId_m << ")" << endl
		<<"CauseType: "<<causeType_m << endl
		<<"CauseValue: "<<causeValue_m <<endl 
		<< "PDU Session Resource Release Request List:Prst" <<pduSessResListRelReqPrst_m <<endl;
		
	return oss.str();
}
