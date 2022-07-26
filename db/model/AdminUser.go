package model

import (
	"Vectorquest/db"
	u "Vectorquest/tools/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Admin_User struct {
	Id       int
	Name     string
	Surname  string
	Login    string
	Email    string
	Password string
	Status   string
}

type ClaimsAdmin struct {
	UserId     int
	UserCat    []string
	UserStatus string
	Create     int
	LifeTime   int64
	jwt.StandardClaims
}

type Logsql struct {
	Category string
}

func (admin *Admin_User) Validate() (map[string]interface{}, bool) {
	temp := &Admin_User{}

	//проверка на наличие ошибок и дубликатов электронных писем
	err := db.GetDB().Table("admin_users").Where("email = ?", admin.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(true, "Requirement passed"), true
}

func (admin *Admin_User) Create() map[string]interface{} {
	if resp, ok := admin.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	admin.Password = string(hashedPassword)
	admin.Status = "N"
	db.GetDB().Create(&admin)

	response := u.Message(true, "Account has been created")
	return response
}

func (admin *Admin_User) CheckStatus(login string) map[string]interface{} {
	err := db.GetDB().Table("admin_users").Where("login = ?", login).First(&admin).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Login not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	if admin.Status == "Y" {
		return u.Message(true, "success")
	} else {
		return u.Message(false, "access denied")
	}

}

func (admin *Admin_User) Signin(login, password string) map[string]interface{} {
	err := db.GetDB().Table("admin_users").Where("login = ?", login).First(&admin).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Login not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Password incorrect")
	}

	return u.Message(true, "success")
}

func (admin *Admin_User) GetAll() []Admin_User {
	a := []Admin_User{}

	db.GetDB().Find(&a)

	return a
}
