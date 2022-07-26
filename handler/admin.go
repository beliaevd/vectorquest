package handler

import (
	"html/template"
	"log"
	"net/http"
)

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/" {
		ErrorHandlerAdmin(w, r, 404)
		return
	}

	// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
	// файл base должен быть *последним* файлом в срезе.

	files := []string{
		"./web/template/admin/pages/index.tmpl",
		"./web/template/admin/base.tmpl",
	}

	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	//Сюда идут данные из бд

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}
}

func AdminSettings(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/settings" {
		ErrorHandlerAdmin(w, r, 404)
		return
	}

	files := []string{
		"./web/template/admin/pages/settings.tmpl",
		"./web/template/admin/base.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}
}

func AdminDocuments(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/doc" {
		ErrorHandlerAdmin(w, r, 404)
		return
	}

	files := []string{
		"./web/template/admin/pages/docs.tmpl",
		"./web/template/admin/base.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}
}

func AdminLogin(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/admin/login" {
		ErrorHandlerAdmin(w, r, 404)
		return
	}

	files := []string{
		"./web/template/admin/authorization/login.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}
}

func AdminRegistration(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/admin/reg" {
		ErrorHandlerAdmin(w, r, 404)
		return
	}
	files := []string{
		"./web/template/admin/authorization/signup.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}
}

func AdminReset(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/admin/reset_password" {
		ErrorHandlerAdmin(w, r, 404)
		return
	}

	files := []string{
		"./web/template/admin/authorization/reset-password.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/users" {
		ErrorHandlerAdmin(w, r, 404)
		return
	}

	files := []string{
		"./web/template/admin/private/users.tmpl",
		"./web/template/admin/base.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}
}

func Webinar(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/template/admin/pages/webinar.tmpl",
		"./web/template/admin/base.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}
}

func AdminOnline(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/template/admin/online/online.tmpl",
		"./web/template/admin/base.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/template/admin/create/create.tmpl",
		"./web/template/admin/base.tmpl",
	}
	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}
}

func WebinarAdd(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./web/template/admin/create/webinarAdd.tmpl",
		"./web/template/admin/base.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}

}

func TestAdd(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./web/template/admin/create/testAdd.tmpl",
		"./web/template/admin/base.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}
}

func AdminAccessDenied(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./web/template/admin/error/accessdenied.tmpl",
	}

	ts, _ := template.ParseFiles(files...)
	err := ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		ErrorHandlerAdmin(w, r, 500)
	}
}

func ErrorHandlerAdmin(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		files := []string{
			"./web/template/admin/error/404.tmpl",
		}

		ts, _ := template.ParseFiles(files...)
		_ = ts.Execute(w, nil)
	}
	if status == http.StatusInternalServerError {
		files := []string{
			"./web/template/admin/error/505.tmpl",
		}

		ts, _ := template.ParseFiles(files...)
		_ = ts.Execute(w, nil)
	}
}
