package ctrl

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tieba-sign/src/db"
	"tieba-sign/src/web/rule"
)

type Bduss struct {
	gorm.Model
	Id    string
	Bduss string
}

func (Bduss) TableName() string {
	return "bduss"
}

func HandleBduss(context *gin.Context) rule.Resp {
	var bduss []Bduss
	db.Db.Find(&bduss)
	list := make([]map[string]string, 0)
	for _, bduss := range bduss {
		list = append(list, func() map[string]string {
			return map[string]string{
				"id":         bduss.Id,
				"name":       "name",
				"bduss":      bduss.Bduss,
				"signStatus": "0",
				"signCount":  "02",
			}
		}())
	}
	return map[string]interface{}{
		"list": list,
	}
}
