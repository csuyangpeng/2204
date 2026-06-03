#ifndef __ngapOssPdu_h__INCLUDED__
#define __ngapOssPdu_h__INCLUDED__

#include "ngapOss.h"
#include "ngapOssCtxt.h"


// Forward Reference
class NgapOssCtxt;

class NgapOssPdu
{
public:
  enum
  {
      encodeSuccessCode_c                 = 0,
      defaultFailureCode_c                = 255,
      maxnoofTACs_c                       = 256,
      maxnoofTimePeriods_c                = 2,
      maxnoofNGConnectionsToReset_c       = 65536,
      maxnoofBPLMNs_c                     = 12,
      maxnoofPLMNs_c                      = 12,
      maxnoofSlices_c                     = 4,
      maxRanNameValue_c                   = 150,
      maxAMFNameValue_c                   = 150,
      maxnoofServedGUAMIs_c               = 256,
      maxnoofPDUSessions_c                = 256,
      maxnoofQoSFlows_c                   = 64,
      maxnoofTAIforPaging                 = 16,
      maxnoofMultiConnectivityMinusOne_c  = 3,
      maxNasPduValue_c                    = 256,
      maxnoofRecommendedCells_c           = 16,
      maxnoofRecommendedRANNodes_c        = 16,
      maxnoofAoI_c                        = 64,
      maxnoofTAIinAoI_c                   = 16,
      maxnoofCellinAoI_c                  = 256,
      maxnoofRANNodeinAoI_c               = 64,
      maxnoofAllowedAreas_c               = 16,
      maxnoofForbTACs_c                   = 4096,
      maxnoofEPLMNsPlusOne_c              = 16,
      maxnoofEPLMNs_c                     = 15,
      maxnoofCellsUEMovingTrajectory_c    = 16,
      maxnoofTAIforInactive_c             = 16,
      maxMaskedIMEISV_c                   = 8,
      maxUeRadioCapabilityValue_c         = 4096,
      maxSecurityKeyValue_c               = 32,
      maxnoofNSSAIs_c                     = 8,
      maxNa_PduValue_c                    = 8,
      maxGnbIdValue_c                     = 4,
      maxPlmnIdLen_c                      = 3,
      maxNrCellIdLength_c                 = 5,
      maxRrcEstablishmentCauseValue_c     = 128,
      maxAmfSetIdValue_c                  = 2,
      maxUeIdentityIndexValue_c           = 2,
      maxnoofAllowedSnssaiValue_c         = 8,
      maxnoofMultiConnectivity            = 4,
  };
  
  NgapOssPdu();
  virtual ~NgapOssPdu();
  
  ////////////////////////////////////////////////////////////////////////
  /// @brief reset_v
  ///
  /// This fucntion sets the primary and the secondary buffers.
  ///
  /// @param None
  ///
  /// @warning - None.
  ///
  /// @returns - None
  ///
  ////////////////////////////////////////////////////////////////////////
  virtual void reset_v();

  ////////////////////////////////////////////////////////////////////////
  /// @brief Get the reference to the encoded data
  ///
  /// This fucntion returns the reference to the encoded data.
  ///
  /// @param None
  ///
  /// @warning - None.
  ///
  /// @returns - reference to the encoded OssBuf
  ///
  ////////////////////////////////////////////////////////////////////////
  OssBuf *getPerBufRef();
  
  ////////////////////////////////////////////////////////////////////////
  /// @brief Get the reference to the decoded data
  ///
  /// This fucntion returns the reference to the decoded data.
  ///
  /// @param None
  ///
  /// @warning - None.
  ///
  /// @returns - reference to the OssBuf
  ///
  ////////////////////////////////////////////////////////////////////////
  OssBuf *getOssBufRef();

  ////////////////////////////////////////////////////////////////////////
  /// @brief set the PerBufRef
  ///
  /// This fucntion sets the perBufref_m with the encoded message.
  ///
  /// @param &bufRef_r - input - reference to the encoded message.
  ///
  /// @warning - None.
  ///
  /// @returns - None.
  ///
  ////////////////////////////////////////////////////////////////////////
  void setPerBufRef_v(const OssBuf &bufRef_r);
  
