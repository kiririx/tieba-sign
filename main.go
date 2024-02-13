package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/kiririx/krutils/algox"
	"github.com/kiririx/krutils/httpx"
	"github.com/kiririx/krutils/jsonx"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	LikeUrl = "https://tieba.baidu.com/mo/q/newmoindex"
	TbsUrl  = "http://tieba.baidu.com/dc/common/tbs"
	SignUrl = "http://c.tieba.baidu.com/c/c/forum/sign"
)

var Logrus *logrus.Logger

func init() {
	Logrus = logrus.New()

	// 使用MultiWriter同时输出到文件和控制台
	Logrus.SetOutput(io.MultiWriter(&lumberjack.Logger{
		Filename:   "./logs/logfile.log", // 日志文件路径，lumberjack会自动根据配置轮转文件
		MaxSize:    10,                   // 文件最大大小（MB）
		MaxBackups: 3,                    // 保留旧文件的最大个数
		MaxAge:     28,                   // 保留旧文件的最大天数
		Compress:   false,                // 是否压缩/归档旧文件
	}, os.Stdout))

	// 设置日志格式为JSON，也可以使用logrus的TextFormatter
	Logrus.SetFormatter(&logrus.JSONFormatter{})
}

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
	defer func() {
		err := recover()
		if err != nil {
			Logrus.Log(logrus.ErrorLevel, err)
			// recover start
			Sign(bduss)
		}
	}()
	// 通过id查询bduss
	followNum, follow, success, signed := getFollowTieba(bduss)
	tbs := getTbs(bduss)

	failTb := make([]string, 0)

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
				Logrus.Info("签到成功：" + rotation)
			} else {
				err := errors.New("签到失败: " + rotation)
				failTb = append(failTb, rotation)
				Logrus.Error(err)
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
	sendNotice(func() string {
		failTbStr := ""
		for _, ft := range failTb {
			failTbStr += "- " + ft + "\n"
		}
		return fmt.Sprintf(`
#### 签到成功%v个吧，失败%v个吧

%s`, len(success), len(failTb), failTbStr)
	}())
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
		Logrus.Info("获取tbs成功")
		tbs = content["tbs"].(string)
	}
	return tbs
}

// getFollowTieba 注入关注的贴吧
func getFollowTieba(bduss string) (followNum int, follow []string, success []string, signed map[string]int) {
	signed = map[string]int{}
	if content, err := httpx.Client().Headers(getHttpHeader(bduss)).GetJSON(LikeUrl, nil); err == nil {
		Logrus.Info("获取关注列表成功")
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
				success = append(success, forumName)
			}
		}
		return followNum, follow, success, signed
	} else {
		panic(err)
	}
}

// send notice to the mobile by pushdeer
func sendNotice(desc string) {
	protocol := os.Getenv("push_protocol")
	if protocol == "" {
		protocol = "http"
	}
	host := os.Getenv("push_host")
	port := os.Getenv("push_port")
	key := os.Getenv("push_key")
	url := fmt.Sprintf("%s://%s:%s/message/push", protocol, host, port)
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("pushkey", key)
	_ = writer.WriteField("text", "百度贴吧:签到成功")
	_ = writer.WriteField("desp", desc)
	err := writer.Close()
	if err != nil {
		logrus.Error(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		logrus.Error(err)
		return
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Error(err)
		return
	}
	result, err := jsonx.JSON2Map(string(body))
	if err != nil {
		logrus.Error(err)
		return
	}
	code := result["code"]
	if code != float64(0) {
		errMsg := result["error"]
		logrus.Error(errMsg)
	}

}
