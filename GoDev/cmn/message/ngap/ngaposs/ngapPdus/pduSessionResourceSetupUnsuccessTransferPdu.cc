#include <string.h>
#include <iostream>

#include "pduSessionResourceSetupUnsuccessTransferPdu.h"

PduSessResSetupUnsuccTransferPdu::PduSessResSetupUnsuccTransferPdu()
{
    reset_v();
}

PduSessResSetupUnsuccTransferPdu::~PduSessResSetupUnsuccTransferPdu()
{
}

void PduSessResSetupUnsuccTransferPdu::reset_v()
{
	cause_m.choice = misc_chosen;
	cause_m.u.misc = CauseMisc_unspecified;
}

unsigned int 
PduSessResSetupUnsuccTransferPdu::decodePerToOss(NgapOssCtxt* ctxt)
{
   if(NULL == ctxt)
   {
       return -1;
   }
   
   OssBuf inputBuf;
   OssBuf outputBuf;
   int pduNum = PDUSessionResourceSetupUnsuccessfulTransfer_PDU;

   inputBuf = perBufRef_m;

   ctxt->reset_v();
   outputBuf = ctxt->getPrimaryCodecBuf();

   // default action -- automatically fully decode the message
   unsigned int result = ctxt->decodePerToOss(&pduNum, &inputBuf, &outputBuf);

   if (result == 0)
   {
      ossBufRef_m = outputBuf;
      getDecodedOssDataPtrs_v();
   }

   return result;
}

unsigned int
PduSessResSetupUnsuccTransferPdu::encodeOssToPer(NgapOssCtxt *ctxt)
{
    if (NULL == ctxt){
        return -1;
    }

	PDUSessionResourceSetupUnsuccessfulTransfer ngApPdu;
	ngApPdu.bit_mask = 0;
	ngApPdu.iE_Extensions = NULL;
	memcpy(&ngApPdu.cause, &cause_m, sizeof(cause_m));
	
    //Convert Structures to PER
    OssBuf outputBuf;
    unsigned int result;

    ctxt->reset_v();
    outputBuf = ctxt->getPrimaryCodecBuf();

    result = ctxt->encodeOssToPer(PDUSessionResourceSetupUnsuccessfulTransfer_PDU, &ngApPdu, &outputBuf);

    if (result == encodeSuccessCode_c){
        perBufRef_m = outputBuf;
    }

    return result;
}

void PduSessResSetupUnsuccTransferPdu::getDecodedOssDataPtrs_v()
{
    reset_v();

	PDUSessionResourceSetupUnsuccessfulTransfer* ptr =  
		reinterpret_cast<PDUSessionResourceSetupUnsuccessfulTransfer *>(ossBufRef_m.value);

    if (ptr == NULL){
        return;
    }
    cause_mp = &ptr->cause;

}

Cause *
PduSessResSetupUnsuccTransferPdu::getCauseEncodePtr()
{
    return &cause_m;
}

Cause *
PduSessResSetupUnsuccTransferPdu::getCauseDecodePtr()
{
    return cause_mp;
}

