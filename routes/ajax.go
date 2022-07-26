package routes

import (
	"Vectorquest/db/model"
	"encoding/json"
	"net/http"
	"strconv"
)

//TODO: Вывод ошибок в шаблон
func School_Add(w http.ResponseWriter, r *http.Request) {

	school := model.Schools{}
	s := r.FormValue("name")
	school.Create(s)
}

func School_Get(w http.ResponseWriter, r *http.Request) {
	type Viewdata struct {
		Schools []model.Schools `json:"schools"`
	}
	if r.Method == "GET" {
		school := model.Schools{}
		data := Viewdata{
			school.GetAll(),
		}

		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			panic(err)
		}

		// set header to 'application/json'
		w.Header().Set("Content-Type", "application/json")
		// write the response
		w.Write(b)
	}
}

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		admin := model.Admin_User{
			Email:    r.FormValue("email"),
			Login:    r.FormValue("login"),
			Surname:  r.FormValue("surname"),
			Name:     r.FormValue("name"),
			Password: r.FormValue("password"),
		}
		admin.Create()
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		school := model.Schools{}
		rank := model.Rank{}
		class := model.Class{}
		rank.GetFirstRank()
		school.GetSchool(r.FormValue("school"))
		c := r.FormValue("class")
		gc, _ := strconv.Atoi(c)
		class.GetClass(gc)

		user := model.User{
			Password:    r.FormValue("password"),
			Name:        r.FormValue("name"),
			Surname:     r.FormValue("surname"),
			Middle_name: r.FormValue("patronymic"),
			Login:       "GGG",
			Email:       r.FormValue("email"),
			School_id:   school.ID,
			Class_id:    class.ID,
			Rank_id:     rank.ID,
		}

		user.Create()
	}

}

func UserTableGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		type Viewdata struct {
			U []model.Admin_User          `json:"u"`
			L []model.Admin_category_link `json:"l"`
			C []model.Admin_user_category `json:"c"`
		}
		user := model.Admin_User{}
		link := model.Admin_category_link{}
		category := model.Admin_user_category{}
		data := Viewdata{
			U: user.GetAll(),
			L: link.GetAll(),
			C: category.GetAll(),
		}

		b, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		// set header to 'application/json'
		w.Header().Set("Content-Type", "application/json")
		// write the response
		w.Write(b)
	}
}

func GetDir(w http.ResponseWriter, r *http.Request) {

}
