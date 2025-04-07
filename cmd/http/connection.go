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

}
