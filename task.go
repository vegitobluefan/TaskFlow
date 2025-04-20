package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type TaskStatus string

const (
	StatusPending TaskStatus = "ожидание"
	StatusRunning TaskStatus = "выполняется"
	StatusDone    TaskStatus = "завершена"
)

type Task struct {
	ID     string     `json:"id"`
	Status TaskStatus `json:"status"`
	Result string     `json:"result,omitempty"`
}

type TaskStore struct {
	sync.RWMutex
	tasks map[string]*Task
}

var taskStore = &TaskStore{
	tasks: make(map[string]*Task),
}

func createTask() string {
	id := generateID()
	task := &Task{
		ID:     id,
		Status: StatusPending,
	}

	taskStore.Lock()
	taskStore.tasks[id] = task
	taskStore.Unlock()

	go runTask(task)
	return id
}

func getTask(id string) (*Task, bool) {
	taskStore.RLock()
	defer taskStore.RUnlock()
	task, ok := taskStore.tasks[id]
	return task, ok
}

func runTask(task *Task) {
	updateStatus(task.ID, StatusRunning)
	log.Printf("Начато выполнение задачи: %s", task.ID)

	duration := time.Duration(3+rand.Intn(3)) * time.Minute
	log.Printf("Выполнение задачи %s займёт %v", task.ID, duration)
	time.Sleep(duration)

	taskStore.Lock()
	task.Status = StatusDone
	task.Result = fmt.Sprintf("Результат задачи %s", task.ID)
	taskStore.Unlock()

	log.Printf("✅ Задача выполнена: %s", task.ID)
}

func updateStatus(id string, status TaskStatus) {
	taskStore.Lock()
	if task, exists := taskStore.tasks[id]; exists {
		task.Status = status
	}
	taskStore.Unlock()
}
