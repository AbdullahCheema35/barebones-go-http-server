package main

import (
	"fmt"
	"log"

	"github.com/AbdullahCheema35/barebones-go-http-server.git/cmd/http"
)

func handleRoot(a http.ResponseWriter, b *http.Request) {
	const statusCode = 200
	const statusText = "OK"
	const body = "Hello World"

	a.Write([]byte(fmt.Sprintf("HTTP/1.0 %d %s\r\n", statusCode, statusText)))
	a.Write([]byte(fmt.Sprintf("Content-Type: text/plain\r\n")))
	a.Write([]byte(fmt.Sprintf("Content-Length: %d\r\n", len(body))))
	a.Write([]byte(fmt.Sprintf("\r\n")))
	a.Write([]byte(body))
}

func handleUsers(w http.ResponseWriter, b *http.Request) {
	log.Println("Handling Users")
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/users", handleUsers)
	http.StartHttpServer(":8080", nil)
}
