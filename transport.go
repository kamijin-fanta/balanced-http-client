package balanced_http_client

import "net/http"

var _ http.RoundTripper = &Transport{}

type Transport struct {
	Base http.RoundTripper
}

func (t *Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	return t.Base.RoundTrip(request)
}
