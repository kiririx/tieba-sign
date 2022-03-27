package exec

import (
	"tieba-sign/src/db"
	"tieba-sign/src/web/ctrl"
)

func GetBduss(id string) ctrl.Bduss {
	var bduss ctrl.Bduss
	db.Db.Find(&bduss, id)
	return bduss
}
