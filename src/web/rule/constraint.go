package rule

import (
	"github.com/gin-gonic/gin"
	"tieba-sign/src/util"
)

type Resp = map[string]interface{}

type Req struct {
	Ctx *gin.Context
}

func (r *Req) GetParam(key string) string {
	return r.Ctx.Param(key)
}

func (r *Req) GetUIntParam(key string) uint {
	v := r.GetParam(key)
	return util.Cvt.StrToUint(v)
}

func (r *Req) GetQuery(key string) string {
	return r.Ctx.Query(key)
}
