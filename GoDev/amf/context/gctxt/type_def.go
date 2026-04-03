package gctxt

type ImsiKey uint64
type GutiKey string
type StmsiKey string
type AmfUeNgApId uint64
type KeyType uint8

const (
	ImsiType KeyType = iota
	GutiType
	StmsiType
	AmfUeNgApIdType
	RanUeNgApIdType
	RanUeKeyType
)
