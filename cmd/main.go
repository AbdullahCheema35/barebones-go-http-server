package main

import (
	"fmt"

	"github.com/AbdullahCheema35/barebones-go-http-server.git/cmd/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	const body = "Hello World"
	w.WriteHeader(200)
	w.Write([]byte(body))
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	users := []string{"Alice", "Bob", "Charlie"}
	var responseBody = fmt.Appendf(nil, "{\"users\": %v}", users)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(responseBody)))
	w.Write(responseBody)
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/users", handleUsers)
	http.StartHttpServer(":8080", nil)
}
