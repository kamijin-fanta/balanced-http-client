package balanced_http_client

import (
	"errors"
	"net/http"
	"net/url"
)

var (
	ErrNoHealthyUpstream = errors.New("no healty upstream")
)

type Balancer interface {
	Select(req *http.Request) (*url.URL, error)
}
