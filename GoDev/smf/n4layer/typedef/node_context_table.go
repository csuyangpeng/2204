package typedef

import (
	"fmt"
	"lite5gc/cmn/syncmap"
)

var NodePool syncmap.SyncMap //map[string]*Node

func ValuesOfNodeTbl() (CxtList []*Node, err error) {

	NodePool.Range(func(key, value interface{}) bool {
		//fmt.Println(key, value)
		ctxt, ok := value.(*Node)
		if !ok {
			err = fmt.Errorf("invalid node type")
			return false
		}
		CxtList = append(CxtList, ctxt)
		return true
	})

	return
}

func AddNode(key string, ctxt *Node) error {

	var err error

	err = NodePool.Set(key, ctxt)
	if err != nil {
		err = fmt.Errorf("failed to set key(%s),err(%s)", key, err)
	}

	return err
}

func GetNode(key string) (n *Node, err error) {

	val := NodePool.Get(key)
	if val == nil {
		err = fmt.Errorf("failed to find Node with peerIp key(%s)", key)
		//rlogger.Trace(types.SmfN4Layer, rlogger.ERROR, nil, err)
		return
	}
	ctxt, ok := val.(*Node)
	if !ok {
		err = fmt.Errorf("invalid node type")
		//rlogger.Trace(types.SmfN4Layer, rlogger.ERROR, nil, err)
		return
	}
	n = ctxt
	//rlogger.Trace(types.SmfN4Layer, rlogger.ERROR, nil, err)
	return
}

func UpdateNode(key string, n *Node) error {

	if n == nil {
		return fmt.Errorf("invalid input parameter, nil Node")
	}

	NodePool.Update(key, n)

	return nil
}

func DeleteNode(key string) error {

	NodePool.Del(key)

	return nil
}

func LengthOfNodeTbl(key string) uint64 {
	var length uint64
	length = NodePool.Length64()

	return length
}
