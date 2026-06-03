#ifndef __snssaiPdu_h__INCLUDED__
#define __snssaiPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class SnssaiPdu
//-- **************************************************************
//--
//-- SNSSAI
//--
//-- **************************************************************
//
//typedef struct S_NSSAI {
//    unsigned char   bit_mask;
//#       define      sD_present 0x80
//#       define      S_NSSAI_iE_Extensions_present 0x40
//    SST             sST;
//    SD              sD;  /* optional; set in bit_mask sD_present if present */
//    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
//                                      * S_NSSAI_iE_Extensions_present if
//                                     * present */
//} S_NSSAI;
//
///////////////////////////////////////////////////////////////////////////////

class SnssaiPdu : public NgapOssPdu
{
  public:
    SnssaiPdu();
    //virtual
    ~SnssaiPdu();

    /////////////////////////////////////////////////////////////
    /// @brief reset_v
    ///  resets all data members to default value
    /////////////////////////////////////////////////////////////
    //virtual
    void reset_v();

    /////////////////////////////////////////////////////////////
    /// @brief encode Oss to PER using ASN.1
    ///  builds and processes the OSS PDU to encode
    ///  it to PER using ASN.1
    /////////////////////////////////////////////////////////////
    //virtual
    unsigned int encodeOssToPer(NgapOssCtxt *ctxt);

    //overwrite the decode methord
    unsigned int decodePerToOss(NgapOssCtxt* ctxt);
    void getDecodedOssDataPtrs_v();

    S_NSSAI * getSnssaiEncodePtr();
    S_NSSAI * getSnssaiDecodePtr();

  private:

    S_NSSAI snssai_m;
    S_NSSAI* snssai_mp;
};

#endif
