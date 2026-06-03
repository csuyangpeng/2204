package pfcpv1

import (
	"fmt"
	"lite5gc/cmn/message/pfcp"
)

func (m *Message) Unmarshal(data []byte) (err error) {
	//todo
	m.Header.UnmarshalBinary(data)
	//	Node Related Messages
	if m.Header.SFlag == 0 {
		data = data[8:]
	} else {
		data = data[16:]
	}

	switch m.Header.MessageType {
	/*case pfcp.PFCP_Session_Establishment_Request:
	//	todo  调用对应的消息解码函数
	objModule := pfcp.RegMsgObject[pfcp.PFCP_Session_Establishment_Request]
	//	map[PFCP_Session_Establishment_Request] = msg handler{Unmarshal,marshal}
	// 新建对象
	obj := pfcp.CreateObject(objModule)
	pfcpMsg, ok := obj.(pfcp.PfcpMsgInterface)
	if !ok {
		//log
		return fmt.Errorf("type error")
	}
	err := pfcpMsg.UnmarshalBinary(data)
	if err != nil {
		fmt.Printf("msgid %d,err:%s\n", m.Header.MessageType, err)
	}
	m.Body = pfcpMsg*/
	default:
		//	调用对应的消息解码函数
		objModule := pfcp.RegMsgObject[int(m.Header.MessageType)]
		//	map[PFCP_Session_Establishment_Request] = msg handler{Unmarshal,marshal}
		// 新建对象
		obj := pfcp.CreateObject(objModule)
		pfcpMsg, ok := obj.(pfcp.PfcpMsgInterface)
		if !ok {
			//log
			return fmt.Errorf("type error")
		}
		err := pfcpMsg.UnmarshalBinary(data)
		if err != nil {
			fmt.Printf("msgid %d,err:%s\n", m.Header.MessageType, err)
			return err
		}
		m.Body = pfcpMsg
	}
	return nil
}
