package greetings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloName(t *testing.T) {
	expected := "Hi, Bruce Wayne"
	msg := Hello("Bruce Wayne")
	assert.Equal(t, expected, msg)

	expected = "Hello anonymous!"
	msg = Hello("")
	assert.Equal(t, expected, msg)
}
