package util

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var envCache = make(map[string]string)

func GetConfig() (map[string]string, error) {
	if len(envCache) > 0 {
		return envCache, nil
	}
	file, err := os.Open("./env.properties")
	if err != nil {
		fmt.Println("open file err :", err)
		return nil, err
	}
	defer file.Close()
	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return nil, err
		}
		content = append(content, buf[:n]...)
	}
	props := string(content)
	propArr := strings.Split(props, "\n")

	for _, prop := range propArr {
		prop = strings.TrimSpace(prop)
		if len(prop) > 2 && !strings.HasPrefix(prop, "#") {
			key := prop[:strings.Index(prop, "=")]
			val := prop[strings.Index(prop, "=")+1:]
			envCache[key] = val
		}
	}
	return envCache, nil
}
