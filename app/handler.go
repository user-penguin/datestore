package app

import (
	"datestore/helper"
	"datestore/models"
	"encoding/json"
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
		req := models.Date{}
		err := helper.Parse(w, r, &req) // fill the model
		if err != nil {
			log.Println(r.Proto, r.Method, r.RemoteAddr, r.UserAgent(), "can't parse body")
			_, err := fmt.Fprint(w, http.StatusBadRequest, " Bad request")
			if err != nil {
				log.Println(err)
			}
		}
		// here we are start to read data to DB
		id, err := a.DB.PostYear(&req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		req.ID = id
		bytesDate, err := json.Marshal(req)
		fmt.Fprint(w, string(bytesDate))
	}
}
