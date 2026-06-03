package types3gpp

import "fmt"

type CauseType uint8

const (
	CT_RadioNetwork CauseType = 1
	CT_Transport    CauseType = 2
	CT_Nas          CauseType = 3
	CT_Protocal     CauseType = 4
	CT_Misc         CauseType = 5
	CT_Extension    CauseType = 6
)

type CauseValue uint8

// cause for radio network
const (
	Radiok_CauseRadioNetwork_unspecified                                   CauseValue = 0
	Radiok_txnrelocoverall_expiry                                          CauseValue = 1
	Radiok_successful_handover                                             CauseValue = 2
	Radiok_release_due_to_ngran_generated_reason                           CauseValue = 3
	Radiok_release_due_to_5gc_generated_reason                             CauseValue = 4
	Radiok_handover_cancelled                                              CauseValue = 5
	Radiok_partial_handover                                                CauseValue = 6
	Radiok_ho_failure_in_target_5GC_ngran_node_or_target_system            CauseValue = 7
	Radiok_ho_target_not_allowed                                           CauseValue = 8
	Radiok_tngrelocoverall_expiry                                          CauseValue = 9
	Radiok_tngrelocprep_expiry                                             CauseValue = 10
	Radiok_cell_not_available                                              CauseValue = 11
	Radiok_unknown_targetID                                                CauseValue = 12
	Radiok_no_radio_resources_available_in_target_cell                     CauseValue = 13
	Radiok_unknown_local_UE_NGAP_ID                                        CauseValue = 14
	Radiok_inconsistent_remote_UE_NGAP_ID                                  CauseValue = 15
	Radiok_handover_desirable_for_radio_reason                             CauseValue = 16
	Radiok_time_critical_handover                                          CauseValue = 17
	Radiok_resource_optimisation_handover                                  CauseValue = 18
	Radiok_reduce_load_in_serving_cell                                     CauseValue = 19
	Radiok_user_inactivity                                                 CauseValue = 20
	Radiok_radio_connection_with_ue_lost                                   CauseValue = 21
	Radiok_load_balancing_tau_required                                     CauseValue = 22
	Radiok_radio_resources_not_available                                   CauseValue = 23
	Radiok_invalid_qos_combination                                         CauseValue = 24
	Radiok_failure_in_radio_interface_procedure                            CauseValue = 25
	Radiok_interaction_with_other_procedure                                CauseValue = 26
	Radiok_unknown_PDU_session_ID                                          CauseValue = 27
	Radiok_unkown_qos_flow_ID                                              CauseValue = 28
	Radiok_multiple_PDU_session_ID_instances                               CauseValue = 29
	Radiok_multiple_qos_flow_ID_instances                                  CauseValue = 30
	Radiok_encryption_and_or_integrity_protection_algorithms_not_supported CauseValue = 31
	Radiok_ng_intra_system_handover_triggered                              CauseValue = 32
	Radiok_ng_inter_system_handover_triggered                              CauseValue = 33
	Radiok_xn_handover_triggered                                           CauseValue = 34
	Radiok_not_supported_5QI_value                                         CauseValue = 35
	Radiok_ue_context_transfer                                             CauseValue = 36
	Radiok_ims_voice_eps_fallback_or_rat_fallback_triggered                CauseValue = 37
	Radiok_up_integrity_protection_not_possible                            CauseValue = 38
	Radiok_up_confidentiality_protection_not_possible                      CauseValue = 39
	Radiok_slice_not_supported                                             CauseValue = 40
	Radiok_ue_in_rrc_inactive_state_not_reachable                          CauseValue = 41
	Radiok_redirection                                                     CauseValue = 42
	Radiok_resources_not_available_for_the_slice                           CauseValue = 43
	Radiok_ue_max_integrity_protected_data_rate_reason                     CauseValue = 44
	Radiok_release_due_to_cn_detected_mobility                             CauseValue = 45
)

//for transport cause
const (
	Transport_resource_unavailable CauseValue = 0
	Transport_unspecified          CauseValue = 1
)

//for nas cause
const (
	Nas_normal_release         CauseValue = 0 //an release
	Nas_authentication_failure CauseValue = 1
	Nas_deregister             CauseValue = 2 // de register
	Nas_unspecified            CauseValue = 3
)

//for protocol
const (
	Protocol_transfer_syntax_error                             CauseValue = 0
	Protocol_abstract_syntax_error_reject                      CauseValue = 1
	Protocol_abstract_syntax_error_ignore_and_notify           CauseValue = 2
	Protocol_message_not_compatible_with_receiver_state        CauseValue = 3
	Protocol_semantic_error                                    CauseValue = 4
	Protocol_abstract_syntax_error_falsely_constructed_message CauseValue = 5
	Protocol_unspecified                                       CauseValue = 6
)

//for misc
const (
	Misc_control_processing_overload                CauseValue = 0
	Misc_not_enough_user_plane_processing_resources CauseValue = 1
	Misc_hardware_failure                           CauseValue = 2
	Misc_om_intervention                            CauseValue = 3
	Misc_unknown_PLMN                               CauseValue = 4
	Misc_unspecified                                CauseValue = 5
)

type NgapCause struct {
	Type  CauseType
	Value CauseValue
}

func (p NgapCause) String() string {
	var str string
	switch p.Type {
	case CT_RadioNetwork:
		str = "RadioNetwork"
	case CT_Transport:
		str = "Transport"
	case CT_Nas:
		str = "Nas"
	case CT_Protocal:
		str = "Protocal"
	case CT_Misc:
		str = "Misc"
	case CT_Extension:
		str = "Extension"
	}
	str += fmt.Sprintf("_%d", p.Value)

	return str
}
