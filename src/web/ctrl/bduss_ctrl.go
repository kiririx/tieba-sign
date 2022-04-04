package ctrl

import (
	"tieba-sign/src/db"
	"tieba-sign/src/model"
	"tieba-sign/src/web/rule"
)

func init() {
	err := db.Db.AutoMigrate(&model.Bduss{})
	if err != nil {
		panic(err)
	}
}

func getBdussResp(bduss *model.Bduss) rule.Resp {
	return map[string]interface{}{
		"id":         bduss.ID,
		"name":       bduss.Name,
		"bduss":      bduss.Bduss,
		"signStatus": bduss.SignStatus,
		"signCount":  bduss.SignCount,
	}
}

func GetBdusses(req *rule.Req) rule.Resp {
	var bduss []model.Bduss
	db.Db.Find(&bduss)
	list := make([]map[string]interface{}, 0)
	for _, bduss := range bduss {
		list = append(list, func() map[string]interface{} {
			return getBdussResp(&bduss)
		}())
	}
	return map[string]interface{}{
		"list": list,
	}
}

func GetBduss(req *rule.Req) rule.Resp {
	id := req.GetUIntParam("id")
	bduss := model.Bduss{}
	db.Db.First(&bduss, id)
	return getBdussResp(&bduss)
}
