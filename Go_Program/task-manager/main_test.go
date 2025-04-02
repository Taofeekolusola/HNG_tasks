package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"task-manager/database"
	"task-manager/handlers"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *mux.Router {
	// Initialize database connection (mock or real)
	database.ConnectDB()
	db := database.DB

	// Create router
	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.Register(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.Login(w, r, db)
	}).Methods("POST")

	// File upload route (Public)
	uploadRouter := r.PathPrefix("/uploads").Subrouter()
	uploadRouter.HandleFunc("/upload", handlers.UploadFile).Methods("POST")

	// Serve static files
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	// Task routes
	taskRouter := r.PathPrefix("/tasks").Subrouter()
	taskRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateTask(w, r, db)
	}).Methods("POST")

	taskRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasks(w, r, db)
	}).Methods("GET")

	taskRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateTask(w, r, db)
	}).Methods("PUT")

	taskRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTask(w, r, db)
	}).Methods("DELETE")

	return r
}

func TestRegister(t *testing.T) {
	router := setupRouter() // Make sure this function correctly initializes your router

	payload := map[string]string{"username": "testuser", "password": "password123"}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder() // This captures the response
	router.ServeHTTP(rr, req)    // Process the request

	assert.Equal(t, http.StatusCreated, rr.Code) // Use rr.Code instead of res.Code
}

// Test user login
func TestLogin(t *testing.T) {
	router := setupRouter()

	payload := map[string]string{"username": "testuser", "password": "password123"}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

// Test uploading a file
func TestUploadFile(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("POST", "/uploads/upload", nil) // Mock file upload needed
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code) // Since middleware is used
}

// Test creating a task
func TestCreateTask(t *testing.T) {
	router := setupRouter()

	payload := map[string]string{"title": "New Task", "description": "This is a test task"}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code) // Middleware blocks unauthenticated requests
}

// Test getting tasks
func TestGetTasks(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code) // Middleware blocks unauthenticated requests
}

// Test updating a task
func TestUpdateTask(t *testing.T) {
	router := setupRouter()

	payload := map[string]string{"title": "Updated Task", "description": "Updated task description"}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(body)) // Assuming task with ID 1 exists
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code) // Middleware blocks unauthenticated requests
}

// Test deleting a task
func TestDeleteTask(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("DELETE", "/tasks/1", nil) // Assuming task with ID 1 exists
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code) // Middleware blocks unauthenticated requests
}