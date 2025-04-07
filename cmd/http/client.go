package http

import (
	"fmt"
	"log"
	"net"
)

func handleClient(conn net.Conn, handler Handler) {
	defer conn.Close()

	var buff = make([]byte, 1024)

	n, err := conn.Read(buff)
	if err != nil {
		fmt.Fprintf(conn, "HTTP/1.0 400 Bad Request\r\n\r\n")
		return
	}

	log.Printf("Received Request: %s\n", string(buff[:n]))

	req := parseRequest(string(buff[:n]))

	resp := Response{
		Conn: conn,
	}

	handler.ServeHTTP(&resp, req)
}
