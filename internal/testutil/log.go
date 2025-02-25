package testutil

import (
	"net/http"
)

type testingT interface {
	Logf(format string, args ...any)
}

// LogTransport logs all requests to testing logger.
type LogTransport struct {
	t    testingT
	prev http.RoundTripper
}

var _ http.RoundTripper = (*LogTransport)(nil)

// NewLogTransport creates a new [LogTransport].
func NewLogTransport(t testingT, prev http.RoundTripper) *LogTransport {
	if prev == nil {
		prev = http.DefaultTransport
	}
	return &LogTransport{t: t, prev: prev}
}

// RoundTrip implements [http.RoundTripper].
func (l *LogTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	l.t.Logf("Request %q", req.URL.String())
	return l.prev.RoundTrip(req)
}

// TestClient creates a new http client with logging transport.
func TestClient(t testingT) *http.Client {
	return &http.Client{
		Transport: NewLogTransport(t, http.DefaultTransport),
	}
}
