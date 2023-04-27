package defaultHandler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Test function for default endpoint.
func TestDefaultHandler(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Current working directory: %s", wd)

	err = os.Chdir("..\\..\\")
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodGet, "/", nil)
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
}
