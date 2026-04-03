/** Copyright(C),2020-2022
* Author: Jaytan
* Date: 11/23/20 9:28 PM
* Description: db search results mapping to struct
 */
package model

type ViewForAccessAndMobilitySubscriptionData struct {
	Sst           int    `gorm:"column:sst"`
	Sd            string `gorm:"column:sd"`
	UeAmbrUl      string `gorm:"column:ue_ambr_uplink"`
	UeAmbrDl      string `gorm:"column:ue_ambr_downlink"`
	RfspIndexId   int    `gorm:"column:rfsp_index"`
	SubsRegTimer  int    `gorm:"column:subs_periodic_reg_timer"`
	MpsPriority   string `gorm:"column:mps_priority_ind"`
	McsPriority   string `gorm:"column:mcs_priority_ind"`
	ActiveTime    int    `gorm:"column:active_time"`
	DlPacketCount int    `gorm:"column:dl_packet_count"`
	MicoAllowed   string `gorm:"column:mico_allowed"`
	SupFeatures   string `gorm:"column:supported_feature"`
}

type SnssaiSup struct {
	Id          int    `gorm:"column:id";PRIMARY_KEY`
	SupFeatures string `gorm:"column:supported_feature"`
	Sst         int    `gorm:"column:sst"`
	Sd          string `gorm:"column:sd"`
}

type DNNInfo struct {
	Dnn               string `gorm:"column:dnn"`
	DefaultDnnInd     bool   `gorm:"column:default_dnn_ind"`
	LboRoamingAllowed bool   `gorm:"column:lbo_roaming_allowed"`
	IwkEpsInd         bool   `gorm:"column:iwk_eps_ind"`
}

func (*DNNInfo) TableName() string {
	return "dnn_info"
}

type DNNConfiguration struct {
	Supi                string `gorm:"column:supi"`
	SnssaiId            int    `gorm:"column:snssai_id""`
	Dnn                 string `gorm:"column:dnn"`
	DefPduSessType      string `gorm:"column:def_pdu_sess_type"`
	AllowedPduSessTypes string `gorm:"column:allowed_pdu_sess_type"`
	DefSscMode          string `gorm:"column:def_ssc_mode"`
	AllowedSscMode      string `gorm:"column:allowed_ssc_mode"`
	IwkEpsInd           string `gorm:"column:iwk_eps_ind"`
	LadnInd             string `gorm:"column:ladn_ind"`
	SessAmbrUplink      string `gorm:"column:sess_ambr_uplink"`
	SessAmbrDownlink    string `gorm:"column:sess_ambr_downlink"`
	ChargingChart       string `gorm:"column:3gpp_charging_characteristic"`
	UpSecurityIntegr    string `gorm:"column:up_security_integrity"`
	UpSecurityConfid    string `gorm:"column:up_security_confidentiality"`
	FiveQI              int    `gorm:"column:5qi"`
	PreemptCap          string `gorm:"column:preempt_cap"`
	PreemptVuln         string `gorm:"column:preempt_vuln"`
	StaticIpv4Addr      string `gorm:"column:static_ipv4_address"`
	StaticIpv6Addr      string `gorm:"column:static_ipv6_address"`
	PriorityLevel       int    `gorm:"column:priority_level"`
	ArpPriorityLevel    int    `gorm:"column:arp_priority_level"`
}

type AuthData struct {
	Supi string `gorm:"column:supi"`
	Ki   string `gorm:"column:ki"`
	Opc  string `gorm:"column:opc"`
	Op   string `gorm:"column:op"`
	Amf  string `gorm:"column:amf"`
}

type Amf3gppAccessRegistration struct {
	Id                          int    `gorm:"id";PRIMARY_KEY`
	SupiId                      int    `gorm:"supi_id"`
	AmfInstanceId               string `gorm:"amf_instance_id"`
	DeregCallbackUri            string `gorm:"dereg_callback_uri"`
	GamAmfId                    string `gorm:"gam_amf_id"`
	GamPlmnId                   string `gorm:"gam_plmn_id"`
	RatType                     string `gorm:"rat_type"`
	SupportedFeatures           string `gorm:"supported_features"`
	PurgeFlag                   int    `gorm:"purge_flag"`
	Pei                         string `gorm:"pei"`
	ImsVoPs                     string `gorm:"ims_vo_ps"`
	AmfServiceNameDereg         string `gorm:"amf_service_name_dereg"`
	PcscfRestorationCallbackUri string `gorm:"pcscf_restoration_callback_uri"`
	AmfServiceNamePcscfRest     string `gorm:"amf_service_name_pcscf_rest"`
	InitialRegistrationInd      int    `gorm:"initial_registration_ind"`
	DrFlag                      int    `gorm:"dr_flag"`
}

func (*Amf3gppAccessRegistration) TableName() string {
	return "amf_3gpp_access_registration"
}

type AmfBackupInfo struct {
	Id          int     `gorm:"id";PRIMARY_KEY`
	AmfGppRegId int     `gorm:"amf_gpp_reg_id"`
	BackUpAmf   string  `gorm:"back_up_amf"`
	Guami       []Guami `gorm:"foreignkey:AmfBackUpId"`
}

func (*AmfBackupInfo) TableName() string {
	return "amf_backup_info"
}

type Guami struct {
	Id          int    `gorm:"id";PRIMARY_KEY`
	AmfGppRegId int    `gorm:"amf_gpp_reg_id"`
	AmfBackUpId int    `gorm:"amf_back_up_id"`
	AmfId       string `gorm:"amf_id"`
	PlmnId      string `gorm:"plmn_id"`
}

func (*Guami) TableName() string {
	return "guami"
}

type Supi struct {
	Id   int    `gorm:"id";PRIMARY_KEY` // index, pk
	Supi string `gorm:"supi"`           // supi£º either an IMSI or an NAI Pattern: '^(imsi-[0-9]{5,15}|nai-.+|.+)$'
}

func (*Supi) TableName() string {
	return "supi"
}

type SmfRegistration struct {
	Id                          int       `gorm:"id";PRIMARY_KEY`
	Supi                        string    `gorm:"supi"`
	SmfInstanceId               string    `gorm:"smf_instance_id"`
	SupportedFeatures           string    `gorm:"supported_features"`
	PduSessionId                int32     `gorm:"pdu_session_id"`
	PcscfRestorationCallbackUri string    `gorm:"pcscf_restoration_callback_uri"`
	PlmnId                      string    `gorm:"plmn_id"`
	PgwFqdn                     string    `gorm:"pgw_fqdn"`
	SmfSnssai                   SmfSnssai `gorm:"foreignkey:SmfRegId"`
}

func (*SmfRegistration) TableName() string {
	return "smf_registration"
}

type SmfSnssai struct {
	Id       int    `gorm:"id"` // pk
	SmfRegId int    `gorm:"smf_reg_id"`
	Sst      int32  `gorm:"sst"`
	Sd       string `gorm:"sd"`
}

func (*SmfSnssai) TableName() string {
	return "smf_snssai"
}
