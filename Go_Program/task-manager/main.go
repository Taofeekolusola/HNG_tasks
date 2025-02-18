package main

import (
	"log"
	"net/http"
	"task-manager/database"
	"task-manager/handlers"
	"task-manager/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to database
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

	// Serve static files (uploaded files)
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	// File upload route (Public)
	uploadRouter := r.PathPrefix("/uploads").Subrouter()
	uploadRouter.Use(middleware.AuthMiddleware)
	uploadRouter.HandleFunc("/upload", handlers.UploadFile).Methods("POST")

	// Task routes
	taskRouter := r.PathPrefix("/tasks").Subrouter()
	taskRouter.Use(middleware.AuthMiddleware)
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

	// Start server
	log.Println("Server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}



// package main

// import (
// 	"log"
// 	"net/http"
// 	"task-manager/database"
// 	"task-manager/handlers"
// 	"task-manager/middleware"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	// Connect to database
// 	database.ConnectDB()
// 	db := database.DB

// 	// Create router
// 	r := mux.NewRouter()

// 	// Auth routes
// 	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
// 		handlers.Register(w, r, db)
// 	}).Methods("POST")

// 	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
// 		handlers.Login(w, r, db)
// 	}).Methods("POST")

// 	// Serve static files (uploaded files)
// 	upoloadRouter = r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

// 	// File upload route (Public)
// 	upoloadRouter.HandleFunc("/upload", handlers.UploadFile).Methods("POST")

// 	// Task routes
// 	taskRouter := r.PathPrefix("/tasks").Subrouter()
// 	taskRouter.Use(middleware.AuthMiddleware)
// 	taskRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
// 		handlers.CreateTask(w, r, db)
// 	}).Methods("POST")

// 	taskRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
// 		handlers.GetTasks(w, r, db)
// 	}).Methods("GET")

// 	taskRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
// 		handlers.UpdateTask(w, r, db)
// 	}).Methods("PUT")

// 	taskRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
// 		handlers.DeleteTask(w, r, db)
// 	}).Methods("DELETE")

// 	// Start server
// 	log.Println("Server running on port :8080")
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }