package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	result, err := SayHello("taufik", false)

	if err != nil {
		t.Fatal("error occured")
	}

	if result != "hello taufik" {
		t.Fatal("result must be 'hello taufik'")
	}
}

func TestSayHello2(t *testing.T) {
	result, err := SayHello("taufik", false)
	assert.NoError(t, err, "error occured")
	assert.Equal(t, "hello taufik", result, "result must be 'hello taufik'")
}

func TestSayHello3(t *testing.T) {
	t.Run("no error check", func(t *testing.T) {
		result, err := SayHello("taufik", false)
		assert.NoError(t, err, "error occured")
		assert.Equal(t, "hello taufik", result, "result must be 'hello taufik'")
	})

	t.Run("when error check", func(t *testing.T) {
		result, err := SayHello("taufik", true)
		assert.Error(t, err, "no error occured")
		assert.Empty(t, result, "still got the result")
	})

}
