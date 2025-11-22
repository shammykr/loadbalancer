package balancer

import "sync"

type RoundRobin struct {
    servers []string
    index   int
    mu      sync.Mutex
}

func NewRoundRobin(servers []string) *RoundRobin {
    return &RoundRobin{
        servers: servers,
        index:   0,
    }
}

func (r *RoundRobin) NextServer() string {
    r.mu.Lock()
    defer r.mu.Unlock()
    server := r.servers[r.index]
    r.index = (r.index + 1) % len(r.servers)
    return server
}
