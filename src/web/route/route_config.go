package route

import (
	"github.com/gin-gonic/gin"
	"tieba-sign/src/web/ctrl"
	"tieba-sign/src/web/rule"
)

const (
	GET = iota
	POST
	PUT
	DELETE
)

type RtHandlerFunc func(ctx *gin.Context) rule.Resp

type RtConf struct {
	Method      int
	Url         string
	HandlerFunc RtHandlerFunc
}

func build(method int, url string, handlerFunc RtHandlerFunc) RtConf {
	return RtConf{Method: method, Url: url, HandlerFunc: handlerFunc}
}

var Routes = make([]RtConf, 0)

func init() {
	Routes = append(Routes, build(GET, "/api/bduss", ctrl.HandleBduss))
}
