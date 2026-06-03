package pfcpv1

import (
	"bytes"
	"fmt"
	"lite5gc/cmn/message/pfcp"
)

func (m *Message) Marshal() (data []byte, err error) {
	var body []byte
	switch m.Header.MessageType {
	//case pfcp.PFCP_Session_Establishment_Request:
	//	调用对应的消息解码函数
	//	map[PFCP_Session_Establishment_Request] = msg handler{Unmarshal,marshal}
	default:
		//	调用对应的消息解码函数
		/*objModule := pfcpType.RegMsgObject[pfcp.PFCP_Session_Establishment_Response]
		// 新建对象
		obj := pfcp.CreateObject(objModule)*/
		pfcpMsg, ok := m.Body.(pfcp.PfcpMsgInterface)
		if !ok {
			//log
			return nil, fmt.Errorf("type error")
		}
		body, err = pfcpMsg.MarshalBinary()
		if err != nil {
			return nil, err
		}

	}
	m.Header.Length = uint16(len(body))
	h, err := m.Header.MarshalBinary()
	if err != nil {
		//rlogger.Trace(tError)
	}
	encBuf := bytes.NewBuffer(nil)
	//data = append(data, h...)
	_, err = encBuf.Write(h)
	if err != nil {
		return nil, err
	}
	//data = append(data, body...)

	_, err = encBuf.Write(body)
	if err != nil {
		return nil, err
	}
	return encBuf.Bytes(), nil
}
