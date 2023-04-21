package utils

import (
	"io"
	"log"
	"net/http"
	"os"
)

func DisplayHTML(w http.ResponseWriter, filePath string) {
	// Set content type to html
	w.Header().Set("Content-Type", "text/html")

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) { // If file does not exist
			log.Println("File does not exist: " + filePath)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else { // If file exists but another error occurred
			log.Printf("Failed to open file: %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	// Close the file when the function returns
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Failed to close file: %s", err)
		}
	}()

	// Copy file contents to response writer
	_, err = io.Copy(w, file)
	if err != nil {
		log.Printf("Failed to copy file contents to response writer: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
