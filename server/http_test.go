package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: write more test cases
func TestFetchAllTodo(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/todos/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateTodo(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	data := url.Values{}
	data.Add("title", "test 1")
	data.Add("completed", "0")

	req, _ := http.NewRequest("POST", "/api/v1/todos/", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}