package fstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFstring(t *testing.T) {
	got := FString("{ a } {b}", M{"a": "a"})
	assert.Equal(t, "a {b}", got)
}
