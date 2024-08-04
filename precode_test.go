package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.NotEmpty(t, responseRecorder.Body)
	assert.Len(t, list, totalCount)

}
func TestMainHandlerValidRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	require.Equal(t, status, http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body)
}
func TestMainHandlerWhenCityNotMoscow(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=mosc", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	require.Equal(t, status, http.StatusBadRequest)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())

}
