package main

import (
	"github.com/gin-gonic/gin"
	"tieba-sign/src/web/route"
)

func main() {
	r := gin.Default()
	for _, rt := range route.Routes {
		var callbackFunc = func(context *gin.Context) {
			resp := rt.HandlerFunc(context)
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
