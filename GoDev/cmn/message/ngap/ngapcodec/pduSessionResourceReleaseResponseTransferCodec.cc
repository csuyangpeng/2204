#include <iostream>
#include <sstream>
#include <string>

#include "pduSessionResourceReleaseResponseTransferCodec.h"

using namespace std;

PduSessResRelRespTransferCodec::PduSessResRelRespTransferCodec() {
  reset();
}

PduSessResRelRespTransferCodec::~PduSessResRelRespTransferCodec() {}

bool
PduSessResRelRespTransferCodec::isSecRatUsageInfoPrst()
{
    return isSecRatUsageInfoPrst_m;
}

void
PduSessResRelRespTransferCodec::setSecRatUsageInfo(SecRatUsageInformation& secRatUsage)
{
    secRatUsageInfo_m = secRatUsage;
    isSecRatUsageInfoPrst_m = true;
}

SecRatUsageInformation&
PduSessResRelRespTransferCodec::getSecRatUsageInfo()
{
    return secRatUsageInfo_m;
}



bool PduSessResRelRespTransferCodec::reset(){
    msgPdu_m.reset_v();
    isSecRatUsageInfoPrst_m = false;
    secRatUsageInfo_m.pduSessUsageReportPrst = false;
    secRatUsageInfo_m.pduSessUsageReport.ratType = 0;
    secRatUsageInfo_m.pduSessUsageReport.volumeTimeReportList.clear();
    secRatUsageInfo_m.qosFlowUsageReportListPrst = false;
    secRatUsageInfo_m.qosFlowUsageReportList.clear();
}

//bool PduSessResRelRespTransferCodec::encode(NgapOssCtxt* ctxt, MsgBuffer msgBuffer)
MsgBuffer PduSessResRelRespTransferCodec::encode(NgapOssCtxt *ctxt)
{
    MsgBuffer msgBuffer;
    msgBuffer.length = 0;
    msgBuffer.value = NULL;

    //optional IEs
    //secondaryRatUsageInformation
    if (isSecRatUsageInfoPrst_m) {
        SecondaryRATUsageInformation *secRatUsagePtr = msgPdu_m.getSecondaryRATUsageInformationEncodePtr();
        if (secRatUsageInfo_m.pduSessUsageReportPrst) {
            secRatUsagePtr->bit_mask |= pDUSessionUsageReport_present;
            secRatUsagePtr->pDUSessionUsageReport.rATType = _enum3(secRatUsageInfo_m.pduSessUsageReport.ratType);
            Uint32 num = secRatUsageInfo_m.pduSessUsageReport.volumeTimeReportList.size();
            msgPdu_m.setNumOfVolumeTimedReport_v(num);
            for (int i = 0; i < num; i++)
            {
                VolumeTimedReport_Item* item = msgPdu_m.getVolumeTimedReportEncodePtr(i);
                VolumeTimeReport* volumeTimeReportItem = &secRatUsageInfo_m.pduSessUsageReport.volumeTimeReportList[i];

                item->startTimeStamp.length = sizeof(volumeTimeReportItem->startTimeStamp);
                memcpy(item->startTimeStamp.value,volumeTimeReportItem->startTimeStamp,item->startTimeStamp.length);
                item->endTimeStamp.length = sizeof(volumeTimeReportItem->endTimeStamp);
                memcpy(item->endTimeStamp.value,volumeTimeReportItem->endTimeStamp,item->endTimeStamp.length);
                item->usageCountUL = volumeTimeReportItem->usageCountUL;
                item->usageCountDL = volumeTimeReportItem->usageCountDL;
            }
        }
        if (secRatUsageInfo_m.qosFlowUsageReportListPrst) {
            secRatUsagePtr->bit_mask |= qosFlowsUsageReportList_present;
            unsigned int num = secRatUsageInfo_m.qosFlowUsageReportList.size();
            msgPdu_m.setNumOfQosFlowsUsageReport_v(num);
            for (int i = 0;i < num; i++)
            {
                QoSFlowsUsageReport_Item* item = msgPdu_m.getQosFlowsUsageReportEncodePtr(i);

                unsigned int numofQosFlowTimeReport = secRatUsageInfo_m.qosFlowUsageReportList[i].qosFlowsTimeReportList.size();
                msgPdu_m.setNumOfQosFlowTimeReport_v(numofQosFlowTimeReport,i);
            }
            for (int i = 0; i <num; i++) {
                QoSFlowsUsageReport_Item* qosFlowUsageReportPtr = msgPdu_m.getQosFlowsUsageReportEncodePtr(i);
                if (NULL == qosFlowUsageReportPtr)
                {
                    std::cout << "NULL Pointer for qosFlowUsageReportPtr" <<std::endl;
                    return msgBuffer;
                }

                qosFlowUsageReportPtr->qosFlowIdentifier = (&secRatUsageInfo_m.qosFlowUsageReportList[i])->qosFlowId;
                qosFlowUsageReportPtr->rATType = _enum3((&secRatUsageInfo_m.qosFlowUsageReportList[i])->ratType);

                Uint32 numofQosFlowTimeReport = (&secRatUsageInfo_m.qosFlowUsageReportList[i])->qosFlowsTimeReportList.size();
                for (int j = 0;j < numofQosFlowTimeReport;j++) {
                    VolumeTimedReport_Item* qosFlowTimeReportPtr = msgPdu_m.getQosFlowTimeReportEncodePtr(i,j);
                    if (NULL == qosFlowTimeReportPtr) {
                        std::cout << "NULL Pointer for qosFlowTimeReportPtr" <<std::endl;
                        return msgBuffer;
                    }
                    qosFlowTimeReportPtr->startTimeStamp.length = sizeof((&(&secRatUsageInfo_m.qosFlowUsageReportList[i])->qosFlowsTimeReportList[j])->startTimeStamp);
                    memcpy(qosFlowTimeReportPtr->startTimeStamp.value,(&(&secRatUsageInfo_m.qosFlowUsageReportList[i])->qosFlowsTimeReportList[j])->startTimeStamp,qosFlowTimeReportPtr->startTimeStamp.length);
                    qosFlowTimeReportPtr->endTimeStamp.length = sizeof((&(&secRatUsageInfo_m.qosFlowUsageReportList[i])->qosFlowsTimeReportList[j])->endTimeStamp);
                    memcpy(qosFlowTimeReportPtr->endTimeStamp.value,(&(&secRatUsageInfo_m.qosFlowUsageReportList[i])->qosFlowsTimeReportList[j])->endTimeStamp,qosFlowTimeReportPtr->endTimeStamp.length);
                    qosFlowTimeReportPtr->usageCountUL = (&(&secRatUsageInfo_m.qosFlowUsageReportList[i])->qosFlowsTimeReportList[j])->usageCountUL;
                    qosFlowTimeReportPtr->usageCountDL = (&(&secRatUsageInfo_m.qosFlowUsageReportList[i])->qosFlowsTimeReportList[j])->usageCountDL;
                }
            }

        }

    }
    msgPdu_m.encodeOssToPer(ctxt);

    OssBuf *perBuf = msgPdu_m.getPerBufRef();

    msgBuffer.length = perBuf->length;
    msgBuffer.value = perBuf->value;

    return msgBuffer;
}

