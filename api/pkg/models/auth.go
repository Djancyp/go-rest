package models

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Token    string `gorm:"not null" json:"token"`
	IsActive bool   `gorm:"not null" json:"is_active"`
	Role     []Role `gorm:"foreignKey:ID"`
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
	var jwtKey = []byte("my_secret_key")
	password, _ := bcrypt.GenerateFromPassword([]byte(e.Password), 12)
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString(jwtKey)
	user := User{
		Email:    e.Email,
		Password: string(password),
		Token:    tokenString,
	}
	db.NewRecord(user)
	db.Create(&user)
	return &user
}

func GetUserById(Id uint64) (*User, *gorm.DB) {
	var user User
	db := db.Where("id = ?", Id).Find(&user)
	fmt.Println(user)
	return &user, db
}

func EmailValidate(email string) (bool, *User) {
	var user User
	db := db.Where("email = ?", email).First(&User{})

	if db.First(&user).RecordNotFound() {
		return false, nil
	}
	return true, &user
}
