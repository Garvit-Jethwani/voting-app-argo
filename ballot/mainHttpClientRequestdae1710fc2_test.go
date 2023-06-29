// Test generated by RoostGPT for test test1 using AI Model gpt

package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpClientRequestdae1710fc2(t *testing.T) {
	// Test case 1: Successful request
	t.Run("successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Write([]byte(`OK`))
		}))
		defer server.Close()

		statusCode, body, err := httpClientRequest("GET", server.URL, "/", nil)
		if err != nil {
			t.Error(err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("expected status code to be %d, but got %d", http.StatusOK, statusCode)
		}
		if string(body) != "OK" {
			t.Errorf("expected body to be %s, but got %s", "OK", string(body))
		}
	})

	// Test case 2: Unsuccessful request due to invalid URL
	t.Run("unsuccessful request due to invalid URL", func(t *testing.T) {
		_, _, err := httpClientRequest("GET", "invalid url", "/", nil)
		if err == nil {
			t.Error("expected error due to invalid url, but got nil")
		}
	})

	// Test case 3: Unsuccessful request due to non-2xx/3xx status code
	t.Run("unsuccessful request due to non-2xx/3xx status code", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			http.Error(rw, "Not Found", http.StatusNotFound)
		}))
		defer server.Close()

		statusCode, _, err := httpClientRequest("GET", server.URL, "/", nil)
		if err == nil {
			t.Error("expected error due to non-2xx/3xx status code, but got nil")
		}
		if statusCode != http.StatusNotFound {
			t.Errorf("expected status code to be %d, but got %d", http.StatusNotFound, statusCode)
		}
	})

	// Test case 4: Successful request with POST method and JSON body
	t.Run("successful request with POST method and JSON body", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			body, _ := ioutil.ReadAll(req.Body)
			if !strings.Contains(string(body), "test") {
				t.Error("request body does not contain expected value")
			}
			rw.Write([]byte(`OK`))
		}))
		defer server.Close()

		jsonBody := bytes.NewBuffer([]byte(`{"key":"test"}`))
		statusCode, body, err := httpClientRequest("POST", server.URL, "/", jsonBody)
		if err != nil {
			t.Error(err)
		}
		if statusCode != http.StatusOK {
			t.Errorf("expected status code to be %d, but got %d", http.StatusOK, statusCode)
		}
		if string(body) != "OK" {
			t.Errorf("expected body to be %s, but got %s", "OK", string(body))
		}
	})
}
