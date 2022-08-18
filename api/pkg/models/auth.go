package models

import (
	"errors"
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
	Email              string    `gorm:"primaryKey;unique;not null" json:"email"`
	Password           string    `gorm:"not null" json:"password"`
	Token              string    `gorm:"not null" json:"token"`
	IsActive           bool      `gorm:"not null" json:"is_active"`
	Roles              []Role    `gorm:"many2many:user_role;" json:"roles"`
	ForgotenPassworJWT string    `json:"forgoten_password_jwt"`
}
type User_Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Role struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Role        string    `gorm:"unique" json:"role"`
	Description string    `json:"description"`
}

type User_Role struct {
	User_id uint64 `json:"user_id"`
	Role_id uint64 `json:"role_id"`
}
type Login struct {
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type LoginResult struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Roles []Role `json:"roles"`
}

func (e *Login) Login() (*LoginResult, *gorm.DB) {
	var user User
	db := db.Table("users").Where("email = ?", e.Email).Scan(&user)
	db.Preload("Roles").First(&user)
	if db.First(&user).RecordNotFound() {
		return nil, db
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(e.Password))
	if err != nil {
		return nil, db
	}
	result := LoginResult{
		ID:    user.ID,
		Email: user.Email,
		Roles: user.Roles,
	}
	return &result, db
}
func (e *User) Register() (*User, error) {
	err, _ := EmailValidate(e.Email)
	if err == true {
		return nil, fmt.Errorf("Email already exists")
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(e.Password), 12)
	token := utils.CreateJwtString()
	var roles = []Role{}
	if err := db.First(&User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		roles = []Role{
			{ID: 1, Role: "superuser", Description: "Use for superprevilage. Access level 1"},
		}
	} else {
		roles = []Role{
			{ID: 3, Role: "user", Description: "User privilage"},
		}
	}
	user := User{
		Email:    e.Email,
		Password: string(password),
		Token:    token,
		Roles:    roles,
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