bool 
PduSessResRelRespTransferCodec::decode(NgapOssCtxt *ctxt, MsgBuffer msgBuf){

    reset();

    OssBuf ossbuf;
    ossbuf.length = msgBuf.length;
    ossbuf.value = msgBuf.value;

    msgPdu_m.setPerBufRef_v(ossbuf);

    if (0 != msgPdu_m.decodePerToOss(ctxt)) {
      return false;
    }

    //optional IEs
    //secondaryRatUsageInformation
    isSecRatUsageInfoPrst_m = msgPdu_m.isSecondaryRatUsageInfoPrst();
    if (true == isSecRatUsageInfoPrst_m)
    {
        SecondaryRATUsageInformation* secRatUsageInfoPtr =
                msgPdu_m.getSecondaryRATUsageInformationDecodedPtr();
        if(msgPdu_m.getSecondaryRATUsageInformationDecodedPtr()->bit_mask & pDUSessionUsageReport_present)
        {
            secRatUsageInfo_m.pduSessUsageReport.ratType = secRatUsageInfoPtr->pDUSessionUsageReport.rATType;
            for(Uint32 i =0; i<msgPdu_m.getNumOfVolumeTimedReport(); i++)
            {
                VolumeTimedReport_Item* volTimeReportItemPtr = msgPdu_m.getVolumeTimedReportDecodedPtr(i);
                if(NULL == volTimeReportItemPtr)
                {
                    return false;
                }
                VolumeTimeReport volumeTimeRprt_m;
                memcpy(volumeTimeRprt_m.startTimeStamp,volTimeReportItemPtr->startTimeStamp.value,volTimeReportItemPtr->startTimeStamp.length);
                memcpy(volumeTimeRprt_m.endTimeStamp,volTimeReportItemPtr->endTimeStamp.value,volTimeReportItemPtr->endTimeStamp.length);
                volumeTimeRprt_m.usageCountUL = volTimeReportItemPtr->usageCountUL;
                volumeTimeRprt_m.usageCountDL = volTimeReportItemPtr->usageCountDL;

                secRatUsageInfo_m.pduSessUsageReport.volumeTimeReportList.push_back(volumeTimeRprt_m);
            }
            secRatUsageInfo_m.pduSessUsageReportPrst = true;
        }
        if (msgPdu_m.getSecondaryRATUsageInformationDecodedPtr()->bit_mask & qosFlowsUsageReportList_present)
        {
            for (Uint32 i = 0; i <msgPdu_m.getNumOfQosFlowsUsageReport();i++)
            {
                QoSFlowsUsageReport_Item* qosFlowUsageRprtPtr = msgPdu_m.getQosFlowsUsageReportDecodedPtr(i);
                if(NULL == qosFlowUsageRprtPtr)
                {
                    return false;
                }
                QosFlowUsageReport qosFlowUsageRprtItem_m;
                qosFlowUsageRprtItem_m.qosFlowId = qosFlowUsageRprtPtr->qosFlowIdentifier;
                qosFlowUsageRprtItem_m.ratType = qosFlowUsageRprtPtr->rATType;
                for(Uint32 j = 0; j <msgPdu_m.getNumOfQosFlowTimeReport(i);j++)
                {
                    VolumeTimedReport_Item* qosFlowTimeRprtPtr = msgPdu_m.getQosFlowTimeReportDecodedPtr(qosFlowUsageRprtPtr,j);
                    if(NULL == qosFlowTimeRprtPtr)
                    {
                        return false;
                    }
                    VolumeTimeReport qosFlowTimeRprtItem;
                    memcpy(qosFlowTimeRprtItem.startTimeStamp,qosFlowTimeRprtPtr->startTimeStamp.value,qosFlowTimeRprtPtr->startTimeStamp.length);
                    memcpy(qosFlowTimeRprtItem.endTimeStamp,qosFlowTimeRprtPtr->endTimeStamp.value,qosFlowTimeRprtPtr->endTimeStamp.length);
                    qosFlowTimeRprtItem.usageCountUL = qosFlowTimeRprtPtr->usageCountUL;
                    qosFlowTimeRprtItem.usageCountDL = qosFlowTimeRprtPtr->usageCountDL;
                    qosFlowUsageRprtItem_m.qosFlowsTimeReportList.push_back(qosFlowTimeRprtItem);
                }
                secRatUsageInfo_m.qosFlowUsageReportList.push_back(qosFlowUsageRprtItem_m);
            }
            secRatUsageInfo_m.qosFlowUsageReportListPrst = true;
        }
        //secRatUsageInfo_m = *(msgPdu_m.getSecondaryRATUsageInformationDecodedPtr());
    }
}


