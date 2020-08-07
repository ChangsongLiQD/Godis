package main

type RespType byte

//const (
//	MaxBulkByteLen = 1024 * 1024 * 512
//	MaxArrayLen    = 1024 * 1024
//)

const (
	TypeString    = '+'
	TypeError     = '-'
	TypeInt       = ':'
	TypeBulkBytes = '$'
	TypeArray     = '*'
)

var (
	RespNil = []byte("nil")
	RespOk  = []byte("OK")
)

type Resp struct {
	Type  byte
	Value []byte
	Array []*Resp
}

func GetStringResponse(b []byte) []byte {
	resp := Resp{Type: TypeString, Value: b}
	return Encode(resp)
}

func GetErrorResponse(b []byte) []byte {
	resp := Resp{Type: TypeError, Value: b}
	return Encode(resp)
}

func GetBulkBytesResponse(b []byte) []byte {
	resp := Resp{Type: TypeBulkBytes, Value: b}
	return Encode(resp)
}
