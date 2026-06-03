// Package idmgr manage IDs in the application,
// store the ids classfied with types, borrow or return operation is suppored.
package idmgr64

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"sync"
)

// IDMgr is manager for the ids
type IDMgr struct {
	idMgrContainer map[string]*idMap
	idMgrRwMu      sync.RWMutex
}

// for global id manger, however, still can create id manger in module
var manager = NewIDMgr()

// GetInst return a pointer for the global id manager.
func GetInst() *IDMgr {
	return manager
}

// NewIDMgr create a new id manager, return a IDMgr pointer
func NewIDMgr() *IDMgr {
	idMgr := &IDMgr{
		idMgrContainer: make(map[string]*idMap),
	}
	return idMgr
}

// RegisterIDMgr register a type of id with types and max id numbers
func (p *IDMgr) RegisterIDMgr(idType string, maxID uint64) error {
	p.idMgrRwMu.Lock()
	defer p.idMgrRwMu.Unlock()

	if _, ok := p.idMgrContainer[idType]; ok {
		return fmt.Errorf("the id type [%s]  is already registered", idType)
	}

	//create a new type id mgr base
	idMgrBase := &idMap{
		idMap: make(map[uint64]bool),
		maxID: maxID,
	}

	p.idMgrContainer[idType] = idMgrBase

	return nil
}

// BorrowID return a id with id type, will set the flag to busy for the id
func (p *IDMgr) BorrowID(idType string) (id uint64, err error) {
	p.idMgrRwMu.Lock()
	defer p.idMgrRwMu.Unlock()

	idMap, exist := p.idMgrContainer[idType]
	if !exist {
		err = fmt.Errorf("The IDMgr Type %s is NOT registered", idType)
		return
	}

	id, err = idMap.getID()
	return
}

// ReturnID return the id and set the flag to avaible
func (p *IDMgr) ReturnID(idType string, id uint64) error {
	p.idMgrRwMu.Lock()
	defer p.idMgrRwMu.Unlock()

	idMap, exist := p.idMgrContainer[idType]
	if !exist {
		return fmt.Errorf("The IdMgr Type %s is NOT registered", idType)
	}

	return idMap.returnID(id)
}

// GetIDList return a id list in used with id types
func (p *IDMgr) GetIDList(idType types.ModuleName) (idList []uint64, err error) {
	p.idMgrRwMu.RLock()
	defer p.idMgrRwMu.RUnlock()

	idMap, exist := p.idMgrContainer[string(idType)]
	if !exist {
		err = fmt.Errorf("The IdMgr Type %s is NOT registered", idType)
		return
	}

	idList = idMap.getIDList()
	return
}

// DumpIDList show all the id in used
func (p *IDMgr) DumpIDList(idType string) {
	p.idMgrRwMu.RLock()
	defer p.idMgrRwMu.RUnlock()

	idMap, exist := p.idMgrContainer[idType]
	if exist {
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "IDMap [ %s ] dump info: ", idType)
		idMap.display()
	}
	return
}
