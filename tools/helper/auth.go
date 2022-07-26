package helper

import (
	"Vectorquest/db"
	"Vectorquest/db/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

func GetAdminSession(sessionId string) (*model.ClaimsAdmin, error) {
	s := model.Session{}
	err := db.GetDB().Table("sessions").Where("sessionid = ?", sessionId).Find(&s).Error
	if err != nil {
		return nil, err
	}
	token := s.Sessionstore

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrInvalidKeyType
		}
		return []byte("iota"), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &model.ClaimsAdmin{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, err) {
			return nil, err
		}
		return nil, err
	}
	claims, ok := jwtToken.Claims.(*model.ClaimsAdmin)
	if !ok {
		return nil, jwt.ErrInvalidKeyType
	}

	return claims, nil
}

func GetUserSession(sessionId string) (*model.Claims, error) {
	s := model.Session{}
	err := db.GetDB().Table("sessions").Where("sessionid = ?", sessionId).Find(&s).Error
	if err != nil {
		return nil, err
	}
	token := s.Sessionstore

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrInvalidKeyType
		}
		return []byte("iota"), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &model.Claims{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, err) {
			return nil, err
		}
		return nil, err
	}
	claims, ok := jwtToken.Claims.(*model.Claims)
	if !ok {
		return nil, jwt.ErrInvalidKeyType
	}

	return claims, nil
}
