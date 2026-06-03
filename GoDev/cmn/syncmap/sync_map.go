package syncmap

import (
	"fmt"
	"sync"
)

type SyncMap struct{ sync.Map }

func (p *SyncMap) Get(key interface{}) interface{} {
	v, ok := p.Load(key)
	if !ok {
		return nil
	}
	return v
}

func (p *SyncMap) Set(key interface{}, val interface{}) error {
	_, err := p.LoadOrStore(key, val)
	if err == false { // false if stored the new value
		return nil
	} else {
		//true if the value was load
		return fmt.Errorf("key exist")
	}
}

func (p *SyncMap) Update(key interface{}, val interface{}) {
	p.Store(key, val)
}

func (p *SyncMap) Del(key interface{}) {
	p.Delete(key)
}

func (p *SyncMap) IsExist(key interface{}) bool {
	_, ok := p.Load(key) // true is exist
	return ok
}

func (p *SyncMap) Length() uint32 {
	length := uint32(0)
	p.Range(func(key, value interface{}) bool {
		length += 1
		return true
	})
	return length
}

// n4 seid
func (p *SyncMap) Length64() uint64 {
	length := uint64(0)
	p.Range(func(key, value interface{}) bool {
		length += 1
		return true
	})
	return length
}
