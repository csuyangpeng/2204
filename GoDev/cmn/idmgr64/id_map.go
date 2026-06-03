package idmgr64

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

type idMap struct {
	idMap          map[uint64]bool
	maxID          uint64
	reserveList    []uint64
	currentId      uint64
	allocatedIdNum uint64
}

func (p *idMap) idMapInit(maxNum uint64) {
	p.idMap = make(map[uint64]bool)
	p.maxID = maxNum
	p.currentId = 0
	p.allocatedIdNum = 0
}

func (p *idMap) reserve(id uint64) error {
	if p.idMap[id] == true {
		return fmt.Errorf("id(%d) is not idle, cannot be reserved.", id)
	}

	p.idMap[id] = true
	p.allocatedIdNum++

	return nil
}

func (p *idMap) unReserve(id uint64) error {
	if p.idMap[id] == false {
		return fmt.Errorf("id(%d) is idle, cannot be unreserved.", id)
	}

	p.idMap[id] = false
	p.allocatedIdNum--

	return nil
}

func (p *idMap) getID() (id uint64, err error) {
	////find the availbe id from current map
	//for key, value := range p.idMap {
	//	// if key == 0 { //skip 0
	//	// 	continue
	//	// }
	//	if value == false {
	//		id = key
	//		p.idMap[key] = true
	//		return
	//	}
	//}
	//
	////no avaiable id
	//id = uint64(len(p.idMap)) // + 1) //skip 0
	//if id == p.maxID {
	//	err = fmt.Errorf("Reach max id(%d) allowed", id)
	//	return
	//}
	//
	////mark the id with busy status
	//p.idMap[id] = true
	//
	//return

	curId := p.currentId

	loopFound := true
	for loopFound {
		if p.idMap[curId] == false {
			loopFound = false
		} else {
			curId++
			if curId >= p.maxID {
				if p.allocatedIdNum > p.maxID {
					// no available id resource
					return 0, fmt.Errorf("no available id resource.")
				}

				curId = 0 // restart loop from begin and find the first available id
			}
		}
	}

	//should be found here
	p.currentId = curId
	p.allocatedIdNum++

	p.idMap[curId] = true
	return curId, nil

}

func (p *idMap) returnID(id uint64) error {
	_, ok := p.idMap[id]
	if ok != true {
		return fmt.Errorf("invalid id %d", id)
	}

	//mark the id with available status
	p.idMap[id] = false
	p.allocatedIdNum--
	return nil
}

func (p *idMap) display() {
	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil,
		"IdMap Info: numOfAllocId(%d),CurrentId(%d),MaxNum(%d)",
		p.allocatedIdNum,
		p.currentId,
		p.maxID)

	for key, value := range p.idMap {
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "  %d - %v", key, value)
	}
}

func (p *idMap) getIDList() (idList []uint64) {
	for key, value := range p.idMap {
		if value == true {
			idList = append(idList, key)
		}
	}
	return
}
