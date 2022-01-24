package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// Define api host and port.
var host string = "http://localhost"

// Request URI.
func doRequest(method string, uri string, data string, t *testing.T) (response string, code int) {
	payload := strings.NewReader(data)
	req, _ := http.NewRequest(method, host+uri, payload)
	if method == http.MethodPost {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("GET %s/%s failed: %v", host, uri, err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body), res.StatusCode
}

//GET method.
func TestGetMethod(t *testing.T) {
	_, code := doRequest(http.MethodGet, "/top", "", t)
	if code != http.StatusMethodNotAllowed {
		t.Fatalf("GET /top should return status 405, and returned %d", code)
	}
}

//POST without vars.
func TestPostWithoutVars(t *testing.T) {
	_, code := doRequest(http.MethodPost, "/top", "", t)
	if code != http.StatusBadRequest {
		t.Fatalf("GET /top should return status 400, and returned %d", code)
	}
}

// POST without empty n.
func TestPostWithEmptyN(t *testing.T) {
	_, code := doRequest(http.MethodPost, "/top", "&n=", t)
	if code != http.StatusBadRequest {
		t.Fatalf("GET /top should return status 400, and returned %d", code)
	}
}

// POST with NaN(n).
func TestPostWithNaNN(t *testing.T) {
	_, code := doRequest(http.MethodPost, "/top", "text=text&n=nan", t)
	if code != http.StatusBadRequest {
		t.Fatalf("GET /top should return status 400, and returned %d", code)
	}
}

// POST with n <= 0.
func TestPostWithZeroOrNegativeN(t *testing.T) {
	_, code := doRequest(http.MethodPost, "/top", "text=text&n=0", t)
	if code != http.StatusBadRequest {
		t.Fatalf("GET /top should return status 400, and returned %d", code)
	}
}

// POST with high n.
func TestPostWithHighN(t *testing.T) {
	_, code := doRequest(http.MethodPost, "/top", "n=1000000", t)
	if code != http.StatusBadRequest {
		t.Fatalf("GET /top should return status 400, and returned %d", code)
	}
}

// POST with valid n and text (just check response code).
func TestPostWithValidN(t *testing.T) {
	_, code := doRequest(http.MethodPost, "/top", "text=text&n=1", t)
	if code != http.StatusOK {
		t.Fatalf("GET /top should return status 200, and returned %d", code)
	}
}
