#ifndef __typesCodec_h__
#define __typesCodec_h__

#include <string>
#include <vector>

//common types definition
typedef unsigned char Uint8;
typedef unsigned short Uint16;
typedef unsigned int Uint32;
typedef unsigned long Uint64;

//static const definition
static const Uint8 tacSize_c = 3;
static const Uint8 plmnSize_c = 3;
static const Uint8 gnbIdSize_c = 4;
static const Uint8 sstSize_c = 1;
static const Uint8 sdSize_c = 3;
static const Uint8 ranNameSize_c = 150;
static const Uint8 amfReginIdBitLen = 8;
static const Uint8 amfSetIdBitLen = 10;
static const Uint8 amfPointerBitLen = 6;
static const Uint8 NRCellIdentityBitLen = 36;
static const Uint8 amfSetIdLen = 2;
static const Uint8 amfPointerLen = 1;
static const Uint8 maxNrCellIdLen = 5;
static const Uint8 ngranTraceIdSize_c=8;
static const Uint8 bitLen = 8;
static const Uint8 transportLayerAddrBitLen = 160;
static const Uint8 nRencryptionAlgBitLen=16;
static const Uint8 nRintegrityProtectionAlgBitLen=16;
static const Uint8 eUTRAencryptionAlgBitLen=16;
static const Uint8 eUTRAintegrityProtectionAlgBitLen=16;
static const Uint16 securityKeyBitLen=256;
static const Uint8 interfacesToTraceBitLen=8;
static const Uint8 traceColEntityIPAddrBitLen=160;
static const Uint8 maskedIMEISVBitLen=64;
static const Uint8 uEIdentityIndexValueBitLen=10;
static const Uint8 perRegUpdateTimerBitLen=8;
static const Uint16 maxSmfTransferLen=512;
static const Uint8 maxSnssaiStrLen=4;
static const Uint16 maxNasLen=256;
static const Uint8 UeIdentIndexValueSize_c = 2;
static const Uint16 maxNtwrkInstanceLen=256;

typedef struct SmfTransferPdu_ {
        unsigned char data[maxSmfTransferLen];
} SmfTransferPdu;

typedef struct SnssaiStr_ {
        unsigned char data[amfSetIdBitLen];
} SnssaiStr;

typedef struct NasPdu_{
        unsigned char data[maxNasLen];
}NasPdu;

typedef struct CommonNetworkInstance_ {
    unsigned char data[maxNtwrkInstanceLen];
}CmmnNtwrkInstance;

typedef struct MsgBuffer
{
        long length;
        unsigned char *value;
} MsgBuffer;

//plmn
typedef Uint8 PlmnId[plmnSize_c];

//NGRANTraceID
typedef Uint8 NGRANTraceId[ngranTraceIdSize_c];

//UeIdentityIndexValue
typedef Uint8 UeIdentIndexValue[UeIdentIndexValueSize_c];

//GnbId
// typedef Uint8 GnbIdVal[gnbIdSize_c];
typedef struct GnbId
{
        Uint16 len;
        Uint8 val[gnbIdSize_c];
} GnbId;

//Global ID
typedef struct GGnbId
{
        PlmnId plmnid;
        GnbId gnbId;
} GGnbId;

//TAC
typedef Uint8 Tac[tacSize_c];

//TAI
typedef struct TaiType
{
        PlmnId plmn;
        Tac tac;
} TaiType;

//SNssai List
typedef Uint8 SST_[sstSize_c];
typedef Uint8 SD_[sdSize_c];
typedef struct SNssai
{
        SST_ sst;
        bool sdPresent;
        SD_ sd;
} SNssai;
typedef std::vector<SNssai> SliceSupList;
typedef std::vector<SNssai> AllowedNssaiList;

//BroadcastPlmnList
typedef struct BPlmnItem
{
        PlmnId plmnid;
        SliceSupList ssList;
} BPlmnItem;
typedef std::vector<BPlmnItem> BPlmnList;

//Supported TAI List
typedef struct SupTAItem
{
        Tac tac;
        BPlmnList bplmnList;
} SupTAItem;
typedef std::vector<SupTAItem> SupTAList;

//AMF Identifier
typedef struct AmfIdentifier
{
        Uint8 regionId;
        Uint8 setId[2];
        Uint8 pointer;
} AmfIdentifier;

