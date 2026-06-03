#include <iostream>
#include <sstream>
#include <string>

#include "pduSessionResourceModifyRequestCodec.h"

using namespace std;

PduSessResModReqCodec::PduSessResModReqCodec()
{
    reset();
}

PduSessResModReqCodec::~PduSessResModReqCodec()
{

}

void PduSessResModReqCodec::setRanUeNgapId(Uint32 id) {
  ranUeNgapId_m = id;
}

Uint32 PduSessResModReqCodec::getRanUeNgapId() {
  return ranUeNgapId_m;
}

void PduSessResModReqCodec::setAmfUeNgapId(Uint64 id) {
  amfUeNgapId_m = id;
}

Uint64 PduSessResModReqCodec::getAmfUeNgapId() {
  return amfUeNgapId_m;
}

void PduSessResModReqCodec::setRanPagingPriority(    Uint16 val) {
  ranPagingPriorityPresent_m = true;
  ranPagingPriority_m = val;
}

Uint16 PduSessResModReqCodec::getRanPagingPriority() {
  return ranPagingPriority_m;
}

bool PduSessResModReqCodec::isRanPagingPriorityPresent() {
  return ranPagingPriorityPresent_m;
}

void PduSessResModReqCodec::addPduSessResModReqList(PduSessResModReqItem& item)
{
  pduSessResModReqList_m.push_back(item);
}

 PduSessResModReqList& 
 PduSessResModReqCodec::getPduSessResModReqList() {
  return pduSessResModReqList_m;
}

bool PduSessResModReqCodec::reset() {
  msgPdu_m.reset_v();

  amfUeNgapId_m = 0;
  ranUeNgapId_m = 0;

  ranPagingPriority_m = 0;
  ranPagingPriorityPresent_m = false;

  pduSessResModReqList_m.clear();
}

MsgBuffer PduSessResModReqCodec::encode(NgapOssCtxt* ctxt)
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

  // encode ran paging priority
  if (ranPagingPriorityPresent_m) {
    *(msgPdu_m.getRANPagingPriorityEncodePtr()) = ranPagingPriority_m;
  }

  // encode PDUSessionResourceSetupListCxtReq
  int num = pduSessResModReqList_m.size();
  msgPdu_m.setNumofPduSessResModifyItemModReq_v(num);
  
  for (int i = 0; i < num; i++) {
    PDUSessionResourceModifyItemModReq *pduPtr =
        msgPdu_m.getPduSessResModifyItemModReqEncodePtr(i);
    if (NULL == pduPtr) {
      return msgBuffer;
    }

    PduSessResModReqItem* pduSessItem = &(pduSessResModReqList_m[i]);

    // psi
    pduPtr->pDUSessionID = pduSessItem->pduSessId;

    // nas
    if (pduSessItem->nasPduPrst) {
      pduPtr->bit_mask |= PDUSessionResourceModifyItemModReq_nAS_PDU_present;
      pduPtr->nAS_PDU.length =pduSessItem->nasPdu.length();
      memcpy(pduPtr->nAS_PDU.value,
             pduSessItem->nasPdu.c_str(),
             pduPtr->nAS_PDU.length);
    }

    // snssai
    if (pduSessItem->snssaiPrst){
//        pduPtr->bit_mask |= PDUSessionResourceModifyItemModReq_iE_Extensions_present;
//
//        S_NSSAI* SnssaiPtr = msgPdu_m.getSnssaiEncodePtr();
//        SnssaiPtr->sST.length = sizeof(pduSessItem->snssai.sst);
//        memcpy(SnssaiPtr->sST.value,
//            pduSessItem->snssai.sst,
//            SnssaiPtr->sST.length);
//
//        if (pduSessItem->snssai.sdPresent) {
//            SnssaiPtr->bit_mask |= sD_present;
//            SnssaiPtr->sD.length =sizeof(pduSessItem->snssai.sd);
//            memcpy(SnssaiPtr->sD.value,
//                pduSessItem->snssai.sd,
//                SnssaiPtr->sD.length);
//           }
     }

    //pdu session resource setup transfer
    pduPtr->pDUSessionResourceModifyRequestTransfer.encoded.length 
      = pduSessItem->pduSessResModReqTransfer.length();
    memcpy(pduPtr->pDUSessionResourceModifyRequestTransfer.encoded.value,
           pduSessItem->pduSessResModReqTransfer.c_str(),
           pduPtr->pDUSessionResourceModifyRequestTransfer.encoded.length); 
  }
  
  msgPdu_m.encodeOssToPer(ctxt);

  OssBuf *perBuf = msgPdu_m.getPerBufRef();

  msgBuffer.length = perBuf->length;
  msgBuffer.value = perBuf->value;

  return msgBuffer;
}

