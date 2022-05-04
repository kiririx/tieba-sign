package main

import (
	"flag"
	"tieba-sign/src/exec"
	"time"
)

func main() {
	startHour := flag.Int("h", -1, "输入小时，0-23，默认6")
	bduss := flag.String("b", "", "输入bduss，可以去cookie中查看")
	flag.Parse()
	if *bduss == "" {
		panic("请设置bduss参数 -b")
	}
	if *startHour == -1 {
		exec.Sign(*bduss)
		return
	}
	exec.Sign(*bduss)
	ticker := time.NewTicker(time.Hour)
	var taskFunc = func() {
		if *startHour == time.Now().Hour() {
			exec.Sign(*bduss)
		}
	}
	taskFunc()
	for _ = range ticker.C {
		taskFunc()
	}
}
