package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Task represents a task object
type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// In-memory store for tasks
var tasks []Task

// GetTasks handles GET requests to fetch all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// CreateTask handles POST requests to create a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	tasks = append(tasks, task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// GetTask handles GET requests to fetch a single task by ID
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, task := range tasks {
		if task.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	http.NotFound(w, r)
}

// DeleteTask handles DELETE requests to remove a task by ID
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, task := range tasks {
		if task.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	// Initialize a mux router
	router := mux.NewRouter()

	// Add some sample tasks
	tasks = append(tasks, Task{ID: "1", Title: "Learn Go", Done: false})
	tasks = append(tasks, Task{ID: "2", Title: "Build Go API", Done: false})

	// Define routes
	router.HandleFunc("/tasks", GetTasks).Methods("GET")
router.HandleFunc("/tasks", CreateTask).Methods("POST")
router.HandleFunc("/tasks/{id}", GetTask).Methods("GET")
router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

	// Start the server
	fmt.Println("Server starting on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}