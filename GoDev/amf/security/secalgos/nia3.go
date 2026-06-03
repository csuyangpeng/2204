package secalgos

import "encoding/binary"

func Nia3ZMAC(key []byte, count []byte, bearer byte, direction byte, Msg []byte, bLength int) (
	zmac []byte, err error) {
	//Eia3V1(ik []byte, count uint32, bearer uint32, direction uint32,
	//	length uint32, m []uint32, mac *uint32)
	// 大端输入，大端输出
	Key := key
	Count := binary.BigEndian.Uint32(count) // BigEndian
	Bearer := uint32(bearer)
	Direction := uint32(direction)

	Length := uint32(bLength)
	PlaintextByte := ZeroPadding(Msg, 4)
	Plaintext, _ := BytesToUint32ArrayV1(PlaintextByte)

	Mac := uint32(0)

	Eia3V1(Key, Count, Bearer, Direction, Length, Plaintext, &Mac)
	tmpMac := make([]byte, 4)
	//binary.BigEndian.PutUint32(tmpMac, Mac)
	// landslide 使用小端流
	binary.LittleEndian.PutUint32(tmpMac, Mac)

	return tmpMac, nil
}
