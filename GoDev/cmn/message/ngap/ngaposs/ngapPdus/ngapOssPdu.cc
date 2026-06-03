#include "ngapOssPdu.h"

NgapOssPdu::NgapOssPdu():
                ngapPdu_mp(NULL)
{
   ossBufRef_m.length = 0;
   ossBufRef_m.value = NULL;
   perBufRef_m.length = 0;
   perBufRef_m.value = NULL;
}

NgapOssPdu::~NgapOssPdu()
{
}

void NgapOssPdu::reset_v()
{
   ossBufRef_m.length = 0;
   ossBufRef_m.value = NULL;
   perBufRef_m.length = 0;
   perBufRef_m.value = NULL;
}

OssBuf *NgapOssPdu::getPerBufRef()
{
   return(&perBufRef_m);
}

OssBuf *NgapOssPdu::getOssBufRef()
{
   return(&ossBufRef_m);
}

void NgapOssPdu::setPerBufRef_v(const OssBuf &bufRef_r)
{
   perBufRef_m = bufRef_r;
}

void NgapOssPdu::setOssBufRef_v(const OssBuf &bufRef_r)
{
   ossBufRef_m = bufRef_r;
}

unsigned int NgapOssPdu::decodePerToOss(NgapOssCtxt* ctxt)
{
   if(NULL == ctxt)
   {
       return -1;
   }
   
   OssBuf inputBuf;
   OssBuf outputBuf;
   int pduNum = NGAP_PDU_PDU;

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

unsigned int NgapOssPdu::encodeOssToPer()
{
   return 0;
}

unsigned int NgapOssPdu::decodeNgapPdu(NgapOssCtxt* ctxt)
{
   if(NULL == ctxt)
   {
       return -1;
   }
   
   OssBuf inputBuf;
   OssBuf outputBuf;
   unsigned int result;
   int pduNum = NGAP_PDU_PDU;

   inputBuf = perBufRef_m;

   ctxt->reset_v();
   outputBuf = ctxt->getPrimaryCodecBuf();

   // do not use auto decode, but manual decode.
   result = ctxt->decodePerToOss(&pduNum, &inputBuf, &outputBuf, false);

   if (result == 0)
   {
      ngapPdu_mp = (NGAP_PDU *) outputBuf.value;
      ossBufRef_m = outputBuf;
   }

   return result;
}

unsigned int NgapOssPdu::getMsgType()
{
   NGAP_PDU *ngApMsg_p;

   ngApMsg_p = getNgapPduDecodedPtr();

   if (ngApMsg_p == NULL)
   {
      return(0); // don't get msg type
   }
   return(ngApMsg_p->choice);
}

NGAP_PDU *NgapOssPdu::getNgapPduEncodePtr()
{
   return(&ngapPdu_m);
}

NGAP_PDU *NgapOssPdu::getNgapPduDecodedPtr()
{
   return(ngapPdu_mp);
}

ProcedureCode NgapOssPdu::getProcedureCode()
{
   NGAP_PDU *ngApMsg_p;

   ngApMsg_p = getNgapPduDecodedPtr();

   if (ngApMsg_p == NULL)
   {
      return(0); // error happend
   }

   switch(ngApMsg_p->choice)
   {
      case initiatingMessage_chosen:
      {
         return(ngApMsg_p->u.initiatingMessage.procedureCode);
         break;
      }

     case successfulOutcome_chosen:
     {
        return(ngApMsg_p->u.successfulOutcome.procedureCode);
        break;
     }

     case unsuccessfulOutcome_chosen:
     {
        return(ngApMsg_p->u.unsuccessfulOutcome.procedureCode);
        break;
     }

     default:
     {
        return(0);
        break;
     }
   }

   return(0);
}

void NgapOssPdu::getDecodedOssDataPtrs_v()
{
   ngapPdu_mp = (NGAP_PDU *) ossBufRef_m.value;
}

