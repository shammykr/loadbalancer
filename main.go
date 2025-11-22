package main

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "loadbalancer/balancer"
)

func main() {
    servers := []string{
        "http://localhost:9001",
        "http://localhost:9002",
        "http://localhost:9003",
    }

    rr := balancer.NewRoundRobin(servers)

    handler := func(w http.ResponseWriter, r *http.Request) {
        target := rr.NextServer()
        remoteURL, _ := url.Parse(target)

        proxy := httputil.NewSingleHostReverseProxy(remoteURL)

        proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {
            http.Error(w, "Backend server unreachable", http.StatusServiceUnavailable)
        }

        proxy.ServeHTTP(w, r)
    }

    http.HandleFunc("/", handler)

    log.Println("Load balancer started at :8080")
    http.ListenAndServe(":8080", nil)
}
