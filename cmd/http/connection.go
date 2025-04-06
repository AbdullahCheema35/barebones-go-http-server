package http

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	var buff = make([]byte, 1024)

	n, err := conn.Read(buff)
	if err != nil {
		fmt.Fprintf(conn, "HTTP/1.0 400 Bad Request\r\n\r\n")
		return
	}

	log.Printf("Received Request: %s\n", string(buff[:n]))

	const statusCode = 200
	const statusText = "OK"
	const body = "Hello World"

	fmt.Fprintf(conn, "HTTP/1.0 %d %s\r\n", statusCode, statusText)
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprint(conn, body)
}
