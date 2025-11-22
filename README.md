I tried to implement a basic HTTP Load Balancer in Go using a Round Robin algorithm to distribute incoming client requests across multiple backend servers.

✅ Features in this project

Forwards HTTP requests to backend servers
Round-robin request distribution
Reverse proxy using Go standard library
Concurrency-safe server selection
Multiple backend servers running independently

How to run?
Run each server in different terminal.
go run backend/server1.go
go run backend/server2.go
go run backend/server3.go

Run the load Balancer
go run main.go

In browser
http://localhost:8080

Keep on refreshing the browser and you'll see output from different servers.
