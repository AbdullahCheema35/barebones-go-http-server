package main

import "github.com/AbdullahCheema35/barebones-go-http-server.git/cmd/server"

func main() {
	server := server.NewServer(":8080")
	server.Start()
}
