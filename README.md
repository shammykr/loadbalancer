Load Balancer in Go — (Round Robin)

I tried to implement a basic HTTP Load Balancer in Go using a Round Robin algorithm to distribute incoming client requests across multiple backend servers.

✅ Features in Level 1

Forwards HTTP requests to backend servers

Round-robin request distribution

Reverse proxy using Go standard library

Concurrency-safe server selection

Multiple backend servers running independently

📂 Project Structure

loadbalancer/
 ├── main.go                  # Load balancer entry point
 ├── go.mod                   # Go module definition
 ├── balancer/
 │     ├── balancer.go        # LoadBalancer interface
 │     └── roundrobin.go      # Round Robin implementation
 └── backend/
       ├── server1.go         # Backend server 1 (port 9001)
       ├── server2.go         # Backend server 2 (port 9002)
       └── server3.go         # Backend server 3 (port 9003)


🚀 How to Run the Project

✅ 1. Run each backend server in separate terminal

go run backend/server1.go
go run backend/server2.go
go run backend/server3.go


✅ 2. Run the Load Balancer

go run main.go


✅ 3. Send requests

In browser or curl:
http://localhost:8080

✅ Expected rotating output:

Response from Server 1
Response from Server 2
Response from Server 3
Response from Server 1
...


🧠 How It Works (Short Explanation)

Client sends request to Load Balancer (localhost:8080)

Load balancer selects backend using Round Robin 3.  Reverse proxy forwards request to selected backend

Backend responds

Load balancer returns response to client

🛠 Technologies Used

Go (Golang)

net/http

httputil.ReverseProxy

Mutex for concurrency safety




“Let’s start Level 2.”
