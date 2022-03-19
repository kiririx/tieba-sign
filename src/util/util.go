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

//func ConvInt(v interface{}) int{
//	switch reflect.TypeOf(v).Kind() {
//	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int16, reflect.Int8:
//		return v.(int)
//	case reflect.Float32, reflect.Float64:
//
//	}
//}
