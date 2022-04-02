package ctrl

import (
	"github.com/gin-gonic/gin"
	"tieba-sign/src/exec"
	"tieba-sign/src/util"
	"tieba-sign/src/web/rule"
)

func DoSign(ctx *gin.Context) rule.Resp {
	id := ctx.Query("id")
	idArr := []uint{util.Cvt.StrToUint(id)}
	exec.Sign(idArr)
	return map[string]interface{}{}
}
