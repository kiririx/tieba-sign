package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/kiririx/krutils/algox"
	"github.com/kiririx/krutils/httpx"
	"github.com/kiririx/krutils/logx"
	"strings"
	"time"
)

const (
	LikeUrl = "https://tieba.baidu.com/mo/q/newmoindex"
	TbsUrl  = "http://tieba.baidu.com/dc/common/tbs"
	SignUrl = "http://c.tieba.baidu.com/c/c/forum/sign"
)

func main() {
	startHour := flag.Int("h", -1, "输入小时，0-23，默认6")
	bduss := flag.String("b", "", "输入bduss，可以去cookie中查看")
	flag.Parse()
	if *bduss == "" {
		panic("请设置bduss参数 -b")
	}
	Sign(*bduss)
	if *startHour > 0 {
		ticker := time.NewTicker(time.Hour)
		var taskFunc = func() {
			if *startHour == time.Now().Hour() {
				Sign(*bduss)
			}
		}
		taskFunc()
		for range ticker.C {
			taskFunc()
		}
	}

}

func Sign(bduss string) {
	// 通过id查询bduss
	followNum, follow, success, signed := getFollowTieba(bduss)
	tbs := getTbs(bduss)
	var signFunc = func(tieba string) error {
		if _, ok := signed[tieba]; ok {
			return nil
		}
		rotation := strings.Replace(tieba, "%2B", "+", -1)
		params := fmt.Sprintf("kw=%s&tbs=%s&sign=%s", tieba, tbs, algox.MD5("kw="+rotation+"tbs="+tbs+"tiebaclient!!!"))
		if resp, _err := httpx.Client().Headers(getHttpHeader(bduss)).PostStringGetJSON(SignUrl, params); _err != nil {
			return _err
		} else {
			if resp["error_code"] == "0" {
				success = append(success, rotation)
				logx.INFO("签到成功：" + rotation)
			} else {
				err := errors.New("签到失败: " + rotation)
				logx.ERR(err)
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
	content, err := httpx.Client().Headers(getHttpHeader(bduss)).GetJSON(TbsUrl, nil)
	if err != nil {
		panic(err)
	}
	if int(content["is_login"].(float64)) == 1 {
		logx.INFO("获取tbs成功")
		tbs = content["tbs"].(string)
	}
	return tbs
}

// getFollowTieba 注入关注的贴吧
func getFollowTieba(bduss string) (followNum int, follow []string, success []string, signed map[string]int) {
	signed = map[string]int{}
	if content, err := httpx.Client().Headers(getHttpHeader(bduss)).GetJSON(LikeUrl, nil); err == nil {
		logx.INFO("获取关注列表成功")
		data := content["data"].(map[string]interface{})
		dataList := data["like_forum"].([]interface{})
		followNum = len(dataList)
		for _, _data := range dataList {
			v := _data.(map[string]interface{})
			forumName := v["forum_name"].(string)
			isSign := v["is_sign"]
			if int(isSign.(float64)) == 0 {
				follow = append(follow, strings.Replace(forumName, "+", "%2B", -1))
			} else {
				signed[forumName] = 1
				fmt.Println("已过签到：" + forumName)
				success = append(success, forumName)
			}
		}
		return followNum, follow, success, signed
	} else {
		panic(err)
	}
}
