package server

import (
	"fmt"
	"httpfromtcp/internal/headers"
	"httpfromtcp/internal/request"
	"httpfromtcp/internal/response"
	"io"
	"net"
)



type  HandlerError struct{
	StatusCode  response.StatusCode
	Message     string
}


type Handler func(w *response.Writer, req *request.Request) 


type Server struct {
	closed bool
	handler Handler
}

func (s *Server) Close() {
	panic("unimplemented")
}

func runConnection(s *Server, conn io.ReadWriteCloser) {
	defer conn.Close()

	responseWriter:=response.NewWriter(conn)
	
	r,err:=request.RequestFromReader(conn)
	if err!=nil {
		responseWriter.WriteStatusLine(response.StatusBADReaquest)
	    responseWriter.WriteHeaders(*headers.NewHeaders())
        return
	}
	
	s.handler(responseWriter,r)
    
	
}

func runServer(s *Server, listener net.Listener) error {
	for {
		conn, err := listener.Accept()
		if s.closed {
			return nil
		}
		if err != nil {
			return err
		}

		go runConnection(s, conn)
	}

}

func Serve(port uint16,handler Handler) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	server := &Server{
		closed: false,
	    handler: handler,
	}
	go runServer(server, listener)

	return server, nil
}

func (s *Server) close() error {
	s.closed = true
	return nil
}

func (s *Server) listeb() error {
	for {

	}

}

func (s *Server) handle(conn net.Conn) {

}
