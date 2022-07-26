package model

import "Vectorquest/db"

type Admin_user_category struct {
	ID   int
	Name string
}

func (category *Admin_user_category) GetAll() []Admin_user_category {
	c := []Admin_user_category{}

	db.GetDB().Find(&c)
	return c
}
