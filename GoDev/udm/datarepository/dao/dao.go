/** Copyright(C),2020-2022
* Author: Jaytan
* Date: 11/23/20 9:31 PM
* Description:
 */
package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lite5gc/udm/datarepository/model"
	"lite5gc/udm/dbmgr"
)

func GetAmsdDataBySupi(supi string) (amsd model.ViewForAccessAndMobilitySubscriptionData, err error) {
	err = dbmgr.DBGorm.Table("supi").Where("supi = ?", supi).Joins("" +
		"left join access_mobility_subscription amsd on amsd.supi_id = supi.id").Joins("" +
		"left join nssai on nssai.id = amsd.nssai_id").Joins("" +
		"left join snssai on snssai.id = nssai.default_snssai_id").Select("" +
		"snssai.sst,snssai.sd,amsd.ue_ambr_uplink,amsd.ue_ambr_downlink," +
		"amsd.rfsp_index,amsd.subs_periodic_reg_timer,amsd.mps_priority_ind,amsd.mcs_priority_ind," +
		"amsd.active_time,amsd.dl_packet_count,amsd.mico_allowed,amsd.supported_feature").Find(&amsd).Error
	if gorm.IsRecordNotFoundError(err) {
		return amsd, fmt.Errorf("failed to get amsd data with supi, key(%s), error(%s)", supi, err)
	}
	return amsd, nil
}

// Snssai Data
func GetSnssaiBySupi(supi string) (ss model.SnssaiSup, err error) {
	err = dbmgr.DBGorm.Table("supi").Where("supi = ?", supi).Joins("" +
		"left join access_mobility_subscription amsd on amsd.supi_id = supi.id").Joins("" +
		"left join nssai on nssai.id = amsd.nssai_id").Joins("" +
		"left join snssai on snssai.id = nssai.default_snssai_id").Select("" +
		"snssai.id,snssai.sst,snssai.sd," +
		"amsd.supported_feature").Find(&ss).Error
	if gorm.IsRecordNotFoundError(err) {
		return ss, fmt.Errorf("failed to get snssai data with supi, key(%s), error(%s)", supi, err)
	}
	return ss, nil
}

// DNN Information Data
func GetDnnInfoBySnssaiID(id int) (dn []model.DNNInfo, err error) {
	err = dbmgr.DBGorm.Table("snssai").Where("snssai.id = ?", id).Joins("" +
		"left join snssai_info on snssai_info.snssai_id = snssai.id").Joins("" +
		"left join dnn_info on dnn_info.id = snssai_info.dnn_info_id").Select("" +
		"dnn_info.dnn,dnn_info.default_dnn_ind,dnn_info.lbo_roaming_allowed,dnn_info.iwk_eps_ind").Find(&dn).Error
	if gorm.IsRecordNotFoundError(err) {
		return dn, fmt.Errorf("failed to get dnn data with snssai id, key(%s), error(%s)", id, err)
	}
	return dn, nil
}

// DNN Configuration Data
func GetDnnConfigByDnnName(supi string) (dnc model.DNNConfiguration, err error) {
	err = dbmgr.DBGorm.Table("supi").Where("supi = ?", supi).Joins("" +
		"left join session_manage_info on supi.id = session_manage_info.supi_id").Joins("" +
		"left join dnn_info on session_manage_info.dnn_id = dnn_info.id").Joins("" +
		"left join dnn_configuration dnnc on dnnc.id = session_manage_info.ue_dnn_config_id").Joins("" +
		"left join qos_profile_5g on qos_profile_5g.id = dnnc.5g_qos_profile_id").Joins("" +
		"left join arp on arp.id = qos_profile_5g.arp_id").Select("" +
		"dnn_info.dnn," +
		"supi.supi," +
		"session_manage_info.static_ipv4_address,session_manage_info.static_ipv6_address,session_manage_info.snssai_id," +
		"qos_profile_5g.5qi,qos_profile_5g.priority_level," +
		"arp.priority_level as arp_priority_level,arp.preempt_cap,arp.preempt_vuln," +
		"dnnc.def_pdu_sess_type,dnnc.allowed_pdu_sess_type,dnnc.def_ssc_mode,dnnc.allowed_ssc_mode," +
		"dnnc.iwk_eps_ind,dnnc.ladn_ind,dnnc.sess_ambr_uplink,dnnc.sess_ambr_downlink," +
		"dnnc.3gpp_charging_characteristic,dnnc.up_security_integrity,dnnc.up_security_confidentiality").Find(&dnc).Error
	if gorm.IsRecordNotFoundError(err) {
		return dnc, fmt.Errorf("failed to get dnn config data with dnn name, key(%s), error(%s)", supi, err)
	}
	//fmt.Println("static ipv4:",dnc.StaticIpv4Addr)
	return dnc, nil
}

