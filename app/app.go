package app

import (
	"datestore/database"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     database.DateDB
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/years", a.PostYearHandler()).Methods("POST")
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}
