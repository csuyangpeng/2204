#ifndef __pduSessionResourceReleaseResponseTransferPdu_h__INCLUDED__
#define __pduSessionResourceReleaseResponseTransferPdu_h__INCLUDED__

#include "ngapOssPdu.h"
#include <ngapMsgTypes.h>
///////////////////////////////////////////////////////////////////////////////
// @class PDUSessionResourceReleaseResponseTransferPdu
//-- **************************************************************
//--
//-- PDU Session Resource Release Response Transfer 
//--
//-- **************************************************************
//
// PDUSessionResourceReleaseResponseTransfer ::= SEQUENCE {
//  protocolIEs     ProtocolIE-Container    { {PDUSessionResourceReleaseResponseTransfer-IEs} },
// { ID id-SecondaryRATUsageInformation     CRITICALITY ignore  EXTENSION SecondaryRATUsageInformation      PRESENCE optional   },
//  ...
//}
// PDUSessionResourceReleaseResponseTransfer-IEs NGAP-PROTOCOL-IES ::= {
//  ...
//}
///////////////////////////////////////////////////////////////////////////////
class PduSesResRelRespTransferPdu : public NgapOssPdu
{
  public:
    PduSesResRelRespTransferPdu();
    //virtual
    ~PduSesResRelRespTransferPdu();

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

    /////////////////////////////////////////////////////////////
    /// @brief get Decoded Oss Data Pointers
    ///  processes the decoded OSS buffer and populates
    ///  the data access pointers in the class for the IEs present in the
    ///  decoded OSS buffer
    /////////////////////////////////////////////////////////////
    //virtual
    void getDecodedOssDataPtrs_v();

    /////////////////////////////////////////////////////////////
    /// @brief get PDUSessionResourceReleaseResponseTransfer pointer
    /// returns the pointer to the
    /// PDUSessionResourceReleaseResponseTransfer in the decoded message
    /////////////////////////////////////////////////////////////
    PDUSessionResourceReleaseResponseTransfer *getPduSessResRelRespTransferDecodedPtr() const;

    SecondaryRATUsageInformation *getSecondaryRATUsageInformationDecodedPtr() const;
    SecondaryRATUsageInformation *getSecondaryRATUsageInformationEncodePtr();
    bool isSecondaryRatUsageInfoPrst();

    VolumeTimedReport_Item *getVolumeTimedReportDecodedPtr(unsigned int index) const;
    VolumeTimedReport_Item *getVolumeTimedReportEncodePtr(unsigned int index);
    void setNumOfVolumeTimedReport_v(unsigned int num);
    unsigned int getNumOfVolumeTimedReport() const;

    QoSFlowsUsageReport_Item *getQosFlowsUsageReportDecodedPtr(unsigned int index) const;
    QoSFlowsUsageReport_Item *getQosFlowsUsageReportEncodePtr(unsigned int index);
    void setNumOfQosFlowsUsageReport_v(unsigned int num);
    unsigned int getNumOfQosFlowsUsageReport() const;

    VolumeTimedReport_Item* getQosFlowTimeReportDecodedPtr(QoSFlowsUsageReport_Item * qosPtr,
                                                unsigned int indQosFlowTime) const;
    VolumeTimedReport_Item* getQosFlowTimeReportEncodePtr(unsigned int indTACs,
                                               unsigned int indQosFlowTime);
    void setNumOfQosFlowTimeReport_v(unsigned int num, unsigned int index);
    unsigned int getNumOfQosFlowTimeReport(unsigned int index) const;

  private:

    typedef PduSessResRelRespTransferIe ProtocolIeType;

    //OSS structures
    PDUSessionResourceReleaseResponseTransfer* pduSessResRelRespTransfer_mp;

    //Secondary RAT Usage Information
    SecondaryRATUsageInformation secRATUsageInfo_m;
    SecondaryRATUsageInformation *secRATUsageInfo_mp;
    bool secondaryRATUsageInfoPresent_m;

    //Volume Timed Report List
    unsigned int numOfVolumeTimedReportItem_ma;
    struct VolumeTimedReportList_ volumeTimedReportList_ma[maxnoofTimePeriods_c];

    //Qos Flow Usage Report List
    unsigned int numOfQosFlowUsageReportItem_ma;
    struct QoSFlowsUsageReportList_ qosFlowUsageReportList_ma[maxnoofQoSFlows_c];

    //Volume Timed Report List
    unsigned int numOfVolumeTimedReportItem_mb[maxnoofQoSFlows_c];
    struct VolumeTimedReportList_ volumeTimedReportList_mb[maxnoofQoSFlows_c][maxnoofTimePeriods_c];
};

#endif