package ctrl

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"tieba-sign/src/db"
	"tieba-sign/src/web/rule"
)

type Bduss struct {
	gorm.Model
	Bduss      string
	Name       string
	SignCount  int
	SignStatus bool
}

func (Bduss) TableName() string {
	return "bduss"
}

func init() {
	err := db.Db.AutoMigrate(&Bduss{})
	if err != nil {
		panic(err)
	}
}

func HandleBduss(context *gin.Context) rule.Resp {
	var bduss []Bduss
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
