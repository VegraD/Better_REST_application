package notificationHandler

import (
	"fmt"
	"net/http"
)

func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println("Get request")
	case http.MethodPost:
		fmt.Println("Post request")
	case http.MethodDelete:
		fmt.Println("Delete request")
	default:
		http.Error(w, r.Method+" not supported, use "+http.MethodGet+"!", http.StatusMethodNotAllowed)
	}
}
