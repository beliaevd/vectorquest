package model

import "Vectorquest/db"

type Webinar_template struct {
	ID          int
	Name        string
	Description string
	Maxpeople   int
	Img         string
}

func (wt *Webinar_template) GetAll() []Webinar_template {
	temp := []Webinar_template{}

	db.GetDB().Find(&wt)
	return temp
}
