package app

import (
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
