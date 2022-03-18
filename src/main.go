package main

import (
	"fmt"
	"strconv"
	"strings"
	go_util "tieba-sign/src/go-util"
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

var followNum = 200

func main() {
	//bduss := os.Args[1]
	cookie := Cookie{""}
	wireTbs()
	wireFollow()

}

/**
注入tbs
*/
func wireTbs() error {
	if content, err := doGet(tbs_url); err == nil {
		if content["is_login"] == strconv.Itoa(1) {
			info("获取tbs成功")
			tbs = content["tbs"].(string)
		}
		return nil
	} else {
		return err
	}
}

/**
注入关注的贴吧
*/
func wireFollow() error {
	if content, err := doGet(like_url); err == nil {
		info("获取关注列表成功")
		dataList := content["data"].(map[string]interface{})["like_forum"].([]map[string]interface{})
		followNum = len(dataList)
		for _, data := range dataList {
			forumName := data["forum_name"].(string)
			if data["is_sign"] == "0" {
				follow = append(follow, strings.Replace(forumName, "+", "%2B", -1))
			} else {
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
			requestBody := make(map[string]interface{})
			requestBody["kw"] = tieba
			requestBody["tbs"] = tbs
			requestBody["sign"] = go_util.MD5("kw=" + rotation + "tbs=" + tbs + "tiebaclient!!!")
			if resp, _err := doPost(sign_url, requestBody); _err != nil {
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
