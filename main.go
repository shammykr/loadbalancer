package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"loadbalancer/balancer"
)

func main() {
	servers := []string{
		"http://localhost:9001",
		"http://localhost:9002",
		"http://localhost:9003",
	}

	rr := balancer.NewRoundRobin()
	hc := balancer.NewHealthChecker(servers)

	// Start health checking every 3 seconds
	hc.Start(3 * time.Second)

	handler := func(w http.ResponseWriter, r *http.Request) {
		healthyServers := hc.HealthyServers()

		if len(healthyServers) == 0 {
			http.Error(w, "No backend servers available", http.StatusServiceUnavailable)
			return
		}

		// Pick next healthy server
		target := rr.NextServer(healthyServers)
		remoteURL, _ := url.Parse(target)

		proxy := httputil.NewSingleHostReverseProxy(remoteURL)

		proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {
			http.Error(w, "Backend unreachable", http.StatusServiceUnavailable)
		}

		proxy.ServeHTTP(w, r)
	}

	http.HandleFunc("/", handler)

	log.Println("Load Balancer running on :8080")
	http.ListenAndServe(":8080", nil)
}
