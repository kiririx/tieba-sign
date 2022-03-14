package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	resp, err := http.Get("http://tieba.baidu.com/dc/common/tbs")
	v, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	} else {
		t.Log(string(v))
	}
	var vt map[string]interface{}
	if json.Unmarshal(v, &vt) == nil {
		t.Log(vt)
	}
}
