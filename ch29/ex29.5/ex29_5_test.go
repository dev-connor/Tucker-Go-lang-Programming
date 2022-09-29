package main

import (
	assert2 "github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	assert := assert2.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil) // 1. / 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code) // 3. Code 확인
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestBarHandler(t *testing.T) {
	assert := assert2.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello Bar", string(data))
}
