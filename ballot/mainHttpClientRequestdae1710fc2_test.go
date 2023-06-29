// Test generated by RoostGPT for test test1 using AI Model gpt

package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpClientRequestdae1710fc2(t *testing.T) {
	// Test case 1: Check successful request
	t.Run("Successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(`OK`))
		}))
		defer server.Close()

		status, body, err := httpClientRequest("GET", server.URL, "", nil)
		if err != nil {
			t.Error("Expected no error, got ", err)
		}
		if status != http.StatusOK {
			t.Error("Expected status OK, got ", status)
		}
		if string(body) != "OK" {
			t.Error("Expected body 'OK', got ", string(body))
		}
	})

	// Test case 2: Check unsuccessful request
	t.Run("Unsuccessful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(`Bad request`))
		}))
		defer server.Close()

		status, _, err := httpClientRequest("GET", server.URL, "", nil)
		if err == nil {
			t.Error("Expected error, got none")
		}
		if status != http.StatusBadRequest {
			t.Error("Expected status Bad Request, got ", status)
		}
	})

	// Test case 3: Check invalid URL
	t.Run("Invalid URL", func(t *testing.T) {
		_, _, err := httpClientRequest("GET", ":", "", nil)
		if err == nil {
			t.Error("Expected error, got none")
		}
		if !errors.Is(err, http.ErrInvalidHost) {
			t.Error("Expected invalid host error, got ", err)
		}
	})

	// Test case 4: Check non-existent server
	t.Run("Non-existent server", func(t *testing.T) {
		_, _, err := httpClientRequest("GET", "http://localhost:12345", "", nil)
		if err == nil {
			t.Error("Expected error, got none")
		}
	})
}