package gctxt

type ImsiKey uint64

type SessionIdKey uint32

type N4SessionIDKey uint64

type SeidKey uint32

type KeyType uint8

const (
	ImsiType           KeyType = 0
	N4SessionIDCxtType KeyType = 1 // 3GPP TS 23.501 V15.3.0 (2018-09) 5.8.2.11.2	N4 Session Context
	SeidType           KeyType = 2
)
