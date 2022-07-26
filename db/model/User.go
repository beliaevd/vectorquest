package model

import (
	"Vectorquest/db"
	u "Vectorquest/tools/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	ID          int
	Name        string
	Surname     string
	Middle_name string
	Login       string
	Email       string
	Level       int
	School_id   int
	Schools     Schools
	Class_id    int
	Class       Class
	Rank_id     int
	Rank        Rank
	Password    string
}

func (user *User) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}
	if len(user.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Email должен быть уникальным
	temp := &User{}

	//проверка на наличие ошибок и дубликатов электронных писем
	err := db.GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(true, "Requirement passed"), true
}

func (user *User) Create() map[string]interface{} {

	if resp, ok := user.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	db.GetDB().Create(user)

	return u.Message(true, "Account has been created")

}

func (user *User) Signin(email, password string) map[string]interface{} {

	err := db.GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Пароль не совпадает!!
		return u.Message(false, "Password incorrect")
	}
	return u.Message(true, "success")
}