//ServedGuami List
typedef struct ServedGuamiItem
{
        PlmnId plmnId;
        AmfIdentifier amfId;
        std::string backupAmfName;
        bool backupAmfNamePrst;
} ServedGuamiItem;
typedef std::vector<ServedGuamiItem> ServedGuamiList;

// NR_CGI
typedef struct NrCgi{
        PlmnId plmn;
        Uint8 nrCellId[maxNrCellIdLen];
} NrCgi;

//Ue Security Capabilities
typedef struct UeSecurityCap{
    Uint8 nrEncryptionAlg[2];
    Uint8 nrIntegrityProAlg[2];
    Uint8 eutraEncryptionAlg[2];
    Uint8 eutraIntegrityProAlg[2];
}UeSecurityCap;

//TraceAct
typedef struct TraceAct{
     NGRANTraceId traceID;
     Uint8 interfaces;
     Uint8 traceDepth;
     Uint8 traceEntityIPAddr[20];
}TraceAct;

//CoreNetworkAssistanceInformation
typedef struct CoreNwkAssInfo{
     Uint8 indexLen[2];
         Uint8 perRegUpdateTimer;
}CoreNwkAssInfo;

//GUAMI
typedef struct Guami{
     PlmnId plmnId;
     AmfIdentifier  amfId;
}Guami;

//EmergencyFallbackRequestIndicator
typedef struct EmergeFBReqInd{
         Uint8 indicator;
         Uint8 serviceTargetCn;
}EmergeFBReqInd;

typedef std::vector<TaiType> TaiList;
typedef std::vector<Tac> TacList;
//ExpectedUEActivityBehaviour
typedef struct ExpUeActivityBehaviour{
         bool  expActPeriodPrst;
         Uint8 expActPeriod;
         bool  expIdlePeriodPrst;
         Uint8 expIdlePeriod;
         bool  srcActBhvrPrst;
     Uint8 srcUeActBhvrInfo;
}ExpUeActivityBehaviour;

//ExpectedUEMovingTrajectory
typedef struct ExpUeMovingTrajectory{
          bool  timeStayedPrst;
      Uint8 timeStayedInCell;
          NrCgi nrcgi;
}ExpUeMovingTrajectory;

typedef std::vector<ExpUeMovingTrajectory> ExpUEMovingList;

//ExpectedUEBehaviour
typedef struct ExpUeBehaviour{
        bool    expUeActBhvrPrst;
        ExpUeActivityBehaviour expUeActBhvr;
        bool    expHoIntervalPrst;
    Uint8 expHoInterval;
        bool    expUeMobPrst;
        Uint8 expUeMob;
        bool    expUeMovingTracePrst;
    ExpUEMovingList expUeMovingTraceList;
}ExpUeBehaviour;        

//CoreNetworkAssistanceInformation
typedef struct CoreNetworkAssInfo{
         Uint8 ueIdentityIndValue[2];
         bool  pagingDrxPrst;
     Uint8 pagingDrx;
         Uint8 PerRegUpdateTimer;
         bool  micoModeIndPrst;
         Uint8 micoModeInd;
         TaiList rrcInactTaiList;
         bool   expUeBhvrPrst;
         ExpUeBehaviour expUeBhvr;
}CoreNetworkAssInfo;

//CNAssistedRANParametersTuning
typedef struct CnAssRanParamsTuning{
    bool   expUeBhvrPrst;
    ExpUeBehaviour expUeBhvr;
}CnAssRanParamsTuning;

typedef std::vector<PlmnId> PlmnList;

//RatRestrictions
typedef struct RatRestric{
          PlmnId plmnId;
      Uint8 ratRestricInfo;
}RatRestric;
typedef std::vector<RatRestric> RatRestricList;

//ForbiddenAreaInfoList
typedef struct ForbiddenAreaInfo{
      PlmnId plmnId;
          TacList forbiddenTacList;
}ForbiddenAreaInfo;
typedef std::vector<ForbiddenAreaInfo> ForbiddenAreaInfoList;

//ServiceAreaInformation
typedef struct ServiceAreaInfoItem{
          PlmnId plmnId;
      TacList allowedTacList;
          TacList notAllowedTacList;
}ServiceAreaInfoItem;
typedef std::vector<ServiceAreaInfoItem> ServiceAreaInfoList;

