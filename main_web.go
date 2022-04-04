package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tieba-sign/src/exec"
	"tieba-sign/src/model"
	"tieba-sign/src/util"
	"tieba-sign/src/web/route"
	"tieba-sign/src/web/rule"
	"time"
)

func task() {
	ticker := time.NewTicker(time.Hour)
	var taskFunc = func() {
		env := util.GetConfig()
		t, _ := strconv.Atoi(env["task.start.hour"])
		if t == time.Now().Hour() {
			bdussModels := exec.GetAllBduss()
			for _, bdussModel := range bdussModels {
				go func(bduss model.Bduss) {
					exec.Sign([]uint{bduss.ID})
				}(bdussModel)
			}
		}
	}
	taskFunc()
	for _ = range ticker.C {
		taskFunc()
	}

}

func main() {
	go task()
	r := gin.Default()
	for _, rt := range route.Routes {
		var callbackFunc = func(context *gin.Context) {
			rt := route.RoutesCache[context.FullPath()]
			req := rule.Req{Ctx: context}
			resp := rt.HandlerFunc(&req)
			context.JSON(200, gin.H{
				"data": resp,
			})
		}
		switch rt.Method {
		case route.GET:
			r.GET(rt.Url, callbackFunc)
		case route.POST:
			r.POST(rt.Url, callbackFunc)
		case route.PUT:
			r.PUT(rt.Url, callbackFunc)
		case route.DELETE:
			r.DELETE(rt.Url, callbackFunc)
		}
	}
	r.Run(":8080")
}
