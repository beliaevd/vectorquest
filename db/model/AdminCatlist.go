package model

import "Vectorquest/db"

type Admin_category_link struct {
	ID          int
	User_id     int
	Admin_user  Admin_User
	Category_id int
}

func (link *Admin_category_link) GetAll() []Admin_category_link {
	l := []Admin_category_link{}

	db.GetDB().Find(&l)

	return l
}
