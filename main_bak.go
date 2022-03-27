package main

//func _main() {
//	var signFunc = func() {
//		bdussArr := getBDUSS()
//		for _, bduss := range bdussArr {
//			err := wireTbs(bduss)
//			if err != nil {
//				return
//			}
//			err = wireFollow(bduss)
//			if err != nil {
//				return
//			}
//			doSign(bduss)
//		}
//	}
//	var taskFunc = func() {
//		env, err := util.GetConfig()
//		if err != nil {
//			errLog(err.Error())
//			return
//		}
//		t, _ := strconv.Atoi(env["task.start.hour"])
//		if t == time.Now().Hour() {
//			signFunc()
//		}
//	}
//	signFunc()
//	ticker := time.NewTicker(time.Hour)
//	for _ = range ticker.C {
//		taskFunc()
//	}
//}