// Authentication Data
func GetAuthDataBySupi(supi string) (auth model.AuthData, err error) {
	err = dbmgr.DBGorm.Table("supi").Where("supi = ?", supi).Joins("" +
		"left join auth_data on auth_data.supi_id = supi.id").Joins("" +
		"left join amf on amf.id = auth_data.amf_id").Joins("" +
		"left join k4 on k4.id = auth_data.ki_k4_id").Joins("" +
		"left join op on op.id = auth_data.op_id").Select("" +
		"amf.amf," +
		"auth_data.ki,auth_data.opc," +
		"supi.supi," +
		"op.op").Find(&auth).Error
	if gorm.IsRecordNotFoundError(err) {
		return auth, fmt.Errorf("failed to get amsd data with supi, key(%s), error(%s)", supi, err)
	}
	return auth, nil
}

func InsertAmf3gppRegInfo(tx *gorm.DB, amf3GppAccessRegistration *model.Amf3gppAccessRegistration) (int, error) {

	create := tx.Create(amf3GppAccessRegistration)

	fmt.Println(amf3GppAccessRegistration.Id)

	return amf3GppAccessRegistration.Id, create.Error
}

func InsertSupi(tx *gorm.DB, sp *model.Supi) (int, error) {

	cr := tx.Create(sp)

	return sp.Id, cr.Error
}

func GetSupiIdBySupi(tx *gorm.DB, supi string) (int, error) {

	sp := model.Supi{}

	err := tx.Table("supi").Where("supi= ?", supi).First(&sp).Error

	if gorm.IsRecordNotFoundError(err) {
		return 0, fmt.Errorf("failed to get amsd data with supi, key(%s), error(%s)", supi, err)
	}
	return sp.Id, nil
}

func InsertAmfBackInfo(tx *gorm.DB, bk *model.AmfBackupInfo) (int, error) {

	cr := tx.Create(bk)

	return bk.Id, cr.Error
}

func InsertSmfReg(tx *gorm.DB, sr *model.SmfRegistration) (int, error) {

	cr := tx.Create(sr)

	return sr.Id, cr.Error
}

func InsertSmfSnssai(tx *gorm.DB, ss *model.SmfSnssai) error {

	cr := tx.Create(ss)
	return cr.Error
}

func Get3gppAmfRegData(tx *gorm.DB, gpsi string) (amf3gpp model.Amf3gppAccessRegistration, err error) {

	err = tx.Table("gpsi").Where("gpsi = ?", gpsi).
		Joins("LEFT JOIN map_supi_gpsi mp ON gpsi.id = mp.gpsi_id").
		Joins("left join amf_3gpp_access_registration amf3gpp on amf3gpp.supi_id = mp.supi_id and amf3gpp.purge_flag != 1").
		Select("amf3gpp.*").
		Find(&amf3gpp).Error

	if gorm.IsRecordNotFoundError(err) {
		return amf3gpp, fmt.Errorf("failed to get amf3gppreg data with supi, key(%s), error(%s)", gpsi, err)
	}
	return amf3gpp, nil
}

func Delete3gppreg(supi string) error {

	err := dbmgr.DBGorm.Exec("update amf_3gpp_access_registration ar set ar.purge_flag = 1 where ar.supi_id = (select supi_id from supi where supi = " + supi + ")").Error

	return err
}

func DeleteSmfReg(supi string, pdussid int32) (err error) {

	tx := dbmgr.DBGorm.Begin()

	err = tx.Exec("delete smf_registration where supi = " + supi + " and pdu_session_id = " + string(pdussid) + ")").Error

	err = tx.Exec("delete smf_snssai where smf_reg_id = (select id from smf_registration where supi = " + supi + ")").Error

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete smf_registration data with supi, key(%s),key(%d), error(%s)", supi, pdussid, err)
	}

	tx.Commit()

	return err
}

func GetSmfReg(gpsi string) (registration *model.SmfRegistration, err error) {

	sql := "SELECT srg.supi,srg.smf_instance_id,srg.supported_features,srg.pdu_session_id,srg.pcscf_restoration_callback_uri,srg.plmn_id,srg.pgw_fqdn,ssi.sst,ssi.sd from smf_registration srg LEFT JOIN smf_snssai ssi on srg.id = ssi.smf_reg_id where srg.supi = (SELECT supi FROM view_supi_gpsi WHERE gpsi = ? )"

	err = dbmgr.DBGorm.Exec(sql, gpsi).Find(registration).Error

	if gorm.IsRecordNotFoundError(err) {
		return registration, fmt.Errorf("failed to get Smf data with gpsi, key(%s), error(%s)", gpsi, err)
	}

	return registration, nil
}
