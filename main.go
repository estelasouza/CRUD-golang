package main

import (
	"github/estelasouza/api-star-wars/web"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	ctrl := web.NewController()
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/ping").HandlerFunc(ctrl.HandlePing)

	r.Methods(http.MethodPost).Path("/person").HandlerFunc(ctrl.HandleCreatePerson)
	r.Methods(http.MethodGet).Path("/person").HandlerFunc(ctrl.HandleListPeople)

	log.Println("starting server...")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Printf("[ERROR] %s \n", err.Error())
		os.Exit(1)
	}

}
