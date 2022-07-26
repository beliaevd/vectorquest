package model

import (
	"Vectorquest/db"
	u "Vectorquest/tools/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type Session struct {
	ID           int
	Sessionid    string
	Sessionstore string
	Expire       int64
}

type Claims struct {
	UserId int
	jwt.StandardClaims
}

func (session *Session) DestroySession() map[string]interface{} {
	err := db.GetDB().Where("sessionId = ?", session.Sessionid).Delete(&session).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Session not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	return u.Message(true, "success")
}

func (session *Session) AddSession() map[string]interface{} {

	err := db.GetDB().Create(&session)
	if err != nil {
		return u.Message(false, "Connection error. Please retry")
	}
	return u.Message(true, "success")
}

func (session *Session) GetSession() *Session {

	db.GetDB().Table("sessions").Where("sessionId = ?", session.Sessionid).First(session)

	return session
}
