package algorithm

import (
	"crypto/md5"
	"fmt"
)

type Algorithm struct {
}

func (*Algorithm) MD5(source string) string {
	temp := []byte(source)
	return fmt.Sprintf("%x", md5.Sum(temp))
}