  ////////////////////////////////////////////////////////////////////////
  /// @brief set the OssBufRef
  ///
  /// This fucntion sets the ossBufRef_m with the decoded message.
  ///
  /// @param &bufRef_r - input - reference to the decoded message.
  ///
  /// @warning - None.
  ///
  /// @returns - None.
  ///
  ////////////////////////////////////////////////////////////////////////
  void setOssBufRef_v(const OssBuf &bufRef_r);
  
  ////////////////////////////////////////////////////////////////////////
  /// @brief decode from PER to OSS
  ///
  /// This fucntion initializes the input buffer and declares output buffer
  /// which are then used to call decode fucntion to automatically fully decode 
  /// the message.
  ///
  /// @param None.
  ///
  /// @warning - None.
  ///
  /// @returns - Return Code - 0 = success.
  ///
  ////////////////////////////////////////////////////////////////////////
  virtual unsigned int decodePerToOss(NgapOssCtxt* ctxt);
  
  ////////////////////////////////////////////////////////////////////////
  /// @brief encode OSS to PER
  ///
  /// This is a blank implementation of the fucntion left for the derived 
  /// calsses to be defined as needed.
  ///
  /// @param None.
  ///
  /// @warning - All mandatory IEs should be set before calling this method.
  ///
  /// @returns - Return Code - 0 = success.
  ///
  ////////////////////////////////////////////////////////////////////////
  virtual unsigned int encodeOssToPer();
  
  ////////////////////////////////////////////////////////////////////////
  /// @brief get the data pointer for the decoded data
  ///
  /// This function sets the NgApPdu_mp with the pointer to the decoded data.
  ///
  /// @param None.
  ///
  /// @warning - None.
  ///
  /// @returns - None.
  ///
  ////////////////////////////////////////////////////////////////////////
  virtual void getDecodedOssDataPtrs_v();

  ////////////////////////////////////////////////////////////////////////
  /// @brief decode Ngap Pdu
  ///
  /// This fucntion initializes the input buffer and declares output buffer 
  /// which are then used to call decode fucntion to do a manual decode 
  /// on the message.
  ///
  /// @param None.
  ///
  /// @warning - None.
  ///
  /// @returns - Return Code - 0 = success.
  ///
  ////////////////////////////////////////////////////////////////////////
  unsigned int decodeNgapPdu(NgapOssCtxt* ctxt);
  
  ////////////////////////////////////////////////////////////////////////
  /// @brief get the message type
  ///
  /// This fucntion gets the message type i.e choice which is a member of
  /// the structure NGAP_PDU
  ///
  /// @param None.
  ///
  /// @warning - None.
  ///
  /// @returns - Return Code - 0 = failure.
  ///
  ////////////////////////////////////////////////////////////////////////
  unsigned int getMsgType();
  
  ////////////////////////////////////////////////////////////////////////
  /// @brief get the procedude code
  ///
  /// This fucntion gets the procedure code depending on the choice which is
  /// a member of the structure NGAP_PDU
  ///
  /// @param None.
  ///
  /// @warning - None.
  ///
  /// @returns - Return Code - 0 = failure.
  ///
  ////////////////////////////////////////////////////////////////////////
  ProcedureCode getProcedureCode();
  
  ////////////////////////////////////////////////////////////////////////
  /// @brief get encode data pointer
  ///
  /// This fucntion gets the reference to the encode data pointer
  ///
  /// @param None.
  ///
  /// @warning - None.
  ///
  /// @returns - reference to the NGAP_PDU
  ///
  ////////////////////////////////////////////////////////////////////////
  NGAP_PDU *getNgapPduEncodePtr();
  
  ////////////////////////////////////////////////////////////////////////
  /// @brief get decode data pointer
  ///
  /// This fucntion gets the reference to the decode data pointer
  ///
  /// @param None.
  ///
  /// @warning - None.
  ///
  /// @returns - pointer to the NGAP_PDU
  ///
  ////////////////////////////////////////////////////////////////////////
  NGAP_PDU *getNgapPduDecodedPtr();

protected:

  NGAP_PDU ngapPdu_m;
  NGAP_PDU *ngapPdu_mp;

  OssBuf perBufRef_m;
  OssBuf ossBufRef_m;
};

#endif // __NgapOssPdu_h__INCLUDED__

