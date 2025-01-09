package main

import (
	"errors"
	"fmt"
	"github.com/kiririx/krutils/ut"
	"github.com/tidwall/gjson"
	"log"
	"os"
	"strings"
)

const (
	LikeUrl = "https://tieba.baidu.com/mo/q/newmoindex"
	TbsUrl  = "http://tieba.baidu.com/dc/common/tbs"
	SignUrl = "http://c.tieba.baidu.com/c/c/forum/sign"
)

func main() {
	Sign()
}

func Sign() {
	defer func() {
		err := recover()
		if err != nil {
			Sign()
		}
	}()
	bduss := os.Getenv("bduss")
	if bduss == "" {
		log.Println("请设置bduss环境变量")
		return
	}
	// 通过id查询bduss
	success := make([]string, 0)
	followNum, follow, signed := getFollowTieba(bduss)
	tbs := getTbs(bduss)
	var signFunc = func(tieba string) error {
		if _, ok := signed[tieba]; ok {
			return nil
		}
		rotation := strings.Replace(tieba, "%2B", "+", -1)
		params := fmt.Sprintf("kw=%s&tbs=%s&sign=%s", tieba, tbs, ut.Algorithm().MD5("kw="+rotation+"tbs="+tbs+"tiebaclient!!!"))
		if resp, _err := ut.HttpClient().Headers(getHttpHeader(bduss)).PostStringGetJSON(SignUrl, params); _err != nil {
			return _err
		} else {
			if resp["error_code"] == "0" {
				success = append(success, rotation)
				log.Println("签到成功：" + rotation)
			} else {
				err := errors.New("签到失败: " + rotation)
				log.Println(err)
				return err
			}
		}
		return nil
	}
	retry := 0
	for signIndex := 0; len(success) < followNum && retry < 5; signIndex++ {
		retry++
		for _, tieba := range follow {
			_ = signFunc(tieba)
		}
		if len(success) == len(follow) {
			break
		}
	}
	pushMsg(success, follow)
}

func removeElements(a, b []string) []string {
	result := []string{}
	for _, v := range a {
		if !contains(b, v) {
			result = append(result, v)
		}
	}
	return result
}
func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func pushMsg(success []string, follow []string) {
	addr := os.Getenv("pushdeer.addr")
	if addr == "" {
		return
	}
	title := "百度贴吧签到完成"
	subtitle := fmt.Sprintf("成功: %v 失败: %v", len(success), len(follow)-len(success))
	// 得到签到失败的贴吧, 拼接字符串
	desc := fmt.Sprintf("失败: %s", strings.Join(removeElements(follow, success), ","))
	result, _ := ut.HttpClient().PostFormGetJSON(addr, map[string]string{
		"pushkey":  os.Getenv("pushdeer.pushkey"),
		"subtitle": subtitle,
		"text":     title,
		"desp":     desc,
	})
	log.Println(result)
}

func getHttpHeader(bduss string) map[string]string {
	return map[string]string{
		"connection":   "keep-alive",
		"Content-Type": "application/x-www-form-urlencoded",
		"charset":      "UTF-8",
		"User-Agent":   "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36",
		"Cookie":       "BDUSS=" + bduss,
	}
}

// getTbs 注入tbs
func getTbs(bduss string) string {
	var tbs string
	content, err := ut.HttpClient().Headers(getHttpHeader(bduss)).GetJSON(TbsUrl, nil)
	if err != nil {
		panic(err)
	}
	if int(content["is_login"].(float64)) == 1 {
		log.Println("获取tbs成功")
		tbs = content["tbs"].(string)
	}
	return tbs
}

// getFollowTieba 注入关注的贴吧
func getFollowTieba(bduss string) (followNum int, follow []string, signed map[string]int) {
	signed = map[string]int{}
	if content, err := ut.HttpClient().Headers(getHttpHeader(bduss)).GetString(LikeUrl, nil); err == nil {
		log.Println("获取关注列表成功")
		data := gjson.Get(content, "data").Map()
		dataList := data["like_forum"].Array()
		followNum = len(dataList)
		for _, _data := range dataList {
			v := _data.Map()
			forumName := v["forum_name"].String()
			isSign := v["is_sign"].Float()
			if isSign == 0 {
				follow = append(follow, strings.Replace(forumName, "+", "%2B", -1))
			} else {
				signed[forumName] = 1
				log.Println("已过签到：" + forumName)
			}
		}
		return followNum, follow, signed
	} else {
		panic(err)
	}
}
