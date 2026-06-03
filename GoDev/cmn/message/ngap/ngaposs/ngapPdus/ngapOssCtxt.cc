#include <iostream>

#include "ngapOssCtxt.h"

NgapOssCtxt::NgapOssCtxt()
{
   ossTempBuffer_m.length = ossTempBufferSize_c;
   ossTempBuffer_m.value  = ossTempBuffer_ma;

   ossinit(&automatic_m, ngap_procedures);
   ossSetFlags (&automatic_m, AUTOMATIC_ENCDEC | NOTRAPPING);
   ossSetTemporaryBuffer(&automatic_m, &ossTempBuffer_m);

   // not automatic encode/decode by default
   ossinit(&manual_m, ngap_procedures);
   ossSetFlags(&manual_m, NOTRAPPING);
   ossSetTemporaryBuffer(&manual_m, &ossTempBuffer_m);

   //set DEBUGPDU flag to enable the ossPrintPDU() output
   ossSetEncodingFlags(&manual_m, ossGetEncodingFlags(&manual_m) | DEBUGPDU );
   ossSetDecodingFlags(&manual_m, ossGetDecodingFlags(&manual_m) | DEBUGPDU );
}

NgapOssCtxt::~NgapOssCtxt()
{
   
   // Added ossterm() for cleaning up the resources used by the OSS compiler 
   ossterm(&automatic_m);
   ossterm(&manual_m);
}

 NgapOssCtxt* NgapOssCtxt::instance()
 {
    static NgapOssCtxt instance;
    return &instance;
 }

void NgapOssCtxt::reset_v()
{
   primaryBuffer_m.length = ossCodecBufferSize_c;
   primaryBuffer_m.value  = primaryCodecBuffer_ma;

   secondaryBuffer_m.length = ossCodecBufferSize_c;
   secondaryBuffer_m.value  = secondaryCodecBuffer_ma;
}

OssBuf& NgapOssCtxt::getPrimaryCodecBuf()
{
   return primaryBuffer_m;
}

OssBuf& NgapOssCtxt::getSecondaryCodecBuf()
{
   return secondaryBuffer_m;
}

unsigned int NgapOssCtxt::encodeOssToPer(int         pduNum,
                                        void*        ossInput_p,
                                        OssBuf*      perOutput_p,
                                        bool         automatic)
{
   ossGlobal *world_p = &automatic_m;
   if (automatic != true)
   {
      return(manualEncodeNotSupported_c);
   }

#ifdef OSSPRINT
   ossPrintPDU(world_p, pduNum, ossInput_p);
#endif

   unsigned int rc = ossEncode(world_p, pduNum, ossInput_p, perOutput_p);

   if (rc != successfulOperation_c)
   {
      ossPrint(world_p, "ossEncode failed.. result = (%d):  %s\n",
               rc, ossGetErrMsg(world_p));
   }

#ifdef OSSPRINT
  	ossPrintHex(&automatic_m, (char *)perOutput_p->value, perOutput_p->length);
#endif

   return rc;
}

unsigned int NgapOssCtxt::decodePerToOss(int*          pduNum_p,
                                         OssBuf*       perInput_p,
                                         OssBuf*       ossOutput_p,
                                         bool          automatic)
{
   ossGlobal *world_p = &automatic_m;
   if (automatic != true)
   {
      world_p = &manual_m;
   }

   void* output = ossOutput_p->value;
   ossSetDecodingLength(world_p, ossOutput_p->length);
   unsigned int rc = ossDecode(world_p, pduNum_p, perInput_p, &output);
   ossOutput_p->length = ossGetDecodingLength(world_p);

   if (rc != successfulOperation_c)
   {
      ossPrint(world_p, "ossDecode failed.. result = (%d):  %s\n",
               rc, ossGetErrMsg(world_p));
   }
   else
   {
#ifdef OSSPRINT
      ossPrintPDU(world_p, (int)(*pduNum_p), output);
#endif
   }

   return rc;
}

