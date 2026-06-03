#include <string.h>
#include <iostream>

#include "pduSessionResourceReleaseResponseTransferPdu.h"

PduSesResRelRespTransferPdu::PduSesResRelRespTransferPdu()
{
    reset_v();
}

PduSesResRelRespTransferPdu::~PduSesResRelRespTransferPdu()
{
}

void PduSesResRelRespTransferPdu::reset_v()
{
    pduSessResRelRespTransfer_mp = NULL;
    secRATUsageInfo_mp = NULL;
    secondaryRATUsageInfoPresent_m = false;
    secRATUsageInfo_m.bit_mask = 0;
    secRATUsageInfo_m.iE_Extension = NULL;
    secRATUsageInfo_m.pDUSessionUsageReport.bit_mask = 0;
    secRATUsageInfo_m.pDUSessionUsageReport.iE_Extensions = NULL;
    secRATUsageInfo_m.pDUSessionUsageReport.rATType = nr;
    for (unsigned int i = 0; i < maxnoofTimePeriods; i++)
    {
        volumeTimedReportList_ma[i].next = NULL;
        volumeTimedReportList_ma[i].value.bit_mask = 0;
        volumeTimedReportList_ma[i].value.iE_Extensions = NULL;
        volumeTimedReportList_ma[i].value.startTimeStamp.length = 0;
        memset(volumeTimedReportList_ma[i].value.startTimeStamp.value,0xff,
               sizeof(volumeTimedReportList_ma[i].value.startTimeStamp.value));
        volumeTimedReportList_ma[i].value.endTimeStamp.length = 0;
        memset(volumeTimedReportList_ma[i].value.endTimeStamp.value,0xff,
               sizeof(volumeTimedReportList_ma[i].value.endTimeStamp.value));
        volumeTimedReportList_ma[i].value.usageCountUL = 0;
        volumeTimedReportList_ma[i].value.usageCountDL = 0;
    }
    for (unsigned int i = 0; i < maxnoofQosFlows; i++)
    {
        qosFlowUsageReportList_ma[i].next = NULL;
        qosFlowUsageReportList_ma[i].value.bit_mask = 0;
        qosFlowUsageReportList_ma[i].value.iE_Extensions = NULL;
        qosFlowUsageReportList_ma[i].value.qosFlowIdentifier = 0;
        qosFlowUsageReportList_ma[i].value.rATType = nr;

        numOfVolumeTimedReportItem_mb[i] = 0;
        qosFlowUsageReportList_ma[i].value.qoSFlowsTimedReportList = NULL;
        for(unsigned int j = 0; j < maxnoofTimePeriods; j++)
        {
            volumeTimedReportList_mb[i][j].next = NULL;
            volumeTimedReportList_mb[i][j].value.bit_mask = 0;
            volumeTimedReportList_mb[i][j].value.iE_Extensions = NULL;
            volumeTimedReportList_mb[i][j].value.startTimeStamp.length = 0;
            memset(volumeTimedReportList_mb[i][j].value.startTimeStamp.value,0xff,
                   sizeof(volumeTimedReportList_mb[i][j].value.startTimeStamp.value));
            volumeTimedReportList_mb[i][j].value.endTimeStamp.length = 0;
            memset(volumeTimedReportList_mb[i][j].value.endTimeStamp.value,0xff,
                   sizeof(volumeTimedReportList_mb[i][j].value.endTimeStamp.value));
            volumeTimedReportList_mb[i][j].value.usageCountUL = 0;
            volumeTimedReportList_mb[i][j].value.usageCountDL = 0;

        }
    }

}

