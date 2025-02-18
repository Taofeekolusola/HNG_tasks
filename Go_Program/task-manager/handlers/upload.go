package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// UploadFile handles file uploads
func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Limit request body size to prevent large uploads
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // 10MB limit

	// Parse the form with file
	err := r.ParseMultipartForm(10 << 20) // 10MB max form size
	if err != nil {
		http.Error(w, "File too large", http.StatusRequestEntityTooLarge)
		return
	}

	// Retrieve the file from form-data
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file upload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate file type (e.g., only allow PDF and images)
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".pdf"}
	ext := strings.ToLower(filepath.Ext(handler.Filename))
	valid := false
	for _, ae := range allowedExtensions {
		if ext == ae {
			valid = true
			break
		}
	}

	if !valid {
		http.Error(w, "Invalid file type. Allowed: .jpg, .jpeg, .png, .pdf", http.StatusBadRequest)
		return
	}

	// Create a new file path
	filePath := filepath.Join("uploads", handler.Filename)
	destFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer destFile.Close()

	// Copy file content
	_, err = io.Copy(destFile, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// Send response with file URL
	fileURL := fmt.Sprintf("/uploads/%s", handler.Filename)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{"file_url": "%s"}`, fileURL)))
}