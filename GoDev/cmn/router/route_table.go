package router

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/syncmap"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
	"math/rand"
	"strings"
)

// route table
type CtrlChannel chan *ControlMsg
type DataChannel chan *IpcMsg

type InstList []uint32

var routeTableSMap syncmap.SyncMap // Type_Inst - DataChannel
var routeInstSMap syncmap.SyncMap  // Type - []Inst

//type InstChannMap map[uint32]DataChannel
//type RouteTable map[InstType]InstChannMap

//func (p RouteTable) register(addr RouteAddr, chann DataChannel) error {
func routerRegister(addr RouteAddr, chann DataChannel) error {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)

	//key := addr.String()
	//if _, ok := p[key]; ok {
	//	rlogger.Trace(types.ModuleCmnRouter, rlogger.WARN, nil,  "route address (%s) already exist in route table", key)
	//	return fmt.Errorf("duplicated, failed to register for (%s)", key)
	//}
	//
	//p[key] = chann

	//if instChannMap, ok := p[addr.Type]; !ok {
	//	rlogger.Trace(types.ModuleCmnRouter, rlogger.DEBUG, nil,  "route address type (%s) is not exist "+
	//		"in route table, create the type", addr.Type.String())
	//	icMap := make(InstChannMap)
	//	icMap[addr.Id] = chann
	//	p[addr.Type] = icMap
	//} else {
	//	if _, ok := instChannMap[addr.Id]; ok {
	//		rlogger.Trace(types.ModuleCmnRouter, rlogger.WARN, nil,  "route address (%s) already exist in route table", addr)
	//		return fmt.Errorf("duplicated, failed to register for (%s)", addr)
	//	}
	//
	//	instChannMap[addr.Id] = chann
	//}

	value := routeInstSMap.Get(addr.Type)
	if value != nil {
		instList := value.(InstList)
		instList = append(instList, addr.Id)
		routeInstSMap.Update(addr.Type, instList)
	} else {
		instList := InstList{addr.Id}
		routeInstSMap.Set(addr.Type, instList)
	}

	routeTableSMap.Set(addr.String(), chann)

	return nil
}

//func (p RouteTable) deregister(addr RouteAddr) error {
func routerDeregister(addr RouteAddr) error {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)

	//delete(p, addr.String())
	//if instChannMap, ok := p[addr.Type]; !ok {
	//	rlogger.Trace(types.ModuleCmnRouter, rlogger.DEBUG, nil,  "route address type (%s) is not exist "+
	//		"in route table", addr.Type.String())
	//} else {
	//	if _, ok := instChannMap[addr.Id]; ok {
	//		delete(instChannMap, addr.Id)
	//	}
	//}

	value := routeInstSMap.Get(addr.Type)
	if value != nil {
		instList := value.(InstList)
		//instList = append(instList, addr.Id)
		for index, value := range instList {
			if value == addr.Id {
				instList = append(instList[:index], instList[index+1:]...)
				break
			}
		}
		routeInstSMap.Update(addr.Type, instList)
	}

	routeTableSMap.Del(addr.String())
	return nil
}

