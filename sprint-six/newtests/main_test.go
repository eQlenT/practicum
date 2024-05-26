package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertExample(t *testing.T) {
	expected := "42"
	result := SomeFunction()
	assert.NotEmpty(t, result)
	assert.Equal(t, result, expected)
}
