package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetContacts).Methods("GET")
	r.HandleFunc("/users/{id}", GetContact).Methods("GET")
	r.HandleFunc("/users", CreateContact).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateContact).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteContact).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitialMigration()
	initializeRouter()
}
