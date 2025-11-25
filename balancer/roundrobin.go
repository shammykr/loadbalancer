package balancer

import "sync"

type RoundRobin struct {
	index int
	mu    sync.Mutex
}

func NewRoundRobin() *RoundRobin {
	return &RoundRobin{}
}

func (r *RoundRobin) NextServer(servers []string) string {
	r.mu.Lock()
	defer r.mu.Unlock()

	// ALWAYS wrap correctly based on current healthy server count
	if len(servers) == 0 {
		return ""
	}

	if r.index >= len(servers) {
		r.index = 0
	}

	server := servers[r.index]
	r.index++

	return server
}
