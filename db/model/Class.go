package model

import "Vectorquest/db"

type Class struct {
	ID     int
	Number int
}

func (class *Class) GetClass(u int) *Class {
	db.GetDB().Where("number = ?", u).First(&class)

	return class
}

func (class *Class) GetAllClass() []Class {
	c := []Class{}

	db.GetDB().Find(&c)

	return c
}
