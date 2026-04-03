package idmgr

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

const MODULE_ID = types.ModuleCmnIdMgr

type idMap struct {
	idMap          map[uint32]bool
	maxID          uint32
	reserveList    []uint32
	currentId      uint32
	allocatedIdNum uint32
}

func (p *idMap) idMapInit(maxNum uint32) {
	p.idMap = make(map[uint32]bool)
	p.maxID = maxNum
	p.reserveList = make([]uint32, 8)
	p.currentId = 0
	p.allocatedIdNum = 0
}

func (p *idMap) reserve(id uint32) error {
	if p.idMap[id] == true {
		return fmt.Errorf("id(%d) is not idle, cannot be reserved.", id)
	}

	p.idMap[id] = true
	p.allocatedIdNum++

	return nil
}

func (p *idMap) unReserve(id uint32) error {
	if p.idMap[id] == false {
		return fmt.Errorf("id(%d) is idle, cannot be unreserved.", id)
	}

	p.idMap[id] = false
	p.allocatedIdNum--

	return nil
}

func (p *idMap) getID() (uint32, error) {

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

func (p *idMap) returnID(id uint32) error {
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
	rlogger.Trace(MODULE_ID, rlogger.DEBUG, nil,
		"IdMap Info: numOfAllocId(%d),CurrentId(%d),MaxNum(%d)",
		p.allocatedIdNum,
		p.currentId,
		p.maxID)

	for key, value := range p.idMap {
		rlogger.Trace(MODULE_ID, rlogger.DEBUG, nil, "  %d - %v", key, value)
	}
}

func (p *idMap) getIDList() (idList []uint32) {
	for key, value := range p.idMap {
		if value == true {
			idList = append(idList, key)
		}
	}
	return
}
