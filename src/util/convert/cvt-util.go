package cvt

import (
	"strconv"
	"tieba-sign/src/log"
)

func IntToStr(v int) string {
	return ""
}

func StrToUint(v string) uint {
	iv, err := strconv.Atoi(v)
	if err != nil {
		log.ERROR("convert string to uint fail: ", err)
	}
	return uint(iv)
}
