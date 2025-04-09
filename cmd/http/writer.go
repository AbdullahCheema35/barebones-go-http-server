package http

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(int)
}

type respWriter struct {
	conn                net.Conn
	response            *Response
	isWriteHeaderCalled bool
	request             *Request
}

func (rw *respWriter) Write(data []byte) (int, error) {
	if !rw.isWriteHeaderCalled {
		rw.WriteHeader(200)
	}
	return rw.writeToConn(data)
}

func (rw *respWriter) writeToConn(data []byte) (int, error) {
	log.Printf("Response Body: %s", data)
	return fmt.Fprintf(rw.conn, "%s", data)
}

func (rw *respWriter) WriteHeader(code int) {
	defer rw.setWriteHeaderCalled()

	// set the code
	rw.response.setResponseStatus(code)

	// get the status header and other headers
	escapedStatusHeader := rw.response.getEscapedStatusHeader()
	otherHeaders := rw.response.getOtherHeaders()
	rw.writeHeaderToConn(escapedStatusHeader, otherHeaders)
}

func (rw *respWriter) writeHeaderToConn(escapedStatusHeader string, otherHeaders []string) (int, error) {
	formattedHeaders := formatHTTPHeaders(otherHeaders)
	responseHeaders := fmt.Sprintf("%s%s\r\n", escapedStatusHeader, formattedHeaders)
	log.Printf("responseHeaders: %s", responseHeaders)
	n, err := fmt.Fprintf(rw.conn, "%s", responseHeaders)
	return n, err
}

func (rw *respWriter) Header() Header {
	return rw.response.Headers
}

func (rw *respWriter) setWriteHeaderCalled() {
	rw.isWriteHeaderCalled = true
}

func newResponseWriter(conn net.Conn, request *Request) *respWriter {
	return &respWriter{
		conn:     conn,
		request:  request,
		response: newResponse(request),
	}
}

// formats the headers for the HTTP response.
func formatHTTPHeaders(headers []string) string {
	if len(headers) == 0 {
		return ""
	}
	return fmt.Sprintf("%s\r\n", strings.Join(headers, "\r\n"))
}
