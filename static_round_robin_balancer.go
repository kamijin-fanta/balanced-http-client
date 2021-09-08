package balanced_http_client

import (
	"net/http"
	"net/url"
	"sync"
)

var _ Balancer = &StaticRoundRobinBalancer{}

type StaticRoundRobinBalancer struct {
	sync.Mutex
	backends []*url.URL
	index    int
}

// todo discovery

func (r *StaticRoundRobinBalancer) Select(*http.Request) (*url.URL, error) {
	r.Lock()
	defer r.Unlock()

	for i := 0; i < len(r.backends); i++ {
		candidate := r.backends[r.index]
		r.index = (r.index + 1) % len(r.backends)
		return candidate, nil
	}
	return nil, ErrNoHealthyUpstream
}
