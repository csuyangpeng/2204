#include <iostream>
#include <sstream>
#include <string>

#include "pduSessionResourceSetupRequestCodec.h"

using namespace std;

PduSessResSetupReqCodec::PduSessResSetupReqCodec() {
  reset();
}

PduSessResSetupReqCodec::~PduSessResSetupReqCodec() {}

void PduSessResSetupReqCodec::setRanUeNgapId(Uint32 id) {
  ranUeNgapId_m = id;
}

Uint32 PduSessResSetupReqCodec::getRanUeNgapId() {
  return ranUeNgapId_m;
}

void PduSessResSetupReqCodec::setAmfUeNgapId(Uint64 id) {
  amfUeNgapId_m = id;
}

Uint64 PduSessResSetupReqCodec::getAmfUeNgapId() {
  return amfUeNgapId_m;
}

void PduSessResSetupReqCodec::setNasPdu(std::string naspdu) {
  nasPduPrst_m = true;
  nasPdu_m = naspdu;
}

std::string PduSessResSetupReqCodec::getNasPdu() {
  return nasPdu_m;
}

bool PduSessResSetupReqCodec::isNasPduPrst() {
  return nasPduPrst_m;
}

void PduSessResSetupReqCodec::setRanPagingPriority(    Uint16 val) {
  ranPagingPriPrst_m = true;
  ranPagingPriority_m = val;
}

Uint16 PduSessResSetupReqCodec::getRanPagingPriority() {
  return ranPagingPriority_m;
}

bool PduSessResSetupReqCodec::isRanPagPriPrst() {
  return ranPagingPriPrst_m;
}

void PduSessResSetupReqCodec::setUeAmbr(UeAmbr& ambr) {
    ueAMBRPrst_m=true;
    ueAmbr_m = ambr;
}
UeAmbr& PduSessResSetupReqCodec::getUeAmbr() {
    return ueAmbr_m;
}
bool PduSessResSetupReqCodec::isUeAMBRPrst(){
    return ueAMBRPrst_m;
}

void PduSessResSetupReqCodec::addPduSessResSetupReqList(PduSessResSetupReqItem& item)
{
  pduSessResSetupReqList_m.push_back(item);
}

 PduSessResSetupReqList& 
 PduSessResSetupReqCodec::getPduSessResSetupReqList() {
  return pduSessResSetupReqList_m;
}

bool PduSessResSetupReqCodec::reset() {
  msgPdu_m.reset_v();

  amfUeNgapId_m = 0;
  ranUeNgapId_m = 0;

  nasPdu_m.clear();
  nasPduPrst_m = false;

  memset(&ueAmbr_m, 0, sizeof(ueAmbr_m));
  ueAMBRPrst_m=false;

  ranPagingPriority_m = 0;
  ranPagingPriPrst_m = false;

  pduSessResSetupReqList_m.clear();
}

//bool PduSessResSetupReqCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer PduSessResSetupReqCodec::encode(NgapOssCtxt* ctxt)
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

  // encode NAS PDU
  if (nasPduPrst_m) {
    NAS_PDU *nasPduPtr = msgPdu_m.getNasPduEncodePtr();
    nasPduPtr->length = nasPdu_m.length();
    memcpy(nasPduPtr->value, nasPdu_m.c_str(), nasPduPtr->length);
  }

  // encode Ue ambr
  if(ueAMBRPrst_m){
      UEAggregateMaximumBitRate* ueAmbrPtr =
          msgPdu_m.getUeAggregateMaximumBitRateEncodePtr();
      ueAmbrPtr->uEAggregateMaximumBitRateDL = ueAmbr_m.downlink;
      ueAmbrPtr->uEAggregateMaximumBitRateUL = ueAmbr_m.uplink;
  }

  // encode ran paging priority
  if (ranPagingPriPrst_m) {
    *(msgPdu_m.getRanPagePriorityEncodePtr()) = ranPagingPriority_m;
  }

  // encode PDUSessionResourceSetupListCxtReq
  int num = pduSessResSetupReqList_m.size();
  msgPdu_m.setNumofPduSessResSetupListSUReqItem_v(num);
  
  for (int i = 0; i < num; i++) {
    PDUSessionResourceSetupItemSUReq *pduPtr =
        msgPdu_m.getPduSessResSetupSUReqItemEncodePtr(i);
    if (NULL == pduPtr) {
      return msgBuffer;
    }
	PduSessResSetupReqItem* pduSessItem = &(pduSessResSetupReqList_m[i]);
	//psi
    pduPtr->pDUSessionID = pduSessItem->pduSessId;
	//nas
    if (pduSessItem->nasPduPrst) {
      pduPtr->bit_mask |= pDUSessionNAS_PDU_present;
      pduPtr->pDUSessionNAS_PDU.length =pduSessItem->nasPdu.length();
      memcpy(pduPtr->pDUSessionNAS_PDU.value,
             pduSessItem->nasPdu.c_str(),
             pduPtr->pDUSessionNAS_PDU.length);
    }

    // snssai
    pduPtr->s_NSSAI.sST.length =sizeof(pduSessItem->snssai.sst);
    memcpy(pduPtr->s_NSSAI.sST.value,
           pduSessItem->snssai.sst,
           pduPtr->s_NSSAI.sST.length);

    if (pduSessItem->snssai.sdPresent) {
      pduPtr->s_NSSAI.bit_mask |= sD_present;
      pduPtr->s_NSSAI.sD.length =sizeof(pduSessItem->snssai.sd);
      memcpy(pduPtr->s_NSSAI.sD.value,
             pduSessItem->snssai.sd,
             pduPtr->s_NSSAI.sD.length);
    }
	//pdu session resource setup transfer
    pduPtr->pDUSessionResourceSetupRequestTransfer.encoded.length 
      = pduSessItem->pduSessResSetupReqTransfer.length();
    memcpy(pduPtr->pDUSessionResourceSetupRequestTransfer.encoded.value,
           pduSessItem->pduSessResSetupReqTransfer.c_str(),
           pduPtr->pDUSessionResourceSetupRequestTransfer.encoded.length);	
  }
  
  msgPdu_m.encodeOssToPer(ctxt);

  OssBuf *perBuf = msgPdu_m.getPerBufRef();

  msgBuffer.length = perBuf->length;
  msgBuffer.value = perBuf->value;

  return msgBuffer;
}

