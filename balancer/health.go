package balancer

import (
	"net/http"
	"sync"
	"time"
)

type ServerStatus struct {
	URL      string
	Healthy  bool
	Failures int
}

type HealthChecker struct {
	servers []*ServerStatus
	mu      sync.Mutex
}

func NewHealthChecker(urls []string) *HealthChecker {
	servers := make([]*ServerStatus, len(urls))
	for i, url := range urls {
		servers[i] = &ServerStatus{
			URL:     url,
			Healthy: true,
		}
	}
	return &HealthChecker{servers: servers}
}

func (hc *HealthChecker) Start(interval time.Duration) {
	go func() {
		for {
			hc.checkAll()
			time.Sleep(interval)
		}
	}()
}

func (hc *HealthChecker) checkAll() {
	hc.mu.Lock()
	defer hc.mu.Unlock()

	for _, s := range hc.servers {
		resp, err := http.Get(s.URL + "/health")
		if err != nil || resp.StatusCode != 200 {
			s.Healthy = false
		} else {
			s.Healthy = true
			s.Failures = 0
		}
	}
}

func (hc *HealthChecker) HealthyServers() []string {
	hc.mu.Lock()
	defer hc.mu.Unlock()

	healthy := []string{}
	for _, s := range hc.servers {
		if s.Healthy {
			healthy = append(healthy, s.URL)
		}
	}
	return healthy
}
