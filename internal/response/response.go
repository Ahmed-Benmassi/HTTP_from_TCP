package response

import (
	"fmt"
	"httpfromtcp/internal/headers"
	"io"

	
)


type Response struct{

}


type StatusCode int 
const(
	StatusOK                  StatusCode=200
	StatusBADReaquest         StatusCode=400
	StatusInternalServerError StatusCode=500
)



func GetDefaultHeaders(contentLen int) *headers.Headers{
	h:=headers.NewHeaders()
	h.Replace("Content-Length",fmt.Sprintf("%d",contentLen))
	h.Replace("Connection","close")
	h.Replace("Content-Type","text/plain")
	
	return h
}



type Writer struct{
	writer io.Writer

}
func (w *Writer) WriteStatusLine(statusCode StatusCode) error {
	statusline:=[]byte{}
	
	switch statusCode{
	case StatusOK :  statusline=[]byte("HTTP/1.1 200 OK\r\n")
	case StatusBADReaquest : statusline= []byte("HTTP/1.1 400 Bad Request\r\n")
	case StatusInternalServerError : statusline= []byte("HTTP/1.1 500 Internal Server Error\r\n")
	default:
		return fmt.Errorf("unrecognized error code")
	}
	_,err:=w.writer.Write(statusline)                                                                      // doing this cause write return err and int 
	return err
}
func NewWriter(writer io.Writer)  *Writer{
	return &Writer{writer:writer}
} 

func (w *Writer) WriteHeaders(h headers.Headers) error{
	var err error=nil                                          // (like a network connection or a file) in the proper HTTP format
	b:=[]byte{}
	h.ForEach(func (n,v string)  {             //for each name in value usinf a function to parsing he header again 
        b=fmt.Appendf(b,"%s: %s\r\n",n, v)     //fmt.appendf designed for efficient string building by appending formatted data directly to a byte slice  
    })                                         //fmt.Appendf is much faster than fmt.Sprintf in loops or when building large strings incrementally
	
	b=fmt.Append(b,"\r\n")                     //adding \r\n in the next line
	_,err=w.writer.Write(b)                         //This writes all the header bytes to w and returns any write error.
	return err

}
func (w *Writer) WriteBody(p []byte) (int, error){
	n,err:=w.writer.Write(p)
	return n,err
}
