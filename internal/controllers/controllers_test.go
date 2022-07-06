package controllers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGreeting(t *testing.T) {
	router := InitializeRoutes()
	request, error := http.NewRequest("GET", "/api/hello/Name", nil)
	assert.Nil(t, error)

	writer := httptest.NewRecorder()
	request.Header.Set(ContentType, ApplicationJSON)
	router.ServeHTTP(writer, request)

	assert.Equal(t, 200, writer.Code)
	assert.Equal(t, "\"Hello Name!\"\n", string(writer.Body.Bytes()))
}

func TestGetGreetingShortName(t *testing.T) {
	router := InitializeRoutes()
	request, error := http.NewRequest("GET", "/api/hello/a", nil)
	assert.Nil(t, error)

	writer := httptest.NewRecorder()
	request.Header.Set(ContentType, ApplicationJSON)
	router.ServeHTTP(writer, request)

	assert.Equal(t, 400, writer.Code)
	assert.Equal(t, "{\"message\":\"Name parameter should be bigger than two symbols.\"}\n", string(writer.Body.Bytes()))
}
