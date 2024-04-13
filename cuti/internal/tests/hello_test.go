package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Log("Hello")
	assert.True(t, true)
}
