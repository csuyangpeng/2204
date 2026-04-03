#include <iostream>
#include <sstream>
#include <string>

#include "pduSessionResourceSetupResponseCodec.h"

using namespace std;

PduSessResSetupRespCodec::PduSessResSetupRespCodec()
{
    reset();
}

PduSessResSetupRespCodec::~PduSessResSetupRespCodec()
{
}

void PduSessResSetupRespCodec::setRanUeNgapId(Uint32 id)
{
    ranUeNgapId_m = id;
}

Uint32
PduSessResSetupRespCodec::getRanUeNgapId()
{
    return ranUeNgapId_m;
}

void PduSessResSetupRespCodec::setAmfUeNgapId(Uint64 id)
{
    amfUeNgapId_m = id;
}

Uint64
PduSessResSetupRespCodec::getAmfUeNgapId()
{
    return amfUeNgapId_m;
}

void 
PduSessResSetupRespCodec::addPduSessResSetupRespList(PduSessResSetupRespItem &item)
{
	pduSessResSetupRespList_m.push_back(item);
}

PduSessResSetupRespList& 
PduSessResSetupRespCodec::getPduSessResSetupRespList()
{
	return pduSessResSetupRespList_m;
}

void 
PduSessResSetupRespCodec::addPduSessResFailedSetupList(PduSessResFailedSetupItem &item)
{
	pduSessResFailedSetupList_m.push_back(item);
}
PduSessResFailedSetupList& 
PduSessResSetupRespCodec::getPduSessResFailedToSetupList()
{
	return pduSessResFailedSetupList_m;
}

//bool PduSessResSetupRespCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer PduSessResSetupRespCodec::encode(NgapOssCtxt *ctxt)
{
	MsgBuffer msgBuffer;
	msgBuffer.length = 0;
	msgBuffer.value = NULL;
	
	// encode RAN UE NGAP ID
	RAN_UE_NGAP_ID *ranUeIdPtr = msgPdu_m.getRanUeNgapIdEncodePtr();
	*ranUeIdPtr = ranUeNgapId_m;
	
	// encode AMF UE NGAP ID
	AMF_UE_NGAP_ID *amfUeIdPtr = msgPdu_m.getAmfUeNgapIdEncodePtr();
	*amfUeIdPtr = amfUeNgapId_m;
	
	// encode pduSessResSetupRespList_m
	int num = pduSessResSetupRespList_m.size();
	msgPdu_m.setNumofSetupListItem_v(num);	
	for (int i = 0; i < num; i++) {
	  PDUSessionResourceSetupItemSURes *pduPtr =
		  msgPdu_m.getPduSessResSetupItemEncodePtr(i);
	  if (NULL == pduPtr) {
		return msgBuffer;
	  }
	  PduSessResSetupRespItem* pduSessItem = &(pduSessResSetupRespList_m[i]);
	  
	  //psi
	  pduPtr->pDUSessionID = pduSessItem->pduSessId;
	  pduPtr->pDUSessionResourceSetupResponseTransfer.encoded.length
	  	= pduSessItem->pduSessResSetupRespTransfer.length();
	  memcpy(pduPtr->pDUSessionResourceSetupResponseTransfer.encoded.value,
	  	     pduSessItem->pduSessResSetupRespTransfer.c_str(),
	  	     pduPtr->pDUSessionResourceSetupResponseTransfer.encoded.length);
	}

	// encode pdu session resource failed to setup list
	num = pduSessResFailedSetupList_m.size();
	msgPdu_m.setNumofFailedSetupItem_v(num);
	for (int i = 0; i < num; i++) {
	  PDUSessionResourceFailedToSetupItemSURes *pduPtr =
		  msgPdu_m.getPduSessResFailedSetupItemEncodePtr(i);
	  if (NULL == pduPtr) {
		return msgBuffer;
	  }
	  PduSessResFailedSetupItem* pduSessItem = &(pduSessResFailedSetupList_m[i]);
	  
	  //psi
	  pduPtr->pDUSessionID = pduSessItem->pduSessId;
	  pduPtr->pDUSessionResourceSetupUnsuccessfulTransfer.encoded.length
	  	= pduSessItem->pduSessResSetupUnSussTransfer.length();
	  memcpy(pduPtr->pDUSessionResourceSetupUnsuccessfulTransfer.encoded.value,
	  	     pduSessItem->pduSessResSetupUnSussTransfer.c_str(),
	  	     pduPtr->pDUSessionResourceSetupUnsuccessfulTransfer.encoded.length);
	}
	
	msgPdu_m.encodeOssToPer(ctxt);
	
	OssBuf *perBuf = msgPdu_m.getPerBufRef();
	
	msgBuffer.length = perBuf->length;
	msgBuffer.value = perBuf->value;
	
	return msgBuffer;

}
bool 
PduSessResSetupRespCodec::decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf)
{
	reset();
	
	OssBuf ossbuf;
	ossbuf.length = msgBuf.length;
	ossbuf.value = msgBuf.value;
	
	msgPdu_m.setPerBufRef_v(ossbuf);
	
	if (0 != msgPdu_m.decodePerToOss(ctxt)) {
		return false;
	}
	
	// decode RAN UE NGAP ID
	ranUeNgapId_m = *(msgPdu_m.getRanUeNgapIdDecodedPtr());
	
	// decode amf ue ngap id
	amfUeNgapId_m = *(msgPdu_m.getAmfUeNgapIdDecodedPtr());
  
	// decode PDUSessionResourceSetupListCxtReq
	for (int i = 0; i < msgPdu_m.getNumofSetupListItem();i++) {
		PDUSessionResourceSetupItemSURes *pduPtr =
			msgPdu_m.getPduSessResSetupItemDecodedPtr(i);
		if (NULL == pduPtr) {
		  return false;
		}
		
		PduSessResSetupRespItem item;
		//psi
		item.pduSessId =pduPtr->pDUSessionID;
		//pdu session resource setup response transfer
		OssBuf outputBuf = ctxt->getSecondaryCodecBuf();
		PDUSessionResourceSetupResponseTransfer* transferPdu;
		transferPdu = pduPtr->pDUSessionResourceSetupResponseTransfer.decoded;
		unsigned int result = 
			ctxt->encodeOssToPer(PDUSessionResourceSetupResponseTransfer_PDU, transferPdu, &outputBuf);
		if (result != 0){
			std::cout << "failed to encode transfer"<< std::endl;
			return false;
		}
		item.pduSessResSetupRespTransfer.assign(
			  (char *)(outputBuf.value),
			  outputBuf.length);	

		pduSessResSetupRespList_m.push_back(item);
	}

	// decode PDUSessionResourceSetupListCxtReq
	for (int i = 0; i < msgPdu_m.getNumofFailedSetupItem();i++) {
		PDUSessionResourceFailedToSetupItemSURes *pduPtr =
			  msgPdu_m.getPduSessResFailedSetupItemDecodedPtr(i);
		if (NULL == pduPtr) {
			return false;
		}
		  
		PduSessResFailedSetupItem item;
		//psi
		item.pduSessId =pduPtr->pDUSessionID;

		//transfer
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
	  
		pduSessResFailedSetupList_m.push_back(item);
	}

    return true;
}

bool 
PduSessResSetupRespCodec::reset()
{
    msgPdu_m.reset_v();

    amfUeNgapId_m = 0;
    ranUeNgapId_m = 0;
	pduSessResSetupRespList_m.clear();
	pduSessResFailedSetupList_m.clear();
}
std::string PduSessResSetupRespCodec::toString()
{
   return "";
}

