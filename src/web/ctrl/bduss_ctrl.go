package ctrl

import (
	"github.com/gin-gonic/gin"
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

func HandleBduss(context *gin.Context) rule.Resp {
	var bduss []model.Bduss
	db.Db.Find(&bduss)
	list := make([]map[string]interface{}, 0)
	for _, bduss := range bduss {
		list = append(list, func() map[string]interface{} {
			return map[string]interface{}{
				"id":         bduss.ID,
				"name":       bduss.Name,
				"bduss":      bduss.Bduss,
				"signStatus": bduss.SignStatus,
				"signCount":  bduss.SignCount,
			}
		}())
	}
	return map[string]interface{}{
		"list": list,
	}
}
