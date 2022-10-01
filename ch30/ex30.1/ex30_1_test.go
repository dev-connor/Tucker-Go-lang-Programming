package main

import (
	"encoding/json"
	assert2 "github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJsonHandler1(t *testing.T) {
	assert := assert2.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil)

	// 1. /students 경로 테스트
	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(2, len(list))
	assert.Equal("aaa", list[0].Name)
	assert.Equal("bbb", list[1].Name)

}

func TestJsonHandler2(t *testing.T) {
	var student Student

	assert := assert2.New(t)
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil)
	mux.ServeHTTP(res, req)

	assert2.Equal(t, http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert2.Nil(t, err)
	assert2.Equal(t, "aaa", student.Name)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/2", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err = json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("bbb", student.Name)
}

func TestJsonHandler3(t *testing.T) {
	var student Student

	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students",
		strings.NewReader(`{"Id":0,"Name":"ccc","Age":15,"Score":78}`))

	// 1. 새로운 학생 데이터
	mux.ServeHTTP(res, req)

	assert2.Equal(t, http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/3", nil)

	// 3. 추가된 학생 데이터
	mux.ServeHTTP(res, req)

	err := json.NewDecoder(res.Body).Decode(&student)
	assert2.Nil(t, err)
	assert2.Equal(t, "ccc", student.Name)
}
