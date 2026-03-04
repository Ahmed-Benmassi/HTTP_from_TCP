# Build an HTTP Server from Scratch in Go

The web is built on HTTP, and there's no better way to understand how 
something works than to implement it yourself.

A NOT simple (hhh) Go project that demonstrates how to handle HTTP requests over raw TCP connections.

Unlike a typical Go HTTP server (which uses `net/http`), this project shows how TCP sockets can be used directly to listen for and respond to HTTP traffic — illustrating the fundamentals of networking and protocol layering.

This repository complements the **Chirpy JSON API** by exploring how to accept and parse HTTP over low-level TCP.

---
## 🚀 Goals

### This project is intended to:

- Demonstrate how HTTP can be implemented over TCP sockets
- Teach how the application layer (HTTP) interacts with the transport layer (TCP)
- Provide a simple foundation for building custom servers
- Support integration with higher-level REST APIs (like the Chirpy backend)

---

## 📁 Folder Structure

- `cmd/httpserver`: Simple HTTP server using raw TCP
- `cmd/tcplistener`: TCP listener that can accept and manage socket connections
- `internal`: Shared utilities and network helpers
- `go.mod`, `go.sum`: Go module dependencies

---

## 🛠 Tech Stack

- Go (Golang)
- TCP sockets (standard library)
- Plain HTTP parsing

---


## 🚧 How It Works

- A TCP listener is started on a port (e.g., localhost:8081)

- Incoming TCP connections are accepted

- Raw HTTP requests are read and parsed manually

- Responses are written back over the socket

- Demonstrates request/response flow without middleware or framework

-------
## Getting Started

 > [!NOTE]
 > Go: Make sure you have Go installed. 
 > You can download it from [golang](golang.org)  

### Running the Server

1. Clone the repository:

```bash
git clone https://github.com/Ahmed-Benmassi/HTTP_from_TCP.git
cd HTTP_from_TCP 
```

 2. Run the application:

~~~ sh 

  go run main.go

~~~

3. Test your server using curl or your web browser:

~~~ sh 

  curl -v http://localhost:8080
~~~


-----------------
## 💡 Future Ideas

- Integrate with Chirpy_Project through TCP forwarding

- Add TLS support (HTTPS over TLS handshake) (Working ON It)

- Provide example clients
