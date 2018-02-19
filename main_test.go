package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	server *httptest.Server
	usrUrl string
)

func init() {
	server = httptest.NewServer(Routes())
}

func doReq(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func checkResponseCode(expected, actual int) error {
	if expected != actual {
		return fmt.Errorf("Expected Status Code %d. Got %d\n", expected, actual)
	}
	return nil
}
func TestNotFound(t *testing.T) {
	usrUrl = fmt.Sprintf("%s/api/v1/you/not/found/me", server.URL)
	res, err := doReq(usrUrl)
	if err != nil {
		t.Error(err)
	}
	err = checkResponseCode(http.StatusNotFound, res.StatusCode)
	if err != nil {
		t.Error(err)
	}
}
