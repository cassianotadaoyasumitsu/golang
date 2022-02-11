package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", servidor.UserCreate).Methods(http.MethodPost)
	router.HandleFunc("/users", servidor.UsersSearch).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", servidor.UserSearch).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", servidor.UserUpdate).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", servidor.UserDelete).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta 5001")
	log.Fatal(http.ListenAndServe(":5001", router))
}
