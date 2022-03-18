package util

import (
	"crypto/md5"
	"fmt"
)

func MD5(source string) string {
	temp := []byte(source)
	return fmt.Sprintf("%x", md5.Sum(temp))
}

func Base64(source string) string {
	return source
}
