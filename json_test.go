package jgstr

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestJsonEncode(t *testing.T) {
	assert.Equal(t, `{"a":1}`, JsonEncode(map[string]int{"a": 1}))
}
