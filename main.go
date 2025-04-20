package main

import (
	"fmt"
	"log"
	"net/http"

	"taskflow/api"

	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	r := mux.NewRouter()
	r.HandleFunc("/task", api.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/task/{id}", api.GetTaskHandler).Methods("GET")

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
