package utils

import (
	"assignment-2/constants"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// DisplayDefaultPage is used for displaying HTML files.
func DisplayDefaultPage(w http.ResponseWriter, filePath string) {
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

// DisplayReadme is used for displaying the README.md file. It uses the GitHub Markdown API.
// https://docs.github.com/en/rest/reference/markdown#render-a-markdown-document-in-raw-mode
func DisplayReadme(w http.ResponseWriter, filePath string) error {
	// Set content type to html
	w.Header().Set("Content-Type", "text/html")

	// Open the file
	file, err := OpenFile(filePath)
	if err != nil {
		return err
	}

	// Wrapper for closing the file
	defer CloseFile(file)

	// Get the markdown file contents
	markdownBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Create a JSON object with a "text" property containing the markdown data
	jsonData := map[string]string{"text": string(markdownBytes)}
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	// Create a new HTTP request to the GitHub Markdown API
	req, err := http.NewRequest(http.MethodPost, constants.MarkdownToHTMLApi, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}

	// Set the request headers
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "School project")

	// Send the request and get the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Read the response body
	htmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Write the DOCTYPE declaration, head tag with link tag, charset and resulting HTML in body tag to response writer
	_, err = w.Write([]byte(fmt.Sprintf(constants.ReadmeHtml, constants.DefaultCss, htmlBytes)))
	if err != nil {
		return err
	}

	return err
}
