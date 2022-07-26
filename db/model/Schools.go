package model

import (
	"Vectorquest/db"
	u "Vectorquest/tools/utils"
	"github.com/jinzhu/gorm"
)

type Schools struct {
	ID   int
	Name string
}

func (school *Schools) Create(name string) map[string]interface{} {

	temp := &Schools{}
	//проверка на наличие ошибок и дубликатов электронных писем
	err := db.GetDB().Table("schools").Where("name = ?", name).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry")
	}
	if temp.Name != "" {
		return u.Message(false, "Школа уже существует.")
	}
	school.Name = name
	db.GetDB().Create(&school)
	return u.Message(true, "Requirement passed")

}

func (school *Schools) GetSchool(u string) *Schools {
	db.GetDB().Where("name = ?", u).First(&school)
	if school.Name == "" { //Школа не найдена
		return nil
	}
	return school
}

func (school *Schools) GetAll() []Schools {
	s := []Schools{}

	db.GetDB().Find(&s)

	return s
}