//func (p RouteTable) getDestChannel(dest *RouteAddr) ([]DataChannel, error) {
func getDestChannel(dest *RouteAddr) ([]DataChannel, error) {
	//rlogger.FuncEntry(types.ModuleCmnRouter, nil)

	//key := dest.String()
	//
	//v, ok := p[key]
	//if !ok {
	//	rlogger.Trace(types.ModuleCmnRouter, rlogger.WARN, nil,  "route address (%s) is NOT exist in route table", key)
	//	return nil, fmt.Errorf("failed, route address (%s) is NOT exist in route table", key)
	//}
	//
	//return v, nil

	//rtIpcChann := make([]DataChannel, 0)

	//if instChannMap, ok := p[dest.Type]; !ok {
	//	rlogger.Trace(types.ModuleCmnRouter, rlogger.ERROR, nil,  "route address type (%s) is not exist "+
	//		"in route table", dest.Type)
	//	return nil, fmt.Errorf("failed, the dest addr type is not registed")
	//} else {
	//	switch dest.Id {
	//	case UnknownId: //rand select a destination id
	//		//get all the inst id
	//		instIds := make([]uint32, 0)
	//		for inst, _ := range instChannMap {
	//			instIds = append(instIds, inst)
	//		}
	//		// select the dest inst id random
	//		selectId := rand.Intn(len(instIds))
	//		selectKey := instIds[selectId]
	//
	//		rtIpcChann = append(rtIpcChann, instChannMap[selectKey])
	//		return rtIpcChann, nil
	//
	//	case BroadcastId:
	//		//get all the inst id
	//		for _, val := range instChannMap {
	//			rtIpcChann = append(rtIpcChann, val)
	//		}
	//		return rtIpcChann, nil
	//
	//	case InvalidId:
	//		return nil, fmt.Errorf("failed, the dest addr id is invalid")
	//
	//	default:
	//		if vch, ok := instChannMap[dest.Id]; ok {
	//			rtIpcChann = append(rtIpcChann, vch)
	//			return rtIpcChann, nil
	//		} else {
	//			rlogger.Trace(types.ModuleCmnRouter, rlogger.WARN, nil,  "route address (%s) is NOT exist in route table", dest)
	//			return nil, fmt.Errorf("failed, route address (%s) is NOT exist in route table", *dest)
	//		}
	//	}
	//}

	rtIpcChann := make([]DataChannel, 0)

	value := routeTableSMap.Get(dest.String())
	if value != nil {
		rtIpcChann = append(rtIpcChann, value.(DataChannel))
		return rtIpcChann, nil
	} else {
		switch dest.Id {
		case UnknownId: //rand select a destination id
			value := routeInstSMap.Get(dest.Type)
			if value == nil {
				return nil, fmt.Errorf("failed, the dest addr type %s is not registed", &dest)
			}
			instIds := value.(InstList)

			// select the dest inst id random
			index := rand.Intn(len(instIds))
			id := instIds[index]
			selectKey := fmt.Sprintf("%s.%d", dest.Type, id)
			value = routeTableSMap.Get(selectKey)
			if value != nil {
				rtIpcChann = append(rtIpcChann, value.(DataChannel))
				return rtIpcChann, nil
			} else {
				return nil, fmt.Errorf("failed, the dest addr %s is not register", selectKey)
			}
		case BroadcastId:
			value := routeInstSMap.Get(dest.Type)
			if value == nil {
				return nil, fmt.Errorf("failed, the dest addr type %s is not registed", &dest)
			}
			instIds := value.(InstList)

			for _, id := range instIds {
				key := fmt.Sprintf("%s.%d", dest.Type, id)
				value = routeTableSMap.Get(key)
				if value != nil {
					rtIpcChann = append(rtIpcChann, value.(DataChannel))
				} else {
					rlogger.Trace(types.ModuleCmnRouter, rlogger.DEBUG, nil, "failed, the dest addr %s is not register", key)
				}
			}
			return rtIpcChann, nil

		case InvalidId:
			return nil, fmt.Errorf("failed, the dest addr id is invalid")

		default:
			rlogger.Trace(types.ModuleCmnRouter, rlogger.WARN, nil, "route address (%s) is NOT exist in route table", *dest)
			return nil, fmt.Errorf("failed, route address (%v) is NOT exist in route table", *dest)
		}
	}
}

func ShowRouteTable() string {
	var dumpString string

	dumpString = fmt.Sprintf("worker %d\nRoute Table Information {", utils.Goid())

	f := func(k, v interface{}) bool {
		dumpString = dumpString + fmt.Sprintf("[type - %s, inst - ", k.(InstType))
		idList := v.(InstList)
		for id, _ := range idList {
			dumpString = dumpString + fmt.Sprintf("%d,", id)
		}
		dumpString = strings.TrimSuffix(dumpString, ",") + "], "
		return true
	}
	routeInstSMap.Range(f)

	dumpString = strings.TrimSuffix(dumpString, ", ") + "}\n"
	return dumpString
}

func ShowTablesMap() string {
	var dumpString string

	dumpString = fmt.Sprintf("Router Map:\n")

	f := func(k, v interface{}) bool {
		dumpString = dumpString + fmt.Sprintf("[DestAddr - %s, channel - %v ]\n", k.(string), v)

		return true
	}
	routeTableSMap.Range(f)
	return dumpString
}

//func (p RouteTable) showRouteTable() {
//	rlogger.FuncEntry(types.ModuleCmnRouter, nil)
//
//	var dumpString string
//
//	dumpString = fmt.Sprintf("worker %d, Route Table Information {", utils.Goid())
//
//	for key, val := range p {
//		dumpString = dumpString + fmt.Sprintf("[type - %s, inst - ", key)
//		for instId, _ := range val {
//			dumpString = dumpString + fmt.Sprintf("%d,", instId)
//		}
//		dumpString = strings.TrimSuffix(dumpString, ",") + "], "
//	}
//	dumpString = strings.TrimSuffix(dumpString, ", ") + "}"
//	rlogger.Trace(types.ModuleCmnRouter, rlogger.DEBUG, nil,  "%s", dumpString)
//}
