package utils

import (
	"io"
	"log"
	"net/http"
)

// DisplayHTML is used for displaying HTML files.
func DisplayHTML(w http.ResponseWriter, filePath string) {
	// Set content type to html
	w.Header().Set("Content-Type", "text/html")

	// Open the file
	file, err := OpenFile(filePath)
	if err != nil {
		log.Printf("Failed to open file: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Wrapper for closing the file
	defer CloseFile(file)

	// TODO: Put the copying into a separate function in the /utils/fileHandler.go file?
	// Copy file contents to response writer
	_, err = io.Copy(w, file)
	if err != nil {
		log.Printf("Failed to copy file contents to response writer: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