unsigned int 
PduSesResRelRespTransferPdu::decodePerToOss(NgapOssCtxt* ctxt)
{
   if(NULL == ctxt)
   {
       return -1;
   }
   
   OssBuf inputBuf;
   OssBuf outputBuf;
   int pduNum = PDUSessionResourceReleaseResponseTransfer_PDU;

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
PduSesResRelRespTransferPdu::encodeOssToPer(NgapOssCtxt *ctxt)
{
    if (NULL == ctxt)
    {
        return -1;
    }

    ProtocolIeType *nextIe_p = NULL;

    ProtocolIeType secRATUsageInfoIe;
    if(secondaryRATUsageInfoPresent_m == true)
    {
        if (numOfVolumeTimedReportItem_ma > maxnoofTimePeriods || numOfVolumeTimedReportItem_ma < 0) {
            return defaultFailureCode_c;
        } else {
            struct VolumeTimedReportList_ *nextVolumeTimeReport_p = NULL;
            for (unsigned int i = 0; i < numOfVolumeTimedReportItem_ma;i++)
            {
                volumeTimedReportList_ma[i].next = nextVolumeTimeReport_p;
                nextVolumeTimeReport_p = &volumeTimedReportList_ma[i];
            }
            secRATUsageInfo_m.pDUSessionUsageReport.pDUSessionTimedReportList = nextVolumeTimeReport_p;

            struct QoSFlowsUsageReportList_ *nextQosFlowsUsageReport_p = NULL;
            for (unsigned int i = 0; i < numOfQosFlowUsageReportItem_ma;i++)
            {
                struct VolumeTimedReportList_ *nextQosFlowTimeReport_p = NULL;
                for(unsigned int j = 0; j < numOfVolumeTimedReportItem_mb[i];j++) {
                    volumeTimedReportList_mb[i][j].next = nextQosFlowTimeReport_p;
                    nextQosFlowTimeReport_p = &volumeTimedReportList_mb[i][j];
                }
                qosFlowUsageReportList_ma[i].value.qoSFlowsTimedReportList = nextQosFlowTimeReport_p;
                qosFlowUsageReportList_ma[i].next = nextQosFlowsUsageReport_p;
                nextQosFlowsUsageReport_p = &qosFlowUsageReportList_ma[i];
            };
            secRATUsageInfo_m.qosFlowsUsageReportList = nextQosFlowsUsageReport_p;
        }
        secRATUsageInfoIe.next = nextIe_p;
        secRATUsageInfoIe.value.id = id_SecondaryRATUsageInformation;
        secRATUsageInfoIe.value.criticality = ignore;
        secRATUsageInfoIe.value.extensionValue.pduNum = PDU_PDUSessionResourceNotifyTransfer_ExtIEs_Extension_SecondaryRATUsageInformation;
        secRATUsageInfoIe.value.extensionValue.decoded.pdu_SecondaryRATUsageInformation = &secRATUsageInfo_m;
        secRATUsageInfoIe.value.extensionValue.encoded.length = 0;
        secRATUsageInfoIe.value.extensionValue.encoded.value = NULL;
        nextIe_p = &secRATUsageInfoIe;
    }

    //message
    PDUSessionResourceReleaseResponseTransfer ngApPdu;
    if (secondaryRATUsageInfoPresent_m == false)
    {
        ngApPdu.bit_mask = 0;
        ngApPdu.iE_Extensions = NULL;
    } else {
        ngApPdu.bit_mask |= PDUSessionResourceReleaseResponseTransfer_iE_Extensions_present;
        ngApPdu.iE_Extensions = nextIe_p;
    }

    //Convert Structures to PER
    OssBuf outputBuf;
    unsigned int result;

    ctxt->reset_v();
    outputBuf = ctxt->getPrimaryCodecBuf();

    result = ctxt->encodeOssToPer(PDUSessionResourceReleaseResponseTransfer_PDU, &ngApPdu, &outputBuf);

    if (result == encodeSuccessCode_c){
        perBufRef_m = outputBuf;
    }

    return result;
}

void PduSesResRelRespTransferPdu::getDecodedOssDataPtrs_v()
{
    ProtocolIeType *currIe_p;
    reset_v();

    pduSessResRelRespTransfer_mp = reinterpret_cast<PDUSessionResourceReleaseResponseTransfer *>(ossBufRef_m.value);

    if (pduSessResRelRespTransfer_mp == NULL)
    {
        return;
    }

    currIe_p = pduSessResRelRespTransfer_mp->iE_Extensions;
    while (currIe_p != NULL)
    {
        if (currIe_p->value.id == id_SecondaryRATUsageInformation
            && currIe_p->value.extensionValue.decoded.pdu_SecondaryRATUsageInformation != NULL)
        {
            secRATUsageInfo_mp = currIe_p->value.extensionValue.decoded.pdu_SecondaryRATUsageInformation;

            SecondaryRATUsageInformation curr_p = *secRATUsageInfo_mp;
            VolumeTimedReportList numPduSessTimeRprt_p;
            QoSFlowsUsageReportList numQosFlowUsageRprt_p;
            VolumeTimedReportList numQosFlowTimeRprt_p;
            numPduSessTimeRprt_p = curr_p.pDUSessionUsageReport.pDUSessionTimedReportList;
            while (numPduSessTimeRprt_p != NULL)
            {
                numOfVolumeTimedReportItem_ma ++;
                numPduSessTimeRprt_p = numPduSessTimeRprt_p->next;
            }
            numQosFlowUsageRprt_p = curr_p.qosFlowsUsageReportList;
            while (numQosFlowUsageRprt_p != NULL)
            {
                numQosFlowTimeRprt_p = numQosFlowUsageRprt_p->value.qoSFlowsTimedReportList;
                while (numQosFlowTimeRprt_p != NULL)
                {
                    numOfVolumeTimedReportItem_mb[numOfQosFlowUsageReportItem_ma]++;
                    numQosFlowTimeRprt_p = numQosFlowTimeRprt_p->next;
                }
                numOfQosFlowUsageReportItem_ma++;
                numQosFlowUsageRprt_p = numQosFlowUsageRprt_p->next;
            }
            secondaryRATUsageInfoPresent_m = true;
        }
        currIe_p = currIe_p->next;
    }
}

PDUSessionResourceReleaseResponseTransfer *
PduSesResRelRespTransferPdu::getPduSessResRelRespTransferDecodedPtr() const
{
    return pduSessResRelRespTransfer_mp;
}

SecondaryRATUsageInformation *
PduSesResRelRespTransferPdu::getSecondaryRATUsageInformationDecodedPtr() const
{
    return secRATUsageInfo_mp;
}

SecondaryRATUsageInformation *
PduSesResRelRespTransferPdu::getSecondaryRATUsageInformationEncodePtr()
{
    secondaryRATUsageInfoPresent_m = true;
    return &secRATUsageInfo_m;
}

bool
PduSesResRelRespTransferPdu::isSecondaryRatUsageInfoPrst()
{
    return secondaryRATUsageInfoPresent_m;
}

VolumeTimedReport_Item*
PduSesResRelRespTransferPdu::getVolumeTimedReportDecodedPtr(unsigned int index) const
{
    VolumeTimedReport_Item* item_p = NULL;
    VolumeTimedReportList volumeTimedReport_p = NULL;
    if (secRATUsageInfo_mp->pDUSessionUsageReport.pDUSessionTimedReportList != NULL)
    {
        volumeTimedReport_p = secRATUsageInfo_mp->pDUSessionUsageReport.pDUSessionTimedReportList;
    }
    if (volumeTimedReport_p != NULL && (index >= 0 && index <numOfVolumeTimedReportItem_ma))
    {
        unsigned int i = 0;
        for(; i < index && volumeTimedReport_p != NULL; i++ ) {
            volumeTimedReport_p = volumeTimedReport_p->next;
        }
        if (volumeTimedReport_p != NULL && i == index) {
            item_p = &(volumeTimedReport_p->value);
        }
    }

    return item_p;

}

VolumeTimedReport_Item*
PduSesResRelRespTransferPdu::getVolumeTimedReportEncodePtr(unsigned int index)
{
    if ((index >= 0 && index < maxnoofTimePeriods))
    {
        return &(volumeTimedReportList_ma[index].value);
    }
    else
    {
        return NULL;
    }
}

void
PduSesResRelRespTransferPdu::setNumOfVolumeTimedReport_v(unsigned int num)
{
    numOfVolumeTimedReportItem_ma = num;
}

unsigned int
PduSesResRelRespTransferPdu::getNumOfVolumeTimedReport() const
{
    return numOfVolumeTimedReportItem_ma;
}

QoSFlowsUsageReport_Item*
PduSesResRelRespTransferPdu::getQosFlowsUsageReportDecodedPtr(unsigned int index) const
{
    QoSFlowsUsageReport_Item* item_p = NULL;
    QoSFlowsUsageReportList qosFlowUsage_p = NULL;
    if (secRATUsageInfo_mp->qosFlowsUsageReportList != NULL)
    {
        qosFlowUsage_p = secRATUsageInfo_mp->qosFlowsUsageReportList;
    }
    if (qosFlowUsage_p != NULL && (index >= 0 && index <numOfQosFlowUsageReportItem_ma))
    {
        unsigned int i = 0;
        for(; i < index && qosFlowUsage_p != NULL; i++ ) {
            qosFlowUsage_p = qosFlowUsage_p->next;
        }
        if (qosFlowUsage_p != NULL && i == index) {
            item_p = &(qosFlowUsage_p->value);
        }
    }

    return item_p;

}

QoSFlowsUsageReport_Item*
PduSesResRelRespTransferPdu::getQosFlowsUsageReportEncodePtr(unsigned int index)
{
    if ((index >= 0 && index < maxnoofQosFlows))
    {
        return &(qosFlowUsageReportList_ma[index].value);
    }
    else
    {
        return NULL;
    }
}

void
PduSesResRelRespTransferPdu::setNumOfQosFlowsUsageReport_v(unsigned int num)
{
    numOfQosFlowUsageReportItem_ma = num;
}

unsigned int
PduSesResRelRespTransferPdu::getNumOfQosFlowsUsageReport() const
{
    return numOfQosFlowUsageReportItem_ma;
}

VolumeTimedReport_Item*
PduSesResRelRespTransferPdu::getQosFlowTimeReportDecodedPtr(QoSFlowsUsageReport_Item * qosPtr,
                                                unsigned int indQosFlowTime) const
{
    VolumeTimedReport_Item* item_p = NULL;
    VolumeTimedReportList qosFlowTimeReport_p = NULL;
    if (qosPtr != NULL) {
        qosFlowTimeReport_p = qosPtr->qoSFlowsTimedReportList;
        if (qosFlowTimeReport_p != NULL &&(indQosFlowTime >= 0 && indQosFlowTime < maxnoofTimePeriods))
        {
            unsigned int i = 0;
            for (;i < indQosFlowTime && qosFlowTimeReport_p != NULL; i++)
            {
                qosFlowTimeReport_p = qosFlowTimeReport_p->next;
            }

            if(qosFlowTimeReport_p != NULL && i == indQosFlowTime)
            {
                item_p = &(qosFlowTimeReport_p->value);
            }
        }
    }

    return item_p;

}

VolumeTimedReport_Item*
PduSesResRelRespTransferPdu::getQosFlowTimeReportEncodePtr(unsigned int index,
                                               unsigned int indQosFlowTime)
{
    if ((index >= 0 && index < maxnoofQosFlows && (indQosFlowTime >= 0 && indQosFlowTime < maxnoofTimePeriods)))
    {
        return &(volumeTimedReportList_mb[index][indQosFlowTime].value);
    }
    else
    {
        return NULL;
    }
}

void
PduSesResRelRespTransferPdu::setNumOfQosFlowTimeReport_v(unsigned int num,unsigned int index)
{
    if(index >= 0 && index < maxnoofTimePeriods)
    {
        numOfVolumeTimedReportItem_mb[index] = num;
    }
}

unsigned int
PduSesResRelRespTransferPdu::getNumOfQosFlowTimeReport(unsigned int index) const
{
    if (index >= 0 &&index < maxnoofTimePeriods)
    {
        return numOfVolumeTimedReportItem_mb[index];
    }
    else
    {
        return 0;
    }

}
