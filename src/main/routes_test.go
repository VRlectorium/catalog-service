package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestRoutes(t *testing.T) {
	port := 1234
	defer startServerOnPort(t, port, Routes().Handler).Close()
	getVersionEndpoint(t)
	getNotFound(t)
}

func getVersionEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:1234/version", nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d. Got %d\n", http.StatusOK, res.StatusCode)
	}
}

func getNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:1234/you/will/not/found/me", nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expected %d. Got %d\n", http.StatusNotFound, res.StatusCode)
	}
}

func startServerOnPort(t *testing.T, port int, h fasthttp.RequestHandler) io.Closer {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		t.Fatalf("cannot start tcp server on port %d: %s", port, err)
	}
	go fasthttp.Serve(ln, h)
	return ln
}
