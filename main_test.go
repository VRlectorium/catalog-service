package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var (
	server *httptest.Server
	usrUrl string
)

func init() {
	server = httptest.NewServer(Routes())
}

type TestRequest struct {
	url string
}

func (r *TestRequest) doGet() (*http.Response, error) {
	req, err := http.NewRequest("GET", r.url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *TestRequest) doPost(body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", r.url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type ExpectedResponse struct {
	ContentType string
	StatusCode  int
}

func NewExpectedResponse(statusCode int) *ExpectedResponse {
	return &ExpectedResponse{StatusCode: statusCode, ContentType: "application/json"}
}

func (r *ExpectedResponse) checkResponseCode(actual int) error {
	if r.StatusCode != actual {
		return fmt.Errorf("Expected Status Code %d. Got %d\n", r.StatusCode, actual)
	}
	return nil
}

func (r *ExpectedResponse) checkResponse(res *http.Response) error {
	err := r.checkResponseCode(res.StatusCode)
	if err != nil {
		return err
	}
	actualContentType := res.Header.Get("Content-Type")
	if r.ContentType != actualContentType {
		return fmt.Errorf("Expected Content-Type %s. Actual %s\n", r.ContentType, actualContentType)
	}
	return nil
}
func TestNotFound(t *testing.T) {
	expectedResponse := NewExpectedResponse(http.StatusNotFound)
	testReq := TestRequest{url: fmt.Sprintf("%s/api/v1/you/not/found/me", server.URL)}
	res, err := testReq.doGet()
	if err != nil {
		t.Error(err)
	}
	err = expectedResponse.checkResponse(res)
	if err != nil {
		t.Error(err)
	}
}

func TestGetCourses(t *testing.T) {
	expectedResponse := NewExpectedResponse(http.StatusOK)
	tesReq := TestRequest{url: fmt.Sprintf("%s/api/v1/courses", server.URL)}
	res, err := tesReq.doGet()
	if err != nil {
		t.Error(err)
	}
	err = expectedResponse.checkResponse(res)
	if err != nil {
		t.Error(err)
	}
}

func TestPostCourse(t *testing.T) {
	expectedResponse := NewExpectedResponse(http.StatusCreated)
	testReq := TestRequest{url: fmt.Sprintf("%s/api/v1/courses", server.URL)}
	res, err := testReq.doPost(strings.NewReader("name=GOLANG"))
	if err != nil {
		t.Error(err)
	}
	err = expectedResponse.checkResponse(res)
	if err != nil {
		t.Error(err)
	}
}

func TestGetSubject(t *testing.T) {
	expectedResponse := NewExpectedResponse(http.StatusOK)
	tesReq := TestRequest{url: fmt.Sprintf("%s/api/v1/subjects", server.URL)}
	res, err := tesReq.doGet()
	if err != nil {
		t.Error(err)
	}
	err = expectedResponse.checkResponse(res)
	if err != nil {
		t.Error(err)
	}
}

func TestPostSubject(t *testing.T) {
	expectedResponse := NewExpectedResponse(http.StatusCreated)
	testReq := TestRequest{url: fmt.Sprintf("%s/api/v1/subjects", server.URL)}
	form := url.Values{}
	form.Add("courseid", "1")
	form.Add("name", "GOLANG")
	res, err := testReq.doPost(strings.NewReader(form.Encode()))
	if err != nil {
		t.Error(err)
	}
	err = expectedResponse.checkResponse(res)
	if err != nil {
		t.Error(err)
	}
}
