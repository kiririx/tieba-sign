package main

import (
	"fmt"
	"strconv"
	"strings"
	util "tieba-sign/src/util"
	"time"
)

const (
	like_url = "https://tieba.baidu.com/mo/q/newmoindex"
	tbs_url  = "http://tieba.baidu.com/dc/common/tbs"
	sign_url = "http://c.tieba.baidu.com/c/c/forum/sign"
)

var follow []string
var success []string
var tbs string
var BDUSS string

var followNum = 200

func initBDUSS() {
	BDUSS = ""
}

func main() {
	initBDUSS()
	err := wireTbs()
	if err != nil {
		return
	}
	err = wireFollow()
	if err != nil {
		return
	}
	//doSign()
}

/**
注入tbs
*/
func wireTbs() error {
	content, err := util.DoGet(tbs_url, util.ReqParam{Bduss: BDUSS})
	if err != nil {
		return err
	}
	if content["is_login"] == strconv.Itoa(1) {
		info("获取tbs成功")
		tbs = content["tbs"].(string)
	}
	return nil
}

/**
注入关注的贴吧
*/
func wireFollow() error {
	if content, err := util.DoGet(like_url, util.ReqParam{Bduss: BDUSS}); err == nil {
		info("获取关注列表成功")
		data := content["data"].(map[string]interface{})
		dataList := data["like_forum"].([]interface{})
		followNum = len(dataList)
		for _, _data := range dataList {
			v := _data.(map[string]interface{})
			forumName := v["forum_name"].(string)
			isSign := v["is_sign"]
			if int(isSign.(float64)) == 0 {
				fmt.Println("未签到：" + forumName)
				follow = append(follow, strings.Replace(forumName, "+", "%2B", -1))
			} else {
				fmt.Println("已签到：" + forumName)
				success = append(success, forumName)
			}
		}
		return nil
	} else {
		return err
	}
}

/**
开始签到
*/
func doSign() {
	retryNum := 5
	for signIndex := 0; (signIndex < retryNum) && len(success) < followNum; signIndex++ {
		for _, tieba := range follow {
			rotation := strings.Replace(tieba, "%2B", "+", -1)
			requestBody := make(map[string]string)
			requestBody["kw"] = tieba
			requestBody["tbs"] = tbs
			requestBody["sign"] = util.MD5("kw=" + rotation + "tbs=" + tbs + "tiebaclient!!!")
			if resp, _err := util.DoPost(sign_url, util.ReqParam{Bduss: BDUSS, Params: requestBody}); _err != nil {
				err(_err.Error())
			} else {
				if resp["error_code"] == "0" {
					success = append(success, rotation)
				} else {
					err("签到失败")
				}
			}
		}
		if len(success) < len(follow) {
			time.Sleep(time.Minute * 5)
			wireTbs()
		}
		retryNum--

	}

}

func info(msg string) {
	fmt.Println("INFO === " + msg)
}

func warn(msg string) {
	fmt.Println("WARN === " + msg)
}

func err(msg string) {
	fmt.Println("ERROR === " + msg)
}
