#include <string.h>
#include <iostream>

#include "pduSessionResourceModifyUnSuccTransferPdu.h"

PduSessResMdfyUnsuccTransferPdu::PduSessResMdfyUnsuccTransferPdu()
{
    reset_v();
}

PduSessResMdfyUnsuccTransferPdu::~PduSessResMdfyUnsuccTransferPdu()
{
}

void PduSessResMdfyUnsuccTransferPdu::reset_v()
{
	cause_m.choice = misc_chosen;
	cause_m.u.misc = CauseMisc_unspecified;
}

unsigned int
PduSessResMdfyUnsuccTransferPdu::decodePerToOss(NgapOssCtxt* ctxt)
{
   if(NULL == ctxt)
   {
       return -1;
   }

   OssBuf inputBuf;
   OssBuf outputBuf;
   int pduNum = PDUSessionResourceModifyUnsuccessfulTransfer_PDU;

   inputBuf = perBufRef_m;

   ctxt->reset_v();
   outputBuf = ctxt->getPrimaryCodecBuf();

   unsigned int result = ctxt->decodePerToOss(&pduNum, &inputBuf, &outputBuf);

   if (result == 0)
   {
      ossBufRef_m = outputBuf;
      getDecodedOssDataPtrs_v();
   }

   return result;
}

unsigned int
PduSessResMdfyUnsuccTransferPdu::encodeOssToPer(NgapOssCtxt *ctxt)
{
    if (NULL == ctxt){
        return -1;
    }

	PDUSessionResourceModifyUnsuccessfulTransfer ngApPdu;
	ngApPdu.bit_mask = 0;
	ngApPdu.iE_Extensions = NULL;
	memcpy(&ngApPdu.cause, &cause_m, sizeof(cause_m));

    //Convert Structures to PER
    OssBuf outputBuf;
    unsigned int result;

    ctxt->reset_v();
    outputBuf = ctxt->getPrimaryCodecBuf();

    result = ctxt->encodeOssToPer(PDUSessionResourceModifyUnsuccessfulTransfer_PDU, &ngApPdu, &outputBuf);

    if (result == encodeSuccessCode_c){
        perBufRef_m = outputBuf;
    }

    return result;
}

void PduSessResMdfyUnsuccTransferPdu::getDecodedOssDataPtrs_v()
{
    reset_v();

	PDUSessionResourceModifyUnsuccessfulTransfer* ptr =
		reinterpret_cast<PDUSessionResourceModifyUnsuccessfulTransfer *>(ossBufRef_m.value);

    if (ptr == NULL){
        return;
    }
    cause_mp = &ptr->cause;

}

Cause *
PduSessResMdfyUnsuccTransferPdu::getCauseEncodePtr()
{
    return &cause_m;
}


Cause *
PduSessResMdfyUnsuccTransferPdu::getCauseDecodePtr()
{
    return cause_mp;
}

