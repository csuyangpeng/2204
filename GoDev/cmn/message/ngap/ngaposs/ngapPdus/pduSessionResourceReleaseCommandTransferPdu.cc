#include <string.h>
#include <iostream>

#include "pduSessionResourceReleaseCommandTransferPdu.h"

PduSessResRelCmdTransferPdu::PduSessResRelCmdTransferPdu()
{
    reset_v();
}

PduSessResRelCmdTransferPdu::~PduSessResRelCmdTransferPdu()
{
}

void PduSessResRelCmdTransferPdu::reset_v()
{
    pduSessResRelCmdTransfer_mp = NULL;
	cause_m.choice = misc_chosen;
	cause_m.u.misc = CauseMisc_unspecified;
}

unsigned int 
PduSessResRelCmdTransferPdu::decodePerToOss(NgapOssCtxt* ctxt)
{
   if(NULL == ctxt)
   {
       return -1;
   }
   
   OssBuf inputBuf;
   OssBuf outputBuf;
   int pduNum = PDUSessionResourceReleaseCommandTransfer_PDU;

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
PduSessResRelCmdTransferPdu::encodeOssToPer(NgapOssCtxt *ctxt)
{
    if (NULL == ctxt){
        return -1;
    }

	PDUSessionResourceReleaseCommandTransfer ngApPdu;
	ngApPdu.bit_mask = 0;
	ngApPdu.iE_Extensions = NULL;
	memcpy(&ngApPdu.cause, &cause_m, sizeof(cause_m));
	
    //Convert Structures to PER
    OssBuf outputBuf;
    unsigned int result;

    ctxt->reset_v();
    outputBuf = ctxt->getPrimaryCodecBuf();

    result = ctxt->encodeOssToPer(PDUSessionResourceReleaseCommandTransfer_PDU, &ngApPdu, &outputBuf);

    if (result == encodeSuccessCode_c){
        perBufRef_m = outputBuf;
    }

    return result;
}

void PduSessResRelCmdTransferPdu::getDecodedOssDataPtrs_v()
{
    reset_v();

	pduSessResRelCmdTransfer_mp =
		reinterpret_cast<PDUSessionResourceReleaseCommandTransfer *>(ossBufRef_m.value);

    if (pduSessResRelCmdTransfer_mp == NULL){
        return;
    }
    cause_mp = &pduSessResRelCmdTransfer_mp->cause;

}

Cause *
PduSessResRelCmdTransferPdu::getCauseEncodePtr()
{
    return &cause_m;
}

Cause *
PduSessResRelCmdTransferPdu::getCauseDecodePtr()
{
    return cause_mp;
}
