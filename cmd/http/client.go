package http

import (
	"log"
	"net"
)

func handleClient(conn net.Conn, handler Handler) {
	defer conn.Close()

	var buff = make([]byte, 1024)

	n, err := conn.Read(buff)
	if err != nil {
		log.Println("Error reading from client connection:", err)
		return
	}

	log.Printf("Received Request: %s\n", string(buff[:n]))

	req := parseRequest(string(buff[:n]))
	rw := newResponseWriter(conn, req)
	// if req.Proto == "" {
	// 	resp.response = &Response{
	// 		Proto: "HTTP/1.0",
	// 	}
	// } else {
	// 	resp.response = &Response{
	// 		Proto: req.Proto,
	// 	}
	// }
	// resp.response.Headers = make(Header)
	// resp.response.Headers["Content-Type"] = "text/html"
	// resp.response.Headers["Server"] = "Go HTTP Server"
	// resp.response.Headers["Date"] = "Wed, 21 Oct 2015 07:28:00 GMT"
	// resp.response.Headers["Connection"] = "close"
	// resp.response.Headers["Content-Length"] = "0"
	// resp.response.Headers["Accept"] = req.Header["Accept"]
	// resp.response.Headers["Accept-Encoding"] = req.Header["Accept-Encoding"]
	// resp.response.Headers["Accept-Language"] = req.Header["Accept-Language"]
	// resp.response.Headers["Accept-Charset"] = req.Header["Accept-Charset"]
	// resp.response.Headers["User-Agent"] = req.Header["User-Agent"]

	handler.ServeHTTP(rw, req)
}
