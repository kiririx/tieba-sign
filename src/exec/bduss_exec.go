package exec

import (
	"tieba-sign/src/db"
	"tieba-sign/src/model"
)

func GetBduss(id uint) model.Bduss {
	var bduss model.Bduss
	db.Db.Find(&bduss, id)
	return bduss
}

func GetAllBduss() []model.Bduss {
	var bdusses []model.Bduss
	db.Db.Find(&bdusses)
	return bdusses
}
