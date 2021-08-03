package main

import (
	"github/estelasouza/api-star-wars/platform"
	"github/estelasouza/api-star-wars/web"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	repo := platform.NewSWPDiscussionRepository()
	ctrl := web.NewController(repo)
	router := mux.NewRouter()

	router.Methods(http.MethodPost).Path("/discussion").HandlerFunc(ctrl.HandleCreateDiscussion)
	router.Methods(http.MethodGet).Path("/discussion").HandlerFunc(ctrl.HandleListDiscussions)
	router.Methods(http.MethodGet).Path("/discussion/{id}").HandlerFunc(ctrl.HandleGetDiscussion)
	router.Methods(http.MethodPut).Path("/discussion/{id}").HandlerFunc(ctrl.HandleUpdateDiscussion)
	router.Methods(http.MethodDelete).Path("/discussion/{id}").HandlerFunc(ctrl.HandleDelete)

	log.Println("starting server...")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Printf("[ERROR] %s \n", err.Error())
		os.Exit(1)
	}

}
