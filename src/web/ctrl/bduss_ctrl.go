package ctrl

import (
	"github.com/gin-gonic/gin"
	"tieba-sign/src/web/rule"
)

func HandleBduss(context *gin.Context) rule.Resp {
	list := make([]map[string]string, 0)
	list = append(list, func() map[string]string {
		return map[string]string{
			"name":       "kiririx",
			"bduss":      "xx",
			"signStatus": "0",
			"signCount":  "02",
		}
	}())
	return map[string]interface{}{
		"list": list,
	}
}
