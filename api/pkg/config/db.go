package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	host := GetConfig("MYSQL_HOST")
	database := GetConfig("MYSQL_DATABASE")
	user := GetConfig("MYSQL_USER")
	password := GetConfig("MYSQL_PASSWORD")
	project_name := GetConfig("PROJECT_NAME")
	var err error
	db, err = gorm.Open(host, user+":"+password+"@("+project_name+"-db)/"+database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
