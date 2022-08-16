package models

import (
	"fmt"
	"github.com/Djancyp/go-rest/pkg/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID                 uint64    `gorm:"primaryKey" json:"id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Email              string    `gorm:"unique;not null" json:"email"`
	Password           string    `gorm:"not null" json:"password"`
	Token              string    `gorm:"not null" json:"token"`
	IsActive           bool      `gorm:"not null" json:"is_active"`
	Role               []Role    `gorm:"many2many:user_role;"`
	ForgotenPassworJWT string    `json:"forgoten_password_jwt"`
}
type Role struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Role      string    `json:"role"`
	Uuid      string    `gorm:"unique;not null" json:"uuid"`
}

type Login struct {
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (e *Login) Login() (*User, *gorm.DB) {
	var user User
	db := db.Table("users").Where("email = ?", e.Email).Scan(&user)
	if db.First(&user).RecordNotFound() {
		return nil, db
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(e.Password))
	if err != nil {
		return nil, db
	}
	return &user, db
}
func (e *User) Register() (*User, error) {
	err, _ := EmailValidate(e.Email)
	if err == true {
		return nil, fmt.Errorf("Email already exists")
	}
	fmt.Println(err)
	password, _ := bcrypt.GenerateFromPassword([]byte(e.Password), 12)
	token := utils.CreateJwtString()
	user := User{
		Email:    e.Email,
		Password: string(password),
		Token:    token,
	}
	db.NewRecord(user)
	db.Create(&user)
	return &user, nil
}
func (e *User) UpdatePassword() (*User, *gorm.DB) {
	var user User
	db := db.Table("users").Where("email = ?", e.Email).Scan(&user)
	if db.First(&user).RecordNotFound() {
		return nil, db
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(e.Password), 12)
	userSave := User{
		Email:    e.Email,
		Password: string(password),
	}

	db.Model(&user).Where("email = ?", e.Email).Updates(userSave)
	return &user, db
}
func GetUserById(Id uint64) (*User, *gorm.DB) {
	var user User
	db := db.Where("id = ?", Id).Find(&user)
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
