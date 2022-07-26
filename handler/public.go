package handler

import (
	"Vectorquest/db/model"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
	// файл base должен быть *последним* файлом в срезе.
	files := []string{
		"./web/template/users/pages/main.tmpl",
		"./web/template/users/base.tmpl",
	}

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func Video(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/template/users/video/video.tmpl",
		"./web/template/users/base.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/template/users/authorization/sign_in.tmpl",
		"./web/template/users/base.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func Video_details(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/template/users/video/details.tmpl",
		"./web/template/users/base.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func Sign_up(w http.ResponseWriter, r *http.Request) {

	type Viewdata struct {
		C []model.Class
	}

	files := []string{
		"./web/template/users/authorization/sign_up.tmpl",
		"./web/template/users/base.tmpl",
	}
	class := model.Class{}

	data := Viewdata{
		class.GetAllClass(),
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func Account(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/template/users/lk/lk.tmpl",
		"./web/template/users/base.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func Calendar(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/template/users/lk/calendar.tmpl",
		"./web/template/users/base.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func Calendarid(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/template/shared/calendar.tmpl",
		"./web/template/admin/base.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internet Server Error", 500)
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internet Server Error", 500)
	}
}

//Blank for Stream

func TestVideo(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/template/public/vd.tmpl",
		"./web/template/users/base.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func Chat(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/template/users/chat/chat.tmpl",
		"./web/template/users/base.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func Testfile(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/template/shared/test.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

//TODO Error handler upload
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
	if status == http.StatusInternalServerError {
		fmt.Fprint(w, "custom 500")
	}
}
