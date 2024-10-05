package main_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Task struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

var (
	tasks   []Task
	taskID  int
	taskMux sync.Mutex
)

func CreateTask(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Error reading request body!", http.StatusInternalServerError)
		return
	}
	var task Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(res, "Error unmarshalling task!", http.StatusInternalServerError)
		return
	}
	taskMux.Lock()
	taskID++
	task.ID = taskID
	tasks = append(tasks, task)
	taskMux.Unlock()

	fmt.Printf("Task added: %+v\n", task)
	fmt.Printf("Total tasks: %d\n", len(tasks))

	res.Header().Set("Content-Type", "application/json")
	taskJSON, err := json.Marshal(task)
	if err != nil {
		http.Error(res, "Error Marshalling the task", http.StatusInternalServerError)
		return
	}
	res.Write(taskJSON)
}

func ShowTasks(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	taskMux.Lock()
	defer taskMux.Unlock()

	fmt.Printf("Number of tasks: %d\n", len(tasks))

	res.Header().Set("Content-Type", "application/json")

	if len(tasks) == 0 {
		fmt.Println("No tasks found, returning empty array")
		res.Write([]byte("[]"))
		return
	}

	taskJSON, err := json.Marshal(tasks)
	if err != nil {
		fmt.Printf("Error marshalling tasks: %v\n", err)
		http.Error(res, "Error loading the tasks!!", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Returning JSON: %s\n", string(taskJSON))

	_, err = res.Write(taskJSON)
	if err != nil {
		fmt.Printf("Error writing response: %v\n", err)
	}
}