bool PduSessResSetupReqCodec::decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf) {
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

  // decode NAS PDU
  if (msgPdu_m.isNasPduPresent()) {
    NAS_PDU *nasPduPtr = msgPdu_m.getNasPduDecodedPtr();
    nasPdu_m.assign((char *)(nasPduPtr->value), nasPduPtr->length);
    nasPduPrst_m = true;
  }

  // decode Ue Ambr
  if(msgPdu_m.isUeAMBRPresent()) {
      ueAmbr_m.uplink =
          msgPdu_m.getUeAggregateMaximumBitRateDecodedPtr()->uEAggregateMaximumBitRateUL;
      ueAmbr_m.downlink =
          msgPdu_m.getUeAggregateMaximumBitRateDecodedPtr()->uEAggregateMaximumBitRateDL;
      ueAMBRPrst_m=true;
  }

  // decode ran paging priority
  if (msgPdu_m.isRanPagePriorityPrst()) {
    ranPagingPriority_m = *msgPdu_m.getRanPagePriorityDecodedPtr();
    ranPagingPriPrst_m = true;
  }
  
  // decode PDUSessionResourceSetupListCxtReq
  for (int i = 0; i < msgPdu_m.getNumofPduSessResSetupListSUReqItem();i++) {
    PDUSessionResourceSetupItemSUReq *pduPtr =
        msgPdu_m.getPduSessResSetupSUReqItemDecodedPtr(i);
    if (NULL == pduPtr) {
      return false;
    }
	
     PduSessResSetupReqItem	 pduSessResItem_m;
    //psi
    pduSessResItem_m.pduSessId =pduPtr->pDUSessionID;
    //nas
    if (pduPtr->bit_mask ==pDUSessionNAS_PDU_present) {
      pduSessResItem_m.nasPduPrst = true;
      pduSessResItem_m.nasPdu.assign(
          (char *)(pduPtr->pDUSessionNAS_PDU.value),
          pduPtr->pDUSessionNAS_PDU.length);
    }
    //snssai
    memcpy(pduSessResItem_m.snssai.sst,
           pduPtr->s_NSSAI.sST.value,
           sizeof(pduPtr->s_NSSAI.sST.value));
    if (pduPtr->s_NSSAI.bit_mask == sD_present) {
      pduSessResItem_m.snssai.sdPresent = true;
      memcpy(pduSessResItem_m.snssai.sd,
             pduPtr->s_NSSAI.sD.value,
             sizeof(pduPtr->s_NSSAI.sD.value));
    }
	//pdu session resource setup request transfer
    
   

	OssBuf outputBuf = ctxt->getSecondaryCodecBuf();
    PDUSessionResourceSetupRequestTransfer* transferPdu;
	transferPdu = pduPtr->pDUSessionResourceSetupRequestTransfer.decoded;
    unsigned int result = 
		ctxt->encodeOssToPer(PDUSessionResourceSetupRequestTransfer_PDU, transferPdu, &outputBuf);
	if (result != 0){
		std::cout << "failed to encode transfer"<< std::endl;
		return false;
	}
    pduSessResItem_m.pduSessResSetupReqTransfer.assign(
          (char *)(outputBuf.value),
          outputBuf.length);	
	
    pduSessResSetupReqList_m.push_back(pduSessResItem_m);
  }
  
  return true;
}

string PduSessResSetupReqCodec::toString() {

  ostringstream oss;

  oss << "PduSessResSetupReqCodec:" << endl
      << "AMF UE NGAP ID:(" << amfUeNgapId_m << ")" << endl
      << "RAN UE NGAP ID:(" << ranUeNgapId_m << ")" << endl
      << "NAS PDU:Flag(" << nasPduPrst_m << "),value(" << nasPdu_m << ")" << endl
      << "Ran Paging Priority:Flag(" << ranPagingPriPrst_m << "),value("
      << ranPagingPriority_m << ")" << endl
      << "4.UeAmbr:(flag:"<< ueAMBRPrst_m << ", downlink - " << ueAmbr_m.downlink << ", uplink - "
      << ueAmbr_m.uplink << ")" << endl
      << "PDU Session Resource Setup Request List:" << endl;

  return oss.str();
}