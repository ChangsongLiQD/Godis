package sds

func sdsNewLen(init string) Sds{
	sds := &sdsHdr{
		buf: []byte(init),
	}
	return sds
}

func SdsEmpty() Sds{
	return sdsNewLen("")
}

func SdsNew(init string) Sds{
	return sdsNewLen(init)
}

func SdsDup(s Sds) Sds{
	return sdsNewLen(string(s.sdsBuf()))
}

func SdsClear(s Sds) {
	s.clearBuf()
}

