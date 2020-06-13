package sds

/**
没有像REDIS一样进行区分长度(8, 16, 32, 64)，Redis通过拿buf的前
一个char，结合mask获得具体的长度类型，再去拿len和alloc。十分巧妙，
这里将进行简单实现，在考虑后期优化。
 */
type sdsHdr struct {
	buf   []byte
}

type Sds interface {
	sdsLen() int
	sdsAvail() int
	sdsAlloc() int
	sdsBuf() []byte
	clearBuf()
}

func (s *sdsHdr)sdsLen() int{
	return len(s.buf)
}

func (s *sdsHdr)sdsAvail() int{
	return cap(s.buf) - len(s.buf)
}

/* sdsAlloc() = sdsAvail() + sdsLen() */
func (s *sdsHdr)sdsAlloc() int{
	return cap(s.buf)
}

func (s *sdsHdr)sdsBuf() []byte{
	return s.buf
}

func (s *sdsHdr)clearBuf(){
	s.buf = nil
}

