%module ngapcodec
%{
#include "ngapCodec.h"
#include "ngSetupRequestCodec.h"
#include "ngSetupResponseCodec.h"
#include "ngSetupFailureCodec.h"
#include "ngResetCodec.h"
#include "ngResetAckCodec.h"
#include "initialUEMessageCodec.h"
#include "downlinkNASTransportCodec.h"
#include "uplinkNASTransportCodec.h"
#include "initialContextSetupRequestCodec.h"
#include "initialContextSetupResponseCodec.h"
#include "initialContextSetupFailureCodec.h"
#include "pduSessionResourceSetupRequestTransferCodec.h"
#include "pduSessionResourceSetupResponseTransferCodec.h"
#include "pduSessionResourceSetupRequestCodec.h"
#include "pduSessionResourceSetupUnSuccTransferCodec.h"
#include "pduSessionResourceSetupResponseCodec.h"
#include "ueContextReleaseRequestCodec.h"
#include "ueContextReleaseCommandCodec.h"
#include "ueContextReleaseCompleteCodec.h"
#include "ueRadioCapabilityInfoIndicationCodec.h"
#include "pduSessionResourceReleaseCommandTransferCodec.h"
#include "pduSessionResourceReleaseCommandCodec.h"
#include "pduSessionResourceReleaseResponseCodec.h"
#include "errorIndicationCodec.h"
#include "pagingCodec.h"
#include "pduSessionResourceReleaseResponseTransferCodec.h"
#include "pduSessionResourceModifyResponseCodec.h"
#include "pduSessionResourceModifyResponseTransferCodec.h"
#include "pduSessionResourceModifyUnSuccTransferCodec.h"
#include "pduSessionResourceModifyRequestCodec.h"
#include "pduSessionResourceModifyRequestTransferCodec.h"
#include "nasNonDeliveryIndicationCodec.h"
#include "snssaiCodec.h"
%}

%include <typemaps.i>
%include "std_string.i"
%include "std_vector.i"


// This will create 2 wrapped types in Go called
// "StringVector" and "ByteVector" for their respective
// types.
namespace std {
   %template(StringVector) vector<string>;
   %template(ByteVector) vector<char>;
   %template(LongVector) vector<long>;
   %template(UShortVector) vector<unsigned short>;
   %template(UIntVector) vector<unsigned int>;
}

/* Let's just grab the original header file here */
%include "../ngaposs/msgTypes/ngapMsgTypes.h"
%include "ngapCodec.h"
%include "ngSetupRequestCodec.h"
%include "ngSetupResponseCodec.h"
%include "ngSetupFailureCodec.h"
%include "ngResetCodec.h"
%include "ngResetAckCodec.h"
%include "initialUEMessageCodec.h"
%include "downlinkNASTransportCodec.h"
%include "uplinkNASTransportCodec.h"
%include "initialContextSetupRequestCodec.h"
%include "initialContextSetupResponseCodec.h"
%include "initialContextSetupFailureCodec.h"
%include "pduSessionResourceSetupRequestTransferCodec.h"
%include "pduSessionResourceSetupResponseTransferCodec.h"
%include "pduSessionResourceSetupRequestCodec.h"
%include "pduSessionResourceSetupUnSuccTransferCodec.h"
%include "pduSessionResourceSetupResponseCodec.h"
%include "ueContextReleaseRequestCodec.h"
%include "ueContextReleaseCommandCodec.h"
%include "ueContextReleaseCompleteCodec.h"
%include "ueRadioCapabilityInfoIndicationCodec.h"
%include "pduSessionResourceReleaseCommandTransferCodec.h"
%include "pduSessionResourceReleaseCommandCodec.h"
%include "pduSessionResourceReleaseResponseCodec.h"
%include "errorIndicationCodec.h"
%include "pagingCodec.h"
%include "pduSessionResourceReleaseResponseTransferCodec.h"
%include "pduSessionResourceModifyResponseCodec.h"
%include "pduSessionResourceModifyUnSuccTransferCodec.h"
%include "pduSessionResourceModifyResponseTransferCodec.h"
%include "pduSessionResourceModifyRequestCodec.h"
%include "pduSessionResourceModifyRequestTransferCodec.h"
%include "nasNonDeliveryIndicationCodec.h"
%include "snssaiCodec.h"

namespace std{
    %template(SNssaiVector) vector<SNssai>;
    %template(BPlmnItemVector) vector<BPlmnItem>;
    %template(SupTAItemVector) vector<SupTAItem>;
    %template(ServedGuamiItemVector) vector<ServedGuamiItem>;
    %template(QosFlowSetupReqItemVector) vector<QosFlowSetupReqItem>;
    %template(PduSessResSetupReqItemVector) vector<PduSessResSetupReqItem>;
    %template(PduSessResModifyReqItemVector) vector<PduSessResModReqItem>;
    %template(QosFlowCodecItemVector) vector<QosFlowCodecItem>;
    %template(PduSessResSetupRespItemVector) vector<PduSessResSetupRespItem>;
    %template(PduSessResFailedSetupItemVector) vector<PduSessResFailedSetupItem>;
    %template(RecommandCellItemVector) vector<RecommandCellItem>;
    %template(RecommandRanNodeItemVector) vector<RecommandRanNodeItem>;
    %template(PduSessResRelCmdItemVector) vector<PduSessResRelCmdItem>;
    %template(PduSessResRelRespItemVector) vector<PduSessResRelRespItem>;
    %template(TaiTypeVector) vector<TaiType>;
    %template(AssQosFlowItemVector) vector<AssQosFlowItem>;
    %template(UeAssLogicalNgConnVector) vector<UeAssLogicalNgConn>;
    %template(VolumeTimeReportVector) vector<VolumeTimeReport>;
    %template(QosFlowUsageReportVector) vector<QosFlowUsageReport>;
    %template(PduSessResModifyRespItemVector) vector<PduSessResModifyRespItem>;
    %template(PduSessResFailedMdfyRespItemVector) vector<PduSessResFailedMdfyRespItem>;
    %template(QosFlowAddOrModReqVector) vector<QosFlowAddOrModReqItem>;
    %template(UlNguUpTnlModifyVector) vector<UlNguUpTnlModifyItem>;
    %template(GtpTunnelVector) vector<GtpTunnel>;
    %template(AddQosFlowPerTNLInfoVector) vector<AddQosFlowPerTNLInfo>;
}
