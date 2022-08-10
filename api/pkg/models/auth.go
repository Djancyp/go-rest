package models

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	UserName     string
	Password     string
	RefreshToken string
	Role         []Roles `gorm:"many2many:user_role;"`
}
type Roles struct {
	gorm.Model
	Name string
	User []*Users `gorm:"many2many:user_role;"`
}
