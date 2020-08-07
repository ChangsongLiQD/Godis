package main

import "strconv"

func Encode(resp Resp) []byte {
	switch resp.Type {
	case TypeString, TypeInt, TypeError:
		return encodeTextByBytes(resp)
	case TypeBulkBytes:
		return encodeBulkBytesByBytes(resp)
	default:
		return nil
	}
}

func encodeTextByBytes(resp Resp) []byte {
	size := len(resp.Value) + 3
	b := make([]byte, size)
	b[0] = resp.Type
	copy(b[1:], resp.Value)
	b[size-2] = '\r'
	b[size-1] = '\n'
	return b
}

func encodeBulkBytesByBytes(resp Resp) []byte {
	l := len(resp.Value)
	count := []byte(strconv.Itoa(l))
	countLen := len(count)
	size := l + countLen + 5
	b := make([]byte, size)

	b[0] = resp.Type
	copy(b[1:], count)
	b[countLen+1] = '\r'
	b[countLen+2] = '\n'

	copy(b[countLen+3:], resp.Value)
	b[size-2] = '\r'
	b[size-1] = '\n'
	return b
}
