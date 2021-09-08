package balanced_http_client

import "net/http"

func NewClient() *http.Client {
	return &http.Client{
		Transport: &Transport{
			Base: http.DefaultTransport,
		},
	}
}
