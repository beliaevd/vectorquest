package routes

import (
	Api "Vectorquest/api"
	"Vectorquest/docs"
	"Vectorquest/handler"
	"Vectorquest/middleware"
	"Vectorquest/tools/logs"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Route struct {
	Router *mux.Router
	logger *logrus.Logger
}

func (router *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.Router.ServeHTTP(w, r)
}

func (*Route) ConfigureRouter() *Route {

	router := &Route{
		Router: mux.NewRouter(),
		logger: logrus.New(),
	}

	/*Public routes*/

	router.Router.Use(router.logRequest)
	router.Router.HandleFunc("/", handler.Home)
	router.Router.HandleFunc("/video", handler.Video)
	router.Router.HandleFunc("/test", handler.Home)
	router.Router.HandleFunc("/about", handler.Home)
	router.Router.HandleFunc("/login", handler.Login)
	router.Router.HandleFunc("/registration", handler.Sign_up)
	router.Router.HandleFunc("/video_details", handler.Video_details)
	router.Router.HandleFunc("/lk", handler.Account)
	router.Router.Handle("/auth", middleware.Log_in(http.HandlerFunc(handler.Home)))
	router.Router.HandleFunc("/chat", handler.Chat)
	router.Router.HandleFunc("/lk/calendar", handler.Calendar)
	router.Router.HandleFunc("/calendar", handler.Calendarid)
	router.Router.HandleFunc("/filetest", handler.Testfile)
	/*ajax routers*/

	ajax := router.Router.PathPrefix("/ajax").Subrouter()
	ajax.HandleFunc("/schoolAdd", School_Add)
	ajax.HandleFunc("/schoolGet", School_Get)
	ajax.HandleFunc("/createUser", CreateUser)
	ajax.HandleFunc("/createAdmin", CreateAdmin)
	ajax.HandleFunc("/userTable", UserTableGet)
	ajax.HandleFunc("/addFile", docs.UploadFile)
	ajax.HandleFunc("/makeDir", docs.MakeDir)

	/*api routes*/
	api := router.Router.PathPrefix("/api").Subrouter()

	api.Handle("/v1/addFile", middleware.CheckAccess(http.HandlerFunc(Api.UploadFile)))
	/*admin routes*/

	admin := router.Router.PathPrefix("/admin").Subrouter()

	admin.Handle("/", middleware.CheckAccess(http.HandlerFunc(handler.AdminIndex)))
	admin.Handle("/doc", middleware.CheckAccess(http.HandlerFunc(handler.AdminDocuments)))
	admin.HandleFunc("/login", handler.AdminLogin)
	admin.Handle("/settings", middleware.CheckAccess(http.HandlerFunc(handler.AdminSettings)))
	admin.HandleFunc("/reg", handler.AdminRegistration)
	admin.HandleFunc("/reset_password", handler.AdminReset)
	admin.Handle("/users", middleware.CheckAdminUserAccess(http.HandlerFunc(handler.Users)))
	admin.Handle("/auth", middleware.LoginAdmin(http.HandlerFunc(handler.AdminIndex)))
	admin.Handle("/webinar", middleware.CheckAccess(http.HandlerFunc(handler.Webinar)))
	admin.Handle("/online", middleware.CheckAccess(http.HandlerFunc(handler.AdminOnline)))
	admin.Handle("/create", middleware.CheckAccess(http.HandlerFunc(handler.Create)))
	admin.Handle("/webinar/add", middleware.CheckAccess(http.HandlerFunc(handler.WebinarAdd)))
	admin.Handle("/test/add", middleware.CheckAccess(http.HandlerFunc(handler.TestAdd)))

	/*static files*/
	router.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))
	router.Router.PathPrefix("/disk/").Handler(http.StripPrefix("/disk/", http.FileServer(http.Dir("./disk/"))))

	return router

}

func (router *Route) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := router.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &logs.ResponseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		var level logrus.Level
		switch {
		case rw.Code >= 500:
			level = logrus.ErrorLevel
		case rw.Code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"completed with %d %s in %v",
			rw.Code,
			http.StatusText(rw.Code),
			time.Now().Sub(start),
		)
	})
}
