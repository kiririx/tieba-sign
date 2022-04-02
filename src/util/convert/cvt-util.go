package cvt

import (
	"strconv"
	"tieba-sign/src/log"
)

type Convert struct {
}

func (*Convert) IntToStr(v int) string {
	return ""
}

func (*Convert) StrToUint(v string) uint {
	iv, err := strconv.Atoi(v)
	if err != nil {
		log.ERROR("convert string to uint fail: ", err)
	}
	return uint(iv)
}
