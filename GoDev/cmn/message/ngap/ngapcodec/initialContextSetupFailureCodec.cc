#include <iostream>
#include <sstream>
#include <string>

#include "initialContextSetupFailureCodec.h"

using namespace std;

InitialContextSetupFailureCodec::InitialContextSetupFailureCodec() 
{
    reset();
}

InitialContextSetupFailureCodec::~InitialContextSetupFailureCodec()
{
}

void 
InitialContextSetupFailureCodec::setRanUeNgapId(Uint32 id) 
{
    ranUeNgapId_m = id;
}

Uint32 
InitialContextSetupFailureCodec::getRanUeNgapId() 
{
    return ranUeNgapId_m;
}

void 
InitialContextSetupFailureCodec::setAmfUeNgapId(Uint64 id) 
{
    amfUeNgapId_m = id;
}

Uint64 
InitialContextSetupFailureCodec::getAmfUeNgapId() 
{
    return amfUeNgapId_m;
}

void 
InitialContextSetupFailureCodec::setCauseType(Uint16 cType)
{
	causeType_m = cType;
}

Uint16 
InitialContextSetupFailureCodec::getCauseType()
{
	return causeType_m;
}

void 
InitialContextSetupFailureCodec::setCauseValue(Uint16 value)
{
	causeValue_m = value;
}

Uint16 
InitialContextSetupFailureCodec::getCauseValue()
{
	return causeValue_m;
}

void 
InitialContextSetupFailureCodec::addPduSessResFailToSetupList(PduSessResFailedSetupItem& val) 
{
    pduSessResFailToSetupList_m.push_back(val);
}

PduSessResFailedSetupList& 
InitialContextSetupFailureCodec::getPduSessResFailToSetupList() 
{
    return pduSessResFailToSetupList_m;
}


bool 
InitialContextSetupFailureCodec::reset() 
{
    msgPdu_m.reset_v();

    amfUeNgapId_m = 0;
    ranUeNgapId_m = 0;

    causeType_m = 0;
    causeValue_m = 0;

    pduSessResFailToSetupList_m.clear();
}

MsgBuffer InitialContextSetupFailureCodec::encode(NgapOssCtxt* ctxt) 
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

    //encode PDUSessionResourceFailedToSetupListCxtFail 
    Uint32 num = pduSessResFailToSetupList_m.size();
    msgPdu_m.setNumofPDUSessionResourceSetupListItemCxtFail_v(num);

    for (int i = 0; i < num; i++) {
        PDUSessionResourceFailedToSetupItemCxtFail *pduPtr = 
            msgPdu_m.getPDUSessionResourceSetupItemCxtFailEncodePtr(i);
        if (NULL == pduPtr) {
	    std::cout <<"invalid pduPtr"<< std::endl;
            return msgBuffer;
        }

		PduSessResFailedSetupItem *itemPtr = &(pduSessResFailToSetupList_m[i]);
		
        pduPtr->pDUSessionID = itemPtr->pduSessId;

	    //pdu session resource fail to setup transfer
        pduPtr->pDUSessionResourceSetupUnsuccessfulTransfer.encoded.length 
          = itemPtr->pduSessResSetupUnSussTransfer.length();
        memcpy(pduPtr->pDUSessionResourceSetupUnsuccessfulTransfer.encoded.value,
               itemPtr->pduSessResSetupUnSussTransfer.c_str(),
               itemPtr->pduSessResSetupUnSussTransfer.length());	
    }

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool InitialContextSetupFailureCodec::decode(NgapOssCtxt* ctxt,MsgBuffer msgBuf) 
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

    // decode Cause Type and Cause Value
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
	
	//decode PDUSessionResourceFailedToSetupListCxtFail
	for (int i = 0; i < msgPdu_m.getNumofPDUSessionResourceSetupListItemCxtFail();i++) {
		PDUSessionResourceFailedToSetupItemCxtFail *pduPtr =
			msgPdu_m.getPDUSessionResourceSetupItemCxtFailDecodedPtr(i);
		if (NULL == pduPtr) {
			return false;
		}
		  
		PduSessResFailedSetupItem item;
		item.pduSessId = pduPtr->pDUSessionID;
	
		//pdu session resource fail to setup transfer
		OssBuf outputBuf = ctxt->getSecondaryCodecBuf();
		PDUSessionResourceSetupUnsuccessfulTransfer* transferPdu;
		transferPdu = pduPtr->pDUSessionResourceSetupUnsuccessfulTransfer.decoded;
		unsigned int result = 
			ctxt->encodeOssToPer(PDUSessionResourceSetupUnsuccessfulTransfer_PDU, transferPdu, &outputBuf);
		if (result != 0){
			std::cout << "failed to encode transfer"<< std::endl;
			return false;
		}
		item.pduSessResSetupUnSussTransfer.assign(
			(char *)(outputBuf.value),
			outputBuf.length);	  
	
		pduSessResFailToSetupList_m.push_back(item);
	}

  return true;
}

string InitialContextSetupFailureCodec::toString(){
   ostringstream oss;

   oss << "InitialContextSetupFailureCodec:" << endl
       << "1.AMF UE NGAP ID:(" << amfUeNgapId_m << ")" << endl
       << "2.RAN UE NGAP ID:(" << ranUeNgapId_m << ")" << endl
       <<"3.CauseType:"<<causeType_m << endl
	   <<"4.CauseValue:"<<causeValue_m << endl
       <<"5.PDU Session Resource Setup Request List" << endl;
    vector<PduSessResFailedSetupItem>::iterator it;
    for (it=pduSessResFailToSetupList_m.begin(); it<pduSessResFailToSetupList_m.end() ; it++)
	{
		oss<<"PSI:" << int(it->pduSessId) << endl;
	}
   return oss.str(); 
}