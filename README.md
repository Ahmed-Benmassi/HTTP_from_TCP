# Build an HTTP Server from Scratch in Go

The web is built on HTTP, and there's no better way to understand how 
something works than to implement it yourself.

In this project, we explore the nitty-gritty details of the HTTP/1.1 
protocol by building a fully functional HTTP server from scratch using Golang. 
By the end of this journey, you will have a deep, practical understanding of 
how web transport works under the hood—no third-party web frameworks required.

## What You Will Learn: 
- The fundamentals of HTTP/1.1 and network transport.
- How to read, parse, and process byte streams.
- Parsing HTTP request lines, headers, and bodies.
- Constructing proper HTTP responses.
- The difference between TCP and UDP, and why 
HTTP relies on TCP.

## Getting Started

 > [!NOTE]
 > Go: Make sure you have Go installed. 
 > You can download it from [golang](golang.org)  

### Running the Server
 
 1. Clone the repository:

~~~ sh

  git clone https://github.com/yourusername/http-server-go.git
  cd http-server-go

~~~

 2. Run the application:

~~~ sh 

  go run main.go

~~~

3. Test your server using curl or your web browser:

~~~ sh 

  curl -v http://localhost:8080

~~~

