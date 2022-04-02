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

var RoutesCache = make(map[string]RtConf)

func appendRoute(routes *[]RtConf, method int, url string, handler func(context *gin.Context) rule.Resp) {
	rt := build(method, url, handler)
	*routes = append(*routes, rt)
	RoutesCache[url] = rt
}

func init() {
	appendRoute(&Routes, GET, "/api/bduss", ctrl.HandleBduss)
	appendRoute(&Routes, POST, "/api/sign", ctrl.DoSign)
}
