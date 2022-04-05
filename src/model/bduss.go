package model

import "gorm.io/gorm"

type Bduss struct {
	gorm.Model
	Bduss      string
	Name       string
	SignCount  int
	SignStatus bool
	Auto       bool
}

func (Bduss) TableName() string {
	return "bduss"
}
