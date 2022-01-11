package app

import (
	"datestore/helper"
	"datestore/models"
	"fmt"
	"log"
	"net/http"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "This is Datestore. You're welcome!")
		if err != nil {
			return
		}
		log.Println(request.Proto, request.Method, request.RemoteAddr, request.UserAgent())
	}
}

func (a *App) PostYearHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.DateRequest{}
		err := helper.Parse(w, r, &req)
		if err != nil {
			log.Println(r.Proto, r.Method, r.RemoteAddr, r.UserAgent(), "can't parse body")
			_, err := fmt.Fprint(w, http.StatusBadRequest)
			if err != nil {
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
