package defaultHandler

import (
	"assignment-2/constants"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Test cases for testing the default endpoint, these are static and will not change therefore declared here.
var testCases = []struct {
	name string
	path string
}{
	{name: "DefaultEP", path: constants.DefaultEP},
	{name: "BasePath", path: constants.BasePath},
	{name: "BasePath + DefaultEP", path: constants.BasePath + constants.DefaultEP},
}

func TestMain(m *testing.M) {
	// Get current working dir.
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// If the current working dir is not the top lvl, change it.
	if wd[len(wd)-3:] != "assignment-2" {
		err = os.Chdir("..\\..\\")
		if err != nil {
			log.Fatal(err)
		}
	}
	// Run tests
	m.Run()

}

// TestDefaultHandler tests that the http status code is 200 when the default endpoint is called.
func TestDefaultHandler(t *testing.T) {

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(DefaultHandler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		})
	}
}

// TestGetDefault tests that the http status code is 200 when the default endpoint is called.
func TestGetDefault(t *testing.T) {

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(getDefault)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		})
	}
}

// TestGetDefaultBody tests that the body of the response is the same as the default html file.
func TestGetDefaultBody(t *testing.T) {

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(getDefault)

			handler.ServeHTTP(rr, req)

			// Get the content of the defEndpoint.html file as a string.
			defaultHtml, err := os.ReadFile(constants.DefaultHtml)
			if err != nil {
				t.Fatal(err)
			}

			// Checks that the body of the response is the same as the contents of the default html file.
			if rr.Body.String() != string(defaultHtml) {
				t.Errorf("handler returned wrong body: got %v want %v",
					rr.Body.String(), string(defaultHtml))
			}
		})
	}
}