std::string 
PduSessResRelRespTransferCodec::toString(){
  ostringstream oss;
  oss << "secondaryRatUsageInformation" << endl
      << "IE optional:" << isSecRatUsageInfoPrst_m << endl
      << "pdu Session Usage Report Prst:" << secRatUsageInfo_m.pduSessUsageReportPrst << endl
      << "ratType value:" << secRatUsageInfo_m.pduSessUsageReport.ratType << endl
      << "the volume list:" <<endl;
  vector<VolumeTimeReport>::iterator it;
  for(it=secRatUsageInfo_m.pduSessUsageReport.volumeTimeReportList.begin();
      it<secRatUsageInfo_m.pduSessUsageReport.volumeTimeReportList.end();it++)
  {
      cout << "start time stamp:" << it->startTimeStamp[0] << it->startTimeStamp[1]
                                  << it->startTimeStamp[2] << it->startTimeStamp[3] << endl;
      cout << "end time stamp:" << it->endTimeStamp[0] << it->endTimeStamp[1]
                                << it->endTimeStamp[2] << it->endTimeStamp[3] << endl;
      cout << "usage count uplink:" << it->usageCountUL << endl;
      cout << "usage count downlink:" << it->usageCountDL << endl;
  }
  oss << "qos Flow Usage Reoport List Prst:" << secRatUsageInfo_m.qosFlowUsageReportListPrst << endl
      << "qos Flow Usage Report List:" << endl;
  vector<QosFlowUsageReport>::iterator qosIt;
  for(qosIt=secRatUsageInfo_m.qosFlowUsageReportList.begin();qosIt<secRatUsageInfo_m.qosFlowUsageReportList.end();qosIt++)
  {
      cout << "qosFlowId:" << qosIt->qosFlowId <<endl;
      cout << "ratType:" << int(qosIt->ratType) << endl;
      cout << "qos flow time report list:" << endl;
      VolumeTimeReportList qosflowTimeReportList = qosIt->qosFlowsTimeReportList;
      vector<VolumeTimeReport>::iterator qosTimeIt;
      for(qosTimeIt=qosflowTimeReportList.begin(); qosTimeIt<qosflowTimeReportList.end();qosTimeIt++)
      {
          cout << "start time stamp:" << qosTimeIt->startTimeStamp[0] << qosTimeIt->startTimeStamp[1]
                                      << qosTimeIt->startTimeStamp[2] << qosTimeIt->startTimeStamp[3] << endl;
          cout << "end time stamp:" << qosTimeIt->endTimeStamp[0] << qosTimeIt->endTimeStamp[1]
                                    << qosTimeIt->endTimeStamp[2] << qosTimeIt->endTimeStamp[3] << endl;
          cout << "usage count uplink:" << qosTimeIt->usageCountUL << endl;
          cout << "usage count downlink:" << qosTimeIt->usageCountDL << endl;
      }
  }
  return oss.str();
}