//MobilityRestrictionList
typedef struct MobRestricList{
      PlmnId plmnId;
          PlmnList equPLMNsList;
          RatRestricList ratRestrictionsItemList;
          ForbiddenAreaInfoList forbiddenAreaInfoList;
          ServiceAreaInfoList serviceAreaInfoList;
}MobRestricList;

//5G Stmsi
typedef struct Stmsi5G
{
        Uint8 setId[2];
        Uint8 pointer;
        Uint8 tmsi[4];
} Stmsi5G;

// timeStamp
typedef Uint8 TimeStampType[4];

// User Location info
typedef struct UserLocationInfo
{
        NrCgi nRCgi;
        TaiType tai;
        bool isTimeStampPrst;
        TimeStampType timeStamp;
        bool psCellInfoPrst;
        NrCgi psCellInfo;
} UserLocationInfo;

// AmfSetId
typedef Uint8 AmfSetId[2];

//Ue Ambr
typedef struct UeAmbr
{
        long uplink;
        long downlink;
} UeAmbr;

//Transport Layer Address
typedef struct TransLayerAddr
{
        Uint8 transLayerAddr[20];
        Uint8 addTransLayerAddr[20];
} TransLayerAddr;

typedef struct EnhanceTransLayerAddr
{
        Uint8 dlTransLayerAddr[20];
        Uint8 ulTransLayerAddr[20];
        Uint8 addTransLayerAddr[20];
}EnhanceTransLayerAddr;

//GBRQosInformation
typedef struct GBRQosInformation
{
        long maxFlowBitRateDL;
        long maxFlowBitRateUL;
        long guaFlowBitRateDL;
        long guaFlowBitRateUL;
        bool notiControlPresent;
        Uint8 notiControl;
        bool maxPacketLossRateDLPresent;
        Uint32 maxPacketLossRateDL;
        bool maxPacketLossRateULPresent;
        Uint32 maxPacketLossRateUL;
} GBRQosInformation;

//AllocationAndRetentionPriority
typedef struct ARP
{
    Uint16 priorityLevel;
    Uint8 emptionCapability;
    Uint8 emptionVulnerability;
} ARP;

//nonDynamic5QI
typedef struct NonDynamic5QI
{
    long          fiveQI;

        bool isPriorityLevelQosPrst;
        long priorityLevelQos;  /* optional; set in bit_mask
                                         * priorityLevelQos_present if
                                         * present */
    bool isAveWindowPrst;
    long averagingWindow;  /* optional; set in bit_mask
                           * NonDynamic5QIDescriptor_averagingWindow_present if
                           * present */
    bool isMaxDataBusrtVolPrst;
    long maximumDataBurstVolume;  /* optional; set in bit_mask*/

} NonDynamic5QI;

typedef struct PktErrRate {
    long            pERScalar;
    long            pERExponent;
} PktErrRate;

typedef struct Dynamic5QI
{
    long priorityLevelQos;
    long packetDelayBudget;
    PktErrRate packetErrorRate;
        
        bool is5QIPrst;
    long          fiveQI;  /* optional; set in bit_mask fiveQI_present if
                              * present */
        bool isDelayCriticalPrst;                                                       
    Uint8   delayCritical;  /* optional; set in bit_mask
                                     * delayCritical_present if present */
        bool isAveWindowPrst;
    long averagingWindow;  /* optional; set in bit_mask
                              * Dynamic5QIDescriptor_averagingWindow_present if
                              * present */
    bool isMaxDataBusrtVolPrst;
    long maximumDataBurstVolume;  /* optional; set in bit_mask
                       * Dynamic5QIDescriptor_maximumDataBurstVolume_present if
                       * present */
} Dynamic5QI;

typedef struct QosCharacter{
        bool isDynamic;
        NonDynamic5QI nonDynamic5qi;
        Dynamic5QI dynamic5qi;
} QosCharacter;

//QosFlowLevelQosParameters
typedef struct QosFlowLevelQosPara
{
        QosCharacter qosCharacter;
        ARP arp;
        bool gBRQosInfoPresent;
        GBRQosInformation gBRQosInfo;
        bool refQosAttrPresent;
        Uint8 refQosAttr;
        bool addQosFlowInfoPresent;
        Uint8 addQosFlowInfo;
} QosFlowLevelQosPara;

