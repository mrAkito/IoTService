package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	router.NewRouter()

	req := httptest.NewRequest("GET", "http://localhost:8000/users", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, helloMessage, rec.Body.String())
}