bool PduSessResModReqCodec::decode(NgapOssCtxt* ctxt, MsgBuffer msgBuf)
{
  reset();
        
  OssBuf ossbuf;
  ossbuf.length = msgBuf.length;
  ossbuf.value = msgBuf.value;
        
  msgPdu_m.setPerBufRef_v(ossbuf);
  std::cout <<"point test" << std::endl;
  if (0 != msgPdu_m.decodePerToOss(ctxt)) {
     return false;
  }

  std::cout <<"test point" <<std::endl;
        
  // decode RAN UE NGAP ID
  ranUeNgapId_m = *(msgPdu_m.getRanUeNgapIdDecodedPtr());
        
  // decode amf ue ngap id
  amfUeNgapId_m = *(msgPdu_m.getAmfUeNgapIdDecodedPtr());
        
  // decode ran paging priority
  if (msgPdu_m.isRANPagingPriorityPresent()) {
     ranPagingPriority_m = *msgPdu_m.getRANPagingPriorityDecodedPtr();
     ranPagingPriorityPresent_m = true;
  }
          
  // decode PDUSessionResourceSetupListCxtReq
  for (int i = 0; i < msgPdu_m.getNumofPduSessResModifyItemModReq();i++) {
     PDUSessionResourceModifyItemModReq *pduPtr =
                msgPdu_m.getPduSessResModifyItemModReqDecodedPtr(i);
     if (NULL == pduPtr) {
        return false;
     }
                
     PduSessResModReqItem pduSessModItem_m;
     //psi
     pduSessModItem_m.pduSessId =pduPtr->pDUSessionID;
     //nas
     if (pduPtr->bit_mask & PDUSessionResourceModifyItemModReq_nAS_PDU_present) {
        pduSessModItem_m.nasPduPrst = true;
        pduSessModItem_m.nasPdu.assign(
                  (char *)(pduPtr->nAS_PDU.value),
                  pduPtr->nAS_PDU.length);
     }
     //snssai
     if (pduPtr->bit_mask & PDUSessionResourceModifyItemModReq_iE_Extensions_present) {
//         SnssaiIeValue* SnssaiPtr = msgPdu_m.getIeExtenIeTypeDecodedPtr();
//
//         memcpy(pduSessModItem_m.snssai.sst,
//             SnssaiPtr->value.extensionValue.decoded.pdu_S_NSSAI->sST.value,
//             sizeof(SnssaiPtr->value.extensionValue.decoded.pdu_S_NSSAI->sST.value));
//         if (SnssaiPtr->value.extensionValue.decoded.pdu_S_NSSAI->bit_mask | sD_present) {
//             pduSessModItem_m.snssai.sdPresent = true;
//             memcpy(pduSessModItem_m.snssai.sd,
//             SnssaiPtr->value.extensionValue.decoded.pdu_S_NSSAI->sD.value,
//             sizeof(SnssaiPtr->value.extensionValue.decoded.pdu_S_NSSAI->sD.value));
//         }
     }

     //pdu session resource setup request transfer
     OssBuf outputBuf = ctxt->getSecondaryCodecBuf();
     PDUSessionResourceModifyRequestTransfer* transferPdu;
     transferPdu = pduPtr->pDUSessionResourceModifyRequestTransfer.decoded;
     unsigned int result = 
        ctxt->encodeOssToPer(PDUSessionResourceModifyRequestTransfer_PDU, transferPdu, &outputBuf);
     if (result != 0){
        std::cout << "failed to encode transfer"<< std::endl;
        return false;
     }

     pduSessModItem_m.pduSessResModReqTransfer.assign(
                  (char *)(outputBuf.value),
                  outputBuf.length);    
     pduSessResModReqList_m.push_back(pduSessModItem_m);
  }

  return true;
}

string PduSessResModReqCodec::toString() {

  ostringstream oss;

  oss << "PduSessResModifyReqCodec:" << endl
      << "AMF UE NGAP ID:(" << amfUeNgapId_m << ")" << endl
      << "RAN UE NGAP ID:(" << ranUeNgapId_m << ")" << endl
      << "Ran Paging Priority:Flag(" << ranPagingPriorityPresent_m << "),value("
      << ranPagingPriority_m << ")" << endl
      << "PDU Session Resource Modify Request List:" << endl;

  return oss.str();
}



