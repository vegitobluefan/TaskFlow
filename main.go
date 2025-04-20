package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	router := mux.NewRouter()
	router.HandleFunc("/task", createTaskHandler).Methods("POST")
	router.HandleFunc("/task/{id}", getTaskHandler).Methods("GET")

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
