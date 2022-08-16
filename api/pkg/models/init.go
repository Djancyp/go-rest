package models

import (
	"errors"
	"fmt"

	"github.com/Djancyp/go-rest/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Roles struct {
	Role []Role
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Example{})
	db.AutoMigrate(&User{})
	// Add default roles if roles has no value
	db.AutoMigrate(&Role{})
	if err := db.First(&Role{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// data := Roles{[]Role{
		// 	{Role: "superuser", Uuid: "xsuper"},
		// 	{Role: "admin", Uuid: "xadmin"},
		// }}
		// for _, v := range data.Role {
		// 	db.Create(&v)
		//
		// }
		db.Create(&User{
			Email:    "superuser@gmail.com",
			Password: "admin",
			Role: []Role{
				{Role: "superuser", Uuid: "xsuper"},
				{Role: "admin", Uuid: "xadmin"},
			},
		})
		fmt.Println("Insert seed data")
	}
	// }
}
