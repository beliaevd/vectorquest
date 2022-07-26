package server

import (
	"Vectorquest/routes"
	"errors"
	"fmt"
	"net/http"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not authenticated")
)

func Start() error {

	r := routes.Route{}
	router := r.ConfigureRouter()
	fmt.Printf("Server successfully start \n")

	return http.ListenAndServe(":4000", router)
}
