package pfcp

import "bytes"

// 3GPP TS 29.244 V15.5.0 (2019-03)
// N4 消息

// 7.5.6	PFCP Session Deletion Request
// IE is null
type IEsSessionDelRequest struct {
}

func (i *IEsSessionDelRequest) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	return encBuf.Bytes(), nil
}

func (i *IEsSessionDelRequest) Decode(data []byte) error {
	return nil
}

func (i *IEsSessionDelRequest) Len() int {
	return 0
}

func (i *IEsSessionDelRequest) SetObject(t uint16, l uint16) error {
	return nil
}

func (i *IEsSessionDelRequest) SetObjectToParent(child interface{}) error {
	return nil
}
