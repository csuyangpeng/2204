/*************************************************************/
/* Copyright (C) 2020 OSS Nokalva, Inc.  All rights reserved.*/
/*************************************************************/

/* THIS FILE IS PROPRIETARY MATERIAL OF OSS NOKALVA, INC.
 * AND MAY BE USED ONLY BY DIRECT LICENSEES OF OSS NOKALVA, INC.
 * THIS FILE MAY NOT BE DISTRIBUTED.
 * THIS COPYRIGHT STATEMENT MAY NOT BE REMOVED. */

/* Generated for: Keybridge Technology (Hong Kong) Co. Ltd, Guanghhzou, Guangdong, China - Project-based, License 69615Z 69615Z,
 * only for project "Keybridge Technology LTE signaling collection and analysis system". */
/* Abstract syntax: ngap_procedures */
/* Created: Thu Nov  5 02:27:06 2020 */
/* ASN.1 compiler version: 10.1.1 */
/* Code generated for runtime version 10.1.1 or later */
/* Compiler operating system: Linux with 64-bit support */
/* Compiler machine type: AMD64 */
/* Target operating system: Linux with 64-bit support */
/* Target machine type: AMD64 */
/* C compiler options required: -m64 */
/* ASN.1 compiler options and file names specified:
 * -codefile ngapToed_x86.cc -output ngapToed_x86 -constraints -c++ -2008 -per
 * -autoencdec
 * /usr/local/ossasn1/linux-glibc2.3-amd64/10.1.1/asn1dflt.linux-amd64
 * ngap-common.asn ngap-constants.asn ngap-containers.asn ngap-ies.asn
 * ngap-pdus.asn ngap-procedures.asn
 */

#ifndef OSS_ngap_procedures
#define OSS_ngap_procedures

#ifdef __cplusplus
extern "C" {
#endif /* __cplusplus */

// 9.4.3	Elementary Procedure Definitions
// ASN1START
// **************************************************************
//
// Elementary Procedure definitions
//
// **************************************************************

#include "ossasn1.h"

#define          AdditionalDLUPTNLInformationForHOList_PDU 1
#define          AllowedNSSAI_PDU 2
#define          AMFName_PDU 3
#define          AMFSetID_PDU 4
#define          AMF_TNLAssociationSetupList_PDU 5
#define          AMF_TNLAssociationToAddList_PDU 6
#define          AMF_TNLAssociationToRemoveList_PDU 7
#define          AMF_TNLAssociationToUpdateList_PDU 8
#define          AMF_UE_NGAP_ID_PDU 9
#define          AssistanceDataForPaging_PDU 10
#define          AssociatedQosFlowList_PDU 11
#define          BroadcastCancelledAreaList_PDU 12
#define          BroadcastCompletedAreaList_PDU 13
#define          CancelAllWarningMessages_PDU 14
#define          Cause_PDU 15
#define          CellIDListForRestart_PDU 16
#define          CNAssistedRANTuning_PDU 17
#define          CNTypeRestrictionsForEquivalent_PDU 18
#define          CNTypeRestrictionsForServing_PDU 19
#define          CommonNetworkInstance_PDU 20
#define          ConcurrentWarningMessageInd_PDU 21
#define          CoreNetworkAssistanceInformationForInactive_PDU 22
#define          CriticalityDiagnostics_PDU 23
#define          DataCodingScheme_PDU 24
#define          DataForwardingNotPossible_PDU 25
#define          DirectForwardingPathAvailability_PDU 26
#define          EmergencyAreaIDListForRestart_PDU 27
#define          EmergencyFallbackIndicator_PDU 28
#define          EN_DCSONConfigurationTransfer_PDU 29
#define          FiveG_S_TMSI_PDU 30
#define          GlobalRANNodeID_PDU 31
#define          GUAMI_PDU 32
#define          HandoverCommandTransfer_PDU 33
#define          HandoverFlag_PDU 34
#define          HandoverPreparationUnsuccessfulTransfer_PDU 35
#define          HandoverRequestAcknowledgeTransfer_PDU 36
#define          HandoverRequiredTransfer_PDU 37
#define          HandoverResourceAllocationUnsuccessfulTransfer_PDU 38
#define          HandoverType_PDU 39
#define          IMSVoiceSupportIndicator_PDU 40
#define          IndexToRFSP_PDU 41
#define          InfoOnRecommendedCellsAndRANNodesForPaging_PDU 42
#define          LocationReportingRequestType_PDU 43
#define          MaskedIMEISV_PDU 44
#define          MessageIdentifier_PDU 45
#define          MaximumIntegrityProtectedDataRate_PDU 46
#define          MobilityRestrictionList_PDU 47
#define          NAS_PDU_PDU 48
#define          NASSecurityParametersFromNGRAN_PDU 49
#define          NetworkInstance_PDU 50
#define          NewSecurityContextInd_PDU 51
#define          NGRAN_CGI_PDU 52
#define          NGRANTraceID_PDU 53
#define          NRPPa_PDU_PDU 54
#define          NumberOfBroadcastsRequested_PDU 55
#define          OverloadResponse_PDU 56
#define          OverloadStartNSSAIList_PDU 57
#define          PagingDRX_PDU 58
#define          PagingOrigin_PDU 59
#define          PagingPriority_PDU 60
#define          PathSwitchRequestAcknowledgeTransfer_PDU 61
#define          PathSwitchRequestSetupFailedTransfer_PDU 62
#define          PathSwitchRequestTransfer_PDU 63
#define          PathSwitchRequestUnsuccessfulTransfer_PDU 64
#define          PDUSessionAggregateMaximumBitRate_PDU 65
#define          PDUSessionResourceAdmittedList_PDU 66
#define          PDUSessionResourceFailedToModifyListModCfm_PDU 67
#define          PDUSessionResourceFailedToModifyListModRes_PDU 68
#define          PDUSessionResourceFailedToSetupListCxtFail_PDU 69
#define          PDUSessionResourceFailedToSetupListCxtRes_PDU 70
#define          PDUSessionResourceFailedToSetupListHOAck_PDU 71
#define          PDUSessionResourceFailedToSetupListPSReq_PDU 72
#define          PDUSessionResourceFailedToSetupListSURes_PDU 73
#define          PDUSessionResourceHandoverList_PDU 74
#define          PDUSessionResourceListCxtRelCpl_PDU 75
#define          PDUSessionResourceListCxtRelReq_PDU 76
#define          PDUSessionResourceListHORqd_PDU 77
#define          PDUSessionResourceModifyConfirmTransfer_PDU 78
#define          PDUSessionResourceModifyIndicationUnsuccessfulTransfer_PDU 79
#define          PDUSessionResourceModifyRequestTransfer_PDU 80
#define          PDUSessionResourceModifyResponseTransfer_PDU 81
#define          PDUSessionResourceModifyIndicationTransfer_PDU 82
#define          PDUSessionResourceModifyListModCfm_PDU 83
#define          PDUSessionResourceModifyListModInd_PDU 84
#define          PDUSessionResourceModifyListModReq_PDU 85
#define          PDUSessionResourceModifyListModRes_PDU 86
#define          PDUSessionResourceModifyUnsuccessfulTransfer_PDU 87
#define          PDUSessionResourceNotifyList_PDU 88
#define          PDUSessionResourceNotifyReleasedTransfer_PDU 89
#define          PDUSessionResourceNotifyTransfer_PDU 90
#define          PDUSessionResourceReleaseCommandTransfer_PDU 91
#define          PDUSessionResourceReleasedListNot_PDU 92
#define          PDUSessionResourceReleasedListPSAck_PDU 93
#define          PDUSessionResourceReleasedListPSFail_PDU 94
#define          PDUSessionResourceReleasedListRelRes_PDU 95
#define          PDUSessionResourceReleaseResponseTransfer_PDU 96
#define          PDUSessionResourceSecondaryRATUsageList_PDU 97
#define          PDUSessionResourceSetupListCxtReq_PDU 98
#define          PDUSessionResourceSetupListCxtRes_PDU 99
#define          PDUSessionResourceSetupListHOReq_PDU 100
#define          PDUSessionResourceSetupListSUReq_PDU 101
#define          PDUSessionResourceSetupListSURes_PDU 102
#define          PDUSessionResourceSetupRequestTransfer_PDU 103
#define          PDUSessionResourceSetupResponseTransfer_PDU 104
#define          PDUSessionResourceSetupUnsuccessfulTransfer_PDU 105
#define          PDUSessionResourceSwitchedList_PDU 106
#define          PDUSessionResourceToBeSwitchedDLList_PDU 107
#define          PDUSessionResourceToReleaseListHOCmd_PDU 108
#define          PDUSessionResourceToReleaseListRelCmd_PDU 109
#define          PDUSessionType_PDU 110
#define          PLMNIdentity_PDU 111
#define          PLMNSupportList_PDU 112
#define          PWSFailedCellIDList_PDU 113
#define          QosFlowAddOrModifyRequestList_PDU 114
#define          QosFlowListWithCause_PDU 115
#define          QosFlowPerTNLInformationList_PDU 116
#define          QosFlowSetupRequestList_PDU 117
#define          RANNodeName_PDU 118
#define          RANPagingPriority_PDU 119
#define          RANStatusTransfer_TransparentContainer_PDU 120
#define          RAN_UE_NGAP_ID_PDU 121
#define          RedirectionVoiceFallback_PDU 122
#define          RelativeAMFCapacity_PDU 123
#define          RepetitionPeriod_PDU 124
#define          ResetType_PDU 125
#define          RoutingID_PDU 126
#define          RRCEstablishmentCause_PDU 127
#define          RRCInactiveTransitionReportRequest_PDU 128
#define          RRCState_PDU 129
#define          SecondaryRATUsageInformation_PDU 130
#define          SecondaryRATDataUsageReportTransfer_PDU 131
#define          SecurityContext_PDU 132
#define          SecurityIndication_PDU 133
#define          SecurityKey_PDU 134
#define          SecurityResult_PDU 135
#define          SerialNumber_PDU 136
#define          ServedGUAMIList_PDU 137
#define          S_NSSAI_PDU 138
#define          SONConfigurationTransfer_PDU 139
#define          SourceToTarget_TransparentContainer_PDU 140
#define          SupportedTAList_PDU 141
#define          TAIListForPaging_PDU 142
#define          TAIListForRestart_PDU 143
#define          TargetID_PDU 144
#define          TargetToSource_TransparentContainer_PDU 145
#define          TimeToWait_PDU 146
#define          TNLAssociationList_PDU 147
#define          TraceActivation_PDU 148
#define          TrafficLoadReductionIndication_PDU 149
#define          TransportLayerAddress_PDU 150
#define          UEAggregateMaximumBitRate_PDU 151
#define          UE_associatedLogicalNG_connectionList_PDU 152
#define          UEContextRequest_PDU 153
#define          UE_NGAP_IDs_PDU 154
#define          UEPagingIdentity_PDU 155
#define          UEPresenceInAreaOfInterestList_PDU 156
#define          UERadioCapability_PDU 157
#define          UERadioCapabilityForPaging_PDU 158
#define          UERetentionInformation_PDU 159
#define          UESecurityCapabilities_PDU 160
#define          UL_NGU_UP_TNLModifyList_PDU 161
#define          UnavailableGUAMIList_PDU 162
#define          UPTransportLayerInformation_PDU 163
#define          UPTransportLayerInformationList_PDU 164
#define          UPTransportLayerInformationPairList_PDU 165
#define          UserLocationInformation_PDU 166
#define          WarningAreaCoordinates_PDU 167
#define          WarningAreaList_PDU 168
#define          WarningMessageContents_PDU 169
#define          WarningSecurityInfo_PDU 170
#define          WarningType_PDU 171
#define          PDUSessionResourceSetupRequest_PDU 172
#define          PDUSessionResourceSetupResponse_PDU 173
#define          PDUSessionResourceReleaseCommand_PDU 174
#define          PDUSessionResourceReleaseResponse_PDU 175
#define          PDUSessionResourceModifyRequest_PDU 176
#define          PDUSessionResourceModifyResponse_PDU 177
#define          PDUSessionResourceNotify_PDU 178
#define          PDUSessionResourceModifyIndication_PDU 179
#define          PDUSessionResourceModifyConfirm_PDU 180
#define          InitialContextSetupRequest_PDU 181
#define          InitialContextSetupResponse_PDU 182
#define          InitialContextSetupFailure_PDU 183
#define          UEContextReleaseRequest_PDU 184
#define          UEContextReleaseCommand_PDU 185
#define          UEContextReleaseComplete_PDU 186
#define          UEContextModificationRequest_PDU 187
#define          UEContextModificationResponse_PDU 188
#define          UEContextModificationFailure_PDU 189
#define          RRCInactiveTransitionReport_PDU 190
#define          HandoverRequired_PDU 191
#define          HandoverCommand_PDU 192
#define          HandoverPreparationFailure_PDU 193
#define          HandoverRequest_PDU 194
#define          HandoverRequestAcknowledge_PDU 195
#define          HandoverFailure_PDU 196
#define          HandoverNotify_PDU 197
#define          PathSwitchRequest_PDU 198
#define          PathSwitchRequestAcknowledge_PDU 199
#define          PathSwitchRequestFailure_PDU 200
#define          HandoverCancel_PDU 201
#define          HandoverCancelAcknowledge_PDU 202
#define          UplinkRANStatusTransfer_PDU 203
#define          DownlinkRANStatusTransfer_PDU 204
#define          Paging_PDU 205
#define          InitialUEMessage_PDU 206
#define          DownlinkNASTransport_PDU 207
#define          UplinkNASTransport_PDU 208
#define          NASNonDeliveryIndication_PDU 209
#define          RerouteNASRequest_PDU 210
#define          NGSetupRequest_PDU 211
#define          NGSetupResponse_PDU 212
#define          NGSetupFailure_PDU 213
#define          RANConfigurationUpdate_PDU 214
#define          RANConfigurationUpdateAcknowledge_PDU 215
#define          RANConfigurationUpdateFailure_PDU 216
#define          AMFConfigurationUpdate_PDU 217
#define          AMFConfigurationUpdateAcknowledge_PDU 218
#define          AMFConfigurationUpdateFailure_PDU 219
#define          AMFStatusIndication_PDU 220
#define          NGReset_PDU 221
#define          NGResetAcknowledge_PDU 222
#define          ErrorIndication_PDU 223
#define          OverloadStart_PDU 224
#define          OverloadStop_PDU 225
#define          UplinkRANConfigurationTransfer_PDU 226
#define          DownlinkRANConfigurationTransfer_PDU 227
#define          WriteReplaceWarningRequest_PDU 228
#define          WriteReplaceWarningResponse_PDU 229
#define          PWSCancelRequest_PDU 230
#define          PWSCancelResponse_PDU 231
#define          PWSRestartIndication_PDU 232
#define          PWSFailureIndication_PDU 233
#define          DownlinkUEAssociatedNRPPaTransport_PDU 234
#define          UplinkUEAssociatedNRPPaTransport_PDU 235
#define          DownlinkNonUEAssociatedNRPPaTransport_PDU 236
#define          UplinkNonUEAssociatedNRPPaTransport_PDU 237
#define          TraceStart_PDU 238
#define          TraceFailureIndication_PDU 239
#define          DeactivateTrace_PDU 240
#define          CellTrafficTrace_PDU 241
#define          LocationReportingControl_PDU 242
#define          LocationReportingFailureIndication_PDU 243
#define          LocationReport_PDU 244
#define          UETNLABindingReleaseRequest_PDU 245
#define          UERadioCapabilityInfoIndication_PDU 246
#define          UERadioCapabilityCheckRequest_PDU 247
#define          UERadioCapabilityCheckResponse_PDU 248
#define          PrivateMessage_PDU 249
#define          SecondaryRATDataUsageReport_PDU 250
#define          NGAP_PDU_PDU 251
#define          PDUSessionResourceItemCxtRelCpl_ExtIEs_EXTENSION_PDU 252
#define          RerouteNASRequest_IEs_TYPE_PDU 253
#define          AdditionalDLUPTNLInformationForHOItem_ExtIEs_OSET 1
#define          AllocationAndRetentionPriority_ExtIEs_OSET 2
#define          AllowedNSSAI_Item_ExtIEs_OSET 3
#define          AMFPagingTarget_ExtIEs_OSET 4
#define          AMF_TNLAssociationSetupItem_ExtIEs_OSET 5
#define          AMF_TNLAssociationToAddItem_ExtIEs_OSET 6
#define          AMF_TNLAssociationToRemoveItem_ExtIEs_OSET 7
#define          AMF_TNLAssociationToUpdateItem_ExtIEs_OSET 8
#define          AreaOfInterest_ExtIEs_OSET 9
#define          AreaOfInterestCellItem_ExtIEs_OSET 10
#define          AreaOfInterestItem_ExtIEs_OSET 11
#define          AreaOfInterestRANNodeItem_ExtIEs_OSET 12
#define          AreaOfInterestTAIItem_ExtIEs_OSET 13
#define          AssistanceDataForPaging_ExtIEs_OSET 14
#define          AssistanceDataForRecommendedCells_ExtIEs_OSET 15
#define          AssociatedQosFlowItem_ExtIEs_OSET 16
#define          BroadcastCancelledAreaList_ExtIEs_OSET 17
#define          BroadcastCompletedAreaList_ExtIEs_OSET 18
#define          BroadcastPLMNItem_ExtIEs_OSET 19
#define          CancelledCellsInEAI_EUTRA_Item_ExtIEs_OSET 20
#define          CancelledCellsInEAI_NR_Item_ExtIEs_OSET 21
#define          CancelledCellsInTAI_EUTRA_Item_ExtIEs_OSET 22
#define          CancelledCellsInTAI_NR_Item_ExtIEs_OSET 23
#define          Cause_ExtIEs_OSET 24
#define          CellIDBroadcastEUTRA_Item_ExtIEs_OSET 25
#define          CellIDBroadcastNR_Item_ExtIEs_OSET 26
#define          CellIDCancelledEUTRA_Item_ExtIEs_OSET 27
#define          CellIDCancelledNR_Item_ExtIEs_OSET 28
#define          CellIDListForRestart_ExtIEs_OSET 29
#define          CNAssistedRANTuning_ExtIEs_OSET 30
#define          CNTypeRestrictionsForEquivalentItem_ExtIEs_OSET 31
#define          CompletedCellsInEAI_EUTRA_Item_ExtIEs_OSET 32
#define          CompletedCellsInEAI_NR_Item_ExtIEs_OSET 33
#define          CompletedCellsInTAI_EUTRA_Item_ExtIEs_OSET 34
#define          CompletedCellsInTAI_NR_Item_ExtIEs_OSET 35
#define          CoreNetworkAssistanceInformationForInactive_ExtIEs_OSET 36
#define          COUNTValueForPDCP_SN12_ExtIEs_OSET 37
#define          COUNTValueForPDCP_SN18_ExtIEs_OSET 38
#define          CPTransportLayerInformation_ExtIEs_OSET 39
#define          CriticalityDiagnostics_ExtIEs_OSET 40
#define          CriticalityDiagnostics_IE_Item_ExtIEs_OSET 41
#define          DataForwardingResponseDRBItem_ExtIEs_OSET 42
#define          DRBsSubjectToStatusTransferItem_ExtIEs_OSET 43
#define          DRBStatusDL_ExtIEs_OSET 44
#define          DRBStatusDL12_ExtIEs_OSET 45
#define          DRBStatusDL18_ExtIEs_OSET 46
#define          DRBStatusUL_ExtIEs_OSET 47
#define          DRBStatusUL12_ExtIEs_OSET 48
#define          DRBStatusUL18_ExtIEs_OSET 49
#define          Dynamic5QIDescriptor_ExtIEs_OSET 50
#define          EmergencyAreaIDBroadcastEUTRA_Item_ExtIEs_OSET 51
#define          EmergencyAreaIDBroadcastNR_Item_ExtIEs_OSET 52
#define          EmergencyAreaIDCancelledEUTRA_Item_ExtIEs_OSET 53
#define          EmergencyAreaIDCancelledNR_Item_ExtIEs_OSET 54
#define          EmergencyFallbackIndicator_ExtIEs_OSET 55
#define          EPS_TAI_ExtIEs_OSET 56
#define          EUTRA_CGI_ExtIEs_OSET 57
#define          ExpectedUEActivityBehaviour_ExtIEs_OSET 58
#define          ExpectedUEBehaviour_ExtIEs_OSET 59
#define          ExpectedUEMovingTrajectoryItem_ExtIEs_OSET 60
#define          FiveG_S_TMSI_ExtIEs_OSET 61
#define          ForbiddenAreaInformation_Item_ExtIEs_OSET 62
#define          GBR_QosInformation_ExtIEs_OSET 63
#define          GlobalGNB_ID_ExtIEs_OSET 64
#define          GlobalN3IWF_ID_ExtIEs_OSET 65
#define          GlobalNgENB_ID_ExtIEs_OSET 66
#define          GlobalRANNodeID_ExtIEs_OSET 67
#define          GNB_ID_ExtIEs_OSET 68
#define          GTPTunnel_ExtIEs_OSET 69
#define          GUAMI_ExtIEs_OSET 70
#define          HandoverCommandTransfer_ExtIEs_OSET 71
#define          HandoverPreparationUnsuccessfulTransfer_ExtIEs_OSET 72
#define          HandoverRequestAcknowledgeTransfer_ExtIEs_OSET 73
#define          HandoverRequiredTransfer_ExtIEs_OSET 74
#define          HandoverResourceAllocationUnsuccessfulTransfer_ExtIEs_OSET 75
#define          InfoOnRecommendedCellsAndRANNodesForPaging_ExtIEs_OSET 76
#define          LocationReportingRequestType_ExtIEs_OSET 77
#define          MobilityRestrictionList_ExtIEs_OSET 78
#define          N3IWF_ID_ExtIEs_OSET 79
#define          NgENB_ID_ExtIEs_OSET 80
#define          NGRAN_CGI_ExtIEs_OSET 81
#define          NonDynamic5QIDescriptor_ExtIEs_OSET 82
#define          NR_CGI_ExtIEs_OSET 83
#define          OverloadResponse_ExtIEs_OSET 84
#define          OverloadStartNSSAIItem_ExtIEs_OSET 85
#define          PacketErrorRate_ExtIEs_OSET 86
#define          PagingAttemptInformation_ExtIEs_OSET 87
#define          PathSwitchRequestAcknowledgeTransfer_ExtIEs_OSET 88
#define          PathSwitchRequestSetupFailedTransfer_ExtIEs_OSET 89
#define          PathSwitchRequestTransfer_ExtIEs_OSET 90
#define          PathSwitchRequestUnsuccessfulTransfer_ExtIEs_OSET 91
#define          PDUSessionAggregateMaximumBitRate_ExtIEs_OSET 92
#define          PDUSessionResourceAdmittedItem_ExtIEs_OSET 93
#define          PDUSessionResourceFailedToModifyItemModCfm_ExtIEs_OSET 94
#define          PDUSessionResourceFailedToModifyItemModRes_ExtIEs_OSET 95
#define          PDUSessionResourceFailedToSetupItemCxtFail_ExtIEs_OSET 96
#define          PDUSessionResourceFailedToSetupItemCxtRes_ExtIEs_OSET 97
#define          PDUSessionResourceFailedToSetupItemHOAck_ExtIEs_OSET 98
#define          PDUSessionResourceFailedToSetupItemPSReq_ExtIEs_OSET 99
#define          PDUSessionResourceFailedToSetupItemSURes_ExtIEs_OSET 100
#define          PDUSessionResourceHandoverItem_ExtIEs_OSET 101
#define          PDUSessionResourceItemCxtRelCpl_ExtIEs_OSET 102
#define          PDUSessionResourceItemCxtRelReq_ExtIEs_OSET 103
#define          PDUSessionResourceItemHORqd_ExtIEs_OSET 104
#define          PDUSessionResourceModifyConfirmTransfer_ExtIEs_OSET 105
#define          PDUSessionResourceModifyIndicationUnsuccessfulTransfer_ExtIEs_OSET 106
#define          PDUSessionResourceModifyRequestTransferIEs_OSET 107
#define          PDUSessionResourceModifyResponseTransfer_ExtIEs_OSET 108
#define          PDUSessionResourceModifyIndicationTransfer_ExtIEs_OSET 109
#define          PDUSessionResourceModifyItemModCfm_ExtIEs_OSET 110
#define          PDUSessionResourceModifyItemModInd_ExtIEs_OSET 111
#define          PDUSessionResourceModifyItemModReq_ExtIEs_OSET 112
#define          PDUSessionResourceModifyItemModRes_ExtIEs_OSET 113
#define          PDUSessionResourceModifyUnsuccessfulTransfer_ExtIEs_OSET 114
#define          PDUSessionResourceNotifyItem_ExtIEs_OSET 115
#define          PDUSessionResourceNotifyReleasedTransfer_ExtIEs_OSET 116
#define          PDUSessionResourceNotifyTransfer_ExtIEs_OSET 117
#define          PDUSessionResourceReleaseCommandTransfer_ExtIEs_OSET 118
#define          PDUSessionResourceReleasedItemNot_ExtIEs_OSET 119
#define          PDUSessionResourceReleasedItemPSAck_ExtIEs_OSET 120
#define          PDUSessionResourceReleasedItemPSFail_ExtIEs_OSET 121
#define          PDUSessionResourceReleasedItemRelRes_ExtIEs_OSET 122
#define          PDUSessionResourceReleaseResponseTransfer_ExtIEs_OSET 123
#define          PDUSessionResourceSecondaryRATUsageItem_ExtIEs_OSET 124
#define          PDUSessionResourceSetupItemCxtReq_ExtIEs_OSET 125
#define          PDUSessionResourceSetupItemCxtRes_ExtIEs_OSET 126
#define          PDUSessionResourceSetupItemHOReq_ExtIEs_OSET 127
#define          PDUSessionResourceSetupItemSUReq_ExtIEs_OSET 128
#define          PDUSessionResourceSetupItemSURes_ExtIEs_OSET 129
#define          PDUSessionResourceSetupRequestTransferIEs_OSET 130
#define          PDUSessionResourceSetupResponseTransfer_ExtIEs_OSET 131
#define          PDUSessionResourceSetupUnsuccessfulTransfer_ExtIEs_OSET 132
#define          PDUSessionResourceSwitchedItem_ExtIEs_OSET 133
#define          PDUSessionResourceToBeSwitchedDLItem_ExtIEs_OSET 134
#define          PDUSessionResourceToReleaseItemHOCmd_ExtIEs_OSET 135
#define          PDUSessionResourceToReleaseItemRelCmd_ExtIEs_OSET 136
#define          PDUSessionUsageReport_ExtIEs_OSET 137
#define          PLMNSupportItem_ExtIEs_OSET 138
#define          PWSFailedCellIDList_ExtIEs_OSET 139
#define          QosCharacteristics_ExtIEs_OSET 140
#define          QosFlowAcceptedItem_ExtIEs_OSET 141
#define          QosFlowAddOrModifyRequestItem_ExtIEs_OSET 142
#define          QosFlowAddOrModifyResponseItem_ExtIEs_OSET 143
#define          QosFlowLevelQosParameters_ExtIEs_OSET 144
#define          QosFlowWithCauseItem_ExtIEs_OSET 145
#define          QosFlowModifyConfirmItem_ExtIEs_OSET 146
#define          QosFlowNotifyItem_ExtIEs_OSET 147
#define          QosFlowPerTNLInformation_ExtIEs_OSET 148
#define          QosFlowPerTNLInformationItem_ExtIEs_OSET 149
#define          QosFlowSetupRequestItem_ExtIEs_OSET 150
#define          QosFlowItemWithDataForwarding_ExtIEs_OSET 151
#define          QosFlowToBeForwardedItem_ExtIEs_OSET 152
#define          QoSFlowsUsageReport_Item_ExtIEs_OSET 153
#define          RANStatusTransfer_TransparentContainer_ExtIEs_OSET 154
#define          RATRestrictions_Item_ExtIEs_OSET 155
#define          RecommendedCellsForPaging_ExtIEs_OSET 156
#define          RecommendedCellItem_ExtIEs_OSET 157
#define          RecommendedRANNodesForPaging_ExtIEs_OSET 158
#define          RecommendedRANNodeItem_ExtIEs_OSET 159
#define          ResetType_ExtIEs_OSET 160
#define          SecondaryRATUsageInformation_ExtIEs_OSET 161
#define          SecondaryRATDataUsageReportTransfer_ExtIEs_OSET 162
#define          SecurityContext_ExtIEs_OSET 163
#define          SecurityIndication_ExtIEs_OSET 164
#define          SecurityResult_ExtIEs_OSET 165
#define          ServedGUAMIItem_ExtIEs_OSET 166
#define          ServiceAreaInformation_Item_ExtIEs_OSET 167
#define          SliceOverloadItem_ExtIEs_OSET 168
#define          SliceSupportItem_ExtIEs_OSET 169
#define          S_NSSAI_ExtIEs_OSET 170
#define          SONConfigurationTransfer_ExtIEs_OSET 171
#define          SONInformation_ExtIEs_OSET 172
#define          SONInformationReply_ExtIEs_OSET 173
#define          SourceRANNodeID_ExtIEs_OSET 174
#define          SupportedTAItem_ExtIEs_OSET 175
#define          TAI_ExtIEs_OSET 176
#define          TAIBroadcastEUTRA_Item_ExtIEs_OSET 177
#define          TAIBroadcastNR_Item_ExtIEs_OSET 178
#define          TAICancelledEUTRA_Item_ExtIEs_OSET 179
#define          TAICancelledNR_Item_ExtIEs_OSET 180
#define          TAIListForInactiveItem_ExtIEs_OSET 181
#define          TAIListForPagingItem_ExtIEs_OSET 182
#define          TargeteNB_ID_ExtIEs_OSET 183
#define          TargetID_ExtIEs_OSET 184
#define          TargetRANNodeID_ExtIEs_OSET 185
#define          TNLAssociationItem_ExtIEs_OSET 186
#define          TraceActivation_ExtIEs_OSET 187
#define          UEAggregateMaximumBitRate_ExtIEs_OSET 188
#define          UE_associatedLogicalNG_connectionItem_ExtIEs_OSET 189
#define          UEIdentityIndexValue_ExtIEs_OSET 190
#define          UE_NGAP_IDs_ExtIEs_OSET 191
#define          UE_NGAP_ID_pair_ExtIEs_OSET 192
#define          UEPagingIdentity_ExtIEs_OSET 193
#define          UEPresenceInAreaOfInterestItem_ExtIEs_OSET 194
#define          UERadioCapabilityForPaging_ExtIEs_OSET 195
#define          UESecurityCapabilities_ExtIEs_OSET 196
#define          UL_NGU_UP_TNLModifyItem_ExtIEs_OSET 197
#define          UnavailableGUAMIItem_ExtIEs_OSET 198
#define          UPTransportLayerInformation_ExtIEs_OSET 199
#define          UPTransportLayerInformationItem_ExtIEs_OSET 200
#define          UPTransportLayerInformationPairItem_ExtIEs_OSET 201
#define          UserLocationInformation_ExtIEs_OSET 202
#define          UserLocationInformationEUTRA_ExtIEs_OSET 203
#define          UserLocationInformationN3IWF_ExtIEs_OSET 204
#define          UserLocationInformationNR_ExtIEs_OSET 205
#define          UserPlaneSecurityInformation_ExtIEs_OSET 206
#define          VolumeTimedReport_Item_ExtIEs_OSET 207
#define          WarningAreaList_ExtIEs_OSET 208
#define          XnExtTLA_Item_ExtIEs_OSET 209
#define          XnTNLConfigurationInfo_ExtIEs_OSET 210
#define          PDUSessionResourceSetupRequestIEs_OSET 211
#define          PDUSessionResourceSetupResponseIEs_OSET 212
#define          PDUSessionResourceReleaseCommandIEs_OSET 213
#define          PDUSessionResourceReleaseResponseIEs_OSET 214
#define          PDUSessionResourceModifyRequestIEs_OSET 215
#define          PDUSessionResourceModifyResponseIEs_OSET 216
#define          PDUSessionResourceNotifyIEs_OSET 217
#define          PDUSessionResourceModifyIndicationIEs_OSET 218
#define          PDUSessionResourceModifyConfirmIEs_OSET 219
#define          InitialContextSetupRequestIEs_OSET 220
#define          InitialContextSetupResponseIEs_OSET 221
#define          InitialContextSetupFailureIEs_OSET 222
#define          UEContextReleaseRequest_IEs_OSET 223
#define          UEContextReleaseCommand_IEs_OSET 224
#define          UEContextReleaseComplete_IEs_OSET 225
#define          UEContextModificationRequestIEs_OSET 226
#define          UEContextModificationResponseIEs_OSET 227
#define          UEContextModificationFailureIEs_OSET 228
#define          RRCInactiveTransitionReportIEs_OSET 229
#define          HandoverRequiredIEs_OSET 230
#define          HandoverCommandIEs_OSET 231
#define          HandoverPreparationFailureIEs_OSET 232
#define          HandoverRequestIEs_OSET 233
#define          HandoverRequestAcknowledgeIEs_OSET 234
#define          HandoverFailureIEs_OSET 235
#define          HandoverNotifyIEs_OSET 236
#define          PathSwitchRequestIEs_OSET 237
#define          PathSwitchRequestAcknowledgeIEs_OSET 238
#define          PathSwitchRequestFailureIEs_OSET 239
#define          HandoverCancelIEs_OSET 240
#define          HandoverCancelAcknowledgeIEs_OSET 241
#define          UplinkRANStatusTransferIEs_OSET 242
#define          DownlinkRANStatusTransferIEs_OSET 243
#define          PagingIEs_OSET 244
#define          InitialUEMessage_IEs_OSET 245
#define          DownlinkNASTransport_IEs_OSET 246
#define          UplinkNASTransport_IEs_OSET 247
#define          NASNonDeliveryIndication_IEs_OSET 248
#define          RerouteNASRequest_IEs_OSET 249
#define          NGSetupRequestIEs_OSET 250
#define          NGSetupResponseIEs_OSET 251
#define          NGSetupFailureIEs_OSET 252
#define          RANConfigurationUpdateIEs_OSET 253
#define          RANConfigurationUpdateAcknowledgeIEs_OSET 254
#define          RANConfigurationUpdateFailureIEs_OSET 255
#define          AMFConfigurationUpdateIEs_OSET 256
#define          AMFConfigurationUpdateAcknowledgeIEs_OSET 257
#define          AMFConfigurationUpdateFailureIEs_OSET 258
#define          AMFStatusIndicationIEs_OSET 259
#define          NGResetIEs_OSET 260
#define          NGResetAcknowledgeIEs_OSET 261
#define          ErrorIndicationIEs_OSET 262
#define          OverloadStartIEs_OSET 263
#define          OverloadStopIEs_OSET 264
#define          UplinkRANConfigurationTransferIEs_OSET 265
#define          DownlinkRANConfigurationTransferIEs_OSET 266
#define          WriteReplaceWarningRequestIEs_OSET 267
#define          WriteReplaceWarningResponseIEs_OSET 268
#define          PWSCancelRequestIEs_OSET 269
#define          PWSCancelResponseIEs_OSET 270
#define          PWSRestartIndicationIEs_OSET 271
#define          PWSFailureIndicationIEs_OSET 272
#define          DownlinkUEAssociatedNRPPaTransportIEs_OSET 273
#define          UplinkUEAssociatedNRPPaTransportIEs_OSET 274
#define          DownlinkNonUEAssociatedNRPPaTransportIEs_OSET 275
#define          UplinkNonUEAssociatedNRPPaTransportIEs_OSET 276
#define          TraceStartIEs_OSET 277
#define          TraceFailureIndicationIEs_OSET 278
#define          DeactivateTraceIEs_OSET 279
#define          CellTrafficTraceIEs_OSET 280
#define          LocationReportingControlIEs_OSET 281
#define          LocationReportingFailureIndicationIEs_OSET 282
#define          LocationReportIEs_OSET 283
#define          UETNLABindingReleaseRequestIEs_OSET 284
#define          UERadioCapabilityInfoIndicationIEs_OSET 285
#define          UERadioCapabilityCheckRequestIEs_OSET 286
#define          UERadioCapabilityCheckResponseIEs_OSET 287
#define          PrivateMessageIEs_OSET 288
#define          SecondaryRATDataUsageReportIEs_OSET 289
#define          NGAP_ELEMENTARY_PROCEDURES_OSET 290
#define          NGAP_ELEMENTARY_PROCEDURES_CLASS_1_OSET 291
#define          NGAP_ELEMENTARY_PROCEDURES_CLASS_2_OSET 292

typedef struct ObjectID {
    unsigned short  length;
    unsigned char   *value;
} ObjectID;

typedef enum Criticality {
    reject = 0,
    ignore = 1,
    notify = 2
} Criticality;

typedef enum Presence {
    optional = 0,
    conditional = 1,
    mandatory = 2
} Presence;

typedef struct PrivateIE_ID {
    unsigned short  choice;
#       define      local_chosen 1
#       define      global_chosen 2
    union _union {
        unsigned short  local;  /* to choose, set choice to local_chosen */
        ObjectID        global;  /* to choose, set choice to global_chosen */
    } u;
} PrivateIE_ID;

typedef unsigned short  ProcedureCode;

typedef unsigned short  ProtocolExtensionID;

typedef unsigned short  ProtocolIE_ID;

typedef enum TriggeringMessage {
    initiating_message = 0,
    successful_outcome = 1,
    unsuccessfull_outcome = 2
} TriggeringMessage;

// **************************************************************
//
// IE parameter types from other modules.
//
// **************************************************************
// **************************************************************
//
// Class Definition for Protocol IEs
//
// **************************************************************
typedef struct NGAP_PROTOCOL_IES {
    ProtocolIE_ID   id;
    Criticality     criticality;
    unsigned short  Value;
    Presence        presence;
} NGAP_PROTOCOL_IES;

// **************************************************************
//
// Class Definition for Protocol Extensions
//
// **************************************************************
typedef struct NGAP_PROTOCOL_EXTENSION {
    ProtocolExtensionID id;
    Criticality     criticality;
    unsigned short  Extension;
    Presence        presence;
} NGAP_PROTOCOL_EXTENSION;

// **************************************************************
//
// Class Definition for Private IEs
//
// **************************************************************
typedef struct NGAP_PRIVATE_IES {
    PrivateIE_ID    id;
    Criticality     criticality;
    unsigned short  Value;
    Presence        presence;
    long            _oss_unique_index;
} NGAP_PRIVATE_IES;

typedef unsigned long   AMF_UE_NGAP_ID;

// C
typedef enum CancelAllWarningMessages {
    CancelAllWarningMessages_true = 0
} CancelAllWarningMessages;

typedef enum CNTypeRestrictionsForServing {
    CNTypeRestrictionsForServing_epc_forbidden = 0
} CNTypeRestrictionsForServing;

typedef enum ConcurrentWarningMessageInd {
    ConcurrentWarningMessageInd_true = 0
} ConcurrentWarningMessageInd;

typedef enum DataForwardingNotPossible {
    data_forwarding_not_possible = 0
} DataForwardingNotPossible;

typedef enum DirectForwardingPathAvailability {
    direct_path_available = 0
} DirectForwardingPathAvailability;

typedef enum HandoverFlag {
    handover_preparation = 0
} HandoverFlag;

typedef enum HandoverType {
    intra5gs = 0,
    fivegs_to_eps = 1,
    eps_to_5gs = 2
} HandoverType;

// I
typedef enum IMSVoiceSupportIndicator {
    supported = 0,
    not_supported = 1
} IMSVoiceSupportIndicator;

typedef long            IndexToRFSP;

typedef enum MaximumIntegrityProtectedDataRate {
    bitrate64kbs = 0,
    maximum_UE_rate = 1
} MaximumIntegrityProtectedDataRate;

typedef long            NetworkInstance;

typedef enum NewSecurityContextInd {
    NewSecurityContextInd_true = 0
} NewSecurityContextInd;

typedef unsigned short  NumberOfBroadcastsRequested;

typedef enum PagingDRX {
    v32 = 0,
    v64 = 1,
    v128 = 2,
    v256 = 3
} PagingDRX;

typedef enum PagingOrigin {
    non_3gpp = 0
} PagingOrigin;

typedef enum PagingPriority {
    priolevel1 = 0,
    priolevel2 = 1,
    priolevel3 = 2,
    priolevel4 = 3,
    priolevel5 = 4,
    priolevel6 = 5,
    priolevel7 = 6,
    priolevel8 = 7
} PagingPriority;

typedef enum PDUSessionType {
    ipv4 = 0,
    ipv6 = 1,
    ipv4v6 = 2,
    ethernet = 3,
    unstructured = 4
} PDUSessionType;

typedef unsigned short  RANPagingPriority;

typedef unsigned int    RAN_UE_NGAP_ID;

typedef enum RedirectionVoiceFallback {
    possible = 0,
    not_possible = 1
} RedirectionVoiceFallback;

typedef unsigned short  RelativeAMFCapacity;

typedef unsigned int    RepetitionPeriod;

typedef enum RRCEstablishmentCause {
    emergency = 0,
    highPriorityAccess = 1,
    mt_Access = 2,
    mo_Signalling = 3,
    mo_Data = 4,
    mo_VoiceCall = 5,
    mo_VideoCall = 6,
    mo_SMS = 7,
    mps_PriorityAccess = 8,
    mcs_PriorityAccess = 9,
    notAvailable = 10
} RRCEstablishmentCause;

typedef enum RRCInactiveTransitionReportRequest {
    subsequent_state_transition_report = 0,
    single_rrc_connected_state_report = 1,
    cancel_report = 2
} RRCInactiveTransitionReportRequest;

typedef enum RRCState {
    inactive = 0,
    connected = 1
} RRCState;

typedef enum TimeToWait {
    v1s = 0,
    v2s = 1,
    v5s = 2,
    v10s = 3,
    v20s = 4,
    v60s = 5
} TimeToWait;

typedef unsigned short  TrafficLoadReductionIndication;

typedef enum UEContextRequest {
    requested = 0
} UEContextRequest;

typedef enum UERetentionInformation {
    ues_retained = 0
} UERetentionInformation;

typedef struct _octet1 {
    unsigned int    length;
    unsigned char   *value;
} _octet1;

typedef struct PDUSessionResourceItemCxtRelCpl_ExtIEs_EXTENSION {
    /* ContentsConstraint is applied to
     * PDUSessionResourceItemCxtRelCpl_ExtIEs_EXTENSION */
    _octet1         encoded;
    struct PDUSessionResourceReleaseResponseTransfer *decoded;
} PDUSessionResourceItemCxtRelCpl_ExtIEs_EXTENSION;

typedef _octet1 RerouteNASRequest_IEs_TYPE;

union All_PDU_Types_union {
    struct AdditionalDLUPTNLInformationForHOList_ **pdu_AdditionalDLUPTNLInformationForHOList;                          /* AdditionalDLUPTNLInformationForHOList_PDU */
    struct AllowedNSSAI_ **pdu_AllowedNSSAI;  /* AllowedNSSAI_PDU */
    struct AMFName  *pdu_AMFName;  /* AMFName_PDU */
    struct AMFSetID *pdu_AMFSetID;  /* AMFSetID_PDU */
    struct AMF_TNLAssociationSetupList_ **pdu_AMF_TNLAssociationSetupList;  
                                        /* AMF_TNLAssociationSetupList_PDU */
    struct AMF_TNLAssociationToAddList_ **pdu_AMF_TNLAssociationToAddList;  
                                        /* AMF_TNLAssociationToAddList_PDU */
    struct AMF_TNLAssociationToRemoveList_ **pdu_AMF_TNLAssociationToRemoveList;                                        /* AMF_TNLAssociationToRemoveList_PDU */
    struct AMF_TNLAssociationToUpdateList_ **pdu_AMF_TNLAssociationToUpdateList;                                        /* AMF_TNLAssociationToUpdateList_PDU */
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  /* AMF_UE_NGAP_ID_PDU */
    struct AssistanceDataForPaging *pdu_AssistanceDataForPaging;  
                                              /* AssistanceDataForPaging_PDU */
    struct AssociatedQosFlowList_ **pdu_AssociatedQosFlowList;  
                                                /* AssociatedQosFlowList_PDU */
    struct BroadcastCancelledAreaList *pdu_BroadcastCancelledAreaList;  
                                           /* BroadcastCancelledAreaList_PDU */
    struct BroadcastCompletedAreaList *pdu_BroadcastCompletedAreaList;  
                                           /* BroadcastCompletedAreaList_PDU */
    CancelAllWarningMessages *pdu_CancelAllWarningMessages;  
                                             /* CancelAllWarningMessages_PDU */
    struct Cause    *pdu_Cause;  /* Cause_PDU */
    struct CellIDListForRestart *pdu_CellIDListForRestart;  
                                                 /* CellIDListForRestart_PDU */
    struct CNAssistedRANTuning *pdu_CNAssistedRANTuning;  
                                                  /* CNAssistedRANTuning_PDU */
    struct CNTypeRestrictionsForEquivalent_ **pdu_CNTypeRestrictionsForEquivalent;                                      /* CNTypeRestrictionsForEquivalent_PDU */
    CNTypeRestrictionsForServing *pdu_CNTypeRestrictionsForServing;  
                                         /* CNTypeRestrictionsForServing_PDU */
    struct CommonNetworkInstance *pdu_CommonNetworkInstance;  
                                                /* CommonNetworkInstance_PDU */
    ConcurrentWarningMessageInd *pdu_ConcurrentWarningMessageInd;  
                                          /* ConcurrentWarningMessageInd_PDU */
    struct CoreNetworkAssistanceInformationForInactive *pdu_CoreNetworkAssistanceInformationForInactive;                /* CoreNetworkAssistanceInformationForInactive_PDU */
    struct CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
                                               /* CriticalityDiagnostics_PDU */
    struct DataCodingScheme *pdu_DataCodingScheme;  /* DataCodingScheme_PDU */
    DataForwardingNotPossible *pdu_DataForwardingNotPossible;  
                                            /* DataForwardingNotPossible_PDU */
    DirectForwardingPathAvailability *pdu_DirectForwardingPathAvailability;  
                                        /* DirectForwardingPathAvailability_PDU */
    struct EmergencyAreaIDListForRestart_ **pdu_EmergencyAreaIDListForRestart;                                          /* EmergencyAreaIDListForRestart_PDU */
    struct EmergencyFallbackIndicator *pdu_EmergencyFallbackIndicator;  
                                           /* EmergencyFallbackIndicator_PDU */
    struct EN_DCSONConfigurationTransfer *pdu_EN_DCSONConfigurationTransfer;  
                                        /* EN_DCSONConfigurationTransfer_PDU */
    struct FiveG_S_TMSI *pdu_FiveG_S_TMSI;  /* FiveG_S_TMSI_PDU */
    struct GlobalRANNodeID *pdu_GlobalRANNodeID;  /* GlobalRANNodeID_PDU */
    struct GUAMI    *pdu_GUAMI;  /* GUAMI_PDU */
    struct HandoverCommandTransfer *pdu_HandoverCommandTransfer;  
                                              /* HandoverCommandTransfer_PDU */
    HandoverFlag    *pdu_HandoverFlag;  /* HandoverFlag_PDU */
    struct HandoverPreparationUnsuccessfulTransfer *pdu_HandoverPreparationUnsuccessfulTransfer;                        /* HandoverPreparationUnsuccessfulTransfer_PDU */
    struct HandoverRequestAcknowledgeTransfer *pdu_HandoverRequestAcknowledgeTransfer;                                  /* HandoverRequestAcknowledgeTransfer_PDU */
    struct HandoverRequiredTransfer *pdu_HandoverRequiredTransfer;  
                                             /* HandoverRequiredTransfer_PDU */
    struct HandoverResourceAllocationUnsuccessfulTransfer *pdu_HandoverResourceAllocationUnsuccessfulTransfer;          /* HandoverResourceAllocationUnsuccessfulTransfer_PDU */
    HandoverType    *pdu_HandoverType;  /* HandoverType_PDU */
    IMSVoiceSupportIndicator *pdu_IMSVoiceSupportIndicator;  
                                             /* IMSVoiceSupportIndicator_PDU */
    IndexToRFSP     *pdu_IndexToRFSP;  /* IndexToRFSP_PDU */
    struct InfoOnRecommendedCellsAndRANNodesForPaging *pdu_InfoOnRecommendedCellsAndRANNodesForPaging;                  /* InfoOnRecommendedCellsAndRANNodesForPaging_PDU */
    struct LocationReportingRequestType *pdu_LocationReportingRequestType;  
                                        /* LocationReportingRequestType_PDU */
    struct MaskedIMEISV *pdu_MaskedIMEISV;  /* MaskedIMEISV_PDU */
    struct MessageIdentifier *pdu_MessageIdentifier;  
                                                    /* MessageIdentifier_PDU */
    MaximumIntegrityProtectedDataRate *pdu_MaximumIntegrityProtectedDataRate;                                           /* MaximumIntegrityProtectedDataRate_PDU */
    struct MobilityRestrictionList *pdu_MobilityRestrictionList;  
                                              /* MobilityRestrictionList_PDU */
    struct NAS_PDU  *pdu_NAS_PDU;  /* NAS_PDU_PDU */
    struct NASSecurityParametersFromNGRAN *pdu_NASSecurityParametersFromNGRAN;                                          /* NASSecurityParametersFromNGRAN_PDU */
    NetworkInstance *pdu_NetworkInstance;  /* NetworkInstance_PDU */
    NewSecurityContextInd *pdu_NewSecurityContextInd;  
                                                /* NewSecurityContextInd_PDU */
    struct NGRAN_CGI *pdu_NGRAN_CGI;  /* NGRAN_CGI_PDU */
    struct NGRANTraceID *pdu_NGRANTraceID;  /* NGRANTraceID_PDU */
    struct NRPPa_PDU *pdu_NRPPa_PDU;  /* NRPPa_PDU_PDU */
    NumberOfBroadcastsRequested *pdu_NumberOfBroadcastsRequested;  
                                          /* NumberOfBroadcastsRequested_PDU */
    struct OverloadResponse *pdu_OverloadResponse;  /* OverloadResponse_PDU */
    struct OverloadStartNSSAIList_ **pdu_OverloadStartNSSAIList;  
                                               /* OverloadStartNSSAIList_PDU */
    PagingDRX       *pdu_PagingDRX;  /* PagingDRX_PDU */
    PagingOrigin    *pdu_PagingOrigin;  /* PagingOrigin_PDU */
    PagingPriority  *pdu_PagingPriority;  /* PagingPriority_PDU */
    struct PathSwitchRequestAcknowledgeTransfer *pdu_PathSwitchRequestAcknowledgeTransfer;                              /* PathSwitchRequestAcknowledgeTransfer_PDU */
    struct PathSwitchRequestSetupFailedTransfer *pdu_PathSwitchRequestSetupFailedTransfer;                              /* PathSwitchRequestSetupFailedTransfer_PDU */
    struct PathSwitchRequestTransfer *pdu_PathSwitchRequestTransfer;  
                                            /* PathSwitchRequestTransfer_PDU */
    struct PathSwitchRequestUnsuccessfulTransfer *pdu_PathSwitchRequestUnsuccessfulTransfer;                            /* PathSwitchRequestUnsuccessfulTransfer_PDU */
    struct PDUSessionAggregateMaximumBitRate *pdu_PDUSessionAggregateMaximumBitRate;                                    /* PDUSessionAggregateMaximumBitRate_PDU */
    struct PDUSessionResourceAdmittedList_ **pdu_PDUSessionResourceAdmittedList;                                        /* PDUSessionResourceAdmittedList_PDU */
    struct PDUSessionResourceFailedToModifyListModCfm_ **pdu_PDUSessionResourceFailedToModifyListModCfm;                /* PDUSessionResourceFailedToModifyListModCfm_PDU */
    struct PDUSessionResourceFailedToModifyListModRes_ **pdu_PDUSessionResourceFailedToModifyListModRes;                /* PDUSessionResourceFailedToModifyListModRes_PDU */
    struct PDUSessionResourceFailedToSetupListCxtFail_ **pdu_PDUSessionResourceFailedToSetupListCxtFail;                /* PDUSessionResourceFailedToSetupListCxtFail_PDU */
    struct PDUSessionResourceFailedToSetupListCxtRes_ **pdu_PDUSessionResourceFailedToSetupListCxtRes;                  /* PDUSessionResourceFailedToSetupListCxtRes_PDU */
    struct PDUSessionResourceFailedToSetupListHOAck_ **pdu_PDUSessionResourceFailedToSetupListHOAck;                    /* PDUSessionResourceFailedToSetupListHOAck_PDU */
    struct PDUSessionResourceFailedToSetupListPSReq_ **pdu_PDUSessionResourceFailedToSetupListPSReq;                    /* PDUSessionResourceFailedToSetupListPSReq_PDU */
    struct PDUSessionResourceFailedToSetupListSURes_ **pdu_PDUSessionResourceFailedToSetupListSURes;                    /* PDUSessionResourceFailedToSetupListSURes_PDU */
    struct PDUSessionResourceHandoverList_ **pdu_PDUSessionResourceHandoverList;                                        /* PDUSessionResourceHandoverList_PDU */
    struct PDUSessionResourceListCxtRelCpl_ **pdu_PDUSessionResourceListCxtRelCpl;                                      /* PDUSessionResourceListCxtRelCpl_PDU */
    struct PDUSessionResourceListCxtRelReq_ **pdu_PDUSessionResourceListCxtRelReq;                                      /* PDUSessionResourceListCxtRelReq_PDU */
    struct PDUSessionResourceListHORqd_ **pdu_PDUSessionResourceListHORqd;  
                                        /* PDUSessionResourceListHORqd_PDU */
    struct PDUSessionResourceModifyConfirmTransfer *pdu_PDUSessionResourceModifyConfirmTransfer;                        /* PDUSessionResourceModifyConfirmTransfer_PDU */
    struct PDUSessionResourceModifyIndicationUnsuccessfulTransfer *pdu_PDUSessionResourceModifyIndicationUnsuccessfulTransfer;          /* PDUSessionResourceModifyIndicationUnsuccessfulTransfer_PDU */
    struct PDUSessionResourceModifyRequestTransfer *pdu_PDUSessionResourceModifyRequestTransfer;                        /* PDUSessionResourceModifyRequestTransfer_PDU */
    struct PDUSessionResourceModifyResponseTransfer *pdu_PDUSessionResourceModifyResponseTransfer;                      /* PDUSessionResourceModifyResponseTransfer_PDU */
    struct PDUSessionResourceModifyIndicationTransfer *pdu_PDUSessionResourceModifyIndicationTransfer;                  /* PDUSessionResourceModifyIndicationTransfer_PDU */
    struct PDUSessionResourceModifyListModCfm_ **pdu_PDUSessionResourceModifyListModCfm;                                /* PDUSessionResourceModifyListModCfm_PDU */
    struct PDUSessionResourceModifyListModInd_ **pdu_PDUSessionResourceModifyListModInd;                                /* PDUSessionResourceModifyListModInd_PDU */
    struct PDUSessionResourceModifyListModReq_ **pdu_PDUSessionResourceModifyListModReq;                                /* PDUSessionResourceModifyListModReq_PDU */
    struct PDUSessionResourceModifyListModRes_ **pdu_PDUSessionResourceModifyListModRes;                                /* PDUSessionResourceModifyListModRes_PDU */
    struct PDUSessionResourceModifyUnsuccessfulTransfer *pdu_PDUSessionResourceModifyUnsuccessfulTransfer;              /* PDUSessionResourceModifyUnsuccessfulTransfer_PDU */
    struct PDUSessionResourceNotifyList_ **pdu_PDUSessionResourceNotifyList;  
                                        /* PDUSessionResourceNotifyList_PDU */
    struct PDUSessionResourceNotifyReleasedTransfer *pdu_PDUSessionResourceNotifyReleasedTransfer;                      /* PDUSessionResourceNotifyReleasedTransfer_PDU */
    struct PDUSessionResourceNotifyTransfer *pdu_PDUSessionResourceNotifyTransfer;                                      /* PDUSessionResourceNotifyTransfer_PDU */
    struct PDUSessionResourceReleaseCommandTransfer *pdu_PDUSessionResourceReleaseCommandTransfer;                      /* PDUSessionResourceReleaseCommandTransfer_PDU */
    struct PDUSessionResourceReleasedListNot_ **pdu_PDUSessionResourceReleasedListNot;                                  /* PDUSessionResourceReleasedListNot_PDU */
    struct PDUSessionResourceReleasedListPSAck_ **pdu_PDUSessionResourceReleasedListPSAck;                              /* PDUSessionResourceReleasedListPSAck_PDU */
    struct PDUSessionResourceReleasedListPSFail_ **pdu_PDUSessionResourceReleasedListPSFail;                            /* PDUSessionResourceReleasedListPSFail_PDU */
    struct PDUSessionResourceReleasedListRelRes_ **pdu_PDUSessionResourceReleasedListRelRes;                            /* PDUSessionResourceReleasedListRelRes_PDU */
    struct PDUSessionResourceReleaseResponseTransfer *pdu_PDUSessionResourceReleaseResponseTransfer;                    /* PDUSessionResourceReleaseResponseTransfer_PDU */
    struct PDUSessionResourceSecondaryRATUsageList_ **pdu_PDUSessionResourceSecondaryRATUsageList;                      /* PDUSessionResourceSecondaryRATUsageList_PDU */
    struct PDUSessionResourceSetupListCxtReq_ **pdu_PDUSessionResourceSetupListCxtReq;                                  /* PDUSessionResourceSetupListCxtReq_PDU */
    struct PDUSessionResourceSetupListCxtRes_ **pdu_PDUSessionResourceSetupListCxtRes;                                  /* PDUSessionResourceSetupListCxtRes_PDU */
    struct PDUSessionResourceSetupListHOReq_ **pdu_PDUSessionResourceSetupListHOReq;                                    /* PDUSessionResourceSetupListHOReq_PDU */
    struct PDUSessionResourceSetupListSUReq_ **pdu_PDUSessionResourceSetupListSUReq;                                    /* PDUSessionResourceSetupListSUReq_PDU */
    struct PDUSessionResourceSetupListSURes_ **pdu_PDUSessionResourceSetupListSURes;                                    /* PDUSessionResourceSetupListSURes_PDU */
    struct PDUSessionResourceSetupRequestTransfer *pdu_PDUSessionResourceSetupRequestTransfer;                          /* PDUSessionResourceSetupRequestTransfer_PDU */
    struct PDUSessionResourceSetupResponseTransfer *pdu_PDUSessionResourceSetupResponseTransfer;                        /* PDUSessionResourceSetupResponseTransfer_PDU */
    struct PDUSessionResourceSetupUnsuccessfulTransfer *pdu_PDUSessionResourceSetupUnsuccessfulTransfer;                /* PDUSessionResourceSetupUnsuccessfulTransfer_PDU */
    struct PDUSessionResourceSwitchedList_ **pdu_PDUSessionResourceSwitchedList;                                        /* PDUSessionResourceSwitchedList_PDU */
    struct PDUSessionResourceToBeSwitchedDLList_ **pdu_PDUSessionResourceToBeSwitchedDLList;                            /* PDUSessionResourceToBeSwitchedDLList_PDU */
    struct PDUSessionResourceToReleaseListHOCmd_ **pdu_PDUSessionResourceToReleaseListHOCmd;                            /* PDUSessionResourceToReleaseListHOCmd_PDU */
    struct PDUSessionResourceToReleaseListRelCmd_ **pdu_PDUSessionResourceToReleaseListRelCmd;                          /* PDUSessionResourceToReleaseListRelCmd_PDU */
    PDUSessionType  *pdu_PDUSessionType;  /* PDUSessionType_PDU */
    struct PLMNIdentity *pdu_PLMNIdentity;  /* PLMNIdentity_PDU */
    struct PLMNSupportList_ **pdu_PLMNSupportList;  /* PLMNSupportList_PDU */
    struct PWSFailedCellIDList *pdu_PWSFailedCellIDList;  
                                                  /* PWSFailedCellIDList_PDU */
    struct QosFlowAddOrModifyRequestList_ **pdu_QosFlowAddOrModifyRequestList;                                          /* QosFlowAddOrModifyRequestList_PDU */
    struct QosFlowListWithCause_ **pdu_QosFlowListWithCause;  
                                                 /* QosFlowListWithCause_PDU */
    struct QosFlowPerTNLInformationList_ **pdu_QosFlowPerTNLInformationList;  
                                        /* QosFlowPerTNLInformationList_PDU */
    struct QosFlowSetupRequestList_ **pdu_QosFlowSetupRequestList;  
                                              /* QosFlowSetupRequestList_PDU */
    struct RANNodeName *pdu_RANNodeName;  /* RANNodeName_PDU */
    RANPagingPriority *pdu_RANPagingPriority;  /* RANPagingPriority_PDU */
    struct RANStatusTransfer_TransparentContainer *pdu_RANStatusTransfer_TransparentContainer;                          /* RANStatusTransfer_TransparentContainer_PDU */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  /* RAN_UE_NGAP_ID_PDU */
    RedirectionVoiceFallback *pdu_RedirectionVoiceFallback;  
                                             /* RedirectionVoiceFallback_PDU */
    RelativeAMFCapacity *pdu_RelativeAMFCapacity;  
                                                  /* RelativeAMFCapacity_PDU */
    RepetitionPeriod *pdu_RepetitionPeriod;  /* RepetitionPeriod_PDU */
    struct ResetType *pdu_ResetType;  /* ResetType_PDU */
    struct RoutingID *pdu_RoutingID;  /* RoutingID_PDU */
    RRCEstablishmentCause *pdu_RRCEstablishmentCause;  
                                                /* RRCEstablishmentCause_PDU */
    RRCInactiveTransitionReportRequest *pdu_RRCInactiveTransitionReportRequest;                                         /* RRCInactiveTransitionReportRequest_PDU */
    RRCState        *pdu_RRCState;  /* RRCState_PDU */
    struct SecondaryRATUsageInformation *pdu_SecondaryRATUsageInformation;  
                                        /* SecondaryRATUsageInformation_PDU */
    struct SecondaryRATDataUsageReportTransfer *pdu_SecondaryRATDataUsageReportTransfer;                                /* SecondaryRATDataUsageReportTransfer_PDU */
    struct SecurityContext *pdu_SecurityContext;  /* SecurityContext_PDU */
    struct SecurityIndication *pdu_SecurityIndication;  
                                                   /* SecurityIndication_PDU */
    struct SecurityKey *pdu_SecurityKey;  /* SecurityKey_PDU */
    struct SecurityResult *pdu_SecurityResult;  /* SecurityResult_PDU */
    struct SerialNumber *pdu_SerialNumber;  /* SerialNumber_PDU */
    struct ServedGUAMIList_ **pdu_ServedGUAMIList;  /* ServedGUAMIList_PDU */
    struct S_NSSAI  *pdu_S_NSSAI;  /* S_NSSAI_PDU */
    struct SONConfigurationTransfer *pdu_SONConfigurationTransfer;  
                                             /* SONConfigurationTransfer_PDU */
    struct SourceToTarget_TransparentContainer *pdu_SourceToTarget_TransparentContainer;                                /* SourceToTarget_TransparentContainer_PDU */
    struct SupportedTAList_ **pdu_SupportedTAList;  /* SupportedTAList_PDU */
    struct TAIListForPaging_ **pdu_TAIListForPaging;  
                                                     /* TAIListForPaging_PDU */
    struct TAIListForRestart_ **pdu_TAIListForRestart;  
                                                    /* TAIListForRestart_PDU */
    struct TargetID *pdu_TargetID;  /* TargetID_PDU */
    struct TargetToSource_TransparentContainer *pdu_TargetToSource_TransparentContainer;                                /* TargetToSource_TransparentContainer_PDU */
    TimeToWait      *pdu_TimeToWait;  /* TimeToWait_PDU */
    struct TNLAssociationList_ **pdu_TNLAssociationList;  
                                                   /* TNLAssociationList_PDU */
    struct TraceActivation *pdu_TraceActivation;  /* TraceActivation_PDU */
    TrafficLoadReductionIndication *pdu_TrafficLoadReductionIndication;  
                                       /* TrafficLoadReductionIndication_PDU */
    struct TransportLayerAddress *pdu_TransportLayerAddress;  
                                                /* TransportLayerAddress_PDU */
    struct UEAggregateMaximumBitRate *pdu_UEAggregateMaximumBitRate;  
                                            /* UEAggregateMaximumBitRate_PDU */
    struct UE_associatedLogicalNG_connectionList_ **pdu_UE_associatedLogicalNG_connectionList;                          /* UE_associatedLogicalNG_connectionList_PDU */
    UEContextRequest *pdu_UEContextRequest;  /* UEContextRequest_PDU */
    struct UE_NGAP_IDs *pdu_UE_NGAP_IDs;  /* UE_NGAP_IDs_PDU */
    struct UEPagingIdentity *pdu_UEPagingIdentity;  /* UEPagingIdentity_PDU */
    struct UEPresenceInAreaOfInterestList_ **pdu_UEPresenceInAreaOfInterestList;                                        /* UEPresenceInAreaOfInterestList_PDU */
    struct UERadioCapability *pdu_UERadioCapability;  
                                                    /* UERadioCapability_PDU */
    struct UERadioCapabilityForPaging *pdu_UERadioCapabilityForPaging;  
                                           /* UERadioCapabilityForPaging_PDU */
    UERetentionInformation *pdu_UERetentionInformation;  
                                               /* UERetentionInformation_PDU */
    struct UESecurityCapabilities *pdu_UESecurityCapabilities;  
                                               /* UESecurityCapabilities_PDU */
    struct UL_NGU_UP_TNLModifyList_ **pdu_UL_NGU_UP_TNLModifyList;  
                                              /* UL_NGU_UP_TNLModifyList_PDU */
    struct UnavailableGUAMIList_ **pdu_UnavailableGUAMIList;  
                                                 /* UnavailableGUAMIList_PDU */
    struct UPTransportLayerInformation *pdu_UPTransportLayerInformation;  
                                        /* UPTransportLayerInformation_PDU */
    struct UPTransportLayerInformationList_ **pdu_UPTransportLayerInformationList;                                      /* UPTransportLayerInformationList_PDU */
    struct UPTransportLayerInformationPairList_ **pdu_UPTransportLayerInformationPairList;                              /* UPTransportLayerInformationPairList_PDU */
    struct UserLocationInformation *pdu_UserLocationInformation;  
                                              /* UserLocationInformation_PDU */
    struct WarningAreaCoordinates *pdu_WarningAreaCoordinates;  
                                               /* WarningAreaCoordinates_PDU */
    struct WarningAreaList *pdu_WarningAreaList;  /* WarningAreaList_PDU */
    struct WarningMessageContents *pdu_WarningMessageContents;  
                                               /* WarningMessageContents_PDU */
    struct WarningSecurityInfo *pdu_WarningSecurityInfo;  
                                                  /* WarningSecurityInfo_PDU */
    struct WarningType *pdu_WarningType;  /* WarningType_PDU */
    struct PDUSessionResourceSetupRequest *pdu_PDUSessionResourceSetupRequest;                                          /* PDUSessionResourceSetupRequest_PDU */
    struct PDUSessionResourceSetupResponse *pdu_PDUSessionResourceSetupResponse;                                        /* PDUSessionResourceSetupResponse_PDU */
    struct PDUSessionResourceReleaseCommand *pdu_PDUSessionResourceReleaseCommand;                                      /* PDUSessionResourceReleaseCommand_PDU */
    struct PDUSessionResourceReleaseResponse *pdu_PDUSessionResourceReleaseResponse;                                    /* PDUSessionResourceReleaseResponse_PDU */
    struct PDUSessionResourceModifyRequest *pdu_PDUSessionResourceModifyRequest;                                        /* PDUSessionResourceModifyRequest_PDU */
    struct PDUSessionResourceModifyResponse *pdu_PDUSessionResourceModifyResponse;                                      /* PDUSessionResourceModifyResponse_PDU */
    struct PDUSessionResourceNotify *pdu_PDUSessionResourceNotify;  
                                             /* PDUSessionResourceNotify_PDU */
    struct PDUSessionResourceModifyIndication *pdu_PDUSessionResourceModifyIndication;                                  /* PDUSessionResourceModifyIndication_PDU */
    struct PDUSessionResourceModifyConfirm *pdu_PDUSessionResourceModifyConfirm;                                        /* PDUSessionResourceModifyConfirm_PDU */
    struct InitialContextSetupRequest *pdu_InitialContextSetupRequest;  
                                           /* InitialContextSetupRequest_PDU */
    struct InitialContextSetupResponse *pdu_InitialContextSetupResponse;  
                                        /* InitialContextSetupResponse_PDU */
    struct InitialContextSetupFailure *pdu_InitialContextSetupFailure;  
                                           /* InitialContextSetupFailure_PDU */
    struct UEContextReleaseRequest *pdu_UEContextReleaseRequest;  
                                              /* UEContextReleaseRequest_PDU */
    struct UEContextReleaseCommand *pdu_UEContextReleaseCommand;  
                                              /* UEContextReleaseCommand_PDU */
    struct UEContextReleaseComplete *pdu_UEContextReleaseComplete;  
                                             /* UEContextReleaseComplete_PDU */
    struct UEContextModificationRequest *pdu_UEContextModificationRequest;  
                                        /* UEContextModificationRequest_PDU */
    struct UEContextModificationResponse *pdu_UEContextModificationResponse;  
                                        /* UEContextModificationResponse_PDU */
    struct UEContextModificationFailure *pdu_UEContextModificationFailure;  
                                        /* UEContextModificationFailure_PDU */
    struct RRCInactiveTransitionReport *pdu_RRCInactiveTransitionReport;  
                                        /* RRCInactiveTransitionReport_PDU */
    struct HandoverRequired *pdu_HandoverRequired;  /* HandoverRequired_PDU */
    struct HandoverCommand *pdu_HandoverCommand;  /* HandoverCommand_PDU */
    struct HandoverPreparationFailure *pdu_HandoverPreparationFailure;  
                                           /* HandoverPreparationFailure_PDU */
    struct HandoverRequest *pdu_HandoverRequest;  /* HandoverRequest_PDU */
    struct HandoverRequestAcknowledge *pdu_HandoverRequestAcknowledge;  
                                           /* HandoverRequestAcknowledge_PDU */
    struct HandoverFailure *pdu_HandoverFailure;  /* HandoverFailure_PDU */
    struct HandoverNotify *pdu_HandoverNotify;  /* HandoverNotify_PDU */
    struct PathSwitchRequest *pdu_PathSwitchRequest;  
                                                    /* PathSwitchRequest_PDU */
    struct PathSwitchRequestAcknowledge *pdu_PathSwitchRequestAcknowledge;  
                                        /* PathSwitchRequestAcknowledge_PDU */
    struct PathSwitchRequestFailure *pdu_PathSwitchRequestFailure;  
                                             /* PathSwitchRequestFailure_PDU */
    struct HandoverCancel *pdu_HandoverCancel;  /* HandoverCancel_PDU */
    struct HandoverCancelAcknowledge *pdu_HandoverCancelAcknowledge;  
                                            /* HandoverCancelAcknowledge_PDU */
    struct UplinkRANStatusTransfer *pdu_UplinkRANStatusTransfer;  
                                              /* UplinkRANStatusTransfer_PDU */
    struct DownlinkRANStatusTransfer *pdu_DownlinkRANStatusTransfer;  
                                            /* DownlinkRANStatusTransfer_PDU */
    struct Paging   *pdu_Paging;  /* Paging_PDU */
    struct InitialUEMessage *pdu_InitialUEMessage;  /* InitialUEMessage_PDU */
    struct DownlinkNASTransport *pdu_DownlinkNASTransport;  
                                                 /* DownlinkNASTransport_PDU */
    struct UplinkNASTransport *pdu_UplinkNASTransport;  
                                                   /* UplinkNASTransport_PDU */
    struct NASNonDeliveryIndication *pdu_NASNonDeliveryIndication;  
                                             /* NASNonDeliveryIndication_PDU */
    struct RerouteNASRequest *pdu_RerouteNASRequest;  
                                                    /* RerouteNASRequest_PDU */
    struct NGSetupRequest *pdu_NGSetupRequest;  /* NGSetupRequest_PDU */
    struct NGSetupResponse *pdu_NGSetupResponse;  /* NGSetupResponse_PDU */
    struct NGSetupFailure *pdu_NGSetupFailure;  /* NGSetupFailure_PDU */
    struct RANConfigurationUpdate *pdu_RANConfigurationUpdate;  
                                               /* RANConfigurationUpdate_PDU */
    struct RANConfigurationUpdateAcknowledge *pdu_RANConfigurationUpdateAcknowledge;                                    /* RANConfigurationUpdateAcknowledge_PDU */
    struct RANConfigurationUpdateFailure *pdu_RANConfigurationUpdateFailure;  
                                        /* RANConfigurationUpdateFailure_PDU */
    struct AMFConfigurationUpdate *pdu_AMFConfigurationUpdate;  
                                               /* AMFConfigurationUpdate_PDU */
    struct AMFConfigurationUpdateAcknowledge *pdu_AMFConfigurationUpdateAcknowledge;                                    /* AMFConfigurationUpdateAcknowledge_PDU */
    struct AMFConfigurationUpdateFailure *pdu_AMFConfigurationUpdateFailure;  
                                        /* AMFConfigurationUpdateFailure_PDU */
    struct AMFStatusIndication *pdu_AMFStatusIndication;  
                                                  /* AMFStatusIndication_PDU */
    struct NGReset  *pdu_NGReset;  /* NGReset_PDU */
    struct NGResetAcknowledge *pdu_NGResetAcknowledge;  
                                                   /* NGResetAcknowledge_PDU */
    struct ErrorIndication *pdu_ErrorIndication;  /* ErrorIndication_PDU */
    struct OverloadStart *pdu_OverloadStart;  /* OverloadStart_PDU */
    struct OverloadStop *pdu_OverloadStop;  /* OverloadStop_PDU */
    struct UplinkRANConfigurationTransfer *pdu_UplinkRANConfigurationTransfer;                                          /* UplinkRANConfigurationTransfer_PDU */
    struct DownlinkRANConfigurationTransfer *pdu_DownlinkRANConfigurationTransfer;                                      /* DownlinkRANConfigurationTransfer_PDU */
    struct WriteReplaceWarningRequest *pdu_WriteReplaceWarningRequest;  
                                           /* WriteReplaceWarningRequest_PDU */
    struct WriteReplaceWarningResponse *pdu_WriteReplaceWarningResponse;  
                                        /* WriteReplaceWarningResponse_PDU */
    struct PWSCancelRequest *pdu_PWSCancelRequest;  /* PWSCancelRequest_PDU */
    struct PWSCancelResponse *pdu_PWSCancelResponse;  
                                                    /* PWSCancelResponse_PDU */
    struct PWSRestartIndication *pdu_PWSRestartIndication;  
                                                 /* PWSRestartIndication_PDU */
    struct PWSFailureIndication *pdu_PWSFailureIndication;  
                                                 /* PWSFailureIndication_PDU */
    struct DownlinkUEAssociatedNRPPaTransport *pdu_DownlinkUEAssociatedNRPPaTransport;                                  /* DownlinkUEAssociatedNRPPaTransport_PDU */
    struct UplinkUEAssociatedNRPPaTransport *pdu_UplinkUEAssociatedNRPPaTransport;                                      /* UplinkUEAssociatedNRPPaTransport_PDU */
    struct DownlinkNonUEAssociatedNRPPaTransport *pdu_DownlinkNonUEAssociatedNRPPaTransport;                            /* DownlinkNonUEAssociatedNRPPaTransport_PDU */
    struct UplinkNonUEAssociatedNRPPaTransport *pdu_UplinkNonUEAssociatedNRPPaTransport;                                /* UplinkNonUEAssociatedNRPPaTransport_PDU */
    struct TraceStart *pdu_TraceStart;  /* TraceStart_PDU */
    struct TraceFailureIndication *pdu_TraceFailureIndication;  
                                               /* TraceFailureIndication_PDU */
    struct DeactivateTrace *pdu_DeactivateTrace;  /* DeactivateTrace_PDU */
    struct CellTrafficTrace *pdu_CellTrafficTrace;  /* CellTrafficTrace_PDU */
    struct LocationReportingControl *pdu_LocationReportingControl;  
                                             /* LocationReportingControl_PDU */
    struct LocationReportingFailureIndication *pdu_LocationReportingFailureIndication;                                  /* LocationReportingFailureIndication_PDU */
    struct LocationReport *pdu_LocationReport;  /* LocationReport_PDU */
    struct UETNLABindingReleaseRequest *pdu_UETNLABindingReleaseRequest;  
                                        /* UETNLABindingReleaseRequest_PDU */
    struct UERadioCapabilityInfoIndication *pdu_UERadioCapabilityInfoIndication;                                        /* UERadioCapabilityInfoIndication_PDU */
    struct UERadioCapabilityCheckRequest *pdu_UERadioCapabilityCheckRequest;  
                                        /* UERadioCapabilityCheckRequest_PDU */
    struct UERadioCapabilityCheckResponse *pdu_UERadioCapabilityCheckResponse;                                          /* UERadioCapabilityCheckResponse_PDU */
    struct PrivateMessage *pdu_PrivateMessage;  /* PrivateMessage_PDU */
    struct SecondaryRATDataUsageReport *pdu_SecondaryRATDataUsageReport;  
                                        /* SecondaryRATDataUsageReport_PDU */
    struct NGAP_PDU *pdu_NGAP_PDU;  /* NGAP_PDU_PDU */
    PDUSessionResourceItemCxtRelCpl_ExtIEs_EXTENSION *pdu_PDUSessionResourceItemCxtRelCpl_ExtIEs_EXTENSION;             /* PDUSessionResourceItemCxtRelCpl_ExtIEs_EXTENSION_PDU */
    _octet1         *pdu_RerouteNASRequest_IEs_TYPE;  
                                           /* RerouteNASRequest_IEs_TYPE_PDU */
};

enum PrivateMessageIEs_Value_PDUs {
    PDU_PrivateMessageIEs_Value_UNKNOWN = 0,
    PDU_PrivateMessageIEs_Value_MAX = INT_MAX
};

union PrivateMessageIEs_Value_union {
    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PrivateMessageIEs_Value {
    PrivateMessageIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PrivateMessageIEs_Value_union decoded;
} PrivateMessageIEs_Value;

typedef struct PrivateIE_Field {
    PrivateIE_ID    id;
    Criticality     criticality;
    PrivateMessageIEs_Value value;
} PrivateIE_Field;

typedef struct TransportLayerAddress {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} TransportLayerAddress;

typedef struct GTP_TEID {
    unsigned short  length;
    unsigned char   value[4];
} GTP_TEID;

typedef struct _seq1 {
    ProtocolExtensionID id;
    Criticality     criticality;
    PrivateMessageIEs_Value extensionValue;
} _seq1;

typedef struct _seqof1_ {
    struct _seqof1_ *next;
    _seq1           value;
} *_seqof1;

typedef struct GTPTunnel {
    unsigned char   bit_mask;
#       define      GTPTunnel_iE_Extensions_present 0x80
    TransportLayerAddress transportLayerAddress;
    GTP_TEID        gTP_TEID;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * GTPTunnel_iE_Extensions_present if
                                      * present */
} GTPTunnel;

typedef struct ProtocolIE_SingleContainer {
    ProtocolIE_ID   id;
    Criticality     criticality;
    PrivateMessageIEs_Value value;
} ProtocolIE_SingleContainer;

typedef struct UPTransportLayerInformation {
    unsigned short  choice;
#       define      gTPTunnel_chosen 1
#       define      UPTransportLayerInformation_choice_Extensions_chosen 2
    union _union {
        GTPTunnel       gTPTunnel;  /* to choose, set choice to
                                     * gTPTunnel_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to
                      * UPTransportLayerInformation_choice_Extensions_chosen */
    } u;
} UPTransportLayerInformation;

typedef struct AdditionalDLUPTNLInformationForHOItem {
    unsigned char   bit_mask;
#       define      additionalDLForwardingUPTNLInformation_present 0x80
#       define      AdditionalDLUPTNLInformationForHOItem_iE_Extensions_present 0x40
    UPTransportLayerInformation additionalDL_NGU_UP_TNLInformation;
    struct QosFlowListWithDataForwarding_ *additionalQosFlowSetupResponseList;
    UPTransportLayerInformation additionalDLForwardingUPTNLInformation;  
                                  /* optional; set in bit_mask
                            * additionalDLForwardingUPTNLInformation_present if
                            * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
               * AdditionalDLUPTNLInformationForHOItem_iE_Extensions_present if
               * present */
} AdditionalDLUPTNLInformationForHOItem;

// A
typedef struct AdditionalDLUPTNLInformationForHOList_ {
    struct AdditionalDLUPTNLInformationForHOList_ *next;
    AdditionalDLUPTNLInformationForHOItem value;
} *AdditionalDLUPTNLInformationForHOList;

typedef enum AdditionalQosFlowInformation {
    more_likely = 0
} AdditionalQosFlowInformation;

typedef unsigned short  PriorityLevelARP;

typedef enum Pre_emptionCapability {
    shall_not_trigger_pre_emption = 0,
    may_trigger_pre_emption = 1
} Pre_emptionCapability;

typedef enum Pre_emptionVulnerability {
    not_pre_emptable = 0,
    pre_emptable = 1
} Pre_emptionVulnerability;

typedef struct AllocationAndRetentionPriority {
    unsigned char   bit_mask;
#       define      AllocationAndRetentionPriority_iE_Extensions_present 0x80
    PriorityLevelARP priorityLevelARP;
    Pre_emptionCapability pre_emptionCapability;
    Pre_emptionVulnerability pre_emptionVulnerability;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                      * AllocationAndRetentionPriority_iE_Extensions_present if
                      * present */
} AllocationAndRetentionPriority;

// This IE includes a transparent container from the source RAN node to the target RAN node. 
// The octets of the OCTET STRING are encoded according to the specifications of the target system.
typedef struct SST {
    unsigned short  length;
    unsigned char   value[1];
} SST;

// S
typedef struct SD {
    unsigned short  length;
    unsigned char   value[3];
} SD;

typedef struct S_NSSAI {
    unsigned char   bit_mask;
#       define      sD_present 0x80
#       define      S_NSSAI_iE_Extensions_present 0x40
    SST             sST;
    SD              sD;  /* optional; set in bit_mask sD_present if present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * S_NSSAI_iE_Extensions_present if
                                      * present */
} S_NSSAI;

typedef struct AllowedNSSAI_Item {
    unsigned char   bit_mask;
#       define      AllowedNSSAI_Item_iE_Extensions_present 0x80
    S_NSSAI         s_NSSAI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * AllowedNSSAI_Item_iE_Extensions_present
                                      * if present */
} AllowedNSSAI_Item;

typedef struct AllowedNSSAI_ {
    struct AllowedNSSAI_ *next;
    AllowedNSSAI_Item value;
} *AllowedNSSAI;

// T
typedef struct TAC {
    unsigned short  length;
    unsigned char   value[3];
} TAC;

typedef struct AllowedTACs_ {
    struct AllowedTACs_ *next;
    TAC             value;
} *AllowedTACs;

typedef struct AMFName {
    unsigned short  length;
    char            *value;
} AMFName;

typedef struct PLMNIdentity {
    unsigned short  length;
    unsigned char   value[3];
} PLMNIdentity;

typedef struct _bit1 {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} _bit1;

typedef struct GNB_ID {
    unsigned short  choice;
#       define      gNB_ID_chosen 1
#       define      GNB_ID_choice_Extensions_chosen 2
    union _union {
        _bit1           gNB_ID;  /* to choose, set choice to gNB_ID_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                        * to GNB_ID_choice_Extensions_chosen */
    } u;
} GNB_ID;

typedef struct GlobalGNB_ID {
    unsigned char   bit_mask;
#       define      GlobalGNB_ID_iE_Extensions_present 0x80
    PLMNIdentity    pLMNIdentity;
    GNB_ID          gNB_ID;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * GlobalGNB_ID_iE_Extensions_present if
                                      * present */
} GlobalGNB_ID;

typedef struct NgENB_ID {
    unsigned short  choice;
#       define      macroNgENB_ID_chosen 1
#       define      shortMacroNgENB_ID_chosen 2
#       define      longMacroNgENB_ID_chosen 3
#       define      NgENB_ID_choice_Extensions_chosen 4
    union _union {
        _bit1           macroNgENB_ID;  /* to choose, set choice to
                                         * macroNgENB_ID_chosen */
        _bit1           shortMacroNgENB_ID;  /* to choose, set choice to
                                              * shortMacroNgENB_ID_chosen */
        _bit1           longMacroNgENB_ID;  /* to choose, set choice to
                                             * longMacroNgENB_ID_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                      * to NgENB_ID_choice_Extensions_chosen */
    } u;
} NgENB_ID;

typedef struct GlobalNgENB_ID {
    unsigned char   bit_mask;
#       define      GlobalNgENB_ID_iE_Extensions_present 0x80
    PLMNIdentity    pLMNIdentity;
    NgENB_ID        ngENB_ID;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * GlobalNgENB_ID_iE_Extensions_present if
                                      * present */
} GlobalNgENB_ID;

// N
typedef struct N3IWF_ID {
    unsigned short  choice;
#       define      n3IWF_ID_chosen 1
#       define      N3IWF_ID_choice_Extensions_chosen 2
    union _union {
        _bit1           n3IWF_ID;  /* to choose, set choice to
                                    * n3IWF_ID_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                      * to N3IWF_ID_choice_Extensions_chosen */
    } u;
} N3IWF_ID;

typedef struct GlobalN3IWF_ID {
    unsigned char   bit_mask;
#       define      GlobalN3IWF_ID_iE_Extensions_present 0x80
    PLMNIdentity    pLMNIdentity;
    N3IWF_ID        n3IWF_ID;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * GlobalN3IWF_ID_iE_Extensions_present if
                                      * present */
} GlobalN3IWF_ID;

typedef struct GlobalRANNodeID {
    unsigned short  choice;
#       define      globalGNB_ID_chosen 1
#       define      globalNgENB_ID_chosen 2
#       define      globalN3IWF_ID_chosen 3
#       define      GlobalRANNodeID_choice_Extensions_chosen 4
    union _union {
        GlobalGNB_ID    globalGNB_ID;  /* to choose, set choice to
                                        * globalGNB_ID_chosen */
        GlobalNgENB_ID  globalNgENB_ID;  /* to choose, set choice to
                                          * globalNgENB_ID_chosen */
        GlobalN3IWF_ID  globalN3IWF_ID;  /* to choose, set choice to
                                          * globalN3IWF_ID_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to
                                  * GlobalRANNodeID_choice_Extensions_chosen */
    } u;
} GlobalRANNodeID;

typedef struct TAI {
    unsigned char   bit_mask;
#       define      TAI_iE_Extensions_present 0x80
    PLMNIdentity    pLMNIdentity;
    TAC             tAC;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * TAI_iE_Extensions_present if present */
} TAI;

typedef struct AMFPagingTarget {
    unsigned short  choice;
#       define      globalRANNodeID_chosen 1
#       define      tAI_chosen 2
#       define      AMFPagingTarget_choice_Extensions_chosen 3
    union _union {
        GlobalRANNodeID globalRANNodeID;  /* to choose, set choice to
                                           * globalRANNodeID_chosen */
        TAI             tAI;  /* to choose, set choice to tAI_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to
                                  * AMFPagingTarget_choice_Extensions_chosen */
    } u;
} AMFPagingTarget;

typedef struct AMFPointer {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} AMFPointer;

typedef struct AMFRegionID {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} AMFRegionID;

typedef struct AMFSetID {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} AMFSetID;

typedef struct CPTransportLayerInformation {
    unsigned short  choice;
#       define      endpointIPAddress_chosen 1
#       define      CPTransportLayerInformation_choice_Extensions_chosen 2
    union _union {
        TransportLayerAddress endpointIPAddress;  /* to choose, set choice to
                                                  * endpointIPAddress_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to
                      * CPTransportLayerInformation_choice_Extensions_chosen */
    } u;
} CPTransportLayerInformation;

typedef struct AMF_TNLAssociationSetupItem {
    unsigned char   bit_mask;
#       define      AMF_TNLAssociationSetupItem_iE_Extensions_present 0x80
    CPTransportLayerInformation aMF_TNLAssociationAddress;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                         * AMF_TNLAssociationSetupItem_iE_Extensions_present if
                         * present */
} AMF_TNLAssociationSetupItem;

typedef struct AMF_TNLAssociationSetupList_ {
    struct AMF_TNLAssociationSetupList_ *next;
    AMF_TNLAssociationSetupItem value;
} *AMF_TNLAssociationSetupList;

typedef enum TNLAssociationUsage {
    ue = 0,
    non_ue = 1,
    both = 2
} TNLAssociationUsage;

typedef unsigned short  TNLAddressWeightFactor;

typedef struct AMF_TNLAssociationToAddItem {
    unsigned char   bit_mask;
#       define      AMF_TNLAssociationToAddItem_tNLAssociationUsage_present 0x80
#       define      AMF_TNLAssociationToAddItem_iE_Extensions_present 0x40
    CPTransportLayerInformation aMF_TNLAssociationAddress;
    TNLAssociationUsage tNLAssociationUsage;  /* optional; set in bit_mask
                   * AMF_TNLAssociationToAddItem_tNLAssociationUsage_present if
                   * present */
    TNLAddressWeightFactor tNLAddressWeightFactor;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                         * AMF_TNLAssociationToAddItem_iE_Extensions_present if
                         * present */
} AMF_TNLAssociationToAddItem;

typedef struct AMF_TNLAssociationToAddList_ {
    struct AMF_TNLAssociationToAddList_ *next;
    AMF_TNLAssociationToAddItem value;
} *AMF_TNLAssociationToAddList;

typedef struct AMF_TNLAssociationToRemoveItem {
    unsigned char   bit_mask;
#       define      AMF_TNLAssociationToRemoveItem_iE_Extensions_present 0x80
    CPTransportLayerInformation aMF_TNLAssociationAddress;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                      * AMF_TNLAssociationToRemoveItem_iE_Extensions_present if
                      * present */
} AMF_TNLAssociationToRemoveItem;

typedef struct AMF_TNLAssociationToRemoveList_ {
    struct AMF_TNLAssociationToRemoveList_ *next;
    AMF_TNLAssociationToRemoveItem value;
} *AMF_TNLAssociationToRemoveList;

typedef struct AMF_TNLAssociationToUpdateItem {
    unsigned char   bit_mask;
#       define      AMF_TNLAssociationToUpdateItem_tNLAssociationUsage_present 0x80
#       define      tNLAddressWeightFactor_present 0x40
#       define      AMF_TNLAssociationToUpdateItem_iE_Extensions_present 0x20
    CPTransportLayerInformation aMF_TNLAssociationAddress;
    TNLAssociationUsage tNLAssociationUsage;  /* optional; set in bit_mask
                * AMF_TNLAssociationToUpdateItem_tNLAssociationUsage_present if
                * present */
    TNLAddressWeightFactor tNLAddressWeightFactor;  /* optional; set in bit_mask
                                            * tNLAddressWeightFactor_present if
                                            * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                      * AMF_TNLAssociationToUpdateItem_iE_Extensions_present if
                      * present */
} AMF_TNLAssociationToUpdateItem;

typedef struct AMF_TNLAssociationToUpdateList_ {
    struct AMF_TNLAssociationToUpdateList_ *next;
    AMF_TNLAssociationToUpdateItem value;
} *AMF_TNLAssociationToUpdateList;

typedef struct AreaOfInterest {
    unsigned char   bit_mask;
#       define      areaOfInterestTAIList_present 0x80
#       define      areaOfInterestCellList_present 0x40
#       define      areaOfInterestRANNodeList_present 0x20
#       define      AreaOfInterest_iE_Extensions_present 0x10
    struct AreaOfInterestTAIList_ *areaOfInterestTAIList;  /* optional; set in
                                   * bit_mask areaOfInterestTAIList_present if
                                   * present */
    struct AreaOfInterestCellList_ *areaOfInterestCellList;  /* optional; set in
                                   * bit_mask areaOfInterestCellList_present if
                                   * present */
    struct AreaOfInterestRANNodeList_ *areaOfInterestRANNodeList;  /* optional;
                                   * set in bit_mask
                                   * areaOfInterestRANNodeList_present if
                                   * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * AreaOfInterest_iE_Extensions_present if
                                      * present */
} AreaOfInterest;

typedef struct NRCellIdentity {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} NRCellIdentity;

typedef struct NR_CGI {
    unsigned char   bit_mask;
#       define      NR_CGI_iE_Extensions_present 0x80
    PLMNIdentity    pLMNIdentity;
    NRCellIdentity  nRCellIdentity;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * NR_CGI_iE_Extensions_present if
                                      * present */
} NR_CGI;

typedef struct EUTRACellIdentity {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} EUTRACellIdentity;

typedef struct EUTRA_CGI {
    unsigned char   bit_mask;
#       define      EUTRA_CGI_iE_Extensions_present 0x80
    PLMNIdentity    pLMNIdentity;
    EUTRACellIdentity eUTRACellIdentity;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * EUTRA_CGI_iE_Extensions_present if
                                      * present */
} EUTRA_CGI;

typedef struct NGRAN_CGI {
    unsigned short  choice;
#       define      nR_CGI_chosen 1
#       define      eUTRA_CGI_chosen 2
#       define      NGRAN_CGI_choice_Extensions_chosen 3
    union _union {
        NR_CGI          nR_CGI;  /* to choose, set choice to nR_CGI_chosen */
        EUTRA_CGI       eUTRA_CGI;  /* to choose, set choice to
                                     * eUTRA_CGI_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                     * to NGRAN_CGI_choice_Extensions_chosen */
    } u;
} NGRAN_CGI;

typedef struct AreaOfInterestCellItem {
    unsigned char   bit_mask;
#       define      AreaOfInterestCellItem_iE_Extensions_present 0x80
    NGRAN_CGI       nGRAN_CGI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                              * AreaOfInterestCellItem_iE_Extensions_present if
                              * present */
} AreaOfInterestCellItem;

typedef struct AreaOfInterestCellList_ {
    struct AreaOfInterestCellList_ *next;
    AreaOfInterestCellItem value;
} *AreaOfInterestCellList;

typedef long            LocationReportingReferenceID;

typedef struct AreaOfInterestItem {
    unsigned char   bit_mask;
#       define      AreaOfInterestItem_iE_Extensions_present 0x80
    AreaOfInterest  areaOfInterest;
    LocationReportingReferenceID locationReportingReferenceID;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * AreaOfInterestItem_iE_Extensions_present
                                      * if present */
} AreaOfInterestItem;

typedef struct AreaOfInterestList_ {
    struct AreaOfInterestList_ *next;
    AreaOfInterestItem value;
} *AreaOfInterestList;

typedef struct AreaOfInterestRANNodeItem {
    unsigned char   bit_mask;
#       define      AreaOfInterestRANNodeItem_iE_Extensions_present 0x80
    GlobalRANNodeID globalRANNodeID;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                           * AreaOfInterestRANNodeItem_iE_Extensions_present if
                           * present */
} AreaOfInterestRANNodeItem;

typedef struct AreaOfInterestRANNodeList_ {
    struct AreaOfInterestRANNodeList_ *next;
    AreaOfInterestRANNodeItem value;
} *AreaOfInterestRANNodeList;

typedef struct AreaOfInterestTAIItem {
    unsigned char   bit_mask;
#       define      AreaOfInterestTAIItem_iE_Extensions_present 0x80
    TAI             tAI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                               * AreaOfInterestTAIItem_iE_Extensions_present if
                               * present */
} AreaOfInterestTAIItem;

typedef struct AreaOfInterestTAIList_ {
    struct AreaOfInterestTAIList_ *next;
    AreaOfInterestTAIItem value;
} *AreaOfInterestTAIList;

typedef struct RecommendedCellsForPaging {
    unsigned char   bit_mask;
#       define      RecommendedCellsForPaging_iE_Extensions_present 0x80
    struct RecommendedCellList_ *recommendedCellList;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                           * RecommendedCellsForPaging_iE_Extensions_present if
                           * present */
} RecommendedCellsForPaging;

typedef struct AssistanceDataForRecommendedCells {
    unsigned char   bit_mask;
#       define      AssistanceDataForRecommendedCells_iE_Extensions_present 0x80
    RecommendedCellsForPaging recommendedCellsForPaging;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                   * AssistanceDataForRecommendedCells_iE_Extensions_present if
                   * present */
} AssistanceDataForRecommendedCells;

typedef long            PagingAttemptCount;

typedef long            IntendedNumberOfPagingAttempts;

typedef enum NextPagingAreaScope {
    same = 0,
    changed = 1
} NextPagingAreaScope;

typedef struct PagingAttemptInformation {
    unsigned char   bit_mask;
#       define      nextPagingAreaScope_present 0x80
#       define      PagingAttemptInformation_iE_Extensions_present 0x40
    PagingAttemptCount pagingAttemptCount;
    IntendedNumberOfPagingAttempts intendedNumberOfPagingAttempts;
    NextPagingAreaScope nextPagingAreaScope;  /* optional; set in bit_mask
                                               * nextPagingAreaScope_present if
                                               * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                            * PagingAttemptInformation_iE_Extensions_present if
                            * present */
} PagingAttemptInformation;

typedef struct AssistanceDataForPaging {
    unsigned char   bit_mask;
#       define      assistanceDataForRecommendedCells_present 0x80
#       define      pagingAttemptInformation_present 0x40
#       define      AssistanceDataForPaging_iE_Extensions_present 0x20
    AssistanceDataForRecommendedCells assistanceDataForRecommendedCells;  
                                        /* optional; set in bit_mask
                                 * assistanceDataForRecommendedCells_present if
                                 * present */
    PagingAttemptInformation pagingAttemptInformation;  /* optional; set in
                                   * bit_mask pagingAttemptInformation_present
                                   * if present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                             * AssistanceDataForPaging_iE_Extensions_present if
                             * present */
} AssistanceDataForPaging;

typedef long            QosFlowIdentifier;

typedef enum _enum1 {
    ul = 0,
    dl = 1
} _enum1;

typedef struct AssociatedQosFlowItem {
    unsigned char   bit_mask;
#       define      qosFlowMappingIndication_present 0x80
#       define      AssociatedQosFlowItem_iE_Extensions_present 0x40
    QosFlowIdentifier qosFlowIdentifier;
    _enum1          qosFlowMappingIndication;  /* optional; set in bit_mask
                                          * qosFlowMappingIndication_present if
                                          * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                               * AssociatedQosFlowItem_iE_Extensions_present if
                               * present */
} AssociatedQosFlowItem;

typedef struct AssociatedQosFlowList_ {
    struct AssociatedQosFlowList_ *next;
    AssociatedQosFlowItem value;
} *AssociatedQosFlowList;

typedef long            AveragingWindow;

// B
typedef long            BitRate;

typedef struct BroadcastCancelledAreaList {
    unsigned short  choice;
#       define      cellIDCancelledEUTRA_chosen 1
#       define      tAICancelledEUTRA_chosen 2
#       define      emergencyAreaIDCancelledEUTRA_chosen 3
#       define      cellIDCancelledNR_chosen 4
#       define      tAICancelledNR_chosen 5
#       define      emergencyAreaIDCancelledNR_chosen 6
#       define      BroadcastCancelledAreaList_choice_Extensions_chosen 7
    union _union {
        struct CellIDCancelledEUTRA_ *cellIDCancelledEUTRA;  /* to choose, set
                                     * choice to cellIDCancelledEUTRA_chosen */
        struct TAICancelledEUTRA_ *tAICancelledEUTRA;  /* to choose, set choice
                                               * to tAICancelledEUTRA_chosen */
        struct EmergencyAreaIDCancelledEUTRA_ *emergencyAreaIDCancelledEUTRA;                                           /* to choose, set choice to
                                      * emergencyAreaIDCancelledEUTRA_chosen */
        struct CellIDCancelledNR_ *cellIDCancelledNR;  /* to choose, set choice
                                               * to cellIDCancelledNR_chosen */
        struct TAICancelledNR_ *tAICancelledNR;  /* to choose, set choice to
                                                  * tAICancelledNR_chosen */
        struct EmergencyAreaIDCancelledNR_ *emergencyAreaIDCancelledNR;  /* to
                                   * choose, set choice to
                                   * emergencyAreaIDCancelledNR_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to
                       * BroadcastCancelledAreaList_choice_Extensions_chosen */
    } u;
} BroadcastCancelledAreaList;

typedef struct BroadcastCompletedAreaList {
    unsigned short  choice;
#       define      cellIDBroadcastEUTRA_chosen 1
#       define      tAIBroadcastEUTRA_chosen 2
#       define      emergencyAreaIDBroadcastEUTRA_chosen 3
#       define      cellIDBroadcastNR_chosen 4
#       define      tAIBroadcastNR_chosen 5
#       define      emergencyAreaIDBroadcastNR_chosen 6
#       define      BroadcastCompletedAreaList_choice_Extensions_chosen 7
    union _union {
        struct CellIDBroadcastEUTRA_ *cellIDBroadcastEUTRA;  /* to choose, set
                                     * choice to cellIDBroadcastEUTRA_chosen */
        struct TAIBroadcastEUTRA_ *tAIBroadcastEUTRA;  /* to choose, set choice
                                               * to tAIBroadcastEUTRA_chosen */
        struct EmergencyAreaIDBroadcastEUTRA_ *emergencyAreaIDBroadcastEUTRA;                                           /* to choose, set choice to
                                      * emergencyAreaIDBroadcastEUTRA_chosen */
        struct CellIDBroadcastNR_ *cellIDBroadcastNR;  /* to choose, set choice
                                               * to cellIDBroadcastNR_chosen */
        struct TAIBroadcastNR_ *tAIBroadcastNR;  /* to choose, set choice to
                                                  * tAIBroadcastNR_chosen */
        struct EmergencyAreaIDBroadcastNR_ *emergencyAreaIDBroadcastNR;  /* to
                                   * choose, set choice to
                                   * emergencyAreaIDBroadcastNR_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to
                       * BroadcastCompletedAreaList_choice_Extensions_chosen */
    } u;
} BroadcastCompletedAreaList;

typedef struct BroadcastPLMNItem {
    unsigned char   bit_mask;
#       define      BroadcastPLMNItem_iE_Extensions_present 0x80
    PLMNIdentity    pLMNIdentity;
    struct SliceSupportList_ *tAISliceSupportList;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * BroadcastPLMNItem_iE_Extensions_present
                                      * if present */
} BroadcastPLMNItem;

typedef struct BroadcastPLMNList_ {
    struct BroadcastPLMNList_ *next;
    BroadcastPLMNItem value;
} *BroadcastPLMNList;

typedef unsigned short  NumberOfBroadcasts;

typedef struct CancelledCellsInEAI_EUTRA_Item {
    unsigned char   bit_mask;
#       define      CancelledCellsInEAI_EUTRA_Item_iE_Extensions_present 0x80
    EUTRA_CGI       eUTRA_CGI;
    NumberOfBroadcasts numberOfBroadcasts;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                      * CancelledCellsInEAI_EUTRA_Item_iE_Extensions_present if
                      * present */
} CancelledCellsInEAI_EUTRA_Item;

typedef struct CancelledCellsInEAI_EUTRA_ {
    struct CancelledCellsInEAI_EUTRA_ *next;
    CancelledCellsInEAI_EUTRA_Item value;
} *CancelledCellsInEAI_EUTRA;

typedef struct CancelledCellsInEAI_NR_Item {
    unsigned char   bit_mask;
#       define      CancelledCellsInEAI_NR_Item_iE_Extensions_present 0x80
    NR_CGI          nR_CGI;
    NumberOfBroadcasts numberOfBroadcasts;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                         * CancelledCellsInEAI_NR_Item_iE_Extensions_present if
                         * present */
} CancelledCellsInEAI_NR_Item;

typedef struct CancelledCellsInEAI_NR_ {
    struct CancelledCellsInEAI_NR_ *next;
    CancelledCellsInEAI_NR_Item value;
} *CancelledCellsInEAI_NR;

typedef struct CancelledCellsInTAI_EUTRA_Item {
    unsigned char   bit_mask;
#       define      CancelledCellsInTAI_EUTRA_Item_iE_Extensions_present 0x80
    EUTRA_CGI       eUTRA_CGI;
    NumberOfBroadcasts numberOfBroadcasts;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                      * CancelledCellsInTAI_EUTRA_Item_iE_Extensions_present if
                      * present */
} CancelledCellsInTAI_EUTRA_Item;

typedef struct CancelledCellsInTAI_EUTRA_ {
    struct CancelledCellsInTAI_EUTRA_ *next;
    CancelledCellsInTAI_EUTRA_Item value;
} *CancelledCellsInTAI_EUTRA;

typedef struct CancelledCellsInTAI_NR_Item {
    unsigned char   bit_mask;
#       define      CancelledCellsInTAI_NR_Item_iE_Extensions_present 0x80
    NR_CGI          nR_CGI;
    NumberOfBroadcasts numberOfBroadcasts;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                         * CancelledCellsInTAI_NR_Item_iE_Extensions_present if
                         * present */
} CancelledCellsInTAI_NR_Item;

typedef struct CancelledCellsInTAI_NR_ {
    struct CancelledCellsInTAI_NR_ *next;
    CancelledCellsInTAI_NR_Item value;
} *CancelledCellsInTAI_NR;

typedef enum CauseRadioNetwork {
    CauseRadioNetwork_unspecified = 0,
    txnrelocoverall_expiry = 1,
    successful_handover = 2,
    release_due_to_ngran_generated_reason = 3,
    release_due_to_5gc_generated_reason = 4,
    handover_cancelled = 5,
    partial_handover = 6,
    ho_failure_in_target_5GC_ngran_node_or_target_system = 7,
    ho_target_not_allowed = 8,
    tngrelocoverall_expiry = 9,
    tngrelocprep_expiry = 10,
    cell_not_available = 11,
    unknown_targetID = 12,
    no_radio_resources_available_in_target_cell = 13,
    unknown_local_UE_NGAP_ID = 14,
    inconsistent_remote_UE_NGAP_ID = 15,
    handover_desirable_for_radio_reason = 16,
    time_critical_handover = 17,
    resource_optimisation_handover = 18,
    reduce_load_in_serving_cell = 19,
    user_inactivity = 20,
    radio_connection_with_ue_lost = 21,
    radio_resources_not_available = 22,
    invalid_qos_combination = 23,
    failure_in_radio_interface_procedure = 24,
    interaction_with_other_procedure = 25,
    unknown_PDU_session_ID = 26,
    unkown_qos_flow_ID = 27,
    multiple_PDU_session_ID_instances = 28,
    multiple_qos_flow_ID_instances = 29,
    encryption_and_or_integrity_protection_algorithms_not_supported = 30,
    ng_intra_system_handover_triggered = 31,
    ng_inter_system_handover_triggered = 32,
    xn_handover_triggered = 33,
    not_supported_5QI_value = 34,
    ue_context_transfer = 35,
    ims_voice_eps_fallback_or_rat_fallback_triggered = 36,
    up_integrity_protection_not_possible = 37,
    up_confidentiality_protection_not_possible = 38,
    slice_not_supported = 39,
    ue_in_rrc_inactive_state_not_reachable = 40,
    redirection = 41,
    resources_not_available_for_the_slice = 42,
    ue_max_integrity_protected_data_rate_reason = 43,
    release_due_to_cn_detected_mobility = 44,
    n26_interface_not_available = 45,
    release_due_to_pre_emption = 46
} CauseRadioNetwork;

typedef enum CauseTransport {
    transport_resource_unavailable = 0,
    CauseTransport_unspecified = 1
} CauseTransport;

typedef enum CauseNas {
    normal_release = 0,
    authentication_failure = 1,
    deregister = 2,
    CauseNas_unspecified = 3
} CauseNas;

typedef enum CauseProtocol {
    transfer_syntax_error = 0,
    abstract_syntax_error_reject = 1,
    abstract_syntax_error_ignore_and_notify = 2,
    message_not_compatible_with_receiver_state = 3,
    semantic_error = 4,
    abstract_syntax_error_falsely_constructed_message = 5,
    CauseProtocol_unspecified = 6
} CauseProtocol;

typedef enum CauseMisc {
    control_processing_overload = 0,
    not_enough_user_plane_processing_resources = 1,
    hardware_failure = 2,
    om_intervention = 3,
    unknown_PLMN = 4,
    CauseMisc_unspecified = 5
} CauseMisc;

typedef struct Cause {
    unsigned short  choice;
#       define      radioNetwork_chosen 1
#       define      transport_chosen 2
#       define      nas_chosen 3
#       define      protocol_chosen 4
#       define      misc_chosen 5
#       define      Cause_choice_Extensions_chosen 6
    union _union {
        CauseRadioNetwork radioNetwork;  /* to choose, set choice to
                                          * radioNetwork_chosen */
        CauseTransport  transport;  /* to choose, set choice to
                                     * transport_chosen */
        CauseNas        nas;  /* to choose, set choice to nas_chosen */
        CauseProtocol   protocol;  /* to choose, set choice to
                                    * protocol_chosen */
        CauseMisc       misc;  /* to choose, set choice to misc_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                         * to Cause_choice_Extensions_chosen */
    } u;
} Cause;

typedef struct CellIDBroadcastEUTRA_Item {
    unsigned char   bit_mask;
#       define      CellIDBroadcastEUTRA_Item_iE_Extensions_present 0x80
    EUTRA_CGI       eUTRA_CGI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                           * CellIDBroadcastEUTRA_Item_iE_Extensions_present if
                           * present */
} CellIDBroadcastEUTRA_Item;

typedef struct CellIDBroadcastEUTRA_ {
    struct CellIDBroadcastEUTRA_ *next;
    CellIDBroadcastEUTRA_Item value;
} *CellIDBroadcastEUTRA;

typedef struct CellIDBroadcastNR_Item {
    unsigned char   bit_mask;
#       define      CellIDBroadcastNR_Item_iE_Extensions_present 0x80
    NR_CGI          nR_CGI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                              * CellIDBroadcastNR_Item_iE_Extensions_present if
                              * present */
} CellIDBroadcastNR_Item;

typedef struct CellIDBroadcastNR_ {
    struct CellIDBroadcastNR_ *next;
    CellIDBroadcastNR_Item value;
} *CellIDBroadcastNR;

typedef struct CellIDCancelledEUTRA_Item {
    unsigned char   bit_mask;
#       define      CellIDCancelledEUTRA_Item_iE_Extensions_present 0x80
    EUTRA_CGI       eUTRA_CGI;
    NumberOfBroadcasts numberOfBroadcasts;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                           * CellIDCancelledEUTRA_Item_iE_Extensions_present if
                           * present */
} CellIDCancelledEUTRA_Item;

typedef struct CellIDCancelledEUTRA_ {
    struct CellIDCancelledEUTRA_ *next;
    CellIDCancelledEUTRA_Item value;
} *CellIDCancelledEUTRA;

typedef struct CellIDCancelledNR_Item {
    unsigned char   bit_mask;
#       define      CellIDCancelledNR_Item_iE_Extensions_present 0x80
    NR_CGI          nR_CGI;
    NumberOfBroadcasts numberOfBroadcasts;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                              * CellIDCancelledNR_Item_iE_Extensions_present if
                              * present */
} CellIDCancelledNR_Item;

typedef struct CellIDCancelledNR_ {
    struct CellIDCancelledNR_ *next;
    CellIDCancelledNR_Item value;
} *CellIDCancelledNR;

typedef struct CellIDListForRestart {
    unsigned short  choice;
#       define      eUTRA_CGIListforRestart_chosen 1
#       define      nR_CGIListforRestart_chosen 2
#       define      CellIDListForRestart_choice_Extensions_chosen 3
    union _union {
        struct EUTRA_CGIList_ *eUTRA_CGIListforRestart;  /* to choose, set
                                   * choice to eUTRA_CGIListforRestart_chosen */
        struct NR_CGIList_ *nR_CGIListforRestart;  /* to choose, set choice to
                                               * nR_CGIListforRestart_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to
                             * CellIDListForRestart_choice_Extensions_chosen */
    } u;
} CellIDListForRestart;

typedef long            ExpectedActivityPeriod;

typedef long            ExpectedIdlePeriod;

typedef enum SourceOfUEActivityBehaviourInformation {
    subscription_information = 0,
    statistics = 1
} SourceOfUEActivityBehaviourInformation;

typedef struct ExpectedUEActivityBehaviour {
    unsigned char   bit_mask;
#       define      expectedActivityPeriod_present 0x80
#       define      expectedIdlePeriod_present 0x40
#       define      sourceOfUEActivityBehaviourInformation_present 0x20
#       define      ExpectedUEActivityBehaviour_iE_Extensions_present 0x10
    ExpectedActivityPeriod expectedActivityPeriod;  /* optional; set in bit_mask
                                            * expectedActivityPeriod_present if
                                            * present */
    ExpectedIdlePeriod expectedIdlePeriod;  /* optional; set in bit_mask
                                             * expectedIdlePeriod_present if
                                             * present */
    SourceOfUEActivityBehaviourInformation sourceOfUEActivityBehaviourInformation;                                      /* optional; set in bit_mask
                            * sourceOfUEActivityBehaviourInformation_present if
                            * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                         * ExpectedUEActivityBehaviour_iE_Extensions_present if
                         * present */
} ExpectedUEActivityBehaviour;

typedef enum ExpectedHOInterval {
    sec15 = 0,
    sec30 = 1,
    sec60 = 2,
    sec90 = 3,
    sec120 = 4,
    sec180 = 5,
    long_time = 6
} ExpectedHOInterval;

typedef enum ExpectedUEMobility {
    stationary = 0,
    mobile = 1
} ExpectedUEMobility;

typedef struct ExpectedUEBehaviour {
    unsigned char   bit_mask;
#       define      expectedUEActivityBehaviour_present 0x80
#       define      expectedHOInterval_present 0x40
#       define      expectedUEMobility_present 0x20
#       define      expectedUEMovingTrajectory_present 0x10
#       define      ExpectedUEBehaviour_iE_Extensions_present 0x08
    ExpectedUEActivityBehaviour expectedUEActivityBehaviour;  /* optional; set
                                   * in bit_mask
                                   * expectedUEActivityBehaviour_present if
                                   * present */
    ExpectedHOInterval expectedHOInterval;  /* optional; set in bit_mask
                                             * expectedHOInterval_present if
                                             * present */
    ExpectedUEMobility expectedUEMobility;  /* optional; set in bit_mask
                                             * expectedUEMobility_present if
                                             * present */
    struct ExpectedUEMovingTrajectory_ *expectedUEMovingTrajectory;  
                                  /* optional; set in bit_mask
                                   * expectedUEMovingTrajectory_present if
                                   * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                 * ExpectedUEBehaviour_iE_Extensions_present if
                                 * present */
} ExpectedUEBehaviour;

typedef struct CNAssistedRANTuning {
    unsigned char   bit_mask;
#       define      CNAssistedRANTuning_expectedUEBehaviour_present 0x80
#       define      CNAssistedRANTuning_iE_Extensions_present 0x40
    ExpectedUEBehaviour expectedUEBehaviour;  /* optional; set in bit_mask
                           * CNAssistedRANTuning_expectedUEBehaviour_present if
                           * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                 * CNAssistedRANTuning_iE_Extensions_present if
                                 * present */
} CNAssistedRANTuning;

typedef enum _enum2 {
    cn_Type_epc_forbidden = 0,
    fiveGC_forbidden = 1
} _enum2;

typedef struct CNTypeRestrictionsForEquivalentItem {
    unsigned char   bit_mask;
#       define      CNTypeRestrictionsForEquivalentItem_iE_Extensions_present 0x80
    PLMNIdentity    plmnIdentity;
    _enum2          cn_Type;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                 * CNTypeRestrictionsForEquivalentItem_iE_Extensions_present if
                 * present */
} CNTypeRestrictionsForEquivalentItem;

typedef struct CNTypeRestrictionsForEquivalent_ {
    struct CNTypeRestrictionsForEquivalent_ *next;
    CNTypeRestrictionsForEquivalentItem value;
} *CNTypeRestrictionsForEquivalent;

typedef struct CommonNetworkInstance {
    unsigned int    length;
    unsigned char   *value;
} CommonNetworkInstance;

typedef struct CompletedCellsInEAI_EUTRA_Item {
    unsigned char   bit_mask;
#       define      CompletedCellsInEAI_EUTRA_Item_iE_Extensions_present 0x80
    EUTRA_CGI       eUTRA_CGI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                      * CompletedCellsInEAI_EUTRA_Item_iE_Extensions_present if
                      * present */
} CompletedCellsInEAI_EUTRA_Item;

typedef struct CompletedCellsInEAI_EUTRA_ {
    struct CompletedCellsInEAI_EUTRA_ *next;
    CompletedCellsInEAI_EUTRA_Item value;
} *CompletedCellsInEAI_EUTRA;

typedef struct CompletedCellsInEAI_NR_Item {
    unsigned char   bit_mask;
#       define      CompletedCellsInEAI_NR_Item_iE_Extensions_present 0x80
    NR_CGI          nR_CGI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                         * CompletedCellsInEAI_NR_Item_iE_Extensions_present if
                         * present */
} CompletedCellsInEAI_NR_Item;

typedef struct CompletedCellsInEAI_NR_ {
    struct CompletedCellsInEAI_NR_ *next;
    CompletedCellsInEAI_NR_Item value;
} *CompletedCellsInEAI_NR;

typedef struct CompletedCellsInTAI_EUTRA_Item {
    unsigned char   bit_mask;
#       define      CompletedCellsInTAI_EUTRA_Item_iE_Extensions_present 0x80
    EUTRA_CGI       eUTRA_CGI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                      * CompletedCellsInTAI_EUTRA_Item_iE_Extensions_present if
                      * present */
} CompletedCellsInTAI_EUTRA_Item;

typedef struct CompletedCellsInTAI_EUTRA_ {
    struct CompletedCellsInTAI_EUTRA_ *next;
    CompletedCellsInTAI_EUTRA_Item value;
} *CompletedCellsInTAI_EUTRA;

typedef struct CompletedCellsInTAI_NR_Item {
    unsigned char   bit_mask;
#       define      CompletedCellsInTAI_NR_Item_iE_Extensions_present 0x80
    NR_CGI          nR_CGI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                         * CompletedCellsInTAI_NR_Item_iE_Extensions_present if
                         * present */
} CompletedCellsInTAI_NR_Item;

typedef struct CompletedCellsInTAI_NR_ {
    struct CompletedCellsInTAI_NR_ *next;
    CompletedCellsInTAI_NR_Item value;
} *CompletedCellsInTAI_NR;

typedef enum ConfidentialityProtectionIndication {
    ConfidentialityProtectionIndication_required = 0,
    ConfidentialityProtectionIndication_preferred = 1,
    ConfidentialityProtectionIndication_not_needed = 2
} ConfidentialityProtectionIndication;

typedef enum ConfidentialityProtectionResult {
    ConfidentialityProtectionResult_performed = 0,
    ConfidentialityProtectionResult_not_performed = 1
} ConfidentialityProtectionResult;

typedef struct UEIdentityIndexValue {
    unsigned short  choice;
#       define      indexLength10_chosen 1
#       define      UEIdentityIndexValue_choice_Extensions_chosen 2
    union _union {
        _bit1           indexLength10;  /* to choose, set choice to
                                         * indexLength10_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to
                             * UEIdentityIndexValue_choice_Extensions_chosen */
    } u;
} UEIdentityIndexValue;

typedef struct PeriodicRegistrationUpdateTimer {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} PeriodicRegistrationUpdateTimer;

typedef enum MICOModeIndication {
    MICOModeIndication_true = 0
} MICOModeIndication;

typedef struct CoreNetworkAssistanceInformationForInactive {
    unsigned char   bit_mask;
#       define      uESpecificDRX_present 0x80
#       define      mICOModeIndication_present 0x40
#       define      CoreNetworkAssistanceInformationForInactive_expectedUEBehaviour_present 0x20
#       define      CoreNetworkAssistanceInformationForInactive_iE_Extensions_present 0x10
    UEIdentityIndexValue uEIdentityIndexValue;
    PagingDRX       uESpecificDRX;  /* optional; set in bit_mask
                                     * uESpecificDRX_present if present */
    PeriodicRegistrationUpdateTimer periodicRegistrationUpdateTimer;
    MICOModeIndication mICOModeIndication;  /* optional; set in bit_mask
                                             * mICOModeIndication_present if
                                             * present */
    struct TAIListForInactive_ *tAIListForInactive;
    ExpectedUEBehaviour expectedUEBehaviour;  /* optional; set in bit_mask
   * CoreNetworkAssistanceInformationForInactive_expectedUEBehaviour_present if
   * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
         * CoreNetworkAssistanceInformationForInactive_iE_Extensions_present if
         * present */
} CoreNetworkAssistanceInformationForInactive;

typedef struct COUNTValueForPDCP_SN12 {
    unsigned char   bit_mask;
#       define      COUNTValueForPDCP_SN12_iE_Extensions_present 0x80
    unsigned short  pDCP_SN12;
    unsigned int    hFN_PDCP_SN12;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                              * COUNTValueForPDCP_SN12_iE_Extensions_present if
                              * present */
} COUNTValueForPDCP_SN12;

typedef struct COUNTValueForPDCP_SN18 {
    unsigned char   bit_mask;
#       define      COUNTValueForPDCP_SN18_iE_Extensions_present 0x80
    unsigned int    pDCP_SN18;
    unsigned short  hFN_PDCP_SN18;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                              * COUNTValueForPDCP_SN18_iE_Extensions_present if
                              * present */
} COUNTValueForPDCP_SN18;

typedef struct CriticalityDiagnostics {
    unsigned char   bit_mask;
#       define      procedureCode_present 0x80
#       define      triggeringMessage_present 0x40
#       define      procedureCriticality_present 0x20
#       define      iEsCriticalityDiagnostics_present 0x10
#       define      CriticalityDiagnostics_iE_Extensions_present 0x08
    ProcedureCode   procedureCode;  /* optional; set in bit_mask
                                     * procedureCode_present if present */
    TriggeringMessage triggeringMessage;  /* optional; set in bit_mask
                                           * triggeringMessage_present if
                                           * present */
    Criticality     procedureCriticality;  /* optional; set in bit_mask
                                            * procedureCriticality_present if
                                            * present */
    struct CriticalityDiagnostics_IE_List_ *iEsCriticalityDiagnostics;  
                                  /* optional; set in bit_mask
                                   * iEsCriticalityDiagnostics_present if
                                   * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                              * CriticalityDiagnostics_iE_Extensions_present if
                              * present */
} CriticalityDiagnostics;

typedef enum TypeOfError {
    not_understood = 0,
    missing = 1
} TypeOfError;

typedef struct CriticalityDiagnostics_IE_Item {
    unsigned char   bit_mask;
#       define      CriticalityDiagnostics_IE_Item_iE_Extensions_present 0x80
    Criticality     iECriticality;
    ProtocolIE_ID   iE_ID;
    TypeOfError     typeOfError;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                      * CriticalityDiagnostics_IE_Item_iE_Extensions_present if
                      * present */
} CriticalityDiagnostics_IE_Item;

typedef struct CriticalityDiagnostics_IE_List_ {
    struct CriticalityDiagnostics_IE_List_ *next;
    CriticalityDiagnostics_IE_Item value;
} *CriticalityDiagnostics_IE_List;

// D
typedef struct DataCodingScheme {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} DataCodingScheme;

typedef enum DataForwardingAccepted {
    data_forwarding_accepted = 0
} DataForwardingAccepted;

typedef long            DRB_ID;

typedef struct DataForwardingResponseDRBItem {
    unsigned char   bit_mask;
#       define      DataForwardingResponseDRBItem_dLForwardingUP_TNLInformation_present 0x80
#       define      uLForwardingUP_TNLInformation_present 0x40
#       define      DataForwardingResponseDRBItem_iE_Extensions_present 0x20
    DRB_ID          dRB_ID;
    UPTransportLayerInformation dLForwardingUP_TNLInformation;  /* optional; set
                                   * in bit_mask
       * DataForwardingResponseDRBItem_dLForwardingUP_TNLInformation_present if
       * present */
    UPTransportLayerInformation uLForwardingUP_TNLInformation;  /* optional; set
                                   * in bit_mask
                                   * uLForwardingUP_TNLInformation_present if
                                   * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                       * DataForwardingResponseDRBItem_iE_Extensions_present if
                       * present */
} DataForwardingResponseDRBItem;

typedef struct DataForwardingResponseDRBList_ {
    struct DataForwardingResponseDRBList_ *next;
    DataForwardingResponseDRBItem value;
} *DataForwardingResponseDRBList;

typedef enum DelayCritical {
    delay_critical = 0,
    non_delay_critical = 1
} DelayCritical;

typedef enum DL_NGU_TNLInformationReused {
    DL_NGU_TNLInformationReused_true = 0
} DL_NGU_TNLInformationReused;

typedef struct DRBStatusUL12 {
    unsigned char   bit_mask;
#       define      DRBStatusUL12_receiveStatusOfUL_PDCP_SDUs_present 0x80
#       define      DRBStatusUL12_iE_Extension_present 0x40
    COUNTValueForPDCP_SN12 uL_COUNTValue;
    _bit1           receiveStatusOfUL_PDCP_SDUs;  /* optional; set in bit_mask
                         * DRBStatusUL12_receiveStatusOfUL_PDCP_SDUs_present if
                         * present */
    struct _seqof1_ *iE_Extension;  /* optional; set in bit_mask
                                     * DRBStatusUL12_iE_Extension_present if
                                     * present */
} DRBStatusUL12;

typedef struct DRBStatusUL18 {
    unsigned char   bit_mask;
#       define      DRBStatusUL18_receiveStatusOfUL_PDCP_SDUs_present 0x80
#       define      DRBStatusUL18_iE_Extension_present 0x40
    COUNTValueForPDCP_SN18 uL_COUNTValue;
    struct _bit2 {
        unsigned int    length;  /* number of significant bits */
        unsigned char   *value;
    } receiveStatusOfUL_PDCP_SDUs;  /* optional; set in bit_mask
                         * DRBStatusUL18_receiveStatusOfUL_PDCP_SDUs_present if
                         * present */
    struct _seqof1_ *iE_Extension;  /* optional; set in bit_mask
                                     * DRBStatusUL18_iE_Extension_present if
                                     * present */
} DRBStatusUL18;

typedef struct DRBStatusUL {
    unsigned short  choice;
#       define      dRBStatusUL12_chosen 1
#       define      dRBStatusUL18_chosen 2
#       define      DRBStatusUL_choice_Extensions_chosen 3
    union _union {
        DRBStatusUL12   dRBStatusUL12;  /* to choose, set choice to
                                         * dRBStatusUL12_chosen */
        DRBStatusUL18   dRBStatusUL18;  /* to choose, set choice to
                                         * dRBStatusUL18_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to DRBStatusUL_choice_Extensions_chosen */
    } u;
} DRBStatusUL;

typedef struct DRBStatusDL12 {
    unsigned char   bit_mask;
#       define      DRBStatusDL12_iE_Extension_present 0x80
    COUNTValueForPDCP_SN12 dL_COUNTValue;
    struct _seqof1_ *iE_Extension;  /* optional; set in bit_mask
                                     * DRBStatusDL12_iE_Extension_present if
                                     * present */
} DRBStatusDL12;

typedef struct DRBStatusDL18 {
    unsigned char   bit_mask;
#       define      DRBStatusDL18_iE_Extension_present 0x80
    COUNTValueForPDCP_SN18 dL_COUNTValue;
    struct _seqof1_ *iE_Extension;  /* optional; set in bit_mask
                                     * DRBStatusDL18_iE_Extension_present if
                                     * present */
} DRBStatusDL18;

typedef struct DRBStatusDL {
    unsigned short  choice;
#       define      dRBStatusDL12_chosen 1
#       define      dRBStatusDL18_chosen 2
#       define      DRBStatusDL_choice_Extensions_chosen 3
    union _union {
        DRBStatusDL12   dRBStatusDL12;  /* to choose, set choice to
                                         * dRBStatusDL12_chosen */
        DRBStatusDL18   dRBStatusDL18;  /* to choose, set choice to
                                         * dRBStatusDL18_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to DRBStatusDL_choice_Extensions_chosen */
    } u;
} DRBStatusDL;

enum DRBsSubjectToStatusTransferItem_ExtIEs_Extension_PDUs {
    PDU_DRBsSubjectToStatusTransferItem_ExtIEs_Extension_UNKNOWN = 0,

    PDU_DRBsSubjectToStatusTransferItem_ExtIEs_Extension_AssociatedQosFlowList = AssociatedQosFlowList_PDU,

    PDU_DRBsSubjectToStatusTransferItem_ExtIEs_Extension_MAX = INT_MAX
};

union DRBsSubjectToStatusTransferItem_ExtIEs_Extension_union {
    struct AssociatedQosFlowList_ **pdu_AssociatedQosFlowList;  
 /* PDU_DRBsSubjectToStatusTransferItem_ExtIEs_Extension_AssociatedQosFlowList */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct DRBsSubjectToStatusTransferItem_ExtIEs_Extension {
    DRBsSubjectToStatusTransferItem_ExtIEs_Extension_PDUs pduNum;
    OssBuf          encoded;
    DRBsSubjectToStatusTransferItem_ExtIEs_Extension_union decoded;
} DRBsSubjectToStatusTransferItem_ExtIEs_Extension;

typedef struct DRBsSubjectToStatusTransferItem {
    unsigned char   bit_mask;
#       define      DRBsSubjectToStatusTransferItem_iE_Extension_present 0x80
    DRB_ID          dRB_ID;
    DRBStatusUL     dRBStatusUL;
    DRBStatusDL     dRBStatusDL;
    struct _seqof53 {
        struct _seqof53 *next;
        struct _seq68 {
            ProtocolExtensionID id;
            Criticality     criticality;
            DRBsSubjectToStatusTransferItem_ExtIEs_Extension extensionValue;
        } value;
    } *iE_Extension;  /* optional; set in bit_mask
                       * DRBsSubjectToStatusTransferItem_iE_Extension_present if
                       * present */
} DRBsSubjectToStatusTransferItem;

typedef struct DRBsSubjectToStatusTransferList_ {
    struct DRBsSubjectToStatusTransferList_ *next;
    DRBsSubjectToStatusTransferItem value;
} *DRBsSubjectToStatusTransferList;

typedef long            PriorityLevelQos;

// P
typedef long            PacketDelayBudget;

typedef struct PacketErrorRate {
    unsigned char   bit_mask;
#       define      PacketErrorRate_iE_Extensions_present 0x80
    long            pERScalar;
    long            pERExponent;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * PacketErrorRate_iE_Extensions_present if
                                      * present */
} PacketErrorRate;

typedef long            FiveQI;

typedef long            MaximumDataBurstVolume;

typedef struct Dynamic5QIDescriptor {
    unsigned char   bit_mask;
#       define      fiveQI_present 0x80
#       define      delayCritical_present 0x40
#       define      Dynamic5QIDescriptor_averagingWindow_present 0x20
#       define      Dynamic5QIDescriptor_maximumDataBurstVolume_present 0x10
#       define      Dynamic5QIDescriptor_iE_Extensions_present 0x08
    PriorityLevelQos priorityLevelQos;
    PacketDelayBudget packetDelayBudget;
    PacketErrorRate packetErrorRate;
    FiveQI          fiveQI;  /* optional; set in bit_mask fiveQI_present if
                              * present */
    DelayCritical   delayCritical;  /* optional; set in bit_mask
                                     * delayCritical_present if present */
// The above IE shall be present in case of GBR QoS flow
    AveragingWindow averagingWindow;  /* optional; set in bit_mask
                              * Dynamic5QIDescriptor_averagingWindow_present if
                              * present */
// The above IE shall be present in case of GBR QoS flow
    MaximumDataBurstVolume maximumDataBurstVolume;  /* optional; set in bit_mask
                       * Dynamic5QIDescriptor_maximumDataBurstVolume_present if
                       * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                * Dynamic5QIDescriptor_iE_Extensions_present if
                                * present */
} Dynamic5QIDescriptor;

// E
typedef struct EmergencyAreaID {
    unsigned short  length;
    unsigned char   value[3];
} EmergencyAreaID;

typedef struct EmergencyAreaIDBroadcastEUTRA_Item {
    unsigned char   bit_mask;
#       define      EmergencyAreaIDBroadcastEUTRA_Item_iE_Extensions_present 0x80
    EmergencyAreaID emergencyAreaID;
    struct CompletedCellsInEAI_EUTRA_ *completedCellsInEAI_EUTRA;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                  * EmergencyAreaIDBroadcastEUTRA_Item_iE_Extensions_present if
                  * present */
} EmergencyAreaIDBroadcastEUTRA_Item;

typedef struct EmergencyAreaIDBroadcastEUTRA_ {
    struct EmergencyAreaIDBroadcastEUTRA_ *next;
    EmergencyAreaIDBroadcastEUTRA_Item value;
} *EmergencyAreaIDBroadcastEUTRA;

typedef struct EmergencyAreaIDBroadcastNR_Item {
    unsigned char   bit_mask;
#       define      EmergencyAreaIDBroadcastNR_Item_iE_Extensions_present 0x80
    EmergencyAreaID emergencyAreaID;
    struct CompletedCellsInEAI_NR_ *completedCellsInEAI_NR;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                     * EmergencyAreaIDBroadcastNR_Item_iE_Extensions_present if
                     * present */
} EmergencyAreaIDBroadcastNR_Item;

typedef struct EmergencyAreaIDBroadcastNR_ {
    struct EmergencyAreaIDBroadcastNR_ *next;
    EmergencyAreaIDBroadcastNR_Item value;
} *EmergencyAreaIDBroadcastNR;

typedef struct EmergencyAreaIDCancelledEUTRA_Item {
    unsigned char   bit_mask;
#       define      EmergencyAreaIDCancelledEUTRA_Item_iE_Extensions_present 0x80
    EmergencyAreaID emergencyAreaID;
    struct CancelledCellsInEAI_EUTRA_ *cancelledCellsInEAI_EUTRA;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                  * EmergencyAreaIDCancelledEUTRA_Item_iE_Extensions_present if
                  * present */
} EmergencyAreaIDCancelledEUTRA_Item;

typedef struct EmergencyAreaIDCancelledEUTRA_ {
    struct EmergencyAreaIDCancelledEUTRA_ *next;
    EmergencyAreaIDCancelledEUTRA_Item value;
} *EmergencyAreaIDCancelledEUTRA;

typedef struct EmergencyAreaIDCancelledNR_Item {
    unsigned char   bit_mask;
#       define      EmergencyAreaIDCancelledNR_Item_iE_Extensions_present 0x80
    EmergencyAreaID emergencyAreaID;
    struct CancelledCellsInEAI_NR_ *cancelledCellsInEAI_NR;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                     * EmergencyAreaIDCancelledNR_Item_iE_Extensions_present if
                     * present */
} EmergencyAreaIDCancelledNR_Item;

typedef struct EmergencyAreaIDCancelledNR_ {
    struct EmergencyAreaIDCancelledNR_ *next;
    EmergencyAreaIDCancelledNR_Item value;
} *EmergencyAreaIDCancelledNR;

typedef struct EmergencyAreaIDList_ {
    struct EmergencyAreaIDList_ *next;
    EmergencyAreaID value;
} *EmergencyAreaIDList;

typedef struct EmergencyAreaIDListForRestart_ {
    struct EmergencyAreaIDListForRestart_ *next;
    EmergencyAreaID value;
} *EmergencyAreaIDListForRestart;

typedef enum EmergencyFallbackRequestIndicator {
    emergency_fallback_requested = 0
} EmergencyFallbackRequestIndicator;

typedef enum EmergencyServiceTargetCN {
    fiveGC = 0,
    epc = 1
} EmergencyServiceTargetCN;

typedef struct EmergencyFallbackIndicator {
    unsigned char   bit_mask;
#       define      emergencyServiceTargetCN_present 0x80
#       define      EmergencyFallbackIndicator_iE_Extensions_present 0x40
    EmergencyFallbackRequestIndicator emergencyFallbackRequestIndicator;
    EmergencyServiceTargetCN emergencyServiceTargetCN;  /* optional; set in
                                   * bit_mask emergencyServiceTargetCN_present
                                   * if present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                          * EmergencyFallbackIndicator_iE_Extensions_present if
                          * present */
} EmergencyFallbackIndicator;

typedef struct EN_DCSONConfigurationTransfer {
    unsigned int    length;
    unsigned char   *value;
} EN_DCSONConfigurationTransfer;

typedef struct EquivalentPLMNs_ {
    struct EquivalentPLMNs_ *next;
    PLMNIdentity    value;
} *EquivalentPLMNs;

typedef struct EPS_TAC {
    unsigned short  length;
    unsigned char   value[2];
} EPS_TAC;

typedef struct EPS_TAI {
    unsigned char   bit_mask;
#       define      EPS_TAI_iE_Extensions_present 0x80
    PLMNIdentity    pLMNIdentity;
    EPS_TAC         ePS_TAC;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * EPS_TAI_iE_Extensions_present if
                                      * present */
} EPS_TAI;

typedef long            E_RAB_ID;

typedef struct EUTRA_CGIList_ {
    struct EUTRA_CGIList_ *next;
    EUTRA_CGI       value;
} *EUTRA_CGIList;

typedef struct EUTRA_CGIListForWarning_ {
    struct EUTRA_CGIListForWarning_ *next;
    EUTRA_CGI       value;
} *EUTRA_CGIListForWarning;

typedef struct EUTRAencryptionAlgorithms {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} EUTRAencryptionAlgorithms;

typedef struct EUTRAintegrityProtectionAlgorithms {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} EUTRAintegrityProtectionAlgorithms;

typedef enum EventType {
    direct = 0,
    change_of_serve_cell = 1,
    ue_presence_in_area_of_interest = 2,
    stop_change_of_serve_cell = 3,
    stop_ue_presence_in_area_of_interest = 4,
    cancel_location_reporting_for_the_ue = 5
} EventType;

typedef struct ExpectedUEMovingTrajectoryItem {
    unsigned char   bit_mask;
#       define      ExpectedUEMovingTrajectoryItem_timeStayedInCell_present 0x80
#       define      ExpectedUEMovingTrajectoryItem_iE_Extensions_present 0x40
    NGRAN_CGI       nGRAN_CGI;
    unsigned short  timeStayedInCell;  /* optional; set in bit_mask
                   * ExpectedUEMovingTrajectoryItem_timeStayedInCell_present if
                   * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                      * ExpectedUEMovingTrajectoryItem_iE_Extensions_present if
                      * present */
} ExpectedUEMovingTrajectoryItem;

typedef struct ExpectedUEMovingTrajectory_ {
    struct ExpectedUEMovingTrajectory_ *next;
    ExpectedUEMovingTrajectoryItem value;
} *ExpectedUEMovingTrajectory;

typedef struct FiveG_TMSI {
    unsigned short  length;
    unsigned char   value[4];
} FiveG_TMSI;

// F
typedef struct FiveG_S_TMSI {
    unsigned char   bit_mask;
#       define      FiveG_S_TMSI_iE_Extensions_present 0x80
    AMFSetID        aMFSetID;
    AMFPointer      aMFPointer;
    FiveG_TMSI      fiveG_TMSI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * FiveG_S_TMSI_iE_Extensions_present if
                                      * present */
} FiveG_S_TMSI;

typedef struct ForbiddenAreaInformation_Item {
    unsigned char   bit_mask;
#       define      ForbiddenAreaInformation_Item_iE_Extensions_present 0x80
    PLMNIdentity    pLMNIdentity;
    struct ForbiddenTACs_ *forbiddenTACs;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                       * ForbiddenAreaInformation_Item_iE_Extensions_present if
                       * present */
} ForbiddenAreaInformation_Item;

typedef struct ForbiddenAreaInformation_ {
    struct ForbiddenAreaInformation_ *next;
    ForbiddenAreaInformation_Item value;
} *ForbiddenAreaInformation;

typedef struct ForbiddenTACs_ {
    struct ForbiddenTACs_ *next;
    TAC             value;
} *ForbiddenTACs;

typedef enum NotificationControl {
    notification_requested = 0
} NotificationControl;

typedef long            PacketLossRate;

// G
typedef struct GBR_QosInformation {
    unsigned char   bit_mask;
#       define      notificationControl_present 0x80
#       define      maximumPacketLossRateDL_present 0x40
#       define      maximumPacketLossRateUL_present 0x20
#       define      GBR_QosInformation_iE_Extensions_present 0x10
    BitRate         maximumFlowBitRateDL;
    BitRate         maximumFlowBitRateUL;
    BitRate         guaranteedFlowBitRateDL;
    BitRate         guaranteedFlowBitRateUL;
    NotificationControl notificationControl;  /* optional; set in bit_mask
                                               * notificationControl_present if
                                               * present */
    PacketLossRate  maximumPacketLossRateDL;  /* optional; set in bit_mask
                                               * maximumPacketLossRateDL_present
                                               * if present */
    PacketLossRate  maximumPacketLossRateUL;  /* optional; set in bit_mask
                                               * maximumPacketLossRateUL_present
                                               * if present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * GBR_QosInformation_iE_Extensions_present
                                      * if present */
} GBR_QosInformation;

typedef struct GUAMI {
    unsigned char   bit_mask;
#       define      GUAMI_iE_Extensions_present 0x80
    PLMNIdentity    pLMNIdentity;
    AMFRegionID     aMFRegionID;
    AMFSetID        aMFSetID;
    AMFPointer      aMFPointer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * GUAMI_iE_Extensions_present if
                                      * present */
} GUAMI;

enum HandoverCommandTransfer_ExtIEs_Extension_PDUs {
    PDU_HandoverCommandTransfer_ExtIEs_Extension_UNKNOWN = 0,

    PDU_HandoverCommandTransfer_ExtIEs_Extension_QosFlowPerTNLInformationList = QosFlowPerTNLInformationList_PDU,
    PDU_HandoverCommandTransfer_ExtIEs_Extension_UPTransportLayerInformation = UPTransportLayerInformation_PDU,

    PDU_HandoverCommandTransfer_ExtIEs_Extension_MAX = INT_MAX
};

union HandoverCommandTransfer_ExtIEs_Extension_union {
    struct QosFlowPerTNLInformationList_ **pdu_QosFlowPerTNLInformationList;  
                                        /* PDU_HandoverCommandTransfer_ExtIEs_Extension_QosFlowPerTNLInformationList */
    UPTransportLayerInformation *pdu_UPTransportLayerInformation;  
 /* PDU_HandoverCommandTransfer_ExtIEs_Extension_UPTransportLayerInformation */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct HandoverCommandTransfer_ExtIEs_Extension {
    HandoverCommandTransfer_ExtIEs_Extension_PDUs pduNum;
    OssBuf          encoded;
    HandoverCommandTransfer_ExtIEs_Extension_union decoded;
} HandoverCommandTransfer_ExtIEs_Extension;

// H
typedef struct HandoverCommandTransfer {
    unsigned char   bit_mask;
#       define      HandoverCommandTransfer_dLForwardingUP_TNLInformation_present 0x80
#       define      qosFlowToBeForwardedList_present 0x40
#       define      HandoverCommandTransfer_dataForwardingResponseDRBList_present 0x20
#       define      HandoverCommandTransfer_iE_Extensions_present 0x10
    UPTransportLayerInformation dLForwardingUP_TNLInformation;  /* optional; set
                                   * in bit_mask
             * HandoverCommandTransfer_dLForwardingUP_TNLInformation_present if
             * present */
    struct QosFlowToBeForwardedList_ *qosFlowToBeForwardedList;  /* optional;
                                   * set in bit_mask
                                   * qosFlowToBeForwardedList_present if
                                   * present */
    struct DataForwardingResponseDRBList_ *dataForwardingResponseDRBList;  
                                        /* optional; set in bit_mask
             * HandoverCommandTransfer_dataForwardingResponseDRBList_present if
             * present */
    struct _seqof67 {
        struct _seqof67 *next;
        struct _seq82 {
            ProtocolExtensionID id;
            Criticality     criticality;
            HandoverCommandTransfer_ExtIEs_Extension extensionValue;
        } value;
    } *iE_Extensions;  /* optional; set in bit_mask
                        * HandoverCommandTransfer_iE_Extensions_present if
                        * present */
} HandoverCommandTransfer;

typedef struct HandoverPreparationUnsuccessfulTransfer {
    unsigned char   bit_mask;
#       define      HandoverPreparationUnsuccessfulTransfer_iE_Extensions_present 0x80
    Cause           cause;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
             * HandoverPreparationUnsuccessfulTransfer_iE_Extensions_present if
             * present */
} HandoverPreparationUnsuccessfulTransfer;

typedef enum IntegrityProtectionResult {
    IntegrityProtectionResult_performed = 0,
    IntegrityProtectionResult_not_performed = 1
} IntegrityProtectionResult;

typedef struct SecurityResult {
    unsigned char   bit_mask;
#       define      SecurityResult_iE_Extensions_present 0x80
    IntegrityProtectionResult integrityProtectionResult;
    ConfidentialityProtectionResult confidentialityProtectionResult;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * SecurityResult_iE_Extensions_present if
                                      * present */
} SecurityResult;

enum HandoverRequestAcknowledgeTransfer_ExtIEs_Extension_PDUs {
    PDU_HandoverRequestAcknowledgeTransfer_ExtIEs_Extension_UNKNOWN = 0,

    PDU_HandoverRequestAcknowledgeTransfer_ExtIEs_Extension_AdditionalDLUPTNLInformationForHOList = AdditionalDLUPTNLInformationForHOList_PDU,
    PDU_HandoverRequestAcknowledgeTransfer_ExtIEs_Extension_UPTransportLayerInformation = UPTransportLayerInformation_PDU,

    PDU_HandoverRequestAcknowledgeTransfer_ExtIEs_Extension_MAX = INT_MAX
};

union HandoverRequestAcknowledgeTransfer_ExtIEs_Extension_union {
    struct AdditionalDLUPTNLInformationForHOList_ **pdu_AdditionalDLUPTNLInformationForHOList;                          /* PDU_HandoverRequestAcknowledgeTransfer_ExtIEs_Extension_AdditionalDLUPTNLInformationForHOList */
    UPTransportLayerInformation *pdu_UPTransportLayerInformation;  
          /* PDU_HandoverRequestAcknowledgeTransfer_ExtIEs_Extension_UPTransportLayerInformation */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct HandoverRequestAcknowledgeTransfer_ExtIEs_Extension {
    HandoverRequestAcknowledgeTransfer_ExtIEs_Extension_PDUs pduNum;
    OssBuf          encoded;
    HandoverRequestAcknowledgeTransfer_ExtIEs_Extension_union decoded;
} HandoverRequestAcknowledgeTransfer_ExtIEs_Extension;

typedef struct HandoverRequestAcknowledgeTransfer {
    unsigned char   bit_mask;
#       define      HandoverRequestAcknowledgeTransfer_dLForwardingUP_TNLInformation_present 0x80
#       define      HandoverRequestAcknowledgeTransfer_securityResult_present 0x40
#       define      HandoverRequestAcknowledgeTransfer_qosFlowFailedToSetupList_present 0x20
#       define      HandoverRequestAcknowledgeTransfer_dataForwardingResponseDRBList_present 0x10
#       define      HandoverRequestAcknowledgeTransfer_iE_Extensions_present 0x08
    UPTransportLayerInformation dL_NGU_UP_TNLInformation;
    UPTransportLayerInformation dLForwardingUP_TNLInformation;  /* optional; set
                                   * in bit_mask
  * HandoverRequestAcknowledgeTransfer_dLForwardingUP_TNLInformation_present if
  * present */
    SecurityResult  securityResult;  /* optional; set in bit_mask
                 * HandoverRequestAcknowledgeTransfer_securityResult_present if
                 * present */
    struct QosFlowListWithDataForwarding_ *qosFlowSetupResponseList;
    struct QosFlowListWithCause_ *qosFlowFailedToSetupList;  /* optional; set in
                                   * bit_mask
       * HandoverRequestAcknowledgeTransfer_qosFlowFailedToSetupList_present if
       * present */
    struct DataForwardingResponseDRBList_ *dataForwardingResponseDRBList;  
                                        /* optional; set in bit_mask
  * HandoverRequestAcknowledgeTransfer_dataForwardingResponseDRBList_present if
  * present */
    struct _seqof70 {
        struct _seqof70 *next;
        struct _seq85 {
            ProtocolExtensionID id;
            Criticality     criticality;
            HandoverRequestAcknowledgeTransfer_ExtIEs_Extension extensionValue;
        } value;
    } *iE_Extensions;  /* optional; set in bit_mask
                  * HandoverRequestAcknowledgeTransfer_iE_Extensions_present if
                  * present */
} HandoverRequestAcknowledgeTransfer;

typedef struct HandoverRequiredTransfer {
    unsigned char   bit_mask;
#       define      directForwardingPathAvailability_present 0x80
#       define      HandoverRequiredTransfer_iE_Extensions_present 0x40
    DirectForwardingPathAvailability directForwardingPathAvailability;  
                                  /* optional; set in bit_mask
                                   * directForwardingPathAvailability_present if
                                   * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                            * HandoverRequiredTransfer_iE_Extensions_present if
                            * present */
} HandoverRequiredTransfer;

typedef struct HandoverResourceAllocationUnsuccessfulTransfer {
    unsigned char   bit_mask;
#       define      HandoverResourceAllocationUnsuccessfulTransfer_criticalityDiagnostics_present 0x80
#       define      HandoverResourceAllocationUnsuccessfulTransfer_iE_Extensions_present 0x40
    Cause           cause;
    CriticalityDiagnostics criticalityDiagnostics;  /* optional; set in bit_mask
     * HandoverResourceAllocationUnsuccessfulTransfer_criticalityDiagnostics_present if
     * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
      * HandoverResourceAllocationUnsuccessfulTransfer_iE_Extensions_present if
      * present */
} HandoverResourceAllocationUnsuccessfulTransfer;

typedef struct RecommendedRANNodesForPaging {
    unsigned char   bit_mask;
#       define      RecommendedRANNodesForPaging_iE_Extensions_present 0x80
    struct RecommendedRANNodeList_ *recommendedRANNodeList;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                        * RecommendedRANNodesForPaging_iE_Extensions_present if
                        * present */
} RecommendedRANNodesForPaging;

typedef struct InfoOnRecommendedCellsAndRANNodesForPaging {
    unsigned char   bit_mask;
#       define      InfoOnRecommendedCellsAndRANNodesForPaging_iE_Extensions_present 0x80
    RecommendedCellsForPaging recommendedCellsForPaging;
    RecommendedRANNodesForPaging recommendRANNodesForPaging;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
          * InfoOnRecommendedCellsAndRANNodesForPaging_iE_Extensions_present if
          * present */
} InfoOnRecommendedCellsAndRANNodesForPaging;

typedef enum IntegrityProtectionIndication {
    IntegrityProtectionIndication_required = 0,
    IntegrityProtectionIndication_preferred = 1,
    IntegrityProtectionIndication_not_needed = 2
} IntegrityProtectionIndication;

typedef struct InterfacesToTrace {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} InterfacesToTrace;

typedef enum ReportArea {
    cell = 0
} ReportArea;

typedef struct LocationReportingRequestType {
    unsigned char   bit_mask;
#       define      areaOfInterestList_present 0x80
#       define      locationReportingReferenceIDToBeCancelled_present 0x40
#       define      LocationReportingRequestType_iE_Extensions_present 0x20
    EventType       eventType;
    ReportArea      reportArea;
    struct AreaOfInterestList_ *areaOfInterestList;  /* optional; set in
                                   * bit_mask areaOfInterestList_present if
                                   * present */
    LocationReportingReferenceID locationReportingReferenceIDToBeCancelled;  
                                        /* optional; set in bit_mask
                         * locationReportingReferenceIDToBeCancelled_present if
                         * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                        * LocationReportingRequestType_iE_Extensions_present if
                        * present */
// The above IE shall be present if the event type is set to “stop reporting UE presence in the area of interest”
} LocationReportingRequestType;

// M
typedef struct MaskedIMEISV {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} MaskedIMEISV;

typedef struct MessageIdentifier {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} MessageIdentifier;

enum MobilityRestrictionList_ExtIEs_Extension_PDUs {
    PDU_MobilityRestrictionList_ExtIEs_Extension_UNKNOWN = 0,

    PDU_MobilityRestrictionList_ExtIEs_Extension_PLMNIdentity = PLMNIdentity_PDU,
    PDU_MobilityRestrictionList_ExtIEs_Extension_CNTypeRestrictionsForServing = CNTypeRestrictionsForServing_PDU,
    PDU_MobilityRestrictionList_ExtIEs_Extension_CNTypeRestrictionsForEquivalent = CNTypeRestrictionsForEquivalent_PDU,

    PDU_MobilityRestrictionList_ExtIEs_Extension_MAX = INT_MAX
};

union MobilityRestrictionList_ExtIEs_Extension_union {
    PLMNIdentity    *pdu_PLMNIdentity;  
                /* PDU_MobilityRestrictionList_ExtIEs_Extension_PLMNIdentity */
    CNTypeRestrictionsForServing *pdu_CNTypeRestrictionsForServing;  
/* PDU_MobilityRestrictionList_ExtIEs_Extension_CNTypeRestrictionsForServing */
    struct CNTypeRestrictionsForEquivalent_ **pdu_CNTypeRestrictionsForEquivalent;                                      /* PDU_MobilityRestrictionList_ExtIEs_Extension_CNTypeRestrictionsForEquivalent */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct MobilityRestrictionList_ExtIEs_Extension {
    MobilityRestrictionList_ExtIEs_Extension_PDUs pduNum;
    OssBuf          encoded;
    MobilityRestrictionList_ExtIEs_Extension_union decoded;
} MobilityRestrictionList_ExtIEs_Extension;

typedef struct MobilityRestrictionList {
    unsigned char   bit_mask;
#       define      equivalentPLMNs_present 0x80
#       define      rATRestrictions_present 0x40
#       define      forbiddenAreaInformation_present 0x20
#       define      serviceAreaInformation_present 0x10
#       define      MobilityRestrictionList_iE_Extensions_present 0x08
    PLMNIdentity    servingPLMN;
    struct EquivalentPLMNs_ *equivalentPLMNs;  /* optional; set in bit_mask
                                                * equivalentPLMNs_present if
                                                * present */
    struct RATRestrictions_ *rATRestrictions;  /* optional; set in bit_mask
                                                * rATRestrictions_present if
                                                * present */
    struct ForbiddenAreaInformation_ *forbiddenAreaInformation;  /* optional;
                                   * set in bit_mask
                                   * forbiddenAreaInformation_present if
                                   * present */
    struct ServiceAreaInformation_ *serviceAreaInformation;  /* optional; set in
                                   * bit_mask serviceAreaInformation_present if
                                   * present */
    struct _seqof76 {
        struct _seqof76 *next;
        struct _seq91 {
            ProtocolExtensionID id;
            Criticality     criticality;
            MobilityRestrictionList_ExtIEs_Extension extensionValue;
        } value;
    } *iE_Extensions;  /* optional; set in bit_mask
                        * MobilityRestrictionList_iE_Extensions_present if
                        * present */
} MobilityRestrictionList;

typedef struct NAS_PDU {
    unsigned int    length;
    unsigned char   *value;
} NAS_PDU;

typedef struct NASSecurityParametersFromNGRAN {
    unsigned int    length;
    unsigned char   *value;
} NASSecurityParametersFromNGRAN;

typedef unsigned short  NextHopChainingCount;

typedef struct NGRANTraceID {
    unsigned short  length;
    unsigned char   value[8];
} NGRANTraceID;

typedef struct NonDynamic5QIDescriptor {
    unsigned char   bit_mask;
#       define      priorityLevelQos_present 0x80
#       define      NonDynamic5QIDescriptor_averagingWindow_present 0x40
#       define      NonDynamic5QIDescriptor_maximumDataBurstVolume_present 0x20
#       define      NonDynamic5QIDescriptor_iE_Extensions_present 0x10
    FiveQI          fiveQI;
    PriorityLevelQos priorityLevelQos;  /* optional; set in bit_mask
                                         * priorityLevelQos_present if
                                         * present */
    AveragingWindow averagingWindow;  /* optional; set in bit_mask
                           * NonDynamic5QIDescriptor_averagingWindow_present if
                           * present */
    MaximumDataBurstVolume maximumDataBurstVolume;  /* optional; set in bit_mask
                    * NonDynamic5QIDescriptor_maximumDataBurstVolume_present if
                    * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                             * NonDynamic5QIDescriptor_iE_Extensions_present if
                             * present */
} NonDynamic5QIDescriptor;

typedef struct NotAllowedTACs_ {
    struct NotAllowedTACs_ *next;
    TAC             value;
} *NotAllowedTACs;

typedef enum NotificationCause {
    fulfilled = 0,
    not_fulfilled = 1
} NotificationCause;

typedef struct NR_CGIList_ {
    struct NR_CGIList_ *next;
    NR_CGI          value;
} *NR_CGIList;

typedef struct NR_CGIListForWarning_ {
    struct NR_CGIListForWarning_ *next;
    NR_CGI          value;
} *NR_CGIListForWarning;

typedef struct NRencryptionAlgorithms {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} NRencryptionAlgorithms;

typedef struct NRintegrityProtectionAlgorithms {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} NRintegrityProtectionAlgorithms;

typedef struct NRPPa_PDU {
    unsigned int    length;
    unsigned char   *value;
} NRPPa_PDU;

// O
typedef enum OverloadAction {
    reject_non_emergency_mo_dt = 0,
    reject_rrc_cr_signalling = 1,
    permit_emergency_sessions_and_mobile_terminated_services_only = 2,
    permit_high_priority_sessions_and_mobile_terminated_services_only = 3
} OverloadAction;

typedef struct OverloadResponse {
    unsigned short  choice;
#       define      overloadAction_chosen 1
#       define      OverloadResponse_choice_Extensions_chosen 2
    union _union {
        OverloadAction  overloadAction;  /* to choose, set choice to
                                          * overloadAction_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to
                                 * OverloadResponse_choice_Extensions_chosen */
    } u;
} OverloadResponse;

typedef struct OverloadStartNSSAIItem {
    unsigned char   bit_mask;
#       define      sliceOverloadResponse_present 0x80
#       define      sliceTrafficLoadReductionIndication_present 0x40
#       define      OverloadStartNSSAIItem_iE_Extensions_present 0x20
    struct SliceOverloadList_ *sliceOverloadList;
    OverloadResponse sliceOverloadResponse;  /* optional; set in bit_mask
                                              * sliceOverloadResponse_present if
                                              * present */
    TrafficLoadReductionIndication sliceTrafficLoadReductionIndication;  
                                  /* optional; set in bit_mask
                                   * sliceTrafficLoadReductionIndication_present
                                   * if present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                              * OverloadStartNSSAIItem_iE_Extensions_present if
                              * present */
} OverloadStartNSSAIItem;

typedef struct OverloadStartNSSAIList_ {
    struct OverloadStartNSSAIList_ *next;
    OverloadStartNSSAIItem value;
} *OverloadStartNSSAIList;

enum SecurityIndication_ExtIEs_Extension_PDUs {
    PDU_SecurityIndication_ExtIEs_Extension_UNKNOWN = 0,

    PDU_SecurityIndication_ExtIEs_Extension_MaximumIntegrityProtectedDataRate = MaximumIntegrityProtectedDataRate_PDU,

    PDU_SecurityIndication_ExtIEs_Extension_MAX = INT_MAX
};

union SecurityIndication_ExtIEs_Extension_union {
    MaximumIntegrityProtectedDataRate *pdu_MaximumIntegrityProtectedDataRate;                                           /* PDU_SecurityIndication_ExtIEs_Extension_MaximumIntegrityProtectedDataRate */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct SecurityIndication_ExtIEs_Extension {
    SecurityIndication_ExtIEs_Extension_PDUs pduNum;
    OssBuf          encoded;
    SecurityIndication_ExtIEs_Extension_union decoded;
} SecurityIndication_ExtIEs_Extension;

typedef struct SecurityIndication {
    unsigned char   bit_mask;
#       define      maximumIntegrityProtectedDataRate_UL_present 0x80
#       define      SecurityIndication_iE_Extensions_present 0x40
    IntegrityProtectionIndication integrityProtectionIndication;
    ConfidentialityProtectionIndication confidentialityProtectionIndication;
    MaximumIntegrityProtectedDataRate maximumIntegrityProtectedDataRate_UL;  
                                        /* optional; set in bit_mask
                              * maximumIntegrityProtectedDataRate_UL_present if
                              * present */
// The above IE shall be present if integrity protection is required or preferred
    struct _seqof79 {
        struct _seqof79 *next;
        struct _seq95 {
            ProtocolExtensionID id;
            Criticality     criticality;
            SecurityIndication_ExtIEs_Extension extensionValue;
        } value;
    } *iE_Extensions;  /* optional; set in bit_mask
                        * SecurityIndication_iE_Extensions_present if present */
} SecurityIndication;

enum PathSwitchRequestAcknowledgeTransfer_ExtIEs_Extension_PDUs {
    PDU_PathSwitchRequestAcknowledgeTransfer_ExtIEs_Extension_UNKNOWN = 0,

    PDU_PathSwitchRequestAcknowledgeTransfer_ExtIEs_Extension_UPTransportLayerInformationPairList = UPTransportLayerInformationPairList_PDU,

    PDU_PathSwitchRequestAcknowledgeTransfer_ExtIEs_Extension_MAX = INT_MAX
};

union PathSwitchRequestAcknowledgeTransfer_ExtIEs_Extension_union {
    struct UPTransportLayerInformationPairList_ **pdu_UPTransportLayerInformationPairList;                              /* PDU_PathSwitchRequestAcknowledgeTransfer_ExtIEs_Extension_UPTransportLayerInformationPairList */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PathSwitchRequestAcknowledgeTransfer_ExtIEs_Extension {
    PathSwitchRequestAcknowledgeTransfer_ExtIEs_Extension_PDUs pduNum;
    OssBuf          encoded;
    PathSwitchRequestAcknowledgeTransfer_ExtIEs_Extension_union decoded;
} PathSwitchRequestAcknowledgeTransfer_ExtIEs_Extension;

typedef struct _seq96 {
    ProtocolExtensionID id;
    Criticality     criticality;
    PathSwitchRequestAcknowledgeTransfer_ExtIEs_Extension extensionValue;
} _seq96;

typedef struct _seqof80_ {
    struct _seqof80_ *next;
    _seq96          value;
} *_seqof80;

typedef struct PathSwitchRequestAcknowledgeTransfer {
    unsigned char   bit_mask;
#       define      PathSwitchRequestAcknowledgeTransfer_uL_NGU_UP_TNLInformation_present 0x80
#       define      securityIndication_present 0x40
#       define      PathSwitchRequestAcknowledgeTransfer_iE_Extensions_present 0x20
    UPTransportLayerInformation uL_NGU_UP_TNLInformation;  /* optional; set in
                                   * bit_mask
     * PathSwitchRequestAcknowledgeTransfer_uL_NGU_UP_TNLInformation_present if
     * present */
    SecurityIndication securityIndication;  /* optional; set in bit_mask
                                             * securityIndication_present if
                                             * present */
    struct _seqof80_ *iE_Extensions;  /* optional; set in bit_mask
                * PathSwitchRequestAcknowledgeTransfer_iE_Extensions_present if
                * present */
} PathSwitchRequestAcknowledgeTransfer;

typedef struct PathSwitchRequestSetupFailedTransfer {
    unsigned char   bit_mask;
#       define      PathSwitchRequestSetupFailedTransfer_iE_Extensions_present 0x80
    Cause           cause;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                * PathSwitchRequestSetupFailedTransfer_iE_Extensions_present if
                * present */
} PathSwitchRequestSetupFailedTransfer;

typedef struct UserPlaneSecurityInformation {
    unsigned char   bit_mask;
#       define      UserPlaneSecurityInformation_iE_Extensions_present 0x80
    SecurityResult  securityResult;
    SecurityIndication securityIndication;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                        * UserPlaneSecurityInformation_iE_Extensions_present if
                        * present */
} UserPlaneSecurityInformation;

enum PathSwitchRequestTransfer_ExtIEs_Extension_PDUs {
    PDU_PathSwitchRequestTransfer_ExtIEs_Extension_UNKNOWN = 0,

    PDU_PathSwitchRequestTransfer_ExtIEs_Extension_QosFlowPerTNLInformationList = QosFlowPerTNLInformationList_PDU,

    PDU_PathSwitchRequestTransfer_ExtIEs_Extension_MAX = INT_MAX
};

union PathSwitchRequestTransfer_ExtIEs_Extension_union {
    struct QosFlowPerTNLInformationList_ **pdu_QosFlowPerTNLInformationList;  
                                        /* PDU_PathSwitchRequestTransfer_ExtIEs_Extension_QosFlowPerTNLInformationList */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PathSwitchRequestTransfer_ExtIEs_Extension {
    PathSwitchRequestTransfer_ExtIEs_Extension_PDUs pduNum;
    OssBuf          encoded;
    PathSwitchRequestTransfer_ExtIEs_Extension_union decoded;
} PathSwitchRequestTransfer_ExtIEs_Extension;

typedef struct PathSwitchRequestTransfer {
    unsigned char   bit_mask;
#       define      dL_NGU_TNLInformationReused_present 0x80
#       define      userPlaneSecurityInformation_present 0x40
#       define      PathSwitchRequestTransfer_iE_Extensions_present 0x20
    UPTransportLayerInformation dL_NGU_UP_TNLInformation;
    DL_NGU_TNLInformationReused dL_NGU_TNLInformationReused;  /* optional; set
                                   * in bit_mask
                                   * dL_NGU_TNLInformationReused_present if
                                   * present */
    UserPlaneSecurityInformation userPlaneSecurityInformation;  /* optional; set
                                   * in bit_mask
                                   * userPlaneSecurityInformation_present if
                                   * present */
    struct QosFlowAcceptedList_ *qosFlowAcceptedList;
    struct _seqof83 {
        struct _seqof83 *next;
        struct _seq99 {
            ProtocolExtensionID id;
            Criticality     criticality;
            PathSwitchRequestTransfer_ExtIEs_Extension extensionValue;
        } value;
    } *iE_Extensions;  /* optional; set in bit_mask
                        * PathSwitchRequestTransfer_iE_Extensions_present if
                        * present */
} PathSwitchRequestTransfer;

typedef struct PathSwitchRequestUnsuccessfulTransfer {
    unsigned char   bit_mask;
#       define      PathSwitchRequestUnsuccessfulTransfer_iE_Extensions_present 0x80
    Cause           cause;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
               * PathSwitchRequestUnsuccessfulTransfer_iE_Extensions_present if
               * present */
} PathSwitchRequestUnsuccessfulTransfer;

typedef struct PDUSessionAggregateMaximumBitRate {
    unsigned char   bit_mask;
#       define      PDUSessionAggregateMaximumBitRate_iE_Extensions_present 0x80
    BitRate         pDUSessionAggregateMaximumBitRateDL;
    BitRate         pDUSessionAggregateMaximumBitRateUL;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                   * PDUSessionAggregateMaximumBitRate_iE_Extensions_present if
                   * present */
} PDUSessionAggregateMaximumBitRate;

typedef unsigned short  PDUSessionID;

typedef struct PDUSessionResourceAdmittedItem {
    unsigned char   bit_mask;
#       define      PDUSessionResourceAdmittedItem_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq103 {
        /* ContentsConstraint is applied to
         * handoverRequestAcknowledgeTransfer */
        _octet1         encoded;
        HandoverRequestAcknowledgeTransfer *decoded;
    } handoverRequestAcknowledgeTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                      * PDUSessionResourceAdmittedItem_iE_Extensions_present if
                      * present */
} PDUSessionResourceAdmittedItem;

typedef struct PDUSessionResourceAdmittedList_ {
    struct PDUSessionResourceAdmittedList_ *next;
    PDUSessionResourceAdmittedItem value;
} *PDUSessionResourceAdmittedList;

typedef struct PDUSessionResourceFailedToModifyItemModCfm {
    unsigned char   bit_mask;
#       define      PDUSessionResourceFailedToModifyItemModCfm_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq105 {
        /* ContentsConstraint is applied to
         * pDUSessionResourceModifyIndicationUnsuccessfulTransfer */
        _octet1         encoded;
        struct PDUSessionResourceModifyIndicationUnsuccessfulTransfer *decoded;
    } pDUSessionResourceModifyIndicationUnsuccessfulTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
          * PDUSessionResourceFailedToModifyItemModCfm_iE_Extensions_present if
          * present */
} PDUSessionResourceFailedToModifyItemModCfm;

typedef struct PDUSessionResourceFailedToModifyListModCfm_ {
    struct PDUSessionResourceFailedToModifyListModCfm_ *next;
    PDUSessionResourceFailedToModifyItemModCfm value;
} *PDUSessionResourceFailedToModifyListModCfm;

typedef struct PDUSessionResourceFailedToModifyItemModRes {
    unsigned char   bit_mask;
#       define      PDUSessionResourceFailedToModifyItemModRes_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq107 {
        /* ContentsConstraint is applied to
         * pDUSessionResourceModifyUnsuccessfulTransfer */
        _octet1         encoded;
        struct PDUSessionResourceModifyUnsuccessfulTransfer *decoded;
    } pDUSessionResourceModifyUnsuccessfulTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
          * PDUSessionResourceFailedToModifyItemModRes_iE_Extensions_present if
          * present */
} PDUSessionResourceFailedToModifyItemModRes;

typedef struct PDUSessionResourceFailedToModifyListModRes_ {
    struct PDUSessionResourceFailedToModifyListModRes_ *next;
    PDUSessionResourceFailedToModifyItemModRes value;
} *PDUSessionResourceFailedToModifyListModRes;

typedef struct _seq109 {
    /* ContentsConstraint is applied to _seq109 */
    _octet1         encoded;
    struct PDUSessionResourceSetupUnsuccessfulTransfer *decoded;
} _seq109;

typedef struct PDUSessionResourceFailedToSetupItemCxtFail {
    unsigned char   bit_mask;
#       define      PDUSessionResourceFailedToSetupItemCxtFail_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    _seq109         pDUSessionResourceSetupUnsuccessfulTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
          * PDUSessionResourceFailedToSetupItemCxtFail_iE_Extensions_present if
          * present */
} PDUSessionResourceFailedToSetupItemCxtFail;

typedef struct PDUSessionResourceFailedToSetupListCxtFail_ {
    struct PDUSessionResourceFailedToSetupListCxtFail_ *next;
    PDUSessionResourceFailedToSetupItemCxtFail value;
} *PDUSessionResourceFailedToSetupListCxtFail;

typedef struct PDUSessionResourceFailedToSetupItemCxtRes {
    unsigned char   bit_mask;
#       define      PDUSessionResourceFailedToSetupItemCxtRes_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    _seq109         pDUSessionResourceSetupUnsuccessfulTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
           * PDUSessionResourceFailedToSetupItemCxtRes_iE_Extensions_present if
           * present */
} PDUSessionResourceFailedToSetupItemCxtRes;

typedef struct PDUSessionResourceFailedToSetupListCxtRes_ {
    struct PDUSessionResourceFailedToSetupListCxtRes_ *next;
    PDUSessionResourceFailedToSetupItemCxtRes value;
} *PDUSessionResourceFailedToSetupListCxtRes;

typedef struct PDUSessionResourceFailedToSetupItemHOAck {
    unsigned char   bit_mask;
#       define      PDUSessionResourceFailedToSetupItemHOAck_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq112 {
        /* ContentsConstraint is applied to
         * handoverResourceAllocationUnsuccessfulTransfer */
        _octet1         encoded;
        HandoverResourceAllocationUnsuccessfulTransfer *decoded;
    } handoverResourceAllocationUnsuccessfulTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
            * PDUSessionResourceFailedToSetupItemHOAck_iE_Extensions_present if
            * present */
} PDUSessionResourceFailedToSetupItemHOAck;

typedef struct PDUSessionResourceFailedToSetupListHOAck_ {
    struct PDUSessionResourceFailedToSetupListHOAck_ *next;
    PDUSessionResourceFailedToSetupItemHOAck value;
} *PDUSessionResourceFailedToSetupListHOAck;

typedef struct PDUSessionResourceFailedToSetupItemPSReq {
    unsigned char   bit_mask;
#       define      PDUSessionResourceFailedToSetupItemPSReq_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq114 {
        /* ContentsConstraint is applied to
         * pathSwitchRequestSetupFailedTransfer */
        _octet1         encoded;
        PathSwitchRequestSetupFailedTransfer *decoded;
    } pathSwitchRequestSetupFailedTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
            * PDUSessionResourceFailedToSetupItemPSReq_iE_Extensions_present if
            * present */
} PDUSessionResourceFailedToSetupItemPSReq;

typedef struct PDUSessionResourceFailedToSetupListPSReq_ {
    struct PDUSessionResourceFailedToSetupListPSReq_ *next;
    PDUSessionResourceFailedToSetupItemPSReq value;
} *PDUSessionResourceFailedToSetupListPSReq;

typedef struct PDUSessionResourceFailedToSetupItemSURes {
    unsigned char   bit_mask;
#       define      PDUSessionResourceFailedToSetupItemSURes_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    _seq109         pDUSessionResourceSetupUnsuccessfulTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
            * PDUSessionResourceFailedToSetupItemSURes_iE_Extensions_present if
            * present */
} PDUSessionResourceFailedToSetupItemSURes;

typedef struct PDUSessionResourceFailedToSetupListSURes_ {
    struct PDUSessionResourceFailedToSetupListSURes_ *next;
    PDUSessionResourceFailedToSetupItemSURes value;
} *PDUSessionResourceFailedToSetupListSURes;

typedef struct PDUSessionResourceHandoverItem {
    unsigned char   bit_mask;
#       define      PDUSessionResourceHandoverItem_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq117 {
        /* ContentsConstraint is applied to handoverCommandTransfer */
        _octet1         encoded;
        HandoverCommandTransfer *decoded;
    } handoverCommandTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                      * PDUSessionResourceHandoverItem_iE_Extensions_present if
                      * present */
} PDUSessionResourceHandoverItem;

typedef struct PDUSessionResourceHandoverList_ {
    struct PDUSessionResourceHandoverList_ *next;
    PDUSessionResourceHandoverItem value;
} *PDUSessionResourceHandoverList;

enum PDUSessionResourceItemCxtRelCpl_ExtIEs_Extension_PDUs {
    PDU_PDUSessionResourceItemCxtRelCpl_ExtIEs_Extension_UNKNOWN = 0,

    PDU_PDUSessionResourceItemCxtRelCpl_ExtIEs_Extension_EXTENSION = PDUSessionResourceItemCxtRelCpl_ExtIEs_EXTENSION_PDU,

    PDU_PDUSessionResourceItemCxtRelCpl_ExtIEs_Extension_MAX = INT_MAX
};

union PDUSessionResourceItemCxtRelCpl_ExtIEs_Extension_union {
    PDUSessionResourceItemCxtRelCpl_ExtIEs_EXTENSION *pdu_PDUSessionResourceItemCxtRelCpl_ExtIEs_EXTENSION;             /* PDU_PDUSessionResourceItemCxtRelCpl_ExtIEs_Extension_EXTENSION */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PDUSessionResourceItemCxtRelCpl_ExtIEs_Extension {
    PDUSessionResourceItemCxtRelCpl_ExtIEs_Extension_PDUs pduNum;
    OssBuf          encoded;
    PDUSessionResourceItemCxtRelCpl_ExtIEs_Extension_union decoded;
} PDUSessionResourceItemCxtRelCpl_ExtIEs_Extension;

typedef struct PDUSessionResourceItemCxtRelCpl {
    unsigned char   bit_mask;
#       define      PDUSessionResourceItemCxtRelCpl_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seqof95 {
        struct _seqof95 *next;
        struct _seq118 {
            ProtocolExtensionID id;
            Criticality     criticality;
            PDUSessionResourceItemCxtRelCpl_ExtIEs_Extension extensionValue;
        } value;
    } *iE_Extensions;  /* optional; set in bit_mask
                        * PDUSessionResourceItemCxtRelCpl_iE_Extensions_present
                        * if present */
} PDUSessionResourceItemCxtRelCpl;

typedef struct PDUSessionResourceListCxtRelCpl_ {
    struct PDUSessionResourceListCxtRelCpl_ *next;
    PDUSessionResourceItemCxtRelCpl value;
} *PDUSessionResourceListCxtRelCpl;

typedef struct PDUSessionResourceItemCxtRelReq {
    unsigned char   bit_mask;
#       define      PDUSessionResourceItemCxtRelReq_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                     * PDUSessionResourceItemCxtRelReq_iE_Extensions_present if
                     * present */
} PDUSessionResourceItemCxtRelReq;

typedef struct PDUSessionResourceListCxtRelReq_ {
    struct PDUSessionResourceListCxtRelReq_ *next;
    PDUSessionResourceItemCxtRelReq value;
} *PDUSessionResourceListCxtRelReq;

typedef struct PDUSessionResourceItemHORqd {
    unsigned char   bit_mask;
#       define      PDUSessionResourceItemHORqd_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq121 {
        /* ContentsConstraint is applied to handoverRequiredTransfer */
        _octet1         encoded;
        HandoverRequiredTransfer *decoded;
    } handoverRequiredTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                         * PDUSessionResourceItemHORqd_iE_Extensions_present if
                         * present */
} PDUSessionResourceItemHORqd;

typedef struct PDUSessionResourceListHORqd_ {
    struct PDUSessionResourceListHORqd_ *next;
    PDUSessionResourceItemHORqd value;
} *PDUSessionResourceListHORqd;

typedef struct PDUSessionResourceModifyConfirmTransfer {
    unsigned char   bit_mask;
#       define      additionalNG_UUPTNLInformation_present 0x80
#       define      qosFlowFailedToModifyList_present 0x40
#       define      PDUSessionResourceModifyConfirmTransfer_iE_Extensions_present 0x20
    struct QosFlowModifyConfirmList_ *qosFlowModifyConfirmList;
    UPTransportLayerInformation uLNGU_UP_TNLInformation;
    struct UPTransportLayerInformationPairList_ *additionalNG_UUPTNLInformation;                                        /* optional; set in bit_mask
                                    * additionalNG_UUPTNLInformation_present if
                                    * present */
    struct QosFlowListWithCause_ *qosFlowFailedToModifyList;  /* optional; set
                                   * in bit_mask
                                   * qosFlowFailedToModifyList_present if
                                   * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
             * PDUSessionResourceModifyConfirmTransfer_iE_Extensions_present if
             * present */
} PDUSessionResourceModifyConfirmTransfer;

typedef struct PDUSessionResourceModifyIndicationUnsuccessfulTransfer {
    unsigned char   bit_mask;
#       define      PDUSessionResourceModifyIndicationUnsuccessfulTransfer_iE_Extensions_present 0x80
    Cause           cause;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
    * PDUSessionResourceModifyIndicationUnsuccessfulTransfer_iE_Extensions_present if
    * present */
} PDUSessionResourceModifyIndicationUnsuccessfulTransfer;

enum PDUSessionResourceModifyRequestTransferIEs_Value_PDUs {
    PDU_PDUSessionResourceModifyRequestTransferIEs_Value_UNKNOWN = 0,

    PDU_PDUSessionResourceModifyRequestTransferIEs_Value_PDUSessionAggregateMaximumBitRate = PDUSessionAggregateMaximumBitRate_PDU,
    PDU_PDUSessionResourceModifyRequestTransferIEs_Value_UL_NGU_UP_TNLModifyList = UL_NGU_UP_TNLModifyList_PDU,
    PDU_PDUSessionResourceModifyRequestTransferIEs_Value_NetworkInstance = NetworkInstance_PDU,
    PDU_PDUSessionResourceModifyRequestTransferIEs_Value_QosFlowAddOrModifyRequestList = QosFlowAddOrModifyRequestList_PDU,
    PDU_PDUSessionResourceModifyRequestTransferIEs_Value_QosFlowListWithCause = QosFlowListWithCause_PDU,
    PDU_PDUSessionResourceModifyRequestTransferIEs_Value_UPTransportLayerInformationList = UPTransportLayerInformationList_PDU,
    PDU_PDUSessionResourceModifyRequestTransferIEs_Value_CommonNetworkInstance = CommonNetworkInstance_PDU,

    PDU_PDUSessionResourceModifyRequestTransferIEs_Value_MAX = INT_MAX
};

union PDUSessionResourceModifyRequestTransferIEs_Value_union {
    PDUSessionAggregateMaximumBitRate *pdu_PDUSessionAggregateMaximumBitRate;                                           /* PDU_PDUSessionResourceModifyRequestTransferIEs_Value_PDUSessionAggregateMaximumBitRate */
    struct UL_NGU_UP_TNLModifyList_ **pdu_UL_NGU_UP_TNLModifyList;  
   /* PDU_PDUSessionResourceModifyRequestTransferIEs_Value_UL_NGU_UP_TNLModifyList */
    NetworkInstance *pdu_NetworkInstance;  
     /* PDU_PDUSessionResourceModifyRequestTransferIEs_Value_NetworkInstance */
    struct QosFlowAddOrModifyRequestList_ **pdu_QosFlowAddOrModifyRequestList;                                          /* PDU_PDUSessionResourceModifyRequestTransferIEs_Value_QosFlowAddOrModifyRequestList */
    struct QosFlowListWithCause_ **pdu_QosFlowListWithCause;  
/* PDU_PDUSessionResourceModifyRequestTransferIEs_Value_QosFlowListWithCause */
    struct UPTransportLayerInformationList_ **pdu_UPTransportLayerInformationList;                                      /* PDU_PDUSessionResourceModifyRequestTransferIEs_Value_UPTransportLayerInformationList */
    CommonNetworkInstance *pdu_CommonNetworkInstance;  
 /* PDU_PDUSessionResourceModifyRequestTransferIEs_Value_CommonNetworkInstance */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PDUSessionResourceModifyRequestTransferIEs_Value {
    PDUSessionResourceModifyRequestTransferIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PDUSessionResourceModifyRequestTransferIEs_Value_union decoded;
} PDUSessionResourceModifyRequestTransferIEs_Value;

typedef struct PDUSessionResourceModifyRequestTransfer {
    struct _seqof100 {
        struct _seqof100 *next;
        struct _seq124 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PDUSessionResourceModifyRequestTransferIEs_Value value;
        } value;
    } *protocolIEs;
} PDUSessionResourceModifyRequestTransfer;

typedef struct PDUSessionResourceModifyResponseTransfer {
    unsigned char   bit_mask;
#       define      dL_NGU_UP_TNLInformation_present 0x80
#       define      PDUSessionResourceModifyResponseTransfer_uL_NGU_UP_TNLInformation_present 0x40
#       define      qosFlowAddOrModifyResponseList_present 0x20
#       define      PDUSessionResourceModifyResponseTransfer_additionalDLQosFlowPerTNLInformation_present 0x10
#       define      qosFlowFailedToAddOrModifyList_present 0x08
#       define      PDUSessionResourceModifyResponseTransfer_iE_Extensions_present 0x04
    UPTransportLayerInformation dL_NGU_UP_TNLInformation;  /* optional; set in
                                   * bit_mask dL_NGU_UP_TNLInformation_present
                                   * if present */
    UPTransportLayerInformation uL_NGU_UP_TNLInformation;  /* optional; set in
                                   * bit_mask
 * PDUSessionResourceModifyResponseTransfer_uL_NGU_UP_TNLInformation_present if
 * present */
    struct QosFlowAddOrModifyResponseList_ *qosFlowAddOrModifyResponseList;  
                                        /* optional; set in bit_mask
                                    * qosFlowAddOrModifyResponseList_present if
                                    * present */
    struct QosFlowPerTNLInformationList_ *additionalDLQosFlowPerTNLInformation;                                         /* optional; set in bit_mask
             * PDUSessionResourceModifyResponseTransfer_additionalDLQosFlowPerTNLInformation_present if
             * present */
    struct QosFlowListWithCause_ *qosFlowFailedToAddOrModifyList;  /* optional;
                                   * set in bit_mask
                                   * qosFlowFailedToAddOrModifyList_present if
                                   * present */
    struct _seqof80_ *iE_Extensions;  /* optional; set in bit_mask
            * PDUSessionResourceModifyResponseTransfer_iE_Extensions_present if
            * present */
} PDUSessionResourceModifyResponseTransfer;

typedef struct QosFlowPerTNLInformation {
    unsigned char   bit_mask;
#       define      QosFlowPerTNLInformation_iE_Extensions_present 0x80
    UPTransportLayerInformation uPTransportLayerInformation;
    struct AssociatedQosFlowList_ *associatedQosFlowList;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                            * QosFlowPerTNLInformation_iE_Extensions_present if
                            * present */
} QosFlowPerTNLInformation;

enum PDUSessionResourceModifyIndicationTransfer_ExtIEs_Extension_PDUs {
    PDU_PDUSessionResourceModifyIndicationTransfer_ExtIEs_Extension_UNKNOWN = 0,

    PDU_PDUSessionResourceModifyIndicationTransfer_ExtIEs_Extension_SecondaryRATUsageInformation = SecondaryRATUsageInformation_PDU,
    PDU_PDUSessionResourceModifyIndicationTransfer_ExtIEs_Extension_SecurityResult = SecurityResult_PDU,

    PDU_PDUSessionResourceModifyIndicationTransfer_ExtIEs_Extension_MAX = INT_MAX
};

union PDUSessionResourceModifyIndicationTransfer_ExtIEs_Extension_union {
    struct SecondaryRATUsageInformation *pdu_SecondaryRATUsageInformation;  
                                        /* PDU_PDUSessionResourceModifyIndicationTransfer_ExtIEs_Extension_SecondaryRATUsageInformation */
    SecurityResult  *pdu_SecurityResult;  
     /* PDU_PDUSessionResourceModifyIndicationTransfer_ExtIEs_Extension_SecurityResult */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PDUSessionResourceModifyIndicationTransfer_ExtIEs_Extension {
    PDUSessionResourceModifyIndicationTransfer_ExtIEs_Extension_PDUs pduNum;
    OssBuf          encoded;
    PDUSessionResourceModifyIndicationTransfer_ExtIEs_Extension_union decoded;
} PDUSessionResourceModifyIndicationTransfer_ExtIEs_Extension;

typedef struct PDUSessionResourceModifyIndicationTransfer {
    unsigned char   bit_mask;
#       define      PDUSessionResourceModifyIndicationTransfer_additionalDLQosFlowPerTNLInformation_present 0x80
#       define      PDUSessionResourceModifyIndicationTransfer_iE_Extensions_present 0x40
    QosFlowPerTNLInformation dLQosFlowPerTNLInformation;
    struct QosFlowPerTNLInformationList_ *additionalDLQosFlowPerTNLInformation;                                         /* optional; set in bit_mask
               * PDUSessionResourceModifyIndicationTransfer_additionalDLQosFlowPerTNLInformation_present if
               * present */
    struct _seqof103 {
        struct _seqof103 *next;
        struct _seq127 {
            ProtocolExtensionID id;
            Criticality     criticality;
            PDUSessionResourceModifyIndicationTransfer_ExtIEs_Extension extensionValue;
        } value;
    } *iE_Extensions;  /* optional; set in bit_mask
          * PDUSessionResourceModifyIndicationTransfer_iE_Extensions_present if
          * present */
} PDUSessionResourceModifyIndicationTransfer;

typedef struct PDUSessionResourceModifyItemModCfm {
    unsigned char   bit_mask;
#       define      PDUSessionResourceModifyItemModCfm_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq129 {
        /* ContentsConstraint is applied to
         * pDUSessionResourceModifyConfirmTransfer */
        _octet1         encoded;
        PDUSessionResourceModifyConfirmTransfer *decoded;
    } pDUSessionResourceModifyConfirmTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                  * PDUSessionResourceModifyItemModCfm_iE_Extensions_present if
                  * present */
} PDUSessionResourceModifyItemModCfm;

typedef struct PDUSessionResourceModifyListModCfm_ {
    struct PDUSessionResourceModifyListModCfm_ *next;
    PDUSessionResourceModifyItemModCfm value;
} *PDUSessionResourceModifyListModCfm;

typedef struct PDUSessionResourceModifyItemModInd {
    unsigned char   bit_mask;
#       define      PDUSessionResourceModifyItemModInd_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq131 {
        /* ContentsConstraint is applied to
         * pDUSessionResourceModifyIndicationTransfer */
        _octet1         encoded;
        PDUSessionResourceModifyIndicationTransfer *decoded;
    } pDUSessionResourceModifyIndicationTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                  * PDUSessionResourceModifyItemModInd_iE_Extensions_present if
                  * present */
} PDUSessionResourceModifyItemModInd;

typedef struct PDUSessionResourceModifyListModInd_ {
    struct PDUSessionResourceModifyListModInd_ *next;
    PDUSessionResourceModifyItemModInd value;
} *PDUSessionResourceModifyListModInd;

enum PDUSessionResourceModifyItemModReq_ExtIEs_Extension_PDUs {
    PDU_PDUSessionResourceModifyItemModReq_ExtIEs_Extension_UNKNOWN = 0,

    PDU_PDUSessionResourceModifyItemModReq_ExtIEs_Extension_S_NSSAI = S_NSSAI_PDU,

    PDU_PDUSessionResourceModifyItemModReq_ExtIEs_Extension_MAX = INT_MAX
};

union PDUSessionResourceModifyItemModReq_ExtIEs_Extension_union {
    S_NSSAI         *pdu_S_NSSAI;  
          /* PDU_PDUSessionResourceModifyItemModReq_ExtIEs_Extension_S_NSSAI */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PDUSessionResourceModifyItemModReq_ExtIEs_Extension {
    PDUSessionResourceModifyItemModReq_ExtIEs_Extension_PDUs pduNum;
    OssBuf          encoded;
    PDUSessionResourceModifyItemModReq_ExtIEs_Extension_union decoded;
} PDUSessionResourceModifyItemModReq_ExtIEs_Extension;

typedef struct PDUSessionResourceModifyItemModReq {
    unsigned char   bit_mask;
#       define      PDUSessionResourceModifyItemModReq_nAS_PDU_present 0x80
#       define      PDUSessionResourceModifyItemModReq_iE_Extensions_present 0x40
    PDUSessionID    pDUSessionID;
    NAS_PDU         nAS_PDU;  /* optional; set in bit_mask
                        * PDUSessionResourceModifyItemModReq_nAS_PDU_present if
                        * present */
    struct _seq133 {
        /* ContentsConstraint is applied to
         * pDUSessionResourceModifyRequestTransfer */
        _octet1         encoded;
        PDUSessionResourceModifyRequestTransfer *decoded;
    } pDUSessionResourceModifyRequestTransfer;
    struct _seqof106 {
        struct _seqof106 *next;
        struct _seq132 {
            ProtocolExtensionID id;
            Criticality     criticality;
            PDUSessionResourceModifyItemModReq_ExtIEs_Extension extensionValue;
        } value;
    } *iE_Extensions;  /* optional; set in bit_mask
                  * PDUSessionResourceModifyItemModReq_iE_Extensions_present if
                  * present */
} PDUSessionResourceModifyItemModReq;

typedef struct PDUSessionResourceModifyListModReq_ {
    struct PDUSessionResourceModifyListModReq_ *next;
    PDUSessionResourceModifyItemModReq value;
} *PDUSessionResourceModifyListModReq;

typedef struct PDUSessionResourceModifyItemModRes {
    unsigned char   bit_mask;
#       define      PDUSessionResourceModifyItemModRes_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq135 {
        /* ContentsConstraint is applied to
         * pDUSessionResourceModifyResponseTransfer */
        _octet1         encoded;
        PDUSessionResourceModifyResponseTransfer *decoded;
    } pDUSessionResourceModifyResponseTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                  * PDUSessionResourceModifyItemModRes_iE_Extensions_present if
                  * present */
} PDUSessionResourceModifyItemModRes;

typedef struct PDUSessionResourceModifyListModRes_ {
    struct PDUSessionResourceModifyListModRes_ *next;
    PDUSessionResourceModifyItemModRes value;
} *PDUSessionResourceModifyListModRes;

typedef struct PDUSessionResourceModifyUnsuccessfulTransfer {
    unsigned char   bit_mask;
#       define      PDUSessionResourceModifyUnsuccessfulTransfer_criticalityDiagnostics_present 0x80
#       define      PDUSessionResourceModifyUnsuccessfulTransfer_iE_Extensions_present 0x40
    Cause           cause;
    CriticalityDiagnostics criticalityDiagnostics;  /* optional; set in bit_mask
   * PDUSessionResourceModifyUnsuccessfulTransfer_criticalityDiagnostics_present if
   * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
        * PDUSessionResourceModifyUnsuccessfulTransfer_iE_Extensions_present if
        * present */
} PDUSessionResourceModifyUnsuccessfulTransfer;

typedef struct PDUSessionResourceNotifyItem {
    unsigned char   bit_mask;
#       define      PDUSessionResourceNotifyItem_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq138 {
        /* ContentsConstraint is applied to pDUSessionResourceNotifyTransfer */
        _octet1         encoded;
        struct PDUSessionResourceNotifyTransfer *decoded;
    } pDUSessionResourceNotifyTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                        * PDUSessionResourceNotifyItem_iE_Extensions_present if
                        * present */
} PDUSessionResourceNotifyItem;

typedef struct PDUSessionResourceNotifyList_ {
    struct PDUSessionResourceNotifyList_ *next;
    PDUSessionResourceNotifyItem value;
} *PDUSessionResourceNotifyList;

enum PDUSessionResourceNotifyTransfer_ExtIEs_Extension_PDUs {
    PDU_PDUSessionResourceNotifyTransfer_ExtIEs_Extension_UNKNOWN = 0,

    PDU_PDUSessionResourceNotifyTransfer_ExtIEs_Extension_SecondaryRATUsageInformation = SecondaryRATUsageInformation_PDU,

    PDU_PDUSessionResourceNotifyTransfer_ExtIEs_Extension_MAX = INT_MAX
};

union PDUSessionResourceNotifyTransfer_ExtIEs_Extension_union {
    struct SecondaryRATUsageInformation *pdu_SecondaryRATUsageInformation;  
                                        /* PDU_PDUSessionResourceNotifyTransfer_ExtIEs_Extension_SecondaryRATUsageInformation */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PDUSessionResourceNotifyTransfer_ExtIEs_Extension {
    PDUSessionResourceNotifyTransfer_ExtIEs_Extension_PDUs pduNum;
    OssBuf          encoded;
    PDUSessionResourceNotifyTransfer_ExtIEs_Extension_union decoded;
} PDUSessionResourceNotifyTransfer_ExtIEs_Extension;

typedef struct _seq139 {
    ProtocolExtensionID id;
    Criticality     criticality;
    PDUSessionResourceNotifyTransfer_ExtIEs_Extension extensionValue;
} _seq139;

typedef struct _seqof110_ {
    struct _seqof110_ *next;
    _seq139         value;
} *_seqof110;

typedef struct PDUSessionResourceNotifyReleasedTransfer {
    unsigned char   bit_mask;
#       define      PDUSessionResourceNotifyReleasedTransfer_iE_Extensions_present 0x80
    Cause           cause;
    struct _seqof110_ *iE_Extensions;  /* optional; set in bit_mask
            * PDUSessionResourceNotifyReleasedTransfer_iE_Extensions_present if
            * present */
} PDUSessionResourceNotifyReleasedTransfer;

typedef struct PDUSessionResourceNotifyTransfer {
    unsigned char   bit_mask;
#       define      qosFlowNotifyList_present 0x80
#       define      qosFlowReleasedList_present 0x40
#       define      PDUSessionResourceNotifyTransfer_iE_Extensions_present 0x20
    struct QosFlowNotifyList_ *qosFlowNotifyList;  /* optional; set in bit_mask
                                                    * qosFlowNotifyList_present
                                                    * if present */
    struct QosFlowListWithCause_ *qosFlowReleasedList;  /* optional; set in
                                   * bit_mask qosFlowReleasedList_present if
                                   * present */
    struct _seqof110_ *iE_Extensions;  /* optional; set in bit_mask
                    * PDUSessionResourceNotifyTransfer_iE_Extensions_present if
                    * present */
} PDUSessionResourceNotifyTransfer;

typedef struct PDUSessionResourceReleaseCommandTransfer {
    unsigned char   bit_mask;
#       define      PDUSessionResourceReleaseCommandTransfer_iE_Extensions_present 0x80
    Cause           cause;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
            * PDUSessionResourceReleaseCommandTransfer_iE_Extensions_present if
            * present */
} PDUSessionResourceReleaseCommandTransfer;

typedef struct PDUSessionResourceReleasedItemNot {
    unsigned char   bit_mask;
#       define      PDUSessionResourceReleasedItemNot_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq143 {
        /* ContentsConstraint is applied to
         * pDUSessionResourceNotifyReleasedTransfer */
        _octet1         encoded;
        PDUSessionResourceNotifyReleasedTransfer *decoded;
    } pDUSessionResourceNotifyReleasedTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                   * PDUSessionResourceReleasedItemNot_iE_Extensions_present if
                   * present */
} PDUSessionResourceReleasedItemNot;

typedef struct PDUSessionResourceReleasedListNot_ {
    struct PDUSessionResourceReleasedListNot_ *next;
    PDUSessionResourceReleasedItemNot value;
} *PDUSessionResourceReleasedListNot;

typedef struct _seq145 {
    /* ContentsConstraint is applied to _seq145 */
    _octet1         encoded;
    PathSwitchRequestUnsuccessfulTransfer *decoded;
} _seq145;

typedef struct PDUSessionResourceReleasedItemPSAck {
    unsigned char   bit_mask;
#       define      PDUSessionResourceReleasedItemPSAck_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    _seq145         pathSwitchRequestUnsuccessfulTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                 * PDUSessionResourceReleasedItemPSAck_iE_Extensions_present if
                 * present */
} PDUSessionResourceReleasedItemPSAck;

typedef struct PDUSessionResourceReleasedListPSAck_ {
    struct PDUSessionResourceReleasedListPSAck_ *next;
    PDUSessionResourceReleasedItemPSAck value;
} *PDUSessionResourceReleasedListPSAck;

typedef struct PDUSessionResourceReleasedItemPSFail {
    unsigned char   bit_mask;
#       define      PDUSessionResourceReleasedItemPSFail_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    _seq145         pathSwitchRequestUnsuccessfulTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                * PDUSessionResourceReleasedItemPSFail_iE_Extensions_present if
                * present */
} PDUSessionResourceReleasedItemPSFail;

typedef struct PDUSessionResourceReleasedListPSFail_ {
    struct PDUSessionResourceReleasedListPSFail_ *next;
    PDUSessionResourceReleasedItemPSFail value;
} *PDUSessionResourceReleasedListPSFail;

typedef struct PDUSessionResourceReleasedItemRelRes {
    unsigned char   bit_mask;
#       define      PDUSessionResourceReleasedItemRelRes_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq148 {
        /* ContentsConstraint is applied to
         * pDUSessionResourceReleaseResponseTransfer */
        _octet1         encoded;
        struct PDUSessionResourceReleaseResponseTransfer *decoded;
    } pDUSessionResourceReleaseResponseTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                * PDUSessionResourceReleasedItemRelRes_iE_Extensions_present if
                * present */
} PDUSessionResourceReleasedItemRelRes;

typedef struct PDUSessionResourceReleasedListRelRes_ {
    struct PDUSessionResourceReleasedListRelRes_ *next;
    PDUSessionResourceReleasedItemRelRes value;
} *PDUSessionResourceReleasedListRelRes;

typedef struct PDUSessionResourceReleaseResponseTransfer {
    unsigned char   bit_mask;
#       define      PDUSessionResourceReleaseResponseTransfer_iE_Extensions_present 0x80
    struct _seqof110_ *iE_Extensions;  /* optional; set in bit_mask
           * PDUSessionResourceReleaseResponseTransfer_iE_Extensions_present if
           * present */
} PDUSessionResourceReleaseResponseTransfer;

typedef struct PDUSessionResourceSecondaryRATUsageItem {
    unsigned char   bit_mask;
#       define      PDUSessionResourceSecondaryRATUsageItem_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq151 {
        /* ContentsConstraint is applied to
         * secondaryRATDataUsageReportTransfer */
        _octet1         encoded;
        struct SecondaryRATDataUsageReportTransfer *decoded;
    } secondaryRATDataUsageReportTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
             * PDUSessionResourceSecondaryRATUsageItem_iE_Extensions_present if
             * present */
} PDUSessionResourceSecondaryRATUsageItem;

typedef struct PDUSessionResourceSecondaryRATUsageList_ {
    struct PDUSessionResourceSecondaryRATUsageList_ *next;
    PDUSessionResourceSecondaryRATUsageItem value;
} *PDUSessionResourceSecondaryRATUsageList;

typedef struct _seq153 {
    /* ContentsConstraint is applied to _seq153 */
    _octet1         encoded;
    struct PDUSessionResourceSetupRequestTransfer *decoded;
} _seq153;

typedef struct PDUSessionResourceSetupItemCxtReq {
    unsigned char   bit_mask;
#       define      PDUSessionResourceSetupItemCxtReq_nAS_PDU_present 0x80
#       define      PDUSessionResourceSetupItemCxtReq_iE_Extensions_present 0x40
    PDUSessionID    pDUSessionID;
    NAS_PDU         nAS_PDU;  /* optional; set in bit_mask
                         * PDUSessionResourceSetupItemCxtReq_nAS_PDU_present if
                         * present */
    S_NSSAI         s_NSSAI;
    _seq153         pDUSessionResourceSetupRequestTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                   * PDUSessionResourceSetupItemCxtReq_iE_Extensions_present if
                   * present */
} PDUSessionResourceSetupItemCxtReq;

typedef struct PDUSessionResourceSetupListCxtReq_ {
    struct PDUSessionResourceSetupListCxtReq_ *next;
    PDUSessionResourceSetupItemCxtReq value;
} *PDUSessionResourceSetupListCxtReq;

typedef struct _seq155 {
    /* ContentsConstraint is applied to _seq155 */
    _octet1         encoded;
    struct PDUSessionResourceSetupResponseTransfer *decoded;
} _seq155;

typedef struct PDUSessionResourceSetupItemCxtRes {
    unsigned char   bit_mask;
#       define      PDUSessionResourceSetupItemCxtRes_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    _seq155         pDUSessionResourceSetupResponseTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                   * PDUSessionResourceSetupItemCxtRes_iE_Extensions_present if
                   * present */
} PDUSessionResourceSetupItemCxtRes;

typedef struct PDUSessionResourceSetupListCxtRes_ {
    struct PDUSessionResourceSetupListCxtRes_ *next;
    PDUSessionResourceSetupItemCxtRes value;
} *PDUSessionResourceSetupListCxtRes;

typedef struct PDUSessionResourceSetupItemHOReq {
    unsigned char   bit_mask;
#       define      PDUSessionResourceSetupItemHOReq_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    S_NSSAI         s_NSSAI;
    _seq153         handoverRequestTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                    * PDUSessionResourceSetupItemHOReq_iE_Extensions_present if
                    * present */
} PDUSessionResourceSetupItemHOReq;

typedef struct PDUSessionResourceSetupListHOReq_ {
    struct PDUSessionResourceSetupListHOReq_ *next;
    PDUSessionResourceSetupItemHOReq value;
} *PDUSessionResourceSetupListHOReq;

typedef struct PDUSessionResourceSetupItemSUReq {
    unsigned char   bit_mask;
#       define      pDUSessionNAS_PDU_present 0x80
#       define      PDUSessionResourceSetupItemSUReq_iE_Extensions_present 0x40
    PDUSessionID    pDUSessionID;
    NAS_PDU         pDUSessionNAS_PDU;  /* optional; set in bit_mask
                                         * pDUSessionNAS_PDU_present if
                                         * present */
    S_NSSAI         s_NSSAI;
    _seq153         pDUSessionResourceSetupRequestTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                    * PDUSessionResourceSetupItemSUReq_iE_Extensions_present if
                    * present */
} PDUSessionResourceSetupItemSUReq;

typedef struct PDUSessionResourceSetupListSUReq_ {
    struct PDUSessionResourceSetupListSUReq_ *next;
    PDUSessionResourceSetupItemSUReq value;
} *PDUSessionResourceSetupListSUReq;

typedef struct PDUSessionResourceSetupItemSURes {
    unsigned char   bit_mask;
#       define      PDUSessionResourceSetupItemSURes_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    _seq155         pDUSessionResourceSetupResponseTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                    * PDUSessionResourceSetupItemSURes_iE_Extensions_present if
                    * present */
} PDUSessionResourceSetupItemSURes;

typedef struct PDUSessionResourceSetupListSURes_ {
    struct PDUSessionResourceSetupListSURes_ *next;
    PDUSessionResourceSetupItemSURes value;
} *PDUSessionResourceSetupListSURes;

enum PDUSessionResourceSetupRequestTransferIEs_Value_PDUs {
    PDU_PDUSessionResourceSetupRequestTransferIEs_Value_UNKNOWN = 0,

    PDU_PDUSessionResourceSetupRequestTransferIEs_Value_PDUSessionAggregateMaximumBitRate = PDUSessionAggregateMaximumBitRate_PDU,
    PDU_PDUSessionResourceSetupRequestTransferIEs_Value_UPTransportLayerInformation = UPTransportLayerInformation_PDU,
    PDU_PDUSessionResourceSetupRequestTransferIEs_Value_UPTransportLayerInformationList = UPTransportLayerInformationList_PDU,
    PDU_PDUSessionResourceSetupRequestTransferIEs_Value_DataForwardingNotPossible = DataForwardingNotPossible_PDU,
    PDU_PDUSessionResourceSetupRequestTransferIEs_Value_PDUSessionType = PDUSessionType_PDU,
    PDU_PDUSessionResourceSetupRequestTransferIEs_Value_SecurityIndication = SecurityIndication_PDU,
    PDU_PDUSessionResourceSetupRequestTransferIEs_Value_NetworkInstance = NetworkInstance_PDU,
    PDU_PDUSessionResourceSetupRequestTransferIEs_Value_QosFlowSetupRequestList = QosFlowSetupRequestList_PDU,
    PDU_PDUSessionResourceSetupRequestTransferIEs_Value_CommonNetworkInstance = CommonNetworkInstance_PDU,

    PDU_PDUSessionResourceSetupRequestTransferIEs_Value_MAX = INT_MAX
};

union PDUSessionResourceSetupRequestTransferIEs_Value_union {
    PDUSessionAggregateMaximumBitRate *pdu_PDUSessionAggregateMaximumBitRate;                                           /* PDU_PDUSessionResourceSetupRequestTransferIEs_Value_PDUSessionAggregateMaximumBitRate */
    UPTransportLayerInformation *pdu_UPTransportLayerInformation;  
      /* PDU_PDUSessionResourceSetupRequestTransferIEs_Value_UPTransportLayerInformation */
    struct UPTransportLayerInformationList_ **pdu_UPTransportLayerInformationList;                                      /* PDU_PDUSessionResourceSetupRequestTransferIEs_Value_UPTransportLayerInformationList */
    DataForwardingNotPossible *pdu_DataForwardingNotPossible;  
    /* PDU_PDUSessionResourceSetupRequestTransferIEs_Value_DataForwardingNotPossible */
    PDUSessionType  *pdu_PDUSessionType;  
       /* PDU_PDUSessionResourceSetupRequestTransferIEs_Value_PDUSessionType */
    SecurityIndication *pdu_SecurityIndication;  
   /* PDU_PDUSessionResourceSetupRequestTransferIEs_Value_SecurityIndication */
    NetworkInstance *pdu_NetworkInstance;  
      /* PDU_PDUSessionResourceSetupRequestTransferIEs_Value_NetworkInstance */
    struct QosFlowSetupRequestList_ **pdu_QosFlowSetupRequestList;  
  /* PDU_PDUSessionResourceSetupRequestTransferIEs_Value_QosFlowSetupRequestList */
    CommonNetworkInstance *pdu_CommonNetworkInstance;  
/* PDU_PDUSessionResourceSetupRequestTransferIEs_Value_CommonNetworkInstance */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PDUSessionResourceSetupRequestTransferIEs_Value {
    PDUSessionResourceSetupRequestTransferIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PDUSessionResourceSetupRequestTransferIEs_Value_union decoded;
} PDUSessionResourceSetupRequestTransferIEs_Value;

typedef struct PDUSessionResourceSetupRequestTransfer {
    struct _seqof124 {
        struct _seqof124 *next;
        struct _seq159 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PDUSessionResourceSetupRequestTransferIEs_Value value;
        } value;
    } *protocolIEs;
} PDUSessionResourceSetupRequestTransfer;

typedef struct PDUSessionResourceSetupResponseTransfer {
    unsigned char   bit_mask;
#       define      PDUSessionResourceSetupResponseTransfer_additionalDLQosFlowPerTNLInformation_present 0x80
#       define      PDUSessionResourceSetupResponseTransfer_securityResult_present 0x40
#       define      PDUSessionResourceSetupResponseTransfer_qosFlowFailedToSetupList_present 0x20
#       define      PDUSessionResourceSetupResponseTransfer_iE_Extensions_present 0x10
    QosFlowPerTNLInformation dLQosFlowPerTNLInformation;
    struct QosFlowPerTNLInformationList_ *additionalDLQosFlowPerTNLInformation;                                         /* optional; set in bit_mask
            * PDUSessionResourceSetupResponseTransfer_additionalDLQosFlowPerTNLInformation_present if
            * present */
    SecurityResult  securityResult;  /* optional; set in bit_mask
            * PDUSessionResourceSetupResponseTransfer_securityResult_present if
            * present */
    struct QosFlowListWithCause_ *qosFlowFailedToSetupList;  /* optional; set in
                                   * bit_mask
  * PDUSessionResourceSetupResponseTransfer_qosFlowFailedToSetupList_present if
  * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
             * PDUSessionResourceSetupResponseTransfer_iE_Extensions_present if
             * present */
} PDUSessionResourceSetupResponseTransfer;

typedef struct PDUSessionResourceSetupUnsuccessfulTransfer {
    unsigned char   bit_mask;
#       define      PDUSessionResourceSetupUnsuccessfulTransfer_criticalityDiagnostics_present 0x80
#       define      PDUSessionResourceSetupUnsuccessfulTransfer_iE_Extensions_present 0x40
    Cause           cause;
    CriticalityDiagnostics criticalityDiagnostics;  /* optional; set in bit_mask
  * PDUSessionResourceSetupUnsuccessfulTransfer_criticalityDiagnostics_present if
  * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
         * PDUSessionResourceSetupUnsuccessfulTransfer_iE_Extensions_present if
         * present */
} PDUSessionResourceSetupUnsuccessfulTransfer;

typedef struct PDUSessionResourceSwitchedItem {
    unsigned char   bit_mask;
#       define      PDUSessionResourceSwitchedItem_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq163 {
        /* ContentsConstraint is applied to
         * pathSwitchRequestAcknowledgeTransfer */
        _octet1         encoded;
        PathSwitchRequestAcknowledgeTransfer *decoded;
    } pathSwitchRequestAcknowledgeTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                      * PDUSessionResourceSwitchedItem_iE_Extensions_present if
                      * present */
} PDUSessionResourceSwitchedItem;

typedef struct PDUSessionResourceSwitchedList_ {
    struct PDUSessionResourceSwitchedList_ *next;
    PDUSessionResourceSwitchedItem value;
} *PDUSessionResourceSwitchedList;

typedef struct PDUSessionResourceToBeSwitchedDLItem {
    unsigned char   bit_mask;
#       define      PDUSessionResourceToBeSwitchedDLItem_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq165 {
        /* ContentsConstraint is applied to pathSwitchRequestTransfer */
        _octet1         encoded;
        PathSwitchRequestTransfer *decoded;
    } pathSwitchRequestTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                * PDUSessionResourceToBeSwitchedDLItem_iE_Extensions_present if
                * present */
} PDUSessionResourceToBeSwitchedDLItem;

typedef struct PDUSessionResourceToBeSwitchedDLList_ {
    struct PDUSessionResourceToBeSwitchedDLList_ *next;
    PDUSessionResourceToBeSwitchedDLItem value;
} *PDUSessionResourceToBeSwitchedDLList;

typedef struct PDUSessionResourceToReleaseItemHOCmd {
    unsigned char   bit_mask;
#       define      PDUSessionResourceToReleaseItemHOCmd_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq167 {
        /* ContentsConstraint is applied to
         * handoverPreparationUnsuccessfulTransfer */
        _octet1         encoded;
        HandoverPreparationUnsuccessfulTransfer *decoded;
    } handoverPreparationUnsuccessfulTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                * PDUSessionResourceToReleaseItemHOCmd_iE_Extensions_present if
                * present */
} PDUSessionResourceToReleaseItemHOCmd;

typedef struct PDUSessionResourceToReleaseListHOCmd_ {
    struct PDUSessionResourceToReleaseListHOCmd_ *next;
    PDUSessionResourceToReleaseItemHOCmd value;
} *PDUSessionResourceToReleaseListHOCmd;

typedef struct PDUSessionResourceToReleaseItemRelCmd {
    unsigned char   bit_mask;
#       define      PDUSessionResourceToReleaseItemRelCmd_iE_Extensions_present 0x80
    PDUSessionID    pDUSessionID;
    struct _seq169 {
        /* ContentsConstraint is applied to
         * pDUSessionResourceReleaseCommandTransfer */
        _octet1         encoded;
        PDUSessionResourceReleaseCommandTransfer *decoded;
    } pDUSessionResourceReleaseCommandTransfer;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
               * PDUSessionResourceToReleaseItemRelCmd_iE_Extensions_present if
               * present */
} PDUSessionResourceToReleaseItemRelCmd;

typedef struct PDUSessionResourceToReleaseListRelCmd_ {
    struct PDUSessionResourceToReleaseListRelCmd_ *next;
    PDUSessionResourceToReleaseItemRelCmd value;
} *PDUSessionResourceToReleaseListRelCmd;

typedef enum _enum3 {
    nr = 0,
    eutra = 1
} _enum3;

typedef struct PDUSessionUsageReport {
    unsigned char   bit_mask;
#       define      PDUSessionUsageReport_iE_Extensions_present 0x80
    _enum3          rATType;
    struct VolumeTimedReportList_ *pDUSessionTimedReportList;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                               * PDUSessionUsageReport_iE_Extensions_present if
                               * present */
} PDUSessionUsageReport;

typedef struct PLMNSupportItem {
    unsigned char   bit_mask;
#       define      PLMNSupportItem_iE_Extensions_present 0x80
    PLMNIdentity    pLMNIdentity;
    struct SliceSupportList_ *sliceSupportList;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * PLMNSupportItem_iE_Extensions_present if
                                      * present */
} PLMNSupportItem;

typedef struct PLMNSupportList_ {
    struct PLMNSupportList_ *next;
    PLMNSupportItem value;
} *PLMNSupportList;

typedef struct PortNumber {
    unsigned short  length;
    unsigned char   value[2];
} PortNumber;

typedef struct PWSFailedCellIDList {
    unsigned short  choice;
#       define      eUTRA_CGI_PWSFailedList_chosen 1
#       define      nR_CGI_PWSFailedList_chosen 2
#       define      PWSFailedCellIDList_choice_Extensions_chosen 3
    union _union {
        struct EUTRA_CGIList_ *eUTRA_CGI_PWSFailedList;  /* to choose, set
                                   * choice to eUTRA_CGI_PWSFailedList_chosen */
        struct NR_CGIList_ *nR_CGI_PWSFailedList;  /* to choose, set choice to
                                               * nR_CGI_PWSFailedList_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to
                              * PWSFailedCellIDList_choice_Extensions_chosen */
    } u;
} PWSFailedCellIDList;

// Q
typedef struct QosCharacteristics {
    unsigned short  choice;
#       define      nonDynamic5QI_chosen 1
#       define      dynamic5QI_chosen 2
#       define      QosCharacteristics_choice_Extensions_chosen 3
    union _union {
        NonDynamic5QIDescriptor nonDynamic5QI;  /* to choose, set choice to
                                                 * nonDynamic5QI_chosen */
        Dynamic5QIDescriptor dynamic5QI;  /* to choose, set choice to
                                           * dynamic5QI_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to
                               * QosCharacteristics_choice_Extensions_chosen */
    } u;
} QosCharacteristics;

typedef struct QosFlowAcceptedItem {
    unsigned char   bit_mask;
#       define      QosFlowAcceptedItem_iE_Extensions_present 0x80
    QosFlowIdentifier qosFlowIdentifier;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                 * QosFlowAcceptedItem_iE_Extensions_present if
                                 * present */
} QosFlowAcceptedItem;

typedef struct QosFlowAcceptedList_ {
    struct QosFlowAcceptedList_ *next;
    QosFlowAcceptedItem value;
} *QosFlowAcceptedList;

typedef enum ReflectiveQosAttribute {
    subject_to = 0
} ReflectiveQosAttribute;

typedef struct QosFlowLevelQosParameters {
    unsigned char   bit_mask;
#       define      gBR_QosInformation_present 0x80
#       define      reflectiveQosAttribute_present 0x40
#       define      additionalQosFlowInformation_present 0x20
#       define      QosFlowLevelQosParameters_iE_Extensions_present 0x10
    QosCharacteristics qosCharacteristics;
    AllocationAndRetentionPriority allocationAndRetentionPriority;
    GBR_QosInformation gBR_QosInformation;  /* optional; set in bit_mask
                                             * gBR_QosInformation_present if
                                             * present */
    ReflectiveQosAttribute reflectiveQosAttribute;  /* optional; set in bit_mask
                                            * reflectiveQosAttribute_present if
                                            * present */
    AdditionalQosFlowInformation additionalQosFlowInformation;  /* optional; set
                                   * in bit_mask
                                   * additionalQosFlowInformation_present if
                                   * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                           * QosFlowLevelQosParameters_iE_Extensions_present if
                           * present */
} QosFlowLevelQosParameters;

typedef struct QosFlowAddOrModifyRequestItem {
    unsigned char   bit_mask;
#       define      qosFlowLevelQosParameters_present 0x80
#       define      QosFlowAddOrModifyRequestItem_e_RAB_ID_present 0x40
#       define      QosFlowAddOrModifyRequestItem_iE_Extensions_present 0x20
    QosFlowIdentifier qosFlowIdentifier;
    QosFlowLevelQosParameters qosFlowLevelQosParameters;  /* optional; set in
                                   * bit_mask qosFlowLevelQosParameters_present
                                   * if present */
    E_RAB_ID        e_RAB_ID;  /* optional; set in bit_mask
                                * QosFlowAddOrModifyRequestItem_e_RAB_ID_present
                                * if present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                       * QosFlowAddOrModifyRequestItem_iE_Extensions_present if
                       * present */
} QosFlowAddOrModifyRequestItem;

typedef struct QosFlowAddOrModifyRequestList_ {
    struct QosFlowAddOrModifyRequestList_ *next;
    QosFlowAddOrModifyRequestItem value;
} *QosFlowAddOrModifyRequestList;

typedef struct QosFlowAddOrModifyResponseItem {
    unsigned char   bit_mask;
#       define      QosFlowAddOrModifyResponseItem_iE_Extensions_present 0x80
    QosFlowIdentifier qosFlowIdentifier;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                      * QosFlowAddOrModifyResponseItem_iE_Extensions_present if
                      * present */
} QosFlowAddOrModifyResponseItem;

typedef struct QosFlowAddOrModifyResponseList_ {
    struct QosFlowAddOrModifyResponseList_ *next;
    QosFlowAddOrModifyResponseItem value;
} *QosFlowAddOrModifyResponseList;

typedef struct QosFlowWithCauseItem {
    unsigned char   bit_mask;
#       define      QosFlowWithCauseItem_iE_Extensions_present 0x80
    QosFlowIdentifier qosFlowIdentifier;
    Cause           cause;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                * QosFlowWithCauseItem_iE_Extensions_present if
                                * present */
} QosFlowWithCauseItem;

typedef struct QosFlowListWithCause_ {
    struct QosFlowListWithCause_ *next;
    QosFlowWithCauseItem value;
} *QosFlowListWithCause;

typedef struct QosFlowModifyConfirmItem {
    unsigned char   bit_mask;
#       define      QosFlowModifyConfirmItem_iE_Extensions_present 0x80
    QosFlowIdentifier qosFlowIdentifier;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                            * QosFlowModifyConfirmItem_iE_Extensions_present if
                            * present */
} QosFlowModifyConfirmItem;

typedef struct QosFlowModifyConfirmList_ {
    struct QosFlowModifyConfirmList_ *next;
    QosFlowModifyConfirmItem value;
} *QosFlowModifyConfirmList;

typedef struct QosFlowNotifyItem {
    unsigned char   bit_mask;
#       define      QosFlowNotifyItem_iE_Extensions_present 0x80
    QosFlowIdentifier qosFlowIdentifier;
    NotificationCause notificationCause;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * QosFlowNotifyItem_iE_Extensions_present
                                      * if present */
} QosFlowNotifyItem;

typedef struct QosFlowNotifyList_ {
    struct QosFlowNotifyList_ *next;
    QosFlowNotifyItem value;
} *QosFlowNotifyList;

typedef struct QosFlowPerTNLInformationItem {
    unsigned char   bit_mask;
#       define      QosFlowPerTNLInformationItem_iE_Extensions_present 0x80
    QosFlowPerTNLInformation qosFlowPerTNLInformation;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                        * QosFlowPerTNLInformationItem_iE_Extensions_present if
                        * present */
} QosFlowPerTNLInformationItem;

typedef struct QosFlowPerTNLInformationList_ {
    struct QosFlowPerTNLInformationList_ *next;
    QosFlowPerTNLInformationItem value;
} *QosFlowPerTNLInformationList;

typedef struct QosFlowSetupRequestItem {
    unsigned char   bit_mask;
#       define      QosFlowSetupRequestItem_e_RAB_ID_present 0x80
#       define      QosFlowSetupRequestItem_iE_Extensions_present 0x40
    QosFlowIdentifier qosFlowIdentifier;
    QosFlowLevelQosParameters qosFlowLevelQosParameters;
    E_RAB_ID        e_RAB_ID;  /* optional; set in bit_mask
                                * QosFlowSetupRequestItem_e_RAB_ID_present if
                                * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                             * QosFlowSetupRequestItem_iE_Extensions_present if
                             * present */
} QosFlowSetupRequestItem;

typedef struct QosFlowSetupRequestList_ {
    struct QosFlowSetupRequestList_ *next;
    QosFlowSetupRequestItem value;
} *QosFlowSetupRequestList;

typedef struct QosFlowItemWithDataForwarding {
    unsigned char   bit_mask;
#       define      dataForwardingAccepted_present 0x80
#       define      QosFlowItemWithDataForwarding_iE_Extensions_present 0x40
    QosFlowIdentifier qosFlowIdentifier;
    DataForwardingAccepted dataForwardingAccepted;  /* optional; set in bit_mask
                                            * dataForwardingAccepted_present if
                                            * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                       * QosFlowItemWithDataForwarding_iE_Extensions_present if
                       * present */
} QosFlowItemWithDataForwarding;

typedef struct QosFlowListWithDataForwarding_ {
    struct QosFlowListWithDataForwarding_ *next;
    QosFlowItemWithDataForwarding value;
} *QosFlowListWithDataForwarding;

typedef struct QosFlowToBeForwardedItem {
    unsigned char   bit_mask;
#       define      QosFlowToBeForwardedItem_iE_Extensions_present 0x80
    QosFlowIdentifier qosFlowIdentifier;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                            * QosFlowToBeForwardedItem_iE_Extensions_present if
                            * present */
} QosFlowToBeForwardedItem;

typedef struct QosFlowToBeForwardedList_ {
    struct QosFlowToBeForwardedList_ *next;
    QosFlowToBeForwardedItem value;
} *QosFlowToBeForwardedList;

typedef struct QoSFlowsUsageReport_Item {
    unsigned char   bit_mask;
#       define      QoSFlowsUsageReport_Item_iE_Extensions_present 0x80
    QosFlowIdentifier qosFlowIdentifier;
    _enum3          rATType;
    struct VolumeTimedReportList_ *qoSFlowsTimedReportList;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                            * QoSFlowsUsageReport_Item_iE_Extensions_present if
                            * present */
} QoSFlowsUsageReport_Item;

typedef struct QoSFlowsUsageReportList_ {
    struct QoSFlowsUsageReportList_ *next;
    QoSFlowsUsageReport_Item value;
} *QoSFlowsUsageReportList;

// R
typedef struct RANNodeName {
    unsigned short  length;
    char            *value;
} RANNodeName;

typedef struct RANStatusTransfer_TransparentContainer {
    unsigned char   bit_mask;
#       define      RANStatusTransfer_TransparentContainer_iE_Extensions_present 0x80
    struct DRBsSubjectToStatusTransferList_ *dRBsSubjectToStatusTransferList;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
              * RANStatusTransfer_TransparentContainer_iE_Extensions_present if
              * present */
} RANStatusTransfer_TransparentContainer;

typedef struct RATRestrictionInformation {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} RATRestrictionInformation;

typedef struct RATRestrictions_Item {
    unsigned char   bit_mask;
#       define      RATRestrictions_Item_iE_Extensions_present 0x80
    PLMNIdentity    pLMNIdentity;
    RATRestrictionInformation rATRestrictionInformation;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                * RATRestrictions_Item_iE_Extensions_present if
                                * present */
} RATRestrictions_Item;

typedef struct RATRestrictions_ {
    struct RATRestrictions_ *next;
    RATRestrictions_Item value;
} *RATRestrictions;

typedef struct RecommendedCellItem {
    unsigned char   bit_mask;
#       define      RecommendedCellItem_timeStayedInCell_present 0x80
#       define      RecommendedCellItem_iE_Extensions_present 0x40
    NGRAN_CGI       nGRAN_CGI;
    unsigned short  timeStayedInCell;  /* optional; set in bit_mask
                              * RecommendedCellItem_timeStayedInCell_present if
                              * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                 * RecommendedCellItem_iE_Extensions_present if
                                 * present */
} RecommendedCellItem;

typedef struct RecommendedCellList_ {
    struct RecommendedCellList_ *next;
    RecommendedCellItem value;
} *RecommendedCellList;

typedef struct RecommendedRANNodeItem {
    unsigned char   bit_mask;
#       define      RecommendedRANNodeItem_iE_Extensions_present 0x80
    AMFPagingTarget aMFPagingTarget;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                              * RecommendedRANNodeItem_iE_Extensions_present if
                              * present */
} RecommendedRANNodeItem;

typedef struct RecommendedRANNodeList_ {
    struct RecommendedRANNodeList_ *next;
    RecommendedRANNodeItem value;
} *RecommendedRANNodeList;

typedef enum ResetAll {
    reset_all = 0
} ResetAll;

typedef struct ResetType {
    unsigned short  choice;
#       define      nG_Interface_chosen 1
#       define      partOfNG_Interface_chosen 2
#       define      ResetType_choice_Extensions_chosen 3
    union _union {
        ResetAll        nG_Interface;  /* to choose, set choice to
                                        * nG_Interface_chosen */
        struct UE_associatedLogicalNG_connectionList_ *partOfNG_Interface;  
                                        /* to choose, set choice to
                                         * partOfNG_Interface_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                     * to ResetType_choice_Extensions_chosen */
    } u;
} ResetType;

typedef struct RoutingID {
    unsigned int    length;
    unsigned char   *value;
} RoutingID;

typedef struct SecondaryRATUsageInformation {
    unsigned char   bit_mask;
#       define      pDUSessionUsageReport_present 0x80
#       define      qosFlowsUsageReportList_present 0x40
#       define      SecondaryRATUsageInformation_iE_Extension_present 0x20
    PDUSessionUsageReport pDUSessionUsageReport;  /* optional; set in bit_mask
                                             * pDUSessionUsageReport_present if
                                             * present */
    struct QoSFlowsUsageReportList_ *qosFlowsUsageReportList;  /* optional; set
                                   * in bit_mask qosFlowsUsageReportList_present
                                   * if present */
    struct _seqof1_ *iE_Extension;  /* optional; set in bit_mask
                         * SecondaryRATUsageInformation_iE_Extension_present if
                         * present */
} SecondaryRATUsageInformation;

typedef struct SecondaryRATDataUsageReportTransfer {
    unsigned char   bit_mask;
#       define      secondaryRATUsageInformation_present 0x80
#       define      SecondaryRATDataUsageReportTransfer_iE_Extensions_present 0x40
    SecondaryRATUsageInformation secondaryRATUsageInformation;  /* optional; set
                                   * in bit_mask
                                   * secondaryRATUsageInformation_present if
                                   * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                 * SecondaryRATDataUsageReportTransfer_iE_Extensions_present if
                 * present */
} SecondaryRATDataUsageReportTransfer;

typedef struct SecurityKey {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} SecurityKey;

typedef struct SecurityContext {
    unsigned char   bit_mask;
#       define      SecurityContext_iE_Extensions_present 0x80
    NextHopChainingCount nextHopChainingCount;
    SecurityKey     nextHopNH;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * SecurityContext_iE_Extensions_present if
                                      * present */
} SecurityContext;

typedef struct SerialNumber {
    unsigned short  length;  /* number of significant bits */
    unsigned char   *value;
} SerialNumber;

typedef struct ServedGUAMIItem {
    unsigned char   bit_mask;
#       define      ServedGUAMIItem_backupAMFName_present 0x80
#       define      ServedGUAMIItem_iE_Extensions_present 0x40
    GUAMI           gUAMI;
    AMFName         backupAMFName;  /* optional; set in bit_mask
                                     * ServedGUAMIItem_backupAMFName_present if
                                     * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * ServedGUAMIItem_iE_Extensions_present if
                                      * present */
} ServedGUAMIItem;

typedef struct ServedGUAMIList_ {
    struct ServedGUAMIList_ *next;
    ServedGUAMIItem value;
} *ServedGUAMIList;

typedef struct ServiceAreaInformation_Item {
    unsigned char   bit_mask;
#       define      allowedTACs_present 0x80
#       define      notAllowedTACs_present 0x40
#       define      ServiceAreaInformation_Item_iE_Extensions_present 0x20
    PLMNIdentity    pLMNIdentity;
    struct AllowedTACs_ *allowedTACs;  /* optional; set in bit_mask
                                        * allowedTACs_present if present */
    struct NotAllowedTACs_ *notAllowedTACs;  /* optional; set in bit_mask
                                              * notAllowedTACs_present if
                                              * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                         * ServiceAreaInformation_Item_iE_Extensions_present if
                         * present */
} ServiceAreaInformation_Item;

typedef struct ServiceAreaInformation_ {
    struct ServiceAreaInformation_ *next;
    ServiceAreaInformation_Item value;
} *ServiceAreaInformation;

typedef struct SliceOverloadItem {
    unsigned char   bit_mask;
#       define      SliceOverloadItem_iE_Extensions_present 0x80
    S_NSSAI         s_NSSAI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * SliceOverloadItem_iE_Extensions_present
                                      * if present */
} SliceOverloadItem;

typedef struct SliceOverloadList_ {
    struct SliceOverloadList_ *next;
    SliceOverloadItem value;
} *SliceOverloadList;

typedef struct SliceSupportItem {
    unsigned char   bit_mask;
#       define      SliceSupportItem_iE_Extensions_present 0x80
    S_NSSAI         s_NSSAI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * SliceSupportItem_iE_Extensions_present
                                      * if present */
} SliceSupportItem;

typedef struct SliceSupportList_ {
    struct SliceSupportList_ *next;
    SliceSupportItem value;
} *SliceSupportList;

typedef struct TargetRANNodeID {
    unsigned char   bit_mask;
#       define      TargetRANNodeID_iE_Extensions_present 0x80
    GlobalRANNodeID globalRANNodeID;
    TAI             selectedTAI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * TargetRANNodeID_iE_Extensions_present if
                                      * present */
} TargetRANNodeID;

typedef struct SourceRANNodeID {
    unsigned char   bit_mask;
#       define      SourceRANNodeID_iE_Extensions_present 0x80
    GlobalRANNodeID globalRANNodeID;
    TAI             selectedTAI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * SourceRANNodeID_iE_Extensions_present if
                                      * present */
} SourceRANNodeID;

typedef enum SONInformationRequest {
    xn_TNL_configuration_info = 0
} SONInformationRequest;

typedef struct XnTNLConfigurationInfo {
    unsigned char   bit_mask;
#       define      xnExtendedTransportLayerAddresses_present 0x80
#       define      XnTNLConfigurationInfo_iE_Extensions_present 0x40
    struct XnTLAs_  *xnTransportLayerAddresses;
    struct XnExtTLAs_ *xnExtendedTransportLayerAddresses;  /* optional; set in
                                   * bit_mask
                                   * xnExtendedTransportLayerAddresses_present
                                   * if present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                              * XnTNLConfigurationInfo_iE_Extensions_present if
                              * present */
} XnTNLConfigurationInfo;

typedef struct SONInformationReply {
    unsigned char   bit_mask;
#       define      SONInformationReply_xnTNLConfigurationInfo_present 0x80
#       define      SONInformationReply_iE_Extensions_present 0x40
    XnTNLConfigurationInfo xnTNLConfigurationInfo;  /* optional; set in bit_mask
                        * SONInformationReply_xnTNLConfigurationInfo_present if
                        * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                 * SONInformationReply_iE_Extensions_present if
                                 * present */
} SONInformationReply;

typedef struct SONInformation {
    unsigned short  choice;
#       define      sONInformationRequest_chosen 1
#       define      sONInformationReply_chosen 2
#       define      SONInformation_choice_Extensions_chosen 3
    union _union {
        SONInformationRequest sONInformationRequest;  /* to choose, set choice
                                           * to sONInformationRequest_chosen */
        SONInformationReply sONInformationReply;  /* to choose, set choice to
                                                * sONInformationReply_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to
                                   * SONInformation_choice_Extensions_chosen */
    } u;
} SONInformation;

typedef struct SONConfigurationTransfer {
    unsigned char   bit_mask;
#       define      SONConfigurationTransfer_xnTNLConfigurationInfo_present 0x80
#       define      SONConfigurationTransfer_iE_Extensions_present 0x40
    TargetRANNodeID targetRANNodeID;
    SourceRANNodeID sourceRANNodeID;
    SONInformation  sONInformation;
    XnTNLConfigurationInfo xnTNLConfigurationInfo;  /* optional; set in bit_mask
                   * SONConfigurationTransfer_xnTNLConfigurationInfo_present if
                   * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                            * SONConfigurationTransfer_iE_Extensions_present if
                            * present */
// The above IE shall be present if the SON Information IE contains the SON Information Request IE set to “Xn TNL Configuration Info”
} SONConfigurationTransfer;

typedef struct SourceToTarget_TransparentContainer {
    unsigned int    length;
    unsigned char   *value;
} SourceToTarget_TransparentContainer;

typedef struct SupportedTAItem {
    unsigned char   bit_mask;
#       define      SupportedTAItem_iE_Extensions_present 0x80
    TAC             tAC;
    struct BroadcastPLMNList_ *broadcastPLMNList;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * SupportedTAItem_iE_Extensions_present if
                                      * present */
} SupportedTAItem;

typedef struct SupportedTAList_ {
    struct SupportedTAList_ *next;
    SupportedTAItem value;
} *SupportedTAList;

typedef struct TAIBroadcastEUTRA_Item {
    unsigned char   bit_mask;
#       define      TAIBroadcastEUTRA_Item_iE_Extensions_present 0x80
    TAI             tAI;
    struct CompletedCellsInTAI_EUTRA_ *completedCellsInTAI_EUTRA;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                              * TAIBroadcastEUTRA_Item_iE_Extensions_present if
                              * present */
} TAIBroadcastEUTRA_Item;

typedef struct TAIBroadcastEUTRA_ {
    struct TAIBroadcastEUTRA_ *next;
    TAIBroadcastEUTRA_Item value;
} *TAIBroadcastEUTRA;

typedef struct TAIBroadcastNR_Item {
    unsigned char   bit_mask;
#       define      TAIBroadcastNR_Item_iE_Extensions_present 0x80
    TAI             tAI;
    struct CompletedCellsInTAI_NR_ *completedCellsInTAI_NR;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                 * TAIBroadcastNR_Item_iE_Extensions_present if
                                 * present */
} TAIBroadcastNR_Item;

typedef struct TAIBroadcastNR_ {
    struct TAIBroadcastNR_ *next;
    TAIBroadcastNR_Item value;
} *TAIBroadcastNR;

typedef struct TAICancelledEUTRA_Item {
    unsigned char   bit_mask;
#       define      TAICancelledEUTRA_Item_iE_Extensions_present 0x80
    TAI             tAI;
    struct CancelledCellsInTAI_EUTRA_ *cancelledCellsInTAI_EUTRA;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                              * TAICancelledEUTRA_Item_iE_Extensions_present if
                              * present */
} TAICancelledEUTRA_Item;

typedef struct TAICancelledEUTRA_ {
    struct TAICancelledEUTRA_ *next;
    TAICancelledEUTRA_Item value;
} *TAICancelledEUTRA;

typedef struct TAICancelledNR_Item {
    unsigned char   bit_mask;
#       define      TAICancelledNR_Item_iE_Extensions_present 0x80
    TAI             tAI;
    struct CancelledCellsInTAI_NR_ *cancelledCellsInTAI_NR;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                 * TAICancelledNR_Item_iE_Extensions_present if
                                 * present */
} TAICancelledNR_Item;

typedef struct TAICancelledNR_ {
    struct TAICancelledNR_ *next;
    TAICancelledNR_Item value;
} *TAICancelledNR;

typedef struct TAIListForInactiveItem {
    unsigned char   bit_mask;
#       define      TAIListForInactiveItem_iE_Extensions_present 0x80
    TAI             tAI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                              * TAIListForInactiveItem_iE_Extensions_present if
                              * present */
} TAIListForInactiveItem;

typedef struct TAIListForInactive_ {
    struct TAIListForInactive_ *next;
    TAIListForInactiveItem value;
} *TAIListForInactive;

typedef struct TAIListForPagingItem {
    unsigned char   bit_mask;
#       define      TAIListForPagingItem_iE_Extensions_present 0x80
    TAI             tAI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                * TAIListForPagingItem_iE_Extensions_present if
                                * present */
} TAIListForPagingItem;

typedef struct TAIListForPaging_ {
    struct TAIListForPaging_ *next;
    TAIListForPagingItem value;
} *TAIListForPaging;

typedef struct TAIListForRestart_ {
    struct TAIListForRestart_ *next;
    TAI             value;
} *TAIListForRestart;

typedef struct TAIListForWarning_ {
    struct TAIListForWarning_ *next;
    TAI             value;
} *TAIListForWarning;

typedef struct TargeteNB_ID {
    unsigned char   bit_mask;
#       define      TargeteNB_ID_iE_Extensions_present 0x80
    GlobalNgENB_ID  globalENB_ID;
    EPS_TAI         selected_EPS_TAI;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * TargeteNB_ID_iE_Extensions_present if
                                      * present */
} TargeteNB_ID;

typedef struct TargetID {
    unsigned short  choice;
#       define      targetRANNodeID_chosen 1
#       define      targeteNB_ID_chosen 2
#       define      TargetID_choice_Extensions_chosen 3
    union _union {
        TargetRANNodeID targetRANNodeID;  /* to choose, set choice to
                                           * targetRANNodeID_chosen */
        TargeteNB_ID    targeteNB_ID;  /* to choose, set choice to
                                        * targeteNB_ID_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                      * to TargetID_choice_Extensions_chosen */
    } u;
} TargetID;

typedef struct TargetToSource_TransparentContainer {
    unsigned int    length;
    unsigned char   *value;
} TargetToSource_TransparentContainer;

// This IE includes a transparent container from the target RAN node to the source RAN node. 
// The octets of the OCTET STRING are encoded according to the specifications of the target system.
typedef enum TimerApproachForGUAMIRemoval {
    apply_timer = 0
} TimerApproachForGUAMIRemoval;

typedef struct TimeStamp {
    unsigned short  length;
    unsigned char   value[4];
} TimeStamp;

typedef struct TNLAssociationItem {
    unsigned char   bit_mask;
#       define      TNLAssociationItem_iE_Extensions_present 0x80
    CPTransportLayerInformation tNLAssociationAddress;
    Cause           cause;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * TNLAssociationItem_iE_Extensions_present
                                      * if present */
} TNLAssociationItem;

typedef struct TNLAssociationList_ {
    struct TNLAssociationList_ *next;
    TNLAssociationItem value;
} *TNLAssociationList;

typedef enum TraceDepth {
    minimum = 0,
    medium = 1,
    maximum = 2,
    minimumWithoutVendorSpecificExtension = 3,
    mediumWithoutVendorSpecificExtension = 4,
    maximumWithoutVendorSpecificExtension = 5
} TraceDepth;

typedef struct TraceActivation {
    unsigned char   bit_mask;
#       define      TraceActivation_iE_Extensions_present 0x80
    NGRANTraceID    nGRANTraceID;
    InterfacesToTrace interfacesToTrace;
    TraceDepth      traceDepth;
    TransportLayerAddress traceCollectionEntityIPAddress;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * TraceActivation_iE_Extensions_present if
                                      * present */
} TraceActivation;

// U
typedef struct UEAggregateMaximumBitRate {
    unsigned char   bit_mask;
#       define      UEAggregateMaximumBitRate_iE_Extensions_present 0x80
    BitRate         uEAggregateMaximumBitRateDL;
    BitRate         uEAggregateMaximumBitRateUL;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                           * UEAggregateMaximumBitRate_iE_Extensions_present if
                           * present */
} UEAggregateMaximumBitRate;

typedef struct UE_associatedLogicalNG_connectionItem {
    unsigned char   bit_mask;
#       define      aMF_UE_NGAP_ID_present 0x80
#       define      rAN_UE_NGAP_ID_present 0x40
#       define      UE_associatedLogicalNG_connectionItem_iE_Extensions_present 0x20
    AMF_UE_NGAP_ID  aMF_UE_NGAP_ID;  /* optional; set in bit_mask
                                      * aMF_UE_NGAP_ID_present if present */
    RAN_UE_NGAP_ID  rAN_UE_NGAP_ID;  /* optional; set in bit_mask
                                      * rAN_UE_NGAP_ID_present if present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
               * UE_associatedLogicalNG_connectionItem_iE_Extensions_present if
               * present */
} UE_associatedLogicalNG_connectionItem;

typedef struct UE_associatedLogicalNG_connectionList_ {
    struct UE_associatedLogicalNG_connectionList_ *next;
    UE_associatedLogicalNG_connectionItem value;
} *UE_associatedLogicalNG_connectionList;

typedef struct UE_NGAP_ID_pair {
    unsigned char   bit_mask;
#       define      UE_NGAP_ID_pair_iE_Extensions_present 0x80
    AMF_UE_NGAP_ID  aMF_UE_NGAP_ID;
    RAN_UE_NGAP_ID  rAN_UE_NGAP_ID;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * UE_NGAP_ID_pair_iE_Extensions_present if
                                      * present */
} UE_NGAP_ID_pair;

typedef struct UE_NGAP_IDs {
    unsigned short  choice;
#       define      uE_NGAP_ID_pair_chosen 1
#       define      aMF_UE_NGAP_ID_chosen 2
#       define      UE_NGAP_IDs_choice_Extensions_chosen 3
    union _union {
        UE_NGAP_ID_pair uE_NGAP_ID_pair;  /* to choose, set choice to
                                           * uE_NGAP_ID_pair_chosen */
        AMF_UE_NGAP_ID  aMF_UE_NGAP_ID;  /* to choose, set choice to
                                          * aMF_UE_NGAP_ID_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to UE_NGAP_IDs_choice_Extensions_chosen */
    } u;
} UE_NGAP_IDs;

typedef struct UEPagingIdentity {
    unsigned short  choice;
#       define      fiveG_S_TMSI_chosen 1
#       define      UEPagingIdentity_choice_Extensions_chosen 2
    union _union {
        FiveG_S_TMSI    fiveG_S_TMSI;  /* to choose, set choice to
                                        * fiveG_S_TMSI_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to
                                 * UEPagingIdentity_choice_Extensions_chosen */
    } u;
} UEPagingIdentity;

typedef enum UEPresence {
    in = 0,
    out = 1,
    unknown = 2
} UEPresence;

typedef struct UEPresenceInAreaOfInterestItem {
    unsigned char   bit_mask;
#       define      UEPresenceInAreaOfInterestItem_iE_Extensions_present 0x80
    LocationReportingReferenceID locationReportingReferenceID;
    UEPresence      uEPresence;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                      * UEPresenceInAreaOfInterestItem_iE_Extensions_present if
                      * present */
} UEPresenceInAreaOfInterestItem;

typedef struct UEPresenceInAreaOfInterestList_ {
    struct UEPresenceInAreaOfInterestList_ *next;
    UEPresenceInAreaOfInterestItem value;
} *UEPresenceInAreaOfInterestList;

typedef struct UERadioCapability {
    unsigned int    length;
    unsigned char   *value;
} UERadioCapability;

typedef struct UERadioCapabilityForPagingOfNR {
    unsigned int    length;
    unsigned char   *value;
} UERadioCapabilityForPagingOfNR;

typedef struct UERadioCapabilityForPagingOfEUTRA {
    unsigned int    length;
    unsigned char   *value;
} UERadioCapabilityForPagingOfEUTRA;

typedef struct UERadioCapabilityForPaging {
    unsigned char   bit_mask;
#       define      uERadioCapabilityForPagingOfNR_present 0x80
#       define      uERadioCapabilityForPagingOfEUTRA_present 0x40
#       define      UERadioCapabilityForPaging_iE_Extensions_present 0x20
    UERadioCapabilityForPagingOfNR uERadioCapabilityForPagingOfNR;  
                                  /* optional; set in bit_mask
                                   * uERadioCapabilityForPagingOfNR_present if
                                   * present */
    UERadioCapabilityForPagingOfEUTRA uERadioCapabilityForPagingOfEUTRA;  
                                        /* optional; set in bit_mask
                                 * uERadioCapabilityForPagingOfEUTRA_present if
                                 * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                          * UERadioCapabilityForPaging_iE_Extensions_present if
                          * present */
} UERadioCapabilityForPaging;

typedef struct UESecurityCapabilities {
    unsigned char   bit_mask;
#       define      UESecurityCapabilities_iE_Extensions_present 0x80
    NRencryptionAlgorithms nRencryptionAlgorithms;
    NRintegrityProtectionAlgorithms nRintegrityProtectionAlgorithms;
    EUTRAencryptionAlgorithms eUTRAencryptionAlgorithms;
    EUTRAintegrityProtectionAlgorithms eUTRAintegrityProtectionAlgorithms;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                              * UESecurityCapabilities_iE_Extensions_present if
                              * present */
} UESecurityCapabilities;

typedef struct UL_NGU_UP_TNLModifyItem {
    unsigned char   bit_mask;
#       define      UL_NGU_UP_TNLModifyItem_iE_Extensions_present 0x80
    UPTransportLayerInformation uL_NGU_UP_TNLInformation;
    UPTransportLayerInformation dL_NGU_UP_TNLInformation;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                             * UL_NGU_UP_TNLModifyItem_iE_Extensions_present if
                             * present */
} UL_NGU_UP_TNLModifyItem;

typedef struct UL_NGU_UP_TNLModifyList_ {
    struct UL_NGU_UP_TNLModifyList_ *next;
    UL_NGU_UP_TNLModifyItem value;
} *UL_NGU_UP_TNLModifyList;

typedef struct UnavailableGUAMIItem {
    unsigned char   bit_mask;
#       define      timerApproachForGUAMIRemoval_present 0x80
#       define      UnavailableGUAMIItem_backupAMFName_present 0x40
#       define      UnavailableGUAMIItem_iE_Extensions_present 0x20
    GUAMI           gUAMI;
    TimerApproachForGUAMIRemoval timerApproachForGUAMIRemoval;  /* optional; set
                                   * in bit_mask
                                   * timerApproachForGUAMIRemoval_present if
                                   * present */
    AMFName         backupAMFName;  /* optional; set in bit_mask
                                * UnavailableGUAMIItem_backupAMFName_present if
                                * present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                * UnavailableGUAMIItem_iE_Extensions_present if
                                * present */
} UnavailableGUAMIItem;

typedef struct UnavailableGUAMIList_ {
    struct UnavailableGUAMIList_ *next;
    UnavailableGUAMIItem value;
} *UnavailableGUAMIList;

typedef struct UPTransportLayerInformationItem {
    unsigned char   bit_mask;
#       define      UPTransportLayerInformationItem_iE_Extensions_present 0x80
    UPTransportLayerInformation nGU_UP_TNLInformation;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                     * UPTransportLayerInformationItem_iE_Extensions_present if
                     * present */
} UPTransportLayerInformationItem;

typedef struct UPTransportLayerInformationList_ {
    struct UPTransportLayerInformationList_ *next;
    UPTransportLayerInformationItem value;
} *UPTransportLayerInformationList;

typedef struct UPTransportLayerInformationPairItem {
    unsigned char   bit_mask;
#       define      UPTransportLayerInformationPairItem_iE_Extensions_present 0x80
    UPTransportLayerInformation uL_NGU_UP_TNLInformation;
    UPTransportLayerInformation dL_NGU_UP_TNLInformation;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                 * UPTransportLayerInformationPairItem_iE_Extensions_present if
                 * present */
} UPTransportLayerInformationPairItem;

typedef struct UPTransportLayerInformationPairList_ {
    struct UPTransportLayerInformationPairList_ *next;
    UPTransportLayerInformationPairItem value;
} *UPTransportLayerInformationPairList;

enum UserLocationInformationEUTRA_ExtIEs_Extension_PDUs {
    PDU_UserLocationInformationEUTRA_ExtIEs_Extension_UNKNOWN = 0,

    PDU_UserLocationInformationEUTRA_ExtIEs_Extension_NGRAN_CGI = NGRAN_CGI_PDU,

    PDU_UserLocationInformationEUTRA_ExtIEs_Extension_MAX = INT_MAX
};

union UserLocationInformationEUTRA_ExtIEs_Extension_union {
    NGRAN_CGI       *pdu_NGRAN_CGI;  
              /* PDU_UserLocationInformationEUTRA_ExtIEs_Extension_NGRAN_CGI */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct UserLocationInformationEUTRA_ExtIEs_Extension {
    UserLocationInformationEUTRA_ExtIEs_Extension_PDUs pduNum;
    OssBuf          encoded;
    UserLocationInformationEUTRA_ExtIEs_Extension_union decoded;
} UserLocationInformationEUTRA_ExtIEs_Extension;

typedef struct _seq227 {
    ProtocolExtensionID id;
    Criticality     criticality;
    UserLocationInformationEUTRA_ExtIEs_Extension extensionValue;
} _seq227;

typedef struct _seqof181_ {
    struct _seqof181_ *next;
    _seq227         value;
} *_seqof181;

typedef struct UserLocationInformationEUTRA {
    unsigned char   bit_mask;
#       define      UserLocationInformationEUTRA_timeStamp_present 0x80
#       define      UserLocationInformationEUTRA_iE_Extensions_present 0x40
    EUTRA_CGI       eUTRA_CGI;
    TAI             tAI;
    TimeStamp       timeStamp;  /* optional; set in bit_mask
                            * UserLocationInformationEUTRA_timeStamp_present if
                            * present */
    struct _seqof181_ *iE_Extensions;  /* optional; set in bit_mask
                        * UserLocationInformationEUTRA_iE_Extensions_present if
                        * present */
} UserLocationInformationEUTRA;

typedef struct UserLocationInformationNR {
    unsigned char   bit_mask;
#       define      UserLocationInformationNR_timeStamp_present 0x80
#       define      UserLocationInformationNR_iE_Extensions_present 0x40
    NR_CGI          nR_CGI;
    TAI             tAI;
    TimeStamp       timeStamp;  /* optional; set in bit_mask
                                 * UserLocationInformationNR_timeStamp_present
                                 * if present */
    struct _seqof181_ *iE_Extensions;  /* optional; set in bit_mask
                           * UserLocationInformationNR_iE_Extensions_present if
                           * present */
} UserLocationInformationNR;

typedef struct UserLocationInformationN3IWF {
    unsigned char   bit_mask;
#       define      UserLocationInformationN3IWF_iE_Extensions_present 0x80
    TransportLayerAddress iPAddress;
    PortNumber      portNumber;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                        * UserLocationInformationN3IWF_iE_Extensions_present if
                        * present */
} UserLocationInformationN3IWF;

typedef struct UserLocationInformation {
    unsigned short  choice;
#       define      userLocationInformationEUTRA_chosen 1
#       define      userLocationInformationNR_chosen 2
#       define      userLocationInformationN3IWF_chosen 3
#       define      UserLocationInformation_choice_Extensions_chosen 4
    union _union {
        UserLocationInformationEUTRA userLocationInformationEUTRA;  /* to
                                   * choose, set choice to
                                   * userLocationInformationEUTRA_chosen */
        UserLocationInformationNR userLocationInformationNR;  /* to choose, set
                                   * choice to
                                   * userLocationInformationNR_chosen */
        UserLocationInformationN3IWF userLocationInformationN3IWF;  /* to
                                   * choose, set choice to
                                   * userLocationInformationN3IWF_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to
                          * UserLocationInformation_choice_Extensions_chosen */
    } u;
} UserLocationInformation;

typedef struct _octet2 {
    unsigned short  length;
    unsigned char   value[4];
} _octet2;

typedef struct VolumeTimedReport_Item {
    unsigned char   bit_mask;
#       define      VolumeTimedReport_Item_iE_Extensions_present 0x80
    _octet2         startTimeStamp;
    _octet2         endTimeStamp;
    unsigned long   usageCountUL;
    unsigned long   usageCountDL;
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                              * VolumeTimedReport_Item_iE_Extensions_present if
                              * present */
} VolumeTimedReport_Item;

// V
typedef struct VolumeTimedReportList_ {
    struct VolumeTimedReportList_ *next;
    VolumeTimedReport_Item value;
} *VolumeTimedReportList;

// W
typedef struct WarningAreaCoordinates {
    unsigned short  length;
    unsigned char   *value;
} WarningAreaCoordinates;

typedef struct WarningAreaList {
    unsigned short  choice;
#       define      eUTRA_CGIListForWarning_chosen 1
#       define      nR_CGIListForWarning_chosen 2
#       define      tAIListForWarning_chosen 3
#       define      emergencyAreaIDList_chosen 4
#       define      WarningAreaList_choice_Extensions_chosen 5
    union _union {
        struct EUTRA_CGIListForWarning_ *eUTRA_CGIListForWarning;  /* to choose,
                                   * set choice to
                                   * eUTRA_CGIListForWarning_chosen */
        struct NR_CGIListForWarning_ *nR_CGIListForWarning;  /* to choose, set
                                     * choice to nR_CGIListForWarning_chosen */
        struct TAIListForWarning_ *tAIListForWarning;  /* to choose, set choice
                                               * to tAIListForWarning_chosen */
        struct EmergencyAreaIDList_ *emergencyAreaIDList;  /* to choose, set
                                      * choice to emergencyAreaIDList_chosen */
        ProtocolIE_SingleContainer choice_Extensions;  /* to choose, set choice
                                   * to
                                  * WarningAreaList_choice_Extensions_chosen */
    } u;
} WarningAreaList;

typedef struct WarningMessageContents {
    unsigned short  length;
    unsigned char   *value;
} WarningMessageContents;

typedef struct WarningSecurityInfo {
    unsigned short  length;
    unsigned char   value[50];
} WarningSecurityInfo;

typedef struct WarningType {
    unsigned short  length;
    unsigned char   value[2];
} WarningType;

typedef struct XnExtTLA_Item {
    unsigned char   bit_mask;
#       define      iPsecTLA_present 0x80
#       define      gTP_TLAs_present 0x40
#       define      XnExtTLA_Item_iE_Extensions_present 0x20
    TransportLayerAddress iPsecTLA;  /* optional; set in bit_mask
                                      * iPsecTLA_present if present */
    struct XnGTP_TLAs_ *gTP_TLAs;  /* optional; set in bit_mask gTP_TLAs_present
                                    * if present */
    struct _seqof1_ *iE_Extensions;  /* optional; set in bit_mask
                                      * XnExtTLA_Item_iE_Extensions_present if
                                      * present */
} XnExtTLA_Item;

// X
typedef struct XnExtTLAs_ {
    struct XnExtTLAs_ *next;
    XnExtTLA_Item   value;
} *XnExtTLAs;

typedef struct XnGTP_TLAs_ {
    struct XnGTP_TLAs_ *next;
    TransportLayerAddress value;
} *XnGTP_TLAs;

typedef struct XnTLAs_ {
    struct XnTLAs_  *next;
    TransportLayerAddress value;
} *XnTLAs;

enum PDUSessionResourceSetupRequestIEs_Value_PDUs {
    PDU_PDUSessionResourceSetupRequestIEs_Value_UNKNOWN = 0,

    PDU_PDUSessionResourceSetupRequestIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceSetupRequestIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceSetupRequestIEs_Value_RANPagingPriority = RANPagingPriority_PDU,
    PDU_PDUSessionResourceSetupRequestIEs_Value_NAS_PDU = NAS_PDU_PDU,
    PDU_PDUSessionResourceSetupRequestIEs_Value_PDUSessionResourceSetupListSUReq = PDUSessionResourceSetupListSUReq_PDU,
    PDU_PDUSessionResourceSetupRequestIEs_Value_UEAggregateMaximumBitRate = UEAggregateMaximumBitRate_PDU,

    PDU_PDUSessionResourceSetupRequestIEs_Value_MAX = INT_MAX
};

union PDUSessionResourceSetupRequestIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
               /* PDU_PDUSessionResourceSetupRequestIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
               /* PDU_PDUSessionResourceSetupRequestIEs_Value_RAN_UE_NGAP_ID */
    RANPagingPriority *pdu_RANPagingPriority;  
            /* PDU_PDUSessionResourceSetupRequestIEs_Value_RANPagingPriority */
    NAS_PDU         *pdu_NAS_PDU;  
                      /* PDU_PDUSessionResourceSetupRequestIEs_Value_NAS_PDU */
    struct PDUSessionResourceSetupListSUReq_ **pdu_PDUSessionResourceSetupListSUReq;                                    /* PDU_PDUSessionResourceSetupRequestIEs_Value_PDUSessionResourceSetupListSUReq */
    UEAggregateMaximumBitRate *pdu_UEAggregateMaximumBitRate;  
    /* PDU_PDUSessionResourceSetupRequestIEs_Value_UEAggregateMaximumBitRate */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PDUSessionResourceSetupRequestIEs_Value {
    PDUSessionResourceSetupRequestIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PDUSessionResourceSetupRequestIEs_Value_union decoded;
} PDUSessionResourceSetupRequestIEs_Value;

// **************************************************************
//
// IE parameter types from other modules.
//
// **************************************************************
// **************************************************************
//
// PDU SESSION MANAGEMENT ELEMENTARY PROCEDURES
//
// **************************************************************

// **************************************************************
//
// PDU Session Resource Setup Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// PDU SESSION RESOURCE SETUP REQUEST
//
// **************************************************************
typedef struct PDUSessionResourceSetupRequest {
    struct _seqof186 {
        struct _seqof186 *next;
        struct _seq234 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PDUSessionResourceSetupRequestIEs_Value value;
        } value;
    } *protocolIEs;
} PDUSessionResourceSetupRequest;

enum PDUSessionResourceSetupResponseIEs_Value_PDUs {
    PDU_PDUSessionResourceSetupResponseIEs_Value_UNKNOWN = 0,

    PDU_PDUSessionResourceSetupResponseIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceSetupResponseIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceSetupResponseIEs_Value_PDUSessionResourceSetupListSURes = PDUSessionResourceSetupListSURes_PDU,
    PDU_PDUSessionResourceSetupResponseIEs_Value_PDUSessionResourceFailedToSetupListSURes = PDUSessionResourceFailedToSetupListSURes_PDU,
    PDU_PDUSessionResourceSetupResponseIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_PDUSessionResourceSetupResponseIEs_Value_MAX = INT_MAX
};

union PDUSessionResourceSetupResponseIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
              /* PDU_PDUSessionResourceSetupResponseIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
              /* PDU_PDUSessionResourceSetupResponseIEs_Value_RAN_UE_NGAP_ID */
    struct PDUSessionResourceSetupListSURes_ **pdu_PDUSessionResourceSetupListSURes;                                    /* PDU_PDUSessionResourceSetupResponseIEs_Value_PDUSessionResourceSetupListSURes */
    struct PDUSessionResourceFailedToSetupListSURes_ **pdu_PDUSessionResourceFailedToSetupListSURes;                    /* PDU_PDUSessionResourceSetupResponseIEs_Value_PDUSessionResourceFailedToSetupListSURes */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
      /* PDU_PDUSessionResourceSetupResponseIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PDUSessionResourceSetupResponseIEs_Value {
    PDUSessionResourceSetupResponseIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PDUSessionResourceSetupResponseIEs_Value_union decoded;
} PDUSessionResourceSetupResponseIEs_Value;

// **************************************************************
//
// PDU SESSION RESOURCE SETUP RESPONSE
//
// **************************************************************
typedef struct PDUSessionResourceSetupResponse {
    struct _seqof187 {
        struct _seqof187 *next;
        struct _seq235 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PDUSessionResourceSetupResponseIEs_Value value;
        } value;
    } *protocolIEs;
} PDUSessionResourceSetupResponse;

enum PDUSessionResourceReleaseCommandIEs_Value_PDUs {
    PDU_PDUSessionResourceReleaseCommandIEs_Value_UNKNOWN = 0,

    PDU_PDUSessionResourceReleaseCommandIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceReleaseCommandIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceReleaseCommandIEs_Value_RANPagingPriority = RANPagingPriority_PDU,
    PDU_PDUSessionResourceReleaseCommandIEs_Value_NAS_PDU = NAS_PDU_PDU,
    PDU_PDUSessionResourceReleaseCommandIEs_Value_PDUSessionResourceToReleaseListRelCmd = PDUSessionResourceToReleaseListRelCmd_PDU,

    PDU_PDUSessionResourceReleaseCommandIEs_Value_MAX = INT_MAX
};

union PDUSessionResourceReleaseCommandIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
             /* PDU_PDUSessionResourceReleaseCommandIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
             /* PDU_PDUSessionResourceReleaseCommandIEs_Value_RAN_UE_NGAP_ID */
    RANPagingPriority *pdu_RANPagingPriority;  
          /* PDU_PDUSessionResourceReleaseCommandIEs_Value_RANPagingPriority */
    NAS_PDU         *pdu_NAS_PDU;  
                    /* PDU_PDUSessionResourceReleaseCommandIEs_Value_NAS_PDU */
    struct PDUSessionResourceToReleaseListRelCmd_ **pdu_PDUSessionResourceToReleaseListRelCmd;                          /* PDU_PDUSessionResourceReleaseCommandIEs_Value_PDUSessionResourceToReleaseListRelCmd */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PDUSessionResourceReleaseCommandIEs_Value {
    PDUSessionResourceReleaseCommandIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PDUSessionResourceReleaseCommandIEs_Value_union decoded;
} PDUSessionResourceReleaseCommandIEs_Value;

// **************************************************************
//
// PDU Session Resource Release Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// PDU SESSION RESOURCE RELEASE COMMAND
//
// **************************************************************
typedef struct PDUSessionResourceReleaseCommand {
    struct _seqof188 {
        struct _seqof188 *next;
        struct _seq236 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PDUSessionResourceReleaseCommandIEs_Value value;
        } value;
    } *protocolIEs;
} PDUSessionResourceReleaseCommand;

enum PDUSessionResourceReleaseResponseIEs_Value_PDUs {
    PDU_PDUSessionResourceReleaseResponseIEs_Value_UNKNOWN = 0,

    PDU_PDUSessionResourceReleaseResponseIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceReleaseResponseIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceReleaseResponseIEs_Value_PDUSessionResourceReleasedListRelRes = PDUSessionResourceReleasedListRelRes_PDU,
    PDU_PDUSessionResourceReleaseResponseIEs_Value_UserLocationInformation = UserLocationInformation_PDU,
    PDU_PDUSessionResourceReleaseResponseIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_PDUSessionResourceReleaseResponseIEs_Value_MAX = INT_MAX
};

union PDUSessionResourceReleaseResponseIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
            /* PDU_PDUSessionResourceReleaseResponseIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
            /* PDU_PDUSessionResourceReleaseResponseIEs_Value_RAN_UE_NGAP_ID */
    struct PDUSessionResourceReleasedListRelRes_ **pdu_PDUSessionResourceReleasedListRelRes;                            /* PDU_PDUSessionResourceReleaseResponseIEs_Value_PDUSessionResourceReleasedListRelRes */
    UserLocationInformation *pdu_UserLocationInformation;  
   /* PDU_PDUSessionResourceReleaseResponseIEs_Value_UserLocationInformation */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
    /* PDU_PDUSessionResourceReleaseResponseIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PDUSessionResourceReleaseResponseIEs_Value {
    PDUSessionResourceReleaseResponseIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PDUSessionResourceReleaseResponseIEs_Value_union decoded;
} PDUSessionResourceReleaseResponseIEs_Value;

// **************************************************************
//
// PDU SESSION RESOURCE RELEASE RESPONSE
//
// **************************************************************
typedef struct PDUSessionResourceReleaseResponse {
    struct _seqof189 {
        struct _seqof189 *next;
        struct _seq237 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PDUSessionResourceReleaseResponseIEs_Value value;
        } value;
    } *protocolIEs;
} PDUSessionResourceReleaseResponse;

enum PDUSessionResourceModifyRequestIEs_Value_PDUs {
    PDU_PDUSessionResourceModifyRequestIEs_Value_UNKNOWN = 0,

    PDU_PDUSessionResourceModifyRequestIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceModifyRequestIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceModifyRequestIEs_Value_RANPagingPriority = RANPagingPriority_PDU,
    PDU_PDUSessionResourceModifyRequestIEs_Value_PDUSessionResourceModifyListModReq = PDUSessionResourceModifyListModReq_PDU,

    PDU_PDUSessionResourceModifyRequestIEs_Value_MAX = INT_MAX
};

union PDUSessionResourceModifyRequestIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
              /* PDU_PDUSessionResourceModifyRequestIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
              /* PDU_PDUSessionResourceModifyRequestIEs_Value_RAN_UE_NGAP_ID */
    RANPagingPriority *pdu_RANPagingPriority;  
           /* PDU_PDUSessionResourceModifyRequestIEs_Value_RANPagingPriority */
    struct PDUSessionResourceModifyListModReq_ **pdu_PDUSessionResourceModifyListModReq;                                /* PDU_PDUSessionResourceModifyRequestIEs_Value_PDUSessionResourceModifyListModReq */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PDUSessionResourceModifyRequestIEs_Value {
    PDUSessionResourceModifyRequestIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PDUSessionResourceModifyRequestIEs_Value_union decoded;
} PDUSessionResourceModifyRequestIEs_Value;

// **************************************************************
//
// PDU Session Resource Modify Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// PDU SESSION RESOURCE MODIFY REQUEST
//
// **************************************************************
typedef struct PDUSessionResourceModifyRequest {
    struct _seqof190 {
        struct _seqof190 *next;
        struct _seq238 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PDUSessionResourceModifyRequestIEs_Value value;
        } value;
    } *protocolIEs;
} PDUSessionResourceModifyRequest;

enum PDUSessionResourceModifyResponseIEs_Value_PDUs {
    PDU_PDUSessionResourceModifyResponseIEs_Value_UNKNOWN = 0,

    PDU_PDUSessionResourceModifyResponseIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceModifyResponseIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceModifyResponseIEs_Value_PDUSessionResourceModifyListModRes = PDUSessionResourceModifyListModRes_PDU,
    PDU_PDUSessionResourceModifyResponseIEs_Value_PDUSessionResourceFailedToModifyListModRes = PDUSessionResourceFailedToModifyListModRes_PDU,
    PDU_PDUSessionResourceModifyResponseIEs_Value_UserLocationInformation = UserLocationInformation_PDU,
    PDU_PDUSessionResourceModifyResponseIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_PDUSessionResourceModifyResponseIEs_Value_MAX = INT_MAX
};

union PDUSessionResourceModifyResponseIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
             /* PDU_PDUSessionResourceModifyResponseIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
             /* PDU_PDUSessionResourceModifyResponseIEs_Value_RAN_UE_NGAP_ID */
    struct PDUSessionResourceModifyListModRes_ **pdu_PDUSessionResourceModifyListModRes;                                /* PDU_PDUSessionResourceModifyResponseIEs_Value_PDUSessionResourceModifyListModRes */
    struct PDUSessionResourceFailedToModifyListModRes_ **pdu_PDUSessionResourceFailedToModifyListModRes;                /* PDU_PDUSessionResourceModifyResponseIEs_Value_PDUSessionResourceFailedToModifyListModRes */
    UserLocationInformation *pdu_UserLocationInformation;  
    /* PDU_PDUSessionResourceModifyResponseIEs_Value_UserLocationInformation */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
     /* PDU_PDUSessionResourceModifyResponseIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PDUSessionResourceModifyResponseIEs_Value {
    PDUSessionResourceModifyResponseIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PDUSessionResourceModifyResponseIEs_Value_union decoded;
} PDUSessionResourceModifyResponseIEs_Value;

// **************************************************************
//
// PDU SESSION RESOURCE MODIFY RESPONSE
//
// **************************************************************
typedef struct PDUSessionResourceModifyResponse {
    struct _seqof191 {
        struct _seqof191 *next;
        struct _seq239 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PDUSessionResourceModifyResponseIEs_Value value;
        } value;
    } *protocolIEs;
} PDUSessionResourceModifyResponse;

enum PDUSessionResourceNotifyIEs_Value_PDUs {
    PDU_PDUSessionResourceNotifyIEs_Value_UNKNOWN = 0,

    PDU_PDUSessionResourceNotifyIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceNotifyIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceNotifyIEs_Value_PDUSessionResourceNotifyList = PDUSessionResourceNotifyList_PDU,
    PDU_PDUSessionResourceNotifyIEs_Value_PDUSessionResourceReleasedListNot = PDUSessionResourceReleasedListNot_PDU,
    PDU_PDUSessionResourceNotifyIEs_Value_UserLocationInformation = UserLocationInformation_PDU,

    PDU_PDUSessionResourceNotifyIEs_Value_MAX = INT_MAX
};

union PDUSessionResourceNotifyIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                     /* PDU_PDUSessionResourceNotifyIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                     /* PDU_PDUSessionResourceNotifyIEs_Value_RAN_UE_NGAP_ID */
    struct PDUSessionResourceNotifyList_ **pdu_PDUSessionResourceNotifyList;  
                                        /* PDU_PDUSessionResourceNotifyIEs_Value_PDUSessionResourceNotifyList */
    struct PDUSessionResourceReleasedListNot_ **pdu_PDUSessionResourceReleasedListNot;                                  /* PDU_PDUSessionResourceNotifyIEs_Value_PDUSessionResourceReleasedListNot */
    UserLocationInformation *pdu_UserLocationInformation;  
            /* PDU_PDUSessionResourceNotifyIEs_Value_UserLocationInformation */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PDUSessionResourceNotifyIEs_Value {
    PDUSessionResourceNotifyIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PDUSessionResourceNotifyIEs_Value_union decoded;
} PDUSessionResourceNotifyIEs_Value;

// **************************************************************
//
// PDU Session Resource Notify Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// PDU SESSION RESOURCE NOTIFY
//
// **************************************************************
typedef struct PDUSessionResourceNotify {
    struct _seqof192 {
        struct _seqof192 *next;
        struct _seq240 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PDUSessionResourceNotifyIEs_Value value;
        } value;
    } *protocolIEs;
} PDUSessionResourceNotify;

enum PDUSessionResourceModifyIndicationIEs_Value_PDUs {
    PDU_PDUSessionResourceModifyIndicationIEs_Value_UNKNOWN = 0,

    PDU_PDUSessionResourceModifyIndicationIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceModifyIndicationIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceModifyIndicationIEs_Value_PDUSessionResourceModifyListModInd = PDUSessionResourceModifyListModInd_PDU,

    PDU_PDUSessionResourceModifyIndicationIEs_Value_MAX = INT_MAX
};

union PDUSessionResourceModifyIndicationIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
           /* PDU_PDUSessionResourceModifyIndicationIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
           /* PDU_PDUSessionResourceModifyIndicationIEs_Value_RAN_UE_NGAP_ID */
    struct PDUSessionResourceModifyListModInd_ **pdu_PDUSessionResourceModifyListModInd;                                /* PDU_PDUSessionResourceModifyIndicationIEs_Value_PDUSessionResourceModifyListModInd */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PDUSessionResourceModifyIndicationIEs_Value {
    PDUSessionResourceModifyIndicationIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PDUSessionResourceModifyIndicationIEs_Value_union decoded;
} PDUSessionResourceModifyIndicationIEs_Value;

// **************************************************************
//
// PDU Session Resource Modify Indication Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// PDU SESSION RESOURCE MODIFY INDICATION
//
// **************************************************************
typedef struct PDUSessionResourceModifyIndication {
    struct _seqof193 {
        struct _seqof193 *next;
        struct _seq241 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PDUSessionResourceModifyIndicationIEs_Value value;
        } value;
    } *protocolIEs;
} PDUSessionResourceModifyIndication;

enum PDUSessionResourceModifyConfirmIEs_Value_PDUs {
    PDU_PDUSessionResourceModifyConfirmIEs_Value_UNKNOWN = 0,

    PDU_PDUSessionResourceModifyConfirmIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceModifyConfirmIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_PDUSessionResourceModifyConfirmIEs_Value_PDUSessionResourceModifyListModCfm = PDUSessionResourceModifyListModCfm_PDU,
    PDU_PDUSessionResourceModifyConfirmIEs_Value_PDUSessionResourceFailedToModifyListModCfm = PDUSessionResourceFailedToModifyListModCfm_PDU,
    PDU_PDUSessionResourceModifyConfirmIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_PDUSessionResourceModifyConfirmIEs_Value_MAX = INT_MAX
};

union PDUSessionResourceModifyConfirmIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
              /* PDU_PDUSessionResourceModifyConfirmIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
              /* PDU_PDUSessionResourceModifyConfirmIEs_Value_RAN_UE_NGAP_ID */
    struct PDUSessionResourceModifyListModCfm_ **pdu_PDUSessionResourceModifyListModCfm;                                /* PDU_PDUSessionResourceModifyConfirmIEs_Value_PDUSessionResourceModifyListModCfm */
    struct PDUSessionResourceFailedToModifyListModCfm_ **pdu_PDUSessionResourceFailedToModifyListModCfm;                /* PDU_PDUSessionResourceModifyConfirmIEs_Value_PDUSessionResourceFailedToModifyListModCfm */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
      /* PDU_PDUSessionResourceModifyConfirmIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PDUSessionResourceModifyConfirmIEs_Value {
    PDUSessionResourceModifyConfirmIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PDUSessionResourceModifyConfirmIEs_Value_union decoded;
} PDUSessionResourceModifyConfirmIEs_Value;

// **************************************************************
//
// PDU SESSION RESOURCE MODIFY CONFIRM
//
// **************************************************************
typedef struct PDUSessionResourceModifyConfirm {
    struct _seqof194 {
        struct _seqof194 *next;
        struct _seq242 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PDUSessionResourceModifyConfirmIEs_Value value;
        } value;
    } *protocolIEs;
} PDUSessionResourceModifyConfirm;

enum InitialContextSetupRequestIEs_Value_PDUs {
    PDU_InitialContextSetupRequestIEs_Value_UNKNOWN = 0,

    PDU_InitialContextSetupRequestIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_InitialContextSetupRequestIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_InitialContextSetupRequestIEs_Value_AMFName = AMFName_PDU,
    PDU_InitialContextSetupRequestIEs_Value_UEAggregateMaximumBitRate = UEAggregateMaximumBitRate_PDU,
    PDU_InitialContextSetupRequestIEs_Value_CoreNetworkAssistanceInformationForInactive = CoreNetworkAssistanceInformationForInactive_PDU,
    PDU_InitialContextSetupRequestIEs_Value_GUAMI = GUAMI_PDU,
    PDU_InitialContextSetupRequestIEs_Value_PDUSessionResourceSetupListCxtReq = PDUSessionResourceSetupListCxtReq_PDU,
    PDU_InitialContextSetupRequestIEs_Value_AllowedNSSAI = AllowedNSSAI_PDU,
    PDU_InitialContextSetupRequestIEs_Value_UESecurityCapabilities = UESecurityCapabilities_PDU,
    PDU_InitialContextSetupRequestIEs_Value_SecurityKey = SecurityKey_PDU,
    PDU_InitialContextSetupRequestIEs_Value_TraceActivation = TraceActivation_PDU,
    PDU_InitialContextSetupRequestIEs_Value_MobilityRestrictionList = MobilityRestrictionList_PDU,
    PDU_InitialContextSetupRequestIEs_Value_UERadioCapability = UERadioCapability_PDU,
    PDU_InitialContextSetupRequestIEs_Value_IndexToRFSP = IndexToRFSP_PDU,
    PDU_InitialContextSetupRequestIEs_Value_MaskedIMEISV = MaskedIMEISV_PDU,
    PDU_InitialContextSetupRequestIEs_Value_NAS_PDU = NAS_PDU_PDU,
    PDU_InitialContextSetupRequestIEs_Value_EmergencyFallbackIndicator = EmergencyFallbackIndicator_PDU,
    PDU_InitialContextSetupRequestIEs_Value_RRCInactiveTransitionReportRequest = RRCInactiveTransitionReportRequest_PDU,
    PDU_InitialContextSetupRequestIEs_Value_UERadioCapabilityForPaging = UERadioCapabilityForPaging_PDU,
    PDU_InitialContextSetupRequestIEs_Value_RedirectionVoiceFallback = RedirectionVoiceFallback_PDU,
    PDU_InitialContextSetupRequestIEs_Value_LocationReportingRequestType = LocationReportingRequestType_PDU,
    PDU_InitialContextSetupRequestIEs_Value_CNAssistedRANTuning = CNAssistedRANTuning_PDU,

    PDU_InitialContextSetupRequestIEs_Value_MAX = INT_MAX
};

union InitialContextSetupRequestIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                   /* PDU_InitialContextSetupRequestIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                   /* PDU_InitialContextSetupRequestIEs_Value_RAN_UE_NGAP_ID */
    AMFName         *pdu_AMFName;  
                          /* PDU_InitialContextSetupRequestIEs_Value_AMFName */
    UEAggregateMaximumBitRate *pdu_UEAggregateMaximumBitRate;  
        /* PDU_InitialContextSetupRequestIEs_Value_UEAggregateMaximumBitRate */
    CoreNetworkAssistanceInformationForInactive *pdu_CoreNetworkAssistanceInformationForInactive;                       /* PDU_InitialContextSetupRequestIEs_Value_CoreNetworkAssistanceInformationForInactive */
    GUAMI           *pdu_GUAMI;  
                            /* PDU_InitialContextSetupRequestIEs_Value_GUAMI */
    struct PDUSessionResourceSetupListCxtReq_ **pdu_PDUSessionResourceSetupListCxtReq;                                  /* PDU_InitialContextSetupRequestIEs_Value_PDUSessionResourceSetupListCxtReq */
    struct AllowedNSSAI_ **pdu_AllowedNSSAI;  
                     /* PDU_InitialContextSetupRequestIEs_Value_AllowedNSSAI */
    UESecurityCapabilities *pdu_UESecurityCapabilities;  
           /* PDU_InitialContextSetupRequestIEs_Value_UESecurityCapabilities */
    SecurityKey     *pdu_SecurityKey;  
                      /* PDU_InitialContextSetupRequestIEs_Value_SecurityKey */
    TraceActivation *pdu_TraceActivation;  
                  /* PDU_InitialContextSetupRequestIEs_Value_TraceActivation */
    MobilityRestrictionList *pdu_MobilityRestrictionList;  
          /* PDU_InitialContextSetupRequestIEs_Value_MobilityRestrictionList */
    UERadioCapability *pdu_UERadioCapability;  
                /* PDU_InitialContextSetupRequestIEs_Value_UERadioCapability */
    IndexToRFSP     *pdu_IndexToRFSP;  
                      /* PDU_InitialContextSetupRequestIEs_Value_IndexToRFSP */
    MaskedIMEISV    *pdu_MaskedIMEISV;  
                     /* PDU_InitialContextSetupRequestIEs_Value_MaskedIMEISV */
    NAS_PDU         *pdu_NAS_PDU;  
                          /* PDU_InitialContextSetupRequestIEs_Value_NAS_PDU */
    EmergencyFallbackIndicator *pdu_EmergencyFallbackIndicator;  
       /* PDU_InitialContextSetupRequestIEs_Value_EmergencyFallbackIndicator */
    RRCInactiveTransitionReportRequest *pdu_RRCInactiveTransitionReportRequest;                                         /* PDU_InitialContextSetupRequestIEs_Value_RRCInactiveTransitionReportRequest */
    UERadioCapabilityForPaging *pdu_UERadioCapabilityForPaging;  
       /* PDU_InitialContextSetupRequestIEs_Value_UERadioCapabilityForPaging */
    RedirectionVoiceFallback *pdu_RedirectionVoiceFallback;  
         /* PDU_InitialContextSetupRequestIEs_Value_RedirectionVoiceFallback */
    LocationReportingRequestType *pdu_LocationReportingRequestType;  
     /* PDU_InitialContextSetupRequestIEs_Value_LocationReportingRequestType */
    CNAssistedRANTuning *pdu_CNAssistedRANTuning;  
              /* PDU_InitialContextSetupRequestIEs_Value_CNAssistedRANTuning */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct InitialContextSetupRequestIEs_Value {
    InitialContextSetupRequestIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    InitialContextSetupRequestIEs_Value_union decoded;
} InitialContextSetupRequestIEs_Value;

// **************************************************************
//
// UE CONTEXT MANAGEMENT ELEMENTARY PROCEDURES
//
// **************************************************************

// **************************************************************
//
// Initial Context Setup Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// INITIAL CONTEXT SETUP REQUEST
//
// **************************************************************
typedef struct InitialContextSetupRequest {
    struct _seqof195 {
        struct _seqof195 *next;
        struct _seq243 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            InitialContextSetupRequestIEs_Value value;
        } value;
    } *protocolIEs;
} InitialContextSetupRequest;

enum InitialContextSetupResponseIEs_Value_PDUs {
    PDU_InitialContextSetupResponseIEs_Value_UNKNOWN = 0,

    PDU_InitialContextSetupResponseIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_InitialContextSetupResponseIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_InitialContextSetupResponseIEs_Value_PDUSessionResourceSetupListCxtRes = PDUSessionResourceSetupListCxtRes_PDU,
    PDU_InitialContextSetupResponseIEs_Value_PDUSessionResourceFailedToSetupListCxtRes = PDUSessionResourceFailedToSetupListCxtRes_PDU,
    PDU_InitialContextSetupResponseIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_InitialContextSetupResponseIEs_Value_MAX = INT_MAX
};

union InitialContextSetupResponseIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                  /* PDU_InitialContextSetupResponseIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                  /* PDU_InitialContextSetupResponseIEs_Value_RAN_UE_NGAP_ID */
    struct PDUSessionResourceSetupListCxtRes_ **pdu_PDUSessionResourceSetupListCxtRes;                                  /* PDU_InitialContextSetupResponseIEs_Value_PDUSessionResourceSetupListCxtRes */
    struct PDUSessionResourceFailedToSetupListCxtRes_ **pdu_PDUSessionResourceFailedToSetupListCxtRes;                  /* PDU_InitialContextSetupResponseIEs_Value_PDUSessionResourceFailedToSetupListCxtRes */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
          /* PDU_InitialContextSetupResponseIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct InitialContextSetupResponseIEs_Value {
    InitialContextSetupResponseIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    InitialContextSetupResponseIEs_Value_union decoded;
} InitialContextSetupResponseIEs_Value;

// **************************************************************
//
// INITIAL CONTEXT SETUP RESPONSE
//
// **************************************************************
typedef struct InitialContextSetupResponse {
    struct _seqof196 {
        struct _seqof196 *next;
        struct _seq244 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            InitialContextSetupResponseIEs_Value value;
        } value;
    } *protocolIEs;
} InitialContextSetupResponse;

enum InitialContextSetupFailureIEs_Value_PDUs {
    PDU_InitialContextSetupFailureIEs_Value_UNKNOWN = 0,

    PDU_InitialContextSetupFailureIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_InitialContextSetupFailureIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_InitialContextSetupFailureIEs_Value_PDUSessionResourceFailedToSetupListCxtFail = PDUSessionResourceFailedToSetupListCxtFail_PDU,
    PDU_InitialContextSetupFailureIEs_Value_Cause = Cause_PDU,
    PDU_InitialContextSetupFailureIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_InitialContextSetupFailureIEs_Value_MAX = INT_MAX
};

union InitialContextSetupFailureIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                   /* PDU_InitialContextSetupFailureIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                   /* PDU_InitialContextSetupFailureIEs_Value_RAN_UE_NGAP_ID */
    struct PDUSessionResourceFailedToSetupListCxtFail_ **pdu_PDUSessionResourceFailedToSetupListCxtFail;                /* PDU_InitialContextSetupFailureIEs_Value_PDUSessionResourceFailedToSetupListCxtFail */
    Cause           *pdu_Cause;  
                            /* PDU_InitialContextSetupFailureIEs_Value_Cause */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
           /* PDU_InitialContextSetupFailureIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct InitialContextSetupFailureIEs_Value {
    InitialContextSetupFailureIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    InitialContextSetupFailureIEs_Value_union decoded;
} InitialContextSetupFailureIEs_Value;

// **************************************************************
//
// INITIAL CONTEXT SETUP FAILURE
//
// **************************************************************
typedef struct InitialContextSetupFailure {
    struct _seqof197 {
        struct _seqof197 *next;
        struct _seq245 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            InitialContextSetupFailureIEs_Value value;
        } value;
    } *protocolIEs;
} InitialContextSetupFailure;

enum UEContextReleaseRequest_IEs_Value_PDUs {
    PDU_UEContextReleaseRequest_IEs_Value_UNKNOWN = 0,

    PDU_UEContextReleaseRequest_IEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_UEContextReleaseRequest_IEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_UEContextReleaseRequest_IEs_Value_PDUSessionResourceListCxtRelReq = PDUSessionResourceListCxtRelReq_PDU,
    PDU_UEContextReleaseRequest_IEs_Value_Cause = Cause_PDU,

    PDU_UEContextReleaseRequest_IEs_Value_MAX = INT_MAX
};

union UEContextReleaseRequest_IEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                     /* PDU_UEContextReleaseRequest_IEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                     /* PDU_UEContextReleaseRequest_IEs_Value_RAN_UE_NGAP_ID */
    struct PDUSessionResourceListCxtRelReq_ **pdu_PDUSessionResourceListCxtRelReq;                                      /* PDU_UEContextReleaseRequest_IEs_Value_PDUSessionResourceListCxtRelReq */
    Cause           *pdu_Cause;  
                              /* PDU_UEContextReleaseRequest_IEs_Value_Cause */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct UEContextReleaseRequest_IEs_Value {
    UEContextReleaseRequest_IEs_Value_PDUs pduNum;
    OssBuf          encoded;
    UEContextReleaseRequest_IEs_Value_union decoded;
} UEContextReleaseRequest_IEs_Value;

// **************************************************************
//
// UE Context Release Request Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// UE CONTEXT RELEASE REQUEST
//
// **************************************************************
typedef struct UEContextReleaseRequest {
    struct _seqof198 {
        struct _seqof198 *next;
        struct _seq246 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            UEContextReleaseRequest_IEs_Value value;
        } value;
    } *protocolIEs;
} UEContextReleaseRequest;

enum UEContextReleaseCommand_IEs_Value_PDUs {
    PDU_UEContextReleaseCommand_IEs_Value_UNKNOWN = 0,

    PDU_UEContextReleaseCommand_IEs_Value_UE_NGAP_IDs = UE_NGAP_IDs_PDU,
    PDU_UEContextReleaseCommand_IEs_Value_Cause = Cause_PDU,

    PDU_UEContextReleaseCommand_IEs_Value_MAX = INT_MAX
};

union UEContextReleaseCommand_IEs_Value_union {
    UE_NGAP_IDs     *pdu_UE_NGAP_IDs;  
                        /* PDU_UEContextReleaseCommand_IEs_Value_UE_NGAP_IDs */
    Cause           *pdu_Cause;  
                              /* PDU_UEContextReleaseCommand_IEs_Value_Cause */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct UEContextReleaseCommand_IEs_Value {
    UEContextReleaseCommand_IEs_Value_PDUs pduNum;
    OssBuf          encoded;
    UEContextReleaseCommand_IEs_Value_union decoded;
} UEContextReleaseCommand_IEs_Value;

// **************************************************************
//
// UE Context Release Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// UE CONTEXT RELEASE COMMAND
//
// **************************************************************
typedef struct UEContextReleaseCommand {
    struct _seqof199 {
        struct _seqof199 *next;
        struct _seq247 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            UEContextReleaseCommand_IEs_Value value;
        } value;
    } *protocolIEs;
} UEContextReleaseCommand;

enum UEContextReleaseComplete_IEs_Value_PDUs {
    PDU_UEContextReleaseComplete_IEs_Value_UNKNOWN = 0,

    PDU_UEContextReleaseComplete_IEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_UEContextReleaseComplete_IEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_UEContextReleaseComplete_IEs_Value_UserLocationInformation = UserLocationInformation_PDU,
    PDU_UEContextReleaseComplete_IEs_Value_InfoOnRecommendedCellsAndRANNodesForPaging = InfoOnRecommendedCellsAndRANNodesForPaging_PDU,
    PDU_UEContextReleaseComplete_IEs_Value_PDUSessionResourceListCxtRelCpl = PDUSessionResourceListCxtRelCpl_PDU,
    PDU_UEContextReleaseComplete_IEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_UEContextReleaseComplete_IEs_Value_MAX = INT_MAX
};

union UEContextReleaseComplete_IEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                    /* PDU_UEContextReleaseComplete_IEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                    /* PDU_UEContextReleaseComplete_IEs_Value_RAN_UE_NGAP_ID */
    UserLocationInformation *pdu_UserLocationInformation;  
           /* PDU_UEContextReleaseComplete_IEs_Value_UserLocationInformation */
    InfoOnRecommendedCellsAndRANNodesForPaging *pdu_InfoOnRecommendedCellsAndRANNodesForPaging;                         /* PDU_UEContextReleaseComplete_IEs_Value_InfoOnRecommendedCellsAndRANNodesForPaging */
    struct PDUSessionResourceListCxtRelCpl_ **pdu_PDUSessionResourceListCxtRelCpl;                                      /* PDU_UEContextReleaseComplete_IEs_Value_PDUSessionResourceListCxtRelCpl */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
            /* PDU_UEContextReleaseComplete_IEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct UEContextReleaseComplete_IEs_Value {
    UEContextReleaseComplete_IEs_Value_PDUs pduNum;
    OssBuf          encoded;
    UEContextReleaseComplete_IEs_Value_union decoded;
} UEContextReleaseComplete_IEs_Value;

// **************************************************************
//
// UE CONTEXT RELEASE COMPLETE
//
// **************************************************************
typedef struct UEContextReleaseComplete {
    struct _seqof200 {
        struct _seqof200 *next;
        struct _seq248 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            UEContextReleaseComplete_IEs_Value value;
        } value;
    } *protocolIEs;
} UEContextReleaseComplete;

enum UEContextModificationRequestIEs_Value_PDUs {
    PDU_UEContextModificationRequestIEs_Value_UNKNOWN = 0,

    PDU_UEContextModificationRequestIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_UEContextModificationRequestIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_UEContextModificationRequestIEs_Value_RANPagingPriority = RANPagingPriority_PDU,
    PDU_UEContextModificationRequestIEs_Value_SecurityKey = SecurityKey_PDU,
    PDU_UEContextModificationRequestIEs_Value_IndexToRFSP = IndexToRFSP_PDU,
    PDU_UEContextModificationRequestIEs_Value_UEAggregateMaximumBitRate = UEAggregateMaximumBitRate_PDU,
    PDU_UEContextModificationRequestIEs_Value_UESecurityCapabilities = UESecurityCapabilities_PDU,
    PDU_UEContextModificationRequestIEs_Value_CoreNetworkAssistanceInformationForInactive = CoreNetworkAssistanceInformationForInactive_PDU,
    PDU_UEContextModificationRequestIEs_Value_EmergencyFallbackIndicator = EmergencyFallbackIndicator_PDU,
    PDU_UEContextModificationRequestIEs_Value_RRCInactiveTransitionReportRequest = RRCInactiveTransitionReportRequest_PDU,
    PDU_UEContextModificationRequestIEs_Value_GUAMI = GUAMI_PDU,
    PDU_UEContextModificationRequestIEs_Value_CNAssistedRANTuning = CNAssistedRANTuning_PDU,

    PDU_UEContextModificationRequestIEs_Value_MAX = INT_MAX
};

union UEContextModificationRequestIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                 /* PDU_UEContextModificationRequestIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                 /* PDU_UEContextModificationRequestIEs_Value_RAN_UE_NGAP_ID */
    RANPagingPriority *pdu_RANPagingPriority;  
              /* PDU_UEContextModificationRequestIEs_Value_RANPagingPriority */
    SecurityKey     *pdu_SecurityKey;  
                    /* PDU_UEContextModificationRequestIEs_Value_SecurityKey */
    IndexToRFSP     *pdu_IndexToRFSP;  
                    /* PDU_UEContextModificationRequestIEs_Value_IndexToRFSP */
    UEAggregateMaximumBitRate *pdu_UEAggregateMaximumBitRate;  
      /* PDU_UEContextModificationRequestIEs_Value_UEAggregateMaximumBitRate */
    UESecurityCapabilities *pdu_UESecurityCapabilities;  
         /* PDU_UEContextModificationRequestIEs_Value_UESecurityCapabilities */
    CoreNetworkAssistanceInformationForInactive *pdu_CoreNetworkAssistanceInformationForInactive;                       /* PDU_UEContextModificationRequestIEs_Value_CoreNetworkAssistanceInformationForInactive */
    EmergencyFallbackIndicator *pdu_EmergencyFallbackIndicator;  
     /* PDU_UEContextModificationRequestIEs_Value_EmergencyFallbackIndicator */
    RRCInactiveTransitionReportRequest *pdu_RRCInactiveTransitionReportRequest;                                         /* PDU_UEContextModificationRequestIEs_Value_RRCInactiveTransitionReportRequest */
    GUAMI           *pdu_GUAMI;  
                          /* PDU_UEContextModificationRequestIEs_Value_GUAMI */
    CNAssistedRANTuning *pdu_CNAssistedRANTuning;  
            /* PDU_UEContextModificationRequestIEs_Value_CNAssistedRANTuning */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct UEContextModificationRequestIEs_Value {
    UEContextModificationRequestIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    UEContextModificationRequestIEs_Value_union decoded;
} UEContextModificationRequestIEs_Value;

// **************************************************************
//
// UE Context Modification Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// UE CONTEXT MODIFICATION REQUEST
//
// **************************************************************
typedef struct UEContextModificationRequest {
    struct _seqof201 {
        struct _seqof201 *next;
        struct _seq249 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            UEContextModificationRequestIEs_Value value;
        } value;
    } *protocolIEs;
} UEContextModificationRequest;

enum UEContextModificationResponseIEs_Value_PDUs {
    PDU_UEContextModificationResponseIEs_Value_UNKNOWN = 0,

    PDU_UEContextModificationResponseIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_UEContextModificationResponseIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_UEContextModificationResponseIEs_Value_RRCState = RRCState_PDU,
    PDU_UEContextModificationResponseIEs_Value_UserLocationInformation = UserLocationInformation_PDU,
    PDU_UEContextModificationResponseIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_UEContextModificationResponseIEs_Value_MAX = INT_MAX
};

union UEContextModificationResponseIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                /* PDU_UEContextModificationResponseIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                /* PDU_UEContextModificationResponseIEs_Value_RAN_UE_NGAP_ID */
    RRCState        *pdu_RRCState;  
                      /* PDU_UEContextModificationResponseIEs_Value_RRCState */
    UserLocationInformation *pdu_UserLocationInformation;  
       /* PDU_UEContextModificationResponseIEs_Value_UserLocationInformation */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
        /* PDU_UEContextModificationResponseIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct UEContextModificationResponseIEs_Value {
    UEContextModificationResponseIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    UEContextModificationResponseIEs_Value_union decoded;
} UEContextModificationResponseIEs_Value;

// **************************************************************
//
// UE CONTEXT MODIFICATION RESPONSE
//
// **************************************************************
typedef struct UEContextModificationResponse {
    struct _seqof202 {
        struct _seqof202 *next;
        struct _seq250 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            UEContextModificationResponseIEs_Value value;
        } value;
    } *protocolIEs;
} UEContextModificationResponse;

enum UEContextModificationFailureIEs_Value_PDUs {
    PDU_UEContextModificationFailureIEs_Value_UNKNOWN = 0,

    PDU_UEContextModificationFailureIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_UEContextModificationFailureIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_UEContextModificationFailureIEs_Value_Cause = Cause_PDU,
    PDU_UEContextModificationFailureIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_UEContextModificationFailureIEs_Value_MAX = INT_MAX
};

union UEContextModificationFailureIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                 /* PDU_UEContextModificationFailureIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                 /* PDU_UEContextModificationFailureIEs_Value_RAN_UE_NGAP_ID */
    Cause           *pdu_Cause;  
                          /* PDU_UEContextModificationFailureIEs_Value_Cause */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
         /* PDU_UEContextModificationFailureIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct UEContextModificationFailureIEs_Value {
    UEContextModificationFailureIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    UEContextModificationFailureIEs_Value_union decoded;
} UEContextModificationFailureIEs_Value;

typedef struct _seq251 {
    ProtocolIE_ID   id;
    Criticality     criticality;
    UEContextModificationFailureIEs_Value value;
} _seq251;

typedef struct _seqof203_ {
    struct _seqof203_ *next;
    _seq251         value;
} *_seqof203;

// **************************************************************
//
// UE CONTEXT MODIFICATION FAILURE
//
// **************************************************************
typedef struct UEContextModificationFailure {
    struct _seqof203_ *protocolIEs;
} UEContextModificationFailure;

enum RRCInactiveTransitionReportIEs_Value_PDUs {
    PDU_RRCInactiveTransitionReportIEs_Value_UNKNOWN = 0,

    PDU_RRCInactiveTransitionReportIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_RRCInactiveTransitionReportIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_RRCInactiveTransitionReportIEs_Value_RRCState = RRCState_PDU,
    PDU_RRCInactiveTransitionReportIEs_Value_UserLocationInformation = UserLocationInformation_PDU,

    PDU_RRCInactiveTransitionReportIEs_Value_MAX = INT_MAX
};

union RRCInactiveTransitionReportIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                  /* PDU_RRCInactiveTransitionReportIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                  /* PDU_RRCInactiveTransitionReportIEs_Value_RAN_UE_NGAP_ID */
    RRCState        *pdu_RRCState;  
                        /* PDU_RRCInactiveTransitionReportIEs_Value_RRCState */
    UserLocationInformation *pdu_UserLocationInformation;  
         /* PDU_RRCInactiveTransitionReportIEs_Value_UserLocationInformation */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct RRCInactiveTransitionReportIEs_Value {
    RRCInactiveTransitionReportIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    RRCInactiveTransitionReportIEs_Value_union decoded;
} RRCInactiveTransitionReportIEs_Value;

// **************************************************************
//
// RRC INACTIVE TRANSITION REPORT
//
// **************************************************************
typedef struct RRCInactiveTransitionReport {
    struct _seqof204 {
        struct _seqof204 *next;
        struct _seq252 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            RRCInactiveTransitionReportIEs_Value value;
        } value;
    } *protocolIEs;
} RRCInactiveTransitionReport;

enum HandoverRequiredIEs_Value_PDUs {
    PDU_HandoverRequiredIEs_Value_UNKNOWN = 0,

    PDU_HandoverRequiredIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_HandoverRequiredIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_HandoverRequiredIEs_Value_HandoverType = HandoverType_PDU,
    PDU_HandoverRequiredIEs_Value_Cause = Cause_PDU,
    PDU_HandoverRequiredIEs_Value_TargetID = TargetID_PDU,
    PDU_HandoverRequiredIEs_Value_DirectForwardingPathAvailability = DirectForwardingPathAvailability_PDU,
    PDU_HandoverRequiredIEs_Value_PDUSessionResourceListHORqd = PDUSessionResourceListHORqd_PDU,
    PDU_HandoverRequiredIEs_Value_SourceToTarget_TransparentContainer = SourceToTarget_TransparentContainer_PDU,

    PDU_HandoverRequiredIEs_Value_MAX = INT_MAX
};

union HandoverRequiredIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                             /* PDU_HandoverRequiredIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                             /* PDU_HandoverRequiredIEs_Value_RAN_UE_NGAP_ID */
    HandoverType    *pdu_HandoverType;  
                               /* PDU_HandoverRequiredIEs_Value_HandoverType */
    Cause           *pdu_Cause;  /* PDU_HandoverRequiredIEs_Value_Cause */
    TargetID        *pdu_TargetID;  
                                   /* PDU_HandoverRequiredIEs_Value_TargetID */
    DirectForwardingPathAvailability *pdu_DirectForwardingPathAvailability;  
                                        /* PDU_HandoverRequiredIEs_Value_DirectForwardingPathAvailability */
    struct PDUSessionResourceListHORqd_ **pdu_PDUSessionResourceListHORqd;  
                                        /* PDU_HandoverRequiredIEs_Value_PDUSessionResourceListHORqd */
    SourceToTarget_TransparentContainer *pdu_SourceToTarget_TransparentContainer;                                       /* PDU_HandoverRequiredIEs_Value_SourceToTarget_TransparentContainer */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct HandoverRequiredIEs_Value {
    HandoverRequiredIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    HandoverRequiredIEs_Value_union decoded;
} HandoverRequiredIEs_Value;

// **************************************************************
//
// UE MOBILITY MANAGEMENT ELEMENTARY PROCEDURES
//
// **************************************************************

// **************************************************************
//
// Handover Preparation Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// HANDOVER REQUIRED
//
// **************************************************************
typedef struct HandoverRequired {
    struct _seqof205 {
        struct _seqof205 *next;
        struct _seq253 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            HandoverRequiredIEs_Value value;
        } value;
    } *protocolIEs;
} HandoverRequired;

enum HandoverCommandIEs_Value_PDUs {
    PDU_HandoverCommandIEs_Value_UNKNOWN = 0,

    PDU_HandoverCommandIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_HandoverCommandIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_HandoverCommandIEs_Value_HandoverType = HandoverType_PDU,
    PDU_HandoverCommandIEs_Value_NASSecurityParametersFromNGRAN = NASSecurityParametersFromNGRAN_PDU,
    PDU_HandoverCommandIEs_Value_PDUSessionResourceHandoverList = PDUSessionResourceHandoverList_PDU,
    PDU_HandoverCommandIEs_Value_PDUSessionResourceToReleaseListHOCmd = PDUSessionResourceToReleaseListHOCmd_PDU,
    PDU_HandoverCommandIEs_Value_TargetToSource_TransparentContainer = TargetToSource_TransparentContainer_PDU,
    PDU_HandoverCommandIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_HandoverCommandIEs_Value_MAX = INT_MAX
};

union HandoverCommandIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                              /* PDU_HandoverCommandIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                              /* PDU_HandoverCommandIEs_Value_RAN_UE_NGAP_ID */
    HandoverType    *pdu_HandoverType;  
                                /* PDU_HandoverCommandIEs_Value_HandoverType */
    NASSecurityParametersFromNGRAN *pdu_NASSecurityParametersFromNGRAN;  
              /* PDU_HandoverCommandIEs_Value_NASSecurityParametersFromNGRAN */
    struct PDUSessionResourceHandoverList_ **pdu_PDUSessionResourceHandoverList;                                        /* PDU_HandoverCommandIEs_Value_PDUSessionResourceHandoverList */
    struct PDUSessionResourceToReleaseListHOCmd_ **pdu_PDUSessionResourceToReleaseListHOCmd;                            /* PDU_HandoverCommandIEs_Value_PDUSessionResourceToReleaseListHOCmd */
    TargetToSource_TransparentContainer *pdu_TargetToSource_TransparentContainer;                                       /* PDU_HandoverCommandIEs_Value_TargetToSource_TransparentContainer */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
                      /* PDU_HandoverCommandIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct HandoverCommandIEs_Value {
    HandoverCommandIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    HandoverCommandIEs_Value_union decoded;
} HandoverCommandIEs_Value;

// **************************************************************
//
// HANDOVER COMMAND
//
// **************************************************************
typedef struct HandoverCommand {
    struct _seqof206 {
        struct _seqof206 *next;
        struct _seq254 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            HandoverCommandIEs_Value value;
        } value;
    } *protocolIEs;
} HandoverCommand;

// **************************************************************
//
// HANDOVER PREPARATION FAILURE
//
// **************************************************************
typedef struct HandoverPreparationFailure {
    struct _seqof203_ *protocolIEs;
} HandoverPreparationFailure;

enum HandoverRequestIEs_Value_PDUs {
    PDU_HandoverRequestIEs_Value_UNKNOWN = 0,

    PDU_HandoverRequestIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_HandoverRequestIEs_Value_HandoverType = HandoverType_PDU,
    PDU_HandoverRequestIEs_Value_Cause = Cause_PDU,
    PDU_HandoverRequestIEs_Value_UEAggregateMaximumBitRate = UEAggregateMaximumBitRate_PDU,
    PDU_HandoverRequestIEs_Value_CoreNetworkAssistanceInformationForInactive = CoreNetworkAssistanceInformationForInactive_PDU,
    PDU_HandoverRequestIEs_Value_UESecurityCapabilities = UESecurityCapabilities_PDU,
    PDU_HandoverRequestIEs_Value_SecurityContext = SecurityContext_PDU,
    PDU_HandoverRequestIEs_Value_NewSecurityContextInd = NewSecurityContextInd_PDU,
    PDU_HandoverRequestIEs_Value_NAS_PDU = NAS_PDU_PDU,
    PDU_HandoverRequestIEs_Value_PDUSessionResourceSetupListHOReq = PDUSessionResourceSetupListHOReq_PDU,
    PDU_HandoverRequestIEs_Value_AllowedNSSAI = AllowedNSSAI_PDU,
    PDU_HandoverRequestIEs_Value_TraceActivation = TraceActivation_PDU,
    PDU_HandoverRequestIEs_Value_MaskedIMEISV = MaskedIMEISV_PDU,
    PDU_HandoverRequestIEs_Value_SourceToTarget_TransparentContainer = SourceToTarget_TransparentContainer_PDU,
    PDU_HandoverRequestIEs_Value_MobilityRestrictionList = MobilityRestrictionList_PDU,
    PDU_HandoverRequestIEs_Value_LocationReportingRequestType = LocationReportingRequestType_PDU,
    PDU_HandoverRequestIEs_Value_RRCInactiveTransitionReportRequest = RRCInactiveTransitionReportRequest_PDU,
    PDU_HandoverRequestIEs_Value_GUAMI = GUAMI_PDU,
    PDU_HandoverRequestIEs_Value_RedirectionVoiceFallback = RedirectionVoiceFallback_PDU,
    PDU_HandoverRequestIEs_Value_CNAssistedRANTuning = CNAssistedRANTuning_PDU,

    PDU_HandoverRequestIEs_Value_MAX = INT_MAX
};

union HandoverRequestIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                              /* PDU_HandoverRequestIEs_Value_AMF_UE_NGAP_ID */
    HandoverType    *pdu_HandoverType;  
                                /* PDU_HandoverRequestIEs_Value_HandoverType */
    Cause           *pdu_Cause;  /* PDU_HandoverRequestIEs_Value_Cause */
    UEAggregateMaximumBitRate *pdu_UEAggregateMaximumBitRate;  
                   /* PDU_HandoverRequestIEs_Value_UEAggregateMaximumBitRate */
    CoreNetworkAssistanceInformationForInactive *pdu_CoreNetworkAssistanceInformationForInactive;                       /* PDU_HandoverRequestIEs_Value_CoreNetworkAssistanceInformationForInactive */
    UESecurityCapabilities *pdu_UESecurityCapabilities;  
                      /* PDU_HandoverRequestIEs_Value_UESecurityCapabilities */
    SecurityContext *pdu_SecurityContext;  
                             /* PDU_HandoverRequestIEs_Value_SecurityContext */
    NewSecurityContextInd *pdu_NewSecurityContextInd;  
                       /* PDU_HandoverRequestIEs_Value_NewSecurityContextInd */
    NAS_PDU         *pdu_NAS_PDU;  /* PDU_HandoverRequestIEs_Value_NAS_PDU */
    struct PDUSessionResourceSetupListHOReq_ **pdu_PDUSessionResourceSetupListHOReq;                                    /* PDU_HandoverRequestIEs_Value_PDUSessionResourceSetupListHOReq */
    struct AllowedNSSAI_ **pdu_AllowedNSSAI;  
                                /* PDU_HandoverRequestIEs_Value_AllowedNSSAI */
    TraceActivation *pdu_TraceActivation;  
                             /* PDU_HandoverRequestIEs_Value_TraceActivation */
    MaskedIMEISV    *pdu_MaskedIMEISV;  
                                /* PDU_HandoverRequestIEs_Value_MaskedIMEISV */
    SourceToTarget_TransparentContainer *pdu_SourceToTarget_TransparentContainer;                                       /* PDU_HandoverRequestIEs_Value_SourceToTarget_TransparentContainer */
    MobilityRestrictionList *pdu_MobilityRestrictionList;  
                     /* PDU_HandoverRequestIEs_Value_MobilityRestrictionList */
    LocationReportingRequestType *pdu_LocationReportingRequestType;  
                /* PDU_HandoverRequestIEs_Value_LocationReportingRequestType */
    RRCInactiveTransitionReportRequest *pdu_RRCInactiveTransitionReportRequest;                                         /* PDU_HandoverRequestIEs_Value_RRCInactiveTransitionReportRequest */
    GUAMI           *pdu_GUAMI;  /* PDU_HandoverRequestIEs_Value_GUAMI */
    RedirectionVoiceFallback *pdu_RedirectionVoiceFallback;  
                    /* PDU_HandoverRequestIEs_Value_RedirectionVoiceFallback */
    CNAssistedRANTuning *pdu_CNAssistedRANTuning;  
                         /* PDU_HandoverRequestIEs_Value_CNAssistedRANTuning */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct HandoverRequestIEs_Value {
    HandoverRequestIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    HandoverRequestIEs_Value_union decoded;
} HandoverRequestIEs_Value;

// **************************************************************
//
// Handover Resource Allocation Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// HANDOVER REQUEST
//
// **************************************************************
typedef struct HandoverRequest {
    struct _seqof208 {
        struct _seqof208 *next;
        struct _seq256 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            HandoverRequestIEs_Value value;
        } value;
    } *protocolIEs;
} HandoverRequest;

enum HandoverRequestAcknowledgeIEs_Value_PDUs {
    PDU_HandoverRequestAcknowledgeIEs_Value_UNKNOWN = 0,

    PDU_HandoverRequestAcknowledgeIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_HandoverRequestAcknowledgeIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_HandoverRequestAcknowledgeIEs_Value_PDUSessionResourceAdmittedList = PDUSessionResourceAdmittedList_PDU,
    PDU_HandoverRequestAcknowledgeIEs_Value_PDUSessionResourceFailedToSetupListHOAck = PDUSessionResourceFailedToSetupListHOAck_PDU,
    PDU_HandoverRequestAcknowledgeIEs_Value_TargetToSource_TransparentContainer = TargetToSource_TransparentContainer_PDU,
    PDU_HandoverRequestAcknowledgeIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_HandoverRequestAcknowledgeIEs_Value_MAX = INT_MAX
};

union HandoverRequestAcknowledgeIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                   /* PDU_HandoverRequestAcknowledgeIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                   /* PDU_HandoverRequestAcknowledgeIEs_Value_RAN_UE_NGAP_ID */
    struct PDUSessionResourceAdmittedList_ **pdu_PDUSessionResourceAdmittedList;                                        /* PDU_HandoverRequestAcknowledgeIEs_Value_PDUSessionResourceAdmittedList */
    struct PDUSessionResourceFailedToSetupListHOAck_ **pdu_PDUSessionResourceFailedToSetupListHOAck;                    /* PDU_HandoverRequestAcknowledgeIEs_Value_PDUSessionResourceFailedToSetupListHOAck */
    TargetToSource_TransparentContainer *pdu_TargetToSource_TransparentContainer;                                       /* PDU_HandoverRequestAcknowledgeIEs_Value_TargetToSource_TransparentContainer */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
           /* PDU_HandoverRequestAcknowledgeIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct HandoverRequestAcknowledgeIEs_Value {
    HandoverRequestAcknowledgeIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    HandoverRequestAcknowledgeIEs_Value_union decoded;
} HandoverRequestAcknowledgeIEs_Value;

// **************************************************************
//
// HANDOVER REQUEST ACKNOWLEDGE
//
// **************************************************************
typedef struct HandoverRequestAcknowledge {
    struct _seqof209 {
        struct _seqof209 *next;
        struct _seq257 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            HandoverRequestAcknowledgeIEs_Value value;
        } value;
    } *protocolIEs;
} HandoverRequestAcknowledge;

enum HandoverFailureIEs_Value_PDUs {
    PDU_HandoverFailureIEs_Value_UNKNOWN = 0,

    PDU_HandoverFailureIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_HandoverFailureIEs_Value_Cause = Cause_PDU,
    PDU_HandoverFailureIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_HandoverFailureIEs_Value_MAX = INT_MAX
};

union HandoverFailureIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                              /* PDU_HandoverFailureIEs_Value_AMF_UE_NGAP_ID */
    Cause           *pdu_Cause;  /* PDU_HandoverFailureIEs_Value_Cause */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
                      /* PDU_HandoverFailureIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct HandoverFailureIEs_Value {
    HandoverFailureIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    HandoverFailureIEs_Value_union decoded;
} HandoverFailureIEs_Value;

// **************************************************************
//
// HANDOVER FAILURE
//
// **************************************************************
typedef struct HandoverFailure {
    struct _seqof210 {
        struct _seqof210 *next;
        struct _seq258 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            HandoverFailureIEs_Value value;
        } value;
    } *protocolIEs;
} HandoverFailure;

enum HandoverNotifyIEs_Value_PDUs {
    PDU_HandoverNotifyIEs_Value_UNKNOWN = 0,

    PDU_HandoverNotifyIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_HandoverNotifyIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_HandoverNotifyIEs_Value_UserLocationInformation = UserLocationInformation_PDU,

    PDU_HandoverNotifyIEs_Value_MAX = INT_MAX
};

union HandoverNotifyIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                               /* PDU_HandoverNotifyIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                               /* PDU_HandoverNotifyIEs_Value_RAN_UE_NGAP_ID */
    UserLocationInformation *pdu_UserLocationInformation;  
                      /* PDU_HandoverNotifyIEs_Value_UserLocationInformation */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct HandoverNotifyIEs_Value {
    HandoverNotifyIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    HandoverNotifyIEs_Value_union decoded;
} HandoverNotifyIEs_Value;

// **************************************************************
//
// Handover Notification Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// HANDOVER NOTIFY
//
// **************************************************************
typedef struct HandoverNotify {
    struct _seqof211 {
        struct _seqof211 *next;
        struct _seq259 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            HandoverNotifyIEs_Value value;
        } value;
    } *protocolIEs;
} HandoverNotify;

enum PathSwitchRequestIEs_Value_PDUs {
    PDU_PathSwitchRequestIEs_Value_UNKNOWN = 0,

    PDU_PathSwitchRequestIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_PathSwitchRequestIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_PathSwitchRequestIEs_Value_UserLocationInformation = UserLocationInformation_PDU,
    PDU_PathSwitchRequestIEs_Value_UESecurityCapabilities = UESecurityCapabilities_PDU,
    PDU_PathSwitchRequestIEs_Value_PDUSessionResourceToBeSwitchedDLList = PDUSessionResourceToBeSwitchedDLList_PDU,
    PDU_PathSwitchRequestIEs_Value_PDUSessionResourceFailedToSetupListPSReq = PDUSessionResourceFailedToSetupListPSReq_PDU,

    PDU_PathSwitchRequestIEs_Value_MAX = INT_MAX
};

union PathSwitchRequestIEs_Value_union {
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                            /* PDU_PathSwitchRequestIEs_Value_RAN_UE_NGAP_ID */
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                            /* PDU_PathSwitchRequestIEs_Value_AMF_UE_NGAP_ID */
    UserLocationInformation *pdu_UserLocationInformation;  
                   /* PDU_PathSwitchRequestIEs_Value_UserLocationInformation */
    UESecurityCapabilities *pdu_UESecurityCapabilities;  
                    /* PDU_PathSwitchRequestIEs_Value_UESecurityCapabilities */
    struct PDUSessionResourceToBeSwitchedDLList_ **pdu_PDUSessionResourceToBeSwitchedDLList;                            /* PDU_PathSwitchRequestIEs_Value_PDUSessionResourceToBeSwitchedDLList */
    struct PDUSessionResourceFailedToSetupListPSReq_ **pdu_PDUSessionResourceFailedToSetupListPSReq;                    /* PDU_PathSwitchRequestIEs_Value_PDUSessionResourceFailedToSetupListPSReq */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PathSwitchRequestIEs_Value {
    PathSwitchRequestIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PathSwitchRequestIEs_Value_union decoded;
} PathSwitchRequestIEs_Value;

// **************************************************************
//
// Path Switch Request Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// PATH SWITCH REQUEST
//
// **************************************************************
typedef struct PathSwitchRequest {
    struct _seqof212 {
        struct _seqof212 *next;
        struct _seq260 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PathSwitchRequestIEs_Value value;
        } value;
    } *protocolIEs;
} PathSwitchRequest;

enum PathSwitchRequestAcknowledgeIEs_Value_PDUs {
    PDU_PathSwitchRequestAcknowledgeIEs_Value_UNKNOWN = 0,

    PDU_PathSwitchRequestAcknowledgeIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_PathSwitchRequestAcknowledgeIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_PathSwitchRequestAcknowledgeIEs_Value_UESecurityCapabilities = UESecurityCapabilities_PDU,
    PDU_PathSwitchRequestAcknowledgeIEs_Value_SecurityContext = SecurityContext_PDU,
    PDU_PathSwitchRequestAcknowledgeIEs_Value_NewSecurityContextInd = NewSecurityContextInd_PDU,
    PDU_PathSwitchRequestAcknowledgeIEs_Value_PDUSessionResourceSwitchedList = PDUSessionResourceSwitchedList_PDU,
    PDU_PathSwitchRequestAcknowledgeIEs_Value_PDUSessionResourceReleasedListPSAck = PDUSessionResourceReleasedListPSAck_PDU,
    PDU_PathSwitchRequestAcknowledgeIEs_Value_AllowedNSSAI = AllowedNSSAI_PDU,
    PDU_PathSwitchRequestAcknowledgeIEs_Value_CoreNetworkAssistanceInformationForInactive = CoreNetworkAssistanceInformationForInactive_PDU,
    PDU_PathSwitchRequestAcknowledgeIEs_Value_RRCInactiveTransitionReportRequest = RRCInactiveTransitionReportRequest_PDU,
    PDU_PathSwitchRequestAcknowledgeIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,
    PDU_PathSwitchRequestAcknowledgeIEs_Value_RedirectionVoiceFallback = RedirectionVoiceFallback_PDU,
    PDU_PathSwitchRequestAcknowledgeIEs_Value_CNAssistedRANTuning = CNAssistedRANTuning_PDU,

    PDU_PathSwitchRequestAcknowledgeIEs_Value_MAX = INT_MAX
};

union PathSwitchRequestAcknowledgeIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                 /* PDU_PathSwitchRequestAcknowledgeIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                 /* PDU_PathSwitchRequestAcknowledgeIEs_Value_RAN_UE_NGAP_ID */
    UESecurityCapabilities *pdu_UESecurityCapabilities;  
         /* PDU_PathSwitchRequestAcknowledgeIEs_Value_UESecurityCapabilities */
    SecurityContext *pdu_SecurityContext;  
                /* PDU_PathSwitchRequestAcknowledgeIEs_Value_SecurityContext */
    NewSecurityContextInd *pdu_NewSecurityContextInd;  
          /* PDU_PathSwitchRequestAcknowledgeIEs_Value_NewSecurityContextInd */
    struct PDUSessionResourceSwitchedList_ **pdu_PDUSessionResourceSwitchedList;                                        /* PDU_PathSwitchRequestAcknowledgeIEs_Value_PDUSessionResourceSwitchedList */
    struct PDUSessionResourceReleasedListPSAck_ **pdu_PDUSessionResourceReleasedListPSAck;                              /* PDU_PathSwitchRequestAcknowledgeIEs_Value_PDUSessionResourceReleasedListPSAck */
    struct AllowedNSSAI_ **pdu_AllowedNSSAI;  
                   /* PDU_PathSwitchRequestAcknowledgeIEs_Value_AllowedNSSAI */
    CoreNetworkAssistanceInformationForInactive *pdu_CoreNetworkAssistanceInformationForInactive;                       /* PDU_PathSwitchRequestAcknowledgeIEs_Value_CoreNetworkAssistanceInformationForInactive */
    RRCInactiveTransitionReportRequest *pdu_RRCInactiveTransitionReportRequest;                                         /* PDU_PathSwitchRequestAcknowledgeIEs_Value_RRCInactiveTransitionReportRequest */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
         /* PDU_PathSwitchRequestAcknowledgeIEs_Value_CriticalityDiagnostics */
    RedirectionVoiceFallback *pdu_RedirectionVoiceFallback;  
       /* PDU_PathSwitchRequestAcknowledgeIEs_Value_RedirectionVoiceFallback */
    CNAssistedRANTuning *pdu_CNAssistedRANTuning;  
            /* PDU_PathSwitchRequestAcknowledgeIEs_Value_CNAssistedRANTuning */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PathSwitchRequestAcknowledgeIEs_Value {
    PathSwitchRequestAcknowledgeIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PathSwitchRequestAcknowledgeIEs_Value_union decoded;
} PathSwitchRequestAcknowledgeIEs_Value;

// **************************************************************
//
// PATH SWITCH REQUEST ACKNOWLEDGE
//
// **************************************************************
typedef struct PathSwitchRequestAcknowledge {
    struct _seqof213 {
        struct _seqof213 *next;
        struct _seq261 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PathSwitchRequestAcknowledgeIEs_Value value;
        } value;
    } *protocolIEs;
} PathSwitchRequestAcknowledge;

enum PathSwitchRequestFailureIEs_Value_PDUs {
    PDU_PathSwitchRequestFailureIEs_Value_UNKNOWN = 0,

    PDU_PathSwitchRequestFailureIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_PathSwitchRequestFailureIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_PathSwitchRequestFailureIEs_Value_PDUSessionResourceReleasedListPSFail = PDUSessionResourceReleasedListPSFail_PDU,
    PDU_PathSwitchRequestFailureIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_PathSwitchRequestFailureIEs_Value_MAX = INT_MAX
};

union PathSwitchRequestFailureIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                     /* PDU_PathSwitchRequestFailureIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                     /* PDU_PathSwitchRequestFailureIEs_Value_RAN_UE_NGAP_ID */
    struct PDUSessionResourceReleasedListPSFail_ **pdu_PDUSessionResourceReleasedListPSFail;                            /* PDU_PathSwitchRequestFailureIEs_Value_PDUSessionResourceReleasedListPSFail */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
             /* PDU_PathSwitchRequestFailureIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PathSwitchRequestFailureIEs_Value {
    PathSwitchRequestFailureIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PathSwitchRequestFailureIEs_Value_union decoded;
} PathSwitchRequestFailureIEs_Value;

// **************************************************************
//
// PATH SWITCH REQUEST FAILURE
//
// **************************************************************
typedef struct PathSwitchRequestFailure {
    struct _seqof214 {
        struct _seqof214 *next;
        struct _seq262 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PathSwitchRequestFailureIEs_Value value;
        } value;
    } *protocolIEs;
} PathSwitchRequestFailure;

enum HandoverCancelIEs_Value_PDUs {
    PDU_HandoverCancelIEs_Value_UNKNOWN = 0,

    PDU_HandoverCancelIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_HandoverCancelIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_HandoverCancelIEs_Value_Cause = Cause_PDU,

    PDU_HandoverCancelIEs_Value_MAX = INT_MAX
};

union HandoverCancelIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                               /* PDU_HandoverCancelIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                               /* PDU_HandoverCancelIEs_Value_RAN_UE_NGAP_ID */
    Cause           *pdu_Cause;  /* PDU_HandoverCancelIEs_Value_Cause */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct HandoverCancelIEs_Value {
    HandoverCancelIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    HandoverCancelIEs_Value_union decoded;
} HandoverCancelIEs_Value;

typedef struct _seq263 {
    ProtocolIE_ID   id;
    Criticality     criticality;
    HandoverCancelIEs_Value value;
} _seq263;

typedef struct _seqof215_ {
    struct _seqof215_ *next;
    _seq263         value;
} *_seqof215;

// **************************************************************
//
// Handover Cancellation Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// HANDOVER CANCEL
//
// **************************************************************
typedef struct HandoverCancel {
    struct _seqof215_ *protocolIEs;
} HandoverCancel;

enum HandoverCancelAcknowledgeIEs_Value_PDUs {
    PDU_HandoverCancelAcknowledgeIEs_Value_UNKNOWN = 0,

    PDU_HandoverCancelAcknowledgeIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_HandoverCancelAcknowledgeIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_HandoverCancelAcknowledgeIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_HandoverCancelAcknowledgeIEs_Value_MAX = INT_MAX
};

union HandoverCancelAcknowledgeIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                    /* PDU_HandoverCancelAcknowledgeIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                    /* PDU_HandoverCancelAcknowledgeIEs_Value_RAN_UE_NGAP_ID */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
            /* PDU_HandoverCancelAcknowledgeIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct HandoverCancelAcknowledgeIEs_Value {
    HandoverCancelAcknowledgeIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    HandoverCancelAcknowledgeIEs_Value_union decoded;
} HandoverCancelAcknowledgeIEs_Value;

// **************************************************************
//
// HANDOVER CANCEL ACKNOWLEDGE
//
// **************************************************************
typedef struct HandoverCancelAcknowledge {
    struct _seqof216 {
        struct _seqof216 *next;
        struct _seq264 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            HandoverCancelAcknowledgeIEs_Value value;
        } value;
    } *protocolIEs;
} HandoverCancelAcknowledge;

enum UplinkRANStatusTransferIEs_Value_PDUs {
    PDU_UplinkRANStatusTransferIEs_Value_UNKNOWN = 0,

    PDU_UplinkRANStatusTransferIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_UplinkRANStatusTransferIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_UplinkRANStatusTransferIEs_Value_RANStatusTransfer_TransparentContainer = RANStatusTransfer_TransparentContainer_PDU,

    PDU_UplinkRANStatusTransferIEs_Value_MAX = INT_MAX
};

union UplinkRANStatusTransferIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                      /* PDU_UplinkRANStatusTransferIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                      /* PDU_UplinkRANStatusTransferIEs_Value_RAN_UE_NGAP_ID */
    RANStatusTransfer_TransparentContainer *pdu_RANStatusTransfer_TransparentContainer;                                 /* PDU_UplinkRANStatusTransferIEs_Value_RANStatusTransfer_TransparentContainer */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct UplinkRANStatusTransferIEs_Value {
    UplinkRANStatusTransferIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    UplinkRANStatusTransferIEs_Value_union decoded;
} UplinkRANStatusTransferIEs_Value;

typedef struct _seq265 {
    ProtocolIE_ID   id;
    Criticality     criticality;
    UplinkRANStatusTransferIEs_Value value;
} _seq265;

typedef struct _seqof217_ {
    struct _seqof217_ *next;
    _seq265         value;
} *_seqof217;

// **************************************************************
//
// Uplink RAN Status Transfer Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// UPLINK RAN STATUS TRANSFER
//
// **************************************************************
typedef struct UplinkRANStatusTransfer {
    struct _seqof217_ *protocolIEs;
} UplinkRANStatusTransfer;

// **************************************************************
//
// Downlink RAN Status Transfer Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// DOWNLINK RAN STATUS TRANSFER
//
// **************************************************************
typedef struct DownlinkRANStatusTransfer {
    struct _seqof217_ *protocolIEs;
} DownlinkRANStatusTransfer;

enum PagingIEs_Value_PDUs {
    PDU_PagingIEs_Value_UNKNOWN    = 0,

    PDU_PagingIEs_Value_UEPagingIdentity = UEPagingIdentity_PDU,
    PDU_PagingIEs_Value_PagingDRX  = PagingDRX_PDU,
    PDU_PagingIEs_Value_TAIListForPaging = TAIListForPaging_PDU,
    PDU_PagingIEs_Value_PagingPriority = PagingPriority_PDU,
    PDU_PagingIEs_Value_UERadioCapabilityForPaging = UERadioCapabilityForPaging_PDU,
    PDU_PagingIEs_Value_PagingOrigin = PagingOrigin_PDU,
    PDU_PagingIEs_Value_AssistanceDataForPaging = AssistanceDataForPaging_PDU,

    PDU_PagingIEs_Value_MAX        = INT_MAX
};

union PagingIEs_Value_union {
    UEPagingIdentity *pdu_UEPagingIdentity;  
                                     /* PDU_PagingIEs_Value_UEPagingIdentity */
    PagingDRX       *pdu_PagingDRX;  /* PDU_PagingIEs_Value_PagingDRX */
    struct TAIListForPaging_ **pdu_TAIListForPaging;  
                                     /* PDU_PagingIEs_Value_TAIListForPaging */
    PagingPriority  *pdu_PagingPriority;  
                                       /* PDU_PagingIEs_Value_PagingPriority */
    UERadioCapabilityForPaging *pdu_UERadioCapabilityForPaging;  
                           /* PDU_PagingIEs_Value_UERadioCapabilityForPaging */
    PagingOrigin    *pdu_PagingOrigin;  /* PDU_PagingIEs_Value_PagingOrigin */
    AssistanceDataForPaging *pdu_AssistanceDataForPaging;  
                              /* PDU_PagingIEs_Value_AssistanceDataForPaging */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PagingIEs_Value {
    PagingIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PagingIEs_Value_union decoded;
} PagingIEs_Value;

// **************************************************************
//
// PAGING ELEMENTARY PROCEDURE
//
// **************************************************************

// **************************************************************
//
// PAGING
//
// **************************************************************
typedef struct Paging {
    struct _seqof219 {
        struct _seqof219 *next;
        struct _seq267 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PagingIEs_Value value;
        } value;
    } *protocolIEs;
} Paging;

enum InitialUEMessage_IEs_Value_PDUs {
    PDU_InitialUEMessage_IEs_Value_UNKNOWN = 0,

    PDU_InitialUEMessage_IEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_InitialUEMessage_IEs_Value_NAS_PDU = NAS_PDU_PDU,
    PDU_InitialUEMessage_IEs_Value_UserLocationInformation = UserLocationInformation_PDU,
    PDU_InitialUEMessage_IEs_Value_RRCEstablishmentCause = RRCEstablishmentCause_PDU,
    PDU_InitialUEMessage_IEs_Value_FiveG_S_TMSI = FiveG_S_TMSI_PDU,
    PDU_InitialUEMessage_IEs_Value_AMFSetID = AMFSetID_PDU,
    PDU_InitialUEMessage_IEs_Value_UEContextRequest = UEContextRequest_PDU,
    PDU_InitialUEMessage_IEs_Value_AllowedNSSAI = AllowedNSSAI_PDU,

    PDU_InitialUEMessage_IEs_Value_MAX = INT_MAX
};

union InitialUEMessage_IEs_Value_union {
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                            /* PDU_InitialUEMessage_IEs_Value_RAN_UE_NGAP_ID */
    NAS_PDU         *pdu_NAS_PDU;  /* PDU_InitialUEMessage_IEs_Value_NAS_PDU */
    UserLocationInformation *pdu_UserLocationInformation;  
                   /* PDU_InitialUEMessage_IEs_Value_UserLocationInformation */
    RRCEstablishmentCause *pdu_RRCEstablishmentCause;  
                     /* PDU_InitialUEMessage_IEs_Value_RRCEstablishmentCause */
    FiveG_S_TMSI    *pdu_FiveG_S_TMSI;  
                              /* PDU_InitialUEMessage_IEs_Value_FiveG_S_TMSI */
    AMFSetID        *pdu_AMFSetID;  
                                  /* PDU_InitialUEMessage_IEs_Value_AMFSetID */
    UEContextRequest *pdu_UEContextRequest;  
                          /* PDU_InitialUEMessage_IEs_Value_UEContextRequest */
    struct AllowedNSSAI_ **pdu_AllowedNSSAI;  
                              /* PDU_InitialUEMessage_IEs_Value_AllowedNSSAI */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct InitialUEMessage_IEs_Value {
    InitialUEMessage_IEs_Value_PDUs pduNum;
    OssBuf          encoded;
    InitialUEMessage_IEs_Value_union decoded;
} InitialUEMessage_IEs_Value;

// **************************************************************
//
// NAS TRANSPORT ELEMENTARY PROCEDURES
//
// **************************************************************

// **************************************************************
//
// INITIAL UE MESSAGE
//
// **************************************************************
typedef struct InitialUEMessage {
    struct _seqof220 {
        struct _seqof220 *next;
        struct _seq268 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            InitialUEMessage_IEs_Value value;
        } value;
    } *protocolIEs;
} InitialUEMessage;

enum DownlinkNASTransport_IEs_Value_PDUs {
    PDU_DownlinkNASTransport_IEs_Value_UNKNOWN = 0,

    PDU_DownlinkNASTransport_IEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_DownlinkNASTransport_IEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_DownlinkNASTransport_IEs_Value_AMFName = AMFName_PDU,
    PDU_DownlinkNASTransport_IEs_Value_RANPagingPriority = RANPagingPriority_PDU,
    PDU_DownlinkNASTransport_IEs_Value_NAS_PDU = NAS_PDU_PDU,
    PDU_DownlinkNASTransport_IEs_Value_MobilityRestrictionList = MobilityRestrictionList_PDU,
    PDU_DownlinkNASTransport_IEs_Value_IndexToRFSP = IndexToRFSP_PDU,
    PDU_DownlinkNASTransport_IEs_Value_UEAggregateMaximumBitRate = UEAggregateMaximumBitRate_PDU,
    PDU_DownlinkNASTransport_IEs_Value_AllowedNSSAI = AllowedNSSAI_PDU,

    PDU_DownlinkNASTransport_IEs_Value_MAX = INT_MAX
};

union DownlinkNASTransport_IEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                        /* PDU_DownlinkNASTransport_IEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                        /* PDU_DownlinkNASTransport_IEs_Value_RAN_UE_NGAP_ID */
    AMFName         *pdu_AMFName;  
                               /* PDU_DownlinkNASTransport_IEs_Value_AMFName */
    RANPagingPriority *pdu_RANPagingPriority;  
                     /* PDU_DownlinkNASTransport_IEs_Value_RANPagingPriority */
    NAS_PDU         *pdu_NAS_PDU;  
                               /* PDU_DownlinkNASTransport_IEs_Value_NAS_PDU */
    MobilityRestrictionList *pdu_MobilityRestrictionList;  
               /* PDU_DownlinkNASTransport_IEs_Value_MobilityRestrictionList */
    IndexToRFSP     *pdu_IndexToRFSP;  
                           /* PDU_DownlinkNASTransport_IEs_Value_IndexToRFSP */
    UEAggregateMaximumBitRate *pdu_UEAggregateMaximumBitRate;  
             /* PDU_DownlinkNASTransport_IEs_Value_UEAggregateMaximumBitRate */
    struct AllowedNSSAI_ **pdu_AllowedNSSAI;  
                          /* PDU_DownlinkNASTransport_IEs_Value_AllowedNSSAI */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct DownlinkNASTransport_IEs_Value {
    DownlinkNASTransport_IEs_Value_PDUs pduNum;
    OssBuf          encoded;
    DownlinkNASTransport_IEs_Value_union decoded;
} DownlinkNASTransport_IEs_Value;

// **************************************************************
//
// DOWNLINK NAS TRANSPORT
//
// **************************************************************
typedef struct DownlinkNASTransport {
    struct _seqof221 {
        struct _seqof221 *next;
        struct _seq269 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            DownlinkNASTransport_IEs_Value value;
        } value;
    } *protocolIEs;
} DownlinkNASTransport;

enum UplinkNASTransport_IEs_Value_PDUs {
    PDU_UplinkNASTransport_IEs_Value_UNKNOWN = 0,

    PDU_UplinkNASTransport_IEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_UplinkNASTransport_IEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_UplinkNASTransport_IEs_Value_NAS_PDU = NAS_PDU_PDU,
    PDU_UplinkNASTransport_IEs_Value_UserLocationInformation = UserLocationInformation_PDU,

    PDU_UplinkNASTransport_IEs_Value_MAX = INT_MAX
};

union UplinkNASTransport_IEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                          /* PDU_UplinkNASTransport_IEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                          /* PDU_UplinkNASTransport_IEs_Value_RAN_UE_NGAP_ID */
    NAS_PDU         *pdu_NAS_PDU;  
                                 /* PDU_UplinkNASTransport_IEs_Value_NAS_PDU */
    UserLocationInformation *pdu_UserLocationInformation;  
                 /* PDU_UplinkNASTransport_IEs_Value_UserLocationInformation */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct UplinkNASTransport_IEs_Value {
    UplinkNASTransport_IEs_Value_PDUs pduNum;
    OssBuf          encoded;
    UplinkNASTransport_IEs_Value_union decoded;
} UplinkNASTransport_IEs_Value;

// **************************************************************
//
// UPLINK NAS TRANSPORT
//
// **************************************************************
typedef struct UplinkNASTransport {
    struct _seqof222 {
        struct _seqof222 *next;
        struct _seq270 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            UplinkNASTransport_IEs_Value value;
        } value;
    } *protocolIEs;
} UplinkNASTransport;

enum NASNonDeliveryIndication_IEs_Value_PDUs {
    PDU_NASNonDeliveryIndication_IEs_Value_UNKNOWN = 0,

    PDU_NASNonDeliveryIndication_IEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_NASNonDeliveryIndication_IEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_NASNonDeliveryIndication_IEs_Value_NAS_PDU = NAS_PDU_PDU,
    PDU_NASNonDeliveryIndication_IEs_Value_Cause = Cause_PDU,

    PDU_NASNonDeliveryIndication_IEs_Value_MAX = INT_MAX
};

union NASNonDeliveryIndication_IEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                    /* PDU_NASNonDeliveryIndication_IEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                    /* PDU_NASNonDeliveryIndication_IEs_Value_RAN_UE_NGAP_ID */
    NAS_PDU         *pdu_NAS_PDU;  
                           /* PDU_NASNonDeliveryIndication_IEs_Value_NAS_PDU */
    Cause           *pdu_Cause;  
                             /* PDU_NASNonDeliveryIndication_IEs_Value_Cause */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct NASNonDeliveryIndication_IEs_Value {
    NASNonDeliveryIndication_IEs_Value_PDUs pduNum;
    OssBuf          encoded;
    NASNonDeliveryIndication_IEs_Value_union decoded;
} NASNonDeliveryIndication_IEs_Value;

// **************************************************************
//
// NAS NON DELIVERY INDICATION
//
// **************************************************************
typedef struct NASNonDeliveryIndication {
    struct _seqof223 {
        struct _seqof223 *next;
        struct _seq271 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            NASNonDeliveryIndication_IEs_Value value;
        } value;
    } *protocolIEs;
} NASNonDeliveryIndication;

enum RerouteNASRequest_IEs_Value_PDUs {
    PDU_RerouteNASRequest_IEs_Value_UNKNOWN = 0,

    PDU_RerouteNASRequest_IEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_RerouteNASRequest_IEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_RerouteNASRequest_IEs_Value_TYPE = RerouteNASRequest_IEs_TYPE_PDU,
    PDU_RerouteNASRequest_IEs_Value_AMFSetID = AMFSetID_PDU,
    PDU_RerouteNASRequest_IEs_Value_AllowedNSSAI = AllowedNSSAI_PDU,

    PDU_RerouteNASRequest_IEs_Value_MAX = INT_MAX
};

union RerouteNASRequest_IEs_Value_union {
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                           /* PDU_RerouteNASRequest_IEs_Value_RAN_UE_NGAP_ID */
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                           /* PDU_RerouteNASRequest_IEs_Value_AMF_UE_NGAP_ID */
    _octet1         *pdu_RerouteNASRequest_IEs_TYPE;  
                                     /* PDU_RerouteNASRequest_IEs_Value_TYPE */
    AMFSetID        *pdu_AMFSetID;  
                                 /* PDU_RerouteNASRequest_IEs_Value_AMFSetID */
    struct AllowedNSSAI_ **pdu_AllowedNSSAI;  
                             /* PDU_RerouteNASRequest_IEs_Value_AllowedNSSAI */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct RerouteNASRequest_IEs_Value {
    RerouteNASRequest_IEs_Value_PDUs pduNum;
    OssBuf          encoded;
    RerouteNASRequest_IEs_Value_union decoded;
} RerouteNASRequest_IEs_Value;

// **************************************************************
//
// REROUTE NAS REQUEST
//
// **************************************************************
typedef struct RerouteNASRequest {
    struct _seqof224 {
        struct _seqof224 *next;
        struct _seq272 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            RerouteNASRequest_IEs_Value value;
        } value;
    } *protocolIEs;
} RerouteNASRequest;

enum NGSetupRequestIEs_Value_PDUs {
    PDU_NGSetupRequestIEs_Value_UNKNOWN = 0,

    PDU_NGSetupRequestIEs_Value_GlobalRANNodeID = GlobalRANNodeID_PDU,
    PDU_NGSetupRequestIEs_Value_RANNodeName = RANNodeName_PDU,
    PDU_NGSetupRequestIEs_Value_SupportedTAList = SupportedTAList_PDU,
    PDU_NGSetupRequestIEs_Value_PagingDRX = PagingDRX_PDU,
    PDU_NGSetupRequestIEs_Value_UERetentionInformation = UERetentionInformation_PDU,

    PDU_NGSetupRequestIEs_Value_MAX = INT_MAX
};

union NGSetupRequestIEs_Value_union {
    GlobalRANNodeID *pdu_GlobalRANNodeID;  
                              /* PDU_NGSetupRequestIEs_Value_GlobalRANNodeID */
    RANNodeName     *pdu_RANNodeName;  
                                  /* PDU_NGSetupRequestIEs_Value_RANNodeName */
    struct SupportedTAList_ **pdu_SupportedTAList;  
                              /* PDU_NGSetupRequestIEs_Value_SupportedTAList */
    PagingDRX       *pdu_PagingDRX;  
                                    /* PDU_NGSetupRequestIEs_Value_PagingDRX */
    UERetentionInformation *pdu_UERetentionInformation;  
                       /* PDU_NGSetupRequestIEs_Value_UERetentionInformation */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct NGSetupRequestIEs_Value {
    NGSetupRequestIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    NGSetupRequestIEs_Value_union decoded;
} NGSetupRequestIEs_Value;

// **************************************************************
//
// INTERFACE MANAGEMENT ELEMENTARY PROCEDURES
//
// **************************************************************

// **************************************************************
//
// NG Setup Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// NG SETUP REQUEST
//
// **************************************************************
typedef struct NGSetupRequest {
    struct _seqof225 {
        struct _seqof225 *next;
        struct _seq273 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            NGSetupRequestIEs_Value value;
        } value;
    } *protocolIEs;
} NGSetupRequest;

enum NGSetupResponseIEs_Value_PDUs {
    PDU_NGSetupResponseIEs_Value_UNKNOWN = 0,

    PDU_NGSetupResponseIEs_Value_AMFName = AMFName_PDU,
    PDU_NGSetupResponseIEs_Value_ServedGUAMIList = ServedGUAMIList_PDU,
    PDU_NGSetupResponseIEs_Value_RelativeAMFCapacity = RelativeAMFCapacity_PDU,
    PDU_NGSetupResponseIEs_Value_PLMNSupportList = PLMNSupportList_PDU,
    PDU_NGSetupResponseIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,
    PDU_NGSetupResponseIEs_Value_UERetentionInformation = UERetentionInformation_PDU,

    PDU_NGSetupResponseIEs_Value_MAX = INT_MAX
};

union NGSetupResponseIEs_Value_union {
    AMFName         *pdu_AMFName;  /* PDU_NGSetupResponseIEs_Value_AMFName */
    struct ServedGUAMIList_ **pdu_ServedGUAMIList;  
                             /* PDU_NGSetupResponseIEs_Value_ServedGUAMIList */
    RelativeAMFCapacity *pdu_RelativeAMFCapacity;  
                         /* PDU_NGSetupResponseIEs_Value_RelativeAMFCapacity */
    struct PLMNSupportList_ **pdu_PLMNSupportList;  
                             /* PDU_NGSetupResponseIEs_Value_PLMNSupportList */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
                      /* PDU_NGSetupResponseIEs_Value_CriticalityDiagnostics */
    UERetentionInformation *pdu_UERetentionInformation;  
                      /* PDU_NGSetupResponseIEs_Value_UERetentionInformation */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct NGSetupResponseIEs_Value {
    NGSetupResponseIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    NGSetupResponseIEs_Value_union decoded;
} NGSetupResponseIEs_Value;

// **************************************************************
//
// NG SETUP RESPONSE
//
// **************************************************************
typedef struct NGSetupResponse {
    struct _seqof226 {
        struct _seqof226 *next;
        struct _seq274 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            NGSetupResponseIEs_Value value;
        } value;
    } *protocolIEs;
} NGSetupResponse;

enum NGSetupFailureIEs_Value_PDUs {
    PDU_NGSetupFailureIEs_Value_UNKNOWN = 0,

    PDU_NGSetupFailureIEs_Value_Cause = Cause_PDU,
    PDU_NGSetupFailureIEs_Value_TimeToWait = TimeToWait_PDU,
    PDU_NGSetupFailureIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_NGSetupFailureIEs_Value_MAX = INT_MAX
};

union NGSetupFailureIEs_Value_union {
    Cause           *pdu_Cause;  /* PDU_NGSetupFailureIEs_Value_Cause */
    TimeToWait      *pdu_TimeToWait;  
                                   /* PDU_NGSetupFailureIEs_Value_TimeToWait */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
                       /* PDU_NGSetupFailureIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct NGSetupFailureIEs_Value {
    NGSetupFailureIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    NGSetupFailureIEs_Value_union decoded;
} NGSetupFailureIEs_Value;

typedef struct _seq275 {
    ProtocolIE_ID   id;
    Criticality     criticality;
    NGSetupFailureIEs_Value value;
} _seq275;

typedef struct _seqof227_ {
    struct _seqof227_ *next;
    _seq275         value;
} *_seqof227;

// **************************************************************
//
// NG SETUP FAILURE
//
// **************************************************************
typedef struct NGSetupFailure {
    struct _seqof227_ *protocolIEs;
} NGSetupFailure;

enum RANConfigurationUpdateIEs_Value_PDUs {
    PDU_RANConfigurationUpdateIEs_Value_UNKNOWN = 0,

    PDU_RANConfigurationUpdateIEs_Value_RANNodeName = RANNodeName_PDU,
    PDU_RANConfigurationUpdateIEs_Value_SupportedTAList = SupportedTAList_PDU,
    PDU_RANConfigurationUpdateIEs_Value_PagingDRX = PagingDRX_PDU,
    PDU_RANConfigurationUpdateIEs_Value_GlobalRANNodeID = GlobalRANNodeID_PDU,

    PDU_RANConfigurationUpdateIEs_Value_MAX = INT_MAX
};

union RANConfigurationUpdateIEs_Value_union {
    RANNodeName     *pdu_RANNodeName;  
                          /* PDU_RANConfigurationUpdateIEs_Value_RANNodeName */
    struct SupportedTAList_ **pdu_SupportedTAList;  
                      /* PDU_RANConfigurationUpdateIEs_Value_SupportedTAList */
    PagingDRX       *pdu_PagingDRX;  
                            /* PDU_RANConfigurationUpdateIEs_Value_PagingDRX */
    GlobalRANNodeID *pdu_GlobalRANNodeID;  
                      /* PDU_RANConfigurationUpdateIEs_Value_GlobalRANNodeID */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct RANConfigurationUpdateIEs_Value {
    RANConfigurationUpdateIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    RANConfigurationUpdateIEs_Value_union decoded;
} RANConfigurationUpdateIEs_Value;

// **************************************************************
//
// RAN Configuration Update Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// RAN CONFIGURATION UPDATE 
//
// **************************************************************
typedef struct RANConfigurationUpdate {
    struct _seqof228 {
        struct _seqof228 *next;
        struct _seq276 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            RANConfigurationUpdateIEs_Value value;
        } value;
    } *protocolIEs;
} RANConfigurationUpdate;

enum RANConfigurationUpdateAcknowledgeIEs_Value_PDUs {
    PDU_RANConfigurationUpdateAcknowledgeIEs_Value_UNKNOWN = 0,

    PDU_RANConfigurationUpdateAcknowledgeIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_RANConfigurationUpdateAcknowledgeIEs_Value_MAX = INT_MAX
};

union RANConfigurationUpdateAcknowledgeIEs_Value_union {
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
    /* PDU_RANConfigurationUpdateAcknowledgeIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct RANConfigurationUpdateAcknowledgeIEs_Value {
    RANConfigurationUpdateAcknowledgeIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    RANConfigurationUpdateAcknowledgeIEs_Value_union decoded;
} RANConfigurationUpdateAcknowledgeIEs_Value;

// **************************************************************
//
// RAN CONFIGURATION UPDATE ACKNOWLEDGE
//
// **************************************************************
typedef struct RANConfigurationUpdateAcknowledge {
    struct _seqof229 {
        struct _seqof229 *next;
        struct _seq277 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            RANConfigurationUpdateAcknowledgeIEs_Value value;
        } value;
    } *protocolIEs;
} RANConfigurationUpdateAcknowledge;

// **************************************************************
//
// RAN CONFIGURATION UPDATE FAILURE
//
// **************************************************************
typedef struct RANConfigurationUpdateFailure {
    struct _seqof227_ *protocolIEs;
} RANConfigurationUpdateFailure;

enum AMFConfigurationUpdateIEs_Value_PDUs {
    PDU_AMFConfigurationUpdateIEs_Value_UNKNOWN = 0,

    PDU_AMFConfigurationUpdateIEs_Value_AMFName = AMFName_PDU,
    PDU_AMFConfigurationUpdateIEs_Value_ServedGUAMIList = ServedGUAMIList_PDU,
    PDU_AMFConfigurationUpdateIEs_Value_RelativeAMFCapacity = RelativeAMFCapacity_PDU,
    PDU_AMFConfigurationUpdateIEs_Value_PLMNSupportList = PLMNSupportList_PDU,
    PDU_AMFConfigurationUpdateIEs_Value_AMF_TNLAssociationToAddList = AMF_TNLAssociationToAddList_PDU,
    PDU_AMFConfigurationUpdateIEs_Value_AMF_TNLAssociationToRemoveList = AMF_TNLAssociationToRemoveList_PDU,
    PDU_AMFConfigurationUpdateIEs_Value_AMF_TNLAssociationToUpdateList = AMF_TNLAssociationToUpdateList_PDU,

    PDU_AMFConfigurationUpdateIEs_Value_MAX = INT_MAX
};

union AMFConfigurationUpdateIEs_Value_union {
    AMFName         *pdu_AMFName;  
                              /* PDU_AMFConfigurationUpdateIEs_Value_AMFName */
    struct ServedGUAMIList_ **pdu_ServedGUAMIList;  
                      /* PDU_AMFConfigurationUpdateIEs_Value_ServedGUAMIList */
    RelativeAMFCapacity *pdu_RelativeAMFCapacity;  
                  /* PDU_AMFConfigurationUpdateIEs_Value_RelativeAMFCapacity */
    struct PLMNSupportList_ **pdu_PLMNSupportList;  
                      /* PDU_AMFConfigurationUpdateIEs_Value_PLMNSupportList */
    struct AMF_TNLAssociationToAddList_ **pdu_AMF_TNLAssociationToAddList;  
                                        /* PDU_AMFConfigurationUpdateIEs_Value_AMF_TNLAssociationToAddList */
    struct AMF_TNLAssociationToRemoveList_ **pdu_AMF_TNLAssociationToRemoveList;                                        /* PDU_AMFConfigurationUpdateIEs_Value_AMF_TNLAssociationToRemoveList */
    struct AMF_TNLAssociationToUpdateList_ **pdu_AMF_TNLAssociationToUpdateList;                                        /* PDU_AMFConfigurationUpdateIEs_Value_AMF_TNLAssociationToUpdateList */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct AMFConfigurationUpdateIEs_Value {
    AMFConfigurationUpdateIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    AMFConfigurationUpdateIEs_Value_union decoded;
} AMFConfigurationUpdateIEs_Value;

// **************************************************************
//
// AMF Configuration Update Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// AMF CONFIGURATION UPDATE 
//
// **************************************************************
typedef struct AMFConfigurationUpdate {
    struct _seqof231 {
        struct _seqof231 *next;
        struct _seq279 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            AMFConfigurationUpdateIEs_Value value;
        } value;
    } *protocolIEs;
} AMFConfigurationUpdate;

enum AMFConfigurationUpdateAcknowledgeIEs_Value_PDUs {
    PDU_AMFConfigurationUpdateAcknowledgeIEs_Value_UNKNOWN = 0,

    PDU_AMFConfigurationUpdateAcknowledgeIEs_Value_AMF_TNLAssociationSetupList = AMF_TNLAssociationSetupList_PDU,
    PDU_AMFConfigurationUpdateAcknowledgeIEs_Value_TNLAssociationList = TNLAssociationList_PDU,
    PDU_AMFConfigurationUpdateAcknowledgeIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_AMFConfigurationUpdateAcknowledgeIEs_Value_MAX = INT_MAX
};

union AMFConfigurationUpdateAcknowledgeIEs_Value_union {
    struct AMF_TNLAssociationSetupList_ **pdu_AMF_TNLAssociationSetupList;  
                                        /* PDU_AMFConfigurationUpdateAcknowledgeIEs_Value_AMF_TNLAssociationSetupList */
    struct TNLAssociationList_ **pdu_TNLAssociationList;  
        /* PDU_AMFConfigurationUpdateAcknowledgeIEs_Value_TNLAssociationList */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
    /* PDU_AMFConfigurationUpdateAcknowledgeIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct AMFConfigurationUpdateAcknowledgeIEs_Value {
    AMFConfigurationUpdateAcknowledgeIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    AMFConfigurationUpdateAcknowledgeIEs_Value_union decoded;
} AMFConfigurationUpdateAcknowledgeIEs_Value;

// **************************************************************
//
// AMF CONFIGURATION UPDATE ACKNOWLEDGE
//
// **************************************************************
typedef struct AMFConfigurationUpdateAcknowledge {
    struct _seqof232 {
        struct _seqof232 *next;
        struct _seq280 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            AMFConfigurationUpdateAcknowledgeIEs_Value value;
        } value;
    } *protocolIEs;
} AMFConfigurationUpdateAcknowledge;

// **************************************************************
//
// AMF CONFIGURATION UPDATE FAILURE
//
// **************************************************************
typedef struct AMFConfigurationUpdateFailure {
    struct _seqof227_ *protocolIEs;
} AMFConfigurationUpdateFailure;

enum AMFStatusIndicationIEs_Value_PDUs {
    PDU_AMFStatusIndicationIEs_Value_UNKNOWN = 0,

    PDU_AMFStatusIndicationIEs_Value_UnavailableGUAMIList = UnavailableGUAMIList_PDU,

    PDU_AMFStatusIndicationIEs_Value_MAX = INT_MAX
};

union AMFStatusIndicationIEs_Value_union {
    struct UnavailableGUAMIList_ **pdu_UnavailableGUAMIList;  
                    /* PDU_AMFStatusIndicationIEs_Value_UnavailableGUAMIList */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct AMFStatusIndicationIEs_Value {
    AMFStatusIndicationIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    AMFStatusIndicationIEs_Value_union decoded;
} AMFStatusIndicationIEs_Value;

// **************************************************************
//
// AMF Status Indication Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// AMF STATUS INDICATION
//
// **************************************************************
typedef struct AMFStatusIndication {
    struct _seqof234 {
        struct _seqof234 *next;
        struct _seq282 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            AMFStatusIndicationIEs_Value value;
        } value;
    } *protocolIEs;
} AMFStatusIndication;

enum NGResetIEs_Value_PDUs {
    PDU_NGResetIEs_Value_UNKNOWN   = 0,

    PDU_NGResetIEs_Value_Cause     = Cause_PDU,
    PDU_NGResetIEs_Value_ResetType = ResetType_PDU,

    PDU_NGResetIEs_Value_MAX       = INT_MAX
};

union NGResetIEs_Value_union {
    Cause           *pdu_Cause;  /* PDU_NGResetIEs_Value_Cause */
    ResetType       *pdu_ResetType;  /* PDU_NGResetIEs_Value_ResetType */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct NGResetIEs_Value {
    NGResetIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    NGResetIEs_Value_union decoded;
} NGResetIEs_Value;

// **************************************************************
//
// NG Reset Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// NG RESET
//
// **************************************************************
typedef struct NGReset {
    struct _seqof235 {
        struct _seqof235 *next;
        struct _seq283 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            NGResetIEs_Value value;
        } value;
    } *protocolIEs;
} NGReset;

enum NGResetAcknowledgeIEs_Value_PDUs {
    PDU_NGResetAcknowledgeIEs_Value_UNKNOWN = 0,

    PDU_NGResetAcknowledgeIEs_Value_UE_associatedLogicalNG_connectionList = UE_associatedLogicalNG_connectionList_PDU,
    PDU_NGResetAcknowledgeIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_NGResetAcknowledgeIEs_Value_MAX = INT_MAX
};

union NGResetAcknowledgeIEs_Value_union {
    struct UE_associatedLogicalNG_connectionList_ **pdu_UE_associatedLogicalNG_connectionList;                          /* PDU_NGResetAcknowledgeIEs_Value_UE_associatedLogicalNG_connectionList */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
                   /* PDU_NGResetAcknowledgeIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct NGResetAcknowledgeIEs_Value {
    NGResetAcknowledgeIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    NGResetAcknowledgeIEs_Value_union decoded;
} NGResetAcknowledgeIEs_Value;

// **************************************************************
//
// NG RESET ACKNOWLEDGE
//
// **************************************************************
typedef struct NGResetAcknowledge {
    struct _seqof236 {
        struct _seqof236 *next;
        struct _seq284 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            NGResetAcknowledgeIEs_Value value;
        } value;
    } *protocolIEs;
} NGResetAcknowledge;

// **************************************************************
//
// Error Indication Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// ERROR INDICATION
//
// **************************************************************
typedef struct ErrorIndication {
    struct _seqof203_ *protocolIEs;
} ErrorIndication;

enum OverloadStartIEs_Value_PDUs {
    PDU_OverloadStartIEs_Value_UNKNOWN = 0,

    PDU_OverloadStartIEs_Value_OverloadResponse = OverloadResponse_PDU,
    PDU_OverloadStartIEs_Value_TrafficLoadReductionIndication = TrafficLoadReductionIndication_PDU,
    PDU_OverloadStartIEs_Value_OverloadStartNSSAIList = OverloadStartNSSAIList_PDU,

    PDU_OverloadStartIEs_Value_MAX = INT_MAX
};

union OverloadStartIEs_Value_union {
    OverloadResponse *pdu_OverloadResponse;  
                              /* PDU_OverloadStartIEs_Value_OverloadResponse */
    TrafficLoadReductionIndication *pdu_TrafficLoadReductionIndication;  
                /* PDU_OverloadStartIEs_Value_TrafficLoadReductionIndication */
    struct OverloadStartNSSAIList_ **pdu_OverloadStartNSSAIList;  
                        /* PDU_OverloadStartIEs_Value_OverloadStartNSSAIList */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct OverloadStartIEs_Value {
    OverloadStartIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    OverloadStartIEs_Value_union decoded;
} OverloadStartIEs_Value;

// **************************************************************
//
// OVERLOAD START
//
// **************************************************************
typedef struct OverloadStart {
    struct _seqof238 {
        struct _seqof238 *next;
        struct _seq286 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            OverloadStartIEs_Value value;
        } value;
    } *protocolIEs;
} OverloadStart;

// **************************************************************
//
// OVERLOAD STOP
//
// **************************************************************
typedef struct OverloadStop {
    struct _seqof239 {
        struct _seqof239 *next;
        struct _seq287 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PrivateMessageIEs_Value value;
        } value;
    } *protocolIEs;
} OverloadStop;

enum UplinkRANConfigurationTransferIEs_Value_PDUs {
    PDU_UplinkRANConfigurationTransferIEs_Value_UNKNOWN = 0,

    PDU_UplinkRANConfigurationTransferIEs_Value_SONConfigurationTransfer = SONConfigurationTransfer_PDU,
    PDU_UplinkRANConfigurationTransferIEs_Value_EN_DCSONConfigurationTransfer = EN_DCSONConfigurationTransfer_PDU,

    PDU_UplinkRANConfigurationTransferIEs_Value_MAX = INT_MAX
};

union UplinkRANConfigurationTransferIEs_Value_union {
    SONConfigurationTransfer *pdu_SONConfigurationTransfer;  
     /* PDU_UplinkRANConfigurationTransferIEs_Value_SONConfigurationTransfer */
    EN_DCSONConfigurationTransfer *pdu_EN_DCSONConfigurationTransfer;  
/* PDU_UplinkRANConfigurationTransferIEs_Value_EN_DCSONConfigurationTransfer */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct UplinkRANConfigurationTransferIEs_Value {
    UplinkRANConfigurationTransferIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    UplinkRANConfigurationTransferIEs_Value_union decoded;
} UplinkRANConfigurationTransferIEs_Value;

typedef struct _seq288 {
    ProtocolIE_ID   id;
    Criticality     criticality;
    UplinkRANConfigurationTransferIEs_Value value;
} _seq288;

typedef struct _seqof240_ {
    struct _seqof240_ *next;
    _seq288         value;
} *_seqof240;

// **************************************************************
//
// CONFIGURATION TRANSFER ELEMENTARY PROCEDURES
//
// **************************************************************

// **************************************************************
//
// UPLINK RAN CONFIGURATION TRANSFER
//
// **************************************************************
typedef struct UplinkRANConfigurationTransfer {
    struct _seqof240_ *protocolIEs;
} UplinkRANConfigurationTransfer;

// **************************************************************
//
// DOWNLINK RAN CONFIGURATION TRANSFER
//
// **************************************************************
typedef struct DownlinkRANConfigurationTransfer {
    struct _seqof240_ *protocolIEs;
} DownlinkRANConfigurationTransfer;

enum WriteReplaceWarningRequestIEs_Value_PDUs {
    PDU_WriteReplaceWarningRequestIEs_Value_UNKNOWN = 0,

    PDU_WriteReplaceWarningRequestIEs_Value_MessageIdentifier = MessageIdentifier_PDU,
    PDU_WriteReplaceWarningRequestIEs_Value_SerialNumber = SerialNumber_PDU,
    PDU_WriteReplaceWarningRequestIEs_Value_WarningAreaList = WarningAreaList_PDU,
    PDU_WriteReplaceWarningRequestIEs_Value_RepetitionPeriod = RepetitionPeriod_PDU,
    PDU_WriteReplaceWarningRequestIEs_Value_NumberOfBroadcastsRequested = NumberOfBroadcastsRequested_PDU,
    PDU_WriteReplaceWarningRequestIEs_Value_WarningType = WarningType_PDU,
    PDU_WriteReplaceWarningRequestIEs_Value_WarningSecurityInfo = WarningSecurityInfo_PDU,
    PDU_WriteReplaceWarningRequestIEs_Value_DataCodingScheme = DataCodingScheme_PDU,
    PDU_WriteReplaceWarningRequestIEs_Value_WarningMessageContents = WarningMessageContents_PDU,
    PDU_WriteReplaceWarningRequestIEs_Value_ConcurrentWarningMessageInd = ConcurrentWarningMessageInd_PDU,
    PDU_WriteReplaceWarningRequestIEs_Value_WarningAreaCoordinates = WarningAreaCoordinates_PDU,

    PDU_WriteReplaceWarningRequestIEs_Value_MAX = INT_MAX
};

union WriteReplaceWarningRequestIEs_Value_union {
    MessageIdentifier *pdu_MessageIdentifier;  
                /* PDU_WriteReplaceWarningRequestIEs_Value_MessageIdentifier */
    SerialNumber    *pdu_SerialNumber;  
                     /* PDU_WriteReplaceWarningRequestIEs_Value_SerialNumber */
    WarningAreaList *pdu_WarningAreaList;  
                  /* PDU_WriteReplaceWarningRequestIEs_Value_WarningAreaList */
    RepetitionPeriod *pdu_RepetitionPeriod;  
                 /* PDU_WriteReplaceWarningRequestIEs_Value_RepetitionPeriod */
    NumberOfBroadcastsRequested *pdu_NumberOfBroadcastsRequested;  
      /* PDU_WriteReplaceWarningRequestIEs_Value_NumberOfBroadcastsRequested */
    WarningType     *pdu_WarningType;  
                      /* PDU_WriteReplaceWarningRequestIEs_Value_WarningType */
    WarningSecurityInfo *pdu_WarningSecurityInfo;  
              /* PDU_WriteReplaceWarningRequestIEs_Value_WarningSecurityInfo */
    DataCodingScheme *pdu_DataCodingScheme;  
                 /* PDU_WriteReplaceWarningRequestIEs_Value_DataCodingScheme */
    WarningMessageContents *pdu_WarningMessageContents;  
           /* PDU_WriteReplaceWarningRequestIEs_Value_WarningMessageContents */
    ConcurrentWarningMessageInd *pdu_ConcurrentWarningMessageInd;  
      /* PDU_WriteReplaceWarningRequestIEs_Value_ConcurrentWarningMessageInd */
    WarningAreaCoordinates *pdu_WarningAreaCoordinates;  
           /* PDU_WriteReplaceWarningRequestIEs_Value_WarningAreaCoordinates */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct WriteReplaceWarningRequestIEs_Value {
    WriteReplaceWarningRequestIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    WriteReplaceWarningRequestIEs_Value_union decoded;
} WriteReplaceWarningRequestIEs_Value;

// **************************************************************
//
// WARNING MESSAGE TRANSMISSION ELEMENTARY PROCEDURES 
//
// **************************************************************

// **************************************************************
//
// Write-Replace Warning Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// WRITE-REPLACE WARNING REQUEST
//
// **************************************************************
typedef struct WriteReplaceWarningRequest {
    struct _seqof242 {
        struct _seqof242 *next;
        struct _seq290 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            WriteReplaceWarningRequestIEs_Value value;
        } value;
    } *protocolIEs;
} WriteReplaceWarningRequest;

enum WriteReplaceWarningResponseIEs_Value_PDUs {
    PDU_WriteReplaceWarningResponseIEs_Value_UNKNOWN = 0,

    PDU_WriteReplaceWarningResponseIEs_Value_MessageIdentifier = MessageIdentifier_PDU,
    PDU_WriteReplaceWarningResponseIEs_Value_SerialNumber = SerialNumber_PDU,
    PDU_WriteReplaceWarningResponseIEs_Value_BroadcastCompletedAreaList = BroadcastCompletedAreaList_PDU,
    PDU_WriteReplaceWarningResponseIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_WriteReplaceWarningResponseIEs_Value_MAX = INT_MAX
};

union WriteReplaceWarningResponseIEs_Value_union {
    MessageIdentifier *pdu_MessageIdentifier;  
               /* PDU_WriteReplaceWarningResponseIEs_Value_MessageIdentifier */
    SerialNumber    *pdu_SerialNumber;  
                    /* PDU_WriteReplaceWarningResponseIEs_Value_SerialNumber */
    BroadcastCompletedAreaList *pdu_BroadcastCompletedAreaList;  
      /* PDU_WriteReplaceWarningResponseIEs_Value_BroadcastCompletedAreaList */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
          /* PDU_WriteReplaceWarningResponseIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct WriteReplaceWarningResponseIEs_Value {
    WriteReplaceWarningResponseIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    WriteReplaceWarningResponseIEs_Value_union decoded;
} WriteReplaceWarningResponseIEs_Value;

// **************************************************************
//
// WRITE-REPLACE WARNING RESPONSE
//
// **************************************************************
typedef struct WriteReplaceWarningResponse {
    struct _seqof243 {
        struct _seqof243 *next;
        struct _seq291 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            WriteReplaceWarningResponseIEs_Value value;
        } value;
    } *protocolIEs;
} WriteReplaceWarningResponse;

enum PWSCancelRequestIEs_Value_PDUs {
    PDU_PWSCancelRequestIEs_Value_UNKNOWN = 0,

    PDU_PWSCancelRequestIEs_Value_MessageIdentifier = MessageIdentifier_PDU,
    PDU_PWSCancelRequestIEs_Value_SerialNumber = SerialNumber_PDU,
    PDU_PWSCancelRequestIEs_Value_WarningAreaList = WarningAreaList_PDU,
    PDU_PWSCancelRequestIEs_Value_CancelAllWarningMessages = CancelAllWarningMessages_PDU,

    PDU_PWSCancelRequestIEs_Value_MAX = INT_MAX
};

union PWSCancelRequestIEs_Value_union {
    MessageIdentifier *pdu_MessageIdentifier;  
                          /* PDU_PWSCancelRequestIEs_Value_MessageIdentifier */
    SerialNumber    *pdu_SerialNumber;  
                               /* PDU_PWSCancelRequestIEs_Value_SerialNumber */
    WarningAreaList *pdu_WarningAreaList;  
                            /* PDU_PWSCancelRequestIEs_Value_WarningAreaList */
    CancelAllWarningMessages *pdu_CancelAllWarningMessages;  
                   /* PDU_PWSCancelRequestIEs_Value_CancelAllWarningMessages */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PWSCancelRequestIEs_Value {
    PWSCancelRequestIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PWSCancelRequestIEs_Value_union decoded;
} PWSCancelRequestIEs_Value;

// **************************************************************
//
// PWS Cancel Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// PWS CANCEL REQUEST
//
// **************************************************************
typedef struct PWSCancelRequest {
    struct _seqof244 {
        struct _seqof244 *next;
        struct _seq292 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PWSCancelRequestIEs_Value value;
        } value;
    } *protocolIEs;
} PWSCancelRequest;

enum PWSCancelResponseIEs_Value_PDUs {
    PDU_PWSCancelResponseIEs_Value_UNKNOWN = 0,

    PDU_PWSCancelResponseIEs_Value_MessageIdentifier = MessageIdentifier_PDU,
    PDU_PWSCancelResponseIEs_Value_SerialNumber = SerialNumber_PDU,
    PDU_PWSCancelResponseIEs_Value_BroadcastCancelledAreaList = BroadcastCancelledAreaList_PDU,
    PDU_PWSCancelResponseIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_PWSCancelResponseIEs_Value_MAX = INT_MAX
};

union PWSCancelResponseIEs_Value_union {
    MessageIdentifier *pdu_MessageIdentifier;  
                         /* PDU_PWSCancelResponseIEs_Value_MessageIdentifier */
    SerialNumber    *pdu_SerialNumber;  
                              /* PDU_PWSCancelResponseIEs_Value_SerialNumber */
    BroadcastCancelledAreaList *pdu_BroadcastCancelledAreaList;  
                /* PDU_PWSCancelResponseIEs_Value_BroadcastCancelledAreaList */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
                    /* PDU_PWSCancelResponseIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PWSCancelResponseIEs_Value {
    PWSCancelResponseIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PWSCancelResponseIEs_Value_union decoded;
} PWSCancelResponseIEs_Value;

// **************************************************************
//
// PWS CANCEL RESPONSE
//
// **************************************************************
typedef struct PWSCancelResponse {
    struct _seqof245 {
        struct _seqof245 *next;
        struct _seq293 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PWSCancelResponseIEs_Value value;
        } value;
    } *protocolIEs;
} PWSCancelResponse;

enum PWSRestartIndicationIEs_Value_PDUs {
    PDU_PWSRestartIndicationIEs_Value_UNKNOWN = 0,

    PDU_PWSRestartIndicationIEs_Value_CellIDListForRestart = CellIDListForRestart_PDU,
    PDU_PWSRestartIndicationIEs_Value_GlobalRANNodeID = GlobalRANNodeID_PDU,
    PDU_PWSRestartIndicationIEs_Value_TAIListForRestart = TAIListForRestart_PDU,
    PDU_PWSRestartIndicationIEs_Value_EmergencyAreaIDListForRestart = EmergencyAreaIDListForRestart_PDU,

    PDU_PWSRestartIndicationIEs_Value_MAX = INT_MAX
};

union PWSRestartIndicationIEs_Value_union {
    CellIDListForRestart *pdu_CellIDListForRestart;  
                   /* PDU_PWSRestartIndicationIEs_Value_CellIDListForRestart */
    GlobalRANNodeID *pdu_GlobalRANNodeID;  
                        /* PDU_PWSRestartIndicationIEs_Value_GlobalRANNodeID */
    struct TAIListForRestart_ **pdu_TAIListForRestart;  
                      /* PDU_PWSRestartIndicationIEs_Value_TAIListForRestart */
    struct EmergencyAreaIDListForRestart_ **pdu_EmergencyAreaIDListForRestart;                                          /* PDU_PWSRestartIndicationIEs_Value_EmergencyAreaIDListForRestart */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PWSRestartIndicationIEs_Value {
    PWSRestartIndicationIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PWSRestartIndicationIEs_Value_union decoded;
} PWSRestartIndicationIEs_Value;

// **************************************************************
//
// PWS Restart Indication Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// PWS RESTART INDICATION
//
// **************************************************************
typedef struct PWSRestartIndication {
    struct _seqof246 {
        struct _seqof246 *next;
        struct _seq294 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PWSRestartIndicationIEs_Value value;
        } value;
    } *protocolIEs;
} PWSRestartIndication;

enum PWSFailureIndicationIEs_Value_PDUs {
    PDU_PWSFailureIndicationIEs_Value_UNKNOWN = 0,

    PDU_PWSFailureIndicationIEs_Value_PWSFailedCellIDList = PWSFailedCellIDList_PDU,
    PDU_PWSFailureIndicationIEs_Value_GlobalRANNodeID = GlobalRANNodeID_PDU,

    PDU_PWSFailureIndicationIEs_Value_MAX = INT_MAX
};

union PWSFailureIndicationIEs_Value_union {
    PWSFailedCellIDList *pdu_PWSFailedCellIDList;  
                    /* PDU_PWSFailureIndicationIEs_Value_PWSFailedCellIDList */
    GlobalRANNodeID *pdu_GlobalRANNodeID;  
                        /* PDU_PWSFailureIndicationIEs_Value_GlobalRANNodeID */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct PWSFailureIndicationIEs_Value {
    PWSFailureIndicationIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    PWSFailureIndicationIEs_Value_union decoded;
} PWSFailureIndicationIEs_Value;

// **************************************************************
//
// PWS Failure Indication Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// PWS FAILURE INDICATION
//
// **************************************************************
typedef struct PWSFailureIndication {
    struct _seqof247 {
        struct _seqof247 *next;
        struct _seq295 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            PWSFailureIndicationIEs_Value value;
        } value;
    } *protocolIEs;
} PWSFailureIndication;

enum DownlinkUEAssociatedNRPPaTransportIEs_Value_PDUs {
    PDU_DownlinkUEAssociatedNRPPaTransportIEs_Value_UNKNOWN = 0,

    PDU_DownlinkUEAssociatedNRPPaTransportIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_DownlinkUEAssociatedNRPPaTransportIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_DownlinkUEAssociatedNRPPaTransportIEs_Value_RoutingID = RoutingID_PDU,
    PDU_DownlinkUEAssociatedNRPPaTransportIEs_Value_NRPPa_PDU = NRPPa_PDU_PDU,

    PDU_DownlinkUEAssociatedNRPPaTransportIEs_Value_MAX = INT_MAX
};

union DownlinkUEAssociatedNRPPaTransportIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
           /* PDU_DownlinkUEAssociatedNRPPaTransportIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
           /* PDU_DownlinkUEAssociatedNRPPaTransportIEs_Value_RAN_UE_NGAP_ID */
    RoutingID       *pdu_RoutingID;  
                /* PDU_DownlinkUEAssociatedNRPPaTransportIEs_Value_RoutingID */
    NRPPa_PDU       *pdu_NRPPa_PDU;  
                /* PDU_DownlinkUEAssociatedNRPPaTransportIEs_Value_NRPPa_PDU */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct DownlinkUEAssociatedNRPPaTransportIEs_Value {
    DownlinkUEAssociatedNRPPaTransportIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    DownlinkUEAssociatedNRPPaTransportIEs_Value_union decoded;
} DownlinkUEAssociatedNRPPaTransportIEs_Value;

typedef struct _seq296 {
    ProtocolIE_ID   id;
    Criticality     criticality;
    DownlinkUEAssociatedNRPPaTransportIEs_Value value;
} _seq296;

typedef struct _seqof248_ {
    struct _seqof248_ *next;
    _seq296         value;
} *_seqof248;

// **************************************************************
//
// NRPPA TRANSPORT ELEMENTARY PROCEDURES
//
// **************************************************************

// **************************************************************
//
// DOWNLINK UE ASSOCIATED NRPPA TRANSPORT
//
// **************************************************************
typedef struct DownlinkUEAssociatedNRPPaTransport {
    struct _seqof248_ *protocolIEs;
} DownlinkUEAssociatedNRPPaTransport;

// **************************************************************
//
// UPLINK UE ASSOCIATED NRPPA TRANSPORT
//
// **************************************************************
typedef struct UplinkUEAssociatedNRPPaTransport {
    struct _seqof248_ *protocolIEs;
} UplinkUEAssociatedNRPPaTransport;

enum DownlinkNonUEAssociatedNRPPaTransportIEs_Value_PDUs {
    PDU_DownlinkNonUEAssociatedNRPPaTransportIEs_Value_UNKNOWN = 0,

    PDU_DownlinkNonUEAssociatedNRPPaTransportIEs_Value_RoutingID = RoutingID_PDU,
    PDU_DownlinkNonUEAssociatedNRPPaTransportIEs_Value_NRPPa_PDU = NRPPa_PDU_PDU,

    PDU_DownlinkNonUEAssociatedNRPPaTransportIEs_Value_MAX = INT_MAX
};

union DownlinkNonUEAssociatedNRPPaTransportIEs_Value_union {
    RoutingID       *pdu_RoutingID;  
             /* PDU_DownlinkNonUEAssociatedNRPPaTransportIEs_Value_RoutingID */
    NRPPa_PDU       *pdu_NRPPa_PDU;  
             /* PDU_DownlinkNonUEAssociatedNRPPaTransportIEs_Value_NRPPa_PDU */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct DownlinkNonUEAssociatedNRPPaTransportIEs_Value {
    DownlinkNonUEAssociatedNRPPaTransportIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    DownlinkNonUEAssociatedNRPPaTransportIEs_Value_union decoded;
} DownlinkNonUEAssociatedNRPPaTransportIEs_Value;

typedef struct _seq298 {
    ProtocolIE_ID   id;
    Criticality     criticality;
    DownlinkNonUEAssociatedNRPPaTransportIEs_Value value;
} _seq298;

typedef struct _seqof250_ {
    struct _seqof250_ *next;
    _seq298         value;
} *_seqof250;

// **************************************************************
//
// DOWNLINK NON UE ASSOCIATED NRPPA TRANSPORT
//
// **************************************************************
typedef struct DownlinkNonUEAssociatedNRPPaTransport {
    struct _seqof250_ *protocolIEs;
} DownlinkNonUEAssociatedNRPPaTransport;

// **************************************************************
//
// UPLINK NON UE ASSOCIATED NRPPA TRANSPORT
//
// **************************************************************
typedef struct UplinkNonUEAssociatedNRPPaTransport {
    struct _seqof250_ *protocolIEs;
} UplinkNonUEAssociatedNRPPaTransport;

enum TraceStartIEs_Value_PDUs {
    PDU_TraceStartIEs_Value_UNKNOWN = 0,

    PDU_TraceStartIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_TraceStartIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_TraceStartIEs_Value_TraceActivation = TraceActivation_PDU,

    PDU_TraceStartIEs_Value_MAX    = INT_MAX
};

union TraceStartIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                                   /* PDU_TraceStartIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                                   /* PDU_TraceStartIEs_Value_RAN_UE_NGAP_ID */
    TraceActivation *pdu_TraceActivation;  
                                  /* PDU_TraceStartIEs_Value_TraceActivation */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct TraceStartIEs_Value {
    TraceStartIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    TraceStartIEs_Value_union decoded;
} TraceStartIEs_Value;

// **************************************************************
//
// TRACE ELEMENTARY PROCEDURES
//
// **************************************************************

// **************************************************************
//
// TRACE START
//
// **************************************************************
typedef struct TraceStart {
    struct _seqof252 {
        struct _seqof252 *next;
        struct _seq300 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            TraceStartIEs_Value value;
        } value;
    } *protocolIEs;
} TraceStart;

enum TraceFailureIndicationIEs_Value_PDUs {
    PDU_TraceFailureIndicationIEs_Value_UNKNOWN = 0,

    PDU_TraceFailureIndicationIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_TraceFailureIndicationIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_TraceFailureIndicationIEs_Value_NGRANTraceID = NGRANTraceID_PDU,
    PDU_TraceFailureIndicationIEs_Value_Cause = Cause_PDU,

    PDU_TraceFailureIndicationIEs_Value_MAX = INT_MAX
};

union TraceFailureIndicationIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                       /* PDU_TraceFailureIndicationIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                       /* PDU_TraceFailureIndicationIEs_Value_RAN_UE_NGAP_ID */
    NGRANTraceID    *pdu_NGRANTraceID;  
                         /* PDU_TraceFailureIndicationIEs_Value_NGRANTraceID */
    Cause           *pdu_Cause;  
                                /* PDU_TraceFailureIndicationIEs_Value_Cause */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct TraceFailureIndicationIEs_Value {
    TraceFailureIndicationIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    TraceFailureIndicationIEs_Value_union decoded;
} TraceFailureIndicationIEs_Value;

// **************************************************************
//
// TRACE FAILURE INDICATION
//
// **************************************************************
typedef struct TraceFailureIndication {
    struct _seqof253 {
        struct _seqof253 *next;
        struct _seq301 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            TraceFailureIndicationIEs_Value value;
        } value;
    } *protocolIEs;
} TraceFailureIndication;

enum DeactivateTraceIEs_Value_PDUs {
    PDU_DeactivateTraceIEs_Value_UNKNOWN = 0,

    PDU_DeactivateTraceIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_DeactivateTraceIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_DeactivateTraceIEs_Value_NGRANTraceID = NGRANTraceID_PDU,

    PDU_DeactivateTraceIEs_Value_MAX = INT_MAX
};

union DeactivateTraceIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                              /* PDU_DeactivateTraceIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                              /* PDU_DeactivateTraceIEs_Value_RAN_UE_NGAP_ID */
    NGRANTraceID    *pdu_NGRANTraceID;  
                                /* PDU_DeactivateTraceIEs_Value_NGRANTraceID */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct DeactivateTraceIEs_Value {
    DeactivateTraceIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    DeactivateTraceIEs_Value_union decoded;
} DeactivateTraceIEs_Value;

// **************************************************************
//
// DEACTIVATE TRACE
//
// **************************************************************
typedef struct DeactivateTrace {
    struct _seqof254 {
        struct _seqof254 *next;
        struct _seq302 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            DeactivateTraceIEs_Value value;
        } value;
    } *protocolIEs;
} DeactivateTrace;

enum CellTrafficTraceIEs_Value_PDUs {
    PDU_CellTrafficTraceIEs_Value_UNKNOWN = 0,

    PDU_CellTrafficTraceIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_CellTrafficTraceIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_CellTrafficTraceIEs_Value_NGRANTraceID = NGRANTraceID_PDU,
    PDU_CellTrafficTraceIEs_Value_NGRAN_CGI = NGRAN_CGI_PDU,
    PDU_CellTrafficTraceIEs_Value_TransportLayerAddress = TransportLayerAddress_PDU,

    PDU_CellTrafficTraceIEs_Value_MAX = INT_MAX
};

union CellTrafficTraceIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                             /* PDU_CellTrafficTraceIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                             /* PDU_CellTrafficTraceIEs_Value_RAN_UE_NGAP_ID */
    NGRANTraceID    *pdu_NGRANTraceID;  
                               /* PDU_CellTrafficTraceIEs_Value_NGRANTraceID */
    NGRAN_CGI       *pdu_NGRAN_CGI;  
                                  /* PDU_CellTrafficTraceIEs_Value_NGRAN_CGI */
    TransportLayerAddress *pdu_TransportLayerAddress;  
                      /* PDU_CellTrafficTraceIEs_Value_TransportLayerAddress */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct CellTrafficTraceIEs_Value {
    CellTrafficTraceIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    CellTrafficTraceIEs_Value_union decoded;
} CellTrafficTraceIEs_Value;

// **************************************************************
//
// CELL TRAFFIC TRACE
//
// **************************************************************
typedef struct CellTrafficTrace {
    struct _seqof255 {
        struct _seqof255 *next;
        struct _seq303 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            CellTrafficTraceIEs_Value value;
        } value;
    } *protocolIEs;
} CellTrafficTrace;

enum LocationReportingControlIEs_Value_PDUs {
    PDU_LocationReportingControlIEs_Value_UNKNOWN = 0,

    PDU_LocationReportingControlIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_LocationReportingControlIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_LocationReportingControlIEs_Value_LocationReportingRequestType = LocationReportingRequestType_PDU,

    PDU_LocationReportingControlIEs_Value_MAX = INT_MAX
};

union LocationReportingControlIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                     /* PDU_LocationReportingControlIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                     /* PDU_LocationReportingControlIEs_Value_RAN_UE_NGAP_ID */
    LocationReportingRequestType *pdu_LocationReportingRequestType;  
       /* PDU_LocationReportingControlIEs_Value_LocationReportingRequestType */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct LocationReportingControlIEs_Value {
    LocationReportingControlIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    LocationReportingControlIEs_Value_union decoded;
} LocationReportingControlIEs_Value;

// **************************************************************
//
// LOCATION REPORTING ELEMENTARY PROCEDURES
//
// **************************************************************

// **************************************************************
//
// LOCATION REPORTING CONTROL
//
// **************************************************************
typedef struct LocationReportingControl {
    struct _seqof256 {
        struct _seqof256 *next;
        struct _seq304 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            LocationReportingControlIEs_Value value;
        } value;
    } *protocolIEs;
} LocationReportingControl;

// **************************************************************
//
// LOCATION REPORTING FAILURE INDICATION
//
// **************************************************************
typedef struct LocationReportingFailureIndication {
    struct _seqof215_ *protocolIEs;
} LocationReportingFailureIndication;

enum LocationReportIEs_Value_PDUs {
    PDU_LocationReportIEs_Value_UNKNOWN = 0,

    PDU_LocationReportIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_LocationReportIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_LocationReportIEs_Value_UserLocationInformation = UserLocationInformation_PDU,
    PDU_LocationReportIEs_Value_UEPresenceInAreaOfInterestList = UEPresenceInAreaOfInterestList_PDU,
    PDU_LocationReportIEs_Value_LocationReportingRequestType = LocationReportingRequestType_PDU,

    PDU_LocationReportIEs_Value_MAX = INT_MAX
};

union LocationReportIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                               /* PDU_LocationReportIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                               /* PDU_LocationReportIEs_Value_RAN_UE_NGAP_ID */
    UserLocationInformation *pdu_UserLocationInformation;  
                      /* PDU_LocationReportIEs_Value_UserLocationInformation */
    struct UEPresenceInAreaOfInterestList_ **pdu_UEPresenceInAreaOfInterestList;                                        /* PDU_LocationReportIEs_Value_UEPresenceInAreaOfInterestList */
    LocationReportingRequestType *pdu_LocationReportingRequestType;  
                 /* PDU_LocationReportIEs_Value_LocationReportingRequestType */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct LocationReportIEs_Value {
    LocationReportIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    LocationReportIEs_Value_union decoded;
} LocationReportIEs_Value;

// **************************************************************
//
// LOCATION REPORT
//
// **************************************************************
typedef struct LocationReport {
    struct _seqof258 {
        struct _seqof258 *next;
        struct _seq306 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            LocationReportIEs_Value value;
        } value;
    } *protocolIEs;
} LocationReport;

enum UETNLABindingReleaseRequestIEs_Value_PDUs {
    PDU_UETNLABindingReleaseRequestIEs_Value_UNKNOWN = 0,

    PDU_UETNLABindingReleaseRequestIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_UETNLABindingReleaseRequestIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,

    PDU_UETNLABindingReleaseRequestIEs_Value_MAX = INT_MAX
};

union UETNLABindingReleaseRequestIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                  /* PDU_UETNLABindingReleaseRequestIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                  /* PDU_UETNLABindingReleaseRequestIEs_Value_RAN_UE_NGAP_ID */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct UETNLABindingReleaseRequestIEs_Value {
    UETNLABindingReleaseRequestIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    UETNLABindingReleaseRequestIEs_Value_union decoded;
} UETNLABindingReleaseRequestIEs_Value;

// **************************************************************
//
// UE TNLA BINDING ELEMENTARY PROCEDURES
//
// **************************************************************

// **************************************************************
//
// UE TNLA BINDING RELEASE REQUEST
//
// **************************************************************
typedef struct UETNLABindingReleaseRequest {
    struct _seqof259 {
        struct _seqof259 *next;
        struct _seq307 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            UETNLABindingReleaseRequestIEs_Value value;
        } value;
    } *protocolIEs;
} UETNLABindingReleaseRequest;

enum UERadioCapabilityInfoIndicationIEs_Value_PDUs {
    PDU_UERadioCapabilityInfoIndicationIEs_Value_UNKNOWN = 0,

    PDU_UERadioCapabilityInfoIndicationIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_UERadioCapabilityInfoIndicationIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_UERadioCapabilityInfoIndicationIEs_Value_UERadioCapability = UERadioCapability_PDU,
    PDU_UERadioCapabilityInfoIndicationIEs_Value_UERadioCapabilityForPaging = UERadioCapabilityForPaging_PDU,

    PDU_UERadioCapabilityInfoIndicationIEs_Value_MAX = INT_MAX
};

union UERadioCapabilityInfoIndicationIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
              /* PDU_UERadioCapabilityInfoIndicationIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
              /* PDU_UERadioCapabilityInfoIndicationIEs_Value_RAN_UE_NGAP_ID */
    UERadioCapability *pdu_UERadioCapability;  
           /* PDU_UERadioCapabilityInfoIndicationIEs_Value_UERadioCapability */
    UERadioCapabilityForPaging *pdu_UERadioCapabilityForPaging;  
  /* PDU_UERadioCapabilityInfoIndicationIEs_Value_UERadioCapabilityForPaging */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct UERadioCapabilityInfoIndicationIEs_Value {
    UERadioCapabilityInfoIndicationIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    UERadioCapabilityInfoIndicationIEs_Value_union decoded;
} UERadioCapabilityInfoIndicationIEs_Value;

// **************************************************************
//
// UE RADIO CAPABILITY MANAGEMENT ELEMENTARY PROCEDURES
//
// **************************************************************

// **************************************************************
//
// UE RADIO CAPABILITY INFO INDICATION
//
// **************************************************************
typedef struct UERadioCapabilityInfoIndication {
    struct _seqof260 {
        struct _seqof260 *next;
        struct _seq308 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            UERadioCapabilityInfoIndicationIEs_Value value;
        } value;
    } *protocolIEs;
} UERadioCapabilityInfoIndication;

enum UERadioCapabilityCheckRequestIEs_Value_PDUs {
    PDU_UERadioCapabilityCheckRequestIEs_Value_UNKNOWN = 0,

    PDU_UERadioCapabilityCheckRequestIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_UERadioCapabilityCheckRequestIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_UERadioCapabilityCheckRequestIEs_Value_UERadioCapability = UERadioCapability_PDU,

    PDU_UERadioCapabilityCheckRequestIEs_Value_MAX = INT_MAX
};

union UERadioCapabilityCheckRequestIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                /* PDU_UERadioCapabilityCheckRequestIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                /* PDU_UERadioCapabilityCheckRequestIEs_Value_RAN_UE_NGAP_ID */
    UERadioCapability *pdu_UERadioCapability;  
             /* PDU_UERadioCapabilityCheckRequestIEs_Value_UERadioCapability */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct UERadioCapabilityCheckRequestIEs_Value {
    UERadioCapabilityCheckRequestIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    UERadioCapabilityCheckRequestIEs_Value_union decoded;
} UERadioCapabilityCheckRequestIEs_Value;

// **************************************************************
//
// UE Radio Capability Check Elementary Procedure
//
// **************************************************************

// **************************************************************
//
// UE RADIO CAPABILITY CHECK REQUEST
//
// **************************************************************
typedef struct UERadioCapabilityCheckRequest {
    struct _seqof261 {
        struct _seqof261 *next;
        struct _seq309 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            UERadioCapabilityCheckRequestIEs_Value value;
        } value;
    } *protocolIEs;
} UERadioCapabilityCheckRequest;

enum UERadioCapabilityCheckResponseIEs_Value_PDUs {
    PDU_UERadioCapabilityCheckResponseIEs_Value_UNKNOWN = 0,

    PDU_UERadioCapabilityCheckResponseIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_UERadioCapabilityCheckResponseIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_UERadioCapabilityCheckResponseIEs_Value_IMSVoiceSupportIndicator = IMSVoiceSupportIndicator_PDU,
    PDU_UERadioCapabilityCheckResponseIEs_Value_CriticalityDiagnostics = CriticalityDiagnostics_PDU,

    PDU_UERadioCapabilityCheckResponseIEs_Value_MAX = INT_MAX
};

union UERadioCapabilityCheckResponseIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
               /* PDU_UERadioCapabilityCheckResponseIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
               /* PDU_UERadioCapabilityCheckResponseIEs_Value_RAN_UE_NGAP_ID */
    IMSVoiceSupportIndicator *pdu_IMSVoiceSupportIndicator;  
     /* PDU_UERadioCapabilityCheckResponseIEs_Value_IMSVoiceSupportIndicator */
    CriticalityDiagnostics *pdu_CriticalityDiagnostics;  
       /* PDU_UERadioCapabilityCheckResponseIEs_Value_CriticalityDiagnostics */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct UERadioCapabilityCheckResponseIEs_Value {
    UERadioCapabilityCheckResponseIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    UERadioCapabilityCheckResponseIEs_Value_union decoded;
} UERadioCapabilityCheckResponseIEs_Value;

// **************************************************************
//
// UE RADIO CAPABILITY CHECK RESPONSE
//
// **************************************************************
typedef struct UERadioCapabilityCheckResponse {
    struct _seqof262 {
        struct _seqof262 *next;
        struct _seq310 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            UERadioCapabilityCheckResponseIEs_Value value;
        } value;
    } *protocolIEs;
} UERadioCapabilityCheckResponse;

typedef struct PrivateIE_Container_ {
    struct PrivateIE_Container_ *next;
    PrivateIE_Field value;
} *PrivateIE_Container;

// **************************************************************
//
// PRIVATE MESSAGE ELEMENTARY PROCEDURE
//
// **************************************************************

// **************************************************************
//
// PRIVATE MESSAGE
//
// **************************************************************
typedef struct PrivateMessage {
    struct PrivateIE_Container_ *privateIEs;
} PrivateMessage;

enum SecondaryRATDataUsageReportIEs_Value_PDUs {
    PDU_SecondaryRATDataUsageReportIEs_Value_UNKNOWN = 0,

    PDU_SecondaryRATDataUsageReportIEs_Value_AMF_UE_NGAP_ID = AMF_UE_NGAP_ID_PDU,
    PDU_SecondaryRATDataUsageReportIEs_Value_RAN_UE_NGAP_ID = RAN_UE_NGAP_ID_PDU,
    PDU_SecondaryRATDataUsageReportIEs_Value_PDUSessionResourceSecondaryRATUsageList = PDUSessionResourceSecondaryRATUsageList_PDU,
    PDU_SecondaryRATDataUsageReportIEs_Value_HandoverFlag = HandoverFlag_PDU,

    PDU_SecondaryRATDataUsageReportIEs_Value_MAX = INT_MAX
};

union SecondaryRATDataUsageReportIEs_Value_union {
    AMF_UE_NGAP_ID  *pdu_AMF_UE_NGAP_ID;  
                  /* PDU_SecondaryRATDataUsageReportIEs_Value_AMF_UE_NGAP_ID */
    RAN_UE_NGAP_ID  *pdu_RAN_UE_NGAP_ID;  
                  /* PDU_SecondaryRATDataUsageReportIEs_Value_RAN_UE_NGAP_ID */
    struct PDUSessionResourceSecondaryRATUsageList_ **pdu_PDUSessionResourceSecondaryRATUsageList;                      /* PDU_SecondaryRATDataUsageReportIEs_Value_PDUSessionResourceSecondaryRATUsageList */
    HandoverFlag    *pdu_HandoverFlag;  
                    /* PDU_SecondaryRATDataUsageReportIEs_Value_HandoverFlag */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct SecondaryRATDataUsageReportIEs_Value {
    SecondaryRATDataUsageReportIEs_Value_PDUs pduNum;
    OssBuf          encoded;
    SecondaryRATDataUsageReportIEs_Value_union decoded;
} SecondaryRATDataUsageReportIEs_Value;

// **************************************************************
//
// DATA USAGE REPORTING ELEMENTARY PROCEDURES
//
// **************************************************************

// **************************************************************
//
// SECONDARY RAT DATA USAGE REPORT
//
// **************************************************************
typedef struct SecondaryRATDataUsageReport {
    struct _seqof264 {
        struct _seqof264 *next;
        struct _seq311 {
            ProtocolIE_ID   id;
            Criticality     criticality;
            SecondaryRATDataUsageReportIEs_Value value;
        } value;
    } *protocolIEs;
} SecondaryRATDataUsageReport;

// **************************************************************
//
// IE parameter types from other modules.
//
// **************************************************************
// **************************************************************
//
// Interface Elementary Procedure Class
//
// **************************************************************
typedef struct NGAP_ELEMENTARY_PROCEDURE {
    unsigned char   bit_mask;
#       define      SuccessfulOutcome_present 0x80
#       define      UnsuccessfulOutcome_present 0x40
#       define      criticality_present 0x20
    unsigned short  InitiatingMessage;
    unsigned short  SuccessfulOutcome;  /* optional; set in bit_mask
                                         * SuccessfulOutcome_present if
                                         * present */
    unsigned short  UnsuccessfulOutcome;  /* optional; set in bit_mask
                                           * UnsuccessfulOutcome_present if
                                           * present */
    ProcedureCode   procedureCode;
    Criticality     criticality;  /* criticality_present not set in bit_mask
                                   * implies value is ignore */
} NGAP_ELEMENTARY_PROCEDURE;

enum NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PDUs {
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UNKNOWN = 0,

    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_AMFConfigurationUpdate = AMFConfigurationUpdate_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_HandoverCancel = HandoverCancel_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_HandoverRequired = HandoverRequired_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_HandoverRequest = HandoverRequest_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_InitialContextSetupRequest = InitialContextSetupRequest_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_NGReset = NGReset_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_NGSetupRequest = NGSetupRequest_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PathSwitchRequest = PathSwitchRequest_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PDUSessionResourceModifyRequest = PDUSessionResourceModifyRequest_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PDUSessionResourceModifyIndication = PDUSessionResourceModifyIndication_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PDUSessionResourceReleaseCommand = PDUSessionResourceReleaseCommand_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PDUSessionResourceSetupRequest = PDUSessionResourceSetupRequest_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PWSCancelRequest = PWSCancelRequest_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_RANConfigurationUpdate = RANConfigurationUpdate_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UEContextModificationRequest = UEContextModificationRequest_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UEContextReleaseCommand = UEContextReleaseCommand_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UERadioCapabilityCheckRequest = UERadioCapabilityCheckRequest_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_WriteReplaceWarningRequest = WriteReplaceWarningRequest_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_AMFStatusIndication = AMFStatusIndication_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_CellTrafficTrace = CellTrafficTrace_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_DeactivateTrace = DeactivateTrace_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_DownlinkNASTransport = DownlinkNASTransport_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_DownlinkNonUEAssociatedNRPPaTransport = DownlinkNonUEAssociatedNRPPaTransport_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_DownlinkRANConfigurationTransfer = DownlinkRANConfigurationTransfer_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_DownlinkRANStatusTransfer = DownlinkRANStatusTransfer_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_DownlinkUEAssociatedNRPPaTransport = DownlinkUEAssociatedNRPPaTransport_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_ErrorIndication = ErrorIndication_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_HandoverNotify = HandoverNotify_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_InitialUEMessage = InitialUEMessage_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_LocationReport = LocationReport_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_LocationReportingControl = LocationReportingControl_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_LocationReportingFailureIndication = LocationReportingFailureIndication_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_NASNonDeliveryIndication = NASNonDeliveryIndication_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_OverloadStart = OverloadStart_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_OverloadStop = OverloadStop_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_Paging = Paging_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PDUSessionResourceNotify = PDUSessionResourceNotify_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PrivateMessage = PrivateMessage_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PWSFailureIndication = PWSFailureIndication_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PWSRestartIndication = PWSRestartIndication_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_RerouteNASRequest = RerouteNASRequest_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_RRCInactiveTransitionReport = RRCInactiveTransitionReport_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_SecondaryRATDataUsageReport = SecondaryRATDataUsageReport_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_TraceFailureIndication = TraceFailureIndication_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_TraceStart = TraceStart_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UEContextReleaseRequest = UEContextReleaseRequest_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UERadioCapabilityInfoIndication = UERadioCapabilityInfoIndication_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UETNLABindingReleaseRequest = UETNLABindingReleaseRequest_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UplinkNASTransport = UplinkNASTransport_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UplinkNonUEAssociatedNRPPaTransport = UplinkNonUEAssociatedNRPPaTransport_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UplinkRANConfigurationTransfer = UplinkRANConfigurationTransfer_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UplinkRANStatusTransfer = UplinkRANStatusTransfer_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UplinkUEAssociatedNRPPaTransport = UplinkUEAssociatedNRPPaTransport_PDU,

    PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_MAX = INT_MAX
};

union NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_union {
    AMFConfigurationUpdate *pdu_AMFConfigurationUpdate;  
  /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_AMFConfigurationUpdate */
    HandoverCancel  *pdu_HandoverCancel;  
          /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_HandoverCancel */
    HandoverRequired *pdu_HandoverRequired;  
        /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_HandoverRequired */
    HandoverRequest *pdu_HandoverRequest;  
         /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_HandoverRequest */
    InitialContextSetupRequest *pdu_InitialContextSetupRequest;  
  /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_InitialContextSetupRequest */
    NGReset         *pdu_NGReset;  
                 /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_NGReset */
    NGSetupRequest  *pdu_NGSetupRequest;  
          /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_NGSetupRequest */
    PathSwitchRequest *pdu_PathSwitchRequest;  
       /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PathSwitchRequest */
    PDUSessionResourceModifyRequest *pdu_PDUSessionResourceModifyRequest;  
                                        /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PDUSessionResourceModifyRequest */
    PDUSessionResourceModifyIndication *pdu_PDUSessionResourceModifyIndication;                                         /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PDUSessionResourceModifyIndication */
    PDUSessionResourceReleaseCommand *pdu_PDUSessionResourceReleaseCommand;  
                                        /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PDUSessionResourceReleaseCommand */
    PDUSessionResourceSetupRequest *pdu_PDUSessionResourceSetupRequest;  
      /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PDUSessionResourceSetupRequest */
    PWSCancelRequest *pdu_PWSCancelRequest;  
        /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PWSCancelRequest */
    RANConfigurationUpdate *pdu_RANConfigurationUpdate;  
  /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_RANConfigurationUpdate */
    UEContextModificationRequest *pdu_UEContextModificationRequest;  
    /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UEContextModificationRequest */
    UEContextReleaseCommand *pdu_UEContextReleaseCommand;  
 /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UEContextReleaseCommand */
    UERadioCapabilityCheckRequest *pdu_UERadioCapabilityCheckRequest;  
     /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UERadioCapabilityCheckRequest */
    WriteReplaceWarningRequest *pdu_WriteReplaceWarningRequest;  
  /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_WriteReplaceWarningRequest */
    AMFStatusIndication *pdu_AMFStatusIndication;  
     /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_AMFStatusIndication */
    CellTrafficTrace *pdu_CellTrafficTrace;  
        /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_CellTrafficTrace */
    DeactivateTrace *pdu_DeactivateTrace;  
         /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_DeactivateTrace */
    DownlinkNASTransport *pdu_DownlinkNASTransport;  
    /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_DownlinkNASTransport */
    DownlinkNonUEAssociatedNRPPaTransport *pdu_DownlinkNonUEAssociatedNRPPaTransport;                                   /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_DownlinkNonUEAssociatedNRPPaTransport */
    DownlinkRANConfigurationTransfer *pdu_DownlinkRANConfigurationTransfer;  
                                        /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_DownlinkRANConfigurationTransfer */
    DownlinkRANStatusTransfer *pdu_DownlinkRANStatusTransfer;  
 /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_DownlinkRANStatusTransfer */
    DownlinkUEAssociatedNRPPaTransport *pdu_DownlinkUEAssociatedNRPPaTransport;                                         /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_DownlinkUEAssociatedNRPPaTransport */
    ErrorIndication *pdu_ErrorIndication;  
         /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_ErrorIndication */
    HandoverNotify  *pdu_HandoverNotify;  
          /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_HandoverNotify */
    InitialUEMessage *pdu_InitialUEMessage;  
        /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_InitialUEMessage */
    LocationReport  *pdu_LocationReport;  
          /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_LocationReport */
    LocationReportingControl *pdu_LocationReportingControl;  
/* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_LocationReportingControl */
    LocationReportingFailureIndication *pdu_LocationReportingFailureIndication;                                         /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_LocationReportingFailureIndication */
    NASNonDeliveryIndication *pdu_NASNonDeliveryIndication;  
/* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_NASNonDeliveryIndication */
    OverloadStart   *pdu_OverloadStart;  
           /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_OverloadStart */
    OverloadStop    *pdu_OverloadStop;  
            /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_OverloadStop */
    Paging          *pdu_Paging;  
                  /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_Paging */
    PDUSessionResourceNotify *pdu_PDUSessionResourceNotify;  
/* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PDUSessionResourceNotify */
    PrivateMessage  *pdu_PrivateMessage;  
          /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PrivateMessage */
    PWSFailureIndication *pdu_PWSFailureIndication;  
    /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PWSFailureIndication */
    PWSRestartIndication *pdu_PWSRestartIndication;  
    /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PWSRestartIndication */
    RerouteNASRequest *pdu_RerouteNASRequest;  
       /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_RerouteNASRequest */
    RRCInactiveTransitionReport *pdu_RRCInactiveTransitionReport;  
   /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_RRCInactiveTransitionReport */
    SecondaryRATDataUsageReport *pdu_SecondaryRATDataUsageReport;  
   /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_SecondaryRATDataUsageReport */
    TraceFailureIndication *pdu_TraceFailureIndication;  
  /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_TraceFailureIndication */
    TraceStart      *pdu_TraceStart;  
              /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_TraceStart */
    UEContextReleaseRequest *pdu_UEContextReleaseRequest;  
 /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UEContextReleaseRequest */
    UERadioCapabilityInfoIndication *pdu_UERadioCapabilityInfoIndication;  
                                        /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UERadioCapabilityInfoIndication */
    UETNLABindingReleaseRequest *pdu_UETNLABindingReleaseRequest;  
   /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UETNLABindingReleaseRequest */
    UplinkNASTransport *pdu_UplinkNASTransport;  
      /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UplinkNASTransport */
    UplinkNonUEAssociatedNRPPaTransport *pdu_UplinkNonUEAssociatedNRPPaTransport;                                       /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UplinkNonUEAssociatedNRPPaTransport */
    UplinkRANConfigurationTransfer *pdu_UplinkRANConfigurationTransfer;  
      /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UplinkRANConfigurationTransfer */
    UplinkRANStatusTransfer *pdu_UplinkRANStatusTransfer;  
 /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UplinkRANStatusTransfer */
    UplinkUEAssociatedNRPPaTransport *pdu_UplinkUEAssociatedNRPPaTransport;  
                                        /* PDU_NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_UplinkUEAssociatedNRPPaTransport */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage {
    NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_PDUs pduNum;
    OssBuf          encoded;
    NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage_union decoded;
} NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage;

typedef struct InitiatingMessage {
    ProcedureCode   procedureCode;
    Criticality     criticality;
    NGAP_ELEMENTARY_PROCEDURES_InitiatingMessage value;
} InitiatingMessage;

enum NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PDUs {
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_UNKNOWN = 0,

    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_AMFConfigurationUpdateAcknowledge = AMFConfigurationUpdateAcknowledge_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_HandoverCancelAcknowledge = HandoverCancelAcknowledge_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_HandoverCommand = HandoverCommand_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_HandoverRequestAcknowledge = HandoverRequestAcknowledge_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_InitialContextSetupResponse = InitialContextSetupResponse_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_NGResetAcknowledge = NGResetAcknowledge_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_NGSetupResponse = NGSetupResponse_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PathSwitchRequestAcknowledge = PathSwitchRequestAcknowledge_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PDUSessionResourceModifyResponse = PDUSessionResourceModifyResponse_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PDUSessionResourceModifyConfirm = PDUSessionResourceModifyConfirm_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PDUSessionResourceReleaseResponse = PDUSessionResourceReleaseResponse_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PDUSessionResourceSetupResponse = PDUSessionResourceSetupResponse_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PWSCancelResponse = PWSCancelResponse_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_RANConfigurationUpdateAcknowledge = RANConfigurationUpdateAcknowledge_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_UEContextModificationResponse = UEContextModificationResponse_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_UEContextReleaseComplete = UEContextReleaseComplete_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_UERadioCapabilityCheckResponse = UERadioCapabilityCheckResponse_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_WriteReplaceWarningResponse = WriteReplaceWarningResponse_PDU,

    PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_MAX = INT_MAX
};

union NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_union {
    AMFConfigurationUpdateAcknowledge *pdu_AMFConfigurationUpdateAcknowledge;                                           /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_AMFConfigurationUpdateAcknowledge */
    HandoverCancelAcknowledge *pdu_HandoverCancelAcknowledge;  
 /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_HandoverCancelAcknowledge */
    HandoverCommand *pdu_HandoverCommand;  
         /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_HandoverCommand */
    HandoverRequestAcknowledge *pdu_HandoverRequestAcknowledge;  
  /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_HandoverRequestAcknowledge */
    InitialContextSetupResponse *pdu_InitialContextSetupResponse;  
   /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_InitialContextSetupResponse */
    NGResetAcknowledge *pdu_NGResetAcknowledge;  
      /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_NGResetAcknowledge */
    NGSetupResponse *pdu_NGSetupResponse;  
         /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_NGSetupResponse */
    PathSwitchRequestAcknowledge *pdu_PathSwitchRequestAcknowledge;  
    /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PathSwitchRequestAcknowledge */
    PDUSessionResourceModifyResponse *pdu_PDUSessionResourceModifyResponse;  
                                        /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PDUSessionResourceModifyResponse */
    PDUSessionResourceModifyConfirm *pdu_PDUSessionResourceModifyConfirm;  
                                        /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PDUSessionResourceModifyConfirm */
    PDUSessionResourceReleaseResponse *pdu_PDUSessionResourceReleaseResponse;                                           /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PDUSessionResourceReleaseResponse */
    PDUSessionResourceSetupResponse *pdu_PDUSessionResourceSetupResponse;  
                                        /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PDUSessionResourceSetupResponse */
    PWSCancelResponse *pdu_PWSCancelResponse;  
       /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PWSCancelResponse */
    RANConfigurationUpdateAcknowledge *pdu_RANConfigurationUpdateAcknowledge;                                           /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_RANConfigurationUpdateAcknowledge */
    UEContextModificationResponse *pdu_UEContextModificationResponse;  
     /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_UEContextModificationResponse */
    UEContextReleaseComplete *pdu_UEContextReleaseComplete;  
/* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_UEContextReleaseComplete */
    UERadioCapabilityCheckResponse *pdu_UERadioCapabilityCheckResponse;  
      /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_UERadioCapabilityCheckResponse */
    WriteReplaceWarningResponse *pdu_WriteReplaceWarningResponse;  
   /* PDU_NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_WriteReplaceWarningResponse */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome {
    NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_PDUs pduNum;
    OssBuf          encoded;
    NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome_union decoded;
} NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome;

typedef struct SuccessfulOutcome {
    ProcedureCode   procedureCode;
    Criticality     criticality;
    NGAP_ELEMENTARY_PROCEDURES_SuccessfulOutcome value;
} SuccessfulOutcome;

enum NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_PDUs {
    PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_UNKNOWN = 0,

    PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_AMFConfigurationUpdateFailure = AMFConfigurationUpdateFailure_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_HandoverPreparationFailure = HandoverPreparationFailure_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_HandoverFailure = HandoverFailure_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_InitialContextSetupFailure = InitialContextSetupFailure_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_NGSetupFailure = NGSetupFailure_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_PathSwitchRequestFailure = PathSwitchRequestFailure_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_RANConfigurationUpdateFailure = RANConfigurationUpdateFailure_PDU,
    PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_UEContextModificationFailure = UEContextModificationFailure_PDU,

    PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_MAX = INT_MAX
};

union NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_union {
    AMFConfigurationUpdateFailure *pdu_AMFConfigurationUpdateFailure;  
       /* PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_AMFConfigurationUpdateFailure */
    HandoverPreparationFailure *pdu_HandoverPreparationFailure;  
    /* PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_HandoverPreparationFailure */
    HandoverFailure *pdu_HandoverFailure;  
       /* PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_HandoverFailure */
    InitialContextSetupFailure *pdu_InitialContextSetupFailure;  
    /* PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_InitialContextSetupFailure */
    NGSetupFailure  *pdu_NGSetupFailure;  
        /* PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_NGSetupFailure */
    PathSwitchRequestFailure *pdu_PathSwitchRequestFailure;  
  /* PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_PathSwitchRequestFailure */
    RANConfigurationUpdateFailure *pdu_RANConfigurationUpdateFailure;  
       /* PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_RANConfigurationUpdateFailure */
    UEContextModificationFailure *pdu_UEContextModificationFailure;  
      /* PDU_NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_UEContextModificationFailure */

    All_PDU_Types_union other;  /* for known extensions */
};

typedef struct NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome {
    NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_PDUs pduNum;
    OssBuf          encoded;
    NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome_union decoded;
} NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome;

typedef struct UnsuccessfulOutcome {
    ProcedureCode   procedureCode;
    Criticality     criticality;
    NGAP_ELEMENTARY_PROCEDURES_UnsuccessfulOutcome value;
} UnsuccessfulOutcome;

// **************************************************************
//
// Interface PDU Definition
//
// **************************************************************
typedef struct NGAP_PDU {
    unsigned short  choice;
#       define      initiatingMessage_chosen 1
#       define      successfulOutcome_chosen 2
#       define      unsuccessfulOutcome_chosen 3
    union _union {
        InitiatingMessage initiatingMessage;  /* to choose, set choice to
                                               * initiatingMessage_chosen */
        SuccessfulOutcome successfulOutcome;  /* to choose, set choice to
                                               * successfulOutcome_chosen */
        UnsuccessfulOutcome unsuccessfulOutcome;  /* to choose, set choice to
                                                * unsuccessfulOutcome_chosen */
    } u;
} NGAP_PDU;

#ifndef _OSSNOVALUES

// **************************************************************
//
// IE parameter types from other modules.
//
// **************************************************************
// **************************************************************
//
// Elementary Procedures
//
// **************************************************************
extern const ProcedureCode id_AMFConfigurationUpdate;

extern const ProcedureCode id_AMFStatusIndication;

extern const ProcedureCode id_CellTrafficTrace;

extern const ProcedureCode id_DeactivateTrace;

extern const ProcedureCode id_DownlinkNASTransport;

extern const ProcedureCode id_DownlinkNonUEAssociatedNRPPaTransport;

extern const ProcedureCode id_DownlinkRANConfigurationTransfer;

extern const ProcedureCode id_DownlinkRANStatusTransfer;

extern const ProcedureCode id_DownlinkUEAssociatedNRPPaTransport;

extern const ProcedureCode id_ErrorIndication;

extern const ProcedureCode id_HandoverCancel;

extern const ProcedureCode id_HandoverNotification;

extern const ProcedureCode id_HandoverPreparation;

extern const ProcedureCode id_HandoverResourceAllocation;

extern const ProcedureCode id_InitialContextSetup;

extern const ProcedureCode id_InitialUEMessage;

extern const ProcedureCode id_LocationReportingControl;

extern const ProcedureCode id_LocationReportingFailureIndication;

extern const ProcedureCode id_LocationReport;

extern const ProcedureCode id_NASNonDeliveryIndication;

extern const ProcedureCode id_NGReset;

extern const ProcedureCode id_NGSetup;

extern const ProcedureCode id_OverloadStart;

extern const ProcedureCode id_OverloadStop;

extern const ProcedureCode id_Paging;

extern const ProcedureCode id_PathSwitchRequest;

extern const ProcedureCode id_PDUSessionResourceModify;

extern const ProcedureCode id_PDUSessionResourceModifyIndication;

extern const ProcedureCode id_PDUSessionResourceRelease;

extern const ProcedureCode id_PDUSessionResourceSetup;

extern const ProcedureCode id_PDUSessionResourceNotify;

extern const ProcedureCode id_PrivateMessage;

extern const ProcedureCode id_PWSCancel;

extern const ProcedureCode id_PWSFailureIndication;

extern const ProcedureCode id_PWSRestartIndication;

extern const ProcedureCode id_RANConfigurationUpdate;

extern const ProcedureCode id_RerouteNASRequest;

extern const ProcedureCode id_RRCInactiveTransitionReport;

extern const ProcedureCode id_TraceFailureIndication;

extern const ProcedureCode id_TraceStart;

extern const ProcedureCode id_UEContextModification;

extern const ProcedureCode id_UEContextRelease;

extern const ProcedureCode id_UEContextReleaseRequest;

extern const ProcedureCode id_UERadioCapabilityCheck;

extern const ProcedureCode id_UERadioCapabilityInfoIndication;

extern const ProcedureCode id_UETNLABindingRelease;

extern const ProcedureCode id_UplinkNASTransport;

extern const ProcedureCode id_UplinkNonUEAssociatedNRPPaTransport;

extern const ProcedureCode id_UplinkRANConfigurationTransfer;

extern const ProcedureCode id_UplinkRANStatusTransfer;

extern const ProcedureCode id_UplinkUEAssociatedNRPPaTransport;

extern const ProcedureCode id_WriteReplaceWarning;

extern const ProcedureCode id_SecondaryRATDataUsageReport;

// **************************************************************
//
// Extension constants
//
// **************************************************************
extern const int maxPrivateIEs;

extern const int maxProtocolExtensions;

extern const int maxProtocolIEs;

// **************************************************************
//
// Lists
//
// **************************************************************
extern const int maxnoofAllowedAreas;

extern const int maxnoofAllowedS_NSSAIs;

extern const int maxnoofBPLMNs;

extern const int maxnoofCellIDforWarning;

extern const int maxnoofCellinAoI;

extern const int maxnoofCellinEAI;

extern const int maxnoofCellinTAI;

extern const int maxnoofCellsingNB;

extern const int maxnoofCellsinngeNB;

extern const int maxnoofCellsUEMovingTrajectory;

extern const int maxnoofDRBs;

extern const int maxnoofEmergencyAreaID;

extern const int maxnoofEAIforRestart;

extern const int maxnoofEPLMNs;

extern const int maxnoofEPLMNsPlusOne;

extern const int maxnoofErrors;

extern const int maxnoofForbTACs;

extern const int maxnoofMultiConnectivity;

extern const int maxnoofMultiConnectivityMinusOne;

extern const int maxnoofNGConnectionsToReset;

extern const int maxnoofPDUSessions;

extern const int maxnoofPLMNs;

extern const int maxnoofQosFlows;

extern const int maxnoofRANNodeinAoI;

extern const int maxnoofRecommendedCells;

extern const int maxnoofRecommendedRANNodes;

extern const int maxnoofAoI;

extern const int maxnoofServedGUAMIs;

extern const int maxnoofSliceItems;

extern const int maxnoofTACs;

extern const int maxnoofTAIforInactive;

extern const int maxnoofTAIforPaging;

extern const int maxnoofTAIforRestart;

extern const int maxnoofTAIforWarning;

extern const int maxnoofTAIinAoI;

extern const int maxnoofTimePeriods;

extern const int maxnoofTNLAssociations;

extern const int maxnoofXnExtTLAs;

extern const int maxnoofXnGTP_TLAs;

extern const int maxnoofXnTLAs;

// **************************************************************
//
// IEs
//
// **************************************************************
extern const ProtocolIE_ID id_AllowedNSSAI;

extern const ProtocolIE_ID id_AMFName;

extern const ProtocolIE_ID id_AMFOverloadResponse;

extern const ProtocolIE_ID id_AMFSetID;

extern const ProtocolIE_ID id_AMF_TNLAssociationFailedToSetupList;

extern const ProtocolIE_ID id_AMF_TNLAssociationSetupList;

extern const ProtocolIE_ID id_AMF_TNLAssociationToAddList;

extern const ProtocolIE_ID id_AMF_TNLAssociationToRemoveList;

extern const ProtocolIE_ID id_AMF_TNLAssociationToUpdateList;

extern const ProtocolIE_ID id_AMFTrafficLoadReductionIndication;

extern const ProtocolIE_ID id_AMF_UE_NGAP_ID;

extern const ProtocolIE_ID id_AssistanceDataForPaging;

extern const ProtocolIE_ID id_BroadcastCancelledAreaList;

extern const ProtocolIE_ID id_BroadcastCompletedAreaList;

extern const ProtocolIE_ID id_CancelAllWarningMessages;

extern const ProtocolIE_ID id_Cause;

extern const ProtocolIE_ID id_CellIDListForRestart;

extern const ProtocolIE_ID id_ConcurrentWarningMessageInd;

extern const ProtocolIE_ID id_CoreNetworkAssistanceInformationForInactive;

extern const ProtocolIE_ID id_CriticalityDiagnostics;

extern const ProtocolIE_ID id_DataCodingScheme;

extern const ProtocolIE_ID id_DefaultPagingDRX;

extern const ProtocolIE_ID id_DirectForwardingPathAvailability;

extern const ProtocolIE_ID id_EmergencyAreaIDListForRestart;

extern const ProtocolIE_ID id_EmergencyFallbackIndicator;

extern const ProtocolIE_ID id_FiveG_S_TMSI;

extern const ProtocolIE_ID id_GlobalRANNodeID;

extern const ProtocolIE_ID id_GUAMI;

extern const ProtocolIE_ID id_HandoverType;

extern const ProtocolIE_ID id_IMSVoiceSupportIndicator;

extern const ProtocolIE_ID id_IndexToRFSP;

extern const ProtocolIE_ID id_InfoOnRecommendedCellsAndRANNodesForPaging;

extern const ProtocolIE_ID id_LocationReportingRequestType;

extern const ProtocolIE_ID id_MaskedIMEISV;

extern const ProtocolIE_ID id_MessageIdentifier;

extern const ProtocolIE_ID id_MobilityRestrictionList;

extern const ProtocolIE_ID id_NASC;

extern const ProtocolIE_ID id_NAS_PDU;

extern const ProtocolIE_ID id_NASSecurityParametersFromNGRAN;

extern const ProtocolIE_ID id_NewAMF_UE_NGAP_ID;

extern const ProtocolIE_ID id_NewSecurityContextInd;

extern const ProtocolIE_ID id_NGAP_Message;

extern const ProtocolIE_ID id_NGRAN_CGI;

extern const ProtocolIE_ID id_NGRANTraceID;

extern const ProtocolIE_ID id_NRPPa_PDU;

extern const ProtocolIE_ID id_NumberOfBroadcastsRequested;

extern const ProtocolIE_ID id_OldAMF;

extern const ProtocolIE_ID id_OverloadStartNSSAIList;

extern const ProtocolIE_ID id_PagingDRX;

extern const ProtocolIE_ID id_PagingOrigin;

extern const ProtocolIE_ID id_PagingPriority;

extern const ProtocolIE_ID id_PDUSessionResourceAdmittedList;

extern const ProtocolIE_ID id_PDUSessionResourceFailedToModifyListModRes;

extern const ProtocolIE_ID id_PDUSessionResourceFailedToSetupListCxtRes;

extern const ProtocolIE_ID id_PDUSessionResourceFailedToSetupListHOAck;

extern const ProtocolIE_ID id_PDUSessionResourceFailedToSetupListPSReq;

extern const ProtocolIE_ID id_PDUSessionResourceFailedToSetupListSURes;

extern const ProtocolIE_ID id_PDUSessionResourceHandoverList;

extern const ProtocolIE_ID id_PDUSessionResourceListCxtRelCpl;

extern const ProtocolIE_ID id_PDUSessionResourceListHORqd;

extern const ProtocolIE_ID id_PDUSessionResourceModifyListModCfm;

extern const ProtocolIE_ID id_PDUSessionResourceModifyListModInd;

extern const ProtocolIE_ID id_PDUSessionResourceModifyListModReq;

extern const ProtocolIE_ID id_PDUSessionResourceModifyListModRes;

extern const ProtocolIE_ID id_PDUSessionResourceNotifyList;

extern const ProtocolIE_ID id_PDUSessionResourceReleasedListNot;

extern const ProtocolIE_ID id_PDUSessionResourceReleasedListPSAck;

extern const ProtocolIE_ID id_PDUSessionResourceReleasedListPSFail;

extern const ProtocolIE_ID id_PDUSessionResourceReleasedListRelRes;

extern const ProtocolIE_ID id_PDUSessionResourceSetupListCxtReq;

extern const ProtocolIE_ID id_PDUSessionResourceSetupListCxtRes;

extern const ProtocolIE_ID id_PDUSessionResourceSetupListHOReq;

extern const ProtocolIE_ID id_PDUSessionResourceSetupListSUReq;

extern const ProtocolIE_ID id_PDUSessionResourceSetupListSURes;

extern const ProtocolIE_ID id_PDUSessionResourceToBeSwitchedDLList;

extern const ProtocolIE_ID id_PDUSessionResourceSwitchedList;

extern const ProtocolIE_ID id_PDUSessionResourceToReleaseListHOCmd;

extern const ProtocolIE_ID id_PDUSessionResourceToReleaseListRelCmd;

extern const ProtocolIE_ID id_PLMNSupportList;

extern const ProtocolIE_ID id_PWSFailedCellIDList;

extern const ProtocolIE_ID id_RANNodeName;

extern const ProtocolIE_ID id_RANPagingPriority;

extern const ProtocolIE_ID id_RANStatusTransfer_TransparentContainer;

extern const ProtocolIE_ID id_RAN_UE_NGAP_ID;

extern const ProtocolIE_ID id_RelativeAMFCapacity;

extern const ProtocolIE_ID id_RepetitionPeriod;

extern const ProtocolIE_ID id_ResetType;

extern const ProtocolIE_ID id_RoutingID;

extern const ProtocolIE_ID id_RRCEstablishmentCause;

extern const ProtocolIE_ID id_RRCInactiveTransitionReportRequest;

extern const ProtocolIE_ID id_RRCState;

extern const ProtocolIE_ID id_SecurityContext;

extern const ProtocolIE_ID id_SecurityKey;

extern const ProtocolIE_ID id_SerialNumber;

extern const ProtocolIE_ID id_ServedGUAMIList;

extern const ProtocolIE_ID id_SONConfigurationTransferDL;

extern const ProtocolIE_ID id_SONConfigurationTransferUL;

extern const ProtocolIE_ID id_SourceAMF_UE_NGAP_ID;

extern const ProtocolIE_ID id_SourceToTarget_TransparentContainer;

extern const ProtocolIE_ID id_SupportedTAList;

extern const ProtocolIE_ID id_TAIListForPaging;

extern const ProtocolIE_ID id_TAIListForRestart;

extern const ProtocolIE_ID id_TargetID;

extern const ProtocolIE_ID id_TargetToSource_TransparentContainer;

extern const ProtocolIE_ID id_TimeToWait;

extern const ProtocolIE_ID id_TraceActivation;

extern const ProtocolIE_ID id_TraceCollectionEntityIPAddress;

extern const ProtocolIE_ID id_UEAggregateMaximumBitRate;

extern const ProtocolIE_ID id_UE_associatedLogicalNG_connectionList;

extern const ProtocolIE_ID id_UEContextRequest;

extern const ProtocolIE_ID id_UE_NGAP_IDs;

extern const ProtocolIE_ID id_UEPagingIdentity;

extern const ProtocolIE_ID id_UEPresenceInAreaOfInterestList;

extern const ProtocolIE_ID id_UERadioCapability;

extern const ProtocolIE_ID id_UERadioCapabilityForPaging;

extern const ProtocolIE_ID id_UESecurityCapabilities;

extern const ProtocolIE_ID id_UnavailableGUAMIList;

extern const ProtocolIE_ID id_UserLocationInformation;

extern const ProtocolIE_ID id_WarningAreaList;

extern const ProtocolIE_ID id_WarningMessageContents;

extern const ProtocolIE_ID id_WarningSecurityInfo;

extern const ProtocolIE_ID id_WarningType;

extern const ProtocolIE_ID id_AdditionalUL_NGU_UP_TNLInformation;

extern const ProtocolIE_ID id_DataForwardingNotPossible;

extern const ProtocolIE_ID id_NetworkInstance;

extern const ProtocolIE_ID id_PDUSessionAggregateMaximumBitRate;

extern const ProtocolIE_ID id_PDUSessionResourceFailedToModifyListModCfm;

extern const ProtocolIE_ID id_PDUSessionResourceFailedToSetupListCxtFail;

extern const ProtocolIE_ID id_PDUSessionResourceListCxtRelReq;

extern const ProtocolIE_ID id_PDUSessionType;

extern const ProtocolIE_ID id_QosFlowAddOrModifyRequestList;

extern const ProtocolIE_ID id_QosFlowSetupRequestList;

extern const ProtocolIE_ID id_QosFlowToReleaseList;

extern const ProtocolIE_ID id_SecurityIndication;

extern const ProtocolIE_ID id_UL_NGU_UP_TNLInformation;

extern const ProtocolIE_ID id_UL_NGU_UP_TNLModifyList;

extern const ProtocolIE_ID id_WarningAreaCoordinates;

extern const ProtocolIE_ID id_PDUSessionResourceSecondaryRATUsageList;

extern const ProtocolIE_ID id_HandoverFlag;

extern const ProtocolIE_ID id_SecondaryRATUsageInformation;

extern const ProtocolIE_ID id_PDUSessionResourceReleaseResponseTransfer;

extern const ProtocolIE_ID id_RedirectionVoiceFallback;

extern const ProtocolIE_ID id_UERetentionInformation;

extern const ProtocolIE_ID id_S_NSSAI;

extern const ProtocolIE_ID id_PSCellInformation;

extern const ProtocolIE_ID id_LastEUTRAN_PLMNIdentity;

extern const ProtocolIE_ID id_MaximumIntegrityProtectedDataRate_DL;

extern const ProtocolIE_ID id_AdditionalDLForwardingUPTNLInformation;

extern const ProtocolIE_ID id_AdditionalDLUPTNLInformationForHOList;

extern const ProtocolIE_ID id_AdditionalNGU_UP_TNLInformation;

extern const ProtocolIE_ID id_AdditionalDLQosFlowPerTNLInformation;

extern const ProtocolIE_ID id_SecurityResult;

extern const ProtocolIE_ID id_ENDC_SONConfigurationTransferDL;

extern const ProtocolIE_ID id_ENDC_SONConfigurationTransferUL;

extern const ProtocolIE_ID id_OldAssociatedQosFlowList_ULendmarkerexpected;

extern const ProtocolIE_ID id_CNTypeRestrictionsForEquivalent;

extern const ProtocolIE_ID id_CNTypeRestrictionsForServing;

extern const ProtocolIE_ID id_NewGUAMI;

extern const ProtocolIE_ID id_ULForwardingUP_TNLInformation;

extern const ProtocolIE_ID id_CNAssistedRANTuning;

extern const ProtocolIE_ID id_CommonNetworkInstance;

// **************************************************************
//
// Interface Elementary Procedures
//
// **************************************************************
extern NGAP_ELEMENTARY_PROCEDURE aMFConfigurationUpdate;

extern NGAP_ELEMENTARY_PROCEDURE aMFStatusIndication;

extern NGAP_ELEMENTARY_PROCEDURE cellTrafficTrace;

extern NGAP_ELEMENTARY_PROCEDURE deactivateTrace;

extern NGAP_ELEMENTARY_PROCEDURE downlinkNASTransport;

extern NGAP_ELEMENTARY_PROCEDURE downlinkNonUEAssociatedNRPPaTransport;

extern NGAP_ELEMENTARY_PROCEDURE downlinkRANConfigurationTransfer;

extern NGAP_ELEMENTARY_PROCEDURE downlinkRANStatusTransfer;

extern NGAP_ELEMENTARY_PROCEDURE downlinkUEAssociatedNRPPaTransport;

extern NGAP_ELEMENTARY_PROCEDURE errorIndication;

extern NGAP_ELEMENTARY_PROCEDURE handoverCancel;

extern NGAP_ELEMENTARY_PROCEDURE handoverNotification;

extern NGAP_ELEMENTARY_PROCEDURE handoverPreparation;

extern NGAP_ELEMENTARY_PROCEDURE handoverResourceAllocation;

extern NGAP_ELEMENTARY_PROCEDURE initialContextSetup;

extern NGAP_ELEMENTARY_PROCEDURE initialUEMessage;

extern NGAP_ELEMENTARY_PROCEDURE locationReport;

extern NGAP_ELEMENTARY_PROCEDURE locationReportingControl;

extern NGAP_ELEMENTARY_PROCEDURE locationReportingFailureIndication;

extern NGAP_ELEMENTARY_PROCEDURE nASNonDeliveryIndication;

extern NGAP_ELEMENTARY_PROCEDURE nGReset;

extern NGAP_ELEMENTARY_PROCEDURE nGSetup;

extern NGAP_ELEMENTARY_PROCEDURE overloadStart;

extern NGAP_ELEMENTARY_PROCEDURE overloadStop;

extern NGAP_ELEMENTARY_PROCEDURE paging;

extern NGAP_ELEMENTARY_PROCEDURE pathSwitchRequest;

extern NGAP_ELEMENTARY_PROCEDURE pDUSessionResourceModify;

extern NGAP_ELEMENTARY_PROCEDURE pDUSessionResourceModifyIndication;

extern NGAP_ELEMENTARY_PROCEDURE pDUSessionResourceNotify;

extern NGAP_ELEMENTARY_PROCEDURE pDUSessionResourceRelease;

extern NGAP_ELEMENTARY_PROCEDURE pDUSessionResourceSetup;

extern NGAP_ELEMENTARY_PROCEDURE privateMessage;

extern NGAP_ELEMENTARY_PROCEDURE pWSCancel;

extern NGAP_ELEMENTARY_PROCEDURE pWSFailureIndication;

extern NGAP_ELEMENTARY_PROCEDURE pWSRestartIndication;

extern NGAP_ELEMENTARY_PROCEDURE rANConfigurationUpdate;

extern NGAP_ELEMENTARY_PROCEDURE rerouteNASRequest;

extern NGAP_ELEMENTARY_PROCEDURE rRCInactiveTransitionReport;

extern NGAP_ELEMENTARY_PROCEDURE secondaryRATDataUsageReport;

extern NGAP_ELEMENTARY_PROCEDURE traceFailureIndication;

extern NGAP_ELEMENTARY_PROCEDURE traceStart;

extern NGAP_ELEMENTARY_PROCEDURE uEContextModification;

extern NGAP_ELEMENTARY_PROCEDURE uEContextRelease;

extern NGAP_ELEMENTARY_PROCEDURE uEContextReleaseRequest;

extern NGAP_ELEMENTARY_PROCEDURE uERadioCapabilityCheck;

extern NGAP_ELEMENTARY_PROCEDURE uERadioCapabilityInfoIndication;

extern NGAP_ELEMENTARY_PROCEDURE uETNLABindingRelease;

extern NGAP_ELEMENTARY_PROCEDURE uplinkNASTransport;

extern NGAP_ELEMENTARY_PROCEDURE uplinkNonUEAssociatedNRPPaTransport;

extern NGAP_ELEMENTARY_PROCEDURE uplinkRANConfigurationTransfer;

extern NGAP_ELEMENTARY_PROCEDURE uplinkRANStatusTransfer;

extern NGAP_ELEMENTARY_PROCEDURE uplinkUEAssociatedNRPPaTransport;

extern NGAP_ELEMENTARY_PROCEDURE writeReplaceWarning;

#endif  /* #ifndef _OSSNOVALUES */


extern void *ngap_procedures;    /* encoder-decoder control table */
#ifdef __cplusplus
}	/* extern "C" */
#endif /* __cplusplus */
#endif /* OSS_ngap_procedures */

// ASN1STOP
