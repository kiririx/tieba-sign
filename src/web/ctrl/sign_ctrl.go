package ctrl

import (
	"tieba-sign/src/exec"
	"tieba-sign/src/util"
	"tieba-sign/src/web/rule"
)

func DoSign(req *rule.Req) rule.Resp {
	id := req.GetQuery("id")
	idArr := []uint{util.Cvt.StrToUint(id)}
	exec.Sign(idArr)
	return map[string]interface{}{}
}
