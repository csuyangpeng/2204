package n4layer

import (
	"container/list"
	"lite5gc/cmn/rlogger"
	"sync"
)

const moduleTag rlogger.ModuleTag = "n4layer"

// buffer size
const (
	buffer_CHAN_CAP = 1000 // rawsocket receive channel buffer
)

var SequenceNumber uint32 = 1

var UpfN4Layer N4Layer = N4Layer{BufferMsg: make(chan []byte, buffer_CHAN_CAP)}

type N4Layer struct {
	UpfIp     string // 用于本地节点标识
	N3Ip      string
	BufferMsg chan []byte // 用于传递缓存的数据
}

// 发送buffer引用到sendingList, 最大长度是N4Cxt的长度
type SendList struct {
	Rw       sync.RWMutex
	State    chan struct{}
	SendList *list.List
}

var SendingList SendList = SendList{State: make(chan struct{}, 1), SendList: list.New()}
