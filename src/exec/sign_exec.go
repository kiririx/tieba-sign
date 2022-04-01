package exec

import (
	"errors"
	"fmt"
	"strings"
	"tieba-sign/src/log"
	"tieba-sign/src/util"
)

var signFailErr = errors.New("签到失败")

func Sign(ids []uint) {
	for _, id := range ids {
		// 通过id查询bduss
		bduss := GetBduss(id)
		followNum, follow, success := getFollowTieba(bduss.Bduss)
		tbs := getTbs(bduss.Bduss)
		var signFunc = func(tieba string) error {
			rotation := strings.Replace(tieba, "%2B", "+", -1)
			requestBody := make(map[string]string)
			requestBody["kw"] = tieba
			requestBody["tbs"] = tbs
			requestBody["sign"] = util.MD5("kw=" + rotation + "tbs=" + tbs + "tiebaclient!!!")
			if resp, _err := util.DoPost(util.SignUrl, util.ReqParam{Bduss: bduss.Bduss, Params: requestBody}); _err != nil {
				return _err
			} else {
				if resp["error_code"] == "0" {
					success = append(success, rotation)
					log.INFO("签到成功：" + rotation)
				} else {
					return signFailErr
				}
			}
			return nil
		}
		for signIndex := 0; len(success) < followNum; signIndex++ {
			for _, tieba := range follow {
				signFunc(tieba)
			}
			if len(success) < len(follow) {
				// 没签完
			}
		}
	}
}

/**
注入tbs
*/
func getTbs(bduss string) string {
	var tbs string
	content, err := util.DoGet(util.TbsUrl, util.ReqParam{Bduss: bduss})
	if err != nil {
		panic(err)
	}
	if int(content["is_login"].(float64)) == 1 {
		log.INFO("获取tbs成功")
		tbs = content["tbs"].(string)
	}
	return tbs
}

/**
注入关注的贴吧
*/
func getFollowTieba(bduss string) (followNum int, follow []string, success []string) {
	if content, err := util.DoGet(util.LikeUrl, util.ReqParam{Bduss: bduss}); err == nil {
		log.INFO("获取关注列表成功")
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
				fmt.Println("已过签到：" + forumName)
				success = append(success, forumName)
			}
		}
		return followNum, follow, success
	} else {
		panic(err)
	}
}
