package models

import (
	"errors"
	"fmt"

	"github.com/Djancyp/go-rest/pkg/config"
	"github.com/jinzhu/gorm"
	// "golang.org/x/crypto/bcrypt"
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

		roles := []Role{
			{Role: "superuser", Description: "Use for superprevilage. Access level 1"},
			{Role: "admin"},
			{Role: "user"},
		}
		for _, role := range roles {
			db.Create(&role)
		}
		fmt.Println("Insert seed data")
	}
}
