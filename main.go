package main

import (
	"datestore/app"
	"log"
	"net/http"
	"os"
)

func main() {
	server := app.New()
	http.HandleFunc("/", server.Router.ServeHTTP)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
