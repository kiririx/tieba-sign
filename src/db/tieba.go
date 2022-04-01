package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"tieba-sign/src/log"
	"tieba-sign/src/util"
)

func GetBDUSS() []string {
	dbConfig := util.GetConfig()
	username := dbConfig["mysql.conn.user"]
	password := dbConfig["mysql.conn.pass"]
	database := dbConfig["mysql.conn.database"]
	ip := dbConfig["mysql.conn.ip"]
	port := dbConfig["mysql.conn.port"]
	db, err := sqlx.Open("mysql", username+":"+password+"@tcp("+ip+":"+port+")/"+database)
	if err != nil {
		log.ERROR("get database connection occur error: %s", err)
	}
	var bduss []string
	err = db.Select(&bduss, "select bduss from bduss")
	if err != nil {
		log.ERROR("query the bduss table occur error", err)
	}
	return bduss
}
