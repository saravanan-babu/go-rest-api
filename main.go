package main

import (
	"log"
	"net/http"

	"./api"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", api.Welcome).Methods("GET")
	router.HandleFunc("/users", api.GetUsers).Methods("GET")
	router.HandleFunc("/users", api.UpdateUsers).Methods("POST", "PUT")

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
