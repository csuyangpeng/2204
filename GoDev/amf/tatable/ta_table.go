/** Copyright(C),2020-2022
* Author: zmj
* Date: 11/27/20 3:57 PM
* Description:
 */
package tatable

import (
	"fmt"
	"lite5gc/cmn/redisclt"
	logger "lite5gc/cmn/rlogger"
	"lite5gc/cmn/syncmap"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

type TaiInfo struct {
	Gnbs map[string]types3gpp.GnbInformation // key is gnb ip address
}

// TaTable, key is tai string, value is TaiInfo
// type TaTable map[string]*TaiInfo
var taTable syncmap.SyncMap

func DeleteTai(taiStr string) {
	logger.FuncEntry(types.ModuleAmfTatbl, nil)

	taTable.Del(taiStr)
}

func AddTai(tai types3gpp.TAI, gnbinfo *types3gpp.GnbInformation) error {
	logger.FuncEntry(types.ModuleAmfTatbl, nil)

	// check the input parameter
	if gnbinfo == nil {
		return fmt.Errorf("invalid gnbinfo pointer")
	}

	// tai already checked, supported by amf
	taiStr := tai.String()
	val := taTable.Get(taiStr)
	if val == nil {
		taiInfo := &TaiInfo{
			Gnbs: make(map[string]types3gpp.GnbInformation),
		}
		taiInfo.Gnbs[gnbinfo.GnbIp] = *gnbinfo

		err := taTable.Set(taiStr, taiInfo)
		if err != nil {
			return fmt.Errorf("failed to set key(%v),err(%s)", taiStr, err)
		}
		err = SaveRecordInRedis(taiStr, taiInfo)
		if err != nil {
			logger.Trace(types.ModuleAmfTatbl, logger.ERROR, nil, "failed save,error(%s)", err)
			return err
		}
	} else {
		data, ok := val.(*TaiInfo)
		if !ok {
			return fmt.Errorf("invalid GnbInfo type")
		}

		v, ok := data.Gnbs[gnbinfo.GnbIp]
		if ok {
			logger.Trace(types.ModuleAmfTatbl, logger.DEBUG, nil,
				"gnb ip address already exist in tatable, overwrite the old(%s) with new(%s)",
				v.String(), gnbinfo.String())
		}

		data.Gnbs[gnbinfo.GnbIp] = *gnbinfo
		err := SaveRecordInRedis(taiStr, data)
		if err != nil {
			logger.Trace(types.ModuleAmfTatbl, logger.ERROR, nil, "failed save,error(%s)", err)
			return err
		}
	}

	return nil
}

func SaveRecordInRedis(taiStr string, taiInfo *TaiInfo) error {
	_, err := redisclt.Agent.HSet("tatable", taiStr, *taiInfo)
	if err != nil {
		return fmt.Errorf("failed to save tai info in redis server")
	}
	return nil
}

func GetGnbInsts(taiStr string) ([]uint32, error) {
	logger.FuncEntry(types.ModuleAmfTatbl, nil)

	val := taTable.Get(taiStr)
	if val == nil {
		err := fmt.Errorf("failed to find tab info with tai(%s)", taiStr)
		return nil, err
	}

	data, ok := val.(*TaiInfo)
	if !ok {
		return nil, fmt.Errorf("invalid GnbInfo type")
	}

	var gnbinsts []uint32
	for _, v := range data.Gnbs {
		gnbinsts = append(gnbinsts, v.GnbInstId)
	}

	return gnbinsts, nil
}

// check whether duplicate ip
func CheckGnbIdUniqueness(gnbId uint32) (bool, error) {
	logger.FuncEntry(types.ModuleAmfTatbl, nil)

	var err error
	var exist bool
	taTable.Range(func(key, value interface{}) bool {
		//fmt.Println(key, value)
		taiInfo, ok := value.(*TaiInfo)
		if !ok {
			err = fmt.Errorf("invalid tai info type")
			return false
		}

		for _, v := range taiInfo.Gnbs {
			if v.GnbId == gnbId {
				exist = true
				return false
			}
		}
		return true
	})

	return exist, err
}

func CheckGnbIpAddr(gnbip string) (bool, error) {
	logger.FuncEntry(types.ModuleAmfTatbl, nil)

	logger.Trace(types.ModuleAmfTatbl, logger.DEBUG, nil, "check gnb ip %s", gnbip)

	var err error
	var exist bool
	taTable.Range(func(key, value interface{}) bool {
		//fmt.Println(key, value)
		taiInfo, ok := value.(*TaiInfo)
		if !ok {
			err = fmt.Errorf("invalid tai info type")

			return false
		}

		for k, _ := range taiInfo.Gnbs {
			if k == gnbip {
				exist = true
				return false
			}
		}
		return true
	})

	return exist, err
}

func RemoveGnbInfo(ip string) {
	logger.FuncEntry(types.ModuleAmfTatbl, nil)

	taTable.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		taiInfo, ok := value.(*TaiInfo)
		if !ok {
			logger.Trace(types.ModuleAmfTatbl, logger.DEBUG, nil, "invalid tai info type")
			return true
		}

		for k, _ := range taiInfo.Gnbs {
			if k == ip {
				keyStr, _ := key.(string)

				DeleteTai(keyStr)

				logger.Trace(types.ModuleAmfTatbl, logger.DEBUG, nil, "found the gnb in ta table, key(%s)", keyStr)
				return true
			}
		}
		return true
	})

}
