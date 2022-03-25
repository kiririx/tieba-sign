package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tieba-sign/src/util"
)

var Db *gorm.DB

func init() {
	cfg, err := util.GetConfig()
	if err != nil {
		panic("获取配置信息失败！")
	}
	username := cfg["mysql.conn.user"]
	password := cfg["mysql.conn.pass"]
	database := cfg["mysql.conn.database"]
	ip := cfg["mysql.conn.ip"]
	port := cfg["mysql.conn.port"]
	dsn := username + ":" + password + "@tcp(" + ip + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("获取数据库连接失败！")
	}

}
