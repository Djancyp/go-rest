package models

import (
	"github.com/Djancyp/go-rest/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Example{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})

}
