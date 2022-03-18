package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

func doGet(url string) (map[string]interface{}, error) {
	if resp, err := http.Get(url); err == nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				warn("http关闭失败")
			}
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

func doPost(url string, body map[string]interface{}) (map[string]interface{}, error) {
	if resp, err := http.Post(url, "application/x-www-form-urlencoded"); err == nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				warn("http关闭失败")
			}
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
