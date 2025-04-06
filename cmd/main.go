package main

import (
	"github.com/AbdullahCheema35/barebones-go-http-server.git/cmd/http"
)

func handleRoot(res, req string) {
	// Implement
}

func handleEnpoint(res, req string) {
	// Implement
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/users", handleEnpoint)
	http.StartHttpServer(":8080", nil)
}
