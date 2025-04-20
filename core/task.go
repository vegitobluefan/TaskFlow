package core

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

var (
	tasks = make(map[string]*Task)
	mu    sync.RWMutex
)

func CreateTask() string {
	id := generateID()
	task := &Task{
		ID:     id,
		Status: StatusPending,
	}

	mu.Lock()
	tasks[id] = task
	mu.Unlock()

	go runTask(task)

	return id
}

func GetTask(id string) (*Task, bool) {
	mu.RLock()
	defer mu.RUnlock()
	t, ok := tasks[id]
	return t, ok
}

func runTask(task *Task) {
	updateStatus(task.ID, StatusRunning)
	log.Printf("Начато выполнение задачи: %s", task.ID)

	duration := time.Duration(3+rand.Intn(3)) * time.Minute
	log.Printf("Выполнение задачи %s займёт %v", task.ID, duration)
	time.Sleep(duration)

	mu.Lock()
	task.Status = StatusDone
	task.Result = fmt.Sprintf("Результат задачи %s", task.ID)
	mu.Unlock()

	log.Printf("✅ Выполнена задача: %s", task.ID)
}

func updateStatus(id string, status TaskStatus) {
	mu.Lock()
	if task, exists := tasks[id]; exists {
		task.Status = status
	}
	mu.Unlock()
}
