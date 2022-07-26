package middleware

import (
	"Vectorquest/db/model"
	u "Vectorquest/tools/utils"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func Log_in(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		password := r.FormValue("password")
		email := r.FormValue("email")

		//Достать модели
		user := &model.User{}
		session := &model.Session{}

		resp := user.Signin(email, password)

		if resp["status"] != true {
			message := resp["message"].(string)
			w.WriteHeader(401)
			w.Write([]byte(message))
			return
		}

		//Создать id сессии
		sessionId := u.CryptSession(user.Name)

		//Создать токен JWT
		tk := &model.Claims{UserId: user.ID}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		jwtKey := []byte("iota")
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
		}
		//Записать сессию в бд
		session.Sessionid = sessionId
		session.Sessionstore = tokenString
		session.Expire = time.Now().Unix() + 3600*3
		session.AddSession()

		//Установить куки
		http.SetCookie(w, &http.Cookie{
			Name:  "sessionId",
			Value: sessionId,
		})
		http.Redirect(w, r, "/", 302)
	})
}

func LoginAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		password := r.FormValue("password")
		login := r.FormValue("login")

		//Достать модели
		admin := model.Admin_User{}
		session := model.Session{}

		resp := admin.Signin(login, password)

		if resp["status"] != true {
			message := resp["message"].(string)
			w.WriteHeader(401)
			w.Write([]byte(message))
			return
		}
		//Создать id сессии
		sessionId := u.CryptSession(admin.Name)
		userCat := u.GetAdminCat(admin.Id)

		//Создать токен JWT
		tk := model.ClaimsAdmin{
			UserId:     admin.Id,
			UserStatus: admin.Status,
			UserCat:    userCat,
			LifeTime:   time.Now().Unix(),
		}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		jwtKey := []byte("iota")
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
		}

		//Записать сессию в бд
		session.Sessionid = sessionId
		session.Sessionstore = tokenString
		session.Expire = time.Now().Unix() + 3600*3
		session.AddSession()

		//Установить куки
		http.SetCookie(w, &http.Cookie{
			Name:  "sessionId",
			Value: sessionId,
		})
		/*next.ServeHTTP(w, r)*/
		http.Redirect(w, r, "/admin/", 302)
	})
}
