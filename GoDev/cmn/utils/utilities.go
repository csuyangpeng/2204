package utils

import "C"
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"unsafe"
)

//IsDigitString check the string is digits string or not
func IsDigitString(str string) bool {
	for _, v := range str {
		if v < '0' || v > '9' {
			return false
		}
	}

	return true
}

func IsValidNaiString(str string) bool {
	return true
}

func GetBitValue(b byte, index uint8) (bool, error) {
	if index > 8 || index < 1 {
		return false, fmt.Errorf("invalid input parameter, index(%d)", index)
	}

	var val byte = 1
	val = val << (index - 1)
	if b&val != 0 {
		return true, nil
	} else {
		return false, nil
	}
}

// fromIdx: 1 to 8
func GetBitsValue(b byte, fromIdx uint8, toIdx uint8) (byte, error) {
	if fromIdx > 8 || fromIdx < 1 ||
		toIdx > 8 || toIdx < 1 ||
		fromIdx > toIdx {
		return 0, fmt.Errorf("invalid input parameter, fromIdx(%d), toIdx(%d)", fromIdx, toIdx)
	}

	var val byte = 0
	for i := fromIdx; i <= toIdx; i++ {
		var tmp byte = 1
		val |= tmp << (i - 1)
	}

	return b & val, nil
}

func BoolToByte(b bool) (bt byte) {
	bt = 1
	if b == false {
		bt = 0
	}
	return bt
}

func BoolToString(b bool) string {
	if b {
		return "True"
	}
	return "False"
}

func Uint16ToBytes(n uint16) []byte {
	return []byte{
		byte(n),
		byte(n >> 8),
	}
}

func BytesToUint16(array []byte) uint16 {
	var data uint16 = 0
	for i := 0; i < len(array); i++ {
		data = data + uint16(uint(array[i])<<uint(8*i))
	}
	return data
}

func Goid() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

// Uint32 to ip
func Long2ip(ip uint32) net.IP {
	a := byte((ip >> 24) & 0xFF)
	b := byte((ip >> 16) & 0xFF)
	c := byte((ip >> 8) & 0xFF)
	d := byte(ip & 0xFF)
	return net.IPv4(a, b, c, d)
}

// ip to uint32 store ad BigEndian
func Ip2long(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip[12:16])
}

func StoreSupportedFeatures(val string) (uint64, error) {
	value, err := strconv.ParseUint(val, 0, 64) //'^[0-9]{5,6}-.+$'
	if err != nil {
		return 0, err
	} else {
		return value, nil
	}
}

func GetInterfaceIp(ifstr string) string {

	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		if i.Name == ifstr {
			addrs, _ := i.Addrs()
			for _, addr := range addrs {
				match, _ := regexp.MatchString(`^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+/[0-9]+$`, addr.String())
				if !match {
					continue
				}
				slit := strings.Split(addr.String(), "/")
				return slit[0]
			}
		}
	}
	return ""
}

func GenRanUeKey(gnbInfo string, ranUeNgApId uint32) string {
	return fmt.Sprintf("%s-ngapid-%d", gnbInfo, ranUeNgApId)
}

func IpToUint32(ip net.IP) uint32 {
	ip4 := ip.To4()
	return binary.BigEndian.Uint32(ip4)
}

// 原子操作
func Uint32AtomicAdd(addr *uint32, delta uint32) {
	atomic.AddUint32(addr, 1)
}

// Conv2ByteSlice get the byte slice from byte pointer
func Conv2ByteSlice(bp *byte, length int) []byte {
	return C.GoBytes(unsafe.Pointer(bp), C.int(length))
}

func SkipIe(msgBuf *bytes.Reader) {
	length, _ := msgBuf.ReadByte()
	leftBytes := make([]byte, length)
	binary.Read(msgBuf, binary.BigEndian, leftBytes)
}

func SkipIeOneByte(msgBuf *bytes.Reader) {
	msgBuf.ReadByte()
}

func SkipBigIe(msgBuf *bytes.Reader) {

	lengthBytes := make([]byte, 2)
	binary.Read(msgBuf, binary.BigEndian, lengthBytes)
	length := binary.BigEndian.Uint16(lengthBytes)

	leftBytes := make([]byte, length)
	binary.Read(msgBuf, binary.BigEndian, leftBytes)
}

func Convert2Bool(v int8) bool {
	switch v {
	case 0:
		return false
	default:
		return true
	}

}