//QosFlowSetupRequestItem
typedef struct QosFlowSetupReqItem
{
        Uint32 qosFlowInd;
        QosFlowLevelQosPara qosFlowLevQosPara;
        bool eRABIdPresent;
        Uint32 eRABId;
} QosFlowSetupReqItem;
typedef std::vector<QosFlowSetupReqItem> QosFlowSetupReqList;

//SecurityInd
typedef struct SecurityInd
{
        Uint8 inteProIndication;
        Uint8 confProIndication;
        bool  maxiProDataRatePresent;
        Uint8 maxiProDataRate;
} SecurityInd;

//gTPTunnel
typedef struct GtpTunnelInfo
{
        Uint8 ipType;
        Uint8 transportLayerAddr[20];
        Uint8 gtpTeid[4];
} GtpTunnel;


typedef std::vector<GtpTunnel> AddUpTransportLayerInfoList;



// UlNguUpTnlModifyItem
typedef struct UlNguUpTnlModifyItem
{
       GtpTunnel uplinkGtpTunnel;
       GtpTunnel downlinkGtpTunnel;
} UlNguUpTnlModifyItem;
typedef std::vector<UlNguUpTnlModifyItem> UlNguUpTnlModifyList;

//PDUSessionResourceSetupListCxtReq
typedef struct PduSessResSetupReqItem
{
        Uint8 pduSessId;
        bool nasPduPrst;
        std::string nasPdu;
        SNssai snssai;
        std::string pduSessResSetupReqTransfer;
} PduSessResSetupReqItem;

typedef std::vector<PduSessResSetupReqItem> PduSessResSetupReqList;

typedef struct AssQosFlowItem
{
        Uint32 qosFlowInd;
        bool qosFlowMapIndPrst;
        Uint8 qosFlowMapInd;
}AssQosFlowItem;

//QosFlowIdentifierList
typedef std::vector<AssQosFlowItem> QosFlowIndList;

//QosFlowSetupRequestItem
typedef struct QosFlowCodecItem
{
        Uint32 qosFlowInd;
        Uint8 causeType;
        Uint8 causeValue;
} QosFlowCodecItem;
typedef std::vector<QosFlowCodecItem> QosFlowCodecList;

//PDU Session Resource Setup Response Item
typedef struct PduSessResSetupRespItem
{
        Uint8 pduSessId;
        std::string pduSessResSetupRespTransfer;
}PduSessResSetupRespItem;

typedef std::vector<PduSessResSetupRespItem> PduSessResSetupRespList;

//Pdu session resource failed to setup Item
typedef struct PduSessResFailedSetupItem
{
        Uint8 pduSessId;
        std::string pduSessResSetupUnSussTransfer;
}PduSessResFailedSetupItem;

typedef std::vector<PduSessResFailedSetupItem> PduSessResFailedSetupList;

typedef struct AddQosFlowPerTNLInfo
{
    GtpTunnelInfo        upTransportLayerInfo;
    QosFlowIndList       assQosFlowList;
}AddQosFlowPerTNLInfo;

typedef std::vector<AddQosFlowPerTNLInfo> AddQosFlowPerTNLInfoList;

typedef std::vector<long> QosFlowRespList;

//Pdu Session Resource Modify Response Item
typedef struct PduSessResModifyRespItem
{
    Uint8       pduSessId;
    bool        pduSessResMdfyRespTransferPrst;
    std::string pduSessResMdfyRespTransfer;
}PduSessResModifyRespItem;

typedef std::vector<PduSessResModifyRespItem> PduSessResModifyRespList;

//Pdu Session Resource Failed to Modify Response Item
typedef struct PduSessResFailedMdfyRespItem
{
    Uint32 pduSessId;
    std::string pduSessResFailedMdfyRespTransfer;
}PduSessResFailedMdfyRespItem;

typedef std::vector<PduSessResFailedMdfyRespItem> PduSessResFailedModifyList;

typedef std::vector<Uint32> QosFlowAddOrMdfyRespList;

typedef struct RecommandCellItem
{
    bool  timeStayedPrst;
        Uint16 timeStayedInCell;
        NrCgi nrcgi;
}RecommandCellItem;
typedef std::vector<RecommandCellItem> RecommandCellList;

typedef struct RecommandRanNodeItem
{
        Uint8   TargetChoice;
        GGnbId  ggnbId;
        TaiType tai;
}RecommandRanNodeItem;
typedef std::vector<RecommandRanNodeItem> RecommandRanNodeList;

