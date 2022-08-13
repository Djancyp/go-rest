package models

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email           string `gorm:"unique;not null" json:"email"`
	Password        string `gorm:"not null" json:"password"`
	Activation_code string `gorm:"not null" json:"refresh_token"`
	IsActive        bool   `gorm:"not null" json:"is_active"`
	Role            []Role `gorm:"foreignKey:ID"`
}
type Role struct {
	gorm.Model
	Name string
}

type Login struct {
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password" json:"-"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (e *Login) Login() (*User, *gorm.DB) {
	var user User
	db := db.Where(map[string]interface{}{"email": e.Email}).First(&user)
	if db.First(&user).RecordNotFound() {
		return nil, db
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(e.Password))
	if err != nil {
		return nil, db
	}
	return &user, db
}
func (e *User) Register() *User {
	password, _ := bcrypt.GenerateFromPassword([]byte(e.Password), 12)
	user := User{
		Email:    e.Email,
		Password: string(password),
	}
	db.NewRecord(user)
	db.Create(&user)
	return &user
}

func EmailValidate(email string) (bool, *User) {
	var user User
	db := db.Where("email = ?", email).First(&User{})

	if db.First(&user).RecordNotFound() {
		return false, nil
	}
	return true, &user
}
