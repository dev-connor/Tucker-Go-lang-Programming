package ch28

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare1(t *testing.T) {
	assert2 := assert2.New(t)                              // 테스트 객체 생성
	assert2.Equal(81, square(9), "square(9) should be 81") // 테스트 함수 호출
}

func TestSquare2(t *testing.T) {
	assert2 := assert2.New(t)
	assert2.Equal(9, square(3), "square(3) should be 9")
}
