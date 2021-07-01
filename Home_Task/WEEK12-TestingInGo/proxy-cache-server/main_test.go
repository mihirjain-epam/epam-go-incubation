package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestProxyOrCacheRequest(t *testing.T) {
	// init start
	go http.ListenAndServe(":8080", nil)
	go http.ListenAndServeTLS(":9443", "RootCA.crt", "RootCA.key", nil)
	time.Sleep(5 * time.Second) // let the end point start
	// init end

	// parallel tests for different ports
	t.Run("Ports", func(t *testing.T) {
		t.Run("Port=8080", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/github.com/stretchr/testify/assert", nil)
			res := httptest.NewRecorder()

			proxyOrCacheRequest(res, req)

			if res.Result().StatusCode != http.StatusOK {
				t.Errorf("Expected status code:%d, received status code:%d", http.StatusOK, res.Result().StatusCode)
			}
		})
		t.Run("Port=9443", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://localhost:9443/github.com/stretchr/testify/assert", nil)
			res := httptest.NewRecorder()

			proxyOrCacheRequest(res, req)

			if res.Result().StatusCode != http.StatusOK {
				t.Errorf("Expected status code:%d, received status code:%d", http.StatusOK, res.Result().StatusCode)
			}
		})
		t.Run("Port=8081", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://localhost:8081/github.com/stretchr/testify/assert", nil)
			res := httptest.NewRecorder()

			proxyOrCacheRequest(res, req)

			if res.Result().StatusCode != http.StatusBadRequest {
				t.Errorf("Expected status code:%d, received status code:%d", http.StatusBadRequest, res.Result().StatusCode)
			}
		})
	})
	// parallel tests for different methods
	t.Run("Methods", func(t *testing.T) {
		t.Run("UnsafeMethods", func(t *testing.T) {
			t.Run("UnsafeMethod=POST", func(t *testing.T) {
				req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/github.com/stretchr/testify/assert", nil)
				res := httptest.NewRecorder()

				proxyOrCacheRequest(res, req)

				if res.Result().StatusCode != http.StatusMethodNotAllowed {
					t.Errorf("Expected status code:%d, received status code:%d", http.StatusMethodNotAllowed, res.Result().StatusCode)
				}
			})
			t.Run("UnsafeMethod=PUT", func(t *testing.T) {
				req := httptest.NewRequest(http.MethodPut, "http://localhost:8080/github.com/stretchr/testify/assert", nil)
				res := httptest.NewRecorder()

				proxyOrCacheRequest(res, req)

				if res.Result().StatusCode != http.StatusMethodNotAllowed {
					t.Errorf("Expected status code:%d, received status code:%d", http.StatusMethodNotAllowed, res.Result().StatusCode)
				}
			})
			t.Run("UnsafeMethod=DELETE", func(t *testing.T) {
				req := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/github.com/stretchr/testify/assert", nil)
				res := httptest.NewRecorder()

				proxyOrCacheRequest(res, req)

				if res.Result().StatusCode != http.StatusMethodNotAllowed {
					t.Errorf("Expected status code:%d, received status code:%d", http.StatusMethodNotAllowed, res.Result().StatusCode)
				}
			})
			// a lot of other unsafe methods exist, similar tests can be written for all
		})
	})
	// tests for request url
	t.Run("BadUrl", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/example.com/stretchr/testify/assert", nil)
		res := httptest.NewRecorder()

		proxyOrCacheRequest(res, req)
		if res.Result().StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code:%d, received status code:%d", http.StatusBadRequest, res.Result().StatusCode)
		}
	})
}
