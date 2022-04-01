package ctrl

import (
	"github.com/gin-gonic/gin"
	"tieba-sign/src/exec"
	cvt "tieba-sign/src/util/convert"
	"tieba-sign/src/web/rule"
)

func DoSign(ctx *gin.Context) rule.Resp {
	id := ctx.Query("id")
	idArr := []uint{cvt.StrToUint(id)}
	exec.Sign(idArr)
	return map[string]interface{}{}
}
