package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := createTask()
	log.Printf("Получена задача: %s", id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"task_id": id})
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	task, ok := getTask(id)
	if !ok {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		log.Printf("❌ Задача не найдена: %s", id)
		return
	}

	log.Printf("Запрошен статус задачи: %s (Статус: %s)", task.ID, task.Status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
