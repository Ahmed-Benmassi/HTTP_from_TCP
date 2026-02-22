package request

import (
	"bytes"
	"fmt"
	"httpfromtcp/internal/headers"
	"io"
	
)

type parserState string
const (
	StateInit  parserState = "init"
	StateHeaders parserState="header"
	StateDone  parserState = "done"
	StateErorr parserState = "error"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

type Request struct {
	RequestLine RequestLine
	Headers     headers.Headers
	state       parserState
}

func (r *Request) done() bool {
	return r.state == StateDone 
}

var ERROR_BAD_START_LINE = fmt.Errorf("Malformed request-line")
var ERROR_UNSUPPORTED_HTTP_VERSION = fmt.Errorf("unsupported http version")
var ERROR_REQUEST_ERROR_STATE = fmt.Errorf("request in error state")
var SEPARATOR = []byte("\r\n")

func newRequest() *Request {
	return &Request{
		state: StateInit,
		Headers: *headers.NewHeaders(),
	}
}

func (r *RequestLine) validHTTP() bool {
	return r.HttpVersion == "HTTP/1.1"
}

func parseRequestline(b []byte) (*RequestLine, int, error) {
	idx := bytes.Index(b, SEPARATOR)
	if idx == -1 {
		return nil, 0, nil

	}
	startLine := b[:idx]
	read := idx + len(SEPARATOR)

	parts := bytes.Split(startLine, []byte(" "))

	if len(parts) != 3 {
		return nil, 0, ERROR_BAD_START_LINE
	}

	httpParts := bytes.Split(parts[2], []byte("/"))
	if len(httpParts) != 2 || string(httpParts[0]) != "HTTP" || string(httpParts[1]) != "1.1" {
		return nil, 0, ERROR_BAD_START_LINE
	}

	rl := &RequestLine{
		Method:        string(parts[0]),
		RequestTarget: string(parts[1]),
		HttpVersion:   string(httpParts[1]),
	}
	return rl, read, nil

}

func (r *Request) parse(data[]byte) (int, error) {
	read := 0
outer:
	for {
		currentData:=data[read:]
		
		switch r.state {
		case StateErorr:
			return 0, ERROR_REQUEST_ERROR_STATE

		case StateInit:
			
			rl, n, err := parseRequestline(currentData)
			if err != nil {
				r.state = StateErorr
				return 0, err
			}
			if n == 0 {
				break outer
			}

			r.RequestLine = *rl
			read += n
			r.state = StateHeaders
			break

        case StateHeaders:
			
			n,done,err:=r.Headers.Parse(currentData)
			if err!= nil{
				return 0,err
			}

			if n==0 {
				break outer
			}

			read+=n
			
			if done{
				r.state=StateDone
			}

		case StateDone:
			break outer
		default:
			panic("somehow we haveprogrammed poorly")
		}
	}

	return read, nil

}

func (r *Request) error() bool {
	return r.state == StateDone || r.state == StateErorr
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	request := newRequest()
	// note : buffer could get overrun .... a header that exceeds 1k would do that or the bpdy
	buf := make([]byte, 1024)
	bufLen := 0
	for !request.done() {
		n, err := reader.Read(buf[bufLen:])
		if err != nil {
			return nil, err
		}

		bufLen += n

		readN, err := request.parse(buf[:bufLen])
		if err != nil {
			return nil, err
		}

		copy(buf, buf[readN:bufLen])
		bufLen -= readN

	}

	return request, nil
}
