#include <iostream>
#include <sstream>
#include <string>

#include "pduSessionResourceReleaseCommandCodec.h"

using namespace std;

PduSessResRelCmdCodec::PduSessResRelCmdCodec() 
{
    reset();
}

PduSessResRelCmdCodec::~PduSessResRelCmdCodec()
{
}

void PduSessResRelCmdCodec::setRanUeNgapId(Uint32 id) 
{
    ranUeNgapId_m = id;
}

Uint32 PduSessResRelCmdCodec::getRanUeNgapId() 
{
    return ranUeNgapId_m;
}

void PduSessResRelCmdCodec::setAmfUeNgapId(Uint64 id) 
{
    amfUeNgapId_m = id;
}

Uint64 PduSessResRelCmdCodec::getAmfUeNgapId() 
{
    return amfUeNgapId_m;
}

void PduSessResRelCmdCodec::setNasPdu(std::string naspdu) 
{
    nasPduPrst_m=true;
    nasPdu_m = naspdu;
}

std::string PduSessResRelCmdCodec::getNasPdu() 
{ 
    return nasPdu_m; 
}

bool PduSessResRelCmdCodec::isNasPduPrst()
{
    return nasPduPrst_m;
}

void 
PduSessResRelCmdCodec::setRanPagingPriority(Uint16 ranPagingPrt)
{
   ranPagingPriorityPresent_m = true;
   ranPagingPriority_m = ranPagingPrt;
}

Uint16
PduSessResRelCmdCodec::getRanPagingPriority()
{
  return ranPagingPriority_m;
}

bool 
PduSessResRelCmdCodec::isRanPagingPriorityPresent()
{
  return ranPagingPriority_m;
}

void PduSessResRelCmdCodec::addPduSessResRelCmdList(PduSessResRelCmdItem& val) 
{
    pduSessResRelCmdList_m.push_back(val);
}

PduSessResRelCmdList& PduSessResRelCmdCodec::getPduSessResRelCmdList() 
{
    return pduSessResRelCmdList_m;
}

bool PduSessResRelCmdCodec::reset() 
{
    msgPdu_m.reset_v();

    amfUeNgapId_m = 0;
    ranUeNgapId_m = 0;
    nasPdu_m.clear();
    nasPduPrst_m=false;

    ranPagingPriority_m = 0;
    ranPagingPriorityPresent_m = false;

    pduSessResRelCmdList_m.clear();
}

//bool PduSessResRelCmdCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer PduSessResRelCmdCodec::encode(NgapOssCtxt* ctxt)
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

    // encode NAS PDU
    if(nasPduPrst_m){
        NAS_PDU* nasPduPtr = msgPdu_m.getNasPduEncodePtr();
        nasPduPtr->length = nasPdu_m.length();
        memcpy(nasPduPtr->value, nasPdu_m.c_str(), nasPduPtr->length);
    }

    //encode ran paging priority
    if(ranPagingPriorityPresent_m)
    {
        RANPagingPriority* ranPagingPriorityPtr = msgPdu_m.getRANPagingPriorityEncodePtr();
        *ranPagingPriorityPtr = ranPagingPriority_m;
    }

    //encode PDUSessionResourceReleaseListCxtRelCmd 
    Uint32 num = pduSessResRelCmdList_m.size();
    msgPdu_m.setNumofPduSessResToRelItemRelCmd_v(num);

    for (int i = 0; i < num; i++) {
        PDUSessionResourceToReleaseItemRelCmd *pduPtr = 
            msgPdu_m.getPduSessResToRelItemRelCmdEncodePtr(i);
        if (NULL == pduPtr) {
            return msgBuffer;
        }

		PduSessResRelCmdItem *itemPtr = &(pduSessResRelCmdList_m[i]);
		
        pduPtr->pDUSessionID = itemPtr->pduSessId;
//cout <<"debug encode transfer~~~~~~~~~~~~~~~~~~~~~" <<endl;
	    //pdu session resource release command transfer
        pduPtr->pDUSessionResourceReleaseCommandTransfer.encoded.length 
          = itemPtr->pduSessResRelCmdTransfer.length();
        memcpy(pduPtr->pDUSessionResourceReleaseCommandTransfer.encoded.value,
               itemPtr->pduSessResRelCmdTransfer.c_str(),
               itemPtr->pduSessResRelCmdTransfer.length());	
    }

    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf* perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool PduSessResRelCmdCodec::decode(NgapOssCtxt* ctxt,MsgBuffer msgBuf) 
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

    // decode NAS PDU
    if(msgPdu_m.isNasPduPresent()) {
        NAS_PDU* nasPduPtr = msgPdu_m.getNasPduDecodedPtr();
        nasPdu_m.assign((char*)(nasPduPtr->value), nasPduPtr->length);
        nasPduPrst_m=true;
    }

    //decode ran paging priority
    if(msgPdu_m.isRANPagingPriorityPresent())
    {
       RANPagingPriority* ranPagingPriorityPtr = msgPdu_m.getRANPagingPriorityDecodedPtr();
       ranPagingPriority_m = *ranPagingPriorityPtr;
       ranPagingPriorityPresent_m = true;
    }

	//decode PDUSessionResourceReleaseListCxtRelCmd 
	for (int i = 0; i < msgPdu_m.getNumofPduSessResToRelItemRelCmd();i++) {
		PDUSessionResourceToReleaseItemRelCmd *pduPtr =
			msgPdu_m.getPduSessResToRelItemRelCmdDecodedPtr(i);
		if (NULL == pduPtr) {
			return false;
		}
		  
		PduSessResRelCmdItem item;
		item.pduSessId = pduPtr->pDUSessionID;
	
		//pdu session resource release command transfer
		OssBuf outputBuf = ctxt->getSecondaryCodecBuf();
		PDUSessionResourceReleaseCommandTransfer* transferPdu;
		transferPdu = pduPtr->pDUSessionResourceReleaseCommandTransfer.decoded;
		unsigned int result = 
			ctxt->encodeOssToPer(PDUSessionResourceReleaseCommandTransfer_PDU, transferPdu, &outputBuf);
		if (result != 0){
			std::cout << "failed to encode transfer"<< std::endl;
			return false;
		}
		item.pduSessResRelCmdTransfer.assign(
			(char *)(outputBuf.value),
			outputBuf.length);	  
	
		pduSessResRelCmdList_m.push_back(item);
	}
  return true;
}

string PduSessResRelCmdCodec::toString(){
   ostringstream oss;

   oss << "PduSessResRelCmdCodec:" << endl
       << "1.AMF UE NGAP ID:(" << amfUeNgapId_m << ")" << endl
       << "2.RAN UE NGAP ID:(" << ranUeNgapId_m << ")" << endl
       << "3.NAS PDU:(" << nasPdu_m << ")" << endl
       << "4.RanPagingPriority, prst(" << ranPagingPriorityPresent_m 
       << ") value(" << ranPagingPriority_m <<")" << endl
       << "PDU Session Resource Release Command List:" << endl;

    vector<PduSessResRelCmdItem>::iterator it;
    for(it = pduSessResRelCmdList_m.begin(); it<pduSessResRelCmdList_m.end(); it++)
    {
      oss<< "psi:" << it->pduSessId << endl;
    }
   return oss.str(); 
}