typedef struct PduSessResRelCmdItem
{
        Uint8 pduSessId;
        std::string pduSessResRelCmdTransfer;
}PduSessResRelCmdItem;
typedef std::vector<PduSessResRelCmdItem> PduSessResRelCmdList;

typedef struct PduSessResRelRespItem
{
        Uint8 pduSessId;
        std::string pduSessResRelRespTransfer;
}PduSessResRelRespItem;
typedef std::vector<PduSessResRelRespItem> PduSessResRelRespList;

typedef struct PduSessResRelCplItem
{
    Uint8 pduSessId;
    bool pduSessResRelRespTransferPrst;
    std::string pduSessResRelRespTransfer;
}PduSessResRelCplItem;
typedef std::vector<PduSessResRelCplItem> PduSessResRelCplList;

typedef struct UePagingIdentity
{
        Uint8 setId[2];
        Uint8 pointer;
        Uint8 tmsi[4];
}UePagingIdentity;

typedef struct PagingAttemptInfo
{
        Uint8 count;
        Uint8 intendedNum;
        Uint8 nextAreaScope;
        bool  nextAreaScopePrst;
}PagingAttemptInfo;

typedef struct AssDateForPaging
{
        RecommandCellList  recommandCellList;
        bool               recommandCellListPrst;
        PagingAttemptInfo  pagingAttemptInfo;
        bool               pagingAttemptInfoPrst;
}AssDateForPaging;

//Tai List For Paging
// typedef struct TaiPagingItem
// {
//      TaiType tai;
// }TaiPagingItem;
typedef std::vector<TaiType> TaiPagingList;

typedef struct UeRadioCapForPaging
{
    std::string ueRadioCapForPagingNr;
    bool        ueRadioCapForPagingNrPrst;
    std::string ueRadioCapForPagingEutra;
    bool        ueRadioCapForPagingEutraPrst;
}UeRadioCapForPaging;

typedef struct UeAssLogicalNgConn
{
    Uint64 amfUeNgapId;
    bool   amfUeNgapIdPrst;
    Uint32 ranUeNgapId;
    bool   ranUeNgapIdPrst;
}UeAssLogicalNgConn;
typedef std::vector<UeAssLogicalNgConn> UeAssLogicalNgConnList;

typedef struct ResetTypeChoice
{
    Uint8                  ngInterface;
    bool                   ngInterfacePrst;
    UeAssLogicalNgConnList ueAssLogicalNgConnList;
    bool                   ueAssLogicalNgConnListPrst;
}ResetTypeChoice;

typedef struct VolumeTimeReport{
    Uint8 startTimeStamp[4];
    Uint8 endTimeStamp[4];
    Uint64 usageCountUL;
    Uint64 usageCountDL;
}VolumeTimeReport;
typedef std::vector<VolumeTimeReport> VolumeTimeReportList; 

typedef struct PduSessUsageReport{
    Uint8 ratType;
    VolumeTimeReportList volumeTimeReportList;
}PduSessUsageReport;

typedef struct QosFlowUsageReport{
    Uint32 qosFlowId;
    Uint8  ratType;
    VolumeTimeReportList qosFlowsTimeReportList;
}QosFlowUsageReport;

typedef std::vector<QosFlowUsageReport> QosFlowUsageReportList;

typedef struct SecRatUsageInformation
{
    bool                   pduSessUsageReportPrst;
    PduSessUsageReport     pduSessUsageReport;
    bool                   qosFlowUsageReportListPrst;
    QosFlowUsageReportList qosFlowUsageReportList; 
}SecRatUsageInformation;

typedef struct PduSessResModReqItem
{
    Uint8 pduSessId;
    bool nasPduPrst;
    std::string nasPdu;
    std::string pduSessResModReqTransfer;
    bool snssaiPrst;
    SNssai snssai;
}PduSessResModReqItem;
typedef std::vector<PduSessResModReqItem> PduSessResModReqList;


typedef struct QosFlowAddOrModReqItem
{
    Uint32 qosFlowInd;
    bool qosFlowParaPresent;
    QosFlowLevelQosPara qosFlowLevQosPara;
    bool eRABIdPresent;
    Uint32 eRABId;
}QosFlowAddOrModReqItem;
typedef std::vector<QosFlowAddOrModReqItem> QosFlowAddOrModReqList;

#endif
