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

	req := Request{
		Req: string(buff[:n]),
	}

	log.Printf("Received Request: %s\n", req.Req)

	resp := Response{
		Conn: conn,
	}

	handler.ServeHTTP(&resp, &req)

}
