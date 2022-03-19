package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io"
	"os"
	"strings"
)

func getDBConfig() (map[string]string, error) {
	file, err := os.Open("./env.properties")
	if err != nil {
		fmt.Println("open file err :", err)
		return nil, err
	}
	defer file.Close()
	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return nil, err
		}
		content = append(content, buf[:n]...)
	}
	props := string(content)
	propArr := strings.Split(props, "\n")

	config := map[string]string{}
	for _, prop := range propArr {
		prop = strings.TrimSpace(prop)
		if len(prop) > 2 && !strings.HasPrefix(prop, "#") {
			key := prop[:strings.Index(prop, "=")]
			val := prop[strings.Index(prop, "=")+1:]
			config[key] = val
		}
	}
	return config, nil
}

func GetBDUSS() ([]string, error) {
	dbConfig, err := getDBConfig()
	if err != nil {
		return nil, err
	}
	username := dbConfig["mysql.conn.user"]
	password := dbConfig["mysql.conn.pass"]
	database := dbConfig["mysql.conn.database"]
	ip := dbConfig["mysql.conn.ip"]
	port := dbConfig["mysql.conn.port"]
	db, err := sqlx.Open("mysql", username+":"+password+"@tcp("+ip+":"+port+")/"+database)
	if err != nil {
		return nil, err
	}
	var bduss []string
	err = db.Select(&bduss, "select bduss from bduss")
	if err != nil {
		return nil, err
	}
	return bduss, nil
}
