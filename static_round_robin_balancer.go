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

func (r *StaticRoundRobinBalancer) SetBackends(replace []*url.URL) {
	r.Lock()
	defer r.Unlock()
	r.index = 0
	r.backends = replace
}

func (r *StaticRoundRobinBalancer) Select(*http.Request) (*url.URL, error) {
	r.Lock()
	defer r.Unlock()

	for i := 0; len(r.backends) > i; i++ {
		candidate := r.backends[r.index]
		r.index = (r.index + 1) % len(r.backends)
		return candidate, nil
	}
	return nil, ErrNoHealthyUpstream
}
