package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saravanan-babu/go-rest-api/api"
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
