package main

import (
	"datestore/app"
	"datestore/database"
	"log"
	"net/http"
	"os"
)

func main() {
	server := app.New()
	server.DB = &database.DB{}
	err := server.DB.Open()
	if err != nil {
		panic(err)
	}
	defer func(DB database.DateDB) {
		err := DB.Close()
		if err != nil {
			log.Println(err)
		}
	}(server.DB)

	http.HandleFunc("/", server.Router.ServeHTTP)
	log.Println("App running")
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
