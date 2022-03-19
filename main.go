package main

import (
	"fmt"
	"strings"
	"tieba-sign/src/db"
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

var followNum = 200

func getBDUSS() []string {
	bdussArr, err := db.GetBDUSS()
	if err != nil {
		fmt.Println("bduss获取失败")
		return nil
	}
	return bdussArr
}

func main() {
	bdussArr := getBDUSS()
	for _, bduss := range bdussArr {
		err := wireTbs(bduss)
		if err != nil {
			return
		}
		err = wireFollow(bduss)
		if err != nil {
			return
		}
		doSign(bduss)
	}

}

/**
注入tbs
*/
func wireTbs(bduss string) error {
	content, err := util.DoGet(tbs_url, util.ReqParam{Bduss: bduss})
	if err != nil {
		return err
	}
	if int(content["is_login"].(float64)) == 1 {
		info("获取tbs成功")
		tbs = content["tbs"].(string)
	}
	return nil
}

/**
注入关注的贴吧
*/
func wireFollow(bduss string) error {
	if content, err := util.DoGet(like_url, util.ReqParam{Bduss: bduss}); err == nil {
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
func doSign(bduss string) {
	retryNum := 5
	var signFunc = func(tieba string) {
		rotation := strings.Replace(tieba, "%2B", "+", -1)
		requestBody := make(map[string]string)
		requestBody["kw"] = tieba
		requestBody["tbs"] = tbs
		requestBody["sign"] = util.MD5("kw=" + rotation + "tbs=" + tbs + "tiebaclient!!!")
		if resp, _err := util.DoPost(sign_url, util.ReqParam{Bduss: bduss, Params: requestBody}); _err != nil {
			err(_err.Error())
		} else {
			if resp["error_code"] == "0" {
				success = append(success, rotation)
			} else {
				err("签到失败")
			}
		}
	}
	for signIndex := 0; (signIndex < retryNum) && len(success) < followNum; signIndex++ {
		for _, tieba := range follow {
			signFunc(tieba)
		}
		if len(success) < len(follow) {
			time.Sleep(time.Minute * 5)
			wireTbs(bduss)
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
