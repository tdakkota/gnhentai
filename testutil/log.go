package testutil

import (
	"net/http"
	"testing"
)

type LogTransport struct {
	t    *testing.T
	prev http.RoundTripper
}

func NewLogTransport(t *testing.T, prev http.RoundTripper) LogTransport {
	if prev == nil {
		prev = http.DefaultTransport
	}
	return LogTransport{t: t, prev: prev}
}

func (l LogTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	l.t.Log(req.URL.String())
	return l.prev.RoundTrip(req)
}

func TestClient(t *testing.T) *http.Client {
	return &http.Client{
		Transport: LogTransport{t, http.DefaultTransport},
	}
}
