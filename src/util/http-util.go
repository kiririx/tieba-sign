package util

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type ReqParam struct {
	Bduss  string
	Params map[string]string
}

func DoGet(url string, param ReqParam) (map[string]interface{}, error) {
	client := new(http.Client)
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		return nil, err
	}
	initHeader(&req.Header, param.Bduss)
	if resp, err := client.Do(req); err == nil {
		defer func(Body io.ReadCloser) {
			Body.Close()
		}(resp.Body)
		var content map[string]interface{}
		if v, err := ioutil.ReadAll(resp.Body); err == nil {
			if err = json.Unmarshal(v, &content); err == nil {
				return content, nil
			}
		}
	}
	return nil, errors.New("http fail")
}

func DoPost(url string, param ReqParam) (map[string]interface{}, error) {
	var bodyString string
	i := 0
	for k, v := range param.Params {
		if i > 0 {
			bodyString += "&"
		}
		bodyString += k + "=" + v
		i++
	}
	client := new(http.Client)
	req, err := http.NewRequest("POST", url, strings.NewReader(bodyString))
	if err != nil {
		return nil, err
	}
	initHeader(&req.Header, param.Bduss)
	{
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer func(Body io.ReadCloser) {
			Body.Close()
		}(resp.Body)
		var content map[string]interface{}
		v, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(v, &content)
		if err != nil {
			return nil, err
		}
		return content, nil
	}
}

func initHeader(header *http.Header, bduss string) {
	header.Set("connection", "keep-alive")
	header.Set("Content-Type", "application/x-www-form-urlencoded")
	header.Set("charset", "UTF-8")
	header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	header.Set("Cookie", "BDUSS="+bduss)
}
